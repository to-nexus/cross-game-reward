// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title IRewardPoolCode
 * @notice Interface for retrieving RewardPool's creation bytecode
 * @dev Used in the Code Contract pattern to deploy RewardPool instances via factory
 */
interface IRewardPoolCode {
    /// @notice Returns the creation bytecode of RewardPool contract
    function code() external pure returns (bytes memory);
}

/**
 * @title IRewardPool
 * @notice Interface for season-based reward distribution pools
 * @dev Each StakingPool has a connected RewardPool that:
 *      - Holds reward tokens for each season
 *      - Distributes rewards proportionally based on points
 *      - Prevents double-claiming through claim tracking
 *      - Supports multiple reward tokens per season
 */
interface IRewardPool {
    // ============ Events ============

    /// @notice Emitted when rewards are deposited for a season
    event SeasonFunded(uint indexed season, address indexed token, uint amount, uint actualReceived);

    /// @notice Emitted when a user claims rewards
    event RewardPaid(address indexed user, uint indexed season, address indexed token, uint amount);

    // ============ Season-based Rewards ============

    /**
     * @notice Deposits reward tokens for a specific season
     * @param season Season number to fund
     * @param token Reward token address
     * @param amount Amount of tokens to deposit
     * @dev Can be called multiple times to add more rewards to the same season/token
     */
    function fundSeason(uint season, address token, uint amount) external;

    /**
     * @notice Pays out rewards to a user (called by StakingPool only)
     * @param user User to receive rewards
     * @param season Season number to claim from
     * @param token Reward token address
     * @param userPoints User's points in the season
     * @param totalPoints Total points in the season
     * @dev Uses proportional distribution: reward = totalReward × (userPoints / totalPoints)
     */
    function payUser(address user, uint season, address token, uint userPoints, uint totalPoints) external;

    // ============ View Functions ============

    /// @notice Returns unclaimed rewards remaining for a season/token
    function getRemainingRewards(uint season, address token) external view returns (uint);

    /// @notice Calculates expected reward for a user in a season
    function getExpectedReward(address user, uint season, address token) external view returns (uint);

    // Season reward token information

    /// @notice Returns list of all reward tokens used in a season
    function getSeasonRewardTokens(uint season) external view returns (address[] memory tokens);

    /// @notice Returns reward information for a specific season/token combination
    function getSeasonTokenInfo(uint season, address token)
        external
        view
        returns (uint total, uint claimed, uint remaining);

    /// @notice Returns all reward information for a season
    function getSeasonAllRewards(uint season)
        external
        view
        returns (address[] memory tokens, uint[] memory totals, uint[] memory claimeds, uint[] memory remainings);

    // ============ Convenience Functions ============

    /// @notice Deposits rewards for the current season (wrapper for fundSeason)
    function depositReward(address token, uint amount) external;

    /// @notice Returns pending rewards for a user in the current season
    function getUserPendingReward(address user, address token) external view returns (uint);

    // ============ Admin Functions ============

    /// @notice Emergency function to recover tokens (protocol only)
    function sweep(address token, address to, uint amount) external;
}
