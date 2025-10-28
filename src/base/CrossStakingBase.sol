// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControlDefaultAdminRules} from
    "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuardTransient} from "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

/**
 * @title CrossStakingBase
 * @notice Base abstract contract for all Cross Staking contracts
 * @dev Provides common functionality for all staking system contracts:
 *      - Access control with 3-day timelock for admin changes (AccessControlDefaultAdminRules)
 *      - Reentrancy protection using transient storage (ReentrancyGuardTransient/EIP-1153)
 *      - Input validation helpers
 *      - Safe ERC20 operations
 *
 *      Inheritance hierarchy:
 *      CrossStakingBase
 *       ├─ StakingPoolBase -> StakingPool
 *       └─ RewardPoolBase -> RewardPool
 */
abstract contract CrossStakingBase is AccessControlDefaultAdminRules, ReentrancyGuardTransient {
    using SafeERC20 for IERC20;

    // ============ Errors ============

    /// @notice Thrown when a zero address is provided where not allowed
    error CrossStakingBaseCanNotZeroAddress();

    /// @notice Thrown when a zero amount is provided where not allowed
    error CrossStakingBaseInvalidAmount();

    /// @notice Thrown when caller lacks required authorization
    error CrossStakingBaseNotAuthorized();

    /// @notice Thrown when attempting to set a value that's already been set
    error CrossStakingBaseAlreadySet();

    /// @notice Thrown when a parameter value is outside acceptable range
    error CrossStakingBaseInvalidParameter();

    // ============ Validation Helpers ============

    /**
     * @notice Validates that an address is not zero
     * @param addr Address to validate
     * @dev Reverts with CrossStakingBaseCanNotZeroAddress if addr is zero
     */
    function _validateAddress(address addr) internal pure {
        require(addr != address(0), CrossStakingBaseCanNotZeroAddress());
    }

    /**
     * @notice Validates that an amount is not zero
     * @param amount Amount to validate
     * @dev Reverts with CrossStakingBaseInvalidAmount if amount is zero
     */
    function _validateAmount(uint amount) internal pure {
        require(amount != 0, CrossStakingBaseInvalidAmount());
    }

    /**
     * @notice Validates multiple addresses at once
     * @param addrs Array of addresses to validate
     * @dev Reverts on first zero address encountered
     */
    function _validateAddresses(address[] memory addrs) internal pure {
        for (uint i = 0; i < addrs.length; i++) {
            _validateAddress(addrs[i]);
        }
    }

    /**
     * @notice Validates that a value is within an acceptable range
     * @param value Value to validate
     * @param min Minimum acceptable value (inclusive)
     * @param max Maximum acceptable value (inclusive)
     * @dev Reverts with CrossStakingBaseInvalidParameter if value is out of range
     */
    function _validateRange(uint value, uint min, uint max) internal pure {
        require(value >= min && value <= max, CrossStakingBaseInvalidParameter());
    }

    // ============ Constructor ============

    /**
     * @notice Initializes the base contract with access control
     * @param admin Initial admin address (receives DEFAULT_ADMIN_ROLE)
     * @dev Sets up 3-day timelock for admin role transfers via AccessControlDefaultAdminRules
     */
    constructor(address admin) AccessControlDefaultAdminRules(3 days, admin) {
        _validateAddress(admin);
    }
}
