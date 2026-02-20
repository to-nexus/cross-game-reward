// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "./base/CrossGameRewardPoolV2Base.t.sol";

/**
 * @title CrossGameRewardPoolV2RoundTest
 * @notice Tests for V2 pool round-based reward distribution
 */
contract CrossGameRewardPoolV2RoundTest is CrossGameRewardPoolV2Base {
    // ==================== Round Creation Tests ====================

    function test_CreateRound_Success() public {
        uint256 amount = 150000 ether;
        uint256 startBlock = block.number + 100;
        uint256 duration = 1000;

        vm.startPrank(sponsor);
        crossdToken.approve(address(poolV2), amount);
        uint256 roundId = poolV2.createRound(amount, startBlock, duration);
        vm.stopPrank();

        assertEq(roundId, 1);
        assertEq(poolV2.nextRoundId(), 2);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        assertEq(round.roundId, 1);
        assertEq(round.creator, sponsor);
        assertEq(round.totalReward, amount);
        assertEq(round.startBlock, startBlock);
        assertEq(round.endBlock, startBlock + duration);
        assertEq(round.rewardPerBlock, (amount / duration / 1e10) * 1e10);
        assertEq(round.remainderReward, amount - round.rewardPerBlock * duration);
        assertEq(round.lastRewardBlock, startBlock);
        assertFalse(round.isCancelled);
    }

    function test_CreateRound_MultipleRounds() public {
        _createRound(100000 ether, 100, 1000);
        _createRound(200000 ether, 200, 2000);
        _createRound(300000 ether, 300, 3000);

        assertEq(poolV2.nextRoundId(), 4);
        assertEq(poolV2.getActiveRoundCount(), 3);
    }

    function test_CreateRound_OnlySponsor() public {
        vm.startPrank(user1);
        crossdToken.approve(address(poolV2), 1000 ether);

        vm.expectRevert();
        poolV2.createRound(1000 ether, block.number + 100, 1000);
        vm.stopPrank();
    }

    function test_CreateRound_InvalidStartBlock() public {
        vm.startPrank(sponsor);
        crossdToken.approve(address(poolV2), 1000 ether);

        // Start block in the past
        vm.expectRevert(
            abi.encodeWithSelector(CrossGameRewardPoolV2.CGRP2InvalidStartBlock.selector, block.number, block.number)
        );
        poolV2.createRound(1000 ether, block.number, 1000);
        vm.stopPrank();
    }

    function test_CreateRound_ZeroAmount() public {
        vm.startPrank(sponsor);

        vm.expectRevert(CrossGameRewardPoolV2.CGRP2CanNotZeroValue.selector);
        poolV2.createRound(0, block.number + 100, 1000);
        vm.stopPrank();
    }

    function test_CreateRound_ZeroDuration() public {
        vm.startPrank(sponsor);
        crossdToken.approve(address(poolV2), 1000 ether);

        vm.expectRevert(CrossGameRewardPoolV2.CGRP2InvalidDuration.selector);
        poolV2.createRound(1000 ether, block.number + 100, 0);
        vm.stopPrank();
    }

    // ==================== Round Cancellation Tests ====================

    function test_CancelRound_Success() public {
        uint256 amount = 100000 ether;
        uint256 roundId = _createRound(amount, 100, 1000);

        uint256 sponsorBalanceBefore = crossdToken.balanceOf(sponsor);

        vm.prank(sponsor);
        poolV2.cancelRound(roundId);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        assertTrue(round.isCancelled);
        assertEq(poolV2.getActiveRoundCount(), 0);
        assertEq(crossdToken.balanceOf(sponsor), sponsorBalanceBefore + amount);
    }

    function test_CancelRound_AlreadyStarted() public {
        uint256 roundId = _createRound(100000 ether, 10, 1000);

        // Advance past start block
        _advanceBlocks(20);

        vm.prank(sponsor);
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPoolV2.CGRP2RoundAlreadyStarted.selector, roundId));
        poolV2.cancelRound(roundId);
    }

    function test_CancelRound_AlreadyCancelled() public {
        uint256 roundId = _createRound(100000 ether, 100, 1000);

        vm.prank(sponsor);
        poolV2.cancelRound(roundId);

        vm.prank(sponsor);
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPoolV2.CGRP2RoundAlreadyCancelled.selector, roundId));
        poolV2.cancelRound(roundId);
    }

    function test_CancelRound_NotFound() public {
        vm.prank(sponsor);
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPoolV2.CGRP2RoundNotFound.selector, 999));
        poolV2.cancelRound(999);
    }

    function test_CancelRound_OnlyCreator() public {
        uint256 roundId = _createRound(100000 ether, 100, 1000);

        // sponsor2 has SPONSOR_ROLE but is not the creator
        vm.prank(sponsor2);
        vm.expectRevert(
            abi.encodeWithSelector(
                CrossGameRewardPoolV2.CGRP2OnlyRoundCreator.selector, roundId, sponsor2, sponsor
            )
        );
        poolV2.cancelRound(roundId);

        // Original creator can cancel
        vm.prank(sponsor);
        poolV2.cancelRound(roundId);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        assertTrue(round.isCancelled);
    }

    function test_CancelRoundToRecipient_OnlyCreator() public {
        uint256 roundId = _createRound(100000 ether, 100, 1000);

        address recipient = address(0x9999);

        // sponsor2 cannot cancel to a recipient
        vm.prank(sponsor2);
        vm.expectRevert(
            abi.encodeWithSelector(
                CrossGameRewardPoolV2.CGRP2OnlyRoundCreator.selector, roundId, sponsor2, sponsor
            )
        );
        poolV2.cancelRoundToRecipient(roundId, recipient);

        // Creator can cancel to a custom recipient
        uint256 recipientBalBefore = crossdToken.balanceOf(recipient);

        vm.prank(sponsor);
        poolV2.cancelRoundToRecipient(roundId, recipient);

        assertEq(crossdToken.balanceOf(recipient), recipientBalBefore + 100000 ether);
    }

    function test_CreateRoundFromReserve_StoresCallerAsCreator() public {
        // sponsor2 creates a round pulling tokens from sponsor (reserve)
        vm.prank(sponsor);
        crossdToken.approve(address(poolV2), 100000 ether);

        vm.prank(sponsor2);
        uint256 roundId = poolV2.createRoundFromReserve(sponsor, 100000 ether, block.number + 100, 1000);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        assertEq(round.creator, sponsor2);

        // sponsor (the reserve, not the creator) cannot cancel
        vm.prank(sponsor);
        vm.expectRevert(
            abi.encodeWithSelector(
                CrossGameRewardPoolV2.CGRP2OnlyRoundCreator.selector, roundId, sponsor, sponsor2
            )
        );
        poolV2.cancelRound(roundId);

        // sponsor2 (the creator) can cancel
        vm.prank(sponsor2);
        poolV2.cancelRound(roundId);
        assertTrue(poolV2.getRound(roundId).isCancelled);
    }

    function test_CancelRound_AfterRoleRevoked() public {
        uint256 roundId = _createRound(100000 ether, 100, 1000);

        // Revoke sponsor's SPONSOR_ROLE
        crossGameReward.revokeSponsorRole(poolId, sponsor);

        // Creator can still cancel even without SPONSOR_ROLE
        uint256 balBefore = crossdToken.balanceOf(sponsor);

        vm.prank(sponsor);
        poolV2.cancelRound(roundId);

        assertEq(crossdToken.balanceOf(sponsor), balBefore + 100000 ether);
        assertTrue(poolV2.getRound(roundId).isCancelled);
    }

    function test_CancelRound_WhenPaused() public {
        uint256 roundId = _createRound(100000 ether, 100, 1000);

        // Pause the pool
        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Paused);

        // Creator can still cancel even when pool is paused
        uint256 balBefore = crossdToken.balanceOf(sponsor);

        vm.prank(sponsor);
        poolV2.cancelRound(roundId);

        assertEq(crossdToken.balanceOf(sponsor), balBefore + 100000 ether);
        assertTrue(poolV2.getRound(roundId).isCancelled);
    }

    // ==================== Reward Distribution Tests ====================

    function test_RewardDistribution_SingleUser() public {
        // Setup: User deposits before round starts
        _userDeposit(user1, 100 ether);

        // Create round: 10000 CROSSD over 100 blocks = 100 CROSSD/block
        uint256 roundId = _createRound(10000 ether, 10, 100);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        // Advance to start of round
        _advanceToBlock(round.startBlock);

        // Advance 50 blocks into the round
        _advanceBlocks(50);

        // Check pending rewards: 50 blocks * 100 CROSSD/block = 5000 CROSSD
        uint256 pending = _getPendingReward(user1);
        assertEq(pending, 5000 ether);
    }

    function test_RewardDistribution_MultipleUsers_EqualDeposits() public {
        // Two users deposit equal amounts
        _userDeposit(user1, 100 ether);
        _userDeposit(user2, 100 ether);

        // Create round: 10000 CROSSD over 100 blocks
        uint256 roundId = _createRound(10000 ether, 10, 100);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(50);

        // Each user should get 50% of 5000 = 2500 CROSSD
        assertEq(_getPendingReward(user1), 2500 ether);
        assertEq(_getPendingReward(user2), 2500 ether);
    }

    function test_RewardDistribution_MultipleUsers_DifferentDeposits() public {
        // user1: 100, user2: 300 (total: 400, ratio 25%:75%)
        _userDeposit(user1, 100 ether);
        _userDeposit(user2, 300 ether);

        // Create round: 10000 CROSSD over 100 blocks
        uint256 roundId = _createRound(10000 ether, 10, 100);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(100); // Complete round

        // user1: 25% of 10000 = 2500
        // user2: 75% of 10000 = 7500
        assertEq(_getPendingReward(user1), 2500 ether);
        assertEq(_getPendingReward(user2), 7500 ether);
    }

    function test_RewardDistribution_DepositAfterRoundStart() public {
        // Create round first
        uint256 roundId = _createRound(10000 ether, 10, 100);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        // Advance to middle of round
        _advanceToBlock(round.startBlock);
        _advanceBlocks(50);

        // User deposits in the middle
        _userDeposit(user1, 100 ether);

        // Advance to end
        _advanceBlocks(50);

        // User should only get rewards from block 60 onwards (50 blocks)
        // 50 blocks * 100 CROSSD/block = 5000 CROSSD
        uint256 pending = _getPendingReward(user1);
        assertEq(pending, 5000 ether);
    }

    function test_RewardDistribution_ClaimRewards() public {
        _userDeposit(user1, 100 ether);

        uint256 roundId = _createRound(10000 ether, 10, 100);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(50);

        uint256 balanceBefore = crossdToken.balanceOf(user1);

        vm.prank(user1);
        poolV2.claimRewards();

        assertEq(crossdToken.balanceOf(user1), balanceBefore + 5000 ether);
        assertEq(_getPendingReward(user1), 0);
    }

    function test_RewardDistribution_ClaimAndContinue() public {
        _userDeposit(user1, 100 ether);

        uint256 roundId = _createRound(10000 ether, 10, 100);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(50);

        // First claim
        vm.prank(user1);
        poolV2.claimRewards();
        assertEq(crossdToken.balanceOf(user1), 5000 ether);

        // Continue to accumulate
        _advanceBlocks(25);

        // Second claim
        vm.prank(user1);
        poolV2.claimRewards();
        assertEq(crossdToken.balanceOf(user1), 7500 ether);
    }

    // ==================== Multiple Rounds Tests ====================

    function test_MultipleRounds_Overlapping() public {
        _userDeposit(user1, 100 ether);

        // Round 1: 10000 CROSSD over 100 blocks starting at block 10
        _createRound(10000 ether, 10, 100);

        // Round 2: 20000 CROSSD over 100 blocks starting at block 50
        _createRound(20000 ether, 50, 100);

        // Advance to block 60 (both rounds active)
        _advanceBlocks(60);

        // Round 1: 50 blocks * 100 = 5000
        // Round 2: 10 blocks * 200 = 2000
        // Total: 7000
        uint256 pending = _getPendingReward(user1);
        assertEq(pending, 7000 ether);
    }

    function test_MultipleRounds_Sequential() public {
        _userDeposit(user1, 100 ether);

        // Round 1: ends at block 110
        _createRound(10000 ether, 10, 100);

        // Round 2: starts at block 200
        _createRound(20000 ether, 200, 100);

        // After round 1 completes
        _advanceBlocks(150);
        assertEq(_getPendingReward(user1), 10000 ether);

        // After round 2 starts and progresses
        _advanceBlocks(100);
        assertEq(_getPendingReward(user1), 10000 ether + 10000 ether); // Round 2: 50 blocks * 200
    }

    // ==================== Reclaimable Amount Tests ====================

    function test_Reclaimable_NoDepositors() public {
        // Create round with no depositors
        uint256 roundId = _createRound(10000 ether, 10, 100);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        _advanceToBlock(round.startBlock);

        // Advance through entire round with no depositors
        _advanceBlocks(100);

        // Trigger pool update by depositing
        _userDeposit(user1, 100 ether);

        // All rewards should be reclaimable
        assertEq(poolV2.reclaimableAmount(), 10000 ether);
    }

    function test_Reclaimable_PartialPeriod() public {
        // Create round
        uint256 roundId = _createRound(10000 ether, 10, 100);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        _advanceToBlock(round.startBlock);

        // Advance 50 blocks with no depositors
        _advanceBlocks(50);

        // User deposits
        _userDeposit(user1, 100 ether);

        // First 50 blocks should be reclaimable
        assertEq(poolV2.reclaimableAmount(), 5000 ether);

        // Advance remaining 50 blocks
        _advanceBlocks(50);

        // User should get the last 50 blocks
        assertEq(_getPendingReward(user1), 5000 ether);
    }

    function test_ReclaimTokens() public {
        // Create round with no depositors
        uint256 roundId = _createRound(10000 ether, 10, 100);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(100);

        // Trigger pool update
        _userDeposit(user1, 100 ether);

        // Reclaim tokens
        address recipient = address(0x9999);
        crossGameReward.reclaimFromPool(poolId, crossdToken, recipient);

        assertEq(crossdToken.balanceOf(recipient), 10000 ether);
        assertEq(poolV2.reclaimableAmount(), 0);
    }

    // ==================== Withdraw With Rewards Tests ====================

    function test_WithdrawClaimsRewards() public {
        _userDeposit(user1, 100 ether);

        uint256 roundId = _createRound(10000 ether, 10, 100);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(100);

        uint256 rewardBalanceBefore = crossdToken.balanceOf(user1);
        uint256 depositBalanceBefore = gameToken.balanceOf(user1);

        vm.prank(user1);
        poolV2.withdraw(0); // Withdraw all

        // Should receive both deposit and rewards
        assertEq(gameToken.balanceOf(user1), depositBalanceBefore + 100 ether);
        assertEq(crossdToken.balanceOf(user1), rewardBalanceBefore + 10000 ether);
    }

    // ==================== Round Query Tests ====================

    function test_GetActiveRounds() public {
        _createRound(10000 ether, 10, 100);
        _createRound(20000 ether, 20, 200);
        _createRound(30000 ether, 30, 300);

        ICrossGameRewardPoolV2.Round[] memory rounds = poolV2.getActiveRounds();
        assertEq(rounds.length, 3);
    }

    function test_GetActiveRoundIds() public {
        _createRound(10000 ether, 10, 100);
        _createRound(20000 ether, 20, 200);

        uint256[] memory ids = poolV2.getActiveRoundIds();
        assertEq(ids.length, 2);
        assertEq(ids[0], 1);
        assertEq(ids[1], 2);
    }

    function test_RoundRemovedAfterCompletion() public {
        _userDeposit(user1, 100 ether);

        uint256 roundId = _createRound(10000 ether, 10, 100);

        assertEq(poolV2.getActiveRoundCount(), 1);

        // Advance past round end
        _advanceBlocks(200);

        // Trigger pool update
        vm.prank(user1);
        poolV2.claimRewards();

        // Round should be removed from active list
        assertEq(poolV2.getActiveRoundCount(), 0);
    }

    // ==================== Audit Fix: Remainder / Dust Token Lock Tests ====================

    function test_CreateRound_RemainderNotLocked() public {
        // 5 ether / 3 blocks → truncated rewardPerBlock, remainder distributed in last block
        uint256 amount = 5 ether;
        uint256 duration = 3;

        uint256 devBalanceBefore = crossdToken.balanceOf(sponsor);

        vm.startPrank(sponsor);
        crossdToken.approve(address(poolV2), amount);
        uint256 roundId = poolV2.createRound(amount, block.number + 10, duration);
        vm.stopPrank();

        uint256 devBalanceAfter = crossdToken.balanceOf(sponsor);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        uint256 expectedRpb = (amount / duration / 1e10) * 1e10;
        uint256 expectedRemainder = amount - (expectedRpb * duration);

        assertEq(round.totalReward, amount);
        assertEq(round.rewardPerBlock, expectedRpb);
        assertEq(round.remainderReward, expectedRemainder);
        assertGt(expectedRemainder, 0);

        // Full amount is transferred from sponsor
        assertEq(devBalanceBefore - devBalanceAfter, amount);
        assertEq(crossdToken.balanceOf(address(poolV2)), amount);

        // Verify full distribution: deposit user, run through entire round
        _userDeposit(user1, 100 ether);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(duration);

        uint256 pending = _getPendingReward(user1);
        assertEq(pending, amount);

        // Claim and verify no tokens are stuck
        vm.prank(user1);
        poolV2.claimRewards();

        assertEq(crossdToken.balanceOf(address(poolV2)), 0);
    }

    function test_CreateRound_AmountLessThanDuration_Reverts() public {
        // amount < durationBlocks => rewardPerBlock = 0 => should revert
        uint256 amount = 99; // 99 wei
        uint256 duration = 100;

        vm.startPrank(sponsor);
        crossdToken.approve(address(poolV2), amount);

        vm.expectRevert(CrossGameRewardPoolV2.CGRP2RewardPerBlockZero.selector);
        poolV2.createRound(amount, block.number + 10, duration);
        vm.stopPrank();
    }

    function test_CreateRound_ExactDivision_NoRemainder() public {
        // 10000 CROSSD / 100 blocks = 100 CROSSD/block, no remainder
        uint256 amount = 10000 ether;
        uint256 duration = 100;

        uint256 devBalanceBefore = crossdToken.balanceOf(sponsor);

        vm.startPrank(sponsor);
        crossdToken.approve(address(poolV2), amount);
        uint256 roundId = poolV2.createRound(amount, block.number + 10, duration);
        vm.stopPrank();

        uint256 devBalanceAfter = crossdToken.balanceOf(sponsor);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        assertEq(round.totalReward, amount);
        assertEq(round.rewardPerBlock, 100 ether);
        assertEq(round.remainderReward, 0);
        assertEq(devBalanceBefore - devBalanceAfter, amount);
    }

    function test_CreateRound_LargeRemainder_Wei() public {
        // Edge case: 999 wei / 1000 blocks => rewardPerBlock = 0 => reverts
        vm.startPrank(sponsor);
        crossdToken.approve(address(poolV2), 999);

        vm.expectRevert(CrossGameRewardPoolV2.CGRP2RewardPerBlockZero.selector);
        poolV2.createRound(999, block.number + 10, 1000);
        vm.stopPrank();
    }

    function test_CreateRound_NonDivisibleAmount() public {
        // 7 ether / 3 blocks:
        // rewardPerBlock truncated to 1e10: (7e18 / 3 / 1e10) * 1e10 = 2333333330000000000
        // remainder = 7e18 - 2333333330000000000 * 3 = 10000000000
        uint256 amount = 7 ether;
        uint256 duration = 3;

        uint256 devBalanceBefore = crossdToken.balanceOf(sponsor);

        vm.startPrank(sponsor);
        crossdToken.approve(address(poolV2), amount);
        uint256 roundId = poolV2.createRound(amount, block.number + 10, duration);
        vm.stopPrank();

        uint256 devBalanceAfter = crossdToken.balanceOf(sponsor);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        uint256 expectedRpb = (amount / duration / 1e10) * 1e10;
        uint256 expectedRemainder = amount - (expectedRpb * duration);

        assertGt(expectedRpb, 0);
        assertGt(expectedRemainder, 0);
        assertEq(round.rewardPerBlock, expectedRpb);
        assertEq(round.totalReward, amount);
        assertEq(round.remainderReward, expectedRemainder);

        // Full amount transferred from sponsor
        assertEq(devBalanceBefore - devBalanceAfter, amount);
    }

    // ==================== Truncation & Remainder Edge Case Tests ====================

    function test_CreateRound_TruncatedRewardPerBlock() public {
        // 5 ether / 3 blocks:
        // raw = 1666666666666666666, truncated to 1e10 = 1666666660000000000
        uint256 amount = 5 ether;
        uint256 duration = 3;

        uint256 roundId = _createRound(amount, 10, duration);
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        uint256 expectedRpb = 1666666660000000000;
        assertEq(round.rewardPerBlock, expectedRpb);
        assertEq(round.rewardPerBlock % 1e10, 0);
    }

    function test_CreateRound_RemainderInLastBlock() public {
        // 5 ether / 3 blocks → remainder distributed in final block
        uint256 amount = 5 ether;
        uint256 duration = 3;

        _userDeposit(user1, 100 ether);

        uint256 roundId = _createRound(amount, 10, duration);
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        _advanceToBlock(round.startBlock);

        // After 2 blocks (not yet last block): 2 * rewardPerBlock
        _advanceBlocks(2);
        uint256 pendingMid = _getPendingReward(user1);
        assertEq(pendingMid, round.rewardPerBlock * 2);

        // After last block: totalReward fully distributed
        _advanceBlocks(1);
        uint256 pendingEnd = _getPendingReward(user1);
        assertEq(pendingEnd, amount);

        // Claim and verify zero dust
        vm.prank(user1);
        poolV2.claimRewards();
        assertEq(crossdToken.balanceOf(address(poolV2)), 0);
    }

    function test_CancelRound_FullRefundWithRemainder() public {
        // Non-divisible amount: cancel should refund full amount
        uint256 amount = 7 ether;
        uint256 duration = 3;

        uint256 balBefore = crossdToken.balanceOf(sponsor);

        uint256 roundId = _createRound(amount, 100, duration);

        uint256 balAfterCreate = crossdToken.balanceOf(sponsor);
        assertEq(balBefore - balAfterCreate, amount);

        vm.prank(sponsor);
        poolV2.cancelRound(roundId);

        assertEq(crossdToken.balanceOf(sponsor), balBefore);
    }

    function test_CreateRound_TruncationCausesRevert() public {
        // amount / durationBlocks < 1e10 after truncation → revert
        // 15e9 wei / 3 blocks = 5e9 per block → truncated to 0
        uint256 amount = 15e9;
        uint256 duration = 3;

        vm.startPrank(sponsor);
        crossdToken.approve(address(poolV2), amount);

        vm.expectRevert(CrossGameRewardPoolV2.CGRP2RewardPerBlockZero.selector);
        poolV2.createRound(amount, block.number + 10, duration);
        vm.stopPrank();
    }

    function test_RewardDistribution_RemainderMultipleUsers() public {
        // Two users with equal deposits; remainder shared proportionally at last block
        _userDeposit(user1, 100 ether);
        _userDeposit(user2, 100 ether);

        uint256 amount = 5 ether;
        uint256 duration = 3;

        uint256 roundId = _createRound(amount, 10, duration);
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(duration);

        uint256 pending1 = _getPendingReward(user1);
        uint256 pending2 = _getPendingReward(user2);

        // Each user gets 50% of total reward
        assertEq(pending1, pending2);
        assertEq(pending1 + pending2, amount);
    }

    function test_RewardDistribution_FullRoundWithRemainder() public {
        // Single user gets exactly totalReward (= original amount) after full round
        uint256 amount = 11 ether;
        uint256 duration = 7;

        _userDeposit(user1, 100 ether);

        uint256 roundId = _createRound(amount, 10, duration);
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        assertGt(round.remainderReward, 0);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(duration);

        assertEq(_getPendingReward(user1), amount);

        uint256 balBefore = crossdToken.balanceOf(user1);
        vm.prank(user1);
        poolV2.claimRewards();

        assertEq(crossdToken.balanceOf(user1), balBefore + amount);
        assertEq(crossdToken.balanceOf(address(poolV2)), 0);
    }

    function test_Reclaimable_WithRemainder() public {
        // No depositors: entire amount including remainder goes to reclaimable
        uint256 amount = 5 ether;
        uint256 duration = 3;

        uint256 roundId = _createRound(amount, 10, duration);
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        assertGt(round.remainderReward, 0);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(duration);

        // Trigger pool update
        _userDeposit(user1, 100 ether);

        assertEq(poolV2.reclaimableAmount(), amount);
    }

    // ==================== Audit Fix: claimReward Token Validation Tests ====================

    function test_ClaimReward_WrongToken_Reverts() public {
        _userDeposit(user1, 100 ether);

        IERC20 wrongToken = IERC20(address(gameToken)); // Use deposit token as wrong reward token

        vm.prank(user1);
        vm.expectRevert(
            abi.encodeWithSelector(
                CrossGameRewardPoolV2.CGRP2InvalidRewardToken.selector,
                address(wrongToken),
                address(crossdToken)
            )
        );
        poolV2.claimReward(wrongToken);
    }

    function test_ClaimRewardFor_WrongToken_Reverts() public {
        // Setup router
        crossGameReward.setRouter(address(this));

        _userDeposit(user1, 100 ether);

        IERC20 wrongToken = IERC20(address(gameToken));

        vm.expectRevert(
            abi.encodeWithSelector(
                CrossGameRewardPoolV2.CGRP2InvalidRewardToken.selector,
                address(wrongToken),
                address(crossdToken)
            )
        );
        poolV2.claimRewardFor(user1, wrongToken);
    }

    function test_ClaimReward_CorrectToken_Success() public {
        _userDeposit(user1, 100 ether);

        uint256 roundId = _createRound(10000 ether, 10, 100);
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(100);

        uint256 balanceBefore = crossdToken.balanceOf(user1);

        vm.prank(user1);
        poolV2.claimReward(crossdToken); // Correct token

        assertEq(crossdToken.balanceOf(user1), balanceBefore + 10000 ether);
    }

    // ==================== Emergency Sync Pagination Tests ====================

    function test_SyncRounds_PartialPagination() public {
        _userDeposit(user1, 100 ether);

        for (uint256 i = 0; i < 10; i++) {
            _createRound(10000 ether, 100, 1000);
        }
        assertEq(poolV2.getActiveRoundCount(), 10);

        _advanceBlocks(1200);

        vm.startPrank(sponsor);
        (uint256 processed, uint256 removed) = poolV2.syncRounds(5);
        assertEq(processed, 5);
        assertEq(removed, 5);
        assertEq(poolV2.getActiveRoundCount(), 5);

        (processed, removed) = poolV2.syncRounds(5);
        assertEq(processed, 5);
        assertEq(removed, 5);
        assertEq(poolV2.getActiveRoundCount(), 0);
        vm.stopPrank();
    }

    function test_SyncRounds_ZeroMeansAll() public {
        _userDeposit(user1, 100 ether);

        for (uint256 i = 0; i < 8; i++) {
            _createRound(10000 ether, 100, 1000);
        }
        assertEq(poolV2.getActiveRoundCount(), 8);

        _advanceBlocks(1200);

        vm.prank(sponsor);
        (uint256 processed, uint256 removed) = poolV2.syncRounds(0);
        assertEq(processed, 8);
        assertEq(removed, 8);
        assertEq(poolV2.getActiveRoundCount(), 0);
    }

    function test_SyncRounds_MaxGreaterThanActive() public {
        _userDeposit(user1, 100 ether);
        _createRound(10000 ether, 10, 100);
        _createRound(10000 ether, 10, 100);

        _advanceBlocks(200);

        vm.prank(sponsor);
        (uint256 processed,) = poolV2.syncRounds(100);
        assertEq(processed, 2);
        assertEq(poolV2.getActiveRoundCount(), 0);
    }

    function test_SyncRounds_EmptyActiveSet() public {
        vm.prank(sponsor);
        (uint256 processed, uint256 removed) = poolV2.syncRounds(10);
        assertEq(processed, 0);
        assertEq(removed, 0);
    }

    function test_SyncRounds_MixedStates() public {
        _userDeposit(user1, 100 ether);

        uint256 round1 = _createRound(10000 ether, 10, 100);
        _createRound(10000 ether, 500, 100);
        _createRound(10000 ether, 10, 100);

        vm.prank(sponsor);
        poolV2.cancelRound(round1);

        _advanceBlocks(200);

        // round1: cancelled (already removed from active set by cancelRound)
        // round2: not started yet (startBlock=current+500), stays
        // round3: completed (startBlock+10, duration 100), removed
        vm.prank(sponsor);
        (uint256 processed, uint256 removed) = poolV2.syncRounds(0);
        assertEq(processed, 2);
        assertEq(removed, 1);
        assertEq(poolV2.getActiveRoundCount(), 1);
    }

    function test_SyncRounds_EquivalenceWithUserPath() public {
        _userDeposit(user1, 100 ether);
        _userDeposit(user2, 100 ether);

        for (uint256 i = 0; i < 5; i++) {
            _createRound(10000 ether, 50, 200);
        }
        _advanceBlocks(300);

        uint256 snapshot = vm.snapshot();

        vm.prank(user1);
        poolV2.claimRewards();
        uint256 directReward = crossdToken.balanceOf(user1);
        uint256 directGlobalAcc = poolV2.globalAccRewardPerShare();

        vm.revertTo(snapshot);

        vm.prank(sponsor);
        poolV2.syncRounds(0);

        vm.prank(user1);
        poolV2.claimRewards();
        uint256 syncedReward = crossdToken.balanceOf(user1);
        uint256 syncedGlobalAcc = poolV2.globalAccRewardPerShare();

        assertEq(syncedGlobalAcc, directGlobalAcc);
        assertEq(syncedReward, directReward);
    }

    function test_SyncRounds_SteppedEquivalence() public {
        _userDeposit(user1, 100 ether);

        for (uint256 i = 0; i < 6; i++) {
            _createRound(10000 ether, 50, 200);
        }
        _advanceBlocks(300);

        uint256 snapshot = vm.snapshot();

        vm.prank(user1);
        poolV2.claimRewards();
        uint256 fullReward = crossdToken.balanceOf(user1);

        vm.revertTo(snapshot);

        vm.startPrank(sponsor);
        poolV2.syncRounds(2);
        poolV2.syncRounds(2);
        poolV2.syncRounds(2);
        vm.stopPrank();

        vm.prank(user1);
        poolV2.claimRewards();
        uint256 steppedReward = crossdToken.balanceOf(user1);

        assertEq(steppedReward, fullReward);
    }

    function test_SyncRounds_NoDoubleCountAfterPartialSync() public {
        _userDeposit(user1, 100 ether);

        _createRound(10000 ether, 10, 100);
        _advanceBlocks(60);

        vm.prank(sponsor);
        poolV2.syncRounds(1);
        uint256 globalAfterSync = poolV2.globalAccRewardPerShare();

        _advanceBlocks(50);

        vm.prank(sponsor);
        poolV2.syncRounds(1);
        uint256 globalAfterSync2 = poolV2.globalAccRewardPerShare();

        assertGt(globalAfterSync2, globalAfterSync);

        vm.prank(user1);
        poolV2.claimRewards();

        assertEq(crossdToken.balanceOf(user1), 10000 ether);
    }

    function test_SyncRounds_LargeRoundCount_GasComparison() public {
        _userDeposit(user1, 100 ether);

        uint256 roundCount = 50;
        for (uint256 i = 0; i < roundCount; i++) {
            _createRound(10000 ether, 100, 500);
        }
        assertEq(poolV2.getActiveRoundCount(), roundCount);

        _advanceBlocks(700);

        vm.prank(sponsor);
        poolV2.syncRounds(0);
        assertEq(poolV2.getActiveRoundCount(), 0);

        uint256 pending = _getPendingReward(user1);
        assertEq(pending, 10000 ether * roundCount);
    }

    function test_SyncRounds_NotStartedRoundsSkipped() public {
        _userDeposit(user1, 100 ether);

        _createRound(10000 ether, 10, 100);
        _createRound(10000 ether, 9999, 100);

        _advanceBlocks(200);

        vm.prank(sponsor);
        (uint256 processed, uint256 removed) = poolV2.syncRounds(0);
        assertEq(processed, 2);
        assertEq(removed, 1);
        assertEq(poolV2.getActiveRoundCount(), 1);
    }

    function test_SyncRounds_ReclaimableWithNoDepositors() public {
        for (uint256 i = 0; i < 5; i++) {
            _createRound(10000 ether, 10, 100);
        }

        _advanceBlocks(200);

        vm.prank(sponsor);
        poolV2.syncRounds(0);

        assertEq(poolV2.reclaimableAmount(), 50000 ether);
        assertEq(poolV2.getActiveRoundCount(), 0);
    }

    function test_SyncRounds_FactoryOwnerCanSync() public {
        _createRound(10000 ether, 10, 100);
        _advanceBlocks(200);

        // owner = address(this) = crossGameReward.owner()
        (uint256 processed, uint256 removed) = poolV2.syncRounds(0);
        assertEq(processed, 1);
        assertEq(removed, 1);
    }

    function test_SyncRounds_UnauthorizedReverts() public {
        vm.prank(user1);
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPoolV2.CGRP2SyncNotAuthorized.selector));
        poolV2.syncRounds(10);
    }
}
