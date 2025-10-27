// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title WCROSS
 * @notice Wrapped CROSS - ERC20 wrapper for native CROSS token
 * @dev Native CROSS를 ERC20 토큰으로 래핑
 * @dev Transfer 이벤트는 ERC20 표준을 준수하여 deposit/withdraw 시에도 발생 (address(0) 사용)
 */
contract WCROSS is IERC20 {
    // ============ 상태 변수 ============

    string public constant name = "Wrapped CROSS";
    string public constant symbol = "WCROSS";
    uint8 public constant decimals = 18;

    mapping(address => uint) public balanceOf;
    mapping(address => mapping(address => uint)) public allowance;
    uint public totalSupply;

    // ============ 이벤트 ============

    event Deposit(address indexed dst, uint amount);
    event Withdrawal(address indexed src, uint amount);

    // ============ 입출금 함수 ============

    /**
     * @notice Native CROSS를 입금하고 WCROSS 발행
     */
    function deposit() public payable {
        balanceOf[msg.sender] += msg.value;
        totalSupply += msg.value;
        emit Deposit(msg.sender, msg.value);
    }

    /**
     * @notice WCROSS를 소각하고 Native CROSS 출금
     * @param amount 출금할 수량
     */
    function withdraw(uint amount) external {
        require(balanceOf[msg.sender] >= amount, "Insufficient balance");
        balanceOf[msg.sender] -= amount;
        totalSupply -= amount;
        (bool success,) = msg.sender.call{value: amount}("");
        require(success, "Transfer failed");
        emit Withdrawal(msg.sender, amount);
    }

    /**
     * @notice Native CROSS를 받을 수 있도록 설정
     */
    receive() external payable {
        deposit();
    }

    fallback() external payable {
        revert("Fallback not allowed");
    }

    // ============ ERC20 함수 ============

    function approve(address spender, uint amount) external returns (bool) {
        allowance[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    function transfer(address to, uint amount) external returns (bool) {
        return transferFrom(msg.sender, to, amount);
    }

    function transferFrom(address from, address to, uint amount) public returns (bool) {
        require(balanceOf[from] >= amount, "Insufficient balance");

        if (from != msg.sender) {
            require(allowance[from][msg.sender] >= amount, "Insufficient allowance");
            allowance[from][msg.sender] -= amount;
        }

        balanceOf[from] -= amount;
        balanceOf[to] += amount;

        emit Transfer(from, to, amount);
        return true;
    }
}
