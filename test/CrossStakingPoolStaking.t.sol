// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/CrossStakingPoolBase.t.sol";

/**
 * @title CrossStakingPoolStakingTest
 * @notice Stake and unstake behaviour tests
 */
contract CrossStakingPoolStakingTest is CrossStakingPoolBase {
    // ==================== Basic staking tests ====================

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
        uint stakeAmount = 0.5 ether; // below minimum requirement

        vm.startPrank(user1);
        crossToken.approve(address(pool), stakeAmount);

        vm.expectRevert(CrossStakingPool.CSPBelowMinimumStakeAmount.selector);
        pool.stake(stakeAmount);
        vm.stopPrank();
    }

    function testStakeMinimumAmountExact() public {
        uint stakeAmount = 1 ether; // exactly the minimum amount

        vm.startPrank(user1);
        crossToken.approve(address(pool), stakeAmount);
        pool.stake(stakeAmount);
        vm.stopPrank();

        assertEq(pool.balances(user1), stakeAmount, "Should allow exact minimum");
    }

    function testStakeVerySmall() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 1 ether);
        pool.stake(1 ether); // exactly the minimum
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

    // ==================== Additional staking tests ====================

    function testAdditionalStakeAccumulates() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 100 ether);

        // Initial stake
        pool.stake(10 ether);
        uint amountBefore = pool.balances(user1);

        _warpSeconds(100);

        // Additional stake
        pool.stake(20 ether);
        uint amountAfter = pool.balances(user1);

        vm.stopPrank();

        assertEq(amountAfter, amountBefore + 20 ether, "Amount should accumulate");
        assertEq(amountAfter, 30 ether, "Total should be 30 CROSS");
    }

    function testAdditionalStakeDoesNotClaimRewards() public {
        // First stake
        _userStake(user1, 10 ether);

        _warpSeconds(100);

        // Deposit rewards
        _depositReward(address(rewardToken1), 100 ether);

        vm.startPrank(user1);
        uint rewardBalanceBefore = rewardToken1.balanceOf(user1);
        assertEq(rewardBalanceBefore, 0, "Should have no claimed rewards yet");

        // Additional stake should not trigger a claim
        crossToken.approve(address(pool), 20 ether);
        pool.stake(20 ether);

        uint rewardBalanceAfter = rewardToken1.balanceOf(user1);
        vm.stopPrank();

        assertEq(rewardBalanceAfter, 0, "Additional stake should not auto-claim rewards");
    }

    // ==================== Unstaking tests ====================

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

        // Unstake immediately with no time passing
        uint balanceBefore = crossToken.balanceOf(user1);
        pool.unstake();
        uint balanceAfter = crossToken.balanceOf(user1);
        vm.stopPrank();

        assertEq(balanceAfter - balanceBefore, 10 ether, "Should receive full amount immediately");
    }

    function testStakeAfterUnstake() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 100 ether);

        // First cycle
        pool.stake(10 ether);
        _warpSeconds(50);
        pool.unstake();

        // Second cycle
        pool.stake(20 ether);
        vm.stopPrank();

        assertEq(pool.balances(user1), 20 ether, "New stake amount should be recorded");
    }

    // ==================== Error scenarios ====================

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

    // ==================== State tracking tests ====================

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

        // Advance a long period (~10 years)
        vm.warp(block.timestamp + 3650 days);

        assertEq(pool.balances(user1), 1000 ether, "Balance should not overflow or change");
    }

    // ==================== Basic view function tests ====================

    function testStakingTokenAddress() public view {
        assertEq(address(pool.stakingToken()), address(crossToken), "Staking token should be CROSS");
    }

    function testRewardTokenCount() public view {
        assertEq(pool.rewardTokenCount(), 2, "Should have 2 reward tokens");
    }
}
