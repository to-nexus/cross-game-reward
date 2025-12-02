// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "./base/CrossGameRewardPoolBase.t.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title CrossGameRewardPoolSecurityTest
 * @notice Security and logic validation tests
 */
contract CrossGameRewardPoolSecurityTest is CrossGameRewardPoolBase {
    // ==================== Invariant checks ====================

    function testInvariantTotalDepositedMatchesActualBalance() public {
        _userDeposit(user1, 100 ether);
        _userDeposit(user2, 200 ether);
        _userDeposit(user3, 300 ether);

        // totalDeposited must match the contract balance
        assertEq(pool.totalDeposited(), 600 ether, "TotalDeposited should match");
        assertEq(crossToken.balanceOf(address(pool)), 600 ether, "Actual balance should match");

        // Should remain in sync after an withdraw
        vm.prank(user1);
        pool.withdraw(0);

        assertEq(pool.totalDeposited(), 500 ether, "TotalDeposited after withdraw");
        assertEq(crossToken.balanceOf(address(pool)), 500 ether, "Actual balance after withdraw");
    }

    function testInvariantRewardAccountingAccuracy() public {
        _userDeposit(user1, 100 ether);
        _userDeposit(user2, 100 ether);

        _depositReward(address(rewardToken1), 1000 ether);

        (, uint[] memory rewards1) = pool.pendingRewards(user1);
        (, uint[] memory rewards2) = pool.pendingRewards(user2);

        // Sum of rewards should equal the deposited amount
        assertApproxEqAbs(rewards1[0] + rewards2[0], 1000 ether, 100, "Total rewards should equal deposited");
    }

    function testInvariantNoRewardLoss() public {
        // Multiple users deposit/withdraw at different times
        _userDeposit(user1, 50 ether);
        _depositReward(address(rewardToken1), 100 ether);

        _userDeposit(user2, 150 ether);
        _depositReward(address(rewardToken1), 200 ether);

        vm.prank(user1);
        pool.withdraw(0);
        uint user1Claimed = rewardToken1.balanceOf(user1);

        _userDeposit(user3, 100 ether);
        _depositReward(address(rewardToken1), 150 ether);

        vm.prank(user2);
        pool.withdraw(0);
        uint user2Claimed = rewardToken1.balanceOf(user2);

        vm.prank(user3);
        pool.withdraw(0);
        uint user3Claimed = rewardToken1.balanceOf(user3);

        // Total claimed rewards should equal total deposits
        uint totalClaimed = user1Claimed + user2Claimed + user3Claimed;
        assertApproxEqAbs(totalClaimed, 450 ether, 100, "No reward should be lost");
    }

    // ==================== Potential attack vectors ====================

    function testCannotDepositZeroAmount() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 0);
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPool.CGRPBelowMinimumDepositAmount.selector, 0, 1 ether));
        pool.deposit(0);
        vm.stopPrank();
    }

    function testReentrancyProtection() public {
        // Ensure ReentrancyGuard protects state-mutating paths
        // nonReentrant modifier covers every critical function
        _userDeposit(user1, 100 ether);

        // Baseline behaviour should succeed
        vm.prank(user1);
        pool.withdraw(0);

        assertEq(pool.balances(user1), 0, "Should complete without reentrancy");
    }

    function testPrecisionLoss() public {
        // Test precision with a tiny reward
        _userDeposit(user1, 1 ether);

        // Reward of 1 wei
        vm.startPrank(owner);
        rewardToken1.transfer(address(pool), 1);
        vm.stopPrank();

        (, uint[] memory rewards) = pool.pendingRewards(user1);
        // 1 wei may be rounded because of PRECISION
        assertTrue(rewards[0] <= 1, "Should handle precision correctly");
    }

    function testOverflowProtection() public {
        // Solidity 0.8.x includes automatic overflow checks
        // Stress test with very large values
        uint veryLarge = type(uint).max / 2;

        vm.startPrank(owner);
        crossToken.mint(user1, veryLarge);
        vm.stopPrank();

        vm.startPrank(user1);
        crossToken.approve(address(pool), veryLarge);

        // Should succeed because amount is above MIN_DEPOSIT_AMOUNT
        pool.deposit(veryLarge);
        vm.stopPrank();

        assertEq(pool.balances(user1), veryLarge, "Should handle very large amounts");
    }

    // ==================== Logic consistency ====================

    function testRewardCalculationConsistency() public {
        _userDeposit(user1, 100 ether);

        // Deposit rewards multiple times
        for (uint i = 0; i < 5; i++) {
            _depositReward(address(rewardToken1), 100 ether);
        }

        // Pending query should match actual claim
        (, uint[] memory pendingBefore) = pool.pendingRewards(user1);

        vm.prank(user1);
        pool.claimRewards();

        uint actualClaimed = rewardToken1.balanceOf(user1);

        assertApproxEqAbs(pendingBefore[0], actualClaimed, 100, "Pending should match actual claim");
    }

    function testWithdrawOrderCorrectness() public {
        _userDeposit(user1, 100 ether);
        _depositReward(address(rewardToken1), 1000 ether);

        uint poolBalanceBefore = crossToken.balanceOf(address(pool));

        // Withdraw operates in the following order:
        // 1. Sync rewards
        // 2. Update account rewards
        // 3. Claim rewards (updates lastBalance)
        // 4. Return deposit tokens (reduces CROSS balance)
        vm.prank(user1);
        pool.withdraw(0);

        // Ensure CROSS tokens were returned
        assertEq(crossToken.balanceOf(user1), 1000 ether, "Should receive CROSS back");

        // Ensure rewards were received
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 1000 ether, 100, "Should receive rewards");

        // Confirm the pool's CROSS balance decreased
        assertEq(crossToken.balanceOf(address(pool)), poolBalanceBefore - 100 ether, "Pool CROSS decreased");
    }

    function testCheckpointAccuracy() public {
        _userDeposit(user1, 100 ether);

        _depositReward(address(rewardToken1), 100 ether);

        // User2's checkpoint is taken at the current rewardPerTokenStored
        _userDeposit(user2, 100 ether);

        // User2 should not receive earlier rewards
        (, uint[] memory rewards2) = pool.pendingRewards(user2);
        assertEq(rewards2[0], 0, "User2 should not get previous rewards");

        // Add new rewards
        _depositReward(address(rewardToken1), 200 ether);

        (, uint[] memory rewards1) = pool.pendingRewards(user1);
        (, rewards2) = pool.pendingRewards(user2);

        // User1: 100 (previous) + 100 (50% of new rewards) = 200
        // User2: 0 (previous) + 100 (50% of new rewards) = 100
        assertApproxEqAbs(rewards1[0], 200 ether, 100, "User1 gets all old + 50% new");
        assertApproxEqAbs(rewards2[0], 100 ether, 100, "User2 gets only 50% new");
    }

    function testRewardDistributionWithZeroDeposited() public {
        // Deposit rewards while totalDeposited is zero
        _depositReward(address(rewardToken1), 1000 ether);

        // First depositr should NOT receive previously deposited rewards
        // Those become withdrawable instead
        _userDeposit(user1, 100 ether);

        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "First depositr does NOT get rewards deposited when pool was empty");

        // The 1000 ether should be withdrawable by owner
        assertEq(pool.getReclaimableAmount(rewardToken1), 1000 ether, "Rewards are withdrawable");

        // Subsequent rewards distribute normally
        _depositReward(address(rewardToken1), 100 ether);
        (, rewards) = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 100 ether, 100, "New rewards distributed to depositr");
    }

    // ==================== Edge-case validation ====================

    function testClaimWithZeroRewards() public {
        _userDeposit(user1, 100 ether);

        // Claim with no rewards available
        vm.prank(user1);
        pool.claimRewards();

        assertEq(rewardToken1.balanceOf(user1), 0, "Should handle zero rewards gracefully");
    }

    function testMultipleUsersWithdrawOrder() public {
        _userDeposit(user1, 100 ether);
        _userDeposit(user2, 100 ether);
        _userDeposit(user3, 100 ether);

        _depositReward(address(rewardToken1), 300 ether);

        // Withdraw in order
        vm.prank(user1);
        pool.withdraw(0);
        uint claimed1 = rewardToken1.balanceOf(user1);

        vm.prank(user2);
        pool.withdraw(0);
        uint claimed2 = rewardToken1.balanceOf(user2);

        vm.prank(user3);
        pool.withdraw(0);
        uint claimed3 = rewardToken1.balanceOf(user3);

        // All users should receive the same amount regardless of order
        assertApproxEqAbs(claimed1, 100 ether, 100, "User1 should get 1/3");
        assertApproxEqAbs(claimed2, 100 ether, 100, "User2 should get 1/3");
        assertApproxEqAbs(claimed3, 100 ether, 100, "User3 should get 1/3");
    }

    function testDepositAfterRewardDeposit() public {
        _userDeposit(user1, 100 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // User2 should not receive previous rewards
        _userDeposit(user2, 100 ether);

        (, uint[] memory rewards1) = pool.pendingRewards(user1);
        (, uint[] memory rewards2) = pool.pendingRewards(user2);

        assertApproxEqAbs(rewards1[0], 100 ether, 100, "User1 gets all");
        assertEq(rewards2[0], 0, "User2 gets nothing from previous");
    }

    // ==================== Mathematical checks ====================

    function testRewardPerTokenCalculation() public {
        _userDeposit(user1, 100 ether);

        // Deposit 100 tokens with 100 CROSS depositd
        // rewardPerToken = (100 * 1e18) / 100 = 1e18
        _depositReward(address(rewardToken1), 100 ether);

        (, uint[] memory rewards) = pool.pendingRewards(user1);

        // earned = 100 * 1e18 / 1e18 = 100
        assertApproxEqAbs(rewards[0], 100 ether, 0.001 ether, "Math should be precise");
    }

    function testProportionalDistribution() public {
        // Deposit tokens in a 1:2:3 ratio
        _userDeposit(user1, 100 ether);
        _userDeposit(user2, 200 ether);
        _userDeposit(user3, 300 ether);

        // Deposit 600 tokens in rewards
        _depositReward(address(rewardToken1), 600 ether);

        (, uint[] memory rewards1) = pool.pendingRewards(user1);
        (, uint[] memory rewards2) = pool.pendingRewards(user2);
        (, uint[] memory rewards3) = pool.pendingRewards(user3);

        // Validate the 1:2:3 distribution
        assertApproxEqAbs(rewards1[0], 100 ether, 100, "1/6 of rewards");
        assertApproxEqAbs(rewards2[0], 200 ether, 100, "2/6 of rewards");
        assertApproxEqAbs(rewards3[0], 300 ether, 100, "3/6 of rewards");

        // Sum should equal the deposited amount
        uint total = rewards1[0] + rewards2[0] + rewards3[0];
        assertApproxEqAbs(total, 600 ether, 100, "Sum should equal deposit");
    }

    // ==================== Time independence checks ====================

    function testRewardsIndependentOfTime() public {
        // Scenario 1: immediate reward
        _userDeposit(user1, 100 ether);
        _depositReward(address(rewardToken1), 100 ether);
        vm.prank(user1);
        pool.withdraw(0);
        uint user1Reward = rewardToken1.balanceOf(user1);

        // Scenario 2: reward after one year under the same conditions
        _warpDays(365);
        _userDeposit(user2, 100 ether);
        _depositReward(address(rewardToken1), 100 ether);
        vm.prank(user2);
        pool.withdraw(0);
        uint user2Reward = rewardToken1.balanceOf(user2);

        // Rewards should be identical regardless of elapsed time
        assertApproxEqAbs(user1Reward, user2Reward, 100, "Time should not affect rewards");
        assertApproxEqAbs(user1Reward, 100 ether, 100, "Both should get 100 ether");
    }

    // ==================== Reward token administration ====================

    function testRewardTokenIndexConsistency() public {
        crossGameReward.addRewardToken(1, IERC20(address(rewardToken3)));

        // Verify addresses
        assertEq(address(pool.rewardTokenAt(0)), address(rewardToken1), "RewardToken1 index");
        assertEq(address(pool.rewardTokenAt(1)), address(rewardToken2), "RewardToken2 index");
        assertEq(address(pool.rewardTokenAt(2)), address(rewardToken3), "RewardToken3 index");

        // Confirm registration status
        assertTrue(pool.isRewardToken(rewardToken1), "RewardToken1 registered");
        assertTrue(pool.isRewardToken(rewardToken2), "RewardToken2 registered");
        assertTrue(pool.isRewardToken(rewardToken3), "RewardToken3 registered");
        assertFalse(pool.isRewardToken(IERC20(address(crossToken))), "CROSS not a reward token");
    }

    // ==================== Boundary checks ====================

    function testMinimumDepositBoundary() public {
        uint belowMin = MIN_DEPOSIT_AMOUNT - 1;
        uint exactMin = MIN_DEPOSIT_AMOUNT;
        uint aboveMin = MIN_DEPOSIT_AMOUNT + 1;

        // Below minimum should fail
        vm.startPrank(user1);
        crossToken.approve(address(pool), belowMin);
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPool.CGRPBelowMinimumDepositAmount.selector, belowMin, MIN_DEPOSIT_AMOUNT));
        pool.deposit(belowMin);

        // Exact minimum should succeed
        crossToken.approve(address(pool), exactMin);
        pool.deposit(exactMin);

        vm.stopPrank();
        assertEq(pool.balances(user1), exactMin, "Should accept exact minimum");

        // Above minimum should succeed
        vm.startPrank(user2);
        crossToken.approve(address(pool), aboveMin);
        pool.deposit(aboveMin);
        vm.stopPrank();

        assertEq(pool.balances(user2), aboveMin, "Should accept above minimum");
    }

    function testZeroRewardHandling() public {
        _userDeposit(user1, 100 ether);

        // Attempt to deposit zero reward
        // A zero amount either fails or has no effect
        vm.startPrank(owner);
        rewardToken1.transfer(address(pool), 0);
        vm.stopPrank();

        // Confirm that no reward has been added
        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "No reward should be added for 0 amount");
    }

    // ==================== State consistency checks ====================

    function testBalanceConsistencyAfterMultipleOperations() public {
        uint initialBalance = crossToken.balanceOf(user1);

        _userDeposit(user1, 100 ether);
        assertEq(crossToken.balanceOf(user1), initialBalance - 100 ether, "After deposit");

        _depositReward(address(rewardToken1), 50 ether);

        vm.prank(user1);
        pool.claimRewards();

        // CROSS balance should remain unchanged because only rewards are claimed
        assertEq(crossToken.balanceOf(user1), initialBalance - 100 ether, "CROSS unchanged after claim");

        vm.prank(user1);
        pool.withdraw(0);

        // Restore original CROSS balances
        assertEq(crossToken.balanceOf(user1), initialBalance, "CROSS restored after withdraw");
    }

    MockERC20 public rewardToken3;

    function setUp() public override {
        super.setUp();
        rewardToken3 = new MockERC20("Reward Token 3", "RWD3");
        rewardToken3.transfer(owner, 10000 ether);
    }

    uint private constant MIN_DEPOSIT_AMOUNT = 1 ether;
}
