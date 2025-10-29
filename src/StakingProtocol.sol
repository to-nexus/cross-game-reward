// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {RewardPool} from "./RewardPool.sol";
import {StakingPool} from "./StakingPool.sol";
import {IRewardPool} from "./interfaces/IRewardPool.sol";
import {IStakingPool} from "./interfaces/IStakingPool.sol";
import {IRewardPoolCode, IStakingPoolCode, IStakingProtocol} from "./interfaces/IStakingProtocol.sol";
import {AccessControlDefaultAdminRules} from
    "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuardTransient} from "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";

/**
 * @title StakingProtocol
 * @notice Central factory contract for creating and managing project-specific staking pools
 * @dev Implements the Code Contract pattern with CREATE2 for gas-efficient deployments
 *
 *      Architecture:
 *      - Acts as factory for StakingPool and RewardPool instances
 *      - Uses Code Contracts to store creation bytecode separately
 *      - Deploys pools via CREATE2 for deterministic addresses
 *      - Manages global settings and project metadata
 *
 *      Code Contract Pattern:
 *      1. StakingPoolCode and RewardPoolCode store creation bytecode
 *      2. Factory retrieves bytecode via code() function
 *      3. Combines with constructor args
 *      4. Deploys using CREATE2 with deterministic salt
 *
 *      Benefits:
 *      - Reduces factory contract size (avoids 24KB limit)
 *      - Predictable pool addresses (useful for frontends)
 *      - Cross-chain deployment with same addresses
 *      - Upgradeable pool logic (by deploying new Code contracts)
 *
 *      CREATE2 Salt Structure:
 *      - StakingPool: keccak256(projectName, "StakingPool")
 *      - RewardPool: keccak256(projectName, "RewardPool")
 *
 *      Access Control:
 *      - DEFAULT_ADMIN_ROLE: Protocol-level admin (3-day timelock)
 *      - Project creators: Can fund their own projects
 *      - Project admins: Can modify their project's pool settings
 */
