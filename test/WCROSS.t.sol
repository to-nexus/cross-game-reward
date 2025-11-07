// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {CrossGameReward} from "../src/CrossGameReward.sol";
import {CrossGameRewardPool} from "../src/CrossGameRewardPool.sol";

import {CrossGameRewardRouter} from "../src/CrossGameRewardRouter.sol";
import {WCROSS} from "../src/WCROSS.sol";
import {ICrossGameRewardPool} from "../src/interfaces/ICrossGameRewardPool.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Test} from "forge-std/Test.sol";

import {IERC20Errors} from "@openzeppelin/contracts/interfaces/draft-IERC6093.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract WCROSSTest is Test {
    CrossGameReward public crossGameReward;
    CrossGameRewardPool public poolImplementation;
    CrossGameRewardRouter public router;
    WCROSS public wcross;

    address public owner;
    address public user1;
    address public user2;

    function setUp() public {
        owner = address(this);
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        // Give users native CROSS
        vm.deal(user1, 1000 ether);
        vm.deal(user2, 1000 ether);

        // Deploy system
        poolImplementation = new CrossGameRewardPool();

        CrossGameReward implementation = new CrossGameReward();
        bytes memory initData =
            abi.encodeCall(CrossGameReward.initialize, (ICrossGameRewardPool(address(poolImplementation)), owner, 2 days));
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), initData);
        crossGameReward = CrossGameReward(address(proxy));

        router = new CrossGameRewardRouter(address(crossGameReward));
        wcross = WCROSS(payable(address(crossGameReward.wcross())));
        crossGameReward.setRouter(address(router));
    }

    // ==================== Deposit (router only) ====================

    function testDepositViaRouter() public {
        vm.deal(address(router), 10 ether);

        vm.prank(address(router));
        wcross.deposit{value: 10 ether}();

        assertEq(wcross.balanceOf(address(router)), 10 ether, "WCROSS minted");
        assertEq(address(wcross).balance, 10 ether, "Contract balance");
    }

    function testReceiveViaRouter() public {
        vm.deal(address(router), 5 ether);

        vm.prank(address(router));
        (bool success,) = address(wcross).call{value: 5 ether}("");
        assertTrue(success, "Transfer succeeded");

        assertEq(wcross.balanceOf(address(router)), 5 ether, "WCROSS minted via receive");
    }

    function testCannotDepositByNonRouter() public {
        vm.prank(user1);
        vm.expectRevert(WCROSS.WCROSSUnauthorized.selector);
        wcross.deposit{value: 10 ether}();
    }

    function testCannotDepositZero() public {
        vm.prank(address(router));
        vm.expectRevert(WCROSS.WCROSSInsufficientBalance.selector);
        wcross.deposit{value: 0}();
    }

    // ==================== Withdraw (router only) ====================

    function testWithdrawToViaRouter() public {
        // Deposit first
        vm.deal(address(router), 10 ether);
        vm.prank(address(router));
        wcross.deposit{value: 10 ether}();

        // Withdraw directly to user (withdrawTo)
        uint balanceBefore = user1.balance;
        vm.prank(address(router));
        wcross.withdrawTo(user1, 10 ether);

        assertEq(wcross.balanceOf(address(router)), 0, "WCROSS burned from router");
        assertEq(user1.balance, balanceBefore + 10 ether, "Native CROSS sent to user");
    }

    function testCannotWithdrawByNonRouter() public {
        // Setup: transfer some WCROSS to user1
        vm.deal(address(router), 10 ether);
        vm.prank(address(router));
        wcross.deposit{value: 10 ether}();

        vm.prank(address(router));
        wcross.transfer(user1, 5 ether);

        // User1 tries to withdraw
        vm.prank(user1);
        vm.expectRevert(WCROSS.WCROSSUnauthorized.selector);
        wcross.withdraw(5 ether);
    }

    function testCannotWithdrawMoreThanBalance() public {
        vm.deal(address(router), 5 ether);
        vm.prank(address(router));
        wcross.deposit{value: 5 ether}();

        vm.prank(address(router));
        vm.expectRevert(
            abi.encodeWithSelector(IERC20Errors.ERC20InsufficientBalance.selector, address(router), 5 ether, 10 ether)
        );
        wcross.withdraw(10 ether);
    }

    // ==================== Transfer (ERC20 standard) ====================

    function testTransferBetweenUsers() public {
        // Router deposits
        vm.deal(address(router), 10 ether);
        vm.prank(address(router));
        wcross.deposit{value: 10 ether}();

        // Router transfers to user1
        vm.prank(address(router));
        wcross.transfer(user1, 4 ether);

        // User1 transfers to user2
        vm.prank(user1);
        wcross.transfer(user2, 2 ether);

        assertEq(wcross.balanceOf(address(router)), 6 ether, "Router balance");
        assertEq(wcross.balanceOf(user1), 2 ether, "User1 balance");
        assertEq(wcross.balanceOf(user2), 2 ether, "User2 balance");
    }

    // ==================== Integration with Router ====================

    function testDepositForIntegration() public {
        uint poolId;
        ICrossGameRewardPool poolAddress;

        // Create pool
        (poolId, poolAddress) = crossGameReward.createPool(IERC20(address(wcross)), 1 ether);

        // User deposits via router
        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(poolId);
        vm.stopPrank();

        // Verify WCROSS was minted and depositd
        CrossGameRewardPool pool = CrossGameRewardPool(address(poolAddress));
        assertEq(pool.balances(user1), 10 ether, "Deposited via router");
    }

    function testWithdrawForIntegration() public {
        // Setup: deposit first
        (uint poolId, ICrossGameRewardPool poolAddress) = crossGameReward.createPool(IERC20(address(wcross)), 1 ether);

        vm.startPrank(user1);
        router.depositNative{value: 10 ether}(poolId);
        vm.stopPrank();

        // Withdraw
        uint balanceBefore = user1.balance;
        vm.prank(user1);
        router.withdrawNative(poolId);

        // Verify native CROSS returned
        assertEq(user1.balance, balanceBefore + 10 ether, "Native CROSS returned");

        CrossGameRewardPool pool = CrossGameRewardPool(address(poolAddress));
        assertEq(pool.balances(user1), 0, "Withdrawn");
    }
}

// Receive function for testing
contract RouterMock {
    receive() external payable {}
}
