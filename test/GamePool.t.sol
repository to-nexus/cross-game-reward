// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "./base/GamePoolBase.t.sol";

/**
 * @title GamePoolTest
 * @notice Tests for basic GamePool functionality (deposit, withdraw, claim)
 */
contract GamePoolTest is GamePoolBase {
    // ==================== Initialization Tests ====================

    function test_Initialize_Success() public view {
        assertEq(address(gamePool.depositToken()), address(gameToken));
        assertEq(address(gamePool.rewardToken()), address(crossdToken));
        assertEq(gamePool.minDepositAmount(), MIN_DEPOSIT);
        assertEq(uint(gamePool.poolStatus()), uint(ICrossGameRewardPool.PoolStatus.Active));
        assertEq(gamePool.nextRoundId(), 1);
    }

    function test_Initialize_PoolType() public view {
        assertEq(uint(crossGameReward.getPoolType(poolId)), uint(ICrossGameReward.PoolType.GamePool));
    }

    function test_Initialize_SponsorRole() public view {
        assertTrue(gamePool.hasRole(gamePool.SPONSOR_ROLE(), sponsor));
        assertFalse(gamePool.hasRole(gamePool.SPONSOR_ROLE(), user1));
    }

    // ==================== Deposit Tests ====================

    function test_Deposit_Success() public {
        uint depositAmount = 100 ether;

        _userDeposit(user1, depositAmount);

        assertEq(gamePool.balances(user1), depositAmount);
        assertEq(gamePool.totalDeposited(), depositAmount);
    }

    function test_Deposit_MultipleUsers() public {
        _userDeposit(user1, 100 ether);
        _userDeposit(user2, 200 ether);
        _userDeposit(user3, 300 ether);

        assertEq(gamePool.balances(user1), 100 ether);
        assertEq(gamePool.balances(user2), 200 ether);
        assertEq(gamePool.balances(user3), 300 ether);
        assertEq(gamePool.totalDeposited(), 600 ether);
    }

    function test_Deposit_BelowMinimum_Reverts() public {
        vm.startPrank(user1);
        gameToken.approve(address(gamePool), MIN_DEPOSIT - 1);

        vm.expectRevert(
            abi.encodeWithSelector(
                GamePool.GPBelowMinimumDepositAmount.selector, MIN_DEPOSIT - 1, MIN_DEPOSIT
            )
        );
        gamePool.deposit(MIN_DEPOSIT - 1);
        vm.stopPrank();
    }

    function test_Deposit_AdditionalDeposit() public {
        _userDeposit(user1, 100 ether);
        _userDeposit(user1, 50 ether);

        assertEq(gamePool.balances(user1), 150 ether);
        assertEq(gamePool.totalDeposited(), 150 ether);
    }

    // ==================== Withdraw Tests ====================

    function test_Withdraw_Full() public {
        _userDeposit(user1, 100 ether);

        uint balanceBefore = gameToken.balanceOf(user1);

        vm.prank(user1);
        gamePool.withdraw(0); // 0 means withdraw all

        assertEq(gamePool.balances(user1), 0);
        assertEq(gamePool.totalDeposited(), 0);
        assertEq(gameToken.balanceOf(user1), balanceBefore + 100 ether);
    }

    function test_Withdraw_Partial() public {
        _userDeposit(user1, 100 ether);

        vm.prank(user1);
        gamePool.withdraw(40 ether);

        assertEq(gamePool.balances(user1), 60 ether);
        assertEq(gamePool.totalDeposited(), 60 ether);
    }

    function test_Withdraw_NoDeposit_Reverts() public {
        vm.prank(user1);
        vm.expectRevert(abi.encodeWithSelector(GamePool.GPNoDepositFound.selector, user1));
        gamePool.withdraw(0);
    }

    function test_Withdraw_InsufficientBalance_Reverts() public {
        _userDeposit(user1, 100 ether);

        vm.prank(user1);
        vm.expectRevert(
            abi.encodeWithSelector(GamePool.GPInsufficientBalance.selector, 100 ether, 150 ether)
        );
        gamePool.withdraw(150 ether);
    }

    // ==================== Claim Tests (without rounds) ====================

    function test_ClaimRewards_NoRewards() public {
        _userDeposit(user1, 100 ether);

        // No rounds created, so no rewards to claim
        vm.prank(user1);
        gamePool.claimRewards();

        assertEq(crossdToken.balanceOf(user1), 0);
    }

    // ==================== View Function Tests ====================

    function test_GetRewardTokens() public view {
        address[] memory tokens = gamePool.getRewardTokens();
        assertEq(tokens.length, 1);
        assertEq(tokens[0], address(crossdToken));
    }

    function test_RewardTokenCount() public view {
        assertEq(gamePool.rewardTokenCount(), 1);
    }

    function test_IsRewardToken() public view {
        assertTrue(gamePool.isRewardToken(crossdToken));
        assertFalse(gamePool.isRewardToken(gameToken));
    }

    function test_GetRemovedRewardTokens() public view {
        address[] memory removed = gamePool.getRemovedRewardTokens();
        assertEq(removed.length, 0);
    }

    function test_PendingRewards_Format() public {
        _userDeposit(user1, 100 ether);

        (address[] memory tokens, uint[] memory rewards) = gamePool.pendingRewards(user1);

        assertEq(tokens.length, 1);
        assertEq(tokens[0], address(crossdToken));
        assertEq(rewards.length, 1);
        assertEq(rewards[0], 0); // No rounds, no rewards
    }

    // ==================== Pool Status Tests ====================

    function test_SetPoolStatus_Inactive() public {
        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Inactive);

        assertEq(uint(gamePool.poolStatus()), uint(ICrossGameRewardPool.PoolStatus.Inactive));

        // Deposit should fail
        vm.startPrank(user1);
        gameToken.approve(address(gamePool), 100 ether);
        vm.expectRevert(
            abi.encodeWithSelector(
                GamePool.GPDepositNotAllowed.selector, ICrossGameRewardPool.PoolStatus.Inactive
            )
        );
        gamePool.deposit(100 ether);
        vm.stopPrank();
    }

    function test_SetPoolStatus_Paused() public {
        _userDeposit(user1, 100 ether);

        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Paused);

        // Withdraw should fail when paused
        vm.prank(user1);
        vm.expectRevert();
        gamePool.withdraw(0);
    }

    // ==================== Min Deposit Amount Tests ====================

    function test_UpdateMinDepositAmount() public {
        uint newMin = 5 ether;
        crossGameReward.updateMinDepositAmount(poolId, newMin);

        assertEq(gamePool.minDepositAmount(), newMin);
    }

    // ==================== Sponsor Role Tests ====================

    function test_GrantSponsorRole() public {
        address newSponsor = address(0x5000);

        crossGameReward.grantSponsorRole(poolId, newSponsor);

        assertTrue(gamePool.hasRole(gamePool.SPONSOR_ROLE(), newSponsor));
    }

    function test_RevokeSponsorRole() public {
        crossGameReward.revokeSponsorRole(poolId, sponsor);

        assertFalse(gamePool.hasRole(gamePool.SPONSOR_ROLE(), sponsor));
    }

    function test_SponsorRole_OnlyForGamePool() public {
        (uint v1PoolId,) = crossGameReward.createPool("V1 Pool", IERC20(address(gameToken)), MIN_DEPOSIT);

        vm.expectRevert(
            abi.encodeWithSelector(
                CrossGameReward.CGRInvalidPoolType.selector,
                v1PoolId,
                ICrossGameReward.PoolType.GamePool,
                ICrossGameReward.PoolType.CrossPool
            )
        );
        crossGameReward.grantSponsorRole(v1PoolId, sponsor);
    }
}
