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
        crossStaking.addRewardToken(1, address(rewardToken3));

        assertEq(pool.rewardTokenCount(), 3, "Should have 3 reward tokens");
        assertTrue(pool.isRewardToken(address(rewardToken3)), "Should be registered as reward token");
    }

    function testAddRewardTokenOnlyByManager() public {
        vm.prank(user1);
        vm.expectRevert();
        pool.addRewardToken(address(rewardToken3));
    }

    function testCannotAddSameRewardTokenTwice() public {
        crossStaking.addRewardToken(1, address(rewardToken3));

        vm.expectRevert(CrossStakingPool.CSPRewardTokenAlreadyAdded.selector);
        crossStaking.addRewardToken(1, address(rewardToken3));
    }

    function testCannotAddZeroAddressAsRewardToken() public {
        vm.expectRevert(CrossStakingPool.CSPCanNotZeroAddress.selector);
        crossStaking.addRewardToken(1, address(0));
    }

    function testCannotAddStakingTokenAsReward() public {
        vm.expectRevert(CrossStakingPool.CSPCanNotUseStakingToken.selector);
        crossStaking.addRewardToken(1, address(crossToken));
    }

    function testRewardTokenIndexMapping() public {
        crossStaking.addRewardToken(1, address(rewardToken3));

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

        // Remove reward token via CrossStaking
        crossStaking.removeRewardToken(1, address(rewardToken1));

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
        vm.expectRevert(CrossStakingPool.CSPInvalidRewardToken.selector);
        crossStaking.removeRewardToken(1, address(rewardToken3));
    }

    function testOnlyRewardManagerCanRemove() public {
        vm.prank(user1);
        vm.expectRevert();
        crossStaking.removeRewardToken(1, address(rewardToken1));
    }

    function testRemovedTokenNoNewRewards() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Remove token via CrossStaking
        crossStaking.removeRewardToken(1, address(rewardToken1));

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

        // Remove token via CrossStaking
        crossStaking.removeRewardToken(1, address(rewardToken1));

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

        // Remove token (removed balance = 100) via CrossStaking
        crossStaking.removeRewardToken(1, address(rewardToken1));

        // 실수로 50 추가 입금 (total balance = 150)
        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // User claim (balance 감소, distributed amount도 차감됨)
        vm.prank(user1);
        pool.claimReward(address(rewardToken1));

        // 출금 가능 = 현재잔액 - (제거시점잔액 - claim된금액)
        uint withdrawable = pool.getEmergencyWithdrawableAmount(address(rewardToken1));
        assertEq(withdrawable, 50 ether, "Still 50 withdrawable after user claim");

        // 비상 출금 (owner가 DEFAULT_ADMIN_ROLE을 가지고 있으므로 직접 호출 가능)
        uint ownerBalanceBefore = rewardToken1.balanceOf(owner);
        pool.emergencyWithdraw(address(rewardToken1), owner);
        assertEq(rewardToken1.balanceOf(owner) - ownerBalanceBefore, 50 ether, "Withdraw only extra deposits");
    }

    function testCannotEmergencyWithdrawNonRemovedToken() public view {
        // 제거되지 않은 토큰은 출금 불가
        uint withdrawable = pool.getEmergencyWithdrawableAmount(address(rewardToken1));
        assertEq(withdrawable, 0, "Non-removed token has 0 withdrawable");
    }

    function testCannotEmergencyWithdrawZero() public {
        // Remove but no extra deposits via CrossStaking
        crossStaking.removeRewardToken(1, address(rewardToken1));

        // No extra deposits - owner can call emergencyWithdraw directly
        vm.expectRevert(CrossStakingPool.CSPNoWithdrawableAmount.selector);
        pool.emergencyWithdraw(address(rewardToken1), owner);
    }

    function testOnlyAdminCanEmergencyWithdraw() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        crossStaking.removeRewardToken(1, address(rewardToken1));

        rewardToken1.mint(owner, 50 ether);
        rewardToken1.transfer(address(pool), 50 ether);

        // Non-admin cannot withdraw
        vm.prank(user1);
        vm.expectRevert();
        pool.emergencyWithdraw(address(rewardToken1), user1);
    }

    // ==================== Pause 기능 테스트 ====================

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

    // ==================== 역할 기반 접근 제어 테스트 ====================

    function testOwnerHasDefaultAdminRole() public view {
        // owner는 CrossStaking 컨트랙트
        assertTrue(pool.hasRole(pool.DEFAULT_ADMIN_ROLE(), owner), "CrossStaking should have DEFAULT_ADMIN_ROLE");
    }

    function testCrossStakingHasStakingRootRole() public view {
        // CrossStaking 컨트랙트가 STAKING_ROOT_ROLE을 가짐
        assertTrue(
            pool.hasRole(pool.STAKING_ROOT_ROLE(), address(crossStaking)), "CrossStaking should have STAKING_ROOT_ROLE"
        );
    }

    function testOnlyStakingRootCanPause() public {
        // user1은 STAKING_ROOT_ROLE이 없으므로 pause 불가
        vm.prank(user1);
        vm.expectRevert();
        pool.pause();

        // CrossStaking만 pause 가능 (setPoolActive를 통해)
        crossStaking.setPoolActive(1, false);
        assertTrue(pool.paused(), "CrossStaking with STAKING_ROOT_ROLE can pause");
    }

    function testOnlyStakingRootCanAddRewardToken() public {
        // user1은 STAKING_ROOT_ROLE이 없으므로 보상 토큰 추가 불가
        vm.prank(user1);
        vm.expectRevert();
        pool.addRewardToken(address(rewardToken3));

        // CrossStaking을 통해 추가 가능
        crossStaking.addRewardToken(1, address(rewardToken3));
        assertEq(pool.rewardTokenCount(), 3, "CrossStaking with STAKING_ROOT_ROLE can add reward token");
    }

    function testGrantStakingRootRole() public {
        // Owner(DEFAULT_ADMIN)가 다른 주소에 STAKING_ROOT_ROLE 부여 가능
        bytes32 role = pool.STAKING_ROOT_ROLE();
        pool.grantRole(role, user1);

        assertTrue(pool.hasRole(role, user1), "User1 should have STAKING_ROOT_ROLE");

        // user1이 pause 가능
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
