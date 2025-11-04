// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {
    AccessControlDefaultAdminRulesUpgradeable as AccessControl,
    AccessControlUpgradeable,
    IAccessControl
} from "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import {Initializable, UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardTransientUpgradeable} from
    "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardTransientUpgradeable.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {ICrossStaking, ICrossStakingPool} from "./interfaces/ICrossStaking.sol";

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
 *
 * === Fairness ===
 *
 * Early stakers:
 *   - Receive rewards for a longer duration
 *   - Each reward is distributed proportionally to current stake amounts
 *
 * === Upgradeability ===
 *
 * UUPS Pattern:
 *   - Managed via AccessControlDefaultAdminRules
 *   - Built-in Pausable functionality
 */
contract CrossStakingPool is
    Initializable,
    AccessControl,
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

    /// @notice Thrown when attempting emergency withdraw with no withdrawable amount
    error CSPNoWithdrawableAmount();

    /// @notice Thrown when attempting to call a function that is not allowed
    error CSPNotAllowedFunction();

    // ==================== Roles ====================

    /// @notice Role identifier for CrossStaking contract (root admin for pool management)
    bytes32 public constant STAKING_ROOT_ROLE = keccak256("STAKING_ROOT_ROLE");

    // ==================== Constants ====================

    /// @notice Precision multiplier for reward calculations
    uint private constant PRECISION = 1e18;

    // ==================== State Variables ====================

    /// @notice The staking token (e.g., CROSS or WCROSS)
    IERC20 public stakingToken;

    /// @notice CrossStaking contract address (used for router validation)
    address public crossStaking;

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

    /// @notice Emitted when admin performs emergency withdrawal
    /// @param token Address of the reward token
    /// @param to Address receiving the withdrawn tokens
    /// @param amount Amount withdrawn
    event EmergencyWithdraw(IERC20 indexed token, address indexed to, uint amount);

    // ==================== Initializer ====================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Fallback function to reject direct calls
     */
    fallback() external {
        revert CSPNotAllowedFunction();
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
        crossStaking = msg.sender;

        // Initialize with 0 delay - admin is CrossStaking's owner (via owner() override)
        __AccessControlDefaultAdminRules_init(0, AccessControl(crossStaking).defaultAdmin());
        __Pausable_init();
        __ReentrancyGuardTransient_init();
        __UUPSUpgradeable_init();

        stakingToken = _stakingToken;
        minStakeAmount = _minStakeAmount;

        // Grant STAKING_ROOT_ROLE to CrossStaking contract for pool management
        _grantRole(STAKING_ROOT_ROLE, crossStaking);
    }

    /**
     * @notice Returns the owner of this pool
     * @dev Overrides AccessControlDefaultAdminRules to return CrossStaking's owner
     *      This ensures DEFAULT_ADMIN_ROLE belongs to CrossStaking's admin, not the contract itself
     * @return Address of the CrossStaking contract's default admin
     */
    function owner() public view override returns (address) {
        return AccessControl(crossStaking).defaultAdmin();
    }

    // ==================== External Functions ====================

    /**
     * @notice Stakes tokens into the pool
     * @dev Additional stakes are accumulated to existing balance
     * @param amount Amount of tokens to stake
     */
    function stake(uint amount) external nonReentrant whenNotPaused {
        _stake(msg.sender, msg.sender, amount);
    }

    /**
     * @notice Stakes tokens on behalf of another account (Router only)
     * @dev Verifies that msg.sender is the registered router
     * @param account Address of the account to stake for
     * @param amount Amount of tokens to stake
     */
    function stakeFor(address account, uint amount) external nonReentrant whenNotPaused {
        _checkDelegate(account);
        _stake(msg.sender, account, amount);
    }

    /**
     * @notice Unstakes all staked tokens and claims all rewards
     * @dev Automatically claims all accumulated rewards
     */
    function unstake() external nonReentrant whenNotPaused {
        _unstake(msg.sender, msg.sender);
    }

    /**
     * @notice Unstakes tokens on behalf of another account (Router only)
     * @dev Verifies that msg.sender is the registered router
     * @param account Address of the account to unstake for
     */
    function unstakeFor(address account) external nonReentrant whenNotPaused {
        _checkDelegate(account);
        _unstake(msg.sender, account);
    }

    /**
     * @notice Claims all pending rewards without unstaking
     * @dev Staked tokens remain in the pool
     */
    function claimRewards() external nonReentrant whenNotPaused {
        require(balances[msg.sender] > 0, CSPNoStakeFound());

        _syncRewards();
        _updateRewards(msg.sender);
        _claimRewards(msg.sender);
    }

    /**
     * @notice Claims pending rewards for a specific reward token
     * @dev Can claim rewards even for removed tokens
     * @param token Address of the reward token to claim
     */
    function claimReward(IERC20 token) external nonReentrant whenNotPaused {
        require(balances[msg.sender] > 0, CSPNoStakeFound());
        // Allow claiming even for removed tokens by checking only _rewardTokenData existence
        require(address(_rewardTokenData[token].token) != address(0), CSPInvalidRewardToken());

        // Only sync for tokens that haven't been removed
        if (_rewardTokenAddresses.contains(address(token))) _syncReward(token);

        _updateReward(token, msg.sender);
        _claimReward(token, msg.sender);
    }

    /**
     * @notice Retrieves pending rewards for a user
     * @param user Address of the user to query
     * @return rewards Array of pending reward amounts for each reward token
     */
    function pendingRewards(address user) external view returns (uint[] memory rewards) {
        uint length = _rewardTokenAddresses.length();
        rewards = new uint[](length);

        for (uint i = 0; i < length;) {
            IERC20 token = IERC20(_rewardTokenAddresses.at(i));
            rewards[i] = _calculatePendingReward(token, user);
            unchecked {
                ++i;
            }
        }
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

    // ==================== Admin Functions ====================

    /**
     * @notice Adds a new reward token to the pool
     * @dev Only callable by CrossStaking contract
     *      Cannot add staking token as reward token
     * @param token Address of the reward token to add
     */
    function addRewardToken(IERC20 token) external onlyRole(STAKING_ROOT_ROLE) {
        require(address(token) != address(0), CSPCanNotZeroAddress());
        require(address(token) != address(stakingToken), CSPCanNotUseStakingToken());
        require(_rewardTokenAddresses.add(address(token)), CSPRewardTokenAlreadyAdded());
        _removedRewardTokenAddresses.remove(address(token));

        _rewardTokenData[token] = RewardToken({
            token: token,
            rewardPerTokenStored: 0,
            lastBalance: 0,
            removedDistributedAmount: 0,
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
    function removeRewardToken(IERC20 token) external onlyRole(STAKING_ROOT_ROLE) {
        // Remove from EnumerableSet
        require(_rewardTokenAddresses.remove(address(token)), CSPInvalidRewardToken());

        // Perform final synchronization
        _syncReward(token);

        // Store the actual balance at removal time
        RewardToken storage rt = _rewardTokenData[token];
        uint currentBalance = token.balanceOf(address(this));
        rt.removedDistributedAmount = currentBalance;
        rt.isRemoved = true;

        _removedRewardTokenAddresses.add(address(token));

        emit RewardTokenRemoved(token);
    }

    /**
     * @notice Retrieves the amount available for emergency withdrawal
     * @dev Only tokens deposited after removal can be withdrawn
     * @param token Address of the reward token
     * @return Amount available for emergency withdrawal
     */
    function getEmergencyWithdrawableAmount(IERC20 token) public view returns (uint) {
        RewardToken storage rt = _rewardTokenData[token];
        if (!rt.isRemoved) return 0; // Token not removed

        uint currentBalance = IERC20(token).balanceOf(address(this));
        return currentBalance > rt.removedDistributedAmount ? currentBalance - rt.removedDistributedAmount : 0;
    }

    /**
     * @notice Performs emergency withdrawal of excess tokens
     * @dev Only callable by CrossStaking's DEFAULT_ADMIN (direct call)
     *      Only withdraws tokens deposited after the token was removed
     * @param token Address of the reward token
     * @param to Address to receive the withdrawn tokens
     */
    function emergencyWithdraw(IERC20 token, address to) external onlyRole(DEFAULT_ADMIN_ROLE) {
        uint amount = getEmergencyWithdrawableAmount(token);
        require(amount > 0, CSPNoWithdrawableAmount());
        require(to != address(0), CSPCanNotZeroAddress());

        IERC20(token).safeTransfer(to, amount);
        emit EmergencyWithdraw(token, to, amount);
    }

    /**
     * @notice Pauses the pool
     * @dev Only callable by CrossStaking contract
     *      Blocks stake, unstake, and claim operations
     */
    function pause() external onlyRole(STAKING_ROOT_ROLE) {
        _pause();
    }

    /**
     * @notice Unpauses the pool
     * @dev Only callable by CrossStaking contract
     *      Re-enables stake, unstake, and claim operations
     */
    function unpause() external onlyRole(STAKING_ROOT_ROLE) {
        _unpause();
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
        // Don't synchronize if there's no staking
        if (totalStaked == 0) return;

        RewardToken storage rt = _rewardTokenData[token];

        uint currentBalance = rt.token.balanceOf(address(this));

        if (currentBalance > rt.lastBalance) {
            uint newReward = currentBalance - rt.lastBalance;
            rt.rewardPerTokenStored += (newReward * PRECISION) / totalStaked;
            emit RewardSynced(rt.token, newReward, totalStaked);
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
     * @dev Calculates and updates user rewards and checkpoint
     * @param token Address of the reward token
     * @param user Address of the user
     */
    function _updateReward(IERC20 token, address user) internal {
        RewardToken storage rt = _rewardTokenData[token];
        UserReward storage ur = userRewards[user][token];

        uint userBalance = balances[user];

        if (userBalance > 0) {
            uint earned = (userBalance * (rt.rewardPerTokenStored - ur.rewardPerTokenPaid)) / PRECISION;
            ur.rewards += earned;
        }

        ur.rewardPerTokenPaid = rt.rewardPerTokenStored;
    }

    /**
     * @dev Calculates pending rewards (view function)
     * @param token Address of the reward token
     * @param user Address of the user
     * @return Calculated pending rewards
     */
    function _calculatePendingReward(IERC20 token, address user) internal view returns (uint) {
        UserReward storage ur = userRewards[user][token];
        RewardToken storage rt = _rewardTokenData[token];

        uint userBalance = balances[user];
        if (userBalance == 0) return ur.rewards;

        uint currentBalance = IERC20(rt.token).balanceOf(address(this));
        uint currentRewardPerToken = rt.rewardPerTokenStored;

        if (currentBalance > rt.lastBalance && totalStaked > 0) {
            uint newReward = currentBalance - rt.lastBalance;
            currentRewardPerToken += (newReward * PRECISION) / totalStaked;
        }

        uint earned = (userBalance * (currentRewardPerToken - ur.rewardPerTokenPaid)) / PRECISION;
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

            // Deduct from removedDistributedAmount if token was removed
            if (rt.isRemoved) rt.removedDistributedAmount -= reward;

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
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}

    /**
     * @dev Overrides the hasRole function to return the owner of the pool
     * @param role The role to check
     * @param account The account to check
     * @return True if the account has the role, false otherwise
     */
    function hasRole(bytes32 role, address account)
        public
        view
        override(AccessControlUpgradeable, IAccessControl)
        returns (bool)
    {
        if (role == DEFAULT_ADMIN_ROLE) return account == owner();
        return super.hasRole(role, account);
    }

    // ==================== Storage Gap ====================

    /**
     * @dev Storage gap for future upgrades
     *      Currently used: 9 slots (stakingToken, crossStaking, minStakeAmount, _rewardTokenAddresses, _removedRewardTokenAddresses, _rewardTokenData, balances, userRewards, totalStaked)
     *      Gap: 50 - 9 = 41 slots
     */
    uint[41] private __gap;
}
