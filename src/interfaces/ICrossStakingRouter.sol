// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title ICrossStakingRouter
 * @notice Interface for the CrossStakingRouter contract
 * @dev Defines functions for staking native CROSS and ERC20 tokens
 */
interface ICrossStakingRouter {
    /// @notice Stakes native CROSS tokens
    function stakeNative(uint poolId) external payable;

    /// @notice Unstakes and returns native CROSS tokens
    function unstakeNative(uint poolId) external;

    /// @notice Stakes ERC20 tokens
    function stakeERC20(uint poolId, uint amount) external;

    /// @notice Unstakes ERC20 tokens
    function unstakeERC20(uint poolId) external;

    /// @notice Retrieves user's staking information
    function getUserStakingInfo(uint poolId, address user)
        external
        view
        returns (uint stakedAmount, address[] memory rewardTokens, uint[] memory pendingRewards);

    /// @notice Checks if a pool is a native CROSS pool
    function isNativePool(uint poolId) external view returns (bool);
}
