// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../interfaces/IStakingAddon.sol";

/**
 * @title RankingAddon
 * @notice 완전 On-chain Top 10 랭킹 관리 애드온
 * @dev 서버 없는 완전한 탈중앙화 스테이킹 서비스를 위한 랭킹 시스템
 *      - Top 10 랭커만 on-chain에서 실시간 관리
 *      - 정렬된 상태 유지 (내림차순: 높은 점수 → 낮은 점수)
 *      - 가스 효율적인 삽입 정렬 알고리즘
 *      - 모든 랭킹 데이터 on-chain 저장 및 조회
 */
contract RankingAddon is IStakingAddon {
    // ============================================
    // Constants
    // ============================================

    /// @notice Top 랭커 최대 인원 (탈중앙화를 위해 Top 10만 관리)
    uint public constant MAX_TOP_RANKERS = 10;

    // ============================================
    // Structs
    // ============================================

    /// @notice 랭커 정보
    struct Ranker {
        address user;
        uint score;
    }

    // ============================================
    // State Variables
    // ============================================

    /// @notice 연결된 StakingPool
    address public immutable stakingPool;

    /// @notice 시즌별 모든 유저 랭킹 점수
    mapping(uint => mapping(address => uint)) public rankingScores;

    /// @notice 시즌별 Top 10 랭커 배열 (정렬된 상태: 내림차순)
    mapping(uint => Ranker[]) private topRankers;

    /// @notice 시즌별 유저가 Top 10에 포함되어 있는지 여부
    mapping(uint => mapping(address => bool)) public isTopRanker;

    /// @notice 시즌별 유저의 Top 10 내 인덱스 (isTopRanker가 true일 때만 유효)
    mapping(uint => mapping(address => uint)) private rankerIndex;

    /// @notice 시즌별 총 참여자 수
    mapping(uint => uint) public totalParticipants;

    // ============================================
    // Events
    // ============================================

    event RankingUpdated(uint indexed season, address indexed user, uint score);
    event TopRankerAdded(uint indexed season, address indexed user, uint rank, uint score);
    event TopRankerRemoved(uint indexed season, address indexed user, uint oldRank);
    event RankChanged(uint indexed season, address indexed user, uint oldRank, uint newRank);
    event SeasonRankingFinalized(uint indexed season, uint totalParticipants);

    // ============================================
    // Constructor
    // ============================================

    constructor(address _stakingPool) {
        require(_stakingPool != address(0), "Invalid pool");
        stakingPool = _stakingPool;
    }

    // ============================================
    // Modifier
    // ============================================

    modifier onlyPool() {
        require(msg.sender == stakingPool, "Only pool");
        _;
    }

    // ============================================
    // IStakingAddon Implementation
    // ============================================

    /**
     * @notice 스테이킹 시 랭킹 점수 업데이트 및 Top 10 갱신
     * @dev 실시간으로 Top 10 정렬 상태 유지
     */
    function onStake(address user, uint, /* amount */ uint, /* oldBalance */ uint newBalance, uint season)
        external
        onlyPool
    {
        // 첫 스테이킹이면 참여자 수 증가
        if (rankingScores[season][user] == 0) totalParticipants[season]++;

        // 랭킹 점수 업데이트 (현재 잔액 기준)
        uint oldScore = rankingScores[season][user];
        rankingScores[season][user] = newBalance;

        // Top 10 업데이트
        _updateTopRanking(season, user, oldScore, newBalance);

        emit RankingUpdated(season, user, newBalance);
    }

    /**
     * @notice 출금 시 랭킹 점수 감소 및 Top 10 갱신
     */
    function onWithdraw(address user, uint amount, uint season) external onlyPool {
        uint oldScore = rankingScores[season][user];
        uint newScore = oldScore >= amount ? oldScore - amount : 0;

        rankingScores[season][user] = newScore;

        // Top 10 업데이트
        _updateTopRanking(season, user, oldScore, newScore);

        emit RankingUpdated(season, user, newScore);
    }

    /**
     * @notice 시즌 종료 시 최종 랭킹 확정
     */
    function onSeasonEnd(uint season, uint, /* totalStaked */ uint /* totalPoints */ ) external onlyPool {
        emit SeasonRankingFinalized(season, totalParticipants[season]);
    }

    /**
     * @notice 보상 청구 시 (구현 선택사항)
     */
    function onClaim(address, /* user */ uint, /* season */ uint, /* points */ uint /* rewardAmount */ )
        external
        onlyPool
    {
        // Top 10 랭킹에는 영향 없음
    }

    // ============================================
    // Internal Functions - Top 10 관리
    // ============================================

    /**
     * @notice Top 10 랭킹 업데이트 (핵심 알고리즘)
     * @dev 가스 효율적인 삽입/삭제/갱신 로직
     */
    function _updateTopRanking(uint season, address user, uint oldScore, uint newScore) private {
        Ranker[] storage rankers = topRankers[season];
        bool wasInTop10 = isTopRanker[season][user];

        // Case 1: 이미 Top 10에 있었고, 점수가 변경됨
        if (wasInTop10) _updateExistingRanker(season, user, oldScore, newScore);
        // Case 2: Top 10에 없었고, 새로 진입 가능성 체크
        else _tryAddToTop10(season, user, newScore);
    }

    /**
     * @notice 기존 Top 10 랭커의 점수 업데이트
     */
    function _updateExistingRanker(uint season, address user, uint oldScore, uint newScore) private {
        Ranker[] storage rankers = topRankers[season];
        uint currentIndex = rankerIndex[season][user];

        // 점수가 0이 되면 Top 10에서 제거
        if (newScore == 0) {
            _removeFromTop10(season, user, currentIndex);
            return;
        }

        // 점수 업데이트
        rankers[currentIndex].score = newScore;

        // 순위 재조정
        if (newScore > oldScore) {
            // 점수 증가: 앞으로 이동 (bubble up)
            _bubbleUp(season, currentIndex);
        } else if (newScore < oldScore) {
            // 점수 감소: 뒤로 이동 (bubble down)
            _bubbleDown(season, currentIndex);

            // Top 10 밖으로 밀려났는지 체크
            if (rankers.length > MAX_TOP_RANKERS) _removeLastRanker(season);
        }
    }

    /**
     * @notice 새로운 유저를 Top 10에 추가 시도
     */
    function _tryAddToTop10(uint season, address user, uint newScore) private {
        Ranker[] storage rankers = topRankers[season];

        // 점수가 0이면 추가하지 않음
        if (newScore == 0) return;

        // Top 10이 아직 다 차지 않았으면 무조건 추가
        if (rankers.length < MAX_TOP_RANKERS) {
            _insertRanker(season, user, newScore);
            return;
        }

        // Top 10이 다 찬 경우: 최하위 랭커보다 점수가 높으면 교체
        uint lastIndex = rankers.length - 1;
        if (newScore > rankers[lastIndex].score) {
            // 최하위 랭커 제거
            _removeLastRanker(season);
            // 새 랭커 삽입
            _insertRanker(season, user, newScore);
        }
    }

    /**
     * @notice 정렬된 위치에 랭커 삽입 (이진 탐색 활용)
     */
    function _insertRanker(uint season, address user, uint score) private {
        Ranker[] storage rankers = topRankers[season];

        // 삽입 위치 찾기 (내림차순)
        uint insertPos = _findInsertPosition(season, score);

        // 배열 끝에 추가
        rankers.push(Ranker({user: user, score: score}));

        // 삽입 위치로 이동 (뒤에서부터 swap)
        for (uint i = rankers.length - 1; i > insertPos; i--) {
            _swapRankers(season, i, i - 1);
        }

        // 메타데이터 업데이트
        isTopRanker[season][user] = true;
        rankerIndex[season][user] = insertPos;

        emit TopRankerAdded(season, user, insertPos + 1, score);
    }

    /**
     * @notice 삽입 위치 찾기 (선형 탐색, Top 10이라 충분히 효율적)
     */
    function _findInsertPosition(uint season, uint score) private view returns (uint) {
        Ranker[] storage rankers = topRankers[season];
        uint length = rankers.length;

        for (uint i = 0; i < length; i++) {
            if (score > rankers[i].score) return i;
        }

        return length;
    }

    /**
     * @notice 랭커를 앞으로 이동 (점수 증가 시)
     */
    function _bubbleUp(uint season, uint index) private {
        Ranker[] storage rankers = topRankers[season];

        while (index > 0 && rankers[index].score > rankers[index - 1].score) {
            _swapRankers(season, index, index - 1);
            index--;
        }
    }

    /**
     * @notice 랭커를 뒤로 이동 (점수 감소 시)
     */
    function _bubbleDown(uint season, uint index) private {
        Ranker[] storage rankers = topRankers[season];
        uint length = rankers.length;

        while (index < length - 1 && rankers[index].score < rankers[index + 1].score) {
            _swapRankers(season, index, index + 1);
            index++;
        }
    }

    /**
     * @notice 두 랭커 위치 교환
     */
    function _swapRankers(uint season, uint i, uint j) private {
        Ranker[] storage rankers = topRankers[season];

        // 데이터 교환
        Ranker memory temp = rankers[i];
        rankers[i] = rankers[j];
        rankers[j] = temp;

        // 인덱스 맵 업데이트
        rankerIndex[season][rankers[i].user] = i;
        rankerIndex[season][rankers[j].user] = j;
    }

    /**
     * @notice Top 10에서 특정 랭커 제거
     */
    function _removeFromTop10(uint season, address user, uint index) private {
        Ranker[] storage rankers = topRankers[season];
        uint lastIndex = rankers.length - 1;

        emit TopRankerRemoved(season, user, index + 1);

        // 마지막 요소가 아니면 마지막 요소로 교체
        if (index != lastIndex) {
            rankers[index] = rankers[lastIndex];
            rankerIndex[season][rankers[index].user] = index;
        }

        // 마지막 요소 제거
        rankers.pop();

        // 메타데이터 정리
        isTopRanker[season][user] = false;
        delete rankerIndex[season][user];

        // 교체 후 정렬 복구 (bubble down)
        if (index < rankers.length) _bubbleDown(season, index);
    }

    /**
     * @notice 최하위 랭커 제거
     */
    function _removeLastRanker(uint season) private {
        Ranker[] storage rankers = topRankers[season];
        if (rankers.length == 0) return;

        uint lastIndex = rankers.length - 1;
        address lastUser = rankers[lastIndex].user;

        emit TopRankerRemoved(season, lastUser, lastIndex + 1);

        // 메타데이터 정리
        isTopRanker[season][lastUser] = false;
        delete rankerIndex[season][lastUser];

        // 배열에서 제거
        rankers.pop();
    }

    // ============================================
    // View Functions - On-chain 조회
    // ============================================

    /**
     * @notice 유저 랭킹 점수 조회
     */
    function getUserRankingScore(uint season, address user) external view returns (uint) {
        return rankingScores[season][user];
    }

    /**
     * @notice Top 10 랭커 전체 조회 (정렬된 상태)
     */
    function getTopRankers(uint season) external view returns (address[] memory users, uint[] memory scores) {
        Ranker[] storage rankers = topRankers[season];
        uint length = rankers.length;

        users = new address[](length);
        scores = new uint[](length);

        for (uint i = 0; i < length; i++) {
            users[i] = rankers[i].user;
            scores[i] = rankers[i].score;
        }
    }

    /**
     * @notice 특정 유저의 순위 조회 (1-based, Top 10 밖이면 0 반환)
     */
    function getUserRank(uint season, address user) external view returns (uint rank) {
        if (!isTopRanker[season][user]) return 0;
        return rankerIndex[season][user] + 1; // 1-based ranking
    }

    /**
     * @notice 특정 순위의 유저 조회 (1-based)
     */
    function getRankerAt(uint season, uint rank) external view returns (address user, uint score) {
        require(rank > 0, "Rank must be > 0");
        Ranker[] storage rankers = topRankers[season];
        require(rank <= rankers.length, "Rank out of bounds");

        uint index = rank - 1;
        return (rankers[index].user, rankers[index].score);
    }

    /**
     * @notice 현재 Top 10 인원수 조회
     */
    function getTopRankersCount(uint season) external view returns (uint) {
        return topRankers[season].length;
    }

    /**
     * @notice 시즌별 총 참여자 수 조회
     */
    function getTotalParticipants(uint season) external view returns (uint) {
        return totalParticipants[season];
    }

    /**
     * @notice Top 10 진입 최소 점수 조회
     */
    function getMinScoreForTop10(uint season) external view returns (uint) {
        Ranker[] storage rankers = topRankers[season];
        if (rankers.length == 0) return 0;
        if (rankers.length < MAX_TOP_RANKERS) return 1; // 아직 여유 있음
        return rankers[rankers.length - 1].score;
    }

    /**
     * @notice 유저가 Top 10에 진입하려면 필요한 추가 점수 조회
     */
    function getScoreNeededForTop10(uint season, address user) external view returns (uint needed) {
        if (isTopRanker[season][user]) return 0; // 이미 Top 10

        Ranker[] storage rankers = topRankers[season];
        uint currentScore = rankingScores[season][user];

        // Top 10이 아직 다 차지 않음
        if (rankers.length < MAX_TOP_RANKERS) return currentScore > 0 ? 0 : 1; // 1점만 있어도 진입 가능

        // 최하위 점수보다 높아야 함
        uint minScore = rankers[rankers.length - 1].score;
        if (currentScore > minScore) return 0; // 이미 충분
        return minScore - currentScore + 1; // 1점 더 필요
    }
}
