// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {WCROSS} from "./WCROSS.sol";
import {IStakingPool} from "./interfaces/IStakingPool.sol";
import {IStakingProtocol} from "./interfaces/IStakingProtocol.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuardTransient} from "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

/**
 * @title StakingRouter
 * @notice User-friendly router for native CROSS staking operations
 * @dev Provides convenient functions that automatically handle WCROSS wrapping/unwrapping
 *
 *      Key Features:
 *      - Stake with native CROSS (auto-wraps to WCROSS)
 *      - Unstake to native CROSS (auto-unwraps from WCROSS)
 *      - Batch claim operations
 *      - Proxy staking/withdrawing on behalf of users
 *
 *      User Flow:
 *      1. User sends native CROSS to stake()
 *      2. Router wraps to WCROSS
 *      3. Router approves and calls StakingPool.stakeFor()
 *      4. User's stake is recorded
 *
 *      Withdrawal Flow:
 *      1. Router calls StakingPool.withdrawAllFor(user)
 *      2. Router receives WCROSS
 *      3. Router unwraps to native CROSS
 *      4. Router sends native CROSS to user
 *
 *      Security:
 *      - Uses ReentrancyGuardTransient (EIP-1153)
 *      - Validates all project IDs
 *      - Safe ERC20 operations
 */
