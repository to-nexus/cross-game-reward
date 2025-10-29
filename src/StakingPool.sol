// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {StakingPoolBase} from "./base/StakingPoolBase.sol";
import {IRewardPool} from "./interfaces/IRewardPool.sol";
import {IStakingProtocol} from "./interfaces/IStakingProtocol.sol";
import {PointsLib} from "./libraries/PointsLib.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";

/**
 * @title StakingPoolCode
 * @notice Code contract that returns StakingPool's creation bytecode
 * @dev Part of Code Contract pattern for gas-efficient factory deployments
 */
contract StakingPoolCode {
    /**
     * @notice Returns the creation bytecode of StakingPool
     * @return Creation bytecode for StakingPool contract
     */
    function code() external pure returns (bytes memory) {
        return type(StakingPool).creationCode;
    }
}

/**
 * @title StakingPool
 * @notice Project-specific $CROSS staking pool implementation
 * @dev Inherits from StakingPoolBase and adds:
 *      - Pausable functionality for emergency stops
 *      - Reward claiming logic with RewardPool integration
 *      - Advanced view functions for points and season data
 *      - Batch finalization for gas optimization
 *
 *      Core Responsibilities:
 *      - Manages user stakes and balances
 *      - Calculates points based on stake × time
 *      - Handles automatic season rollovers
 *      - Processes reward claims via RewardPool
 *      - Provides comprehensive view functions
 *
 *      Key Features:
 *      - O(1) total points calculation via aggregation
 *      - Lazy snapshot system for gas efficiency
 *      - Pre-deposit support for season 1
 *      - Virtual season calculations for views
 *      - Support for up to 50 automatic season rollovers
 *
 *      Security:
 *      - Pausable by admin for emergencies
 *      - ReentrancyGuardTransient (EIP-1153)
 *      - Role-based access control
 *      - Safe ERC20 operations
 */
