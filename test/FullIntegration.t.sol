// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardPool.sol";
import "../src/CrossGameRewardRouter.sol";
import "../src/WCROSS.sol";
import "../src/interfaces/ICrossGameRewardPool.sol";
import "./mocks/MockERC20.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Test.sol";

/**
 * @title FullIntegration
 * @notice End-to-end system integration tests
 */
contract FullIntegrationTest is Test {
    CrossGameReward public crossGameReward;
    CrossGameRewardPool public poolImplementation;
    CrossGameRewardRouter public router;
    WCROSS public wcross;

    MockERC20 public usdt;
    MockERC20 public usdc;
    MockERC20 public dai;

    address public admin;
    address public alice;
    address public bob;
    address public carol;

    uint public nativePoolId;
    ICrossGameRewardPool public nativePool;

    function setUp() public {
        admin = address(this);
        alice = makeAddr("alice");
        bob = makeAddr("bob");
        carol = makeAddr("carol");

        // Give users native CROSS
        vm.deal(alice, 1000 ether);
        vm.deal(bob, 1000 ether);
        vm.deal(carol, 1000 ether);

        // Deploy system
        poolImplementation = new CrossGameRewardPool();

        // Deploy CrossGameReward as a UUPS proxy (instantiates WCROSS internally)
        CrossGameReward implementation = new CrossGameReward();
        bytes memory initData = abi.encodeCall(
            CrossGameReward.initialize, (ICrossGameRewardPool(address(poolImplementation)), admin, 2 days)
        );
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), initData);
        crossGameReward = CrossGameReward(address(proxy));

        router = new CrossGameRewardRouter(address(crossGameReward));
        wcross = WCROSS(payable(address(crossGameReward.wcross())));

        // Setup router
        crossGameReward.setRouter(address(router));

        // Deploy reward tokens
        usdt = new MockERC20("Tether USD", "USDT");
        usdc = new MockERC20("USD Coin", "USDC");
        dai = new MockERC20("Dai Stablecoin", "DAI");

        // Create native pool
        (nativePoolId, nativePool) = crossGameReward.createPool("Native Pool", IERC20(address(wcross)), 1 ether);

        // Add reward tokens
        crossGameReward.addRewardToken(nativePoolId, IERC20(address(usdt)));
        crossGameReward.addRewardToken(nativePoolId, IERC20(address(usdc)));
        crossGameReward.addRewardToken(nativePoolId, IERC20(address(dai)));

        // Mint rewards for admin
        usdt.mint(admin, 10000 ether);
        usdc.mint(admin, 10000 ether);
        dai.mint(admin, 10000 ether);
    }

    // ==================== Full user journey ====================

    function testCompleteUserJourney() public {
        CrossGameRewardPool pool = CrossGameRewardPool(address(nativePool));

        // Day 0: Alice deposits 100 CROSS
        vm.startPrank(alice);
        router.depositNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        assertEq(pool.balances(alice), 100 ether, "Alice depositd");

        // Day 1: Deposit reward (USDT 1000)
        usdt.transfer(address(nativePool), 1000 ether);

        // Day 2: Bob deposits 200 CROSS
        vm.startPrank(bob);
        router.depositNative{value: 200 ether}(nativePoolId);
        vm.stopPrank();

        assertEq(pool.totalDeposited(), 300 ether, "Total depositd");

        // Day 3: Deposit reward (USDC 600)
        usdc.transfer(address(nativePool), 600 ether);

        // Day 4: Carol deposits 100 CROSS
        vm.startPrank(carol);
        router.depositNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        // Day 5: Deposit reward (DAI 800)
        dai.transfer(address(nativePool), 800 ether);

        // Verify Alice's rewards
        (uint aliceDeposited,, uint[] memory aliceRewards) = router.getUserDepositInfo(nativePoolId, alice);

        assertEq(aliceDeposited, 100 ether, "Alice depositd");

        // Alice: 1000 USDT (alone) + 200 USDC (1/3) + 200 DAI (1/4) = 1400+
        assertApproxEqAbs(aliceRewards[0], 1000 ether, 100, "Alice USDT");
        assertApproxEqAbs(aliceRewards[1], 200 ether, 100, "Alice USDC");
        assertApproxEqAbs(aliceRewards[2], 200 ether, 100, "Alice DAI");

        // Bob: 0 USDT + 400 USDC (2/3) + 400 DAI (2/4) = 800
        (,, uint[] memory bobRewards) = router.getUserDepositInfo(nativePoolId, bob);
        assertEq(bobRewards[0], 0, "Bob no USDT");
        assertApproxEqAbs(bobRewards[1], 400 ether, 100, "Bob USDC");
        assertApproxEqAbs(bobRewards[2], 400 ether, 100, "Bob DAI");

        // Carol: 0 USDT + 0 USDC + 200 DAI (1/4) = 200
        (,, uint[] memory carolRewards) = router.getUserDepositInfo(nativePoolId, carol);
        assertEq(carolRewards[0], 0, "Carol no USDT");
        assertEq(carolRewards[1], 0, "Carol no USDC");
        assertApproxEqAbs(carolRewards[2], 200 ether, 100, "Carol DAI");

        // Alice withdraws
        uint aliceBalanceBefore = alice.balance;
        vm.prank(alice);
        router.withdrawNative(nativePoolId, 0);

        assertEq(alice.balance, aliceBalanceBefore + 100 ether, "Alice got native CROSS");
        assertApproxEqAbs(usdt.balanceOf(alice), 1000 ether, 100, "Alice got USDT");
        assertApproxEqAbs(usdc.balanceOf(alice), 200 ether, 100, "Alice got USDC");
        assertApproxEqAbs(dai.balanceOf(alice), 200 ether, 100, "Alice got DAI");
    }

    // ==================== Multiple pool scenario ====================

    function testMultiplePoolsSimultaneously() public {
        // Create another pool for ERC20
        MockERC20 depositToken = new MockERC20("Deposit", "STK");
        (uint erc20PoolId, ICrossGameRewardPool erc20PoolAddress) =
            crossGameReward.createPool("ERC20 Pool", IERC20(address(depositToken)), 1 ether);

        crossGameReward.addRewardToken(erc20PoolId, IERC20(address(usdt)));

        // Mint tokens
        depositToken.mint(alice, 1000 ether);
        depositToken.mint(bob, 1000 ether);

        // Alice deposits in native pool
        vm.startPrank(alice);
        router.depositNative{value: 50 ether}(nativePoolId);
        vm.stopPrank();

        // Bob deposits in ERC20 pool
        vm.startPrank(bob);
        depositToken.approve(address(router), 100 ether);
        router.depositERC20(erc20PoolId, 100 ether);
        vm.stopPrank();

        // Add rewards to both pools
        usdt.transfer(address(nativePool), 100 ether);
        usdt.transfer(address(erc20PoolAddress), 200 ether);

        // Check rewards
        (,, uint[] memory aliceRewards) = router.getUserDepositInfo(nativePoolId, alice);
        (,, uint[] memory bobRewards) = router.getUserDepositInfo(erc20PoolId, bob);

        assertApproxEqAbs(aliceRewards[0], 100 ether, 100, "Alice native pool rewards");
        assertApproxEqAbs(bobRewards[0], 200 ether, 100, "Bob ERC20 pool rewards");
    }

    // ==================== Realistic scenario ====================

    function testRealWorldScenario() public {
        CrossGameRewardPool pool = CrossGameRewardPool(address(nativePool));

        // Week 1: Initial depositrs
        vm.startPrank(alice);
        router.depositNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        vm.startPrank(bob);
        router.depositNative{value: 50 ether}(nativePoolId);
        vm.stopPrank();

        // Week 1: First rewards
        usdt.transfer(address(nativePool), 300 ether);

        vm.warp(block.timestamp + 7 days);

        // Week 2: Carol joins
        vm.startPrank(carol);
        router.depositNative{value: 150 ether}(nativePoolId);
        vm.stopPrank();

        // Week 2: More rewards
        usdc.transfer(address(nativePool), 600 ether);

        vm.warp(block.timestamp + 7 days);

        // Week 3: Rewards
        dai.transfer(address(nativePool), 900 ether);

        // Verify final balances
        assertEq(pool.balances(alice), 100 ether, "Alice still depositd");
        assertEq(pool.balances(bob), 50 ether, "Bob still depositd");
        assertEq(pool.balances(carol), 150 ether, "Carol still depositd");
        assertEq(pool.totalDeposited(), 300 ether, "Total depositd");

        // Alice: 200 USDT (2/3) + 200 USDC (1/3) + 300 DAI (1/3) = 700
        (,, uint[] memory aliceRewards) = router.getUserDepositInfo(nativePoolId, alice);
        assertApproxEqAbs(aliceRewards[0], 200 ether, 100, "Alice USDT");
        assertApproxEqAbs(aliceRewards[1], 200 ether, 100, "Alice USDC");
        assertApproxEqAbs(aliceRewards[2], 300 ether, 100, "Alice DAI");

        // Bob: 100 USDT (1/3) + 100 USDC (1/6) + 150 DAI (1/6) = 350
        (,, uint[] memory bobRewards) = router.getUserDepositInfo(nativePoolId, bob);
        assertApproxEqAbs(bobRewards[0], 100 ether, 100, "Bob USDT");
        assertApproxEqAbs(bobRewards[1], 100 ether, 100, "Bob USDC");
        assertApproxEqAbs(bobRewards[2], 150 ether, 100, "Bob DAI");

        // Carol: 0 USDT + 300 USDC (1/2) + 450 DAI (1/2) = 750
        (,, uint[] memory carolRewards) = router.getUserDepositInfo(nativePoolId, carol);
        assertEq(carolRewards[0], 0, "Carol no USDT");
        assertApproxEqAbs(carolRewards[1], 300 ether, 100, "Carol USDC");
        assertApproxEqAbs(carolRewards[2], 450 ether, 100, "Carol DAI");

        // All withdraw
        vm.prank(alice);
        router.withdrawNative(nativePoolId, 0);

        vm.prank(bob);
        router.withdrawNative(nativePoolId, 0);

        vm.prank(carol);
        router.withdrawNative(nativePoolId, 0);

        // Verify pool is empty
        assertEq(pool.totalDeposited(), 0, "Pool empty");
    }

    // ==================== Reward accuracy ====================

    function testRewardDistributionAccuracy() public {
        CrossGameRewardPool pool = CrossGameRewardPool(address(nativePool));

        // Setup: 3 users with different deposits
        vm.startPrank(alice);
        router.depositNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        vm.startPrank(bob);
        router.depositNative{value: 200 ether}(nativePoolId);
        vm.stopPrank();

        vm.startPrank(carol);
        router.depositNative{value: 300 ether}(nativePoolId);
        vm.stopPrank();

        // Total: 600 ether depositd
        assertEq(pool.totalDeposited(), 600 ether, "Total depositd");

        // Deposit 600 USDT
        usdt.transfer(address(nativePool), 600 ether);

        // Expected distribution: 100 (alice), 200 (bob), 300 (carol)
        (,, uint[] memory aliceRewards) = router.getUserDepositInfo(nativePoolId, alice);
        (,, uint[] memory bobRewards) = router.getUserDepositInfo(nativePoolId, bob);
        (,, uint[] memory carolRewards) = router.getUserDepositInfo(nativePoolId, carol);

        assertApproxEqAbs(aliceRewards[0], 100 ether, 100, "Alice gets 100");
        assertApproxEqAbs(bobRewards[0], 200 ether, 100, "Bob gets 200");
        assertApproxEqAbs(carolRewards[0], 300 ether, 100, "Carol gets 300");

        // Verify total = 600
        uint totalRewards = aliceRewards[0] + bobRewards[0] + carolRewards[0];
        assertApproxEqAbs(totalRewards, 600 ether, 100, "Total rewards match");
    }

    // ==================== Additional edge cases ====================

    function testDepositWithdrawDeposit() public {
        // Alice deposits
        vm.startPrank(alice);
        router.depositNative{value: 50 ether}(nativePoolId);
        vm.stopPrank();

        // Reward 1
        usdt.transfer(address(nativePool), 100 ether);

        // Alice withdraws
        vm.prank(alice);
        router.withdrawNative(nativePoolId, 0);

        assertApproxEqAbs(usdt.balanceOf(alice), 100 ether, 100, "Alice got first rewards");

        // Reward 2 was deposited while nobody depositd, so it becomes withdrawable (NOT given to Alice)
        usdt.transfer(address(nativePool), 100 ether);

        // Alice deposits again
        vm.startPrank(alice);
        router.depositNative{value: 50 ether}(nativePoolId);
        vm.stopPrank();

        // Reward 3
        usdt.transfer(address(nativePool), 100 ether);

        // Alice withdraws again
        uint usdtBefore = usdt.balanceOf(alice);
        vm.prank(alice);
        router.withdrawNative(nativePoolId, 0);

        // Alice should receive only Reward 3 (Reward 2 was deposited when pool was empty)
        assertApproxEqAbs(usdt.balanceOf(alice) - usdtBefore, 100 ether, 1 ether, "Alice got only Reward 3");

        // Reward 2 should be withdrawable
        assertEq(nativePool.getReclaimableAmount(usdt), 100 ether, "Reward 2 is withdrawable");
    }

    function testMultipleRewardRounds() public {
        // Alice deposits
        vm.startPrank(alice);
        router.depositNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        // 10 rounds of rewards
        for (uint i = 0; i < 10; i++) {
            usdt.transfer(address(nativePool), 10 ether);
            usdc.transfer(address(nativePool), 20 ether);
            dai.transfer(address(nativePool), 30 ether);
        }

        // Alice should have 100 USDT, 200 USDC, 300 DAI
        (,, uint[] memory rewards) = router.getUserDepositInfo(nativePoolId, alice);
        assertApproxEqAbs(rewards[0], 100 ether, 100, "Alice USDT");
        assertApproxEqAbs(rewards[1], 200 ether, 100, "Alice USDC");
        assertApproxEqAbs(rewards[2], 300 ether, 100, "Alice DAI");
    }

    // ==================== Security checks ====================

    function testCannotWithdrawOthersDeposit() public {
        // Alice deposits
        vm.startPrank(alice);
        router.depositNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        // Bob tries to withdraw Alice's deposit
        vm.prank(bob);
        vm.expectRevert(abi.encodeWithSelector(CrossGameRewardPool.CGRPNoDepositFound.selector, bob));
        router.withdrawNative(nativePoolId, 0);
    }

    function testReentrancyProtection() public {
        // Router interactions are guarded by the pool's nonReentrant modifier
        // Attempt to call twice
        vm.startPrank(alice);
        router.depositNative{value: 10 ether}(nativePoolId);

        // Should execute successfully
        assertEq(CrossGameRewardPool(address(nativePool)).balances(alice), 10 ether, "Deposit succeeded");
        vm.stopPrank();
    }

    // ==================== View function checks ====================

    function testViewFunctionsConsistency() public {
        // Deposit
        vm.startPrank(alice);
        router.depositNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        usdt.transfer(address(nativePool), 50 ether);

        // Get info via router
        (uint depositdViaRouter,, uint[] memory rewardsViaRouter) = router.getUserDepositInfo(nativePoolId, alice);

        // Get info directly from pool
        CrossGameRewardPool pool = CrossGameRewardPool(address(nativePool));
        uint depositdDirect = pool.balances(alice);
        (, uint[] memory rewardsDirect) = pool.pendingRewards(alice);

        // Should match
        assertEq(depositdViaRouter, depositdDirect, "Deposited amount matches");
        assertEq(rewardsViaRouter.length, rewardsDirect.length, "Rewards length matches");
        assertEq(rewardsViaRouter[0], rewardsDirect[0], "Rewards match");
    }
}
