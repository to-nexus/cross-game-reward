// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "./base/CrossGameRewardPoolV2Base.t.sol";

/**
 * @title CrossGameRewardPoolV2Test
 * @notice Tests for basic V2 pool functionality (deposit, withdraw, claim)
 */
contract CrossGameRewardPoolV2Test is CrossGameRewardPoolV2Base {
    // ==================== Initialization Tests ====================

    function test_Initialize_Success() public view {
        assertEq(address(poolV2.depositToken()), address(gameToken));
        assertEq(address(poolV2.rewardToken()), address(crossdToken));
        assertEq(poolV2.minDepositAmount(), MIN_DEPOSIT);
        assertEq(uint(poolV2.poolStatus()), uint(ICrossGameRewardPool.PoolStatus.Active));
        assertEq(poolV2.nextRoundId(), 1);
    }

    function test_Initialize_PoolType() public view {
        assertEq(uint(crossGameReward.getPoolType(poolId)), uint(ICrossGameReward.PoolType.GamePool));
    }

    function test_Initialize_DeveloperRole() public view {
        assertTrue(poolV2.hasDeveloperRole(developer));
        assertFalse(poolV2.hasDeveloperRole(user1));
    }

    // ==================== Deposit Tests ====================

    function test_Deposit_Success() public {
        uint depositAmount = 100 ether;

        _userDeposit(user1, depositAmount);

        assertEq(poolV2.balances(user1), depositAmount);
        assertEq(poolV2.totalDeposited(), depositAmount);
    }

    function test_Deposit_MultipleUsers() public {
        _userDeposit(user1, 100 ether);
        _userDeposit(user2, 200 ether);
        _userDeposit(user3, 300 ether);

        assertEq(poolV2.balances(user1), 100 ether);
        assertEq(poolV2.balances(user2), 200 ether);
        assertEq(poolV2.balances(user3), 300 ether);
        assertEq(poolV2.totalDeposited(), 600 ether);
    }

    function test_Deposit_BelowMinimum_Reverts() public {
        vm.startPrank(user1);
        gameToken.approve(address(poolV2), MIN_DEPOSIT - 1);

        vm.expectRevert(
            abi.encodeWithSelector(
                CrossGameRewardPoolV2.CGRP2BelowMinimumDepositAmount.selector, MIN_DEPOSIT - 1, MIN_DEPOSIT
            )
        );
        poolV2.deposit(MIN_DEPOSIT - 1);
        vm.stopPrank();
    }

    function test_Deposit_AdditionalDeposit() public {
        _userDeposit(user1, 100 ether);
        _userDeposit(user1, 50 ether);

        assertEq(poolV2.balances(user1), 150 ether);
        assertEq(poolV2.totalDeposited(), 150 ether);
    }

    // ==================== Withdraw Tests ====================

    function test_Withdraw_Full() public {
        _userDeposit(user1, 100 ether);

        uint balanceBefore = gameToken.balanceOf(user1);

        vm.prank(user1);
        poolV2.withdraw(0); // 0 means withdraw all

        assertEq(poolV2.balances(user1), 0);
        assertEq(poolV2.totalDeposited(), 0);
        assertEq(gameToken.balanceOf(user1), balanceBefore + 100 ether);
    }

    function test_Withdraw_Partial() public {
        _userDeposit(user1, 100 ether);

        vm.prank(user1);
        poolV2.withdraw(40 ether);

        assertEq(poolV2.balances(user1), 60 ether);
        assertEq(poolV2.totalDeposited(), 60 ether);
    }

    function test_Withdraw_NoDeposit_Reverts() public {
        vm.prank(user1);
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPoolV2.CGRP2NoDepositFound.selector, user1));
        poolV2.withdraw(0);
    }

    function test_Withdraw_InsufficientBalance_Reverts() public {
        _userDeposit(user1, 100 ether);

        vm.prank(user1);
        vm.expectRevert(
            abi.encodeWithSelector(CrossGameRewardPoolV2.CGRP2InsufficientBalance.selector, 100 ether, 150 ether)
        );
        poolV2.withdraw(150 ether);
    }

    // ==================== Claim Tests (without rounds) ====================

    function test_ClaimRewards_NoRewards() public {
        _userDeposit(user1, 100 ether);

        // No rounds created, so no rewards to claim
        vm.prank(user1);
        poolV2.claimRewards();

        assertEq(crossdToken.balanceOf(user1), 0);
    }

    // ==================== View Function Tests ====================

    function test_GetRewardTokens() public view {
        address[] memory tokens = poolV2.getRewardTokens();
        assertEq(tokens.length, 1);
        assertEq(tokens[0], address(crossdToken));
    }

    function test_RewardTokenCount() public view {
        assertEq(poolV2.rewardTokenCount(), 1);
    }

    function test_IsRewardToken() public view {
        assertTrue(poolV2.isRewardToken(crossdToken));
        assertFalse(poolV2.isRewardToken(gameToken));
    }

    function test_GetRemovedRewardTokens() public view {
        address[] memory removed = poolV2.getRemovedRewardTokens();
        assertEq(removed.length, 0);
    }

    function test_PendingRewards_Format() public {
        _userDeposit(user1, 100 ether);

        (address[] memory tokens, uint[] memory rewards) = poolV2.pendingRewards(user1);

        assertEq(tokens.length, 1);
        assertEq(tokens[0], address(crossdToken));
        assertEq(rewards.length, 1);
        assertEq(rewards[0], 0); // No rounds, no rewards
    }

    // ==================== Pool Status Tests ====================

    function test_SetPoolStatus_Inactive() public {
        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Inactive);

        assertEq(uint(poolV2.poolStatus()), uint(ICrossGameRewardPool.PoolStatus.Inactive));

        // Deposit should fail
        vm.startPrank(user1);
        gameToken.approve(address(poolV2), 100 ether);
        vm.expectRevert(
            abi.encodeWithSelector(
                CrossGameRewardPoolV2.CGRP2DepositNotAllowed.selector, ICrossGameRewardPool.PoolStatus.Inactive
            )
        );
        poolV2.deposit(100 ether);
        vm.stopPrank();
    }

    function test_SetPoolStatus_Paused() public {
        _userDeposit(user1, 100 ether);

        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Paused);

        // Withdraw should fail when paused
        vm.prank(user1);
        vm.expectRevert();
        poolV2.withdraw(0);
    }

    // ==================== Min Deposit Amount Tests ====================

    function test_UpdateMinDepositAmount() public {
        uint newMin = 5 ether;
        crossGameReward.updateMinDepositAmount(poolId, newMin);

        assertEq(poolV2.minDepositAmount(), newMin);
    }

    // ==================== Developer Role Tests ====================

    function test_GrantDeveloperRole() public {
        address newDev = address(0x5000);

        crossGameReward.grantDeveloperRole(poolId, newDev);

        assertTrue(poolV2.hasDeveloperRole(newDev));
    }

    function test_RevokeDeveloperRole() public {
        crossGameReward.revokeDeveloperRole(poolId, developer);

        assertFalse(poolV2.hasDeveloperRole(developer));
    }

    function test_DeveloperRole_OnlyForV2Pool() public {
        // Create a V1 pool
        (uint v1PoolId,) = crossGameReward.createPool("V1 Pool", IERC20(address(gameToken)), MIN_DEPOSIT);

        // Trying to grant developer role to V1 pool should fail
        vm.expectRevert(
            abi.encodeWithSelector(
                CrossGameReward.CGRInvalidPoolType.selector,
                v1PoolId,
                ICrossGameReward.PoolType.GamePool,
                ICrossGameReward.PoolType.CrossPool
            )
        );
        crossGameReward.grantDeveloperRole(v1PoolId, developer);
    }
}
