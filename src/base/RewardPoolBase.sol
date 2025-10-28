// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IRewardPool} from "../interfaces/IRewardPool.sol";
import {PointsLib} from "../libraries/PointsLib.sol";
import {CrossStakingBase} from "./CrossStakingBase.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title RewardPoolBase
 * @notice Base abstract contract for reward distribution pools
 * @dev Provides core functionality for season-based reward management:
 *      - Multi-token reward support per season
 *      - Proportional distribution based on points
 *      - Double-claim prevention
 *      - Hook system for custom reward logic
 *
 *      Key Features:
 *      - Season rewards can be funded multiple times (accumulative)
 *      - Each user can claim once per season per token
 *      - Bonus reward system via hooks (extensible)
 *      - Safe ERC20 operations for all token transfers
 */
abstract contract RewardPoolBase is IRewardPool, CrossStakingBase {
    using SafeERC20 for IERC20;

    // ============ Constants ============

    /// @notice Role identifier for StakingPool (only StakingPool can call payUser)
    bytes32 public constant STAKING_POOL_ROLE = keccak256("STAKING_POOL_ROLE");

    // ============ Errors ============

    /// @notice Thrown when reward pool has insufficient balance to pay rewards
    error RewardPoolBaseInsufficientBalance();

    /// @notice Thrown when user attempts to claim rewards already claimed
    error RewardPoolBaseAlreadyClaimed();

    /// @notice Thrown when no rewards are available for the request
    error RewardPoolBaseNoRewards();

    // ============ State Variables ============

    /// @notice Total rewards deposited for each season and token
    /// @dev Mapping: seasonNumber => tokenAddress => totalAmount
    mapping(uint => mapping(address => uint)) public seasonRewards;

    /// @notice Total rewards claimed for each season and token
    /// @dev Mapping: seasonNumber => tokenAddress => claimedAmount
    mapping(uint => mapping(address => uint)) public seasonClaimed;

    /// @notice Tracks if a user has claimed rewards for a season/token combination
    /// @dev Mapping: userAddress => seasonNumber => tokenAddress => hasClaimed
    mapping(address => mapping(uint => mapping(address => bool))) public hasClaimedSeasonReward;

    /// @notice List of reward tokens used in each season (for enumeration)
    /// @dev Mapping: seasonNumber => array of token addresses
    mapping(uint => address[]) private seasonRewardTokens;

    /// @notice Index of each token in seasonRewardTokens array (1-based, 0 = not present)
    /// @dev Mapping: seasonNumber => tokenAddress => arrayIndex+1
    mapping(uint => mapping(address => uint)) private seasonTokenIndex;

    // ============ Events ============

    /// @notice Emitted when bonus rewards are paid (in addition to base rewards)
    event BonusRewardPaid(address indexed user, uint indexed season, address indexed token, uint bonusAmount);

    // ============ Constructor ============

    /**
     * @notice Initializes the reward pool base contract
     * @param admin Initial admin address
     */
    constructor(address admin) CrossStakingBase(admin) {}

    // ============ Core Functions ============

    /**
     * @notice Deposits reward tokens for a specific season
     * @param season Season number to fund
     * @param token Reward token address
     * @param amount Amount of tokens to deposit
     * @dev - Transfers tokens from caller to this contract
     *      - Can be called multiple times to add more rewards
     *      - Automatically tracks new tokens in seasonRewardTokens array
     *      - Emits SeasonFunded event
     */
    function fundSeason(uint season, address token, uint amount) public virtual nonReentrant {
        _validateAddress(token);
        _validateAmount(amount);

        _beforeFundSeason(season, token, amount);

        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);

        if (seasonTokenIndex[season][token] == 0) {
            seasonRewardTokens[season].push(token);
            seasonTokenIndex[season][token] = seasonRewardTokens[season].length;
        }

        seasonRewards[season][token] += amount;

        emit SeasonFunded(season, token, amount, amount);

        _afterFundSeason(season, token, amount);
    }

    /**
     * @notice Pays out rewards to a user based on their points (StakingPool only)
     * @param user User to receive rewards
     * @param season Season number
     * @param token Reward token address
     * @param userPoints User's points in the season
     * @param totalPoints Total points in the season
     * @dev - Only callable by StakingPool (STAKING_POOL_ROLE)
     *      - Calculates proportional reward: (totalReward × userPoints) / totalPoints
     *      - Includes bonus rewards via _calculateBonusReward hook
     *      - Prevents double-claiming via hasClaimedSeasonReward mapping
     *      - Emits RewardPaid and BonusRewardPaid events
     */
    function payUser(address user, uint season, address token, uint userPoints, uint totalPoints)
        external
        virtual
        onlyRole(STAKING_POOL_ROLE)
        nonReentrant
    {
        require(!hasClaimedSeasonReward[user][season][token], RewardPoolBaseAlreadyClaimed());
        require(userPoints != 0 && totalPoints != 0, RewardPoolBaseNoRewards());

        uint totalReward = seasonRewards[season][token];
        require(totalReward != 0, RewardPoolBaseNoRewards());

        uint baseReward = PointsLib.calculateProRata(userPoints, totalPoints, totalReward);
        require(baseReward != 0, RewardPoolBaseNoRewards());

        uint bonusReward = _calculateBonusReward(user, season, token, userPoints, totalPoints, baseReward);

        uint totalPayout = baseReward + bonusReward;

        uint balance = IERC20(token).balanceOf(address(this));
        require(balance >= totalPayout, RewardPoolBaseInsufficientBalance());

        hasClaimedSeasonReward[user][season][token] = true;
        seasonClaimed[season][token] += totalPayout;

        _beforePayUser(user, season, token, totalPayout);

        IERC20(token).safeTransfer(user, totalPayout);

        _afterPayUser(user, season, token, totalPayout);

        if (bonusReward > 0) emit BonusRewardPaid(user, season, token, bonusReward);
    }

    // ============ Hook Functions (Template Method Pattern) ============

    /**
     * @notice Hook called before funding a season
     * @dev Override to add custom validation or logic before rewards are deposited
     */
    function _beforeFundSeason(uint season, address token, uint amount) internal virtual {}

    /**
     * @notice Hook called after funding a season
     * @dev Override to add custom logic after rewards are deposited
     */
    function _afterFundSeason(uint season, address token, uint amount) internal virtual {}

    /**
     * @notice Hook for calculating bonus rewards (default: 0)
     * @dev Override to implement custom bonus logic (e.g., long-term staking bonus)
     * @return bonusAmount Additional rewards to pay on top of base reward
     */
    function _calculateBonusReward(
        address, /* user */
        uint, /* season */
        address, /* token */
        uint, /* userPoints */
        uint, /* totalPoints */
        uint /* baseReward */
    ) internal virtual returns (uint bonusAmount) {
        return 0;
    }

    /**
     * @notice Hook called before paying rewards to a user
     * @dev Override to add custom logic before token transfer
     */
    function _beforePayUser(address user, uint season, address token, uint amount) internal virtual {}

    /**
     * @notice Hook called after paying rewards to a user
     * @dev Override to add custom logic after token transfer
     */
    function _afterPayUser(address user, uint season, address token, uint amount) internal virtual {}

    // ============ View Functions ============

    /**
     * @notice Returns remaining unclaimed rewards for a season/token
     * @param season Season number
     * @param token Reward token address
     * @return Remaining reward amount (total - claimed)
     */
    function getRemainingRewards(uint season, address token) external view virtual returns (uint) {
        return seasonRewards[season][token] - seasonClaimed[season][token];
    }

    /**
     * @notice Calculates expected reward for a user (must be overridden)
     * @param user User address
     * @param season Season number
     * @param token Reward token address
     * @return expectedReward Calculated expected reward amount
     * @dev Must be implemented by inheriting contract to query points from StakingPool
     */
    function getExpectedReward(address user, uint season, address token)
        external
        view
        virtual
        returns (uint expectedReward);

    /**
     * @notice Returns all reward tokens used in a season
     * @param season Season number
     * @return tokens Array of token addresses
     */
    function getSeasonRewardTokens(uint season) external view returns (address[] memory tokens) {
        return seasonRewardTokens[season];
    }

    /**
     * @notice Returns detailed information for a specific season/token
     * @param season Season number
     * @param token Reward token address
     * @return total Total rewards deposited
     * @return claimed Total rewards claimed
     * @return remaining Unclaimed rewards
     */
    function getSeasonTokenInfo(uint season, address token)
        external
        view
        returns (uint total, uint claimed, uint remaining)
    {
        total = seasonRewards[season][token];
        claimed = seasonClaimed[season][token];
        remaining = total > claimed ? total - claimed : 0;
    }

    /**
     * @notice Returns complete reward information for all tokens in a season
     * @param season Season number
     * @return tokens Array of reward token addresses
     * @return totals Array of total rewards per token
     * @return claimeds Array of claimed rewards per token
     * @return remainings Array of remaining rewards per token
     */
    function getSeasonAllRewards(uint season)
        external
        view
        returns (address[] memory tokens, uint[] memory totals, uint[] memory claimeds, uint[] memory remainings)
    {
        tokens = seasonRewardTokens[season];
        uint length = tokens.length;

        totals = new uint[](length);
        claimeds = new uint[](length);
        remainings = new uint[](length);

        for (uint i = 0; i < length; i++) {
            address token = tokens[i];
            totals[i] = seasonRewards[season][token];
            claimeds[i] = seasonClaimed[season][token];
            remainings[i] = totals[i] > claimeds[i] ? totals[i] - claimeds[i] : 0;
        }
    }
}
