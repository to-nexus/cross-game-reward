// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC5313} from "@openzeppelin/contracts/interfaces/IERC5313.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title ICrossStakingPool
 * @notice Interface for the CrossStakingPool contract
 * @dev Defines the structure and functions for staking and reward management
 *      Implements IERC5313 for standard owner() function
 */
interface ICrossStakingPool is IERC5313 {
    /**
     * @notice Pool status enum
     * @param Active All operations allowed (stake, unstake, claim)
     * @param Inactive Only unstake and claim allowed
     * @param Paused All operations stopped
     */
    enum PoolStatus {
        Active,
        Inactive,
        Paused
    }

    /**
     * @notice Reward token information structure
     * @param token Address of the reward token
     * @param rewardPerTokenStored Accumulated reward per token staked
     * @param lastBalance Last known balance of the reward token
     * @param withdrawableAmount Amount that owner can withdraw (from zero-stake deposits)
     * @param distributedAmount Amount distributed to users (for removed tokens, tracks claimable balance)
     * @param isRemoved Whether the reward token has been removed
     */
    struct RewardToken {
        IERC20 token;
        uint rewardPerTokenStored;
        uint lastBalance;
        uint withdrawableAmount;
        uint distributedAmount;
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

    /// @notice Returns pending rewards for a user across all active reward tokens
    function pendingRewards(address account) external view returns (address[] memory tokens, uint[] memory rewards);

    /// @notice Returns pending reward for a specific token
    function pendingReward(address account, IERC20 token) external view returns (uint amount);

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

    /// @notice Returns all removed reward token addresses
    function getRemovedRewardTokens() external view returns (address[] memory);

    /// @notice Returns the number of removed reward tokens
    function removedRewardTokenCount() external view returns (uint);

    /// @notice Checks if a token is a removed reward token
    function isRemovedRewardToken(IERC20 token) external view returns (bool);

    /// @notice Returns user's claimable rewards for removed tokens
    function getRemovedTokenRewards(address user)
        external
        view
        returns (address[] memory tokens, uint[] memory rewards);

    /// @notice Adds a reward token to the pool
    function addRewardToken(IERC20 token) external;

    /// @notice Removes a reward token from the pool
    function removeRewardToken(IERC20 token) external;

    /// @notice Sets the pool status
    function setPoolStatus(PoolStatus status) external;

    /// @notice Returns the withdrawable amount for a removed reward token
    function getWithdrawableAmount(IERC20 token) external view returns (uint);

    /// @notice Withdraws undistributed rewards that were deposited after token removal
    function withdraw(IERC20 token, address to) external;

    /// @notice Updates the minimum stake amount for the pool
    function updateMinStakeAmount(uint amount) external;
}
