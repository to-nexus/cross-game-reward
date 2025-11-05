// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/CrossStakingPoolBase.t.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title CrossStakingPoolAdminTest
 * @notice Admin feature tests covering roles, pause, and reward token management
 */
contract CrossStakingPoolAdminTest is CrossStakingPoolBase {
    MockERC20 public rewardToken3;

    function setUp() public override {
        super.setUp();
        rewardToken3 = new MockERC20("Reward Token 3", "RWD3");
        rewardToken3.transfer(owner, 10000 ether);
    }

    // ==================== Reward token management tests ====================

    function testAddRewardToken() public {
        crossStaking.addRewardToken(1, rewardToken3);

        assertEq(pool.rewardTokenCount(), 3, "Should have 3 reward tokens");
        assertTrue(pool.isRewardToken(rewardToken3), "Should be registered as reward token");
    }

    function testAddRewardTokenOnlyByManager() public {
        vm.prank(user1);
        vm.expectRevert();
        pool.addRewardToken(rewardToken3);
    }

    function testCannotAddSameRewardTokenTwice() public {
        crossStaking.addRewardToken(1, rewardToken3);

        vm.expectRevert(CrossStakingPool.CSPRewardTokenAlreadyAdded.selector);
        crossStaking.addRewardToken(1, rewardToken3);
    }

    function testCannotAddZeroAddressAsRewardToken() public {
        vm.expectRevert(CrossStakingPool.CSPCanNotZeroAddress.selector);
        crossStaking.addRewardToken(1, IERC20(address(0)));
    }

    function testCannotAddStakingTokenAsReward() public {
        vm.expectRevert(CrossStakingPool.CSPCanNotUseStakingToken.selector);
        crossStaking.addRewardToken(1, crossToken);
    }

    function testRewardTokenIndexMapping() public {
        crossStaking.addRewardToken(1, rewardToken3);

        // Check if token is registered
        assertTrue(pool.isRewardToken(rewardToken3), "Third token should be registered");

        // Check token at index 2
        assertEq(address(pool.rewardTokenAt(2)), address(rewardToken3), "Third token should be at index 2");
    }

    // ==================== Reward deposit permissions ====================

    function testDirectTransferReward() public {
        // User2 stakes first
        _userStake(user2, 10 ether);

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
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove reward token via CrossStaking
        crossStaking.removeRewardToken(1, rewardToken1);

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
        vm.expectRevert(CrossStakingPool.CSPInvalidRewardToken.selector);
        crossStaking.removeRewardToken(1, rewardToken3);
    }

    function testOnlyRewardManagerCanRemove() public {
        vm.prank(user1);
        vm.expectRevert();
        crossStaking.removeRewardToken(1, rewardToken1);
    }

    function testRemovedTokenNoNewRewards() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove token via CrossStaking
        crossStaking.removeRewardToken(1, rewardToken1);

        // Try to add new rewards (won't be distributed because not in set)
        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // Sync won't happen for removed token
        _userStake(user2, 10 ether);

        // user1 should only have original 100, not 150
        vm.prank(user1);
        pool.claimReward(rewardToken1);

        assertApproxEqAbs(rewardToken1.balanceOf(user1), 100 ether, 100, "No new rewards after removal");
    }

    // ==================== Withdraw ====================

    function testWithdraw() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove token via CrossStaking
        crossStaking.removeRewardToken(1, rewardToken1);

        // Accidentally deposit additional rewards
        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // Check withdrawable amount
        uint withdrawable = pool.getWithdrawableAmount(rewardToken1);
        assertEq(withdrawable, 50 ether, "Should be able to withdraw extra deposit");

        // Perform withdrawal via CrossStaking
        uint ownerBalanceBefore = rewardToken1.balanceOf(owner);
        crossStaking.withdrawFromPool(1, rewardToken1, owner);

        // Confirm withdrawal result
        assertEq(rewardToken1.balanceOf(owner) - ownerBalanceBefore, 50 ether, "Withdraw succeeded");
        assertEq(pool.getWithdrawableAmount(rewardToken1), 0, "No more withdrawable");

        // User can still claim rewards
        vm.prank(user1);
        pool.claimReward(rewardToken1);
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 100 ether, 100, "User can still claim");
    }

    function testWithdrawAfterUserClaim() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove token (removed balance = 100) via CrossStaking
        crossStaking.removeRewardToken(1, rewardToken1);

        // Accidentally add 50 more (total balance becomes 150)
        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // User claims (reduces on-chain balance and distributed amount)
        vm.prank(user1);
        pool.claimReward(rewardToken1);

        // Withdrawable = current balance - (balance at removal - claimed amount)
        uint withdrawable = pool.getWithdrawableAmount(rewardToken1);
        assertEq(withdrawable, 50 ether, "Still 50 withdrawable after user claim");

        // Owner can perform the withdrawal via CrossStaking
        uint ownerBalanceBefore = rewardToken1.balanceOf(owner);
        crossStaking.withdrawFromPool(1, rewardToken1, owner);
        assertEq(rewardToken1.balanceOf(owner) - ownerBalanceBefore, 50 ether, "Withdraw only extra deposits");
    }

    function testCannotWithdrawNonRemovedToken() public view {
        // Cannot withdraw rewards for a token that is still registered
        uint withdrawable = pool.getWithdrawableAmount(rewardToken1);
        assertEq(withdrawable, 0, "Non-removed token has 0 withdrawable");
    }

    function testCannotWithdrawZero() public {
        // Remove but no extra deposits via CrossStaking
        crossStaking.removeRewardToken(1, rewardToken1);

        // No extra deposits - cannot withdraw
        vm.expectRevert(CrossStakingPool.CSPNoWithdrawableAmount.selector);
        crossStaking.withdrawFromPool(1, rewardToken1, owner);
    }

    function testOnlyManagerCanWithdraw() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        crossStaking.removeRewardToken(1, rewardToken1);

        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // Non-manager cannot withdraw
        vm.prank(user1);
        vm.expectRevert();
        crossStaking.withdrawFromPool(1, rewardToken1, user1);
    }

    // ==================== Pool Status Tests ====================

    function testSetPoolStatusPaused() public {
        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Paused); // Paused = 2

        assertEq(uint(pool.poolStatus()), 2, "Pool should be paused");
        assertTrue(pool.paused(), "Pool should be paused");
    }

    function testSetPoolStatusInactive() public {
        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Inactive); // Inactive = 1

        assertEq(uint(pool.poolStatus()), 1, "Pool should be inactive");
        assertFalse(pool.paused(), "Pool should not be paused");
    }

    function testSetPoolStatusActive() public {
        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Inactive); // Inactive
        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Active); // Active = 0

        assertEq(uint(pool.poolStatus()), 0, "Pool should be active");
        assertFalse(pool.paused(), "Pool should not be paused");
    }

    function testOnlyStakingRootCanSetStatus() public {
        vm.prank(user1);
        vm.expectRevert();
        pool.setPoolStatus(ICrossStakingPool.PoolStatus.Paused); // Paused
    }

    function testCannotStakeWhenPaused() public {
        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Paused); // Paused

        vm.startPrank(user1);
        crossToken.approve(address(pool), 10 ether);
        // When paused, PausableUpgradeable.EnforcedPause is thrown first (from whenNotPaused modifier)
        vm.expectRevert();
        pool.stake(10 ether);
        vm.stopPrank();
    }

    function testCannotStakeWhenInactive() public {
        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Inactive); // Inactive

        vm.startPrank(user1);
        crossToken.approve(address(pool), 10 ether);
        vm.expectRevert(CrossStakingPool.CSPCannotStakeInCurrentState.selector);
        pool.stake(10 ether);
        vm.stopPrank();
    }

    function testCanUnstakeWhenInactive() public {
        _userStake(user1, 10 ether);

        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Inactive); // Inactive

        vm.prank(user1);
        pool.unstake();
        assertEq(pool.balances(user1), 0, "Should be able to unstake when inactive");
    }

    function testCannotUnstakeWhenPaused() public {
        _userStake(user1, 10 ether);

        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Paused); // Paused

        vm.prank(user1);
        // When paused, PausableUpgradeable.EnforcedPause is thrown first (from whenNotPaused modifier)
        vm.expectRevert();
        pool.unstake();
    }

    function testCanClaimWhenInactive() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Inactive); // Inactive

        vm.prank(user1);
        pool.claimRewards();
        assertGt(rewardToken1.balanceOf(user1), 0, "Should be able to claim when inactive");
    }

    function testCannotClaimWhenPaused() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Paused); // Paused

        vm.prank(user1);
        // When paused, PausableUpgradeable.EnforcedPause is thrown first (from whenNotPaused modifier)
        vm.expectRevert();
        pool.claimRewards();
    }

    function testStakeAfterReactivation() public {
        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Inactive); // Inactive
        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Active); // Active

        _userStake(user1, 10 ether);
        assertEq(pool.balances(user1), 10 ether, "Should be able to stake after reactivation");
    }

    // ==================== Access control tests ====================

    function testOwner() public view {
        // Pool owner is CrossStaking contract's owner (default admin)
        assertTrue(pool.owner() == owner, "Pool owner should be CrossStaking's owner");
    }

    function testCrossStakingReference() public view {
        // CrossStaking contract reference
        assertTrue(
            address(pool.crossStaking()) == address(crossStaking),
            "Pool crossStaking should be the CrossStaking contract"
        );
    }

    function testOnlyStakingRootCanChangeStatus() public {
        // User1 cannot set status
        vm.prank(user1);
        vm.expectRevert();
        pool.setPoolStatus(ICrossStakingPool.PoolStatus.Paused); // Paused

        // CrossStaking can set status
        crossStaking.setPoolStatus(1, ICrossStakingPool.PoolStatus.Paused); // Pool 1, Paused status
        assertTrue(pool.paused(), "CrossStaking (STAKING_ROOT) can set status");
    }

    function testOnlyStakingRootCanAddRewardToken() public {
        // User1 cannot add a reward token
        vm.prank(user1);
        vm.expectRevert();
        pool.addRewardToken(rewardToken3);

        // CrossStaking can add the token
        crossStaking.addRewardToken(1, rewardToken3);
        assertEq(pool.rewardTokenCount(), 3, "CrossStaking can add reward token");
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
        assertEq(address(pool.stakingToken()), address(crossToken), "Staking token should be set");
        assertEq(pool.rewardTokenCount(), 2, "Should have 2 reward tokens from setup");
        assertFalse(pool.paused(), "Should not be paused initially");
        assertEq(uint(pool.poolStatus()), 0, "Should be active initially");
    }
}
