// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/CrossStakingPoolBase.t.sol";

/**
 * @title CrossStakingPoolAdminTest
 * @notice 관리자 기능 테스트 (권한, pause, 보상 토큰 관리)
 */
contract CrossStakingPoolAdminTest is CrossStakingPoolBase {
    MockERC20 public rewardToken3;

    function setUp() public override {
        super.setUp();
        rewardToken3 = new MockERC20("Reward Token 3", "RWD3");
        rewardToken3.transfer(owner, 10000 ether);
    }

    // ==================== 보상 토큰 관리 테스트 ====================

    function testAddRewardToken() public {
        vm.prank(owner);
        pool.addRewardToken(address(rewardToken3));

        assertEq(pool.rewardTokenCount(), 3, "Should have 3 reward tokens");
        assertTrue(pool.isRewardToken(address(rewardToken3)), "Should be registered as reward token");
    }

    function testAddRewardTokenOnlyByManager() public {
        vm.prank(user1);
        vm.expectRevert();
        pool.addRewardToken(address(rewardToken3));
    }

    function testCannotAddSameRewardTokenTwice() public {
        vm.startPrank(owner);
        pool.addRewardToken(address(rewardToken3));

        vm.expectRevert(CrossStakingPool.CSPRewardTokenAlreadyAdded.selector);
        pool.addRewardToken(address(rewardToken3));
        vm.stopPrank();
    }

    function testCannotAddZeroAddressAsRewardToken() public {
        vm.prank(owner);
        vm.expectRevert(CrossStakingPool.CSPCanNotZeroAddress.selector);
        pool.addRewardToken(address(0));
    }

    function testCannotAddStakingTokenAsReward() public {
        vm.prank(owner);
        vm.expectRevert(CrossStakingPool.CSPCanNotUseStakingToken.selector);
        pool.addRewardToken(address(crossToken));
    }

    function testRewardTokenIndexMapping() public {
        vm.prank(owner);
        pool.addRewardToken(address(rewardToken3));

        // Check if token is registered
        assertTrue(pool.isRewardToken(address(rewardToken3)), "Third token should be registered");

        // Check token at index 2
        assertEq(pool.rewardTokenAt(2), address(rewardToken3), "Third token should be at index 2");
    }

    // ==================== 보상 입금 권한 테스트 ====================

    function testDirectTransferReward() public {
        // 먼저 user2가 스테이킹
        _userStake(user2, 10 ether);

        // 누구나 직접 transfer로 보상 입금 가능
        vm.startPrank(user1);
        rewardToken1.mint(user1, 100 ether);
        rewardToken1.transfer(address(pool), 100 ether);
        vm.stopPrank();

        // 검증: 보상이 입금되었는지
        uint[] memory rewards = pool.pendingRewards(user2);
        assertApproxEqAbs(rewards[0], 100 ether, 1 ether, "Anyone can transfer rewards");
    }

    // ==================== Remove Reward Token ====================

    function testRemoveRewardToken() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove reward token
        vm.prank(owner);
        pool.removeRewardToken(address(rewardToken1));

        // Verify removed from set
        assertFalse(pool.isRewardToken(address(rewardToken1)), "Should be removed from set");
        assertEq(pool.rewardTokenCount(), 1, "Count decreased"); // rewardToken2 남음

        // But can still claim existing rewards
        vm.prank(user1);
        pool.claimReward(address(rewardToken1));

        // Verify claimed
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 100 ether, 1 ether, "Can still claim after removal");
    }

    function testCannotRemoveNonExistentToken() public {
        vm.prank(owner);
        vm.expectRevert(CrossStakingPool.CSPInvalidRewardToken.selector);
        pool.removeRewardToken(address(rewardToken3));
    }

    function testOnlyRewardManagerCanRemove() public {
        vm.prank(user1);
        vm.expectRevert();
        pool.removeRewardToken(address(rewardToken1));
    }

    function testRemovedTokenNoNewRewards() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove token
        vm.prank(owner);
        pool.removeRewardToken(address(rewardToken1));

        // Try to add new rewards (won't be distributed because not in set)
        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // Sync won't happen for removed token
        _userStake(user2, 10 ether);

        // user1 should only have original 100, not 150
        vm.prank(user1);
        pool.claimReward(address(rewardToken1));
        
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 100 ether, 1 ether, "No new rewards after removal");
    }

    // ==================== Emergency Withdraw ====================

    function testEmergencyWithdraw() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove token
        vm.prank(owner);
        pool.removeRewardToken(address(rewardToken1));

        // 실수로 추가 입금
        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // 비상 출금 가능 확인
        uint withdrawable = pool.getEmergencyWithdrawableAmount(address(rewardToken1));
        assertEq(withdrawable, 50 ether, "Should be able to withdraw extra deposit");

        // 비상 출금
        uint ownerBalanceBefore = rewardToken1.balanceOf(owner);
        vm.prank(owner);
        pool.emergencyWithdraw(address(rewardToken1), owner);

        // 출금 확인
        assertEq(rewardToken1.balanceOf(owner) - ownerBalanceBefore, 50 ether, "Emergency withdraw succeeded");
        assertEq(pool.getEmergencyWithdrawableAmount(address(rewardToken1)), 0, "No more withdrawable");

        // 사용자는 여전히 claim 가능
        vm.prank(user1);
        pool.claimReward(address(rewardToken1));
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 100 ether, 1 ether, "User can still claim");
    }

    function testEmergencyWithdrawAfterUserClaim() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove token (removed balance = 100)
        vm.prank(owner);
        pool.removeRewardToken(address(rewardToken1));

        // 실수로 50 추가 입금 (total balance = 150)
        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // User claim (balance 감소, distributed amount도 차감됨)
        vm.prank(user1);
        pool.claimReward(address(rewardToken1));

        // 출금 가능 = 현재잔액 - (제거시점잔액 - claim된금액)
        uint withdrawable = pool.getEmergencyWithdrawableAmount(address(rewardToken1));
        assertEq(withdrawable, 50 ether, "Still 50 withdrawable after user claim");

        // 비상 출금
        uint ownerBalanceBefore = rewardToken1.balanceOf(owner);
        vm.prank(owner);
        pool.emergencyWithdraw(address(rewardToken1), owner);
        assertEq(rewardToken1.balanceOf(owner) - ownerBalanceBefore, 50 ether, "Withdraw only extra deposits");
    }

    function testCannotEmergencyWithdrawNonRemovedToken() public view {
        // 제거되지 않은 토큰은 출금 불가
        uint withdrawable = pool.getEmergencyWithdrawableAmount(address(rewardToken1));
        assertEq(withdrawable, 0, "Non-removed token has 0 withdrawable");
    }

    function testCannotEmergencyWithdrawZero() public {
        // Remove but no extra deposits
        vm.prank(owner);
        pool.removeRewardToken(address(rewardToken1));

        // No extra deposits
        vm.prank(owner);
        vm.expectRevert(CrossStakingPool.CSPNoWithdrawableAmount.selector);
        pool.emergencyWithdraw(address(rewardToken1), owner);
    }

    function testOnlyAdminCanEmergencyWithdraw() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        vm.prank(owner);
        pool.removeRewardToken(address(rewardToken1));

        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // Non-admin cannot withdraw
        vm.prank(user1);
        vm.expectRevert();
        pool.emergencyWithdraw(address(rewardToken1), user1);
    }

    // ==================== Pause 기능 테스트 ====================

    function testPause() public {
        vm.prank(owner);
        pool.pause();

        assertTrue(pool.paused(), "Pool should be paused");
    }

    function testUnpause() public {
        vm.startPrank(owner);
        pool.pause();
        pool.unpause();
        vm.stopPrank();

        assertFalse(pool.paused(), "Pool should be unpaused");
    }

    function testPauseOnlyByPauserRole() public {
        vm.prank(user1);
        vm.expectRevert();
        pool.pause();
    }

    function testUnpauseOnlyByPauserRole() public {
        vm.prank(owner);
        pool.pause();

        vm.prank(user1);
        vm.expectRevert();
        pool.unpause();
    }

    function testCannotStakeWhenPaused() public {
        vm.prank(owner);
        pool.pause();

        vm.startPrank(user1);
        crossToken.approve(address(pool), 10 ether);
        vm.expectRevert();
        pool.stake(10 ether);
        vm.stopPrank();
    }

    function testCannotUnstakeWhenPaused() public {
        _userStake(user1, 10 ether);

        vm.prank(owner);
        pool.pause();

        vm.prank(user1);
        vm.expectRevert();
        pool.unstake();
    }

    function testCannotClaimWhenPaused() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        vm.prank(owner);
        pool.pause();

        vm.prank(user1);
        vm.expectRevert();
        pool.claimRewards();
    }

    function testStakeAfterUnpause() public {
        vm.startPrank(owner);
        pool.pause();
        pool.unpause();
        vm.stopPrank();

        _userStake(user1, 10 ether);
        assertEq(pool.balances(user1), 10 ether, "Should be able to stake after unpause");
    }

    // ==================== 역할 기반 접근 제어 테스트 ====================

    function testOwnerHasDefaultAdminRole() public view {
        assertTrue(pool.hasRole(pool.DEFAULT_ADMIN_ROLE(), owner), "Owner should have DEFAULT_ADMIN_ROLE");
    }

    function testOwnerHasPauserRole() public view {
        assertTrue(pool.hasRole(pool.PAUSER_ROLE(), owner), "Owner should have PAUSER_ROLE");
    }

    function testOwnerHasRewardManagerRole() public view {
        assertTrue(pool.hasRole(pool.REWARD_MANAGER_ROLE(), owner), "Owner should have REWARD_MANAGER_ROLE");
    }

    function testGrantPauserRole() public {
        vm.startPrank(owner);
        bytes32 role = pool.PAUSER_ROLE();
        pool.grantRole(role, user1);
        vm.stopPrank();

        assertTrue(pool.hasRole(role, user1), "User1 should have PAUSER_ROLE");

        // user1이 pause 가능
        vm.prank(user1);
        pool.pause();

        assertTrue(pool.paused(), "User1 with PAUSER_ROLE should be able to pause");
    }

    function testGrantRewardManagerRole() public {
        vm.startPrank(owner);
        bytes32 role = pool.REWARD_MANAGER_ROLE();
        pool.grantRole(role, user1);
        vm.stopPrank();

        assertTrue(pool.hasRole(role, user1), "User1 should have REWARD_MANAGER_ROLE");

        // user1이 보상 토큰 추가 가능
        vm.prank(user1);
        pool.addRewardToken(address(rewardToken3));

        assertEq(pool.rewardTokenCount(), 3, "User1 with REWARD_MANAGER_ROLE should be able to add reward token");
    }

    function testRevokeRole() public {
        vm.startPrank(owner);
        pool.grantRole(pool.PAUSER_ROLE(), user1);
        pool.revokeRole(pool.PAUSER_ROLE(), user1);
        vm.stopPrank();

        assertFalse(pool.hasRole(pool.PAUSER_ROLE(), user1), "User1 should not have PAUSER_ROLE");

        vm.prank(user1);
        vm.expectRevert();
        pool.pause();
    }

    // ==================== UUPS 업그레이드 권한 테스트 ====================

    function testUpgradeAuthorization() public view {
        // owner는 DEFAULT_ADMIN_ROLE을 가지고 있으므로 업그레이드 가능
        assertTrue(pool.hasRole(pool.DEFAULT_ADMIN_ROLE(), owner), "Owner should be able to upgrade");
    }

    function testNonAdminCannotUpgrade() public view {
        // user1은 DEFAULT_ADMIN_ROLE이 없으므로 업그레이드 불가
        assertFalse(pool.hasRole(pool.DEFAULT_ADMIN_ROLE(), user1), "User1 should not be able to upgrade");
    }

    // ==================== 초기화 테스트 ====================

    function testInitialConfiguration() public view {
        assertEq(address(pool.stakingToken()), address(crossToken), "Staking token should be set");
        assertEq(pool.rewardTokenCount(), 2, "Should have 2 reward tokens from setup");
        assertFalse(pool.paused(), "Should not be paused initially");
    }
}
