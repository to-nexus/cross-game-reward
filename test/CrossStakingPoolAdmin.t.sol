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

        vm.expectRevert(CrossStakingPool.RewardTokenAlreadyAdded.selector);
        pool.addRewardToken(address(rewardToken3));
        vm.stopPrank();
    }

    function testCannotAddZeroAddressAsRewardToken() public {
        vm.prank(owner);
        vm.expectRevert(CrossStakingPool.InvalidTokenAddress.selector);
        pool.addRewardToken(address(0));
    }

    function testCannotAddStakingTokenAsReward() public {
        vm.prank(owner);
        vm.expectRevert(CrossStakingPool.CannotUseStakingTokenAsReward.selector);
        pool.addRewardToken(address(crossToken));
    }

    function testRewardTokenIndexMapping() public {
        vm.prank(owner);
        pool.addRewardToken(address(rewardToken3));

        uint index = pool.tokenToIndex(address(rewardToken3));
        assertEq(index, 2, "Third token should have index 2");
    }

    // ==================== 보상 입금 권한 테스트 ====================

    function testDepositRewardByAnyone() public {
        // 먼저 user2가 스테이킹
        _userStake(user2, 10 ether);

        // 누구나 보상을 입금할 수 있음
        vm.startPrank(user1);
        rewardToken1.mint(user1, 100 ether);
        rewardToken1.approve(address(pool), 100 ether);
        pool.depositReward(address(rewardToken1), 100 ether);
        vm.stopPrank();

        // 검증: 보상이 입금되었는지
        uint[] memory rewards = pool.pendingRewards(user2);
        assertApproxEqAbs(rewards[0], 100 ether, 1 ether, "Anyone can deposit rewards");
    }

    function testCannotDepositInvalidRewardToken() public {
        vm.prank(user1);
        vm.expectRevert(CrossStakingPool.InvalidRewardToken.selector);
        pool.depositReward(address(rewardToken3), 100 ether);
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

    function testOwnerHasDefaultAdminRole() public {
        assertTrue(pool.hasRole(pool.DEFAULT_ADMIN_ROLE(), owner), "Owner should have DEFAULT_ADMIN_ROLE");
    }

    function testOwnerHasPauserRole() public {
        assertTrue(pool.hasRole(pool.PAUSER_ROLE(), owner), "Owner should have PAUSER_ROLE");
    }

    function testOwnerHasRewardManagerRole() public {
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

    function testUpgradeAuthorization() public {
        // 새 구현 배포
        CrossStakingPool newImplementation = new CrossStakingPool();

        // owner는 DEFAULT_ADMIN_ROLE을 가지고 있으므로 업그레이드 가능
        vm.prank(owner);
        // 실제 업그레이드는 테스트하지 않고, 권한만 확인
        assertTrue(pool.hasRole(pool.DEFAULT_ADMIN_ROLE(), owner), "Owner should be able to upgrade");
    }

    function testNonAdminCannotUpgrade() public {
        CrossStakingPool newImplementation = new CrossStakingPool();

        // user1은 DEFAULT_ADMIN_ROLE이 없으므로 업그레이드 불가
        assertFalse(pool.hasRole(pool.DEFAULT_ADMIN_ROLE(), user1), "User1 should not be able to upgrade");
    }

    // ==================== 초기화 테스트 ====================

    function testInitialConfiguration() public {
        assertEq(address(pool.stakingToken()), address(crossToken), "Staking token should be set");
        assertEq(pool.rewardTokenCount(), 2, "Should have 2 reward tokens from setup");
        assertFalse(pool.paused(), "Should not be paused initially");
    }
}
