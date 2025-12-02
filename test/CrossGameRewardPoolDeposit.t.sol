// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "./base/CrossGameRewardPoolBase.t.sol";

/**
 * @title CrossGameRewardPoolDepositTest
 * @notice Deposit and withdraw behaviour tests
 */
contract CrossGameRewardPoolDepositTest is CrossGameRewardPoolBase {
    // ==================== Basic deposit tests ====================

    function testDepositBasic() public {
        uint depositAmount = 10 ether;

        vm.startPrank(user1);
        crossToken.approve(address(pool), depositAmount);
        pool.deposit(depositAmount);
        vm.stopPrank();

        assertEq(pool.balances(user1), depositAmount, "User should have depositd amount recorded");
        assertEq(crossToken.balanceOf(address(pool)), depositAmount, "Pool should hold CROSS tokens");
    }

    function testDepositMinimumAmount() public {
        uint depositAmount = 0.5 ether; // below minimum requirement

        vm.startPrank(user1);
        crossToken.approve(address(pool), depositAmount);

        vm.expectRevert(
            abi.encodeWithSelector(CrossGameRewardPool.CGRPBelowMinimumDepositAmount.selector, depositAmount, 1 ether)
        );
        pool.deposit(depositAmount);
        vm.stopPrank();
    }

    function testDepositMinimumAmountExact() public {
        uint depositAmount = 1 ether; // exactly the minimum amount

        vm.startPrank(user1);
        crossToken.approve(address(pool), depositAmount);
        pool.deposit(depositAmount);
        vm.stopPrank();

        assertEq(pool.balances(user1), depositAmount, "Should allow exact minimum");
    }

    function testDepositVerySmall() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 1 ether);
        pool.deposit(1 ether); // exactly the minimum
        vm.stopPrank();

        assertEq(pool.balances(user1), 1 ether, "Should accept minimum deposit");
    }

    function testDepositVeryLarge() public {
        uint largeAmount = 1000 ether;

        vm.startPrank(user1);
        crossToken.approve(address(pool), largeAmount);
        pool.deposit(largeAmount);
        vm.stopPrank();

        assertEq(pool.balances(user1), largeAmount, "Should accept large deposit");
    }

    // ==================== Additional deposit tests ====================

    function testAdditionalDepositAccumulates() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 100 ether);

        // Initial deposit
        pool.deposit(10 ether);
        uint amountBefore = pool.balances(user1);

        _warpSeconds(100);

        // Additional deposit
        pool.deposit(20 ether);
        uint amountAfter = pool.balances(user1);

        vm.stopPrank();

        assertEq(amountAfter, amountBefore + 20 ether, "Amount should accumulate");
        assertEq(amountAfter, 30 ether, "Total should be 30 CROSS");
    }

    function testAdditionalDepositDoesNotClaimRewards() public {
        // First deposit
        _userDeposit(user1, 10 ether);

        _warpSeconds(100);

        // Deposit rewards
        _depositReward(address(rewardToken1), 100 ether);

        vm.startPrank(user1);
        uint rewardBalanceBefore = rewardToken1.balanceOf(user1);
        assertEq(rewardBalanceBefore, 0, "Should have no claimed rewards yet");

        // Additional deposit should not trigger a claim
        crossToken.approve(address(pool), 20 ether);
        pool.deposit(20 ether);

        uint rewardBalanceAfter = rewardToken1.balanceOf(user1);
        vm.stopPrank();

        assertEq(rewardBalanceAfter, 0, "Additional deposit should not auto-claim rewards");
    }

    // ==================== Withdrawing tests ====================

    function testWithdrawFullAmount() public {
        uint depositAmount = 10 ether;

        _userDeposit(user1, depositAmount);
        _warpSeconds(100);

        vm.startPrank(user1);
        uint balanceBefore = crossToken.balanceOf(user1);
        pool.withdraw(0);
        uint balanceAfter = crossToken.balanceOf(user1);
        vm.stopPrank();

        assertEq(balanceAfter - balanceBefore, depositAmount, "Should receive full deposit amount");
        assertEq(pool.balances(user1), 0, "User deposit should be cleared");
    }

    function testWithdrawWithRewards() public {
        _userDeposit(user1, 10 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 100 ether);

        vm.startPrank(user1);
        uint rewardBalance1Before = rewardToken1.balanceOf(user1);
        uint crossBalanceBefore = crossToken.balanceOf(user1);

        pool.withdraw(0);

        uint rewardBalance1After = rewardToken1.balanceOf(user1);
        uint crossBalanceAfter = crossToken.balanceOf(user1);
        vm.stopPrank();

        assertEq(crossBalanceAfter - crossBalanceBefore, 10 ether, "Should receive depositd CROSS");
        assertApproxEqAbs(
            rewardBalance1After - rewardBalance1Before, 100 ether, 1 ether, "Should receive rewards on withdraw"
        );
    }

    function testImmediateWithdraw() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 10 ether);
        pool.deposit(10 ether);

        // Withdraw immediately with no time passing
        uint balanceBefore = crossToken.balanceOf(user1);
        pool.withdraw(0);
        uint balanceAfter = crossToken.balanceOf(user1);
        vm.stopPrank();

        assertEq(balanceAfter - balanceBefore, 10 ether, "Should receive full amount immediately");
    }

    function testDepositAfterWithdraw() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 100 ether);

        // First cycle
        pool.deposit(10 ether);
        _warpSeconds(50);
        pool.withdraw(0);

        // Second cycle
        pool.deposit(20 ether);
        vm.stopPrank();

        assertEq(pool.balances(user1), 20 ether, "New deposit amount should be recorded");
    }

    // ==================== Error scenarios ====================

    function testCannotWithdrawWithoutDeposit() public {
        vm.prank(user1);
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPool.CGRPNoDepositFound.selector, user1));
        pool.withdraw(0);
    }

    function testCannotClaimWithoutDeposit() public {
        vm.prank(user1);
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPool.CGRPNoDepositFound.selector, user1));
        pool.claimRewards();
    }

    // ==================== State tracking tests ====================

    function testUserBalanceTracking() public {
        _userDeposit(user1, 10 ether);

        assertEq(pool.balances(user1), 10 ether, "User balance should be 10");

        _warpSeconds(100);

        assertEq(pool.balances(user1), 10 ether, "Balance should not change over time");
    }

    function testTotalDepositedCalculation() public {
        _userDeposit(user1, 10 ether);
        assertEq(pool.totalDeposited(), 10 ether, "Total depositd should be 10");

        _userDeposit(user2, 20 ether);
        assertEq(pool.totalDeposited(), 30 ether, "Total depositd should be 30");

        vm.prank(user1);
        pool.withdraw(0);
        assertEq(pool.totalDeposited(), 20 ether, "Total depositd should decrease to 20");
    }

    function testBalanceDoesNotOverflow() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 1000 ether);
        pool.deposit(1000 ether);
        vm.stopPrank();

        // Advance a long period (~10 years)
        vm.warp(block.timestamp + 3650 days);

        assertEq(pool.balances(user1), 1000 ether, "Balance should not overflow or change");
    }

    // ==================== Basic view function tests ====================

    function testDepositingTokenAddress() public view {
        assertEq(address(pool.depositToken()), address(crossToken), "Deposit token should be CROSS");
    }

    function testRewardTokenCount() public view {
        assertEq(pool.rewardTokenCount(), 2, "Should have 2 reward tokens");
    }

    // ==================== Partial withdrawal tests ====================

    function testPartialWithdraw() public {
        uint depositAmount = 100 ether;
        uint withdrawAmount = 30 ether;

        _userDeposit(user1, depositAmount);
        _warpSeconds(100);

        vm.startPrank(user1);
        uint balanceBefore = crossToken.balanceOf(user1);
        pool.withdraw(withdrawAmount);
        uint balanceAfter = crossToken.balanceOf(user1);
        vm.stopPrank();

        assertEq(balanceAfter - balanceBefore, withdrawAmount, "Should receive partial withdraw amount");
        assertEq(pool.balances(user1), depositAmount - withdrawAmount, "Remaining balance should be correct");
        assertEq(
            pool.totalDeposited(), depositAmount - withdrawAmount, "Total deposited should decrease by withdraw amount"
        );
    }

    function testPartialWithdrawMultipleTimes() public {
        uint depositAmount = 100 ether;

        _userDeposit(user1, depositAmount);
        _warpSeconds(100);

        vm.startPrank(user1);

        // First partial withdraw
        pool.withdraw(20 ether);
        assertEq(pool.balances(user1), 80 ether, "Balance after first withdraw");

        // Second partial withdraw
        pool.withdraw(30 ether);
        assertEq(pool.balances(user1), 50 ether, "Balance after second withdraw");

        // Third partial withdraw
        pool.withdraw(50 ether);
        assertEq(pool.balances(user1), 0, "Balance after final withdraw");

        vm.stopPrank();
    }

    function testPartialWithdrawWithRewards() public {
        uint depositAmount = 100 ether;
        uint withdrawAmount = 40 ether;

        _userDeposit(user1, depositAmount);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 200 ether);

        vm.startPrank(user1);
        uint crossBalanceBefore = crossToken.balanceOf(user1);
        uint rewardBalanceBefore = rewardToken1.balanceOf(user1);

        pool.withdraw(withdrawAmount);

        uint crossBalanceAfter = crossToken.balanceOf(user1);
        uint rewardBalanceAfter = rewardToken1.balanceOf(user1);
        vm.stopPrank();

        // Check deposit token
        assertEq(crossBalanceAfter - crossBalanceBefore, withdrawAmount, "Should receive partial deposit");
        assertEq(pool.balances(user1), depositAmount - withdrawAmount, "Remaining deposit balance");

        // Check rewards (all rewards should be claimed on partial withdraw)
        assertApproxEqAbs(
            rewardBalanceAfter - rewardBalanceBefore,
            200 ether,
            1 ether,
            "Should receive all rewards on partial withdraw"
        );
    }

    function testCannotPartialWithdrawMoreThanBalance() public {
        uint depositAmount = 100 ether;
        uint withdrawAmount = 150 ether; // More than deposited

        _userDeposit(user1, depositAmount);
        _warpSeconds(100);

        vm.prank(user1);
        vm.expectRevert(
            abi.encodeWithSelector(CrossGameRewardPool.CGRPInsufficientBalance.selector, depositAmount, withdrawAmount)
        );
        pool.withdraw(withdrawAmount);
    }

    function testZeroAmountWithdrawsAll() public {
        uint depositAmount = 100 ether;

        _userDeposit(user1, depositAmount);
        _warpSeconds(100);

        vm.startPrank(user1);
        uint balanceBefore = crossToken.balanceOf(user1);
        pool.withdraw(0); // 0 = withdraw all
        uint balanceAfter = crossToken.balanceOf(user1);
        vm.stopPrank();

        assertEq(balanceAfter - balanceBefore, depositAmount, "Should withdraw full amount with 0");
        assertEq(pool.balances(user1), 0, "Balance should be zero");
    }

    function testPartialWithdrawAndContinueEarning() public {
        // User1 deposits 100 ether
        _userDeposit(user1, 100 ether);
        _warpSeconds(100);

        // Add first reward batch
        _depositReward(address(rewardToken1), 100 ether);

        // User1 partially withdraws 50 ether
        vm.prank(user1);
        pool.withdraw(50 ether);

        assertEq(pool.balances(user1), 50 ether, "Should have 50 ether left");

        // Add second reward batch (should only be distributed to remaining 50 ether)
        _depositReward(address(rewardToken1), 50 ether);

        // Check pending rewards
        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 50 ether, 1 ether, "Should have pending rewards from second batch");

        // Final withdraw
        vm.prank(user1);
        pool.withdraw(0);

        // Total rewards should be: 100 (first batch) + 50 (second batch) = 150 ether
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 150 ether, 2 ether, "Total rewards collected");
    }
}
