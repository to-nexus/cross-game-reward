// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {Initializable, UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardTransientUpgradeable} from
    "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardTransientUpgradeable.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {ICrossGameReward} from "./interfaces/ICrossGameReward.sol";
import {ICrossGameRewardPool} from "./interfaces/ICrossGameRewardPool.sol";
import {ICrossGameRewardPoolV2} from "./interfaces/ICrossGameRewardPoolV2.sol";

/**
 * @title CrossGameRewardPoolV2
 * @notice Game reward pool with round-based linear reward distribution
 * @dev Implements block-per-block reward distribution with UUPS upgradeability
 *
 * === Core Principles ===
 *
 * Reward Distribution:
 *   - Rewards are distributed linearly over blocks within each round
 *   - Multiple rounds can be active simultaneously (overlapping)
 *   - Users receive rewards proportionally based on their deposit share
 *
 * === Round System ===
 *
 * - Developers create rounds by depositing reward tokens
 * - Each round has a start block, end block, and reward per block
 * - Rounds can be cancelled before they start (full refund)
 * - Completed rounds are automatically cleaned up
 *
 * === Roles ===
 *
 * - Owner (CrossGameReward's admin): Full control, emergency functions
 * - RewardRoot (CrossGameReward contract): Pool management
 * - Developer: Round creation and cancellation
 */
