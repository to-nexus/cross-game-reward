// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title PointsLib
 * @notice Pure logic library for points calculation
 * @dev Provides gas-efficient point calculation functions
 *      Points = balance × time (1 token × 1 second = 1 point)
 */
library PointsLib {
    /// @notice Decimals for display purposes
    /// @dev Frontend divides raw points by 10^POINT_DECIMALS for display
    ///      Raw points are in wei-seconds, divide by 1e18 to get token-seconds
    uint public constant POINT_DECIMALS = 18;

    /**
     * @notice Calculates points earned based on staking amount and time
     * @param balance Amount staked (in wei)
     * @param fromTime Starting timestamp
     * @param toTime Ending timestamp
     * @return points Calculated points (balance × timeElapsed)
     * @dev Formula: points = balance × timeElapsed
     *
     *      This gives us "token-seconds" as the unit:
     *      - 1 token (1e18 wei) × 1 second = 1e18 points
     *      - 100 tokens × 1 hour (3600s) = 360,000e18 points
     *
     *      For display: divide by 1e18 to get human-readable token-seconds
     *
     *      Overflow Safety Analysis:
     *      - Uses Solidity 0.8+ built-in overflow checks
     *      - Maximum realistic: balance(~10^27) × time(~10^9) = ~10^36
     *      - uint256 max value: ~10^77, providing 10^41 safety margin
     *
     *      Example: 100 tokens staked for 1 hour (3600 seconds)
     *      = 100e18 × 3600 = 3.6e23 raw points
     *      = 360,000 token-seconds (display)
     */
    function calculatePoints(uint balance, uint fromTime, uint toTime) internal pure returns (uint points) {
        if (fromTime >= toTime || balance == 0) return 0;

        uint timeElapsed = toTime - fromTime;

        return balance * timeElapsed;
    }

    /**
     * @notice Calculates proportional distribution (for reward distribution)
     * @param userAmount User's amount (e.g., user points)
     * @param totalAmount Total amount (e.g., total points)
     * @param rewardAmount Total reward to distribute
     * @return userReward User's proportional share of the reward
     * @dev Formula: userReward = (rewardAmount × userAmount) / totalAmount
     *
     *      Overflow Safety:
     *      - Guaranteed: userAmount ≤ totalAmount
     *      - Therefore: rewardAmount × userAmount ≤ rewardAmount × totalAmount
     *      - Solidity 0.8+ automatic overflow protection ensures safety
     *
     *      Example: If total reward is 1000 tokens, user has 400 points out of 1000 total
     *      = (1000 × 400) / 1000 = 400 tokens
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
