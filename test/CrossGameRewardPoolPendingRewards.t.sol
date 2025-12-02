// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {CrossGameRewardPool} from "../src/CrossGameRewardPool.sol";
import {CrossGameRewardPoolBase} from "./base/CrossGameRewardPoolBase.t.sol";

/**
 * @title CrossGameRewardPoolPendingRewardsTest
 * @notice Tests for pendingRewards() accuracy with reclaimableAmount logic
 * @dev Validates that _calculatePendingReward matches _syncReward logic
 */
contract CrossGameRewardPoolPendingRewardsTest is CrossGameRewardPoolBase {
    /**
     * @notice Test that pendingRewards correctly excludes reclaimableAmount
     * @dev This test validates the fix for _calculatePendingReward logic
     */
    function testPendingRewardsExcludesWithdrawableAmount() public {
        // 1. Deposit 100 ether when pool is empty (becomes withdrawable)
        _depositReward(address(rewardToken1), 100 ether);

        // 2. User deposits (triggers sync)
        _userDeposit(user1, 10 ether);

        // Verify withdrawable amount
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether, "100 ether withdrawable");

        // 3. User should have 0 pending rewards (100 ether is withdrawable, not distributed)
        (, uint[] memory pending) = pool.pendingRewards(user1);
        assertEq(pending[0], 0, "No rewards for user1 - all is withdrawable");

        // 4. Deposit 50 more ether (should be distributed)
        _depositReward(address(rewardToken1), 50 ether);

        // 5. Check pending rewards - should be 50 ether
        (, pending) = pool.pendingRewards(user1);
        assertApproxEqAbs(pending[0], 50 ether, 100, "User1 should have 50 ether pending");

        // 6. Verify withdrawable amount unchanged
        assertEq(pool.getReclaimableAmount(rewardToken1), 100 ether, "Withdrawable still 100 ether");
    }

    /**
     * @notice Test pendingRewards with multiple users and reclaimableAmount
     */
    function testPendingRewardsWithMultipleUsersAndWithdrawable() public {
        // 1. Deposit when empty
        _depositReward(address(rewardToken1), 200 ether);

        // 2. User1 deposits
        _userDeposit(user1, 10 ether);
        assertEq(pool.getReclaimableAmount(rewardToken1), 200 ether);

        // 3. Deposit more (will be distributed)
        _depositReward(address(rewardToken1), 100 ether);

        // 4. User2 deposits
        _userDeposit(user2, 10 ether);

        // 5. Check pending rewards before claim
        (, uint[] memory pending1) = pool.pendingRewards(user1);
        (, uint[] memory pending2) = pool.pendingRewards(user2);

        // User1: got all 100 ether
        assertApproxEqAbs(pending1[0], 100 ether, 100, "User1: 100 ether");
        // User2: got nothing yet
        assertEq(pending2[0], 0, "User2: 0 ether");

        // 6. Deposit more
        _depositReward(address(rewardToken1), 50 ether);

        // 7. Check pending again - should split 50 ether
        (, pending1) = pool.pendingRewards(user1);
        (, pending2) = pool.pendingRewards(user2);

        assertApproxEqAbs(pending1[0], 125 ether, 100, "User1: 100 + 25 = 125 ether");
        assertApproxEqAbs(pending2[0], 25 ether, 100, "User2: 25 ether");

        // 8. Withdrawable unchanged
        assertEq(pool.getReclaimableAmount(rewardToken1), 200 ether, "Withdrawable still 200 ether");
    }

    /**
     * @notice Test pendingRewards after partial withdraw of reclaimableAmount
     */
    function testPendingRewardsAfterPartialTokensReclaimed() public {
        // 1. Setup: 100 withdrawable
        _depositReward(address(rewardToken1), 100 ether);
        _userDeposit(user1, 10 ether);

        // 2. Add distributed rewards
        _depositReward(address(rewardToken1), 50 ether);

        // 3. Check pending (need to trigger sync by depositing more)
        _userDeposit(user1, 1 ether); // Trigger sync

        (, uint[] memory pending) = pool.pendingRewards(user1);
        assertApproxEqAbs(pending[0], 50 ether, 100, "50 ether pending");

        // 4. Owner withdraws reclaimableAmount
        uint withdrawable = pool.getReclaimableAmount(rewardToken1);
        assertEq(withdrawable, 100 ether, "100 ether withdrawable before withdraw");

        // After owner withdraws, pending should still be accurate
        crossGameReward.reclaimFromPool(1, rewardToken1, owner);

        // User's pending should be unchanged (they get distributed, not withdrawable)
        (, pending) = pool.pendingRewards(user1);
        assertApproxEqAbs(pending[0], 50 ether, 100, "User pending unchanged after owner withdraw");
    }

    /**
     * @notice Test pendingRewards consistency before and after claim
     */
    function testPendingRewardsConsistencyWithClaim() public {
        // Setup
        _depositReward(address(rewardToken1), 50 ether);
        _userDeposit(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Check pending
        (, uint[] memory pending) = pool.pendingRewards(user1);
        uint expectedAmount = pending[0];
        assertApproxEqAbs(expectedAmount, 100 ether, 100, "100 ether pending");

        // Claim
        vm.prank(user1);
        pool.claimRewards();

        // Verify claimed amount matches pending
        assertApproxEqAbs(rewardToken1.balanceOf(user1), expectedAmount, 100, "Claimed matches pending");

        // Pending should now be 0
        (, pending) = pool.pendingRewards(user1);
        assertEq(pending[0], 0, "No pending after claim");
    }

    /**
     * @notice Test pendingRewards with removed token
     */
    function testPendingRewardsWithRemovedToken() public {
        // 1. Setup with withdrawable
        _depositReward(address(rewardToken1), 100 ether);
        _userDeposit(user1, 10 ether);

        // 2. Add distributed
        _depositReward(address(rewardToken1), 50 ether);

        // 3. Trigger sync by depositing
        _userDeposit(user1, 1 ether);

        // Check pending before removal
        (, uint[] memory pendingBefore) = pool.pendingRewards(user1);
        assertApproxEqAbs(pendingBefore[0], 50 ether, 100, "50 ether pending before removal");

        // 4. Claim first (to store rewards)
        vm.prank(user1);
        pool.claimRewards();

        // 5. Deposit more and remove token
        _depositReward(address(rewardToken1), 30 ether);
        _userDeposit(user1, 1 ether); // Trigger sync

        crossGameReward.removeRewardToken(1, rewardToken1);

        // 6. pendingRewards will only show active tokens now
        // But we can check the user's stored rewards
        (, uint storedRewards) = pool.userRewards(user1, rewardToken1);
        assertApproxEqAbs(storedRewards, 30 ether, 100, "30 ether stored after removal");

        // 7. User can still claim removed token
        vm.prank(user1);
        pool.claimReward(rewardToken1);

        // Total claimed: 50 (first) + 30 (second) = 80 ether
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 80 ether, 200, "Claimed all rewards");
    }

    /**
     * @notice Test M-01: pendingReward should NOT include deposits made after token removal
     * @dev Validates that removed tokens do not include new deposits in pending calculation
     */
    function testPendingRewardDoesNotIncludePostRemovalDeposits() public {
        // 1. User deposits
        _userDeposit(user1, 10 ether);

        // 2. Deposit initial rewards
        _depositReward(address(rewardToken1), 100 ether);

        // 3. Trigger sync and check pending
        _userDeposit(user2, 10 ether);
        uint pendingBefore = pool.pendingReward(user1, rewardToken1);
        assertApproxEqAbs(pendingBefore, 100 ether, 100, "User1 has 100 ether pending");

        // 4. Remove the reward token
        crossGameReward.removeRewardToken(1, rewardToken1);

        // 5. Add MORE rewards after removal
        rewardToken1.mint(owner, 200 ether);
        rewardToken1.transfer(address(pool), 200 ether);

        // 6. Check pendingReward - should NOT include the 200 ether deposited after removal
        uint pendingAfter = pool.pendingReward(user1, rewardToken1);
        assertApproxEqAbs(
            pendingAfter, 100 ether, 100, "User1 still has only 100 ether - post-removal deposits excluded"
        );

        // 7. Verify actual claim matches the pending calculation
        uint balanceBefore = rewardToken1.balanceOf(user1);
        vm.prank(user1);
        pool.claimReward(rewardToken1);
        uint claimed = rewardToken1.balanceOf(user1) - balanceBefore;

        assertApproxEqAbs(claimed, pendingAfter, 100, "Claimed amount matches pendingReward");
        assertApproxEqAbs(claimed, 100 ether, 100, "Only original 100 ether claimed, not the 200 added after removal");
    }

    /**
     * @notice Test M-02: getRemovedTokenRewards should calculate actual pending rewards
     * @dev Validates that removed token rewards are accurately calculated even without _updateReward
     */
    function testGetRemovedTokenRewardsCalculatesActualPending() public {
        // 1. Two users deposit to establish initial state
        _userDeposit(user1, 10 ether);
        _userDeposit(user2, 10 ether); // totalDeposited = 20 ether

        // 2. Deposit rewards (will be distributed 50/50)
        _depositReward(address(rewardToken1), 100 ether);

        // 3. Trigger sync to update rewardPerTokenStored
        // Use another small deposit to trigger sync without affecting distribution much
        _userDeposit(user1, 1 ether); // This triggers sync

        // 4. Remove the reward token BEFORE user2's rewards are updated
        // At this point: user1 has claimed/updated, but user2 has not
        crossGameReward.removeRewardToken(1, rewardToken1);

        // 5. Check getRemovedTokenRewards for user2 - should return actual pending rewards
        (address[] memory tokens, uint[] memory rewards) = pool.getRemovedTokenRewards(user2);

        assertEq(tokens.length, 1, "Should have 1 removed token");
        assertEq(tokens[0], address(rewardToken1), "Should be rewardToken1");

        // User2 should get approximately 50% of 100 ether = 50 ether (before user1's 1 ether deposit)
        // More precisely: 10/20 * 100 = 50 ether
        assertApproxEqAbs(
            rewards[0], 50 ether, 1 ether, "Should calculate actual pending rewards including unsettled amount"
        );

        // 6. Verify actual claim matches the reported amount
        uint balanceBefore = rewardToken1.balanceOf(user2);
        vm.prank(user2);
        pool.claimReward(rewardToken1);
        uint claimed = rewardToken1.balanceOf(user2) - balanceBefore;

        assertApproxEqAbs(claimed, rewards[0], 100, "Claimed amount should match getRemovedTokenRewards");
    }

    /**
     * @notice Test M-02: getRemovedTokenRewards with zero depositd should return stored rewards
     */
    function testGetRemovedTokenRewardsAfterWithdraw() public {
        // 1. User deposits and earns rewards
        _userDeposit(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // 2. Trigger sync and update
        vm.prank(user1);
        pool.claimRewards(); // This updates rewards

        // 3. Remove token
        crossGameReward.removeRewardToken(1, rewardToken1);

        // 4. Add more rewards after removal (should not affect user)
        _depositReward(address(rewardToken1), 50 ether);

        // 5. User withdraws (balance becomes 0)
        vm.prank(user1);
        pool.withdraw(0);

        // 6. getRemovedTokenRewards should return 0 since user has no deposit and already claimed
        (, uint[] memory rewards) = pool.getRemovedTokenRewards(user1);

        assertEq(rewards[0], 0, "Should return 0 for user with no deposit and no pending rewards");
    }

    /**
     * @notice Stress test: many deposits with withdrawable
     */
    function testPendingRewardsWithManyDeposits() public {
        // Initial withdrawable deposit
        _depositReward(address(rewardToken1), 1000 ether);
        _userDeposit(user1, 100 ether);

        uint totalExpected = 0;

        // Many small deposits
        for (uint i = 1; i <= 10; i++) {
            _depositReward(address(rewardToken1), i * 10 ether);
            totalExpected += i * 10 ether;

            // Check pending matches expected
            (, uint[] memory pending) = pool.pendingRewards(user1);
            assertApproxEqAbs(pending[0], totalExpected, 1000, "Pending accumulates correctly");
        }

        // Withdrawable should be unchanged
        assertEq(pool.getReclaimableAmount(rewardToken1), 1000 ether, "Withdrawable unchanged");

        // Total distributed: 10+20+30+...+100 = 550 ether
        assertApproxEqAbs(totalExpected, 550 ether, 10, "Total 550 ether distributed");
    }
}
