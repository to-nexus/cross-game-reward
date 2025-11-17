// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardPool.sol";
import "../src/CrossGameRewardRouter.sol";
import "../src/WCROSS.sol";
import "../src/interfaces/ICrossGameRewardPool.sol";
import "./mocks/MockERC20.sol";
import "./mocks/MockERC20Permit.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Test.sol";

contract CrossGameRewardRouterTest is Test {
    CrossGameReward public crossGameReward;
    CrossGameRewardPool public poolImplementation;
    CrossGameRewardRouter public router;
    WCROSS public wcross;

    MockERC20 public rewardToken;
    MockERC20 public depositToken;
    MockERC20Permit public permitToken; // EIP-2612 compatible token

    address public owner;
    address public user1;
    address public user2;
    uint public user1PrivateKey;

    uint public nativePoolId;
    ICrossGameRewardPool public nativePool;

    uint public erc20PoolId;
    ICrossGameRewardPool public erc20Pool;

    uint public permitPoolId;
    ICrossGameRewardPool public permitPool;

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
        poolImplementation = new CrossGameRewardPool();

        // Deploy CrossGameReward as a UUPS proxy (instantiates WCROSS)
        CrossGameReward implementation = new CrossGameReward();
        bytes memory initData = abi.encodeCall(
            CrossGameReward.initialize, (ICrossGameRewardPool(address(poolImplementation)), owner, 2 days)
        );
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), initData);
        crossGameReward = CrossGameReward(address(proxy));

        router = new CrossGameRewardRouter(address(crossGameReward));
        wcross = WCROSS(payable(address(crossGameReward.wcross())));

        // Setup router
        crossGameReward.setRouter(address(router));

        // Create test tokens
        rewardToken = new MockERC20("Reward", "RWD");
        depositToken = new MockERC20("Deposit", "STK");
        permitToken = new MockERC20Permit("Permit Token", "PTK");

        // Create pools
        (nativePoolId, nativePool) = crossGameReward.createPool(IERC20(address(wcross)), 1 ether);
        (erc20PoolId, erc20Pool) = crossGameReward.createPool(IERC20(address(depositToken)), 1 ether);
        (permitPoolId, permitPool) = crossGameReward.createPool(IERC20(address(permitToken)), 1 ether);

        // Add reward tokens
        crossGameReward.addRewardToken(nativePoolId, IERC20(address(rewardToken)));
        crossGameReward.addRewardToken(erc20PoolId, IERC20(address(rewardToken)));
        crossGameReward.addRewardToken(permitPoolId, IERC20(address(rewardToken)));

        // Mint deposit tokens for users
        depositToken.mint(user1, 1000 ether);
        depositToken.mint(user2, 1000 ether);
        permitToken.mint(user1, 1000 ether);
        permitToken.mint(user2, 1000 ether);
    }

    // ==================== Native CROSS deposit ====================

    function testDepositNative() public {
        uint amount = 10 ether;

        vm.startPrank(user1);

        router.depositNative{value: amount}(nativePoolId);
        vm.stopPrank();

        // Verify
        CrossGameRewardPool pool = CrossGameRewardPool(address(nativePool));
        assertEq(pool.balances(user1), amount, "User depositd");
        assertEq(pool.totalDeposited(), amount, "Total depositd");
    }

    function testDepositNativeMultipleTimes() public {
        vm.startPrank(user1);

        router.depositNative{value: 5 ether}(nativePoolId);
        router.depositNative{value: 3 ether}(nativePoolId);
        router.depositNative{value: 2 ether}(nativePoolId);
        vm.stopPrank();

        CrossGameRewardPool pool = CrossGameRewardPool(address(nativePool));
        assertEq(pool.balances(user1), 10 ether, "Total user deposit");
    }

    function testCannotDepositNativeZero() public {
        vm.prank(user1);
        vm.expectRevert(CrossGameRewardRouter.CSRInvalidAmount.selector);
        router.depositNative{value: 0}(nativePoolId);
    }

    function testCannotDepositNativeOnERC20Pool() public {
        vm.prank(user1);
        vm.expectRevert(CrossGameRewardRouter.CSRNotWCROSSPool.selector);
        router.depositNative{value: 10 ether}(erc20PoolId);
    }

    function testUndepositNative() public {
        // Deposit first
        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(nativePool), 100 ether);

        // Withdraw
        uint balanceBefore = user1.balance;
        vm.startPrank(user1);
        router.withdrawNative(nativePoolId);
        vm.stopPrank();

        // Verify
        CrossGameRewardPool pool = CrossGameRewardPool(address(nativePool));
        assertEq(pool.balances(user1), 0, "Withdrawn");
        assertEq(user1.balance, balanceBefore + 10 ether, "Native CROSS returned");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 100 ether, 10, "Rewards claimed");
    }

    function testCannotUndepositNativeWithoutDeposit() public {
        vm.prank(user1);
        vm.expectRevert(CrossGameRewardRouter.CSRNoDepositFound.selector);
        router.withdrawNative(nativePoolId);
    }

    // ==================== ERC20 deposit ====================

    function testDepositERC20() public {
        uint amount = 50 ether;

        vm.startPrank(user1);
        depositToken.approve(address(router), amount);
        router.depositERC20(erc20PoolId, amount);
        vm.stopPrank();

        CrossGameRewardPool pool = CrossGameRewardPool(address(erc20Pool));
        assertEq(pool.balances(user1), amount, "User depositd");
    }

    function testDepositERC20MultipleTimes() public {
        vm.startPrank(user1);
        depositToken.approve(address(router), 100 ether);

        router.depositERC20(erc20PoolId, 30 ether);
        router.depositERC20(erc20PoolId, 20 ether);
        router.depositERC20(erc20PoolId, 10 ether);
        vm.stopPrank();

        CrossGameRewardPool pool = CrossGameRewardPool(address(erc20Pool));
        assertEq(pool.balances(user1), 60 ether, "Total deposit");
    }

    function testUndepositERC20() public {
        // Deposit first
        vm.startPrank(user1);
        depositToken.approve(address(router), 50 ether);
        router.depositERC20(erc20PoolId, 50 ether);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(erc20Pool), 100 ether);

        // Withdraw
        uint balanceBefore = depositToken.balanceOf(user1);
        vm.prank(user1);
        router.withdrawERC20(erc20PoolId);

        CrossGameRewardPool pool = CrossGameRewardPool(address(erc20Pool));
        assertEq(pool.balances(user1), 0, "Withdrawn");
        assertEq(depositToken.balanceOf(user1), balanceBefore + 50 ether, "Tokens returned");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 100 ether, 10, "Rewards claimed");
    }

    function testCannotUndepositERC20WithoutDeposit() public {
        vm.prank(user1);
        vm.expectRevert(CrossGameRewardRouter.CSRNoDepositFound.selector);
        router.withdrawERC20(erc20PoolId);
    }

    function testCannotDepositERC20Zero() public {
        vm.prank(user1);
        vm.expectRevert(CrossGameRewardRouter.CSRInvalidAmount.selector);
        router.depositERC20(erc20PoolId, 0);
    }

    // ==================== View functions ====================

    function testGetUserDepositInfo() public {
        // Deposit
        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 50 ether);
        rewardToken.transfer(address(nativePool), 50 ether);

        (uint depositedAmount,, uint[] memory pendingRewards) = router.getUserDepositInfo(nativePoolId, user1);

        assertEq(depositedAmount, 10 ether, "Deposited amount");
        assertEq(pendingRewards.length, 1, "1 reward token");
        assertApproxEqAbs(pendingRewards[0], 50 ether, 10, "Pending reward");
    }

    function testIsNativePool() public view {
        assertTrue(router.isNativePool(nativePoolId), "Native pool");
        assertFalse(router.isNativePool(erc20PoolId), "Not native pool");
    }

    // ==================== Complex scenarios ====================

    function testMultiUserNativeDepositing() public {
        // User1 deposits
        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // User2 deposits
        vm.startPrank(user2);
        router.depositNative{value: 20 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 90 ether);
        rewardToken.transfer(address(nativePool), 90 ether);

        // User1 withdraws
        uint user1BalanceBefore = user1.balance;
        vm.prank(user1);
        router.withdrawNative(nativePoolId);

        assertEq(user1.balance, user1BalanceBefore + 10 ether, "User1 got native CROSS");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 30 ether, 10, "User1 rewards (1/3)");

        // User2 withdraws
        uint user2BalanceBefore = user2.balance;
        vm.prank(user2);
        router.withdrawNative(nativePoolId);

        assertEq(user2.balance, user2BalanceBefore + 20 ether, "User2 got native CROSS");
        assertApproxEqAbs(rewardToken.balanceOf(user2), 60 ether, 10, "User2 rewards (2/3)");
    }

    function testMixedPoolUsage() public {
        // Native pool depositing
        vm.startPrank(user1);
        router.depositNative{value: 5 ether}(nativePoolId);
        vm.stopPrank();

        // ERC20 pool depositing
        vm.startPrank(user2);
        depositToken.approve(address(router), 50 ether);
        router.depositERC20(erc20PoolId, 50 ether);
        vm.stopPrank();

        // Verify both pools
        CrossGameRewardPool nativePoolContract = CrossGameRewardPool(address(nativePool));
        CrossGameRewardPool erc20PoolContract = CrossGameRewardPool(address(erc20Pool));

        assertEq(nativePoolContract.balances(user1), 5 ether, "Native pool deposit");
        assertEq(erc20PoolContract.balances(user2), 50 ether, "ERC20 pool deposit");
    }

    // ==================== ERC20 permit deposit ====================

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

    function testDepositERC20WithPermit() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1 hours;

        // Generate permit signature
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);

        // Deposit with permit (no prior approval needed!)
        vm.prank(user1);
        router.depositERC20WithPermit(permitPoolId, amount, deadline, v, r, s);

        // Verify
        CrossGameRewardPool pool = CrossGameRewardPool(address(permitPool));
        assertEq(pool.balances(user1), amount, "User depositd with permit");
        assertEq(pool.totalDeposited(), amount, "Total depositd");
        assertEq(permitToken.balanceOf(user1), 950 ether, "Tokens deducted");
    }

    function testDepositERC20WithPermitMultipleTimes() public {
        uint deadline = block.timestamp + 1 hours;

        // First deposit
        (uint8 v1, bytes32 r1, bytes32 s1) =
            _getPermitSignature(address(permitToken), user1, address(router), 20 ether, deadline, user1PrivateKey);
        vm.prank(user1);
        router.depositERC20WithPermit(permitPoolId, 20 ether, deadline, v1, r1, s1);

        // Second deposit (nonce increased)
        (uint8 v2, bytes32 r2, bytes32 s2) =
            _getPermitSignature(address(permitToken), user1, address(router), 30 ether, deadline, user1PrivateKey);
        vm.prank(user1);
        router.depositERC20WithPermit(permitPoolId, 30 ether, deadline, v2, r2, s2);

        // Verify
        CrossGameRewardPool pool = CrossGameRewardPool(address(permitPool));
        assertEq(pool.balances(user1), 50 ether, "Total depositd with permit");
    }

    function testCannotDepositERC20WithPermitZeroAmount() public {
        uint deadline = block.timestamp + 1 hours;

        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), 0, deadline, user1PrivateKey);

        vm.prank(user1);
        vm.expectRevert(CrossGameRewardRouter.CSRInvalidAmount.selector);
        router.depositERC20WithPermit(permitPoolId, 0, deadline, v, r, s);
    }

    function testCannotDepositERC20WithPermitInvalidSignature() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1 hours;

        // Generate signature with wrong private key
        uint wrongPrivateKey = 0xBAD;
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, wrongPrivateKey);

        vm.prank(user1);
        vm.expectRevert(); // ERC20Permit will revert with invalid signature
        router.depositERC20WithPermit(permitPoolId, amount, deadline, v, r, s);
    }

    function testCannotDepositERC20WithPermitExpiredDeadline() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp - 1; // Already expired

        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);

        vm.prank(user1);
        vm.expectRevert(); // ERC20Permit will revert with expired deadline
        router.depositERC20WithPermit(permitPoolId, amount, deadline, v, r, s);
    }

    function testDepositERC20WithPermitAndWithdraw() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1 hours;

        // Deposit with permit
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);
        vm.prank(user1);
        router.depositERC20WithPermit(permitPoolId, amount, deadline, v, r, s);

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(permitPool), 100 ether);

        // Withdraw
        uint balanceBefore = permitToken.balanceOf(user1);
        vm.prank(user1);
        router.withdrawERC20(permitPoolId);

        // Verify
        CrossGameRewardPool pool = CrossGameRewardPool(address(permitPool));
        assertEq(pool.balances(user1), 0, "Withdrawn");
        assertEq(permitToken.balanceOf(user1), balanceBefore + amount, "Tokens returned");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 100 ether, 10, "Rewards claimed");
    }

    function testMultiUserDepositWithPermit() public {
        uint deadline = block.timestamp + 1 hours;

        // User1 deposits with permit
        (uint8 v1, bytes32 r1, bytes32 s1) =
            _getPermitSignature(address(permitToken), user1, address(router), 30 ether, deadline, user1PrivateKey);
        vm.prank(user1);
        router.depositERC20WithPermit(permitPoolId, 30 ether, deadline, v1, r1, s1);

        // User2 deposits normally (without permit)
        vm.startPrank(user2);
        permitToken.approve(address(router), 70 ether);
        router.depositERC20(permitPoolId, 70 ether);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(permitPool), 100 ether);

        // Both users withdraw
        vm.prank(user1);
        router.withdrawERC20(permitPoolId);

        vm.prank(user2);
        router.withdrawERC20(permitPoolId);

        // Verify rewards distribution (30:70 ratio)
        assertApproxEqAbs(rewardToken.balanceOf(user1), 30 ether, 10, "User1 rewards (30%)");
        assertApproxEqAbs(rewardToken.balanceOf(user2), 70 ether, 10, "User2 rewards (70%)");
    }

    /// @notice Tests permit with amount mismatch (signature for different amount)
    function testCannotDepositERC20WithPermitAmountMismatch() public {
        uint signedAmount = 50 ether;
        uint actualAmount = 30 ether; // Different from signed amount
        uint deadline = block.timestamp + 1 hours;

        // Generate signature for 50 ether
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), signedAmount, deadline, user1PrivateKey);

        // Try to deposit 30 ether with signature for 50 ether
        vm.prank(user1);
        vm.expectRevert(); // Will fail because allowance is 50 but trying to transferFrom 30
        router.depositERC20WithPermit(permitPoolId, actualAmount, deadline, v, r, s);
    }

    /// @notice Tests permit with non-permit supporting token (should revert)
    function testCannotDepositERC20WithPermitOnNonPermitToken() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1 hours;

        // Generate a valid signature (doesn't matter, depositToken doesn't support permit)
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);

        // Try to call permit on a non-permit token (depositToken = regular MockERC20)
        vm.prank(user1);
        vm.expectRevert(); // Should revert because depositToken doesn't have permit function
        router.depositERC20WithPermit(erc20PoolId, amount, deadline, v, r, s);
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
        router.depositERC20WithPermit(permitPoolId, amount, deadline, v, r, s);

        CrossGameRewardPool pool = CrossGameRewardPool(address(permitPool));
        assertEq(pool.balances(user1), 30 ether, "First deposit succeeded");

        // Try to reuse same signature - should fail (nonce increased)
        vm.prank(user1);
        vm.expectRevert(); // Nonce mismatch
        router.depositERC20WithPermit(permitPoolId, amount, deadline, v, r, s);
    }

    /// @notice Tests permit with wrong spender address
    function testCannotDepositERC20WithPermitWrongSpender() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1 hours;
        address wrongSpender = address(0xBADBEEF);

        // Generate signature for wrong spender
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, wrongSpender, amount, deadline, user1PrivateKey);

        // Try to deposit with router (signature was for different spender)
        vm.prank(user1);
        vm.expectRevert(); // transferFrom will fail - no allowance
        router.depositERC20WithPermit(permitPoolId, amount, deadline, v, r, s);
    }

    /// @notice Tests permit with insufficient balance
    function testCannotDepositERC20WithPermitInsufficientBalance() public {
        uint amount = 2000 ether; // More than user1's balance (1000 ether)
        uint deadline = block.timestamp + 1 hours;

        // Generate valid signature
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);

        // Try to deposit more than balance
        vm.prank(user1);
        vm.expectRevert(); // ERC20: transfer amount exceeds balance
        router.depositERC20WithPermit(permitPoolId, amount, deadline, v, r, s);
    }

    /// @notice Tests permit with very short deadline (edge case)
    function testDepositERC20WithPermitVeryShortDeadline() public {
        uint amount = 50 ether;
        uint deadline = block.timestamp + 1; // 1 second deadline

        // Generate signature
        (uint8 v, bytes32 r, bytes32 s) =
            _getPermitSignature(address(permitToken), user1, address(router), amount, deadline, user1PrivateKey);

        // Should succeed if executed immediately
        vm.prank(user1);
        router.depositERC20WithPermit(permitPoolId, amount, deadline, v, r, s);

        CrossGameRewardPool pool = CrossGameRewardPool(address(permitPool));
        assertEq(pool.balances(user1), amount, "Deposited with short deadline");
    }

    // ==================== Claim Rewards ====================

    function testClaimRewardsNativePool() public {
        // Deposit
        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(nativePool), 100 ether);

        // Claim rewards (without withdrawing deposit)
        vm.prank(user1);
        router.claimRewards(nativePoolId);

        // Verify rewards claimed but deposit remains
        CrossGameRewardPool pool = CrossGameRewardPool(address(nativePool));
        assertEq(pool.balances(user1), 10 ether, "Deposit still in pool");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 100 ether, 10, "Rewards claimed");
    }

    function testClaimRewardsERC20Pool() public {
        // Deposit
        vm.startPrank(user1);
        depositToken.approve(address(router), 50 ether);
        router.depositERC20(erc20PoolId, 50 ether);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 200 ether);
        rewardToken.transfer(address(erc20Pool), 200 ether);

        // Claim rewards
        vm.prank(user1);
        router.claimRewards(erc20PoolId);

        // Verify
        CrossGameRewardPool pool = CrossGameRewardPool(address(erc20Pool));
        assertEq(pool.balances(user1), 50 ether, "Deposit still in pool");
        assertApproxEqAbs(rewardToken.balanceOf(user1), 200 ether, 10, "Rewards claimed");
    }

    function testClaimSpecificRewardToken() public {
        // Create second reward token
        MockERC20 rewardToken2 = new MockERC20("Reward2", "RWD2");
        crossGameReward.addRewardToken(nativePoolId, IERC20(address(rewardToken2)));

        // Deposit
        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add both reward tokens
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(nativePool), 100 ether);

        rewardToken2.mint(owner, 50 ether);
        rewardToken2.transfer(address(nativePool), 50 ether);

        // Claim only first reward token
        vm.prank(user1);
        router.claimReward(nativePoolId, address(rewardToken));

        // Verify only first reward claimed
        assertApproxEqAbs(rewardToken.balanceOf(user1), 100 ether, 10, "First reward claimed");
        assertEq(rewardToken2.balanceOf(user1), 0, "Second reward not claimed yet");

        // Claim second reward token
        vm.prank(user1);
        router.claimReward(nativePoolId, address(rewardToken2));

        assertApproxEqAbs(rewardToken2.balanceOf(user1), 50 ether, 10, "Second reward claimed");
    }

    function testMultipleClaimsSamePool() public {
        // Deposit
        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add first batch of rewards
        rewardToken.mint(owner, 50 ether);
        rewardToken.transfer(address(nativePool), 50 ether);

        // First claim
        vm.prank(user1);
        router.claimRewards(nativePoolId);
        assertApproxEqAbs(rewardToken.balanceOf(user1), 50 ether, 10, "First claim");

        // Add second batch of rewards
        rewardToken.mint(owner, 30 ether);
        rewardToken.transfer(address(nativePool), 30 ether);

        // Second claim
        vm.prank(user1);
        router.claimRewards(nativePoolId);
        assertApproxEqAbs(rewardToken.balanceOf(user1), 80 ether, 10, "Total after second claim");
    }

    function testClaimRewardsMultipleUsers() public {
        // User1 deposits 30 ether
        vm.startPrank(user1);
        router.depositNative{value: 30 ether}(nativePoolId);
        vm.stopPrank();

        // User2 deposits 70 ether
        vm.startPrank(user2);
        router.depositNative{value: 70 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards (should be distributed 30:70)
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(nativePool), 100 ether);

        // User1 claims
        vm.prank(user1);
        router.claimRewards(nativePoolId);
        assertApproxEqAbs(rewardToken.balanceOf(user1), 30 ether, 10, "User1 rewards (30%)");

        // User2 claims
        vm.prank(user2);
        router.claimRewards(nativePoolId);
        assertApproxEqAbs(rewardToken.balanceOf(user2), 70 ether, 10, "User2 rewards (70%)");

        // Verify deposits unchanged
        CrossGameRewardPool pool = CrossGameRewardPool(address(nativePool));
        assertEq(pool.balances(user1), 30 ether, "User1 deposit unchanged");
        assertEq(pool.balances(user2), 70 ether, "User2 deposit unchanged");
    }

    function testClaimWithoutDeposit() public {
        // User has no deposit
        vm.prank(user1);
        vm.expectRevert(); // Should revert - no deposit or stored rewards
        router.claimRewards(nativePoolId);
    }

    function testClaimAfterPartialWithdraw() public {
        // User1 deposits
        vm.startPrank(user1);
        router.depositNative{value: 20 ether}(nativePoolId);
        vm.stopPrank();

        // User2 deposits
        vm.startPrank(user2);
        router.depositNative{value: 80 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(nativePool), 100 ether);

        // User2 withdraws (gets rewards + deposit back)
        vm.prank(user2);
        router.withdrawNative(nativePoolId);
        assertApproxEqAbs(rewardToken.balanceOf(user2), 80 ether, 10, "User2 rewards from withdraw");

        // Add more rewards (all should go to user1 now)
        rewardToken.mint(owner, 50 ether);
        rewardToken.transfer(address(nativePool), 50 ether);

        // User1 claims
        vm.prank(user1);
        router.claimRewards(nativePoolId);

        // User1 should have: 20 ether from first batch + 50 ether from second batch
        assertApproxEqAbs(rewardToken.balanceOf(user1), 70 ether, 10, "User1 total rewards");
    }

    function testGetPendingRewards() public {
        // Deposit
        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(nativePool), 100 ether);

        // Check pending rewards via router
        (address[] memory tokens, uint[] memory amounts) = router.getPendingRewards(nativePoolId, user1);

        assertEq(tokens.length, 1, "One reward token");
        assertEq(tokens[0], address(rewardToken), "Correct reward token");
        assertApproxEqAbs(amounts[0], 100 ether, 10, "Correct pending amount");
    }

    function testGetPendingReward() public {
        // Deposit
        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add rewards
        rewardToken.mint(owner, 100 ether);
        rewardToken.transfer(address(nativePool), 100 ether);

        // Check pending reward for specific token
        uint amount = router.getPendingReward(nativePoolId, user1, address(rewardToken));

        assertApproxEqAbs(amount, 100 ether, 10, "Correct pending amount");
    }

    function testGetRemovedTokenRewardsViaRouter() public {
        // Setup deposit
        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Add and remove reward token with pending rewards
        MockERC20 removedToken = new MockERC20("Removed", "REM");
        crossGameReward.addRewardToken(nativePoolId, IERC20(address(removedToken)));

        uint rewardAmount = 25 ether;
        removedToken.mint(owner, rewardAmount);
        removedToken.transfer(address(nativePool), rewardAmount);

        crossGameReward.removeRewardToken(nativePoolId, IERC20(address(removedToken)));

        (address[] memory tokens, uint[] memory amounts) = router.getRemovedTokenRewards(nativePoolId, user1);

        assertEq(tokens.length, 1, "One removed token");
        assertEq(tokens[0], address(removedToken), "Removed token address");
        assertApproxEqAbs(amounts[0], rewardAmount, 10, "Pending amount for removed token");
    }

    function testGetAllPendingRewardsIncludesRemovedTokens() public {
        // Deposit
        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(nativePoolId);
        vm.stopPrank();

        // Active reward
        rewardToken.mint(owner, 40 ether);
        rewardToken.transfer(address(nativePool), 40 ether);

        // Removed reward
        MockERC20 removedToken = new MockERC20("Removed", "REM");
        crossGameReward.addRewardToken(nativePoolId, IERC20(address(removedToken)));
        uint removedRewardAmount = 60 ether;
        removedToken.mint(owner, removedRewardAmount);
        removedToken.transfer(address(nativePool), removedRewardAmount);
        crossGameReward.removeRewardToken(nativePoolId, IERC20(address(removedToken)));

        (address[] memory tokens, uint[] memory amounts) = router.getAllPendingRewards(nativePoolId, user1);

        assertEq(tokens.length, 2, "Active + removed tokens reported");

        // Active token should be first, removed second (order defined by helper)
        assertEq(tokens[0], address(rewardToken), "Active token first");
        assertApproxEqAbs(amounts[0], 40 ether, 10, "Active reward amount");
        assertEq(tokens[1], address(removedToken), "Removed token second");
        assertApproxEqAbs(amounts[1], removedRewardAmount, 10, "Removed reward amount");

        vm.prank(user1);
        router.claimRewards(nativePoolId);

        assertApproxEqAbs(rewardToken.balanceOf(user1), 40 ether, 10, "Active token claimed");
        assertApproxEqAbs(removedToken.balanceOf(user1), removedRewardAmount, 10, "Removed token claimed");
    }
}
