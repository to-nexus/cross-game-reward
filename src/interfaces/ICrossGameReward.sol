// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {IERC5313} from "@openzeppelin/contracts/interfaces/IERC5313.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {ICrossGameRewardPool} from "./ICrossGameRewardPool.sol";
import {IWCROSS} from "./IWCROSS.sol";

/**
 * @title ICrossGameReward
 * @notice Interface for the CrossGameReward factory contract
 * @dev Defines the structure and functions for managing multiple game reward pools
 *      Implements IERC5313 for standard owner() function
 */
interface ICrossGameReward is IERC5313 {
    /**
     * @notice Pool type enum
     * @param CrossPool V1 pool - CROSS deposit, game token rewards (real-time distribution)
     * @param GamePool V2 pool - Game token deposit, CROSSD rewards (round-based distribution)
     */
    enum PoolType {
        CrossPool, // V1: CROSS -> Game Token
        GamePool // V2: Game Token -> CROSSD
    }

    /**
     * @notice Pool information structure
     * @param poolId Unique identifier for the pool
     * @param pool Address of the pool contract
     * @param name Name of the pool
     * @param depositToken Address of the token that can be deposited
     * @param createdAt Timestamp when the pool was created
     */
    struct PoolInfo {
        uint poolId;
        ICrossGameRewardPool pool;
        string name;
        IERC20 depositToken;
        uint createdAt;
    }

    /// @notice Returns the address of the WCROSS token
    function wcross() external view returns (IWCROSS);

    /// @notice Returns the address of the router contract
    function router() external view returns (address);

    /// @notice Returns the address of the pool implementation (V1)
    function poolImplementation() external view returns (ICrossGameRewardPool);

    /// @notice Returns the address of the pool implementation V2
    function poolImplementationV2() external view returns (ICrossGameRewardPool);

    /// @notice Returns the next pool ID to be assigned
    function nextPoolId() external view returns (uint);

    /// @notice Creates a new game reward pool (V1 - CrossPool)
    function createPool(string calldata name, IERC20 depositToken, uint minDepositAmount)
        external
        returns (uint poolId, ICrossGameRewardPool poolAddress);

    /// @notice Creates a new game reward pool V2 (GamePool)
    function createPoolV2(string calldata name, IERC20 depositToken, IERC20 rewardToken, uint minDepositAmount)
        external
        returns (uint poolId, ICrossGameRewardPool poolAddress);

    /// @notice Adds a reward token to a pool
    function addRewardToken(uint poolId, IERC20 token) external;

    /// @notice Removes a reward token from a pool
    function removeRewardToken(uint poolId, IERC20 token) external;

    /// @notice Sets the router address
    function setRouter(address _router) external;

    /// @notice Returns the pool type for a given pool ID
    function getPoolType(uint poolId) external view returns (PoolType);

    /// @notice Grants developer role to an account for a V2 pool
    function grantDeveloperRole(uint poolId, address developer) external;

    /// @notice Revokes developer role from an account for a V2 pool
    function revokeDeveloperRole(uint poolId, address developer) external;

    /// @notice Retrieves pool information by pool ID
    function getPoolInfo(uint poolId) external view returns (PoolInfo memory);

    /// @notice Retrieves pool ID by index
    function poolAt(uint index) external view returns (uint);

    /// @notice Retrieves the number of pools for a specific deposit token
    function getPoolCountByDepositToken(IERC20 token) external view returns (uint);

    /// @notice Retrieves all pool IDs for a specific deposit token
    function getPoolIdsByDepositToken(IERC20 token) external view returns (uint[] memory);

    /// @notice Retrieves pool ID by deposit token and index
    function poolByDepositTokenAt(IERC20 token, uint index) external view returns (uint);

    /// @notice Retrieves the total number of pools
    function getTotalPoolCount() external view returns (uint);

    /// @notice Retrieves all pool IDs
    function getAllPoolIds() external view returns (uint[] memory);

    /// @notice Retrieves pool address by pool ID
    function getPoolAddress(uint poolId) external view returns (ICrossGameRewardPool);

    /// @notice Retrieves pool ID by pool address
    function getPoolId(ICrossGameRewardPool pool) external view returns (uint);

    /// @notice Retrieves only active pool IDs
    function getActivePoolIds() external view returns (uint[] memory);

    /// @notice Retrieves pool IDs by pool type
    function getPoolIdsByType(PoolType poolType) external view returns (uint[] memory);
}
