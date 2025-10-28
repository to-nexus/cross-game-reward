// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title WCROSS
 * @notice Wrapped CROSS - ERC20 wrapper for native CROSS token
 * @dev Standard WETH-style implementation with 1:1 wrapping ratio
 *
 *      Key features:
 *      - Wrap: Send native CROSS to receive WCROSS ERC20 tokens
 *      - Unwrap: Burn WCROSS to receive native CROSS
 *      - Full ERC20 compatibility
 *      - Automatic wrapping via receive() function
 *
 *      Use in Cross Staking:
 *      - StakingPools accept WCROSS (ERC20)
 *      - StakingRouter auto-wraps native CROSS for user convenience
 *      - Users can also interact with WCROSS directly
 */
contract WCROSS is IERC20 {
    // ============ Errors ============

    /// @notice Thrown when user has insufficient WCROSS balance
    error WCROSSInsufficientBalance();

    /// @notice Thrown when spender has insufficient allowance
    error WCROSSInsufficientAllowance();

    /// @notice Thrown when native CROSS transfer fails
    error WCROSSTransferFailed();

    /// @notice Thrown when fallback is called (use receive or deposit)
    error WCROSSFallbackNotAllowed();

    // ============ ERC20 Metadata ============

    /// @notice Token name
    string public constant name = "Wrapped CROSS";

    /// @notice Token symbol
    string public constant symbol = "WCROSS";

    /// @notice Token decimals (matches native CROSS)
    uint8 public constant decimals = 18;

    // ============ State Variables ============

    /// @notice Maps addresses to their WCROSS balances
    mapping(address => uint) public balanceOf;

    /// @notice Maps owner => spender => allowance
    mapping(address => mapping(address => uint)) public allowance;

    /// @notice Total supply of WCROSS (equals native CROSS held by contract)
    uint public totalSupply;

    // ============ Events ============

    /// @notice Emitted when native CROSS is deposited (wrapped)
    event Deposit(address indexed dst, uint amount);

    /// @notice Emitted when WCROSS is withdrawn (unwrapped) to native CROSS
    event Withdrawal(address indexed src, uint amount);

    // ============ Wrapping Functions ============

    /**
     * @notice Wraps native CROSS into WCROSS
     * @dev Mints WCROSS equal to msg.value and adds to sender's balance
     *      Can also send native CROSS directly to contract (triggers receive)
     */
    function deposit() public payable {
        balanceOf[msg.sender] += msg.value;
        totalSupply += msg.value;
        emit Deposit(msg.sender, msg.value);
    }

    /**
     * @notice Unwraps WCROSS back to native CROSS
     * @param amount Amount of WCROSS to unwrap
     * @dev Burns WCROSS and sends equivalent native CROSS to caller
     */
    function withdraw(uint amount) external {
        require(balanceOf[msg.sender] >= amount, WCROSSInsufficientBalance());
        balanceOf[msg.sender] -= amount;
        totalSupply -= amount;
        (bool success,) = msg.sender.call{value: amount}("");
        require(success, WCROSSTransferFailed());
        emit Withdrawal(msg.sender, amount);
    }

    /**
     * @notice Automatically wraps native CROSS sent to contract
     * @dev Sending native CROSS directly to this contract will mint WCROSS
     */
    receive() external payable {
        deposit();
    }

    /**
     * @notice Fallback function (reverts to prevent accidental calls)
     * @dev Use deposit() or send native CROSS to trigger receive()
     */
    fallback() external payable {
        revert WCROSSFallbackNotAllowed();
    }

    // ============ ERC20 Functions ============

    /**
     * @notice Approves spender to spend tokens on behalf of caller
     * @param spender Address to approve
     * @param amount Amount to approve
     * @return True on success
     */
    function approve(address spender, uint amount) external returns (bool) {
        allowance[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    /**
     * @notice Transfers tokens to another address
     * @param to Recipient address
     * @param amount Amount to transfer
     * @return True on success
     */
    function transfer(address to, uint amount) external returns (bool) {
        return transferFrom(msg.sender, to, amount);
    }

    /**
     * @notice Transfers tokens from one address to another
     * @param from Sender address
     * @param to Recipient address
     * @param amount Amount to transfer
     * @return True on success
     * @dev Checks allowance if caller is not the sender
     */
    function transferFrom(address from, address to, uint amount) public returns (bool) {
        require(balanceOf[from] >= amount, WCROSSInsufficientBalance());

        if (from != msg.sender) {
            require(allowance[from][msg.sender] >= amount, WCROSSInsufficientAllowance());
            allowance[from][msg.sender] -= amount;
        }

        balanceOf[from] -= amount;
        balanceOf[to] += amount;

        emit Transfer(from, to, amount);
        return true;
    }
}
