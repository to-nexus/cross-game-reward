// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {ICrossStaking} from "./interfaces/ICrossStaking.sol";
import {ICrossStakingPool} from "./interfaces/ICrossStakingPool.sol";
import {ICrossStakingRouter} from "./interfaces/ICrossStakingRouter.sol";
import {IWCROSS} from "./interfaces/IWCROSS.sol";

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
    ICrossStaking public immutable crossStaking;

    /// @notice WCROSS token reference
    IWCROSS public immutable wcross;

    // ==================== Constructor ====================

    /**
     * @notice Initializes the CrossStakingRouter
     * @param _crossStaking Address of the CrossStaking contract
     */
    constructor(address _crossStaking) {
        require(_crossStaking != address(0), CSRCanNotZeroAddress());

        crossStaking = ICrossStaking(_crossStaking);
        wcross = IWCROSS(crossStaking.wcross());
    }

    // ==================== Native CROSS Staking ====================

    /**
     * @notice Stakes native CROSS tokens
     * @dev Wraps native CROSS to WCROSS and stakes into the pool
     * @param poolId ID of the pool to stake into
     */
    function stakeNative(uint poolId) external payable {
        require(msg.value > 0, CSRInvalidAmount());

        ICrossStakingPool pool = _getPoolAndValidateWCROSS(poolId);

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
        ICrossStakingPool pool = _getPoolAndValidateWCROSS(poolId);

        uint stakedAmount = pool.balances(msg.sender);
        require(stakedAmount > 0, CSRNoStakeFound());

        // Unstake from pool (WCROSS sent to Router, rewards sent to msg.sender)
        pool.unstakeFor(msg.sender);

        // Router unwraps and sends native CROSS directly to user
        wcross.withdrawTo(stakedAmount, msg.sender);

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

        ICrossStakingPool pool = _getPool(poolId);
        IERC20 stakingToken = pool.stakingToken();

        _stakeERC20(poolId, pool, stakingToken, amount);
    }

    /**
     * @notice Stakes ERC20 tokens using EIP-2612 permit
     * @dev Performs permit + transfer + stake in a single transaction (for tokens supporting EIP-2612)
     * @param poolId ID of the pool to stake into
     * @param amount Amount of tokens to stake
     * @param deadline Permit signature deadline
     * @param v Permit signature v
     * @param r Permit signature r
     * @param s Permit signature s
     */
    function stakeERC20WithPermit(uint poolId, uint amount, uint deadline, uint8 v, bytes32 r, bytes32 s) external {
        require(amount > 0, CSRInvalidAmount());

        ICrossStakingPool pool = _getPool(poolId);
        IERC20 stakingToken = pool.stakingToken();

        // Approve Router via EIP-2612 permit
        IERC20Permit(address(stakingToken)).permit(msg.sender, address(this), amount, deadline, v, r, s);

        _stakeERC20(poolId, pool, stakingToken, amount);
    }

    /**
     * @notice Unstakes ERC20 tokens
     * @dev Unstakes all staked tokens and claims rewards
     * @param poolId ID of the pool to unstake from
     */
    function unstakeERC20(uint poolId) external {
        ICrossStakingPool pool = _getPool(poolId);
        IERC20 stakingToken = pool.stakingToken();

        uint stakedAmount = pool.balances(msg.sender);
        require(stakedAmount > 0, CSRNoStakeFound());

        // Unstake (staking tokens sent to Router, rewards to msg.sender)
        pool.unstakeFor(msg.sender);

        // Transfer staking tokens from Router to user
        stakingToken.safeTransfer(msg.sender, stakedAmount);

        emit UnstakedERC20(msg.sender, poolId, address(stakingToken), stakedAmount);
    }

    // ==================== View Functions ====================

    /**
     * @notice Retrieves user's staking information
     * @param poolId ID of the pool
     * @param user Address of the user
     * @return stakedAmount Amount of tokens staked
     * @return rewardTokens Array of reward token addresses
     * @return pendingRewards Array of pending rewards for each reward token
     */
    function getUserStakingInfo(uint poolId, address user)
        external
        view
        returns (uint stakedAmount, address[] memory rewardTokens, uint[] memory pendingRewards)
    {
        ICrossStakingPool pool = _getPool(poolId);
        stakedAmount = pool.balances(user);
        (rewardTokens, pendingRewards) = pool.pendingRewards(user);
    }

    /**
     * @notice Checks if a pool is a native CROSS pool
     * @param poolId ID of the pool
     * @return True if the pool uses WCROSS as staking token
     */
    function isNativePool(uint poolId) external view returns (bool) {
        ICrossStakingPool pool = _getPool(poolId);
        return address(pool.stakingToken()) == address(wcross);
    }

    // ==================== Internal Functions ====================

    /**
     * @dev Retrieves pool address and returns pool instance
     * @param poolId ID of the pool
     * @return Pool contract instance
     */
    function _getPool(uint poolId) internal view returns (ICrossStakingPool) {
        return crossStaking.getPoolAddress(poolId);
    }

    /**
     * @dev Retrieves pool and validates it's a WCROSS pool
     * @param poolId ID of the pool
     * @return pool Pool contract instance
     */
    function _getPoolAndValidateWCROSS(uint poolId) internal view returns (ICrossStakingPool pool) {
        pool = _getPool(poolId);
        require(address(pool.stakingToken()) == address(wcross), CSRNotWCROSSPool());
    }

    /**
     * @dev Stakes ERC20 tokens from user and stakes into the pool
     * @param poolId ID of the pool to stake into
     * @param pool Pool contract instance
     * @param token ERC20 token contract instance
     * @param amount Amount of tokens to stake
     */
    function _stakeERC20(uint poolId, ICrossStakingPool pool, IERC20 token, uint amount) internal {
        // Transfer tokens from user and stake to pool
        token.safeTransferFrom(msg.sender, address(this), amount);
        token.forceApprove(address(pool), amount);
        pool.stakeFor(msg.sender, amount);

        emit StakedERC20(msg.sender, poolId, address(token), amount);
    }
}
