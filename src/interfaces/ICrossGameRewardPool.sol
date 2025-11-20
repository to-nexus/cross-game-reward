// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {IERC5313} from "@openzeppelin/contracts/interfaces/IERC5313.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title ICrossGameRewardPool
 * @notice Interface for the CrossGameRewardPool contract
 * @dev Defines the structure and functions for deposit and reward management
 *      Implements IERC5313 for standard owner() function
 */
interface ICrossGameRewardPool is IERC5313 {
    /**
     * @notice Pool status enum
     * @param Active All operations allowed (deposit, withdraw, claim)
     * @param Inactive Only withdraw and claim allowed
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
     * @param rewardPerTokenStored Accumulated reward per token deposited
     * @param lastBalance Last known balance of the reward token
     * @param reclaimableAmount Amount that owner can reclaim (from zero-deposit rewards)
     * @param distributedAmount Amount distributed to users (for removed tokens, tracks claimable balance)
     * @param isRemoved Whether the reward token has been removed
     */
    struct RewardToken {
        IERC20 token;
        uint rewardPerTokenStored;
        uint lastBalance;
        uint reclaimableAmount;
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

    /// @notice Returns the pool status
    function poolStatus() external view returns (PoolStatus);

    /// @notice Returns the deposit token
    function depositToken() external view returns (IERC20);

    /// @notice Returns the deposited balance of a user
    function balances(address account) external view returns (uint);

    /// @notice Returns the total amount deposited in the pool
    function totalDeposited() external view returns (uint);

    /// @notice Returns the reward information for a user and token
    function userRewards(address account, IERC20 token) external view returns (uint rewardPerTokenPaid, uint rewards);

    /// @notice Deposits tokens into the pool
    function deposit(uint amount) external;

    /// @notice Deposits tokens on behalf of another account
    function depositFor(address account, uint amount) external;

    /// @notice Withdraws all deposited tokens and claims rewards
    function withdraw() external;

    /// @notice Withdraws tokens on behalf of another account
    function withdrawFor(address account) external;

    /// @notice Claims all pending rewards
    function claimRewards() external;

    /// @notice Claims all pending rewards on behalf of another account
    function claimRewardsFor(address account) external;

    /// @notice Claims pending rewards for a specific token
    function claimReward(IERC20 token) external;

    /// @notice Claims pending reward for a specific token on behalf of another account
    function claimRewardFor(address account, IERC20 token) external;

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

    /// @notice Returns the reclaimable amount for a removed reward token
    function getReclaimableAmount(IERC20 token) external view returns (uint);

    /// @notice Reclaims undistributed rewards that were deposited after token removal
    function reclaimTokens(IERC20 token, address to) external;

    /// @notice Updates the minimum deposit amount for the pool
    function updateMinDepositAmount(uint amount) external;
}
