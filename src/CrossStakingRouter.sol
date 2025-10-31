// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {CrossStaking} from "./CrossStaking.sol";
import {CrossStakingPool} from "./CrossStakingPool.sol";
import {IWCROSS} from "./interfaces/IWCROSS.sol";

import {ICrossStakingRouter} from "./interfaces/ICrossStakingRouter.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title CrossStakingRouter
 * @notice Interface between users and staking pools
 * @dev Handles native CROSS wrapping and staking operations
 *
 * Key Features:
 * - Wraps native CROSS to WCROSS and stakes
 * - Unstakes WCROSS and returns native CROSS
 * - Supports general ERC20 token staking
 */
contract CrossStakingRouter is ICrossStakingRouter {
    using SafeERC20 for IERC20;

    // ==================== Errors ====================

    /// @notice Thrown when an invalid amount is provided
    error CSRInvalidAmount();

    /// @notice Thrown when a zero address is provided where it's not allowed
    error CSRCanNotZeroAddress();

    /// @notice Thrown when a transfer fails
    error CSRTransferFailed();

    /// @notice Thrown when accessing a non-existent pool
    error CSRPoolNotFound();

    /// @notice Thrown when attempting native operations on a non-WCROSS pool
    error CSRNotWCROSSPool();

    /// @notice Thrown when attempting to unstake with no active stake
    error CSRNoStakeFound();

    // ==================== Events ====================

    /// @notice Emitted when a user stakes native CROSS
    /// @param user Address of the user who staked
    /// @param poolId ID of the pool staked into
    /// @param amount Amount of native CROSS staked
    event StakedNative(address indexed user, uint indexed poolId, uint amount);

    /// @notice Emitted when a user unstakes to native CROSS
    /// @param user Address of the user who unstaked
    /// @param poolId ID of the pool unstaked from
    /// @param amount Amount of native CROSS unstaked
    event UnstakedNative(address indexed user, uint indexed poolId, uint amount);

    /// @notice Emitted when a user stakes ERC20 tokens
    /// @param user Address of the user who staked
    /// @param poolId ID of the pool staked into
    /// @param token Address of the staked token
    /// @param amount Amount of tokens staked
    event StakedERC20(address indexed user, uint indexed poolId, address token, uint amount);

    /// @notice Emitted when a user unstakes ERC20 tokens
    /// @param user Address of the user who unstaked
    /// @param poolId ID of the pool unstaked from
    /// @param token Address of the unstaked token
    /// @param amount Amount of tokens unstaked
    event UnstakedERC20(address indexed user, uint indexed poolId, address token, uint amount);

    // ==================== State Variables ====================

    /// @notice CrossStaking contract reference
    CrossStaking public immutable crossStaking;

    /// @notice WCROSS token reference
    IWCROSS public immutable wcross;

    // ==================== Constructor ====================

    /**
     * @notice Initializes the CrossStakingRouter
     * @param _crossStaking Address of the CrossStaking contract
     */
    constructor(address _crossStaking) {
        require(_crossStaking != address(0), CSRCanNotZeroAddress());

        crossStaking = CrossStaking(_crossStaking);
        wcross = IWCROSS(crossStaking.wcross());
    }

    // ==================== Receive Function ====================

    /**
     * @notice Receives native CROSS
     * @dev Required for unstaking native CROSS
     */
    receive() external payable {}

    // ==================== Native CROSS Staking ====================

    /**
     * @notice Stakes native CROSS tokens
     * @dev Wraps native CROSS to WCROSS and stakes into the pool
     * @param poolId ID of the pool to stake into
     */
    function stakeNative(uint poolId) external payable {
        require(msg.value > 0, CSRInvalidAmount());

        CrossStakingPool pool = _getPoolAndValidateWCROSS(poolId);

        // Router wraps native CROSS to WCROSS
        wcross.deposit{value: msg.value}();

        // Router stakes to pool on behalf of msg.sender
        IERC20(address(wcross)).forceApprove(address(pool), msg.value);
        pool.stakeFor(msg.sender, msg.value);

        emit StakedNative(msg.sender, poolId, msg.value);
    }

    /**
     * @notice Unstakes and returns native CROSS tokens
     * @dev Unstakes WCROSS from pool and unwraps to native CROSS
     * @param poolId ID of the pool to unstake from
     */
    function unstakeNative(uint poolId) external {
        CrossStakingPool pool = _getPoolAndValidateWCROSS(poolId);

        uint stakedAmount = pool.balances(msg.sender);
        require(stakedAmount > 0, CSRNoStakeFound());

        // Unstake from pool (WCROSS + rewards sent to msg.sender)
        pool.unstakeFor(msg.sender);

        // Router takes msg.sender's WCROSS and unwraps
        uint wcrossBalance = IERC20(address(wcross)).balanceOf(msg.sender);
        require(wcrossBalance >= stakedAmount, CSRTransferFailed());

        IERC20(address(wcross)).safeTransferFrom(msg.sender, address(this), wcrossBalance);
        wcross.withdraw(wcrossBalance);

        // Send native CROSS to msg.sender
        (bool success,) = msg.sender.call{value: wcrossBalance}("");
        require(success, CSRTransferFailed());

        emit UnstakedNative(msg.sender, poolId, stakedAmount);
    }

    // ==================== ERC20 Staking (General Tokens) ====================

    /**
     * @notice Stakes ERC20 tokens
     * @dev Transfers tokens from user and stakes into the pool
     * @param poolId ID of the pool to stake into
     * @param amount Amount of tokens to stake
     */
    function stakeERC20(uint poolId, uint amount) external {
        require(amount > 0, CSRInvalidAmount());

        CrossStakingPool pool = _getPool(poolId);
        IERC20 stakingToken = pool.stakingToken();

        // Transfer tokens from msg.sender and stake to pool
        stakingToken.safeTransferFrom(msg.sender, address(this), amount);
        stakingToken.forceApprove(address(pool), amount);
        pool.stakeFor(msg.sender, amount);

        emit StakedERC20(msg.sender, poolId, address(stakingToken), amount);
    }

    /**
     * @notice Unstakes ERC20 tokens
     * @dev Unstakes all staked tokens and claims rewards
     * @param poolId ID of the pool to unstake from
     */
    function unstakeERC20(uint poolId) external {
        CrossStakingPool pool = _getPool(poolId);

        uint stakedAmount = pool.balances(msg.sender);
        require(stakedAmount > 0, CSRNoStakeFound());

        // Unstake (includes rewards)
        pool.unstakeFor(msg.sender);

        emit UnstakedERC20(msg.sender, poolId, address(pool.stakingToken()), stakedAmount);
    }

    // ==================== View Functions ====================

    /**
     * @notice Retrieves user's staking information
     * @param poolId ID of the pool
     * @param user Address of the user
     * @return stakedAmount Amount of tokens staked
     * @return pendingRewards Array of pending rewards for each reward token
     */
    function getUserStakingInfo(uint poolId, address user)
        external
        view
        returns (uint stakedAmount, uint[] memory pendingRewards)
    {
        CrossStakingPool pool = _getPool(poolId);
        stakedAmount = pool.balances(user);
        pendingRewards = pool.pendingRewards(user);
    }

    /**
     * @notice Checks if a pool is a native CROSS pool
     * @param poolId ID of the pool
     * @return True if the pool uses WCROSS as staking token
     */
    function isNativePool(uint poolId) external view returns (bool) {
        CrossStakingPool pool = _getPool(poolId);
        return address(pool.stakingToken()) == address(wcross);
    }

    // ==================== Internal Functions ====================

    /**
     * @dev Retrieves pool address and returns pool instance
     * @param poolId ID of the pool
     * @return Pool contract instance
     */
    function _getPool(uint poolId) internal view returns (CrossStakingPool) {
        return CrossStakingPool(crossStaking.getPoolAddress(poolId));
    }

    /**
     * @dev Retrieves pool and validates it's a WCROSS pool
     * @param poolId ID of the pool
     * @return pool Pool contract instance
     */
    function _getPoolAndValidateWCROSS(uint poolId) internal view returns (CrossStakingPool pool) {
        pool = _getPool(poolId);
        require(address(pool.stakingToken()) == address(wcross), CSRNotWCROSSPool());
    }
}