contract StakingProtocol is IStakingProtocol, AccessControlDefaultAdminRules, ReentrancyGuardTransient {
    using SafeERC20 for IERC20;

    // ============ Constants ============

    /// @notice Role for pool managers
    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");

    // ============ Events ============

    /// @notice Emitted when global router approval changes
    event GlobalRouterApprovalUpdated(address indexed router, bool approved);

    // ============ Errors ============

    /// @notice Thrown when zero address provided
    error StakingProtocolCanNotZeroAddress();

    /// @notice Thrown when project name is empty
    error StakingProtocolEmptyProjectName();

    /// @notice Thrown when project name already exists
    error StakingProtocolProjectNameExists();

    /// @notice Thrown when project ID is invalid
    error StakingProtocolInvalidProjectID();

    /// @notice Thrown when season blocks parameter is invalid
    error StakingProtocolInvalidSeasonBlocks();

    /// @notice Thrown when pool deployment fails
    error StakingProtocolDeploymentFailed();

    /// @notice Thrown when caller is not authorized
    error StakingProtocolNotAuthorized();

    /// @notice Thrown when amount is zero
    error StakingProtocolAmountMustBeGreaterThanZero();

    /// @notice Thrown when reward pool not found
    error StakingProtocolRewardPoolNotFound();

    /// @notice Thrown when code retrieval from Code contract fails
    error StakingProtocolCodeRetrievalFailed();

    // ============ Immutable State ============

    /// @notice CROSS token address (WCROSS)
    address public immutable crossToken;

    /// @notice Code contract for StakingPool deployments
    IStakingPoolCode public immutable stakingPoolCodeContract;

    /// @notice Code contract for RewardPool deployments
    IRewardPoolCode public immutable rewardPoolCodeContract;

    // ============ State Variables ============

    /// @notice Mapping of project ID to project information
    mapping(uint => IStakingProtocol.ProjectInfo) public projects;

    /// @notice Total number of projects created
    uint public projectCount;

    /// @notice Mapping of project name to project ID (ensures uniqueness)
    mapping(string => uint) public projectIDByName;

    /// @notice Mapping of admin address to their project IDs
    mapping(address => uint[]) public projectsByAdmin;

    /// @notice Default season length in blocks (≈30 days at 1 sec/block)
    uint public defaultSeasonBlocks = 2592000;

    /// @notice Mapping of globally approved routers (can interact with any project)
    mapping(address => bool) public globalApprovedRouters;

    // ============ Modifiers ============

    /**
     * @notice Restricts access to project admin or protocol admin
     * @param projectID Project ID to check
     * @dev Allows either:
     *      - The project's admin address
     *      - An address with DEFAULT_ADMIN_ROLE
     */
    modifier onlyProjectAdminOrProtocolAdmin(uint projectID) {
        require(projectID > 0 && projectID <= projectCount, StakingProtocolInvalidProjectID());
        require(
            msg.sender == projects[projectID].admin || hasRole(DEFAULT_ADMIN_ROLE, msg.sender),
            StakingProtocolNotAuthorized()
        );
        _;
    }

    // ============ Constructor ============

    /**
     * @notice Initializes the StakingProtocol factory
     * @param crossTokenAddr Address of CROSS token (WCROSS)
     * @param _stakingPoolCodeContract Address of StakingPoolCode contract
     * @param _rewardPoolCodeContract Address of RewardPoolCode contract
     * @param _admin Initial protocol admin address
     * @dev - Validates all addresses are non-zero
     *      - Validates Code contracts can return bytecode
     *      - Sets up 3-day timelock for admin role
     *      - Stores immutable references
     */
    constructor(
        address crossTokenAddr,
        address _stakingPoolCodeContract,
        address _rewardPoolCodeContract,
        address _admin
    ) AccessControlDefaultAdminRules(3 days, _admin) {
        require(crossTokenAddr != address(0), StakingProtocolCanNotZeroAddress());
        require(_stakingPoolCodeContract != address(0), StakingProtocolCanNotZeroAddress());
        require(_rewardPoolCodeContract != address(0), StakingProtocolCanNotZeroAddress());
        require(_admin != address(0), StakingProtocolCanNotZeroAddress());

        (bool success1, bytes memory data1) = _stakingPoolCodeContract.staticcall(abi.encodeWithSignature("code()"));
        require(success1 && data1.length > 0, StakingProtocolCodeRetrievalFailed());

        (bool success2, bytes memory data2) = _rewardPoolCodeContract.staticcall(abi.encodeWithSignature("code()"));
        require(success2 && data2.length > 0, StakingProtocolCodeRetrievalFailed());

        crossToken = crossTokenAddr;
        stakingPoolCodeContract = IStakingPoolCode(_stakingPoolCodeContract);
        rewardPoolCodeContract = IRewardPoolCode(_rewardPoolCodeContract);
    }

    // ============ Project Creation ============

    /**
     * @notice Creates a new staking project with dedicated pools
     * @param projectName Unique project name
     * @param seasonDuration Number of blocks per season (0 = use default)
     * @param firstSeasonStartTime Block when season 1 starts
     * @param poolEndTime Block when pool ends (0 = infinite)
     * @param projectAdmin Admin address for the project
     * @param preDepositStartTime Block when pre-deposit starts (0 = disabled)
     * @return projectID Newly created project ID
     * @return stakingPool Address of deployed StakingPool
     * @return rewardPool Address of deployed RewardPool
     * @dev Process:
     *      1. Validates project name is unique and non-empty
     *      2. Uses default season blocks if 0 provided
     *      3. Increments project counter
     *      4. Deploys StakingPool via CREATE2 (deterministic address)
     *      5. Deploys RewardPool via CREATE2 (deterministic address)
     *      6. Connects pools to each other
     *      7. Saves project metadata
     *      8. Emits ProjectCreated event
     *
     *      CREATE2 ensures:
     *      - Address predictable before deployment
     *      - Same project name = same addresses across chains
     *      - Useful for frontend pre-display and cross-chain consistency
     */
    function createProject(
        string calldata projectName,
        uint seasonDuration,
        uint firstSeasonStartTime,
        uint poolEndTime,
        address projectAdmin,
        uint preDepositStartTime
    ) external nonReentrant returns (uint projectID, address stakingPool, address rewardPool) {
        require(bytes(projectName).length != 0, StakingProtocolEmptyProjectName());
        require(projectIDByName[projectName] == 0, StakingProtocolProjectNameExists());
        require(projectAdmin != address(0), StakingProtocolCanNotZeroAddress());

        if (seasonDuration == 0) seasonDuration = defaultSeasonBlocks;
        require(seasonDuration > 0, StakingProtocolInvalidSeasonBlocks());

        projectCount++;
        projectID = projectCount;
        bytes32 salt = keccak256(abi.encode(projectName));

        // Deploy both pools
        stakingPool =
            _deployStakingPool(projectID, salt, seasonDuration, firstSeasonStartTime, poolEndTime, preDepositStartTime);
        rewardPool = _deployRewardPool(projectID, salt, stakingPool);

        // Connect pools
        _setupPools(stakingPool, rewardPool);

        // Save project info
        _saveProjectInfo(projectID, projectName, stakingPool, rewardPool, projectAdmin);

        emit ProjectCreated(projectID, projectName, stakingPool, rewardPool, msg.sender, projectAdmin);
    }

    // ============ Deployment Helpers ============

    /**
     * @notice Deploys StakingPool using CREATE2 for deterministic address
     * @param projectID Project ID
     * @param salt Salt for deterministic address
     * @param seasonDuration Blocks per season
     * @param firstSeasonStartTime First season start block
     * @param poolEndTime Pool end block
     * @param preDepositStartTime Pre-deposit start block
     * @return pool Deployed StakingPool address
     * @dev Salt: keccak256(abi.encode(projectName))
     *      This makes addresses predictable based on project name
     */
    function _deployStakingPool(
        uint projectID,
        bytes32 salt,
        uint seasonDuration,
        uint firstSeasonStartTime,
        uint poolEndTime,
        uint preDepositStartTime
    ) internal returns (address pool) {
        bytes memory code = _getStakingPoolCode();
        bytes memory bytecode = bytes.concat(
            code,
            abi.encode(
                projectID,
                crossToken,
                address(this),
                seasonDuration,
                firstSeasonStartTime,
                poolEndTime,
                preDepositStartTime
            )
        );
        pool = Create2.deploy(0, salt, bytecode);
    }

    /**
     * @notice Deploys RewardPool using CREATE2 for deterministic address
     * @param salt Salt for deterministic address
     * @param stakingPool Address of already-deployed StakingPool
     * @return pool Deployed RewardPool address
     * @dev Salt: keccak256(abi.encode(projectName))
     *      This makes addresses predictable based on project name
     */
    function _deployRewardPool(uint, /* projectID */ bytes32 salt, address stakingPool)
        internal
        returns (address pool)
    {
        bytes memory code = _getRewardPoolCode();
        bytes memory bytecode = bytes.concat(code, abi.encode(stakingPool, address(this)));
        pool = Create2.deploy(0, salt, bytecode);
    }

    /**
     * @notice Connects StakingPool and RewardPool
     * @param stakingPool StakingPool address
     * @param rewardPool RewardPool address
     * @dev Calls StakingPool.setRewardPool() to establish connection
     */
    function _setupPools(address stakingPool, address rewardPool) internal {
        IStakingPool(stakingPool).setRewardPool(IRewardPool(rewardPool));
    }

    /**
     * @notice Saves project information to storage
     * @param projectID Project ID
     * @param projectName Project name
     * @param stakingPool StakingPool address
     * @param rewardPool RewardPool address
     * @param projectAdmin Project admin address
     * @dev Updates three mappings:
     *      - projects[projectID] = full project info
     *      - projectIDByName[name] = projectID
     *      - projectsByAdmin[admin] += projectID
     */
    function _saveProjectInfo(
        uint projectID,
        string calldata projectName,
        address stakingPool,
        address rewardPool,
        address projectAdmin
    ) internal {
        projects[projectID] = IStakingProtocol.ProjectInfo({
            stakingPool: stakingPool,
            rewardPool: rewardPool,
            name: projectName,
            isActive: true,
            createdAt: block.timestamp,
            creator: msg.sender,
            admin: projectAdmin
        });

        projectIDByName[projectName] = projectID;
        projectsByAdmin[projectAdmin].push(projectID);
    }

    // ============ Code Retrieval Functions ============

    /**
     * @notice Retrieves StakingPool creation bytecode from Code contract
     * @return Creation bytecode for StakingPool
     */
    function _getStakingPoolCode() internal view returns (bytes memory) {
        return stakingPoolCodeContract.code();
    }

    /**
     * @notice Retrieves RewardPool creation bytecode from Code contract
     * @return Creation bytecode for RewardPool
     */
    function _getRewardPoolCode() internal view returns (bytes memory) {
        return rewardPoolCodeContract.code();
    }

    // ============ Admin Functions ============

    /**
     * @notice Changes project admin (current admin only)
     * @param projectID Project ID
     * @param newAdmin New admin address
     * @dev Only the current project admin can transfer admin rights
     */
    function setProjectAdmin(uint projectID, address newAdmin) external {
        _checkProjectID(projectID);
        require(newAdmin != address(0), StakingProtocolCanNotZeroAddress());
        require(msg.sender == projects[projectID].admin, StakingProtocolNotAuthorized());

        address oldAdmin = projects[projectID].admin;
        projects[projectID].admin = newAdmin;
        emit ProjectAdminUpdated(projectID, oldAdmin, newAdmin);
    }

    /**
     * @notice Sets project active status (protocol admin only)
     * @param projectID Project ID
     * @param isActive New active status
     * @dev Allows protocol admin to pause/unpause projects
     */
    function setProjectStatus(uint projectID, bool isActive) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _checkProjectID(projectID);
        projects[projectID].isActive = isActive;
        emit ProjectStatusUpdated(projectID, isActive);
    }

    /**
     * @notice Sets globally approved router (protocol admin only)
     * @param router Router address
     * @param approved Approval status
     * @dev Global routers can interact with any project without per-project approval
     */
    function setGlobalApprovedRouter(address router, bool approved) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(router != address(0), StakingProtocolCanNotZeroAddress());
        globalApprovedRouters[router] = approved;
        emit GlobalRouterApprovalUpdated(router, approved);
    }

    /**
     * @notice Approves router for a specific project
     * @param projectID Project ID
     * @param router Router address
     * @param approved Approval status
     * @dev Can be called by project admin or protocol admin
     */
    function setApprovedRouter(uint projectID, address router, bool approved)
        external
        onlyProjectAdminOrProtocolAdmin(projectID)
    {
        IStakingPool(projects[projectID].stakingPool).setApprovedRouter(router, approved);
    }

    /**
     * @notice Sets default season blocks for new projects
     * @param blocks Number of blocks per season
     * @dev Only affects newly created projects
     */
    function setDefaultSeasonBlocks(uint blocks) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(blocks > 0, StakingProtocolInvalidSeasonBlocks());
        emit DefaultSeasonBlocksUpdated(defaultSeasonBlocks, blocks);
        defaultSeasonBlocks = blocks;
    }

    /**
     * @notice Sets next season start block for a project's pool
     * @param projectID Project ID
     * @param startTime Start block number
     */
    function setPoolNextSeasonStart(uint projectID, uint startTime)
        external
        onlyProjectAdminOrProtocolAdmin(projectID)
    {
        IStakingPool(projects[projectID].stakingPool).setNextSeasonStart(startTime);
    }

    /**
     * @notice Sets pool end block for a project
     * @param projectID Project ID
     * @param endTime End block number (0 = infinite)
     */
    function setPoolEndBlock(uint projectID, uint endTime) external onlyProjectAdminOrProtocolAdmin(projectID) {
        IStakingPool(projects[projectID].stakingPool).setPoolEndBlock(endTime);
    }

    /**
     * @notice Changes reward pool for a project (emergency only)
     * @param projectID Project ID
     * @param rewardPool New RewardPool address
     * @dev Protocol admin only, should rarely be used
     */
    function setPoolRewardPool(uint projectID, IRewardPool rewardPool) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _checkProjectID(projectID);
        IStakingPool(projects[projectID].stakingPool).setRewardPool(rewardPool);
    }

    /**
     * @notice Manually rolls over seasons for a project
     * @param projectID Project ID
     * @param maxRollovers Maximum seasons to rollover
     * @return rolloversPerformed Number of seasons actually rolled over
     * @dev Used when >50 seasons are pending
     */
    function manualRolloverSeasons(uint projectID, uint maxRollovers)
        external
        onlyProjectAdminOrProtocolAdmin(projectID)
        returns (uint rolloversPerformed)
    {
        return IStakingPool(projects[projectID].stakingPool).manualRolloverSeasons(maxRollovers);
    }

    /**
     * @notice Emergency function to recover tokens from reward pool
     * @param projectID Project ID
     * @param token Token address
     * @param to Recipient address
     * @param amount Amount to recover
     * @dev Protocol admin only
     */
    function sweepRewardPool(uint projectID, address token, address to, uint amount)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _checkProjectID(projectID);
        require(to != address(0), StakingProtocolCanNotZeroAddress());
        require(amount > 0, StakingProtocolAmountMustBeGreaterThanZero());

        address rewardPoolAddress = projects[projectID].rewardPool;
        require(rewardPoolAddress != address(0), StakingProtocolRewardPoolNotFound());

        IRewardPool(rewardPoolAddress).sweep(token, to, amount);
    }

    /**
     * @notice Returns number of pending season rollovers for a project
     * @param projectID Project ID
     * @return pendingSeasons Number of pending rollovers
     */
    function getPendingSeasonRollovers(uint projectID) external view returns (uint pendingSeasons) {
        _checkProjectID(projectID);
        return IStakingPool(projects[projectID].stakingPool).getPendingSeasonRollovers();
    }

    // ============ View Functions ============

    /**
     * @notice Checks if a router is globally approved
     * @param router Router address
     * @return True if globally approved
     */
    function isGlobalApprovedRouter(address router) external view returns (bool) {
        return globalApprovedRouters[router];
    }

    /**
     * @notice Returns project information
     * @param projectID Project ID
     * @return Project information struct
     */
    function getProject(uint projectID) external view returns (IStakingProtocol.ProjectInfo memory) {
        _checkProjectID(projectID);
        return projects[projectID];
    }

    /**
     * @notice Returns all projects managed by an admin
     * @param admin Admin address
     * @return Array of project IDs
     */
    function getProjectsByAdmin(address admin) external view returns (uint[] memory) {
        return projectsByAdmin[admin];
    }

    /**
     * @notice Computes StakingPool address before deployment (CREATE2 prediction)
     * @param projectName Project name
     * @param projectID Project ID
     * @param seasonDuration Season length in blocks
     * @param firstSeasonStartTime First season start block
     * @param poolEndTime Pool end block
     * @return Predicted StakingPool address
     * @dev Uses CREATE2 formula with deterministic salt
     *      Useful for frontends to display address before creation
     */
    function computeStakingPoolAddress(
        string calldata projectName,
        uint projectID,
        uint seasonDuration,
        uint firstSeasonStartTime,
        uint poolEndTime
    ) external view returns (address) {
        bytes memory code = _getStakingPoolCode();
        bytes memory bytecode = bytes.concat(
            code, abi.encode(projectID, crossToken, address(this), seasonDuration, firstSeasonStartTime, poolEndTime)
        );

        bytes32 salt = keccak256(abi.encode(projectName, "StakingPool"));
        bytes32 hash = keccak256(abi.encodePacked(bytes1(0xff), address(this), salt, keccak256(bytecode)));

        return address(uint160(uint(hash)));
    }

    /**
     * @notice Computes RewardPool address before deployment (CREATE2 prediction)
     * @param projectName Project name
     * @param stakingPool StakingPool address (must compute StakingPool address first)
     * @return Predicted RewardPool address
     */
    function computeRewardPoolAddress(string calldata projectName, uint, /* projectID */ address stakingPool)
        external
        view
        returns (address)
    {
        bytes memory code = _getRewardPoolCode();
        bytes memory bytecode = bytes.concat(code, abi.encode(stakingPool, address(this)));

        bytes32 salt = keccak256(abi.encode(projectName, "RewardPool"));
        bytes32 hash = keccak256(abi.encodePacked(bytes1(0xff), address(this), salt, keccak256(bytecode)));

        return address(uint160(uint(hash)));
    }

    /**
     * @notice Returns paginated list of projects
     * @param offset Starting project ID (1-indexed)
     * @param limit Maximum number of projects to return
     * @return projectList Array of project information
     * @return total Total number of projects
     * @dev Supports pagination for frontend display
     *      offset is 1-indexed (project IDs start at 1)
     */
    function getProjects(uint offset, uint limit)
        external
        view
        returns (IStakingProtocol.ProjectInfo[] memory projectList, uint total)
    {
        total = projectCount;
        if (offset < 1 || offset > total) return (new IStakingProtocol.ProjectInfo[](0), total);

        uint end = offset + limit - 1;
        if (end > total) end = total;

        uint count = end - offset + 1;
        projectList = new IStakingProtocol.ProjectInfo[](count);

        for (uint i = 0; i < count; i++) {
            projectList[i] = projects[offset + i];
        }
    }

    /**
     * @notice Validates project ID is within valid range
     * @param projectID Project ID to check
     * @dev Reverts if project ID is 0 or greater than projectCount
     */
    function _checkProjectID(uint projectID) internal view {
        require(projectID > 0 && projectID <= projectCount, StakingProtocolInvalidProjectID());
    }
}
