// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./IRewardPool.sol";
import "./IStakingAddon.sol";

/**
 * @title IStakingProtocol
 * @notice StakingProtocol 컨트랙트의 인터페이스
 */
interface IStakingProtocol {
    // ============ 구조체 ============

    struct ProjectInfo {
        address stakingPool;
        address rewardPool;
        string name;
        bool isActive;
        uint createdAt;
        address creator;
        address admin; // 프로젝트 관리자 (pool 설정 변경 가능)
    }

    // ============ 이벤트 ============

    event ProjectCreated(
        uint indexed projectID, string name, address stakingPool, address rewardPool, address creator, address admin
    );
    event ProjectAdminUpdated(uint indexed projectID, address indexed oldAdmin, address indexed newAdmin);
    event ProjectStatusUpdated(uint indexed projectID, bool isActive);
    event DefaultSeasonBlocksUpdated(uint oldValue, uint newValue);

    // ============ 프로젝트 생성 ============

    function createProject(
        string calldata projectName,
        uint seasonBlocks,
        uint firstSeasonStartBlock,
        uint poolEndBlock,
        address projectAdmin
    ) external returns (uint projectID, address stakingPool, address rewardPool);

    // ============ 보상 관리 ============

    function fundProjectSeason(uint projectID, uint season, address token, uint amount) external;

    // ============ 관리 함수 ============

    function setProjectAdmin(uint projectID, address newAdmin) external;
    function setProjectStatus(uint projectID, bool isActive) external;
    function setApprovedRouter(uint projectID, address router, bool approved) external;
    function setDefaultSeasonBlocks(uint blocks) external;
    function setPoolPointsTimeUnit(uint projectID, uint timeUnit) external;
    function setPoolBlockTime(uint projectID, uint blockTime) external;
    function setPoolNextSeasonStart(uint projectID, uint startBlock) external;
    function setPoolEndBlock(uint projectID, uint endBlock) external;
    function setPoolRewardPool(uint projectID, IRewardPool rewardPool) external;
    function setPoolStakingAddon(uint projectID, IStakingAddon addon) external;
    function setPoolAddonApproved(uint projectID, IStakingAddon addon, bool approved) external;

    // ============ 조회 함수 ============

    function crossToken() external view returns (address); // IERC20를 address로 반환 (호환성)
    function projectCount() external view returns (uint);
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
    function projectIDByName(string calldata name) external view returns (uint);
    function getProject(uint projectID) external view returns (ProjectInfo memory);
    function getProjectsByCreator(address creator) external view returns (uint[] memory);
    function getProjects(uint offset, uint limit)
        external
        view
        returns (ProjectInfo[] memory projectList, uint total);
}
