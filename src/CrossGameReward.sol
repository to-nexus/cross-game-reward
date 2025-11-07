// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControlDefaultAdminRulesUpgradeable as AccessControl} from
    "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import {Initializable, UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import {IERC5313} from "@openzeppelin/contracts/interfaces/IERC5313.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {CrossGameRewardPool} from "./CrossGameRewardPool.sol";

import {WCROSS} from "./WCROSS.sol";
import {ICrossGameReward} from "./interfaces/ICrossGameReward.sol";
import {ICrossGameRewardPool} from "./interfaces/ICrossGameRewardPool.sol";
import {IWCROSS} from "./interfaces/IWCROSS.sol";
/**
 * @title CrossGameReward
 * @notice Factory contract for managing multiple game reward pools
 * @dev UUPS upgradeable pattern with pool creation and management capabilities
 *      Serves as the central hub for all CrossGameReward pools and WCROSS token
 */

contract CrossGameReward is Initializable, AccessControl, UUPSUpgradeable, ICrossGameReward {
    using EnumerableSet for EnumerableSet.AddressSet;
    using EnumerableSet for EnumerableSet.UintSet;

    // ==================== Roles ====================

    /// @notice Role identifier for pool management operations
    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");

    // ==================== Errors ====================

    /// @notice Thrown when attempting to access a non-existent pool
    error CSPoolNotFound();

    /// @notice Thrown when a zero address is provided where it's not allowed
    error CSCanNotZeroAddress();

    /// @notice Thrown when a zero value is provided where it's not allowed
    error CSCanNotZeroValue();

    // ==================== Events ====================

    /// @notice Emitted when a new game reward pool is created
    /// @param poolId The ID of the newly created pool
    /// @param poolAddress The address of the newly created pool
    /// @param depositToken The token that can be deposited in this pool
    event PoolCreated(uint indexed poolId, address indexed poolAddress, address indexed depositToken);

    /// @notice Emitted when the pool implementation is updated
    /// @param implementation The new implementation address
    event PoolImplementationSet(ICrossGameRewardPool indexed implementation);

    /// @notice Emitted when the router address is set
    /// @param router The new router address
    event RouterSet(address indexed router);

    /// @notice Emitted when owner reclaims tokens from a pool
    /// @param poolId The ID of the pool
    /// @param token The address of the token reclaimed
    /// @param to The address receiving the reclaimed tokens
    /// @param amount The amount reclaimed
    event ReclaimedFromPool(uint indexed poolId, IERC20 indexed token, address indexed to, uint amount);

    // ==================== State Variables ====================

    /// @notice Block number when the contract was initialized
    uint public initializedAt;

    /// @notice Address of the WCROSS token contract
    IWCROSS public wcross;

    /// @notice Address of the router contract for game reward operations
    address public router;

    /// @notice Implementation address for CrossGameRewardPool (UUPS proxy pattern)
    ICrossGameRewardPool public poolImplementation;

    /// @notice Next pool ID to be assigned
    uint public nextPoolId;

    /// @notice Mapping from pool ID to pool information
    mapping(uint => PoolInfo) public pools;

    /// @notice Mapping from pool address to pool ID
    mapping(ICrossGameRewardPool => uint) public poolIds;

    /// @notice Mapping from deposit token to set of pool IDs
    mapping(IERC20 => EnumerableSet.UintSet) private _poolsByDepositToken;

    /// @notice Set of all pool IDs
    EnumerableSet.UintSet private _allPoolIds;

    // ==================== Constructor ====================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initializes the CrossGameReward contract
     * @dev Deploys WCROSS token and sets up initial roles
     * @param _poolImplementation Address of the CrossGameRewardPool implementation contract
     * @param _admin Address of the initial admin
     * @param _initialDelay Delay in seconds for admin role transfers
     */
    function initialize(ICrossGameRewardPool _poolImplementation, address _admin, uint48 _initialDelay)
        external
        initializer
    {
        require(address(_poolImplementation) != address(0), CSCanNotZeroAddress());
        require(_admin != address(0), CSCanNotZeroAddress());

        __AccessControlDefaultAdminRules_init(_initialDelay, _admin);
        __UUPSUpgradeable_init();

        initializedAt = block.number;
        poolImplementation = _poolImplementation;
        wcross = new WCROSS();
        nextPoolId = 1;

        // Grant default roles
        _grantRole(MANAGER_ROLE, _admin);
    }

    // ==================== Pool Management Functions ====================

    /**
     * @notice Creates a new game reward pool
     * @dev Deploys a new UUPS proxy pointing to the pool implementation
     *      Pool's DEFAULT_ADMIN_ROLE references CrossGameReward's DEFAULT_ADMIN_ROLE
     *      CrossGameReward receives REWARD_ROOT_ROLE for pool management
     * @param depositToken Address of the token to be deposited in the pool
     * @param minDepositAmount Minimum amount required for depositing (in wei)
     * @return poolId ID of the newly created pool
     * @return pool Address of the newly created pool
     */
    function createPool(IERC20 depositToken, uint minDepositAmount)
        external
        onlyRole(MANAGER_ROLE)
        returns (uint poolId, ICrossGameRewardPool pool)
    {
        require(address(depositToken) != address(0), CSCanNotZeroAddress());
        require(minDepositAmount > 0, CSCanNotZeroValue());

        poolId = nextPoolId++;

        // Deploy pool as UUPS proxy
        // Pool will set CrossGameReward as msg.sender and get owner from defaultAdmin()
        bytes memory initData = abi.encodeCall(CrossGameRewardPool.initialize, (depositToken, minDepositAmount));

        ERC1967Proxy proxy = new ERC1967Proxy(address(poolImplementation), initData);
        pool = ICrossGameRewardPool(address(proxy));

        // Store pool information
        pools[poolId] = PoolInfo({poolId: poolId, pool: pool, depositToken: depositToken, createdAt: block.timestamp});

        poolIds[pool] = poolId;
        _allPoolIds.add(poolId);
        _poolsByDepositToken[depositToken].add(poolId);

        emit PoolCreated(poolId, address(pool), address(depositToken));
    }

    /**
     * @notice Adds a reward token to a pool
     * @dev Can be called directly since CrossGameReward has REWARD_ROOT_ROLE
     * @param poolId ID of the pool
     * @param token Address of the reward token to add
     */
    function addRewardToken(uint poolId, IERC20 token) external onlyRole(MANAGER_ROLE) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());

        pools[poolId].pool.addRewardToken(token);
        // Event emitted by CrossGameRewardPool
    }

    /**
     * @notice Removes a reward token from a pool
     * @dev Can be called directly since CrossGameReward has REWARD_ROOT_ROLE
     * @param poolId ID of the pool
     * @param token Address of the reward token to remove
     */
    function removeRewardToken(uint poolId, IERC20 token) external onlyRole(MANAGER_ROLE) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());

        pools[poolId].pool.removeRewardToken(token);
        // Event emitted by CrossGameRewardPool
    }

    /**
     * @notice Updates the minimum deposit amount for a pool
     * @dev Only callable by MANAGER_ROLE
     * @param poolId ID of the pool
     * @param amount Minimum deposit amount
     */
    function updateMinDepositAmount(uint poolId, uint amount) external onlyRole(MANAGER_ROLE) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());
        pools[poolId].pool.updateMinDepositAmount(amount);
        // Event emitted by CrossGameRewardPool
    }

    /**
     * @notice Reclaims unallocated rewards from a pool
     * @dev Only callable by MANAGER_ROLE
     *      Can reclaim rewards deposited when totalDeposited was 0 or after token removal
     * @param poolId ID of the pool
     * @param token Address of the reward token
     * @param to Address to receive the reclaimed tokens
     */
    function reclaimFromPool(uint poolId, IERC20 token, address to) external onlyRole(MANAGER_ROLE) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());

        uint amount = pools[poolId].pool.getReclaimableAmount(token);
        pools[poolId].pool.reclaimTokens(token, to);
        emit ReclaimedFromPool(poolId, token, to, amount);
    }

    /**
     * @notice Sets the pool status
     * @dev Controls pool operations based on status:
     *      Active: all operations allowed
     *      Inactive: only claim and withdraw allowed
     *      Paused: no operations allowed
     * @param poolId ID of the pool
     * @param status New pool status
     */
    function setPoolStatus(uint poolId, ICrossGameRewardPool.PoolStatus status) external onlyRole(MANAGER_ROLE) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());

        // Set pool status in the pool contract
        pools[poolId].pool.setPoolStatus(status);
        // Event emitted by CrossGameRewardPool
    }

    /**
     * @notice Updates the pool implementation address
     * @dev Only affects newly created pools, existing pools remain unchanged
     * @param newImplementation Address of the new implementation contract
     */
    function setPoolImplementation(ICrossGameRewardPool newImplementation) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(address(newImplementation) != address(0), CSCanNotZeroAddress());
        poolImplementation = newImplementation;
        emit PoolImplementationSet(newImplementation);
    }

    /**
     * @notice Sets the router address for WCROSS operations
     * @param _router Address of the router contract
     */
    function setRouter(address _router) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_router != address(0), CSCanNotZeroAddress());
        router = _router;
        emit RouterSet(_router);
    }

    // ==================== View Functions ====================

    /**
     * @notice Retrieves pool information by pool ID
     * @param poolId ID of the pool
     * @return Pool information struct
     */
    function getPoolInfo(uint poolId) external view returns (PoolInfo memory) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());
        return pools[poolId];
    }

    /**
     * @notice Retrieves pool ID by index
     * @param index Index in the pool list
     * @return poolId Pool ID at the specified index
     */
    function poolAt(uint index) external view returns (uint) {
        return _allPoolIds.at(index);
    }

    /**
     * @notice Retrieves the number of pools for a specific deposit token
     * @param depositToken Address of the deposit token
     * @return Number of pools using this deposit token
     */
    function getPoolCountByDepositToken(IERC20 depositToken) external view returns (uint) {
        return _poolsByDepositToken[depositToken].length();
    }

    /**
     * @notice Retrieves all pool IDs for a specific deposit token
     * @param depositToken Address of the deposit token
     * @return Array of pool IDs
     */
    function getPoolIdsByDepositToken(IERC20 depositToken) external view returns (uint[] memory) {
        return _poolsByDepositToken[depositToken].values();
    }

    /**
     * @notice Retrieves pool ID by deposit token and index
     * @param depositToken Address of the deposit token
     * @param index Index in the deposit token's pool list
     * @return poolId Pool ID at the specified index
     */
    function poolByDepositTokenAt(IERC20 depositToken, uint index) external view returns (uint) {
        return _poolsByDepositToken[depositToken].at(index);
    }

    /**
     * @notice Retrieves the total number of pools
     * @return Total pool count
     */
    function getTotalPoolCount() external view returns (uint) {
        return _allPoolIds.length();
    }

    /**
     * @notice Retrieves all pool IDs
     * @return Array of all pool IDs
     */
    function getAllPoolIds() external view returns (uint[] memory) {
        return _allPoolIds.values();
    }

    /**
     * @notice Retrieves pool address by pool ID
     * @param poolId ID of the pool
     * @return Address of the pool contract
     */
    function getPoolAddress(uint poolId) external view returns (ICrossGameRewardPool) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());
        return pools[poolId].pool;
    }

    /**
     * @notice Retrieves pool ID by pool address
     * @param pool Address of the pool contract
     * @return Pool ID
     */
    function getPoolId(ICrossGameRewardPool pool) external view returns (uint) {
        uint poolId = poolIds[pool];
        require(poolId != 0, CSPoolNotFound());
        return poolId;
    }

    /**
     * @notice Retrieves only active pool IDs
     * @return Array of active pool IDs
     */
    function getActivePoolIds() external view returns (uint[] memory) {
        uint totalCount = _allPoolIds.length();
        uint[] memory tempIds = new uint[](totalCount);
        uint activeCount = 0;

        for (uint i = 0; i < totalCount; i++) {
            uint poolId = _allPoolIds.at(i);
            if (pools[poolId].pool.poolStatus() == ICrossGameRewardPool.PoolStatus.Active) {
                tempIds[activeCount] = poolId;
                activeCount++;
            }
        }

        // Resize array
        uint[] memory activeIds = new uint[](activeCount);
        for (uint i = 0; i < activeCount; i++) {
            activeIds[i] = tempIds[i];
        }

        return activeIds;
    }

    // ==================== UUPS ====================

    /**
     * @dev Authorizes contract upgrades
     * @param newImplementation Address of the new implementation
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}

    // ==================== Storage Gap ====================

    /**
     * @dev Storage gap for future upgrades
     */
    uint[41] private __gap;
}
