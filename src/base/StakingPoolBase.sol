// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IStakingPool} from "../interfaces/IStakingPool.sol";
import {PointsLib} from "../libraries/PointsLib.sol";
import {SeasonLib} from "../libraries/SeasonLib.sol";
import {CrossStakingBase} from "./CrossStakingBase.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title StakingPoolBase
 * @notice Base abstract contract for project-specific staking pools
 * @dev Implements core staking pool functionality with advanced features:
 *
 *      Season System:
 *      - Time-based seasons with automatic rollovers (up to 50 at once)
 *      - Lazy snapshot system for gas-efficient user data finalization
 *      - Virtual season support for view functions
 *
 *      Points System:
 *      - O(1) aggregation for total points calculation
 *      - Real-time points calculation without state changes
 *      - Points = balance × time × PRECISION / timeUnit
 *
 *      Pre-deposit Feature:
 *      - Allows staking before season 1 starts
 *      - Points accumulate from season start time
 *      - Optional feature (disabled if preDepositStartTime = 0)
 *
 *      Gas Optimizations:
 *      - Incremental aggregation updates (O(1) complexity)
 *      - Lazy user season snapshots (only when needed)
 *      - Unchecked arithmetic where safe
 *      - Transient storage for reentrancy guard (EIP-1153)
 */
