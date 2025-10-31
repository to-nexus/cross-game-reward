// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface ICrossStakingPool {
    struct RewardToken {
        address tokenAddress;
        uint rewardPerTokenStored;
        uint lastBalance;
        uint removedDistributedAmount;
        bool isRemoved;
    }

    struct UserReward {
        uint rewardPerTokenPaid;
        uint rewards;
    }

    function stakingToken() external view returns (IERC20);
    function balances(address user) external view returns (uint);
    function totalStaked() external view returns (uint);
    function userRewards(address user, address token) external view returns (uint rewardPerTokenPaid, uint rewards);
    
    function stake(uint amount) external;
    function stakeFor(address account, uint amount) external;
    function unstake() external;
    function unstakeFor(address account) external;
    
    function claimRewards() external;
    function claimReward(address tokenAddress) external;
    
    function pendingRewards(address user) external view returns (uint[] memory);
    
    function rewardTokenAt(uint index) external view returns (address);
    function getRewardToken(address tokenAddress) external view returns (RewardToken memory);
    function isRewardToken(address tokenAddress) external view returns (bool);
    function getRewardTokens() external view returns (address[] memory);
    function rewardTokensLength() external view returns (uint);
    function rewardTokenCount() external view returns (uint);
}
