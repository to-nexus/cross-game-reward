// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title ICrossStaking
 * @notice Interface for the CrossStaking factory contract
 * @dev Defines the structure and functions for managing multiple staking pools
 */
interface ICrossStaking {
    /**
     * @notice Pool information structure
     * @param poolId Unique identifier for the pool
     * @param poolAddress Address of the pool contract
     * @param stakingToken Address of the token that can be staked
     * @param createdAt Timestamp when the pool was created
     * @param active Whether the pool is currently active
     */
    struct PoolInfo {
        uint poolId;
        address poolAddress;
        address stakingToken;
        uint createdAt;
        bool active;
    }

    /// @notice Returns the address of the WCROSS token
    function wcross() external view returns (address);

    /// @notice Returns the address of the router contract
    function router() external view returns (address);

    /// @notice Returns the address of the pool implementation
    function poolImplementation() external view returns (address);

    /// @notice Returns the next pool ID to be assigned
    function nextPoolId() external view returns (uint);

    /// @notice Creates a new staking pool
    function createPool(address stakingToken, uint minStakeAmount)
        external
        returns (uint poolId, address poolAddress);

    /// @notice Adds a reward token to a pool
    function addRewardToken(uint poolId, address rewardToken) external;

    /// @notice Removes a reward token from a pool
    function removeRewardToken(uint poolId, address rewardToken) external;

    /// @notice Sets the active status of a pool
    function setPoolActive(uint poolId, bool active) external;

    /// @notice Sets the router address
    function setRouter(address _router) external;

    /// @notice Retrieves pool information by pool ID
    function getPoolInfo(uint poolId) external view returns (PoolInfo memory);

    /// @notice Retrieves pool ID by index
    function poolAt(uint index) external view returns (uint);

    /// @notice Retrieves the number of pools for a specific staking token
    function getPoolCountByStakingToken(address stakingToken) external view returns (uint);

    /// @notice Retrieves all pool IDs for a specific staking token
    function getPoolIdsByStakingToken(address stakingToken) external view returns (uint[] memory);

    /// @notice Retrieves pool ID by staking token and index
    function poolByStakingTokenAt(address stakingToken, uint index) external view returns (uint);

    /// @notice Retrieves the total number of pools
    function getTotalPoolCount() external view returns (uint);

    /// @notice Retrieves all pool IDs
    function getAllPoolIds() external view returns (uint[] memory);

    /// @notice Retrieves pool address by pool ID
    function getPoolAddress(uint poolId) external view returns (address);

    /// @notice Retrieves pool ID by pool address
    function getPoolId(address poolAddress) external view returns (uint);

    /// @notice Retrieves only active pool IDs
    function getActivePoolIds() external view returns (uint[] memory);
}
