// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardTransientUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

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
    UUPSUpgradeable
{
    using SafeERC20 for IERC20;

    // ==================== 커스텀 에러 ====================

    error BelowMinimumStakeAmount();
    error NoStakeFound();
    error InvalidRewardTokenIndex();
    error InvalidTokenAddress();
    error AmountMustBeGreaterThanZero();
    error RewardTokenAlreadyAdded();
    error InvalidRewardToken();
    error CannotUseStakingTokenAsReward();

    // ==================== Roles ====================

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant REWARD_MANAGER_ROLE = keccak256("REWARD_MANAGER_ROLE");

    // ==================== 상수 ====================

    uint public constant MIN_STAKE_AMOUNT = 1 ether;
    uint private constant PRECISION = 1e18;

    // ==================== 상태 변수 ====================

    // 스테이킹 토큰 (CROSS)
    IERC20 public stakingToken;

    /**
     * @dev 보상 토큰 정보
     */
    struct RewardToken {
        address tokenAddress;
        uint rewardPerTokenStored; // 누적 토큰당 보상
        uint lastBalance; // 마지막 기록 잔액
    }

    /**
     * @dev 사용자별 보상 정보
     */
    struct UserReward {
        uint rewardPerTokenPaid; // 예치 시점의 rewardPerToken
        uint rewards; // 누적 보상
    }

    // 보상 토큰 배열
    RewardToken[] public rewardTokens;

    // 보상 토큰 주소 -> 인덱스 매핑 (O(1) lookup)
    mapping(address => uint) public tokenToIndex;

    // 보상 토큰 등록 여부
    mapping(address => bool) public isRewardToken;

    // 사용자별 예치 금액
    mapping(address => uint) public balances;

    // 사용자별, 보상토큰별 정보
    mapping(address => mapping(uint => UserReward)) public userRewards;

    // 전체 예치량
    uint public totalStaked;

    // ==================== 이벤트 ====================

    event Staked(address indexed user, uint amount);
    event Unstaked(address indexed user, uint amount);
    event RewardClaimed(address indexed user, address indexed rewardToken, uint amount);
    event RewardTokenAdded(address indexed rewardToken);
    event RewardDeposited(address indexed sender, address indexed rewardToken, uint amount);
    event RewardDistributed(address indexed rewardToken, uint amount);

    // ==================== Initializer ====================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice 컨트랙트 초기화
     * @param _stakingToken CROSS 토큰 주소
     * @param _admin 관리자 주소
     * @param _initialDelay 관리자 변경 딜레이 (초)
     */
    function initialize(IERC20 _stakingToken, address _admin, uint48 _initialDelay) external initializer {
        require(address(_stakingToken) != address(0), InvalidTokenAddress());
        require(_admin != address(0), "Invalid admin address");

        __AccessControlDefaultAdminRules_init(_initialDelay, _admin);
        __Pausable_init();
        __ReentrancyGuardTransient_init();
        __UUPSUpgradeable_init();

        stakingToken = _stakingToken;

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
        require(amount >= MIN_STAKE_AMOUNT, BelowMinimumStakeAmount());

        _syncReward();
        _updateRewards(msg.sender);

        stakingToken.safeTransferFrom(msg.sender, address(this), amount);

        balances[msg.sender] += amount;
        totalStaked += amount;

        emit Staked(msg.sender, amount);
    }

    /**
     * @notice 전체 수량 unstake 및 보상 자동 클레임
     * @dev 누적된 모든 보상을 함께 수령
     */
    function unstake() external nonReentrant whenNotPaused {
        require(balances[msg.sender] > 0, NoStakeFound());

        uint amount = balances[msg.sender];

        _syncReward();
        _updateRewards(msg.sender);
        _claimRewards(msg.sender);

        stakingToken.safeTransfer(msg.sender, amount);

        totalStaked -= amount;
        delete balances[msg.sender];

        emit Unstaked(msg.sender, amount);
    }

    /**
     * @notice 보상만 클레임 (스테이킹 유지)
     * @dev CROSS 토큰은 그대로 예치 상태로 유지
     */
    function claimRewards() external nonReentrant whenNotPaused {
        require(balances[msg.sender] > 0, NoStakeFound());

        _syncReward();
        _updateRewards(msg.sender);
        _claimRewards(msg.sender);
    }

    /**
     * @notice 특정 보상 토큰만 클레임
     * @param rewardTokenIndex 클레임할 보상 토큰 인덱스
     */
    function claimReward(uint rewardTokenIndex) external nonReentrant whenNotPaused {
        require(balances[msg.sender] > 0, NoStakeFound());
        require(rewardTokenIndex < rewardTokens.length, InvalidRewardTokenIndex());

        _syncReward(rewardTokenIndex);
        _updateReward(rewardTokenIndex, msg.sender);
        _claimReward(rewardTokenIndex, msg.sender);
    }

    /**
     * @notice 사용자의 pending 보상 조회
     * @param user 조회할 사용자 주소
     * @return rewards 각 보상 토큰별 pending 금액 배열
     */
    function pendingRewards(address user) external view returns (uint[] memory rewards) {
        rewards = new uint[](rewardTokens.length);

        for (uint i = 0; i < rewardTokens.length; i++) {
            rewards[i] = _calculatePendingReward(i, user);
        }
    }

    // ==================== 관리자 함수 ====================

    /**
     * @notice 보상 토큰 추가
     * @param tokenAddress 추가할 보상 토큰 주소
     */
    function addRewardToken(address tokenAddress) external onlyRole(REWARD_MANAGER_ROLE) {
        require(tokenAddress != address(0), InvalidTokenAddress());
        require(tokenAddress != address(stakingToken), CannotUseStakingTokenAsReward());
        require(!isRewardToken[tokenAddress], RewardTokenAlreadyAdded());

        uint index = rewardTokens.length;
        rewardTokens.push(RewardToken({tokenAddress: tokenAddress, rewardPerTokenStored: 0, lastBalance: 0}));

        tokenToIndex[tokenAddress] = index;
        isRewardToken[tokenAddress] = true;

        emit RewardTokenAdded(tokenAddress);
    }

    /**
     * @notice 보상 입금
     * @param tokenAddress 보상 토큰 주소
     * @param amount 입금할 수량
     * @dev 직접 transfer된 금액도 다음 동기화 시 자동 반영됨
     */
    function depositReward(address tokenAddress, uint amount) external {
        require(amount > 0, AmountMustBeGreaterThanZero());
        require(isRewardToken[tokenAddress], InvalidRewardToken());

        IERC20(tokenAddress).safeTransferFrom(msg.sender, address(this), amount);

        uint index = tokenToIndex[tokenAddress];
        _syncReward(index);

        emit RewardDeposited(msg.sender, tokenAddress, amount);
    }

    /**
     * @notice 보상 토큰 개수 조회
     */
    function rewardTokenCount() external view returns (uint) {
        return rewardTokens.length;
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
        for (uint i = 0; i < rewardTokens.length; i++) {
            _syncReward(i);
        }
    }

    /**
     * @dev 새 보상 감지 및 rewardPerToken 업데이트
     * @param rewardTokenIndex 보상 토큰 인덱스
     */
    function _syncReward(uint rewardTokenIndex) internal {
        RewardToken storage rt = rewardTokens[rewardTokenIndex];

        uint currentBalance = IERC20(rt.tokenAddress).balanceOf(address(this));

        if (currentBalance > rt.lastBalance) {
            uint newReward = currentBalance - rt.lastBalance;

            if (totalStaked > 0) {
                rt.rewardPerTokenStored += (newReward * PRECISION) / totalStaked;
                emit RewardDistributed(rt.tokenAddress, newReward);
            }
        }

        rt.lastBalance = currentBalance;
    }

    // ==================== 내부 함수: 보상 업데이트 ====================

    /**
     * @dev 모든 보상 토큰에 대해 사용자 보상 업데이트
     */
    function _updateRewards(address user) internal {
        for (uint i = 0; i < rewardTokens.length; i++) {
            _updateReward(i, user);
        }
    }

    /**
     * @dev 사용자 보상 계산 및 체크포인트 갱신
     * @param rewardTokenIndex 보상 토큰 인덱스
     * @param user 사용자 주소
     */
    function _updateReward(uint rewardTokenIndex, address user) internal {
        RewardToken storage rt = rewardTokens[rewardTokenIndex];
        UserReward storage ur = userRewards[user][rewardTokenIndex];

        uint userBalance = balances[user];

        if (userBalance > 0) {
            uint earned = (userBalance * (rt.rewardPerTokenStored - ur.rewardPerTokenPaid)) / PRECISION;
            ur.rewards += earned;
        }

        ur.rewardPerTokenPaid = rt.rewardPerTokenStored;
    }

    /**
     * @dev Pending 보상 계산 (view)
     * @param rewardTokenIndex 보상 토큰 인덱스
     * @param user 사용자 주소
     * @return 계산된 pending 보상
     */
    function _calculatePendingReward(uint rewardTokenIndex, address user) internal view returns (uint) {
        UserReward storage ur = userRewards[user][rewardTokenIndex];
        RewardToken storage rt = rewardTokens[rewardTokenIndex];

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
        for (uint i = 0; i < rewardTokens.length; i++) {
            _claimReward(i, user);
        }
    }

    /**
     * @dev 보상 전송 및 잔액 동기화
     * @param rewardTokenIndex 보상 토큰 인덱스
     * @param user 사용자 주소
     */
    function _claimReward(uint rewardTokenIndex, address user) internal {
        UserReward storage ur = userRewards[user][rewardTokenIndex];
        uint reward = ur.rewards;

        if (reward > 0) {
            ur.rewards = 0;

            RewardToken storage rt = rewardTokens[rewardTokenIndex];
            IERC20(rt.tokenAddress).safeTransfer(user, reward);

            rt.lastBalance = IERC20(rt.tokenAddress).balanceOf(address(this));

            emit RewardClaimed(user, rt.tokenAddress, reward);
        }
    }

    // ==================== UUPS ====================

    /**
     * @dev 업그레이드 권한 체크
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}

    // ==================== Storage Gap ====================

    /**
     * @dev 향후 업그레이드를 위한 storage gap
     */
    uint[43] private __gap;
}
