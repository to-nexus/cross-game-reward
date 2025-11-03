// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/CrossStaking.sol";
import "../src/CrossStakingPool.sol";
import "../src/WCROSS.sol";
import "./mocks/MockERC20.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "forge-std/Test.sol";

contract CrossStakingTest is Test {
    CrossStaking public crossStaking;
    CrossStakingPool public poolImplementation;
    WCROSS public wcross;

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
        bytes memory initData = abi.encodeCall(CrossStaking.initialize, (address(poolImplementation), owner, 2 days));
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), initData);
        crossStaking = CrossStaking(address(proxy));

        // WCROSS는 CrossStaking이 생성함
        wcross = WCROSS(payable(crossStaking.wcross()));

        // Deploy test tokens
        token1 = new MockERC20("Token1", "TK1");
        token2 = new MockERC20("Token2", "TK2");
    }

    // ==================== 초기화 ====================

    function testInitialization() public view {
        assertTrue(address(crossStaking.wcross()) != address(0), "WCROSS created");
        assertEq(crossStaking.poolImplementation(), address(poolImplementation), "Implementation set");
        assertEq(crossStaking.nextPoolId(), 1, "Next pool ID is 1");

        // Check roles
        assertTrue(crossStaking.hasRole(crossStaking.DEFAULT_ADMIN_ROLE(), owner), "Owner has admin role");
        assertTrue(crossStaking.hasRole(crossStaking.MANAGER_ROLE(), owner), "Owner has pool manager role");
    }

    // ==================== 풀 생성 ====================

    function testCreatePool() public {
        (uint poolId, address poolAddress) = crossStaking.createPool(address(wcross), 1 ether);

        assertEq(poolId, 1, "First pool ID should be 1");
        assertTrue(poolAddress != address(0), "Pool address should be set");

        CrossStakingPool pool = CrossStakingPool(poolAddress);
        assertEq(address(pool.stakingToken()), address(wcross), "Staking token");

        // Verify pool info
        CrossStaking.PoolInfo memory info = crossStaking.getPoolInfo(poolId);
        assertEq(info.poolId, poolId, "Pool ID");
        assertEq(info.poolAddress, poolAddress, "Pool address");
        assertEq(info.stakingToken, address(wcross), "Staking token");
        assertTrue(info.active, "Active by default");
    }

    function testCreateMultiplePools() public {
        (uint poolId1, address pool1) = crossStaking.createPool(address(wcross), 1 ether);
        (uint poolId2, address pool2) = crossStaking.createPool(address(token1), 1 ether);
        (uint poolId3, address pool3) = crossStaking.createPool(address(token2), 1 ether);

        assertEq(poolId1, 1, "Pool 1 ID");
        assertEq(poolId2, 2, "Pool 2 ID");
        assertEq(poolId3, 3, "Pool 3 ID");

        assertTrue(pool1 != pool2, "Different pools");
        assertTrue(pool2 != pool3, "Different pools");
    }

    function testOnlyPoolManagerCanCreatePool() public {
        vm.prank(user1);
        vm.expectRevert();
        crossStaking.createPool(address(wcross), 1 ether);
    }

    function testCannotCreatePoolWithZeroAddress() public {
        vm.expectRevert(CrossStaking.CSCanNotZeroAddress.selector);
        crossStaking.createPool(address(0), 1 ether);
    }

    function testMultiplePoolsWithSameStakingToken() public {
        // 같은 스테이킹 토큰으로 여러 풀 생성 가능
        (uint poolId1,) = crossStaking.createPool(address(wcross), 1 ether);
        (uint poolId2,) = crossStaking.createPool(address(wcross), 1 ether);

        assertEq(poolId1, 1, "Pool 1");
        assertEq(poolId2, 2, "Pool 2");

        // 같은 토큰의 풀 개수 확인
        assertEq(crossStaking.getPoolCountByStakingToken(address(wcross)), 2, "2 pools for WCROSS");
    }

    // ==================== 보상 토큰 관리 ====================

    function testAddRewardToken() public {
        (uint poolId, address poolAddress) = crossStaking.createPool(address(wcross), 1 ether);

        MockERC20 rewardToken = new MockERC20("Reward", "RWD");
        crossStaking.addRewardToken(poolId, address(rewardToken));

        // Verify
        CrossStakingPool pool = CrossStakingPool(poolAddress);
        assertTrue(pool.isRewardToken(address(rewardToken)), "Reward token added");
        assertEq(pool.rewardTokenCount(), 1, "1 reward token");
    }

    function testOnlyPoolManagerCanAddRewardToken() public {
        (uint poolId,) = crossStaking.createPool(address(wcross), 1 ether);
        MockERC20 rewardToken = new MockERC20("Reward", "RWD");

        vm.prank(user1);
        vm.expectRevert();
        crossStaking.addRewardToken(poolId, address(rewardToken));
    }

    function testCannotAddRewardTokenToNonExistentPool() public {
        MockERC20 rewardToken = new MockERC20("Reward", "RWD");

        vm.expectRevert(CrossStaking.CSPoolNotFound.selector);
        crossStaking.addRewardToken(999, address(rewardToken));
    }

    // ==================== 풀 활성화/비활성화 ====================

    function testSetPoolActive() public {
        (uint poolId, address poolAddress) = crossStaking.createPool(address(wcross), 1 ether);

        CrossStakingPool pool = CrossStakingPool(poolAddress);
        assertFalse(pool.paused(), "Initially not paused");

        // Deactivate
        crossStaking.setPoolActive(poolId, false);

        CrossStaking.PoolInfo memory info = crossStaking.getPoolInfo(poolId);
        assertFalse(info.active, "Pool deactivated");
        assertTrue(pool.paused(), "Pool paused");

        // Reactivate
        crossStaking.setPoolActive(poolId, true);

        info = crossStaking.getPoolInfo(poolId);
        assertTrue(info.active, "Pool reactivated");
        assertFalse(pool.paused(), "Pool unpaused");
    }

    function testOnlyPoolManagerCanSetPoolActive() public {
        (uint poolId,) = crossStaking.createPool(address(wcross), 1 ether);

        vm.prank(user1);
        vm.expectRevert();
        crossStaking.setPoolActive(poolId, false);
    }

    function testCannotSetActiveOnNonExistentPool() public {
        vm.expectRevert(CrossStaking.CSPoolNotFound.selector);
        crossStaking.setPoolActive(999, false);
    }

    // ==================== 풀 조회 ====================

    function testGetPoolInfo() public {
        (uint poolId, address poolAddress) = crossStaking.createPool(address(wcross), 1 ether);

        CrossStaking.PoolInfo memory info = crossStaking.getPoolInfo(poolId);

        assertEq(info.poolId, poolId, "Pool ID");
        assertEq(info.poolAddress, poolAddress, "Pool address");
        assertEq(info.stakingToken, address(wcross), "Staking token");
        assertTrue(info.active, "Active");
    }

    function testPoolAt() public {
        crossStaking.createPool(address(wcross), 1 ether);
        crossStaking.createPool(address(token1), 1 ether);
        crossStaking.createPool(address(token2), 1 ether);

        assertEq(crossStaking.poolAt(0), 1, "Pool at index 0");
        assertEq(crossStaking.poolAt(1), 2, "Pool at index 1");
        assertEq(crossStaking.poolAt(2), 3, "Pool at index 2");
    }

    function testGetPoolAddress() public {
        (uint poolId, address poolAddress) = crossStaking.createPool(address(wcross), 1 ether);

        assertEq(crossStaking.getPoolAddress(poolId), poolAddress, "Pool address lookup");
    }

    function testGetPoolId() public {
        (uint poolId, address poolAddress) = crossStaking.createPool(address(wcross), 1 ether);

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

        crossStaking.createPool(address(wcross), 1 ether);
        assertEq(crossStaking.getTotalPoolCount(), 1, "After 1 pool");

        crossStaking.createPool(address(token1), 1 ether);
        assertEq(crossStaking.getTotalPoolCount(), 2, "After 2 pools");
    }

    function testGetAllPoolIds() public {
        crossStaking.createPool(address(wcross), 1 ether);
        crossStaking.createPool(address(token1), 1 ether);
        crossStaking.createPool(address(token2), 1 ether);

        uint[] memory allIds = crossStaking.getAllPoolIds();
        assertEq(allIds.length, 3, "3 pools");
        assertEq(allIds[0], 1, "Pool 1");
        assertEq(allIds[1], 2, "Pool 2");
        assertEq(allIds[2], 3, "Pool 3");
    }

    function testGetPoolIdsByStakingToken() public {
        crossStaking.createPool(address(wcross), 1 ether);
        crossStaking.createPool(address(token1), 1 ether);
        crossStaking.createPool(address(wcross), 1 ether);

        uint[] memory wcrossIds = crossStaking.getPoolIdsByStakingToken(address(wcross));
        assertEq(wcrossIds.length, 2, "2 WCROSS pools");
        assertEq(wcrossIds[0], 1, "Pool 1");
        assertEq(wcrossIds[1], 3, "Pool 3");

        uint[] memory token1Ids = crossStaking.getPoolIdsByStakingToken(address(token1));
        assertEq(token1Ids.length, 1, "1 Token1 pool");
        assertEq(token1Ids[0], 2, "Pool 2");
    }

    function testPoolByStakingTokenAt() public {
        crossStaking.createPool(address(wcross), 1 ether);
        crossStaking.createPool(address(wcross), 1 ether);

        assertEq(crossStaking.poolByStakingTokenAt(address(wcross), 0), 1, "First WCROSS pool");
        assertEq(crossStaking.poolByStakingTokenAt(address(wcross), 1), 2, "Second WCROSS pool");
    }

    function testGetActivePoolIds() public {
        (uint pool1,) = crossStaking.createPool(address(wcross), 1 ether);
        (uint pool2,) = crossStaking.createPool(address(token1), 1 ether);
        (uint pool3,) = crossStaking.createPool(address(token2), 1 ether);

        // All active initially
        uint[] memory activeIds = crossStaking.getActivePoolIds();
        assertEq(activeIds.length, 3, "3 active pools");

        // Deactivate pool2
        crossStaking.setPoolActive(pool2, false);

        activeIds = crossStaking.getActivePoolIds();
        assertEq(activeIds.length, 2, "2 active pools");
        assertEq(activeIds[0], pool1, "Pool 1 still active");
        assertEq(activeIds[1], pool3, "Pool 3 still active");
    }

    // ==================== Router 관리 ====================

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

    // ==================== 관리자 함수 ====================

    function testSetPoolImplementation() public {
        address newImpl = address(new CrossStakingPool());

        crossStaking.setPoolImplementation(newImpl);
        assertEq(crossStaking.poolImplementation(), newImpl, "Implementation updated");
    }

    function testOnlyAdminCanSetPoolImplementation() public {
        address newImpl = address(new CrossStakingPool());

        vm.prank(user1);
        vm.expectRevert();
        crossStaking.setPoolImplementation(newImpl);
    }

    // ==================== 역할 관리 ====================

    function testGrantPoolManagerRole() public {
        assertFalse(crossStaking.hasRole(crossStaking.MANAGER_ROLE(), user1), "User1 not pool manager");

        // Grant role
        crossStaking.grantRole(crossStaking.MANAGER_ROLE(), user1);
        assertTrue(crossStaking.hasRole(crossStaking.MANAGER_ROLE(), user1), "User1 is pool manager");

        // User1 can create pool
        vm.prank(user1);
        (, address poolAddress) = crossStaking.createPool(address(token1), 1 ether);
        assertTrue(poolAddress != address(0), "Pool created by user1");
    }

    function testRevokePoolManagerRole() public {
        crossStaking.grantRole(crossStaking.MANAGER_ROLE(), user1);
        assertTrue(crossStaking.hasRole(crossStaking.MANAGER_ROLE(), user1), "User1 has role");

        crossStaking.revokeRole(crossStaking.MANAGER_ROLE(), user1);
        assertFalse(crossStaking.hasRole(crossStaking.MANAGER_ROLE(), user1), "Role revoked");

        vm.prank(user1);
        vm.expectRevert();
        crossStaking.createPool(address(token1), 1 ether);
    }

    // ==================== Integration ====================

    function testPoolsAreIndependent() public {
        // Create 2 pools
        (, address pool1Addr) = crossStaking.createPool(address(wcross), 1 ether);
        (, address pool2Addr) = crossStaking.createPool(address(token1), 1 ether);

        CrossStakingPool pool1 = CrossStakingPool(pool1Addr);
        CrossStakingPool pool2 = CrossStakingPool(pool2Addr);

        // Different staking tokens
        assertEq(address(pool1.stakingToken()), address(wcross), "Pool1 staking token");
        assertEq(address(pool2.stakingToken()), address(token1), "Pool2 staking token");

        // Different total staked
        token1.mint(user1, 100 ether);
        vm.startPrank(user1);
        token1.approve(pool2Addr, 50 ether);
        pool2.stake(50 ether);
        vm.stopPrank();

        assertEq(pool1.totalStaked(), 0, "Pool1 empty");
        assertEq(pool2.totalStaked(), 50 ether, "Pool2 has stake");
    }

    function testPoolActiveStatusAffectsPause() public {
        (uint poolId, address poolAddress) = crossStaking.createPool(address(wcross), 1 ether);
        CrossStakingPool pool = CrossStakingPool(poolAddress);

        // Deactivate pool
        crossStaking.setPoolActive(poolId, false);
        assertTrue(pool.paused(), "Pool should be paused");

        // Reactivate pool
        crossStaking.setPoolActive(poolId, true);
        assertFalse(pool.paused(), "Pool should be unpaused");
    }

    // ==================== UUPS 업그레이드 ====================

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
