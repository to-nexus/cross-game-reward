// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "./base/GamePoolBase.t.sol";

/**
 * @title GamePoolAuditFixesTest
 * @notice Tests for CertiK audit issue fixes: CGR-05, CGR-07, CGR-08
 */
contract GamePoolAuditFixesTest is GamePoolBase {
    // ==================== CGR-05: Rounding Dust Tests ====================

    function test_CGR05_DustGoesToReclaimable() public {
        // Use a deposit amount that will cause truncation dust
        // reward = 7 wei per block, totalDeposited = 3 wei
        // rewardPerShare = (7 * 1e18) / 3 = 2333333333333333333 (truncated)
        // distributed = (2333333333333333333 * 3) / 1e18 = 6 (dust = 1 wei)
        _userDeposit(user1, 3 ether);

        uint256 amount = 7 ether;
        uint256 duration = 1;
        uint256 roundId = _createRound(amount, 10, duration);

        IGamePool.Round memory round = gamePool.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(duration);

        // Trigger pool update
        vm.prank(user1);
        gamePool.claimRewards();

        // There should be some dust in reclaimableAmount
        uint256 claimed = crossdToken.balanceOf(user1);
        uint256 reclaimable = gamePool.reclaimableAmount();

        // Total accounting: claimed + reclaimable == amount (no tokens stuck)
        assertEq(claimed + reclaimable, amount, "No tokens should be permanently stuck");
        assertGt(reclaimable, 0, "Dust should be reclaimable");
    }

    function test_CGR05_RewardPerShareZero_FullRewardReclaimable() public {
        // totalDeposited is very large, reward per block is tiny
        // so rewardPerShare truncates to 0
        uint256 hugeTotalDeposited = 1e30; // 1 trillion tokens (1e12 * 1e18)

        gameToken.mint(user1, hugeTotalDeposited);

        vm.startPrank(user1);
        gameToken.approve(address(gamePool), hugeTotalDeposited);
        gamePool.deposit(hugeTotalDeposited);
        vm.stopPrank();

        // Create round with tiny reward: rewardPerBlock will be small
        // reward = 1e10 (minimum granularity from REWARD_PER_BLOCK_PRECISION)
        // rewardPerShare = (1e10 * 1e18) / 1e30 = 1e-2 => 0 (truncated)
        uint256 amount = 1e10;
        uint256 duration = 1;

        vm.startPrank(sponsor);
        crossdToken.approve(address(gamePool), amount);
        uint256 roundId = gamePool.createRound(amount, block.number + 10, duration);
        vm.stopPrank();

        IGamePool.Round memory round = gamePool.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(duration);

        // Trigger pool update
        vm.prank(user1);
        gamePool.claimRewards();

        // rewardPerShare was 0 => entire reward goes to reclaimable
        assertEq(gamePool.reclaimableAmount(), amount, "Full reward should be reclaimable when rewardPerShare == 0");
        assertEq(crossdToken.balanceOf(user1), 0, "User should get 0 when rewardPerShare truncates to 0");
    }

    function test_CGR05_DustReclaimable_MultipleBlocks() public {
        _userDeposit(user1, 7 ether);

        // Create round: 10 ether over 3 blocks
        // rewardPerBlock = (10e18 / 3 / 1e10) * 1e10 = 3333333330000000000
        // actual distributed per block = 3333333330000000000
        // remainder per block truncation will produce dust each block
        uint256 amount = 10 ether;
        uint256 duration = 3;
        uint256 roundId = _createRound(amount, 10, duration);

        IGamePool.Round memory round = gamePool.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(duration);

        vm.prank(user1);
        gamePool.claimRewards();

        uint256 claimed = crossdToken.balanceOf(user1);
        uint256 reclaimable = gamePool.reclaimableAmount();

        assertEq(claimed + reclaimable, amount, "All tokens must be accounted for");
    }

    function test_CGR05_DustReclaimable_MultipleUsers() public {
        _userDeposit(user1, 100 ether);
        _userDeposit(user2, 200 ether);
        _userDeposit(user3, 300 ether);
        // totalDeposited = 600 ether

        uint256 amount = 1000 ether;
        uint256 duration = 7;
        uint256 roundId = _createRound(amount, 10, duration);

        IGamePool.Round memory round = gamePool.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(duration);

        // All users claim
        vm.prank(user1);
        gamePool.claimRewards();
        vm.prank(user2);
        gamePool.claimRewards();
        vm.prank(user3);
        gamePool.claimRewards();

        uint256 totalClaimed =
            crossdToken.balanceOf(user1) + crossdToken.balanceOf(user2) + crossdToken.balanceOf(user3);
        uint256 reclaimable = gamePool.reclaimableAmount();

        assertEq(totalClaimed + reclaimable, amount, "All tokens must be distributed or reclaimable");
    }

    function test_CGR05_DustReclaimTokens_CanWithdraw() public {
        _userDeposit(user1, 3 ether);

        uint256 amount = 7 ether;
        uint256 roundId = _createRound(amount, 10, 1);

        IGamePool.Round memory round = gamePool.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(1);

        vm.prank(user1);
        gamePool.claimRewards();

        uint256 reclaimable = gamePool.reclaimableAmount();
        assertGt(reclaimable, 0, "Should have reclaimable dust");

        // Reclaim dust
        address recipient = address(0x9999);
        crossGameReward.reclaimFromPool(poolId, crossdToken, recipient);

        assertEq(crossdToken.balanceOf(recipient), reclaimable);
        assertEq(gamePool.reclaimableAmount(), 0);
    }

    function test_CGR05_NoDustWhenPerfectlyDivisible() public {
        // 1000 ether / 10 blocks = 100 ether/block, no remainder
        // 100 ether * 1e18 / 100 ether deposit = 1e18, no truncation
        _userDeposit(user1, 100 ether);

        uint256 amount = 1000 ether;
        uint256 duration = 10;
        uint256 roundId = _createRound(amount, 10, duration);

        IGamePool.Round memory round = gamePool.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(duration);

        vm.prank(user1);
        gamePool.claimRewards();

        assertEq(crossdToken.balanceOf(user1), amount, "All rewards should go to user");
        assertEq(gamePool.reclaimableAmount(), 0, "No dust when perfectly divisible");
    }

    function test_CGR05_DustAccumulation_OverMultipleRounds() public {
        _userDeposit(user1, 3 ether);

        // Create 5 rounds, each with dust
        uint256 totalAmount;
        for (uint256 i = 0; i < 5; i++) {
            uint256 amount = 7 ether;
            totalAmount += amount;
            _createRound(amount, 10 + i * 20, 1);
        }

        // Advance past all rounds
        _advanceBlocks(200);

        vm.prank(user1);
        gamePool.claimRewards();

        uint256 claimed = crossdToken.balanceOf(user1);
        uint256 reclaimable = gamePool.reclaimableAmount();

        assertEq(claimed + reclaimable, totalAmount, "All tokens accounted across multiple rounds");
    }

    function test_CGR05_PendingRewardsView_ConsistentWithClaim() public {
        _userDeposit(user1, 3 ether);

        uint256 amount = 7 ether;
        uint256 roundId = _createRound(amount, 10, 1);

        IGamePool.Round memory round = gamePool.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(1);

        // Check pending rewards view
        uint256 pending = _getPendingReward(user1);

        // Claim
        vm.prank(user1);
        gamePool.claimRewards();

        uint256 claimed = crossdToken.balanceOf(user1);

        assertEq(pending, claimed, "pendingRewards view should match actual claim amount");
    }

    // ==================== CGR-07: Max Active Rounds Tests ====================

    function test_CGR07_InitialMaxActiveRounds() public view {
        assertEq(gamePool.maxActiveRounds(), 50, "Default max active rounds should be 50");
    }

    function test_CGR07_SetMaxActiveRounds() public {
        crossGameReward.setMaxActiveRounds(poolId, 100);
        assertEq(gamePool.maxActiveRounds(), 100);
    }

    function test_CGR07_SetMaxActiveRounds_EmitsEvent() public {
        vm.expectEmit(false, false, false, true);
        emit GamePool.MaxActiveRoundsUpdated(50, 100);
        crossGameReward.setMaxActiveRounds(poolId, 100);
    }

    function test_CGR07_SetMaxActiveRounds_Zero_Reverts() public {
        vm.expectRevert(GamePool.GPInvalidMaxActiveRounds.selector);
        crossGameReward.setMaxActiveRounds(poolId, 0);
    }

    function test_CGR07_SetMaxActiveRounds_OnlyManager() public {
        vm.prank(user1);
        vm.expectRevert();
        crossGameReward.setMaxActiveRounds(poolId, 100);
    }

    function test_CGR07_SetMaxActiveRounds_OnlyGamePool() public {
        (uint256 v1PoolId,) = crossGameReward.createPool("V1 Pool", IERC20(address(gameToken)), MIN_DEPOSIT);

        vm.expectRevert(
            abi.encodeWithSelector(
                CrossGameReward.CGRInvalidPoolType.selector,
                v1PoolId,
                ICrossGameReward.PoolType.GamePool,
                ICrossGameReward.PoolType.CrossPool
            )
        );
        crossGameReward.setMaxActiveRounds(v1PoolId, 100);
    }

    function test_CGR07_CreateRound_ExceedsMax_Reverts() public {
        crossGameReward.setMaxActiveRounds(poolId, 3);

        _createRound(10000 ether, 100, 1000);
        _createRound(10000 ether, 100, 1000);
        _createRound(10000 ether, 100, 1000);

        assertEq(gamePool.getActiveRoundCount(), 3);

        // 4th round should revert
        vm.startPrank(sponsor);
        crossdToken.approve(address(gamePool), 10000 ether);
        vm.expectRevert(abi.encodeWithSelector(GamePool.GPMaxActiveRoundsReached.selector, 3, 3));
        gamePool.createRound(10000 ether, block.number + 100, 1000);
        vm.stopPrank();
    }

    function test_CGR07_CreateRound_AtExactMax() public {
        crossGameReward.setMaxActiveRounds(poolId, 2);

        _createRound(10000 ether, 100, 1000);
        _createRound(10000 ether, 100, 1000);

        // 3rd round should fail
        vm.startPrank(sponsor);
        crossdToken.approve(address(gamePool), 10000 ether);
        vm.expectRevert(abi.encodeWithSelector(GamePool.GPMaxActiveRoundsReached.selector, 2, 2));
        gamePool.createRound(10000 ether, block.number + 100, 1000);
        vm.stopPrank();
    }

    function test_CGR07_CreateRound_AfterRoundCompletion() public {
        crossGameReward.setMaxActiveRounds(poolId, 2);

        _userDeposit(user1, 100 ether);

        uint256 roundId1 = _createRound(10000 ether, 10, 50);
        _createRound(10000 ether, 100, 1000);

        // Advance past first round
        IGamePool.Round memory round1 = gamePool.getRound(roundId1);
        _advanceToBlock(round1.endBlock + 1);

        // Trigger update to remove completed round
        vm.prank(user1);
        gamePool.claimRewards();

        assertEq(gamePool.getActiveRoundCount(), 1, "Completed round should be removed");

        // Now creating a new round should work
        _createRound(10000 ether, 10, 1000);
        assertEq(gamePool.getActiveRoundCount(), 2);
    }

    function test_CGR07_CreateRound_AfterCancellation() public {
        crossGameReward.setMaxActiveRounds(poolId, 2);

        _createRound(10000 ether, 100, 1000);
        uint256 roundId2 = _createRound(10000 ether, 100, 1000);

        assertEq(gamePool.getActiveRoundCount(), 2);

        // Cancel one round
        vm.prank(sponsor);
        gamePool.cancelRound(roundId2);

        assertEq(gamePool.getActiveRoundCount(), 1);

        // Now can create again
        _createRound(10000 ether, 100, 1000);
        assertEq(gamePool.getActiveRoundCount(), 2);
    }

    function test_CGR07_IncreaseMax_AllowsMoreRounds() public {
        crossGameReward.setMaxActiveRounds(poolId, 2);

        _createRound(10000 ether, 100, 1000);
        _createRound(10000 ether, 100, 1000);

        // Should fail at max
        vm.startPrank(sponsor);
        crossdToken.approve(address(gamePool), 10000 ether);
        vm.expectRevert(abi.encodeWithSelector(GamePool.GPMaxActiveRoundsReached.selector, 2, 2));
        gamePool.createRound(10000 ether, block.number + 100, 1000);
        vm.stopPrank();

        // Increase max
        crossGameReward.setMaxActiveRounds(poolId, 5);

        // Now should succeed
        _createRound(10000 ether, 100, 1000);
        assertEq(gamePool.getActiveRoundCount(), 3);
    }

    function test_CGR07_DecreaseMax_DoesNotAffectExisting() public {
        _createRound(10000 ether, 100, 1000);
        _createRound(10000 ether, 100, 1000);
        _createRound(10000 ether, 100, 1000);

        assertEq(gamePool.getActiveRoundCount(), 3);

        // Decrease max below current count
        crossGameReward.setMaxActiveRounds(poolId, 1);

        // Existing rounds are unaffected
        assertEq(gamePool.getActiveRoundCount(), 3);

        // But cannot create new ones
        vm.startPrank(sponsor);
        crossdToken.approve(address(gamePool), 10000 ether);
        vm.expectRevert(abi.encodeWithSelector(GamePool.GPMaxActiveRoundsReached.selector, 3, 1));
        gamePool.createRound(10000 ether, block.number + 100, 1000);
        vm.stopPrank();
    }

    function test_CGR07_SetMaxActiveRounds_Multiple_Updates() public {
        crossGameReward.setMaxActiveRounds(poolId, 10);
        assertEq(gamePool.maxActiveRounds(), 10);

        crossGameReward.setMaxActiveRounds(poolId, 200);
        assertEq(gamePool.maxActiveRounds(), 200);

        crossGameReward.setMaxActiveRounds(poolId, 1);
        assertEq(gamePool.maxActiveRounds(), 1);
    }

    // ==================== CGR-08: userRewards Token Validation Tests ====================

    function test_CGR08_UserRewards_CorrectToken() public {
        _userDeposit(user1, 100 ether);

        (uint256 rewardPerTokenPaid, uint256 rewards) = gamePool.userRewards(user1, crossdToken);

        assertEq(rewardPerTokenPaid, 0);
        assertEq(rewards, 0);
    }

    function test_CGR08_UserRewards_WrongToken_Reverts() public {
        _userDeposit(user1, 100 ether);

        vm.expectRevert(
            abi.encodeWithSelector(
                GamePool.GPInvalidRewardToken.selector, address(gameToken), address(crossdToken)
            )
        );
        gamePool.userRewards(user1, gameToken);
    }

    function test_CGR08_UserRewards_ZeroAddress_Reverts() public {
        _userDeposit(user1, 100 ether);

        vm.expectRevert(
            abi.encodeWithSelector(
                GamePool.GPInvalidRewardToken.selector, address(0), address(crossdToken)
            )
        );
        gamePool.userRewards(user1, IERC20(address(0)));
    }

    function test_CGR08_UserRewards_RandomToken_Reverts() public {
        _userDeposit(user1, 100 ether);

        MockERC20GP randomToken = new MockERC20GP("Random", "RND");

        vm.expectRevert(
            abi.encodeWithSelector(
                GamePool.GPInvalidRewardToken.selector, address(randomToken), address(crossdToken)
            )
        );
        gamePool.userRewards(user1, IERC20(address(randomToken)));
    }

    function test_CGR08_UserRewards_WithActiveRewards_CorrectToken() public {
        _userDeposit(user1, 100 ether);

        uint256 roundId = _createRound(10000 ether, 10, 100);
        IGamePool.Round memory round = gamePool.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(50);

        // Should succeed with correct token
        (uint256 rewardPerTokenPaid, uint256 rewards) = gamePool.userRewards(user1, crossdToken);

        // Should not revert; values not critical here, just ensuring no revert
        assertEq(rewardPerTokenPaid, 0);
        assertEq(rewards, 0);
    }

    function test_CGR08_UserRewards_ConsistencyWithPendingReward() public {
        _userDeposit(user1, 100 ether);

        uint256 roundId = _createRound(10000 ether, 10, 100);
        IGamePool.Round memory round = gamePool.getRound(roundId);
        _advanceToBlock(round.startBlock);
        _advanceBlocks(100);

        // Trigger update
        vm.prank(user1);
        gamePool.claimRewards();

        // Create new round for accumulated rewards
        uint256 roundId2 = _createRound(5000 ether, 10, 50);
        IGamePool.Round memory round2 = gamePool.getRound(roundId2);
        _advanceToBlock(round2.startBlock);
        _advanceBlocks(25);

        // Trigger update to store pending rewards
        _userDeposit(user1, 1 ether); // deposit triggers _updateUser

        (, uint256 rewards) = gamePool.userRewards(user1, crossdToken);

        // userRewards returns _userPendingRewards[account] which should be non-zero
        // after _updateUser was called
        assertGt(rewards, 0, "Should have accumulated pending rewards");
    }
}
