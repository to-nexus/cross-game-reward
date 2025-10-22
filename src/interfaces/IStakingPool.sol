// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./IRewardPool.sol";
import "./IStakingAddon.sol";

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

    // ============ 포인트 ============

    function updatePoints(address user) external;

    // ============ 관리 함수 ============

    function setRewardPool(IRewardPool rewardPool) external;
    function setApprovedRouter(address router, bool approved) external;
    function setPointsTimeUnit(uint timeUnit) external;
    function setBlockTime(uint blockTime) external;
    function setNextSeasonStart(uint startBlock) external;
    function setPoolEndBlock(uint endBlock) external;
    function setStakingAddon(IStakingAddon addon) external;
    function setAddonApproved(IStakingAddon addon, bool approved) external;

    // ============ 조회 함수 ============

    function getStakingPower(address user) external view returns (uint);
    function getTotalStakingPower() external view returns (uint);
    function getUserPoints(address user) external view returns (uint);
    function getStakePosition(address user) external view returns (uint balance, uint points, uint lastUpdateBlock);
    function getCurrentSeasonInfo()
        external
        view
        returns (uint season, uint startBlock, uint endBlock, uint blocksElapsed);
    function getSeasonUserPoints(uint season, address user) external view returns (uint);
    function seasonTotalPointsSnapshot(uint season) external view returns (uint);
    function currentSeason() external view returns (uint);
    function totalStaked() external view returns (uint);
    function getExpectedSeasonPoints(uint season, address user) external view returns (uint);
    function getExpectedSeasonReward(uint season, address user, address rewardToken) external view returns (uint);
    function isSeasonActive() external view returns (bool);
    function poolEndBlock() external view returns (uint);
    function nextSeasonStartBlock() external view returns (uint);
}