contract StakingPool is StakingPoolBase, Pausable {
    // ============================================
    // Errors
    // ============================================

    /// @notice Thrown when user attempts to claim already-claimed rewards
    error StakingPoolAlreadyClaimed();

    // ============================================
    // Immutable State
    // ============================================

    /// @notice Project ID this pool belongs to
    uint public immutable projectID;

    /// @notice StakingProtocol factory contract
    IStakingProtocol public immutable protocol;

    /// @notice Connected RewardPool for this project
    IRewardPool public rewardPool;

    // ============================================
    // Events
    // ============================================

    /// @notice Emitted when reward pool is set
    event RewardPoolSet(IRewardPool indexed rewardPool);

    /// @notice Emitted when router approval changes
    event RouterApprovalSet(address indexed router, bool approved);

    /// @notice Emitted when points time unit is updated

    /// @notice Emitted when block time is updated

    /// @notice Emitted when next season start is updated
    event NextSeasonStartUpdated(uint oldValue, uint newValue);

    /// @notice Emitted when pool end block is updated
    event PoolEndBlockUpdated(uint oldValue, uint newValue);

    // ============================================
    // Constructor
    // ============================================

    /**
     * @notice Initializes the staking pool for a project
     * @param _projectID Project ID this pool belongs to
     * @param _stakingToken Staking token address (WCROSS)
     * @param _protocol StakingProtocol factory address
     * @param _seasonDuration Number of blocks per season
     * @param _firstSeasonStartTime Block when first season starts
     * @param _poolEndTime Block when pool ends (0 = infinite)
     * @param _preDepositStartTime Block when pre-deposit starts (0 = disabled)
     * @dev Calls StakingPoolBase constructor with all parameters
     */
    constructor(
        uint _projectID,
        IERC20 _stakingToken,
        IStakingProtocol _protocol,
        uint _seasonDuration,
        uint _firstSeasonStartTime,
        uint _poolEndTime,
        uint _preDepositStartTime
    )
        StakingPoolBase(
            _stakingToken,
            address(_protocol),
            _seasonDuration,
            _firstSeasonStartTime,
            _poolEndTime,
            _preDepositStartTime
        )
    {
        projectID = _projectID;
        protocol = _protocol;
    }

    // ============================================
    // Pausable Functions
    // ============================================

    /**
     * @notice Emergency pause (admin only)
     * @dev Prevents all stake/unstake operations when paused
     */
    function pause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
    }

    /**
     * @notice Unpause (admin only)
     * @dev Re-enables stake/unstake operations
     */
    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
    }

    // ============================================
    // Override Functions with Pausable
    // ============================================

    /**
     * @notice Stakes tokens (pausable version)
     * @param amount Amount to stake
     * @dev Overrides base to add whenNotPaused check
     */
    function stake(uint amount) external override nonReentrant whenNotPaused {
        _stakeFor(msg.sender, amount, msg.sender);
    }

    /**
     * @notice Stakes on behalf of user (pausable version)
     * @param user User to stake for
     * @param amount Amount to stake
     * @dev Overrides base to add whenNotPaused check
     */
    function stakeFor(address user, uint amount) external override nonReentrant whenNotPaused {
        require(_isApprovedRouter(msg.sender), CrossStakingBaseNotAuthorized());
        _stakeFor(user, amount, msg.sender);
    }

    /**
     * @notice Withdraws all staked tokens (pausable version)
     * @dev Overrides base to add whenNotPaused check
     */
    function withdrawAll() external override nonReentrant whenNotPaused {
        _withdrawAll(msg.sender, msg.sender);
    }

    function withdrawAllFor(address user) external override nonReentrant whenNotPaused {
        require(_isApprovedRouter(msg.sender), CrossStakingBaseNotAuthorized());
        _withdrawAll(user, msg.sender);
    }

    /**
     * @notice Manually triggers season rollover
     * @dev Validates:
     *      - currentSeason > 0 (at least one season exists)
     *      - Current block > season end block
     *      - Next season start block reached (if set)
     *
     *      Usually not needed as _ensureSeason handles automatic rollovers
     */
    function rolloverSeason() external override {
        require(currentSeason != 0, StakingPoolBaseNoActiveSeason());
        Season storage current = seasons[currentSeason];
        require(block.timestamp > current.endTime, StakingPoolBaseSeasonNotEnded());
        require(nextSeasonStartTime == 0 || block.timestamp >= nextSeasonStartTime, StakingPoolBaseSeasonNotEnded());
        _rolloverSeason();
    }

    /**
     * @notice Claims rewards for a specific season
     * @param seasonNumber Season to claim from
     * @param rewardToken Reward token address
     * @dev Calls internal _claimSeasonFor for msg.sender
     *      Ensures season rollover before claim (auto-finalizes ended seasons)
     */
    function claimSeason(uint seasonNumber, address rewardToken) external nonReentrant {
        _ensureSeason(); // Auto-rollover to finalize ended seasons
        _claimSeasonFor(msg.sender, seasonNumber, rewardToken);
    }

    /**
     * @notice Claims rewards on behalf of a user (router only)
     * @param user User to claim for
     * @param seasonNumber Season to claim from
     * @param rewardToken Reward token address
     * @dev Ensures season rollover before claim (auto-finalizes ended seasons)
     */
    function claimSeasonFor(address user, uint seasonNumber, address rewardToken) external nonReentrant {
        require(_isApprovedRouter(msg.sender), CrossStakingBaseNotAuthorized());
        _ensureSeason(); // Auto-rollover to finalize ended seasons
        _claimSeasonFor(user, seasonNumber, rewardToken);
    }

    /**
     * @notice Updates user's points to current block
     * @param user User address
     * @dev Can be called by RewardPool or admin
     *      Useful for accurate reward previews
     */
    function updatePoints(address user) external {
        require(
            hasRole(REWARD_POOL_ROLE, msg.sender) || hasRole(DEFAULT_ADMIN_ROLE, msg.sender),
            CrossStakingBaseNotAuthorized()
        );
        _updatePoints(user);
    }

    // ============================================
    // Internal Functions
    // ============================================

    /**
     * @notice Internal function to update user's points to current block
     * @param user User address
     * @dev Process:
     *      1. Skip if user has no balance
     *      2. Skip if already updated at current block
     *      3. Ensure all previous seasons are finalized
     *      4. Update current season user data
     *      5. Calculate additional points since last update
     *      6. Reset position points (moved to season data)
     *      7. Emit PointsUpdated event
     */
    function _updatePoints(address user) internal {
        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return;

        uint lastUpdate = position.lastUpdateTime;
        uint currentBlock = block.timestamp;

        if (lastUpdate >= currentBlock) return;

        _ensureUserAllPreviousSeasons(user);

        Season storage current = seasons[currentSeason];

        UserSeasonData storage userData = userSeasonData[currentSeason][user];
        if (userData.balance == 0) {
            uint joinTime = lastUpdate < current.startTime ? current.startTime : lastUpdate;
            userData.balance = position.balance;
            userData.joinTime = joinTime;
            userData.lastPointsTime = currentBlock;
        } else {
            userData.balance = position.balance;
        }

        uint effectiveStart = lastUpdate < current.startTime ? current.startTime : lastUpdate;
        uint newPoints = PointsLib.calculatePoints(position.balance, effectiveStart, currentBlock);

        userData.points += newPoints;
        userData.lastPointsTime = currentBlock;

        position.points = 0;
        position.lastUpdateTime = currentBlock;
        emit PointsUpdated(user, userData.points);
    }

    /**
     * @notice Internal function to claim rewards for a user
     * @param user User to claim for
     * @param seasonNumber Season to claim from
     * @param rewardToken Reward token address
     * @dev Process:
     *      1. Validates season exists and has ended
     *      2. Ensures user's season snapshot exists (lazy finalization)
     *      3. Validates not already claimed
     *      4. Gets user points from snapshot
     *      5. Gets or calculates total points
     *      6. Marks as claimed
     *      7. Calls RewardPool.payUser() to transfer rewards
     *      8. Emits SeasonClaimed event
     *
     *      Season End Validation:
     *      - Accepts if season.isFinalized (rolled over on-chain)
     *      - OR if season ended AND (past season OR pool ended)
     *      This enables claims for last season when pool ends
     *
     *      Lazy snapshot ensures user's points are accurately calculated
     *      even if they didn't interact during the season
     */
    function _claimSeasonFor(address user, uint seasonNumber, address rewardToken) internal {
        Season storage season = seasons[seasonNumber];
        require(seasonNumber > 0 && season.seasonNumber == seasonNumber, StakingPoolBaseSeasonNotEnded());

        // Check if season has ended:
        // 1. Season is finalized (rolled over) OR
        // 2. Season ended AND (it's a past season OR pool ended)
        bool seasonEnded = season.isFinalized
            || (
                block.timestamp > season.endTime
                    && (seasonNumber < currentSeason || (poolEndTime > 0 && block.timestamp >= poolEndTime))
            );

        require(seasonEnded, StakingPoolBaseSeasonNotEnded());

        _ensureUserSeasonSnapshot(user, seasonNumber);

        UserSeasonData storage userData = userSeasonData[seasonNumber][user];
        // Note: claimed check is now in RewardPool.payUser (per-token basis)

        uint userPoints = userData.points;

        uint totalPoints = season.totalPoints;
        if (totalPoints == 0) {
            totalPoints = _calculateSeasonTotalPoints(seasonNumber);
            season.totalPoints = totalPoints;
        }

        if (userPoints > 0 && totalPoints > 0) {
            // RewardPool will check hasClaimedSeasonReward[user][season][token] and revert if already claimed
            rewardPool.payUser(user, seasonNumber, rewardToken, userPoints, totalPoints);
            emit SeasonClaimed(user, seasonNumber, rewardToken, userPoints);
        }
    }

    // ============================================
    // Points Calculation (View)
    // ============================================

    /**
     * @notice Calculates user's points in current season (real-time, no state change)
     * @param user User address
     * @return Current season points
     * @dev Handles multiple scenarios:
     *      - Virtual season (before rollover)
     *      - Current season with UserSeasonData
     *      - Current season with only position data
     *      - Pre-deposit period
     *
     *      Uses lastPointsTime to prevent double-counting
     */
    function _calculateCurrentSeasonPoints(address user) internal view returns (uint) {
        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return 0;

        if (nextSeasonStartTime > 0 && (currentSeason == 0 || seasons[currentSeason].isFinalized)) {
            if (block.timestamp < nextSeasonStartTime) return 0;

            uint posLastUpdate = position.lastUpdateTime;
            uint seasonStart = nextSeasonStartTime;

            uint virtualSeasonNum = currentSeason + 1;
            UserSeasonData storage virtualData = userSeasonData[virtualSeasonNum][user];

            if (posLastUpdate < seasonStart) {
                return virtualData.points + PointsLib.calculatePoints(position.balance, seasonStart, block.timestamp);
            } else {
                return virtualData.points + PointsLib.calculatePoints(position.balance, posLastUpdate, block.timestamp);
            }
        }

        if (currentSeason == 0) return 0;

        Season storage current = seasons[currentSeason];
        UserSeasonData storage userData = userSeasonData[currentSeason][user];

        if (userData.balance > 0) {
            // lastPointsTime부터 추가 포인트 계산 (이중 계산 방지)
            uint effectiveStart = userData.lastPointsTime > 0 ? userData.lastPointsTime : userData.joinTime;
            if (effectiveStart < current.startTime) effectiveStart = current.startTime;

            uint additionalPoints = PointsLib.calculatePoints(userData.balance, effectiveStart, block.timestamp);
            return userData.points + additionalPoints;
        }

        uint lastUpdate = position.lastUpdateTime;
        if (lastUpdate < current.startTime && position.balance > 0) {
            return PointsLib.calculatePoints(position.balance, current.startTime, block.timestamp);
        }

        return 0;
    }

    /**
     * @notice Calculates total points for a season (with forfeited points deduction)
     * @param seasonNum Season number
     * @return Total points minus forfeited points
     * @dev Handles three cases:
     *
     *      1. Finalized season:
     *         - Use cached totalPoints if available
     *         - Otherwise use aggregatedPoints
     *         - Subtract forfeitedPoints
     *
     *      2. Current season (ongoing):
     *         - Base: aggregatedPoints (from last update)
     *         - Add: points since last aggregation update
     *         - Subtract: forfeitedPoints
     *
     *      3. Past season (not finalized yet):
     *         - Calculate up to season.endTime
     *         - Subtract forfeitedPoints
     *
     *      Forfeited points: Points lost when users withdraw during a season
     */
    function _calculateSeasonTotalPoints(uint seasonNum) internal view returns (uint) {
        Season storage season = seasons[seasonNum];

        // Finalized season: return cached totalPoints (forfeited 제외)
        if (season.isFinalized) {
            // totalPoints가 있으면 그것 사용 (forfeited 이미 제외됨)
            if (season.totalPoints > 0) return season.totalPoints;

            // totalPoints가 0이면 aggregatedPoints 사용 (forfeited 제외)
            if (season.aggregatedPoints > 0) {
                return season.aggregatedPoints > season.forfeitedPoints
                    ? season.aggregatedPoints - season.forfeitedPoints
                    : 0;
            }

            // Edge case: finalized but totalPoints/aggregatedPoints = 0
            // This can happen if season was rolled over without proper aggregation
            // Calculate properly using seasonTotalStaked
            if (season.startTime > 0 && season.endTime > 0 && season.seasonTotalStaked > 0) {
                uint calculatedPoints =
                    PointsLib.calculatePoints(season.seasonTotalStaked, season.startTime, season.endTime);
                return calculatedPoints > season.forfeitedPoints ? calculatedPoints - season.forfeitedPoints : 0;
            }

            return 0;
        }

        // 시즌이 생성되었으면 (startTime > 0) 계산
        if (season.startTime > 0) {
            // 실제 현재 시즌 확인 (가상 롤오버 포함)
            (uint actualCurrentSeason,,,) = getCurrentSeasonInfo();

            uint totalAggregated = season.aggregatedPoints;

            // 계산 종료 시간 결정
            uint calculationEndTime;
            if (seasonNum < actualCurrentSeason) {
                // 과거 시즌: endTime 사용
                calculationEndTime = season.endTime;
            } else if (seasonNum == actualCurrentSeason) {
                // 현재 시즌: min(block.timestamp, endTime)
                calculationEndTime = block.timestamp < season.endTime ? block.timestamp : season.endTime;
            } else {
                // 미래 시즌: 아직 시작 안됨
                return 0;
            }

            // lastAggregatedTime부터 calculationEndTime까지 추가 포인트 계산
            if (calculationEndTime > season.lastAggregatedTime) {
                uint additionalPoints =
                    PointsLib.calculatePoints(season.seasonTotalStaked, season.lastAggregatedTime, calculationEndTime);
                totalAggregated += additionalPoints;
            }

            return totalAggregated > season.forfeitedPoints ? totalAggregated - season.forfeitedPoints : 0;
        }

        return 0;
    }

    // ============================================
    // Public View Functions
    // ============================================

    /**
     * @notice Returns user's stake position with real-time points
     * @param user User address
     * @return balance Current staked amount
     * @return points Real-time points in current season
     * @return lastUpdateTime Last block when position was updated
     */
    function getStakePosition(address user) external view returns (uint balance, uint points, uint lastUpdateTime) {
        StakePosition storage position = userStakes[user];
        uint currentPoints = getUserPoints(user);
        return (position.balance, currentPoints, position.lastUpdateTime);
    }

    /**
     * @notice Returns current season information (handles virtual seasons)
     * @return season Current season number (may be virtual if not rolled over)
     * @return startTime Season start block
     * @return endTime Season end block
     * @return timeElapsed Blocks elapsed since season start
     * @dev Virtual season calculation:
     *      - If many seasons passed without rollover, calculates which season we're in
     *      - Based on: (blocks since first season) / seasonDuration
     *      - Allows view functions to work even with pending rollovers
     *      - poolEndTime takes priority over all calculations (checked once at start)
     *
     *      Three cases handled:
     *      1. No season started yet (currentSeason = 0)
     *      2. Current season active
     *      3. Current season ended (calculates virtual next season)
     */
    function getCurrentSeasonInfo() public view returns (uint season, uint startTime, uint endTime, uint timeElapsed) {
        // OPTIMIZATION: Check poolEndTime once at the start
        bool isPoolEnded = poolEndTime > 0 && block.timestamp >= poolEndTime;

        // Early return for pool ended with on-chain season
        if (isPoolEnded && currentSeason > 0) {
            Season storage poolEndedSeason = seasons[currentSeason];
            return (
                currentSeason,
                poolEndedSeason.startTime,
                poolEndedSeason.endTime,
                poolEndedSeason.endTime - poolEndedSeason.startTime + 1
            );
        }

        // Handle virtual season case (currentSeason == 0)
        if (currentSeason == 0 && nextSeasonStartTime > 0) {
            if (block.timestamp >= nextSeasonStartTime) {
                // Calculate which virtual season we're in
                uint timeSinceStart = block.timestamp - nextSeasonStartTime;
                uint seasonIndex = timeSinceStart / seasonDuration; // 0-based index

                season = seasonIndex + 1; // Convert to 1-based season number
                startTime = nextSeasonStartTime + (seasonIndex * seasonDuration);
                endTime = _calculateEndTime(startTime);

                // Pool ended: clamp timeElapsed to season length
                if (isPoolEnded) timeElapsed = endTime >= startTime ? endTime - startTime + 1 : 0;
                else timeElapsed = block.timestamp - startTime;
            } else {
                // Season not started yet
                return (0, 0, 0, 0);
            }
            return (season, startTime, endTime, timeElapsed);
        }

        // No season at all
        if (currentSeason == 0) return (0, 0, 0, 0);

        // Case: On-chain season exists - calculate actual current season
        Season storage current = seasons[currentSeason];

        // Check if we're still in the current season
        if (block.timestamp <= current.endTime) {
            // Still in current season
            season = currentSeason;
            startTime = current.startTime;
            endTime = current.endTime;
            timeElapsed = block.timestamp - startTime;
        } else {
            // Current season ended, calculate which season we're actually in
            // Each season starts at: current.startTime + (n * seasonDuration)
            // where n = 0, 1, 2, 3...

            uint timeSinceFirstSeason = block.timestamp - current.startTime;
            uint seasonIndex = timeSinceFirstSeason / seasonDuration;

            // Calculate the actual season number
            season = currentSeason + seasonIndex;

            // Calculate start and end blocks for this season
            startTime = current.startTime + (seasonIndex * seasonDuration);

            // Use _calculateEndTime to respect poolEndTime
            endTime = _calculateEndTime(startTime);

            timeElapsed = block.timestamp - startTime;
        }
    }

    /**
     * @notice Returns user's points for a specific season
     * @param seasonNumber Season number
     * @param user User address
     * @return userPoints User's points in the season
     * @return totalPoints Total points in the season
     * @dev Prioritizes finalized data, falls back to real-time calculation
     *      Handles both on-chain and virtual seasons
     */
    function getSeasonUserPoints(uint seasonNumber, address user)
        public
        view
        returns (uint userPoints, uint totalPoints)
    {
        UserSeasonData storage userData = userSeasonData[seasonNumber][user];

        // Return finalized data if available
        if (userData.finalized) return (userData.points, seasonTotalPointsSnapshot(seasonNumber));

        // 과거 시즌: finalize되지 않았어도 계산해서 반환
        Season storage season = seasons[seasonNumber];

        // 현재 시즌이면서 아직 종료되지 않은 경우만 실시간 포인트 반환
        if (seasonNumber == currentSeason && currentSeason > 0 && block.timestamp <= season.endTime) {
            return (getUserPoints(user), seasonTotalPointsSnapshot(seasonNumber));
        }

        // 시즌이 존재하지 않으면 가상 시즌 계산
        if (season.startTime == 0) {
            // 가상 시즌의 endTime 계산
            uint virtualStartTime = nextSeasonStartTime + ((seasonNumber - 1) * seasonDuration);
            uint virtualEndTime = _calculateEndTime(virtualStartTime);

            // 현재 블록이 가상 시즌 범위 내인지 확인
            if (block.timestamp < virtualStartTime) return (0, 0);

            StakePosition storage pos = userStakes[user];
            if (pos.balance == 0) {
                userPoints = 0;
            } else {
                uint posLastUpdate = pos.lastUpdateTime;
                if (posLastUpdate <= virtualEndTime) {
                    uint fromTime = posLastUpdate > virtualStartTime ? posLastUpdate : virtualStartTime;
                    // 과거 가상 시즌: virtualEndTime 사용, 현재/미래 가상 시즌: min(block.timestamp, virtualEndTime) 사용
                    // 중요: currentSeason이 아닌 actualCurrentSeason과 비교해야 함 (가상 시즌 고려)
                    (uint actualCurrentSeason,,,) = getCurrentSeasonInfo();
                    uint toTime = (seasonNumber < actualCurrentSeason || block.timestamp >= virtualEndTime)
                        ? virtualEndTime
                        : block.timestamp;

                    userPoints = PointsLib.calculatePoints(pos.balance, fromTime, toTime);
                }
            }
            return (userPoints, seasonTotalPointsSnapshot(seasonNumber));
        }

        // 아직 시작하지 않은 경우
        if (block.timestamp < season.startTime) return (0, 0);

        // 과거 시즌에서 해당 시즌의 잔액이 있으면 그것으로 계산
        if (userData.balance > 0) {
            uint basePoints = userData.points;
            uint effectiveStart = userData.joinTime > season.startTime ? userData.joinTime : season.startTime;

            // 전체 기간 계산
            uint calculatedPoints = PointsLib.calculatePoints(userData.balance, effectiveStart, season.endTime);

            // basePoints는 중간 스냅샷일 수 있으므로, 더 큰 값 사용
            userPoints = calculatedPoints > basePoints ? calculatedPoints : basePoints;
            return (userPoints, seasonTotalPointsSnapshot(seasonNumber));
        }

        // userData가 없으면 position으로 계산 (자동 참여 케이스)
        StakePosition storage position = userStakes[user];
        if (position.balance == 0) {
            userPoints = 0;
        } else {
            uint lastUpdate = position.lastUpdateTime;

            // 시즌 종료 후 스테이킹한 경우
            if (lastUpdate > season.endTime) {
                userPoints = 0;
            } else {
                // 과거 시즌인지 확인 (롤오버 되었지만 finalize 안된 경우)
                // 과거 시즌이면 endTime 사용, 현재 시즌이면 현재 블록 사용
                // 중요: currentSeason이 아닌 actualCurrentSeason과 비교해야 함 (가상 시즌 고려)
                (uint actualCurrentSeason,,,) = getCurrentSeasonInfo();
                uint calculationEndBlock = (seasonNumber < actualCurrentSeason) ? season.endTime : block.timestamp;
                if (calculationEndBlock > season.endTime) calculationEndBlock = season.endTime;

                // 시즌 시작 전에 스테이킹한 경우: 시즌 전체 기간 계산
                if (lastUpdate < season.startTime) {
                    userPoints = PointsLib.calculatePoints(position.balance, season.startTime, calculationEndBlock);
                } else {
                    // 시즌 중간에 스테이킹한 경우: lastUpdate부터 계산
                    userPoints = PointsLib.calculatePoints(position.balance, lastUpdate, calculationEndBlock);
                }
            }
        }

        return (userPoints, seasonTotalPointsSnapshot(seasonNumber));
    }

    /**
     * @notice Returns finalized total points for a season
     * @param seasonNumber Season number
     * @return Total points snapshot
     * @dev - For finalized seasons: returns cached totalPoints
     *      - For current season: calculates real-time total
     *      - For virtual seasons: calculates estimated total
     *
     *      This is the authoritative source for reward distribution
     */
    function seasonTotalPointsSnapshot(uint seasonNumber) public view returns (uint) {
        Season storage season = seasons[seasonNumber];

        if (season.isFinalized) {
            if (season.totalPoints == 0) return _calculateSeasonTotalPoints(seasonNumber);
            return season.totalPoints;
        }

        // 시즌이 생성되었으면 해당 데이터로 계산 (현재 또는 과거 시즌)
        if (season.startTime > 0) return _calculateSeasonTotalPoints(seasonNumber);

        // 가상 시즌 계산 (시즌이 생성 안됨)
        uint virtualStartTime = nextSeasonStartTime + ((seasonNumber - 1) * seasonDuration);
        uint virtualEndTime = _calculateEndTime(virtualStartTime);

        if (totalStaked == 0) return 0;

        // 과거 가상 시즌: virtualEndTime 사용, 현재/미래 가상 시즌: min(block.timestamp, virtualEndTime) 사용
        uint toTime =
            (seasonNumber < currentSeason || block.timestamp >= virtualEndTime) ? virtualEndTime : block.timestamp;

        return PointsLib.calculatePoints(totalStaked, virtualStartTime, toTime);
    }

    /**
     * @notice Returns expected reward for a user in a past season
     * @param seasonNumber Season number (must be ended)
     * @param user User address
     * @param rewardToken Reward token address
     * @return Expected reward amount
     * @dev Returns 0 if:
     *      - No reward pool connected
     *      - Season is current or future (not ended)
     *      - User has no points
     */
    function getExpectedSeasonReward(uint seasonNumber, address user, address rewardToken)
        external
        view
        returns (uint)
    {
        if (address(rewardPool) == address(0)) return 0;

        // Current season not ended yet
        (uint actualSeason,,,) = getCurrentSeasonInfo();
        if (seasonNumber >= actualSeason) return 0;

        return rewardPool.getExpectedReward(user, seasonNumber, rewardToken);
    }

    /**
     * @notice Returns user's current season points (real-time)
     * @param user User address
     * @return Current season points
     */
    function getUserPoints(address user) public view returns (uint) {
        return _calculateCurrentSeasonPoints(user);
    }

    /**
     * @notice Returns detailed user data for a specific season
     * @param seasonNumber Season number
     * @param user User address
     * @return points User's points
     * @return balance User's balance
     * @return joinTime Block when user joined
     * @return finalized Whether snapshot finalized
     * @dev claimed status is tracked per-token in RewardPool.hasClaimedSeasonReward
     */
    function getUserSeasonData(uint seasonNumber, address user)
        external
        view
        returns (uint points, uint balance, uint joinTime, bool finalized)
    {
        UserSeasonData storage data = userSeasonData[seasonNumber][user];
        return (data.points, data.balance, data.joinTime, data.finalized);
    }

    /**
     * @notice Previews claim information for a user
     * @param seasonNumber Season number
     * @param user User address
     * @param rewardToken Reward token address
     * @return userPoints User's points
     * @return totalPoints Total points
     * @return expectedReward Expected reward amount
     * @return alreadyClaimed Whether already claimed
     * @return canClaim Whether claiming is possible
     * @dev Comprehensive preview for frontend display before claim transaction
     */
    function previewClaim(uint seasonNumber, address user, address rewardToken)
        external
        view
        returns (uint userPoints, uint totalPoints, uint expectedReward, bool alreadyClaimed, bool canClaim)
    {
        if (address(rewardPool) == address(0)) return (0, 0, 0, false, false);

        Season storage season = seasons[seasonNumber];

        // 온체인 시즌 (rollover 완료)
        if (season.startTime > 0) {
            canClaim = season.isFinalized;

            // Check if already claimed from RewardPool (per-token)
            alreadyClaimed = rewardPool.hasClaimedSeasonReward(user, seasonNumber, rewardToken);

            UserSeasonData storage userData = userSeasonData[seasonNumber][user];
            if (userData.finalized) {
                userPoints = userData.points;
                totalPoints = season.totalPoints > 0 ? season.totalPoints : seasonTotalPointsSnapshot(seasonNumber);
            } else {
                (userPoints, totalPoints) = getSeasonUserPoints(seasonNumber, user);
            }

            if (userPoints > 0 && totalPoints > 0 && !alreadyClaimed) {
                expectedReward = rewardPool.getExpectedReward(user, seasonNumber, rewardToken);
            }
            return (userPoints, totalPoints, expectedReward, alreadyClaimed, canClaim);
        }

        // 가상 시즌 (rollover 안됨)
        (uint actualSeason,,,) = getCurrentSeasonInfo();
        if (seasonNumber > actualSeason) return (0, 0, 0, false, false);

        // 가상 시즌 포인트 계산
        (userPoints, totalPoints) = getSeasonUserPoints(seasonNumber, user);

        canClaim = false; // 가상 시즌은 claim 불가
        alreadyClaimed = false;

        if (userPoints > 0 && totalPoints > 0) {
            expectedReward = rewardPool.getExpectedReward(user, seasonNumber, rewardToken);
        }
    }

    /**
     * @notice Returns paginated list of stakers
     * @param offset Starting index
     * @param limit Maximum number to return
     * @return stakerList Array of staker addresses
     * @return total Total number of stakers
     * @dev Note: Array includes users who have unstaked (balance = 0)
     *      Filter by getStakePosition(user).balance > 0 for active stakers
     */
    function getStakers(uint offset, uint limit) external view returns (address[] memory stakerList, uint total) {
        total = stakers.length;
        if (offset >= total) return (new address[](0), total);

        uint end = offset + limit;
        if (end > total) end = total;

        uint count = end - offset;
        stakerList = new address[](count);

        for (uint i = 0; i < count; i++) {
            stakerList[i] = stakers[offset + i];
        }
    }

    /**
     * @notice Batch finalizes user's previous seasons (gas optimization)
     * @param user User address
     * @param maxSeasons Maximum seasons to process in one call
     * @return processed Number of seasons actually processed
     * @dev Useful for users who haven't interacted for many seasons
     *      Allows pre-processing to reduce gas on subsequent operations
     *
     *      Example:
     *      User staked in season 1, didn't interact until season 20
     *      Call finalizeUserSeasonsBatch(user, 10) twice to process all 19 seasons
     */
    function finalizeUserSeasonsBatch(address user, uint maxSeasons) external returns (uint processed) {
        if (currentSeason == 0) return 0;

        uint startSeason = lastFinalizedSeason[user];
        if (startSeason == 0) startSeason = 1;

        uint endSeason = currentSeason > 0 ? currentSeason - 1 : 0;
        if (startSeason > endSeason) return 0;

        processed = 0;
        uint lastProcessed = startSeason - 1;

        for (uint i = startSeason; i <= endSeason && processed < maxSeasons;) {
            _ensureUserSeasonSnapshot(user, i);
            lastProcessed = i;
            unchecked {
                ++i;
                ++processed;
            }
        }

        // 루프 끝에서 한 번만 업데이트 (가스 절약)
        if (processed > 0) lastFinalizedSeason[user] = lastProcessed;

        return processed;
    }

    /**
     * @notice Checks if a router is approved (override to add global check)
     * @param router Router address
     * @return True if router is globally approved or has ROUTER_ROLE
     */
    function _isApprovedRouter(address router) internal view override returns (bool) {
        return protocol.isGlobalApprovedRouter(router) || hasRole(ROUTER_ROLE, router);
    }

    // ============================================
    // Admin Configuration Functions
    // ============================================

    /**
     * @notice Sets the reward pool (one-time only)
     * @param _rewardPool RewardPool address
     * @dev Can only be set once (reverts if already set)
     *      Grants REWARD_POOL_ROLE to the reward pool
     */
    function setRewardPool(IRewardPool _rewardPool) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(address(rewardPool) == address(0), CrossStakingBaseAlreadySet());
        _validateAddress(address(_rewardPool));
        rewardPool = _rewardPool;
        _grantRole(REWARD_POOL_ROLE, address(_rewardPool));
        emit RewardPoolSet(_rewardPool);
    }

    /**
     * @notice Approves/revokes a router
     * @param router Router address
     * @param approved Approval status
     * @dev Grants or revokes ROUTER_ROLE
     */
    function setApprovedRouter(address router, bool approved) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (approved) _grantRole(ROUTER_ROLE, router);
        else _revokeRole(ROUTER_ROLE, router);
        emit RouterApprovalSet(router, approved);
    }

    /**
     * @notice Sets when the next season should start
     * @param _startTime Start block (0 = use automatic calculation)
     * @dev Used to create gaps between seasons or delay next season
     */
    function setNextSeasonStart(uint _startTime) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_startTime == 0 || _startTime > block.timestamp, StakingPoolBaseInvalidStartTime());
        uint oldValue = nextSeasonStartTime;
        nextSeasonStartTime = _startTime;
        emit NextSeasonStartUpdated(oldValue, _startTime);
    }

    /**
     * @notice Sets when the pool should end
     * @param _endTime End block (0 = infinite)
     * @dev After end block, no new seasons are created
     *      Existing stakes can still be withdrawn
     */
    function setPoolEndBlock(uint _endTime) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_endTime == 0 || _endTime > block.timestamp, StakingPoolBaseInvalidEndTime());
        uint oldValue = poolEndTime;
        poolEndTime = _endTime;
        emit PoolEndBlockUpdated(oldValue, _endTime);
    }
}
