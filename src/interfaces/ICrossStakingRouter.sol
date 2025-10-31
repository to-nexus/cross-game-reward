// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

interface ICrossStakingRouter {
    function stakeNative(uint poolId) external payable;
    function unstakeNative(uint poolId) external;
    
    function stakeERC20(uint poolId, uint amount) external;
    function unstakeERC20(uint poolId) external;
    
    function getUserStakingInfo(uint poolId, address user) 
        external 
        view 
        returns (uint stakedAmount, uint[] memory pendingRewards);
    
    function isNativePool(uint poolId) external view returns (bool);
}

