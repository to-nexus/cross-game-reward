// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../interfaces/IStakingAddon.sol";

/**
 * @title RankingAddon
 * @notice 스테이킹 랭킹 집계 애드온
 * @dev IStakingAddon 인터페이스 구현
 *      가스 효율을 위해 정렬은 off-chain에서 수행 권장
 *      On-chain에서는 데이터만 기록하고, 조회 시 off-chain 정렬 사용
 */
contract RankingAddon is IStakingAddon {
    // ============================================
    // State Variables
    // ============================================

    /// @notice 연결된 StakingPool
    address public immutable stakingPool;

    /// @notice 시즌별 사용자별 랭킹 점수
    mapping(uint => mapping(address => uint)) public rankingScores;

    /// @notice 시즌별 Top 스테이커 목록 (간단한 구현)
    mapping(uint => address[]) public topStakers;
    mapping(uint => mapping(address => bool)) private isTopStaker;

    /// @notice 최대 Top 스테이커 수
    uint public constant MAX_TOP_STAKERS = 100;

    // ============================================
    // Events
    // ============================================

    event RankingUpdated(uint indexed season, address indexed user, uint score);
    event TopStakerAdded(uint indexed season, address indexed user);
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
     * @notice 스테이킹 시 랭킹 점수 업데이트
     * @dev 스테이킹 금액 * 시간 가중치 개념으로 점수 계산
     */
    function onStake(address user, uint amount, uint oldBalance, uint newBalance, uint season) external onlyPool {
        // 랭킹 점수 증가 (단순 예: 새 잔액 기준)
        rankingScores[season][user] = newBalance;

        // Top 스테이커 목록 업데이트
        _updateTopStakers(season, user, newBalance);

        emit RankingUpdated(season, user, newBalance);
    }

    /**
     * @notice 출금 시 랭킹 점수 업데이트
     */
    function onWithdraw(address user, uint amount, uint season) external onlyPool {
        // 출금 시 점수 감소 또는 제거
        if (rankingScores[season][user] >= amount) rankingScores[season][user] -= amount;
        else rankingScores[season][user] = 0;

        emit RankingUpdated(season, user, rankingScores[season][user]);
    }

    /**
     * @notice 시즌 종료 시 최종 랭킹 확정
     */
    function onSeasonEnd(uint season, uint totalStaked, uint totalPoints) external onlyPool {
        emit SeasonRankingFinalized(season, topStakers[season].length);
    }

    /**
     * @notice 보상 청구 시 (선택적 구현)
     */
    function onClaim(address user, uint season, uint points, uint rewardAmount) external onlyPool {
        // 필요 시 청구 기록 추가 가능
    }

    // ============================================
    // Internal Functions
    // ============================================

    /**
     * @notice Top 스테이커 목록 업데이트
     */
    function _updateTopStakers(uint season, address user, uint newBalance) private {
        address[] storage topList = topStakers[season];

        // 이미 Top 목록에 있으면 스킵
        if (isTopStaker[season][user]) return;

        // Top 목록이 아직 가득 차지 않았으면 추가
        if (topList.length < MAX_TOP_STAKERS) {
            topList.push(user);
            isTopStaker[season][user] = true;
            emit TopStakerAdded(season, user);
            return;
        }

        // 간단한 구현: 마지막 사용자와 비교 (실제로는 정렬 필요)
        // 실제 프로덕션에서는 MinHeap 등 효율적 자료구조 사용 권장
    }

    // ============================================
    // View Functions
    // ============================================

    /**
     * @notice 사용자 랭킹 점수 조회
     */
    function getUserRankingScore(uint season, address user) external view returns (uint) {
        return rankingScores[season][user];
    }

    /**
     * @notice Top 스테이커 목록 조회
     */
    function getTopStakers(uint season, uint offset, uint limit)
        external
        view
        returns (address[] memory stakers, uint[] memory scores)
    {
        address[] storage topList = topStakers[season];
        uint total = topList.length;

        if (offset >= total) return (new address[](0), new uint[](0));

        uint end = offset + limit;
        if (end > total) end = total;

        uint count = end - offset;
        stakers = new address[](count);
        scores = new uint[](count);

        for (uint i = 0; i < count; i++) {
            address staker = topList[offset + i];
            stakers[i] = staker;
            scores[i] = rankingScores[season][staker];
        }
    }

    /**
     * @notice 시즌별 총 참여자 수
     */
    function getTotalParticipants(uint season) external view returns (uint) {
        return topStakers[season].length;
    }
}
