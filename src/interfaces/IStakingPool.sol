// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./IRewardPool.sol";

/**
 * @title IStakingPoolCode
 * @notice Interface for retrieving StakingPool's creation bytecode
 * @dev Used in the Code Contract pattern to deploy StakingPool instances via factory
 */
interface IStakingPoolCode {
    /// @notice Returns the creation bytecode of StakingPool contract
    function code() external pure returns (bytes memory);
}

/**
 * @title IStakingPool
 * @notice Interface for project-specific staking pools
 * @dev Each project has its own StakingPool instance managing:
 *      - $CROSS token staking/unstaking
 *      - Season-based points accumulation
 *      - Reward claims through connected RewardPool
 *      - Automatic season rollovers (up to 50 seasons at once)
 */
interface IStakingPool {
    // ============ Staking Functions ============

    /// @notice Stakes tokens for msg.sender
    function stake(uint amount) external;

    /// @notice Stakes tokens on behalf of a user (router only)
    function stakeFor(address user, uint amount) external;

    /// @notice Withdraws all staked tokens for msg.sender
    function withdrawAll() external;

    /// @notice Withdraws all staked tokens on behalf of a user (router only)
    function withdrawAllFor(address user) external;

    // ============ Season Management ============

    /// @notice Manually triggers season rollover (automatic in most cases)
    function rolloverSeason() external;

    /// @notice Claims rewards for a specific season
    function claimSeason(uint season, address rewardToken) external;

    /// @notice Claims rewards on behalf of a user (router only)
    function claimSeasonFor(address user, uint season, address rewardToken) external;

    // ============ Points ============

    /// @notice Updates a user's points to current block
    function updatePoints(address user) external;

    // ============ Admin Functions ============

    /// @notice Sets the connected reward pool (one-time only)
    function setRewardPool(IRewardPool rewardPool) external;

    /// @notice Approves/revokes a router for proxy staking
    function setApprovedRouter(address router, bool approved) external;

    /// @notice Sets the time unit for points calculation (e.g., 3600 for 1 hour)

    /// @notice Sets the block time in seconds (e.g., 1 for 1 second/block)

    /// @notice Sets the start block for the next season
    function setNextSeasonStart(uint startTime) external;

    /// @notice Sets the pool end block (0 = infinite)
    function setPoolEndBlock(uint endTime) external;

    /// @notice Manually rolls over multiple seasons (for catching up)
    function manualRolloverSeasons(uint maxRollovers) external returns (uint rolloversPerformed);

    /// @notice Finalizes user's season snapshots in batch (gas optimization)
    function finalizeUserSeasonsBatch(address user, uint maxSeasons) external returns (uint processed);

    // ============ View Functions ============

    /// @notice Returns user's staking power (same as balance)
    function getStakingPower(address user) external view returns (uint);

    /// @notice Returns total staking power (same as totalStaked)
    function getTotalStakingPower() external view returns (uint);

    /// @notice Returns user's current season points (real-time calculation)
    function getUserPoints(address user) external view returns (uint);

    /// @notice Returns user's stake position details
    function getStakePosition(address user) external view returns (uint balance, uint points, uint lastUpdateBlock);

    /// @notice Returns current season information (handles virtual seasons)
    function getCurrentSeasonInfo()
        external
        view
        returns (uint season, uint startTime, uint endTime, uint timeElapsed);

    /// @notice Returns user's points for a specific season
    function getSeasonUserPoints(uint season, address user) external view returns (uint userPoints, uint totalPoints);

    /// @notice Returns finalized total points for a season
    function seasonTotalPointsSnapshot(uint season) external view returns (uint);

    /// @notice Returns current season number
    function currentSeason() external view returns (uint);

    /// @notice Returns total amount currently staked
    function totalStaked() external view returns (uint);

    /// @notice Returns expected reward for a user in a season
    function getExpectedSeasonReward(uint season, address user, address rewardToken) external view returns (uint);

    /// @notice Checks if a season is currently active
    function isSeasonActive() external view returns (bool);

    /// @notice Returns the pool end block (0 = infinite)
    function poolEndTime() external view returns (uint);

    /// @notice Returns the next season start block
    function nextSeasonStartTime() external view returns (uint);

    /// @notice Returns the pre-deposit start block (0 = disabled)
    function preDepositStartTime() external view returns (uint);

    /// @notice Returns number of pending season rollovers
    function getPendingSeasonRollovers() external view returns (uint pendingSeasons);

    /// @notice Returns user's season data (for lazy snapshot system)
    /// @dev claimed status is tracked per-token in RewardPool.hasClaimedSeasonReward
    function getUserSeasonData(uint season, address user)
        external
        view
        returns (uint points, uint balance, uint joinBlock, bool finalized);

    /// @notice Previews claim information for a user
    function previewClaim(uint season, address user, address rewardToken)
        external
        view
        returns (uint userPoints, uint totalPoints, uint expectedReward, bool alreadyClaimed, bool canClaim);

    // ============ Configuration View Functions ============

    /// @notice Returns block time in seconds

    /// @notice Returns points calculation time unit in seconds

    /// @notice Returns number of blocks per season
    function seasonDuration() external view returns (uint);
}
