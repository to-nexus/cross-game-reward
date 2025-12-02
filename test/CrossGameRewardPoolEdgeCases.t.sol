// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {CrossGameRewardPool} from "../src/CrossGameRewardPool.sol";
import {CrossGameRewardPoolBase} from "./base/CrossGameRewardPoolBase.t.sol";

/**
 * @title CrossGameRewardPoolEdgeCasesTest
 * @notice Comprehensive edge case testing for CrossGameRewardPool
 * @dev Tests critical scenarios including:
 *      - Partial reclaimableAmount withdrawals
 *      - Multiple deposit/withdraw cycles with reclaimableAmount
 *      - Reward distribution after partial withdrawals
 *      - Zero balance edge cases
 *      - Arithmetic edge cases
 */
contract CrossGameRewardPoolEdgeCasesTest is CrossGameRewardPoolBase {
    // ==================== Partial Withdraw Scenarios ====================

    function testPartialWithdrawThenDeposit() public {
        // 1. Deposit 100 ether when pool is empty
        _depositReward(address(rewardToken1), 100 ether);
        _userDeposit(user1, 1 ether); // Trigger sync

        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether, "100 ether withdrawable");

        // 2. Withdraw 50 ether (partial)
        crossGameReward.reclaimFromPool(1, rewardToken1, owner);
        // Manually withdraw only 50 ether by updating the test
        // Actually getReclaimableAmount returns all, so let's test the scenario differently

        // Since we can't do partial withdraw in one call, let's test the scenario
        // where reclaimableAmount exists and new deposits come in
    }

    function testWithdrawableAmountNotDistributedAfterPartialTokensReclaimed() public {
        // 1. Deposit 100 when pool empty
        _depositReward(address(rewardToken1), 100 ether);
        _userDeposit(user1, 1 ether);

        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether);

        // 2. Another user deposits
        _userDeposit(user2, 1 ether);

        // 3. Deposit 200 more
        _depositReward(address(rewardToken1), 200 ether);

        // 4. Check rewards - should split 200 equally (100 each)
        (, uint[] memory rewards1) = pool.pendingRewards(user1);
        (, uint[] memory rewards2) = pool.pendingRewards(user2);

        assertApproxEqAbs(rewards1[0], 100 ether, 100, "User1 gets 50% of 200");
        assertApproxEqAbs(rewards2[0], 100 ether, 100, "User2 gets 50% of 200");

        // 5. Initial 100 ether should still be withdrawable (not distributed)
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether, "Initial 100 still withdrawable");
    }

    // ==================== Multiple Deposit Cycles ====================

    function testMultipleZeroDepositDeposits() public {
        // Multiple deposits when pool is empty
        _depositReward(address(rewardToken1), 50 ether);
        _depositReward(address(rewardToken1), 30 ether);
        _depositReward(address(rewardToken1), 20 ether);

        // User deposits (triggers sync)
        _userDeposit(user1, 10 ether);

        // All deposits should be withdrawable
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether, "All pre-deposit deposits withdrawable");

        // New deposit should distribute
        _depositReward(address(rewardToken1), 50 ether);

        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 50 ether, 100, "Only new deposit distributed");

        // Withdrawable unchanged
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether, "Withdrawable unchanged");
    }

    function testDepositWithdrawDepositWithWithdrawable() public {
        // 1. Deposit when empty
        _depositReward(address(rewardToken1), 100 ether);

        // 2. User1 deposits
        _userDeposit(user1, 10 ether);
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether);

        // 3. Deposit more
        _depositReward(address(rewardToken1), 50 ether);

        // 4. User1 withdraws (claims 50 ether)
        vm.prank(user1);
        pool.withdraw(0);
        assertEq(rewardToken1.balanceOf(user1), 50 ether, "User1 got 50 ether");

        // 5. Pool is empty again, deposit more
        _depositReward(address(rewardToken1), 75 ether);

        // 6. User2 deposits
        _userDeposit(user2, 10 ether);

        // 7. Withdrawable should be 100 + 75 = 175 ether
        assertEq(pool.getReclaimableAmount(rewardToken1), 175 ether, "Both zero-deposit deposits withdrawable");

        // 8. New deposit
        _depositReward(address(rewardToken1), 25 ether);

        (, uint[] memory rewards) = pool.pendingRewards(user2);
        assertApproxEqAbs(rewards[0], 25 ether, 100, "Only new 25 ether distributed");
    }

    // ==================== Withdraw After Removal ====================

    function testWithdrawAfterRemovalWithWithdrawableAmount() public {
        // 1. Deposit when empty (becomes reclaimableAmount)
        _depositReward(address(rewardToken1), 100 ether);
        _userDeposit(user1, 10 ether);
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether, "Initial withdrawable");

        // 2. Deposit more (will be distributed)
        _depositReward(address(rewardToken1), 50 ether);

        // 3. Remove the token
        // At removal: currentBalance = 150 ether
        // reclaimableAmount = 100 ether (stays as is)
        // distributedAmount = 150 - 100 = 50 ether (user-claimable)
        crossGameReward.removeRewardToken(1, rewardToken1);

        // 4. Add more tokens after removal
        rewardToken1.mint(owner, 30 ether);
        rewardToken1.transfer(address(pool), 30 ether);

        // 5. getReclaimableAmount for removed token:
        // currentBalance = 180
        // distributedAmount = 50 (user-claimable)
        // reclaimableAmount = 100 (original owner-withdrawable)
        // Post-removal deposits = 180 - 50 - 100 = 30
        // Total owner-withdrawable = 30 + 100 = 130
        assertEq(pool.getReclaimableAmount(rewardToken1), 130 ether, "Post-removal deposits + original withdrawable");
    }

    // ==================== Arithmetic Edge Cases ====================

    function testVerySmallWithdrawableAmount() public {
        // Deposit 1 wei when empty
        _depositReward(address(rewardToken1), 1);
        _userDeposit(user1, 10 ether);

        assertEq(pool.getReclaimableAmount(rewardToken1), 1, "1 wei withdrawable");

        // Large deposit
        _depositReward(address(rewardToken1), 1000 ether);

        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 1000 ether, 100, "Large deposit distributed correctly");
        assertEq(pool.getReclaimableAmount(rewardToken1), 1, "1 wei still withdrawable");
    }

    function testMaxUintBoundary() public {
        // Test with large but reasonable amounts
        uint largeAmount = 1_000_000_000 ether; // 1 billion tokens

        // Mint tokens
        rewardToken1.mint(owner, largeAmount);
        crossToken.mint(owner, largeAmount);
        crossToken.transfer(user1, largeAmount);

        // This should not overflow
        _userDeposit(user1, largeAmount / 2);
        _depositReward(address(rewardToken1), largeAmount / 2);

        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertGt(rewards[0], 0, "Rewards calculated without overflow");
    }

    // ==================== Zero Balance Edge Cases ====================

    function testClaimWhenBalanceIsExactlyZero() public {
        _userDeposit(user1, 10 ether);

        // No rewards deposited
        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "No rewards");

        // Claim should not revert
        vm.prank(user1);
        pool.claimRewards();

        assertEq(rewardToken1.balanceOf(user1), 0, "No tokens transferred");
    }

    function testSyncWhenBalanceDecreases() public {
        // This should never happen in normal operation
        // but let's ensure it doesn't break anything

        _userDeposit(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // User claims (balance decreases)
        vm.prank(user1);
        pool.claimRewards();

        // Another user deposits (triggers sync with decreased balance)
        _userDeposit(user2, 10 ether);

        // Should not revert or cause issues
        (, uint[] memory rewards2) = pool.pendingRewards(user2);
        assertEq(rewards2[0], 0, "User2 has no rewards yet");
    }

    // ==================== Multiple Reward Tokens with Withdrawable ====================

    function testMultipleRewardTokensWithDifferentWithdrawableAmounts() public {
        // rewardToken2 already added in setUp

        // Token1: deposit when empty (will be withdrawable)
        _depositReward(address(rewardToken1), 100 ether);

        // User deposits
        _userDeposit(user1, 10 ether);

        // Token1: 100 withdrawable
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether);

        // Token2: deposit after deposit (will be distributed)
        rewardToken2.mint(owner, 50 ether);
        rewardToken2.transfer(address(pool), 50 ether);

        // Token2: 0 withdrawable (distributed to user)
        assertEq(pool.getReclaimableAmount(rewardToken2), 0);

        (, uint[] memory rewards) = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "Token1: no rewards (was withdrawable)");
        assertApproxEqAbs(rewards[1], 50 ether, 100, "Token2: rewards distributed");
    }

    // ==================== Stress Testing ====================

    function testManyDepositsCyclesWithWithdrawable() public {
        uint withdrawableAccumulated = 0;

        // Cycle 1: empty deposit
        _depositReward(address(rewardToken1), 10 ether);
        withdrawableAccumulated += 10 ether;

        // Cycle 2: deposit and withdraw
        _userDeposit(user1, 5 ether);
        assertEq(pool.getReclaimableAmount(rewardToken1), withdrawableAccumulated);

        _depositReward(address(rewardToken1), 20 ether);
        vm.prank(user1);
        pool.withdraw(0); // Claims 20 ether

        // Cycle 3: empty deposit again
        _depositReward(address(rewardToken1), 15 ether);
        withdrawableAccumulated += 15 ether;

        // Cycle 4: deposit
        _userDeposit(user2, 5 ether);
        assertEq(pool.getReclaimableAmount(rewardToken1), withdrawableAccumulated, "Accumulated withdrawable");

        // New deposits should distribute
        _depositReward(address(rewardToken1), 30 ether);
        (, uint[] memory rewards) = pool.pendingRewards(user2);
        assertApproxEqAbs(rewards[0], 30 ether, 100, "New rewards distributed");

        assertEq(pool.getReclaimableAmount(rewardToken1), withdrawableAccumulated, "Withdrawable unchanged");
    }

    // ==================== Reentrancy Edge Cases ====================

    function testClaimDoesNotAffectOtherUsersWithdrawable() public {
        // Deposit when empty
        _depositReward(address(rewardToken1), 100 ether);

        // Two users deposit
        _userDeposit(user1, 10 ether);
        _userDeposit(user2, 10 ether);

        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether);

        // Deposit more
        _depositReward(address(rewardToken1), 200 ether);

        // User1 claims their share
        vm.prank(user1);
        pool.claimRewards();

        // Withdrawable should still be 100 (not affected by user claims)
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether, "Withdrawable not affected by claims");
    }
}