contract CrossGameRewardPoolV2 is
    Initializable,
    PausableUpgradeable,
    ReentrancyGuardTransientUpgradeable,
    UUPSUpgradeable,
    ICrossGameRewardPoolV2
{
    using SafeERC20 for IERC20;
    using EnumerableSet for EnumerableSet.AddressSet;
    using EnumerableSet for EnumerableSet.UintSet;

    // ==================== Custom Errors ====================

    /// @notice Thrown when deposit amount is below minimum required
    error CGRP2BelowMinimumDepositAmount(uint provided, uint minimum);

    /// @notice Thrown when attempting to withdraw with no active deposit
    error CGRP2NoDepositFound(address account);

    /// @notice Thrown when attempting to withdraw more than the deposited amount
    error CGRP2InsufficientBalance(uint depositedAmount, uint withdrawAmount);

    /// @notice Thrown when a zero address is provided where it's not allowed
    error CGRP2CanNotZeroAddress();

    /// @notice Thrown when a zero value is provided where it's not allowed
    error CGRP2CanNotZeroValue();

    /// @notice Thrown when caller is not the authorized router
    error CGRP2OnlyRouter();

    /// @notice Thrown when caller is not the owner of the pool
    error CGRP2OnlyOwner();

    /// @notice Thrown when caller is not the reward root
    error CGRP2OnlyRewardRoot();

    /// @notice Thrown when caller does not have developer role
    error CGRP2OnlyDeveloper();

    /// @notice Thrown when attempting to deposit in an inactive or paused pool
    error CGRP2DepositNotAllowed(PoolStatus currentStatus);

    /// @notice Thrown when attempting an operation not allowed in current pool state
    error CGRP2NotAllowedInCurrentState();

    /// @notice Thrown when round is not found
    error CGRP2RoundNotFound(uint roundId);

    /// @notice Thrown when attempting to cancel an already started round
    error CGRP2RoundAlreadyStarted(uint roundId);

    /// @notice Thrown when attempting to cancel an already cancelled round
    error CGRP2RoundAlreadyCancelled(uint roundId);

    /// @notice Thrown when start block is not in the future
    error CGRP2InvalidStartBlock(uint startBlock, uint currentBlock);

    /// @notice Thrown when duration is zero
    error CGRP2InvalidDuration();

    /// @notice Thrown when there is no reclaimable amount
    error CGRP2NoReclaimableAmount();

    /// @notice Thrown when reward token cannot be deposit token
    error CGRP2RewardIsDepositToken();

    /// @notice Thrown when rewardPerBlock would be zero (amount < durationBlocks)
    error CGRP2RewardPerBlockZero();

    /// @notice Thrown when provided token does not match the pool's reward token
    /// @param provided The token address provided
    /// @param expected The expected reward token address
    error CGRP2InvalidRewardToken(address provided, address expected);

    // ==================== Constants ====================

    /// @notice Precision multiplier for reward calculations
    uint private constant PRECISION = 1e18;

    /// @notice Developer role identifier
    bytes32 public constant DEVELOPER_ROLE = keccak256("DEVELOPER_ROLE");

    // ==================== State Variables ====================

    /// @notice Block number when the contract was initialized
    uint public initializedAt;

    /// @notice The deposit token
    IERC20 public depositToken;

    /// @notice The single reward token for this pool
    IERC20 public rewardToken;

    /// @notice CrossGameReward contract address
    ICrossGameReward public crossGameReward;

    /// @notice Minimum amount required for depositing
    uint public minDepositAmount;

    /// @notice Mapping from user address to deposited balance
    mapping(address => uint) public balances;

    /// @notice Total amount of tokens deposited in the pool
    uint public totalDeposited;

    /// @notice Current status of the pool
    PoolStatus public poolStatus;

    /// @notice Set of accounts with developer role
    EnumerableSet.AddressSet private _developers;

    // ==================== Round Management ====================

    /// @notice Next round ID to be assigned
    uint public nextRoundId;

    /// @notice Mapping from round ID to round data
    mapping(uint => Round) private _rounds;

    /// @notice Set of active round IDs
    EnumerableSet.UintSet private _activeRoundIds;

    /// @notice Set of all round IDs (for history)
    EnumerableSet.UintSet private _allRoundIds;

    // ==================== Reward Tracking ====================

    /// @notice Global accumulated reward per share (sum of all rounds' accRewardPerShare)
    uint public globalAccRewardPerShare;

    /// @notice User's reward debt (last recorded globalAccRewardPerShare * balance)
    mapping(address => uint) private _userRewardDebt;

    /// @notice User's pending rewards (accumulated but not yet claimed)
    mapping(address => uint) private _userPendingRewards;

    /// @notice Reclaimable amount (rewards distributed when totalDeposited was 0)
    uint public reclaimableAmount;

    // ==================== Events ====================

    /// @notice Emitted when a user deposits tokens
    event Deposited(address indexed account, uint amount);

    /// @notice Emitted when a user withdraws tokens
    event Withdrawn(address indexed account, uint amount);

    /// @notice Emitted when a user claims rewards
    event RewardClaimed(address indexed account, IERC20 indexed token, uint amount);

    /// @notice Emitted when a reward claim fails
    event RewardClaimFailed(address indexed account, IERC20 indexed token, uint amount);

    /// @notice Emitted when the minimum deposit amount is updated
    event MinDepositAmountUpdated(uint oldAmount, uint newAmount);

    /// @notice Emitted when the pool status changes
    event PoolStatusChanged(PoolStatus oldStatus, PoolStatus newStatus);

    /// @notice Emitted when tokens are reclaimed
    event TokensReclaimed(IERC20 indexed token, address indexed to, uint amount);

    // ==================== Modifiers ====================

    modifier onlyOwner() {
        require(msg.sender == owner(), CGRP2OnlyOwner());
        _;
    }

    modifier onlyRewardRoot() {
        require(msg.sender == address(crossGameReward), CGRP2OnlyRewardRoot());
        _;
    }

    modifier onlyDeveloper() {
        require(_developers.contains(msg.sender), CGRP2OnlyDeveloper());
        _;
    }

    // ==================== Constructor ====================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // ==================== Initializer ====================

    /**
     * @notice Initializes the CrossGameRewardPoolV2 contract
     * @dev Sets up the pool with deposit and reward tokens
     * @param _depositToken Address of the token to be deposited
     * @param _rewardToken Address of the reward token
     * @param _minDepositAmount Minimum amount required for depositing
     */
    function initialize(IERC20 _depositToken, IERC20 _rewardToken, uint _minDepositAmount) external initializer {
        require(address(_depositToken) != address(0), CGRP2CanNotZeroAddress());
        require(address(_rewardToken) != address(0), CGRP2CanNotZeroAddress());
        require(address(_depositToken) != address(_rewardToken), CGRP2RewardIsDepositToken());
        require(_minDepositAmount > 0, CGRP2CanNotZeroValue());

        crossGameReward = ICrossGameReward(msg.sender);

        __Pausable_init();
        __ReentrancyGuardTransient_init();
        __UUPSUpgradeable_init();

        initializedAt = block.number;
        depositToken = _depositToken;
        rewardToken = _rewardToken;
        minDepositAmount = _minDepositAmount;
        poolStatus = PoolStatus.Active;
        nextRoundId = 1;
    }

    /**
     * @notice Returns the owner of this pool
     * @dev Returns CrossGameReward's owner
     */
    function owner() public view returns (address) {
        return crossGameReward.owner();
    }

    // ==================== Round Management Functions ====================

    /**
     * @notice Creates a new reward distribution round
     * @dev Transfers reward tokens from caller and schedules distribution
     * @param amount Total reward amount to distribute
     * @param startBlock Block number when distribution starts
     * @param durationBlocks Number of blocks over which to distribute rewards
     * @return roundId The ID of the created round
     */
    function createRound(uint amount, uint startBlock, uint durationBlocks)
        external
        nonReentrant
        whenNotPaused
        onlyDeveloper
        returns (uint roundId)
    {
        require(amount > 0, CGRP2CanNotZeroValue());
        require(startBlock > block.number, CGRP2InvalidStartBlock(startBlock, block.number));
        require(durationBlocks > 0, CGRP2InvalidDuration());

        uint rewardPerBlock = amount / durationBlocks;
        require(rewardPerBlock > 0, CGRP2RewardPerBlockZero());

        // Calculate actual distributable reward (remainder stays with developer)
        uint actualReward = rewardPerBlock * durationBlocks;

        // Update pool state before creating new round
        _updatePool();

        roundId = nextRoundId++;

        uint endBlock = startBlock + durationBlocks;

        _rounds[roundId] = Round({
            roundId: roundId,
            totalReward: actualReward,
            startBlock: startBlock,
            endBlock: endBlock,
            rewardPerBlock: rewardPerBlock,
            lastRewardBlock: startBlock,
            accRewardPerShare: 0,
            isCancelled: false
        });

        _activeRoundIds.add(roundId);
        _allRoundIds.add(roundId);

        // Transfer only actualReward from developer (remainder stays with developer)
        rewardToken.safeTransferFrom(msg.sender, address(this), actualReward);

        emit RoundCreated(roundId, actualReward, startBlock, endBlock, rewardPerBlock);
    }

    /**
     * @notice Cancels a round that hasn't started yet
     * @dev Refunds the reward tokens to the caller
     * @param roundId The ID of the round to cancel
     */
    function cancelRound(uint roundId) external nonReentrant whenNotPaused onlyDeveloper {
        Round storage round = _rounds[roundId];

        require(round.roundId != 0, CGRP2RoundNotFound(roundId));
        require(!round.isCancelled, CGRP2RoundAlreadyCancelled(roundId));
        require(block.number < round.startBlock, CGRP2RoundAlreadyStarted(roundId));

        round.isCancelled = true;
        _activeRoundIds.remove(roundId);

        uint refundAmount = round.totalReward;
        rewardToken.safeTransfer(msg.sender, refundAmount);

        emit RoundCancelled(roundId, refundAmount);
    }

    // ==================== View Functions - Round ====================

    /**
     * @notice Returns round information by round ID
     */
    function getRound(uint roundId) external view returns (Round memory) {
        require(_rounds[roundId].roundId != 0, CGRP2RoundNotFound(roundId));
        return _rounds[roundId];
    }

    /**
     * @notice Returns all active rounds
     */
    function getActiveRounds() external view returns (Round[] memory) {
        uint length = _activeRoundIds.length();
        Round[] memory rounds = new Round[](length);

        for (uint i = 0; i < length; i++) {
            rounds[i] = _rounds[_activeRoundIds.at(i)];
        }

        return rounds;
    }

    /**
     * @notice Returns the number of active rounds
     */
    function getActiveRoundCount() external view returns (uint) {
        return _activeRoundIds.length();
    }

    /**
     * @notice Returns all active round IDs
     */
    function getActiveRoundIds() external view returns (uint[] memory) {
        return _activeRoundIds.values();
    }

    // ==================== Deposit Functions ====================

    /**
     * @notice Deposits tokens into the pool
     * @param amount Amount of tokens to deposit
     */
    function deposit(uint amount) external nonReentrant whenNotPaused {
        require(poolStatus == PoolStatus.Active, CGRP2DepositNotAllowed(poolStatus));
        _deposit(msg.sender, msg.sender, amount);
    }

    /**
     * @notice Deposits tokens on behalf of another account (Router only)
     * @param account Address of the account to deposit for
     * @param amount Amount of tokens to deposit
     */
    function depositFor(address account, uint amount) external nonReentrant whenNotPaused {
        require(poolStatus == PoolStatus.Active, CGRP2DepositNotAllowed(poolStatus));
        _checkDelegate(account);
        _deposit(msg.sender, account, amount);
    }

    // ==================== Withdraw Functions ====================

    /**
     * @notice Withdraws deposited tokens and claims all rewards
     * @param amount Amount of tokens to withdraw (0 = withdraw all)
     */
    function withdraw(uint amount) external nonReentrant whenNotPaused {
        require(poolStatus != PoolStatus.Paused, CGRP2NotAllowedInCurrentState());
        _withdraw(msg.sender, msg.sender, amount);
    }

    /**
     * @notice Withdraws tokens on behalf of another account (Router only)
     * @param account Address of the account to withdraw for
     * @param amount Amount of tokens to withdraw (0 = withdraw all)
     */
    function withdrawFor(address account, uint amount) external nonReentrant whenNotPaused {
        require(poolStatus != PoolStatus.Paused, CGRP2NotAllowedInCurrentState());
        _checkDelegate(account);
        _withdraw(msg.sender, account, amount);
    }

    // ==================== Claim Functions ====================

    /**
     * @notice Claims all pending rewards without withdrawing
     */
    function claimRewards() external nonReentrant whenNotPaused {
        require(poolStatus != PoolStatus.Paused, CGRP2NotAllowedInCurrentState());
        _claimRewards(msg.sender);
    }

    /**
     * @notice Claims all pending rewards on behalf of another account (Router only)
     * @param account Address of the account to claim rewards for
     */
    function claimRewardsFor(address account) external nonReentrant whenNotPaused {
        require(poolStatus != PoolStatus.Paused, CGRP2NotAllowedInCurrentState());
        _checkDelegate(account);
        _claimRewards(account);
    }

    /**
     * @notice Claims pending rewards for a specific token
     * @dev In V2, validates that the provided token matches the pool's reward token
     * @param token Address of the reward token to claim (must match pool's rewardToken)
     */
    function claimReward(IERC20 token) external nonReentrant whenNotPaused {
        require(poolStatus != PoolStatus.Paused, CGRP2NotAllowedInCurrentState());
        require(token == rewardToken, CGRP2InvalidRewardToken(address(token), address(rewardToken)));
        _claimRewards(msg.sender);
    }

    /**
     * @notice Claims pending reward for a specific token on behalf of another account
     * @dev In V2, validates that the provided token matches the pool's reward token
     * @param account Address of the account to claim rewards for
     * @param token Address of the reward token to claim (must match pool's rewardToken)
     */
    function claimRewardFor(address account, IERC20 token) external nonReentrant whenNotPaused {
        require(poolStatus != PoolStatus.Paused, CGRP2NotAllowedInCurrentState());
        require(token == rewardToken, CGRP2InvalidRewardToken(address(token), address(rewardToken)));
        _checkDelegate(account);
        _claimRewards(account);
    }

    // ==================== View Functions - Rewards ====================

    /**
     * @notice Retrieves pending rewards for a user
     * @param user Address of the user to query
     * @return tokens Array of reward token addresses (single element)
     * @return rewards Array of pending reward amounts (single element)
     */
    function pendingRewards(address user) external view returns (address[] memory tokens, uint[] memory rewards) {
        tokens = new address[](1);
        rewards = new uint[](1);

        tokens[0] = address(rewardToken);
        rewards[0] = _calculatePendingReward(user);
    }

    /**
     * @notice Retrieves pending reward for a specific token
     * @param user Address of the user to query
     * @return amount Pending reward amount
     */
    function pendingReward(address user, IERC20) external view returns (uint amount) {
        return _calculatePendingReward(user);
    }

    // ==================== View Functions - ICrossGameRewardPool Compatibility ====================

    /**
     * @notice Returns reward token address at index
     * @dev V2 only has one reward token, so only index 0 is valid
     */
    function rewardTokenAt(uint index) external view returns (IERC20) {
        require(index == 0, "Invalid index");
        return rewardToken;
    }

    /**
     * @notice Returns reward token information
     * @dev Returns a RewardToken struct compatible with V1 interface
     */
    function getRewardToken(IERC20 token) external view returns (RewardToken memory) {
        require(token == rewardToken, "Invalid token");

        return RewardToken({
            token: rewardToken,
            rewardPerTokenStored: globalAccRewardPerShare,
            lastBalance: rewardToken.balanceOf(address(this)),
            reclaimableAmount: reclaimableAmount,
            distributedAmount: 0, // Not tracked in V2 the same way
            isRemoved: false
        });
    }

    /**
     * @notice Checks if a token is a registered reward token
     */
    function isRewardToken(IERC20 token) external view returns (bool) {
        return token == rewardToken;
    }

    /**
     * @notice Returns all reward token addresses
     */
    function getRewardTokens() external view returns (address[] memory) {
        address[] memory tokens = new address[](1);
        tokens[0] = address(rewardToken);
        return tokens;
    }

    /**
     * @notice Returns the number of reward tokens
     */
    function rewardTokenCount() external pure returns (uint) {
        return 1;
    }

    /**
     * @notice Returns all removed reward token addresses
     * @dev V2 doesn't support removing reward tokens, always returns empty
     */
    function getRemovedRewardTokens() external pure returns (address[] memory) {
        return new address[](0);
    }

    /**
     * @notice Returns the number of removed reward tokens
     */
    function removedRewardTokenCount() external pure returns (uint) {
        return 0;
    }

    /**
     * @notice Checks if a token is a removed reward token
     */
    function isRemovedRewardToken(IERC20) external pure returns (bool) {
        return false;
    }

    /**
     * @notice Returns user's claimable rewards for removed tokens
     * @dev V2 doesn't support removing reward tokens, always returns empty
     */
    function getRemovedTokenRewards(address) external pure returns (address[] memory tokens, uint[] memory rewards) {
        return (new address[](0), new uint[](0));
    }

    /**
     * @notice Returns user reward info (V1 compatibility)
     */
    function userRewards(address account, IERC20) external view returns (uint rewardPerTokenPaid, uint rewards) {
        uint balance = balances[account];
        if (balance > 0) rewardPerTokenPaid = _userRewardDebt[account] / balance;
        rewards = _userPendingRewards[account];
    }

    // ==================== Admin Functions ====================

    /**
     * @notice Adds a reward token (V1 compatibility - no-op in V2)
     * @dev V2 has a fixed reward token set at initialization
     */
    function addRewardToken(IERC20 token) external view onlyRewardRoot {
        require(token == rewardToken, "Cannot add different reward token");
    }

    /**
     * @notice Removes a reward token (V1 compatibility - no-op in V2)
     * @dev V2 doesn't support removing reward tokens
     */
    function removeRewardToken(IERC20) external view onlyRewardRoot {
        revert("V2 does not support removing reward tokens");
    }

    /**
     * @notice Returns the reclaimable amount
     */
    function getReclaimableAmount(IERC20 token) external view returns (uint) {
        if (token != rewardToken) return 0;
        return reclaimableAmount;
    }

    /**
     * @notice Reclaims tokens that couldn't be distributed
     * @dev Only callable by CrossGameReward contract
     */
    function reclaimTokens(IERC20 token, address to) external onlyRewardRoot {
        require(token == rewardToken, "Invalid token");
        require(reclaimableAmount > 0, CGRP2NoReclaimableAmount());
        require(to != address(0), CGRP2CanNotZeroAddress());

        uint amount = reclaimableAmount;
        reclaimableAmount = 0;

        rewardToken.safeTransfer(to, amount);
        emit TokensReclaimed(rewardToken, to, amount);
    }

    /**
     * @notice Sets the minimum deposit amount
     */
    function updateMinDepositAmount(uint amount) external onlyRewardRoot {
        require(amount > 0, CGRP2CanNotZeroValue());
        emit MinDepositAmountUpdated(minDepositAmount, amount);
        minDepositAmount = amount;
    }

    /**
     * @notice Sets the pool status
     */
    function setPoolStatus(PoolStatus newStatus) external onlyRewardRoot {
        PoolStatus oldStatus = poolStatus;
        require(oldStatus != newStatus, "Pool status unchanged");

        poolStatus = newStatus;

        if (newStatus == PoolStatus.Paused && !paused()) _pause();
        else if (newStatus != PoolStatus.Paused && paused()) _unpause();

        emit PoolStatusChanged(oldStatus, newStatus);
    }

    /**
     * @notice Grants developer role to an account
     */
    function grantDeveloperRole(address developer) external onlyRewardRoot {
        require(developer != address(0), CGRP2CanNotZeroAddress());
        if (_developers.add(developer)) emit DeveloperRoleGranted(developer);
    }

    /**
     * @notice Revokes developer role from an account
     */
    function revokeDeveloperRole(address developer) external onlyRewardRoot {
        if (_developers.remove(developer)) emit DeveloperRoleRevoked(developer);
    }

    /**
     * @notice Checks if an account has developer role
     */
    function hasDeveloperRole(address account) external view returns (bool) {
        return _developers.contains(account);
    }

    // ==================== Internal Functions: Pool Update ====================

    /**
     * @dev Updates all active rounds and accumulates rewards
     */
    function _updatePool() internal {
        uint length = _activeRoundIds.length();
        if (length == 0) return;

        uint[] memory toRemove = new uint[](length);
        uint removeCount = 0;

        for (uint i = 0; i < length; i++) {
            uint roundId = _activeRoundIds.at(i);
            Round storage round = _rounds[roundId];

            // Skip cancelled rounds
            if (round.isCancelled) {
                toRemove[removeCount++] = roundId;
                continue;
            }

            // Skip rounds that haven't started
            if (block.number < round.startBlock) continue;

            // Skip if already updated this block
            if (block.number <= round.lastRewardBlock) continue;

            uint endBlock = block.number < round.endBlock ? block.number : round.endBlock;
            uint multiplier = endBlock - round.lastRewardBlock;
            uint reward = multiplier * round.rewardPerBlock;

            if (totalDeposited == 0) {
                // No depositors - accumulate to reclaimable
                reclaimableAmount += reward;
            } else {
                // Distribute to users
                uint rewardPerShare = (reward * PRECISION) / totalDeposited;
                round.accRewardPerShare += rewardPerShare;
                globalAccRewardPerShare += rewardPerShare;
            }

            round.lastRewardBlock = endBlock;

            // Mark completed rounds for removal
            if (block.number >= round.endBlock) toRemove[removeCount++] = roundId;
        }

        // Remove completed rounds from active set
        for (uint i = 0; i < removeCount; i++) {
            _activeRoundIds.remove(toRemove[i]);
        }
    }

    /**
     * @dev Updates user's pending rewards based on current pool state
     */
    function _updateUser(address account) internal {
        uint userBalance = balances[account];

        if (userBalance > 0) {
            uint pending = (userBalance * globalAccRewardPerShare - _userRewardDebt[account]) / PRECISION;
            if (pending > 0) _userPendingRewards[account] += pending;
        }

        _userRewardDebt[account] = userBalance * globalAccRewardPerShare;
    }

    /**
     * @dev Calculates pending rewards for a user (view function)
     */
    function _calculatePendingReward(address user) internal view returns (uint) {
        uint userBalance = balances[user];
        uint pending = _userPendingRewards[user];

        if (userBalance == 0) return pending;

        // Simulate pool update
        uint simulatedGlobalAcc = globalAccRewardPerShare;
        uint length = _activeRoundIds.length();

        for (uint i = 0; i < length; i++) {
            uint roundId = _activeRoundIds.at(i);
            Round storage round = _rounds[roundId];

            if (round.isCancelled || block.number < round.startBlock) continue;

            if (block.number <= round.lastRewardBlock) continue;

            uint endBlock = block.number < round.endBlock ? block.number : round.endBlock;
            uint multiplier = endBlock - round.lastRewardBlock;
            uint reward = multiplier * round.rewardPerBlock;

            if (totalDeposited > 0) simulatedGlobalAcc += (reward * PRECISION) / totalDeposited;
        }

        uint newPending = (userBalance * simulatedGlobalAcc - _userRewardDebt[user]) / PRECISION;
        return pending + newPending;
    }

    // ==================== Internal Functions: Deposit/Withdraw/Claim ====================

    /**
     * @dev Internal deposit logic
     */
    function _deposit(address payer, address account, uint amount) internal {
        require(amount >= minDepositAmount, CGRP2BelowMinimumDepositAmount(amount, minDepositAmount));

        _updatePool();
        _updateUser(account);

        depositToken.safeTransferFrom(payer, address(this), amount);

        balances[account] += amount;
        totalDeposited += amount;

        // Update debt after balance change
        _userRewardDebt[account] = balances[account] * globalAccRewardPerShare;

        emit Deposited(account, amount);
    }

    /**
     * @dev Internal withdraw logic
     */
    function _withdraw(address caller, address account, uint amount) internal {
        require(balances[account] > 0, CGRP2NoDepositFound(account));

        uint withdrawAmount = amount == 0 ? balances[account] : amount;
        require(withdrawAmount <= balances[account], CGRP2InsufficientBalance(balances[account], withdrawAmount));

        _updatePool();
        _updateUser(account);

        // Claim rewards
        _claimRewardsInternal(account);

        balances[account] -= withdrawAmount;
        totalDeposited -= withdrawAmount;

        // Update debt after balance change
        _userRewardDebt[account] = balances[account] * globalAccRewardPerShare;

        depositToken.safeTransfer(caller, withdrawAmount);

        emit Withdrawn(account, withdrawAmount);
    }

    /**
     * @dev Internal claim logic
     */
    function _claimRewards(address account) internal {
        uint userBalance = balances[account];
        uint pending = _userPendingRewards[account];

        require(userBalance > 0 || pending > 0, CGRP2NoDepositFound(account));

        if (userBalance > 0) {
            _updatePool();
            _updateUser(account);
        }

        _claimRewardsInternal(account);
    }

    /**
     * @dev Transfers pending rewards to user
     */
    function _claimRewardsInternal(address account) internal {
        uint reward = _userPendingRewards[account];

        if (reward > 0) {
            _userPendingRewards[account] = 0;

            bool ok = rewardToken.trySafeTransfer(account, reward);
            if (!ok) {
                _userPendingRewards[account] = reward;
                emit RewardClaimFailed(account, rewardToken, reward);
            } else {
                emit RewardClaimed(account, rewardToken, reward);
            }
        }
    }

    /**
     * @dev Validates that the caller is the authorized router
     */
    function _checkDelegate(address account) internal view {
        require(account != address(0), CGRP2CanNotZeroAddress());
        require(msg.sender == crossGameReward.router(), CGRP2OnlyRouter());
    }

    // ==================== UUPS ====================

    /**
     * @dev Authorizes contract upgrades
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    // ==================== Storage Gap ====================

    /**
     * @dev Storage gap for future upgrades
     */
    uint[30] private __gap;
}
