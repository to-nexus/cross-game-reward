// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./WCROSS.sol";

import "./interfaces/IRewardPool.sol";
import "./interfaces/IStakingPool.sol";
import "./interfaces/IStakingProtocol.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

/**
 * @title StakingRouter
 * @notice StakingProtocol 전체를 지원하는 Native CROSS 스테이킹 라우터
 * @dev 사용자는 Native CROSS만 보내면 자동으로 wrap/unwrap 처리됨
 */
contract StakingRouter is ReentrancyGuardTransient {
    using SafeERC20 for IERC20;

    // ============ 에러 ============

    error StakingRouterCanNotZeroAddress();
    error StakingRouterAmountMustBeGreaterThanZero();
    error StakingRouterNoStake();
    error StakingRouterInvalidProjectID();
    error StakingRouterTransferFailed();
    error StakingRouterOnlyWCROSS();
    error StakingRouterLengthMismatch();
    error StakingRouterInvalidOffset();
    error StakingRouterInvalidRange();

    // ============ Modifiers ============

    /**
     * @notice projectID 검증 modifier
     */
    modifier validProject(uint projectID) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        _;
    }

    // ============ 상태 변수 ============

    /// @notice WCROSS 토큰 주소
    WCROSS public immutable wcross;

    /// @notice StakingProtocol 주소
    IStakingProtocol public immutable protocol;

    // ============ 이벤트 ============

    event StakedWithNative(address indexed user, uint indexed projectID, uint amount);
    event UnstakedToNative(address indexed user, uint indexed projectID, uint amount);
    event RewardClaimed(
        address indexed user, uint indexed projectID, uint indexed season, address rewardToken, uint amount
    );

    // ============ 생성자 ============

    constructor(address _wcross, address _protocol) {
        require(_wcross != address(0), StakingRouterCanNotZeroAddress());
        require(_protocol != address(0), StakingRouterCanNotZeroAddress());

        wcross = WCROSS(payable(_wcross));
        protocol = IStakingProtocol(_protocol);
    }

    // ============ Internal Helpers ============

    /**
     * @notice 프로젝트 정보 조회 헬퍼 (중복 로직 공통화)
     * @param projectID 프로젝트 ID
     * @return stakingPool StakingPool 주소
     * @return rewardPool RewardPool 주소
     */
    function _getProjectPools(uint projectID)
        internal
        view
        validProject(projectID)
        returns (address stakingPool, address rewardPool)
    {
        (stakingPool, rewardPool,,,,,) = protocol.projects(projectID);
    }

    // ============ 스테이킹 함수 (Native CROSS 사용) ============

    /**
     * @notice Native CROSS로 스테이킹 (자동 wrap)
     * @param projectID 프로젝트 ID
     * @dev msg.value를 WCROSS로 변환 후 StakingPool.stakeFor() 호출
     */
    function stake(uint projectID) external payable nonReentrant validProject(projectID) {
        uint amount = msg.value;
        require(amount > 0, StakingRouterAmountMustBeGreaterThanZero());

        // 프로젝트 정보 조회
        (address stakingPool,) = _getProjectPools(projectID);

        // 1. Native CROSS를 WCROSS로 변환
        wcross.deposit{value: amount}();

        // 2. WCROSS를 StakingPool에 approve
        IERC20(address(wcross)).safeIncreaseAllowance(stakingPool, amount);

        // 3. 사용자를 위해 스테이킹 (stakeFor 사용)
        IStakingPool(stakingPool).stakeFor(msg.sender, amount);

        emit StakedWithNative(msg.sender, projectID, amount);
    }

    /**
     * @notice StakingPool에서 출금 후 Native CROSS로 자동 변환
     * @param projectID 프로젝트 ID
     * @dev 1회 요청으로 withdraw + unwrap 처리
     */
    function unstake(uint projectID) external nonReentrant validProject(projectID) {
        // 프로젝트 정보 조회
        (address stakingPool,) = _getProjectPools(projectID);

        // 1. StakingPool에서 사용자를 대신해 출금 (WCROSS를 router가 받음)
        IStakingPool(stakingPool).withdrawAllFor(msg.sender);

        // 2. Router가 받은 WCROSS 수량 확인
        uint wcrossBalance = IERC20(address(wcross)).balanceOf(address(this));
        require(wcrossBalance > 0, StakingRouterNoStake());

        // 3. WCROSS를 Native CROSS로 변환
        wcross.withdraw(wcrossBalance);

        // 4. Native CROSS를 사용자에게 전송
        (bool success,) = msg.sender.call{value: wcrossBalance}("");
        require(success, StakingRouterTransferFailed());

        emit UnstakedToNative(msg.sender, projectID, wcrossBalance);
    }

    /**
     * @notice WCROSS withdraw 시에만 Native CROSS 수신 가능
     */
    receive() external payable {
        require(msg.sender == address(wcross), StakingRouterOnlyWCROSS());
    }

    // ============ 리워드 청구 함수 ============

    /**
     * @notice 특정 시즌의 리워드 청구
     * @param projectID 프로젝트 ID
     * @param season 시즌 번호
     * @param rewardToken 리워드 토큰 주소
     */
    function claimReward(uint projectID, uint season, address rewardToken) external nonReentrant {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());

        (address stakingPool,,,,,,) = protocol.projects(projectID);

        // 사용자를 대신해 리워드 청구 (Router 승인 필요)
        IStakingPool(stakingPool).claimSeasonFor(msg.sender, season, rewardToken);

        emit RewardClaimed(msg.sender, projectID, season, rewardToken, 0);
    }

    /**
     * @notice 여러 시즌의 리워드 일괄 청구
     * @param projectID 프로젝트 ID
     * @param seasons 시즌 번호 배열
     * @param rewardTokens 리워드 토큰 주소 배열
     */
    function claimMultipleRewards(uint projectID, uint[] calldata seasons, address[] calldata rewardTokens)
        external
        nonReentrant
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        require(seasons.length == rewardTokens.length, StakingRouterLengthMismatch());

        (address stakingPool,,,,,,) = protocol.projects(projectID);

        for (uint i = 0; i < seasons.length;) {
            IStakingPool(stakingPool).claimSeasonFor(msg.sender, seasons[i], rewardTokens[i]);
            emit RewardClaimed(msg.sender, projectID, seasons[i], rewardTokens[i], 0);
            unchecked {
                ++i;
            }
        }
    }

    // ============ 시즌 최적화 함수 ============

    /**
     * @notice 유저의 이전 시즌 데이터를 배치로 finalize
     * @param projectID 프로젝트 ID
     * @param user 사용자 주소
     * @param maxSeasons 최대 처리 시즌 수
     * @return processed 처리된 시즌 수
     */
    function finalizeUserSeasonsBatch(uint projectID, address user, uint maxSeasons)
        external
        returns (uint processed)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);

        // StakingPool의 finalizeUserSeasonsBatch 호출 (존재하는 경우)
        // 하드코딩된 시그니처로 호출
        (bool success, bytes memory data) =
            stakingPool.call(abi.encodeWithSignature("finalizeUserSeasonsBatch(address,uint256)", user, maxSeasons));

        if (success && data.length >= 32) processed = abi.decode(data, (uint));

        return processed;
    }

    // (조회 전용 함수들은 StakingViewer로 이전)
}
