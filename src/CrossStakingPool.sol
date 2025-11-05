// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Initializable, UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardTransientUpgradeable} from
    "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardTransientUpgradeable.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {ICrossStaking} from "./interfaces/ICrossStaking.sol";
import {ICrossStakingPool} from "./interfaces/ICrossStakingPool.sol";

/**
 * @title CrossStakingPool
 * @notice Staking pool with real-time reward distribution mechanism
 * @dev Implements rewardPerToken accumulation pattern with UUPS upgradeability
 *
 * === Core Principles ===
 *
 * Reward Distribution:
 *   - Users only receive rewards deposited after their stake
 *   - Rewards are distributed proportionally based on stake amount at deposit time
 *   - Uses cumulative rewardPerToken accounting
 *   - Zero-stake deposits are marked as withdrawable (owner can reclaim)
 *
 * === Fairness ===
 *
 * Early stakers:
 *   - Receive rewards for a longer duration
 *   - Each reward is distributed proportionally to current stake amounts
 *   - Do NOT receive rewards deposited before their stake
 *
 * === Upgradeability ===
 *
 * UUPS Pattern:
 *   - Owner-only upgrade authorization
 *   - Simplified access control (onlyOwner/onlyStakingRoot)
 *   - Built-in Pausable functionality via pool status
 */
