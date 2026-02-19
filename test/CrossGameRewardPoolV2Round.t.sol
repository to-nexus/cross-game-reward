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
        assertEq(round.totalReward, amount);
        assertEq(round.startBlock, startBlock);
        assertEq(round.endBlock, startBlock + duration);
        assertEq(round.rewardPerBlock, amount / duration);
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
        // 10001 CROSSD / 1000 blocks = 10 CROSSD/block, remainder = 1 CROSSD
        uint256 amount = 10001 ether;
        uint256 duration = 1000;

        uint256 devBalanceBefore = crossdToken.balanceOf(sponsor);

        vm.startPrank(sponsor);
        crossdToken.approve(address(poolV2), amount);
        uint256 roundId = poolV2.createRound(amount, block.number + 10, duration);
        vm.stopPrank();

        uint256 devBalanceAfter = crossdToken.balanceOf(sponsor);

        // Verify round stores actualReward (without remainder)
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        uint256 expectedRewardPerBlock = amount / duration; // 10 ether
        uint256 expectedActualReward = expectedRewardPerBlock * duration; // 10000 ether
        assertEq(round.totalReward, expectedActualReward);
        assertEq(round.rewardPerBlock, expectedRewardPerBlock);

        // Developer should only have been charged actualReward (remainder stays)
        assertEq(devBalanceBefore - devBalanceAfter, expectedActualReward);

        // Pool balance should equal actualReward
        assertEq(crossdToken.balanceOf(address(poolV2)), expectedActualReward);

        // Now verify full distribution: deposit user, run through round
        _userDeposit(user1, 100 ether);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(duration);

        uint256 pending = _getPendingReward(user1);
        assertEq(pending, expectedActualReward);

        // Claim and verify no tokens are stuck
        vm.prank(user1);
        poolV2.claimRewards();

        // Pool should have 0 reward tokens left (no dust)
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

        // When exact division: totalReward == amount, full transfer
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        assertEq(round.totalReward, amount);
        assertEq(round.rewardPerBlock, 100 ether);
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
        // rewardPerBlock = 7e18 / 3 = 2333333333333333333
        // actualReward = 2333333333333333333 * 3 = 6999999999999999999 (1 wei remainder)
        uint256 amount = 7 ether;
        uint256 duration = 3;

        uint256 devBalanceBefore = crossdToken.balanceOf(sponsor);

        vm.startPrank(sponsor);
        crossdToken.approve(address(poolV2), amount);
        uint256 roundId = poolV2.createRound(amount, block.number + 10, duration);
        vm.stopPrank();

        uint256 devBalanceAfter = crossdToken.balanceOf(sponsor);

        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);
        uint256 expectedRpb = amount / duration; // 2333333333333333333
        uint256 expectedActual = expectedRpb * duration; // 6999999999999999999
        uint256 remainder = amount - expectedActual; // 1 wei

        assertGt(expectedRpb, 0);
        assertGt(remainder, 0); // There IS a remainder
        assertEq(round.rewardPerBlock, expectedRpb);
        assertEq(round.totalReward, expectedActual);

        // Only actualReward was transferred (remainder stays with sponsor)
        assertEq(devBalanceBefore - devBalanceAfter, expectedActual);
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
}
