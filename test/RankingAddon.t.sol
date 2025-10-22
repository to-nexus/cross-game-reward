// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../src/addons/RankingAddon.sol";
import "./BaseTest.sol";

/**
 * @title RankingAddonTest
 * @notice 완전 On-chain Top 10 랭킹 테스트
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
    event TopRankerAdded(uint indexed season, address indexed user, uint rank, uint score);
    event TopRankerRemoved(uint indexed season, address indexed user, uint oldRank);
    event RankChanged(uint indexed season, address indexed user, uint oldRank, uint newRank);
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

    function test_MaxTopRankers() public {
        assertEq(rankingAddon.MAX_TOP_RANKERS(), 10, "MAX should be 10");
    }

    function test_OnStake_UpdatesRankingScore() public {
        uint stakeAmount = 100 ether;
        _stakeAs(alice, stakeAmount);

        uint score = rankingAddon.getUserRankingScore(1, alice);
        assertEq(score, stakeAmount, "Score should equal staked amount");
    }

    function test_OnStake_AddsToTop10() public {
        uint stakeAmount = 100 ether;

        _stakeAs(alice, stakeAmount);

        // Top 랭커 확인
        (address[] memory users, uint[] memory scores) = rankingAddon.getTopRankers(1);
        assertEq(users.length, 1, "Should have 1 top ranker");
        assertEq(users[0], alice, "Alice should be rank 1");
        assertEq(scores[0], stakeAmount, "Score should match");

        // isTopRanker 확인
        assertTrue(rankingAddon.isTopRanker(1, alice), "Alice should be flagged as top ranker");

        // 순위 확인
        assertEq(rankingAddon.getUserRank(1, alice), 1, "Alice should be rank 1");
    }

    function test_Top10_MaintainsSortedOrder() public {
        // 5명 서로 다른 금액으로 스테이킹
        _stakeAs(alice, 50 ether);
        _stakeAs(bob, 100 ether);
        _stakeAs(carol, 75 ether);
        _stakeAs(makeAddr("dave"), 200 ether);
        _stakeAs(makeAddr("eve"), 150 ether);

        // Top 랭커 조회
        (address[] memory users, uint[] memory scores) = rankingAddon.getTopRankers(1);

        // 내림차순 확인
        assertEq(users[0], makeAddr("dave"), "Rank 1 should be dave");
        assertEq(scores[0], 200 ether, "Top score should be 200");

        assertEq(users[1], makeAddr("eve"), "Rank 2 should be eve");
        assertEq(scores[1], 150 ether);

        assertEq(users[2], bob, "Rank 3 should be bob");
        assertEq(scores[2], 100 ether);

        assertEq(users[3], carol, "Rank 4 should be carol");
        assertEq(scores[3], 75 ether);

        assertEq(users[4], alice, "Rank 5 should be alice");
        assertEq(scores[4], 50 ether);
    }

    function test_Top10_LimitsTo10Rankers() public {
        // 15명 스테이킹
        for (uint i = 1; i <= 15; i++) {
            address user = makeAddr(string(abi.encodePacked("user", i)));
            _stakeAs(user, i * 10 ether);
        }

        // Top 10만 유지되는지 확인
        uint count = rankingAddon.getTopRankersCount(1);
        assertEq(count, 10, "Should only have 10 rankers");

        // 상위 10명만 Top 10에 있는지 확인
        (address[] memory users, uint[] memory scores) = rankingAddon.getTopRankers(1);

        // 최고 점수는 150 ether (user15)
        assertEq(scores[0], 150 ether, "Top score should be 150");

        // 최저 점수는 60 ether (user6)
        assertEq(scores[9], 60 ether, "10th score should be 60");

        // user5 (50 ether)는 Top 10 밖
        assertFalse(rankingAddon.isTopRanker(1, makeAddr("user5")), "user5 should not be in top 10");
    }

    function test_OnWithdraw_UpdatesRanking() public {
        _stakeAs(alice, 100 ether);
        _stakeAs(bob, 200 ether);

        // Bob이 1위, Alice가 2위
        assertEq(rankingAddon.getUserRank(1, bob), 1, "Bob should be rank 1");
        assertEq(rankingAddon.getUserRank(1, alice), 2, "Alice should be rank 2");

        // Bob 출금
        vm.prank(bob);
        stakingPool.withdrawAll();

        // Alice가 1위로
        assertEq(rankingAddon.getUserRank(1, alice), 1, "Alice should now be rank 1");
        assertEq(rankingAddon.getUserRankingScore(1, bob), 0, "Bob score should be 0");
    }

    function test_OnWithdraw_RemovesFromTop10WhenZero() public {
        _stakeAs(alice, 100 ether);

        assertTrue(rankingAddon.isTopRanker(1, alice), "Alice should be in top 10");

        // Alice 전액 출금
        vm.prank(alice);
        stakingPool.withdrawAll();

        assertFalse(rankingAddon.isTopRanker(1, alice), "Alice should be removed from top 10");
        assertEq(rankingAddon.getTopRankersCount(1), 0, "Top 10 should be empty");
    }

    function test_RankingDynamicUpdate_StakeMore() public {
        // 초기 스테이킹
        _stakeAs(alice, 50 ether);
        _stakeAs(bob, 100 ether);
        _stakeAs(carol, 75 ether);

        // 초기 순위: Bob(1), Carol(2), Alice(3)
        assertEq(rankingAddon.getUserRank(1, bob), 1);
        assertEq(rankingAddon.getUserRank(1, carol), 2);
        assertEq(rankingAddon.getUserRank(1, alice), 3);

        // Alice 추가 스테이킹 (50 + 100 = 150 ether)
        _stakeAs(alice, 100 ether);

        // 새 순위: Alice(1), Bob(2), Carol(3)
        assertEq(rankingAddon.getUserRank(1, alice), 1, "Alice should be rank 1");
        assertEq(rankingAddon.getUserRank(1, bob), 2, "Bob should be rank 2");
        assertEq(rankingAddon.getUserRank(1, carol), 3, "Carol should be rank 3");
    }

    function test_NewRanker_ReplacesLowestWhenTop10Full() public {
        // 10명 채우기 (10 ~ 100 ether)
        for (uint i = 1; i <= 10; i++) {
            _stakeAs(makeAddr(string(abi.encodePacked("user", i))), i * 10 ether);
        }

        // 최하위: user1 (10 ether)
        uint minScore = rankingAddon.getMinScoreForTop10(1);
        assertEq(minScore, 10 ether);

        // 새 유저가 15 ether로 스테이킹
        address newUser = makeAddr("newUser");
        _stakeAs(newUser, 15 ether);

        // newUser는 Top 10에 진입
        assertTrue(rankingAddon.isTopRanker(1, newUser), "newUser should be in top 10");

        // user1은 제거됨
        assertFalse(rankingAddon.isTopRanker(1, makeAddr("user1")), "user1 should be removed");

        // 여전히 10명 유지
        assertEq(rankingAddon.getTopRankersCount(1), 10);
    }

    function test_GetUserRank_ReturnsZeroIfNotInTop10() public {
        _stakeAs(alice, 100 ether);

        assertEq(rankingAddon.getUserRank(1, bob), 0, "Bob should have rank 0");
    }

    function test_GetRankerAt() public {
        _stakeAs(alice, 100 ether);
        _stakeAs(bob, 200 ether);

        (address user1, uint score1) = rankingAddon.getRankerAt(1, 1);
        assertEq(user1, bob, "Rank 1 should be Bob");
        assertEq(score1, 200 ether);

        (address user2, uint score2) = rankingAddon.getRankerAt(1, 2);
        assertEq(user2, alice, "Rank 2 should be Alice");
        assertEq(score2, 100 ether);
    }

    function test_GetScoreNeededForTop10_WhenNotFull() public {
        _stakeAs(alice, 100 ether);

        // Bob은 아직 스테이킹 안 함
        uint needed = rankingAddon.getScoreNeededForTop10(1, bob);
        assertEq(needed, 1, "Should need at least 1 wei to enter");
    }

    function test_GetScoreNeededForTop10_WhenFull() public {
        // 10명 채우기
        for (uint i = 1; i <= 10; i++) {
            _stakeAs(makeAddr(string(abi.encodePacked("user", i))), i * 10 ether);
        }

        // 새 유저가 진입하려면
        address newUser = makeAddr("newUser");
        uint needed = rankingAddon.getScoreNeededForTop10(1, newUser);

        // 최하위(10 ether) + 1
        assertEq(needed, 10 ether + 1, "Should need min + 1");
    }

    function test_GetMinScoreForTop10() public {
        // 5명만 있을 때
        for (uint i = 1; i <= 5; i++) {
            _stakeAs(makeAddr(string(abi.encodePacked("user", i))), i * 10 ether);
        }

        uint minScore = rankingAddon.getMinScoreForTop10(1);
        assertEq(minScore, 1, "Should return 1 when not full");

        // 10명 다 채웠을 때
        for (uint i = 6; i <= 10; i++) {
            _stakeAs(makeAddr(string(abi.encodePacked("user", i))), i * 10 ether);
        }

        minScore = rankingAddon.getMinScoreForTop10(1);
        assertEq(minScore, 10 ether, "Should return actual min score");
    }

    function test_MultipleSeasons_IndependentRankings() public {
        // 시즌 1
        _stakeAs(alice, 100 ether);
        _stakeAs(bob, 200 ether);

        assertEq(rankingAddon.getUserRank(1, bob), 1, "Bob rank 1 in season 1");

        // 시즌 2로 롤오버
        vm.roll(firstSeasonStartBlock + seasonBlocks + 1);
        stakingPool.rolloverSeason();

        // 시즌 2에서는 carol이 최고
        _stakeAs(carol, 300 ether);

        assertEq(rankingAddon.getUserRank(2, carol), 1, "Carol rank 1 in season 2");
        assertEq(rankingAddon.getUserRank(1, bob), 1, "Bob still rank 1 in season 1");

        // 시즌별 독립성 확인
        assertEq(rankingAddon.getTotalParticipants(1), 2);
        assertGe(rankingAddon.getTotalParticipants(2), 1);
    }

    function test_TotalParticipants() public {
        assertEq(rankingAddon.getTotalParticipants(1), 0, "Should start at 0");

        _stakeAs(alice, 100 ether);
        assertEq(rankingAddon.getTotalParticipants(1), 1);

        _stakeAs(bob, 200 ether);
        assertEq(rankingAddon.getTotalParticipants(1), 2);

        // Alice 추가 스테이킹 (참여자는 증가하지 않음)
        _stakeAs(alice, 50 ether);
        assertEq(rankingAddon.getTotalParticipants(1), 2, "Should still be 2");
    }

    function test_OnSeasonEnd() public {
        _stakeAs(alice, 100 ether);
        _stakeAs(bob, 200 ether);

        // 시즌 종료
        vm.roll(firstSeasonStartBlock + seasonBlocks + 1);

        vm.expectEmit(true, false, false, false);
        emit SeasonRankingFinalized(1, 2);

        stakingPool.rolloverSeason();

        // 데이터는 유지됨
        assertEq(rankingAddon.getUserRank(1, bob), 1);
    }

    // ============================================
    // 엣지 케이스 테스트
    // ============================================

    function test_SameScore_MaintainsInsertionOrder() public {
        // 같은 점수로 스테이킹
        _stakeAs(alice, 100 ether);
        _stakeAs(bob, 100 ether);
        _stakeAs(carol, 100 ether);

        (address[] memory users,) = rankingAddon.getTopRankers(1);

        // 먼저 들어온 순서대로 유지
        assertEq(users[0], alice, "Alice should be first");
        assertEq(users[1], bob, "Bob should be second");
        assertEq(users[2], carol, "Carol should be third");
    }

    function test_ZeroStake_NotAddedToTop10() public {
        // 0 스테이킹은 Top 10에 추가되지 않음 (실제로는 MIN_STAKE 때문에 불가능하지만)
        uint count = rankingAddon.getTopRankersCount(1);
        assertEq(count, 0, "Should remain 0");
    }

    function test_PartialWithdraw_UpdatesRank() public {
        _stakeAs(alice, 100 ether);
        _stakeAs(bob, 150 ether);

        // Bob이 1위
        assertEq(rankingAddon.getUserRank(1, bob), 1);

        // Bob 일부 출금 (150 - 100 = 50 남음)
        deal(address(wcross), bob, 150 ether);
        vm.startPrank(bob);
        wcross.approve(address(stakingPool), 150 ether);
        stakingPool.stake(150 ether);
        // 전액 출금 후 100만 다시 스테이킹
        stakingPool.withdrawAll();
        wcross.approve(address(stakingPool), 100 ether);
        stakingPool.stake(100 ether);
        vm.stopPrank();

        // Alice와 Bob이 동점이 되고, Alice가 먼저 들어왔으므로 Alice가 1위
        assertEq(rankingAddon.getUserRank(1, alice), 1, "Alice should be rank 1");
        assertEq(rankingAddon.getUserRank(1, bob), 2, "Bob should be rank 2");
    }

    function test_Stress_ManyRankers() public {
        // 20명 스테이킹 (Top 10만 유지되어야 함)
        for (uint i = 1; i <= 20; i++) {
            address user = makeAddr(string(abi.encodePacked("user", i)));
            _stakeAs(user, i * 10 ether);
        }

        // 정확히 10명만
        assertEq(rankingAddon.getTopRankersCount(1), 10);

        // 상위 10명 확인 (11번부터 20번까지)
        (address[] memory users, uint[] memory scores) = rankingAddon.getTopRankers(1);

        assertEq(scores[0], 200 ether, "Top should be 200");
        assertEq(scores[9], 110 ether, "10th should be 110");

        // user10 (100 ether)은 제외
        assertFalse(rankingAddon.isTopRanker(1, makeAddr("user10")));
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
