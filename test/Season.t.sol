// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./BaseTest.sol";

/**
 * @title SeasonTest
 * @notice 시즌 시스템 테스트
 */
contract SeasonTest is BaseTest {
    function test_SeasonRollover() public {
        assertEq(stakingPool.currentSeason(), 1);

        rolloverSeason();

        assertEq(stakingPool.currentSeason(), 2);
    }

    function test_PointsResetAfterRollover() public {
        stakeFor(user1, 10 ether);

        vm.roll(block.number + 50);
        updatePointsFor(user1);

        uint pointsBefore = stakingPool.getUserPoints(user1);
        assertGt(pointsBefore, 0);

        // 이전 시즌 ID 저장
        uint season1 = stakingPool.currentSeason();

        rolloverSeason();

        // 이전 시즌의 totalPoints는 롤오버 시 저장됨
        uint season1Snapshot = stakingPool.seasonTotalPointsSnapshot(season1);
        assertGt(season1Snapshot, 0);

        // 새 시즌 ID 확인
        uint season2 = stakingPool.currentSeason();
        assertEq(season2, season1 + 1);

        // 새 시즌에서의 현재 포인트 (자동 참여로 누적 시작)
        uint pointsInNewSeason = stakingPool.getUserPoints(user1);
        // 새 시즌에서는 다시 0부터 시작 (자동 참여 개념이므로 블록에 따라 누적)
        // 롤오버 직후이므로 이전 시즌보다 적거나 같아야 함
        assertLe(pointsInNewSeason, pointsBefore);
    }

    function test_TotalPointsSnapshot() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 5 ether);

        vm.roll(block.number + 50);
        updatePointsFor(user1);
        updatePointsFor(user2);

        rolloverSeason();

        uint totalSnapshot = stakingPool.seasonTotalPointsSnapshot(1);
        assertGt(totalSnapshot, 0);
    }

    function test_MultipleSeasons() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 5 ether); // 두 명의 유저로 total points 생성

        // Season 1
        assertEq(stakingPool.currentSeason(), 1);

        vm.roll(block.number + 50);
        updatePointsFor(user1);
        updatePointsFor(user2);

        uint season1User1Points = stakingPool.getUserPoints(user1);
        uint season1User2Points = stakingPool.getUserPoints(user2);
        assertGt(season1User1Points, 0); // 시즌 1에서 user1 포인트 확인
        assertGt(season1User2Points, 0); // 시즌 1에서 user2 포인트 확인

        rolloverSeason();
        assertEq(stakingPool.currentSeason(), 2);

        uint snapshot1 = stakingPool.seasonTotalPointsSnapshot(1);

        // Season 2
        vm.roll(block.number + 50);
        updatePointsFor(user1);
        updatePointsFor(user2);

        uint season2User1Points = stakingPool.getUserPoints(user1);
        uint season2User2Points = stakingPool.getUserPoints(user2);
        assertGt(season2User1Points, 0); // 시즌 2에서 user1 포인트 확인
        assertGt(season2User2Points, 0); // 시즌 2에서 user2 포인트 확인

        rolloverSeason();
        assertEq(stakingPool.currentSeason(), 3);

        uint snapshot2 = stakingPool.seasonTotalPointsSnapshot(2);

        assertGt(snapshot1, 0);
        assertGt(snapshot2, 0);
    }

    function test_RevertWhen_RolloverBeforeSeasonEnd() public {
        vm.expectRevert();
        stakingPool.rolloverSeason();
    }

    function test_SeasonInfo() public view {
        (uint season, uint startBlock, uint endBlock, uint blocksElapsed) = stakingPool.getCurrentSeasonInfo();

        assertEq(season, 1);
        assertEq(startBlock, block.number);
        uint seasonBlocks = stakingPool.seasonBlocks();
        assertEq(endBlock, block.number + seasonBlocks);
        assertEq(blocksElapsed, 0);
    }

    function test_StakePreservedAcrossSeasons() public {
        stakeFor(user1, 10 ether);

        rolloverSeason();

        (uint balance,,) = stakingPool.getStakePosition(user1);
        assertEq(balance, 10 ether);
    }
}
