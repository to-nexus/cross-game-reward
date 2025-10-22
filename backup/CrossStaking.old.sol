// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title CrossStaking
 * @notice CROSS Staking Protocol v0.1 - 토큰 스테이킹 및 보상 분배 컨트랙트
 * @dev 스테이킹 파워는 스테이킹 수량과 남은 기간에 비례하여 계산됩니다
 */
contract CrossStaking is Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    // ============ 상수 ============

    /// @notice 최소 스테이킹 기간 (1주)
    uint256 public constant MIN_LOCK_TIME = 1 weeks; // 604800초

    /// @notice 최대 스테이킹 기간 (5주)
    uint256 public constant MAX_LOCK_TIME = 5 weeks; // 3024000초

    /// @notice 스테이킹 주기 (1주)
    uint256 public constant EPOCH_DURATION = 1 weeks;

    // ============ 구조체 ============

    /// @notice 스테이킹 포지션 정보
    struct StakePosition {
        uint256 amount; // 스테이킹한 토큰 수량
        uint256 expiry; // 스테이킹 만료 시간 (Unix timestamp)
        uint256 lastUpdateTime; // 마지막 업데이트 시간
    }

    /// @notice 보상 토큰 정보
    struct RewardToken {
        address tokenAddress; // 보상 토큰 주소
        uint256 rewardPerPower; // 스테이킹 파워당 누적 보상량 (scaled by 1e18)
        bool isActive; // 활성화 여부
    }

    /// @notice 사용자별 보상 정보
    struct UserReward {
        uint256 rewardDebt; // 보상 부채 (이미 계산된 보상)
        uint256 pendingReward; // 대기 중인 보상
    }

    // ============ 상태 변수 ============

    /// @notice 스테이킹 토큰 주소
    IERC20 public immutable stakingToken;

    /// @notice 사용자별 스테이킹 포지션
    mapping(address => StakePosition) public stakePositions;

    /// @notice 전체 스테이킹 파워 (실시간 계산을 위한 캐시)
    uint256 public totalStakingPower;

    /// @notice 마지막 전체 스테이킹 파워 업데이트 시간
    uint256 public lastTotalPowerUpdate;

    /// @notice 보상 토큰 목록
    address[] public rewardTokenList;

    /// @notice 보상 토큰 정보 매핑
    mapping(address => RewardToken) public rewardTokens;

    /// @notice 사용자별 보상 토큰별 보상 정보
    mapping(address => mapping(address => UserReward)) public userRewards;

    /// @notice 현재 에포크 번호
    uint256 public currentEpoch;

    /// @notice 프로토콜 시작 시간
    uint256 public immutable protocolStartTime;

    // ============ 이벤트 ============

    event Staked(address indexed user, uint256 amount, uint256 expiry, uint256 stakingPower);
    event AmountIncreased(address indexed user, uint256 additionalAmount, uint256 newStakingPower);
    event DurationExtended(address indexed user, uint256 newExpiry, uint256 newStakingPower);
    event Unstaked(address indexed user, uint256 amount);
    event RewardTokenAdded(address indexed token);
    event RewardTokenRemoved(address indexed token);
    event RewardsDistributed(address indexed token, uint256 amount, uint256 epoch);
    event RewardsClaimed(address indexed user, address indexed token, uint256 amount);
    event StakingPowerUpdated(address indexed user, uint256 newPower);

    // ============ 생성자 ============

    /**
     * @notice CrossStaking 컨트랙트 생성자
     * @param _stakingToken 스테이킹할 토큰 주소
     * @param _owner 컨트랙트 소유자 주소
     */
    constructor(address _stakingToken, address _owner) Ownable(_owner) {
        require(_stakingToken != address(0), "Invalid staking token");
        stakingToken = IERC20(_stakingToken);
        protocolStartTime = block.timestamp;
        lastTotalPowerUpdate = block.timestamp;
        currentEpoch = 0;
    }

    // ============ 스테이킹 함수 ============

    /**
     * @notice 토큰 스테이킹
     * @param amount 스테이킹할 토큰 수량
     * @param duration 스테이킹 기간 (초 단위)
     */
    function stake(uint256 amount, uint256 duration) external nonReentrant {
        require(amount > 0, "Amount must be greater than 0");
        require(duration >= MIN_LOCK_TIME, "Duration too short");
        require(duration <= MAX_LOCK_TIME, "Duration too long");
        require(stakePositions[msg.sender].amount == 0, "Position already exists");

        // 토큰 전송
        stakingToken.safeTransferFrom(msg.sender, address(this), amount);

        // 만료 시간 계산
        uint256 expiry = block.timestamp + duration;

        // 스테이킹 포지션 생성
        stakePositions[msg.sender] = StakePosition({amount: amount, expiry: expiry, lastUpdateTime: block.timestamp});

        // 스테이킹 파워 계산 및 업데이트
        uint256 stakingPower = _calculateStakingPower(amount, expiry);
        _updateTotalStakingPower(int256(stakingPower));

        // 보상 부채 초기화
        _updateUserRewardDebt(msg.sender);

        emit Staked(msg.sender, amount, expiry, stakingPower);
    }

    /**
     * @notice 스테이킹 수량 증가
     * @param additionalAmount 추가할 토큰 수량
     */
    function increaseAmount(uint256 additionalAmount) external nonReentrant {
        require(additionalAmount > 0, "Amount must be greater than 0");

        StakePosition storage position = stakePositions[msg.sender];
        require(position.amount > 0, "No staking position");
        require(block.timestamp < position.expiry, "Position expired");

        // 기존 스테이킹 파워 계산
        uint256 oldStakingPower = _calculateStakingPower(position.amount, position.expiry);

        // 보상 업데이트
        _updateUserRewards(msg.sender);

        // 토큰 전송
        stakingToken.safeTransferFrom(msg.sender, address(this), additionalAmount);

        // 수량 증가
        position.amount += additionalAmount;
        position.lastUpdateTime = block.timestamp;

        // 새로운 스테이킹 파워 계산
        uint256 newStakingPower = _calculateStakingPower(position.amount, position.expiry);

        // 전체 스테이킹 파워 업데이트
        _updateTotalStakingPower(int256(newStakingPower) - int256(oldStakingPower));

        // 보상 부채 업데이트
        _updateUserRewardDebt(msg.sender);

        emit AmountIncreased(msg.sender, additionalAmount, newStakingPower);
    }

    /**
     * @notice 스테이킹 기간 연장
     * @param newExpiry 새로운 만료 시간 (Unix timestamp)
     */
    function extendDuration(uint256 newExpiry) external nonReentrant {
        StakePosition storage position = stakePositions[msg.sender];
        require(position.amount > 0, "No staking position");
        require(newExpiry > position.expiry, "New expiry must be greater");
        require(newExpiry <= block.timestamp + MAX_LOCK_TIME, "Exceeds max lock time");

        // 기존 스테이킹 파워 계산
        uint256 oldStakingPower = _calculateStakingPower(position.amount, position.expiry);

        // 보상 업데이트
        _updateUserRewards(msg.sender);

        // 만료 시간 연장
        position.expiry = newExpiry;
        position.lastUpdateTime = block.timestamp;

        // 새로운 스테이킹 파워 계산
        uint256 newStakingPower = _calculateStakingPower(position.amount, position.expiry);

        // 전체 스테이킹 파워 업데이트
        _updateTotalStakingPower(int256(newStakingPower) - int256(oldStakingPower));

        // 보상 부채 업데이트
        _updateUserRewardDebt(msg.sender);

        emit DurationExtended(msg.sender, newExpiry, newStakingPower);
    }

    /**
     * @notice 스테이킹 해제
     */
    function unstake() external nonReentrant {
        StakePosition storage position = stakePositions[msg.sender];
        require(position.amount > 0, "No staking position");
        require(block.timestamp >= position.expiry, "Position not expired");

        uint256 amount = position.amount;

        // 보상 업데이트 및 청구
        _updateUserRewards(msg.sender);
        _claimAllRewards(msg.sender);

        // 스테이킹 파워 업데이트 (만료되었으므로 0)
        _updateTotalStakingPower(0);

        // 포지션 삭제
        delete stakePositions[msg.sender];

        // 토큰 반환
        stakingToken.safeTransfer(msg.sender, amount);

        emit Unstaked(msg.sender, amount);
    }

    // ============ 보상 관련 함수 ============

    /**
     * @notice 보상 토큰 추가 (관리자 전용)
     * @param tokenAddress 보상 토큰 주소
     */
    function addRewardToken(address tokenAddress) external onlyOwner {
        require(tokenAddress != address(0), "Invalid token address");
        require(!rewardTokens[tokenAddress].isActive, "Token already added");

        rewardTokens[tokenAddress] = RewardToken({tokenAddress: tokenAddress, rewardPerPower: 0, isActive: true});

        rewardTokenList.push(tokenAddress);

        emit RewardTokenAdded(tokenAddress);
    }

    /**
     * @notice 보상 토큰 제거 (관리자 전용)
     * @param tokenAddress 보상 토큰 주소
     */
    function removeRewardToken(address tokenAddress) external onlyOwner {
        require(rewardTokens[tokenAddress].isActive, "Token not active");

        rewardTokens[tokenAddress].isActive = false;

        emit RewardTokenRemoved(tokenAddress);
    }

    /**
     * @notice 보상 분배 (관리자 전용)
     * @param tokenAddress 보상 토큰 주소
     * @param amount 분배할 보상량
     */
    function distributeRewards(address tokenAddress, uint256 amount) external onlyOwner {
        require(rewardTokens[tokenAddress].isActive, "Token not active");
        require(amount > 0, "Amount must be greater than 0");

        uint256 totalPower = getTotalStakingPower();
        require(totalPower > 0, "No staking power");

        // 보상 토큰 전송
        IERC20(tokenAddress).safeTransferFrom(msg.sender, address(this), amount);

        // 스테이킹 파워당 보상량 계산 (1e18로 스케일링)
        uint256 rewardPerPower = (amount * 1e18) / totalPower;
        rewardTokens[tokenAddress].rewardPerPower += rewardPerPower;

        currentEpoch++;

        emit RewardsDistributed(tokenAddress, amount, currentEpoch);
    }

    /**
     * @notice 특정 보상 토큰 청구
     * @param tokenAddress 보상 토큰 주소
     */
    function claimReward(address tokenAddress) public nonReentrant {
        require(rewardTokens[tokenAddress].isActive, "Token not active");

        _updateUserRewards(msg.sender);

        UserReward storage userReward = userRewards[msg.sender][tokenAddress];
        uint256 pending = userReward.pendingReward;

        if (pending > 0) {
            userReward.pendingReward = 0;

            // 보상 부채 업데이트 (현재 스테이킹 파워 기준)
            uint256 stakingPower = getStakingPower(msg.sender);
            userReward.rewardDebt = (stakingPower * rewardTokens[tokenAddress].rewardPerPower) / 1e18;

            IERC20(tokenAddress).safeTransfer(msg.sender, pending);

            emit RewardsClaimed(msg.sender, tokenAddress, pending);
        }
    }

    /**
     * @notice 모든 보상 토큰 청구
     */
    function claimAllRewards() external nonReentrant {
        _updateUserRewards(msg.sender);
        _claimAllRewards(msg.sender);
    }

    // ============ 조회 함수 ============

    /**
     * @notice 사용자의 현재 스테이킹 파워 조회
     * @param user 사용자 주소
     * @return 스테이킹 파워
     */
    function getStakingPower(address user) public view returns (uint256) {
        StakePosition memory position = stakePositions[user];
        if (position.amount == 0) {
            return 0;
        }
        return _calculateStakingPower(position.amount, position.expiry);
    }

    /**
     * @notice 전체 스테이킹 파워 조회
     * @return 전체 스테이킹 파워
     */
    function getTotalStakingPower() public view returns (uint256) {
        // 실제 구현에서는 모든 사용자의 스테이킹 파워를 합산해야 하지만
        // 가스 효율을 위해 캐시된 값을 사용합니다
        return totalStakingPower;
    }

    /**
     * @notice 사용자의 스테이킹 포지션 조회
     * @param user 사용자 주소
     * @return amount 스테이킹 수량
     * @return expiry 만료 시간
     * @return stakingPower 현재 스테이킹 파워
     */
    function getStakePosition(address user)
        external
        view
        returns (uint256 amount, uint256 expiry, uint256 stakingPower)
    {
        StakePosition memory position = stakePositions[user];
        return (position.amount, position.expiry, getStakingPower(user));
    }

    /**
     * @notice 사용자의 대기 중인 보상 조회
     * @param user 사용자 주소
     * @param tokenAddress 보상 토큰 주소
     * @return 대기 중인 보상량
     */
    function getPendingReward(address user, address tokenAddress) external view returns (uint256) {
        if (!rewardTokens[tokenAddress].isActive) {
            return 0;
        }

        UserReward memory userReward = userRewards[user][tokenAddress];
        uint256 stakingPower = getStakingPower(user);

        if (stakingPower == 0) {
            return userReward.pendingReward;
        }

        uint256 accumulatedReward = (stakingPower * rewardTokens[tokenAddress].rewardPerPower) / 1e18;
        uint256 pending = accumulatedReward - userReward.rewardDebt + userReward.pendingReward;

        return pending;
    }

    /**
     * @notice 보상 토큰 목록 조회
     * @return 보상 토큰 주소 배열
     */
    function getRewardTokenList() external view returns (address[] memory) {
        return rewardTokenList;
    }

    // ============ 내부 함수 ============

    /**
     * @notice 스테이킹 파워 계산
     * @param amount 스테이킹 수량
     * @param expiry 만료 시간
     * @return 스테이킹 파워
     */
    function _calculateStakingPower(uint256 amount, uint256 expiry) internal view returns (uint256) {
        if (block.timestamp >= expiry) {
            return 0;
        }

        uint256 remainingTime = expiry - block.timestamp;
        return (amount * remainingTime) / MAX_LOCK_TIME;
    }

    /**
     * @notice 전체 스테이킹 파워 업데이트
     * @param delta 변화량 (증가는 양수, 감소는 음수)
     */
    function _updateTotalStakingPower(int256 delta) internal {
        if (delta > 0) {
            totalStakingPower += uint256(delta);
        } else if (delta < 0) {
            totalStakingPower -= uint256(-delta);
        }
        lastTotalPowerUpdate = block.timestamp;
    }

    /**
     * @notice 사용자 보상 업데이트
     * @param user 사용자 주소
     */
    function _updateUserRewards(address user) internal {
        uint256 stakingPower = getStakingPower(user);

        for (uint256 i = 0; i < rewardTokenList.length; i++) {
            address tokenAddress = rewardTokenList[i];
            if (!rewardTokens[tokenAddress].isActive) {
                continue;
            }

            UserReward storage userReward = userRewards[user][tokenAddress];

            if (stakingPower > 0) {
                uint256 accumulatedReward = (stakingPower * rewardTokens[tokenAddress].rewardPerPower) / 1e18;
                uint256 pending = accumulatedReward - userReward.rewardDebt;
                userReward.pendingReward += pending;
            }
        }
    }

    /**
     * @notice 사용자 보상 부채 업데이트
     * @param user 사용자 주소
     */
    function _updateUserRewardDebt(address user) internal {
        uint256 stakingPower = getStakingPower(user);

        for (uint256 i = 0; i < rewardTokenList.length; i++) {
            address tokenAddress = rewardTokenList[i];
            if (!rewardTokens[tokenAddress].isActive) {
                continue;
            }

            UserReward storage userReward = userRewards[user][tokenAddress];
            userReward.rewardDebt = (stakingPower * rewardTokens[tokenAddress].rewardPerPower) / 1e18;
        }
    }

    /**
     * @notice 모든 보상 청구 (내부 함수)
     * @param user 사용자 주소
     */
    function _claimAllRewards(address user) internal {
        uint256 stakingPower = getStakingPower(user);

        for (uint256 i = 0; i < rewardTokenList.length; i++) {
            address tokenAddress = rewardTokenList[i];
            if (!rewardTokens[tokenAddress].isActive) {
                continue;
            }

            UserReward storage userReward = userRewards[user][tokenAddress];
            uint256 pending = userReward.pendingReward;

            if (pending > 0) {
                userReward.pendingReward = 0;

                // 보상 부채 업데이트
                userReward.rewardDebt = (stakingPower * rewardTokens[tokenAddress].rewardPerPower) / 1e18;

                IERC20(tokenAddress).safeTransfer(user, pending);

                emit RewardsClaimed(user, tokenAddress, pending);
            }
        }
    }

    // ============ 레거시 함수 (하위 호환성) ============

    /**
     * @notice 사용자의 스테이킹 수량 조회 (레거시)
     * @param user 사용자 주소
     * @return 스테이킹 수량
     */
    function getStakedAmount(address user) public view returns (uint256) {
        return stakePositions[user].amount;
    }

    /**
     * @notice 언스테이킹 가능 여부 조회 (레거시)
     * @param user 사용자 주소
     * @return 언스테이킹 가능 여부 (만료되었으면 스테이킹 수량, 아니면 0)
     */
    function getUnstakedAmount(address user) public view returns (uint256) {
        StakePosition memory position = stakePositions[user];
        if (position.amount > 0 && block.timestamp >= position.expiry) {
            return position.amount;
        }
        return 0;
    }
}
