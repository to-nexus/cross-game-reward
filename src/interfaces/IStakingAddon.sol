// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

/**
 * @title IStakingAddon
 * @notice 스테이킹 풀에 연결 가능한 애드온 인터페이스
 * @dev 랭킹, 배지, 통계 등의 부가 기능을 독립 컨트랙트로 구현
 */
interface IStakingAddon {
    /**
     * @notice 스테이킹 시 호출
     * @param user 사용자 주소
     * @param amount 스테이킹 금액
     * @param oldBalance 이전 잔액
     * @param newBalance 새 잔액
     * @param season 현재 시즌
     */
    function onStake(address user, uint amount, uint oldBalance, uint newBalance, uint season) external;

    /**
     * @notice 출금 시 호출
     * @param user 사용자 주소
     * @param amount 출금 금액
     * @param season 현재 시즌
     */
    function onWithdraw(address user, uint amount, uint season) external;

    /**
     * @notice 시즌 종료 시 호출
     * @param season 종료된 시즌 번호
     * @param totalStaked 총 스테이킹 금액
     * @param totalPoints 총 포인트
     */
    function onSeasonEnd(uint season, uint totalStaked, uint totalPoints) external;

    /**
     * @notice 보상 청구 시 호출
     * @param user 사용자 주소
     * @param season 시즌 번호
     * @param points 포인트
     * @param rewardAmount 보상 금액
     */
    function onClaim(address user, uint season, uint points, uint rewardAmount) external;
}