contract CrossStakingPool is
    Initializable,
    PausableUpgradeable,
    ReentrancyGuardTransientUpgradeable,
    UUPSUpgradeable,
    ICrossStakingPool
{
    using SafeERC20 for IERC20;
    using EnumerableSet for EnumerableSet.AddressSet;

    // ==================== Custom Errors ====================

    /// @notice Thrown when stake amount is below minimum required
    error CSPBelowMinimumStakeAmount();

    /// @notice Thrown when attempting to unstake with no active stake
    error CSPNoStakeFound();

    /// @notice Thrown when a zero address is provided where it's not allowed
    error CSPCanNotZeroAddress();

    /// @notice Thrown when a zero value is provided where it's not allowed
    error CSPCanNotZeroValue();

    /// @notice Thrown when attempting to add an already existing reward token
    error CSPRewardTokenAlreadyAdded();

    /// @notice Thrown when accessing an invalid reward token
    error CSPInvalidRewardToken();

    /// @notice Thrown when attempting to use staking token as reward token
    error CSPCanNotUseStakingToken();

    /// @notice Thrown when caller is not the authorized router
    error CSPOnlyRouter();

    /// @notice Thrown when attempting withdraw with no withdrawable amount
    error CSPNoWithdrawableAmount();

    /// @notice Thrown when attempting to call a function that is not allowed
    error CSPNotAllowedFunction();

    /// @notice Thrown when caller is not the owner of the pool
    error CSPOnlyOwner();

    /// @notice Thrown when caller is not the staking root
    error CSPOnlyStakingRoot();

    /// @notice Thrown when attempting to stake in an inactive or paused pool
    error CSPCannotStakeInCurrentState();

    /// @notice Thrown when attempting an operation not allowed in current pool state
    error CSPOperationNotAllowedInCurrentState();

    // ==================== Constants ====================

    /// @notice Precision multiplier for reward calculations
    uint private constant PRECISION = 1e18;

    // ==================== State Variables ====================

    /// @notice The staking token (e.g., CROSS or WCROSS)
    IERC20 public stakingToken;

    /// @notice CrossStaking contract address (used for router validation)
    ICrossStaking public crossStaking;

    /// @notice Minimum amount required for staking
    uint public minStakeAmount;

    /// @notice Set of active reward token addresses
    EnumerableSet.AddressSet private _rewardTokenAddresses;

    /// @notice Set of removed reward token addresses (still claimable)
    EnumerableSet.AddressSet private _removedRewardTokenAddresses;

    /// @notice Mapping from reward token to reward token data
    mapping(IERC20 => RewardToken) private _rewardTokenData;

    /// @notice Mapping from user address to staked balance
    mapping(address => uint) public balances;

    /// @notice Mapping from user address to reward token to user reward data
    mapping(address => mapping(IERC20 => UserReward)) public userRewards;

    /// @notice Total amount of tokens staked in the pool
    uint public totalStaked;

    /// @notice Current status of the pool
    ICrossStakingPool.PoolStatus public poolStatus;

    // ==================== Events ====================

    /// @notice Emitted when a account stakes tokens
    /// @param account Address of the account who staked
    /// @param amount Amount of tokens staked
    event Staked(address indexed account, uint amount);

    /// @notice Emitted when a account unstakes tokens
    /// @param account Address of the account who unstaked
    /// @param amount Amount of tokens unstaked
    event Unstaked(address indexed account, uint amount);

    /// @notice Emitted when a account claims rewards
    /// @param account Address of the account who claimed
    /// @param token Address of the reward token claimed
    /// @param amount Amount of reward tokens claimed
    event RewardClaimed(address indexed account, IERC20 indexed token, uint amount);

    /// @notice Emitted when a new reward token is added
    /// @param token Address of the added reward token
    event RewardTokenAdded(IERC20 indexed token);

    /// @notice Emitted when a reward token is removed
    /// @param token Address of the removed reward token
    event RewardTokenRemoved(IERC20 indexed token);

    /// @notice Emitted when rewards are synced to the pool
    /// @param token Address of the reward token
    /// @param amount Amount of rewards added to the pool
    /// @param totalStaked Total amount of tokens staked in the pool
    event RewardSynced(IERC20 indexed token, uint amount, uint totalStaked);

    /// @notice Emitted when admin performs withdrawal
    /// @param token Address of the reward token
    /// @param to Address receiving the withdrawn tokens
    /// @param amount Amount withdrawn
    event Withdraw(IERC20 indexed token, address indexed to, uint amount);

    /// @notice Emitted when the minimum stake amount is updated
    /// @param oldAmount Old minimum stake amount
    /// @param newAmount New minimum stake amount
    event MinStakeAmountUpdated(uint oldAmount, uint newAmount);

    /// @notice Emitted when the pool status changes
    /// @param oldStatus Old pool status
    /// @param newStatus New pool status
    event PoolStatusChanged(ICrossStakingPool.PoolStatus oldStatus, ICrossStakingPool.PoolStatus newStatus);

    // ==================== Initializer ====================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    modifier onlyOwner() {
        require(msg.sender == owner(), CSPOnlyOwner());
        _;
    }

    modifier onlyStakingRoot() {
        require(msg.sender == address(crossStaking), CSPOnlyStakingRoot());
        _;
    }

    /**
     * @notice Initializes the CrossStakingPool contract
     * @dev Sets up roles and initializes inherited contracts
     *      msg.sender must be CrossStaking contract
     *      Pool's DEFAULT_ADMIN_ROLE = CrossStaking's owner (via owner() override)
     *      STAKING_ROOT_ROLE is granted to CrossStaking contract for pool management
     * @param _stakingToken Address of the token to be staked
     * @param _minStakeAmount Minimum amount required for staking
     */
    function initialize(IERC20 _stakingToken, uint _minStakeAmount) external initializer {
        require(address(_stakingToken) != address(0), CSPCanNotZeroAddress());
        require(_minStakeAmount > 0, CSPCanNotZeroValue());

        // msg.sender is always CrossStaking contract
        crossStaking = ICrossStaking(msg.sender);

        __Pausable_init();
        __ReentrancyGuardTransient_init();
        __UUPSUpgradeable_init();

        stakingToken = _stakingToken;
        minStakeAmount = _minStakeAmount;
        poolStatus = ICrossStakingPool.PoolStatus.Active;
    }

    /**
     * @notice Returns the owner of this pool
     * @dev Returns CrossStaking's owner (default admin)
     *      This ensures pool owner is the CrossStaking contract's admin
     * @return Address of the CrossStaking contract's default admin
     */
    function owner() public view returns (address) {
        return crossStaking.owner();
    }

    // ==================== External Functions ====================

    /**
     * @notice Stakes tokens into the pool
     * @dev Additional stakes are accumulated to existing balance
     *      Only allowed when pool is Active
     * @param amount Amount of tokens to stake
     */
    function stake(uint amount) external nonReentrant whenNotPaused {
        require(poolStatus == ICrossStakingPool.PoolStatus.Active, CSPCannotStakeInCurrentState());
        _stake(msg.sender, msg.sender, amount);
    }

    /**
     * @notice Stakes tokens on behalf of another account (Router only)
     * @dev Verifies that msg.sender is the registered router
     *      Only allowed when pool is Active
     * @param account Address of the account to stake for
     * @param amount Amount of tokens to stake
     */
    function stakeFor(address account, uint amount) external nonReentrant whenNotPaused {
        require(poolStatus == ICrossStakingPool.PoolStatus.Active, CSPCannotStakeInCurrentState());
        _checkDelegate(account);
        _stake(msg.sender, account, amount);
    }

    /**
     * @notice Unstakes all staked tokens and claims all rewards
     * @dev Automatically claims all accumulated rewards
     *      Allowed in Active and Inactive states, blocked in Paused state
     */
    function unstake() external nonReentrant whenNotPaused {
        require(poolStatus != ICrossStakingPool.PoolStatus.Paused, CSPOperationNotAllowedInCurrentState());
        _unstake(msg.sender, msg.sender);
    }

    /**
     * @notice Unstakes tokens on behalf of another account (Router only)
     * @dev Verifies that msg.sender is the registered router
     *      Allowed in Active and Inactive states, blocked in Paused state
     * @param account Address of the account to unstake for
     */
    function unstakeFor(address account) external nonReentrant whenNotPaused {
        require(poolStatus != ICrossStakingPool.PoolStatus.Paused, CSPOperationNotAllowedInCurrentState());
        _checkDelegate(account);
        _unstake(msg.sender, account);
    }

    /**
     * @notice Claims all pending rewards without unstaking
     * @dev Staked tokens remain in the pool
     *      Allowed in Active and Inactive states, blocked in Paused state
     */
    function claimRewards() external nonReentrant whenNotPaused {
        require(poolStatus != ICrossStakingPool.PoolStatus.Paused, CSPOperationNotAllowedInCurrentState());
        require(balances[msg.sender] > 0, CSPNoStakeFound());

        _syncRewards();
        _updateRewards(msg.sender);
        _claimRewards(msg.sender);
    }

    /**
     * @notice Claims pending rewards for a specific reward token
     * @dev Can claim rewards even for removed tokens
     *      Allowed in Active and Inactive states, blocked in Paused state
     * @param token Address of the reward token to claim
     */
    function claimReward(IERC20 token) external nonReentrant whenNotPaused {
        require(poolStatus != ICrossStakingPool.PoolStatus.Paused, CSPOperationNotAllowedInCurrentState());
        require(balances[msg.sender] > 0, CSPNoStakeFound());
        // Allow claiming even for removed tokens by checking only _rewardTokenData existence
        require(address(_rewardTokenData[token].token) != address(0), CSPInvalidRewardToken());

        // Only sync for tokens that haven't been removed
        if (_rewardTokenAddresses.contains(address(token))) _syncReward(token);

        _updateReward(token, msg.sender);
        _claimReward(token, msg.sender);
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

        for (uint i = 0; i < length;) {
            IERC20 token = IERC20(_rewardTokenAddresses.at(i));
            tokens[i] = address(token);
            rewards[i] = _calculatePendingReward(token, user);
            unchecked {
                ++i;
            }
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
        require(_rewardTokenAddresses.contains(address(token)), CSPInvalidRewardToken());
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

        for (uint i = 0; i < length;) {
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

            unchecked {
                ++i;
            }
        }
    }

    // ==================== Admin Functions ====================

    /**
     * @notice Adds a new reward token to the pool
     * @dev Only callable by CrossStaking contract
     *      Cannot add staking token as reward token
     * @param token Address of the reward token to add
     */
    function addRewardToken(IERC20 token) external onlyStakingRoot {
        require(address(token) != address(0), CSPCanNotZeroAddress());
        require(address(token) != address(stakingToken), CSPCanNotUseStakingToken());
        require(_rewardTokenAddresses.add(address(token)), CSPRewardTokenAlreadyAdded());
        _removedRewardTokenAddresses.remove(address(token));

        _rewardTokenData[token] = RewardToken({
            token: token,
            rewardPerTokenStored: 0,
            lastBalance: 0,
            withdrawableAmount: 0,
            distributedAmount: 0,
            isRemoved: false
        });

        emit RewardTokenAdded(token);
    }

    /**
     * @notice Removes a reward token from the pool
     * @dev Only callable by CrossStaking contract
     *      Accumulated rewards can still be claimed after removal
     * @param token Address of the reward token to remove
     */
    function removeRewardToken(IERC20 token) external onlyStakingRoot {
        // Remove from EnumerableSet
        require(_rewardTokenAddresses.remove(address(token)), CSPInvalidRewardToken());

        // Perform final synchronization
        _syncReward(token);

        // Store the actual balance at removal time
        RewardToken storage rt = _rewardTokenData[token];
        uint currentBalance = token.balanceOf(address(this));

        // distributedAmount = amount that users can still claim
        // withdrawableAmount stays as is (owner-withdrawable from zero-stake deposits)
        rt.distributedAmount = currentBalance - rt.withdrawableAmount;
        rt.isRemoved = true;

        _removedRewardTokenAddresses.add(address(token));

        emit RewardTokenRemoved(token);
    }

    /**
     * @notice Retrieves the amount available for withdrawal by owner
     * @dev Withdrawable amount includes:
     *      1. Tokens deposited when totalStaked was 0 (always in withdrawableAmount)
     *      2. For removed tokens: new deposits after removal (currentBalance - distributedAmount - withdrawableAmount)
     * @param token Address of the reward token
     * @return Amount available for withdrawal by owner
     */
    function getWithdrawableAmount(IERC20 token) public view returns (uint) {
        RewardToken storage rt = _rewardTokenData[token];
        uint currentBalance = token.balanceOf(address(this));

        // Case 1: Token is removed
        if (rt.isRemoved) {
            // withdrawableAmount: owner-withdrawable (zero-stake deposits)
            // distributedAmount: user-claimable (distributed rewards)
            // currentBalance - distributedAmount - withdrawableAmount = post-removal deposits
            uint userClaimable = rt.distributedAmount;
            uint ownerWithdrawable = rt.withdrawableAmount;

            if (currentBalance > userClaimable + ownerWithdrawable) {
                // Post-removal deposits + original withdrawable
                return (currentBalance - userClaimable - ownerWithdrawable) + ownerWithdrawable;
            }
            return ownerWithdrawable;
        }

        // Case 2: Token is active - only zero-stake deposits are withdrawable
        return rt.withdrawableAmount;
    }

    /**
     * @notice Performs withdrawal of excess tokens
     * @dev Only callable by CrossStaking contract (STAKING_ROOT_ROLE)
     *      Withdraws tokens that cannot be distributed:
     *      - Tokens deposited when totalStaked was 0
     *      - Tokens deposited after the reward token was removed
     * @param token Address of the reward token
     * @param to Address to receive the withdrawn tokens
     */
    function withdraw(IERC20 token, address to) external onlyStakingRoot {
        uint amount = getWithdrawableAmount(token);
        require(amount > 0, CSPNoWithdrawableAmount());
        require(to != address(0), CSPCanNotZeroAddress());

        RewardToken storage rt = _rewardTokenData[token];
        uint currentBalance = token.balanceOf(address(this));

        // Update lastBalance
        rt.lastBalance = currentBalance - amount;

        // Update withdrawableAmount
        // For both active and removed tokens, we're withdrawing from withdrawableAmount
        rt.withdrawableAmount = rt.withdrawableAmount > amount ? rt.withdrawableAmount - amount : 0;

        token.safeTransfer(to, amount);
        emit Withdraw(token, to, amount);
    }

    /**
     * @notice Sets the minimum stake amount
     * @dev Only callable by CrossStaking contract
     * @param amount Minimum stake amount
     */
    function updateMinStakeAmount(uint amount) external onlyStakingRoot {
        require(amount > 0, CSPCanNotZeroValue());
        emit MinStakeAmountUpdated(minStakeAmount, amount);
        minStakeAmount = amount;
    }

    /**
     * @notice Sets the pool status
     * @dev Only callable by CrossStaking contract
     *      Active: all operations allowed
     *      Inactive: only claim and unstake allowed
     *      Paused: no operations allowed (also triggers Pausable pause)
     * @param newStatus New pool status (as uint8 to avoid enum type issues)
     */
    function setPoolStatus(ICrossStakingPool.PoolStatus newStatus) external onlyStakingRoot {
        ICrossStakingPool.PoolStatus oldStatus = poolStatus;
        require(oldStatus != newStatus, "Pool status unchanged");

        poolStatus = newStatus;

        // Sync Pausable state with PoolStatus
        if (newStatus == ICrossStakingPool.PoolStatus.Paused && !paused()) _pause();
        else if (newStatus != ICrossStakingPool.PoolStatus.Paused && paused()) _unpause();

        emit PoolStatusChanged(oldStatus, newStatus);
    }

    // ==================== Internal Functions: Reward Synchronization ====================

    /**
     * @dev Synchronizes all reward tokens
     */
    function _syncRewards() internal {
        uint length = _rewardTokenAddresses.length();
        for (uint i = 0; i < length;) {
            IERC20 token = IERC20(_rewardTokenAddresses.at(i));
            _syncReward(token);
            unchecked {
                ++i;
            }
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

            // If there's no staking, treat new rewards as withdrawable amount
            if (totalStaked == 0) {
                rt.withdrawableAmount += newReward;
            } else {
                // When totalStaked > 0, only distribute amount beyond withdrawableAmount
                // Calculate distributable amount (current balance minus what's already withdrawable)
                uint totalDistributable =
                    currentBalance > rt.withdrawableAmount ? currentBalance - rt.withdrawableAmount : 0;

                // Only distribute the newly added portion
                if (totalDistributable > rt.lastBalance - rt.withdrawableAmount) {
                    uint distributableReward = totalDistributable - (rt.lastBalance - rt.withdrawableAmount);
                    rt.rewardPerTokenStored += (distributableReward * PRECISION) / totalStaked;
                    emit RewardSynced(rt.token, distributableReward, totalStaked);
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
        for (uint i = 0; i < length;) {
            IERC20 token = IERC20(_rewardTokenAddresses.at(i));
            _updateReward(token, user);
            unchecked {
                ++i;
            }
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
            if (currentBalance > rt.lastBalance && totalStaked > 0) {
                // Only distribute amount beyond withdrawableAmount (match _syncReward logic)
                uint totalDistributable =
                    currentBalance > rt.withdrawableAmount ? currentBalance - rt.withdrawableAmount : 0;

                if (totalDistributable > rt.lastBalance - rt.withdrawableAmount) {
                    uint distributableReward = totalDistributable - (rt.lastBalance - rt.withdrawableAmount);
                    currentRewardPerToken += (distributableReward * PRECISION) / totalStaked;
                }
            }
        }

        uint earned = _calculateEarned(ur, userBalance, currentRewardPerToken);
        return ur.rewards + earned;
    }

    // ==================== Internal Functions: Reward Claims ====================

    /**
     * @dev Claims all reward tokens
     * @param user Address of the user
     */
    function _claimRewards(address user) internal {
        uint length = _rewardTokenAddresses.length();
        for (uint i = 0; i < length;) {
            IERC20 token = IERC20(_rewardTokenAddresses.at(i));
            _claimReward(token, user);
            unchecked {
                ++i;
            }
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
            ur.rewards = 0;

            RewardToken storage rt = _rewardTokenData[token];
            rt.token.safeTransfer(user, reward);

            rt.lastBalance -= reward;

            // Deduct from distributedAmount if token was removed (user is claiming distributed rewards)
            if (rt.isRemoved) rt.distributedAmount = rt.distributedAmount > reward ? rt.distributedAmount - reward : 0;

            emit RewardClaimed(user, rt.token, reward);
        }
    }

    function _updateRemovedRewards(address user) private {
        uint length = _removedRewardTokenAddresses.length();
        for (uint i = 0; i < length;) {
            IERC20 token = IERC20(_removedRewardTokenAddresses.at(i));
            _updateReward(token, user);
            unchecked {
                ++i;
            }
        }
    }

    function _claimRemovedRewards(address user) private {
        uint length = _removedRewardTokenAddresses.length();
        for (uint i = 0; i < length;) {
            IERC20 token = IERC20(_removedRewardTokenAddresses.at(i));
            _claimReward(token, user);
            unchecked {
                ++i;
            }
        }
    }

    // ==================== Internal Functions: Stake/Unstake ====================

    /**
     * @dev Internal staking logic
     * @param payer Address sending the tokens
     * @param account Address to stake for
     * @param amount Amount to stake
     */
    function _stake(address payer, address account, uint amount) internal {
        require(amount >= minStakeAmount, CSPBelowMinimumStakeAmount());

        _syncRewards();
        _updateRewards(account);
        _updateRemovedRewards(account); // Update removed token rewards before balance changes

        stakingToken.safeTransferFrom(payer, address(this), amount);

        balances[account] += amount;
        totalStaked += amount;

        emit Staked(account, amount);
    }

    /**
     * @dev Internal unstaking logic
     * @param caller Address receiving the unstaked tokens
     * @param account Address to unstake for
     */
    function _unstake(address caller, address account) internal {
        require(balances[account] > 0, CSPNoStakeFound());

        uint amount = balances[account];

        _syncRewards();
        _updateRewards(account);
        _updateRemovedRewards(account);
        _claimRewards(account);
        _claimRemovedRewards(account);

        totalStaked -= amount;
        stakingToken.safeTransfer(caller, amount);

        delete balances[account];

        emit Unstaked(account, amount);
    }

    /**
     * @dev Validates that the caller is the authorized router
     * @param account Address of the account being acted upon
     */
    function _checkDelegate(address account) internal view {
        require(account != address(0), CSPCanNotZeroAddress());
        require(msg.sender == ICrossStaking(crossStaking).router(), CSPOnlyRouter());
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
    uint[38] private __gap;
}
