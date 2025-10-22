// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

/**
 * @title PointsLib
 * @notice 포인트 계산 순수 로직 라이브러리
 * @dev 가스 효율적인 포인트 계산 제공
 */
library PointsLib {
    uint public constant POINTS_PRECISION = 1e6; // 포인트 정밀도: 소수점 6자리

    /**
     * @notice 포인트 계산
     * @param balance 스테이킹 수량
     * @param fromBlock 시작 블록
     * @param toBlock 종료 블록
     * @param blockTime 블록 시간 (초)
     * @param timeUnit 시간 단위 (초)
     * @return points 계산된 포인트
     * @dev Overflow 안정성:
     *      - Solidity 0.8+ 내장 오버플로우 체크 사용
     *      - 최대 연산: balance(~10^27) * timeElapsed(~10^9) * PRECISION(10^6) = ~10^42
     *      - uint256 최대값: ~10^77 이므로 안전함
     *      - 연산 순서: 곱셈 먼저 수행하여 정밀도 유지, 마지막에 나눗셈
     */
    function calculatePoints(uint balance, uint fromBlock, uint toBlock, uint blockTime, uint timeUnit)
        internal
        pure
        returns (uint points)
    {
        if (fromBlock >= toBlock || balance == 0) return 0;

        uint blockElapsed = toBlock - fromBlock;
        uint timeElapsed = blockElapsed * blockTime;

        // 정밀도를 위해 곱셈을 먼저 수행 (Solidity 0.8+ 자동 오버플로우 체크)
        return (balance * timeElapsed * POINTS_PRECISION) / timeUnit;
    }

    /**
     * @notice 집계 기반 포인트 계산
     * @param totalStaked 총 스테이킹 금액
     * @param lastBlock 마지막 집계 블록
     * @param currentBlock 현재 블록
     * @param blockTime 블록 시간
     * @param timeUnit 시간 단위
     * @return additionalPoints 추가 포인트
     */
    function calculateAggregatedPoints(
        uint totalStaked,
        uint lastBlock,
        uint currentBlock,
        uint blockTime,
        uint timeUnit
    ) internal pure returns (uint additionalPoints) {
        if (lastBlock >= currentBlock || totalStaked == 0) return 0;
        return calculatePoints(totalStaked, lastBlock, currentBlock, blockTime, timeUnit);
    }

    /**
     * @notice 비율 계산 (보상 분배용)
     * @param userAmount 사용자 금액
     * @param totalAmount 총 금액
     * @param rewardAmount 보상 금액
     * @return userReward 사용자 보상
     * @dev Overflow 안정성:
     *      - userAmount ≤ totalAmount 보장됨
     *      - rewardAmount * userAmount는 rewardAmount * totalAmount보다 작음
     *      - Solidity 0.8+ 자동 오버플로우 체크로 안전함
     */
    function calculateProRata(uint userAmount, uint totalAmount, uint rewardAmount)
        internal
        pure
        returns (uint userReward)
    {
        if (totalAmount == 0 || userAmount == 0) return 0;
        return (rewardAmount * userAmount) / totalAmount;
    }
}
