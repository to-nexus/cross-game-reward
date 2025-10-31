// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/CrossStaking.sol";
import "../src/CrossStakingPool.sol";
import "../src/CrossStakingRouter.sol";
import "../src/WCROSS.sol";
import "./mocks/MockERC20.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "forge-std/Test.sol";

contract CrossStakingRouterTest is Test {
    CrossStaking public crossStaking;
    CrossStakingPool public poolImplementation;
    CrossStakingRouter public router;
    WCROSS public wcross;

    MockERC20 public rewardToken;
    MockERC20 public stakingToken;

    address public owner;
    address public user1;
    address public user2;

    uint public nativePoolId;
    address public nativePoolAddress;

    uint public erc20PoolId;
    address public erc20PoolAddress;

    function setUp() public {
        owner = address(this);
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        // Give users native CROSS
        vm.deal(user1, 1000 ether);
        vm.deal(user2, 1000 ether);

        // Deploy core contracts
        poolImplementation = new CrossStakingPool();

        // Deploy CrossStaking as UUPS proxy (WCROSS를 생성함)
        CrossStaking implementation = new CrossStaking();
        bytes memory initData =
            abi.encodeWithSelector(CrossStaking.initialize.selector, address(poolImplementation), owner, 2 days);
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), initData);
        crossStaking = CrossStaking(address(proxy));

        router = new CrossStakingRouter(address(crossStaking));
        wcross = WCROSS(payable(crossStaking.wcross()));

        // Setup router
        crossStaking.setRouter(address(router));

        // Create test tokens
        rewardToken = new MockERC20("Reward", "RWD");
        stakingToken = new MockERC20("Staking", "STK");

        // Create pools
        (nativePoolId, nativePoolAddress) = crossStaking.createPool(address(wcross), 2 days);
        (erc20PoolId, erc20PoolAddress) = crossStaking.createPool(address(stakingToken), 2 days);

        // Add reward tokens
        crossStaking.addRewardToken(nativePoolId, address(rewardToken));
        crossStaking.addRewardToken(erc20PoolId, address(rewardToken));

        // Mint staking tokens for users
        stakingToken.mint(user1, 1000 ether);
        stakingToken.mint(user2, 1000 ether);
    }

    // ==================== Native CROSS 스테이킹 ====================

    function testStakeNative() public {
        uint amount = 10 ether;

        vm.startPrank(user1);

        // User needs to approve router for WCROSS
        wcross.approve(address(router), type(uint).max);

        router.stakeNative{value: amount}(nativePoolId);
        vm.stopPrank();

        // Verify
        CrossStakingPool pool = CrossStakingPool(nativePoolAddress);
        assertEq(pool.balances(user1), amount, "User staked");
        assertEq(pool.totalStaked(), amount, "Total staked");
    }

    function testStakeNativeMultipleTimes() public {
        vm.startPrank(user1);
        wcross.approve(address(router), type(uint).max);

        router.stakeNative{value: 5 ether}(nativePoolId);
        router.stakeNative{value: 3 ether}(nativePoolId);
        router.stakeNative{value: 2 ether}(nativePoolId);
        vm.stopPrank();

        CrossStakingPool pool = CrossStakingPool(nativePoolAddress);
        assertEq(pool.balances(user1), 10 ether, "Total user stake");
    }

    function testCannotStakeNativeZero() public {
        vm.prank(user1);
        vm.expectRevert(CrossStakingRouter.CSRInvalidAmount.selector);
        router.stakeNative{value: 0}(nativePoolId);
    }

    function testCannotStakeNativeOnERC20Pool() public {
        vm.prank(user1);
        vm.expectRevert(CrossStakingRouter.CSRNotWCROSSPool.selector);
        router.stakeNative{value: 10 ether}(erc20PoolId);
    }

    function testUnstakeNative() public {
        // Stake first
        vm.startPrank(user1);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(nativePoolAddress, 100 ether);

        // Unstake
        uint balanceBefore = user1.balance;
        vm.startPrank(user1);
        router.unstakeNative(nativePoolId);
        vm.stopPrank();

        // Verify
        CrossStakingPool pool = CrossStakingPool(nativePoolAddress);
        assertEq(pool.balances(user1), 0, "Unstaked");
        assertEq(user1.balance, balanceBefore + 10 ether, "Native CROSS returned");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 100 ether, 1 ether, "Rewards claimed");
    }

    function testCannotUnstakeNativeWithoutStake() public {
        vm.prank(user1);
        vm.expectRevert(CrossStakingRouter.CSRNoStakeFound.selector);
        router.unstakeNative(nativePoolId);
    }

    // ==================== ERC20 스테이킹 ====================

    function testStakeERC20() public {
        uint amount = 50 ether;

        vm.startPrank(user1);
        stakingToken.approve(address(router), amount);
        router.stakeERC20(erc20PoolId, amount);
        vm.stopPrank();

        CrossStakingPool pool = CrossStakingPool(erc20PoolAddress);
        assertEq(pool.balances(user1), amount, "User staked");
    }

    function testStakeERC20MultipleTimes() public {
        vm.startPrank(user1);
        stakingToken.approve(address(router), 100 ether);

        router.stakeERC20(erc20PoolId, 30 ether);
        router.stakeERC20(erc20PoolId, 20 ether);
        router.stakeERC20(erc20PoolId, 10 ether);
        vm.stopPrank();

        CrossStakingPool pool = CrossStakingPool(erc20PoolAddress);
        assertEq(pool.balances(user1), 60 ether, "Total stake");
    }

    function testUnstakeERC20() public {
        // Stake first
        vm.startPrank(user1);
        stakingToken.approve(address(router), 50 ether);
        router.stakeERC20(erc20PoolId, 50 ether);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(erc20PoolAddress, 100 ether);

        // Unstake
        uint balanceBefore = stakingToken.balanceOf(user1);
        vm.prank(user1);
        router.unstakeERC20(erc20PoolId);

        CrossStakingPool pool = CrossStakingPool(erc20PoolAddress);
        assertEq(pool.balances(user1), 0, "Unstaked");
        assertEq(stakingToken.balanceOf(user1), balanceBefore + 50 ether, "Tokens returned");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 100 ether, 1 ether, "Rewards claimed");
    }

    function testCannotUnstakeERC20WithoutStake() public {
        vm.prank(user1);
        vm.expectRevert(CrossStakingRouter.CSRNoStakeFound.selector);
        router.unstakeERC20(erc20PoolId);
    }

    function testCannotStakeERC20Zero() public {
        vm.prank(user1);
        vm.expectRevert(CrossStakingRouter.CSRInvalidAmount.selector);
        router.stakeERC20(erc20PoolId, 0);
    }

    // ==================== View 함수 ====================

    function testGetUserStakingInfo() public {
        // Stake
        vm.startPrank(user1);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 50 ether);
        rewardToken.transfer(nativePoolAddress, 50 ether);

        (uint stakedAmount, uint[] memory pendingRewards) = router.getUserStakingInfo(nativePoolId, user1);

        assertEq(stakedAmount, 10 ether, "Staked amount");
        assertEq(pendingRewards.length, 1, "1 reward token");
        assertApproxEqAbs(pendingRewards[0], 50 ether, 1 ether, "Pending reward");
    }

    function testIsNativePool() public view {
        assertTrue(router.isNativePool(nativePoolId), "Native pool");
        assertFalse(router.isNativePool(erc20PoolId), "Not native pool");
    }

    // ==================== 복잡한 시나리오 ====================

    function testMultiUserNativeStaking() public {
        // User1 stakes
        vm.startPrank(user1);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // User2 stakes
        vm.startPrank(user2);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 20 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 90 ether);
        rewardToken.transfer(nativePoolAddress, 90 ether);

        // User1 unstakes
        uint user1BalanceBefore = user1.balance;
        vm.prank(user1);
        router.unstakeNative(nativePoolId);

        assertEq(user1.balance, user1BalanceBefore + 10 ether, "User1 got native CROSS");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 30 ether, 1 ether, "User1 rewards (1/3)");

        // User2 unstakes
        uint user2BalanceBefore = user2.balance;
        vm.prank(user2);
        router.unstakeNative(nativePoolId);

        assertEq(user2.balance, user2BalanceBefore + 20 ether, "User2 got native CROSS");
        assertApproxEqAbs(rewardToken.balanceOf(user2), 60 ether, 1 ether, "User2 rewards (2/3)");
    }

    function testMixedPoolUsage() public {
        // Native pool staking
        vm.startPrank(user1);
        wcross.approve(address(router), type(uint).max);
        router.stakeNative{value: 5 ether}(nativePoolId);
        vm.stopPrank();

        // ERC20 pool staking
        vm.startPrank(user2);
        stakingToken.approve(address(router), 50 ether);
        router.stakeERC20(erc20PoolId, 50 ether);
        vm.stopPrank();

        // Verify both pools
        CrossStakingPool nativePool = CrossStakingPool(nativePoolAddress);
        CrossStakingPool erc20Pool = CrossStakingPool(erc20PoolAddress);

        assertEq(nativePool.balances(user1), 5 ether, "Native pool stake");
        assertEq(erc20Pool.balances(user2), 50 ether, "ERC20 pool stake");
    }
}
