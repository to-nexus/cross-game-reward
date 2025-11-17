// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {CrossGameReward} from "./CrossGameReward.sol";
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

    /// @notice Thrown when native CROSS transfer fails
    error WCROSSTransferFailed();

    /// @notice Thrown when attempting to withdraw to an invalid address
    error WCROSSInvalidAddress();

    // ==================== Constructor ====================

    /**
     * @notice Initializes the WCROSS token
     * @dev Sets the deployer (CrossGameReward) as the game reward contract reference
     */
    constructor() ERC20("Wrapped CROSS", "WCROSS") {}

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
     *      Mints WCROSS tokens equivalent to the native CROSS sent
     */
    function deposit() public payable {
        if (msg.value != 0) _mint(msg.sender, msg.value);
    }

    /**
     * @notice Unwraps WCROSS to native CROSS (sends to msg.sender)
     *      Burns WCROSS tokens and returns equivalent native CROSS
     * @param amount Amount of WCROSS to unwrap
     */
    function withdraw(uint amount) external {
        withdrawTo(msg.sender, amount);
    }

    /**
     * @notice Unwraps WCROSS to native CROSS and sends to specified address
     *      Burns WCROSS from msg.sender and sends native CROSS to recipient
     * @param to Address to receive the unwrapped native CROSS
     * @param amount Amount of WCROSS to unwrap
     */
    function withdrawTo(address to, uint amount) public {
        require(to != address(0), WCROSSInvalidAddress());
        _burn(msg.sender, amount);

        (bool success,) = to.call{value: amount}("");
        require(success, WCROSSTransferFailed());
    }
}
