// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

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

    // ============ 스테이킹 함수 (Native CROSS 사용) ============

    /**
     * @notice Native CROSS로 스테이킹 (자동 wrap)
     * @param projectID 프로젝트 ID
     * @dev msg.value를 WCROSS로 변환 후 StakingPool.stakeFor() 호출
     */
    function stake(uint projectID) external payable nonReentrant {
        uint amount = msg.value;
        require(amount > 0, StakingRouterAmountMustBeGreaterThanZero());
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());

        // 프로젝트 정보 조회
        (address stakingPool,,,,,,) = protocol.projects(projectID);

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
    function unstake(uint projectID) external nonReentrant {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());

        // 프로젝트 정보 조회
        (address stakingPool,,,,,,) = protocol.projects(projectID);

        // 1. StakingPool에서 사용자를 대신해 출금 (WCROSS를 router가 받음)
        IStakingPool(stakingPool).withdrawAllFor(msg.sender);

        // 2. Router가 받은 WCROSS 수량 확인
        uint wcrossBalance = IERC20(address(wcross)).balanceOf(address(this));
        require(wcrossBalance > 0, StakingRouterNoStake());

        // 3. WCROSS를 Native CROSS로 변환
        wcross.withdraw(wcrossBalance);

        // 4. Native CROSS를 사용자에게 전송
        (bool success,) = msg.sender.call{value: wcrossBalance}("");
        require(success, "Transfer failed");

        emit UnstakedToNative(msg.sender, projectID, wcrossBalance);
    }

    /**
     * @notice WCROSS withdraw 시에만 Native CROSS 수신 가능
     */
    receive() external payable {
        // WCROSS 컨트랙트로부터만 ETH 수신 허용
        require(msg.sender == address(wcross), "Only WCROSS");
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

        // StakingPool을 통해 리워드 청구
        IStakingPool(stakingPool).claimSeason(season, rewardToken);

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
        require(seasons.length == rewardTokens.length, "Length mismatch");

        (address stakingPool,,,,,,) = protocol.projects(projectID);

        for (uint i = 0; i < seasons.length;) {
            IStakingPool(stakingPool).claimSeason(seasons[i], rewardTokens[i]);
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

    // ============ 조회 함수 ============

    /**
     * @notice 사용자의 스테이킹 정보 조회
     * @param projectID 프로젝트 ID
     * @param user 사용자 주소
     * @return balance 스테이킹 수량
     * @return points 누적 포인트
     * @return lastUpdateBlock 마지막 업데이트 블록
     */
    function getStakeInfo(uint projectID, address user)
        external
        view
        returns (uint balance, uint points, uint lastUpdateBlock)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).getStakePosition(user);
    }

    /**
     * @notice 시즌 정보 조회
     * @param projectID 프로젝트 ID
     * @return currentSeason 현재 시즌
     * @return seasonStartBlock 시즌 시작 블록
     * @return seasonEndBlock 시즌 종료 블록
     * @return blocksRemaining 남은 블록 수
     */
    function getSeasonInfo(uint projectID)
        external
        view
        returns (uint currentSeason, uint seasonStartBlock, uint seasonEndBlock, uint blocksRemaining)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).getCurrentSeasonInfo();
    }

    /**
     * @notice 프로젝트 정보 조회
     * @param projectID 프로젝트 ID
     * @return stakingPool StakingPool 주소
     * @return rewardPool RewardPool 주소
     * @return name 프로젝트 이름
     * @return isActive 활성 상태
     * @return createdAt 생성 시간
     */
    function getProjectInfo(uint projectID)
        external
        view
        returns (address stakingPool, address rewardPool, string memory name, bool isActive, uint createdAt)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (stakingPool, rewardPool, name, isActive, createdAt,,) = protocol.projects(projectID);
    }

    /**
     * @notice 사용자의 시즌별 청구 가능 보상 조회
     * @param projectID 프로젝트 ID
     * @param user 사용자 주소
     * @param season 시즌 번호
     * @param rewardToken 보상 토큰 주소
     * @return claimableAmount 청구 가능한 보상 수량
     */
    function getClaimableReward(uint projectID, address user, uint season, address rewardToken)
        external
        view
        returns (uint claimableAmount)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (, address rewardPool,,,,,) = protocol.projects(projectID);
        return IRewardPool(rewardPool).getExpectedReward(user, season, rewardToken);
    }

    /**
     * @notice 프로젝트의 총 스테이킹 수량 조회
     * @param projectID 프로젝트 ID
     * @return totalStaked 총 스테이킹 수량
     */
    function getTotalStaked(uint projectID) external view returns (uint totalStaked) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).totalStaked();
    }

    /**
     * @notice 시즌별 총 포인트 조회
     * @param projectID 프로젝트 ID
     * @param season 시즌 번호
     * @return totalPoints 총 포인트
     */
    function getSeasonTotalPoints(uint projectID, uint season) external view returns (uint totalPoints) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).seasonTotalPointsSnapshot(season);
    }

    /**
     * @notice 시즌별 예상 포인트 조회 (스냅샷 또는 계산)
     * @param projectID 프로젝트 ID
     * @param season 시즌 번호
     * @param user 사용자 주소
     * @return expectedPoints 예상 포인트
     */
    function getExpectedSeasonPoints(uint projectID, uint season, address user)
        external
        view
        returns (uint expectedPoints)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).getExpectedSeasonPoints(season, user);
    }

    /**
     * @notice 시즌별 예상 리워드 조회
     * @param projectID 프로젝트 ID
     * @param season 시즌 번호
     * @param user 사용자 주소
     * @param rewardToken 리워드 토큰 주소
     * @return expectedReward 예상 리워드
     */
    function getExpectedSeasonReward(uint projectID, uint season, address user, address rewardToken)
        external
        view
        returns (uint expectedReward)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).getExpectedSeasonReward(season, user, rewardToken);
    }

    /**
     * @notice 사용자의 스테이킹 파워 조회
     * @param projectID 프로젝트 ID
     * @param user 사용자 주소
     * @return stakingPower 스테이킹 파워
     */
    function getStakingPower(uint projectID, address user) external view returns (uint stakingPower) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).getStakingPower(user);
    }

    /**
     * @notice 사용자의 누적 포인트 조회 (현재 시즌)
     * @param projectID 프로젝트 ID
     * @param user 사용자 주소
     * @return points 누적 포인트
     */
    function getUserPoints(uint projectID, address user) external view returns (uint points) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).getUserPoints(user);
    }

    /**
     * @notice 시즌별 사용자 포인트 조회
     * @param projectID 프로젝트 ID
     * @param season 시즌 번호
     * @param user 사용자 주소
     * @return points 시즌별 포인트
     */
    function getSeasonUserPoints(uint projectID, uint season, address user) external view returns (uint points) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).getSeasonUserPoints(season, user);
    }

    /**
     * @notice 시즌 활성 상태 조회
     * @param projectID 프로젝트 ID
     * @return isActive 시즌 활성 여부
     */
    function isSeasonActive(uint projectID) external view returns (bool isActive) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).isSeasonActive();
    }

    /**
     * @notice 풀 종료 블록 조회
     * @param projectID 프로젝트 ID
     * @return endBlock 풀 종료 블록
     */
    function getPoolEndBlock(uint projectID) external view returns (uint endBlock) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).poolEndBlock();
    }

    /**
     * @notice 다음 시즌 시작 블록 조회
     * @param projectID 프로젝트 ID
     * @return startBlock 다음 시즌 시작 블록
     */
    function getNextSeasonStartBlock(uint projectID) external view returns (uint startBlock) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).nextSeasonStartBlock();
    }

    /**
     * @notice 현재 시즌 번호 조회
     * @param projectID 프로젝트 ID
     * @return season 현재 시즌 번호
     */
    function getCurrentSeason(uint projectID) external view returns (uint season) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).currentSeason();
    }

    /**
     * @notice 프로젝트의 총 스테이킹 파워 조회
     * @param projectID 프로젝트 ID
     * @return totalPower 총 스테이킹 파워
     */
    function getTotalStakingPower(uint projectID) external view returns (uint totalPower) {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);
        return IStakingPool(stakingPool).getTotalStakingPower();
    }

    // ============ 편의 View 함수 (배치 조회) ============

    /**
     * @notice 사용자의 시즌별 데이터 조회
     * @param projectID 프로젝트 ID
     * @param user 사용자 주소
     * @param season 시즌 번호
     * @return points 시즌별 포인트
     * @return balance 시즌별 잔액
     * @return joinBlock 참여 블록
     * @return claimed 청구 여부
     * @return finalized finalize 여부
     */
    function getUserSeasonData(uint projectID, address user, uint season)
        external
        view
        returns (uint points, uint balance, uint joinBlock, bool claimed, bool finalized)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);

        // StakingPool의 getUserSeasonData 호출 (low-level call로 처리)
        (bool success, bytes memory data) =
            stakingPool.staticcall(abi.encodeWithSignature("getUserSeasonData(uint256,address)", season, user));

        if (success && data.length >= 160) {
            // 5 * 32 bytes
            (points, balance, joinBlock, claimed, finalized) = abi.decode(data, (uint, uint, uint, bool, bool));
        }
    }

    /**
     * @notice 리워드 미리보기 (청구 가능 금액 확인)
     * @param projectID 프로젝트 ID
     * @param user 사용자 주소
     * @param season 시즌 번호
     * @param rewardToken 리워드 토큰 주소
     * @return expectedReward 예상 리워드
     * @return userPoints 사용자 포인트
     * @return totalPoints 총 포인트
     * @return alreadyClaimed 이미 청구 여부
     */
    function previewClaim(uint projectID, address user, uint season, address rewardToken)
        external
        view
        returns (uint expectedReward, uint userPoints, uint totalPoints, bool alreadyClaimed)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);

        // StakingPool의 previewClaim 호출
        (bool success, bytes memory data) = stakingPool.staticcall(
            abi.encodeWithSignature("previewClaim(uint256,address,address)", season, user, rewardToken)
        );

        if (success && data.length >= 128) {
            // 4 * 32 bytes
            (expectedReward, userPoints, totalPoints, alreadyClaimed) = abi.decode(data, (uint, uint, uint, bool));
        }
    }

    /**
     * @notice 여러 시즌의 리워드 예상치 일괄 조회
     * @param projectID 프로젝트 ID
     * @param user 사용자 주소
     * @param seasons 시즌 번호 배열
     * @param rewardToken 리워드 토큰 주소
     * @return expectedRewards 예상 리워드 배열
     */
    function getExpectedRewardsBatch(uint projectID, address user, uint[] calldata seasons, address rewardToken)
        external
        view
        returns (uint[] memory expectedRewards)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        (address stakingPool,,,,,,) = protocol.projects(projectID);

        expectedRewards = new uint[](seasons.length);
        for (uint i = 0; i < seasons.length;) {
            expectedRewards[i] = IStakingPool(stakingPool).getExpectedSeasonReward(seasons[i], user, rewardToken);
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice 활성 프로젝트 목록 조회 (페이지네이션)
     * @param offset 시작 인덱스
     * @param limit 조회 개수
     * @return projectIDs 프로젝트 ID 배열
     * @return stakingPools StakingPool 주소 배열
     * @return names 프로젝트 이름 배열
     */
    function getActiveProjects(uint offset, uint limit)
        external
        view
        returns (uint[] memory projectIDs, address[] memory stakingPools, string[] memory names)
    {
        uint totalProjects = protocol.projectCount();
        require(offset < totalProjects, "Invalid offset");

        uint end = offset + limit;
        if (end > totalProjects) end = totalProjects;

        uint count = end - offset;
        projectIDs = new uint[](count);
        stakingPools = new address[](count);
        names = new string[](count);

        uint resultIndex = 0;
        for (uint i = offset; i < end;) {
            uint projectID = i + 1;
            (address stakingPool,, string memory name, bool isActive,,,) = protocol.projects(projectID);

            if (isActive) {
                projectIDs[resultIndex] = projectID;
                stakingPools[resultIndex] = stakingPool;
                names[resultIndex] = name;
                unchecked {
                    ++resultIndex;
                }
            }
            unchecked {
                ++i;
            }
        }

        // 실제 활성 프로젝트 수에 맞춰 배열 크기 조정
        if (resultIndex < count) {
            assembly {
                mstore(projectIDs, resultIndex)
                mstore(stakingPools, resultIndex)
                mstore(names, resultIndex)
            }
        }
    }

    /**
     * @notice 사용자의 모든 프로젝트 스테이킹 요약 정보
     * @param user 사용자 주소
     * @param projectIDs 조회할 프로젝트 ID 배열
     * @return balances 스테이킹 잔액 배열
     * @return points 포인트 배열
     * @return currentSeasons 현재 시즌 배열
     */
    function getUserStakingSummary(address user, uint[] calldata projectIDs)
        external
        view
        returns (uint[] memory balances, uint[] memory points, uint[] memory currentSeasons)
    {
        balances = new uint[](projectIDs.length);
        points = new uint[](projectIDs.length);
        currentSeasons = new uint[](projectIDs.length);

        for (uint i = 0; i < projectIDs.length;) {
            if (projectIDs[i] > 0 && projectIDs[i] <= protocol.projectCount()) {
                (address stakingPool,,,,,,) = protocol.projects(projectIDs[i]);
                (balances[i], points[i],) = IStakingPool(stakingPool).getStakePosition(user);
                currentSeasons[i] = IStakingPool(stakingPool).currentSeason();
            }
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice 프로젝트의 시즌 히스토리 조회
     * @param projectID 프로젝트 ID
     * @param fromSeason 시작 시즌
     * @param toSeason 종료 시즌
     * @return seasons 시즌 배열
     * @return totalPoints 시즌별 총 포인트 배열
     */
    function getSeasonHistory(uint projectID, uint fromSeason, uint toSeason)
        external
        view
        returns (uint[] memory seasons, uint[] memory totalPoints)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), StakingRouterInvalidProjectID());
        require(fromSeason <= toSeason, "Invalid range");

        (address stakingPool,,,,,,) = protocol.projects(projectID);
        uint currentSeason = IStakingPool(stakingPool).currentSeason();

        if (toSeason > currentSeason) toSeason = currentSeason;
        if (fromSeason > toSeason) return (new uint[](0), new uint[](0));

        uint count = toSeason - fromSeason + 1;
        seasons = new uint[](count);
        totalPoints = new uint[](count);

        for (uint i = 0; i < count;) {
            uint season = fromSeason + i;
            seasons[i] = season;
            totalPoints[i] = IStakingPool(stakingPool).seasonTotalPointsSnapshot(season);
            unchecked {
                ++i;
            }
        }
    }
}
