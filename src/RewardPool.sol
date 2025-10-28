// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {RewardPoolBase} from "./base/RewardPoolBase.sol";
import {IStakingPool} from "./interfaces/IStakingPool.sol";
import {IStakingProtocol} from "./interfaces/IStakingProtocol.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title RewardPoolCode
 * @notice Code contract that returns RewardPool's creation bytecode
 * @dev Part of Code Contract pattern for gas-efficient factory deployments
 */
contract RewardPoolCode {
    /**
     * @notice Returns the creation bytecode of RewardPool
     * @return Creation bytecode for RewardPool contract
     */
    function code() external pure returns (bytes memory) {
        return type(RewardPool).creationCode;
    }
}

/**
 * @title RewardPool
 * @notice Season-based reward token management and user claim processing
 * @dev Inherits from RewardPoolBase and implements:
 *      - Connection to StakingPool for points queries
 *      - Proportional reward distribution
 *      - Multi-token support per season
 *      - Double-claim prevention
 *      - Emergency token recovery (sweep)
 *
 *      Reward Flow:
 *      1. Project creator funds seasons via fundSeason()
 *      2. Users stake and earn points in StakingPool
 *      3. Season ends and is finalized
 *      4. Users claim via StakingPool.claimSeason()
 *      5. StakingPool calls RewardPool.payUser()
 *      6. Rewards transferred to user
 */
contract RewardPool is RewardPoolBase {
    using SafeERC20 for IERC20;

    // ============================================
    // Errors
    // ============================================

    /// @notice Thrown when caller is not the protocol contract
    error RewardPoolOnlyProtocol();

    /// @notice Thrown when attempting to deposit before pre-deposit period
    error RewardPoolPreDepositNotAvailable();

    // ============================================
    // Events
    // ============================================

    /// @notice Emitted when tokens are swept (emergency recovery)
    event TokensSwept(address indexed token, address indexed to, uint amount, uint balanceBefore, uint balanceAfter);

    // ============================================
    // Immutable State
    // ============================================

    /// @notice Connected StakingPool (source of points data)
    IStakingPool public immutable stakingPool;

    /// @notice Protocol contract (for access control)
    IStakingProtocol public immutable protocol;

    // ============================================
    // Constructor
    // ============================================

    /**
     * @notice Initializes the reward pool
     * @param _stakingPool Address of the connected StakingPool
     * @param _protocol Address of the StakingProtocol factory
     * @dev - Validates both addresses
     *      - Grants STAKING_POOL_ROLE to the staking pool
     *      - Sets both as immutable (cannot be changed)
     */
    constructor(address _stakingPool, address _protocol) RewardPoolBase(_protocol) {
        _validateAddress(_stakingPool);
        _validateAddress(_protocol);

        stakingPool = IStakingPool(_stakingPool);
        protocol = IStakingProtocol(_protocol);

        _grantRole(STAKING_POOL_ROLE, _stakingPool);
    }

    // ============================================
    // Hook Implementations
    // ============================================

    /**
     * @notice Calculates bonus rewards (currently returns 0)
     * @dev Override this in a derived contract to implement bonus logic
     *      Example: 10% bonus for users who staked the entire season
     */
    function _calculateBonusReward(
        address, /*user*/
        uint, /*season*/
        address, /*token*/
        uint, /*userPoints*/
        uint, /*totalPoints*/
        uint /*baseReward*/
    ) internal virtual override returns (uint bonusAmount) {
        return 0;
    }

    // ============================================
    // View Functions
    // ============================================

    /**
     * @notice Calculates expected reward for a user in a season
     * @param user User address
     * @param season Season number
     * @param token Reward token address
     * @return expectedReward Calculated reward amount
     * @dev Formula: (totalReward × userPoints) / totalPoints
     *      Returns 0 if already claimed or no points
     */
    function getExpectedReward(address user, uint season, address token)
        public
        view
        override
        returns (uint expectedReward)
    {
        if (hasClaimedSeasonReward[user][season][token]) return 0;

        (uint userPoints, uint totalPoints) = stakingPool.getSeasonUserPoints(season, user);

        if (userPoints == 0 || totalPoints == 0) return 0;

        uint totalReward = seasonRewards[season][token];
        if (totalReward == 0) return 0;

        return (totalReward * userPoints) / totalPoints;
    }

    /**
     * @notice Convenience function to deposit rewards for current/next season
     * @param token Reward token address
     * @param amount Amount to deposit
     * @dev - If season 0 (not started): deposits for season 1
     *      - Validates pre-deposit period if applicable
     *      - Otherwise deposits for current season
     */
    function depositReward(address token, uint amount) external {
        uint currentSeason = stakingPool.currentSeason();
        uint targetSeason = currentSeason;

        if (currentSeason == 0) {
            targetSeason = 1;

            uint preDepositStart = stakingPool.preDepositStartBlock();
            require(preDepositStart == 0 || block.number >= preDepositStart, RewardPoolPreDepositNotAvailable());
        }

        fundSeason(targetSeason, token, amount);
    }

    /**
     * @notice Returns pending rewards for user in current season
     * @param user User address
     * @param token Reward token address
     * @return Pending reward amount
     */
    function getUserPendingReward(address user, address token) external view returns (uint) {
        uint currentSeason = stakingPool.currentSeason();
        return getExpectedReward(user, currentSeason, token);
    }

    /**
     * @notice Returns summary of rewards for multiple tokens in a season
     * @param season Season number
     * @param tokens Array of token addresses
     * @return totals Total rewards per token
     * @return claimed Claimed rewards per token
     * @return remaining Remaining rewards per token
     */
    function getSeasonSummary(uint season, address[] calldata tokens)
        external
        view
        returns (uint[] memory totals, uint[] memory claimed, uint[] memory remaining)
    {
        uint len = tokens.length;
        totals = new uint[](len);
        claimed = new uint[](len);
        remaining = new uint[](len);

        for (uint i = 0; i < len; i++) {
            totals[i] = seasonRewards[season][tokens[i]];
            claimed[i] = seasonClaimed[season][tokens[i]];
            remaining[i] = totals[i] - claimed[i];
        }
    }

    /**
     * @notice Checks if user has claimed rewards for a season/token
     * @param user User address
     * @param season Season number
     * @param token Reward token address
     * @return True if already claimed
     */
    function getUserClaimed(address user, uint season, address token) external view returns (bool) {
        return hasClaimedSeasonReward[user][season][token];
    }

    // ============================================
    // Emergency Functions
    // ============================================

    /**
     * @notice Emergency function to recover tokens (protocol only)
     * @param token Token address to recover
     * @param to Recipient address
     * @param amount Amount to recover
     * @dev Only callable by protocol contract
     *      Use cases:
     *      - Recovering accidentally sent tokens
     *      - Recovering unclaimed rewards after extended period
     *      - Emergency fund recovery
     */
    function sweep(address token, address to, uint amount) external {
        require(msg.sender == address(protocol), RewardPoolOnlyProtocol());
        _validateAddress(to);
        _validateAmount(amount);

        uint balanceBefore = IERC20(token).balanceOf(address(this));
        IERC20(token).safeTransfer(to, amount);
        uint balanceAfter = IERC20(token).balanceOf(address(this));

        emit TokensSwept(token, to, amount, balanceBefore, balanceAfter);
    }
}
