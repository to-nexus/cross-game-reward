// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title IWCROSS
 * @notice Interface for the WCROSS token contract
 * @dev Extends IERC20 with wrapping and unwrapping functions
 */
interface IWCROSS is IERC20 {
    /// @notice Wraps native CROSS to WCROSS
    function deposit() external payable;

    /// @notice Unwraps WCROSS to native CROSS
    function withdraw(uint amount) external;
}
