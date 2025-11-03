// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

/**
 * @title MockERC20Permit
 * @notice Mock ERC20 token with EIP-2612 permit support for testing
 */
contract MockERC20Permit is ERC20Permit {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) ERC20Permit(name) {}

    function mint(address to, uint amount) external {
        _mint(to, amount);
    }

    function burn(address from, uint amount) external {
        _burn(from, amount);
    }
}

