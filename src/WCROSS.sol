// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {CrossStaking} from "./CrossStaking.sol";
import {IWCROSS} from "./interfaces/IWCROSS.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title WCROSS
 * @notice Wrapped CROSS Token (Native Token Wrapper)
 * @dev WETH9-style wrapping token with router-only access control
 *
 * Features:
 * - Only the designated router can deposit/withdraw
 * - Native CROSS is automatically wrapped when received
 * - Maintains 1:1 parity with native CROSS
 */
contract WCROSS is ERC20, IWCROSS {
    // ==================== Errors ====================

    /// @notice Thrown when caller is not the authorized router
    error WCROSSUnauthorized();

    /// @notice Thrown when attempting to deposit/withdraw zero or insufficient amount
    error WCROSSInsufficientBalance();

    /// @notice Thrown when native CROSS transfer fails
    error WCROSSTransferFailed();

    // ==================== State Variables ====================

    /// @notice CrossStaking contract reference for router validation
    CrossStaking public staking;

    // ==================== Constructor ====================

    /**
     * @notice Initializes the WCROSS token
     * @dev Sets the deployer (CrossStaking) as the staking contract reference
     */
    constructor() ERC20("Wrapped CROSS", "WCROSS") {
        staking = CrossStaking(msg.sender);
    }

    // ==================== Receive Function ====================

    /**
     * @notice Automatically wraps native CROSS when received
     * @dev Calls deposit() to mint equivalent WCROSS
     */
    receive() external payable {
        deposit();
    }

    // ==================== Public Functions ====================

    /**
     * @notice Wraps native CROSS to WCROSS
     * @dev Only callable by the authorized router
     *      Mints WCROSS tokens equivalent to the native CROSS sent
     */
    function deposit() public payable {
        require(msg.sender == staking.router(), WCROSSUnauthorized());
        require(msg.value > 0, WCROSSInsufficientBalance());
        _mint(msg.sender, msg.value);
    }

    /**
     * @notice Unwraps WCROSS to native CROSS (sends to msg.sender)
     * @dev Only callable by the authorized router
     *      Burns WCROSS tokens and returns equivalent native CROSS
     * @param amount Amount of WCROSS to unwrap
     */
    function withdraw(uint amount) external {
        withdrawTo(amount, msg.sender);
    }

    /**
     * @notice Unwraps WCROSS to native CROSS and sends to specified address
     * @dev Only callable by the authorized router
     *      Burns WCROSS from msg.sender and sends native CROSS to recipient
     * @param amount Amount of WCROSS to unwrap
     * @param to Address to receive the unwrapped native CROSS
     */
    function withdrawTo(uint amount, address to) public {
        require(msg.sender == staking.router(), WCROSSUnauthorized());
        _burn(msg.sender, amount);

        (bool success,) = to.call{value: amount}("");
        require(success, WCROSSTransferFailed());
    }
}
