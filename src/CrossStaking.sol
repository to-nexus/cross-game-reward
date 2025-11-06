// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControlDefaultAdminRulesUpgradeable as AccessControl} from
    "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import {Initializable, UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import {IERC5313} from "@openzeppelin/contracts/interfaces/IERC5313.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {CrossStakingPool} from "./CrossStakingPool.sol";

import {WCROSS} from "./WCROSS.sol";
import {ICrossStaking} from "./interfaces/ICrossStaking.sol";
import {ICrossStakingPool} from "./interfaces/ICrossStakingPool.sol";
import {IWCROSS} from "./interfaces/IWCROSS.sol";
/**
 * @title CrossStaking
 * @notice Factory contract for managing multiple staking pools
 * @dev UUPS upgradeable pattern with pool creation and management capabilities
 *      Serves as the central hub for all CrossStaking pools and WCROSS token
 */

contract CrossStaking is Initializable, AccessControl, UUPSUpgradeable, ICrossStaking {
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

    /// @notice Emitted when a new staking pool is created
    /// @param poolId The ID of the newly created pool
    /// @param poolAddress The address of the newly created pool
    /// @param stakingToken The token that can be staked in this pool
    event PoolCreated(uint indexed poolId, address indexed poolAddress, address indexed stakingToken);

    /// @notice Emitted when the pool implementation is updated
    /// @param implementation The new implementation address
    event PoolImplementationSet(ICrossStakingPool indexed implementation);

    /// @notice Emitted when the router address is set
    /// @param router The new router address
    event RouterSet(address indexed router);

    /// @notice Emitted when owner withdraws from a pool
    /// @param poolId The ID of the pool
    /// @param token The address of the token withdrawn
    /// @param to The address receiving the withdrawn tokens
    /// @param amount The amount withdrawn
    event WithdrawnFromPool(uint indexed poolId, IERC20 indexed token, address indexed to, uint amount);

    // ==================== State Variables ====================

    /// @notice Block number when the contract was initialized
    uint public initializedAt;

    /// @notice Address of the WCROSS token contract
    IWCROSS public wcross;

    /// @notice Address of the router contract for staking operations
    address public router;

    /// @notice Implementation address for CrossStakingPool (UUPS proxy pattern)
    ICrossStakingPool public poolImplementation;

    /// @notice Next pool ID to be assigned
    uint public nextPoolId;

    /// @notice Mapping from pool ID to pool information
    mapping(uint => PoolInfo) public pools;

    /// @notice Mapping from pool address to pool ID
    mapping(ICrossStakingPool => uint) public poolIds;

    /// @notice Mapping from staking token to set of pool IDs
    mapping(IERC20 => EnumerableSet.UintSet) private _poolsByStakingToken;

    /// @notice Set of all pool IDs
    EnumerableSet.UintSet private _allPoolIds;

    // ==================== Constructor ====================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initializes the CrossStaking contract
     * @dev Deploys WCROSS token and sets up initial roles
     * @param _poolImplementation Address of the CrossStakingPool implementation contract
     * @param _admin Address of the initial admin
     * @param _initialDelay Delay in seconds for admin role transfers
     */
    function initialize(ICrossStakingPool _poolImplementation, address _admin, uint48 _initialDelay)
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
     * @notice Creates a new staking pool
     * @dev Deploys a new UUPS proxy pointing to the pool implementation
     *      Pool's DEFAULT_ADMIN_ROLE references CrossStaking's DEFAULT_ADMIN_ROLE
     *      CrossStaking receives STAKING_ROOT_ROLE for pool management
     * @param stakingToken Address of the token to be staked in the pool
     * @param minStakeAmount Minimum amount required for staking (in wei)
     * @return poolId ID of the newly created pool
     * @return pool Address of the newly created pool
     */
    function createPool(IERC20 stakingToken, uint minStakeAmount)
        external
        onlyRole(MANAGER_ROLE)
        returns (uint poolId, ICrossStakingPool pool)
    {
        require(address(stakingToken) != address(0), CSCanNotZeroAddress());
        require(minStakeAmount > 0, CSCanNotZeroValue());

        poolId = nextPoolId++;

        // Deploy pool as UUPS proxy
        // Pool will set CrossStaking as msg.sender and get owner from defaultAdmin()
        bytes memory initData = abi.encodeCall(CrossStakingPool.initialize, (stakingToken, minStakeAmount));

        ERC1967Proxy proxy = new ERC1967Proxy(address(poolImplementation), initData);
        pool = ICrossStakingPool(address(proxy));

        // Store pool information
        pools[poolId] = PoolInfo({poolId: poolId, pool: pool, stakingToken: stakingToken, createdAt: block.timestamp});

        poolIds[pool] = poolId;
        _allPoolIds.add(poolId);
        _poolsByStakingToken[stakingToken].add(poolId);

        emit PoolCreated(poolId, address(pool), address(stakingToken));
    }

    /**
     * @notice Adds a reward token to a pool
     * @dev Can be called directly since CrossStaking has STAKING_ROOT_ROLE
     * @param poolId ID of the pool
     * @param token Address of the reward token to add
     */
    function addRewardToken(uint poolId, IERC20 token) external onlyRole(MANAGER_ROLE) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());

        pools[poolId].pool.addRewardToken(token);
        // Event emitted by CrossStakingPool
    }

    /**
     * @notice Removes a reward token from a pool
     * @dev Can be called directly since CrossStaking has STAKING_ROOT_ROLE
     * @param poolId ID of the pool
     * @param token Address of the reward token to remove
     */
    function removeRewardToken(uint poolId, IERC20 token) external onlyRole(MANAGER_ROLE) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());

        pools[poolId].pool.removeRewardToken(token);
        // Event emitted by CrossStakingPool
    }

    /**
     * @notice Updates the minimum stake amount for a pool
     * @dev Only callable by MANAGER_ROLE
     * @param poolId ID of the pool
     * @param amount Minimum stake amount
     */
    function updateMinStakeAmount(uint poolId, uint amount) external onlyRole(MANAGER_ROLE) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());
        pools[poolId].pool.updateMinStakeAmount(amount);
        // Event emitted by CrossStakingPool
    }

    /**
     * @notice Withdraws unallocated rewards from a pool
     * @dev Only callable by MANAGER_ROLE
     *      Can withdraw rewards deposited when totalStaked was 0 or after token removal
     * @param poolId ID of the pool
     * @param token Address of the reward token
     * @param to Address to receive the withdrawn tokens
     */
    function withdrawFromPool(uint poolId, IERC20 token, address to) external onlyRole(MANAGER_ROLE) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());

        uint amount = pools[poolId].pool.getWithdrawableAmount(token);
        pools[poolId].pool.withdraw(token, to);
        emit WithdrawnFromPool(poolId, token, to, amount);
    }

    /**
     * @notice Sets the pool status
     * @dev Controls pool operations based on status:
     *      Active: all operations allowed
     *      Inactive: only claim and unstake allowed
     *      Paused: no operations allowed
     * @param poolId ID of the pool
     * @param status New pool status
     */
    function setPoolStatus(uint poolId, ICrossStakingPool.PoolStatus status) external onlyRole(MANAGER_ROLE) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());

        // Set pool status in the pool contract
        pools[poolId].pool.setPoolStatus(status);
        // Event emitted by CrossStakingPool
    }

    /**
     * @notice Updates the pool implementation address
     * @dev Only affects newly created pools, existing pools remain unchanged
     * @param newImplementation Address of the new implementation contract
     */
    function setPoolImplementation(ICrossStakingPool newImplementation) external onlyRole(DEFAULT_ADMIN_ROLE) {
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
     * @notice Retrieves the number of pools for a specific staking token
     * @param stakingToken Address of the staking token
     * @return Number of pools using this staking token
     */
    function getPoolCountByStakingToken(IERC20 stakingToken) external view returns (uint) {
        return _poolsByStakingToken[stakingToken].length();
    }

    /**
     * @notice Retrieves all pool IDs for a specific staking token
     * @param stakingToken Address of the staking token
     * @return Array of pool IDs
     */
    function getPoolIdsByStakingToken(IERC20 stakingToken) external view returns (uint[] memory) {
        return _poolsByStakingToken[stakingToken].values();
    }

    /**
     * @notice Retrieves pool ID by staking token and index
     * @param stakingToken Address of the staking token
     * @param index Index in the staking token's pool list
     * @return poolId Pool ID at the specified index
     */
    function poolByStakingTokenAt(IERC20 stakingToken, uint index) external view returns (uint) {
        return _poolsByStakingToken[stakingToken].at(index);
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
    function getPoolAddress(uint poolId) external view returns (ICrossStakingPool) {
        require(address(pools[poolId].pool) != address(0), CSPoolNotFound());
        return pools[poolId].pool;
    }

    /**
     * @notice Retrieves pool ID by pool address
     * @param pool Address of the pool contract
     * @return Pool ID
     */
    function getPoolId(ICrossStakingPool pool) external view returns (uint) {
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
            if (pools[poolId].pool.poolStatus() == ICrossStakingPool.PoolStatus.Active) {
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
