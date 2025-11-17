// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title ICrossGameRewardRouter
 * @notice Interface for the CrossGameRewardRouter contract
 * @dev Defines functions for depositing, withdrawing, claiming and querying rewards
 */
interface ICrossGameRewardRouter {
    /// @notice Deposits native CROSS tokens
    function depositNative(uint poolId) external payable;

    /// @notice Withdraws and returns native CROSS tokens
    function withdrawNative(uint poolId) external;

    /// @notice Deposits ERC20 tokens
    function depositERC20(uint poolId, uint amount) external;

    /// @notice Withdraws ERC20 tokens
    function withdrawERC20(uint poolId) external;

    /// @notice Claims all pending rewards from a pool
    function claimRewards(uint poolId) external;

    /// @notice Claims a specific reward token from a pool
    function claimReward(uint poolId, address token) external;

    /// @notice Retrieves user's deposit information
    function getUserDepositInfo(uint poolId, address user)
        external
        view
        returns (uint depositedAmount, address[] memory rewardTokens, uint[] memory pendingRewards);

    /// @notice Checks if a pool is a native CROSS pool
    function isNativePool(uint poolId) external view returns (bool);

    /// @notice Retrieves user's pending rewards for active reward tokens
    function getPendingRewards(uint poolId, address user)
        external
        view
        returns (address[] memory rewardTokens, uint[] memory pendingRewards);

    /// @notice Retrieves user's pending reward for a specific token
    function getPendingReward(uint poolId, address user, address token) external view returns (uint amount);

    /// @notice Retrieves pending rewards for removed reward tokens
    function getRemovedTokenRewards(uint poolId, address user)
        external
        view
        returns (address[] memory rewardTokens, uint[] memory pendingRewards);

    /// @notice Retrieves all pending rewards (active + removed) filtered by amount > 0
    function getAllPendingRewards(uint poolId, address user)
        external
        view
        returns (address[] memory rewardTokens, uint[] memory pendingRewards);
}
