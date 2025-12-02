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

/**
 * @title CrossGameRewardPool
 * @notice Game reward pool with real-time reward distribution mechanism
 * @dev Implements rewardPerToken accumulation pattern with UUPS upgradeability
 *
 * === Core Principles ===
 *
 * Reward Distribution:
 *   - Users only receive rewards deposited after their deposit
 *   - Rewards are distributed proportionally based on deposit amount at deposit time
 *   - Uses cumulative rewardPerToken accounting
 *   - Zero-deposit rewards are marked as reclaimable (owner can reclaim)
 *
 * === Fairness ===
 *
 * Early depositors:
 *   - Receive rewards for a longer duration
 *   - Each reward is distributed proportionally to current deposit amounts
 *   - Do NOT receive rewards deposited before their deposit
 *
 * === Upgradeability ===
 *
 * UUPS Pattern:
 *   - Owner-only upgrade authorization
 *   - Simplified access control (onlyOwner/onlyRewardRoot)
 *   - Built-in Pausable functionality via pool status
 */
contract CrossGameRewardPool is
    Initializable,
    PausableUpgradeable,
    ReentrancyGuardTransientUpgradeable,
    UUPSUpgradeable,
    ICrossGameRewardPool
{
    using SafeERC20 for IERC20;
    using EnumerableSet for EnumerableSet.AddressSet;

    // ==================== Custom Errors ====================

    /// @notice Thrown when deposit amount is below minimum required
    /// @param provided The amount provided
    /// @param minimum The minimum required amount
    error CGRPBelowMinimumDepositAmount(uint provided, uint minimum);

    /// @notice Thrown when attempting to withdraw with no active deposit
    /// @param account The account with no deposit
    error CGRPNoDepositFound(address account);

    /// @notice Thrown when attempting to withdraw more than the deposited amount
    /// @param depositedAmount The amount deposited
    /// @param withdrawAmount The amount to withdraw
    error CGRPInsufficientBalance(uint depositedAmount, uint withdrawAmount);

    /// @notice Thrown when a zero address is provided where it's not allowed
    error CGRPCanNotZeroAddress();

    /// @notice Thrown when a zero value is provided where it's not allowed
    error CGRPCanNotZeroValue();

    /// @notice Thrown when attempting to add an already existing reward token
    /// @param token The token address that is already added
    error CGRPRewardTokenAlreadyAdded(address token);

    /// @notice Thrown when accessing an invalid reward token
    /// @param token The invalid token address
    error CGRPInvalidRewardToken(address token);

    /// @notice Thrown when attempting to use deposit token as reward token
    error CGRPCanNotUseDepositToken();

    /// @notice Thrown when caller is not the authorized router
    error CGRPOnlyRouter();

    /// @notice Thrown when attempting reclaim with no reclaimable amount
    /// @param token The token with no reclaimable amount
    error CGRPNoReclaimableAmount(address token);

    /// @notice Thrown when attempting to call a function that is not allowed
    error CGRPNotAllowedFunction();

    /// @notice Thrown when caller is not the owner of the pool
    error CGRPOnlyOwner();

    /// @notice Thrown when caller is not the reward root
    error CGRPOnlyRewardRoot();

    /// @notice Thrown when attempting to deposit in an inactive or paused pool
    /// @param currentStatus The current pool status
    error CGRPCannotDepositInCurrentState(PoolStatus currentStatus);

    /// @notice Thrown when attempting an operation not allowed in current pool state
    error CGRPNotAllowedInCurrentState();

    // ==================== Constants ====================

    /// @notice Block number when the contract was initialized
    uint public initializedAt;

    /// @notice Precision multiplier for reward calculations
    uint private constant PRECISION = 1e18;

    // ==================== State Variables ====================

    /// @notice The deposit token
    IERC20 public depositToken;

    /// @notice CrossGameReward contract address (used for router validation)
    ICrossGameReward public crossGameReward;

    /// @notice Minimum amount required for depositing
    uint public minDepositAmount;

    /// @notice Set of active reward token addresses
    EnumerableSet.AddressSet private _rewardTokenAddresses;

    /// @notice Set of removed reward token addresses (still claimable)
    EnumerableSet.AddressSet private _removedRewardTokenAddresses;

    /// @notice Mapping from reward token to reward token data
    mapping(IERC20 => RewardToken) private _rewardTokenData;

    /// @notice Mapping from user address to deposited balance
    mapping(address => uint) public balances;

    /// @notice Mapping from user address to reward token to user reward data
    mapping(address => mapping(IERC20 => UserReward)) public userRewards;

    /// @notice Total amount of tokens deposited in the pool
    uint public totalDeposited;

    /// @notice Current status of the pool
    ICrossGameRewardPool.PoolStatus public poolStatus;

    // ==================== Events ====================

    /// @notice Emitted when a account deposits tokens
    /// @param account Address of the account who deposited
    /// @param amount Amount of tokens deposited
    event Deposited(address indexed account, uint amount);

    /// @notice Emitted when a account withdraws tokens
    /// @param account Address of the account who withdrew
    /// @param amount Amount of tokens withdrawn
    event Withdrawn(address indexed account, uint amount);

    /// @notice Emitted when a account claims rewards
    /// @param account Address of the account who claimed
    /// @param token Address of the reward token claimed
    /// @param amount Amount of reward tokens claimed
    event RewardClaimed(address indexed account, IERC20 indexed token, uint amount);

    /// @notice Emitted when a account claims rewards failed
    /// @param account Address of the account who claimed
    /// @param token Address of the reward token claimed
    /// @param amount Amount of reward tokens claimed
    event RewardClaimFailed(address indexed account, IERC20 indexed token, uint amount);

    /// @notice Emitted when a new reward token is added
    /// @param token Address of the added reward token
    event RewardTokenAdded(IERC20 indexed token);

    /// @notice Emitted when a reward token is removed
    /// @param token Address of the removed reward token
    event RewardTokenRemoved(IERC20 indexed token);

    /// @notice Emitted when rewards are synced to the pool
    /// @param token Address of the reward token
    /// @param amount Amount of rewards added to the pool
    /// @param totalDeposited Total amount of tokens deposited in the pool
    event RewardSynced(IERC20 indexed token, uint amount, uint totalDeposited);

    /// @notice Emitted when admin reclaims tokens
    /// @param token Address of the reward token
    /// @param to Address receiving the reclaimed tokens
    /// @param amount Amount reclaimed
    event TokensReclaimed(IERC20 indexed token, address indexed to, uint amount);

    /// @notice Emitted when the minimum deposit amount is updated
    /// @param oldAmount Old minimum deposit amount
    /// @param newAmount New minimum deposit amount
    event MinDepositAmountUpdated(uint oldAmount, uint newAmount);

    /// @notice Emitted when the pool status changes
    /// @param oldStatus Old pool status
    /// @param newStatus New pool status
    event PoolStatusChanged(ICrossGameRewardPool.PoolStatus oldStatus, ICrossGameRewardPool.PoolStatus newStatus);

    // ==================== Initializer ====================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    modifier onlyOwner() {
        require(msg.sender == owner(), CGRPOnlyOwner());
        _;
    }

    modifier onlyRewardRoot() {
        require(msg.sender == address(crossGameReward), CGRPOnlyRewardRoot());
        _;
    }

    /**
     * @notice Initializes the CrossGameRewardPool contract
     * @dev Sets up roles and initializes inherited contracts
     *      msg.sender must be CrossGameReward contract
     *      Pool's DEFAULT_ADMIN_ROLE = CrossGameReward's owner (via owner() override)
     *      REWARD_ROOT_ROLE is granted to CrossGameReward contract for pool management
     * @param _depositToken Address of the token to be deposited
     * @param _minDepositAmount Minimum amount required for depositing
     */
    function initialize(IERC20 _depositToken, uint _minDepositAmount) external initializer {
        require(address(_depositToken) != address(0), CGRPCanNotZeroAddress());
        require(_minDepositAmount > 0, CGRPCanNotZeroValue());

        // msg.sender is always CrossGameReward contract
        crossGameReward = ICrossGameReward(msg.sender);

        __Pausable_init();
        __ReentrancyGuardTransient_init();
        __UUPSUpgradeable_init();

        initializedAt = block.number;
        depositToken = _depositToken;
        minDepositAmount = _minDepositAmount;
        poolStatus = ICrossGameRewardPool.PoolStatus.Active;
    }

    /**
     * @notice Returns the owner of this pool
     * @dev Returns CrossGameReward's owner (default admin)
     *      This ensures pool owner is the CrossGameReward contract's admin
     * @return Address of the CrossGameReward contract's default admin
     */
    function owner() public view returns (address) {
        return crossGameReward.owner();
    }

    // ==================== Internal Helper Functions ====================

    /**
     * @dev Checks if user has stored rewards in any token
     * @param user Address to check
     * @return True if user has any unclaimed rewards
     */
    function _hasStoredRewards(address user) private view returns (bool) {
        uint length = _rewardTokenAddresses.length();
        for (uint i = 0; i < length; ++i) {
            IERC20 token = IERC20(_rewardTokenAddresses.at(i));
            if (userRewards[user][token].rewards > 0) return true;
        }
        length = _removedRewardTokenAddresses.length();
        for (uint i = 0; i < length; ++i) {
            IERC20 token = IERC20(_removedRewardTokenAddresses.at(i));
            if (userRewards[user][token].rewards > 0) return true;
        }
        return false;
    }

    // ==================== External Functions ====================

    /**
     * @notice Deposits tokens into the pool
     * @dev Additional deposits are accumulated to existing balance
     *      Only allowed when pool is Active
     * @param amount Amount of tokens to deposit
     */
    function deposit(uint amount) external nonReentrant whenNotPaused {
        require(poolStatus == ICrossGameRewardPool.PoolStatus.Active, CGRPCannotDepositInCurrentState(poolStatus));
        _deposit(msg.sender, msg.sender, amount);
    }

    /**
     * @notice Deposits tokens on behalf of another account (Router only)
     * @dev Verifies that msg.sender is the registered router
     *      Only allowed when pool is Active
     * @param account Address of the account to deposit for
     * @param amount Amount of tokens to deposit
     */
    function depositFor(address account, uint amount) external nonReentrant whenNotPaused {
        require(poolStatus == ICrossGameRewardPool.PoolStatus.Active, CGRPCannotDepositInCurrentState(poolStatus));
        _checkDelegate(account);
        _deposit(msg.sender, account, amount);
    }

    /**
     * @notice Withdraws deposited tokens and claims all rewards
     * @dev Automatically claims all accumulated rewards
     *      Allowed in Active and Inactive states, blocked in Paused state
     * @param amount Amount of tokens to withdraw (0 = withdraw all)
     */
    function withdraw(uint amount) external nonReentrant whenNotPaused {
        require(poolStatus != ICrossGameRewardPool.PoolStatus.Paused, CGRPNotAllowedInCurrentState());
        _withdraw(msg.sender, msg.sender, amount);
    }

    /**
     * @notice Withdraws tokens on behalf of another account (Router only)
     * @dev Verifies that msg.sender is the registered router
     *      Allowed in Active and Inactive states, blocked in Paused state
     * @param account Address of the account to withdraw for
     * @param amount Amount of tokens to withdraw (0 = withdraw all)
     */
    function withdrawFor(address account, uint amount) external nonReentrant whenNotPaused {
        require(poolStatus != ICrossGameRewardPool.PoolStatus.Paused, CGRPNotAllowedInCurrentState());
        _checkDelegate(account);
        _withdraw(msg.sender, account, amount);
    }

    /**
     * @notice Claims all pending rewards without withdrawing
     * @dev Deposited tokens remain in the pool
     *      Allowed in Active and Inactive states, blocked in Paused state
     *      Can claim even with zero balance if stored rewards exist (for failed transfer recovery)
     */
    function claimRewards() external nonReentrant whenNotPaused {
        _claimAllRewards(msg.sender);
    }

    /**
     * @notice Claims all pending rewards on behalf of another account (Router only)
     * @dev Verifies that msg.sender is the registered router
     *      Allowed in Active and Inactive states, blocked in Paused state
     * @param account Address of the account to claim rewards for
     */
    function claimRewardsFor(address account) external nonReentrant whenNotPaused {
        _checkDelegate(account);
        _claimAllRewards(account);
    }

    /**
     * @notice Claims pending rewards for a specific reward token
     * @dev Can claim rewards even for removed tokens
     *      Allowed in Active and Inactive states, blocked in Paused state
     *      Can claim even with zero balance if stored rewards exist (for failed transfer recovery)
     * @param token Address of the reward token to claim
     */
    function claimReward(IERC20 token) external nonReentrant whenNotPaused {
        _claimSingleReward(msg.sender, token);
    }

    /**
     * @notice Claims pending reward for a specific token on behalf of another account (Router only)
     * @dev Verifies that msg.sender is the registered router
     *      Allowed in Active and Inactive states, blocked in Paused state
     * @param account Address of the account to claim rewards for
     * @param token Address of the reward token to claim
     */
    function claimRewardFor(address account, IERC20 token) external nonReentrant whenNotPaused {
        _checkDelegate(account);
        _claimSingleReward(account, token);
    }

    /**
     * @notice Retrieves pending rewards for a user across all active reward tokens
     * @param user Address of the user to query
     * @return tokens Array of reward token addresses
     * @return rewards Array of pending reward amounts corresponding to each token
     */
    function pendingRewards(address user) external view returns (address[] memory tokens, uint[] memory rewards) {
        uint length = _rewardTokenAddresses.length();
        tokens = new address[](length);
        rewards = new uint[](length);

        for (uint i = 0; i < length; i++) {
            IERC20 token = IERC20(_rewardTokenAddresses.at(i));
            tokens[i] = address(token);
            rewards[i] = _calculatePendingReward(token, user);
        }
    }

    /**
     * @notice Retrieves pending reward for a specific token
     * @param user Address of the user to query
     * @param token Address of the reward token
     * @return amount Pending reward amount for the specified token
     */
    function pendingReward(address user, IERC20 token) external view returns (uint amount) {
        return _calculatePendingReward(token, user);
    }

    /**
     * @notice Retrieves reward token address by index
     * @param index Index in the reward token list
     * @return Address of the reward token at the specified index
     */
    function rewardTokenAt(uint index) external view returns (IERC20) {
        return IERC20(_rewardTokenAddresses.at(index));
    }

    /**
     * @notice Retrieves reward token data
     * @param token Address of the reward token
     * @return Reward token data struct
     */
    function getRewardToken(IERC20 token) external view returns (RewardToken memory) {
        require(_rewardTokenAddresses.contains(address(token)), CGRPInvalidRewardToken(address(token)));
        return _rewardTokenData[token];
    }

    /**
     * @notice Checks if a token is registered as a reward token
     * @param token Address of the token to check
     * @return True if the token is a registered reward token
     */
    function isRewardToken(IERC20 token) external view returns (bool) {
        return _rewardTokenAddresses.contains(address(token));
    }

    /**
     * @notice Retrieves all reward token addresses
     * @return Array of all reward token addresses
     */
    function getRewardTokens() external view returns (address[] memory) {
        return _rewardTokenAddresses.values();
    }

    /**
     * @notice Retrieves the number of reward tokens
     * @return Number of reward tokens
     */
    function rewardTokenCount() external view returns (uint) {
        return _rewardTokenAddresses.length();
    }

    /**
     * @notice Retrieves all removed reward token addresses
     * @return Array of removed reward token addresses
     */
    function getRemovedRewardTokens() external view returns (address[] memory) {
        return _removedRewardTokenAddresses.values();
    }

    /**
     * @notice Retrieves the number of removed reward tokens
     * @return Number of removed reward tokens
     */
    function removedRewardTokenCount() external view returns (uint) {
        return _removedRewardTokenAddresses.length();
    }

    /**
     * @notice Checks if a token is a removed reward token
     * @param token Address of the token to check
     * @return True if the token is a removed reward token
     */
    function isRemovedRewardToken(IERC20 token) external view returns (bool) {
        return _removedRewardTokenAddresses.contains(address(token));
    }

    /**
     * @notice Retrieves user's claimable rewards for removed tokens
     * @dev Calculates actual pending rewards including accumulated but not yet updated amounts
     * @param user Address of the user to query
     * @return tokens Array of removed reward token addresses
     * @return rewards Array of claimable reward amounts for removed tokens
     */
    function getRemovedTokenRewards(address user)
        external
        view
        returns (address[] memory tokens, uint[] memory rewards)
    {
        uint length = _removedRewardTokenAddresses.length();
        tokens = new address[](length);
        rewards = new uint[](length);

        for (uint i = 0; i < length; i++) {
            IERC20 token = IERC20(_removedRewardTokenAddresses.at(i));
            tokens[i] = address(token);

            // Calculate actual pending rewards (including accumulated but not updated)
            UserReward storage ur = userRewards[user][token];
            RewardToken storage rt = _rewardTokenData[token];
            uint userBalance = balances[user];

            // For removed tokens, rewardPerTokenStored is fixed at removal time
            // Calculate earned rewards since last update
            uint earned = _calculateEarned(ur, userBalance, rt.rewardPerTokenStored);
            rewards[i] = ur.rewards + earned;
        }
    }

    // ==================== Admin Functions ====================

    /**
     * @notice Adds a new reward token to the pool
     * @dev Only callable by CrossGameReward contract
     *      Cannot add deposit token as reward token
     *      Cannot add a removed reward token
     * @param token Address of the reward token to add
     */
    function addRewardToken(IERC20 token) external onlyRewardRoot {
        require(address(token) != address(0), CGRPCanNotZeroAddress());
        require(address(token) != address(depositToken), CGRPCanNotUseDepositToken());
        require(!_removedRewardTokenAddresses.contains(address(token)), CGRPInvalidRewardToken(address(token)));
        require(_rewardTokenAddresses.add(address(token)), CGRPRewardTokenAlreadyAdded(address(token)));

        _rewardTokenData[token] = RewardToken({
            token: token,
            rewardPerTokenStored: 0,
            lastBalance: 0,
            reclaimableAmount: 0,
            distributedAmount: 0,
            isRemoved: false
        });

        emit RewardTokenAdded(token);
    }

    /**
     * @notice Removes a reward token from the pool
     * @dev Only callable by CrossGameReward contract
     *      Accumulated rewards can still be claimed after removal
     * @param token Address of the reward token to remove
     */
    function removeRewardToken(IERC20 token) external onlyRewardRoot {
        // Remove from EnumerableSet
        require(_rewardTokenAddresses.remove(address(token)), CGRPInvalidRewardToken(address(token)));

        // Perform final synchronization
        _syncReward(token);

        // Store the actual balance at removal time
        RewardToken storage rt = _rewardTokenData[token];
        uint currentBalance = token.balanceOf(address(this));

        // distributedAmount = amount that users can still claim
        // reclaimableAmount stays as is (owner-reclaimable from zero-deposit rewards)
        rt.distributedAmount = currentBalance - rt.reclaimableAmount;
        rt.isRemoved = true;

        _removedRewardTokenAddresses.add(address(token));

        emit RewardTokenRemoved(token);
    }

    /**
     * @notice Retrieves the amount available for reclaim by owner
     * @dev Reclaimable amount includes:
     *      1. Tokens deposited when totalDeposited was 0 (always in reclaimableAmount)
     *      2. For removed tokens: new deposits after removal (currentBalance - distributedAmount - reclaimableAmount)
     * @param token Address of the reward token
     * @return Amount available for reclaim by owner
     */
    function getReclaimableAmount(IERC20 token) public view returns (uint) {
        RewardToken storage rt = _rewardTokenData[token];
        uint currentBalance = token.balanceOf(address(this));

        // Case 1: Token is removed
        if (rt.isRemoved) {
            // reclaimableAmount: owner-reclaimable (zero-deposit rewards)
            // distributedAmount: user-claimable (distributed rewards)
            // currentBalance - distributedAmount - reclaimableAmount = post-removal deposits
            uint userClaimable = rt.distributedAmount;
            uint ownerReclaimable = rt.reclaimableAmount;

            if (currentBalance > userClaimable + ownerReclaimable) {
                // Post-removal deposits + original reclaimable
                return (currentBalance - userClaimable - ownerReclaimable) + ownerReclaimable;
            }
            return ownerReclaimable;
        }

        // Case 2: Token is active - only zero-deposit rewards are reclaimable
        return rt.reclaimableAmount;
    }

    /**
     * @notice Reclaims excess tokens
     * @dev Only callable by CrossGameReward contract (REWARD_ROOT_ROLE)
     *      Reclaims tokens that cannot be distributed:
     *      - Tokens deposited when totalDeposited was 0
     *      - Tokens deposited after the reward token was removed
     * @param token Address of the reward token
     * @param to Address to receive the reclaimed tokens
     */
    function reclaimTokens(IERC20 token, address to) external onlyRewardRoot {
        uint amount = getReclaimableAmount(token);
        require(amount > 0, CGRPNoReclaimableAmount(address(token)));
        require(to != address(0), CGRPCanNotZeroAddress());

        RewardToken storage rt = _rewardTokenData[token];
        uint currentBalance = token.balanceOf(address(this));

        // Update lastBalance
        rt.lastBalance = currentBalance - amount;

        // Update reclaimableAmount
        // For both active and removed tokens, we're reclaiming from reclaimableAmount
        rt.reclaimableAmount = rt.reclaimableAmount > amount ? rt.reclaimableAmount - amount : 0;

        token.safeTransfer(to, amount);
        emit TokensReclaimed(token, to, amount);
    }

    /**
     * @notice Sets the minimum deposit amount
     * @dev Only callable by CrossGameReward contract
     * @param amount Minimum deposit amount
     */
    function updateMinDepositAmount(uint amount) external onlyRewardRoot {
        require(amount > 0, CGRPCanNotZeroValue());
        emit MinDepositAmountUpdated(minDepositAmount, amount);
        minDepositAmount = amount;
    }

    /**
     * @notice Sets the pool status
     * @dev Only callable by CrossGameReward contract
     *      Active: all operations allowed
     *      Inactive: only claim and withdraw allowed
     *      Paused: no operations allowed (also triggers Pausable pause)
     * @param newStatus New pool status (as uint8 to avoid enum type issues)
     */
    function setPoolStatus(ICrossGameRewardPool.PoolStatus newStatus) external onlyRewardRoot {
        ICrossGameRewardPool.PoolStatus oldStatus = poolStatus;
        require(oldStatus != newStatus, "Pool status unchanged");

        poolStatus = newStatus;

        // Sync Pausable state with PoolStatus
        if (newStatus == ICrossGameRewardPool.PoolStatus.Paused && !paused()) _pause();
        else if (newStatus != ICrossGameRewardPool.PoolStatus.Paused && paused()) _unpause();

        emit PoolStatusChanged(oldStatus, newStatus);
    }

    // ==================== Internal Functions: Reward Synchronization ====================

    /**
     * @dev Synchronizes all reward tokens
     */
    function _syncRewards() internal {
        uint length = _rewardTokenAddresses.length();
        for (uint i = 0; i < length; i++) {
            IERC20 token = IERC20(_rewardTokenAddresses.at(i));
            _syncReward(token);
        }
    }

    /**
     * @dev Detects new rewards and updates rewardPerToken
     * @param token Address of the reward token
     */
    function _syncReward(IERC20 token) internal {
        RewardToken storage rt = _rewardTokenData[token];
        uint currentBalance = rt.token.balanceOf(address(this));

        if (currentBalance > rt.lastBalance) {
            uint newReward = currentBalance - rt.lastBalance;

            // If there's no deposit, treat new rewards as reclaimable amount
            if (totalDeposited == 0) {
                rt.reclaimableAmount += newReward;
            } else {
                // When totalDeposited > 0, only distribute amount beyond reclaimableAmount
                // Calculate distributable amount (current balance minus what's already reclaimable)
                uint totalDistributable =
                    currentBalance > rt.reclaimableAmount ? currentBalance - rt.reclaimableAmount : 0;

                // Only distribute the newly added portion
                if (totalDistributable > rt.lastBalance - rt.reclaimableAmount) {
                    uint distributableReward = totalDistributable - (rt.lastBalance - rt.reclaimableAmount);
                    rt.rewardPerTokenStored += (distributableReward * PRECISION) / totalDeposited;
                    emit RewardSynced(rt.token, distributableReward, totalDeposited);
                }
            }
        }

        rt.lastBalance = currentBalance;
    }

    // ==================== Internal Functions: Reward Updates ====================

    /**
     * @dev Updates user rewards for all reward tokens
     * @param user Address of the user
     */
    function _updateRewards(address user) internal {
        uint length = _rewardTokenAddresses.length();
        for (uint i = 0; i < length; i++) {
            IERC20 token = IERC20(_rewardTokenAddresses.at(i));
            _updateReward(token, user);
        }
    }

    /**
     * @dev Calculates earned rewards for a user (internal helper to avoid duplication)
     * @param ur UserReward storage reference
     * @param userBalance User's current balance
     * @param rewardPerToken Current reward per token value to use
     * @return Earned rewards amount
     */
    function _calculateEarned(UserReward storage ur, uint userBalance, uint rewardPerToken)
        internal
        view
        returns (uint)
    {
        if (userBalance == 0) return 0;
        return (userBalance * (rewardPerToken - ur.rewardPerTokenPaid)) / PRECISION;
    }

    /**
     * @dev Calculates and updates user rewards and checkpoint
     * @param token Address of the reward token
     * @param user Address of the user
     */
    function _updateReward(IERC20 token, address user) internal {
        RewardToken storage rt = _rewardTokenData[token];
        UserReward storage ur = userRewards[user][token];

        uint userBalance = balances[user];
        uint earned = _calculateEarned(ur, userBalance, rt.rewardPerTokenStored);

        if (earned > 0) ur.rewards += earned;

        ur.rewardPerTokenPaid = rt.rewardPerTokenStored;
    }

    /**
     * @dev Calculates pending rewards (view function)
     * @dev Must match _syncReward logic for accurate calculations
     * @param token Address of the reward token
     * @param user Address of the user
     * @return Calculated pending rewards
     */
    function _calculatePendingReward(IERC20 token, address user) internal view returns (uint) {
        UserReward storage ur = userRewards[user][token];
        RewardToken storage rt = _rewardTokenData[token];

        uint userBalance = balances[user];
        if (userBalance == 0) return ur.rewards;

        uint currentRewardPerToken = rt.rewardPerTokenStored;

        // For removed tokens, do not simulate new deposits
        // Only return already accumulated rewards
        if (!rt.isRemoved) {
            uint currentBalance = rt.token.balanceOf(address(this));

            // Simulate _syncReward logic for accurate calculation
            if (currentBalance > rt.lastBalance && totalDeposited > 0) {
                // Only distribute amount beyond reclaimableAmount (match _syncReward logic)
                uint totalDistributable =
                    currentBalance > rt.reclaimableAmount ? currentBalance - rt.reclaimableAmount : 0;

                if (totalDistributable > rt.lastBalance - rt.reclaimableAmount) {
                    uint distributableReward = totalDistributable - (rt.lastBalance - rt.reclaimableAmount);
                    currentRewardPerToken += (distributableReward * PRECISION) / totalDeposited;
                }
            }
        }

        uint earned = _calculateEarned(ur, userBalance, currentRewardPerToken);
        return ur.rewards + earned;
    }

    // ==================== Internal Functions: Reward Claims ====================

    /**
     * @dev Claims all rewards for an account (both active and removed tokens)
     * @param account Address of the account to claim for
     */
    function _claimAllRewards(address account) internal {
        require(poolStatus != ICrossGameRewardPool.PoolStatus.Paused, CGRPNotAllowedInCurrentState());

        uint userBalance = balances[account];
        bool hasRewards = _hasStoredRewards(account);

        require(userBalance > 0 || hasRewards, CGRPNoDepositFound(account));

        // Only sync and update if user has active balance
        // (no point updating rewards when balance is 0)
        if (userBalance > 0) {
            _syncRewards();
            _updateRewards(account);
            _updateRemovedRewards(account);
        }

        // Claim all rewards (both active and removed)
        _claimRewards(account);
        _claimRemovedRewards(account);
    }

    /**
     * @dev Claims a specific reward token for an account
     * @param account Address of the account to claim for
     * @param token Address of the reward token to claim
     */
    function _claimSingleReward(address account, IERC20 token) internal {
        require(poolStatus != ICrossGameRewardPool.PoolStatus.Paused, CGRPNotAllowedInCurrentState());

        uint userBalance = balances[account];
        uint storedReward = userRewards[account][token].rewards;

        require(userBalance > 0 || storedReward > 0, CGRPNoDepositFound(account));
        require(address(_rewardTokenData[token].token) != address(0), CGRPInvalidRewardToken(address(token)));

        // Only sync and update if user has active balance
        if (userBalance > 0) {
            if (_rewardTokenAddresses.contains(address(token))) _syncReward(token);
            _updateReward(token, account);
        }

        _claimReward(token, account);
    }

    /**
     * @dev Claims all reward tokens
     * @param user Address of the user
     */
    function _claimRewards(address user) internal {
        uint length = _rewardTokenAddresses.length();
        for (uint i = 0; i < length; i++) {
            IERC20 token = IERC20(_rewardTokenAddresses.at(i));
            _claimReward(token, user);
        }
    }

    /**
     * @dev Transfers rewards and synchronizes balance
     * @param token Address of the reward token
     * @param user Address of the user
     */
    function _claimReward(IERC20 token, address user) internal {
        UserReward storage ur = userRewards[user][token];
        uint reward = ur.rewards;

        if (reward > 0) {
            RewardToken storage rt = _rewardTokenData[token];

            bool ok = rt.token.trySafeTransfer(user, reward);
            if (!ok) {
                emit RewardClaimFailed(user, rt.token, reward);
            } else {
                ur.rewards = 0;
                rt.lastBalance -= reward;
                // Deduct from distributedAmount if token was removed (user is claiming distributed rewards)
                if (rt.isRemoved) {
                    rt.distributedAmount = rt.distributedAmount > reward ? rt.distributedAmount - reward : 0;
                }
                emit RewardClaimed(user, rt.token, reward);
            }
        }
    }

    function _updateRemovedRewards(address user) private {
        uint length = _removedRewardTokenAddresses.length();
        for (uint i = 0; i < length; i++) {
            IERC20 token = IERC20(_removedRewardTokenAddresses.at(i));
            _updateReward(token, user);
        }
    }

    function _claimRemovedRewards(address user) private {
        uint length = _removedRewardTokenAddresses.length();
        for (uint i = 0; i < length; i++) {
            IERC20 token = IERC20(_removedRewardTokenAddresses.at(i));
            _claimReward(token, user);
        }
    }

    // ==================== Internal Functions: Deposit/Withdraw ====================

    /**
     * @dev Internal deposit logic
     * @param payer Address sending the tokens
     * @param account Address to deposit for
     * @param amount Amount to deposit
     */
    function _deposit(address payer, address account, uint amount) internal {
        require(amount >= minDepositAmount, CGRPBelowMinimumDepositAmount(amount, minDepositAmount));

        _syncRewards();
        _updateRewards(account);
        _updateRemovedRewards(account); // Update removed token rewards before balance changes

        depositToken.safeTransferFrom(payer, address(this), amount);

        balances[account] += amount;
        totalDeposited += amount;

        emit Deposited(account, amount);
    }

    /**
     * @dev Internal withdraw logic
     * @param caller Address receiving the withdrawn tokens
     * @param account Address to withdraw for
     * @param amount Amount to withdraw (0 = withdraw all)
     */
    function _withdraw(address caller, address account, uint amount) internal {
        require(balances[account] > 0, CGRPNoDepositFound(account));

        // If amount is 0, withdraw all
        uint withdrawAmount = amount == 0 ? balances[account] : amount;
        require(withdrawAmount > 0, CGRPCanNotZeroValue());
        require(withdrawAmount <= balances[account], CGRPInsufficientBalance(balances[account], withdrawAmount));

        _syncRewards();
        _updateRewards(account);
        _updateRemovedRewards(account);
        _claimRewards(account);
        _claimRemovedRewards(account);

        balances[account] -= withdrawAmount;
        totalDeposited -= withdrawAmount;
        depositToken.safeTransfer(caller, withdrawAmount);

        emit Withdrawn(account, withdrawAmount);
    }

    /**
     * @dev Validates that the caller is the authorized router
     * @param account Address of the account being acted upon
     */
    function _checkDelegate(address account) internal view {
        require(account != address(0), CGRPCanNotZeroAddress());
        require(msg.sender == ICrossGameReward(crossGameReward).router(), CGRPOnlyRouter());
    }

    // ==================== UUPS ====================

    /**
     * @dev Authorizes contract upgrades
     * @param newImplementation Address of the new implementation
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    // ==================== Storage Gap ====================

    /**
     * @dev Storage gap for future upgrades
     */
    uint[37] private __gap;
}
