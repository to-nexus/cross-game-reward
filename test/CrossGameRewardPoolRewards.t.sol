// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "./base/CrossGameRewardPoolBase.t.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title CrossGameRewardPoolRewardsTest
 * @notice Reward accrual, distribution, and claim tests
 */
contract CrossGameRewardPoolRewardsTest is CrossGameRewardPoolBase {
    // ==================== Reward accrual tests ====================

    function testRewardAccumulation() public {
        _userDeposit(user1, 10 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 100 ether);
        _depositReward(address(rewardToken2), 50 ether);

        (, uint[] memory rewards) = pool.pendingRewards(user1);

        assertEq(rewards[0], 100 ether, "Reward token 1 should accumulate");
        assertEq(rewards[1], 50 ether, "Reward token 2 should accumulate");
    }

    function testRewardAccumulationWithVerySmallDeposit() public {
        _userDeposit(user1, 1 ether); // minimum deposit amount
        _warpSeconds(1000);

        _depositReward(address(rewardToken1), 1000 ether);

        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 1000 ether, "Small deposit should get all rewards");
    }

    function testRewardAccumulationWithVeryLargeDeposit() public {
        _userDeposit(user1, 1000 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 10000 ether);

        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 10000 ether, "Large deposit should get all rewards");
    }

    // ==================== Reward claim tests ====================

    function testClaimRewards() public {
        _userDeposit(user1, 10 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 100 ether);
        _depositReward(address(rewardToken2), 50 ether);

        vm.startPrank(user1);
        uint rewardBalance1Before = rewardToken1.balanceOf(user1);
        uint rewardBalance2Before = rewardToken2.balanceOf(user1);

        pool.claimRewards();

        uint rewardBalance1After = rewardToken1.balanceOf(user1);
        uint rewardBalance2After = rewardToken2.balanceOf(user1);
        vm.stopPrank();

        assertEq(rewardBalance1After - rewardBalance1Before, 100 ether, "Should receive reward token 1");
        assertEq(rewardBalance2After - rewardBalance2Before, 50 ether, "Should receive reward token 2");
    }

    function testClaimSpecificReward() public {
        _userDeposit(user1, 10 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 100 ether);
        _depositReward(address(rewardToken2), 50 ether);

        // Claim only the first reward token
        vm.prank(user1);
        pool.claimReward(rewardToken1);

        assertEq(rewardToken1.balanceOf(user1), 100 ether, "Claimed reward1");
        assertEq(rewardToken2.balanceOf(user1), 0, "Not claimed reward2 yet");

        // Claim the second reward token
        vm.prank(user1);
        pool.claimReward(rewardToken2);

        assertEq(rewardToken2.balanceOf(user1), 50 ether, "Claimed reward2");
    }

    function testMultipleClaimsAccumulate() public {
        _userDeposit(user1, 10 ether);
        _warpSeconds(50);

        // First reward batch
        _depositReward(address(rewardToken1), 50 ether);

        vm.prank(user1);
        pool.claimRewards();
        assertEq(rewardToken1.balanceOf(user1), 50 ether, "First claim");

        _warpSeconds(50);

        // Second reward batch
        _depositReward(address(rewardToken1), 50 ether);

        vm.prank(user1);
        pool.claimRewards();
        assertEq(rewardToken1.balanceOf(user1), 100 ether, "Claims accumulate");
    }

    // ==================== Multi-user reward distribution tests ====================

    function testMultipleUsersRewardDistribution() public {
        // User1 deposits 10 CROSS
        _userDeposit(user1, 10 ether);
        _warpSeconds(50);

        // User2 deposits 10 CROSS
        _userDeposit(user2, 10 ether);
        _warpSeconds(50);

        // Deposit rewards (20 CROSS total depositd)
        _depositReward(address(rewardToken1), 200 ether);

        (, uint[] memory rewardsUser1) = pool.pendingRewards(user1);
        (, uint[] memory rewardsUser2) = pool.pendingRewards(user2);

        // User1: (10 / 20) × 200 = 100 ether
        // User2: (10 / 20) × 200 = 100 ether
        assertEq(rewardsUser1[0], 100 ether, "User1 gets 50% (equal deposits)");
        assertEq(rewardsUser2[0], 100 ether, "User2 gets 50% (equal deposits)");
    }

    function testThreeUsersComplexScenario() public {
        // User1: 10 CROSS at t=0
        _userDeposit(user1, 10 ether);
        _warpSeconds(50);

        // User2: 20 CROSS at t=50
        _userDeposit(user2, 20 ether);
        _warpSeconds(50);

        // At t=100 deposit rewards (total: 30 CROSS)
        _depositReward(address(rewardToken1), 150 ether);

        (, uint[] memory rewards1) = pool.pendingRewards(user1);
        (, uint[] memory rewards2) = pool.pendingRewards(user2);

        // User1: (10 / 30) × 150 = 50 ether
        // User2: (20 / 30) × 150 = 100 ether
        assertApproxEqAbs(rewards1[0], 50 ether, 10, "User1: 1/3 of rewards");
        assertApproxEqAbs(rewards2[0], 100 ether, 10, "User2: 2/3 of rewards");
    }

    function testRewardDistributionWithUnequalDeposits() public {
        // User1: 30 CROSS
        _userDeposit(user1, 30 ether);

        // User2: 70 CROSS
        _userDeposit(user2, 70 ether);

        // Deposit rewards (total deposit: 100 CROSS)
        _depositReward(address(rewardToken1), 1000 ether);

        (, uint[] memory rewards1) = pool.pendingRewards(user1);
        (, uint[] memory rewards2) = pool.pendingRewards(user2);

        // User1: 30%
        // User2: 70%
        assertApproxEqAbs(rewards1[0], 300 ether, 10, "User1: 30%");
        assertApproxEqAbs(rewards2[0], 700 ether, 10, "User2: 70%");
    }

    // ==================== Multiple reward token tests ====================

    function testMultipleRewardTokens() public {
        _userDeposit(user1, 10 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 100 ether);
        _depositReward(address(rewardToken2), 200 ether);

        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 100 ether, 10, "Reward token 1");
        assertApproxEqAbs(rewards[1], 200 ether, 10, "Reward token 2");
    }

    // ==================== Time-based reward variations ====================

    function testRewardBeforeAndAfterDeposit() public {
        // Day 1: User A deposits 100 CROSS
        _userDeposit(user1, 100 ether);
        _warpDays(1);

        // Day 2: deposit reward #1 (100)
        _depositReward(address(rewardToken1), 100 ether);
        _warpDays(8);

        // Day 10: User B deposits 100 CROSS
        _userDeposit(user2, 100 ether);
        _warpDays(1);

        // Day 11: deposit reward #2 (200)
        _depositReward(address(rewardToken1), 200 ether);
        _warpDays(9);

        // Day 20: evaluate rewards
        (, uint[] memory rewardsA) = pool.pendingRewards(user1);
        (, uint[] memory rewardsB) = pool.pendingRewards(user2);

        // User A: 100 from reward1 (full share) + 100 from reward2 (50% share) = 200
        // User B: 0 from reward1 (joined later) + 100 from reward2 (50% share) = 100
        assertApproxEqAbs(rewardsA[0], 200 ether, 10, "User A: reward1 + 50% reward2");
        assertApproxEqAbs(rewardsB[0], 100 ether, 10, "User B: only 50% reward2");

        // Total distribution check
        assertApproxEqAbs(rewardsA[0] + rewardsB[0], 300 ether, 20, "Total distributed");
    }

    // ==================== Edge cases ====================

    function testZeroDepositrs() public {
        // Deposit rewards with no depositrs
        _depositReward(address(rewardToken1), 100 ether);

        // First depositr should NOT receive rewards accumulated while pool was empty
        // Those rewards are marked as withdrawable instead
        _userDeposit(user1, 10 ether);

        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "First depositr does NOT get rewards deposited when pool was empty");

        // The 100 ether should be withdrawable by owner
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether, "Rewards are withdrawable");
    }

    function testInvalidRewardTokenIndex() public {
        _userDeposit(user1, 10 ether);

        vm.prank(user1);
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPool.CGRPInvalidRewardToken.selector, address(uint160(0xdead))));
        pool.claimReward(IERC20(address(uint160(0xdead))));
    }

    function testZeroAmountTransfer() public {
        _userDeposit(user1, 10 ether);

        // Transferring zero should have no effect
        vm.startPrank(owner);
        rewardToken1.transfer(address(pool), 0);
        vm.stopPrank();

        // Verify no rewards were created
        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "No reward for 0 transfer");
    }

    function testPendingRewardsAfterClaim() public {
        _userDeposit(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Claim
        vm.prank(user1);
        pool.claimRewards();

        // Pending rewards should be zero
        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "Pending rewards should be zero after claim");

        // New reward
        _depositReward(address(rewardToken1), 50 ether);
        (, rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 50 ether, "New rewards should accumulate");
    }

    // ==================== Direct transfer detection tests ====================

    function testDirectTransferDetection() public {
        _userDeposit(user1, 10 ether);

        // Direct transfer without helper
        vm.prank(owner);
        rewardToken1.transfer(address(pool), 100 ether);

        // Detected when another deposit action happens (deposit/withdraw/claim)
        _userDeposit(user2, 10 ether);

        // Reward should reflect the transfer
        (, uint[] memory rewardsAfter) = pool.pendingRewards(user1);
        assertEq(rewardsAfter[0], 100 ether, "Direct transfer detected on next action");
    }

    function testDirectTransferWithDepositReward() public {
        _userDeposit(user1, 10 ether);

        // Perform direct transfer first
        vm.prank(owner);
        rewardToken1.transfer(address(pool), 50 ether);

        // Then deposit via helper
        // RewardSynced event records the full delta (150)
        _depositReward(address(rewardToken1), 100 ether);

        // Verify reward equals 50 + 100 = 150
        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 150 ether, "Total rewards should include direct transfer");
    }

    function testMultipleDirectTransfers() public {
        _userDeposit(user1, 10 ether);

        // Multiple direct transfers
        vm.startPrank(owner);
        rewardToken1.transfer(address(pool), 30 ether);
        rewardToken1.transfer(address(pool), 20 ether);
        rewardToken1.transfer(address(pool), 50 ether);
        vm.stopPrank();

        // One additional action detects them all
        _userDeposit(user2, 10 ether);

        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 100 ether, "All direct transfers should be detected");
    }

    // ==================== Dust redistribution tests ====================

    /// @notice Tests that rounding dust is minimal thanks to PRECISION
    /// @dev With 1e18 PRECISION, dust is measured in wei, not ether!
    function testDustAutoRedistribution() public {
        // Setup: User1 deposits 1, User2 deposits 2 (1:2 ratio, total 3)
        _userDeposit(user1, 1 ether);
        _userDeposit(user2, 2 ether);

        // First distribution: 100 tokens
        _depositReward(address(rewardToken1), 100 ether);

        // Thanks to PRECISION (1e18), distribution is very accurate!
        // rewardPerTokenStored = (100 * 1e18 * 1e18) / 3
        // User1 (1 ether deposit): gets exactly 33.333... ether
        // User2 (2 ether deposit): gets exactly 66.666... ether

        (, uint[] memory rewards1) = pool.pendingRewards(user1);
        (, uint[] memory rewards2) = pool.pendingRewards(user2);

        // Verify almost exact distribution (dust in wei, not ether!)
        assertApproxEqAbs(rewards1[0], 33.333333333333333333 ether, 1, "User1 gets 1/3");
        assertApproxEqAbs(rewards2[0], 66.666666666666666666 ether, 1, "User2 gets 2/3");

        // Track balances before claiming
        uint totalDeposited = 100 ether;

        // Users claim their rewards
        vm.prank(user1);
        pool.claimRewards();
        vm.prank(user2);
        pool.claimRewards();

        uint user1Claimed = rewardToken1.balanceOf(user1);
        uint user2Claimed = rewardToken1.balanceOf(user2);
        uint poolDust = rewardToken1.balanceOf(address(pool));

        // Exact accounting: deposited = claimed + dust
        assertEq(totalDeposited, user1Claimed + user2Claimed + poolDust, "Perfect accounting");

        // Dust should be tiny (wei, not ether!)
        assertLt(poolDust, 2, "Dust is negligible (< 2 wei)");

        // Second distribution: Add 90 tokens
        _depositReward(address(rewardToken1), 90 ether);
        totalDeposited += 90 ether;

        // Claim again
        vm.prank(user1);
        pool.claimRewards();
        vm.prank(user2);
        pool.claimRewards();

        uint user1Total = rewardToken1.balanceOf(user1);
        uint user2Total = rewardToken1.balanceOf(user2);
        uint finalDust = rewardToken1.balanceOf(address(pool));

        // Still perfect accounting
        assertEq(totalDeposited, user1Total + user2Total + finalDust, "Still perfect accounting");

        // Dust remains minimal
        assertLt(finalDust, 3, "Dust still negligible (< 3 wei)");

        // Ratio is maintained
        assertApproxEqRel(user1Total, (totalDeposited * 1) / 3, 0.0001e18, "User1 gets ~1/3");
        assertApproxEqRel(user2Total, (totalDeposited * 2) / 3, 0.0001e18, "User2 gets ~2/3");
    }

    /// @notice Tests dust accumulation remains minimal over many rounds
    /// @dev Multiple rounds with PRECISION means dust stays in wei, not ether
    function testDustEventuallyDistributed() public {
        // Setup: 1:2 ratio
        _userDeposit(user1, 1 ether);
        _userDeposit(user2, 2 ether);

        uint totalDeposited = 0;

        // Do 10 rounds of 100 token distribution
        for (uint i = 0; i < 10; i++) {
            _depositReward(address(rewardToken1), 100 ether);
            totalDeposited += 100 ether;
        }

        // Claim all rewards
        vm.prank(user1);
        pool.claimRewards();
        vm.prank(user2);
        pool.claimRewards();

        uint user1Balance = rewardToken1.balanceOf(user1);
        uint user2Balance = rewardToken1.balanceOf(user2);
        uint poolDust = rewardToken1.balanceOf(address(pool));

        // Perfect accounting
        assertEq(totalDeposited, user1Balance + user2Balance + poolDust, "Total matches");

        // Dust is minimal (wei, not ether!)
        assertLt(poolDust, 10, "Dust < 10 wei after 10 rounds");

        // User1 withdraws (leaves only User2)
        vm.prank(user1);
        pool.withdraw(0);

        // Now add one more reward with only User2 depositing
        _depositReward(address(rewardToken1), 100 ether);
        totalDeposited += 100 ether;

        // User2 should get all the new reward + previous dust
        vm.prank(user2);
        pool.claimRewards();

        // Final accounting
        uint finalUser1 = rewardToken1.balanceOf(user1);
        uint finalUser2 = rewardToken1.balanceOf(user2);
        uint finalDust = rewardToken1.balanceOf(address(pool));

        assertEq(totalDeposited, finalUser1 + finalUser2 + finalDust, "Final total matches");
        assertLt(finalDust, 2, "Final dust < 2 wei");
    }

    /// @notice Tests that dust doesn't grow unbounded over time
    /// @dev Even with worst-case rounding, dust is automatically redistributed
    function testDustDoesNotAccumulateUnbounded() public {
        // Worst case: many users with odd deposit amounts
        _userDeposit(user1, 1 ether);
        _userDeposit(user2, 2 ether);
        _userDeposit(user3, 3 ether);
        // Total: 6 ether

        uint totalDeposited = 0;
        uint totalClaimed = 0;

        // 100 rounds of distribution
        for (uint i = 0; i < 100; i++) {
            _depositReward(address(rewardToken1), 100 ether);
            totalDeposited += 100 ether;
        }

        // Claim all rewards
        vm.prank(user1);
        pool.claimRewards();
        totalClaimed += rewardToken1.balanceOf(user1);

        vm.prank(user2);
        pool.claimRewards();
        totalClaimed += rewardToken1.balanceOf(user2);

        vm.prank(user3);
        pool.claimRewards();
        totalClaimed += rewardToken1.balanceOf(user3);

        uint poolBalance = rewardToken1.balanceOf(address(pool));

        // Dust should be minimal (< 1% of one distribution round)
        assertLt(poolBalance, 1 ether, "Dust should not accumulate significantly");

        // Almost all tokens should be distributed
        assertGt(totalClaimed, totalDeposited - 1 ether, "Most tokens should be claimed");
    }

    /// @notice Tests exact accounting: every wei is tracked
    /// @dev Total deposited = Total claimed + Pool balance (should always hold)
    function testExactDustAccounting() public {
        _userDeposit(user1, 1 ether);
        _userDeposit(user2, 2 ether);

        uint totalDeposited = 0;

        // Deposit and claim multiple times
        for (uint i = 0; i < 5; i++) {
            uint depositAmount = 77 ether; // Odd number to create dust
            _depositReward(address(rewardToken1), depositAmount);
            totalDeposited += depositAmount;

            vm.prank(user1);
            pool.claimRewards();
            vm.prank(user2);
            pool.claimRewards();
        }

        uint user1Balance = rewardToken1.balanceOf(user1);
        uint user2Balance = rewardToken1.balanceOf(user2);
        uint poolBalance = rewardToken1.balanceOf(address(pool));

        // Exact accounting: deposited = claimed + remaining
        assertEq(
            totalDeposited,
            user1Balance + user2Balance + poolBalance,
            "Total deposited must equal total claimed + pool balance"
        );
    }

    // ==================== Zero deposit reward handling tests ====================

    function testRewardDepositWhenNoDepositors() public {
        // No one has depositd yet
        assertEq(pool.totalDeposited(), 0, "No depositrs initially");

        // Deposit rewards when totalDeposited = 0 (rewardToken1 is already registered in setup)
        _depositReward(address(rewardToken1), 100 ether);

        // Trigger a deposit to cause sync (this will make the contract recognize the deposit)
        _userDeposit(user1, 1 ether);

        // Now the rewards should be marked as withdrawable
        uint withdrawable = pool.getReclaimableAmount(rewardToken1);
        assertEq(withdrawable, 100 ether, "Rewards deposited when deposit=0 should be withdrawable");

        // Owner can withdraw via CrossGameReward
        uint ownerBalanceBefore = rewardToken1.balanceOf(owner);
        crossGameReward.reclaimFromPool(1, rewardToken1, owner);
        assertEq(rewardToken1.balanceOf(owner) - ownerBalanceBefore, 100 ether, "Owner withdrew unallocated rewards");

        // User should not have rewards (deposited before their deposit)
        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "User should not have rewards from pre-deposit deposit");
    }

    function testRewardDepositBeforeAndAfterDeposit() public {
        // Deposit before any deposit
        _depositReward(address(rewardToken1), 50 ether);

        // User deposits (this triggers sync and marks the 50 ether as withdrawable)
        _userDeposit(user1, 10 ether);

        // Check withdrawable amount
        assertEq(pool.getReclaimableAmount(rewardToken1), 50 ether, "50 ether withdrawable");

        // Deposit more rewards after deposit
        _depositReward(address(rewardToken1), 100 ether);

        // User should only get the 100 ether deposited after deposit
        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 100 ether, 100, "User gets only post-deposit rewards");

        // Initial 50 ether should still be withdrawable
        assertEq(pool.getReclaimableAmount(rewardToken1), 50 ether, "Initial 50 ether still withdrawable");
    }

    function testWithdrawZeroDepositRewards() public {
        // Deposit when no depositrs
        _depositReward(address(rewardToken1), 200 ether);

        // Deposit to trigger sync
        _userDeposit(user1, 1 ether);

        // Verify withdrawable amount
        assertEq(pool.getReclaimableAmount(rewardToken1), 200 ether, "200 ether withdrawable");

        // Withdraw via CrossGameReward
        uint ownerBalanceBefore = rewardToken1.balanceOf(owner);
        crossGameReward.reclaimFromPool(1, rewardToken1, owner);
        assertEq(rewardToken1.balanceOf(owner) - ownerBalanceBefore, 200 ether, "Withdrew all");

        // No more withdrawable
        assertEq(pool.getReclaimableAmount(rewardToken1), 0, "Nothing left to withdraw");
    }

    function testCannotWithdrawAllocatedRewards() public {
        // User deposits first
        _userDeposit(user1, 10 ether);

        // Deposit rewards (will be allocated to user1)
        _depositReward(address(rewardToken1), 100 ether);

        // No withdrawable amount since rewards are allocated
        assertEq(pool.getReclaimableAmount(rewardToken1), 0, "No withdrawable amount");

        // Cannot withdraw
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPool.CGRPNoReclaimableAmount.selector, address(rewardToken1)));
        crossGameReward.reclaimFromPool(1, rewardToken1, owner);
    }

    function testMultipleUsersAfterZeroDepositDeposit() public {
        // Deposit when no depositrs
        _depositReward(address(rewardToken1), 100 ether);

        // User1 deposits (triggers sync, marks 100 ether as withdrawable)
        _userDeposit(user1, 10 ether);

        // Initial 100 ether should be withdrawable
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether, "Initial deposit withdrawable");

        // Deposit more rewards
        _depositReward(address(rewardToken1), 200 ether);

        // User2 deposits
        _userDeposit(user2, 10 ether);

        // Deposit more rewards
        _depositReward(address(rewardToken1), 300 ether);

        // User1 should get: 200 (alone) + 150 (half of 300) = 350
        // User2 should get: 150 (half of 300)
        (, uint[] memory rewards1) = pool.pendingRewards(user1);
        (, uint[] memory rewards2) = pool.pendingRewards(user2);

        assertApproxEqAbs(rewards1[0], 350 ether, 100, "User1 rewards");
        assertApproxEqAbs(rewards2[0], 150 ether, 100, "User2 rewards");

        // Initial 100 ether still withdrawable
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether, "Initial deposit still withdrawable");
    }
}
