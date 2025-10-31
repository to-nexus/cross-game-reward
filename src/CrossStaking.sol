// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {CrossStakingPool} from "./CrossStakingPool.sol";
import {WCROSS} from "./WCROSS.sol";

import {ICrossStaking} from "./interfaces/ICrossStaking.sol";
import {AccessControlDefaultAdminRulesUpgradeable} from
    "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @title CrossStaking
 * @notice Factory contract for managing multiple staking pools
 * @dev UUPS upgradeable pattern with pool creation and management capabilities
 *      Serves as the central hub for all CrossStaking pools and WCROSS token
 */
contract CrossStaking is Initializable, AccessControlDefaultAdminRulesUpgradeable, UUPSUpgradeable, ICrossStaking {
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

    // ==================== Events ====================

    /// @notice Emitted when a new staking pool is created
    /// @param poolId The ID of the newly created pool
    /// @param poolAddress The address of the newly created pool
    /// @param stakingToken The token that can be staked in this pool
    event PoolCreated(uint indexed poolId, address indexed poolAddress, address indexed stakingToken);

    /// @notice Emitted when the pool implementation is updated
    /// @param implementation The new implementation address
    event PoolImplementationSet(address indexed implementation);

    /// @notice Emitted when a pool's active status changes
    /// @param poolId The ID of the pool
    /// @param active The new active status
    event PoolStatusChanged(uint indexed poolId, bool active);

    /// @notice Emitted when the router address is set
    /// @param router The new router address
    event RouterSet(address indexed router);

    // ==================== State Variables ====================

    /// @notice Address of the WCROSS token contract
    address public wcross;

    /// @notice Address of the router contract for staking operations
    address public router;

    /// @notice Implementation address for CrossStakingPool (UUPS proxy pattern)
    address public poolImplementation;

    /// @notice Next pool ID to be assigned
    uint public nextPoolId;

    /// @notice Mapping from pool ID to pool information
    mapping(uint => PoolInfo) public pools;

    /// @notice Mapping from pool address to pool ID
    mapping(address => uint) public poolIds;

    /// @notice Mapping from staking token to set of pool IDs
    mapping(address => EnumerableSet.UintSet) private _poolsByStakingToken;

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
    function initialize(address _poolImplementation, address _admin, uint48 _initialDelay) external initializer {
        require(_poolImplementation != address(0), CSCanNotZeroAddress());
        require(_admin != address(0), CSCanNotZeroAddress());

        __AccessControlDefaultAdminRules_init(_initialDelay, _admin);
        __UUPSUpgradeable_init();

        poolImplementation = _poolImplementation;
        wcross = address(new WCROSS());
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
     * @return poolId ID of the newly created pool
     * @return poolAddress Address of the newly created pool
     */
    function createPool(address stakingToken)
        external
        onlyRole(MANAGER_ROLE)
        returns (uint poolId, address poolAddress)
    {
        require(stakingToken != address(0), CSCanNotZeroAddress());

        poolId = nextPoolId++;

        // Deploy pool as UUPS proxy
        // Pool will set CrossStaking as msg.sender and get owner from defaultAdmin()
        bytes memory initData = abi.encodeWithSelector(CrossStakingPool.initialize.selector, IERC20(stakingToken));

        ERC1967Proxy proxy = new ERC1967Proxy(poolImplementation, initData);
        poolAddress = address(proxy);

        // Store pool information
        pools[poolId] = PoolInfo({
            poolId: poolId,
            poolAddress: poolAddress,
            stakingToken: stakingToken,
            createdAt: block.timestamp,
            active: true
        });

        poolIds[poolAddress] = poolId;
        _allPoolIds.add(poolId);
        _poolsByStakingToken[stakingToken].add(poolId);

        emit PoolCreated(poolId, poolAddress, stakingToken);
    }

    /**
     * @notice Adds a reward token to a pool
     * @dev Can be called directly since CrossStaking has STAKING_ROOT_ROLE
     * @param poolId ID of the pool
     * @param rewardToken Address of the reward token to add
     */
    function addRewardToken(uint poolId, address rewardToken) external onlyRole(MANAGER_ROLE) {
        require(pools[poolId].poolAddress != address(0), CSPoolNotFound());

        CrossStakingPool pool = CrossStakingPool(pools[poolId].poolAddress);
        pool.addRewardToken(rewardToken);
    }

    /**
     * @notice Removes a reward token from a pool
     * @dev Can be called directly since CrossStaking has STAKING_ROOT_ROLE
     * @param poolId ID of the pool
     * @param rewardToken Address of the reward token to remove
     */
    function removeRewardToken(uint poolId, address rewardToken) external onlyRole(MANAGER_ROLE) {
        require(pools[poolId].poolAddress != address(0), CSPoolNotFound());

        CrossStakingPool pool = CrossStakingPool(pools[poolId].poolAddress);
        pool.removeRewardToken(rewardToken);
    }

    /**
     * @notice Activates or deactivates a pool
     * @dev Pauses/unpauses the actual pool contract
     * @param poolId ID of the pool
     * @param active New active status
     */
    function setPoolActive(uint poolId, bool active) external onlyRole(MANAGER_ROLE) {
        require(pools[poolId].poolAddress != address(0), CSPoolNotFound());

        pools[poolId].active = active;

        // Pause/unpause the actual pool contract
        CrossStakingPool pool = CrossStakingPool(pools[poolId].poolAddress);
        if (active) pool.unpause();
        else pool.pause();

        emit PoolStatusChanged(poolId, active);
    }

    /**
     * @notice Updates the pool implementation address
     * @dev Only affects newly created pools, existing pools remain unchanged
     * @param newImplementation Address of the new implementation contract
     */
    function setPoolImplementation(address newImplementation) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(newImplementation != address(0), CSCanNotZeroAddress());
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
        require(pools[poolId].poolAddress != address(0), CSPoolNotFound());
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
    function getPoolCountByStakingToken(address stakingToken) external view returns (uint) {
        return _poolsByStakingToken[stakingToken].length();
    }

    /**
     * @notice Retrieves all pool IDs for a specific staking token
     * @param stakingToken Address of the staking token
     * @return Array of pool IDs
     */
    function getPoolIdsByStakingToken(address stakingToken) external view returns (uint[] memory) {
        return _poolsByStakingToken[stakingToken].values();
    }

    /**
     * @notice Retrieves pool ID by staking token and index
     * @param stakingToken Address of the staking token
     * @param index Index in the staking token's pool list
     * @return poolId Pool ID at the specified index
     */
    function poolByStakingTokenAt(address stakingToken, uint index) external view returns (uint) {
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
    function getPoolAddress(uint poolId) external view returns (address) {
        require(pools[poolId].poolAddress != address(0), CSPoolNotFound());
        return pools[poolId].poolAddress;
    }

    /**
     * @notice Retrieves pool ID by pool address
     * @param poolAddress Address of the pool contract
     * @return Pool ID
     */
    function getPoolId(address poolAddress) external view returns (uint) {
        uint poolId = poolIds[poolAddress];
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
            if (pools[poolId].active) {
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
     *      Currently used: 8 slots
     *      Gap: 50 - 8 = 42 slots
     */
    uint[42] private __gap;
}
