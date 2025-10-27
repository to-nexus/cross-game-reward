// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/RewardPoolBase.sol";

import "./interfaces/IStakingPool.sol";
import "./interfaces/IStakingProtocol.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title RewardPoolCode
 * @notice RewardPool의 creation code를 반환하는 컨트랙트
 */
contract RewardPoolCode {
    function code() external pure returns (bytes memory) {
        return type(RewardPool).creationCode;
    }
}

/**
 * @title RewardPool
 * @notice 시즌 기반 보상 토큰 관리 및 유저 클레임 처리 (Base 상속 버전)
 * @dev RewardPoolBase를 상속하여 기본 보상 구조 구현
 */
contract RewardPool is RewardPoolBase {
    using SafeERC20 for IERC20;

    // ============================================
    // Errors
    // ============================================

    error RewardPoolOnlyProtocol();

    // ============================================
    // Events
    // ============================================

    event TokensSwept(address indexed token, address indexed to, uint amount, uint balanceBefore, uint balanceAfter);

    // ============================================
    // State Variables
    // ============================================

    /// @notice 연결된 StakingPool
    IStakingPool public immutable stakingPool;

    /// @notice 프로토콜 컨트랙트
    IStakingProtocol public immutable protocol;

    // ============================================
    // Constructor
    // ============================================

    constructor(address _stakingPool, address _protocol) RewardPoolBase(_protocol) {
        _validateAddress(_stakingPool);
        _validateAddress(_protocol);

        stakingPool = IStakingPool(_stakingPool);
        protocol = IStakingProtocol(_protocol);

        // StakingPool에 STAKING_POOL_ROLE 부여
        _grantRole(STAKING_POOL_ROLE, _stakingPool);
    }

    // ============================================
    // Override Hooks (추가보상 없는 기본 구현)
    // ============================================

    /**
     * @notice 추가 보상 계산 (기본 구현: 보너스 없음)
     * @dev 자식 컨트랙트에서 오버라이드하여 추가보상 구현 가능
     */
    function _calculateBonusReward(
        address, /*user*/
        uint, /*season*/
        address, /*token*/
        uint, /*userPoints*/
        uint, /*totalPoints*/
        uint /*baseReward*/
    ) internal virtual override returns (uint bonusAmount) {
        // 기본 구현: 추가 보상 없음
        return 0;
    }

    // ============================================
    // View Functions
    // ============================================

    /**
     * @notice 사용자가 받을 수 있는 예상 보상량
     */
    function getExpectedReward(address user, uint season, address token)
        external
        view
        override
        returns (uint expectedReward)
    {
        if (hasClaimedSeasonReward[user][season][token]) return 0;

        (uint userPoints, uint totalPoints) = stakingPool.getSeasonUserPoints(season, user);

        if (userPoints == 0 || totalPoints == 0) return 0;

        uint totalReward = seasonRewards[season][token];
        if (totalReward == 0) return 0;

        return (totalReward * userPoints) / totalPoints;
    }

    // ============================================
    // 편의 함수
    // ============================================

    /**
     * @notice 현재 시즌에 보상 예치 (사전 예치 지원)
     * @dev 표준 ERC20 토큰만 지원
     * @dev preDepositStartBlock 이후부터 예치 가능
     * @dev currentSeason = 0 (시즌 시작 전)이면 시즌 1에 예치
     */
    function depositReward(address token, uint amount) external {
        uint currentSeason = stakingPool.currentSeason();
        uint targetSeason = currentSeason;

        // 시즌 시작 전(currentSeason = 0)이면 시즌 1에 사전 예치
        if (currentSeason == 0) {
            targetSeason = 1;

            // preDepositStartBlock 체크
            uint preDepositStart = stakingPool.preDepositStartBlock();
            require(preDepositStart == 0 || block.number >= preDepositStart, "Pre-deposit not yet available");
        }

        // fundSeason을 호출하여 토큰 리스트도 자동으로 업데이트
        fundSeason(targetSeason, token, amount);
    }

    /**
     * @notice 사용자의 현재 시즌 예상 보상 조회
     */
    function getUserPendingReward(address user, address token) external view returns (uint) {
        uint currentSeason = stakingPool.currentSeason();
        return this.getExpectedReward(user, currentSeason, token);
    }

    /**
     * @notice 시즌별 토큰 요약 정보 조회
     */
    function getSeasonSummary(uint season, address[] calldata tokens)
        external
        view
        returns (uint[] memory totals, uint[] memory claimed, uint[] memory remaining)
    {
        uint len = tokens.length;
        totals = new uint[](len);
        claimed = new uint[](len);
        remaining = new uint[](len);

        for (uint i = 0; i < len; i++) {
            totals[i] = seasonRewards[season][tokens[i]];
            claimed[i] = seasonClaimed[season][tokens[i]];
            remaining[i] = totals[i] - claimed[i];
        }
    }

    /**
     * @notice 사용자의 시즌별 토큰 청구 여부 조회
     */
    function getUserClaimed(address user, uint season, address token) external view returns (bool) {
        return hasClaimedSeasonReward[user][season][token];
    }

    // ============================================
    // 관리 함수
    // ============================================

    /**
     * @notice 잘못 전송된 토큰 회수
     */
    function sweep(address token, address to, uint amount) external {
        require(msg.sender == address(protocol), RewardPoolOnlyProtocol());
        _validateAddress(to);
        _validateAmount(amount);

        uint balanceBefore = IERC20(token).balanceOf(address(this));
        IERC20(token).safeTransfer(to, amount);
        uint balanceAfter = IERC20(token).balanceOf(address(this));

        emit TokensSwept(token, to, amount, balanceBefore, balanceAfter);
    }
}
