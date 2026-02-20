// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "./base/GamePoolBase.t.sol";
import "../src/CrossGameRewardRouter.sol";

/**
 * @title GamePoolIntegrationTest
 * @notice Integration tests for GamePool with Factory and Router
 */
contract GamePoolIntegrationTest is GamePoolBase {
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

        // GamePool already created in setUp
        // Verify both exist
        assertEq(uint256(crossGameReward.getPoolType(v1PoolId)), uint256(ICrossGameReward.PoolType.CrossPool));
        assertEq(uint256(crossGameReward.getPoolType(poolId)), uint256(ICrossGameReward.PoolType.GamePool));
    }

    function test_Factory_GetPoolIdsByType() public {
        // Create additional pools
        (uint256 v1Pool1,) = crossGameReward.createPool("V1 Pool 1", IERC20(address(gameToken)), MIN_DEPOSIT);

        MockERC20GP anotherGameToken = new MockERC20GP("Another Game", "AGAME");
        MockERC20GP anotherReward = new MockERC20GP("Another Reward", "ARWD");

        (uint256 gp2,) = crossGameReward.createGamePool(
            "Game Pool 2", IERC20(address(anotherGameToken)), IERC20(address(anotherReward)), MIN_DEPOSIT
        );

        // Get V1 pools
        uint256[] memory v1Pools = crossGameReward.getPoolIdsByType(ICrossGameReward.PoolType.CrossPool);
        assertEq(v1Pools.length, 1);
        assertEq(v1Pools[0], v1Pool1);

        // Get GamePools (includes the one from setUp)
        uint256[] memory gamePools = crossGameReward.getPoolIdsByType(ICrossGameReward.PoolType.GamePool);
        assertEq(gamePools.length, 2);
    }

    function test_Factory_TotalPoolCount() public {
        // Already have 1 GamePool from setUp
        crossGameReward.createPool("V1 Pool", IERC20(address(gameToken)), MIN_DEPOSIT);

        assertEq(crossGameReward.getTotalPoolCount(), 2);
    }

    // ==================== Router Integration Tests ====================

    function test_Router_DepositERC20_GamePool() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        assertEq(gamePool.balances(user1), 100 ether);
    }

    function test_Router_WithdrawERC20_GamePool() public {
        // First deposit
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        uint256 balanceBefore = gameToken.balanceOf(user1);

        // Then withdraw
        vm.prank(user1);
        router.withdrawERC20(poolId, 0);

        assertEq(gamePool.balances(user1), 0);
        assertEq(gameToken.balanceOf(user1), balanceBefore + 100 ether);
    }

    function test_Router_ClaimRewards_GamePool() public {
        // Setup: deposit and create round
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        uint256 roundId = _createRound(10000 ether, 10, 100);
        IGamePool.Round memory round = gamePool.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(100);

        uint256 balanceBefore = crossdToken.balanceOf(user1);

        vm.prank(user1);
        router.claimRewards(poolId);

        assertEq(crossdToken.balanceOf(user1), balanceBefore + 10000 ether);
    }

    function test_Router_GetUserDepositInfo_GamePool() public {
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

    function test_Router_GetPendingRewards_GamePool() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        uint256 roundId = _createRound(10000 ether, 10, 100);
        IGamePool.Round memory round = gamePool.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(50);

        (address[] memory tokens, uint256[] memory rewards) = router.getPendingRewards(poolId, user1);

        assertEq(tokens.length, 1);
        assertEq(tokens[0], address(crossdToken));
        assertEq(rewards[0], 5000 ether);
    }

    function test_Router_GetAllPendingRewards_GamePool() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        uint256 roundId = _createRound(10000 ether, 10, 100);
        IGamePool.Round memory round = gamePool.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(50);

        (address[] memory tokens, uint256[] memory rewards) = router.getAllPendingRewards(poolId, user1);

        assertEq(tokens.length, 1);
        assertEq(rewards[0], 5000 ether);
    }

    function test_Router_GetRemovedTokenRewards_GamePool_Empty() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        // GamePool doesn't support removed tokens, should return empty
        (address[] memory tokens, uint256[] memory rewards) = router.getRemovedTokenRewards(poolId, user1);

        assertEq(tokens.length, 0);
        assertEq(rewards.length, 0);
    }

    function test_Router_IsNativePool_GamePool() public view {
        // GamePool is not a native pool (doesn't use WCROSS)
        assertFalse(router.isNativePool(poolId));
    }

    function test_Router_GetTotalDeposited_GamePool() public {
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
        MockERC20GP v1DepositToken = new MockERC20GP("V1 Deposit", "V1D");
        MockERC20GP v1RewardToken = new MockERC20GP("V1 Reward", "V1R");

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

        // Deposit in GamePool via router
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);

        vm.stopPrank();

        // Verify independent balances
        assertEq(v1Pool.balances(user1), 100 ether);
        assertEq(gamePool.balances(user1), 100 ether);
    }

    // ==================== Edge Case Tests ====================

    function test_GamePool_RouterWithdrawAll() public {
        // Create V1 pool
        MockERC20GP v1DepositToken = new MockERC20GP("V1 Deposit", "V1D");
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
        assertEq(gamePool.balances(user1), 0);
    }

    function test_GamePool_ClaimSpecificToken() public {
        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        uint256 roundId = _createRound(10000 ether, 10, 100);
        IGamePool.Round memory round = gamePool.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(100);

        uint256 balanceBefore = crossdToken.balanceOf(user1);

        // Claim specific token via router
        vm.prank(user1);
        router.claimReward(poolId, address(crossdToken));

        assertEq(crossdToken.balanceOf(user1), balanceBefore + 10000 ether);
    }

    // ==================== Pool Status Tests via Factory ====================

    function test_Factory_SetPoolStatus_GamePool() public {
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

        assertEq(gamePool.balances(user1), 0);
    }

    // ==================== Reclaim via Factory ====================

    function test_Factory_ReclaimFromGamePool() public {
        // Create round with no depositors
        uint256 roundId = _createRound(10000 ether, 10, 100);
        IGamePool.Round memory round = gamePool.getRound(roundId);

        _advanceToBlock(round.startBlock);
        _advanceBlocks(100);

        // Trigger pool update
        _userDeposit(user1, 100 ether);

        // Reclaim via factory
        address recipient = address(0x8888);
        crossGameReward.reclaimFromPool(poolId, crossdToken, recipient);

        assertEq(crossdToken.balanceOf(recipient), 10000 ether);
    }

    // ==================== WithdrawAll Large Pool Scenario Tests ====================

    function test_WithdrawAll_ManyMixedPools() public {
        uint256 v1Count = 5;
        uint256 gpCount = 5;

        MockERC20GP[] memory v1Tokens = new MockERC20GP[](v1Count);
        MockERC20GP[] memory v1RewardTokens = new MockERC20GP[](v1Count);
        uint256[] memory v1PoolIds = new uint256[](v1Count);

        for (uint256 i = 0; i < v1Count; i++) {
            v1Tokens[i] = new MockERC20GP(string.concat("V1Dep", vm.toString(i)), "V1D");
            v1RewardTokens[i] = new MockERC20GP(string.concat("V1Rwd", vm.toString(i)), "V1R");
            v1Tokens[i].transfer(user1, 1000 ether);

            (v1PoolIds[i],) = crossGameReward.createPool(
                string.concat("V1 Pool ", vm.toString(i)),
                IERC20(address(v1Tokens[i])),
                MIN_DEPOSIT
            );
            crossGameReward.addRewardToken(v1PoolIds[i], IERC20(address(v1RewardTokens[i])));

            v1RewardTokens[i].transfer(address(crossGameReward.getPoolAddress(v1PoolIds[i])), 100 ether);
        }

        MockERC20GP[] memory gpGameTokens = new MockERC20GP[](gpCount);
        MockERC20GP[] memory gpRewardTokens = new MockERC20GP[](gpCount);
        uint256[] memory gpIds = new uint256[](gpCount);
        GamePool[] memory gamePools = new GamePool[](gpCount);

        for (uint256 i = 0; i < gpCount; i++) {
            gpGameTokens[i] = new MockERC20GP(string.concat("GPGame", vm.toString(i)), "GPG");
            gpRewardTokens[i] = new MockERC20GP(string.concat("GPRwd", vm.toString(i)), "GPR");
            gpGameTokens[i].transfer(user1, 1000 ether);
            gpRewardTokens[i].transfer(sponsor, 100000 ether);

            (gpIds[i],) = crossGameReward.createGamePool(
                string.concat("Game Pool ", vm.toString(i)),
                IERC20(address(gpGameTokens[i])),
                IERC20(address(gpRewardTokens[i])),
                MIN_DEPOSIT
            );
            gamePools[i] = GamePool(address(crossGameReward.getPoolAddress(gpIds[i])));
            crossGameReward.grantSponsorRole(gpIds[i], sponsor);

            vm.startPrank(sponsor);
            gpRewardTokens[i].approve(address(gamePools[i]), 10000 ether);
            gamePools[i].createRound(10000 ether, block.number + 10, 100);
            vm.stopPrank();
        }

        vm.startPrank(user1);
        for (uint256 i = 0; i < v1Count; i++) {
            v1Tokens[i].approve(address(router), 100 ether);
            router.depositERC20(v1PoolIds[i], 100 ether);
        }
        for (uint256 i = 0; i < gpCount; i++) {
            gpGameTokens[i].approve(address(router), 100 ether);
            router.depositERC20(gpIds[i], 100 ether);
        }
        vm.stopPrank();

        _advanceBlocks(200);

        for (uint256 i = 0; i < gpCount; i++) {
            vm.prank(sponsor);
            gamePools[i].syncRounds(0);
        }

        vm.prank(user1);
        router.withdrawAll();

        for (uint256 i = 0; i < v1Count; i++) {
            assertEq(crossGameReward.getPoolAddress(v1PoolIds[i]).balances(user1), 0);
        }
        for (uint256 i = 0; i < gpCount; i++) {
            assertEq(gamePools[i].balances(user1), 0);
        }
    }

    function test_WithdrawAll_ManyPools_NoPreSync() public {
        uint256 poolCount = 8;

        MockERC20GP[] memory gameTokens = new MockERC20GP[](poolCount);
        MockERC20GP[] memory rwdTokens = new MockERC20GP[](poolCount);
        uint256[] memory pIds = new uint256[](poolCount);
        GamePool[] memory pools = new GamePool[](poolCount);

        for (uint256 i = 0; i < poolCount; i++) {
            gameTokens[i] = new MockERC20GP(string.concat("G", vm.toString(i)), "G");
            rwdTokens[i] = new MockERC20GP(string.concat("R", vm.toString(i)), "R");
            gameTokens[i].transfer(user1, 1000 ether);
            rwdTokens[i].transfer(sponsor, 500000 ether);

            (pIds[i],) = crossGameReward.createGamePool(
                string.concat("Pool ", vm.toString(i)),
                IERC20(address(gameTokens[i])),
                IERC20(address(rwdTokens[i])),
                MIN_DEPOSIT
            );
            pools[i] = GamePool(address(crossGameReward.getPoolAddress(pIds[i])));
            crossGameReward.grantSponsorRole(pIds[i], sponsor);

            vm.startPrank(sponsor);
            rwdTokens[i].approve(address(pools[i]), 100000 ether);
            for (uint256 j = 0; j < 5; j++) {
                pools[i].createRound(10000 ether, block.number + 10 + j, 100);
            }
            vm.stopPrank();
        }

        vm.startPrank(user1);
        for (uint256 i = 0; i < poolCount; i++) {
            gameTokens[i].approve(address(router), 100 ether);
            router.depositERC20(pIds[i], 100 ether);
        }
        vm.stopPrank();

        _advanceBlocks(200);

        vm.prank(user1);
        router.withdrawAll();

        for (uint256 i = 0; i < poolCount; i++) {
            assertEq(pools[i].balances(user1), 0);
            assertEq(pools[i].getActiveRoundCount(), 0);
        }
    }

    function test_WithdrawAll_OnlyDeposited_PoolsSkipped() public {
        MockERC20GP extraGameToken = new MockERC20GP("Extra", "EXT");
        MockERC20GP extraRewardToken = new MockERC20GP("ExtraR", "EXTR");

        (uint256 extraPoolId,) = crossGameReward.createGamePool(
            "Extra Pool",
            IERC20(address(extraGameToken)),
            IERC20(address(extraRewardToken)),
            MIN_DEPOSIT
        );

        vm.startPrank(user1);
        gameToken.approve(address(router), 100 ether);
        router.depositERC20(poolId, 100 ether);
        vm.stopPrank();

        vm.prank(user1);
        router.withdrawAll();

        assertEq(gamePool.balances(user1), 0);
        assertEq(crossGameReward.getPoolAddress(extraPoolId).balances(user1), 0);
    }
}
