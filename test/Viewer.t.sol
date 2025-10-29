// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/StakingViewer.sol";
import "./BaseTest.sol";

contract ViewerTest is BaseTest {
    function test_ExpectedSeasonPoints_VirtualAcrossSeasons() public {
        // Arrange: user1 stakes in season 1
        stakeFor(user1, 10 ether);
        assertEq(stakingPool.currentSeason(), 1);

        // Determine season 1 range
        (uint s1,, uint end1,) = stakingPool.getCurrentSeasonInfo();
        assertEq(s1, 1);

        // Move to mid season 3 without causing storage rollover
        uint targetBlock = end1 + SEASON_DURATION + 50; // into season 3
        vm.warp(targetBlock);

        // Act: query via Viewer (virtual expected points)
        StakingViewer viewer = new StakingViewer(address(protocol));

        (uint s2Points, uint s2Total) = viewer.getExpectedSeasonPoints(PROJECT_ID, 2, user1);
        (uint s3Points, uint s3Total) = viewer.getExpectedSeasonPoints(PROJECT_ID, 3, user1);

        // Assert: non-zero as auto-participation continues across seasons
        assertGt(s2Points, 0, "Season 2 expected points should be > 0");
        assertGt(s3Points, 0, "Season 3 expected points should be > 0");

        // Total points should also be calculated for virtual seasons
        assertGt(s2Total, 0, "Season 2 total points should be calculated");
        assertGt(s3Total, 0, "Season 3 total points should be calculated");
    }
}
