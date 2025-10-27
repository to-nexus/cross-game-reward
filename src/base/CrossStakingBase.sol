// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

/**
 * @title CrossStakingBase
 * @notice 모든 Cross Staking 컨트랙트의 기본 추상 컨트랙트
 * @dev 공통 상속, 에러, 검증 로직 제공
 */
abstract contract CrossStakingBase is AccessControlDefaultAdminRules, ReentrancyGuardTransient {
    using SafeERC20 for IERC20;

    // ============================================
    // Common Errors
    // ============================================

    error CrossStakingBaseCanNotZeroAddress();
    error CrossStakingBaseInvalidAmount();
    error CrossStakingBaseNotAuthorized();
    error CrossStakingBaseAlreadySet();
    error CrossStakingBaseInvalidParameter();

    // ============================================
    // Common Validation Functions
    // ============================================

    /**
     * @notice 주소 검증
     */
    function _validateAddress(address addr) internal pure {
        require(addr != address(0), CrossStakingBaseCanNotZeroAddress());
    }

    /**
     * @notice 금액 검증
     */
    function _validateAmount(uint amount) internal pure {
        require(amount != 0, CrossStakingBaseInvalidAmount());
    }

    /**
     * @notice 여러 주소 일괄 검증
     */
    function _validateAddresses(address[] memory addrs) internal pure {
        for (uint i = 0; i < addrs.length; i++) {
            _validateAddress(addrs[i]);
        }
    }

    /**
     * @notice 파라미터 범위 검증
     */
    function _validateRange(uint value, uint min, uint max) internal pure {
        require(value >= min && value <= max, CrossStakingBaseInvalidParameter());
    }

    /**
     * @notice Constructor
     * @dev 3일 지연 시간으로 AccessControl 초기화
     */
    constructor(address admin) AccessControlDefaultAdminRules(3 days, admin) {
        _validateAddress(admin);
    }
}
