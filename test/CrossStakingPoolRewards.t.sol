// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/CrossStakingPoolBase.t.sol";

/**
 * @title CrossStakingPoolRewardsTest
 * @notice 보상 계산, 분배, 클레임 기능 테스트
 */
contract CrossStakingPoolRewardsTest is CrossStakingPoolBase {
    // ==================== 보상 누적 테스트 ====================

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
        _userStake(user1, 1 ether); // 최소 금액
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

    // ==================== 보상 클레임 테스트 ====================

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

        // 첫 번째 보상만 claim
        vm.prank(user1);
        pool.claimReward(address(rewardToken1));

        assertEq(rewardToken1.balanceOf(user1), 100 ether, "Claimed reward1");
        assertEq(rewardToken2.balanceOf(user1), 0, "Not claimed reward2 yet");

        // 두 번째 보상 claim
        vm.prank(user1);
        pool.claimReward(address(rewardToken2));

        assertEq(rewardToken2.balanceOf(user1), 50 ether, "Claimed reward2");
    }

    function testMultipleClaimsAccumulate() public {
        _userStake(user1, 10 ether);
        _warpSeconds(50);

        // 첫 번째 보상
        _depositReward(address(rewardToken1), 50 ether);

        vm.prank(user1);
        pool.claimRewards();
        assertEq(rewardToken1.balanceOf(user1), 50 ether, "First claim");

        _warpSeconds(50);

        // 두 번째 보상
        _depositReward(address(rewardToken1), 50 ether);

        vm.prank(user1);
        pool.claimRewards();
        assertEq(rewardToken1.balanceOf(user1), 100 ether, "Claims accumulate");
    }

    // ==================== 다중 사용자 보상 분배 테스트 ====================

    function testMultipleUsersRewardDistribution() public {
        // User1 stakes 10 CROSS
        _userStake(user1, 10 ether);
        _warpSeconds(50);

        // User2 stakes 10 CROSS
        _userStake(user2, 10 ether);
        _warpSeconds(50);

        // 보상 입금 (total 20 CROSS staked)
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

        // t=100에 보상 입금 (total: 30 CROSS)
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

        // 보상 입금 (total: 100 CROSS)
        _depositReward(address(rewardToken1), 1000 ether);

        uint[] memory rewards1 = pool.pendingRewards(user1);
        uint[] memory rewards2 = pool.pendingRewards(user2);

        // User1: 30%
        // User2: 70%
        assertApproxEqAbs(rewards1[0], 300 ether, 10, "User1: 30%");
        assertApproxEqAbs(rewards2[0], 700 ether, 10, "User2: 70%");
    }

    // ==================== 다중 보상 토큰 테스트 ====================

    function testMultipleRewardTokens() public {
        _userStake(user1, 10 ether);
        _warpSeconds(100);

        _depositReward(address(rewardToken1), 100 ether);
        _depositReward(address(rewardToken2), 200 ether);

        uint[] memory rewards = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 100 ether, 10, "Reward token 1");
        assertApproxEqAbs(rewards[1], 200 ether, 10, "Reward token 2");
    }

    // ==================== 시간에 따른 보상 변화 테스트 ====================

    function testRewardBeforeAndAfterStake() public {
        // Day 1: User A stakes 100 CROSS
        _userStake(user1, 100 ether);
        _warpDays(1);

        // Day 2: 보상1 100 입금
        _depositReward(address(rewardToken1), 100 ether);
        _warpDays(8);

        // Day 10: User B stakes 100 CROSS
        _userStake(user2, 100 ether);
        _warpDays(1);

        // Day 11: 보상2 200 입금
        _depositReward(address(rewardToken1), 200 ether);
        _warpDays(9);

        // Day 20: 보상 확인
        uint[] memory rewardsA = pool.pendingRewards(user1);
        uint[] memory rewardsB = pool.pendingRewards(user2);

        // User A: 100 (보상1, 100% 전체) + 100 (보상2, 50% 절반) = 200
        // User B: 0 (보상1, 예치 전) + 100 (보상2, 50% 절반) = 100
        assertApproxEqAbs(rewardsA[0], 200 ether, 10, "User A: reward1 + 50% reward2");
        assertApproxEqAbs(rewardsB[0], 100 ether, 10, "User B: only 50% reward2");

        // 총합 검증
        assertApproxEqAbs(rewardsA[0] + rewardsB[0], 300 ether, 20, "Total distributed");
    }

    // ==================== 엣지 케이스 ====================

    function testZeroStakers() public {
        // 스테이커 없이 보상 입금
        _depositReward(address(rewardToken1), 100 ether);

        // 첫 스테이커가 들어오면 이전 보상을 받음
        _userStake(user1, 10 ether);

        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 100 ether, "First staker gets all rewards deposited when pool was empty");
    }

    function testInvalidRewardTokenIndex() public {
        _userStake(user1, 10 ether);

        vm.prank(user1);
        vm.expectRevert(CrossStakingPool.CSPInvalidRewardToken.selector);
        pool.claimReward(address(0xdead));
    }

    function testZeroAmountTransfer() public {
        _userStake(user1, 10 ether);

        // 0 금액 transfer는 아무 효과 없음
        vm.startPrank(owner);
        rewardToken1.transfer(address(pool), 0);
        vm.stopPrank();

        // 보상이 없는지 확인
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

    // ==================== 직접 Transfer 감지 테스트 ====================

    function testDirectTransferDetection() public {
        _userStake(user1, 10 ether);

        // 직접 transfer
        vm.prank(owner);
        rewardToken1.transfer(address(pool), 100 ether);

        // 누군가 stake하면 감지됨 (또는 unstake, claim)
        _userStake(user2, 10 ether);

        // 반영됨
        uint[] memory rewardsAfter = pool.pendingRewards(user1);
        assertEq(rewardsAfter[0], 100 ether, "Direct transfer detected on next action");
    }

    function testDirectTransferWithDepositReward() public {
        _userStake(user1, 10 ether);

        // 직접 transfer 먼저
        vm.prank(owner);
        rewardToken1.transfer(address(pool), 50 ether);

        // 그 다음 추가 transfer
        // RewardDistributed 이벤트는 실제 델타(150)를 기록
        _depositReward(address(rewardToken1), 100 ether);

        // 실제 보상 확인 (직접 transfer 50 + 추가 transfer 100 = 150)
        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 150 ether, "Total rewards should include direct transfer");
    }

    function testMultipleDirectTransfers() public {
        _userStake(user1, 10 ether);

        // 여러 번 직접 transfer
        vm.startPrank(owner);
        rewardToken1.transfer(address(pool), 30 ether);
        rewardToken1.transfer(address(pool), 20 ether);
        rewardToken1.transfer(address(pool), 50 ether);
        vm.stopPrank();

        // 한 번의 액션으로 모두 감지
        _userStake(user2, 10 ether);

        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 100 ether, "All direct transfers should be detected");
    }

    // ==================== Dust 자동 재분배 테스트 ====================

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
