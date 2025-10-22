// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./BaseTest.sol";

/**
 * @title MultiPoolTest
 * @notice 여러 풀 생성 및 관리 테스트
 */
contract MultiPoolTest is BaseTest {
    StakingPool public pool2;
    StakingPool public pool3;
    RewardPool public rewardPool2;
    RewardPool public rewardPool3;
    StakingRouter public router2;
    StakingRouter public router3;

    uint constant PROJECT_ID_2 = 2;
    uint constant PROJECT_ID_3 = 3;

    function setUp() public override {
        super.setUp();

        vm.startPrank(owner);

        // Protocol을 통해 추가 프로젝트 생성 (즉시 시작)
        (, address poolAddr2, address rewardPoolAddr2) =
            protocol.createProject("TestProject2", SEASON_BLOCKS * 2, block.number, 0, address(0));
        pool2 = StakingPool(poolAddr2);
        rewardPool2 = RewardPool(rewardPoolAddr2);
        router2 = new StakingRouter(address(wcross), address(protocol));
        protocol.setApprovedRouter(PROJECT_ID_2, address(router2), true);

        (, address poolAddr3, address rewardPoolAddr3) =
            protocol.createProject("TestProject3", SEASON_BLOCKS / 2, block.number, 0, address(0));
        pool3 = StakingPool(poolAddr3);
        rewardPool3 = RewardPool(rewardPoolAddr3);
        router3 = new StakingRouter(address(wcross), address(protocol));
        protocol.setApprovedRouter(PROJECT_ID_3, address(router3), true);

        vm.stopPrank();
    }

    function test_MultiplePoolsCreated() public view {
        assert(address(stakingPool) != address(0));
        assert(address(pool2) != address(0));
        assert(address(pool3) != address(0));
    }

    function test_StakeInDifferentPools() public {
        // Pool 1에 스테이킹
        vm.prank(user1);
        router.stake{value: 10 ether}(PROJECT_ID);

        // Pool 2에 스테이킹
        vm.prank(user1);
        router2.stake{value: 5 ether}(2);

        (uint balance1,,) = stakingPool.getStakePosition(user1);
        (uint balance2,,) = pool2.getStakePosition(user1);

        assertEq(balance1, 10 ether);
        assertEq(balance2, 5 ether);
    }

    function test_IndependentSeasons() public {
        vm.prank(user1);
        router.stake{value: 10 ether}(PROJECT_ID);

        vm.prank(user2);
        router2.stake{value: 10 ether}(2);

        assertEq(stakingPool.currentSeason(), 1);
        assertEq(pool2.currentSeason(), 1);

        // Pool 1 롤오버 (실제 컨트랙트의 seasonBlocks 사용)
        uint seasonBlocks = stakingPool.seasonBlocks();
        vm.roll(block.number + seasonBlocks + 1);
        stakingPool.rolloverSeason();

        assertEq(stakingPool.currentSeason(), 2);
        assertEq(pool2.currentSeason(), 1); // Pool 2는 아직 시즌 1
    }

    function test_DifferentSeasonLengths() public view {
        assertEq(stakingPool.seasonBlocks(), SEASON_BLOCKS);
        assertEq(pool2.seasonBlocks(), SEASON_BLOCKS * 2);
        assertEq(pool3.seasonBlocks(), SEASON_BLOCKS / 2);
    }

    function test_IndependentRewards() public {
        // Pool 1
        vm.prank(user1);
        router.stake{value: 10 ether}(PROJECT_ID);

        uint seasonBlocks = stakingPool.seasonBlocks();
        vm.roll(block.number + seasonBlocks + 1);
        stakingPool.rolloverSeason();

        vm.startPrank(rewardProvider);
        rewardToken.approve(address(rewardPool), 100 ether);
        rewardPool.fundSeason(1, address(rewardToken), 100 ether);
        vm.stopPrank();

        // Pool 2
        vm.prank(user2);
        router2.stake{value: 10 ether}(2);

        vm.roll(block.number + SEASON_BLOCKS + 1);
        pool2.rolloverSeason();

        vm.startPrank(rewardProvider);
        rewardToken.approve(address(rewardPool2), 50 ether);
        rewardPool2.fundSeason(1, address(rewardToken), 50 ether);
        vm.stopPrank();

        // Claims
        vm.prank(user1);
        stakingPool.claimSeason(1, address(rewardToken));

        vm.prank(user2);
        pool2.claimSeason(1, address(rewardToken));

        assertEq(rewardToken.balanceOf(user1), 100 ether);
        assertEq(rewardToken.balanceOf(user2), 50 ether);
    }

    function test_SameUserMultiplePools() public {
        vm.startPrank(user1);

        router.stake{value: 5 ether}(PROJECT_ID);
        router2.stake{value: 3 ether}(2);
        router3.stake{value: 2 ether}(3);

        vm.stopPrank();

        (uint balance1,,) = stakingPool.getStakePosition(user1);
        (uint balance2,,) = pool2.getStakePosition(user1);
        (uint balance3,,) = pool3.getStakePosition(user1);

        assertEq(balance1, 5 ether);
        assertEq(balance2, 3 ether);
        assertEq(balance3, 2 ether);
    }
}
