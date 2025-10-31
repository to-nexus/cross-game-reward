// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControlDefaultAdminRulesUpgradeable} from
    "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardTransientUpgradeable} from
    "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardTransientUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {ICrossStaking} from "./interfaces/ICrossStaking.sol";
import {ICrossStakingPool} from "./interfaces/ICrossStakingPool.sol";

/**
 * @title CrossStakingPool
 * @notice CROSS 토큰 스테이킹 풀 - 실시간 보상 입금 방식
 * @dev rewardPerToken 누적 방식 + UUPS Upgradeable
 *
 * === 핵심 원리 ===
 *
 * 보상 분배:
 *   - 예치 이후에 입금된 보상만 받을 수 있음
 *   - 보상 입금 시점의 금액 비율로 분배
 *   - rewardPerToken 누적 방식
 *
 * === 공정성 ===
 *
 * 먼저 예치한 사람:
 *   - 더 오래 보상을 받음
 *   - 각 보상은 현재 예치 금액 비율대로
 *
 * === 업그레이더블 ===
 *
 * UUPS 패턴:
 *   - AccessControlDefaultAdminRules로 관리
 *   - Pausable 기능 내장
 */
contract CrossStakingPool is
    Initializable,
    AccessControlDefaultAdminRulesUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardTransientUpgradeable,
    UUPSUpgradeable,
    ICrossStakingPool
{
    using SafeERC20 for IERC20;
    using EnumerableSet for EnumerableSet.AddressSet;

    // ==================== 커스텀 에러 ====================

    error CSPBelowMinimumStakeAmount();
    error CSPNoStakeFound();
    error CSPCanNotZeroAddress();
    error CSPCanNotZeroValue();
    error CSPRewardTokenAlreadyAdded();
    error CSPInvalidRewardToken();
    error CSPCanNotUseStakingToken();
    error CSPOnlyRouter();
    error CSPNoWithdrawableAmount();

    // ==================== Roles ====================

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant REWARD_MANAGER_ROLE = keccak256("REWARD_MANAGER_ROLE");

    // ==================== 상수 ====================

    uint public constant MIN_STAKE_AMOUNT = 1 ether;
    uint private constant PRECISION = 1e18;

    // ==================== 상태 변수 ====================

    // 스테이킹 토큰 (CROSS)
    IERC20 public stakingToken;

    // CrossStaking 주소 (Router 확인용)
    address public crossStaking;

    // 보상 토큰 주소 집합 (EnumerableSet 사용)
    EnumerableSet.AddressSet private _rewardTokenAddresses;

    // 보상 토큰 주소 -> 데이터 매핑
    mapping(address => RewardToken) private _rewardTokenData;

    // 사용자별 예치 금액
    mapping(address => uint) public balances;

    // 사용자별, 보상토큰별 정보
    mapping(address => mapping(address => UserReward)) public userRewards;

    // 전체 예치량
    uint public totalStaked;

    // ==================== 이벤트 ====================

    event Staked(address indexed user, uint amount);
    event Unstaked(address indexed user, uint amount);
    event RewardClaimed(address indexed user, address indexed rewardToken, uint amount);
    event RewardTokenAdded(address indexed rewardToken);
    event RewardTokenRemoved(address indexed rewardToken);
    event RewardDeposited(address indexed sender, address indexed rewardToken, uint amount);
    event RewardDistributed(address indexed rewardToken, uint amount);
    event EmergencyWithdraw(address indexed rewardToken, address indexed to, uint amount);

    // ==================== Initializer ====================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    fallback() external {
        revert("Not allowed");
    }

    /**
     * @notice 컨트랙트 초기화
     * @param _stakingToken CROSS 토큰 주소
     * @param _admin 관리자 주소
     * @param _initialDelay 관리자 변경 딜레이 (초)
     */
    function initialize(IERC20 _stakingToken, address _admin, uint48 _initialDelay) external initializer {
        require(address(_stakingToken) != address(0), CSPCanNotZeroAddress());
        require(_admin != address(0), CSPCanNotZeroAddress());

        __AccessControlDefaultAdminRules_init(_initialDelay, _admin);
        __Pausable_init();
        __ReentrancyGuardTransient_init();
        __UUPSUpgradeable_init();

        stakingToken = _stakingToken;
        crossStaking = _admin; // admin은 CrossStaking 컨트랙트

        // 기본 역할 부여
        _grantRole(PAUSER_ROLE, _admin);
        _grantRole(REWARD_MANAGER_ROLE, _admin);
    }

    // ==================== 외부 함수 ====================

    /**
     * @notice CROSS 토큰 스테이킹
     * @param amount 예치할 CROSS 수량
     * @dev 추가 스테이킹 시 기존 금액에 누적됨
     */
    function stake(uint amount) external nonReentrant whenNotPaused {
        _stake(msg.sender, msg.sender, amount);
    }

    /**
     * @notice 다른 계정을 위해 스테이킹 (Router 전용)
     * @param account 스테이킹될 계정
     * @param amount 예치할 CROSS 수량
     * @dev msg.sender가 등록된 Router인지 확인
     */
    function stakeFor(address account, uint amount) external nonReentrant whenNotPaused {
        _checkDelegate(account);
        _stake(msg.sender, account, amount);
    }

    /**
     * @notice 전체 수량 unstake 및 보상 자동 클레임
     * @dev 누적된 모든 보상을 함께 수령
     */
    function unstake() external nonReentrant whenNotPaused {
        _unstake(msg.sender, msg.sender);
    }

    /**
     * @notice 다른 계정을 위해 언스테이킹 (Router 전용)
     * @param account 언스테이킹할 계정
     * @dev msg.sender가 등록된 Router인지 확인
     */
    function unstakeFor(address account) external nonReentrant whenNotPaused {
        _checkDelegate(account);
        _unstake(msg.sender, account);
    }

    /**
     * @notice 보상만 클레임 (스테이킹 유지)
     * @dev CROSS 토큰은 그대로 예치 상태로 유지
     */
    function claimRewards() external nonReentrant whenNotPaused {
        require(balances[msg.sender] > 0, CSPNoStakeFound());

        _syncReward();
        _updateRewards(msg.sender);
        _claimRewards(msg.sender);
    }

    /**
     * @notice 특정 보상 토큰만 클레임
     * @param tokenAddress 클레임할 보상 토큰 주소
     * @dev 제거된 토큰도 기존 보상은 claim 가능
     */
    function claimReward(address tokenAddress) external nonReentrant whenNotPaused {
        require(balances[msg.sender] > 0, CSPNoStakeFound());
        // 제거된 토큰도 claim 가능하도록 _rewardTokenData 존재 여부만 체크
        require(_rewardTokenData[tokenAddress].tokenAddress != address(0), CSPInvalidRewardToken());

        // 제거되지 않은 토큰만 sync
        if (_rewardTokenAddresses.contains(tokenAddress)) _syncReward(tokenAddress);

        _updateReward(tokenAddress, msg.sender);
        _claimReward(tokenAddress, msg.sender);
    }

    /**
     * @notice 사용자의 pending 보상 조회
     * @param user 조회할 사용자 주소
     * @return rewards 각 보상 토큰별 pending 금액 배열
     */
    function pendingRewards(address user) external view returns (uint[] memory rewards) {
        uint length = _rewardTokenAddresses.length();
        rewards = new uint[](length);

        for (uint i = 0; i < length; i++) {
            address tokenAddress = _rewardTokenAddresses.at(i);
            rewards[i] = _calculatePendingReward(tokenAddress, user);
        }
    }

    /**
     * @notice 보상 토큰 주소 조회 (인덱스로)
     * @param index 인덱스
     * @return tokenAddress 보상 토큰 주소
     */
    function rewardTokenAt(uint index) external view returns (address) {
        return _rewardTokenAddresses.at(index);
    }

    /**
     * @notice 보상 토큰 데이터 조회
     * @param tokenAddress 보상 토큰 주소
     * @return 보상 토큰 데이터
     */
    function getRewardToken(address tokenAddress) external view returns (RewardToken memory) {
        require(_rewardTokenAddresses.contains(tokenAddress), CSPInvalidRewardToken());
        return _rewardTokenData[tokenAddress];
    }

    /**
     * @notice 보상 토큰 등록 여부 확인
     * @param tokenAddress 확인할 보상 토큰 주소
     * @return 등록 여부
     */
    function isRewardToken(address tokenAddress) external view returns (bool) {
        return _rewardTokenAddresses.contains(tokenAddress);
    }

    /**
     * @notice 보상 토큰 주소 목록 조회
     * @return 모든 보상 토큰 주소 배열
     */
    function getRewardTokens() external view returns (address[] memory) {
        return _rewardTokenAddresses.values();
    }

    /**
     * @notice 보상 토큰 길이 조회
     * @return 보상 토큰 개수
     */
    function rewardTokensLength() external view returns (uint) {
        return _rewardTokenAddresses.length();
    }

    // ==================== 관리자 함수 ====================

    /**
     * @notice 보상 토큰 추가
     * @param tokenAddress 추가할 보상 토큰 주소
     */
    function addRewardToken(address tokenAddress) external onlyRole(REWARD_MANAGER_ROLE) {
        require(tokenAddress != address(0), CSPCanNotZeroAddress());
        require(tokenAddress != address(stakingToken), CSPCanNotUseStakingToken());
        require(_rewardTokenAddresses.add(tokenAddress), CSPRewardTokenAlreadyAdded());

        _rewardTokenData[tokenAddress] = RewardToken({
            tokenAddress: tokenAddress,
            rewardPerTokenStored: 0,
            lastBalance: 0,
            removedDistributedAmount: 0,
            isRemoved: false
        });

        emit RewardTokenAdded(tokenAddress);
    }

    /**
     * @notice 보상 토큰 개수 조회
     */
    function rewardTokenCount() external view returns (uint) {
        return _rewardTokenAddresses.length();
    }

    /**
     * @notice 보상 토큰 제거
     * @param tokenAddress 제거할 보상 토큰 주소
     * @dev 기존 누적 보상은 여전히 claim 가능
     */
    function removeRewardToken(address tokenAddress) external onlyRole(REWARD_MANAGER_ROLE) {
        require(_rewardTokenAddresses.contains(tokenAddress), CSPInvalidRewardToken());

        // 마지막 동기화
        _syncReward(tokenAddress);

        // 제거 시점의 실제 잔액 저장
        RewardToken storage rt = _rewardTokenData[tokenAddress];
        uint currentBalance = IERC20(tokenAddress).balanceOf(address(this));
        rt.removedDistributedAmount = currentBalance;
        rt.isRemoved = true;

        // EnumerableSet에서 제거
        _rewardTokenAddresses.remove(tokenAddress);

        emit RewardTokenRemoved(tokenAddress);
    }

    /**
     * @notice 비상 출금 가능 금액 조회
     * @param tokenAddress 보상 토큰 주소
     * @return 출금 가능 금액 (제거 이후 입금액)
     */
    function getEmergencyWithdrawableAmount(address tokenAddress) public view returns (uint) {
        RewardToken storage rt = _rewardTokenData[tokenAddress];
        if (!rt.isRemoved) return 0; // 제거되지 않은 토큰

        uint currentBalance = IERC20(tokenAddress).balanceOf(address(this));
        return currentBalance > rt.removedDistributedAmount ? currentBalance - rt.removedDistributedAmount : 0;
    }

    /**
     * @notice 비상 출금 (제거된 토큰의 제거 이후 입금액만)
     * @param tokenAddress 보상 토큰 주소
     * @param to 출금 받을 주소
     */
    function emergencyWithdraw(address tokenAddress, address to) external onlyRole(DEFAULT_ADMIN_ROLE) {
        uint amount = getEmergencyWithdrawableAmount(tokenAddress);
        require(amount > 0, CSPNoWithdrawableAmount());
        require(to != address(0), CSPCanNotZeroAddress());

        IERC20(tokenAddress).safeTransfer(to, amount);
        emit EmergencyWithdraw(tokenAddress, to, amount);
    }

    /**
     * @notice 긴급 정지
     * @dev stake, unstake, claim 기능 차단
     */
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @notice 긴급 정지 해제
     */
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    // ==================== 내부 함수: 보상 동기화 ====================

    /**
     * @dev 모든 보상 토큰 동기화
     */
    function _syncReward() internal {
        uint length = _rewardTokenAddresses.length();
        for (uint i = 0; i < length; i++) {
            address tokenAddress = _rewardTokenAddresses.at(i);
            _syncReward(tokenAddress);
        }
    }

    /**
     * @dev 새 보상 감지 및 rewardPerToken 업데이트
     * @param tokenAddress 보상 토큰 주소
     */
    function _syncReward(address tokenAddress) internal {
        // 스테이킹이 없는상태에서는 동기화하지 않음
        if (totalStaked == 0) return;

        RewardToken storage rt = _rewardTokenData[tokenAddress];

        uint currentBalance = IERC20(rt.tokenAddress).balanceOf(address(this));

        if (currentBalance > rt.lastBalance && totalStaked > 0) {
            uint newReward = currentBalance - rt.lastBalance;
            rt.rewardPerTokenStored += (newReward * PRECISION) / totalStaked;
            emit RewardDistributed(rt.tokenAddress, newReward);
        }

        rt.lastBalance = currentBalance;
    }

    // ==================== 내부 함수: 보상 업데이트 ====================

    /**
     * @dev 모든 보상 토큰에 대해 사용자 보상 업데이트
     */
    function _updateRewards(address user) internal {
        uint length = _rewardTokenAddresses.length();
        for (uint i = 0; i < length; i++) {
            address tokenAddress = _rewardTokenAddresses.at(i);
            _updateReward(tokenAddress, user);
        }
    }

    /**
     * @dev 사용자 보상 계산 및 체크포인트 갱신
     * @param tokenAddress 보상 토큰 주소
     * @param user 사용자 주소
     */
    function _updateReward(address tokenAddress, address user) internal {
        RewardToken storage rt = _rewardTokenData[tokenAddress];
        UserReward storage ur = userRewards[user][tokenAddress];

        uint userBalance = balances[user];

        if (userBalance > 0) {
            uint earned = (userBalance * (rt.rewardPerTokenStored - ur.rewardPerTokenPaid)) / PRECISION;
            ur.rewards += earned;
        }

        ur.rewardPerTokenPaid = rt.rewardPerTokenStored;
    }

    /**
     * @dev Pending 보상 계산 (view)
     * @param tokenAddress 보상 토큰 주소
     * @param user 사용자 주소
     * @return 계산된 pending 보상
     */
    function _calculatePendingReward(address tokenAddress, address user) internal view returns (uint) {
        UserReward storage ur = userRewards[user][tokenAddress];
        RewardToken storage rt = _rewardTokenData[tokenAddress];

        uint userBalance = balances[user];
        if (userBalance == 0) return ur.rewards;

        uint currentBalance = IERC20(rt.tokenAddress).balanceOf(address(this));
        uint currentRewardPerToken = rt.rewardPerTokenStored;

        if (currentBalance > rt.lastBalance && totalStaked > 0) {
            uint newReward = currentBalance - rt.lastBalance;
            currentRewardPerToken += (newReward * PRECISION) / totalStaked;
        }

        uint earned = (userBalance * (currentRewardPerToken - ur.rewardPerTokenPaid)) / PRECISION;
        return ur.rewards + earned;
    }

    // ==================== 내부 함수: 보상 Claim ====================

    /**
     * @dev 모든 보상 토큰 클레임
     */
    function _claimRewards(address user) internal {
        uint length = _rewardTokenAddresses.length();
        for (uint i = 0; i < length; i++) {
            address tokenAddress = _rewardTokenAddresses.at(i);
            _claimReward(tokenAddress, user);
        }
    }

    /**
     * @dev 보상 전송 및 잔액 동기화
     * @param tokenAddress 보상 토큰 주소
     * @param user 사용자 주소
     */
    function _claimReward(address tokenAddress, address user) internal {
        UserReward storage ur = userRewards[user][tokenAddress];
        uint reward = ur.rewards;

        if (reward > 0) {
            ur.rewards = 0;

            RewardToken storage rt = _rewardTokenData[tokenAddress];
            IERC20(rt.tokenAddress).safeTransfer(user, reward);

            rt.lastBalance -= reward;

            // 제거된 토큰이면 분배 예정량 차감
            if (rt.isRemoved) rt.removedDistributedAmount -= reward;

            emit RewardClaimed(user, rt.tokenAddress, reward);
        }
    }

    // ==================== 내부 함수: Stake/Unstake ====================

    /**
     * @dev 내부 스테이킹 로직
     * @param payer 토큰을 전송하는 주소
     * @param account 스테이킹될 계정
     * @param amount 예치할 수량
     */
    function _stake(address payer, address account, uint amount) internal {
        require(amount >= MIN_STAKE_AMOUNT, CSPBelowMinimumStakeAmount());

        _syncReward();
        _updateRewards(account);

        stakingToken.safeTransferFrom(payer, address(this), amount);

        balances[account] += amount;
        totalStaked += amount;

        emit Staked(account, amount);
    }

    /**
     * @dev 내부 언스테이킹 로직
     * @param account 언스테이킹할 계정
     */
    function _unstake(address, /* caller */ address account) internal {
        require(balances[account] > 0, CSPNoStakeFound());

        uint amount = balances[account];

        _syncReward();
        _updateRewards(account);
        _claimRewards(account);

        totalStaked -= amount;
        stakingToken.safeTransfer(account, amount);

        delete balances[account];

        emit Unstaked(account, amount);
    }

    function _checkDelegate(address account) internal view {
        require(account != address(0), CSPCanNotZeroAddress());
        require(msg.sender == ICrossStaking(crossStaking).router(), CSPOnlyRouter());
    }

    // ==================== UUPS ====================

    /**
     * @dev 업그레이드 권한 체크
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}

    // ==================== Storage Gap ====================

    /**
     * @dev 향후 업그레이드를 위한 storage gap
     * 현재 사용: 6 slots
     * Gap: 50 - 6 = 44 slots
     */
    uint[44] private __gap;
}
