// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IRewardPool, IRewardPoolCode} from "./IRewardPool.sol";
import {IStakingPoolCode} from "./IStakingPool.sol";

/**
 * @title IStakingProtocol
 * @notice Interface for the StakingProtocol factory contract
 * @dev StakingProtocol serves as the central factory for creating and managing project-specific staking pools
 *      It uses the Code Contract pattern with CREATE2 for deterministic deployment addresses
 */
interface IStakingProtocol {
    // ============ Structs ============

    /**
     * @notice Information about a staking project
     * @param stakingPool Address of the project's StakingPool contract
     * @param rewardPool Address of the project's RewardPool contract
     * @param name Unique project name
     * @param isActive Whether the project is currently active
     * @param createdAt Timestamp when the project was created
     * @param creator Address that created the project
     * @param admin Project administrator who can modify pool settings
     */
    struct ProjectInfo {
        address stakingPool;
        address rewardPool;
        string name;
        bool isActive;
        uint createdAt;
        address creator;
        address admin;
    }

    // ============ Events ============

    /// @notice Emitted when a new project is created
    event ProjectCreated(
        uint indexed projectID, string name, address stakingPool, address rewardPool, address creator, address admin
    );

    /// @notice Emitted when a project's admin address is changed
    event ProjectAdminUpdated(uint indexed projectID, address indexed oldAdmin, address indexed newAdmin);

    /// @notice Emitted when a project's active status is changed
    event ProjectStatusUpdated(uint indexed projectID, bool isActive);

    /// @notice Emitted when the default season block count is updated
    event DefaultSeasonBlocksUpdated(uint oldValue, uint newValue);

    // ============ Project Creation ============

    /**
     * @notice Creates a new staking project with dedicated StakingPool and RewardPool
     * @param projectName Unique name for the project
     * @param seasonBlocks Number of blocks per season (0 = use default)
     * @param firstSeasonStartBlock Block number when first season starts
     * @param poolEndBlock Block number when pool ends (0 = infinite)
     * @param projectAdmin Address of the project administrator
     * @param preDepositStartBlock Block number when pre-deposit starts (0 = disabled)
     * @return projectID The newly created project's ID
     * @return stakingPool Address of the deployed StakingPool
     * @return rewardPool Address of the deployed RewardPool
     * @dev Uses CREATE2 for deterministic addresses based on project name
     */
    function createProject(
        string calldata projectName,
        uint seasonBlocks,
        uint firstSeasonStartBlock,
        uint poolEndBlock,
        address projectAdmin,
        uint preDepositStartBlock
    ) external returns (uint projectID, address stakingPool, address rewardPool);

    // ============ Admin Functions ============

    /// @notice Changes the administrator of a project (only current admin)
    function setProjectAdmin(uint projectID, address newAdmin) external;

    /// @notice Sets the active status of a project (only protocol admin)
    function setProjectStatus(uint projectID, bool isActive) external;

    /// @notice Approves/revokes a router globally (only protocol admin)
    function setGlobalApprovedRouter(address router, bool approved) external;

    /// @notice Approves/revokes a router for a specific project
    function setApprovedRouter(uint projectID, address router, bool approved) external;

    /// @notice Sets the default season block count for new projects
    function setDefaultSeasonBlocks(uint blocks) external;

    /// @notice Sets the points calculation time unit for a project's pool
    function setPoolPointsTimeUnit(uint projectID, uint timeUnit) external;

    /// @notice Sets the block time for a project's pool
    function setPoolBlockTime(uint projectID, uint blockTime) external;

    /// @notice Sets the next season start block for a project's pool
    function setPoolNextSeasonStart(uint projectID, uint startBlock) external;

    /// @notice Sets the end block for a project's pool
    function setPoolEndBlock(uint projectID, uint endBlock) external;

    /// @notice Changes the reward pool for a project (emergency use only)
    function setPoolRewardPool(uint projectID, IRewardPool rewardPool) external;

    /// @notice Sweeps tokens from a project's reward pool (emergency recovery)
    function sweepRewardPool(uint projectID, address token, address to, uint amount) external;

    // ============ View Functions ============

    /// @notice Returns the address of the CROSS token
    function crossToken() external view returns (address);

    /// @notice Returns the total number of projects created
    function projectCount() external view returns (uint);

    /// @notice Checks if a router is globally approved
    function isGlobalApprovedRouter(address router) external view returns (bool);

    /// @notice Returns project information by ID
    function projects(uint projectID)
        external
        view
        returns (
            address stakingPool,
            address rewardPool,
            string memory name,
            bool isActive,
            uint createdAt,
            address creator,
            address admin
        );

    /// @notice Returns project ID by name
    function projectIDByName(string calldata name) external view returns (uint);

    /// @notice Returns full project information
    function getProject(uint projectID) external view returns (ProjectInfo memory);

    /// @notice Returns paginated list of projects
    function getProjects(uint offset, uint limit)
        external
        view
        returns (ProjectInfo[] memory projectList, uint total);
}
