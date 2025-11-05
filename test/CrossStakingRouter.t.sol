// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/CrossStaking.sol";
import "../src/CrossStakingPool.sol";
import "../src/CrossStakingRouter.sol";
import "../src/WCROSS.sol";
import "../src/interfaces/ICrossStakingPool.sol";
import "./mocks/MockERC20.sol";
import "./mocks/MockERC20Permit.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Test.sol";

contract CrossStakingRouterTest is Test {
    CrossStaking public crossStaking;
    CrossStakingPool public poolImplementation;
    CrossStakingRouter public router;
    WCROSS public wcross;

    MockERC20 public rewardToken;
    MockERC20 public stakingToken;
    MockERC20Permit public permitToken; // EIP-2612 compatible token

    address public owner;
    address public user1;
    address public user2;
    uint public user1PrivateKey;

    uint public nativePoolId;
    ICrossStakingPool public nativePool;

    uint public erc20PoolId;
    ICrossStakingPool public erc20Pool;

    uint public permitPoolId;
    ICrossStakingPool public permitPool;

    function setUp() public {
        owner = address(this);

        // Create user1 with private key (for permit signing)
        user1PrivateKey = 0xA11CE;
        user1 = vm.addr(user1PrivateKey);

        user2 = makeAddr("user2");

        // Give users native CROSS
        vm.deal(user1, 1000 ether);
        vm.deal(user2, 1000 ether);

        // Deploy core contracts
        poolImplementation = new CrossStakingPool();

        // Deploy CrossStaking as a UUPS proxy (instantiates WCROSS)
        CrossStaking implementation = new CrossStaking();
        bytes memory initData =
            abi.encodeCall(CrossStaking.initialize, (ICrossStakingPool(address(poolImplementation)), owner, 2 days));
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), initData);
        crossStaking = CrossStaking(address(proxy));

        router = new CrossStakingRouter(address(crossStaking));
        wcross = WCROSS(payable(address(crossStaking.wcross())));

        // Setup router
        crossStaking.setRouter(address(router));

        // Create test tokens
        rewardToken = new MockERC20("Reward", "RWD");
        stakingToken = new MockERC20("Staking", "STK");
        permitToken = new MockERC20Permit("Permit Token", "PTK");

        // Create pools
        (nativePoolId, nativePool) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);
        (erc20PoolId, erc20Pool) = crossStaking.createPool(IERC20(address(stakingToken)), 1 ether);
        (permitPoolId, permitPool) = crossStaking.createPool(IERC20(address(permitToken)), 1 ether);

        // Add reward tokens
        crossStaking.addRewardToken(nativePoolId, IERC20(address(rewardToken)));
        crossStaking.addRewardToken(erc20PoolId, IERC20(address(rewardToken)));
        crossStaking.addRewardToken(permitPoolId, IERC20(address(rewardToken)));

        // Mint staking tokens for users
        stakingToken.mint(user1, 1000 ether);
        stakingToken.mint(user2, 1000 ether);
        permitToken.mint(user1, 1000 ether);
        permitToken.mint(user2, 1000 ether);
    }

    // ==================== Native CROSS staking ====================

    function testStakeNative() public {
        uint amount = 10 ether;

        vm.startPrank(user1);

        router.stakeNative{value: amount}(nativePoolId);
        vm.stopPrank();

        // Verify
        CrossStakingPool pool = CrossStakingPool(address(nativePool));
        assertEq(pool.balances(user1), amount, "User staked");
        assertEq(pool.totalStaked(), amount, "Total staked");
    }

    function testStakeNativeMultipleTimes() public {
        vm.startPrank(user1);

        router.stakeNative{value: 5 ether}(nativePoolId);
        router.stakeNative{value: 3 ether}(nativePoolId);
        router.stakeNative{value: 2 ether}(nativePoolId);
        vm.stopPrank();

        CrossStakingPool pool = CrossStakingPool(address(nativePool));
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
        router.stakeNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(nativePool), 100 ether);

        // Unstake
        uint balanceBefore = user1.balance;
        vm.startPrank(user1);
        router.unstakeNative(nativePoolId);
        vm.stopPrank();

        // Verify
        CrossStakingPool pool = CrossStakingPool(address(nativePool));
        assertEq(pool.balances(user1), 0, "Unstaked");
        assertEq(user1.balance, balanceBefore + 10 ether, "Native CROSS returned");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 100 ether, 10, "Rewards claimed");
    }

    function testCannotUnstakeNativeWithoutStake() public {
        vm.prank(user1);
        vm.expectRevert(CrossStakingRouter.CSRNoStakeFound.selector);
        router.unstakeNative(nativePoolId);
    }

    // ==================== ERC20 staking ====================

    function testStakeERC20() public {
        uint amount = 50 ether;

        vm.startPrank(user1);
        stakingToken.approve(address(router), amount);
        router.stakeERC20(erc20PoolId, amount);
        vm.stopPrank();

        CrossStakingPool pool = CrossStakingPool(address(erc20Pool));
        assertEq(pool.balances(user1), amount, "User staked");
    }

    function testStakeERC20MultipleTimes() public {
        vm.startPrank(user1);
        stakingToken.approve(address(router), 100 ether);

        router.stakeERC20(erc20PoolId, 30 ether);
        router.stakeERC20(erc20PoolId, 20 ether);
        router.stakeERC20(erc20PoolId, 10 ether);
        vm.stopPrank();

        CrossStakingPool pool = CrossStakingPool(address(erc20Pool));
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
        rewardToken.transfer(address(erc20Pool), 100 ether);

        // Unstake
        uint balanceBefore = stakingToken.balanceOf(user1);
        vm.prank(user1);
        router.unstakeERC20(erc20PoolId);

        CrossStakingPool pool = CrossStakingPool(address(erc20Pool));
        assertEq(pool.balances(user1), 0, "Unstaked");
        assertEq(stakingToken.balanceOf(user1), balanceBefore + 50 ether, "Tokens returned");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 100 ether, 10, "Rewards claimed");
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

    // ==================== View functions ====================

    function testGetUserStakingInfo() public {
        // Stake
        vm.startPrank(user1);
        router.stakeNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 50 ether);
        rewardToken.transfer(address(nativePool), 50 ether);

        (uint stakedAmount,, uint[] memory pendingRewards) = router.getUserStakingInfo(nativePoolId, user1);

        assertEq(stakedAmount, 10 ether, "Staked amount");
        assertEq(pendingRewards.length, 1, "1 reward token");
        assertApproxEqAbs(pendingRewards[0], 50 ether, 10, "Pending reward");
    }

    function testIsNativePool() public view {
        assertTrue(router.isNativePool(nativePoolId), "Native pool");
        assertFalse(router.isNativePool(erc20PoolId), "Not native pool");
    }

    // ==================== Complex scenarios ====================

    function testMultiUserNativeStaking() public {
        // User1 stakes
        vm.startPrank(user1);
        router.stakeNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // User2 stakes
        vm.startPrank(user2);
        router.stakeNative{value: 20 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 90 ether);
        rewardToken.transfer(address(nativePool), 90 ether);

        // User1 unstakes
        uint user1BalanceBefore = user1.balance;
        vm.prank(user1);
        router.unstakeNative(nativePoolId);

        assertEq(user1.balance, user1BalanceBefore + 10 ether, "User1 got native CROSS");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 30 ether, 10, "User1 rewards (1/3)");

        // User2 unstakes
        uint user2BalanceBefore = user2.balance;
        vm.prank(user2);
        router.unstakeNative(nativePoolId);

        assertEq(user2.balance, user2BalanceBefore + 20 ether, "User2 got native CROSS");
        assertApproxEqAbs(rewardToken.balanceOf(user2), 60 ether, 10, "User2 rewards (2/3)");
    }

    function testMixedPoolUsage() public {
        // Native pool staking
        vm.startPrank(user1);
        router.stakeNative{value: 5 ether}(nativePoolId);
        vm.stopPrank();

        // ERC20 pool staking
        vm.startPrank(user2);
        stakingToken.approve(address(router), 50 ether);
        router.stakeERC20(erc20PoolId, 50 ether);
        vm.stopPrank();

        // Verify both pools
        CrossStakingPool nativePoolContract = CrossStakingPool(address(nativePool));
        CrossStakingPool erc20PoolContract = CrossStakingPool(address(erc20Pool));

        assertEq(nativePoolContract.balances(user1), 5 ether, "Native pool stake");
        assertEq(erc20PoolContract.balances(user2), 50 ether, "ERC20 pool stake");
    }

    // ==================== ERC20 permit staking ====================

    /// @notice Helper function to generate EIP-2612 permit signature
    function _getPermitSignature(
        address token,
        address tokenOwner,
        address spender,
        uint value,
        uint deadline,
        uint privateKey
    ) internal view returns (uint8 v, bytes32 r, bytes32 s) {
        bytes32 PERMIT_TYPEHASH =
            keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

        bytes32 domainSeparator = MockERC20Permit(token).DOMAIN_SEPARATOR();
        uint nonce = MockERC20Permit(token).nonces(tokenOwner);

        bytes32 structHash = keccak256(abi.encode(PERMIT_TYPEHASH, tokenOwner, spender, value, nonce, deadline));

        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));

        (v, r, s) = vm.sign(privateKey, digest);
    }

    function testStakeERC20WithPermit() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1 hours;

        // Generate permit signature
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);

        // Stake with permit (no prior approval needed!)
        vm.prank(user1);
        router.stakeERC20WithPermit(permitPoolId, amount, deadline, v, r, s);

        // Verify
        CrossStakingPool pool = CrossStakingPool(address(permitPool));
        assertEq(pool.balances(user1), amount, "User staked with permit");
        assertEq(pool.totalStaked(), amount, "Total staked");
        assertEq(permitToken.balanceOf(user1), 950 ether, "Tokens deducted");
    }

    function testStakeERC20WithPermitMultipleTimes() public {
        uint deadline = block.timestamp + 1 hours;

        // First stake
        (uint8 v1, bytes32 r1, bytes32 s1) =
            _getPermitSignature(address(permitToken), user1, address(router), 20 ether, deadline, user1PrivateKey);
        vm.prank(user1);
        router.stakeERC20WithPermit(permitPoolId, 20 ether, deadline, v1, r1, s1);

        // Second stake (nonce increased)
        (uint8 v2, bytes32 r2, bytes32 s2) =
            _getPermitSignature(address(permitToken), user1, address(router), 30 ether, deadline, user1PrivateKey);
        vm.prank(user1);
        router.stakeERC20WithPermit(permitPoolId, 30 ether, deadline, v2, r2, s2);

        // Verify
        CrossStakingPool pool = CrossStakingPool(address(permitPool));
        assertEq(pool.balances(user1), 50 ether, "Total staked with permit");
    }

    function testCannotStakeERC20WithPermitZeroAmount() public {
        uint deadline = block.timestamp + 1 hours;

        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), 0, deadline, user1PrivateKey);

        vm.prank(user1);
        vm.expectRevert(CrossStakingRouter.CSRInvalidAmount.selector);
        router.stakeERC20WithPermit(permitPoolId, 0, deadline, v, r, s);
    }

    function testCannotStakeERC20WithPermitInvalidSignature() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1 hours;

        // Generate signature with wrong private key
        uint wrongPrivateKey = 0xBAD;
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, wrongPrivateKey);

        vm.prank(user1);
        vm.expectRevert(); // ERC20Permit will revert with invalid signature
        router.stakeERC20WithPermit(permitPoolId, amount, deadline, v, r, s);
    }

    function testCannotStakeERC20WithPermitExpiredDeadline() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp - 1; // Already expired

        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);

        vm.prank(user1);
        vm.expectRevert(); // ERC20Permit will revert with expired deadline
        router.stakeERC20WithPermit(permitPoolId, amount, deadline, v, r, s);
    }

    function testStakeERC20WithPermitAndUnstake() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1 hours;

        // Stake with permit
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);
        vm.prank(user1);
        router.stakeERC20WithPermit(permitPoolId, amount, deadline, v, r, s);

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(permitPool), 100 ether);

        // Unstake
        uint balanceBefore = permitToken.balanceOf(user1);
        vm.prank(user1);
        router.unstakeERC20(permitPoolId);

        // Verify
        CrossStakingPool pool = CrossStakingPool(address(permitPool));
        assertEq(pool.balances(user1), 0, "Unstaked");
        assertEq(permitToken.balanceOf(user1), balanceBefore + amount, "Tokens returned");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 100 ether, 10, "Rewards claimed");
    }

    function testMultiUserStakeWithPermit() public {
        uint deadline = block.timestamp + 1 hours;

        // User1 stakes with permit
        (uint8 v1, bytes32 r1, bytes32 s1) =
            _getPermitSignature(address(permitToken), user1, address(router), 30 ether, deadline, user1PrivateKey);
        vm.prank(user1);
        router.stakeERC20WithPermit(permitPoolId, 30 ether, deadline, v1, r1, s1);

        // User2 stakes normally (without permit)
        vm.startPrank(user2);
        permitToken.approve(address(router), 70 ether);
        router.stakeERC20(permitPoolId, 70 ether);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(permitPool), 100 ether);

        // Both users unstake
        vm.prank(user1);
        router.unstakeERC20(permitPoolId);

        vm.prank(user2);
        router.unstakeERC20(permitPoolId);

        // Verify rewards distribution (30:70 ratio)
        assertApproxEqAbs(rewardToken.balanceOf(user1), 30 ether, 10, "User1 rewards (30%)");
        assertApproxEqAbs(rewardToken.balanceOf(user2), 70 ether, 10, "User2 rewards (70%)");
    }

    /// @notice Tests permit with amount mismatch (signature for different amount)
    function testCannotStakeERC20WithPermitAmountMismatch() public {
        uint signedAmount = 50 ether;
        uint actualAmount = 30 ether; // Different from signed amount
        uint deadline = block.timestamp + 1 hours;

        // Generate signature for 50 ether
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), signedAmount, deadline, user1PrivateKey);

        // Try to stake 30 ether with signature for 50 ether
        vm.prank(user1);
        vm.expectRevert(); // Will fail because allowance is 50 but trying to transferFrom 30
        router.stakeERC20WithPermit(permitPoolId, actualAmount, deadline, v, r, s);
    }

    /// @notice Tests permit with non-permit supporting token (should revert)
    function testCannotStakeERC20WithPermitOnNonPermitToken() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1 hours;

        // Generate a valid signature (doesn't matter, stakingToken doesn't support permit)
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);

        // Try to call permit on a non-permit token (stakingToken = regular MockERC20)
        vm.prank(user1);
        vm.expectRevert(); // Should revert because stakingToken doesn't have permit function
        router.stakeERC20WithPermit(erc20PoolId, amount, deadline, v, r, s);
    }

    /// @notice Tests permit signature reuse protection (nonce increases after use)
    function testCannotReusePermitSignature() public {
        uint amount = 30 ether;
        uint deadline = block.timestamp + 1 hours;

        // Generate signature
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);

        // First use - should succeed
        vm.prank(user1);
        router.stakeERC20WithPermit(permitPoolId, amount, deadline, v, r, s);

        CrossStakingPool pool = CrossStakingPool(address(permitPool));
        assertEq(pool.balances(user1), 30 ether, "First stake succeeded");

        // Try to reuse same signature - should fail (nonce increased)
        vm.prank(user1);
        vm.expectRevert(); // Nonce mismatch
        router.stakeERC20WithPermit(permitPoolId, amount, deadline, v, r, s);
    }

    /// @notice Tests permit with wrong spender address
    function testCannotStakeERC20WithPermitWrongSpender() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1 hours;
        address wrongSpender = address(0xBADBEEF);

        // Generate signature for wrong spender
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, wrongSpender, amount, deadline, user1PrivateKey);

        // Try to stake with router (signature was for different spender)
        vm.prank(user1);
        vm.expectRevert(); // transferFrom will fail - no allowance
        router.stakeERC20WithPermit(permitPoolId, amount, deadline, v, r, s);
    }

    /// @notice Tests permit with insufficient balance
    function testCannotStakeERC20WithPermitInsufficientBalance() public {
        uint amount = 2000 ether; // More than user1's balance (1000 ether)
        uint deadline = block.timestamp + 1 hours;

        // Generate valid signature
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);

        // Try to stake more than balance
        vm.prank(user1);
        vm.expectRevert(); // ERC20: transfer amount exceeds balance
        router.stakeERC20WithPermit(permitPoolId, amount, deadline, v, r, s);
    }

    /// @notice Tests permit with very short deadline (edge case)
    function testStakeERC20WithPermitVeryShortDeadline() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1; // 1 second deadline

        // Generate signature
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);

        // Should succeed if executed immediately
        vm.prank(user1);
        router.stakeERC20WithPermit(permitPoolId, amount, deadline, v, r, s);

        CrossStakingPool pool = CrossStakingPool(address(permitPool));
        assertEq(pool.balances(user1), amount, "Staked with short deadline");
    }
}
