// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./BaseTest.sol";

/**
 * @title PointsTest
 * @notice 포인트 계산 및 관리 테스트
 */
contract PointsTest is BaseTest {
    function test_PointsAccumulation() public {
        stakeFor(user1, 10 ether);

        vm.warp(block.timestamp + 50);
        updatePointsFor(user1);

        uint points = stakingPool.getUserPoints(user1);
        assertGt(points, 0);
    }

    function test_PointsCalculation() public {
        stakeFor(user1, 10 ether);

        vm.warp(block.timestamp + 100);
        updatePointsFor(user1);

        uint points1 = stakingPool.getUserPoints(user1);

        vm.warp(block.timestamp + 100);
        updatePointsFor(user1);

        uint points2 = stakingPool.getUserPoints(user1);

        // 시간이 2배이므로 포인트도 거의 2배
        assertGt(points2, points1);
        assertGt(points2, points1 * 19 / 10); // 1.9배 이상
    }

    // NOTE: test_PointsTimeUnitAffectsCalculation removed - no longer uses timeUnit

    function test_PointsResetOnWithdraw() public {
        stakeFor(user1, 10 ether);

        vm.warp(block.timestamp + 50);
        updatePointsFor(user1);

        uint pointsBefore = stakingPool.getUserPoints(user1);
        assertGt(pointsBefore, 0);

        withdrawFor(user1);

        (, uint pointsAfter,) = stakingPool.getStakePosition(user1);
        assertEq(pointsAfter, 0);
    }

    function test_PointsProportionalToStake() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 5 ether);

        vm.warp(block.timestamp + 100);
        updatePointsFor(user1);
        updatePointsFor(user2);

        uint points1 = stakingPool.getUserPoints(user1);
        uint points2 = stakingPool.getUserPoints(user2);

        // User1이 2배 스테이킹했으므로 포인트도 약 2배
        assertGt(points1, points2);
        assertGt(points1, points2 * 19 / 10); // 1.9배 이상
    }

    function test_NoPointsWithoutStake() public {
        vm.warp(block.timestamp + 100);

        uint points = stakingPool.getUserPoints(user1);
        assertEq(points, 0);
    }

    function test_SnapshotPlusAdditionalPoints() public {
        // 1. 초기 스테이킹
        stakeFor(user1, 100 ether);

        // 2. 50블록 경과
        vm.warp(block.timestamp + 50);

        uint pointsBeforeAdditional = stakingPool.getUserPoints(user1);

        // 3. 추가 스테이킹 (이때 이전 포인트가 스냅샷됨)
        vm.prank(user1);
        router.stake{value: 50 ether}(PROJECT_ID);

        // 4. 추가 50블록 경과
        vm.warp(block.timestamp + 50);

        // 5. 포인트 조회 (스냅샷 + 추가 기간 포인트)
        uint totalPoints = stakingPool.getUserPoints(user1);

        // 검증: 추가 예치 후 포인트는 이전보다 많아야 함
        assertGt(totalPoints, pointsBeforeAdditional, "Points should increase after additional stake");

        // 검증: 대략적인 비율 확인 (100 CROSS × 50블록 + 150 CROSS × 50블록)
        // 비율: pointsBeforeAdditional : additionalPoints = 100:150 = 2:3
        uint additionalPoints = totalPoints - pointsBeforeAdditional;
        assertGt(
            additionalPoints, pointsBeforeAdditional, "Additional period points should be higher due to increased stake"
        );
    }

    function test_ExpectedSeasonPointsWithSnapshot() public {
        // 1. 초기 스테이킹
        stakeFor(user1, 100 ether);

        // 2. 50블록 경과
        vm.warp(block.timestamp + 50);

        uint pointsBeforeAdditional = stakingPool.getUserPoints(user1);

        // 3. 추가 스테이킹 (스냅샷 생성)
        vm.prank(user1);
        router.stake{value: 50 ether}(PROJECT_ID);

        // 4. 시즌 종료까지 대기
        rolloverSeason();

        // 5. 과거 시즌 (시즌 1) 포인트 조회 (스냅샷 + 추가 기간 포함되어야 함)
        (uint season1Points,) = stakingPool.getSeasonUserPoints(1, user1);

        // 검증: 시즌 1의 포인트는 0보다 커야 함
        assertGt(season1Points, 0, "Season 1 points should be greater than 0");

        // 검증: 시즌 1 포인트는 최소한 중간에 기록된 포인트보다는 많아야 함
        assertGt(season1Points, pointsBeforeAdditional, "Season 1 points should include snapshot + additional period");

        // 검증: 스냅샷이 기록되어 있어야 함
        (uint snapshotPoints,) = stakingPool.getSeasonUserPoints(1, user1);
        assertGt(snapshotPoints, 0, "Snapshot points should be recorded");
    }

    function test_TotalPointsDecreaseOnUnstake() public {
        // 두 명의 사용자가 스테이킹
        stakeFor(user1, 10 ether);
        stakeFor(user2, 20 ether);

        // 50초 진행 (시즌 진행 중)
        // SEASON_DURATION = 100, endTime = startTime + 99
        // 따라서 50초 진행하면 시즌 중간
        vm.warp(block.timestamp + 50);

        // 현재 시즌의 총 포인트 계산
        uint totalPointsBefore = stakingPool.seasonTotalPointsSnapshot(1);
        assertGt(totalPointsBefore, 0, "Total points should be greater than 0");

        // user1의 현재 포인트 확인
        uint user1PointsBefore = stakingPool.getUserPoints(user1);
        assertGt(user1PointsBefore, 0, "User1 points should be greater than 0");

        // user1이 unstake (시즌 진행 중이므로 포인트 몰수)
        withdrawFor(user1);

        // unstake 후 user1의 포인트는 0이어야 함
        uint user1PointsAfter = stakingPool.getUserPoints(user1);
        assertEq(user1PointsAfter, 0, "User1 points should be 0 after unstake");

        // unstake 후 총 포인트는 user1의 포인트만큼 감소해야 함 (시즌 진행 중 unstake)
        uint totalPointsAfter = stakingPool.seasonTotalPointsSnapshot(1);
        assertLt(totalPointsAfter, totalPointsBefore, "Total points should decrease after unstake during season");
    }

    function test_MultipleUnstakesTotalPointsDecrease() public {
        // 세 명의 사용자가 스테이킹
        stakeFor(user1, 10 ether);
        stakeFor(user2, 20 ether);
        stakeFor(user3, 30 ether);

        // 50 블록 진행 (시즌 중간)
        vm.warp(block.timestamp + 50);

        // 초기 총 포인트 기록
        uint initialTotalPoints = stakingPool.seasonTotalPointsSnapshot(1);
        assertGt(initialTotalPoints, 0, "Initial total points should be greater than 0");

        // user1 unstake (블록 진행 없음)
        withdrawFor(user1);

        // user1 unstake 직후 총 포인트 (몰수 반영)
        uint totalAfterUser1 = stakingPool.seasonTotalPointsSnapshot(1);
        assertLt(totalAfterUser1, initialTotalPoints, "Total should decrease after user1 unstake");

        // user1의 포인트는 0이 되어야 함
        assertEq(stakingPool.getUserPoints(user1), 0, "User1 should have 0 points");

        // user2 unstake (블록 진행 없음 - 연속 unstake 시나리오)
        withdrawFor(user2);

        // user2 unstake 직후 총 포인트
        uint totalAfterUser2 = stakingPool.seasonTotalPointsSnapshot(1);
        assertLt(totalAfterUser2, totalAfterUser1, "Total should decrease after user2 unstake");

        // user2의 포인트도 0이 되어야 함
        assertEq(stakingPool.getUserPoints(user2), 0, "User2 should have 0 points");

        // user3만 남아있어야 함
        uint user3PointsAfter = stakingPool.getUserPoints(user3);
        assertGt(user3PointsAfter, 0, "User3 should still have points");

        // 최종 총 포인트는 대략 user3의 포인트와 비슷해야 함
        // (약간의 오차 허용 - withdrawAll이 각각 1블록씩 소비하므로 user3 포인트 약간 증가)
        uint finalTotalPoints = stakingPool.seasonTotalPointsSnapshot(1);
        assertGt(finalTotalPoints, 0, "Final total points should still be greater than 0");
        assertApproxEqAbs(
            finalTotalPoints, user3PointsAfter, 60 ether, "Final total should be approximately user3's points"
        );
    }
}
