// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./BaseTest.sol";

/**
 * @title UnstakeEndTimeTest
 * @notice Unstake 시 endTime 체크 검증 테스트
 * @dev 시즌 종료 후 unstake 시 forfeited points가 endTime까지만 계산되는지 검증
 */
contract UnstakeEndTimeTest is BaseTest {
    /**
     * @notice 시즌 중간 unstake 시 정상 작동 검증 (endTime 이전)
     * @dev 시즌이 아직 종료되지 않았을 때 unstake - 정상 케이스
     */
    function test_UnstakeDuringSeasonBeforeEnd() public {
        // 시즌1 시작 시 100 CROSS stake
        stakeFor(user1, 100 ether);

        uint stakeTime = block.timestamp;

        // 시즌1 중간(30초 경과)
        vm.warp(stakeTime + 30);

        uint totalPointsBefore = stakingPool.seasonTotalPointsSnapshot(1);

        // Unstake
        withdrawFor(user1);

        uint totalPointsAfter = stakingPool.seasonTotalPointsSnapshot(1);
        uint forfeitedPoints = totalPointsBefore - totalPointsAfter;

        // 예상: 100 CROSS × 30초
        uint expectedForfeited = 100 ether * 30;

        console.log("Forfeited points (during season):", forfeitedPoints);
        console.log("Expected forfeited:", expectedForfeited);

        assertApproxEqAbs(forfeitedPoints, expectedForfeited, 100 ether, "Should forfeit points for time staked");
    }

    /**
     * @notice 시즌 중간에 unstake 시 정상 작동 검증
     */
    function test_UnstakeDuringSeason_NormalOperation() public {
        // 시즌1 시작 시 100 CROSS stake
        stakeFor(user1, 100 ether);

        (,, uint season1End,) = stakingPool.getCurrentSeasonInfo();

        // 시즌1 중간(50초 경과)
        vm.warp(block.timestamp + 50);

        uint totalPointsBefore = stakingPool.seasonTotalPointsSnapshot(1);

        // Unstake
        withdrawFor(user1);

        uint totalPointsAfter = stakingPool.seasonTotalPointsSnapshot(1);
        uint forfeitedPoints = totalPointsBefore - totalPointsAfter;

        // 예상: 100 CROSS × 50초
        uint expectedForfeited = 100 ether * 50;

        console.log("Forfeited points (mid-season):", forfeitedPoints);
        console.log("Expected forfeited:", expectedForfeited);

        assertApproxEqAbs(forfeitedPoints, expectedForfeited, 100 ether, "Mid-season unstake should work correctly");
    }

    /**
     * @notice 시즌 종료 직전 unstake 검증
     */
    function test_UnstakeJustBeforeSeasonEnd() public {
        stakeFor(user1, 100 ether);

        (,, uint season1End,) = stakingPool.getCurrentSeasonInfo();

        // 시즌 종료 1초 전
        vm.warp(season1End);

        uint totalPointsBefore = stakingPool.seasonTotalPointsSnapshot(1);

        withdrawFor(user1);

        uint totalPointsAfter = stakingPool.seasonTotalPointsSnapshot(1);
        uint forfeitedPoints = totalPointsBefore - totalPointsAfter;

        // 거의 전체 시즌 참여
        uint expectedForfeited = 100 ether * (SEASON_DURATION - 1);

        assertApproxEqAbs(forfeitedPoints, expectedForfeited, 100 ether, "Should forfeit almost full season points");
    }

    /**
     * @notice 시즌 종료 직후 unstake 검증 (롤오버 발생)
     * @dev 시즌 종료 후에는 자동 롤오버되어 새 시즌에서 unstake됨
     */
    function test_UnstakeAfterSeasonEnd_WithRollover() public {
        // 시즌1에 stake
        stakeFor(user1, 100 ether);

        (,, uint season1End,) = stakingPool.getCurrentSeasonInfo();

        // 시즌1 종료 직후
        vm.warp(season1End + 1);

        // Unstake - 이 과정에서 롤오버 발생
        withdrawFor(user1);

        // 롤오버 후 currentSeason 확인
        uint currentSeasonAfter = stakingPool.currentSeason();
        assertEq(currentSeasonAfter, 2, "Should have rolled over to season 2");

        // 시즌1의 포인트는 finalize되어 저장됨
        uint season1Points = stakingPool.seasonTotalPointsSnapshot(1);

        console.log("Season 1 total points:", season1Points);
        console.log("Current season after unstake:", currentSeasonAfter);

        // 시즌1 전체 참여했으므로 포인트가 있어야 함
        assertGt(season1Points, 0, "Season 1 should have points");

        // 대략 100 CROSS × (SEASON_DURATION - 1)초
        uint expectedPoints = 100 ether * (SEASON_DURATION - 1);
        assertApproxEqAbs(season1Points, expectedPoints, 100 ether, "Season 1 points should be correct");
    }

    /**
     * @notice UserSeasonData가 있는 경우 시즌 중 unstake 검증
     */
    function test_UnstakeWithUserSeasonData_DuringSeason() public {
        // 시즌1에 stake
        stakeFor(user1, 100 ether);

        uint firstStakeTime = block.timestamp;

        // 시즌 중간에 추가 stake (UserSeasonData 업데이트)
        vm.warp(firstStakeTime + 30);
        vm.prank(user1);
        router.stake{value: 50 ether}(PROJECT_ID);

        // 추가 20초 후 unstake (시즌 종료 전)
        vm.warp(block.timestamp + 20);

        uint totalPointsBefore = stakingPool.seasonTotalPointsSnapshot(1);

        // Unstake
        withdrawFor(user1);

        uint totalPointsAfter = stakingPool.seasonTotalPointsSnapshot(1);
        uint forfeitedPoints = totalPointsBefore - totalPointsAfter;

        // 예상: 100 CROSS × 30초 + 150 CROSS × 20초
        uint expectedPart1 = 100 ether * 30;
        uint expectedPart2 = 150 ether * 20;
        uint expectedTotal = expectedPart1 + expectedPart2;

        console.log("Forfeited (with UserSeasonData):", forfeitedPoints);
        console.log("Expected:", expectedTotal);

        assertApproxEqAbs(
            forfeitedPoints, expectedTotal, 150 ether, "Should forfeit correct points with UserSeasonData"
        );
    }
}
