// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./interfaces/IRewardPool.sol";

import "./interfaces/IStakingPool.sol";
import "./interfaces/IStakingProtocol.sol";
import "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

import "./RewardPool.sol";
import "./StakingPool.sol";
/**
 * @title StakingProtocol
 * @notice Factory 컨트랙트 - 프로젝트별 Pool 생성 및 관리
 * @dev Code 컨트랙트 패턴 사용으로 코드 사이즈 최소화
 */

contract StakingProtocol is IStakingProtocol, AccessControlDefaultAdminRules, ReentrancyGuardTransient {
    using SafeERC20 for IERC20;

    // ============ Role 정의 ============

    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");

    // ============ 이벤트 ============

    event GlobalRouterApprovalUpdated(address indexed router, bool approved);

    // ============ 에러 ============

    error StakingProtocolCanNotZeroAddress();
    error StakingProtocolEmptyProjectName();
    error StakingProtocolProjectNameExists();
    error StakingProtocolInvalidProjectID();
    error StakingProtocolInvalidSeasonBlocks();
    error StakingProtocolDeploymentFailed();
    error StakingProtocolNotAuthorized();
    error StakingProtocolAmountMustBeGreaterThanZero();
    error StakingProtocolRewardPoolNotFound();
    error StakingProtocolCodeRetrievalFailed();

    // ============ 상태 변수 ============

    /// @notice $CROSS 토큰 주소 (내부용)
    address public immutable crossToken;

    /// @notice 각 Code 컨트랙트 주소들
    IStakingPoolCode public immutable stakingPoolCodeContract;
    IRewardPoolCode public immutable rewardPoolCodeContract;

    /// @notice 프로젝트 목록
    mapping(uint => IStakingProtocol.ProjectInfo) public projects;

    /// @notice 총 프로젝트 수
    uint public projectCount;

    /// @notice 프로젝트 이름으로 ID 조회
    mapping(string => uint) public projectIDByName;

    /// @notice 관리자 주소로 프로젝트 ID 조회
    mapping(address => uint[]) public projectsByAdmin;

    /// @notice 기본 시즌 길이 (블록 수)
    uint public defaultSeasonBlocks = 2592000; // 약 30일 (1초/블록 기준)

    /// @notice 글로벌 승인 Router 목록
    mapping(address => bool) public globalApprovedRouters;

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
        require(_admin != address(0), StakingProtocolCanNotZeroAddress());

        // Code 컨트랙트가 실제로 code() 함수를 가지고 있는지 검증
        (bool success1, bytes memory data1) = _stakingPoolCodeContract.staticcall(abi.encodeWithSignature("code()"));
        require(success1 && data1.length > 0, StakingProtocolCodeRetrievalFailed());

        (bool success2, bytes memory data2) = _rewardPoolCodeContract.staticcall(abi.encodeWithSignature("code()"));
        require(success2 && data2.length > 0, StakingProtocolCodeRetrievalFailed());

        crossToken = crossTokenAddr;
        stakingPoolCodeContract = IStakingPoolCode(_stakingPoolCodeContract);
        rewardPoolCodeContract = IRewardPoolCode(_rewardPoolCodeContract);
    }

    // ============ 프로젝트 생성 ============

    /**
     * @notice 새 프로젝트 풀 생성
     * @param projectName 프로젝트 이름
     * @param seasonBlocks 시즌 길이 (블록 수, 0이면 기본값 사용)
     * @param firstSeasonStartBlock 첫 시즌 시작 블록
     * @param poolEndBlock 풀 종료 블록 (0이면 무한)
     * @param projectAdmin 프로젝트 관리자 (0이면 msg.sender)
     * @param preDepositStartBlock 사전 예치 시작 블록 (0이면 즉시 가능)
     * @dev CREATE2를 사용하여 주소 예측 가능. salt = keccak256(abi.encode(projectName, {Pool Type}))
     */
    function createProject(
        string calldata projectName,
        uint seasonBlocks,
        uint firstSeasonStartBlock,
        uint poolEndBlock,
        address projectAdmin,
        uint preDepositStartBlock
    ) external nonReentrant returns (uint projectID, address stakingPool, address rewardPool) {
        require(bytes(projectName).length != 0, StakingProtocolEmptyProjectName());
        require(projectIDByName[projectName] == 0, StakingProtocolProjectNameExists());
        require(projectAdmin != address(0), StakingProtocolCanNotZeroAddress());

        if (seasonBlocks == 0) seasonBlocks = defaultSeasonBlocks;
        require(seasonBlocks > 0, StakingProtocolInvalidSeasonBlocks());

        projectCount++;
        projectID = projectCount;

        // 각 풀 배포
        stakingPool = _deployStakingPool(
            projectID, projectName, seasonBlocks, firstSeasonStartBlock, poolEndBlock, preDepositStartBlock
        );
        rewardPool = _deployRewardPool(projectID, projectName, stakingPool);

        // 연결 설정
        _setupPools(stakingPool, rewardPool);

        // 프로젝트 정보 저장
        _saveProjectInfo(projectID, projectName, stakingPool, rewardPool, projectAdmin);

        emit ProjectCreated(projectID, projectName, stakingPool, rewardPool, msg.sender, projectAdmin);
    }

    // ============ 배포 헬퍼 함수 ============

    /**
     * @notice StakingPool 배포 (CREATE2)
     * @dev salt = keccak256(abi.encode(projectName, "StakingPool"))
     */
    function _deployStakingPool(
        uint projectID,
        string calldata projectName,
        uint seasonBlocks,
        uint firstSeasonStartBlock,
        uint poolEndBlock,
        uint preDepositStartBlock
    ) internal returns (address pool) {
        bytes memory code = _getStakingPoolCode();
        bytes memory bytecode = bytes.concat(
            code,
            abi.encode(
                projectID,
                crossToken,
                address(this),
                seasonBlocks,
                firstSeasonStartBlock,
                poolEndBlock,
                preDepositStartBlock
            )
        );

        bytes32 salt = keccak256(abi.encode(projectName, "StakingPool"));

        assembly {
            pool := create2(0, add(bytecode, 0x20), mload(bytecode), salt)
        }
        require(pool != address(0), StakingProtocolDeploymentFailed());
    }

    /**
     * @notice RewardPool 배포 (CREATE2)
     * @dev salt = keccak256(abi.encode(projectName, "RewardPool"))
     */
    function _deployRewardPool(uint, /* projectID */ string calldata projectName, address stakingPool)
        internal
        returns (address pool)
    {
        bytes memory code = _getRewardPoolCode();
        bytes memory bytecode = bytes.concat(code, abi.encode(stakingPool, address(this)));

        bytes32 salt = keccak256(abi.encode(projectName, "RewardPool"));

        assembly {
            pool := create2(0, add(bytecode, 0x20), mload(bytecode), salt)
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
        projectsByAdmin[projectAdmin].push(projectID);
    }

    // ============ Code 가져오기 함수 ============

    function _getStakingPoolCode() internal view returns (bytes memory) {
        return stakingPoolCodeContract.code();
    }

    function _getRewardPoolCode() internal view returns (bytes memory) {
        return rewardPoolCodeContract.code();
    }

    // ============ 관리 함수 ============

    /**
     * @notice 프로젝트 admin 변경 (현재 admin만 가능)
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
     * @notice 프로젝트 상태 변경 (프로토콜 admin만 가능)
     */
    function setProjectStatus(uint projectID, bool isActive) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _checkProjectID(projectID);
        projects[projectID].isActive = isActive;
        emit ProjectStatusUpdated(projectID, isActive);
    }

    /**
     * @notice 글로벌 Router 승인 (프로토콜 admin 전용)
     * @param router Router 주소
     * @param approved 승인 여부
     * @dev 글로벌로 승인된 Router는 모든 프로젝트에서 사용 가능
     */
    function setGlobalApprovedRouter(address router, bool approved) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(router != address(0), StakingProtocolCanNotZeroAddress());
        globalApprovedRouters[router] = approved;
        emit GlobalRouterApprovalUpdated(router, approved);
    }

    /**
     * @notice 프로젝트별 Router 승인 (프로젝트 admin 또는 프로토콜 admin)
     * @param projectID 프로젝트 ID
     * @param router Router 주소
     * @param approved 승인 여부
     * @dev 프로젝트별 승인은 해당 프로젝트에서만 유효
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
        _checkProjectID(projectID);
        IStakingPool(projects[projectID].stakingPool).setRewardPool(rewardPool);
    }

    /**
     * @notice 프로젝트별 수동 시즌 롤오버 (프로젝트 admin 또는 프로토콜 admin)
     * @param projectID 프로젝트 ID
     * @param maxRollovers 최대 롤오버 횟수
     * @return rolloversPerformed 실제로 수행된 롤오버 횟수
     * @dev 50개 이상의 시즌이 쌓인 경우 사용
     */
    function manualRolloverSeasons(uint projectID, uint maxRollovers)
        external
        onlyProjectAdminOrProtocolAdmin(projectID)
        returns (uint rolloversPerformed)
    {
        return IStakingPool(projects[projectID].stakingPool).manualRolloverSeasons(maxRollovers);
    }

    /**
     * @notice RewardPool에서 토큰 회수 (프로토콜 admin 전용)
     * @param projectID 프로젝트 ID
     * @param token 회수할 토큰 주소
     * @param to 수신자 주소
     * @param amount 회수할 수량
     * @dev RewardPool.sweep()는 protocol 주소만 호출 가능하므로 이 함수를 통해 실행
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
     * @notice 프로젝트별 롤오버 대기 중인 시즌 개수 조회
     * @param projectID 프로젝트 ID
     * @return pendingSeasons 롤오버가 필요한 시즌 수
     */
    function getPendingSeasonRollovers(uint projectID) external view returns (uint pendingSeasons) {
        _checkProjectID(projectID);
        return IStakingPool(projects[projectID].stakingPool).getPendingSeasonRollovers();
    }

    // ============ 조회 함수 ============
    function isGlobalApprovedRouter(address router) external view returns (bool) {
        return globalApprovedRouters[router];
    }

    function getProject(uint projectID) external view returns (IStakingProtocol.ProjectInfo memory) {
        _checkProjectID(projectID);
        return projects[projectID];
    }

    function getProjectsByAdmin(address admin) external view returns (uint[] memory) {
        return projectsByAdmin[admin];
    }

    /**
     * @notice StakingPool 주소 계산 (CREATE2 예측)
     * @param projectName 프로젝트 이름
     * @param projectID 프로젝트 ID
     * @param seasonBlocks 시즌 길이
     * @param firstSeasonStartBlock 첫 시즌 시작 블록
     * @param poolEndBlock 풀 종료 블록
     * @return 예측된 StakingPool 주소
     */
    function computeStakingPoolAddress(
        string calldata projectName,
        uint projectID,
        uint seasonBlocks,
        uint firstSeasonStartBlock,
        uint poolEndBlock
    ) external view returns (address) {
        bytes memory code = _getStakingPoolCode();
        bytes memory bytecode = bytes.concat(
            code, abi.encode(projectID, crossToken, address(this), seasonBlocks, firstSeasonStartBlock, poolEndBlock)
        );

        bytes32 salt = keccak256(abi.encode(projectName, "StakingPool"));
        bytes32 hash = keccak256(abi.encodePacked(bytes1(0xff), address(this), salt, keccak256(bytecode)));

        return address(uint160(uint(hash)));
    }

    /**
     * @notice RewardPool 주소 계산 (CREATE2 예측)
     * @param projectName 프로젝트 이름
     * @param stakingPool StakingPool 주소
     * @return 예측된 RewardPool 주소
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

    function _checkProjectID(uint projectID) internal view {
        require(projectID > 0 && projectID <= projectCount, StakingProtocolInvalidProjectID());
    }
}
