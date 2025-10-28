// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title SeasonLib
 * @notice Helper library for season management and validation
 * @dev Provides pure/view functions for season state checks and calculations
 *      Used across StakingPool contracts to maintain consistent season logic
 */
library SeasonLib {
    /// @notice Thrown when season block parameters are invalid (endBlock <= startBlock)
    error SeasonLibInvalidSeasonBlocks();

    /**
     * @notice Checks if a season is currently active
     * @param startBlock Season start block number
     * @param endBlock Season end block number
     * @param isFinalized Whether the season has been finalized
     * @return active True if season is active (current block is within range and not finalized)
     * @dev A season is active if:
     *      - Not finalized
     *      - Current block >= startBlock
     *      - Current block <= endBlock
     */
    function isSeasonActive(uint startBlock, uint endBlock, bool isFinalized) internal view returns (bool active) {
        if (isFinalized) return false;
        return block.number >= startBlock && block.number <= endBlock;
    }

    /**
     * @notice Checks if a season has ended
     * @param endBlock Season end block number
     * @return ended True if current block is past the season end block
     */
    function isSeasonEnded(uint endBlock) internal view returns (bool ended) {
        return block.number > endBlock;
    }

    /**
     * @notice Validates that season block numbers are logically correct
     * @param startBlock Season start block number
     * @param endBlock Season end block number
     * @dev Reverts with SeasonLibInvalidSeasonBlocks if endBlock <= startBlock
     */
    function validateSeasonBlocks(uint startBlock, uint endBlock) internal pure {
        require(endBlock > startBlock, SeasonLibInvalidSeasonBlocks());
    }

    /**
     * @notice Checks if a specific block number falls within a season
     * @param blockNumber Block number to check
     * @param startBlock Season start block number
     * @param endBlock Season end block number
     * @return inSeason True if blockNumber is within [startBlock, endBlock] inclusive
     */
    function isBlockInSeason(uint blockNumber, uint startBlock, uint endBlock) internal pure returns (bool inSeason) {
        return blockNumber >= startBlock && blockNumber <= endBlock;
    }

    /**
     * @notice Calculates the effective start block for point calculation
     * @param userJoinBlock Block when user joined/staked
     * @param seasonStartBlock Season start block
     * @return effectiveStart The later of userJoinBlock and seasonStartBlock
     * @dev Used to ensure points are only calculated from when both:
     *      1. The season started
     *      2. The user had a position
     *
     *      Example: If user staked at block 100 but season starts at block 200,
     *      points should only accumulate from block 200 onwards.
     */
    function calculateEffectiveStart(uint userJoinBlock, uint seasonStartBlock)
        internal
        pure
        returns (uint effectiveStart)
    {
        return userJoinBlock > seasonStartBlock ? userJoinBlock : seasonStartBlock;
    }
}
