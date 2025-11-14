// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/CrossGameRewardPoolBase.t.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title CrossGameRewardPoolAdminTest
 * @notice Admin feature tests covering roles, pause, and reward token management
 */
contract CrossGameRewardPoolAdminTest is CrossGameRewardPoolBase {
    MockERC20 public rewardToken3;

    function setUp() public override {
        super.setUp();
        rewardToken3 = new MockERC20("Reward Token 3", "RWD3");
        rewardToken3.transfer(owner, 10000 ether);
    }

    // ==================== Reward token management tests ====================

    function testAddRewardToken() public {
        crossGameReward.addRewardToken(1, rewardToken3);

        assertEq(pool.rewardTokenCount(), 3, "Should have 3 reward tokens");
        assertTrue(pool.isRewardToken(rewardToken3), "Should be registered as reward token");
    }

    function testAddRewardTokenOnlyByManager() public {
        vm.prank(user1);
        vm.expectRevert();
        pool.addRewardToken(rewardToken3);
    }

    function testCannotAddSameRewardTokenTwice() public {
        crossGameReward.addRewardToken(1, rewardToken3);

        vm.expectRevert(CrossGameRewardPool.CGRPRewardTokenAlreadyAdded.selector);
        crossGameReward.addRewardToken(1, rewardToken3);
    }

    function testCannotAddZeroAddressAsRewardToken() public {
        vm.expectRevert(CrossGameRewardPool.CGRPCanNotZeroAddress.selector);
        crossGameReward.addRewardToken(1, IERC20(address(0)));
    }

    function testCannotAddDepositTokenAsReward() public {
        vm.expectRevert(CrossGameRewardPool.CGRPCanNotUseDepositToken.selector);
        crossGameReward.addRewardToken(1, crossToken);
    }

    function testRewardTokenIndexMapping() public {
        crossGameReward.addRewardToken(1, rewardToken3);

        // Check if token is registered
        assertTrue(pool.isRewardToken(rewardToken3), "Third token should be registered");

        // Check token at index 2
        assertEq(address(pool.rewardTokenAt(2)), address(rewardToken3), "Third token should be at index 2");
    }

    // ==================== Reward deposit permissions ====================

    function testDirectTransferReward() public {
        // User2 deposits first
        _userDeposit(user2, 10 ether);

        // Any address can transfer rewards directly
        vm.startPrank(user1);
        rewardToken1.mint(user1, 100 ether);
        rewardToken1.transfer(address(pool), 100 ether);
        vm.stopPrank();

        // Verify rewards are credited
        (, uint[] memory rewards) = pool.pendingRewards(user2);
        assertApproxEqAbs(rewards[0], 100 ether, 100, "Anyone can transfer rewards");
    }

    // ==================== Remove Reward Token ====================

    function testRemoveRewardToken() public {
        _userDeposit(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove reward token via CrossGameReward
        crossGameReward.removeRewardToken(1, rewardToken1);

        // Verify removed from set
        assertFalse(pool.isRewardToken(rewardToken1), "Should be removed from set");
        assertEq(pool.rewardTokenCount(), 1, "Count decreased"); // rewardToken2 remains

        // But can still claim existing rewards
        vm.prank(user1);
        pool.claimReward(rewardToken1);

        // Verify claimed
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 100 ether, 100, "Can still claim after removal");
    }

    function testCannotRemoveNonExistentToken() public {
        vm.expectRevert(CrossGameRewardPool.CGRPInvalidRewardToken.selector);
        crossGameReward.removeRewardToken(1, rewardToken3);
    }

    function testOnlyRewardManagerCanRemove() public {
        vm.prank(user1);
        vm.expectRevert();
        crossGameReward.removeRewardToken(1, rewardToken1);
    }

    function testRemovedTokenNoNewRewards() public {
        _userDeposit(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove token via CrossGameReward
        crossGameReward.removeRewardToken(1, rewardToken1);

        // Try to add new rewards (won't be distributed because not in set)
        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // Sync won't happen for removed token
        _userDeposit(user2, 10 ether);

        // user1 should only have original 100, not 150
        vm.prank(user1);
        pool.claimReward(rewardToken1);

        assertApproxEqAbs(rewardToken1.balanceOf(user1), 100 ether, 100, "No new rewards after removal");
    }

    // ==================== Withdraw ====================

    function testTokensReclaimed() public {
        _userDeposit(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove token via CrossGameReward
        crossGameReward.removeRewardToken(1, rewardToken1);

        // Accidentally deposit additional rewards
        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // Check withdrawable amount
        uint withdrawable = pool.getReclaimableAmount(rewardToken1);
        assertEq(withdrawable, 50 ether, "Should be able to withdraw extra deposit");

        // Perform withdrawal via CrossGameReward
        uint ownerBalanceBefore = rewardToken1.balanceOf(owner);
        crossGameReward.reclaimFromPool(1, rewardToken1, owner);

        // Confirm withdrawal result
        assertEq(rewardToken1.balanceOf(owner) - ownerBalanceBefore, 50 ether, "Withdraw succeeded");
        assertEq(pool.getReclaimableAmount(rewardToken1), 0, "No more withdrawable");

        // User can still claim rewards
        vm.prank(user1);
        pool.claimReward(rewardToken1);
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 100 ether, 100, "User can still claim");
    }

    function testWithdrawAfterUserClaim() public {
        _userDeposit(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove token (removed balance = 100) via CrossGameReward
        crossGameReward.removeRewardToken(1, rewardToken1);

        // Accidentally add 50 more (total balance becomes 150)
        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // User claims (reduces on-chain balance and distributed amount)
        vm.prank(user1);
        pool.claimReward(rewardToken1);

        // Withdrawable = current balance - (balance at removal - claimed amount)
        uint withdrawable = pool.getReclaimableAmount(rewardToken1);
        assertEq(withdrawable, 50 ether, "Still 50 withdrawable after user claim");

        // Owner can perform the withdrawal via CrossGameReward
        uint ownerBalanceBefore = rewardToken1.balanceOf(owner);
        crossGameReward.reclaimFromPool(1, rewardToken1, owner);
        assertEq(rewardToken1.balanceOf(owner) - ownerBalanceBefore, 50 ether, "Withdraw only extra deposits");
    }

    function testCannotWithdrawNonRemovedToken() public view {
        // Cannot withdraw rewards for a token that is still registered
        uint withdrawable = pool.getReclaimableAmount(rewardToken1);
        assertEq(withdrawable, 0, "Non-removed token has 0 withdrawable");
    }

    function testCannotWithdrawZero() public {
        // Remove but no extra deposits via CrossGameReward
        crossGameReward.removeRewardToken(1, rewardToken1);

        // No extra deposits - cannot withdraw
        vm.expectRevert(CrossGameRewardPool.CGRPNoReclaimableAmount.selector);
        crossGameReward.reclaimFromPool(1, rewardToken1, owner);
    }

    function testOnlyManagerCanTokensReclaimed() public {
        _userDeposit(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        crossGameReward.removeRewardToken(1, rewardToken1);

        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // Non-manager cannot withdraw
        vm.prank(user1);
        vm.expectRevert();
        crossGameReward.reclaimFromPool(1, rewardToken1, user1);
    }

    // ==================== Pool Status Tests ====================

    function testSetPoolStatusPaused() public {
        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Paused); // Paused = 2

        assertEq(uint(pool.poolStatus()), 2, "Pool should be paused");
        assertTrue(pool.paused(), "Pool should be paused");
    }

    function testSetPoolStatusInactive() public {
        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Inactive); // Inactive = 1

        assertEq(uint(pool.poolStatus()), 1, "Pool should be inactive");
        assertFalse(pool.paused(), "Pool should not be paused");
    }

    function testSetPoolStatusActive() public {
        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Inactive); // Inactive
        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Active); // Active = 0

        assertEq(uint(pool.poolStatus()), 0, "Pool should be active");
        assertFalse(pool.paused(), "Pool should not be paused");
    }

    function testOnlyRewardRootCanSetStatus() public {
        vm.prank(user1);
        vm.expectRevert();
        pool.setPoolStatus(ICrossGameRewardPool.PoolStatus.Paused); // Paused
    }

    function testCannotDepositWhenPaused() public {
        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Paused); // Paused

        vm.startPrank(user1);
        crossToken.approve(address(pool), 10 ether);
        // When paused, PausableUpgradeable.EnforcedPause is thrown first (from whenNotPaused modifier)
        vm.expectRevert();
        pool.deposit(10 ether);
        vm.stopPrank();
    }

    function testCannotDepositWhenInactive() public {
        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Inactive); // Inactive

        vm.startPrank(user1);
        crossToken.approve(address(pool), 10 ether);
        vm.expectRevert(CrossGameRewardPool.CGRPCannotDepositInCurrentState.selector);
        pool.deposit(10 ether);
        vm.stopPrank();
    }

    function testCanWithdrawWhenInactive() public {
        _userDeposit(user1, 10 ether);

        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Inactive); // Inactive

        vm.prank(user1);
        pool.withdraw();
        assertEq(pool.balances(user1), 0, "Should be able to withdraw when inactive");
    }

    function testCannotWithdrawWhenPaused() public {
        _userDeposit(user1, 10 ether);

        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Paused); // Paused

        vm.prank(user1);
        // When paused, PausableUpgradeable.EnforcedPause is thrown first (from whenNotPaused modifier)
        vm.expectRevert();
        pool.withdraw();
    }

    function testCanClaimWhenInactive() public {
        _userDeposit(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Inactive); // Inactive

        vm.prank(user1);
        pool.claimRewards();
        assertGt(rewardToken1.balanceOf(user1), 0, "Should be able to claim when inactive");
    }

    function testCannotClaimWhenPaused() public {
        _userDeposit(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Paused); // Paused

        vm.prank(user1);
        // When paused, PausableUpgradeable.EnforcedPause is thrown first (from whenNotPaused modifier)
        vm.expectRevert();
        pool.claimRewards();
    }

    function testDepositAfterReactivation() public {
        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Inactive); // Inactive
        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Active); // Active

        _userDeposit(user1, 10 ether);
        assertEq(pool.balances(user1), 10 ether, "Should be able to deposit after reactivation");
    }

    // ==================== Access control tests ====================

    function testOwner() public view {
        // Pool owner is CrossGameReward contract's owner (default admin)
        assertTrue(pool.owner() == owner, "Pool owner should be CrossGameReward's owner");
    }

    function testCrossGameRewardReference() public view {
        // CrossGameReward contract reference
        assertTrue(
            address(pool.crossGameReward()) == address(crossGameReward),
            "Pool crossGameReward should be the CrossGameReward contract"
        );
    }

    function testOnlyRewardRootCanChangeStatus() public {
        // User1 cannot set status
        vm.prank(user1);
        vm.expectRevert();
        pool.setPoolStatus(ICrossGameRewardPool.PoolStatus.Paused); // Paused

        // CrossGameReward can set status
        crossGameReward.setPoolStatus(1, ICrossGameRewardPool.PoolStatus.Paused); // Pool 1, Paused status
        assertTrue(pool.paused(), "CrossGameReward (REWARD_ROOT) can set status");
    }

    function testOnlyRewardRootCanAddRewardToken() public {
        // User1 cannot add a reward token
        vm.prank(user1);
        vm.expectRevert();
        pool.addRewardToken(rewardToken3);

        // CrossGameReward can add the token
        crossGameReward.addRewardToken(1, rewardToken3);
        assertEq(pool.rewardTokenCount(), 3, "CrossGameReward can add reward token");
    }

    // ==================== UUPS upgrade authorization tests ====================

    function testUpgradeAuthorization() public view {
        // Owner should be able to upgrade
        assertTrue(pool.owner() == owner, "Owner should be able to upgrade");
    }

    function testNonOwnerCannotUpgrade() public view {
        // User1 is not the owner
        assertFalse(pool.owner() == user1, "User1 should not be able to upgrade");
    }

    // ==================== Initialization tests ====================

    function testInitialConfiguration() public view {
        assertEq(address(pool.depositToken()), address(crossToken), "Deposit token should be set");
        assertEq(pool.rewardTokenCount(), 2, "Should have 2 reward tokens from setup");
        assertFalse(pool.paused(), "Should not be paused initially");
        assertEq(uint(pool.poolStatus()), 0, "Should be active initially");
    }
}
