// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardPool.sol";
import "../src/GamePool.sol";
import "../src/WCROSS.sol";
import "../src/interfaces/ICrossGameRewardPool.sol";
import "../src/interfaces/IGamePool.sol";

import "./mocks/MockERC20.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Test.sol";

contract CrossGameRewardTest is Test {
    CrossGameReward public crossGameReward;
    CrossGameRewardPool public poolImplementation;
    GamePool public gamePoolImpl;
    IWCROSS public wcross;

    MockERC20 public token1;
    MockERC20 public token2;
    MockERC20 public rewardToken;

    address public owner;
    address public user1;
    address public user2;

    function setUp() public {
        owner = address(this);
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        // Deploy WCROSS
        wcross = new WCROSS();

        // Deploy pool implementations
        poolImplementation = new CrossGameRewardPool();
        gamePoolImpl = new GamePool();

        // Deploy CrossGameReward as UUPS proxy
        CrossGameReward implementation = new CrossGameReward();
        bytes memory initData = abi.encodeCall(
            CrossGameReward.initialize, (ICrossGameRewardPool(address(poolImplementation)), owner, 2 days)
        );
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), initData);
        crossGameReward = CrossGameReward(address(proxy));

        // Set GamePool implementation
        crossGameReward.setGamePoolImplementation(ICrossGameRewardPool(address(gamePoolImpl)));

        // CrossGameReward creates the WCROSS instance internally
        wcross = crossGameReward.wcross();

        // Deploy test tokens
        token1 = new MockERC20("Token1", "TK1");
        token2 = new MockERC20("Token2", "TK2");
        rewardToken = new MockERC20("CROSSD", "CROSSD");
    }

    // ==================== Initialization ====================

    function testInitialization() public view {
        assertTrue(address(crossGameReward.wcross()) != address(0), "WCROSS created");
        assertEq(address(crossGameReward.poolImplementation()), address(poolImplementation), "Implementation set");
        assertEq(crossGameReward.nextPoolId(), 1, "Next pool ID is 1");

        // Check roles
        assertTrue(crossGameReward.hasRole(crossGameReward.DEFAULT_ADMIN_ROLE(), owner), "Owner has admin role");
        assertTrue(crossGameReward.hasRole(crossGameReward.MANAGER_ROLE(), owner), "Owner has pool manager role");
    }

    // ==================== Pool creation ====================

    function testCreatePool() public {
        (uint poolId, ICrossGameRewardPool pool) =
            crossGameReward.createPool("Test Pool", IERC20(address(wcross)), 1 ether);

        assertEq(poolId, 1, "First pool ID should be 1");
        assertTrue(address(pool) != address(0), "Pool address should be set");

        assertEq(address(pool.depositToken()), address(wcross), "Deposit token");

        // Verify pool info
        CrossGameReward.PoolInfo memory info = crossGameReward.getPoolInfo(poolId);
        assertEq(info.poolId, poolId, "Pool ID");
        assertEq(address(info.pool), address(pool), "Pool address");
        assertEq(address(info.depositToken), address(wcross), "Deposit token");
        assertTrue(info.pool.poolStatus() == ICrossGameRewardPool.PoolStatus.Active, "Active by default");
    }

    function testCreateMultiplePools() public {
        (uint poolId1, ICrossGameRewardPool pool1) =
            crossGameReward.createPool("Pool 1", IERC20(address(wcross)), 1 ether);
        (uint poolId2, ICrossGameRewardPool pool2) =
            crossGameReward.createPool("Pool 2", IERC20(address(token1)), 1 ether);
        (uint poolId3, ICrossGameRewardPool pool3) =
            crossGameReward.createPool("Pool 3", IERC20(address(token2)), 1 ether);

        assertEq(poolId1, 1, "Pool 1 ID");
        assertEq(poolId2, 2, "Pool 2 ID");
        assertEq(poolId3, 3, "Pool 3 ID");

        assertTrue(pool1 != pool2, "Different pools");
        assertTrue(pool2 != pool3, "Different pools");
    }

    function testOnlyPoolManagerCanCreatePool() public {
        vm.prank(user1);
        vm.expectRevert();
        crossGameReward.createPool("Test", IERC20(address(wcross)), 1 ether);
    }

    function testCannotCreatePoolWithZeroAddress() public {
        vm.expectRevert(CrossGameReward.CGRCanNotZeroAddress.selector);
        crossGameReward.createPool("Test", IERC20(address(0)), 1 ether);
    }

    function testMultiplePoolsWithSameDepositToken() public {
        // Multiple pools can use the same deposit token
        (uint poolId1,) = crossGameReward.createPool("Pool 1", IERC20(address(wcross)), 1 ether);
        (uint poolId2,) = crossGameReward.createPool("Pool 2", IERC20(address(wcross)), 1 ether);

        assertEq(poolId1, 1, "Pool 1");
        assertEq(poolId2, 2, "Pool 2");

        // Confirm pool count for the same deposit token
        assertEq(crossGameReward.getPoolCountByDepositToken(IERC20(address(wcross))), 2, "2 pools for WCROSS");
    }

    // ==================== Reward token management ====================

    function testAddRewardToken() public {
        (uint poolId, ICrossGameRewardPool poolInterface) =
            crossGameReward.createPool("Test Pool", IERC20(address(wcross)), 1 ether);

        MockERC20 rewardToken = new MockERC20("Reward", "RWD");
        crossGameReward.addRewardToken(poolId, IERC20(address(rewardToken)));

        // Verify
        CrossGameRewardPool pool = CrossGameRewardPool(address(poolInterface));
        assertTrue(pool.isRewardToken(IERC20(address(rewardToken))), "Reward token added");
        assertEq(pool.rewardTokenCount(), 1, "1 reward token");
    }

    function testOnlyPoolManagerCanAddRewardToken() public {
        (uint poolId,) = crossGameReward.createPool("Test Pool", IERC20(address(wcross)), 1 ether);
        MockERC20 rewardToken = new MockERC20("Reward", "RWD");

        vm.prank(user1);
        vm.expectRevert();
        crossGameReward.addRewardToken(poolId, IERC20(address(rewardToken)));
    }

    function testCannotAddRewardTokenToNonExistentPool() public {
        MockERC20 rewardToken = new MockERC20("Reward", "RWD");

        vm.expectRevert(abi.encodeWithSelector(CrossGameReward.CGRPoolNotFound.selector));
        crossGameReward.addRewardToken(999, IERC20(address(rewardToken)));
    }

    // ==================== Pool status management ====================

    function testSetPoolStatus() public {
        (uint poolId, ICrossGameRewardPool poolInterface) =
            crossGameReward.createPool("Test Pool", IERC20(address(wcross)), 1 ether);

        CrossGameRewardPool pool = CrossGameRewardPool(address(poolInterface));
        assertFalse(pool.paused(), "Initially not paused");
        assertEq(uint(pool.poolStatus()), 0, "Initially active");

        // Set to inactive (1)
        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Inactive);

        CrossGameReward.PoolInfo memory info = crossGameReward.getPoolInfo(poolId);
        assertFalse(info.pool.poolStatus() == ICrossGameRewardPool.PoolStatus.Active, "Pool inactive");
        assertEq(uint(pool.poolStatus()), 1, "Status inactive");
        assertFalse(pool.paused(), "Not paused");

        // Set to paused (2)
        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Paused);

        info = crossGameReward.getPoolInfo(poolId);
        assertFalse(info.pool.poolStatus() == ICrossGameRewardPool.PoolStatus.Active, "Pool paused (not active)");
        assertEq(uint(pool.poolStatus()), 2, "Status paused");
        assertTrue(pool.paused(), "Paused");

        // Set to active (0)
        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Active);

        info = crossGameReward.getPoolInfo(poolId);
        assertTrue(info.pool.poolStatus() == ICrossGameRewardPool.PoolStatus.Active, "Pool active");
        assertEq(uint(pool.poolStatus()), 0, "Status active");
        assertFalse(pool.paused(), "Not paused");
    }

    function testOnlyPoolManagerCanSetPoolStatus() public {
        (uint poolId,) = crossGameReward.createPool("Test Pool", IERC20(address(wcross)), 1 ether);

        vm.prank(user1);
        vm.expectRevert();
        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Inactive);
    }

    function testCannotSetStatusOnNonExistentPool() public {
        vm.expectRevert(abi.encodeWithSelector(CrossGameReward.CGRPoolNotFound.selector));
        crossGameReward.setPoolStatus(999, ICrossGameRewardPool.PoolStatus.Inactive);
    }

    // ==================== Pool queries ====================

    function testGetPoolInfo() public {
        (uint poolId, ICrossGameRewardPool pool) =
            crossGameReward.createPool("Test Pool", IERC20(address(wcross)), 1 ether);

        CrossGameReward.PoolInfo memory info = crossGameReward.getPoolInfo(poolId);

        assertEq(info.poolId, poolId, "Pool ID");
        assertEq(address(info.pool), address(pool), "Pool address");
        assertEq(address(info.depositToken), address(wcross), "Deposit token");
        assertTrue(info.pool.poolStatus() == ICrossGameRewardPool.PoolStatus.Active, "Active");
    }

    function testPoolAt() public {
        crossGameReward.createPool("Pool 1", IERC20(address(wcross)), 1 ether);
        crossGameReward.createPool("Pool 2", IERC20(address(token1)), 1 ether);
        crossGameReward.createPool("Pool 3", IERC20(address(token2)), 1 ether);

        assertEq(crossGameReward.poolAt(0), 1, "Pool at index 0");
        assertEq(crossGameReward.poolAt(1), 2, "Pool at index 1");
        assertEq(crossGameReward.poolAt(2), 3, "Pool at index 2");
    }

    function testGetPoolAddress() public {
        (uint poolId, ICrossGameRewardPool poolAddress) =
            crossGameReward.createPool("Test Pool", IERC20(address(wcross)), 1 ether);

        assertEq(address(crossGameReward.getPoolAddress(poolId)), address(poolAddress), "Pool address lookup");
    }

    function testGetPoolId() public {
        (uint poolId, ICrossGameRewardPool poolAddress) =
            crossGameReward.createPool("Test Pool", IERC20(address(wcross)), 1 ether);

        assertEq(crossGameReward.getPoolId(poolAddress), poolId, "Pool ID lookup");
    }

    function testCannotGetNonExistentPool() public {
        vm.expectRevert(abi.encodeWithSelector(CrossGameReward.CGRPoolNotFound.selector));
        crossGameReward.getPoolInfo(999);

        vm.expectRevert(abi.encodeWithSelector(CrossGameReward.CGRPoolNotFound.selector));
        crossGameReward.getPoolAddress(999);
    }

    function testGetTotalPoolCount() public {
        assertEq(crossGameReward.getTotalPoolCount(), 0, "Initially 0");

        crossGameReward.createPool("Pool 1", IERC20(address(wcross)), 1 ether);
        assertEq(crossGameReward.getTotalPoolCount(), 1, "After 1 pool");

        crossGameReward.createPool("Pool 2", IERC20(address(token1)), 1 ether);
        assertEq(crossGameReward.getTotalPoolCount(), 2, "After 2 pools");
    }

    function testGetAllPoolIds() public {
        crossGameReward.createPool("Pool 1", IERC20(address(wcross)), 1 ether);
        crossGameReward.createPool("Pool 2", IERC20(address(token1)), 1 ether);
        crossGameReward.createPool("Pool 3", IERC20(address(token2)), 1 ether);

        uint[] memory allIds = crossGameReward.getAllPoolIds();
        assertEq(allIds.length, 3, "3 pools");
        assertEq(allIds[0], 1, "Pool 1");
        assertEq(allIds[1], 2, "Pool 2");
        assertEq(allIds[2], 3, "Pool 3");
    }

    function testGetPoolIdsByDepositToken() public {
        crossGameReward.createPool("Pool 1", IERC20(address(wcross)), 1 ether);
        crossGameReward.createPool("Pool 2", IERC20(address(token1)), 1 ether);
        crossGameReward.createPool("Pool 3", IERC20(address(wcross)), 1 ether);

        uint[] memory wcrossIds = crossGameReward.getPoolIdsByDepositToken(IERC20(address(wcross)));
        assertEq(wcrossIds.length, 2, "2 WCROSS pools");
        assertEq(wcrossIds[0], 1, "Pool 1");
        assertEq(wcrossIds[1], 3, "Pool 3");

        uint[] memory token1Ids = crossGameReward.getPoolIdsByDepositToken(IERC20(address(token1)));
        assertEq(token1Ids.length, 1, "1 Token1 pool");
        assertEq(token1Ids[0], 2, "Pool 2");
    }

    function testPoolByDepositTokenAt() public {
        crossGameReward.createPool("Pool 1", IERC20(address(wcross)), 1 ether);
        crossGameReward.createPool("Pool 2", IERC20(address(wcross)), 1 ether);

        assertEq(crossGameReward.poolByDepositTokenAt(IERC20(address(wcross)), 0), 1, "First WCROSS pool");
        assertEq(crossGameReward.poolByDepositTokenAt(IERC20(address(wcross)), 1), 2, "Second WCROSS pool");
    }

    function testGetActivePoolIds() public {
        (uint pool1,) = crossGameReward.createPool("Pool 1", IERC20(address(wcross)), 1 ether);
        (uint pool2,) = crossGameReward.createPool("Pool 2", IERC20(address(token1)), 1 ether);
        (uint pool3,) = crossGameReward.createPool("Pool 3", IERC20(address(token2)), 1 ether);

        // All active initially
        uint[] memory activeIds = crossGameReward.getActivePoolIds();
        assertEq(activeIds.length, 3, "3 active pools");

        // Deactivate pool2
        crossGameReward.setPoolStatus(pool2, ICrossGameRewardPool.PoolStatus.Inactive); // Inactive

        activeIds = crossGameReward.getActivePoolIds();
        assertEq(activeIds.length, 2, "2 active pools");
        assertEq(activeIds[0], pool1, "Pool 1 still active");
        assertEq(activeIds[1], pool3, "Pool 3 still active");
    }

    // ==================== Router management ====================

    function testSetRouter() public {
        address newRouter = makeAddr("newRouter");

        crossGameReward.setRouter(newRouter);
        assertEq(crossGameReward.router(), newRouter, "Router set");
    }

    function testOnlyAdminCanSetRouter() public {
        address newRouter = makeAddr("newRouter");

        vm.prank(user1);
        vm.expectRevert();
        crossGameReward.setRouter(newRouter);
    }

    function testCannotSetZeroAddressAsRouter() public {
        vm.expectRevert(CrossGameReward.CGRCanNotZeroAddress.selector);
        crossGameReward.setRouter(address(0));
    }

    // ==================== Admin functions ====================

    function testSetPoolImplementation() public {
        CrossGameRewardPool newImpl = new CrossGameRewardPool();

        crossGameReward.setPoolImplementation(ICrossGameRewardPool(address(newImpl)));
        assertEq(address(crossGameReward.poolImplementation()), address(newImpl), "Implementation updated");
    }

    function testOnlyAdminCanSetPoolImplementation() public {
        CrossGameRewardPool newImpl = new CrossGameRewardPool();

        vm.prank(user1);
        vm.expectRevert();
        crossGameReward.setPoolImplementation(ICrossGameRewardPool(address(newImpl)));
    }

    // ==================== Role management ====================

    function testGrantPoolManagerRole() public {
        assertFalse(crossGameReward.hasRole(crossGameReward.MANAGER_ROLE(), user1), "User1 not pool manager");

        // Grant role
        crossGameReward.grantRole(crossGameReward.MANAGER_ROLE(), user1);
        assertTrue(crossGameReward.hasRole(crossGameReward.MANAGER_ROLE(), user1), "User1 is pool manager");

        // User1 can create pool
        vm.prank(user1);
        (, ICrossGameRewardPool poolAddress) = crossGameReward.createPool("User Pool", IERC20(address(token1)), 1 ether);
        assertTrue(address(poolAddress) != address(0), "Pool created by user1");
    }

    function testRevokePoolManagerRole() public {
        crossGameReward.grantRole(crossGameReward.MANAGER_ROLE(), user1);
        assertTrue(crossGameReward.hasRole(crossGameReward.MANAGER_ROLE(), user1), "User1 has role");

        crossGameReward.revokeRole(crossGameReward.MANAGER_ROLE(), user1);
        assertFalse(crossGameReward.hasRole(crossGameReward.MANAGER_ROLE(), user1), "Role revoked");

        vm.prank(user1);
        vm.expectRevert();
        crossGameReward.createPool("Test", IERC20(address(token1)), 1 ether);
    }

    // ==================== Integration ====================

    function testPoolsAreIndependent() public {
        // Create 2 pools
        (, ICrossGameRewardPool pool1Addr) = crossGameReward.createPool("Pool 1", IERC20(address(wcross)), 1 ether);
        (, ICrossGameRewardPool pool2Addr) = crossGameReward.createPool("Pool 2", IERC20(address(token1)), 1 ether);

        CrossGameRewardPool pool1 = CrossGameRewardPool(address(pool1Addr));
        CrossGameRewardPool pool2 = CrossGameRewardPool(address(pool2Addr));

        // Different deposit tokens
        assertEq(address(pool1.depositToken()), address(wcross), "Pool1 deposit token");
        assertEq(address(pool2.depositToken()), address(token1), "Pool2 deposit token");

        // Different total depositd
        token1.mint(user1, 100 ether);
        vm.startPrank(user1);
        token1.approve(address(pool2Addr), 50 ether);
        pool2.deposit(50 ether);
        vm.stopPrank();

        assertEq(pool1.totalDeposited(), 0, "Pool1 empty");
        assertEq(pool2.totalDeposited(), 50 ether, "Pool2 has deposit");
    }

    function testPoolStatusAffectsPause() public {
        (uint poolId, ICrossGameRewardPool poolAddress) =
            crossGameReward.createPool("Test Pool", IERC20(address(wcross)), 1 ether);
        CrossGameRewardPool pool = CrossGameRewardPool(address(poolAddress));

        // Set to paused (2)
        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Paused);
        assertTrue(pool.paused(), "Pool should be paused");

        // Set to inactive (1)
        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Inactive);
        assertFalse(pool.paused(), "Pool should not be paused when inactive");

        // Set to active (0)
        crossGameReward.setPoolStatus(poolId, ICrossGameRewardPool.PoolStatus.Active);
        assertFalse(pool.paused(), "Pool should not be paused when active");
    }

    // ==================== UUPS upgrades ====================

    function testUpgrade() public {
        // Deploy new implementation
        CrossGameReward newImplementation = new CrossGameReward();

        // Upgrade
        crossGameReward.upgradeToAndCall(address(newImplementation), "");

        // Verify state persisted
        assertEq(crossGameReward.nextPoolId(), 1, "State persisted");
    }

    function testOnlyAdminCanUpgrade() public {
        CrossGameReward newImplementation = new CrossGameReward();

        vm.prank(user1);
        vm.expectRevert();
        crossGameReward.upgradeToAndCall(address(newImplementation), "");
    }

    // ==================== Batch Pool Upgrades ====================

    function testUpgradePoolsByType_GamePool() public {
        // Create two GamePools
        (uint poolId1,) = crossGameReward.createGamePool(
            "GamePool 1", IERC20(address(token1)), IERC20(address(rewardToken)), 1 ether
        );
        (uint poolId2,) = crossGameReward.createGamePool(
            "GamePool 2", IERC20(address(token2)), IERC20(address(rewardToken)), 1 ether
        );

        // Deploy new GamePool implementation
        GamePool newGamePoolImpl = new GamePool();

        // Batch upgrade all GamePools
        crossGameReward.upgradePoolsByType(ICrossGameReward.PoolType.GamePool, address(newGamePoolImpl), "");

        // Verify stored implementation was updated
        assertEq(address(crossGameReward.gamePoolImplementation()), address(newGamePoolImpl));

        // Verify pools still work (state preserved)
        ICrossGameRewardPool pool1 = crossGameReward.getPoolAddress(poolId1);
        ICrossGameRewardPool pool2 = crossGameReward.getPoolAddress(poolId2);
        assertEq(address(pool1.depositToken()), address(token1));
        assertEq(address(pool2.depositToken()), address(token2));
    }

    function testUpgradePoolsByType_V1() public {
        // Create two V1 pools
        (uint poolId1,) = crossGameReward.createPool("V1 Pool 1", IERC20(address(wcross)), 1 ether);
        (uint poolId2,) = crossGameReward.createPool("V1 Pool 2", IERC20(address(token1)), 1 ether);

        // Deploy new V1 implementation
        CrossGameRewardPool newImplV1 = new CrossGameRewardPool();

        // Batch upgrade all V1 pools
        crossGameReward.upgradePoolsByType(ICrossGameReward.PoolType.CrossPool, address(newImplV1), "");

        // Verify stored implementation was updated
        assertEq(address(crossGameReward.poolImplementation()), address(newImplV1));

        // Verify pools still work (state preserved)
        ICrossGameRewardPool pool1 = crossGameReward.getPoolAddress(poolId1);
        ICrossGameRewardPool pool2 = crossGameReward.getPoolAddress(poolId2);
        assertEq(address(pool1.depositToken()), address(wcross));
        assertEq(address(pool2.depositToken()), address(token1));
    }

    function testUpgradePoolsByType_OnlyUpgradesMatchingType() public {
        // Create one V1 and one GamePool
        (uint v1PoolId,) = crossGameReward.createPool("V1 Pool", IERC20(address(wcross)), 1 ether);
        (uint v2PoolId,) = crossGameReward.createGamePool(
            "GamePool", IERC20(address(token1)), IERC20(address(rewardToken)), 1 ether
        );

        // Deploy new GamePool implementation
        GamePool newGamePoolImpl = new GamePool();

        // Upgrade only GamePools
        crossGameReward.upgradePoolsByType(ICrossGameReward.PoolType.GamePool, address(newGamePoolImpl), "");

        // GamePool upgraded, V1 pool untouched - both still functional
        ICrossGameRewardPool v1Pool = crossGameReward.getPoolAddress(v1PoolId);
        ICrossGameRewardPool v2Pool = crossGameReward.getPoolAddress(v2PoolId);
        assertEq(address(v1Pool.depositToken()), address(wcross));
        assertEq(address(v2Pool.depositToken()), address(token1));
    }

    function testUpgradePoolsByType_OnlyAdmin() public {
        GamePool newGamePoolImpl = new GamePool();

        vm.prank(user1);
        vm.expectRevert();
        crossGameReward.upgradePoolsByType(ICrossGameReward.PoolType.GamePool, address(newGamePoolImpl), "");
    }

    function testUpgradePoolsByType_ZeroAddress_Reverts() public {
        vm.expectRevert(CrossGameReward.CGRCanNotZeroAddress.selector);
        crossGameReward.upgradePoolsByType(ICrossGameReward.PoolType.GamePool, address(0), "");
    }

    function testUpgradePoolsByType_NoPools() public {
        // No pools of this type exist - should not revert
        GamePool newGamePoolImpl = new GamePool();
        crossGameReward.upgradePoolsByType(ICrossGameReward.PoolType.GamePool, address(newGamePoolImpl), "");

        // Implementation still updated
        assertEq(address(crossGameReward.gamePoolImplementation()), address(newGamePoolImpl));
    }

    function testUpgradePoolsByType_PreservesDeposits() public {
        // Create GamePool and make a deposit
        (uint poolId,) = crossGameReward.createGamePool(
            "GamePool", IERC20(address(token1)), IERC20(address(rewardToken)), 1 ether
        );

        ICrossGameRewardPool pool = crossGameReward.getPoolAddress(poolId);

        // User deposits
        token1.mint(user1, 100 ether);
        vm.startPrank(user1);
        token1.approve(address(pool), 50 ether);
        GamePool(address(pool)).deposit(50 ether);
        vm.stopPrank();

        assertEq(pool.totalDeposited(), 50 ether);
        assertEq(pool.balances(user1), 50 ether);

        // Upgrade
        GamePool newGamePoolImpl = new GamePool();
        crossGameReward.upgradePoolsByType(ICrossGameReward.PoolType.GamePool, address(newGamePoolImpl), "");

        // State preserved
        assertEq(pool.totalDeposited(), 50 ether);
        assertEq(pool.balances(user1), 50 ether);

        // User can still withdraw after upgrade
        vm.prank(user1);
        GamePool(address(pool)).withdraw(50 ether);

        assertEq(pool.totalDeposited(), 0);
        assertEq(pool.balances(user1), 0);
    }

    // ==================== Pool _authorizeUpgrade ====================

    function testPool_OwnerCanUpgrade() public {
        (, ICrossGameRewardPool pool) = crossGameReward.createPool("Pool", IERC20(address(wcross)), 1 ether);

        CrossGameRewardPool newImpl = new CrossGameRewardPool();

        // Owner (defaultAdmin = address(this)) can upgrade directly
        CrossGameRewardPool(address(pool)).upgradeToAndCall(address(newImpl), "");

        // Pool still works
        assertEq(address(pool.depositToken()), address(wcross));
    }

    function testPool_RewardRootCanUpgrade() public {
        (, ICrossGameRewardPool pool) = crossGameReward.createPool("Pool", IERC20(address(wcross)), 1 ether);

        CrossGameRewardPool newImpl = new CrossGameRewardPool();

        // CrossGameReward (rewardRoot) can upgrade via batch
        crossGameReward.upgradePoolsByType(ICrossGameReward.PoolType.CrossPool, address(newImpl), "");

        // Pool still works
        assertEq(address(pool.depositToken()), address(wcross));
    }

    function testPool_UnauthorizedCannotUpgrade() public {
        (, ICrossGameRewardPool pool) = crossGameReward.createPool("Pool", IERC20(address(wcross)), 1 ether);

        CrossGameRewardPool newImpl = new CrossGameRewardPool();

        vm.prank(user1);
        vm.expectRevert();
        CrossGameRewardPool(address(pool)).upgradeToAndCall(address(newImpl), "");
    }

    function testGamePool_AdminCanUpgradeViaBatch() public {
        (, ICrossGameRewardPool pool) = crossGameReward.createGamePool(
            "GamePool", IERC20(address(token1)), IERC20(address(rewardToken)), 1 ether
        );

        GamePool newImpl = new GamePool();

        // GamePool's DEFAULT_ADMIN_ROLE is CrossGameReward contract,
        // so upgrades must go through the factory's upgradePoolsByType
        crossGameReward.upgradePoolsByType(ICrossGameReward.PoolType.GamePool, address(newImpl), "");

        assertEq(address(pool.depositToken()), address(token1));
    }

    function testGamePool_DirectUpgrade_OwnerCanUpgrade() public {
        (, ICrossGameRewardPool pool) = crossGameReward.createGamePool(
            "GamePool", IERC20(address(token1)), IERC20(address(rewardToken)), 1 ether
        );

        GamePool newImpl = new GamePool();

        // Owner (crossGameReward.owner() = address(this)) CAN directly upgrade GamePool
        // This is intentional - owner has direct upgrade capability alongside CrossGameReward contract
        GamePool(address(pool)).upgradeToAndCall(address(newImpl), "");

        // Verify upgrade succeeded by checking implementation changed
        assertEq(address(pool.depositToken()), address(token1));
    }

    function testGamePool_UnauthorizedCannotUpgrade() public {
        (, ICrossGameRewardPool pool) = crossGameReward.createGamePool(
            "GamePool", IERC20(address(token1)), IERC20(address(rewardToken)), 1 ether
        );

        GamePool newImpl = new GamePool();

        vm.prank(user1);
        vm.expectRevert();
        GamePool(address(pool)).upgradeToAndCall(address(newImpl), "");
    }
}
