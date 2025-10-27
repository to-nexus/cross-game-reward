// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../interfaces/IRewardPool.sol";

import "../libraries/PointsLib.sol";
import "./CrossStakingBase.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title RewardPoolBase
 * @notice 보상 풀의 기본 추상 컨트랙트
 * @dev 확장 가능한 보상 구조 지원 (추가보상 hook 포함)
 */
abstract contract RewardPoolBase is IRewardPool, CrossStakingBase {
    using SafeERC20 for IERC20;
    // ============================================
    // Roles
    // ============================================

    bytes32 public constant STAKING_POOL_ROLE = keccak256("STAKING_POOL_ROLE");

    // ============================================
    // Errors
    // ============================================

    error RewardPoolBaseInsufficientBalance();
    error RewardPoolBaseAlreadyClaimed();
    error RewardPoolBaseNoRewards();

    // ============================================
    // State Variables
    // ============================================

    /// @notice 시즌별 토큰별 보상량
    mapping(uint => mapping(address => uint)) public seasonRewards;

    /// @notice 시즌별 토큰별 총 청구량
    mapping(uint => mapping(address => uint)) public seasonClaimed;

    /// @notice 사용자별 시즌별 토큰별 청구 여부
    mapping(address => mapping(uint => mapping(address => bool))) public hasClaimedSeasonReward;

    /// @notice 시즌별 보상 토큰 리스트 (season => token[])
    mapping(uint => address[]) private seasonRewardTokens;

    /// @notice 시즌별 토큰별 인덱스 (season => token => index+1, 0이면 없음)
    mapping(uint => mapping(address => uint)) private seasonTokenIndex;

    // ============================================
    // Events
    // ============================================

    event BonusRewardPaid(address indexed user, uint indexed season, address indexed token, uint bonusAmount);

    // ============================================
    // Constructor
    // ============================================

    constructor(address admin) CrossStakingBase(admin) {}

    // ============================================
    // Core Functions (Template Pattern)
    // ============================================

    /**
     * @notice 시즌에 보상 예치
     * @param season 시즌 번호
     * @param token 보상 토큰 주소
     * @param amount 예치할 수량
     * @dev 표준 ERC20 토큰만 지원 (수수료형/리베이스형 토큰 미지원)
     */
    function fundSeason(uint season, address token, uint amount) public virtual nonReentrant {
        _validateAddress(token);
        _validateAmount(amount);

        // Hook: 예치 전 검증
        _beforeFundSeason(season, token, amount);

        // 토큰 전송
        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);

        // 토큰이 처음 예치되는 경우 리스트에 추가
        if (seasonTokenIndex[season][token] == 0) {
            seasonRewardTokens[season].push(token);
            seasonTokenIndex[season][token] = seasonRewardTokens[season].length; // index + 1
        }

        // 보상 누적
        seasonRewards[season][token] += amount;

        // 이벤트 발생
        emit SeasonFunded(season, token, amount, amount);

        // Hook: 예치 후 처리
        _afterFundSeason(season, token, amount);
    }

    /**
     * @notice 사용자에게 보상 지급 (StakingPool 전용)
     * @param user 사용자 주소
     * @param season 시즌 번호
     * @param token 보상 토큰 주소
     * @param userPoints 사용자 포인트
     * @param totalPoints 총 포인트
     */
    function payUser(address user, uint season, address token, uint userPoints, uint totalPoints)
        external
        virtual
        onlyRole(STAKING_POOL_ROLE)
        nonReentrant
    {
        require(!hasClaimedSeasonReward[user][season][token], RewardPoolBaseAlreadyClaimed());
        require(userPoints != 0 && totalPoints != 0, RewardPoolBaseNoRewards());

        uint totalReward = seasonRewards[season][token];
        require(totalReward != 0, RewardPoolBaseNoRewards());

        // 기본 보상 계산
        uint baseReward = PointsLib.calculateProRata(userPoints, totalPoints, totalReward);
        require(baseReward != 0, RewardPoolBaseNoRewards());

        // Hook: 추가 보상 계산
        uint bonusReward = _calculateBonusReward(user, season, token, userPoints, totalPoints, baseReward);

        uint totalPayout = baseReward + bonusReward;

        // 잔액 확인
        uint balance = IERC20(token).balanceOf(address(this));
        require(balance >= totalPayout, RewardPoolBaseInsufficientBalance());

        // 청구 기록
        hasClaimedSeasonReward[user][season][token] = true;
        seasonClaimed[season][token] += totalPayout;

        // Hook: 지급 전 처리
        _beforePayUser(user, season, token, totalPayout);

        // 토큰 전송
        IERC20(token).safeTransfer(user, totalPayout);

        // Hook: 지급 후 처리
        _afterPayUser(user, season, token, totalPayout);

        if (bonusReward > 0) emit BonusRewardPaid(user, season, token, bonusReward);
    }

    // ============================================
    // Hook Functions (확장 포인트)
    // ============================================

    /**
     * @notice 예치 전 Hook (오버라이드 가능)
     */
    function _beforeFundSeason(uint season, address token, uint amount) internal virtual {}

    /**
     * @notice 예치 후 Hook (오버라이드 가능)
     */
    function _afterFundSeason(uint season, address token, uint amount) internal virtual {}

    /**
     * @notice 추가 보상 계산 Hook (오버라이드 필수)
     * @dev 프로젝트별 추가보상 로직 구현
     * @return bonusAmount 추가 보상 금액
     */
    function _calculateBonusReward(
        address, /* user */
        uint, /* season */
        address, /* token */
        uint, /* userPoints */
        uint, /* totalPoints */
        uint /* baseReward */
    ) internal virtual returns (uint bonusAmount) {
        // 기본 구현: 추가 보상 없음
        return 0;
    }

    /**
     * @notice 지급 전 Hook (오버라이드 가능)
     */
    function _beforePayUser(address user, uint season, address token, uint amount) internal virtual {}

    /**
     * @notice 지급 후 Hook (오버라이드 가능)
     */
    function _afterPayUser(address user, uint season, address token, uint amount) internal virtual {}

    // ============================================
    // View Functions
    // ============================================

    /**
     * @notice 시즌별 토큰별 남은 보상량
     */
    function getRemainingRewards(uint season, address token) external view virtual returns (uint) {
        return seasonRewards[season][token] - seasonClaimed[season][token];
    }

    /**
     * @notice 사용자가 받을 수 있는 예상 보상량 (순수 view)
     */
    function getExpectedReward(address user, uint season, address token)
        external
        view
        virtual
        returns (uint expectedReward);

    /**
     * @notice 시즌의 보상 토큰 리스트 조회
     * @param season 시즌 번호
     * @return tokens 해당 시즌에 예치된 토큰 주소 배열
     */
    function getSeasonRewardTokens(uint season) external view returns (address[] memory tokens) {
        return seasonRewardTokens[season];
    }

    /**
     * @notice 시즌의 특정 토큰 보상 상세 정보
     * @param season 시즌 번호
     * @param token 토큰 주소
     * @return total 총 예치된 보상
     * @return claimed 청구된 보상
     * @return remaining 남은 보상
     */
    function getSeasonTokenInfo(uint season, address token)
        external
        view
        returns (uint total, uint claimed, uint remaining)
    {
        total = seasonRewards[season][token];
        claimed = seasonClaimed[season][token];
        remaining = total > claimed ? total - claimed : 0;
    }

    /**
     * @notice 시즌의 모든 보상 토큰 정보 일괄 조회
     * @param season 시즌 번호
     * @return tokens 토큰 주소 배열
     * @return totals 각 토큰의 총 보상
     * @return claimeds 각 토큰의 청구된 보상
     * @return remainings 각 토큰의 남은 보상
     */
    function getSeasonAllRewards(uint season)
        external
        view
        returns (address[] memory tokens, uint[] memory totals, uint[] memory claimeds, uint[] memory remainings)
    {
        tokens = seasonRewardTokens[season];
        uint length = tokens.length;

        totals = new uint[](length);
        claimeds = new uint[](length);
        remainings = new uint[](length);

        for (uint i = 0; i < length; i++) {
            address token = tokens[i];
            totals[i] = seasonRewards[season][token];
            claimeds[i] = seasonClaimed[season][token];
            remainings[i] = totals[i] > claimeds[i] ? totals[i] - claimeds[i] : 0;
        }
    }
}
