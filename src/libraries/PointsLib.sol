// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title PointsLib
 * @notice Pure logic library for points calculation
 * @dev Provides gas-efficient point calculation functions
 *      All calculations use fixed-point arithmetic with 6 decimal precision (1e6)
 */
library PointsLib {
    /// @notice Points calculation precision: 6 decimal places (1,000,000)
    /// @dev Used to maintain fractional precision in Solidity's integer arithmetic
    uint public constant POINTS_PRECISION = 1e6;

    /**
     * @notice Calculates points earned based on staking amount and time
     * @param balance Amount staked
     * @param fromBlock Starting block number
     * @param toBlock Ending block number
     * @param blockTime Block time in seconds (e.g., 1 for 1 second per block)
     * @param timeUnit Time unit in seconds for point calculation (e.g., 3600 for 1 hour)
     * @return points Calculated points with POINTS_PRECISION
     * @dev Formula: points = (balance × timeElapsed × POINTS_PRECISION) / timeUnit
     *
     *      Overflow Safety Analysis:
     *      - Uses Solidity 0.8+ built-in overflow checks
     *      - Maximum calculation: balance(~10^27) × timeElapsed(~10^9) × PRECISION(10^6) = ~10^42
     *      - uint256 max value: ~10^77, providing 10^35 safety margin
     *      - Operation order: multiplication first (preserves precision), division last
     *
     *      Example: 100 tokens staked for 3600 seconds with 1-hour time unit
     *      = (100 × 3600 × 1e6) / 3600 = 100,000,000 (100 points in raw format)
     */
    function calculatePoints(uint balance, uint fromBlock, uint toBlock, uint blockTime, uint timeUnit)
        internal
        pure
        returns (uint points)
    {
        if (fromBlock >= toBlock || balance == 0) return 0;

        uint blockElapsed = toBlock - fromBlock;
        uint timeElapsed = blockElapsed * blockTime;

        return (balance * timeElapsed * POINTS_PRECISION) / timeUnit;
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
