// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./BaseTest.sol";

/**
 * @title RewardsTest
 * @notice 보상 시스템 테스트
 */
contract RewardsTest is BaseTest {
    function test_FundSeason() public {
        fundSeason(1, 100 ether);

        uint seasonReward = rewardPool.seasonRewards(1, address(rewardToken));
        assertEq(seasonReward, 100 ether);
    }

    function test_ProportionalRewardDistribution() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 5 ether);

        rolloverSeason();
        fundSeason(1, 150 ether);

        claimSeasonFor(user1, 1);
        claimSeasonFor(user2, 1);

        uint reward1 = rewardToken.balanceOf(user1);
        uint reward2 = rewardToken.balanceOf(user2);

        assertGt(reward1, 0);
        assertGt(reward2, 0);
        assertGt(reward1, reward2);

        // User1이 2배 스테이킹했으므로 2배 보상
        assertGt(reward1, reward2 * 19 / 10); // 1.9배 이상
    }

    function test_EqualStakeEqualReward() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 10 ether);

        rolloverSeason();
        fundSeason(1, 100 ether);

        claimSeasonFor(user1, 1);
        claimSeasonFor(user2, 1);

        uint reward1 = rewardToken.balanceOf(user1);
        uint reward2 = rewardToken.balanceOf(user2);

        // 동일한 스테이킹이므로 비슷한 보상 (1% 오차 허용)
        uint diff = reward1 > reward2 ? reward1 - reward2 : reward2 - reward1;
        assertLt(diff, reward1 / 100);
    }

    function test_MultipleSeasonRewards() public {
        stakeFor(user1, 10 ether);

        // Season 1
        rolloverSeason();
        fundSeason(1, 100 ether);
        claimSeasonFor(user1, 1);

        uint reward1 = rewardToken.balanceOf(user1);

        // Season 2
        rolloverSeason();
        fundSeason(2, 100 ether);
        claimSeasonFor(user1, 2);

        uint totalReward = rewardToken.balanceOf(user1);

        assertGt(reward1, 0);
        assertGt(totalReward, reward1);
    }

    function test_RevertWhen_ClaimBeforeSeasonEnd() public {
        stakeFor(user1, 10 ether);

        vm.prank(user1);
        vm.expectRevert();
        stakingPool.claimSeason(1, address(rewardToken));
    }

    function test_RevertWhen_DuplicateClaim() public {
        stakeFor(user1, 10 ether);

        rolloverSeason();
        fundSeason(1, 100 ether);

        claimSeasonFor(user1, 1);

        vm.prank(user1);
        vm.expectRevert();
        stakingPool.claimSeason(1, address(rewardToken));
    }

    function test_NoRewardForNoStake() public {
        stakeFor(user1, 10 ether);

        rolloverSeason();
        fundSeason(1, 100 ether);

        claimSeasonFor(user1, 1);

        // User2는 스테이킹 안 함
        uint reward2 = rewardToken.balanceOf(user2);
        assertEq(reward2, 0);
    }

    function test_RemainingRewards() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 5 ether);

        rolloverSeason();
        fundSeason(1, 150 ether);

        // User1만 청구
        claimSeasonFor(user1, 1);

        uint claimed = rewardToken.balanceOf(user1);
        uint remaining = rewardPool.getRemainingRewards(1, address(rewardToken));

        assertGt(remaining, 0);
        assertEq(claimed + remaining, 150 ether);
    }
}
