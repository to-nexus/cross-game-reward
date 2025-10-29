// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title SeasonLib
 * @notice Helper library for season management and validation
 * @dev Provides pure/view functions for season state checks and calculations
 *      Used across StakingPool contracts to maintain consistent season logic
 */
library SeasonLib {
    /// @notice Thrown when season time parameters are invalid (endTime <= startTime)
    error SeasonLibInvalidSeasonTime();

    /**
     * @notice Checks if a season is currently active
     * @param startTime Season start timestamp
     * @param endTime Season end timestamp
     * @param isFinalized Whether the season has been finalized
     * @return active True if season is active (current time is within range and not finalized)
     * @dev A season is active if:
     *      - Not finalized
     *      - Current timestamp >= startTime
     *      - Current timestamp <= endTime
     */
    function isSeasonActive(uint startTime, uint endTime, bool isFinalized) internal view returns (bool active) {
        if (isFinalized) return false;
        return block.timestamp >= startTime && block.timestamp <= endTime;
    }

    /**
     * @notice Checks if a season has ended
     * @param endTime Season end timestamp
     * @return ended True if current timestamp is past the season end time
     */
    function isSeasonEnded(uint endTime) internal view returns (bool ended) {
        return block.timestamp > endTime;
    }

    /**
     * @notice Validates that season timestamps are logically correct
     * @param startTime Season start timestamp
     * @param endTime Season end timestamp
     * @dev Reverts with SeasonLibInvalidSeasonTime if endTime <= startTime
     */
    function validateSeasonTime(uint startTime, uint endTime) internal pure {
        require(endTime > startTime, SeasonLibInvalidSeasonTime());
    }

    /**
     * @notice Checks if a specific timestamp falls within a season
     * @param timestamp Timestamp to check
     * @param startTime Season start timestamp
     * @param endTime Season end timestamp
     * @return inSeason True if timestamp is within [startTime, endTime] inclusive
     */
    function isTimeInSeason(uint timestamp, uint startTime, uint endTime) internal pure returns (bool inSeason) {
        return timestamp >= startTime && timestamp <= endTime;
    }

    /**
     * @notice Calculates the effective start time for point calculation
     * @param userJoinTime Timestamp when user joined/staked
     * @param seasonStartTime Season start timestamp
     * @return effectiveStart The later of userJoinTime and seasonStartTime
     * @dev Used to ensure points are only calculated from when both:
     *      1. The season started
     *      2. The user had a position
     *
     *      Example: If user staked at time 100 but season starts at time 200,
     *      points should only accumulate from time 200 onwards.
     */
    function calculateEffectiveStart(uint userJoinTime, uint seasonStartTime)
        internal
        pure
        returns (uint effectiveStart)
    {
        return userJoinTime > seasonStartTime ? userJoinTime : seasonStartTime;
    }
}
