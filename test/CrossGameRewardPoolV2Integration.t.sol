// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "./base/CrossGameRewardPoolV2Base.t.sol";
import "../src/CrossGameRewardRouter.sol";

/**
 * @title CrossGameRewardPoolV2IntegrationTest
 * @notice Integration tests for V2 pool with Factory and Router
 */
contract CrossGameRewardPoolV2IntegrationTest is CrossGameRewardPoolV2Base {
    CrossGameRewardRouter public router;

    function setUp() public override {
        super.setUp();

        // Deploy and setup router
        router = new CrossGameRewardRouter(address(crossGameReward));
        crossGameReward.setRouter(address(router));
    }

    // ==================== Factory Integration Tests ====================

    function test_Factory_CreateBothPoolTypes() public {
        // V1 pool already exists from base setup

        // Create V1 pool
        (uint256 v1PoolId,) =
            crossGameReward.createPool("V1 Test Pool", IERC20(address(gameToken)), MIN_DEPOSIT);

        // V2 pool already created in setUp
        // Verify both exist
        assertEq(uint256(crossGameReward.getPoolType(v1PoolId)), uint256(ICrossGameReward.PoolType.CrossPool));
        assertEq(uint256(crossGameReward.getPoolType(poolId)), uint256(ICrossGameReward.PoolType.GamePool));
    }

    function test_Factory_GetPoolIdsByType() public {
        // Create additional pools
        (uint256 v1Pool1,) = crossGameReward.createPool("V1 Pool 1", IERC20(address(gameToken)), MIN_DEPOSIT);

        MockERC20V2 anotherGameToken = new MockERC20V2("Another Game", "AGAME");
        MockERC20V2 anotherReward = new MockERC20V2("Another Reward", "ARWD");

        (uint256 v2Pool2,) = crossGameReward.createPoolV2(
            "V2 Pool 2", IERC20(address(anotherGameToken)), IERC20(address(anotherReward)), MIN_DEPOSIT
        );

        // Get V1 pools
        uint256[] memory v1Pools = crossGameReward.getPoolIdsByType(ICrossGameReward.PoolType.CrossPool);
        assertEq(v1Pools.length, 1);
        assertEq(v1Pools[0], v1Pool1);

        // Get V2 pools (includes the one from setUp)
        uint256[] memory v2Pools = crossGameReward.getPoolIdsByType(ICrossGameReward.PoolType.GamePool);
        assertEq(v2Pools.length, 2);
    }

    function test_Factory_TotalPoolCount() public {
        // Already have 1 V2 pool from setUp
        crossGameReward.createPool("V1 Pool", IERC20(address(gameToken)), MIN_DEPOSIT);

        assertEq(crossGameReward.getTotalPoolCount(), 2);
    }

    // ==================== Router Integration Tests ====================

    function test_Router_DepositERC20_V2Pool() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        assertEq(poolV2.balances(user1), 100 ether);
    }

    function test_Router_WithdrawERC20_V2Pool() public {
        // First deposit
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        uint256 balanceBefore = gameToken.balanceOf(user1);

        // Then withdraw
        vm.prank(user1);
        router.withdrawERC20(poolId, 0);

        assertEq(poolV2.balances(user1), 0);
        assertEq(gameToken.balanceOf(user1), balanceBefore + 100 ether);
    }

    function test_Router_ClaimRewards_V2Pool() public {
        // Setup: deposit and create round
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        uint256 roundId = _createRound(10000 ether, 10, 100);
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(100);

        uint256 balanceBefore = crossdToken.balanceOf(user1);

        vm.prank(user1);
        router.claimRewards(poolId);

        assertEq(crossdToken.balanceOf(user1), balanceBefore + 10000 ether);
    }

    function test_Router_GetUserDepositInfo_V2Pool() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        (uint256 deposited, address[] memory tokens, uint256[] memory rewards) =
            router.getUserDepositInfo(poolId, user1);

        assertEq(deposited, 100 ether);
        assertEq(tokens.length, 1);
        assertEq(tokens[0], address(crossdToken));
        assertEq(rewards[0], 0); // No rewards yet
    }

    function test_Router_GetPendingRewards_V2Pool() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        uint256 roundId = _createRound(10000 ether, 10, 100);
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(50);

        (address[] memory tokens, uint256[] memory rewards) = router.getPendingRewards(poolId, user1);

        assertEq(tokens.length, 1);
        assertEq(tokens[0], address(crossdToken));
        assertEq(rewards[0], 5000 ether);
    }

    function test_Router_GetAllPendingRewards_V2Pool() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        uint256 roundId = _createRound(10000 ether, 10, 100);
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(50);

        (address[] memory tokens, uint256[] memory rewards) = router.getAllPendingRewards(poolId, user1);

        assertEq(tokens.length, 1);
        assertEq(rewards[0], 5000 ether);
    }

    function test_Router_GetRemovedTokenRewards_V2Pool_Empty() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        // V2 doesn't support removed tokens, should return empty
        (address[] memory tokens, uint256[] memory rewards) = router.getRemovedTokenRewards(poolId, user1);

        assertEq(tokens.length, 0);
        assertEq(rewards.length, 0);
    }

    function test_Router_IsNativePool_V2Pool() public view {
        // V2 pool is not a native pool (doesn't use WCROSS)
        assertFalse(router.isNativePool(poolId));
    }

    function test_Router_GetTotalDeposited_V2Pool() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        vm.startPrank(user2);
        gameToken.approve(address(router), 200 ether);
        router.depositERC20(poolId, 200 ether);
        vm.stopPrank();

        uint256 total = router.getTotalDeposited(address(gameToken));
        assertEq(total, 300 ether);
    }

    // ==================== Mixed Pool Type Tests ====================

    function test_MixedPools_IndependentOperation() public {
        // Create V1 pool
        MockERC20V2 v1DepositToken = new MockERC20V2("V1 Deposit", "V1D");
        MockERC20V2 v1RewardToken = new MockERC20V2("V1 Reward", "V1R");

        (uint256 v1PoolId, ICrossGameRewardPool v1Pool) =
            crossGameReward.createPool("V1 Pool", IERC20(address(v1DepositToken)), MIN_DEPOSIT);

        // Add reward token to V1 pool
        crossGameReward.addRewardToken(v1PoolId, v1RewardToken);

        // Distribute tokens
        v1DepositToken.transfer(user1, 1000 ether);

        // User deposits in both pools
        vm.startPrank(user1);

        // Deposit in V1 pool
        v1DepositToken.approve(address(v1Pool), 100 ether);
        v1Pool.deposit(100 ether);

        // Deposit in V2 pool via router
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);

        vm.stopPrank();

        // Verify independent balances
        assertEq(v1Pool.balances(user1), 100 ether);
        assertEq(poolV2.balances(user1), 100 ether);
    }

    // ==================== Edge Case Tests ====================

    function test_V2Pool_RouterWithdrawAll() public {
        // Create V1 pool
        MockERC20V2 v1DepositToken = new MockERC20V2("V1 Deposit", "V1D");
        (uint256 v1PoolId, ICrossGameRewardPool v1Pool) =
            crossGameReward.createPool("V1 Pool", IERC20(address(v1DepositToken)), MIN_DEPOSIT);

        // Distribute tokens
        v1DepositToken.transfer(user1, 1000 ether);

        // User deposits in both pools via router
        vm.startPrank(user1);

        v1DepositToken.approve(address(router), 100 ether);
        router.depositERC20(v1PoolId, 100 ether);

        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);

        // Withdraw all from all pools
        router.withdrawAll();

        vm.stopPrank();

        assertEq(v1Pool.balances(user1), 0);
        assertEq(poolV2.balances(user1), 0);
    }

    function test_V2Pool_ClaimSpecificToken() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        uint256 roundId = _createRound(10000 ether, 10, 100);
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(100);

        uint256 balanceBefore = crossdToken.balanceOf(user1);

        // Claim specific token via router
        vm.prank(user1);
        router.claimReward(poolId, address(crossdToken));

        assertEq(crossdToken.balanceOf(user1), balanceBefore + 10000 ether);
    }

    // ==================== Pool Status Tests via Factory ====================

    function test_Factory_SetPoolStatus_V2() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        // Set to Inactive
        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Inactive);

        // Should not be able to deposit
        vm.startPrank(user2);
        gameToken.approve(address(router), 100 ether);
        vm.expectRevert();
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        // But should be able to withdraw
        vm.prank(user1);
        router.withdrawERC20(poolId, 0);

        assertEq(poolV2.balances(user1), 0);
    }

    // ==================== Reclaim via Factory ====================

    function test_Factory_ReclaimFromV2Pool() public {
        // Create round with no depositors
        uint256 roundId = _createRound(10000 ether, 10, 100);
        ICrossGameRewardPoolV2.Round memory round = poolV2.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(100);

        // Trigger pool update
        _userDeposit(user1, 100 ether);

        // Reclaim via factory
        address recipient = address(0x8888);
        crossGameReward.reclaimFromPool(poolId, crossdToken, recipient);

        assertEq(crossdToken.balanceOf(recipient), 10000 ether);
    }
}
