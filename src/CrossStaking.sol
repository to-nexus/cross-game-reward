// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {CrossStakingPool} from "./CrossStakingPool.sol";
import {WCROSS} from "./WCROSS.sol";

import {ICrossStaking} from "./interfaces/ICrossStaking.sol";
import {AccessControlDefaultAdminRulesUpgradeable} from
    "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @title CrossStaking
 * @notice 여러 스테이킹 풀을 관리하는 팩토리 컨트랙트
 * @dev UUPS 업그레이더블, 풀 생성/관리, WCROSS 화이트리스트 관리
 */
contract CrossStaking is Initializable, AccessControlDefaultAdminRulesUpgradeable, UUPSUpgradeable, ICrossStaking {
    using EnumerableSet for EnumerableSet.AddressSet;
    using EnumerableSet for EnumerableSet.UintSet;

    // ==================== Roles ====================

    bytes32 public constant POOL_MANAGER_ROLE = keccak256("POOL_MANAGER_ROLE");

    // ==================== 에러 ====================

    error CSPoolNotFound();
    error CSCanNotZeroAddress();

    // ==================== 이벤트 ====================

    event PoolCreated(uint indexed poolId, address indexed poolAddress, address indexed stakingToken);
    event PoolImplementationSet(address indexed implementation);
    event PoolStatusChanged(uint indexed poolId, bool active);
    event RouterSet(address indexed router);

    // ==================== 상태 변수 ====================

    /// @notice WCROSS 토큰 주소
    address public wcross;

    /// @notice Router 주소
    address public router;

    /// @notice CrossStakingPool implementation (UUPS 프록시용)
    address public poolImplementation;

    /// @notice 다음 풀 ID
    uint public nextPoolId;

    /// @notice 풀 ID => 풀 정보
    mapping(uint => PoolInfo) public pools;

    /// @notice 풀 주소 => 풀 ID
    mapping(address => uint) public poolIds;

    /// @notice 스테이킹 토큰 => 풀 ID 목록
    mapping(address => EnumerableSet.UintSet) private _poolsByStakingToken;

    /// @notice 모든 풀 ID 목록
    EnumerableSet.UintSet private _allPoolIds;

    // ==================== 생성자 ====================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice 컨트랙트 초기화
     * @param _poolImplementation CrossStakingPool implementation 주소
     * @param _admin 관리자 주소
     * @param _initialDelay 관리자 변경 딜레이 (초)
     */
    function initialize(address _poolImplementation, address _admin, uint48 _initialDelay) external initializer {
        require(_poolImplementation != address(0), CSCanNotZeroAddress());
        require(_admin != address(0), CSCanNotZeroAddress());

        __AccessControlDefaultAdminRules_init(_initialDelay, _admin);
        __UUPSUpgradeable_init();

        poolImplementation = _poolImplementation;
        wcross = address(new WCROSS());
        nextPoolId = 1;

        // 기본 역할 부여
        _grantRole(POOL_MANAGER_ROLE, _admin);
    }

    // ==================== 풀 관리 함수 ====================

    /**
     * @notice 새 스테이킹 풀 생성
     * @param stakingToken 스테이킹 토큰 주소
     * @param initialDelay 풀 관리자 변경 딜레이
     * @return poolId 생성된 풀 ID
     * @return poolAddress 생성된 풀 주소
     * @dev 풀의 admin은 CrossStaking 자신으로 설정됨
     */
    function createPool(address stakingToken, uint48 initialDelay)
        external
        onlyRole(POOL_MANAGER_ROLE)
        returns (uint poolId, address poolAddress)
    {
        require(stakingToken != address(0), CSCanNotZeroAddress());

        poolId = nextPoolId++;

        // UUPS 프록시로 풀 배포 (admin = CrossStaking 자신)
        bytes memory initData = abi.encodeWithSelector(
            CrossStakingPool.initialize.selector,
            IERC20(stakingToken),
            address(this), // CrossStaking이 풀의 admin
            initialDelay
        );

        ERC1967Proxy proxy = new ERC1967Proxy(poolImplementation, initData);
        poolAddress = address(proxy);

        // 풀 정보 저장
        pools[poolId] = PoolInfo({
            poolId: poolId,
            poolAddress: poolAddress,
            stakingToken: stakingToken,
            createdAt: block.timestamp,
            active: true
        });

        poolIds[poolAddress] = poolId;
        _allPoolIds.add(poolId);
        _poolsByStakingToken[stakingToken].add(poolId);

        emit PoolCreated(poolId, poolAddress, stakingToken);
    }

    /**
     * @notice 풀에 보상 토큰 추가
     * @param poolId 풀 ID
     * @param rewardToken 보상 토큰 주소
     * @dev CrossStaking이 풀의 admin이므로 직접 호출 가능
     */
    function addRewardToken(uint poolId, address rewardToken) external onlyRole(POOL_MANAGER_ROLE) {
        require(pools[poolId].poolAddress != address(0), CSPoolNotFound());

        CrossStakingPool pool = CrossStakingPool(pools[poolId].poolAddress);
        pool.addRewardToken(rewardToken);
    }

    /**
     * @notice 풀 활성화/비활성화
     * @param poolId 풀 ID
     * @param active 활성화 상태
     */
    function setPoolActive(uint poolId, bool active) external onlyRole(POOL_MANAGER_ROLE) {
        require(pools[poolId].poolAddress != address(0), CSPoolNotFound());

        pools[poolId].active = active;

        // 실제 풀 컨트랙트 pause/unpause
        CrossStakingPool pool = CrossStakingPool(pools[poolId].poolAddress);
        if (active) pool.unpause();
        else pool.pause();

        emit PoolStatusChanged(poolId, active);
    }

    /**
     * @notice Pool Implementation 업데이트
     * @param newImplementation 새 implementation 주소
     * @dev 기존 풀은 영향 없음, 새로 생성되는 풀만 적용
     */
    function setPoolImplementation(address newImplementation) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(newImplementation != address(0), CSCanNotZeroAddress());
        poolImplementation = newImplementation;
        emit PoolImplementationSet(newImplementation);
    }

    function setRouter(address _router) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_router != address(0), CSCanNotZeroAddress());
        router = _router;
        emit RouterSet(_router);
    }

    // ==================== View 함수 ====================

    /**
     * @notice 풀 정보 조회
     * @param poolId 풀 ID
     * @return 풀 정보
     */
    function getPoolInfo(uint poolId) external view returns (PoolInfo memory) {
        require(pools[poolId].poolAddress != address(0), CSPoolNotFound());
        return pools[poolId];
    }

    /**
     * @notice 인덱스로 풀 ID 조회
     * @param index 인덱스
     * @return poolId 풀 ID
     */
    function poolAt(uint index) external view returns (uint) {
        return _allPoolIds.at(index);
    }

    /**
     * @notice 특정 스테이킹 토큰의 풀 개수 조회
     * @param stakingToken 스테이킹 토큰 주소
     * @return 풀 개수
     */
    function getPoolCountByStakingToken(address stakingToken) external view returns (uint) {
        return _poolsByStakingToken[stakingToken].length();
    }

    /**
     * @notice 특정 스테이킹 토큰의 풀 ID 목록 조회
     * @param stakingToken 스테이킹 토큰 주소
     * @return 풀 ID 배열
     */
    function getPoolIdsByStakingToken(address stakingToken) external view returns (uint[] memory) {
        return _poolsByStakingToken[stakingToken].values();
    }

    /**
     * @notice 특정 스테이킹 토큰의 인덱스로 풀 ID 조회
     * @param stakingToken 스테이킹 토큰 주소
     * @param index 인덱스
     * @return poolId 풀 ID
     */
    function poolByStakingTokenAt(address stakingToken, uint index) external view returns (uint) {
        return _poolsByStakingToken[stakingToken].at(index);
    }

    /**
     * @notice 전체 풀 개수 조회
     * @return 전체 풀 개수
     */
    function getTotalPoolCount() external view returns (uint) {
        return _allPoolIds.length();
    }

    /**
     * @notice 모든 풀 ID 조회
     * @return 풀 ID 배열
     */
    function getAllPoolIds() external view returns (uint[] memory) {
        return _allPoolIds.values();
    }

    /**
     * @notice 풀 ID로 풀 주소 조회
     * @param poolId 풀 ID
     * @return 풀 주소
     */
    function getPoolAddress(uint poolId) external view returns (address) {
        require(pools[poolId].poolAddress != address(0), CSPoolNotFound());
        return pools[poolId].poolAddress;
    }

    /**
     * @notice 풀 주소로 풀 ID 조회
     * @param poolAddress 풀 주소
     * @return 풀 ID
     */
    function getPoolId(address poolAddress) external view returns (uint) {
        uint poolId = poolIds[poolAddress];
        require(poolId != 0, CSPoolNotFound());
        return poolId;
    }

    /**
     * @notice 활성 풀만 필터링하여 조회
     * @return 활성 풀 ID 배열
     */
    function getActivePoolIds() external view returns (uint[] memory) {
        uint totalCount = _allPoolIds.length();
        uint[] memory tempIds = new uint[](totalCount);
        uint activeCount = 0;

        for (uint i = 0; i < totalCount; i++) {
            uint poolId = _allPoolIds.at(i);
            if (pools[poolId].active) {
                tempIds[activeCount] = poolId;
                activeCount++;
            }
        }

        // Resize array
        uint[] memory activeIds = new uint[](activeCount);
        for (uint i = 0; i < activeCount; i++) {
            activeIds[i] = tempIds[i];
        }

        return activeIds;
    }

    // ==================== UUPS ====================

    /**
     * @dev 업그레이드 권한 체크
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}

    // ==================== Storage Gap ====================

    /**
     * @dev 향후 업그레이드를 위한 storage gap
     * 현재 사용: 8 slots
     * Gap: 50 - 8 = 42 slots
     */
    uint[42] private __gap;
}
