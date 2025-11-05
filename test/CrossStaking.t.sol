// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/CrossStaking.sol";
import "../src/CrossStakingPool.sol";
import "../src/WCROSS.sol";
import "../src/interfaces/ICrossStakingPool.sol";

import "./mocks/MockERC20.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Test.sol";

contract CrossStakingTest is Test {
    CrossStaking public crossStaking;
    CrossStakingPool public poolImplementation;
    IWCROSS public wcross;

    MockERC20 public token1;
    MockERC20 public token2;

    address public owner;
    address public user1;
    address public user2;

    function setUp() public {
        owner = address(this);
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        // Deploy WCROSS
        wcross = new WCROSS();

        // Deploy pool implementation
        poolImplementation = new CrossStakingPool();

        // Deploy CrossStaking as UUPS proxy
        CrossStaking implementation = new CrossStaking();
        bytes memory initData =
            abi.encodeCall(CrossStaking.initialize, (ICrossStakingPool(address(poolImplementation)), owner, 2 days));
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), initData);
        crossStaking = CrossStaking(address(proxy));

        // CrossStaking creates the WCROSS instance internally
        wcross = crossStaking.wcross();

        // Deploy test tokens
        token1 = new MockERC20("Token1", "TK1");
        token2 = new MockERC20("Token2", "TK2");
    }

    // ==================== Initialization ====================

    function testInitialization() public view {
        assertTrue(address(crossStaking.wcross()) != address(0), "WCROSS created");
        assertEq(address(crossStaking.poolImplementation()), address(poolImplementation), "Implementation set");
        assertEq(crossStaking.nextPoolId(), 1, "Next pool ID is 1");

        // Check roles
        assertTrue(crossStaking.hasRole(crossStaking.DEFAULT_ADMIN_ROLE(), owner), "Owner has admin role");
        assertTrue(crossStaking.hasRole(crossStaking.MANAGER_ROLE(), owner), "Owner has pool manager role");
    }

    // ==================== Pool creation ====================

    function testCreatePool() public {
        (uint poolId, ICrossStakingPool pool) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);

        assertEq(poolId, 1, "First pool ID should be 1");
        assertTrue(address(pool) != address(0), "Pool address should be set");

        assertEq(address(pool.stakingToken()), address(wcross), "Staking token");

        // Verify pool info
        CrossStaking.PoolInfo memory info = crossStaking.getPoolInfo(poolId);
        assertEq(info.poolId, poolId, "Pool ID");
        assertEq(address(info.pool), address(pool), "Pool address");
        assertEq(address(info.stakingToken), address(wcross), "Staking token");
        assertTrue(info.active, "Active by default");
    }

    function testCreateMultiplePools() public {
        (uint poolId1, ICrossStakingPool pool1) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        (uint poolId2, ICrossStakingPool pool2) = crossStaking.createPool(IERC20(address(token1)), 1 ether);
        (uint poolId3, ICrossStakingPool pool3) = crossStaking.createPool(IERC20(address(token2)), 1 ether);

        assertEq(poolId1, 1, "Pool 1 ID");
        assertEq(poolId2, 2, "Pool 2 ID");
        assertEq(poolId3, 3, "Pool 3 ID");

        assertTrue(pool1 != pool2, "Different pools");
        assertTrue(pool2 != pool3, "Different pools");
    }

    function testOnlyPoolManagerCanCreatePool() public {
        vm.prank(user1);
        vm.expectRevert();
        crossStaking.createPool(IERC20(address(wcross)), 1 ether);
    }

    function testCannotCreatePoolWithZeroAddress() public {
        vm.expectRevert(CrossStaking.CSCanNotZeroAddress.selector);
        crossStaking.createPool(IERC20(address(0)), 1 ether);
    }

    function testMultiplePoolsWithSameStakingToken() public {
        // Multiple pools can use the same staking token
        (uint poolId1,) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        (uint poolId2,) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);

        assertEq(poolId1, 1, "Pool 1");
        assertEq(poolId2, 2, "Pool 2");

        // Confirm pool count for the same staking token
        assertEq(crossStaking.getPoolCountByStakingToken(IERC20(address(wcross))), 2, "2 pools for WCROSS");
    }

    // ==================== Reward token management ====================

    function testAddRewardToken() public {
        (uint poolId, ICrossStakingPool poolInterface) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);

        MockERC20 rewardToken = new MockERC20("Reward", "RWD");
        crossStaking.addRewardToken(poolId, IERC20(address(rewardToken)));

        // Verify
        CrossStakingPool pool = CrossStakingPool(address(poolInterface));
        assertTrue(pool.isRewardToken(IERC20(address(rewardToken))), "Reward token added");
        assertEq(pool.rewardTokenCount(), 1, "1 reward token");
    }

    function testOnlyPoolManagerCanAddRewardToken() public {
        (uint poolId,) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        MockERC20 rewardToken = new MockERC20("Reward", "RWD");

        vm.prank(user1);
        vm.expectRevert();
        crossStaking.addRewardToken(poolId, IERC20(address(rewardToken)));
    }

    function testCannotAddRewardTokenToNonExistentPool() public {
        MockERC20 rewardToken = new MockERC20("Reward", "RWD");

        vm.expectRevert(CrossStaking.CSPoolNotFound.selector);
        crossStaking.addRewardToken(999, IERC20(address(rewardToken)));
    }

    // ==================== Pool status management ====================

    function testSetPoolStatus() public {
        (uint poolId, ICrossStakingPool poolInterface) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);

        CrossStakingPool pool = CrossStakingPool(address(poolInterface));
        assertFalse(pool.paused(), "Initially not paused");
        assertEq(uint(pool.poolStatus()), 0, "Initially active");

        // Set to inactive (1)
        crossStaking.setPoolStatus(poolId, ICrossStakingPool.PoolStatus.Inactive);

        CrossStaking.PoolInfo memory info = crossStaking.getPoolInfo(poolId);
        assertFalse(info.active, "Pool inactive");
        assertEq(uint(pool.poolStatus()), 1, "Status inactive");
        assertFalse(pool.paused(), "Not paused");

        // Set to paused (2)
        crossStaking.setPoolStatus(poolId, ICrossStakingPool.PoolStatus.Paused);

        info = crossStaking.getPoolInfo(poolId);
        assertFalse(info.active, "Pool paused (not active)");
        assertEq(uint(pool.poolStatus()), 2, "Status paused");
        assertTrue(pool.paused(), "Paused");

        // Set to active (0)
        crossStaking.setPoolStatus(poolId, ICrossStakingPool.PoolStatus.Active);

        info = crossStaking.getPoolInfo(poolId);
        assertTrue(info.active, "Pool active");
        assertEq(uint(pool.poolStatus()), 0, "Status active");
        assertFalse(pool.paused(), "Not paused");
    }

    function testOnlyPoolManagerCanSetPoolStatus() public {
        (uint poolId,) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);

        vm.prank(user1);
        vm.expectRevert();
        crossStaking.setPoolStatus(poolId, ICrossStakingPool.PoolStatus.Inactive);
    }

    function testCannotSetStatusOnNonExistentPool() public {
        vm.expectRevert(CrossStaking.CSPoolNotFound.selector);
        crossStaking.setPoolStatus(999, ICrossStakingPool.PoolStatus.Inactive);
    }

    // ==================== Pool queries ====================

    function testGetPoolInfo() public {
        (uint poolId, ICrossStakingPool pool) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);

        CrossStaking.PoolInfo memory info = crossStaking.getPoolInfo(poolId);

        assertEq(info.poolId, poolId, "Pool ID");
        assertEq(address(info.pool), address(pool), "Pool address");
        assertEq(address(info.stakingToken), address(wcross), "Staking token");
        assertTrue(info.active, "Active");
    }

    function testPoolAt() public {
        crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        crossStaking.createPool(IERC20(address(token1)), 1 ether);
        crossStaking.createPool(IERC20(address(token2)), 1 ether);

        assertEq(crossStaking.poolAt(0), 1, "Pool at index 0");
        assertEq(crossStaking.poolAt(1), 2, "Pool at index 1");
        assertEq(crossStaking.poolAt(2), 3, "Pool at index 2");
    }

    function testGetPoolAddress() public {
        (uint poolId, ICrossStakingPool poolAddress) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);

        assertEq(address(crossStaking.getPoolAddress(poolId)), address(poolAddress), "Pool address lookup");
    }

    function testGetPoolId() public {
        (uint poolId, ICrossStakingPool poolAddress) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);

        assertEq(crossStaking.getPoolId(poolAddress), poolId, "Pool ID lookup");
    }

    function testCannotGetNonExistentPool() public {
        vm.expectRevert(CrossStaking.CSPoolNotFound.selector);
        crossStaking.getPoolInfo(999);

        vm.expectRevert(CrossStaking.CSPoolNotFound.selector);
        crossStaking.getPoolAddress(999);
    }

    function testGetTotalPoolCount() public {
        assertEq(crossStaking.getTotalPoolCount(), 0, "Initially 0");

        crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        assertEq(crossStaking.getTotalPoolCount(), 1, "After 1 pool");

        crossStaking.createPool(IERC20(address(token1)), 1 ether);
        assertEq(crossStaking.getTotalPoolCount(), 2, "After 2 pools");
    }

    function testGetAllPoolIds() public {
        crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        crossStaking.createPool(IERC20(address(token1)), 1 ether);
        crossStaking.createPool(IERC20(address(token2)), 1 ether);

        uint[] memory allIds = crossStaking.getAllPoolIds();
        assertEq(allIds.length, 3, "3 pools");
        assertEq(allIds[0], 1, "Pool 1");
        assertEq(allIds[1], 2, "Pool 2");
        assertEq(allIds[2], 3, "Pool 3");
    }

    function testGetPoolIdsByStakingToken() public {
        crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        crossStaking.createPool(IERC20(address(token1)), 1 ether);
        crossStaking.createPool(IERC20(address(wcross)), 1 ether);

        uint[] memory wcrossIds = crossStaking.getPoolIdsByStakingToken(IERC20(address(wcross)));
        assertEq(wcrossIds.length, 2, "2 WCROSS pools");
        assertEq(wcrossIds[0], 1, "Pool 1");
        assertEq(wcrossIds[1], 3, "Pool 3");

        uint[] memory token1Ids = crossStaking.getPoolIdsByStakingToken(IERC20(address(token1)));
        assertEq(token1Ids.length, 1, "1 Token1 pool");
        assertEq(token1Ids[0], 2, "Pool 2");
    }

    function testPoolByStakingTokenAt() public {
        crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        crossStaking.createPool(IERC20(address(wcross)), 1 ether);

        assertEq(crossStaking.poolByStakingTokenAt(IERC20(address(wcross)), 0), 1, "First WCROSS pool");
        assertEq(crossStaking.poolByStakingTokenAt(IERC20(address(wcross)), 1), 2, "Second WCROSS pool");
    }

    function testGetActivePoolIds() public {
        (uint pool1,) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        (uint pool2,) = crossStaking.createPool(IERC20(address(token1)), 1 ether);
        (uint pool3,) = crossStaking.createPool(IERC20(address(token2)), 1 ether);

        // All active initially
        uint[] memory activeIds = crossStaking.getActivePoolIds();
        assertEq(activeIds.length, 3, "3 active pools");

        // Deactivate pool2
        crossStaking.setPoolStatus(pool2, ICrossStakingPool.PoolStatus.Inactive); // Inactive

        activeIds = crossStaking.getActivePoolIds();
        assertEq(activeIds.length, 2, "2 active pools");
        assertEq(activeIds[0], pool1, "Pool 1 still active");
        assertEq(activeIds[1], pool3, "Pool 3 still active");
    }

    // ==================== Router management ====================

    function testSetRouter() public {
        address newRouter = makeAddr("newRouter");

        crossStaking.setRouter(newRouter);
        assertEq(crossStaking.router(), newRouter, "Router set");
    }

    function testOnlyAdminCanSetRouter() public {
        address newRouter = makeAddr("newRouter");

        vm.prank(user1);
        vm.expectRevert();
        crossStaking.setRouter(newRouter);
    }

    function testCannotSetZeroAddressAsRouter() public {
        vm.expectRevert(CrossStaking.CSCanNotZeroAddress.selector);
        crossStaking.setRouter(address(0));
    }

    // ==================== Admin functions ====================

    function testSetPoolImplementation() public {
        CrossStakingPool newImpl = new CrossStakingPool();

        crossStaking.setPoolImplementation(ICrossStakingPool(address(newImpl)));
        assertEq(address(crossStaking.poolImplementation()), address(newImpl), "Implementation updated");
    }

    function testOnlyAdminCanSetPoolImplementation() public {
        CrossStakingPool newImpl = new CrossStakingPool();

        vm.prank(user1);
        vm.expectRevert();
        crossStaking.setPoolImplementation(ICrossStakingPool(address(newImpl)));
    }

    // ==================== Role management ====================

    function testGrantPoolManagerRole() public {
        assertFalse(crossStaking.hasRole(crossStaking.MANAGER_ROLE(), user1), "User1 not pool manager");

        // Grant role
        crossStaking.grantRole(crossStaking.MANAGER_ROLE(), user1);
        assertTrue(crossStaking.hasRole(crossStaking.MANAGER_ROLE(), user1), "User1 is pool manager");

        // User1 can create pool
        vm.prank(user1);
        (, ICrossStakingPool poolAddress) = crossStaking.createPool(IERC20(address(token1)), 1 ether);
        assertTrue(address(poolAddress) != address(0), "Pool created by user1");
    }

    function testRevokePoolManagerRole() public {
        crossStaking.grantRole(crossStaking.MANAGER_ROLE(), user1);
        assertTrue(crossStaking.hasRole(crossStaking.MANAGER_ROLE(), user1), "User1 has role");

        crossStaking.revokeRole(crossStaking.MANAGER_ROLE(), user1);
        assertFalse(crossStaking.hasRole(crossStaking.MANAGER_ROLE(), user1), "Role revoked");

        vm.prank(user1);
        vm.expectRevert();
        crossStaking.createPool(IERC20(address(token1)), 1 ether);
    }

    // ==================== Integration ====================

    function testPoolsAreIndependent() public {
        // Create 2 pools
        (, ICrossStakingPool pool1Addr) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        (, ICrossStakingPool pool2Addr) = crossStaking.createPool(IERC20(address(token1)), 1 ether);

        CrossStakingPool pool1 = CrossStakingPool(address(pool1Addr));
        CrossStakingPool pool2 = CrossStakingPool(address(pool2Addr));

        // Different staking tokens
        assertEq(address(pool1.stakingToken()), address(wcross), "Pool1 staking token");
        assertEq(address(pool2.stakingToken()), address(token1), "Pool2 staking token");

        // Different total staked
        token1.mint(user1, 100 ether);
        vm.startPrank(user1);
        token1.approve(address(pool2Addr), 50 ether);
        pool2.stake(50 ether);
        vm.stopPrank();

        assertEq(pool1.totalStaked(), 0, "Pool1 empty");
        assertEq(pool2.totalStaked(), 50 ether, "Pool2 has stake");
    }

    function testPoolStatusAffectsPause() public {
        (uint poolId, ICrossStakingPool poolAddress) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        CrossStakingPool pool = CrossStakingPool(address(poolAddress));

        // Set to paused (2)
        crossStaking.setPoolStatus(poolId, ICrossStakingPool.PoolStatus.Paused);
        assertTrue(pool.paused(), "Pool should be paused");

        // Set to inactive (1)
        crossStaking.setPoolStatus(poolId, ICrossStakingPool.PoolStatus.Inactive);
        assertFalse(pool.paused(), "Pool should not be paused when inactive");

        // Set to active (0)
        crossStaking.setPoolStatus(poolId, ICrossStakingPool.PoolStatus.Active);
        assertFalse(pool.paused(), "Pool should not be paused when active");
    }

    // ==================== UUPS upgrades ====================

    function testUpgrade() public {
        // Deploy new implementation
        CrossStaking newImplementation = new CrossStaking();

        // Upgrade
        crossStaking.upgradeToAndCall(address(newImplementation), "");

        // Verify state persisted
        assertEq(crossStaking.nextPoolId(), 1, "State persisted");
    }

    function testOnlyAdminCanUpgrade() public {
        CrossStaking newImplementation = new CrossStaking();

        vm.prank(user1);
        vm.expectRevert();
        crossStaking.upgradeToAndCall(address(newImplementation), "");
    }
}
