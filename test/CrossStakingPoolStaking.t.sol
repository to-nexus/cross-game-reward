// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/CrossStakingPoolBase.t.sol";

/**
 * @title CrossStakingPoolStakingTest
 * @notice 스테이킹 및 언스테이킹 기본 기능 테스트
 */
contract CrossStakingPoolStakingTest is CrossStakingPoolBase {
    // ==================== 기본 스테이킹 테스트 ====================

    function testStakeBasic() public {
        uint stakeAmount = 10 ether;

        vm.startPrank(user1);
        crossToken.approve(address(pool), stakeAmount);
        pool.stake(stakeAmount);
        vm.stopPrank();

        assertEq(pool.balances(user1), stakeAmount, "User should have staked amount recorded");
        assertEq(crossToken.balanceOf(address(pool)), stakeAmount, "Pool should hold CROSS tokens");
    }

    function testStakeMinimumAmount() public {
        uint stakeAmount = 0.5 ether; // 최소 금액보다 적음

        vm.startPrank(user1);
        crossToken.approve(address(pool), stakeAmount);

        vm.expectRevert(CrossStakingPool.CSPBelowMinimumStakeAmount.selector);
        pool.stake(stakeAmount);
        vm.stopPrank();
    }

    function testStakeMinimumAmountExact() public {
        uint stakeAmount = 1 ether; // 정확히 최소 금액

        vm.startPrank(user1);
        crossToken.approve(address(pool), stakeAmount);
        pool.stake(stakeAmount);
        vm.stopPrank();

        assertEq(pool.balances(user1), stakeAmount, "Should allow exact minimum");
    }

    function testStakeVerySmall() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 1 ether);
        pool.stake(1 ether); // 정확히 최소
        vm.stopPrank();

        assertEq(pool.balances(user1), 1 ether, "Should accept minimum stake");
    }

    function testStakeVeryLarge() public {
        uint largeAmount = 1000 ether;

        vm.startPrank(user1);
        crossToken.approve(address(pool), largeAmount);
        pool.stake(largeAmount);
        vm.stopPrank();

        assertEq(pool.balances(user1), largeAmount, "Should accept large stake");
    }

    // ==================== 추가 스테이킹 테스트 ====================

    function testAdditionalStakeAccumulates() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 100 ether);

        // 첫 번째 stake
        pool.stake(10 ether);
        uint amountBefore = pool.balances(user1);

        _warpSeconds(100);

        // 추가 stake
        pool.stake(20 ether);
        uint amountAfter = pool.balances(user1);

        vm.stopPrank();

        assertEq(amountAfter, amountBefore + 20 ether, "Amount should accumulate");
        assertEq(amountAfter, 30 ether, "Total should be 30 CROSS");
    }

    function testAdditionalStakeDoesNotClaimRewards() public {
        // 첫 번째 stake
        _userStake(user1, 10 ether);

        _warpSeconds(100);

        // 보상 입금
        _depositReward(address(rewardToken1), 100 ether);

        vm.startPrank(user1);
        uint rewardBalanceBefore = rewardToken1.balanceOf(user1);
        assertEq(rewardBalanceBefore, 0, "Should have no claimed rewards yet");

        // 추가 stake - 보상은 claim되지 않음
        crossToken.approve(address(pool), 20 ether);
        pool.stake(20 ether);

        uint rewardBalanceAfter = rewardToken1.balanceOf(user1);
        vm.stopPrank();

        assertEq(rewardBalanceAfter, 0, "Additional stake should not auto-claim rewards");
    }

    // ==================== 언스테이킹 테스트 ====================

    function testUnstakeFullAmount() public {
        uint stakeAmount = 10 ether;

        _userStake(user1, stakeAmount);
        _warpSeconds(100);

        vm.startPrank(user1);
        uint balanceBefore = crossToken.balanceOf(user1);
        pool.unstake();
        uint balanceAfter = crossToken.balanceOf(user1);
        vm.stopPrank();

        assertEq(balanceAfter - balanceBefore, stakeAmount, "Should receive full stake amount");
        assertEq(pool.balances(user1), 0, "User stake should be cleared");
    }

    function testUnstakeWithRewards() public {
        _userStake(user1, 10 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 100 ether);

        vm.startPrank(user1);
        uint rewardBalance1Before = rewardToken1.balanceOf(user1);
        uint crossBalanceBefore = crossToken.balanceOf(user1);

        pool.unstake();

        uint rewardBalance1After = rewardToken1.balanceOf(user1);
        uint crossBalanceAfter = crossToken.balanceOf(user1);
        vm.stopPrank();

        assertEq(crossBalanceAfter - crossBalanceBefore, 10 ether, "Should receive staked CROSS");
        assertApproxEqAbs(
            rewardBalance1After - rewardBalance1Before, 100 ether, 1 ether, "Should receive rewards on unstake"
        );
    }

    function testImmediateUnstake() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 10 ether);
        pool.stake(10 ether);

        // 즉시 unstake (시간 경과 없음)
        uint balanceBefore = crossToken.balanceOf(user1);
        pool.unstake();
        uint balanceAfter = crossToken.balanceOf(user1);
        vm.stopPrank();

        assertEq(balanceAfter - balanceBefore, 10 ether, "Should receive full amount immediately");
    }

    function testStakeAfterUnstake() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 100 ether);

        // 첫 번째 사이클
        pool.stake(10 ether);
        _warpSeconds(50);
        pool.unstake();

        // 두 번째 사이클
        pool.stake(20 ether);
        vm.stopPrank();

        assertEq(pool.balances(user1), 20 ether, "New stake amount should be recorded");
    }

    // ==================== 에러 케이스 ====================

    function testCannotUnstakeWithoutStake() public {
        vm.prank(user1);
        vm.expectRevert(CrossStakingPool.CSPNoStakeFound.selector);
        pool.unstake();
    }

    function testCannotClaimWithoutStake() public {
        vm.prank(user1);
        vm.expectRevert(CrossStakingPool.CSPNoStakeFound.selector);
        pool.claimRewards();
    }

    // ==================== 상태 추적 테스트 ====================

    function testUserBalanceTracking() public {
        _userStake(user1, 10 ether);

        assertEq(pool.balances(user1), 10 ether, "User balance should be 10");

        _warpSeconds(100);

        assertEq(pool.balances(user1), 10 ether, "Balance should not change over time");
    }

    function testTotalStakedCalculation() public {
        _userStake(user1, 10 ether);
        assertEq(pool.totalStaked(), 10 ether, "Total staked should be 10");

        _userStake(user2, 20 ether);
        assertEq(pool.totalStaked(), 30 ether, "Total staked should be 30");

        vm.prank(user1);
        pool.unstake();
        assertEq(pool.totalStaked(), 20 ether, "Total staked should decrease to 20");
    }

    function testBalanceDoesNotOverflow() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 1000 ether);
        pool.stake(1000 ether);
        vm.stopPrank();

        // 매우 긴 시간 (10년)
        vm.warp(block.timestamp + 3650 days);

        assertEq(pool.balances(user1), 1000 ether, "Balance should not overflow or change");
    }

    // ==================== 기본 조회 함수 테스트 ====================

    function testStakingTokenAddress() public view {
        assertEq(address(pool.stakingToken()), address(crossToken), "Staking token should be CROSS");
    }

    function testRewardTokenCount() public view {
        assertEq(pool.rewardTokenCount(), 2, "Should have 2 reward tokens");
    }
}
