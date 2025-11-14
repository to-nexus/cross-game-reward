// SPDX-License-Identifier: MIT
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

        vm.expectRevert(CrossGameRewardPool.CGRPBelowMinimumDepositAmount.selector);
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
        pool.withdraw();
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

        pool.withdraw();

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
        pool.withdraw();
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
        pool.withdraw();

        // Second cycle
        pool.deposit(20 ether);
        vm.stopPrank();

        assertEq(pool.balances(user1), 20 ether, "New deposit amount should be recorded");
    }

    // ==================== Error scenarios ====================

    function testCannotWithdrawWithoutDeposit() public {
        vm.prank(user1);
        vm.expectRevert(CrossGameRewardPool.CGRPNoDepositFound.selector);
        pool.withdraw();
    }

    function testCannotClaimWithoutDeposit() public {
        vm.prank(user1);
        vm.expectRevert(CrossGameRewardPool.CGRPNoDepositFound.selector);
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
        pool.withdraw();
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
}
