// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./BaseTest.sol";

/**
 * @title PointsTest
 * @notice 포인트 계산 및 관리 테스트
 */
contract PointsTest is BaseTest {
    function test_PointsAccumulation() public {
        stakeFor(user1, 10 ether);

        vm.roll(block.number + 50);
        updatePointsFor(user1);

        uint points = stakingPool.getUserPoints(user1);
        assertGt(points, 0);
    }

    function test_PointsCalculation() public {
        stakeFor(user1, 10 ether);

        vm.roll(block.number + 100);
        updatePointsFor(user1);

        uint points1 = stakingPool.getUserPoints(user1);

        vm.roll(block.number + 100);
        updatePointsFor(user1);

        uint points2 = stakingPool.getUserPoints(user1);

        // 시간이 2배이므로 포인트도 거의 2배
        assertGt(points2, points1);
        assertGt(points2, points1 * 19 / 10); // 1.9배 이상
    }

    function test_BlockTimeAffectsPoints() public {
        stakeFor(user1, 10 ether);

        // 블록 시간 1초로 설정 (protocol을 통해)
        vm.prank(address(protocol));
        stakingPool.setBlockTime(1);

        vm.roll(block.number + 100);
        updatePointsFor(user1);

        uint points1 = stakingPool.getUserPoints(user1);

        // 새로운 사용자 with 블록 시간 12초
        stakeFor(user2, 10 ether);

        vm.prank(address(protocol));
        stakingPool.setBlockTime(12);

        vm.roll(block.number + 100);
        updatePointsFor(user2);

        uint points2 = stakingPool.getUserPoints(user2);

        // User2가 12배 더 많은 포인트
        assertGt(points2, points1 * 10); // 최소 10배
    }

    function test_PointsTimeUnitAffectsCalculation() public {
        stakeFor(user1, 10 ether);

        // 1시간 단위로 설정 (protocol을 통해)
        vm.prank(address(protocol));
        stakingPool.setPointsTimeUnit(1 hours);

        // blockTime(1초) × pointsTimeUnit(1시간) = 3600블록 필요
        uint pointsTimeUnit = stakingPool.pointsTimeUnit();
        uint blockTime = stakingPool.blockTime();
        uint blocksNeeded = pointsTimeUnit / blockTime; // 3600 / 1 = 3600블록

        vm.roll(block.number + blocksNeeded);
        updatePointsFor(user1);

        uint points = stakingPool.getUserPoints(user1);
        assertGt(points, 0);
    }

    function test_PointsResetOnWithdraw() public {
        stakeFor(user1, 10 ether);

        vm.roll(block.number + 50);
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

        vm.roll(block.number + 100);
        updatePointsFor(user1);
        updatePointsFor(user2);

        uint points1 = stakingPool.getUserPoints(user1);
        uint points2 = stakingPool.getUserPoints(user2);

        // User1이 2배 스테이킹했으므로 포인트도 약 2배
        assertGt(points1, points2);
        assertGt(points1, points2 * 19 / 10); // 1.9배 이상
    }

    function test_NoPointsWithoutStake() public {
        vm.roll(block.number + 100);

        uint points = stakingPool.getUserPoints(user1);
        assertEq(points, 0);
    }

    function test_SnapshotPlusAdditionalPoints() public {
        // 1. 초기 스테이킹
        stakeFor(user1, 100 ether);

        // 2. 50블록 경과
        vm.roll(block.number + 50);

        uint pointsBeforeAdditional = stakingPool.getUserPoints(user1);

        // 3. 추가 스테이킹 (이때 이전 포인트가 스냅샷됨)
        vm.prank(user1);
        router.stake{value: 50 ether}(PROJECT_ID);

        // 4. 추가 50블록 경과
        vm.roll(block.number + 50);

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
        vm.roll(block.number + 50);

        uint pointsBeforeAdditional = stakingPool.getUserPoints(user1);

        // 3. 추가 스테이킹 (스냅샷 생성)
        vm.prank(user1);
        router.stake{value: 50 ether}(PROJECT_ID);

        // 4. 시즌 종료까지 대기
        rolloverSeason();

        // 5. 과거 시즌 (시즌 1) 포인트 조회 (스냅샷 + 추가 기간 포함되어야 함)
        uint season1Points = stakingPool.getExpectedSeasonPoints(1, user1);

        // 검증: 시즌 1의 포인트는 0보다 커야 함
        assertGt(season1Points, 0, "Season 1 points should be greater than 0");

        // 검증: 시즌 1 포인트는 최소한 중간에 기록된 포인트보다는 많아야 함
        assertGt(season1Points, pointsBeforeAdditional, "Season 1 points should include snapshot + additional period");

        // 검증: 스냅샷이 기록되어 있어야 함
        uint snapshotPoints = stakingPool.getSeasonUserPoints(1, user1);
        assertGt(snapshotPoints, 0, "Snapshot points should be recorded");
    }
}
