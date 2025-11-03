// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/CrossStakingPoolBase.t.sol";

/**
 * @title CrossStakingPoolIntegrationTest
 * @notice 복잡한 시나리오 및 통합 테스트
 */
contract CrossStakingPoolIntegrationTest is CrossStakingPoolBase {
    // ==================== 실전 시나리오 테스트 ====================

    function testCompleteUserJourney() public {
        // 초기 CROSS 잔액 기록
        uint user1InitialCross = crossToken.balanceOf(user1);
        uint user2InitialCross = crossToken.balanceOf(user2);
        uint user3InitialCross = crossToken.balanceOf(user3);

        // Day 0: User1 stakes 50 CROSS
        _userStake(user1, 50 ether);
        assertEq(crossToken.balanceOf(user1), user1InitialCross - 50 ether, "User1 CROSS decreased");
        assertEq(pool.balances(user1), 50 ether, "User1 pool balance: 50");
        assertEq(pool.totalStaked(), 50 ether, "Total staked: 50");

        // Day 1: First reward
        _warpDays(1);
        _depositReward(address(rewardToken1), 100 ether);

        // Day 2: User2 stakes 100 CROSS
        _warpDays(1);
        _userStake(user2, 100 ether);
        assertEq(crossToken.balanceOf(user2), user2InitialCross - 100 ether, "User2 CROSS decreased");
        assertEq(pool.balances(user2), 100 ether, "User2 pool balance: 100");
        assertEq(pool.totalStaked(), 150 ether, "Total staked: 150");

        // Day 3: Second reward
        _warpDays(1);
        _depositReward(address(rewardToken1), 150 ether);

        // Day 4: User1 claims (CROSS는 변하지 않음)
        _warpDays(1);
        uint user1CrossBeforeClaim = crossToken.balanceOf(user1);
        vm.prank(user1);
        pool.claimRewards();

        assertEq(crossToken.balanceOf(user1), user1CrossBeforeClaim, "User1 CROSS unchanged after claim");
        assertEq(pool.balances(user1), 50 ether, "User1 stake unchanged");

        uint user1Claimed = rewardToken1.balanceOf(user1);
        assertTrue(user1Claimed > 0, "User1 claimed rewards");

        // Day 5: User3 stakes 150 CROSS
        _warpDays(1);
        _userStake(user3, 150 ether);
        assertEq(crossToken.balanceOf(user3), user3InitialCross - 150 ether, "User3 CROSS decreased");
        assertEq(pool.balances(user3), 150 ether, "User3 pool balance: 150");
        assertEq(pool.totalStaked(), 300 ether, "Total staked: 300");

        // Day 6: Third reward
        _warpDays(1);
        _depositReward(address(rewardToken1), 300 ether);

        // Day 7: Everyone unstakes
        _warpDays(1);

        // User1 unstake
        uint user1RewardBefore = rewardToken1.balanceOf(user1);
        uint user1CrossBefore = crossToken.balanceOf(user1);
        vm.prank(user1);
        pool.unstake();

        assertEq(crossToken.balanceOf(user1), user1CrossBefore + 50 ether, "User1 gets 50 CROSS back");
        assertEq(crossToken.balanceOf(user1), user1InitialCross, "User1 CROSS fully restored");
        assertEq(pool.balances(user1), 0, "User1 balance cleared");
        uint user1Total = rewardToken1.balanceOf(user1) - user1RewardBefore + user1Claimed;

        // User2 unstake
        uint user2CrossBefore = crossToken.balanceOf(user2);
        vm.prank(user2);
        pool.unstake();

        assertEq(crossToken.balanceOf(user2), user2CrossBefore + 100 ether, "User2 gets 100 CROSS back");
        assertEq(crossToken.balanceOf(user2), user2InitialCross, "User2 CROSS fully restored");
        assertEq(pool.balances(user2), 0, "User2 balance cleared");
        uint user2Total = rewardToken1.balanceOf(user2);

        // User3 unstake
        uint user3CrossBefore = crossToken.balanceOf(user3);
        vm.prank(user3);
        pool.unstake();

        assertEq(crossToken.balanceOf(user3), user3CrossBefore + 150 ether, "User3 gets 150 CROSS back");
        assertEq(crossToken.balanceOf(user3), user3InitialCross, "User3 CROSS fully restored");
        assertEq(pool.balances(user3), 0, "User3 balance cleared");
        uint user3Total = rewardToken1.balanceOf(user3);

        // 총 입금: 100 + 150 + 300 = 550 ether
        uint totalRewards = user1Total + user2Total + user3Total;

        assertApproxEqAbs(totalRewards, 550 ether, 100, "Total rewards match deposits");

        // 모든 CROSS 반환 확인
        assertEq(pool.totalStaked(), 0, "All CROSS unstaked from pool");
        assertEq(crossToken.balanceOf(address(pool)), 0, "Pool has no CROSS");
    }

    function testMultipleRewardTokensComplexScenario() public {
        // Users stake
        _userStake(user1, 100 ether);
        _userStake(user2, 200 ether);
        _userStake(user3, 300 ether);

        // Multiple reward deposits over time
        _warpDays(1);
        _depositReward(address(rewardToken1), 600 ether);
        _depositReward(address(rewardToken2), 300 ether);

        _warpDays(1);
        _depositReward(address(rewardToken1), 300 ether);

        _warpDays(1);
        _depositReward(address(rewardToken2), 600 ether);

        // Check rewards
        uint[] memory rewards1 = pool.pendingRewards(user1);
        uint[] memory rewards2 = pool.pendingRewards(user2);
        uint[] memory rewards3 = pool.pendingRewards(user3);

        // Total staked: 600
        // User1: 100/600 = 16.67%
        // User2: 200/600 = 33.33%
        // User3: 300/600 = 50%

        // RewardToken1: 900 total
        assertApproxEqAbs(rewards1[0], 150 ether, 100, "User1 reward1: 16.67%");
        assertApproxEqAbs(rewards2[0], 300 ether, 100, "User2 reward1: 33.33%");
        assertApproxEqAbs(rewards3[0], 450 ether, 100, "User3 reward1: 50%");

        // RewardToken2: 900 total
        assertApproxEqAbs(rewards1[1], 150 ether, 100, "User1 reward2: 16.67%");
        assertApproxEqAbs(rewards2[1], 300 ether, 100, "User2 reward2: 33.33%");
        assertApproxEqAbs(rewards3[1], 450 ether, 100, "User3 reward2: 50%");
    }

    function testDynamicStakingAndUnstaking() public {
        // Initial stakes
        _userStake(user1, 100 ether);
        _userStake(user2, 200 ether);

        _warpDays(1);
        _depositReward(address(rewardToken1), 300 ether);

        // User1 unstakes
        vm.prank(user1);
        pool.unstake();
        uint user1Rewards = rewardToken1.balanceOf(user1);

        // User3 stakes
        _userStake(user3, 300 ether);

        _warpDays(1);
        _depositReward(address(rewardToken1), 500 ether);

        // Now only user2 and user3 share the second reward
        uint[] memory rewards2 = pool.pendingRewards(user2);
        uint[] memory rewards3 = pool.pendingRewards(user3);

        // First reward: user1 got 100/300, user2 got 200/300
        assertApproxEqAbs(user1Rewards, 100 ether, 100, "User1 first reward");

        // Second reward: user2 (200/500) = 200, user3 (300/500) = 300
        // User2 total: 200 (first) + 200 (second) = 400
        assertApproxEqAbs(rewards2[0], 400 ether, 100, "User2 both rewards");
        assertApproxEqAbs(rewards3[0], 300 ether, 100, "User3 only second reward");
    }

    function testRepeatedStakeAndClaim() public {
        for (uint i = 0; i < 5; i++) {
            // Stake
            _userStake(user1, 10 ether);

            // Reward
            _warpDays(1);
            _depositReward(address(rewardToken1), 100 ether);

            // Claim
            vm.prank(user1);
            pool.claimRewards();

            _warpDays(1);
        }

        uint totalClaimed = rewardToken1.balanceOf(user1);
        assertApproxEqAbs(totalClaimed, 500 ether, 0.001 ether, "Should accumulate all claims");

        // Final balance
        assertEq(pool.balances(user1), 50 ether, "Should have accumulated stakes");
    }

    function testLongTermStaking() public {
        _userStake(user1, 100 ether);

        // 1 year of weekly rewards
        for (uint i = 0; i < 52; i++) {
            _warpDays(7);
            _depositReward(address(rewardToken1), 100 ether);
        }

        uint[] memory rewards = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 5200 ether, 0.01 ether, "1 year of rewards");

        // Unstake after 1 year
        vm.prank(user1);
        pool.unstake();

        assertEq(crossToken.balanceOf(user1), 1000 ether, "Should get all CROSS back");
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 5200 ether, 0.01 ether, "Should get all rewards");
    }

    // ==================== 스트레스 테스트 ====================

    function testManyUsersStaking() public {
        // 100명의 사용자가 동일 금액 스테이킹
        address[] memory users = new address[](10); // 간소화: 10명
        for (uint i = 0; i < 10; i++) {
            users[i] = address(uint160(i + 100));
            crossToken.transfer(users[i], 100 ether);
            _userStake(users[i], 10 ether);
        }

        _warpDays(1);
        _depositReward(address(rewardToken1), 1000 ether);

        // 각 사용자는 100 ether씩 받아야 함
        for (uint i = 0; i < 10; i++) {
            uint[] memory rewards = pool.pendingRewards(users[i]);
            assertApproxEqAbs(rewards[0], 100 ether, 100, "Equal distribution");
        }
    }

    function testHighFrequencyRewards() public {
        _userStake(user1, 100 ether);

        // 매일 보상 100번
        for (uint i = 0; i < 100; i++) {
            _depositReward(address(rewardToken1), 10 ether);
        }

        uint[] memory rewards = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 1000 ether, 0.001 ether, "Accumulated many small rewards");
    }

    // ==================== 엣지 케이스 통합 ====================

    function testZeroBalanceAfterMultipleOperations() public {
        _userStake(user1, 100 ether);
        _depositReward(address(rewardToken1), 100 ether);

        vm.startPrank(user1);
        pool.claimRewards();
        pool.unstake();
        vm.stopPrank();

        assertEq(pool.balances(user1), 0, "Balance should be zero");

        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "Pending rewards should be zero");
    }

    function testRewardAccuracyWithPrecision() public {
        // 매우 작은 스테이킹과 큰 보상
        _userStake(user1, 1 ether);
        _depositReward(address(rewardToken1), 10000 ether); // 더 작은 금액으로 조정

        uint[] memory rewards = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 10000 ether, 100, "Should handle large rewards");

        // 매우 큰 스테이킹과 작은 보상
        _userStake(user2, 1000 ether);
        _depositReward(address(rewardToken1), 1 ether);

        uint[] memory rewards2 = pool.pendingRewards(user2);
        // user2는 1000/1001 비율
        assertApproxEqAbs(rewards2[0], 0.999 ether, 0.01 ether, "Should handle small rewards");
    }

    function testSequentialClaimsPreserveAccuracy() public {
        _userStake(user1, 100 ether);

        for (uint i = 0; i < 10; i++) {
            _depositReward(address(rewardToken1), 100 ether);

            vm.prank(user1);
            pool.claimRewards();
        }

        // User1이 총 1000 ether를 받았어야 함
        uint totalClaimed = rewardToken1.balanceOf(user1);
        assertApproxEqAbs(totalClaimed, 1000 ether, 0.001 ether, "Sequential claims should be accurate");
    }

    // ==================== 실제 사용 패턴 시뮬레이션 ====================

    function testTypicalDeFiUsage() public {
        // Week 1: Initial liquidity providers
        _userStake(user1, 500 ether);
        _warpDays(7);

        // Week 2: Protocol starts distributing rewards
        _depositReward(address(rewardToken1), 1000 ether);
        _warpDays(7);

        // Week 3: More users join
        _userStake(user2, 500 ether);
        _depositReward(address(rewardToken1), 2000 ether);
        _warpDays(7);

        // Week 4: User1 claims and adds more
        vm.prank(user1);
        pool.claimRewards();
        _userStake(user1, 200 ether);
        _depositReward(address(rewardToken1), 3000 ether);
        _warpDays(7);

        // Week 5: User2 partially unstakes
        vm.prank(user2);
        pool.unstake();

        // Final state check
        assertTrue(pool.totalStaked() > 0, "Pool should still have stakers");
        assertTrue(rewardToken1.balanceOf(user1) > 0, "User1 should have claimed rewards");
        assertTrue(rewardToken1.balanceOf(user2) > 0, "User2 should have rewards from unstake");
    }
}
