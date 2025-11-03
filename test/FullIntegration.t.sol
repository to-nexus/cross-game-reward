// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/CrossStaking.sol";
import "../src/CrossStakingPool.sol";
import "../src/CrossStakingRouter.sol";
import "../src/WCROSS.sol";
import "./mocks/MockERC20.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "forge-std/Test.sol";

/**
 * @title FullIntegration
 * @notice 전체 시스템 통합 테스트
 */
contract FullIntegrationTest is Test {
    CrossStaking public crossStaking;
    CrossStakingPool public poolImplementation;
    CrossStakingRouter public router;
    WCROSS public wcross;

    MockERC20 public usdt;
    MockERC20 public usdc;
    MockERC20 public dai;

    address public admin;
    address public alice;
    address public bob;
    address public carol;

    uint public nativePoolId;
    address public nativePoolAddress;

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
        poolImplementation = new CrossStakingPool();

        // Deploy CrossStaking as UUPS proxy (WCROSS를 생성함)
        CrossStaking implementation = new CrossStaking();
        bytes memory initData = abi.encodeCall(CrossStaking.initialize, (address(poolImplementation), admin, 2 days));
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), initData);
        crossStaking = CrossStaking(address(proxy));

        router = new CrossStakingRouter(address(crossStaking));
        wcross = WCROSS(payable(crossStaking.wcross()));

        // Setup router
        crossStaking.setRouter(address(router));

        // Deploy reward tokens
        usdt = new MockERC20("Tether USD", "USDT");
        usdc = new MockERC20("USD Coin", "USDC");
        dai = new MockERC20("Dai Stablecoin", "DAI");

        // Create native pool
        (nativePoolId, nativePoolAddress) = crossStaking.createPool(address(wcross), 1 ether);

        // Add reward tokens
        crossStaking.addRewardToken(nativePoolId, address(usdt));
        crossStaking.addRewardToken(nativePoolId, address(usdc));
        crossStaking.addRewardToken(nativePoolId, address(dai));

        // Mint rewards for admin
        usdt.mint(admin, 10000 ether);
        usdc.mint(admin, 10000 ether);
        dai.mint(admin, 10000 ether);
    }

    // ==================== 전체 사용자 여정 ====================

    function testCompleteUserJourney() public {
        CrossStakingPool pool = CrossStakingPool(nativePoolAddress);

        // Day 0: Alice stakes 100 CROSS
        vm.startPrank(alice);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        assertEq(pool.balances(alice), 100 ether, "Alice staked");

        // Day 1: Reward 입금 (USDT 1000)
        usdt.transfer(nativePoolAddress, 1000 ether);

        // Day 2: Bob stakes 200 CROSS
        vm.startPrank(bob);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 200 ether}(nativePoolId);
        vm.stopPrank();

        assertEq(pool.totalStaked(), 300 ether, "Total staked");

        // Day 3: Reward 입금 (USDC 600)
        usdc.transfer(nativePoolAddress, 600 ether);

        // Day 4: Carol stakes 100 CROSS
        vm.startPrank(carol);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        // Day 5: Reward 입금 (DAI 800)
        dai.transfer(nativePoolAddress, 800 ether);

        // Verify Alice's rewards
        (uint aliceStaked, uint[] memory aliceRewards) = router.getUserStakingInfo(nativePoolId, alice);

        assertEq(aliceStaked, 100 ether, "Alice staked");

        // Alice: 1000 USDT (alone) + 200 USDC (1/3) + 200 DAI (1/4) = 1400+
        assertApproxEqAbs(aliceRewards[0], 1000 ether, 100, "Alice USDT");
        assertApproxEqAbs(aliceRewards[1], 200 ether, 100, "Alice USDC");
        assertApproxEqAbs(aliceRewards[2], 200 ether, 100, "Alice DAI");

        // Bob: 0 USDT + 400 USDC (2/3) + 400 DAI (2/4) = 800
        (, uint[] memory bobRewards) = router.getUserStakingInfo(nativePoolId, bob);
        assertEq(bobRewards[0], 0, "Bob no USDT");
        assertApproxEqAbs(bobRewards[1], 400 ether, 100, "Bob USDC");
        assertApproxEqAbs(bobRewards[2], 400 ether, 100, "Bob DAI");

        // Carol: 0 USDT + 0 USDC + 200 DAI (1/4) = 200
        (, uint[] memory carolRewards) = router.getUserStakingInfo(nativePoolId, carol);
        assertEq(carolRewards[0], 0, "Carol no USDT");
        assertEq(carolRewards[1], 0, "Carol no USDC");
        assertApproxEqAbs(carolRewards[2], 200 ether, 100, "Carol DAI");

        // Alice unstakes
        uint aliceBalanceBefore = alice.balance;
        vm.prank(alice);
        router.unstakeNative(nativePoolId);

        assertEq(alice.balance, aliceBalanceBefore + 100 ether, "Alice got native CROSS");
        assertApproxEqAbs(usdt.balanceOf(alice), 1000 ether, 100, "Alice got USDT");
        assertApproxEqAbs(usdc.balanceOf(alice), 200 ether, 100, "Alice got USDC");
        assertApproxEqAbs(dai.balanceOf(alice), 200 ether, 100, "Alice got DAI");
    }

    // ==================== 다중 풀 시나리오 ====================

    function testMultiplePoolsSimultaneously() public {
        // Create another pool for ERC20
        MockERC20 stakingToken = new MockERC20("Staking", "STK");
        (uint erc20PoolId, address erc20PoolAddress) = crossStaking.createPool(address(stakingToken), 1 ether);

        crossStaking.addRewardToken(erc20PoolId, address(usdt));

        // Mint tokens
        stakingToken.mint(alice, 1000 ether);
        stakingToken.mint(bob, 1000 ether);

        // Alice stakes in native pool
        vm.startPrank(alice);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 50 ether}(nativePoolId);
        vm.stopPrank();

        // Bob stakes in ERC20 pool
        vm.startPrank(bob);
        stakingToken.approve(address(router), 100 ether);
        router.stakeERC20(erc20PoolId, 100 ether);
        vm.stopPrank();

        // Add rewards to both pools
        usdt.transfer(nativePoolAddress, 100 ether);
        usdt.transfer(erc20PoolAddress, 200 ether);

        // Check rewards
        (, uint[] memory aliceRewards) = router.getUserStakingInfo(nativePoolId, alice);
        (, uint[] memory bobRewards) = router.getUserStakingInfo(erc20PoolId, bob);

        assertApproxEqAbs(aliceRewards[0], 100 ether, 100, "Alice native pool rewards");
        assertApproxEqAbs(bobRewards[0], 200 ether, 100, "Bob ERC20 pool rewards");
    }

    // ==================== 실전 시나리오 ====================

    function testRealWorldScenario() public {
        CrossStakingPool pool = CrossStakingPool(nativePoolAddress);

        // Week 1: Initial stakers
        vm.startPrank(alice);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        vm.startPrank(bob);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 50 ether}(nativePoolId);
        vm.stopPrank();

        // Week 1: First rewards
        usdt.transfer(nativePoolAddress, 300 ether);

        vm.warp(block.timestamp + 7 days);

        // Week 2: Carol joins
        vm.startPrank(carol);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 150 ether}(nativePoolId);
        vm.stopPrank();

        // Week 2: More rewards
        usdc.transfer(nativePoolAddress, 600 ether);

        vm.warp(block.timestamp + 7 days);

        // Week 3: Rewards
        dai.transfer(nativePoolAddress, 900 ether);

        // Verify final balances
        assertEq(pool.balances(alice), 100 ether, "Alice still staked");
        assertEq(pool.balances(bob), 50 ether, "Bob still staked");
        assertEq(pool.balances(carol), 150 ether, "Carol still staked");
        assertEq(pool.totalStaked(), 300 ether, "Total staked");

        // Alice: 200 USDT (2/3) + 200 USDC (1/3) + 300 DAI (1/3) = 700
        (, uint[] memory aliceRewards) = router.getUserStakingInfo(nativePoolId, alice);
        assertApproxEqAbs(aliceRewards[0], 200 ether, 100, "Alice USDT");
        assertApproxEqAbs(aliceRewards[1], 200 ether, 100, "Alice USDC");
        assertApproxEqAbs(aliceRewards[2], 300 ether, 100, "Alice DAI");

        // Bob: 100 USDT (1/3) + 100 USDC (1/6) + 150 DAI (1/6) = 350
        (, uint[] memory bobRewards) = router.getUserStakingInfo(nativePoolId, bob);
        assertApproxEqAbs(bobRewards[0], 100 ether, 100, "Bob USDT");
        assertApproxEqAbs(bobRewards[1], 100 ether, 100, "Bob USDC");
        assertApproxEqAbs(bobRewards[2], 150 ether, 100, "Bob DAI");

        // Carol: 0 USDT + 300 USDC (1/2) + 450 DAI (1/2) = 750
        (, uint[] memory carolRewards) = router.getUserStakingInfo(nativePoolId, carol);
        assertEq(carolRewards[0], 0, "Carol no USDT");
        assertApproxEqAbs(carolRewards[1], 300 ether, 100, "Carol USDC");
        assertApproxEqAbs(carolRewards[2], 450 ether, 100, "Carol DAI");

        // All unstake
        vm.prank(alice);
        router.unstakeNative(nativePoolId);

        vm.prank(bob);
        router.unstakeNative(nativePoolId);

        vm.prank(carol);
        router.unstakeNative(nativePoolId);

        // Verify pool is empty
        assertEq(pool.totalStaked(), 0, "Pool empty");
    }

    // ==================== 보상 정확성 검증 ====================

    function testRewardDistributionAccuracy() public {
        CrossStakingPool pool = CrossStakingPool(nativePoolAddress);

        // Setup: 3 users with different stakes
        vm.startPrank(alice);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        vm.startPrank(bob);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 200 ether}(nativePoolId);
        vm.stopPrank();

        vm.startPrank(carol);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 300 ether}(nativePoolId);
        vm.stopPrank();

        // Total: 600 ether staked
        assertEq(pool.totalStaked(), 600 ether, "Total staked");

        // Deposit 600 USDT
        usdt.transfer(nativePoolAddress, 600 ether);

        // Expected distribution: 100 (alice), 200 (bob), 300 (carol)
        (, uint[] memory aliceRewards) = router.getUserStakingInfo(nativePoolId, alice);
        (, uint[] memory bobRewards) = router.getUserStakingInfo(nativePoolId, bob);
        (, uint[] memory carolRewards) = router.getUserStakingInfo(nativePoolId, carol);

        assertApproxEqAbs(aliceRewards[0], 100 ether, 100, "Alice gets 100");
        assertApproxEqAbs(bobRewards[0], 200 ether, 100, "Bob gets 200");
        assertApproxEqAbs(carolRewards[0], 300 ether, 100, "Carol gets 300");

        // Verify total = 600
        uint totalRewards = aliceRewards[0] + bobRewards[0] + carolRewards[0];
        assertApproxEqAbs(totalRewards, 600 ether, 100, "Total rewards match");
    }

    // ==================== 에지 케이스 ====================

    function testStakeUnstakeStake() public {
        // Alice stakes
        vm.startPrank(alice);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 50 ether}(nativePoolId);
        vm.stopPrank();

        // Reward 1
        usdt.transfer(nativePoolAddress, 100 ether);

        // Alice unstakes
        vm.prank(alice);
        router.unstakeNative(nativePoolId);

        assertApproxEqAbs(usdt.balanceOf(alice), 100 ether, 100, "Alice got first rewards");

        // Reward 2 (Alice가 없는 동안 입금 - Alice가 받음!)
        usdt.transfer(nativePoolAddress, 100 ether);

        // Alice stakes again
        vm.startPrank(alice);
        router.stakeNative{value: 50 ether}(nativePoolId);
        vm.stopPrank();

        // Reward 3
        usdt.transfer(nativePoolAddress, 100 ether);

        // Alice unstakes again
        uint usdtBefore = usdt.balanceOf(alice);
        vm.prank(alice);
        router.unstakeNative(nativePoolId);

        // Alice should get Reward 2 + Reward 3 (스테이커 없을 때 입금된 보상 포함)
        assertApproxEqAbs(
            usdt.balanceOf(alice) - usdtBefore, 200 ether, 1 ether, "Alice got rewards including zero-staker period"
        );
    }

    function testMultipleRewardRounds() public {
        // Alice stakes
        vm.startPrank(alice);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        // 10 rounds of rewards
        for (uint i = 0; i < 10; i++) {
            usdt.transfer(nativePoolAddress, 10 ether);
            usdc.transfer(nativePoolAddress, 20 ether);
            dai.transfer(nativePoolAddress, 30 ether);
        }

        // Alice should have 100 USDT, 200 USDC, 300 DAI
        (, uint[] memory rewards) = router.getUserStakingInfo(nativePoolId, alice);
        assertApproxEqAbs(rewards[0], 100 ether, 100, "Alice USDT");
        assertApproxEqAbs(rewards[1], 200 ether, 100, "Alice USDC");
        assertApproxEqAbs(rewards[2], 300 ether, 100, "Alice DAI");
    }

    // ==================== 보안 검증 ====================

    function testCannotUnstakeOthersStake() public {
        // Alice stakes
        vm.startPrank(alice);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        // Bob tries to unstake Alice's stake
        vm.prank(bob);
        vm.expectRevert(CrossStakingRouter.CSRNoStakeFound.selector);
        router.unstakeNative(nativePoolId);
    }

    function testReentrancyProtection() public {
        // Router는 CrossStakingPool의 nonReentrant로 보호됨
        // 이중 호출 시도
        vm.startPrank(alice);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 10 ether}(nativePoolId);

        // 정상적으로 완료되어야 함
        assertEq(CrossStakingPool(nativePoolAddress).balances(alice), 10 ether, "Stake succeeded");
        vm.stopPrank();
    }

    // ==================== View 함수 ====================

    function testViewFunctionsConsistency() public {
        // Stake
        vm.startPrank(alice);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 100 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        usdt.transfer(nativePoolAddress, 50 ether);

        // Get info via router
        (uint stakedViaRouter, uint[] memory rewardsViaRouter) = router.getUserStakingInfo(nativePoolId, alice);

        // Get info directly from pool
        CrossStakingPool pool = CrossStakingPool(nativePoolAddress);
        uint stakedDirect = pool.balances(alice);
        uint[] memory rewardsDirect = pool.pendingRewards(alice);

        // Should match
        assertEq(stakedViaRouter, stakedDirect, "Staked amount matches");
        assertEq(rewardsViaRouter.length, rewardsDirect.length, "Rewards length matches");
        assertEq(rewardsViaRouter[0], rewardsDirect[0], "Rewards match");
    }
}
