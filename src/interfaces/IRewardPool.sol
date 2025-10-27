// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title IRewardPoolCode
 * @notice RewardPool의 creation code를 반환하는 인터페이스
 */
interface IRewardPoolCode {
    function code() external pure returns (bytes memory);
}

/**
 * @title IRewardPool
 * @notice RewardPool 컨트랙트의 인터페이스 (Season-based)
 */
interface IRewardPool {
    // ============ 이벤트 ============

    event SeasonFunded(uint indexed season, address indexed token, uint amount, uint actualReceived);
    event RewardPaid(address indexed user, uint indexed season, address indexed token, uint amount);

    // ============ 시즌 기반 보상 ============

    function fundSeason(uint season, address token, uint amount) external;
    function payUser(address user, uint season, address token, uint userPoints, uint totalPoints) external;

    // ============ 조회 함수 ============

    function getRemainingRewards(uint season, address token) external view returns (uint);
    function getExpectedReward(address user, uint season, address token) external view returns (uint);

    // 시즌 보상 토큰 정보
    function getSeasonRewardTokens(uint season) external view returns (address[] memory tokens);
    function getSeasonTokenInfo(uint season, address token)
        external
        view
        returns (uint total, uint claimed, uint remaining);
    function getSeasonAllRewards(uint season)
        external
        view
        returns (address[] memory tokens, uint[] memory totals, uint[] memory claimeds, uint[] memory remainings);

    // ============ 편의 함수 ============

    function depositReward(address token, uint amount) external;
    function getUserPendingReward(address user, address token) external view returns (uint);

    // ============ 관리 함수 ============

    function sweep(address token, address to, uint amount) external;
}
