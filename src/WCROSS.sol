// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {CrossStaking} from "./CrossStaking.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title WCROSS
 * @notice Wrapped CROSS Token (Native Token Wrapper)
 * @dev WETH9 스타일의 래핑 토큰
 *
 * 기능:
 * - 누구나 deposit/withdraw 가능
 * - CrossStaking이 화이트리스트 관리자로 등록됨
 * - 화이트리스트된 주소만 depositFor/withdrawFor 가능
 */
contract WCROSS is ERC20 {
    // ==================== 에러 ====================

    error WCROSSUnauthorized();
    error WCROSSInsufficientBalance();
    error WCROSSTransferFailed();

    // ==================== 상태 변수 ====================

    /// @notice CrossStaking 컨트랙트
    CrossStaking public staking;

    // ==================== 생성자 ====================

    constructor() ERC20("Wrapped CROSS", "WCROSS") {
        staking = CrossStaking(msg.sender);
    }

    // ==================== 수신 함수 ====================

    /**
     * @notice Native CROSS 수신 시 자동 래핑
     */
    receive() external payable {
        deposit();
    }

    // ==================== 사용자 함수 ====================

    /**
     * @notice CROSS를 WCROSS로 래핑
     */
    function deposit() public payable {
        require(msg.sender == staking.router(), WCROSSUnauthorized());
        require(msg.value > 0, WCROSSInsufficientBalance());
        _mint(msg.sender, msg.value);
    }

    /**
     * @notice WCROSS를 CROSS로 언래핑
     * @param amount 언래핑할 수량
     */
    function withdraw(uint amount) public {
        require(msg.sender == staking.router(), WCROSSUnauthorized());
        require(balanceOf(msg.sender) >= amount, WCROSSInsufficientBalance());
        _burn(msg.sender, amount);

        (bool success,) = msg.sender.call{value: amount}("");
        require(success, WCROSSTransferFailed());
    }
}
