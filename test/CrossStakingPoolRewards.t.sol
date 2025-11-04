// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/CrossStakingPoolBase.t.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title CrossStakingPoolRewardsTest
 * @notice Reward accrual, distribution, and claim tests
 */
contract CrossStakingPoolRewardsTest is CrossStakingPoolBase {
    // ==================== Reward accrual tests ====================

    function testRewardAccumulation() public {
        _userStake(user1, 10 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 100 ether);
        _depositReward(address(rewardToken2), 50 ether);

        uint[] memory rewards = pool.pendingRewards(user1);

        assertEq(rewards[0], 100 ether, "Reward token 1 should accumulate");
        assertEq(rewards[1], 50 ether, "Reward token 2 should accumulate");
    }

    function testRewardAccumulationWithVerySmallStake() public {
        _userStake(user1, 1 ether); // minimum stake amount
        _warpSeconds(1000);

        _depositReward(address(rewardToken1), 1000 ether);

        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 1000 ether, "Small stake should get all rewards");
    }

    function testRewardAccumulationWithVeryLargeStake() public {
        _userStake(user1, 1000 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 10000 ether);

        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 10000 ether, "Large stake should get all rewards");
    }

    // ==================== Reward claim tests ====================

    function testClaimRewards() public {
        _userStake(user1, 10 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 100 ether);
        _depositReward(address(rewardToken2), 50 ether);

        vm.startPrank(user1);
        uint rewardBalance1Before = rewardToken1.balanceOf(user1);
        uint rewardBalance2Before = rewardToken2.balanceOf(user1);

        pool.claimRewards();

        uint rewardBalance1After = rewardToken1.balanceOf(user1);
        uint rewardBalance2After = rewardToken2.balanceOf(user1);
        vm.stopPrank();

        assertEq(rewardBalance1After - rewardBalance1Before, 100 ether, "Should receive reward token 1");
        assertEq(rewardBalance2After - rewardBalance2Before, 50 ether, "Should receive reward token 2");
    }

    function testClaimSpecificReward() public {
        _userStake(user1, 10 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 100 ether);
        _depositReward(address(rewardToken2), 50 ether);

        // Claim only the first reward token
        vm.prank(user1);
        pool.claimReward(rewardToken1);

        assertEq(rewardToken1.balanceOf(user1), 100 ether, "Claimed reward1");
        assertEq(rewardToken2.balanceOf(user1), 0, "Not claimed reward2 yet");

        // Claim the second reward token
        vm.prank(user1);
        pool.claimReward(rewardToken2);

        assertEq(rewardToken2.balanceOf(user1), 50 ether, "Claimed reward2");
    }

    function testMultipleClaimsAccumulate() public {
        _userStake(user1, 10 ether);
        _warpSeconds(50);

        // First reward batch
        _depositReward(address(rewardToken1), 50 ether);

        vm.prank(user1);
        pool.claimRewards();
        assertEq(rewardToken1.balanceOf(user1), 50 ether, "First claim");

        _warpSeconds(50);

        // Second reward batch
        _depositReward(address(rewardToken1), 50 ether);

        vm.prank(user1);
        pool.claimRewards();
        assertEq(rewardToken1.balanceOf(user1), 100 ether, "Claims accumulate");
    }

    // ==================== Multi-user reward distribution tests ====================

    function testMultipleUsersRewardDistribution() public {
        // User1 stakes 10 CROSS
        _userStake(user1, 10 ether);
        _warpSeconds(50);

        // User2 stakes 10 CROSS
        _userStake(user2, 10 ether);
        _warpSeconds(50);

        // Deposit rewards (20 CROSS total staked)
        _depositReward(address(rewardToken1), 200 ether);

        uint[] memory rewardsUser1 = pool.pendingRewards(user1);
        uint[] memory rewardsUser2 = pool.pendingRewards(user2);

        // User1: (10 / 20) × 200 = 100 ether
        // User2: (10 / 20) × 200 = 100 ether
        assertEq(rewardsUser1[0], 100 ether, "User1 gets 50% (equal stakes)");
        assertEq(rewardsUser2[0], 100 ether, "User2 gets 50% (equal stakes)");
    }

    function testThreeUsersComplexScenario() public {
        // User1: 10 CROSS at t=0
        _userStake(user1, 10 ether);
        _warpSeconds(50);

        // User2: 20 CROSS at t=50
        _userStake(user2, 20 ether);
        _warpSeconds(50);

        // At t=100 deposit rewards (total: 30 CROSS)
        _depositReward(address(rewardToken1), 150 ether);

        uint[] memory rewards1 = pool.pendingRewards(user1);
        uint[] memory rewards2 = pool.pendingRewards(user2);

        // User1: (10 / 30) × 150 = 50 ether
        // User2: (20 / 30) × 150 = 100 ether
        assertApproxEqAbs(rewards1[0], 50 ether, 10, "User1: 1/3 of rewards");
        assertApproxEqAbs(rewards2[0], 100 ether, 10, "User2: 2/3 of rewards");
    }

    function testRewardDistributionWithUnequalStakes() public {
        // User1: 30 CROSS
        _userStake(user1, 30 ether);

        // User2: 70 CROSS
        _userStake(user2, 70 ether);

        // Deposit rewards (total stake: 100 CROSS)
        _depositReward(address(rewardToken1), 1000 ether);

        uint[] memory rewards1 = pool.pendingRewards(user1);
        uint[] memory rewards2 = pool.pendingRewards(user2);

        // User1: 30%
        // User2: 70%
        assertApproxEqAbs(rewards1[0], 300 ether, 10, "User1: 30%");
        assertApproxEqAbs(rewards2[0], 700 ether, 10, "User2: 70%");
    }

    // ==================== Multiple reward token tests ====================

    function testMultipleRewardTokens() public {
        _userStake(user1, 10 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 100 ether);
        _depositReward(address(rewardToken2), 200 ether);

        uint[] memory rewards = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 100 ether, 10, "Reward token 1");
        assertApproxEqAbs(rewards[1], 200 ether, 10, "Reward token 2");
    }

    // ==================== Time-based reward variations ====================

    function testRewardBeforeAndAfterStake() public {
        // Day 1: User A stakes 100 CROSS
        _userStake(user1, 100 ether);
        _warpDays(1);

        // Day 2: deposit reward #1 (100)
        _depositReward(address(rewardToken1), 100 ether);
        _warpDays(8);

        // Day 10: User B stakes 100 CROSS
        _userStake(user2, 100 ether);
        _warpDays(1);

        // Day 11: deposit reward #2 (200)
        _depositReward(address(rewardToken1), 200 ether);
        _warpDays(9);

        // Day 20: evaluate rewards
        uint[] memory rewardsA = pool.pendingRewards(user1);
        uint[] memory rewardsB = pool.pendingRewards(user2);

        // User A: 100 from reward1 (full share) + 100 from reward2 (50% share) = 200
        // User B: 0 from reward1 (joined later) + 100 from reward2 (50% share) = 100
        assertApproxEqAbs(rewardsA[0], 200 ether, 10, "User A: reward1 + 50% reward2");
        assertApproxEqAbs(rewardsB[0], 100 ether, 10, "User B: only 50% reward2");

        // Total distribution check
        assertApproxEqAbs(rewardsA[0] + rewardsB[0], 300 ether, 20, "Total distributed");
    }

    // ==================== Edge cases ====================

    function testZeroStakers() public {
        // Deposit rewards with no stakers
        _depositReward(address(rewardToken1), 100 ether);

        // First staker should receive rewards accumulated while pool was empty
        _userStake(user1, 10 ether);

        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 100 ether, "First staker gets all rewards deposited when pool was empty");
    }

    function testInvalidRewardTokenIndex() public {
        _userStake(user1, 10 ether);

        vm.prank(user1);
        vm.expectRevert(CrossStakingPool.CSPInvalidRewardToken.selector);
        pool.claimReward(IERC20(address(uint160(0xdead))));
    }

    function testZeroAmountTransfer() public {
        _userStake(user1, 10 ether);

        // Transferring zero should have no effect
        vm.startPrank(owner);
        rewardToken1.transfer(address(pool), 0);
        vm.stopPrank();

        // Verify no rewards were created
        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "No reward for 0 transfer");
    }

    function testPendingRewardsAfterClaim() public {
        _userStake(user1, 10 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // Claim
        vm.prank(user1);
        pool.claimRewards();

        // Pending rewards should be zero
        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "Pending rewards should be zero after claim");

        // New reward
        _depositReward(address(rewardToken1), 50 ether);
        rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 50 ether, "New rewards should accumulate");
    }

    // ==================== Direct transfer detection tests ====================

    function testDirectTransferDetection() public {
        _userStake(user1, 10 ether);

        // Direct transfer without helper
        vm.prank(owner);
        rewardToken1.transfer(address(pool), 100 ether);

        // Detected when another staking action happens (stake/unstake/claim)
        _userStake(user2, 10 ether);

        // Reward should reflect the transfer
        uint[] memory rewardsAfter = pool.pendingRewards(user1);
        assertEq(rewardsAfter[0], 100 ether, "Direct transfer detected on next action");
    }

    function testDirectTransferWithDepositReward() public {
        _userStake(user1, 10 ether);

        // Perform direct transfer first
        vm.prank(owner);
        rewardToken1.transfer(address(pool), 50 ether);

        // Then deposit via helper
        // RewardSynced event records the full delta (150)
        _depositReward(address(rewardToken1), 100 ether);

        // Verify reward equals 50 + 100 = 150
        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 150 ether, "Total rewards should include direct transfer");
    }

    function testMultipleDirectTransfers() public {
        _userStake(user1, 10 ether);

        // Multiple direct transfers
        vm.startPrank(owner);
        rewardToken1.transfer(address(pool), 30 ether);
        rewardToken1.transfer(address(pool), 20 ether);
        rewardToken1.transfer(address(pool), 50 ether);
        vm.stopPrank();

        // One additional action detects them all
        _userStake(user2, 10 ether);

        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 100 ether, "All direct transfers should be detected");
    }

    // ==================== Dust redistribution tests ====================

    /// @notice Tests that rounding dust is minimal thanks to PRECISION
    /// @dev With 1e18 PRECISION, dust is measured in wei, not ether!
    function testDustAutoRedistribution() public {
        // Setup: User1 stakes 1, User2 stakes 2 (1:2 ratio, total 3)
        _userStake(user1, 1 ether);
        _userStake(user2, 2 ether);

        // First distribution: 100 tokens
        _depositReward(address(rewardToken1), 100 ether);

        // Thanks to PRECISION (1e18), distribution is very accurate!
        // rewardPerTokenStored = (100 * 1e18 * 1e18) / 3
        // User1 (1 ether stake): gets exactly 33.333... ether
        // User2 (2 ether stake): gets exactly 66.666... ether

        uint[] memory rewards1 = pool.pendingRewards(user1);
        uint[] memory rewards2 = pool.pendingRewards(user2);

        // Verify almost exact distribution (dust in wei, not ether!)
        assertApproxEqAbs(rewards1[0], 33.333333333333333333 ether, 1, "User1 gets 1/3");
        assertApproxEqAbs(rewards2[0], 66.666666666666666666 ether, 1, "User2 gets 2/3");

        // Track balances before claiming
        uint totalDeposited = 100 ether;

        // Users claim their rewards
        vm.prank(user1);
        pool.claimRewards();
        vm.prank(user2);
        pool.claimRewards();

        uint user1Claimed = rewardToken1.balanceOf(user1);
        uint user2Claimed = rewardToken1.balanceOf(user2);
        uint poolDust = rewardToken1.balanceOf(address(pool));

        // Exact accounting: deposited = claimed + dust
        assertEq(totalDeposited, user1Claimed + user2Claimed + poolDust, "Perfect accounting");

        // Dust should be tiny (wei, not ether!)
        assertLt(poolDust, 2, "Dust is negligible (< 2 wei)");

        // Second distribution: Add 90 tokens
        _depositReward(address(rewardToken1), 90 ether);
        totalDeposited += 90 ether;

        // Claim again
        vm.prank(user1);
        pool.claimRewards();
        vm.prank(user2);
        pool.claimRewards();

        uint user1Total = rewardToken1.balanceOf(user1);
        uint user2Total = rewardToken1.balanceOf(user2);
        uint finalDust = rewardToken1.balanceOf(address(pool));

        // Still perfect accounting
        assertEq(totalDeposited, user1Total + user2Total + finalDust, "Still perfect accounting");

        // Dust remains minimal
        assertLt(finalDust, 3, "Dust still negligible (< 3 wei)");

        // Ratio is maintained
        assertApproxEqRel(user1Total, (totalDeposited * 1) / 3, 0.0001e18, "User1 gets ~1/3");
        assertApproxEqRel(user2Total, (totalDeposited * 2) / 3, 0.0001e18, "User2 gets ~2/3");
    }

    /// @notice Tests dust accumulation remains minimal over many rounds
    /// @dev Multiple rounds with PRECISION means dust stays in wei, not ether
    function testDustEventuallyDistributed() public {
        // Setup: 1:2 ratio
        _userStake(user1, 1 ether);
        _userStake(user2, 2 ether);

        uint totalDeposited = 0;

        // Do 10 rounds of 100 token distribution
        for (uint i = 0; i < 10; i++) {
            _depositReward(address(rewardToken1), 100 ether);
            totalDeposited += 100 ether;
        }

        // Claim all rewards
        vm.prank(user1);
        pool.claimRewards();
        vm.prank(user2);
        pool.claimRewards();

        uint user1Balance = rewardToken1.balanceOf(user1);
        uint user2Balance = rewardToken1.balanceOf(user2);
        uint poolDust = rewardToken1.balanceOf(address(pool));

        // Perfect accounting
        assertEq(totalDeposited, user1Balance + user2Balance + poolDust, "Total matches");

        // Dust is minimal (wei, not ether!)
        assertLt(poolDust, 10, "Dust < 10 wei after 10 rounds");

        // User1 unstakes (leaves only User2)
        vm.prank(user1);
        pool.unstake();

        // Now add one more reward with only User2 staking
        _depositReward(address(rewardToken1), 100 ether);
        totalDeposited += 100 ether;

        // User2 should get all the new reward + previous dust
        vm.prank(user2);
        pool.claimRewards();

        // Final accounting
        uint finalUser1 = rewardToken1.balanceOf(user1);
        uint finalUser2 = rewardToken1.balanceOf(user2);
        uint finalDust = rewardToken1.balanceOf(address(pool));

        assertEq(totalDeposited, finalUser1 + finalUser2 + finalDust, "Final total matches");
        assertLt(finalDust, 2, "Final dust < 2 wei");
    }

    /// @notice Tests that dust doesn't grow unbounded over time
    /// @dev Even with worst-case rounding, dust is automatically redistributed
    function testDustDoesNotAccumulateUnbounded() public {
        // Worst case: many users with odd staking amounts
        _userStake(user1, 1 ether);
        _userStake(user2, 2 ether);
        _userStake(user3, 3 ether);
        // Total: 6 ether

        uint totalDeposited = 0;
        uint totalClaimed = 0;

        // 100 rounds of distribution
        for (uint i = 0; i < 100; i++) {
            _depositReward(address(rewardToken1), 100 ether);
            totalDeposited += 100 ether;
        }

        // Claim all rewards
        vm.prank(user1);
        pool.claimRewards();
        totalClaimed += rewardToken1.balanceOf(user1);

        vm.prank(user2);
        pool.claimRewards();
        totalClaimed += rewardToken1.balanceOf(user2);

        vm.prank(user3);
        pool.claimRewards();
        totalClaimed += rewardToken1.balanceOf(user3);

        uint poolBalance = rewardToken1.balanceOf(address(pool));

        // Dust should be minimal (< 1% of one distribution round)
        assertLt(poolBalance, 1 ether, "Dust should not accumulate significantly");

        // Almost all tokens should be distributed
        assertGt(totalClaimed, totalDeposited - 1 ether, "Most tokens should be claimed");
    }

    /// @notice Tests exact accounting: every wei is tracked
    /// @dev Total deposited = Total claimed + Pool balance (should always hold)
    function testExactDustAccounting() public {
        _userStake(user1, 1 ether);
        _userStake(user2, 2 ether);

        uint totalDeposited = 0;

        // Deposit and claim multiple times
        for (uint i = 0; i < 5; i++) {
            uint depositAmount = 77 ether; // Odd number to create dust
            _depositReward(address(rewardToken1), depositAmount);
            totalDeposited += depositAmount;

            vm.prank(user1);
            pool.claimRewards();
            vm.prank(user2);
            pool.claimRewards();
        }

        uint user1Balance = rewardToken1.balanceOf(user1);
        uint user2Balance = rewardToken1.balanceOf(user2);
        uint poolBalance = rewardToken1.balanceOf(address(pool));

        // Exact accounting: deposited = claimed + remaining
        assertEq(
            totalDeposited,
            user1Balance + user2Balance + poolBalance,
            "Total deposited must equal total claimed + pool balance"
        );
    }
}
