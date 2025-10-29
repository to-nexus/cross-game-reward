// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

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

        vm.warp(block.timestamp + 50);
        updatePointsFor(user1);

        uint pointsBefore = stakingPool.getUserPoints(user1);
        assertGt(pointsBefore, 0);

        // 이전 시즌 ID 저장
        uint season1 = stakingPool.currentSeason();

        // 시즌 종료로 이동 후 롤오버 (헬퍼 함수 사용)
        rolloverSeason();

        // 이전 시즌의 totalPoints는 롤오버 시 저장됨
        uint season1Snapshot = stakingPool.seasonTotalPointsSnapshot(season1);
        assertGt(season1Snapshot, 0);

        // 새 시즌 ID 확인
        uint season2 = stakingPool.currentSeason();
        assertEq(season2, season1 + 1);

        // 새 시즌에서의 포인트는 자동 참여로 계속 누적됨
        // (rolloverSeason 헬퍼가 블록을 이동시키므로 포인트가 계속 누적될 수 있음)
        uint pointsInNewSeason = stakingPool.getUserPoints(user1);
        assertGt(pointsInNewSeason, 0);
    }

    function test_GetCurrentSeasonInfo_NoStake() public {
        // Scenario: Season 1 started, user staked, then no activity for 2+ seasons

        // Initial stake in season 1
        stakeFor(user1, 10 ether);
        assertEq(stakingPool.currentSeason(), 1);

        // Check season info at season 1
        (uint season, uint startTime, uint endTime, uint timeElapsed) = stakingPool.getCurrentSeasonInfo();
        uint season1Start = startTime;
        uint season1End = endTime;
        assertEq(season, 1, "Should be season 1");

        console.log("Season 1 start:", season1Start);
        console.log("Season 1 end:", season1End);
        console.log("SEASON_DURATION:", SEASON_DURATION);

        // Move to season 3 time without any transactions
        // Season 1 ends at season1End (e.g., 100)
        // Season 2: 101~200
        // Season 3: 201~300
        // We want to be in the middle of season 3, so 250
        uint targetBlock = season1End + SEASON_DURATION + 50; // Season 2 (100 blocks) + 50 into season 3
        vm.warp(targetBlock);

        console.log("Current block after skip:", block.timestamp);
        console.log("Blocks since season 1 start:", block.timestamp - season1Start);

        // Now getCurrentSeasonInfo should return season 3, not season 1!
        (season, startTime, endTime, timeElapsed) = stakingPool.getCurrentSeasonInfo();

        console.log("Returned season:", season);
        console.log("Returned startTime:", startTime);
        console.log("Returned endTime:", endTime);
        console.log("Blocks elapsed:", timeElapsed);

        assertEq(season, 3, "Should calculate season 3 based on current block");
        assertEq(timeElapsed, 49, "Should be 49 blocks into season 3 (0-indexed)");
        assertLt(timeElapsed, SEASON_DURATION, "Blocks elapsed should be less than season length");

        // Verify storage season is still 1 (not rolled over)
        assertEq(stakingPool.currentSeason(), 1, "Storage should still be season 1");
    }

    function test_TotalPointsSnapshot() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 5 ether);

        vm.warp(block.timestamp + 50);
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

        vm.warp(block.timestamp + 50);
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
        vm.warp(block.timestamp + 50);
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
        (uint season, uint startTime, uint endTime, uint timeElapsed) = stakingPool.getCurrentSeasonInfo();

        assertEq(season, 1);
        assertEq(startTime, block.timestamp);
        uint seasonDuration = stakingPool.seasonDuration();
        // endTime은 inclusive이므로 startTime + seasonDuration - 1
        assertEq(endTime, block.timestamp + seasonDuration - 1);
        assertEq(timeElapsed, 0);
    }

    function test_StakePreservedAcrossSeasons() public {
        stakeFor(user1, 10 ether);

        rolloverSeason();

        (uint balance,,) = stakingPool.getStakePosition(user1);
        assertEq(balance, 10 ether);
    }

    function test_Season_ManualRollover_WithPendingCheck() public {
        // 시즌 1 시작
        stakeFor(user1, 50 ether);

        // 시즌 1의 종료 블록 확인
        (,, uint season1EndBlock,) = stakingPool.getCurrentSeasonInfo();

        // 5개 시즌이 완전히 지나가도록 블록 증가
        // endTime = startTime + seasonDuration - 1이므로
        // 시즌 2~6이 완료되려면 season1EndBlock + SEASON_DURATION * 5까지 가야 함
        vm.warp(season1EndBlock + SEASON_DURATION * 5);

        // 대기 중인 시즌 개수 확인
        uint pending = protocol.getPendingSeasonRollovers(PROJECT_ID);
        assertEq(pending, 5);

        // 수동 롤오버
        vm.startPrank(owner);
        uint rolled = protocol.manualRolloverSeasons(PROJECT_ID, 50);
        vm.stopPrank();

        assertEq(rolled, 5);
        assertEq(stakingPool.currentSeason(), 6);

        // 더 이상 대기 중인 시즌 없음
        pending = protocol.getPendingSeasonRollovers(PROJECT_ID);
        assertEq(pending, 0);
    }

    function test_Season_ManualRollover_MultipleIterations() public {
        // 시즌 1 시작
        stakeFor(user1, 50 ether);

        // 시즌 1의 종료 블록 확인
        (,, uint season1EndBlock,) = stakingPool.getCurrentSeasonInfo();

        // 120개 시즌이 완전히 지나가도록 블록 증가
        // 시즌 2~121이 완료되려면 season1EndBlock + SEASON_DURATION * 120까지
        vm.warp(season1EndBlock + SEASON_DURATION * 120);

        // Protocol을 통해 수동 롤오버 수행
        vm.startPrank(owner);

        // 첫 번째 호출: 최대 100개 처리
        uint rolled1 = protocol.manualRolloverSeasons(PROJECT_ID, 100);
        assertEq(rolled1, 100);
        assertEq(stakingPool.currentSeason(), 101);

        // 두 번째 호출: 나머지 20개 처리
        uint rolled2 = protocol.manualRolloverSeasons(PROJECT_ID, 100);
        assertEq(rolled2, 20);
        assertEq(stakingPool.currentSeason(), 121);

        // 세 번째 호출: 더 이상 처리할 시즌 없음
        uint rolled3 = protocol.manualRolloverSeasons(PROJECT_ID, 100);
        assertEq(rolled3, 0);

        vm.stopPrank();
    }
}
