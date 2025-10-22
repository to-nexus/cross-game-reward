// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./interfaces/IRewardPool.sol";
import "./interfaces/IStakingPool.sol";
import "./interfaces/IStakingProtocol.sol";
import "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

/**
 * @title StakingProtocol
 * @notice Factory 컨트랙트 - 프로젝트별 Pool 생성 및 관리
 * @dev Code 컨트랙트 패턴 사용으로 코드 사이즈 최소화
 */
contract StakingProtocol is IStakingProtocol, AccessControlDefaultAdminRules, ReentrancyGuardTransient {
    // ============ Role 정의 ============

    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");
    // ============ 에러 ============

    error StakingProtocolCanNotZeroAddress();
    error StakingProtocolEmptyProjectName();
    error StakingProtocolProjectNameExists();
    error StakingProtocolInvalidProjectID();
    error StakingProtocolInvalidSeasonBlocks();
    error StakingProtocolDeploymentFailed();
    error StakingProtocolNotAuthorized();

    // ============ 상태 변수 ============

    /// @notice $CROSS 토큰 주소 (내부용)
    IERC20 private immutable _crossToken;

    /// @notice 각 Code 컨트랙트 주소들
    address public immutable stakingPoolCodeContract;
    address public immutable rewardPoolCodeContract;

    /// @notice 프로젝트 목록
    mapping(uint => IStakingProtocol.ProjectInfo) public projects;

    /// @notice 총 프로젝트 수
    uint public projectCount;

    /// @notice 프로젝트 이름으로 ID 조회
    mapping(string => uint) public projectIDByName;

    /// @notice 주소로 프로젝트 ID 조회
    mapping(address => uint[]) public projectsByAddress;

    /// @notice 기본 시즌 길이 (블록 수)
    uint public defaultSeasonBlocks = 2592000; // 약 30일 (1초/블록 기준)

    // ============ Modifiers ============

    /**
     * @notice 프로젝트 admin 또는 프로토콜 admin만 호출 가능
     */
    modifier onlyProjectAdminOrProtocolAdmin(uint projectID) {
        require(projectID > 0 && projectID <= projectCount, StakingProtocolInvalidProjectID());
        require(
            msg.sender == projects[projectID].admin || hasRole(DEFAULT_ADMIN_ROLE, msg.sender),
            StakingProtocolNotAuthorized()
        );
        _;
    }

    // ============ 생성자 ============

    constructor(
        address crossTokenAddr,
        address _stakingPoolCodeContract,
        address _rewardPoolCodeContract,
        address _admin
    ) AccessControlDefaultAdminRules(3 days, _admin) {
        require(crossTokenAddr != address(0), StakingProtocolCanNotZeroAddress());
        require(_stakingPoolCodeContract != address(0), StakingProtocolCanNotZeroAddress());
        require(_rewardPoolCodeContract != address(0), StakingProtocolCanNotZeroAddress());

        _crossToken = IERC20(crossTokenAddr);
        stakingPoolCodeContract = _stakingPoolCodeContract;
        rewardPoolCodeContract = _rewardPoolCodeContract;
    }

    // ============ 프로젝트 생성 ============

    /**
     * @notice 새 프로젝트 풀 생성
     * @param projectName 프로젝트 이름
     * @param seasonBlocks 시즌 길이 (블록 수, 0이면 기본값 사용)
     * @param firstSeasonStartBlock 첫 시즌 시작 블록
     * @param poolEndBlock 풀 종료 블록 (0이면 무한)
     * @param projectAdmin 프로젝트 관리자 (0이면 msg.sender)
     */
    function createProject(
        string calldata projectName,
        uint seasonBlocks,
        uint firstSeasonStartBlock,
        uint poolEndBlock,
        address projectAdmin
    ) external nonReentrant returns (uint projectID, address stakingPool, address rewardPool) {
        require(bytes(projectName).length != 0, StakingProtocolEmptyProjectName());
        require(projectIDByName[projectName] == 0, StakingProtocolProjectNameExists());

        // projectAdmin이 0이면 msg.sender를 admin으로 설정
        if (projectAdmin == address(0)) projectAdmin = msg.sender;

        if (seasonBlocks == 0) seasonBlocks = defaultSeasonBlocks;
        require(seasonBlocks > 0, StakingProtocolInvalidSeasonBlocks());

        projectCount++;
        projectID = projectCount;

        // 각 풀 배포 (별도 함수로 분리하여 stack 깊이 줄임)
        stakingPool = _deployStakingPool(projectID, seasonBlocks, firstSeasonStartBlock, poolEndBlock);
        rewardPool = _deployRewardPool(stakingPool);

        // 연결 설정
        _setupPools(stakingPool, rewardPool);

        // 프로젝트 정보 저장
        _saveProjectInfo(projectID, projectName, stakingPool, rewardPool, projectAdmin);

        emit ProjectCreated(projectID, projectName, stakingPool, rewardPool, msg.sender, projectAdmin);
    }

    // ============ 배포 헬퍼 함수 ============

    function _deployStakingPool(uint projectID, uint seasonBlocks, uint firstSeasonStartBlock, uint poolEndBlock)
        internal
        returns (address pool)
    {
        bytes memory code = _getStakingPoolCode();
        bytes memory bytecode = bytes.concat(
            code,
            abi.encode(
                projectID, address(_crossToken), address(this), seasonBlocks, firstSeasonStartBlock, poolEndBlock
            )
        );

        assembly {
            pool := create(0, add(bytecode, 0x20), mload(bytecode))
        }
        require(pool != address(0), StakingProtocolDeploymentFailed());
    }

    function _deployRewardPool(address stakingPool) internal returns (address pool) {
        bytes memory code = _getRewardPoolCode();
        bytes memory bytecode = bytes.concat(code, abi.encode(stakingPool, address(this)));

        assembly {
            pool := create(0, add(bytecode, 0x20), mload(bytecode))
        }
        require(pool != address(0), StakingProtocolDeploymentFailed());
    }

    function _setupPools(address stakingPool, address rewardPool) internal {
        IStakingPool(stakingPool).setRewardPool(IRewardPool(rewardPool));
    }

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
        projectsByAddress[msg.sender].push(projectID);
    }

    // ============ Code 가져오기 함수 ============

    function _getStakingPoolCode() internal view returns (bytes memory) {
        (bool success, bytes memory code) = stakingPoolCodeContract.staticcall(abi.encodeWithSignature("code()"));
        require(success, "Failed to get StakingPool code");
        return abi.decode(code, (bytes));
    }

    function _getRewardPoolCode() internal view returns (bytes memory) {
        (bool success, bytes memory code) = rewardPoolCodeContract.staticcall(abi.encodeWithSignature("code()"));
        require(success, "Failed to get RewardPool code");
        return abi.decode(code, (bytes));
    }

    // ============ 보상 관리 함수 ============

    /**
     * @notice 프로젝트 시즌에 보상 예치
     * @param projectID 프로젝트 ID
     * @param season 시즌 번호
     * @param token 보상 토큰 주소
     * @param amount 예치할 수량
     * @dev 프로젝트 creator 또는 admin이 호출 가능
     */
    function fundProjectSeason(uint projectID, uint season, address token, uint amount) external nonReentrant {
        require(projectID > 0 && projectID <= projectCount, StakingProtocolInvalidProjectID());
        IStakingProtocol.ProjectInfo storage project = projects[projectID];
        require(msg.sender == project.creator || msg.sender == project.admin, "Only project creator or admin");
        require(project.isActive, "Project not active");

        // RewardPool에 보상 예치
        IRewardPool(project.rewardPool).fundSeason(season, token, amount);
    }

    // ============ 관리 함수 ============

    /**
     * @notice 프로젝트 admin 변경 (현재 admin만 가능)
     */
    function setProjectAdmin(uint projectID, address newAdmin) external {
        require(projectID > 0 && projectID <= projectCount, StakingProtocolInvalidProjectID());
        require(newAdmin != address(0), StakingProtocolCanNotZeroAddress());
        require(msg.sender == projects[projectID].admin, StakingProtocolNotAuthorized());

        address oldAdmin = projects[projectID].admin;
        projects[projectID].admin = newAdmin;
        emit ProjectAdminUpdated(projectID, oldAdmin, newAdmin);
    }

    /**
     * @notice 프로젝트 상태 변경 (프로토콜 admin만 가능)
     */
    function setProjectStatus(uint projectID, bool isActive) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(projectID > 0 && projectID <= projectCount, StakingProtocolInvalidProjectID());
        projects[projectID].isActive = isActive;
        emit ProjectStatusUpdated(projectID, isActive);
    }

    /**
     * @notice Router 승인 (프로젝트 admin 또는 프로토콜 admin)
     */
    function setApprovedRouter(uint projectID, address router, bool approved)
        external
        onlyProjectAdminOrProtocolAdmin(projectID)
    {
        IStakingPool(projects[projectID].stakingPool).setApprovedRouter(router, approved);
    }

    /**
     * @notice 기본 시즌 길이 설정
     */
    function setDefaultSeasonBlocks(uint blocks) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(blocks > 0, StakingProtocolInvalidSeasonBlocks());
        uint oldValue = defaultSeasonBlocks;
        defaultSeasonBlocks = blocks;
        emit DefaultSeasonBlocksUpdated(oldValue, blocks);
    }

    /**
     * @notice 프로젝트별 포인트 계산 시간 단위 설정 (프로젝트 admin 또는 프로토콜 admin)
     */
    function setPoolPointsTimeUnit(uint projectID, uint timeUnit) external onlyProjectAdminOrProtocolAdmin(projectID) {
        IStakingPool(projects[projectID].stakingPool).setPointsTimeUnit(timeUnit);
    }

    /**
     * @notice 프로젝트별 블록 시간 설정 (프로젝트 admin 또는 프로토콜 admin)
     */
    function setPoolBlockTime(uint projectID, uint blockTime) external onlyProjectAdminOrProtocolAdmin(projectID) {
        IStakingPool(projects[projectID].stakingPool).setBlockTime(blockTime);
    }

    /**
     * @notice 프로젝트별 다음 시즌 시작 블록 설정 (프로젝트 admin 또는 프로토콜 admin)
     */
    function setPoolNextSeasonStart(uint projectID, uint startBlock)
        external
        onlyProjectAdminOrProtocolAdmin(projectID)
    {
        IStakingPool(projects[projectID].stakingPool).setNextSeasonStart(startBlock);
    }

    /**
     * @notice 프로젝트별 풀 종료 블록 설정 (프로젝트 admin 또는 프로토콜 admin)
     */
    function setPoolEndBlock(uint projectID, uint endBlock) external onlyProjectAdminOrProtocolAdmin(projectID) {
        IStakingPool(projects[projectID].stakingPool).setPoolEndBlock(endBlock);
    }

    /**
     * @notice 프로젝트별 RewardPool 설정 (프로토콜 admin만 가능 - 보안 중요)
     */
    function setPoolRewardPool(uint projectID, IRewardPool rewardPool) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(projectID > 0 && projectID <= projectCount, StakingProtocolInvalidProjectID());
        IStakingPool(projects[projectID].stakingPool).setRewardPool(rewardPool);
    }

    /**
     * @notice 프로젝트별 StakingAddon 설정 (프로젝트 admin 또는 프로토콜 admin)
     */
    function setPoolStakingAddon(uint projectID, IStakingAddon addon)
        external
        onlyProjectAdminOrProtocolAdmin(projectID)
    {
        IStakingPool(projects[projectID].stakingPool).setStakingAddon(addon);
    }

    /**
     * @notice 프로젝트별 Addon 승인 설정 (프로젝트 admin 또는 프로토콜 admin)
     */
    function setPoolAddonApproved(uint projectID, IStakingAddon addon, bool approved)
        external
        onlyProjectAdminOrProtocolAdmin(projectID)
    {
        IStakingPool(projects[projectID].stakingPool).setAddonApproved(addon, approved);
    }

    // ============ 조회 함수 ============

    /**
     * @notice $CROSS 토큰 주소 반환 (IStakingProtocol 인터페이스 구현)
     */
    function crossToken() external view returns (address) {
        return address(_crossToken);
    }

    function getProject(uint projectID) external view returns (IStakingProtocol.ProjectInfo memory) {
        require(projectID > 0 && projectID <= projectCount, StakingProtocolInvalidProjectID());
        return projects[projectID];
    }

    function getProjectsByCreator(address creator) external view returns (uint[] memory) {
        return projectsByAddress[creator];
    }

    /**
     * @notice 프로젝트 목록 페이지네이션 조회 (AUDIT 권장)
     * @param offset 시작 ID (1부터 시작)
     * @param limit 최대 개수
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
}
