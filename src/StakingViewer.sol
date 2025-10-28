// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./interfaces/IRewardPool.sol";
import "./interfaces/IStakingPool.sol";
import "./interfaces/IStakingProtocol.sol";
import "./libraries/PointsLib.sol";

/**
 * @title StakingViewer
 * @notice Unified view contract for querying staking and reward data across all projects
 * @dev Provides gas-free read-only functions for:
 *      - User staking positions and points
 *      - Season information and history
 *      - Reward calculations and previews
 *      - Virtual season data (before rollover)
 *      - Batch queries for multiple projects
 *
 *      Key Features:
 *      - No state changes (all view/pure functions)
 *      - Supports virtual seasons (calculates data for seasons not yet rolled over)
 *      - Batch operations for efficient multi-project queries
 *      - Comprehensive reward preview before claiming
 *
 *      Virtual Season Support:
 *      When multiple seasons have passed without rollover on-chain,
 *      this contract calculates what the data would be if rollovers had occurred.
 *      This allows frontends to show accurate information without requiring
 *      expensive rollover transactions.
 *
 *      Use Cases:
 *      - Frontend dashboards
 *      - Reward calculators
 *      - Analytics and reporting
 *      - User position tracking
 */
contract StakingViewer {
    using PointsLib for *;

    /// @notice StakingProtocol factory contract
    IStakingProtocol public immutable protocol;

    /**
     * @notice Initializes the viewer contract
     * @param _protocol StakingProtocol address
     */
    constructor(address _protocol) {
        require(_protocol != address(0), "Invalid protocol");
        protocol = IStakingProtocol(_protocol);
    }

    // ============ Internal Helpers ============

    /**
     * @notice Gets pool addresses for a project
     * @param projectID Project ID
     * @return pool StakingPool interface
     * @return rewardPool RewardPool interface
     */
    function _getPools(uint projectID) internal view returns (IStakingPool pool, IRewardPool rewardPool) {
        (address stakingPool, address rewardPoolAddr,,,,,) = protocol.projects(projectID);
        pool = IStakingPool(stakingPool);
        rewardPool = IRewardPool(rewardPoolAddr);
    }

    /**
     * @notice Calculates the first season's start block
     * @param pool StakingPool contract
     * @return ok True if calculation succeeded
     * @return firstStart First season start block
     * @dev Attempts to derive from current season or nextSeasonStartBlock
     */
    function _calcFirstSeasonStart(IStakingPool pool) internal view returns (bool ok, uint firstStart) {
        (uint season, uint startBlock,,) = pool.getCurrentSeasonInfo();
        if (season > 0 && startBlock > 0) {
            uint sb = pool.seasonBlocks();
            // Reverse calculate: season is 1-based
            firstStart = startBlock - ((season - 1) * sb);
            return (true, firstStart);
        }
        // Season not started yet - use nextSeasonStartBlock
        uint nextStart = pool.nextSeasonStartBlock();
        if (nextStart > 0) return (true, nextStart);
        return (false, 0);
    }

    /**
     * @notice Calculates block range for any season number
     * @param pool StakingPool contract
     * @param season Season number
     * @return ok True if calculation succeeded
     * @return startBlock Season start block
     * @return endBlock Season end block
     * @dev Calculates virtual season ranges even if seasons haven't been rolled over on-chain
     */
    function _calcSeasonRange(IStakingPool pool, uint season)
        internal
        view
        returns (bool ok, uint startBlock, uint endBlock)
    {
        (bool okStart, uint firstStart) = _calcFirstSeasonStart(pool);
        if (!okStart || season == 0) return (false, 0, 0);

        uint sb = pool.seasonBlocks();
        startBlock = firstStart + ((season - 1) * sb);
        endBlock = startBlock + sb - 1;

        uint pe = pool.poolEndBlock();
        if (pe > 0) {
            if (startBlock > pe) return (false, 0, 0);
            if (endBlock > pe) endBlock = pe;
        }
        return (true, startBlock, endBlock);
    }

    /**
     * @notice Calculates virtual season data for seasons not yet rolled over on-chain
     * @param pool StakingPool contract
     * @param season Season number
     * @param user User address (address(0) calculates only totalPoints)
     * @return userPoints User's points in the virtual season
     * @return totalPoints Total points in the virtual season
     * @dev Virtual calculation allows view functions to work without requiring rollover
     *
     *      Process:
     *      1. Calculate season's block range
     *      2. Determine effective end block (past vs current/future seasons)
     *      3. Calculate user points from their position
     *      4. Calculate total points from totalStaked
     *
     *      Use case: User wants to see their points for season 5, but only seasons 1-3
     *      have been rolled over. This calculates season 4-5 data virtually.
     */
    function _calculateVirtualSeasonData(IStakingPool pool, uint season, address user)
        internal
        view
        returns (uint userPoints, uint totalPoints)
    {
        (bool ok, uint startBlock, uint endBlock) = _calcSeasonRange(pool, season);
        if (!ok) return (0, 0);

        if (block.number < startBlock) return (0, 0);

        uint currentSeason = pool.currentSeason();

        // 과거 시즌: endBlock 사용 (종료된 시즌은 고정)
        // 현재/미래 시즌: 현재 블록까지만 반영
        uint toBlock = (season < currentSeason) ? endBlock : (block.number < endBlock ? block.number : endBlock);

        uint blockTime = pool.blockTime();
        uint timeUnit = pool.pointsTimeUnit();

        // 유저 포인트 계산 (user가 address(0)이 아닌 경우만)
        if (user != address(0)) {
            (uint balance,, uint lastUpdate) = pool.getStakePosition(user);

            if (balance > 0) {
                // 유저의 시작 블록: lastUpdate와 시즌 시작 블록 중 큰 값
                uint fromBlock = lastUpdate > startBlock ? lastUpdate : startBlock;

                // lastUpdate가 시즌 종료 후면 참여 안 함
                if (fromBlock <= toBlock) {
                    userPoints = PointsLib.calculatePoints(balance, fromBlock, toBlock, blockTime, timeUnit);
                }
            }
        }

        // 토탈 포인트 계산: 현재 totalStaked 기준
        uint totalStaked = pool.totalStaked();
        if (totalStaked > 0) {
            totalPoints = PointsLib.calculatePoints(totalStaked, startBlock, toBlock, blockTime, timeUnit);
        }

        return (userPoints, totalPoints);
    }

    // ============ Public View Functions ============

    /**
     * @notice Returns user's stake information for a project
     * @param projectID Project ID
     * @param user User address
     * @return balance Staked amount
     * @return points Current season points
     * @return lastUpdateBlock Last update block
     */
    function getStakeInfo(uint projectID, address user)
        external
        view
        returns (uint balance, uint points, uint lastUpdateBlock)
    {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getStakePosition(user);
    }

    /**
     * @notice Returns current season number for a project
     * @param projectID Project ID
     * @return season Current season number
     */
    function getCurrentSeason(uint projectID) external view returns (uint season) {
        (IStakingPool pool,) = _getPools(projectID);
        (season,,,) = pool.getCurrentSeasonInfo();
    }

    /**
     * @notice Returns user's current season points
     * @param projectID Project ID
     * @param user User address
     * @return points Current season points
     */
    function getUserPoints(uint projectID, address user) external view returns (uint points) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getUserPoints(user);
    }

    /**
     * @notice Returns user's staking power (same as balance)
     * @param projectID Project ID
     * @param user User address
     * @return stakingPower Staking power
     */
    function getStakingPower(uint projectID, address user) external view returns (uint stakingPower) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getStakingPower(user);
    }

    /**
     * @notice Returns total staking power for a project
     * @param projectID Project ID
     * @return totalPower Total staking power
     */
    function getTotalStakingPower(uint projectID) external view returns (uint totalPower) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getTotalStakingPower();
    }

    /**
     * @notice Returns current season information
     * @param projectID Project ID
     * @return currentSeason Current season number
     * @return seasonStartBlock Season start block
     * @return seasonEndBlock Season end block
     * @return blocksElapsed Blocks elapsed in current season
     */
    function getSeasonInfo(uint projectID)
        external
        view
        returns (uint currentSeason, uint seasonStartBlock, uint seasonEndBlock, uint blocksElapsed)
    {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getCurrentSeasonInfo();
    }

    /**
     * @notice Returns project information
     * @param projectID Project ID
     * @return stakingPool StakingPool address
     * @return rewardPool RewardPool address
     * @return name Project name
     * @return isActive Whether project is active
     * @return createdAt Creation timestamp
     */
    function getProjectInfo(uint projectID)
        external
        view
        returns (address stakingPool, address rewardPool, string memory name, bool isActive, uint createdAt)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), "InvalidProjectID");
        (stakingPool, rewardPool, name, isActive, createdAt,,) = protocol.projects(projectID);
    }

    /**
     * @notice Returns claimable reward for a user
     * @param projectID Project ID
     * @param user User address
     * @param season Season number
     * @param rewardToken Reward token address
     * @return claimableAmount Claimable reward amount
     */
    function getClaimableReward(uint projectID, address user, uint season, address rewardToken)
        external
        view
        returns (uint claimableAmount)
    {
        (, IRewardPool rewardPool) = _getPools(projectID);
        return rewardPool.getExpectedReward(user, season, rewardToken);
    }

    /**
     * @notice Returns total amount staked in a project
     * @param projectID Project ID
     * @return totalStaked Total staked amount
     */
    function getTotalStaked(uint projectID) external view returns (uint totalStaked) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.totalStaked();
    }

    /**
     * @notice Returns total points for a season (supports virtual seasons)
     * @param projectID Project ID
     * @param season Season number
     * @return totalPoints Total points in the season
     * @dev For virtual seasons (not rolled over), calculates points as if rolled over
     */
    function getSeasonTotalPoints(uint projectID, uint season) external view returns (uint totalPoints) {
        (IStakingPool pool,) = _getPools(projectID);
        uint currentSeason = pool.currentSeason();

        // Virtual season - calculate directly
        if (season > currentSeason) {
            (, totalPoints) = _calculateVirtualSeasonData(pool, season, address(0));
            return totalPoints;
        }

        // On-chain season
        totalPoints = pool.seasonTotalPointsSnapshot(season);

        // If 0 and past season, try virtual calculation
        if (totalPoints == 0 && season < currentSeason) {
            (, totalPoints) = _calculateVirtualSeasonData(pool, season, address(0));
        }

        return totalPoints;
    }

    /**
     * @notice Returns user points for a specific season (supports virtual seasons)
     * @param projectID Project ID
     * @param season Season number
     * @param user User address
     * @return userPoints User's points
     * @return totalPoints Total points
     * @dev Handles both on-chain and virtual seasons transparently
     */
    function getSeasonUserPoints(uint projectID, uint season, address user)
        external
        view
        returns (uint userPoints, uint totalPoints)
    {
        (IStakingPool pool,) = _getPools(projectID);
        uint currentSeason = pool.currentSeason();

        // Virtual season - calculate directly
        if (season > currentSeason) return _calculateVirtualSeasonData(pool, season, user);

        // On-chain season - delegate to pool
        (userPoints, totalPoints) = pool.getSeasonUserPoints(season, user);

        // If totalPoints = 0, try virtual calculation
        if (totalPoints == 0 && season < currentSeason) {
            (, totalPoints) = _calculateVirtualSeasonData(pool, season, address(0));
        }
    }

    /**
     * @notice Checks if a season is currently active
     * @param projectID Project ID
     * @return isActive True if season is active
     */
    function isSeasonActive(uint projectID) external view returns (bool isActive) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.isSeasonActive();
    }

    /**
     * @notice Returns pool end block
     * @param projectID Project ID
     * @return endBlock Pool end block (0 = infinite)
     */
    function getPoolEndBlock(uint projectID) external view returns (uint endBlock) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.poolEndBlock();
    }

    /**
     * @notice Returns next season start block
     * @param projectID Project ID
     * @return startBlock Next season start block
     */
    function getNextSeasonStartBlock(uint projectID) external view returns (uint startBlock) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.nextSeasonStartBlock();
    }

    /**
     * @notice Returns detailed user season data
     * @param projectID Project ID
     * @param user User address
     * @param season Season number
     * @return points User's points
     * @return balance User's balance
     * @return joinBlock Join block
     * @return claimed Whether claimed
     * @return finalized Whether finalized
     */
    function getUserSeasonData(uint projectID, address user, uint season)
        external
        view
        returns (uint points, uint balance, uint joinBlock, bool claimed, bool finalized)
    {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getUserSeasonData(season, user);
    }

    /**
     * @notice Previews claim information
     * @param projectID Project ID
     * @param user User address
     * @param season Season number
     * @param rewardToken Reward token address
     * @return userPoints User's points
     * @return totalPoints Total points
     * @return expectedReward Expected reward
     * @return alreadyClaimed Whether already claimed
     * @return canClaim Whether can claim
     */
    function previewClaim(uint projectID, address user, uint season, address rewardToken)
        external
        view
        returns (uint userPoints, uint totalPoints, uint expectedReward, bool alreadyClaimed, bool canClaim)
    {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.previewClaim(season, user, rewardToken);
    }

    /**
     * @notice Returns expected rewards for multiple seasons (batch query)
     * @param projectID Project ID
     * @param user User address
     * @param seasons Array of season numbers
     * @param rewardToken Reward token address
     * @return expectedRewards Array of expected reward amounts
     * @dev Gas-efficient batch query for multiple seasons at once
     */
    function getExpectedRewardsBatch(uint projectID, address user, uint[] calldata seasons, address rewardToken)
        external
        view
        returns (uint[] memory expectedRewards)
    {
        (IStakingPool pool,) = _getPools(projectID);
        expectedRewards = new uint[](seasons.length);
        for (uint i = 0; i < seasons.length;) {
            expectedRewards[i] = pool.getExpectedSeasonReward(seasons[i], user, rewardToken);
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Returns paginated list of active projects
     * @param offset Starting index
     * @param limit Maximum number to return
     * @return projectIDs Array of project IDs
     * @return stakingPools Array of StakingPool addresses
     * @return names Array of project names
     * @dev Filters only active projects from the full list
     */
    function getActiveProjects(uint offset, uint limit)
        external
        view
        returns (uint[] memory projectIDs, address[] memory stakingPools, string[] memory names)
    {
        uint totalProjects = protocol.projectCount();
        require(offset < totalProjects, "InvalidOffset");

        uint end = offset + limit;
        if (end > totalProjects) end = totalProjects;
        uint count = end - offset;
        projectIDs = new uint[](count);
        stakingPools = new address[](count);
        names = new string[](count);

        uint resultIndex = 0;
        for (uint i = offset; i < end;) {
            uint projectID = i + 1;
            (address stakingPool,, string memory name, bool isActive,,,) = protocol.projects(projectID);
            if (isActive) {
                projectIDs[resultIndex] = projectID;
                stakingPools[resultIndex] = stakingPool;
                names[resultIndex] = name;
                unchecked {
                    ++resultIndex;
                }
            }
            unchecked {
                ++i;
            }
        }

        if (resultIndex < count) {
            assembly {
                mstore(projectIDs, resultIndex)
                mstore(stakingPools, resultIndex)
                mstore(names, resultIndex)
            }
        }
    }

    /**
     * @notice Returns staking summary for a user across multiple projects
     * @param user User address
     * @param projectIDs Array of project IDs to query
     * @return balances Array of staked balances
     * @return points Array of current season points
     * @return currentSeasons Array of current season numbers
     * @dev Efficient batch query for dashboard displays
     */
    function getUserStakingSummary(address user, uint[] calldata projectIDs)
        external
        view
        returns (uint[] memory balances, uint[] memory points, uint[] memory currentSeasons)
    {
        balances = new uint[](projectIDs.length);
        points = new uint[](projectIDs.length);
        currentSeasons = new uint[](projectIDs.length);

        for (uint i = 0; i < projectIDs.length;) {
            if (projectIDs[i] > 0 && projectIDs[i] <= protocol.projectCount()) {
                (address stakingPool,,,,,,) = protocol.projects(projectIDs[i]);
                (balances[i], points[i],) = IStakingPool(stakingPool).getStakePosition(user);
                // Use getCurrentSeasonInfo to get calculated season (supports virtual seasons)
                (currentSeasons[i],,,) = IStakingPool(stakingPool).getCurrentSeasonInfo();
            }
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Returns historical season data
     * @param projectID Project ID
     * @param fromSeason Starting season number
     * @param toSeason Ending season number
     * @return seasons Array of season numbers
     * @return totalPoints Array of total points per season
     * @dev Useful for analytics and trend analysis
     */
    function getSeasonHistory(uint projectID, uint fromSeason, uint toSeason)
        external
        view
        returns (uint[] memory seasons, uint[] memory totalPoints)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), "InvalidProjectID");
        require(fromSeason <= toSeason, "InvalidRange");

        (IStakingPool pool,) = _getPools(projectID);
        uint currentSeason = pool.currentSeason();
        if (toSeason > currentSeason) toSeason = currentSeason;
        if (fromSeason > toSeason) return (new uint[](0), new uint[](0));

        uint count = toSeason - fromSeason + 1;
        seasons = new uint[](count);
        totalPoints = new uint[](count);
        for (uint i = 0; i < count;) {
            uint season = fromSeason + i;
            seasons[i] = season;
            totalPoints[i] = pool.seasonTotalPointsSnapshot(season);
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Returns all reward tokens used in a season
     * @param projectID Project ID
     * @param season Season number
     * @return tokens Array of token addresses
     */
    function getSeasonRewardTokens(uint projectID, uint season) external view returns (address[] memory tokens) {
        (, IRewardPool rewardPool) = _getPools(projectID);
        return rewardPool.getSeasonRewardTokens(season);
    }

    /**
     * @notice Returns reward token information for a season
     * @param projectID Project ID
     * @param season Season number
     * @param token Token address
     * @return total Total rewards
     * @return claimed Claimed rewards
     * @return remaining Remaining rewards
     */
    function getSeasonTokenInfo(uint projectID, uint season, address token)
        external
        view
        returns (uint total, uint claimed, uint remaining)
    {
        (, IRewardPool rewardPool) = _getPools(projectID);
        return rewardPool.getSeasonTokenInfo(season, token);
    }

    /**
     * @notice Returns all reward information for a season
     * @param projectID Project ID
     * @param season Season number
     * @return tokens Array of token addresses
     * @return totals Array of total rewards
     * @return claimeds Array of claimed rewards
     * @return remainings Array of remaining rewards
     */
    function getSeasonAllRewards(uint projectID, uint season)
        external
        view
        returns (address[] memory tokens, uint[] memory totals, uint[] memory claimeds, uint[] memory remainings)
    {
        (, IRewardPool rewardPool) = _getPools(projectID);
        return rewardPool.getSeasonAllRewards(season);
    }

    /**
     * @notice Returns expected points for a season (supports virtual/unfinalized seasons)
     * @param projectID Project ID
     * @param season Season number
     * @param user User address
     * @return userPoints User's expected points
     * @return totalPoints Total expected points
     * @dev Calculates points even for seasons not finalized or rolled over
     *      Useful for showing estimated data to users
     */
    function getExpectedSeasonPoints(uint projectID, uint season, address user)
        external
        view
        returns (uint userPoints, uint totalPoints)
    {
        (IStakingPool pool,) = _getPools(projectID);
        uint currentSeason = pool.currentSeason();

        // 가상 시즌이면 직접 계산
        if (season > currentSeason) return _calculateVirtualSeasonData(pool, season, user);

        // 시즌 범위 계산
        (bool ok, uint startBlock, uint endBlock) = _calcSeasonRange(pool, season);
        if (!ok) return (0, 0);
        if (block.number < startBlock) return (0, 0);

        // 과거 시즌: 항상 endBlock 사용 (종료된 시즌이므로 고정)
        // 현재 시즌: 현재 블록까지만 반영
        uint toBlock = (season < currentSeason) ? endBlock : (block.number < endBlock ? block.number : endBlock);

        if (toBlock <= startBlock) return (0, 0);

        // 유저 포인트 계산
        (uint balance,, uint lastUpdate) = pool.getStakePosition(user);
        if (balance == 0) return (0, 0);
        if (lastUpdate > toBlock) return (0, 0);

        uint fromBlock = lastUpdate > startBlock ? lastUpdate : startBlock;
        userPoints = PointsLib.calculatePoints(balance, fromBlock, toBlock, pool.blockTime(), pool.pointsTimeUnit());

        // 토탈 포인트: 온체인 값 우선, 없으면 계산
        totalPoints = pool.seasonTotalPointsSnapshot(season);
        if (totalPoints == 0) {
            // 가상 계산
            (, totalPoints) = _calculateVirtualSeasonData(pool, season, address(0));
        }

        return (userPoints, totalPoints);
    }

    /**
     * @notice Returns expected reward for a user in a season
     * @param projectID Project ID
     * @param season Season number
     * @param user User address
     * @param rewardToken Reward token address
     * @return expectedReward Expected reward amount
     * @dev Delegates to StakingPool for calculation
     */
    function getExpectedSeasonReward(uint projectID, uint season, address user, address rewardToken)
        external
        view
        returns (uint expectedReward)
    {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getExpectedSeasonReward(season, user, rewardToken);
    }

    /**
     * @notice Returns comprehensive pool configuration and state
     * @param projectID Project ID
     * @return blockTime Block time in seconds
     * @return pointsTimeUnit Points calculation time unit in seconds
     * @return seasonBlocks Number of blocks per season
     * @return poolEndBlock Pool end block (0 = infinite)
     * @return currentSeason Current season number
     * @return preDepositStartBlock Pre-deposit start block (0 = disabled)
     * @return firstSeasonStartBlock First season start block
     * @dev Includes pre-deposit information for frontend UI
     *      firstSeasonStartBlock is calculated from current season or nextSeasonStartBlock
     */
    function getPoolInfo(uint projectID)
        external
        view
        returns (
            uint blockTime,
            uint pointsTimeUnit,
            uint seasonBlocks,
            uint poolEndBlock,
            uint currentSeason,
            uint preDepositStartBlock,
            uint firstSeasonStartBlock
        )
    {
        (IStakingPool pool,) = _getPools(projectID);

        blockTime = pool.blockTime();
        pointsTimeUnit = pool.pointsTimeUnit();
        seasonBlocks = pool.seasonBlocks();
        poolEndBlock = pool.poolEndBlock();
        // Use getCurrentSeasonInfo to get calculated season (supports virtual seasons)
        (currentSeason,,,) = pool.getCurrentSeasonInfo();
        preDepositStartBlock = pool.preDepositStartBlock();

        // firstSeasonStartBlock 계산
        if (currentSeason > 0) {
            // 시즌이 시작된 경우: 현재 시즌 정보에서 역산
            (uint season, uint startBlock,,) = pool.getCurrentSeasonInfo();
            if (season > 0 && startBlock > 0) firstSeasonStartBlock = startBlock - ((season - 1) * seasonBlocks);
        } else {
            // 시즌 시작 전: nextSeasonStartBlock 사용
            firstSeasonStartBlock = pool.nextSeasonStartBlock();
        }

        return (
            blockTime,
            pointsTimeUnit,
            seasonBlocks,
            poolEndBlock,
            currentSeason,
            preDepositStartBlock,
            firstSeasonStartBlock
        );
    }
}