abstract contract StakingPoolBase is IStakingPool, CrossStakingBase {
    using SafeERC20 for IERC20;
    using PointsLib for *;
    using SeasonLib for *;

    // ============ Constants ============

    /// @notice Role for RewardPool contract (can call updatePoints)
    bytes32 public constant REWARD_POOL_ROLE = keccak256("REWARD_POOL_ROLE");

    /// @notice Role for pool managers (can manually rollover seasons)
    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");

    /// @notice Role for routers (can stake/withdraw on behalf of users)
    bytes32 public constant ROUTER_ROLE = keccak256("ROUTER_ROLE");

    /// @notice Minimum stake amount (1 CROSS token)
    uint public constant MIN_STAKE = 1e18;

    /// @notice Maximum seasons that can be rolled over in one transaction
    /// @dev Prevents excessive gas consumption. 50 seasons ≈ 6M gas
    uint public constant MAX_AUTO_ROLLOVERS = 50;

    // ============ Errors ============

    /// @notice Thrown when stake amount is below MIN_STAKE
    error StakingPoolBaseBelowMinStake();

    /// @notice Thrown when user has no staking position
    error StakingPoolBaseNoPosition();

    /// @notice Thrown when season duration parameter is invalid (zero)
    error StakingPoolBaseInvalidSeasonDuration();

    /// @notice Thrown when time unit parameter is invalid (zero)
    error StakingPoolBaseInvalidTimeUnit();

    /// @notice Thrown when attempting action before season has ended
    error StakingPoolBaseSeasonNotEnded();

    /// @notice Thrown when no active season exists
    error StakingPoolBaseNoActiveSeason();

    /// @notice Thrown when start time parameter is invalid
    error StakingPoolBaseInvalidStartTime();

    /// @notice Thrown when end time parameter is invalid
    error StakingPoolBaseInvalidEndTime();

    /// @notice Thrown when too many seasons need to be rolled over (>50)
    error StakingPoolBaseTooManySeasons();

    /// @notice Thrown when external caller tries to call self-only function
    error StakingPoolBaseOnlySelf();

    /// @notice Thrown when manual rollover parameter is invalid
    error StakingPoolBaseInvalidMaxRollovers();

    /// @notice Thrown when pre-deposit block is after first season start
    error StakingPoolBaseInvalidPreDepositTime();

    // ============ Structs ============

    /**
     * @notice User's current staking position (persists across seasons)
     * @param balance Current staked amount
     * @param points Temporary points (moved to UserSeasonData on updates)
     * @param lastUpdateTime Last timestamp when position was updated
     */
    struct StakePosition {
        uint balance;
        uint points;
        uint lastUpdateTime;
    }

    /**
     * @notice Season data structure
     * @param seasonNumber Season identifier (1-indexed)
     * @param startTime First timestamp of the season
     * @param endTime Last timestamp of the season (inclusive)
     * @param isFinalized Whether season has ended and been finalized
     * @param totalPoints Finalized total points (cached, immutable after finalization)
     * @param seasonTotalStaked Current total staked in season (for aggregation)
     * @param lastAggregatedTime Last timestamp when aggregation was updated
     * @param aggregatedPoints Accumulated aggregated points (real-time during season)
     * @param forfeitedPoints Points forfeited from withdrawals during season
     */
    struct Season {
        uint seasonNumber;
        uint startTime;
        uint endTime;
        bool isFinalized;
        uint totalPoints;
        uint seasonTotalStaked;
        uint lastAggregatedTime;
        uint aggregatedPoints;
        uint forfeitedPoints;
    }

    /**
     * @notice User's data for a specific season (lazy snapshot)
     * @param points User's points in this season
     * @param balance User's balance in this season
     * @param joinTime Timestamp when user joined this season
     * @param lastPointsTime Last timestamp when points were calculated
     * @param finalized Whether snapshot has been taken (lazy finalization)
     * @dev claimed status is tracked per-token in RewardPool.hasClaimedSeasonReward
     */
    struct UserSeasonData {
        uint points;
        uint balance;
        uint joinTime;
        uint lastPointsTime;
        bool finalized;
    }

    // ============ Immutable State ============

    /// @notice The token being staked (WCROSS)
    IERC20 public immutable stakingToken;

    // ============ User Data ============

    /// @notice Maps user address to their stake position
    mapping(address => StakePosition) public userStakes;

    /// @notice Total amount of tokens currently staked
    uint public totalStaked;

    /// @notice Array of all stakers (never removed, even if balance becomes 0)
    address[] public stakers;

    /// @notice Tracks if an address has ever staked (optimization)
    mapping(address => bool) public isStaker;

    // ============ Season Data ============

    /// @notice Current active season number (0 = not started)
    uint public currentSeason;

    /// @notice Maps season number to Season data
    mapping(uint => Season) public seasons;

    /// @notice Maps season number => user address => UserSeasonData
    mapping(uint => mapping(address => UserSeasonData)) public userSeasonData;

    /// @notice Tracks last finalized season for each user (lazy snapshot optimization)
    mapping(address => uint) public lastFinalizedSeason;

    // ============ Configuration ============

    /// @notice Number of blocks per season
    uint public seasonDuration;

    /// @notice Block number when pool ends (0 = infinite)
    uint public poolEndTime;

    /// @notice Block number when next season starts (manual override)
    uint public nextSeasonStartTime;

    /// @notice Block number when pre-deposit starts (0 = disabled)
    /// @dev Pre-deposit allows staking before season 1 starts (season 1 only)
    uint public preDepositStartTime;

    /// @notice Time unit for points calculation in seconds (default: 1 hour)

    /// @notice Block time in seconds (default: 1 second per block)

    // ============ Events ============

    /// @notice Emitted when a user stakes tokens
    event Staked(address indexed user, uint amount, uint newBalance);

    /// @notice Emitted when a user withdraws all tokens
    event WithdrawnAll(address indexed user, uint amount);

    /// @notice Emitted when user's points are updated
    event PointsUpdated(address indexed user, uint points);

    /// @notice Emitted when a season is rolled over to the next
    event SeasonRolledOver(uint indexed oldSeason, uint indexed newSeason, uint totalPoints);

    /// @notice Emitted when a user claims rewards for a season
    event SeasonClaimed(address indexed user, uint indexed season, address indexed rewardToken, uint points);

    /// @notice Emitted when manual rollover is completed
    event ManualRolloverCompleted(uint rolloversPerformed, uint fromSeason, uint toSeason);

    /// @notice Emitted when points are forfeited due to withdrawal
    event PointsForfeited(address indexed user, uint indexed season, uint amount);

    /// @notice Emitted when season aggregation is updated
    event SeasonAggregationUpdated(uint indexed season, uint aggregatedPoints, uint lastAggregatedTime);

    // ============ Constructor ============

    /**
     * @notice Initializes the staking pool base contract
     * @param _stakingToken Token to be staked (WCROSS)
     * @param admin Initial admin address
     * @param _seasonDuration Number of blocks per season
     * @param _firstSeasonStartTime Block when first season starts
     * @param _poolEndTime Block when pool ends (0 = infinite)
     * @param _preDepositStartTime Block when pre-deposit starts (0 = disabled)
     * @dev Validates:
     *      - All addresses are non-zero
     *      - Season blocks > 0
     *      - First season start time > 0
     *      - Pool end block > first season end (if not 0)
     *      - Pre-deposit block < first season start (if not 0)
     */
    constructor(
        IERC20 _stakingToken,
        address admin,
        uint _seasonDuration,
        uint _firstSeasonStartTime,
        uint _poolEndTime,
        uint _preDepositStartTime
    ) CrossStakingBase(admin) {
        _validateAddress(address(_stakingToken));
        require(_seasonDuration != 0, StakingPoolBaseInvalidSeasonDuration());
        require(_firstSeasonStartTime != 0, StakingPoolBaseInvalidStartTime());
        require(
            _poolEndTime == 0 || _poolEndTime > _firstSeasonStartTime + _seasonDuration, StakingPoolBaseInvalidEndTime()
        );
        if (_preDepositStartTime > 0 && _preDepositStartTime > _firstSeasonStartTime) {
            revert StakingPoolBaseInvalidPreDepositTime();
        }

        stakingToken = _stakingToken;
        seasonDuration = _seasonDuration;
        nextSeasonStartTime = _firstSeasonStartTime;
        poolEndTime = _poolEndTime;
        preDepositStartTime = _preDepositStartTime;

        _grantRole(MANAGER_ROLE, admin);
    }

    // ============ External Staking Functions ============

    /**
     * @notice Stakes tokens for the caller
     * @param amount Amount of tokens to stake
     * @dev Calls internal _stakeFor function
     */
    function stake(uint amount) external virtual nonReentrant {
        _stakeFor(msg.sender, amount, msg.sender);
    }

    /**
     * @notice Stakes tokens on behalf of a user (router only)
     * @param user User to stake for
     * @param amount Amount of tokens to stake
     * @dev Only approved routers (ROUTER_ROLE) can call this
     */
    function stakeFor(address user, uint amount) external virtual nonReentrant {
        require(_isApprovedRouter(msg.sender), CrossStakingBaseNotAuthorized());
        _stakeFor(user, amount, msg.sender);
    }

    /**
     * @notice Withdraws all staked tokens for the caller
     * @dev - Forfeits current season points
     *      - Previous season rewards can still be claimed
     */
    function withdrawAll() external virtual nonReentrant {
        _withdrawAll(msg.sender, msg.sender);
    }

    /**
     * @notice Withdraws all staked tokens on behalf of a user (router only)
     * @param user User to withdraw for
     * @dev Only approved routers (ROUTER_ROLE) can call this
     */
    function withdrawAllFor(address user) external virtual nonReentrant {
        require(_isApprovedRouter(msg.sender), CrossStakingBaseNotAuthorized());
        _withdrawAll(user, msg.sender);
    }

    // ============ Internal Core Functions ============

    /**
     * @notice Internal function to stake tokens for a user
     * @param user User to stake for
     * @param amount Amount to stake
     * @param from Address to transfer tokens from (msg.sender for direct, router for proxy)
     * @dev Process flow:
     *      1. Validates pool is not ended
     *      2. Ensures season is active (auto-rollover if needed)
     *      3. Validates season is active or pre-deposit period
     *      4. Finalizes all previous seasons for user (lazy snapshot)
     *      5. Calculates and updates points for current balance
     *      6. Validates minimum stake requirement
     *      7. Transfers tokens using SafeERC20
     *      8. Updates user position and season data
     *      9. Updates season aggregation (O(1) total points)
     *      10. Adds user to stakers array if first time
     *      11. Emits Staked event
     *
     *      Pre-deposit handling:
     *      - If currentSeason = 0 and in pre-deposit period, allows staking
     *      - Points will accumulate from season 1 start block
     */
    function _stakeFor(address user, uint amount, address from) internal virtual {
        require(poolEndTime == 0 || block.timestamp < poolEndTime, StakingPoolBaseNoActiveSeason());

        _ensureSeason();

        if (currentSeason == 0) {
            if (preDepositStartTime > 0 && block.timestamp >= preDepositStartTime) {} else {
                require(block.timestamp >= nextSeasonStartTime, StakingPoolBaseNoActiveSeason());
            }
        } else {
            require(isSeasonActive(), StakingPoolBaseNoActiveSeason());
        }

        _ensureUserAllPreviousSeasons(user);

        StakePosition storage position = userStakes[user];
        uint oldBalance = position.balance;

        if (oldBalance > 0) {
            uint effectiveStart = position.lastUpdateTime < seasons[currentSeason].startTime
                ? seasons[currentSeason].startTime
                : position.lastUpdateTime;
            uint additionalPoints = PointsLib.calculatePoints(position.balance, effectiveStart, block.timestamp);
            UserSeasonData storage currentUserData = userSeasonData[currentSeason][user];
            currentUserData.points += additionalPoints;
            currentUserData.lastPointsTime = block.timestamp;
            position.points = 0;
        }

        uint newBalance = oldBalance + amount;
        require(newBalance >= MIN_STAKE, StakingPoolBaseBelowMinStake());

        stakingToken.safeTransferFrom(from, address(this), amount);

        position.balance = newBalance;
        position.lastUpdateTime = block.timestamp;
        totalStaked += amount;

        Season storage currentSeasonData = seasons[currentSeason];
        _updateSeasonAggregation(currentSeason);
        currentSeasonData.seasonTotalStaked += amount;

        UserSeasonData storage seasonData = userSeasonData[currentSeason][user];
        if (seasonData.balance == 0) {
            Season storage current = seasons[currentSeason];
            seasonData.balance = newBalance;
            seasonData.joinTime = block.timestamp < current.startTime ? current.startTime : block.timestamp;
            seasonData.lastPointsTime = block.timestamp;
        } else {
            seasonData.balance = newBalance;
            seasonData.lastPointsTime = block.timestamp;
        }

        if (!isStaker[user]) {
            stakers.push(user);
            isStaker[user] = true;
        }

        emit Staked(user, amount, newBalance);
    }

    /**
     * @notice Internal function to withdraw all staked tokens
     * @param user User to withdraw for
     * @param recipient Address to receive the withdrawn tokens
     * @dev Process flow:
     *      1. Ensures current season is active (auto-rollover)
     *      2. Finalizes all previous seasons (lazy snapshot)
     *      3. Validates user has a position
     *      4. Calculates forfeited points in current season
     *      5. Resets user's current season data (points = 0, balance = 0)
     *      6. Updates season aggregation and forfeited points
     *      7. Resets user's overall position
     *      8. Updates totalStaked
     *      9. Transfers tokens to recipient
     *      10. Emits WithdrawnAll and PointsForfeited events
     *
     *      Important: Current season points are LOST on withdrawal
     *      Previous season rewards can still be claimed after withdrawal
     */
    function _withdrawAll(address user, address recipient) internal virtual {
        StakePosition storage position = userStakes[user];
        require(position.balance != 0, StakingPoolBaseNoPosition());

        uint amount = position.balance;

        // 롤오버 전에 현재 시즌의 forfeitedPoints 먼저 계산하고 추가
        // 단, 시즌이 아직 active한 경우에만 (종료 후 unstake는 포인트 유지)
        if (currentSeason > 0) {
            Season storage currentSeasonData = seasons[currentSeason];
            UserSeasonData storage seasonData = userSeasonData[currentSeason][user];

            _updateSeasonAggregation(currentSeason);

            // 시즌이 아직 진행 중인 경우에만 포인트 몰수 (종료 후에는 유지)
            bool seasonStillActive = block.timestamp <= currentSeasonData.endTime;

            uint userForfeitedPoints = 0;
            if (seasonStillActive) {
                if (seasonData.balance > 0) {
                    userForfeitedPoints = seasonData.points;

                    uint effectiveStart =
                        seasonData.lastPointsTime > 0 ? seasonData.lastPointsTime : seasonData.joinTime;
                    if (effectiveStart < currentSeasonData.startTime) effectiveStart = currentSeasonData.startTime;

                    // Cap at season endTime to prevent calculating points beyond season end
                    uint targetTime = block.timestamp;
                    if (targetTime > currentSeasonData.endTime) targetTime = currentSeasonData.endTime;

                    uint additionalPoints = PointsLib.calculatePoints(seasonData.balance, effectiveStart, targetTime);
                    userForfeitedPoints += additionalPoints;
                } else if (position.balance > 0) {
                    uint lastUpdate = position.lastUpdateTime;
                    uint effectiveStart =
                        lastUpdate < currentSeasonData.startTime ? currentSeasonData.startTime : lastUpdate;

                    // Cap at season endTime to prevent calculating points beyond season end
                    uint targetTime = block.timestamp;
                    if (targetTime > currentSeasonData.endTime) targetTime = currentSeasonData.endTime;

                    userForfeitedPoints = PointsLib.calculatePoints(position.balance, effectiveStart, targetTime);
                }
            }

            seasonData.points = 0;
            seasonData.balance = 0;

            if (userForfeitedPoints > 0) {
                currentSeasonData.forfeitedPoints += userForfeitedPoints;
                emit PointsForfeited(user, currentSeason, userForfeitedPoints);
            }

            currentSeasonData.seasonTotalStaked -= amount;
        }

        // forfeitedPoints 추가 후 롤오버 처리
        _ensureSeason();
        _ensureUserAllPreviousSeasons(user);

        position.balance = 0;
        position.points = 0;
        position.lastUpdateTime = block.timestamp;
        totalStaked -= amount;

        stakingToken.safeTransfer(recipient, amount);

        emit WithdrawnAll(user, amount);
    }

    // ============ Season Management (Internal) ============

    /**
     * @notice Ensures a valid season exists, performing auto-rollovers if needed
     * @dev Called at the start of stake/withdraw operations to maintain season state
     *
     *      Behavior:
     *      1. If pool has ended, returns immediately
     *      2. If currentSeason = 0 and start block reached, starts first season
     *      3. While current season has ended, rolls over to next (max 50 times)
     *
     *      Auto-rollover mechanism:
     *      - Can rollover up to MAX_AUTO_ROLLOVERS (50) seasons in one call
     *      - Each rollover: ~120k gas
     *      - Total for 50: ~6M gas (well within block limit)
     *
     *      Edge cases handled:
     *      - Pool end block reached -> no more seasons created
     *      - Next season start time set -> waits for that block
     *      - Multiple seasons passed -> all rolled over automatically
     *      - >50 seasons pending -> reverts (use manualRolloverSeasons)
     */
    function _ensureSeason() internal virtual {
        if (poolEndTime > 0 && (nextSeasonStartTime == 0 || nextSeasonStartTime <= poolEndTime)) {
            if (block.timestamp >= poolEndTime) return;
        }

        if (currentSeason == 0) {
            if (block.timestamp >= nextSeasonStartTime) _startFirstSeason();
            else return;
        }

        uint maxRollovers = MAX_AUTO_ROLLOVERS;
        uint rolloversPerformed = 0;

        while (currentSeason > 0 && rolloversPerformed < maxRollovers) {
            Season storage current = seasons[currentSeason];

            if (block.timestamp <= current.endTime) break;
            if (nextSeasonStartTime > 0 && block.timestamp < nextSeasonStartTime) break;

            _rolloverSeason();
            unchecked {
                ++rolloversPerformed;
            }
        }

        require(rolloversPerformed < maxRollovers, StakingPoolBaseTooManySeasons());
    }

    /**
     * @notice Starts the first season (season 1)
     * @dev - Only called when currentSeason = 0 and start block reached
     *      - Sets currentSeason = 1
     *      - Creates season 1 with configured parameters
     *      - Resets nextSeasonStartTime to 0
     */
    function _startFirstSeason() internal virtual {
        require(currentSeason == 0, StakingPoolBaseSeasonNotEnded());

        uint startTime = nextSeasonStartTime;
        nextSeasonStartTime = 0;

        currentSeason = 1;
        _createSeason(1, startTime);

        emit SeasonRolledOver(0, 1, 0);
    }

    /**
     * @notice Rolls over from current season to next season
     * @dev Process:
     *      1. Finalizes aggregation for ending season
     *      2. Marks season as finalized (isFinalized = true)
     *      3. Increments season number
     *      4. Calculates next season start (respects manual override)
     *      5. Creates new season
     *      6. Emits SeasonRolledOver event
     *
     *      Note: User data is NOT finalized here (lazy snapshot system)
     */
    function _rolloverSeason() internal virtual {
        uint oldSeasonNumber = currentSeason;
        Season storage oldSeason = seasons[oldSeasonNumber];

        _finalizeSeasonAggregation(oldSeasonNumber);

        oldSeason.isFinalized = true;

        uint newSeasonNumber = oldSeasonNumber + 1;
        currentSeason = newSeasonNumber;

        uint nextStart;
        if (nextSeasonStartTime > 0) {
            nextStart = nextSeasonStartTime;
            nextSeasonStartTime = 0;
        } else {
            nextStart = oldSeason.endTime + 1;
        }

        _createSeason(newSeasonNumber, nextStart);

        emit SeasonRolledOver(oldSeasonNumber, newSeasonNumber, 0);
    }

    /**
     * @notice Calculates season end time based on start time and configuration
     * @param startTime Season start timestamp
     * @return End timestamp for the season (capped by poolEndTime if set)
     */
    function _calculateEndTime(uint startTime) internal view returns (uint) {
        uint endTime = startTime + seasonDuration - 1;
        if (poolEndTime > 0 && endTime > poolEndTime) endTime = poolEndTime;
        return endTime;
    }

    /**
     * @notice Creates a new season with initial values
     * @param seasonNumber Season number (1-indexed)
     * @param startTime Starting timestamp for the season
     * @dev Initializes Season struct with:
     *      - Calculated endTime
     *      - isFinalized = false
     *      - totalPoints = 0 (will be set on finalization)
     *      - seasonTotalStaked = current totalStaked
     *      - lastAggregatedTime = startTime
     *      - aggregatedPoints = 0
     */
    function _createSeason(uint seasonNumber, uint startTime) internal {
        uint endTime = _calculateEndTime(startTime);

        seasons[seasonNumber] = Season({
            seasonNumber: seasonNumber,
            startTime: startTime,
            endTime: endTime,
            isFinalized: false,
            totalPoints: 0,
            seasonTotalStaked: totalStaked,
            lastAggregatedTime: startTime,
            aggregatedPoints: 0,
            forfeitedPoints: 0
        });
    }

    // ============ O(1) Aggregation System ============

    /**
     * @notice Updates season aggregation incrementally (O(1) complexity)
     * @param seasonNum Season number to update
     * @dev Key optimization: Calculates total points without iterating over all users
     *
     *      Mathematical principle:
     *      Σ(each user's points) = Σ(user balance × time) = (total staked × time)
     *
     *      Process:
     *      1. Skip if already up-to-date (lastAggregatedTime >= current block)
     *      2. If no stake, just update lastAggregatedTime
     *      3. Calculate additional points: seasonTotalStaked × blocks elapsed
     *      4. Add to aggregatedPoints
     *      5. Update lastAggregatedTime
     *
     *      Called whenever stake/unstake occurs to maintain accurate running total
     *
     *      Example:
     *      Block 1000: 100 CROSS staked, lastAggregated=1000, aggregated=0
     *      Block 1100: +50 CROSS staked
     *        -> additional = 100 × (1100-1000) = 10,000 points
     *        -> aggregated = 0 + 10,000 = 10,000
     *        -> seasonTotalStaked = 150
     *        -> lastAggregated = 1100
     */
    function _updateSeasonAggregation(uint seasonNum) internal {
        Season storage season = seasons[seasonNum];

        if (season.lastAggregatedTime >= block.timestamp) return;

        // Cap at endTime to prevent overshooting when season has ended
        uint targetBlock = block.timestamp;
        if (season.endTime > 0 && targetBlock > season.endTime) targetBlock = season.endTime;

        if (season.lastAggregatedTime >= targetBlock) return;

        if (season.seasonTotalStaked == 0) {
            season.lastAggregatedTime = targetBlock;
            return;
        }

        uint additionalPoints =
            PointsLib.calculatePoints(season.seasonTotalStaked, season.lastAggregatedTime, targetBlock);

        season.aggregatedPoints += additionalPoints;
        season.lastAggregatedTime = targetBlock;
        emit SeasonAggregationUpdated(seasonNum, season.aggregatedPoints, targetBlock);
    }

    /**
     * @notice Finalizes season aggregation at season end
     * @param seasonNum Season number to finalize
     * @dev - Updates aggregation to season end block
     *      - Copies aggregatedPoints to totalPoints (immutable cache)
     *      - Called during rollover to lock in final total
     *
     *      Handles edge cases:
     *      - lastAggregatedTime > endTime (bug recovery)
     *      - seasonTotalStaked = 0 (no stakers)
     *      - Season already finalized (idempotent)
     *
     *      After finalization:
     *      - totalPoints becomes the source of truth
     *      - Used for all reward calculations
     *      - Never changes again
     */
    function _finalizeSeasonAggregation(uint seasonNum) internal {
        Season storage season = seasons[seasonNum];

        if (season.endTime == 0) return;

        // If already properly finalized, skip
        if (season.totalPoints > 0 && season.lastAggregatedTime >= season.endTime) return;

        uint finalBlock = season.endTime < block.timestamp ? season.endTime : block.timestamp;

        // Handle case where lastAggregatedTime exceeded endTime (bug recovery)
        if (season.lastAggregatedTime > season.endTime) {
            // aggregatedPoints may be overcounted, but we keep it for consistency
            // The overcounted portion will be beyond endTime
            season.totalPoints =
                season.aggregatedPoints > season.forfeitedPoints ? season.aggregatedPoints - season.forfeitedPoints : 0;
            season.lastAggregatedTime = finalBlock;
            return;
        }

        // Normal case: aggregate from lastAggregatedTime to endTime
        if (season.lastAggregatedTime < finalBlock && season.seasonTotalStaked > 0) {
            uint additionalPoints =
                PointsLib.calculatePoints(season.seasonTotalStaked, season.lastAggregatedTime, finalBlock);
            season.aggregatedPoints += additionalPoints;
        }

        season.lastAggregatedTime = finalBlock;
        // totalPoints = aggregatedPoints - forfeitedPoints (몰수 포인트 제외)
        season.totalPoints =
            season.aggregatedPoints > season.forfeitedPoints ? season.aggregatedPoints - season.forfeitedPoints : 0;
    }

    // ============ Lazy Snapshot System ============

    /**
     * @notice Ensures all previous seasons are finalized for a user (lazy evaluation)
     * @param user User address to finalize
     * @dev Gas-efficient approach:
     *      - Doesn't snapshot all users during rollover (would be O(N) gas)
     *      - Instead, snapshots each user only when they interact
     *      - Tracks last finalized season to avoid redundant work
     *
     *      Process:
     *      1. Get last finalized season for user (default = 1)
     *      2. Loop through all seasons from last finalized to current-1
     *      3. Call _ensureUserSeasonSnapshot for each
     *      4. Update lastFinalizedSeason tracker
     *
     *      Called before:
     *      - Staking (to update points before new stake)
     *      - Withdrawing (to snapshot before exit)
     *      - Claiming (to ensure accurate data)
     *
     *      Example:
     *      User stakes in season 1, doesn't interact until season 5
     *      On first season 5 interaction:
     *      - Snapshots seasons 1, 2, 3, 4 retroactively
     *      - Uses stored position data and season end blocks
     *      - User pays gas only for their own snapshots
     */
    function _ensureUserAllPreviousSeasons(address user) internal {
        if (currentSeason == 0) return;

        uint startSeason = lastFinalizedSeason[user];
        if (startSeason == 0) startSeason = 1;

        if (startSeason >= currentSeason) return;

        for (uint i = startSeason; i < currentSeason;) {
            _ensureUserSeasonSnapshot(user, i);
            unchecked {
                ++i;
            }
        }

        lastFinalizedSeason[user] = currentSeason - 1;
    }

    /**
     * @notice Creates snapshot of user's data for a specific season (if not already done)
     * @param user User address
     * @param seasonNum Season number to snapshot
     * @dev Retroactively calculates user's points for a past season using stored data
     *
     *      Preconditions:
     *      - Season must be finalized (isFinalized = true)
     *      - User data not already finalized for this season
     *
     *      Process:
     *      1. Skip if season not finalized or already processed
     *      2. If user has UserSeasonData.balance > 0 (mid-season stake/unstake):
     *         - Use userData.balance and userData.joinTime
     *         - Calculate points from join to season end
     *      3. Else use current position data:
     *         - Use position.balance and position.lastUpdateTime
     *         - If joined before season start, use season.startTime
     *         - Calculate points for full participation
     *      4. Mark userData.finalized = true
     *
     *      Edge cases:
     *      - No balance in season -> finalized with 0 points
     *      - Joined after season ended -> finalized with 0 points
     *      - Joined before season started -> points from season start
     *
     *      This enables lazy evaluation: data computed only when needed (claim, new stake, etc.)
     */
    function _ensureUserSeasonSnapshot(address user, uint seasonNum) internal virtual {
        if (seasonNum == 0 || seasonNum > currentSeason) return;

        Season storage season = seasons[seasonNum];
        if (!season.isFinalized) return;

        UserSeasonData storage userData = userSeasonData[seasonNum][user];
        if (userData.finalized) return;

        StakePosition storage position = userStakes[user];
        uint lastUpdate = position.lastUpdateTime;

        if (userData.balance > 0) {
            uint userJoinBlock = userData.joinTime > 0 ? userData.joinTime : season.startTime;

            if (userJoinBlock < season.startTime) {
                userData.joinTime = season.startTime;
                userData.points = PointsLib.calculatePoints(userData.balance, season.startTime, season.endTime);
            } else if (userJoinBlock <= season.endTime) {
                userData.points = PointsLib.calculatePoints(userData.balance, userJoinBlock, season.endTime);
            }
            userData.finalized = true;
            return;
        }

        if (lastUpdate > season.endTime) {
            userData.finalized = true;
            return;
        }

        if (position.balance == 0) {
            userData.finalized = true;
            return;
        }

        uint positionJoinBlock = lastUpdate;

        if (positionJoinBlock < season.startTime) {
            userData.balance = position.balance;
            userData.joinTime = season.startTime;
            userData.points = PointsLib.calculatePoints(position.balance, season.startTime, season.endTime);
        } else if (positionJoinBlock <= season.endTime) {
            userData.balance = position.balance;
            userData.joinTime = positionJoinBlock;
            userData.points = PointsLib.calculatePoints(position.balance, positionJoinBlock, season.endTime);
        }

        userData.finalized = true;
    }

    // ============ Helper Functions ============

    /**
     * @notice Checks if an address is an approved router
     * @param router Address to check
     * @return True if router has ROUTER_ROLE
     * @dev Can be overridden to add global router approval logic
     */
    function _isApprovedRouter(address router) internal view virtual returns (bool) {
        return hasRole(ROUTER_ROLE, router);
    }

    /**
     * @notice Checks if a season is currently active
     * @return True if there's an active season that accepts stakes
     * @dev Handles multiple cases:
     *      - Pool ended -> false
     *      - Season 0 (not started) -> checks if virtual season is active
     *      - Current season -> checks if within season blocks
     *      - Between seasons -> checks if next virtual season is active
     *
     *      Virtual season support enables view functions to work before rollover
     */
    function isSeasonActive() public view virtual returns (bool) {
        if (poolEndTime > 0 && block.timestamp >= poolEndTime) return false;

        if (currentSeason == 0) {
            if (block.timestamp < nextSeasonStartTime) return false;
            if (poolEndTime > 0 && nextSeasonStartTime >= poolEndTime) return false;

            uint virtualEndTime = _calculateEndTime(nextSeasonStartTime);
            return block.timestamp <= virtualEndTime;
        }

        Season storage season = seasons[currentSeason];

        if (block.timestamp > season.endTime) {
            if (nextSeasonStartTime > 0 && block.timestamp < nextSeasonStartTime) return false;

            uint nextStart = nextSeasonStartTime > 0 ? nextSeasonStartTime : season.endTime + 1;

            if (poolEndTime > 0 && nextStart >= poolEndTime) return false;

            uint virtualEndTime = _calculateEndTime(nextStart);
            return block.timestamp <= virtualEndTime;
        }

        return !season.isFinalized && block.timestamp >= season.startTime && block.timestamp <= season.endTime;
    }

    function getStakingPower(address user) external view virtual returns (uint) {
        return userStakes[user].balance;
    }

    function getTotalStakingPower() external view virtual returns (uint) {
        return totalStaked;
    }

    // ============ Manual Season Management ============

    /**
     * @notice Manually rolls over multiple seasons (for >50 pending seasons)
     * @param maxRollovers Maximum number of seasons to rollover
     * @return rolloversPerformed Actual number of seasons rolled over
     * @dev Used when automatic rollover limit (50) is exceeded
     *
     *      Use case:
     *      - Project abandoned for long period (e.g., 2 years)
     *      - 100+ seasons need rollover
     *      - Cannot happen in single transaction (MAX_AUTO_ROLLOVERS = 50)
     *
     *      Solution:
     *      1. Call manualRolloverSeasons(50) multiple times
     *      2. Each call processes up to 50 seasons
     *      3. Continue until getPendingSeasonRollovers() returns 0
     *
     *      Permissions: MANAGER_ROLE or protocol admin
     *      Max limit: 100 seasons per call (safety cap)
     *
     *      Example:
     *      150 pending seasons:
     *      - Call 1: processes 50, remaining = 100
     *      - Call 2: processes 50, remaining = 50
     *      - Call 3: processes 50, remaining = 0
     */
    function manualRolloverSeasons(uint maxRollovers)
        external
        virtual
        onlyRole(MANAGER_ROLE)
        returns (uint rolloversPerformed)
    {
        require(maxRollovers > 0 && maxRollovers <= 100, StakingPoolBaseInvalidMaxRollovers());

        uint startSeason = currentSeason;

        if (poolEndTime > 0 && block.timestamp >= poolEndTime) {
            emit ManualRolloverCompleted(0, startSeason, currentSeason);
            return 0;
        }

        if (currentSeason == 0) {
            if (block.timestamp >= nextSeasonStartTime) {
                _startFirstSeason();
                rolloversPerformed = 1;
                emit ManualRolloverCompleted(1, 0, currentSeason);
            } else {
                emit ManualRolloverCompleted(0, 0, 0);
                return 0;
            }
        }

        uint count = 0;
        while (currentSeason > 0 && count < maxRollovers) {
            Season storage current = seasons[currentSeason];

            if (block.timestamp <= current.endTime) break;
            if (nextSeasonStartTime > 0 && block.timestamp < nextSeasonStartTime) break;

            _rolloverSeason();

            unchecked {
                ++count;
                ++rolloversPerformed;
            }
        }

        emit ManualRolloverCompleted(rolloversPerformed, startSeason, currentSeason);

        return rolloversPerformed;
    }

    /**
     * @notice Calculates how many seasons are pending rollover
     * @return pendingSeasons Number of seasons that need to be rolled over
     * @dev Complex calculation that handles multiple scenarios:
     *
     *      Scenario 1: No seasons started yet (currentSeason = 0)
     *      - Calculate how many seasons would exist from firstSeasonStart to now
     *
     *      Scenario 2: Current season ended
     *      - Calculate how many new seasons are needed to reach current block
     *
     *      Used by:
     *      - Frontend to show status
     *      - Admins to decide if manual rollover needed
     *      - Monitoring to detect abandoned projects
     *
     *      Returns 0 if:
     *      - Pool has ended
     *      - Current season is still active
     *      - Next season start time not yet reached
     */
    function getPendingSeasonRollovers() external view returns (uint pendingSeasons) {
        if (poolEndTime > 0 && block.timestamp >= poolEndTime) return 0;

        if (currentSeason == 0) {
            if (block.timestamp < nextSeasonStartTime) return 0;

            uint currentBlk = block.timestamp;
            uint startTime = nextSeasonStartTime;
            uint count = 0;

            count = 1;
            uint endTime = _calculateEndTime(startTime);

            if (poolEndTime > 0 && endTime >= poolEndTime) return count;

            if (currentBlk > endTime) {
                uint blocksAfterEnd = currentBlk - endTime;
                uint additionalSeasons = (blocksAfterEnd + seasonDuration - 1) / seasonDuration;

                if (poolEndTime > 0) {
                    uint nextStartBlk = endTime + 1;
                    for (uint i = 0; i < additionalSeasons; i++) {
                        if (nextStartBlk >= poolEndTime) break;
                        count++;
                        nextStartBlk += seasonDuration;
                    }
                } else {
                    count += additionalSeasons;
                }
            }

            return count;
        }

        if (nextSeasonStartTime > 0 && block.timestamp < nextSeasonStartTime) return 0;

        Season storage current = seasons[currentSeason];

        if (block.timestamp <= current.endTime) return 0;

        uint nextStart = nextSeasonStartTime > 0 ? nextSeasonStartTime : current.endTime + 1;

        if (poolEndTime > 0 && nextStart >= poolEndTime) return 0;

        uint rolloverCount = 0;
        uint currentBlock = block.timestamp;

        while (nextStart <= currentBlock) {
            uint endTime = _calculateEndTime(nextStart);
            rolloverCount++;

            if ((poolEndTime > 0 && endTime >= poolEndTime) || currentBlock <= endTime) break;

            nextStart = endTime + 1;

            if (rolloverCount >= 100) break;
        }

        return rolloverCount;
    }
}
