// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

/**
 * @title IRewardPool
 * @notice RewardPool 컨트랙트의 인터페이스 (Season-based)
 */
interface IRewardPool {
    // ============ 이벤트 ============

    event SeasonFunded(uint indexed season, address indexed token, uint amount, address funder);
    event RewardPaid(address indexed user, uint indexed season, address indexed token, uint amount);

    // ============ 시즌 기반 보상 ============

    function fundSeason(uint season, address token, uint amount) external;
    function payUser(address user, uint season, address token, uint userPoints, uint totalPoints) external;

    // ============ 조회 함수 ============

    function getRemainingRewards(uint season, address token) external view returns (uint);
    function getExpectedReward(address user, uint season, address token) external view returns (uint);

    // ============ 편의 함수 ============

    function depositReward(address token, uint amount) external;
    function getUserPendingReward(address user, address token) external view returns (uint);
}
