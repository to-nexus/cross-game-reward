// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {ICrossGameRewardPool} from "./ICrossGameRewardPool.sol";

/**
 * @title ICrossGameRewardPoolV2
 * @notice Interface for the CrossGameRewardPoolV2 contract (Game Pool)
 * @dev Extends ICrossGameRewardPool with round-based reward distribution
 *      Uses OZ AccessControl for role management (SPONSOR_ROLE)
 *
 * Key differences from V1:
 * - Single reward token per pool (set at initialization)
 * - Round-based linear reward distribution over blocks
 * - Sponsor role (via AccessControl) for round management
 */
interface ICrossGameRewardPoolV2 is ICrossGameRewardPool {
    /**
     * @notice Round information structure for reward distribution
     * @param roundId Unique identifier for the round
     * @param creator Address of the sponsor who created this round
     * @param totalReward Total reward amount to be distributed
     * @param startBlock Block number when distribution starts
     * @param endBlock Block number when distribution ends
     * @param rewardPerBlock Reward amount distributed per block (truncated to 1e10 precision)
     * @param lastRewardBlock Last block number when rewards were calculated
     * @param accRewardPerShare Accumulated reward per share (scaled by PRECISION)
     * @param remainderReward Remaining reward added to the last block after truncation
     * @param isCancelled Whether the round has been cancelled
     */
    struct Round {
        uint256 roundId;
        address creator;
        uint256 totalReward;
        uint256 startBlock;
        uint256 endBlock;
        uint256 rewardPerBlock;
        uint256 lastRewardBlock;
        uint256 accRewardPerShare;
        uint256 remainderReward;
        bool isCancelled;
    }

    // ==================== Events ====================

    /// @notice Emitted when a new round is created
    event RoundCreated(
        uint256 indexed roundId,
        address indexed creator,
        uint256 totalReward,
        uint256 startBlock,
        uint256 endBlock,
        uint256 rewardPerBlock
    );

    /// @notice Emitted when a round is cancelled
    event RoundCancelled(uint256 indexed roundId, address indexed recipient, uint256 refundAmount);

    /// @notice Emitted when rounds are synced via emergency pagination
    event RoundsSynced(uint256 processed, uint256 removed);

    // ==================== View Functions ====================

    /// @notice Returns the sponsor role identifier
    function SPONSOR_ROLE() external view returns (bytes32);

    /// @notice Returns the single reward token for this pool
    function rewardToken() external view returns (IERC20);

    /// @notice Returns the next round ID to be assigned
    function nextRoundId() external view returns (uint256);

    /// @notice Returns the global accumulated reward per share
    function globalAccRewardPerShare() external view returns (uint256);

    /// @notice Returns the reclaimable amount (rewards distributed when totalDeposited was 0)
    function reclaimableAmount() external view returns (uint256);

    /// @notice Returns round information by round ID
    function getRound(uint256 roundId) external view returns (Round memory);

    /// @notice Returns all active (ongoing) rounds
    function getActiveRounds() external view returns (Round[] memory);

    /// @notice Returns the number of active rounds
    function getActiveRoundCount() external view returns (uint256);

    /// @notice Returns all active round IDs
    function getActiveRoundIds() external view returns (uint256[] memory);

    // ==================== Round Management Functions ====================

    /**
     * @notice Creates a new reward distribution round (tokens from caller)
     * @dev Only callable by accounts with SPONSOR_ROLE
     * @param amount Total reward amount to distribute
     * @param startBlock Block number when distribution starts (must be > current block)
     * @param durationBlocks Number of blocks over which to distribute rewards
     * @return roundId The ID of the created round
     */
    function createRound(uint256 amount, uint256 startBlock, uint256 durationBlocks)
        external
        returns (uint256 roundId);

    /**
     * @notice Creates a new reward distribution round (tokens from specified reserve)
     * @dev Only callable by accounts with SPONSOR_ROLE
     * @param reserve Address to transfer reward tokens from
     * @param amount Total reward amount to distribute
     * @param startBlock Block number when distribution starts (must be > current block)
     * @param durationBlocks Number of blocks over which to distribute rewards
     * @return roundId The ID of the created round
     */
    function createRoundFromReserve(address reserve, uint256 amount, uint256 startBlock, uint256 durationBlocks)
        external
        returns (uint256 roundId);

    /**
     * @notice Cancels a round that hasn't started yet (refund to caller)
     * @dev Only callable by the round creator
     * @param roundId The ID of the round to cancel
     */
    function cancelRound(uint256 roundId) external;

    /**
     * @notice Cancels a round that hasn't started yet (refund to specified recipient)
     * @dev Only callable by the round creator
     * @param roundId The ID of the round to cancel
     * @param recipient Address to receive the refund
     */
    function cancelRoundToRecipient(uint256 roundId, address recipient) external;

    /**
     * @notice Emergency paginated round sync for backlog resolution
     * @dev Callable by SPONSOR_ROLE holders or the factory owner.
     *      Processes up to maxRounds active rounds (0 = all).
     *      Identical state transitions to internal _updatePool().
     * @param maxRounds Maximum number of rounds to process (0 = unlimited)
     * @return processed Number of rounds iterated
     * @return removed Number of completed/cancelled rounds removed from active set
     */
    function syncRounds(uint256 maxRounds) external returns (uint256 processed, uint256 removed);
}
