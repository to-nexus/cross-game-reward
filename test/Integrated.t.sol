// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./BaseTest.sol";

/**
 * @title IntegratedTest
 * @notice 전체 플로우 통합 테스트
 */
contract IntegratedTest is BaseTest {
    function test_CompleteStakingFlow() public {
        // 1. User1 스테이킹
        stakeFor(user1, 10 ether);
        (uint balance1,,) = stakingPool.getStakePosition(user1);
        assertEq(balance1, 10 ether);

        // 2. 포인트 누적
        vm.roll(block.number + 50);
        updatePointsFor(user1);
        uint points1 = stakingPool.getUserPoints(user1);
        assertGt(points1, 0);

        // 3. User2 스테이킹
        stakeFor(user2, 5 ether);

        // 4. 시즌 종료
        rolloverSeason();
        assertEq(stakingPool.currentSeason(), 2);

        // 5. 보상 예치
        fundSeason(1, 150 ether);

        // 6. 보상 청구
        claimSeasonFor(user1, 1);
        claimSeasonFor(user2, 1);

        // 7. 보상 검증
        uint reward1 = rewardToken.balanceOf(user1);
        uint reward2 = rewardToken.balanceOf(user2);

        assertGt(reward1, 0);
        assertGt(reward2, 0);
        assertGt(reward1, reward2);
    }

    function test_MultiSeasonFlow() public {
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

        assertEq(stakingPool.currentSeason(), 3);
        assertGt(totalReward, reward1);
        assertEq(totalReward, 200 ether);
    }

    function test_StakeWithdrawRestake() public {
        // 첫 번째 스테이킹
        stakeFor(user1, 10 ether);

        vm.roll(block.number + 50);
        updatePointsFor(user1);

        // 출금
        withdrawFor(user1);
        (uint balance1,,) = stakingPool.getStakePosition(user1);
        assertEq(balance1, 0);

        // 재스테이킹
        stakeFor(user1, 5 ether);
        (uint balance2,,) = stakingPool.getStakePosition(user1);
        assertEq(balance2, 5 ether);
    }

    function test_MultiUserMultiSeasonFlow() public {
        // Season 1
        stakeFor(user1, 10 ether);
        stakeFor(user2, 5 ether);

        rolloverSeason();
        fundSeason(1, 150 ether);

        claimSeasonFor(user1, 1);
        claimSeasonFor(user2, 1);

        uint s1_reward1 = rewardToken.balanceOf(user1);
        uint s1_reward2 = rewardToken.balanceOf(user2);

        // Season 2
        stakeFor(user3, 3 ether);

        rolloverSeason();
        fundSeason(2, 180 ether);

        claimSeasonFor(user1, 2);
        claimSeasonFor(user2, 2);
        claimSeasonFor(user3, 2);

        uint s2_total1 = rewardToken.balanceOf(user1);
        uint s2_total2 = rewardToken.balanceOf(user2);
        uint s2_reward3 = rewardToken.balanceOf(user3);

        assertGt(s2_total1, s1_reward1);
        assertGt(s2_total2, s1_reward2);
        assertGt(s2_reward3, 0);
    }

    function test_EdgeCase_ClaimWithoutPoints() public {
        stakeFor(user1, 10 ether);

        // 즉시 롤오버 (포인트 누적 없음)
        rolloverSeason();
        fundSeason(1, 100 ether);

        // 포인트가 거의 없어도 청구 가능
        claimSeasonFor(user1, 1);

        uint reward = rewardToken.balanceOf(user1);
        // 최소한의 포인트는 있을 것
        assertGt(reward, 0);
    }

    function test_FullNativeTokenFlow() public {
        uint initialBalance = user1.balance;

        // Native CROSS로 스테이킹
        vm.prank(user1);
        router.stake{value: 10 ether}(PROJECT_ID);

        assertEq(user1.balance, initialBalance - 10 ether);

        // WCROSS로 변환됨
        (uint staked,,) = stakingPool.getStakePosition(user1);
        assertEq(staked, 10 ether);

        // router.unstake()가 내부에서 withdrawAllFor + unwrap을 처리
        vm.prank(user1);
        router.unstake(PROJECT_ID);

        // Native CROSS 복구 (가스 제외)
        assertGt(user1.balance, initialBalance - 0.01 ether);
    }

    // ============ Protocol 테스트 (Code 패턴) ============

    function test_ProtocolCreateProject() public {
        // 새 프로젝트 생성 테스트
        vm.prank(owner);
        (uint projectId, address stakingPoolAddr, address rewardPoolAddr) =
            protocol.createProject("NewProject", SEASON_BLOCKS * 2, block.number + 1, 0, address(0));

        assertEq(projectId, 2); // 첫 프로젝트는 setUp에서 생성됨
        assertNotEq(stakingPoolAddr, address(0));
        assertNotEq(rewardPoolAddr, address(0));
    }

    function test_ProtocolCompleteFlow() public {
        // Protocol을 통해 새 프로젝트 생성하고 전체 플로우 테스트
        vm.prank(owner);
        (uint projectId, address stakingPoolAddr,) =
            protocol.createProject("NewProject", SEASON_BLOCKS, block.number, 0, address(0));

        StakingPool newPool = StakingPool(stakingPoolAddr);

        // Router 생성 및 승인
        vm.prank(owner);
        StakingRouter newRouter = new StakingRouter(address(wcross), address(protocol));

        vm.prank(owner);
        protocol.setApprovedRouter(projectId, address(newRouter), true);

        // 스테이킹
        vm.prank(user1);
        newRouter.stake{value: 10 ether}(projectId);

        (uint balance,,) = newPool.getStakePosition(user1);
        assertEq(balance, 10 ether);
    }
}
