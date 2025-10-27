// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./IRewardPool.sol";

/**
 * @title IStakingPoolCode
 * @notice StakingPool의 creation code를 반환하는 인터페이스
 */
interface IStakingPoolCode {
    function code() external pure returns (bytes memory);
}

/**
 * @title IStakingPool
 * @notice StakingPool 컨트랙트의 인터페이스
 */
interface IStakingPool {
    // ============ 스테이킹 함수 ============

    function stake(uint amount) external;
    function stakeFor(address user, uint amount) external;
    function withdrawAll() external;
    function withdrawAllFor(address user) external;

    // ============ 시즌 관리 ============

    function rolloverSeason() external;
    function claimSeason(uint season, address rewardToken) external;
    function claimSeasonFor(address user, uint season, address rewardToken) external;

    // ============ 포인트 ============

    function updatePoints(address user) external;

    // ============ 관리 함수 ============

    function setRewardPool(IRewardPool rewardPool) external;
    function setApprovedRouter(address router, bool approved) external;
    function setPointsTimeUnit(uint timeUnit) external;
    function setBlockTime(uint blockTime) external;
    function setNextSeasonStart(uint startBlock) external;
    function setPoolEndBlock(uint endBlock) external;
    function manualRolloverSeasons(uint maxRollovers) external returns (uint rolloversPerformed);

    // ============ 조회 함수 ============

    function getStakingPower(address user) external view returns (uint);
    function getTotalStakingPower() external view returns (uint);
    function getUserPoints(address user) external view returns (uint);
    function getStakePosition(address user) external view returns (uint balance, uint points, uint lastUpdateBlock);
    function getCurrentSeasonInfo()
        external
        view
        returns (uint season, uint startBlock, uint endBlock, uint blocksElapsed);
    function getSeasonUserPoints(uint season, address user) external view returns (uint userPoints, uint totalPoints);
    function seasonTotalPointsSnapshot(uint season) external view returns (uint);
    function currentSeason() external view returns (uint);
    function totalStaked() external view returns (uint);
    function getExpectedSeasonPoints(uint season, address user) external view returns (uint);
    function getExpectedSeasonReward(uint season, address user, address rewardToken) external view returns (uint);
    function isSeasonActive() external view returns (bool);
    function poolEndBlock() external view returns (uint);
    function nextSeasonStartBlock() external view returns (uint);
    function preDepositStartBlock() external view returns (uint);
    function getPendingSeasonRollovers() external view returns (uint pendingSeasons);

    // ============ 설정 조회 함수 ============

    function blockTime() external view returns (uint);
    function pointsTimeUnit() external view returns (uint);
    function seasonBlocks() external view returns (uint);
}
