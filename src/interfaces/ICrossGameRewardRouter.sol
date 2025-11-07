// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title ICrossGameRewardRouter
 * @notice Interface for the CrossGameRewardRouter contract
 * @dev Defines functions for depositing native CROSS and ERC20 tokens
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

    /// @notice Retrieves user's deposit information
    function getUserDepositInfo(uint poolId, address user)
        external
        view
        returns (uint depositedAmount, address[] memory rewardTokens, uint[] memory pendingRewards);

    /// @notice Checks if a pool is a native CROSS pool
    function isNativePool(uint poolId) external view returns (bool);
}
