// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {ICrossGameReward} from "./interfaces/ICrossGameReward.sol";
import {ICrossGameRewardPool} from "./interfaces/ICrossGameRewardPool.sol";
import {ICrossGameRewardRouter} from "./interfaces/ICrossGameRewardRouter.sol";
import {IWCROSS} from "./interfaces/IWCROSS.sol";

/**
 * @title CrossGameRewardRouter
 * @notice Interface between users and game reward pools
 * @dev Handles native CROSS wrapping and deposit operations
 *
 * Key Features:
 * - Wraps native CROSS to WCROSS and deposits
 * - Withdraws WCROSS and returns native CROSS
 * - Supports general ERC20 token deposits
 */
contract CrossGameRewardRouter is ICrossGameRewardRouter {
    using SafeERC20 for IERC20;

    // ==================== Errors ====================

    /// @notice Thrown when an invalid amount is provided
    /// @param provided The amount provided
    error CSRInvalidAmount(uint provided);

    /// @notice Thrown when a zero address is provided where it's not allowed
    error CSRCanNotZeroAddress();

    /// @notice Thrown when attempting native operations on a non-WCROSS pool
    /// @param poolId The pool ID
    /// @param actualToken The actual deposit token of the pool
    error CSRNotWCROSSPool(uint poolId, address actualToken);

    // ==================== Events ====================

    /// @notice Emitted when a user deposits native CROSS
    /// @param user Address of the user who deposited
    /// @param poolId ID of the pool deposited into
    /// @param amount Amount of native CROSS deposited
    event DepositedNative(address indexed user, uint indexed poolId, uint amount);

    /// @notice Emitted when a user withdraws to native CROSS
    /// @param user Address of the user who withdrew
    /// @param poolId ID of the pool withdrawn from
    /// @param amount Amount of native CROSS withdrawn
    event WithdrawnNative(address indexed user, uint indexed poolId, uint amount);

    /// @notice Emitted when a user deposits ERC20 tokens
    /// @param user Address of the user who deposited
    /// @param poolId ID of the pool deposited into
    /// @param token Address of the deposited token
    /// @param amount Amount of tokens deposited
    event DepositedERC20(address indexed user, uint indexed poolId, address token, uint amount);

    /// @notice Emitted when a user withdraws ERC20 tokens
    /// @param user Address of the user who withdrew
    /// @param poolId ID of the pool withdrawn from
    /// @param token Address of the withdrawn token
    /// @param amount Amount of tokens withdrawn
    event WithdrawnERC20(address indexed user, uint indexed poolId, address token, uint amount);

    // ==================== State Variables ====================

    /// @notice Native token placeholder address (use this to represent native CROSS)
    address public constant NATIVE_TOKEN = address(0x1);

    /// @notice CrossGameReward contract reference
    ICrossGameReward public immutable crossGameReward;

    /// @notice WCROSS token reference
    IWCROSS public immutable wcross;

    // ==================== Constructor ====================

    /**
     * @notice Initializes the CrossGameRewardRouter
     * @param _crossGameReward Address of the CrossGameReward contract
     */
    constructor(address _crossGameReward) {
        require(_crossGameReward != address(0), CSRCanNotZeroAddress());

        crossGameReward = ICrossGameReward(_crossGameReward);
        wcross = crossGameReward.wcross();
    }

    // ==================== Withdraw All ====================

    /**
     * @notice Withdraws all deposits from all pools
     * @dev Iterates through all pools and withdraws user's deposits
     *      Native pools return CROSS, ERC20 pools return their deposit tokens
     */
    function withdrawAll() external {
        uint[] memory poolIds = crossGameReward.getAllPoolIds();
        for (uint i = 0; i < poolIds.length; i++) {
            uint poolId = poolIds[i];
            ICrossGameRewardPool pool = _getPool(poolId);
            if (pool.balances(msg.sender) > 0) {
                if (address(pool.depositToken()) == address(wcross)) withdrawNative(poolId);
                else withdrawERC20(poolId);
            }
        }
    }

    // ==================== Native CROSS Deposit ====================

    /**
     * @notice Deposits native CROSS tokens
     * @dev Wraps native CROSS to WCROSS and deposits into the pool
     * @param poolId ID of the pool to deposit into
     */
    function depositNative(uint poolId) external payable {
        require(msg.value > 0, CSRInvalidAmount(msg.value));

        ICrossGameRewardPool pool = _getPoolAndValidateWCROSS(poolId);

        // Router wraps native CROSS to WCROSS
        wcross.deposit{value: msg.value}();

        // Router deposits to pool on behalf of msg.sender
        IERC20(wcross).forceApprove(address(pool), msg.value);
        pool.depositFor(msg.sender, msg.value);

        emit DepositedNative(msg.sender, poolId, msg.value);
    }

    function withdrawNative(uint poolId) public {
        withdrawNative(poolId, 0);
    }

    /**
     * @notice Withdraws and returns native CROSS tokens
     * @dev Withdraws WCROSS from pool and unwraps to native CROSS
     * @param poolId ID of the pool to withdraw from
     * @param amount Amount to withdraw (0 = withdraw all)
     */
    function withdrawNative(uint poolId, uint amount) public {
        ICrossGameRewardPool pool = _getPoolAndValidateWCROSS(poolId);

        // Get Router's WCROSS balance before withdrawal
        uint balanceBefore = IERC20(wcross).balanceOf(address(this));

        // Withdraw from pool (WCROSS sent to Router, rewards sent to msg.sender)
        pool.withdrawFor(msg.sender, amount);

        // Calculate actual withdrawn amount
        uint balanceAfter = IERC20(wcross).balanceOf(address(this));
        uint withdrawnAmount = balanceAfter - balanceBefore;

        // Router unwraps and sends native CROSS directly to user
        wcross.withdrawTo(msg.sender, withdrawnAmount);

        emit WithdrawnNative(msg.sender, poolId, withdrawnAmount);
    }

    // ==================== ERC20 Deposit (General Tokens) ====================

    /**
     * @notice Deposits ERC20 tokens
     * @dev Transfers tokens from user and deposits into the pool
     * @param poolId ID of the pool to deposit into
     * @param amount Amount of tokens to deposit
     */
    function depositERC20(uint poolId, uint amount) external {
        require(amount > 0, CSRInvalidAmount(amount));

        ICrossGameRewardPool pool = _getPool(poolId);
        IERC20 depositToken = pool.depositToken();

        _depositERC20(poolId, pool, depositToken, amount);
    }

    /**
     * @notice Deposits ERC20 tokens using EIP-2612 permit
     * @dev Performs permit + transfer + deposit in a single transaction (for tokens supporting EIP-2612)
     * @param poolId ID of the pool to deposit into
     * @param amount Amount of tokens to deposit
     * @param deadline Permit signature deadline
     * @param v Permit signature v
     * @param r Permit signature r
     * @param s Permit signature s
     */
    function depositERC20WithPermit(uint poolId, uint amount, uint deadline, uint8 v, bytes32 r, bytes32 s) external {
        require(amount > 0, CSRInvalidAmount(amount));

        ICrossGameRewardPool pool = _getPool(poolId);
        IERC20 depositToken = pool.depositToken();

        // Approve Router via EIP-2612 permit
        IERC20Permit(address(depositToken)).permit(msg.sender, address(this), amount, deadline, v, r, s);

        _depositERC20(poolId, pool, depositToken, amount);
    }

    function withdrawERC20(uint poolId) public {
        withdrawERC20(poolId, 0);
    }

    /**
     * @notice Withdraws ERC20 tokens
     * @dev Withdraws deposited tokens and claims rewards
     * @param poolId ID of the pool to withdraw from
     * @param amount Amount to withdraw (0 = withdraw all)
     */
    function withdrawERC20(uint poolId, uint amount) public {
        ICrossGameRewardPool pool = _getPool(poolId);
        IERC20 depositToken = pool.depositToken();

        // Get Router's deposit token balance before withdrawal
        uint balanceBefore = depositToken.balanceOf(address(this));

        // Withdraw (deposit tokens sent to Router, rewards to msg.sender)
        pool.withdrawFor(msg.sender, amount);

        // Calculate actual withdrawn amount
        uint balanceAfter = depositToken.balanceOf(address(this));
        uint withdrawnAmount = balanceAfter - balanceBefore;

        // Transfer deposit tokens from Router to user
        depositToken.safeTransfer(msg.sender, withdrawnAmount);

        emit WithdrawnERC20(msg.sender, poolId, address(depositToken), withdrawnAmount);
    }

    // ==================== Claim Rewards ====================

    /**
     * @notice Claims all pending rewards from a pool
     * @dev Claims all reward tokens and sends them directly to msg.sender
     * @param poolId ID of the pool to claim rewards from
     */
    function claimRewards(uint poolId) external {
        ICrossGameRewardPool pool = _getPool(poolId);

        // Claim all rewards on behalf of msg.sender (rewards sent directly to msg.sender)
        pool.claimRewardsFor(msg.sender);
    }

    /**
     * @notice Claims a specific reward token from a pool
     * @dev Claims only the specified reward token and sends it to msg.sender
     * @param poolId ID of the pool to claim rewards from
     * @param token Address of the reward token to claim
     */
    function claimReward(uint poolId, address token) external {
        require(token != address(0), CSRCanNotZeroAddress());

        ICrossGameRewardPool pool = _getPool(poolId);

        // Claim specific reward token on behalf of msg.sender (reward sent directly to msg.sender)
        pool.claimRewardFor(msg.sender, IERC20(token));
    }

    // ==================== View Functions ====================

    /**
     * @notice Retrieves user's deposit information
     * @param poolId ID of the pool
     * @param user Address of the user
     * @return depositedAmount Amount of tokens deposited
     * @return rewardTokens Array of reward token addresses
     * @return pendingRewards Array of pending rewards for each reward token
     */
    function getUserDepositInfo(uint poolId, address user)
        external
        view
        returns (uint depositedAmount, address[] memory rewardTokens, uint[] memory pendingRewards)
    {
        ICrossGameRewardPool pool = _getPool(poolId);
        depositedAmount = pool.balances(user);
        (rewardTokens, pendingRewards) = pool.pendingRewards(user);
    }

    /**
     * @notice Checks if a pool is a native CROSS pool
     * @param poolId ID of the pool
     * @return True if the pool uses WCROSS as deposit token
     */
    function isNativePool(uint poolId) external view returns (bool) {
        ICrossGameRewardPool pool = _getPool(poolId);
        return address(pool.depositToken()) == address(wcross);
    }

    /**
     * @notice Retrieves user's pending rewards from active reward tokens
     * @param poolId ID of the pool
     * @param user Address of the user
     * @return rewardTokens Array of active reward token addresses
     * @return pendingRewards Array of pending rewards for each active reward token
     */
    function getPendingRewards(uint poolId, address user)
        external
        view
        returns (address[] memory rewardTokens, uint[] memory pendingRewards)
    {
        ICrossGameRewardPool pool = _getPool(poolId);
        return pool.pendingRewards(user);
    }

    /**
     * @notice Retrieves user's pending reward for a specific token
     * @param poolId ID of the pool
     * @param user Address of the user
     * @param token Address of the reward token
     * @return amount Pending reward amount
     */
    function getPendingReward(uint poolId, address user, address token) external view returns (uint amount) {
        ICrossGameRewardPool pool = _getPool(poolId);
        return pool.pendingReward(user, IERC20(token));
    }

    /**
     * @notice Retrieves pending rewards for removed reward tokens
     * @param poolId ID of the pool
     * @param user Address of the user
     * @return rewardTokens Array of removed reward token addresses
     * @return pendingRewards Array of pending rewards for each removed reward token
     */
    function getRemovedTokenRewards(uint poolId, address user)
        external
        view
        returns (address[] memory rewardTokens, uint[] memory pendingRewards)
    {
        ICrossGameRewardPool pool = _getPool(poolId);
        return pool.getRemovedTokenRewards(user);
    }

    /**
     * @notice Retrieves all pending rewards (active + removed tokens) filtered by amount > 0
     * @param poolId ID of the pool
     * @param user Address of the user
     * @return rewardTokens Array of reward token addresses with positive pending amounts
     * @return pendingRewards Array of pending rewards corresponding to rewardTokens
     */
    function getAllPendingRewards(uint poolId, address user)
        external
        view
        returns (address[] memory rewardTokens, uint[] memory pendingRewards)
    {
        ICrossGameRewardPool pool = _getPool(poolId);
        return _collectPendingRewards(pool, user);
    }

    /**
     * @notice Retrieves total deposited amount across all pools grouped by deposit token
     * @dev Aggregates totalDeposited from all pools, grouped by their deposit token
     * @return depositTokens Array of unique deposit token addresses
     * @return totalDeposited Array of total deposited amounts for each token
     */
    function getTotalDeposited() external view returns (address[] memory depositTokens, uint[] memory totalDeposited) {
        uint[] memory poolIds = crossGameReward.getAllPoolIds();

        // First pass: collect unique deposit tokens
        address[] memory tempTokens = new address[](poolIds.length);
        uint[] memory tempAmounts = new uint[](poolIds.length);
        uint uniqueCount = 0;

        for (uint i = 0; i < poolIds.length; i++) {
            ICrossGameRewardPool pool = _getPool(poolIds[i]);
            address depositToken = address(pool.depositToken());
            uint poolDeposited = pool.totalDeposited();

            // Find if this token already exists in our temp arrays
            bool found = false;
            for (uint j = 0; j < uniqueCount; j++) {
                if (tempTokens[j] == depositToken) {
                    tempAmounts[j] += poolDeposited;
                    found = true;
                    break;
                }
            }

            // If not found, add as new unique token
            if (!found) {
                tempTokens[uniqueCount] = depositToken;
                tempAmounts[uniqueCount] = poolDeposited;
                uniqueCount++;
            }
        }

        // Second pass: create properly sized result arrays
        depositTokens = new address[](uniqueCount);
        totalDeposited = new uint[](uniqueCount);
        for (uint i = 0; i < uniqueCount; i++) {
            depositTokens[i] = tempTokens[i];
            totalDeposited[i] = tempAmounts[i];
        }
    }

    /**
     * @notice Retrieves total deposited amount for a specific token across all pools
     * @dev Sums totalDeposited from all pools that use the specified deposit token
     *      Special case: NATIVE_TOKEN (0x1) returns native CROSS (WCROSS) pool deposits
     * @param token Address of the deposit token to filter by (use NATIVE_TOKEN for native CROSS)
     * @return totalDeposited Total amount deposited across all pools using this token
     */
    function getTotalDeposited(address token) external view returns (uint totalDeposited) {
        // NATIVE_TOKEN (0x1) means native CROSS (WCROSS)
        address targetToken = token == NATIVE_TOKEN ? address(wcross) : token;

        uint[] memory poolIds = crossGameReward.getAllPoolIds();
        for (uint i = 0; i < poolIds.length; i++) {
            ICrossGameRewardPool pool = _getPool(poolIds[i]);
            if (address(pool.depositToken()) == targetToken) totalDeposited += pool.totalDeposited();
        }
    }

    // ==================== Internal Functions ====================

    /**
     * @dev Retrieves pool address and returns pool instance
     * @param poolId ID of the pool
     * @return Pool contract instance
     */
    function _getPool(uint poolId) internal view returns (ICrossGameRewardPool) {
        return crossGameReward.getPoolAddress(poolId);
    }

    /**
     * @dev Retrieves pool and validates it's a WCROSS pool
     * @param poolId ID of the pool
     * @return pool Pool contract instance
     */
    function _getPoolAndValidateWCROSS(uint poolId) internal view returns (ICrossGameRewardPool pool) {
        pool = _getPool(poolId);
        require(address(pool.depositToken()) == address(wcross), CSRNotWCROSSPool(poolId, address(pool.depositToken())));
    }

    function _depositERC20(uint poolId, ICrossGameRewardPool pool, IERC20 token, uint amount) internal {
        // Transfer tokens from user and deposit to pool
        token.safeTransferFrom(msg.sender, address(this), amount);
        token.forceApprove(address(pool), amount);
        pool.depositFor(msg.sender, amount);

        emit DepositedERC20(msg.sender, poolId, address(token), amount);
    }

    function _collectPendingRewards(ICrossGameRewardPool pool, address user)
        internal
        view
        returns (address[] memory tokens, uint[] memory amounts)
    {
        (address[] memory activeTokens, uint[] memory activeAmounts) = pool.pendingRewards(user);
        (address[] memory removedTokens, uint[] memory removedAmounts) = pool.getRemovedTokenRewards(user);

        uint totalCount = _countPositive(activeAmounts) + _countPositive(removedAmounts);
        tokens = new address[](totalCount);
        amounts = new uint[](totalCount);

        uint idx;
        for (uint i = 0; i < activeTokens.length; i++) {
            if (activeAmounts[i] == 0) continue;
            tokens[idx] = activeTokens[i];
            amounts[idx] = activeAmounts[i];
            idx++;
        }
        for (uint i = 0; i < removedTokens.length; i++) {
            if (removedAmounts[i] == 0) continue;
            tokens[idx] = removedTokens[i];
            amounts[idx] = removedAmounts[i];
            idx++;
        }

        return (tokens, amounts);
    }

    function _countPositive(uint[] memory values) private pure returns (uint count) {
        for (uint i = 0; i < values.length; i++) {
            if (values[i] > 0) count++;
        }
    }
}
