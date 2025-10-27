// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./BaseTest.sol";

/**
 * @title RewardsTest
 * @notice 보상 시스템 테스트
 */
contract RewardsTest is BaseTest {
    function test_FundSeason() public {
        fundSeason(1, 100 ether);

        uint seasonReward = rewardPool.seasonRewards(1, address(rewardToken));
        assertEq(seasonReward, 100 ether);
    }

    function test_ProportionalRewardDistribution() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 5 ether);

        rolloverSeason();
        fundSeason(1, 150 ether);

        claimSeasonFor(user1, 1);
        claimSeasonFor(user2, 1);

        uint reward1 = rewardToken.balanceOf(user1);
        uint reward2 = rewardToken.balanceOf(user2);

        assertGt(reward1, 0);
        assertGt(reward2, 0);
        assertGt(reward1, reward2);

        // User1이 2배 스테이킹했으므로 2배 보상
        assertGt(reward1, reward2 * 19 / 10); // 1.9배 이상
    }

    function test_EqualStakeEqualReward() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 10 ether);

        rolloverSeason();
        fundSeason(1, 100 ether);

        claimSeasonFor(user1, 1);
        claimSeasonFor(user2, 1);

        uint reward1 = rewardToken.balanceOf(user1);
        uint reward2 = rewardToken.balanceOf(user2);

        // 동일한 스테이킹이므로 비슷한 보상 (1% 오차 허용)
        uint diff = reward1 > reward2 ? reward1 - reward2 : reward2 - reward1;
        assertLt(diff, reward1 / 100);
    }

    function test_MultipleSeasonRewards() public {
        stakeFor(user1, 10 ether);

        // Season 1
        rolloverSeason();
        fundSeason(1, 100 ether);
        claimSeasonFor(user1, 1);

        uint reward1 = rewardToken.balanceOf(user1);

        // Season 2
        rolloverSeason();
        fundSeason(2, 100 ether);
        claimSeasonFor(user1, 2);

        uint totalReward = rewardToken.balanceOf(user1);

        assertGt(reward1, 0);
        assertGt(totalReward, reward1);
    }

    function test_RevertWhen_ClaimBeforeSeasonEnd() public {
        stakeFor(user1, 10 ether);

        vm.prank(user1);
        vm.expectRevert();
        stakingPool.claimSeason(1, address(rewardToken));
    }

    function test_RevertWhen_DuplicateClaim() public {
        stakeFor(user1, 10 ether);

        rolloverSeason();
        fundSeason(1, 100 ether);

        claimSeasonFor(user1, 1);

        vm.prank(user1);
        vm.expectRevert();
        stakingPool.claimSeason(1, address(rewardToken));
    }

    function test_NoRewardForNoStake() public {
        stakeFor(user1, 10 ether);

        rolloverSeason();
        fundSeason(1, 100 ether);

        claimSeasonFor(user1, 1);

        // User2는 스테이킹 안 함
        uint reward2 = rewardToken.balanceOf(user2);
        assertEq(reward2, 0);
    }

    function test_RemainingRewards() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 5 ether);

        rolloverSeason();
        fundSeason(1, 150 ether);

        // User1만 청구
        claimSeasonFor(user1, 1);

        uint claimed = rewardToken.balanceOf(user1);
        uint remaining = rewardPool.getRemainingRewards(1, address(rewardToken));

        assertGt(remaining, 0);
        assertEq(claimed + remaining, 150 ether);
    }

    function test_RewardPool_InsufficientBalance() public {
        // 시즌 1에 100 이더 예치
        fundSeason(1, 100 ether);

        // 사용자 1이 스테이킹
        stakeFor(user1, 50 ether);

        // 시즌 종료
        rolloverSeason();

        // RewardPool에서 토큰 회수 (잔액 부족 상황 생성)
        vm.prank(address(protocol));
        rewardPool.sweep(address(rewardToken), owner, 50 ether);

        // 청구 시도 - 잔액 부족으로 revert
        vm.expectRevert();
        claimSeasonFor(user1, 1);
    }

    function test_RewardPool_ClaimRetryAfterRefund() public {
        // 시즌 1에 100 이더 예치
        fundSeason(1, 100 ether);

        // 사용자 1이 스테이킹
        stakeFor(user1, 50 ether);

        // 시즌 종료
        rolloverSeason();

        // RewardPool에서 토큰 회수 (잔액 부족 상황 생성)
        vm.prank(address(protocol));
        rewardPool.sweep(address(rewardToken), owner, 60 ether);

        // 첫 번째 청구 시도 - 잔액 부족으로 revert
        vm.expectRevert();
        claimSeasonFor(user1, 1);

        // 토큰 재예치
        vm.startPrank(rewardProvider);
        rewardToken.approve(address(rewardPool), 100 ether);
        rewardToken.transfer(address(rewardPool), 100 ether);
        vm.stopPrank();

        // 재시도 - 성공
        claimSeasonFor(user1, 1);

        // 청구 완료 확인
        assertTrue(rewardPool.hasClaimedSeasonReward(user1, 1, address(rewardToken)));
    }

    function test_RewardPool_SeasonFundedEvent() public {
        vm.startPrank(rewardProvider);
        rewardToken.approve(address(rewardPool), 1000 ether);

        // SeasonFunded 이벤트 검증
        vm.expectEmit(true, true, false, true);
        emit IRewardPool.SeasonFunded(1, address(rewardToken), 1000 ether, 1000 ether);

        rewardPool.fundSeason(1, address(rewardToken), 1000 ether);
        vm.stopPrank();
    }

    function test_RewardPool_TokensSweptEvent() public {
        // 토큰 예치
        fundSeason(1, 100 ether);

        uint balanceBefore = rewardToken.balanceOf(address(rewardPool));

        // TokensSwept 이벤트 검증
        vm.expectEmit(true, true, false, true);
        emit RewardPool.TokensSwept(address(rewardToken), owner, 50 ether, balanceBefore, balanceBefore - 50 ether);

        // Sweep
        vm.prank(address(protocol));
        rewardPool.sweep(address(rewardToken), owner, 50 ether);
    }

    function test_PreDepositBeforeSeasonStart() public {
        // 새 프로젝트 생성 (미래 시작 블록, 사전 예치 즉시 가능)
        uint futureStartBlock = block.number + 100;

        vm.prank(owner);
        (, address newStakingPoolAddr, address newRewardPoolAddr) =
            protocol.createProject("PreDepositTest", SEASON_BLOCKS, futureStartBlock, 0, owner, 0);

        StakingPool newPool = StakingPool(newStakingPoolAddr);
        RewardPool newRewardPool = RewardPool(newRewardPoolAddr);

        // 시즌이 아직 시작되지 않음
        assertEq(newPool.currentSeason(), 0, "Season should not have started");

        // 사전 예치 토큰 준비
        deal(address(rewardToken), owner, 1000 ether);
        vm.startPrank(owner);
        rewardToken.approve(newRewardPoolAddr, 1000 ether);

        // 사전 예치 (시즌 1에 예치됨)
        newRewardPool.depositReward(address(rewardToken), 100 ether);

        // 시즌 1에 토큰이 예치되었는지 확인
        address[] memory tokens = newRewardPool.getSeasonRewardTokens(1);
        assertEq(tokens.length, 1, "Should have 1 token");
        assertEq(tokens[0], address(rewardToken), "Token should be reward token");

        (uint total, uint claimed, uint remaining) = newRewardPool.getSeasonTokenInfo(1, address(rewardToken));
        assertEq(total, 100 ether, "Total should be 100 ether");
        assertEq(claimed, 0, "Claimed should be 0");
        assertEq(remaining, 100 ether, "Remaining should be 100 ether");

        vm.stopPrank();
    }

    function test_PreDepositWithBlockRestriction() public {
        // 새 프로젝트 생성 (미래 시작 블록, 사전 예치 미래 블록부터 가능)
        uint futureStartBlock = block.number + 100;
        uint preDepositBlock = block.number + 50;

        vm.prank(owner);
        (,, address newRewardPoolAddr) =
            protocol.createProject("PreDepositRestricted", SEASON_BLOCKS, futureStartBlock, 0, owner, preDepositBlock);

        RewardPool newRewardPool = RewardPool(newRewardPoolAddr);

        // 사전 예치 토큰 준비
        deal(address(rewardToken), owner, 1000 ether);
        vm.startPrank(owner);
        rewardToken.approve(newRewardPoolAddr, 1000 ether);

        // preDepositBlock 이전에 예치 시도 (실패해야 함)
        vm.expectRevert("Pre-deposit not yet available");
        newRewardPool.depositReward(address(rewardToken), 100 ether);

        // preDepositBlock으로 이동
        vm.roll(preDepositBlock);

        // 이제 예치 가능
        newRewardPool.depositReward(address(rewardToken), 100 ether);

        // 시즌 1에 토큰이 예치되었는지 확인
        (uint total,,) = newRewardPool.getSeasonTokenInfo(1, address(rewardToken));
        assertEq(total, 100 ether, "Total should be 100 ether");

        vm.stopPrank();
    }

    function test_PreDepositAndNormalDeposit() public {
        // 새 프로젝트 생성 (미래 시작 블록)
        uint futureStartBlock = block.number + 100;

        vm.prank(owner);
        (uint newProjectID, address newStakingPoolAddr, address newRewardPoolAddr) =
            protocol.createProject("PreDepositAndNormal", SEASON_BLOCKS, futureStartBlock, 0, owner, 0);

        StakingPool newPool = StakingPool(newStakingPoolAddr);
        RewardPool newRewardPool = RewardPool(newRewardPoolAddr);

        // 사전 예치 토큰 준비
        deal(address(rewardToken), owner, 1000 ether);
        vm.startPrank(owner);
        rewardToken.approve(newRewardPoolAddr, 1000 ether);

        // 사전 예치
        newRewardPool.depositReward(address(rewardToken), 100 ether);

        // 시즌 시작 블록으로 이동
        vm.roll(futureStartBlock);

        // 시즌을 시작하기 위해 스테이킹 액션 필요
        // 먼저 StakingRouter 생성 및 승인
        vm.stopPrank();

        StakingRouter newRouter = new StakingRouter(address(wcross), address(protocol));
        vm.prank(owner);
        protocol.setApprovedRouter(newProjectID, address(newRouter), true);

        // 사용자가 스테이킹하여 시즌 시작
        deal(user1, 10 ether);
        vm.prank(user1);
        newRouter.stake{value: 10 ether}(newProjectID);

        // 현재 시즌 확인
        assertEq(newPool.currentSeason(), 1, "Season should have started");

        // 시즌 시작 후 추가 예치
        vm.startPrank(owner);
        newRewardPool.depositReward(address(rewardToken), 50 ether);

        // 총 예치량 확인
        (uint total,,) = newRewardPool.getSeasonTokenInfo(1, address(rewardToken));
        assertEq(total, 150 ether, "Total should be 150 ether (100 + 50)");

        vm.stopPrank();
    }
}