contract StakingRouter is ReentrancyGuardTransient {
    using SafeERC20 for IERC20;

    // ============ Errors ============

    /// @notice Thrown when zero address provided
    error StakingRouterCanNotZeroAddress();

    /// @notice Thrown when zero amount provided
    error StakingRouterAmountMustBeGreaterThanZero();

    /// @notice Thrown when user has no stake to withdraw
    error StakingRouterNoStake();

    /// @notice Thrown when project ID is invalid
    error StakingRouterInvalidProjectID();

    /// @notice Thrown when native CROSS transfer fails
    error StakingRouterTransferFailed();

    /// @notice Thrown when receive() called by non-WCROSS address
    error StakingRouterOnlyWCROSS();

    /// @notice Thrown when array length mismatch in batch operations
    error StakingRouterLengthMismatch();

    /// @notice Thrown when invalid offset provided
    error StakingRouterInvalidOffset();

    /// @notice Thrown when invalid range provided
    error StakingRouterInvalidRange();

    // ============ Modifiers ============

    /**
     * @notice Validates that a project ID exists
     * @param projectID Project ID to validate
     */
    modifier validProject(uint projectID) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        _;
    }

    // ============ Immutable State ============

    /// @notice WCROSS token contract
    WCROSS public immutable wcross;

    /// @notice StakingProtocol factory contract
    IStakingProtocol public immutable protocol;

    // ============ Events ============

    /// @notice Emitted when user stakes with native CROSS
    event StakedWithNative(address indexed user, uint indexed projectID, uint amount);

    /// @notice Emitted when user unstakes to native CROSS
    event UnstakedToNative(address indexed user, uint indexed projectID, uint amount);

    /// @notice Emitted when user claims rewards
    event RewardClaimed(
        address indexed user, uint indexed projectID, uint indexed season, address rewardToken, uint amount
    );

    // ============ Constructor ============

    /**
     * @notice Initializes the staking router
     * @param _wcross WCROSS token address
     * @param _protocol StakingProtocol factory address
     */
    constructor(address _wcross, address _protocol) {
        require(_wcross != address(0), StakingRouterCanNotZeroAddress());
        require(_protocol != address(0), StakingRouterCanNotZeroAddress());

        wcross = WCROSS(payable(_wcross));
        protocol = IStakingProtocol(_protocol);
    }

    // ============ Internal Helpers ============

    /**
     * @notice Gets pool addresses for a project
     * @param projectID Project ID
     * @return stakingPool StakingPool address
     * @return rewardPool RewardPool address
     */
    function _getProjectPools(uint projectID)
        internal
        view
        validProject(projectID)
        returns (address stakingPool, address rewardPool)
    {
        (stakingPool, rewardPool,,,,,) = protocol.projects(projectID);
    }

    // ============ Staking Functions ============

    /**
     * @notice Stakes native CROSS tokens for a project
     * @param projectID Project ID to stake in
     * @dev Process:
     *      1. Validates msg.value > 0
     *      2. Wraps native CROSS to WCROSS
     *      3. Approves StakingPool to spend WCROSS
     *      4. Calls StakingPool.stakeFor() on behalf of msg.sender
     *      5. Emits StakedWithNative event
     *
     *      Example: router.stake{value: 100 ether}(1)
     */
    function stake(uint projectID) external payable nonReentrant validProject(projectID) {
        uint amount = msg.value;
        require(amount > 0, StakingRouterAmountMustBeGreaterThanZero());

        (address stakingPool,) = _getProjectPools(projectID);

        wcross.deposit{value: amount}();
        IERC20(address(wcross)).safeIncreaseAllowance(stakingPool, amount);
        IStakingPool(stakingPool).stakeFor(msg.sender, amount);

        emit StakedWithNative(msg.sender, projectID, amount);
    }

    /**
     * @notice Withdraws all staked tokens as native CROSS
     * @param projectID Project ID to withdraw from
     * @dev Process:
     *      1. Calls StakingPool.withdrawAllFor() (router receives WCROSS)
     *      2. Unwraps WCROSS to native CROSS
     *      3. Sends native CROSS to msg.sender
     *      4. Emits UnstakedToNative event
     *
     *      Warning: Current season points are forfeited!
     *      Previous season rewards can still be claimed.
     */
    function unstake(uint projectID) external nonReentrant validProject(projectID) {
        (address stakingPool,) = _getProjectPools(projectID);

        IStakingPool(stakingPool).withdrawAllFor(msg.sender);

        uint wcrossBalance = IERC20(address(wcross)).balanceOf(address(this));
        require(wcrossBalance > 0, StakingRouterNoStake());

        wcross.withdraw(wcrossBalance);

        (bool success,) = msg.sender.call{value: wcrossBalance}("");
        require(success, StakingRouterTransferFailed());

        emit UnstakedToNative(msg.sender, projectID, wcrossBalance);
    }

    /**
     * @notice Receives native CROSS from WCROSS contract only
     * @dev Only accepts native CROSS from WCROSS.withdraw()
     */
    receive() external payable {
        require(msg.sender == address(wcross), StakingRouterOnlyWCROSS());
    }

    // ============ Reward Claim Functions ============

    /**
     * @notice Claims rewards for a single season
     * @param projectID Project ID
     * @param season Season number
     * @param rewardToken Reward token address
     * @dev Calls StakingPool.claimSeasonFor() which handles the actual claim logic
     */
    function claimReward(uint projectID, uint season, address rewardToken) external nonReentrant {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());

        (address stakingPool,,,,,,) = protocol.projects(projectID);
        IStakingPool(stakingPool).claimSeasonFor(msg.sender, season, rewardToken);

        emit RewardClaimed(msg.sender, projectID, season, rewardToken, 0);
    }

    /**
     * @notice Claims rewards for multiple seasons in one transaction (gas efficient)
     * @param projectID Project ID
     * @param seasons Array of season numbers
     * @param rewardTokens Array of reward token addresses (must match seasons length)
     * @dev Example:
     *      seasons = [1, 2, 3]
     *      tokens = [tokenA, tokenA, tokenB]
     *      Claims season 1 and 2 for tokenA, season 3 for tokenB
     */
    function claimMultipleRewards(uint projectID, uint[] calldata seasons, address[] calldata rewardTokens)
        external
        nonReentrant
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        require(seasons.length == rewardTokens.length, StakingRouterLengthMismatch());

        (address stakingPool,,,,,,) = protocol.projects(projectID);

        for (uint i = 0; i < seasons.length;) {
            IStakingPool(stakingPool).claimSeasonFor(msg.sender, seasons[i], rewardTokens[i]);
            emit RewardClaimed(msg.sender, projectID, seasons[i], rewardTokens[i], 0);
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Finalizes user's season snapshots in batch (gas optimization)
     * @param projectID Project ID
     * @param user User address
     * @param maxSeasons Maximum seasons to process
     * @return processed Number of seasons actually processed
     * @dev Useful for users who haven't interacted for many seasons
     *      Allows pre-processing snapshots to reduce gas on subsequent claims
     */
    function finalizeUserSeasonsBatch(uint projectID, address user, uint maxSeasons)
        external
        returns (uint processed)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);

        (bool success, bytes memory data) =
            stakingPool.call(abi.encodeWithSignature("finalizeUserSeasonsBatch(address,uint256)", user, maxSeasons));

        if (success && data.length >= 32) processed = abi.decode(data, (uint));

        return processed;
    }
}
