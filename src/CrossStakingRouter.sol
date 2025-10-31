// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {CrossStaking} from "./CrossStaking.sol";
import {CrossStakingPool} from "./CrossStakingPool.sol";
import {IWCROSS} from "./interfaces/IWCROSS.sol";

import {ICrossStakingRouter} from "./interfaces/ICrossStakingRouter.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title CrossStakingRouter
 * @notice 사용자와 스테이킹 풀 사이의 인터페이스
 * @dev Native CROSS 래핑 및 스테이킹 처리
 *
 * 주요 기능:
 * - Native CROSS를 WCROSS로 래핑 후 스테이킹
 * - WCROSS 언스테이킹 후 Native CROSS로 반환
 * - 일반 ERC20 토큰 스테이킹도 지원
 */
contract CrossStakingRouter is ICrossStakingRouter {
    using SafeERC20 for IERC20;

    // ==================== 에러 ====================

    error CSRInvalidAmount();
    error CSRCanNotZeroAddress();
    error CSRTransferFailed();
    error CSRPoolNotFound();
    error CSRNotWCROSSPool();
    error CSRNoStakeFound();

    // ==================== 이벤트 ====================

    event StakedNative(address indexed user, uint indexed poolId, uint amount);
    event UnstakedNative(address indexed user, uint indexed poolId, uint amount);
    event StakedERC20(address indexed user, uint indexed poolId, address token, uint amount);
    event UnstakedERC20(address indexed user, uint indexed poolId, address token, uint amount);

    // ==================== 상태 변수 ====================

    /// @notice CrossStaking 컨트랙트
    CrossStaking public immutable crossStaking;

    /// @notice WCROSS 토큰
    IWCROSS public immutable wcross;

    // ==================== 생성자 ====================

    constructor(address _crossStaking) {
        require(_crossStaking != address(0), CSRCanNotZeroAddress());

        crossStaking = CrossStaking(_crossStaking);
        wcross = IWCROSS(crossStaking.wcross());
    }

    // ==================== 수신 함수 ====================

    /**
     * @notice Native CROSS 수신
     */
    receive() external payable {}

    // ==================== Native CROSS 스테이킹 ====================

    /**
     * @notice Native CROSS 스테이킹
     * @param poolId 스테이킹할 풀 ID
     */
    function stakeNative(uint poolId) external payable {
        require(msg.value > 0, CSRInvalidAmount());

        CrossStakingPool pool = _getPoolAndValidateWCROSS(poolId);

        // Router가 Native CROSS를 WCROSS로 래핑
        wcross.deposit{value: msg.value}();

        // Router가 풀에 msg.sender를 위해 스테이킹
        IERC20(address(wcross)).forceApprove(address(pool), msg.value);
        pool.stakeFor(msg.sender, msg.value);

        emit StakedNative(msg.sender, poolId, msg.value);
    }

    /**
     * @notice Native CROSS 언스테이킹
     * @param poolId 언스테이킹할 풀 ID
     */
    function unstakeNative(uint poolId) external {
        CrossStakingPool pool = _getPoolAndValidateWCROSS(poolId);

        uint stakedAmount = pool.balances(msg.sender);
        require(stakedAmount > 0, CSRNoStakeFound());

        // Pool에서 msg.sender 언스테이킹 (WCROSS + 보상이 msg.sender에게 전송됨)
        pool.unstakeFor(msg.sender);

        // msg.sender의 WCROSS를 Router가 가져와서 언래핑
        uint wcrossBalance = IERC20(address(wcross)).balanceOf(msg.sender);
        require(wcrossBalance >= stakedAmount, CSRTransferFailed());

        IERC20(address(wcross)).safeTransferFrom(msg.sender, address(this), wcrossBalance);
        wcross.withdraw(wcrossBalance);

        // Native CROSS를 msg.sender에게 전송
        (bool success,) = msg.sender.call{value: wcrossBalance}("");
        require(success, CSRTransferFailed());

        emit UnstakedNative(msg.sender, poolId, stakedAmount);
    }

    // ==================== ERC20 스테이킹 (일반 토큰) ====================

    /**
     * @notice ERC20 토큰 스테이킹
     * @param poolId 스테이킹할 풀 ID
     * @param amount 스테이킹할 수량
     */
    function stakeERC20(uint poolId, uint amount) external {
        require(amount > 0, CSRInvalidAmount());

        CrossStakingPool pool = _getPool(poolId);
        IERC20 stakingToken = pool.stakingToken();

        // 토큰을 msg.sender에서 가져와서 풀에 스테이킹
        stakingToken.safeTransferFrom(msg.sender, address(this), amount);
        stakingToken.forceApprove(address(pool), amount);
        pool.stakeFor(msg.sender, amount);

        emit StakedERC20(msg.sender, poolId, address(stakingToken), amount);
    }

    /**
     * @notice ERC20 토큰 언스테이킹
     * @param poolId 언스테이킹할 풀 ID
     */
    function unstakeERC20(uint poolId) external {
        CrossStakingPool pool = _getPool(poolId);

        uint stakedAmount = pool.balances(msg.sender);
        require(stakedAmount > 0, CSRNoStakeFound());

        // 언스테이킹 (보상 포함)
        pool.unstakeFor(msg.sender);

        emit UnstakedERC20(msg.sender, poolId, address(pool.stakingToken()), stakedAmount);
    }

    // ==================== View 함수 ====================

    /**
     * @notice 사용자의 스테이킹 정보 조회
     * @param poolId 풀 ID
     * @param user 사용자 주소
     * @return stakedAmount 스테이킹된 수량
     * @return pendingRewards 대기 중인 보상 배열
     */
    function getUserStakingInfo(uint poolId, address user)
        external
        view
        returns (uint stakedAmount, uint[] memory pendingRewards)
    {
        CrossStakingPool pool = _getPool(poolId);
        stakedAmount = pool.balances(user);
        pendingRewards = pool.pendingRewards(user);
    }

    /**
     * @notice 풀이 Native CROSS 풀인지 확인
     * @param poolId 풀 ID
     * @return WCROSS 풀 여부
     */
    function isNativePool(uint poolId) external view returns (bool) {
        CrossStakingPool pool = _getPool(poolId);
        return address(pool.stakingToken()) == address(wcross);
    }

    // ==================== 내부 함수 ====================

    /**
     * @dev 풀 주소 조회 및 Pool 인스턴스 반환
     */
    function _getPool(uint poolId) internal view returns (CrossStakingPool) {
        return CrossStakingPool(crossStaking.getPoolAddress(poolId));
    }

    /**
     * @dev 풀 조회 및 WCROSS 풀 검증
     */
    function _getPoolAndValidateWCROSS(uint poolId) internal view returns (CrossStakingPool pool) {
        pool = _getPool(poolId);
        require(address(pool.stakingToken()) == address(wcross), CSRNotWCROSSPool());
    }
}
