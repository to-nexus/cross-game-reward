// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

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

        uint userPoints = stakingPool.getSeasonUserPoints(season, user);
        uint totalPoints = stakingPool.seasonTotalPointsSnapshot(season);

        if (userPoints == 0 || totalPoints == 0) return 0;

        uint totalReward = seasonRewards[season][token];
        return (totalReward * userPoints) / totalPoints;
    }

    // ============================================
    // 편의 함수
    // ============================================

    /**
     * @notice 현재 시즌에 보상 예치 (간편 함수)
     */
    function depositReward(address token, uint amount) external nonReentrant {
        _validateAddress(token);
        _validateAmount(amount);

        uint currentSeason = stakingPool.currentSeason();

        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        seasonRewards[currentSeason][token] += amount;
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
     * @notice 잘못 전송된 토큰 또는 잔여 토큰 회수 (프로토콜 관리자 전용)
     */
    function sweep(address token, address to, uint amount) external {
        require(msg.sender == address(protocol), RewardPoolOnlyProtocol());
        _validateAddress(to);
        _validateAmount(amount);

        IERC20(token).safeTransfer(to, amount);
    }
}
