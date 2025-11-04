// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title ICrossStakingPool
 * @notice Interface for the CrossStakingPool contract
 * @dev Defines the structure and functions for staking and reward management
 */
interface ICrossStakingPool {
    /**
     * @notice Reward token information structure
     * @param token Address of the reward token
     * @param rewardPerTokenStored Accumulated reward per token staked
     * @param lastBalance Last known balance of the reward token
     * @param removedDistributedAmount Amount allocated for distribution when token was removed
     * @param isRemoved Whether the reward token has been removed
     */
    struct RewardToken {
        IERC20 token;
        uint rewardPerTokenStored;
        uint lastBalance;
        uint removedDistributedAmount;
        bool isRemoved;
    }

    /**
     * @notice User reward information structure
     * @param rewardPerTokenPaid Last recorded rewardPerToken for the user
     * @param rewards Accumulated claimable rewards
     */
    struct UserReward {
        uint rewardPerTokenPaid;
        uint rewards;
    }

    /// @notice Returns the staking token
    function stakingToken() external view returns (IERC20);

    /// @notice Returns the staked balance of a user
    function balances(address account) external view returns (uint);

    /// @notice Returns the total amount staked in the pool
    function totalStaked() external view returns (uint);

    /// @notice Returns the reward information for a user and token
    function userRewards(address account, IERC20 token) external view returns (uint rewardPerTokenPaid, uint rewards);

    /// @notice Stakes tokens into the pool
    function stake(uint amount) external;

    /// @notice Stakes tokens on behalf of another account
    function stakeFor(address account, uint amount) external;

    /// @notice Unstakes all staked tokens and claims rewards
    function unstake() external;

    /// @notice Unstakes tokens on behalf of another account
    function unstakeFor(address account) external;

    /// @notice Claims all pending rewards
    function claimRewards() external;

    /// @notice Claims pending rewards for a specific token
    function claimReward(IERC20 token) external;

    /// @notice Returns pending rewards for a user
    function pendingRewards(address account) external view returns (uint[] memory);

    /// @notice Returns reward token address at a specific index
    function rewardTokenAt(uint index) external view returns (IERC20);

    /// @notice Returns reward token information
    function getRewardToken(IERC20 token) external view returns (RewardToken memory);

    /// @notice Checks if a token is a registered reward token
    function isRewardToken(IERC20 token) external view returns (bool);

    /// @notice Returns all reward token addresses
    function getRewardTokens() external view returns (address[] memory);

    /// @notice Returns the number of reward tokens
    function rewardTokenCount() external view returns (uint);

    /// @notice Adds a reward token to the pool
    function addRewardToken(IERC20 token) external;

    /// @notice Removes a reward token from the pool
    function removeRewardToken(IERC20 token) external;

    /// @notice Pauses the pool
    function pause() external;

    /// @notice Unpauses the pool
    function unpause() external;
}
