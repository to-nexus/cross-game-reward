// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

interface ICrossStaking {
    struct PoolInfo {
        uint poolId;
        address poolAddress;
        address stakingToken;
        uint createdAt;
        bool active;
    }

    function wcross() external view returns (address);
    function router() external view returns (address);
    function poolImplementation() external view returns (address);
    function nextPoolId() external view returns (uint);

    function createPool(address stakingToken, uint48 initialDelay)
        external
        returns (uint poolId, address poolAddress);

    function addRewardToken(uint poolId, address rewardToken) external;
    function setPoolActive(uint poolId, bool active) external;
    function setRouter(address _router) external;

    function getPoolInfo(uint poolId) external view returns (PoolInfo memory);
    function poolAt(uint index) external view returns (uint);
    function getPoolCountByStakingToken(address stakingToken) external view returns (uint);
    function getPoolIdsByStakingToken(address stakingToken) external view returns (uint[] memory);
    function poolByStakingTokenAt(address stakingToken, uint index) external view returns (uint);
    function getTotalPoolCount() external view returns (uint);
    function getAllPoolIds() external view returns (uint[] memory);
    function getPoolAddress(uint poolId) external view returns (address);
    function getPoolId(address poolAddress) external view returns (uint);
    function getActivePoolIds() external view returns (uint[] memory);
}
