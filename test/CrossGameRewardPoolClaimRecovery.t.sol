// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/CrossGameRewardPoolBase.t.sol";

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title MockFailingERC20
 * @notice ERC20 token that can simulate transfer failures
 * @dev Used to test reward claim recovery scenarios
 */
contract MockFailingERC20 is ERC20 {
    bool public transferShouldFail;

    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _mint(msg.sender, 1000000 ether);
    }

    function mint(address to, uint amount) external {
        _mint(to, amount);
    }

    function setTransferShouldFail(bool _shouldFail) external {
        transferShouldFail = _shouldFail;
    }

    function transfer(address to, uint amount) public override returns (bool) {
        if (transferShouldFail) return false;
        return super.transfer(to, amount);
    }

    function transferFrom(address from, address to, uint amount) public override returns (bool) {
        if (transferShouldFail) return false;
        return super.transferFrom(from, to, amount);
    }
}

/**
 * @title CrossGameRewardPoolClaimRecoveryTest
 * @notice Tests reward claim recovery after transfer failures
 * @dev Verifies that users can reclaim rewards after withdraw even if initial transfer failed
 */
contract CrossGameRewardPoolClaimRecoveryTest is CrossGameRewardPoolBase {
    MockFailingERC20 public failingToken;

    function setUp() public override {
        super.setUp();

        // Deploy failing token and register it
        failingToken = new MockFailingERC20("Failing Token", "FAIL");
        crossGameReward.addRewardToken(1, IERC20(address(failingToken)));
    }

    // ==================== Basic Recovery Tests ====================

    /**
     * @notice Tests basic recovery: claim fails → withdraw → fix token → claim succeeds
     */
    function testClaimRecoveryAfterWithdraw() public {
        // Setup: User deposits
        _userDeposit(user1, 10 ether);

        // Deposit rewards
        vm.prank(owner);
        failingToken.transfer(address(pool), 100 ether);

        // Verify rewards are pending
        uint pendingReward = pool.pendingReward(user1, IERC20(address(failingToken)));
        assertEq(pendingReward, 100 ether, "Rewards should be pending");

        // Make token fail
        failingToken.setTransferShouldFail(true);

        // Try to claim - transfer fails but doesn't revert
        vm.prank(user1);
        vm.expectEmit(true, true, false, true);
        emit RewardClaimFailed(user1, IERC20(address(failingToken)), 100 ether);
        pool.claimReward(IERC20(address(failingToken)));

        // Verify rewards are still stored
        (, uint storedRewards) = pool.userRewards(user1, IERC20(address(failingToken)));
        assertEq(storedRewards, 100 ether, "Rewards should still be stored after failed claim");

        // Verify user didn't receive tokens
        assertEq(failingToken.balanceOf(user1), 0, "User should not have received tokens");

        // User withdraws principal
        vm.prank(user1);
        pool.withdraw();

        // Verify balance is now 0
        assertEq(pool.balances(user1), 0, "Balance should be 0 after withdraw");

        // Fix the token
        failingToken.setTransferShouldFail(false);

        // User should still be able to claim rewards even with 0 balance
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));

        // Verify user received rewards
        assertEq(failingToken.balanceOf(user1), 100 ether, "User should receive rewards after recovery");

        // Verify rewards are cleared
        (, storedRewards) = pool.userRewards(user1, IERC20(address(failingToken)));
        assertEq(storedRewards, 0, "Rewards should be cleared after successful claim");
    }

    /**
     * @notice Tests claimRewards() recovery (all tokens at once)
     */
    function testClaimAllRewardsRecoveryAfterWithdraw() public {
        // Setup: User deposits
        _userDeposit(user1, 10 ether);

        // Deposit multiple rewards
        vm.prank(owner);
        failingToken.transfer(address(pool), 100 ether);
        _depositReward(address(rewardToken1), 50 ether);

        // Make failing token fail
        failingToken.setTransferShouldFail(true);

        // Claim all rewards - one fails, one succeeds
        vm.prank(user1);
        pool.claimRewards();

        // Verify partial success
        assertEq(rewardToken1.balanceOf(user1), 50 ether, "Normal token should be claimed");
        assertEq(failingToken.balanceOf(user1), 0, "Failing token should not be claimed");

        // Verify failing token rewards are still stored
        (, uint storedRewards) = pool.userRewards(user1, IERC20(address(failingToken)));
        assertEq(storedRewards, 100 ether, "Failing token rewards should be stored");

        // User withdraws
        vm.prank(user1);
        pool.withdraw();

        // Fix the token
        failingToken.setTransferShouldFail(false);

        // User can claim the failed rewards
        vm.prank(user1);
        pool.claimRewards();

        // Verify all rewards received
        assertEq(failingToken.balanceOf(user1), 100 ether, "Should receive failed rewards");
    }

    /**
     * @notice Tests that users without stored rewards cannot claim with 0 balance
     */
    function testCannotClaimWithZeroBalanceAndNoRewards() public {
        // User has no deposit and no stored rewards
        vm.prank(user1);
        vm.expectRevert(CrossGameRewardPool.CGRPNoDepositFound.selector);
        pool.claimRewards();

        vm.prank(user1);
        vm.expectRevert(CrossGameRewardPool.CGRPNoDepositFound.selector);
        pool.claimReward(IERC20(address(failingToken)));
    }

    /**
     * @notice Tests that users can still claim if they have stored rewards but no balance
     */
    function testCanClaimWithZeroBalanceButStoredRewards() public {
        // Setup: User deposits and earns rewards
        _userDeposit(user1, 10 ether);

        vm.prank(owner);
        failingToken.transfer(address(pool), 100 ether);

        // Make token fail and claim
        failingToken.setTransferShouldFail(true);
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));

        // Verify rewards are stored
        (, uint storedRewards) = pool.userRewards(user1, IERC20(address(failingToken)));
        assertEq(storedRewards, 100 ether, "Rewards should be stored");

        // Withdraw principal
        vm.prank(user1);
        pool.withdraw();

        // Fix token
        failingToken.setTransferShouldFail(false);

        // Should be able to claim with 0 balance but stored rewards
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));

        assertEq(failingToken.balanceOf(user1), 100 ether, "Should receive rewards");
    }

    // ==================== Multiple Users Recovery Tests ====================

    /**
     * @notice Tests recovery with multiple users having failed claims
     */
    function testMultipleUsersRecovery() public {
        // Both users deposit
        _userDeposit(user1, 30 ether);
        _userDeposit(user2, 70 ether);

        // Deposit rewards (100 total)
        vm.prank(owner);
        failingToken.transfer(address(pool), 100 ether);

        // Verify distribution: user1 gets 30%, user2 gets 70%
        assertApproxEqAbs(pool.pendingReward(user1, IERC20(address(failingToken))), 30 ether, 10, "User1: 30%");
        assertApproxEqAbs(pool.pendingReward(user2, IERC20(address(failingToken))), 70 ether, 10, "User2: 70%");

        // Make token fail
        failingToken.setTransferShouldFail(true);

        // Both users try to claim
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));
        vm.prank(user2);
        pool.claimReward(IERC20(address(failingToken)));

        // Both withdraw
        vm.prank(user1);
        pool.withdraw();
        vm.prank(user2);
        pool.withdraw();

        // Fix token
        failingToken.setTransferShouldFail(false);

        // Both can recover
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));
        vm.prank(user2);
        pool.claimReward(IERC20(address(failingToken)));

        // Verify both received correct amounts
        assertApproxEqAbs(failingToken.balanceOf(user1), 30 ether, 10, "User1 recovered 30%");
        assertApproxEqAbs(failingToken.balanceOf(user2), 70 ether, 10, "User2 recovered 70%");
    }

    /**
     * @notice Tests that partial claim failures don't affect other users
     */
    function testPartialFailureDoesNotAffectOtherUsers() public {
        _userDeposit(user1, 50 ether);
        _userDeposit(user2, 50 ether);

        // Deposit rewards
        vm.prank(owner);
        failingToken.transfer(address(pool), 100 ether);

        // Only user1's claim fails
        failingToken.setTransferShouldFail(true);
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));

        // User2's claim succeeds
        failingToken.setTransferShouldFail(false);
        vm.prank(user2);
        pool.claimReward(IERC20(address(failingToken)));

        // Verify user2 received rewards
        assertApproxEqAbs(failingToken.balanceOf(user2), 50 ether, 10, "User2 should receive rewards");

        // Verify user1 still has stored rewards
        (, uint storedRewards) = pool.userRewards(user1, IERC20(address(failingToken)));
        assertApproxEqAbs(storedRewards, 50 ether, 10, "User1 should have stored rewards");

        // User1 can claim before withdrawing
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));

        assertApproxEqAbs(failingToken.balanceOf(user1), 50 ether, 10, "User1 recovered rewards");
    }

    // ==================== Edge Cases ====================

    /**
     * @notice Tests repeated claim attempts with failed token
     */
    function testRepeatedClaimAttempts() public {
        _userDeposit(user1, 10 ether);

        vm.prank(owner);
        failingToken.transfer(address(pool), 100 ether);

        failingToken.setTransferShouldFail(true);

        // Try multiple times while failing
        for (uint i = 0; i < 3; i++) {
            vm.prank(user1);
            pool.claimReward(IERC20(address(failingToken)));

            // Rewards should still be stored
            (, uint storedRewards) = pool.userRewards(user1, IERC20(address(failingToken)));
            assertEq(storedRewards, 100 ether, "Rewards should persist after failed claims");
        }

        // Withdraw
        vm.prank(user1);
        pool.withdraw();

        // Try multiple times with 0 balance
        for (uint i = 0; i < 2; i++) {
            vm.prank(user1);
            pool.claimReward(IERC20(address(failingToken)));

            // Should still work and rewards persist
            (, uint storedRewards) = pool.userRewards(user1, IERC20(address(failingToken)));
            assertEq(storedRewards, 100 ether, "Rewards should persist");
        }

        // Finally fix and claim
        failingToken.setTransferShouldFail(false);
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));

        assertEq(failingToken.balanceOf(user1), 100 ether, "Should receive all rewards");
    }

    /**
     * @notice Tests recovery with removed token
     */
    function testRecoveryWithRemovedToken() public {
        _userDeposit(user1, 10 ether);

        vm.prank(owner);
        failingToken.transfer(address(pool), 100 ether);

        // Remove the token
        crossGameReward.removeRewardToken(1, IERC20(address(failingToken)));

        // Make token fail
        failingToken.setTransferShouldFail(true);

        // Try to claim (should still work for removed tokens)
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));

        // Verify rewards stored
        (, uint storedRewards) = pool.userRewards(user1, IERC20(address(failingToken)));
        assertEq(storedRewards, 100 ether, "Rewards should be stored for removed token");

        // Withdraw
        vm.prank(user1);
        pool.withdraw();

        // Fix and claim
        failingToken.setTransferShouldFail(false);
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));

        assertEq(failingToken.balanceOf(user1), 100 ether, "Should recover removed token rewards");
    }

    /**
     * @notice Tests that new deposits don't interfere with stored rewards
     */
    function testNewDepositDoesNotAffectStoredRewards() public {
        // Initial deposit and failed claim
        _userDeposit(user1, 10 ether);

        vm.prank(owner);
        failingToken.transfer(address(pool), 100 ether);

        failingToken.setTransferShouldFail(true);
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));

        // Verify stored rewards
        (, uint storedRewards) = pool.userRewards(user1, IERC20(address(failingToken)));
        assertEq(storedRewards, 100 ether, "Initial stored rewards");

        // Withdraw - this will try to claim but fail, keeping stored rewards
        // Note: withdraw calls _claimRewards which will fail for failingToken
        vm.prank(user1);
        pool.withdraw();

        // User deposits again (new cycle)
        _userDeposit(user1, 20 ether);

        // Fix token temporarily to add more rewards
        failingToken.setTransferShouldFail(false);
        vm.prank(owner);
        failingToken.transfer(address(pool), 50 ether);
        // Break it again for later test
        failingToken.setTransferShouldFail(true);

        // Stored rewards (100) should still exist plus new pending rewards (50)
        uint totalPending = pool.pendingReward(user1, IERC20(address(failingToken)));
        assertEq(totalPending, 150 ether, "Should have old stored + new pending");

        // Fix and claim all
        failingToken.setTransferShouldFail(false);
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));

        assertEq(failingToken.balanceOf(user1), 150 ether, "Should receive all rewards");
    }

    // ==================== Gas Optimization Tests ====================

    /**
     * @notice Tests that claiming with 0 balance uses less gas (no sync/update)
     */
    function testGasOptimizationWithZeroBalance() public {
        // Setup failed claim and withdraw
        _userDeposit(user1, 10 ether);

        vm.prank(owner);
        failingToken.transfer(address(pool), 100 ether);

        failingToken.setTransferShouldFail(true);
        vm.prank(user1);
        pool.claimReward(IERC20(address(failingToken)));

        vm.prank(user1);
        pool.withdraw();

        failingToken.setTransferShouldFail(false);

        // Measure gas for claim with 0 balance
        vm.prank(user1);
        uint gasBefore = gasleft();
        pool.claimReward(IERC20(address(failingToken)));
        uint gasUsed = gasBefore - gasleft();

        // Should use less gas since no sync/update operations
        // Just a sanity check that it doesn't revert and completes
        assertEq(failingToken.balanceOf(user1), 100 ether, "Claim should succeed");

        // Gas should be reasonable (less than claiming with active balance)
        // This is more of a smoke test
        assertLt(gasUsed, 200000, "Gas should be reasonable");
    }

    // ==================== Events ====================

    event RewardClaimFailed(address indexed account, IERC20 indexed token, uint amount);
    event RewardClaimed(address indexed account, IERC20 indexed token, uint amount);
}
