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
 *      - Block-based seasons with automatic rollovers (up to 50 at once)
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
 *      - Points accumulate from season start block
 *      - Optional feature (disabled if preDepositStartBlock = 0)
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

    /// @notice Thrown when season blocks parameter is invalid (zero)
    error StakingPoolBaseInvalidSeasonBlocks();

    /// @notice Thrown when time unit parameter is invalid (zero)
    error StakingPoolBaseInvalidTimeUnit();

    /// @notice Thrown when attempting action before season has ended
    error StakingPoolBaseSeasonNotEnded();

    /// @notice Thrown when no active season exists
    error StakingPoolBaseNoActiveSeason();

    /// @notice Thrown when start block parameter is invalid
    error StakingPoolBaseInvalidStartBlock();

    /// @notice Thrown when end block parameter is invalid
    error StakingPoolBaseInvalidEndBlock();

    /// @notice Thrown when too many seasons need to be rolled over (>50)
    error StakingPoolBaseTooManySeasons();

    /// @notice Thrown when external caller tries to call self-only function
    error StakingPoolBaseOnlySelf();

    /// @notice Thrown when manual rollover parameter is invalid
    error StakingPoolBaseInvalidMaxRollovers();

    /// @notice Thrown when pre-deposit block is after first season start
    error StakingPoolBaseInvalidPreDepositBlock();

    // ============ Structs ============

    /**
     * @notice User's current staking position (persists across seasons)
     * @param balance Current staked amount
     * @param points Temporary points (moved to UserSeasonData on updates)
     * @param lastUpdateBlock Last block when position was updated
     */
    struct StakePosition {
        uint balance;
        uint points;
        uint lastUpdateBlock;
    }

    /**
     * @notice Season data structure
     * @param seasonNumber Season identifier (1-indexed)
     * @param startBlock First block of the season
     * @param endBlock Last block of the season (inclusive)
     * @param isFinalized Whether season has ended and been finalized
     * @param totalPoints Finalized total points (cached, immutable after finalization)
     * @param seasonTotalStaked Current total staked in season (for aggregation)
     * @param lastAggregatedBlock Last block when aggregation was updated
     * @param aggregatedPoints Accumulated aggregated points (real-time during season)
     * @param forfeitedPoints Points forfeited from withdrawals during season
     */
    struct Season {
        uint seasonNumber;
        uint startBlock;
        uint endBlock;
        bool isFinalized;
        uint totalPoints;
        uint seasonTotalStaked;
        uint lastAggregatedBlock;
        uint aggregatedPoints;
        uint forfeitedPoints;
    }

    /**
     * @notice User's data for a specific season (lazy snapshot)
     * @param points User's points in this season
     * @param balance User's balance in this season
     * @param joinBlock Block when user joined this season
     * @param lastPointsBlock Last block when points were calculated
     * @param claimed Whether user has claimed rewards for this season
     * @param finalized Whether snapshot has been taken (lazy finalization)
     */
    struct UserSeasonData {
        uint points;
        uint balance;
        uint joinBlock;
        uint lastPointsBlock;
        bool claimed;
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
    uint public seasonBlocks;

    /// @notice Block number when pool ends (0 = infinite)
    uint public poolEndBlock;

    /// @notice Block number when next season starts (manual override)
    uint public nextSeasonStartBlock;

    /// @notice Block number when pre-deposit starts (0 = disabled)
    /// @dev Pre-deposit allows staking before season 1 starts (season 1 only)
    uint public preDepositStartBlock;

    /// @notice Time unit for points calculation in seconds (default: 1 hour)
    uint public pointsTimeUnit = 1 hours;

    /// @notice Block time in seconds (default: 1 second per block)
    uint public blockTime = 1;

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
    event SeasonClaimed(address indexed user, uint indexed season, uint points);

    /// @notice Emitted when manual rollover is completed
    event ManualRolloverCompleted(uint rolloversPerformed, uint fromSeason, uint toSeason);

    /// @notice Emitted when points are forfeited due to withdrawal
    event PointsForfeited(address indexed user, uint indexed season, uint amount);

    /// @notice Emitted when season aggregation is updated
    event SeasonAggregationUpdated(uint indexed season, uint aggregatedPoints, uint lastAggregatedBlock);

    // ============ Constructor ============

    /**
     * @notice Initializes the staking pool base contract
     * @param _stakingToken Token to be staked (WCROSS)
     * @param admin Initial admin address
     * @param _seasonBlocks Number of blocks per season
     * @param _firstSeasonStartBlock Block when first season starts
     * @param _poolEndBlock Block when pool ends (0 = infinite)
     * @param _preDepositStartBlock Block when pre-deposit starts (0 = disabled)
     * @dev Validates:
     *      - All addresses are non-zero
     *      - Season blocks > 0
     *      - First season start block > 0
     *      - Pool end block > first season end (if not 0)
     *      - Pre-deposit block < first season start (if not 0)
     */
    constructor(
        IERC20 _stakingToken,
        address admin,
        uint _seasonBlocks,
        uint _firstSeasonStartBlock,
        uint _poolEndBlock,
        uint _preDepositStartBlock
    ) CrossStakingBase(admin) {
        _validateAddress(address(_stakingToken));
        require(_seasonBlocks != 0, StakingPoolBaseInvalidSeasonBlocks());
        require(_firstSeasonStartBlock != 0, StakingPoolBaseInvalidStartBlock());
        require(
            _poolEndBlock == 0 || _poolEndBlock > _firstSeasonStartBlock + _seasonBlocks,
            StakingPoolBaseInvalidEndBlock()
        );
        if (_preDepositStartBlock > 0 && _preDepositStartBlock > _firstSeasonStartBlock) {
            revert StakingPoolBaseInvalidPreDepositBlock();
        }

        stakingToken = _stakingToken;
        seasonBlocks = _seasonBlocks;
        nextSeasonStartBlock = _firstSeasonStartBlock;
        poolEndBlock = _poolEndBlock;
        preDepositStartBlock = _preDepositStartBlock;

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
        require(poolEndBlock == 0 || block.number < poolEndBlock, StakingPoolBaseNoActiveSeason());

        _ensureSeason();

        if (currentSeason == 0) {
            if (preDepositStartBlock > 0 && block.number >= preDepositStartBlock) {} else {
                require(block.number >= nextSeasonStartBlock, StakingPoolBaseNoActiveSeason());
            }
        } else {
            require(isSeasonActive(), StakingPoolBaseNoActiveSeason());
        }

        _ensureUserAllPreviousSeasons(user);

        StakePosition storage position = userStakes[user];
        uint oldBalance = position.balance;

        if (oldBalance > 0) {
            uint effectiveStart = position.lastUpdateBlock < seasons[currentSeason].startBlock
                ? seasons[currentSeason].startBlock
                : position.lastUpdateBlock;
            uint additionalPoints =
                PointsLib.calculatePoints(position.balance, effectiveStart, block.number, blockTime, pointsTimeUnit);
            UserSeasonData storage currentUserData = userSeasonData[currentSeason][user];
            currentUserData.points += additionalPoints;
            currentUserData.lastPointsBlock = block.number;
            position.points = 0;
        }

        uint newBalance = oldBalance + amount;
        require(newBalance >= MIN_STAKE, StakingPoolBaseBelowMinStake());

        stakingToken.safeTransferFrom(from, address(this), amount);

        position.balance = newBalance;
        position.lastUpdateBlock = block.number;
        totalStaked += amount;

        Season storage currentSeasonData = seasons[currentSeason];
        _updateSeasonAggregation(currentSeason);
        currentSeasonData.seasonTotalStaked += amount;

        UserSeasonData storage seasonData = userSeasonData[currentSeason][user];
        if (seasonData.balance == 0) {
            Season storage current = seasons[currentSeason];
            seasonData.balance = newBalance;
            seasonData.joinBlock = block.number < current.startBlock ? current.startBlock : block.number;
            seasonData.lastPointsBlock = block.number;
        } else {
            seasonData.balance = newBalance;
            seasonData.lastPointsBlock = block.number;
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
        _ensureSeason();
        _ensureUserAllPreviousSeasons(user);

        StakePosition storage position = userStakes[user];
        require(position.balance != 0, StakingPoolBaseNoPosition());

        uint amount = position.balance;

        if (currentSeason > 0) {
            Season storage currentSeasonData = seasons[currentSeason];
            UserSeasonData storage seasonData = userSeasonData[currentSeason][user];

            _updateSeasonAggregation(currentSeason);

            uint userForfeitedPoints = 0;
            if (seasonData.balance > 0) {
                userForfeitedPoints = seasonData.points;

                uint effectiveStart = seasonData.lastPointsBlock > 0 ? seasonData.lastPointsBlock : seasonData.joinBlock;
                if (effectiveStart < currentSeasonData.startBlock) effectiveStart = currentSeasonData.startBlock;

                uint additionalPoints = PointsLib.calculatePoints(
                    seasonData.balance, effectiveStart, block.number, blockTime, pointsTimeUnit
                );
                userForfeitedPoints += additionalPoints;
            } else if (position.balance > 0) {
                uint lastUpdate = position.lastUpdateBlock;
                uint effectiveStart =
                    lastUpdate < currentSeasonData.startBlock ? currentSeasonData.startBlock : lastUpdate;
                userForfeitedPoints =
                    PointsLib.calculatePoints(position.balance, effectiveStart, block.number, blockTime, pointsTimeUnit);
            }

            seasonData.points = 0;
            seasonData.balance = 0;

            if (userForfeitedPoints > 0) {
                currentSeasonData.forfeitedPoints += userForfeitedPoints;
                emit PointsForfeited(user, currentSeason, userForfeitedPoints);
            }

            currentSeasonData.seasonTotalStaked -= amount;
        }

        position.balance = 0;
        position.points = 0;
        position.lastUpdateBlock = block.number;
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
     *      - Next season start block set -> waits for that block
     *      - Multiple seasons passed -> all rolled over automatically
     *      - >50 seasons pending -> reverts (use manualRolloverSeasons)
     */
    function _ensureSeason() internal virtual {
        if (poolEndBlock > 0 && (nextSeasonStartBlock == 0 || nextSeasonStartBlock <= poolEndBlock)) {
            if (block.number >= poolEndBlock) return;
        }

        if (currentSeason == 0) {
            if (block.number >= nextSeasonStartBlock) _startFirstSeason();
            else return;
        }

        uint maxRollovers = MAX_AUTO_ROLLOVERS;
        uint rolloversPerformed = 0;

        while (currentSeason > 0 && rolloversPerformed < maxRollovers) {
            Season storage current = seasons[currentSeason];

            if (block.number <= current.endBlock) break;
            if (nextSeasonStartBlock > 0 && block.number < nextSeasonStartBlock) break;

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
     *      - Resets nextSeasonStartBlock to 0
     */
    function _startFirstSeason() internal virtual {
        require(currentSeason == 0, StakingPoolBaseSeasonNotEnded());

        uint startBlock = nextSeasonStartBlock;
        nextSeasonStartBlock = 0;

        currentSeason = 1;
        _createSeason(1, startBlock);

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
        if (nextSeasonStartBlock > 0) {
            nextStart = nextSeasonStartBlock;
            nextSeasonStartBlock = 0;
        } else {
            nextStart = oldSeason.endBlock + 1;
        }

        _createSeason(newSeasonNumber, nextStart);

        emit SeasonRolledOver(oldSeasonNumber, newSeasonNumber, 0);
    }

    /**
     * @notice Calculates season end block based on start block and configuration
     * @param startBlock Season start block
     * @return End block for the season (capped by poolEndBlock if set)
     */
    function _calculateEndBlock(uint startBlock) internal view returns (uint) {
        uint endBlock = startBlock + seasonBlocks - 1;
        if (poolEndBlock > 0 && endBlock > poolEndBlock) endBlock = poolEndBlock;
        return endBlock;
    }

    /**
     * @notice Creates a new season with initial values
     * @param seasonNumber Season number (1-indexed)
     * @param startBlock Starting block for the season
     * @dev Initializes Season struct with:
     *      - Calculated endBlock
     *      - isFinalized = false
     *      - totalPoints = 0 (will be set on finalization)
     *      - seasonTotalStaked = current totalStaked
     *      - lastAggregatedBlock = startBlock
     *      - aggregatedPoints = 0
     */
    function _createSeason(uint seasonNumber, uint startBlock) internal {
        uint endBlock = _calculateEndBlock(startBlock);

        seasons[seasonNumber] = Season({
            seasonNumber: seasonNumber,
            startBlock: startBlock,
            endBlock: endBlock,
            isFinalized: false,
            totalPoints: 0,
            seasonTotalStaked: totalStaked,
            lastAggregatedBlock: startBlock,
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
     *      1. Skip if already up-to-date (lastAggregatedBlock >= current block)
     *      2. If no stake, just update lastAggregatedBlock
     *      3. Calculate additional points: seasonTotalStaked × blocks elapsed
     *      4. Add to aggregatedPoints
     *      5. Update lastAggregatedBlock
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

        if (season.lastAggregatedBlock >= block.number) return;

        // Cap at endBlock to prevent overshooting when season has ended
        uint targetBlock = block.number;
        if (season.endBlock > 0 && targetBlock > season.endBlock) targetBlock = season.endBlock;

        if (season.lastAggregatedBlock >= targetBlock) return;

        if (season.seasonTotalStaked == 0) {
            season.lastAggregatedBlock = targetBlock;
            return;
        }

        uint additionalPoints = PointsLib.calculatePoints(
            season.seasonTotalStaked, season.lastAggregatedBlock, targetBlock, blockTime, pointsTimeUnit
        );

        season.aggregatedPoints += additionalPoints;
        season.lastAggregatedBlock = targetBlock;
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
     *      - lastAggregatedBlock > endBlock (bug recovery)
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

        if (season.endBlock == 0) return;

        // If already properly finalized, skip
        if (season.totalPoints > 0 && season.lastAggregatedBlock >= season.endBlock) return;

        uint finalBlock = season.endBlock < block.number ? season.endBlock : block.number;

        // Handle case where lastAggregatedBlock exceeded endBlock (bug recovery)
        if (season.lastAggregatedBlock > season.endBlock) {
            // aggregatedPoints may be overcounted, but we keep it for consistency
            // The overcounted portion will be beyond endBlock
            season.totalPoints = season.aggregatedPoints;
            season.lastAggregatedBlock = finalBlock;
            return;
        }

        // Normal case: aggregate from lastAggregatedBlock to endBlock
        if (season.lastAggregatedBlock < finalBlock && season.seasonTotalStaked > 0) {
            uint additionalPoints = PointsLib.calculatePoints(
                season.seasonTotalStaked, season.lastAggregatedBlock, finalBlock, blockTime, pointsTimeUnit
            );
            season.aggregatedPoints += additionalPoints;
        }

        season.lastAggregatedBlock = finalBlock;
        season.totalPoints = season.aggregatedPoints;
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
     *         - Use userData.balance and userData.joinBlock
     *         - Calculate points from join to season end
     *      3. Else use current position data:
     *         - Use position.balance and position.lastUpdateBlock
     *         - If joined before season start, use season.startBlock
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
        uint lastUpdate = position.lastUpdateBlock;

        if (userData.balance > 0) {
            uint userJoinBlock = userData.joinBlock > 0 ? userData.joinBlock : season.startBlock;

            if (userJoinBlock < season.startBlock) {
                userData.joinBlock = season.startBlock;
                userData.points = PointsLib.calculatePoints(
                    userData.balance, season.startBlock, season.endBlock, blockTime, pointsTimeUnit
                );
            } else if (userJoinBlock <= season.endBlock) {
                userData.points = PointsLib.calculatePoints(
                    userData.balance, userJoinBlock, season.endBlock, blockTime, pointsTimeUnit
                );
            }
            userData.finalized = true;
            return;
        }

        if (lastUpdate > season.endBlock) {
            userData.finalized = true;
            return;
        }

        if (position.balance == 0) {
            userData.finalized = true;
            return;
        }

        uint positionJoinBlock = lastUpdate;

        if (positionJoinBlock < season.startBlock) {
            userData.balance = position.balance;
            userData.joinBlock = season.startBlock;
            userData.points = PointsLib.calculatePoints(
                position.balance, season.startBlock, season.endBlock, blockTime, pointsTimeUnit
            );
        } else if (positionJoinBlock <= season.endBlock) {
            userData.balance = position.balance;
            userData.joinBlock = positionJoinBlock;
            userData.points = PointsLib.calculatePoints(
                position.balance, positionJoinBlock, season.endBlock, blockTime, pointsTimeUnit
            );
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
        if (poolEndBlock > 0 && block.number >= poolEndBlock) return false;

        if (currentSeason == 0) {
            if (block.number < nextSeasonStartBlock) return false;
            if (poolEndBlock > 0 && nextSeasonStartBlock >= poolEndBlock) return false;

            uint virtualEndBlock = _calculateEndBlock(nextSeasonStartBlock);
            return block.number <= virtualEndBlock;
        }

        Season storage season = seasons[currentSeason];

        if (block.number > season.endBlock) {
            if (nextSeasonStartBlock > 0 && block.number < nextSeasonStartBlock) return false;

            uint nextStart = nextSeasonStartBlock > 0 ? nextSeasonStartBlock : season.endBlock + 1;

            if (poolEndBlock > 0 && nextStart >= poolEndBlock) return false;

            uint virtualEndBlock = _calculateEndBlock(nextStart);
            return block.number <= virtualEndBlock;
        }

        return !season.isFinalized && block.number >= season.startBlock && block.number <= season.endBlock;
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

        if (poolEndBlock > 0 && block.number >= poolEndBlock) {
            emit ManualRolloverCompleted(0, startSeason, currentSeason);
            return 0;
        }

        if (currentSeason == 0) {
            if (block.number >= nextSeasonStartBlock) {
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

            if (block.number <= current.endBlock) break;
            if (nextSeasonStartBlock > 0 && block.number < nextSeasonStartBlock) break;

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
     *      - Next season start block not yet reached
     */
    function getPendingSeasonRollovers() external view returns (uint pendingSeasons) {
        if (poolEndBlock > 0 && block.number >= poolEndBlock) return 0;

        if (currentSeason == 0) {
            if (block.number < nextSeasonStartBlock) return 0;

            uint currentBlk = block.number;
            uint startBlock = nextSeasonStartBlock;
            uint count = 0;

            count = 1;
            uint endBlock = _calculateEndBlock(startBlock);

            if (poolEndBlock > 0 && endBlock >= poolEndBlock) return count;

            if (currentBlk > endBlock) {
                uint blocksAfterEnd = currentBlk - endBlock;
                uint additionalSeasons = (blocksAfterEnd + seasonBlocks - 1) / seasonBlocks;

                if (poolEndBlock > 0) {
                    uint nextStartBlk = endBlock + 1;
                    for (uint i = 0; i < additionalSeasons; i++) {
                        if (nextStartBlk >= poolEndBlock) break;
                        count++;
                        nextStartBlk += seasonBlocks;
                    }
                } else {
                    count += additionalSeasons;
                }
            }

            return count;
        }

        if (nextSeasonStartBlock > 0 && block.number < nextSeasonStartBlock) return 0;

        Season storage current = seasons[currentSeason];

        if (block.number <= current.endBlock) return 0;

        uint nextStart = nextSeasonStartBlock > 0 ? nextSeasonStartBlock : current.endBlock + 1;

        if (poolEndBlock > 0 && nextStart >= poolEndBlock) return 0;

        uint rolloverCount = 0;
        uint currentBlock = block.number;

        while (nextStart <= currentBlock) {
            uint endBlock = _calculateEndBlock(nextStart);
            rolloverCount++;

            if ((poolEndBlock > 0 && endBlock >= poolEndBlock) || currentBlock <= endBlock) break;

            nextStart = endBlock + 1;

            if (rolloverCount >= 100) break;
        }

        return rolloverCount;
    }
}
