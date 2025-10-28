// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./BaseTest.sol";

/**
 * @title AdvancedTest
 * @notice 고급 시나리오 및 엣지 케이스 테스트
 */
contract AdvancedTest is BaseTest {
    /**
     * @notice 풀 재시작 테스트
     */
    function test_PoolRestartAfterEnd() public {
        // 1. 시즌 1 진행
        stakeFor(user1, 10 ether);

        vm.roll(block.number + SEASON_BLOCKS + 1);
        rolloverSeason();

        assertEq(stakingPool.currentSeason(), 2);

        // 2. poolEndBlock 설정하여 시즌 2 종료 예정
        uint endBlock = block.number + 50;
        vm.prank(address(protocol));
        stakingPool.setPoolEndBlock(endBlock);

        // 3. endBlock에 도달 + 1
        vm.roll(endBlock + 1);

        // 4. 시즌이 활성화되지 않음 확인
        assertFalse(stakingPool.isSeasonActive());

        // 5. 풀 재시작 설정 (poolEndBlock을 해제하거나 연장)
        uint restartBlock = block.number + 100;
        vm.startPrank(address(protocol));
        stakingPool.setNextSeasonStart(restartBlock);
        // poolEndBlock을 재시작 이후로 연장하여 재시작 허용
        stakingPool.setPoolEndBlock(restartBlock + SEASON_BLOCKS * 2);
        vm.stopPrank();

        // 6. restartBlock에 도달
        vm.roll(restartBlock);

        // 7. Virtual Season 활성화 확인
        assertTrue(stakingPool.isSeasonActive());

        // 8. 새로운 stake로 시즌 3 시작
        stakeFor(user2, 5 ether);

        // 9. 시즌 3이 생성되었는지 확인
        assertEq(stakingPool.currentSeason(), 3);
    }

    /**
     * @notice Virtual Season에서 Real Season으로 전환 테스트
     */
    function test_VirtualToRealSeasonTransition() public {
        // 새 프로젝트 생성 (미래 시작 블록)
        uint futureStartBlock = block.number + 100;

        vm.prank(owner);
        (, address poolAddr,) = protocol.createProject("FutureProject", SEASON_BLOCKS, futureStartBlock, 0, owner, 0);

        StakingPool newPool = StakingPool(poolAddr);

        // 1. 아직 시작 전 - 시즌 0
        assertEq(newPool.currentSeason(), 0);
        assertFalse(newPool.isSeasonActive());

        // 2. 시작 블록 직전
        vm.roll(futureStartBlock - 1);
        assertFalse(newPool.isSeasonActive());

        // 3. 시작 블록 도달 - Virtual Season 활성화
        vm.roll(futureStartBlock);
        assertTrue(newPool.isSeasonActive());
        assertEq(newPool.currentSeason(), 0); // 아직 구조체는 생성 안됨

        // 4. Virtual Season 정보 조회
        (uint season, uint startBlock,,) = newPool.getCurrentSeasonInfo();
        assertEq(season, 1); // Virtual season 1
        assertEq(startBlock, futureStartBlock);

        // 5. 첫 stake로 Real Season 생성
        vm.deal(user1, 100 ether);
        vm.startPrank(user1);
        wcross.deposit{value: 50 ether}();
        wcross.approve(address(newPool), 10 ether);
        newPool.stake(10 ether);
        vm.stopPrank();

        // 6. Real Season 생성 확인
        assertEq(newPool.currentSeason(), 1);
        assertTrue(newPool.isSeasonActive());
    }

    /**
     * @notice 다중 시즌 자동 참여 테스트
     */
    function test_MultiSeasonAutoParticipation() public {
        // 1. 시즌 1에 stake
        stakeFor(user1, 10 ether);

        // 블록 진행 후 포인트 확인
        vm.roll(block.number + 10);
        uint initialPoints = stakingPool.getUserPoints(user1);
        assertGt(initialPoints, 0);

        // 2. 5개 시즌 연속 롤오버 (행위 없음)
        for (uint i = 0; i < 5; i++) {
            vm.roll(block.number + SEASON_BLOCKS + 1);
            rolloverSeason();
        }

        assertEq(stakingPool.currentSeason(), 6);

        // 3. 각 시즌별 snapshot 확인
        for (uint season = 1; season <= 5; season++) {
            uint snapshot = stakingPool.seasonTotalPointsSnapshot(season);
            assertGt(snapshot, 0, "Season should have points");

            // 각 시즌의 예상 포인트 확인
            (uint userSeasonPoints,) = stakingPool.getSeasonUserPoints(season, user1);
            assertGt(userSeasonPoints, 0, "User should have points in season");
        }

        // 4. 현재 시즌(6)에서도 포인트 누적 중
        vm.roll(block.number + 10);
        uint currentPoints = stakingPool.getUserPoints(user1);
        assertGt(currentPoints, 0);
    }

    /**
     * @notice Admin 긴급 함수 테스트
     */
    function test_AdminEmergencyFunctions() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 5 ether);

        // 1. 긴급 풀 종료 설정 (미래 블록으로)
        uint emergencyEndBlock = block.number + 20;
        vm.prank(address(protocol));
        stakingPool.setPoolEndBlock(emergencyEndBlock);

        assertEq(stakingPool.poolEndBlock(), emergencyEndBlock);

        // 2. 종료 블록 도달 + 1
        vm.roll(emergencyEndBlock + 1);

        // 3. 시즌 비활성화 확인
        assertFalse(stakingPool.isSeasonActive());

        // 4. 새 stake 불가 확인
        vm.expectRevert();
        stakeFor(user3, 10 ether);

        // 5. 기존 stake는 출금 가능
        withdrawFor(user1);
        (uint balance,,) = stakingPool.getStakePosition(user1);
        assertEq(balance, 0);
    }

    /**
     * @notice 시즌 gap 테스트
     */
    function test_SeasonGapBetweenRollovers() public {
        stakeFor(user1, 10 ether);

        // Season 1 정보
        (,, uint season1End,) = stakingPool.getCurrentSeasonInfo();

        // Rollover - NOTE: rolloverSeason() helper moves block forward significantly
        vm.roll(season1End + 1);
        stakingPool.rolloverSeason(); // Don't use helper, just call directly

        // Season 2 정보
        (, uint season2Start,,) = stakingPool.getCurrentSeasonInfo();

        // 1블록 gap 확인
        assertEq(season2Start, season1End + 1);
    }

    /**
     * @notice 시즌 종료 후 claim 테스트
     */
    function test_ClaimAfterSeasonEnd() public {
        stakeFor(user1, 10 ether);

        vm.roll(block.number + 50);
        updatePointsFor(user1);

        uint season1 = stakingPool.currentSeason();

        // Rollover
        rolloverSeason();

        // 다음 시즌에서 이전 시즌 포인트 조회
        (uint expectedPoints,) = stakingPool.getSeasonUserPoints(season1, user1);
        assertGt(expectedPoints, 0);

        // 보상 추가
        fundSeason(season1, 100 ether);

        // Claim
        claimSeasonFor(user1, season1);

        uint reward = rewardToken.balanceOf(user1);
        assertGt(reward, 0);
    }

    /**
     * @notice 0 balance 유저 처리 테스트
     */
    function test_ZeroBalanceUserInRollover() public {
        // user1 stake
        stakeFor(user1, 10 ether);

        // user2 stake 후 즉시 출금
        stakeFor(user2, 5 ether);
        withdrawFor(user2);

        (uint balance,,) = stakingPool.getStakePosition(user2);
        assertEq(balance, 0);

        // Rollover (user2는 balance가 0이지만 stakers 배열에는 존재)
        vm.roll(block.number + SEASON_BLOCKS + 1);
        rolloverSeason();

        // user1만 포인트 있어야 함
        uint snapshot = stakingPool.seasonTotalPointsSnapshot(1);
        assertGt(snapshot, 0);

        (uint user1Points,) = stakingPool.getSeasonUserPoints(1, user1);
        (uint user2Points,) = stakingPool.getSeasonUserPoints(1, user2);

        assertGt(user1Points, 0);
        assertEq(user2Points, 0);
    }

    /**
     * @notice poolEndBlock과 season 경계 테스트
     */
    function test_PoolEndBlockAtSeasonBoundary() public {
        stakeFor(user1, 10 ether);

        // poolEndBlock을 현재 시즌 종료 전으로 설정
        (,, uint endBlock,) = stakingPool.getCurrentSeasonInfo();

        // 시즌 종료 10블록 전으로 설정
        uint earlyEndBlock = endBlock - 10;
        require(earlyEndBlock > block.number, "Test setup error");

        vm.prank(address(protocol));
        stakingPool.setPoolEndBlock(earlyEndBlock);

        // earlyEndBlock 도달 + 1
        vm.roll(earlyEndBlock + 1);

        // 시즌 비활성화
        assertFalse(stakingPool.isSeasonActive());

        // Rollover 불가
        vm.expectRevert();
        stakingPool.rolloverSeason();
    }

    /**
     * @notice 가상 시즌 - poolEndBlock이 시즌 중간에 오는 경우
     */
    function test_VirtualSeason_EarlyPoolEnd() public {
        // 새 프로젝트 생성 (정상적으로 시작)
        vm.startPrank(owner);
        uint futureStart = block.number + 100;

        (, address newPoolAddr,) = protocol.createProject("EarlyEnd", SEASON_BLOCKS, futureStart, 0, owner, 0);

        StakingPool newPool = StakingPool(newPoolAddr);
        vm.stopPrank();

        // 시작 블록 이후 poolEndBlock을 시즌 중간으로 설정
        vm.roll(futureStart);
        assertTrue(newPool.isSeasonActive());

        // poolEndBlock을 현재부터 50블록 후로 설정 (시즌 길이보다 짧음)
        uint earlyEnd = block.number + 50;
        vm.prank(address(protocol));
        newPool.setPoolEndBlock(earlyEnd);

        // earlyEnd + 1 도달 (poolEndBlock은 해당 블록까지 활성)
        vm.roll(earlyEnd + 1);
        assertFalse(newPool.isSeasonActive());

        // 그 이후도 비활성
        vm.roll(earlyEnd + 10);
        assertFalse(newPool.isSeasonActive());
    }

    /**
     * @notice 가상 시즌 - nextSeasonStart가 poolEndBlock 이후인 경우
     */
    function test_VirtualSeason_NextSeasonAfterPoolEnd() public {
        stakeFor(user1, 10 ether);

        // 시즌 1 종료
        vm.roll(block.number + SEASON_BLOCKS + 1);
        rolloverSeason();

        // poolEndBlock 설정
        uint poolEnd = block.number + 50;
        vm.prank(address(protocol));
        stakingPool.setPoolEndBlock(poolEnd);

        // nextSeasonStart를 poolEnd 이후로 설정
        vm.prank(address(protocol));
        stakingPool.setNextSeasonStart(poolEnd + 100);

        // poolEnd + 1 도달 - 비활성
        vm.roll(poolEnd + 1);
        assertFalse(stakingPool.isSeasonActive());

        // nextSeasonStart 도달해도 poolEnd가 지났으므로 비활성
        vm.roll(poolEnd + 100);
        assertFalse(stakingPool.isSeasonActive());
    }

    /**
     * @notice 가상 시즌 - poolEndBlock이 정확히 다음 시즌 시작 블록인 경우
     * @dev poolEndBlock이 nextSeasonStart와 같거나 작으면 가상 시즌은 시작할 수 없음
     */
    function test_VirtualSeason_PoolEndAtNextSeasonStart() public {
        stakeFor(user1, 10 ether);

        // 시즌 1 종료
        (,, uint season1End,) = stakingPool.getCurrentSeasonInfo();
        vm.roll(season1End + 1);
        rolloverSeason();

        // 시즌 2의 endBlock 확인
        (,, uint season2End,) = stakingPool.getCurrentSeasonInfo();

        // poolEndBlock을 다음 시즌 시작 블록으로 설정
        uint nextSeasonStart = season2End + 100;
        vm.prank(address(protocol));
        stakingPool.setNextSeasonStart(nextSeasonStart);

        vm.prank(address(protocol));
        stakingPool.setPoolEndBlock(nextSeasonStart);

        // 시즌 2 종료 이후로 이동
        vm.roll(season2End + 1);
        assertFalse(stakingPool.isSeasonActive());

        // nextSeasonStart - 1 블록: 시즌 2 종료됨, 가상 시즌 미시작
        vm.roll(nextSeasonStart - 1);
        assertFalse(stakingPool.isSeasonActive());

        // nextSeasonStart 블록: poolEndBlock = nextSeasonStart이므로 가상 시즌 시작 불가
        vm.roll(nextSeasonStart);
        assertFalse(stakingPool.isSeasonActive());

        // nextSeasonStart + 1: 비활성 유지
        vm.roll(nextSeasonStart + 1);
        assertFalse(stakingPool.isSeasonActive());
    }
}
