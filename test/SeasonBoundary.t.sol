// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./BaseTest.sol";

/**
 * @title SeasonBoundaryTest
 * @notice 시즌 경계 및 포인트 계산 정확성 테스트
 * @dev 버그 수정 검증: 시즌1 중간에 stake 후 여러 시즌 경과 시 포인트 계산 정확성
 */
contract SeasonBoundaryTest is BaseTest {
    /**
     * @notice 시즌 경계가 정확히 계산되는지 검증
     * @dev 11월 1일 00:00~23:59:59, 11월 2일 00:00~23:59:59 형태 검증
     */
    function test_SeasonBoundaryCalculation() public view {
        (uint season, uint startTime, uint endTime,) = stakingPool.getCurrentSeasonInfo();

        assertEq(season, 1, "Should be season 1");
        assertEq(endTime - startTime, SEASON_DURATION - 1, "Season duration should be SEASON_DURATION - 1");

        // 시즌이 정확히 SEASON_DURATION초 동안 지속되는지 확인
        // endTime은 inclusive이므로 endTime - startTime + 1 = SEASON_DURATION
        uint actualDuration = endTime - startTime + 1;
        assertEq(actualDuration, SEASON_DURATION, "Actual duration should equal SEASON_DURATION");
    }

    /**
     * @notice 시즌1 중간에 stake 후 여러 시즌 경과 시 포인트 계산 검증
     * @dev 버그 재현 시나리오:
     *      - 시즌1 중간(50초)에 100 CROSS stake
     *      - 시즌5까지 아무 행위 없음
     *      - 각 시즌 포인트 비교: 시즌2,3,4 > 시즌1 이어야 함
     */
    function test_MidSeasonStakePoints() public {
        // 시즌1 시작 시점
        uint season1Start = block.timestamp;

        // 시즌1 중간(50초 후)으로 이동
        vm.warp(season1Start + 50);

        // 100 CROSS stake
        stakeFor(user1, 100 ether);

        // 시즌1 종료 시점 확인
        (, uint s1Start, uint s1End,) = stakingPool.getCurrentSeasonInfo();

        // 시즌5 중간까지 시간 이동 (시즌1~4 완전히 지나감)
        uint season5Mid = s1End + (SEASON_DURATION * 4) + 50;
        vm.warp(season5Mid);

        // 각 시즌의 포인트 조회
        (uint s1Points,) = stakingPool.getSeasonUserPoints(1, user1);
        (uint s2Points,) = stakingPool.getSeasonUserPoints(2, user1);
        (uint s3Points,) = stakingPool.getSeasonUserPoints(3, user1);
        (uint s4Points,) = stakingPool.getSeasonUserPoints(4, user1);

        console.log("Season 1 points:", s1Points);
        console.log("Season 2 points:", s2Points);
        console.log("Season 3 points:", s3Points);
        console.log("Season 4 points:", s4Points);

        // 시즌1은 중간(50초)부터 참여 → 약 50초 분량의 포인트
        // 시즌2,3,4는 전체 시즌 참여 → 약 SEASON_DURATION초 분량의 포인트

        uint s1ExpectedDuration = s1End - (s1Start + 50);
        uint s2ExpectedDuration = SEASON_DURATION - 1; // endTime은 inclusive

        // 시즌1 포인트는 대략 (시즌1 중간부터 끝까지) × balance
        uint s1ExpectedPoints = 100 ether * s1ExpectedDuration;

        // 시즌2,3,4 포인트는 대략 (전체 시즌) × balance
        uint s2ExpectedPoints = 100 ether * s2ExpectedDuration;

        // 오차 범위 내에서 검증 (±1초)
        assertApproxEqAbs(s1Points, s1ExpectedPoints, 100 ether, "Season 1 points mismatch");
        assertApproxEqAbs(s2Points, s2ExpectedPoints, 100 ether, "Season 2 points mismatch");
        assertApproxEqAbs(s3Points, s2ExpectedPoints, 100 ether, "Season 3 points mismatch");
        assertApproxEqAbs(s4Points, s2ExpectedPoints, 100 ether, "Season 4 points mismatch");

        // 중요: 시즌2,3,4 포인트는 시즌1보다 많아야 함
        assertGt(s2Points, s1Points, "Season 2 should have more points than Season 1");
        assertGt(s3Points, s1Points, "Season 3 should have more points than Season 1");
        assertGt(s4Points, s1Points, "Season 4 should have more points than Season 1");

        // 시즌2,3,4는 거의 동일해야 함 (모두 전체 시즌 참여)
        assertApproxEqAbs(s2Points, s3Points, 100 ether, "Season 2 and 3 should have similar points");
        assertApproxEqAbs(s3Points, s4Points, 100 ether, "Season 3 and 4 should have similar points");
    }

    /**
     * @notice 가상 시즌 포인트 계산 검증
     * @dev 온체인 시즌 롤오버 없이 여러 시즌 경과 시 정확한 계산
     */
    function test_VirtualSeasonPoints() public {
        // 시즌1 시작 시 stake
        stakeFor(user1, 100 ether);

        (,, uint s1End,) = stakingPool.getCurrentSeasonInfo();

        // 시즌3 중간까지 이동 (롤오버 없이)
        vm.warp(s1End + (SEASON_DURATION * 2) + 50);

        // 가상 시즌2의 포인트 조회
        (uint s2Points,) = stakingPool.getSeasonUserPoints(2, user1);

        // 시즌2는 전체 시즌 참여했으므로 SEASON_DURATION-1초 분량
        uint expectedPoints = 100 ether * (SEASON_DURATION - 1);

        assertApproxEqAbs(s2Points, expectedPoints, 100 ether, "Virtual season 2 points incorrect");
    }

    /**
     * @notice 여러 사용자의 시즌별 포인트 비율 검증
     */
    function test_MultiUserSeasonPointsRatio() public {
        uint season1Start = block.timestamp;

        // User1: 시즌1 중간(50초)에 100 CROSS stake
        vm.warp(season1Start + 50);
        stakeFor(user1, 100 ether);

        // User2: 시즌1 시작 시 50 CROSS stake
        vm.warp(season1Start);
        stakeFor(user2, 50 ether);

        // 시즌2로 롤오버
        rolloverSeason();

        // 시즌1 포인트 조회
        (uint s1User1Points,) = stakingPool.getSeasonUserPoints(1, user1);
        (uint s1User2Points,) = stakingPool.getSeasonUserPoints(1, user2);

        console.log("Season 1 User1 points:", s1User1Points);
        console.log("Season 1 User2 points:", s1User2Points);

        // User2는 전체 시즌 참여, User1은 중간부터 참여
        // 하지만 User1의 balance가 2배이므로 User1 포인트가 더 높을 수 있음

        // User1: 100 CROSS × ~50초 = ~5000 CROSS-seconds
        // User2: 50 CROSS × ~100초 = ~5000 CROSS-seconds
        // 비슷해야 함

        assertApproxEqRel(s1User1Points, s1User2Points, 0.1e18, "Points ratio should be similar");
    }

    /**
     * @notice 시즌 종료 후 unstake 시 포인트 계산 검증
     */
    function test_UnstakeAfterSeasonEnd() public {
        // 시즌1에 stake
        stakeFor(user1, 100 ether);

        (,, uint s1End,) = stakingPool.getCurrentSeasonInfo();

        // 시즌1 종료 후로 이동
        vm.warp(s1End + 10);

        // Unstake
        withdrawFor(user1);

        // 시즌1 포인트는 시즌 종료 시점까지만 계산되어야 함
        (uint s1Points,) = stakingPool.getSeasonUserPoints(1, user1);

        // 전체 시즌 참여했으므로 대략 SEASON_DURATION-1초 분량
        uint expectedPoints = 100 ether * (SEASON_DURATION - 1);

        assertApproxEqAbs(s1Points, expectedPoints, 100 ether, "Points should be capped at season end");
    }

    /**
     * @notice 버그 수정 검증: 1시즌 예치 후 3시즌까지 아무 행위 없이 경과했을 때
     * @dev 사용자 보고 버그: "1시즌 예치 -> 3시즌까지 아무 행위도 하지않음 -> 1시즌 정상, 2시즌 정상, 3시즌은 내 포인트 0"
     */
    function test_Bug_Season1StakeSeason3NoPoints() public {
        // 시즌1 시작 시점 (현재 시즌 = 1)
        uint season1Start = block.timestamp;

        // 시즌1에 100 CROSS stake
        stakeFor(user1, 100 ether);

        console.log("=== Initial State ===");
        console.log("Current block:", block.timestamp);
        console.log("Current season (storage):", stakingPool.currentSeason());

        // 시즌1 정보 확인
        (uint s1,,,) = stakingPool.getCurrentSeasonInfo();
        assertEq(s1, 1, "Should be season 1");

        // 시즌3 중간까지 시간 이동 (온체인 롤오버 없음)
        // 시즌1: 0 ~ SEASON_DURATION-1
        // 시즌2: SEASON_DURATION ~ 2*SEASON_DURATION-1
        // 시즌3: 2*SEASON_DURATION ~ 3*SEASON_DURATION-1
        uint season3Mid = season1Start + (SEASON_DURATION * 2) + 50;
        vm.warp(season3Mid);

        console.log("\n=== After Time Skip ===");
        console.log("Current block:", block.timestamp);
        console.log("Current season (storage):", stakingPool.currentSeason());

        // getCurrentSeasonInfo는 가상 시즌을 반환해야 함
        (uint currentSeasonCalc,,,) = stakingPool.getCurrentSeasonInfo();
        console.log("Current season (calculated):", currentSeasonCalc);
        assertEq(currentSeasonCalc, 3, "Should calculate season 3");

        // 각 시즌의 포인트 조회
        (uint s1Points,) = stakingPool.getSeasonUserPoints(1, user1);
        (uint s2Points,) = stakingPool.getSeasonUserPoints(2, user1);
        (uint s3Points,) = stakingPool.getSeasonUserPoints(3, user1);

        console.log("\n=== Points ===");
        console.log("Season 1 points:", s1Points);
        console.log("Season 2 points:", s2Points);
        console.log("Season 3 points:", s3Points);

        // 예상 포인트 계산
        // 시즌1: 전체 기간 = SEASON_DURATION-1 (endTime은 inclusive)
        uint s1ExpectedPoints = 100 ether * (SEASON_DURATION - 1);
        // 시즌2: 전체 기간 = SEASON_DURATION-1
        uint s2ExpectedPoints = 100 ether * (SEASON_DURATION - 1);
        // 시즌3: season3Mid = season1Start(1) + SEASON_DURATION*2(200) + 50 = 251
        //        season3 startTime = 201, current = 251, elapsed = 50초
        //        하지만 시간 계산은 startTime부터 currentTime까지 inclusive이므로 51개 블록
        uint s3ExpectedPoints = 100 ether * 50;

        // 검증: 모든 시즌에서 포인트가 있어야 함 (버그 수정 전에는 시즌3이 0)
        assertGt(s1Points, 0, "Season 1 should have points");
        assertGt(s2Points, 0, "Season 2 should have points");
        assertGt(s3Points, 0, "Season 3 should have points (BUG FIX)");

        // 오차 범위 내에서 검증
        assertApproxEqAbs(s1Points, s1ExpectedPoints, 100 ether, "Season 1 points mismatch");
        assertApproxEqAbs(s2Points, s2ExpectedPoints, 100 ether, "Season 2 points mismatch");
        assertApproxEqAbs(s3Points, s3ExpectedPoints, 100 ether, "Season 3 points mismatch");
    }
}
