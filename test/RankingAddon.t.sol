// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../src/addons/RankingAddon.sol";
import "./BaseTest.sol";

/**
 * @title RankingAddonTest
 * @notice RankingAddon 테스트
 */
contract RankingAddonTest is BaseTest {
    RankingAddon public rankingAddon;

    // 테스트용 계정 별칭
    address alice;
    address bob;
    address carol;
    uint projectID;
    uint firstSeasonStartBlock;
    uint seasonBlocks;

    event RankingUpdated(uint indexed season, address indexed user, uint score);
    event TopStakerAdded(uint indexed season, address indexed user);
    event SeasonRankingFinalized(uint indexed season, uint totalParticipants);

    function setUp() public override {
        super.setUp();

        // 별칭 설정
        alice = user1;
        bob = user2;
        carol = user3;
        projectID = PROJECT_ID;
        firstSeasonStartBlock = block.number;
        seasonBlocks = SEASON_BLOCKS;

        // RankingAddon 배포
        rankingAddon = new RankingAddon(address(stakingPool));

        // Addon 승인 및 설정 (owner가 protocol admin)
        vm.startPrank(owner);
        protocol.setPoolAddonApproved(projectID, IStakingAddon(address(rankingAddon)), true);
        protocol.setPoolStakingAddon(projectID, IStakingAddon(address(rankingAddon)));
        vm.stopPrank();
    }

    // ============================================
    // 기본 기능 테스트
    // ============================================

    function test_OnStake_UpdatesRankingScore() public {
        // Alice 스테이킹
        uint stakeAmount = 100 ether;
        deal(address(wcross), alice, stakeAmount);

        vm.startPrank(alice);
        wcross.approve(address(stakingPool), stakeAmount);

        vm.expectEmit(true, true, false, true);
        emit RankingUpdated(1, alice, stakeAmount);

        stakingPool.stake(stakeAmount);
        vm.stopPrank();

        // 랭킹 점수 확인
        uint score = rankingAddon.getUserRankingScore(1, alice);
        assertEq(score, stakeAmount, "Ranking score should equal staked amount");
    }

    function test_OnStake_AddsToTopStakers() public {
        // Alice 스테이킹
        uint stakeAmount = 100 ether;
        deal(address(wcross), alice, stakeAmount);

        vm.startPrank(alice);
        wcross.approve(address(stakingPool), stakeAmount);

        vm.expectEmit(true, true, false, false);
        emit TopStakerAdded(1, alice);

        stakingPool.stake(stakeAmount);
        vm.stopPrank();

        // Top 스테이커 목록 확인
        (address[] memory stakers, uint[] memory scores) = rankingAddon.getTopStakers(1, 0, 10);
        assertEq(stakers.length, 1, "Should have 1 top staker");
        assertEq(stakers[0], alice, "Alice should be in top stakers");
        assertEq(scores[0], stakeAmount, "Score should match");
    }

    function test_OnStake_UpdatesExistingScore() public {
        // 첫 번째 스테이킹
        uint firstAmount = 100 ether;
        deal(address(wcross), alice, firstAmount * 2);

        vm.startPrank(alice);
        wcross.approve(address(stakingPool), firstAmount * 2);
        stakingPool.stake(firstAmount);

        // 두 번째 스테이킹
        stakingPool.stake(firstAmount);
        vm.stopPrank();

        // 랭킹 점수 확인 (누적)
        uint score = rankingAddon.getUserRankingScore(1, alice);
        assertEq(score, firstAmount * 2, "Score should be cumulative");
    }

    function test_OnWithdraw_DecreasesScore() public {
        // 스테이킹
        uint stakeAmount = 100 ether;
        deal(address(wcross), alice, stakeAmount);

        vm.startPrank(alice);
        wcross.approve(address(stakingPool), stakeAmount);
        stakingPool.stake(stakeAmount);

        // 출금
        stakingPool.withdrawAll();
        vm.stopPrank();

        // 랭킹 점수 확인
        uint score = rankingAddon.getUserRankingScore(1, alice);
        assertEq(score, 0, "Score should be 0 after withdraw");
    }

    function test_OnSeasonEnd_FinalizesRanking() public {
        // Alice와 Bob 스테이킹
        _stakeAs(alice, 100 ether);
        _stakeAs(bob, 200 ether);

        // 시즌 종료까지 진행
        vm.roll(firstSeasonStartBlock + seasonBlocks + 1);

        // 시즌 롤오버 (onSeasonEnd 호출됨)
        vm.expectEmit(true, false, false, false);
        emit SeasonRankingFinalized(1, 2);

        stakingPool.rolloverSeason();

        // 참여자 수 확인
        uint participants = rankingAddon.getTotalParticipants(1);
        assertEq(participants, 2, "Should have 2 participants");
    }

    // ============================================
    // 조회 함수 테스트
    // ============================================

    function test_GetTopStakers_ReturnsCorrectData() public {
        // 여러 유저 스테이킹
        _stakeAs(alice, 100 ether);
        _stakeAs(bob, 200 ether);
        _stakeAs(carol, 150 ether);

        // Top 스테이커 조회
        (address[] memory stakers, uint[] memory scores) = rankingAddon.getTopStakers(1, 0, 10);

        assertEq(stakers.length, 3, "Should have 3 stakers");
        assertEq(stakers[0], alice, "Alice should be first added");
        assertEq(stakers[1], bob, "Bob should be second added");
        assertEq(stakers[2], carol, "Carol should be third added");

        assertEq(scores[0], 100 ether, "Alice score");
        assertEq(scores[1], 200 ether, "Bob score");
        assertEq(scores[2], 150 ether, "Carol score");
    }

    function test_GetTopStakers_WithPagination() public {
        // 5명 스테이킹
        address[] memory users = new address[](5);
        users[0] = alice;
        users[1] = bob;
        users[2] = carol;
        users[3] = makeAddr("dave");
        users[4] = makeAddr("eve");

        for (uint i = 0; i < 5; i++) {
            _stakeAs(users[i], (i + 1) * 100 ether);
        }

        // 첫 2개 조회
        (address[] memory stakers1, uint[] memory scores1) = rankingAddon.getTopStakers(1, 0, 2);
        assertEq(stakers1.length, 2, "Should return 2 stakers");
        assertEq(stakers1[0], alice, "First should be alice");

        // 다음 2개 조회
        (address[] memory stakers2, uint[] memory scores2) = rankingAddon.getTopStakers(1, 2, 2);
        assertEq(stakers2.length, 2, "Should return 2 stakers");
        assertEq(stakers2[0], carol, "First of second page should be carol");
    }

    function test_GetTopStakers_EmptySeason() public {
        // 빈 시즌 조회
        (address[] memory stakers, uint[] memory scores) = rankingAddon.getTopStakers(999, 0, 10);

        assertEq(stakers.length, 0, "Should return empty array");
        assertEq(scores.length, 0, "Scores should be empty");
    }

    function test_GetUserRankingScore_ReturnsZeroForNonStaker() public {
        uint score = rankingAddon.getUserRankingScore(1, makeAddr("nonStaker"));
        assertEq(score, 0, "Non-staker should have 0 score");
    }

    function test_GetTotalParticipants_ReturnsCorrectCount() public {
        _stakeAs(alice, 100 ether);
        _stakeAs(bob, 200 ether);

        uint participants = rankingAddon.getTotalParticipants(1);
        assertEq(participants, 2, "Should have 2 participants");
    }

    // ============================================
    // 엣지 케이스 테스트
    // ============================================

    function test_OnStake_DoesNotDuplicateInTopStakers() public {
        _stakeAs(alice, 100 ether);

        // Alice 추가 스테이킹
        deal(address(wcross), alice, 100 ether);
        vm.startPrank(alice);
        wcross.approve(address(stakingPool), 100 ether);
        stakingPool.stake(100 ether);
        vm.stopPrank();

        // Top 스테이커 목록 확인
        (address[] memory stakers,) = rankingAddon.getTopStakers(1, 0, 10);
        assertEq(stakers.length, 1, "Alice should appear only once");
    }

    function test_OnSeasonEnd_PreservesDataAcrossSeasons() public {
        // 시즌 1에 스테이킹
        _stakeAs(alice, 100 ether);

        uint season1Score = rankingAddon.getUserRankingScore(1, alice);

        // 시즌 2로 롤오버
        vm.roll(firstSeasonStartBlock + seasonBlocks + 1);
        stakingPool.rolloverSeason();

        // 시즌 1 데이터 확인
        uint season1ScoreAfter = rankingAddon.getUserRankingScore(1, alice);
        assertEq(season1Score, season1ScoreAfter, "Season 1 score should be preserved");

        // 시즌 2에서는 0
        uint season2Score = rankingAddon.getUserRankingScore(2, alice);
        assertEq(season2Score, 0, "Season 2 score should start at 0");
    }

    function test_MultipleSeasons_IndependentRankings() public {
        // 시즌 1
        _stakeAs(alice, 100 ether);
        _stakeAs(bob, 200 ether);

        // 시즌 2로 롤오버
        vm.roll(firstSeasonStartBlock + seasonBlocks + 1);
        stakingPool.rolloverSeason();

        // 시즌 2에서는 carol이 최고
        _stakeAs(carol, 300 ether);

        // 각 시즌의 참여자 수 확인
        uint season1Participants = rankingAddon.getTotalParticipants(1);
        uint season2Participants = rankingAddon.getTotalParticipants(2);

        assertEq(season1Participants, 2, "Season 1 should have 2 participants");
        assertGe(season2Participants, 1, "Season 2 should have at least 1 participant");
    }

    // ============================================
    // 권한 테스트
    // ============================================

    function test_OnlyPool_CanCallCallbacks() public {
        vm.prank(alice);
        vm.expectRevert("Only pool");
        rankingAddon.onStake(alice, 100 ether, 0, 100 ether, 1);

        vm.prank(alice);
        vm.expectRevert("Only pool");
        rankingAddon.onWithdraw(alice, 100 ether, 1);

        vm.prank(alice);
        vm.expectRevert("Only pool");
        rankingAddon.onSeasonEnd(1, 1000 ether, 5000);

        vm.prank(alice);
        vm.expectRevert("Only pool");
        rankingAddon.onClaim(alice, 1, 100, 50 ether);
    }

    // ============================================
    // Helper Functions
    // ============================================

    function _stakeAs(address user, uint amount) internal {
        deal(address(wcross), user, amount);
        vm.startPrank(user);
        wcross.approve(address(stakingPool), amount);
        stakingPool.stake(amount);
        vm.stopPrank();
    }
}
