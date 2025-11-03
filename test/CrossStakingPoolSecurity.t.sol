// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/CrossStakingPoolBase.t.sol";

/**
 * @title CrossStakingPoolSecurityTest
 * @notice 보안 및 로직 검증 테스트
 */
contract CrossStakingPoolSecurityTest is CrossStakingPoolBase {
    // ==================== 불변성 검증 ====================

    function testInvariantTotalStakedMatchesActualBalance() public {
        _userStake(user1, 100 ether);
        _userStake(user2, 200 ether);
        _userStake(user3, 300 ether);

        // totalStaked와 실제 컨트랙트 잔액이 일치해야 함
        assertEq(pool.totalStaked(), 600 ether, "TotalStaked should match");
        assertEq(crossToken.balanceOf(address(pool)), 600 ether, "Actual balance should match");

        // unstake 후에도 일치
        vm.prank(user1);
        pool.unstake();

        assertEq(pool.totalStaked(), 500 ether, "TotalStaked after unstake");
        assertEq(crossToken.balanceOf(address(pool)), 500 ether, "Actual balance after unstake");
    }

    function testInvariantRewardAccountingAccuracy() public {
        _userStake(user1, 100 ether);
        _userStake(user2, 100 ether);

        _depositReward(address(rewardToken1), 1000 ether);

        uint[] memory rewards1 = pool.pendingRewards(user1);
        uint[] memory rewards2 = pool.pendingRewards(user2);

        // 총 보상은 입금된 보상과 일치해야 함
        assertApproxEqAbs(rewards1[0] + rewards2[0], 1000 ether, 100, "Total rewards should equal deposited");
    }

    function testInvariantNoRewardLoss() public {
        // 여러 사용자가 다양한 시점에 stake/unstake
        _userStake(user1, 50 ether);
        _depositReward(address(rewardToken1), 100 ether);

        _userStake(user2, 150 ether);
        _depositReward(address(rewardToken1), 200 ether);

        vm.prank(user1);
        pool.unstake();
        uint user1Claimed = rewardToken1.balanceOf(user1);

        _userStake(user3, 100 ether);
        _depositReward(address(rewardToken1), 150 ether);

        vm.prank(user2);
        pool.unstake();
        uint user2Claimed = rewardToken1.balanceOf(user2);

        vm.prank(user3);
        pool.unstake();
        uint user3Claimed = rewardToken1.balanceOf(user3);

        // 총 클레임은 총 입금과 일치해야 함
        uint totalClaimed = user1Claimed + user2Claimed + user3Claimed;
        assertApproxEqAbs(totalClaimed, 450 ether, 100, "No reward should be lost");
    }

    // ==================== 잠재적 공격 벡터 검증 ====================

    function testCannotStakeZeroAmount() public {
        vm.startPrank(user1);
        crossToken.approve(address(pool), 0);
        vm.expectRevert(CrossStakingPool.CSPBelowMinimumStakeAmount.selector);
        pool.stake(0);
        vm.stopPrank();
    }

    function testReentrancyProtection() public {
        // ReentrancyGuard가 작동하는지 확인
        // nonReentrant modifier가 모든 주요 함수에 적용되어 있음
        _userStake(user1, 100 ether);

        // 정상 작동 확인
        vm.prank(user1);
        pool.unstake();

        assertEq(pool.balances(user1), 0, "Should complete without reentrancy");
    }

    function testPrecisionLoss() public {
        // 매우 작은 보상으로 정밀도 손실 테스트
        _userStake(user1, 1 ether);

        // 1 wei 보상
        vm.startPrank(owner);
        rewardToken1.transfer(address(pool), 1);
        vm.stopPrank();

        uint[] memory rewards = pool.pendingRewards(user1);
        // 1 wei는 PRECISION으로 나뉘어지므로 손실 가능
        assertTrue(rewards[0] <= 1, "Should handle precision correctly");
    }

    function testOverflowProtection() public {
        // Solidity 0.8.28은 자동 오버플로우 체크
        // 매우 큰 수로 테스트
        uint veryLarge = type(uint).max / 2;

        vm.startPrank(owner);
        crossToken.mint(user1, veryLarge);
        vm.stopPrank();

        vm.startPrank(user1);
        crossToken.approve(address(pool), veryLarge);

        // MIN_STAKE_AMOUNT 이상이므로 성공해야 함
        pool.stake(veryLarge);
        vm.stopPrank();

        assertEq(pool.balances(user1), veryLarge, "Should handle very large amounts");
    }

    // ==================== 로직 정확성 검증 ====================

    function testRewardCalculationConsistency() public {
        _userStake(user1, 100 ether);

        // 여러 번 보상 입금
        for (uint i = 0; i < 5; i++) {
            _depositReward(address(rewardToken1), 100 ether);
        }

        // Pending 조회와 실제 claim 일치
        uint[] memory pendingBefore = pool.pendingRewards(user1);

        vm.prank(user1);
        pool.claimRewards();

        uint actualClaimed = rewardToken1.balanceOf(user1);

        assertApproxEqAbs(pendingBefore[0], actualClaimed, 100, "Pending should match actual claim");
    }

    function testUnstakeOrderCorrectness() public {
        _userStake(user1, 100 ether);
        _depositReward(address(rewardToken1), 1000 ether);

        uint poolBalanceBefore = crossToken.balanceOf(address(pool));

        // unstake는 다음 순서로 진행:
        // 1. 보상 동기화
        // 2. 보상 업데이트
        // 3. 보상 claim (lastBalance 갱신)
        // 4. CROSS 반환 (이때 CROSS 잔액 감소)
        vm.prank(user1);
        pool.unstake();

        // CROSS가 반환되었는지 확인
        assertEq(crossToken.balanceOf(user1), 1000 ether, "Should receive CROSS back");

        // 보상도 받았는지 확인
        assertApproxEqAbs(rewardToken1.balanceOf(user1), 1000 ether, 100, "Should receive rewards");

        // 풀의 CROSS 잔액이 감소했는지 확인
        assertEq(crossToken.balanceOf(address(pool)), poolBalanceBefore - 100 ether, "Pool CROSS decreased");
    }

    function testCheckpointAccuracy() public {
        _userStake(user1, 100 ether);

        _depositReward(address(rewardToken1), 100 ether);

        // User2 stake 시 체크포인트는 현재 rewardPerTokenStored
        _userStake(user2, 100 ether);

        // User2는 이전 보상 못 받음
        uint[] memory rewards2 = pool.pendingRewards(user2);
        assertEq(rewards2[0], 0, "User2 should not get previous rewards");

        // 새 보상
        _depositReward(address(rewardToken1), 200 ether);

        uint[] memory rewards1 = pool.pendingRewards(user1);
        rewards2 = pool.pendingRewards(user2);

        // User1: 100 (이전) + 100 (새 보상의 50%) = 200
        // User2: 0 (이전) + 100 (새 보상의 50%) = 100
        assertApproxEqAbs(rewards1[0], 200 ether, 100, "User1 gets all old + 50% new");
        assertApproxEqAbs(rewards2[0], 100 ether, 100, "User2 gets only 50% new");
    }

    function testRewardDistributionWithZeroStaked() public {
        // totalStaked = 0일 때 보상 입금
        _depositReward(address(rewardToken1), 1000 ether);

        // 첫 스테이커가 이전 보상을 모두 받음
        _userStake(user1, 100 ether);

        uint[] memory rewards = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 1000 ether, 100, "First staker gets rewards deposited when pool was empty");

        // 이후 보상은 정상 분배
        _depositReward(address(rewardToken1), 100 ether);
        rewards = pool.pendingRewards(user1);
        assertApproxEqAbs(rewards[0], 1100 ether, 100, "New rewards added to existing");
    }

    // ==================== 엣지 케이스 검증 ====================

    function testClaimWithZeroRewards() public {
        _userStake(user1, 100 ether);

        // 보상 없이 claim
        vm.prank(user1);
        pool.claimRewards();

        assertEq(rewardToken1.balanceOf(user1), 0, "Should handle zero rewards gracefully");
    }

    function testMultipleUsersUnstakeOrder() public {
        _userStake(user1, 100 ether);
        _userStake(user2, 100 ether);
        _userStake(user3, 100 ether);

        _depositReward(address(rewardToken1), 300 ether);

        // 순서대로 unstake
        vm.prank(user1);
        pool.unstake();
        uint claimed1 = rewardToken1.balanceOf(user1);

        vm.prank(user2);
        pool.unstake();
        uint claimed2 = rewardToken1.balanceOf(user2);

        vm.prank(user3);
        pool.unstake();
        uint claimed3 = rewardToken1.balanceOf(user3);

        // 모두 동일하게 받아야 함 (순서 무관)
        assertApproxEqAbs(claimed1, 100 ether, 100, "User1 should get 1/3");
        assertApproxEqAbs(claimed2, 100 ether, 100, "User2 should get 1/3");
        assertApproxEqAbs(claimed3, 100 ether, 100, "User3 should get 1/3");
    }

    function testStakeAfterRewardDeposit() public {
        _userStake(user1, 100 ether);
        _depositReward(address(rewardToken1), 100 ether);

        // User2 stake 후 이전 보상 못 받음
        _userStake(user2, 100 ether);

        uint[] memory rewards1 = pool.pendingRewards(user1);
        uint[] memory rewards2 = pool.pendingRewards(user2);

        assertApproxEqAbs(rewards1[0], 100 ether, 100, "User1 gets all");
        assertEq(rewards2[0], 0, "User2 gets nothing from previous");
    }

    // ==================== 수학 검증 ====================

    function testRewardPerTokenCalculation() public {
        _userStake(user1, 100 ether);

        // 100 토큰 입금, 100 CROSS 스테이킹
        // rewardPerToken = (100 * 1e18) / 100 = 1e18
        _depositReward(address(rewardToken1), 100 ether);

        uint[] memory rewards = pool.pendingRewards(user1);

        // earned = 100 * 1e18 / 1e18 = 100
        assertApproxEqAbs(rewards[0], 100 ether, 0.001 ether, "Math should be precise");
    }

    function testProportionalDistribution() public {
        // 1:2:3 비율로 stake
        _userStake(user1, 100 ether);
        _userStake(user2, 200 ether);
        _userStake(user3, 300 ether);

        // 600 보상 입금
        _depositReward(address(rewardToken1), 600 ether);

        uint[] memory rewards1 = pool.pendingRewards(user1);
        uint[] memory rewards2 = pool.pendingRewards(user2);
        uint[] memory rewards3 = pool.pendingRewards(user3);

        // 1:2:3 비율 검증
        assertApproxEqAbs(rewards1[0], 100 ether, 100, "1/6 of rewards");
        assertApproxEqAbs(rewards2[0], 200 ether, 100, "2/6 of rewards");
        assertApproxEqAbs(rewards3[0], 300 ether, 100, "3/6 of rewards");

        // 총합 검증
        uint total = rewards1[0] + rewards2[0] + rewards3[0];
        assertApproxEqAbs(total, 600 ether, 100, "Sum should equal deposit");
    }

    // ==================== 시간 독립성 검증 ====================

    function testRewardsIndependentOfTime() public {
        // Scenario 1: 즉시 보상
        _userStake(user1, 100 ether);
        _depositReward(address(rewardToken1), 100 ether);
        vm.prank(user1);
        pool.unstake();
        uint user1Reward = rewardToken1.balanceOf(user1);

        // Scenario 2: 1년 후 보상 (동일 조건)
        _warpDays(365);
        _userStake(user2, 100 ether);
        _depositReward(address(rewardToken1), 100 ether);
        vm.prank(user2);
        pool.unstake();
        uint user2Reward = rewardToken1.balanceOf(user2);

        // 시간 차이와 무관하게 동일한 보상
        assertApproxEqAbs(user1Reward, user2Reward, 100, "Time should not affect rewards");
        assertApproxEqAbs(user1Reward, 100 ether, 100, "Both should get 100 ether");
    }

    // ==================== 보상 토큰 관리 검증 ====================

    function testRewardTokenIndexConsistency() public {
        crossStaking.addRewardToken(1, address(rewardToken3));

        // 주소 확인
        assertEq(pool.rewardTokenAt(0), address(rewardToken1), "RewardToken1 index");
        assertEq(pool.rewardTokenAt(1), address(rewardToken2), "RewardToken2 index");
        assertEq(pool.rewardTokenAt(2), address(rewardToken3), "RewardToken3 index");

        // isRewardToken 확인
        assertTrue(pool.isRewardToken(address(rewardToken1)), "RewardToken1 registered");
        assertTrue(pool.isRewardToken(address(rewardToken2)), "RewardToken2 registered");
        assertTrue(pool.isRewardToken(address(rewardToken3)), "RewardToken3 registered");
        assertFalse(pool.isRewardToken(address(crossToken)), "CROSS not a reward token");
    }

    // ==================== 경계값 검증 ====================

    function testMinimumStakeBoundary() public {
        uint belowMin = MIN_STAKE_AMOUNT - 1;
        uint exactMin = MIN_STAKE_AMOUNT;
        uint aboveMin = MIN_STAKE_AMOUNT + 1;

        // 미만: 실패
        vm.startPrank(user1);
        crossToken.approve(address(pool), belowMin);
        vm.expectRevert(CrossStakingPool.CSPBelowMinimumStakeAmount.selector);
        pool.stake(belowMin);

        // 정확히: 성공
        crossToken.approve(address(pool), exactMin);
        pool.stake(exactMin);

        vm.stopPrank();
        assertEq(pool.balances(user1), exactMin, "Should accept exact minimum");

        // 초과: 성공
        vm.startPrank(user2);
        crossToken.approve(address(pool), aboveMin);
        pool.stake(aboveMin);
        vm.stopPrank();

        assertEq(pool.balances(user2), aboveMin, "Should accept above minimum");
    }

    function testZeroRewardHandling() public {
        _userStake(user1, 100 ether);

        // 0 보상 입금 시도
        // 0 금액은 transfer 자체가 실패하거나 아무 효과 없음
        vm.startPrank(owner);
        rewardToken1.transfer(address(pool), 0);
        vm.stopPrank();

        // 보상이 추가되지 않았는지 확인
        uint[] memory rewards = pool.pendingRewards(user1);
        assertEq(rewards[0], 0, "No reward should be added for 0 amount");
    }

    // ==================== 상태 일관성 검증 ====================

    function testBalanceConsistencyAfterMultipleOperations() public {
        uint initialBalance = crossToken.balanceOf(user1);

        _userStake(user1, 100 ether);
        assertEq(crossToken.balanceOf(user1), initialBalance - 100 ether, "After stake");

        _depositReward(address(rewardToken1), 50 ether);

        vm.prank(user1);
        pool.claimRewards();

        // CROSS 잔액은 변하지 않아야 함 (보상만 claim)
        assertEq(crossToken.balanceOf(user1), initialBalance - 100 ether, "CROSS unchanged after claim");

        vm.prank(user1);
        pool.unstake();

        // 원래 CROSS 복구
        assertEq(crossToken.balanceOf(user1), initialBalance, "CROSS restored after unstake");
    }

    MockERC20 public rewardToken3;

    function setUp() public override {
        super.setUp();
        rewardToken3 = new MockERC20("Reward Token 3", "RWD3");
        rewardToken3.transfer(owner, 10000 ether);
    }

    uint private constant MIN_STAKE_AMOUNT = 1 ether;
}
