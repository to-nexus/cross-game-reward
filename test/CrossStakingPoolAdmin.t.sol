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
        uint[] memory rewards = pool.pendingRewards(user2);
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

    // ==================== Emergency Withdraw ====================

    function testEmergencyWithdraw() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove token via CrossStaking
        crossStaking.removeRewardToken(1, rewardToken1);

        // Accidentally deposit additional rewards
        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // Check withdrawable amount
        uint withdrawable = pool.getEmergencyWithdrawableAmount(rewardToken1);
        assertEq(withdrawable, 50 ether, "Should be able to withdraw extra deposit");

        // Perform emergency withdrawal
        uint ownerBalanceBefore = rewardToken1.balanceOf(owner);
        vm.prank(owner);
        pool.emergencyWithdraw(rewardToken1, owner);

        // Confirm withdrawal result
        assertEq(rewardToken1.balanceOf(owner) - ownerBalanceBefore, 50 ether, "Emergency withdraw succeeded");
        assertEq(pool.getEmergencyWithdrawableAmount(rewardToken1), 0, "No more withdrawable");

        // User can still claim rewards
        vm.prank(user1);
        pool.claimReward(rewardToken1);
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 100 ether, 100, "User can still claim");
    }

    function testEmergencyWithdrawAfterUserClaim() public {
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
        uint withdrawable = pool.getEmergencyWithdrawableAmount(rewardToken1);
        assertEq(withdrawable, 50 ether, "Still 50 withdrawable after user claim");

        // Owner with DEFAULT_ADMIN_ROLE can perform the withdrawal directly
        uint ownerBalanceBefore = rewardToken1.balanceOf(owner);
        pool.emergencyWithdraw(rewardToken1, owner);
        assertEq(rewardToken1.balanceOf(owner) - ownerBalanceBefore, 50 ether, "Withdraw only extra deposits");
    }

    function testCannotEmergencyWithdrawNonRemovedToken() public view {
        // Cannot withdraw rewards for a token that is still registered
        uint withdrawable = pool.getEmergencyWithdrawableAmount(rewardToken1);
        assertEq(withdrawable, 0, "Non-removed token has 0 withdrawable");
    }

    function testCannotEmergencyWithdrawZero() public {
        // Remove but no extra deposits via CrossStaking
        crossStaking.removeRewardToken(1, rewardToken1);

        // No extra deposits - owner can call emergencyWithdraw directly
        vm.expectRevert(CrossStakingPool.CSPNoWithdrawableAmount.selector);
        pool.emergencyWithdraw(rewardToken1, owner);
    }

    function testOnlyAdminCanEmergencyWithdraw() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        crossStaking.removeRewardToken(1, rewardToken1);

        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // Non-admin cannot withdraw
        vm.prank(user1);
        vm.expectRevert();
        pool.emergencyWithdraw(rewardToken1, user1);
    }

    // ==================== Pause feature tests ====================

    function testPause() public {
        crossStaking.setPoolActive(1, false);

        assertTrue(pool.paused(), "Pool should be paused");
    }

    function testUnpause() public {
        crossStaking.setPoolActive(1, false);
        crossStaking.setPoolActive(1, true);

        assertFalse(pool.paused(), "Pool should be unpaused");
    }

    function testPauseOnlyByPauserRole() public {
        vm.prank(user1);
        vm.expectRevert();
        pool.pause();
    }

    function testUnpauseOnlyByPauserRole() public {
        crossStaking.setPoolActive(1, false);

        vm.prank(user1);
        vm.expectRevert();
        pool.unpause();
    }

    function testCannotStakeWhenPaused() public {
        crossStaking.setPoolActive(1, false);

        vm.startPrank(user1);
        crossToken.approve(address(pool), 10 ether);
        vm.expectRevert();
        pool.stake(10 ether);
        vm.stopPrank();
    }

    function testCannotUnstakeWhenPaused() public {
        _userStake(user1, 10 ether);

        crossStaking.setPoolActive(1, false);

        vm.prank(user1);
        vm.expectRevert();
        pool.unstake();
    }

    function testCannotClaimWhenPaused() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        crossStaking.setPoolActive(1, false);

        vm.prank(user1);
        vm.expectRevert();
        pool.claimRewards();
    }

    function testStakeAfterUnpause() public {
        crossStaking.setPoolActive(1, false);
        crossStaking.setPoolActive(1, true);

        _userStake(user1, 10 ether);
        assertEq(pool.balances(user1), 10 ether, "Should be able to stake after unpause");
    }

    // ==================== Role-based access control tests ====================

    function testOwnerHasDefaultAdminRole() public view {
        // Owner is the CrossStaking contract
        assertTrue(pool.hasRole(pool.DEFAULT_ADMIN_ROLE(), owner), "CrossStaking should have DEFAULT_ADMIN_ROLE");
    }

    function testCrossStakingHasStakingRootRole() public view {
        // CrossStaking contract holds STAKING_ROOT_ROLE
        assertTrue(
            pool.hasRole(pool.STAKING_ROOT_ROLE(), address(crossStaking)), "CrossStaking should have STAKING_ROOT_ROLE"
        );
    }

    function testOnlyStakingRootCanPause() public {
        // User1 lacks STAKING_ROOT_ROLE so cannot pause
        vm.prank(user1);
        vm.expectRevert();
        pool.pause();

        // Only CrossStaking can pause via setPoolActive
        crossStaking.setPoolActive(1, false);
        assertTrue(pool.paused(), "CrossStaking with STAKING_ROOT_ROLE can pause");
    }

    function testOnlyStakingRootCanAddRewardToken() public {
        // User1 without STAKING_ROOT_ROLE cannot add a reward token
        vm.prank(user1);
        vm.expectRevert();
        pool.addRewardToken(rewardToken3);

        // CrossStaking can add the token via the factory
        crossStaking.addRewardToken(1, rewardToken3);
        assertEq(pool.rewardTokenCount(), 3, "CrossStaking with STAKING_ROOT_ROLE can add reward token");
    }

    function testGrantStakingRootRole() public {
        // Owner (DEFAULT_ADMIN) can grant STAKING_ROOT_ROLE to another address
        bytes32 role = pool.STAKING_ROOT_ROLE();
        pool.grantRole(role, user1);

        assertTrue(pool.hasRole(role, user1), "User1 should have STAKING_ROOT_ROLE");

        // User1 can now pause
        vm.prank(user1);
        pool.pause();
        assertTrue(pool.paused(), "User1 with STAKING_ROOT_ROLE can pause");
    }

    function testRevokeStakingRootRole() public {
        pool.grantRole(pool.STAKING_ROOT_ROLE(), user1);
        pool.revokeRole(pool.STAKING_ROOT_ROLE(), user1);

        assertFalse(pool.hasRole(pool.STAKING_ROOT_ROLE(), user1), "User1 should not have STAKING_ROOT_ROLE");

        vm.prank(user1);
        vm.expectRevert();
        pool.pause();
    }

    // ==================== UUPS upgrade authorization tests ====================

    function testUpgradeAuthorization() public view {
        // Owner has DEFAULT_ADMIN_ROLE, so upgrades are permitted
        assertTrue(pool.hasRole(pool.DEFAULT_ADMIN_ROLE(), owner), "Owner should be able to upgrade");
    }

    function testNonAdminCannotUpgrade() public view {
        // User1 lacks DEFAULT_ADMIN_ROLE, so upgrades are not allowed
        assertFalse(pool.hasRole(pool.DEFAULT_ADMIN_ROLE(), user1), "User1 should not be able to upgrade");
    }

    // ==================== Initialization tests ====================

    function testInitialConfiguration() public view {
        assertEq(address(pool.stakingToken()), address(crossToken), "Staking token should be set");
        assertEq(pool.rewardTokenCount(), 2, "Should have 2 reward tokens from setup");
        assertFalse(pool.paused(), "Should not be paused initially");
    }
}
