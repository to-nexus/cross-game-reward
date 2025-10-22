// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../interfaces/IStakingAddon.sol";
import "../interfaces/IStakingPool.sol";
import "../libraries/PointsLib.sol";
import "../libraries/SeasonLib.sol";
import "./CrossStakingBase.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title StakingPoolBase
 * @notice 스테이킹 풀의 기본 추상 컨트랙트
 * @dev 확장 가능한 스테이킹 로직 제공
 */
abstract contract StakingPoolBase is IStakingPool, CrossStakingBase {
    using SafeERC20 for IERC20;
    using PointsLib for *;
    using SeasonLib for *;

    // ============================================
    // Constants & Roles
    // ============================================

    bytes32 public constant REWARD_POOL_ROLE = keccak256("REWARD_POOL_ROLE");
    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");
    bytes32 public constant ROUTER_ROLE = keccak256("ROUTER_ROLE");

    uint public constant MIN_STAKE = 1e18; // 최소 스테이크: 1 CROSS

    // ============================================
    // Errors
    // ============================================

    error StakingPoolBaseBelowMinStake();
    error StakingPoolBaseNoPosition();
    error StakingPoolBaseInvalidSeasonBlocks();
    error StakingPoolBaseInvalidTimeUnit();
    error StakingPoolBaseSeasonNotEnded();
    error StakingPoolBaseNoActiveSeason();
    error StakingPoolBaseInvalidStartBlock();
    error StakingPoolBaseInvalidEndBlock();
    error StakingPoolBaseAddonNotApproved(); // 승인되지 않은 애드온

    // ============================================
    // Structs
    // ============================================

    /// @notice 스테이킹 포지션
    struct StakePosition {
        uint balance;
        uint points;
        uint lastUpdateBlock;
    }

    /// @notice 시즌 정보
    struct Season {
        uint seasonNumber; // 시즌 번호
        uint startBlock; // 시작 블록
        uint endBlock; // 종료 블록
        bool isFinalized; // 종료 여부
        uint totalPoints; // 총 포인트 (finalize 시 확정, 불변 캐시)
        uint seasonTotalStaked; // 시즌 중 총 스테이킹 (집계용, 변동)
        uint lastAggregatedBlock; // 마지막 집계 블록
        uint aggregatedPoints; // 집계된 포인트 (실시간 누적, finalize 시 totalPoints로 복사)
    }

    /// @notice 시즌별 유저 데이터
    struct UserSeasonData {
        uint points;
        uint balance;
        uint joinBlock;
        bool claimed;
        bool finalized;
    }

    // ============================================
    // State Variables
    // ============================================

    /// @notice 스테이킹 토큰
    IERC20 public immutable stakingToken;

    /// @notice 사용자별 스테이킹 포지션
    mapping(address => StakePosition) public userStakes;

    /// @notice 총 스테이킹 수량
    uint public totalStaked;

    /// @notice 스테이커 목록
    address[] public stakers;
    mapping(address => bool) public isStaker;

    /// @notice 현재 시즌 번호
    uint public currentSeason;

    /// @notice 시즌 기간 (블록 수)
    uint public seasonBlocks;

    /// @notice 풀 종료 블록
    uint public poolEndBlock;

    /// @notice 다음 시즌 시작 블록
    uint public nextSeasonStartBlock;

    /// @notice 포인트 계산 시간 단위
    uint public pointsTimeUnit = 1 hours;

    /// @notice 블록 시간
    uint public blockTime = 1;

    /// @notice 시즌별 정보
    mapping(uint => Season) public seasons;

    /// @notice 시즌별 유저 데이터
    mapping(uint => mapping(address => UserSeasonData)) public userSeasonData;

    /// @notice 사용자별 마지막 finalize된 시즌
    mapping(address => uint) public lastFinalizedSeason;

    /// @notice 연결된 애드온 (선택적)
    IStakingAddon public stakingAddon;

    /// @notice 승인된 애드온 목록 (allowlist)
    mapping(address => bool) public approvedAddons;

    // ============================================
    // Events
    // ============================================

    event Staked(address indexed user, uint amount, uint newBalance);
    event WithdrawnAll(address indexed user, uint amount);
    event PointsUpdated(address indexed user, uint points);
    event SeasonRolledOver(uint indexed oldSeason, uint indexed newSeason, uint totalPoints);
    event SeasonClaimed(address indexed user, uint indexed season, uint points);
    event AddonSet(IStakingAddon indexed oldAddon, IStakingAddon indexed newAddon);
    event AddonCallFailed(IStakingAddon indexed addon, bytes4 selector, string reason);
    event AddonApprovalChanged(IStakingAddon indexed addon, bool approved);

    // ============================================
    // Constructor
    // ============================================

    constructor(
        IERC20 _stakingToken,
        address admin,
        uint _seasonBlocks,
        uint _firstSeasonStartBlock,
        uint _poolEndBlock
    ) CrossStakingBase(admin) {
        _validateAddress(address(_stakingToken));
        require(_seasonBlocks != 0, StakingPoolBaseInvalidSeasonBlocks());
        require(_firstSeasonStartBlock != 0, StakingPoolBaseInvalidStartBlock());
        require(
            _poolEndBlock == 0 || _poolEndBlock > _firstSeasonStartBlock + _seasonBlocks,
            StakingPoolBaseInvalidEndBlock()
        );

        stakingToken = _stakingToken;
        seasonBlocks = _seasonBlocks;
        nextSeasonStartBlock = _firstSeasonStartBlock;
        poolEndBlock = _poolEndBlock;
    }

    // ============================================
    // Core Staking Functions (Template Pattern)
    // ============================================

    /**
     * @notice 토큰 스테이킹
     */
    function stake(uint amount) external virtual nonReentrant {
        _stakeFor(msg.sender, amount, msg.sender);
    }

    /**
     * @notice 다른 사용자를 위한 스테이킹 (Router 전용)
     */
    function stakeFor(address user, uint amount) external virtual nonReentrant onlyRole(ROUTER_ROLE) {
        _stakeFor(user, amount, msg.sender);
    }

    /**
     * @notice 전액 출금
     */
    function withdrawAll() external virtual nonReentrant {
        _withdrawAll(msg.sender, msg.sender);
    }

    /**
     * @notice 유저를 대신해 전액 출금 (Router 전용)
     */
    function withdrawAllFor(address user) external virtual nonReentrant onlyRole(ROUTER_ROLE) {
        _withdrawAll(user, msg.sender);
    }

    // ============================================
    // Internal Core Functions
    // ============================================

    /**
     * @notice 스테이킹 내부 로직
     */
    function _stakeFor(address user, uint amount, address from) internal virtual {
        _ensureSeason();
        require(isSeasonActive(), StakingPoolBaseNoActiveSeason());
        _ensureUserAllPreviousSeasons(user);

        StakePosition storage position = userStakes[user];
        uint oldBalance = position.balance;

        // Hook: 스테이킹 전 처리
        _beforeStake(user, amount, oldBalance);

        if (oldBalance > 0) {
            uint additionalPoints = PointsLib.calculatePoints(
                position.balance,
                position.lastUpdateBlock < seasons[currentSeason].startBlock
                    ? seasons[currentSeason].startBlock
                    : position.lastUpdateBlock,
                block.number,
                blockTime,
                pointsTimeUnit
            );
            userSeasonData[currentSeason][user].points += additionalPoints;
            position.points = 0;
        }

        uint newBalance = oldBalance + amount;
        require(newBalance >= MIN_STAKE, StakingPoolBaseBelowMinStake());

        stakingToken.safeTransferFrom(from, address(this), amount);

        position.balance = newBalance;
        position.lastUpdateBlock = block.number;
        totalStaked += amount;

        // 현재 시즌의 집계용 totalStaked 업데이트
        Season storage currentSeasonData = seasons[currentSeason];
        _updateSeasonAggregation(currentSeason);
        currentSeasonData.seasonTotalStaked += amount;

        UserSeasonData storage seasonData = userSeasonData[currentSeason][user];
        if (seasonData.balance == 0) {
            Season storage current = seasons[currentSeason];
            seasonData.balance = newBalance;
            seasonData.joinBlock = block.number < current.startBlock ? current.startBlock : block.number;
        } else {
            seasonData.balance = newBalance;
        }

        if (!isStaker[user]) {
            stakers.push(user);
            isStaker[user] = true;
        }

        // Hook: 스테이킹 후 처리
        _afterStake(user, amount, newBalance);

        emit Staked(user, amount, newBalance);
    }

    /**
     * @notice 내부 출금 로직
     */
    function _withdrawAll(address user, address recipient) internal virtual {
        _ensureSeason();
        _ensureUserAllPreviousSeasons(user);

        StakePosition storage position = userStakes[user];
        require(position.balance != 0, StakingPoolBaseNoPosition());

        uint amount = position.balance;

        // Hook: 출금 전 처리
        _beforeWithdraw(user, amount);

        if (currentSeason > 0) {
            UserSeasonData storage seasonData = userSeasonData[currentSeason][user];
            seasonData.points = 0;
            seasonData.balance = 0;

            Season storage currentSeasonData = seasons[currentSeason];
            _updateSeasonAggregation(currentSeason);
            currentSeasonData.seasonTotalStaked -= amount;
        }

        position.balance = 0;
        position.points = 0;
        position.lastUpdateBlock = block.number;
        totalStaked -= amount;

        stakingToken.safeTransfer(recipient, amount);

        // Hook: 출금 후 처리
        _afterWithdraw(user, amount, recipient);

        emit WithdrawnAll(user, amount);
    }

    // ============================================
    // Season Management
    // ============================================

    /**
     * @notice 시즌 자동 전환 체크
     */
    function _ensureSeason() internal virtual {
        if (poolEndBlock > 0 && (nextSeasonStartBlock == 0 || nextSeasonStartBlock <= poolEndBlock)) {
            if (block.number >= poolEndBlock) return;
        }

        if (currentSeason == 0) {
            if (block.number >= nextSeasonStartBlock) _startFirstSeason();
            return;
        }

        Season storage current = seasons[currentSeason];
        if (block.number > current.endBlock) {
            if (nextSeasonStartBlock == 0 || block.number >= nextSeasonStartBlock) _rolloverSeason();
        }
    }

    /**
     * @notice 첫 시즌 시작
     */
    function _startFirstSeason() internal virtual {
        require(currentSeason == 0, StakingPoolBaseSeasonNotEnded());

        uint startBlock = nextSeasonStartBlock;
        uint endBlock = startBlock + seasonBlocks;

        if (poolEndBlock > 0 && endBlock > poolEndBlock) endBlock = poolEndBlock;

        currentSeason = 1;
        seasons[1] = Season({
            seasonNumber: 1,
            startBlock: startBlock,
            endBlock: endBlock,
            isFinalized: false,
            totalPoints: 0,
            seasonTotalStaked: totalStaked,
            lastAggregatedBlock: startBlock,
            aggregatedPoints: 0
        });

        nextSeasonStartBlock = 0;

        emit SeasonRolledOver(0, 1, 0);
    }

    /**
     * @notice 시즌 롤오버
     */
    function _rolloverSeason() internal virtual {
        uint oldSeasonNumber = currentSeason;
        Season storage oldSeason = seasons[oldSeasonNumber];

        _finalizeSeasonAggregation(oldSeasonNumber);

        // 시즌 종료 애드온 알림
        _notifySeasonEnd(oldSeasonNumber, oldSeason.seasonTotalStaked, oldSeason.totalPoints);

        oldSeason.isFinalized = true;

        uint newSeasonNumber = oldSeasonNumber + 1;
        currentSeason = newSeasonNumber;

        uint nextStart;
        if (nextSeasonStartBlock > 0) {
            nextStart = nextSeasonStartBlock;
            nextSeasonStartBlock = 0;
        } else {
            nextStart = oldSeason.endBlock + 1;
        }

        uint nextEnd = nextStart + seasonBlocks;

        if (poolEndBlock > 0 && nextStart <= poolEndBlock) if (nextEnd > poolEndBlock) nextEnd = poolEndBlock;

        seasons[newSeasonNumber] = Season({
            seasonNumber: newSeasonNumber,
            startBlock: nextStart,
            endBlock: nextEnd,
            isFinalized: false,
            totalPoints: 0,
            seasonTotalStaked: totalStaked,
            lastAggregatedBlock: nextStart,
            aggregatedPoints: 0
        });

        emit SeasonRolledOver(oldSeasonNumber, newSeasonNumber, 0);
    }

    // ============================================
    // Aggregation Functions
    // ============================================

    /**
     * @notice 시즌 집계 업데이트
     */
    function _updateSeasonAggregation(uint seasonNum) internal {
        Season storage season = seasons[seasonNum];

        if (season.lastAggregatedBlock >= block.number) return;
        if (season.seasonTotalStaked == 0) {
            season.lastAggregatedBlock = block.number;
            return;
        }

        uint additionalPoints = PointsLib.calculatePoints(
            season.seasonTotalStaked, season.lastAggregatedBlock, block.number, blockTime, pointsTimeUnit
        );

        season.aggregatedPoints += additionalPoints;
        season.lastAggregatedBlock = block.number;
    }

    /**
     * @notice 시즌 종료 시 최종 집계
     */
    function _finalizeSeasonAggregation(uint seasonNum) internal {
        Season storage season = seasons[seasonNum];

        if (season.endBlock == 0 || season.lastAggregatedBlock >= season.endBlock) return;

        uint finalBlock = season.endBlock < block.number ? season.endBlock : block.number;

        if (season.seasonTotalStaked > 0) {
            uint additionalPoints = PointsLib.calculatePoints(
                season.seasonTotalStaked, season.lastAggregatedBlock, finalBlock, blockTime, pointsTimeUnit
            );
            season.aggregatedPoints += additionalPoints;
        }

        season.lastAggregatedBlock = finalBlock;
        season.totalPoints = season.aggregatedPoints;
    }

    /**
     * @notice 유저의 이전 시즌들을 스냅샷
     */
    function _ensureUserAllPreviousSeasons(address user) internal {
        if (currentSeason == 0) return;

        uint startSeason = lastFinalizedSeason[user];
        if (startSeason == 0) startSeason = 1;

        if (startSeason >= currentSeason) return;

        for (uint i = startSeason; i < currentSeason;) {
            _ensureUserSeasonSnapshot(user, i);
            unchecked {
                ++i;
            }
        }

        lastFinalizedSeason[user] = currentSeason - 1;
    }

    /**
     * @notice 유저의 특정 시즌 데이터 스냅샷
     */
    function _ensureUserSeasonSnapshot(address user, uint seasonNum) internal virtual {
        if (seasonNum == 0 || seasonNum > currentSeason) return;

        Season storage season = seasons[seasonNum];
        if (!season.isFinalized) return;

        UserSeasonData storage userData = userSeasonData[seasonNum][user];
        if (userData.finalized) return;

        StakePosition storage position = userStakes[user];
        uint lastUpdate = position.lastUpdateBlock;

        if (lastUpdate > season.endBlock) {
            if (userData.balance == 0) {
                userData.finalized = true;
                return;
            }
        }

        if (position.balance == 0 && userData.balance == 0) {
            userData.finalized = true;
            return;
        }

        uint balanceToUse = userData.balance > 0 ? userData.balance : position.balance;
        uint joinBlockToUse = userData.joinBlock > 0 ? userData.joinBlock : lastUpdate;

        if (joinBlockToUse < season.startBlock) {
            userData.balance = balanceToUse;
            userData.joinBlock = season.startBlock;
            userData.points =
                PointsLib.calculatePoints(balanceToUse, season.startBlock, season.endBlock, blockTime, pointsTimeUnit);
        } else if (joinBlockToUse <= season.endBlock) {
            userData.balance = balanceToUse;
            userData.joinBlock = joinBlockToUse;
            userData.points =
                PointsLib.calculatePoints(balanceToUse, joinBlockToUse, season.endBlock, blockTime, pointsTimeUnit);
        }

        userData.finalized = true;
    }

    // ============================================
    // Hook Functions (확장 포인트)
    // ============================================

    function _beforeStake(address user, uint amount, uint oldBalance) internal virtual {}

    function _afterStake(address user, uint amount, uint newBalance) internal virtual {
        // 애드온이 설정되어 있으면 호출
        if (address(stakingAddon) != address(0)) {
            _callAddonSafe(
                stakingAddon,
                abi.encodeWithSignature(
                    "onStake(address,uint256,uint256,uint256,uint256)",
                    user,
                    amount,
                    newBalance - amount, // oldBalance
                    newBalance,
                    currentSeason
                )
            );
        }
    }

    function _beforeWithdraw(address user, uint amount) internal virtual {}

    function _afterWithdraw(address user, uint amount, address /* recipient */ ) internal virtual {
        // 애드온이 설정되어 있으면 호출
        if (address(stakingAddon) != address(0)) {
            _callAddonSafe(
                stakingAddon,
                abi.encodeWithSignature("onWithdraw(address,uint256,uint256)", user, amount, currentSeason)
            );
        }
    }

    /**
     * @notice 시즌 종료 시 애드온 호출
     */
    function _notifySeasonEnd(uint season, uint totalStakedAmount, uint totalPointsAmount) internal {
        if (address(stakingAddon) != address(0)) {
            _callAddonSafe(
                stakingAddon,
                abi.encodeWithSignature(
                    "onSeasonEnd(uint256,uint256,uint256)", season, totalStakedAmount, totalPointsAmount
                )
            );
        }
    }

    /**
     * @notice 애드온 안전 호출 (실패 시 로그만 남김)
     */
    function _callAddonSafe(IStakingAddon addon, bytes memory data) private {
        try this.callAddon(addon, data) {}
        catch Error(string memory reason) {
            emit AddonCallFailed(addon, bytes4(data), reason);
        } catch {
            emit AddonCallFailed(addon, bytes4(data), "Unknown error");
        }
    }

    /**
     * @notice 외부 호출용 헬퍼 (try/catch를 위해 external 필요)
     */
    function callAddon(IStakingAddon addon, bytes memory data) external {
        require(msg.sender == address(this), "Only self");
        (bool success, bytes memory returnData) = address(addon).call(data);
        if (!success) {
            if (returnData.length > 0) {
                assembly {
                    revert(add(32, returnData), mload(returnData))
                }
            } else {
                revert("Addon call failed");
            }
        }
    }

    // ============================================
    // View Functions
    // ============================================

    function isSeasonActive() public view virtual returns (bool) {
        // 가상 시즌 처리 (nextSeasonStartBlock이 설정되어 있을 때)
        if (nextSeasonStartBlock > 0) {
            bool seasonEnded = currentSeason == 0 || seasons[currentSeason].isFinalized
                || block.number > seasons[currentSeason].endBlock;

            if (seasonEnded) {
                if (block.number < nextSeasonStartBlock) return false;

                uint virtualEndBlock = nextSeasonStartBlock + seasonBlocks;

                if (poolEndBlock > 0 && poolEndBlock > nextSeasonStartBlock && virtualEndBlock > poolEndBlock) {
                    virtualEndBlock = poolEndBlock;
                }

                return block.number <= virtualEndBlock;
            }
        }

        if (currentSeason == 0) return false;

        Season storage season = seasons[currentSeason];

        if (poolEndBlock > 0 && season.startBlock <= poolEndBlock) if (block.number >= poolEndBlock) return false;

        return !season.isFinalized && block.number >= season.startBlock && block.number <= season.endBlock;
    }

    function getStakingPower(address user) external view virtual returns (uint) {
        return userStakes[user].balance;
    }

    function getTotalStakingPower() external view virtual returns (uint) {
        return totalStaked;
    }

    // ============================================
    // Addon Management
    // ============================================

    /**
     * @notice 애드온 승인 상태 변경 (관리자 전용)
     * @param addon 애드온 주소
     * @param approved 승인 여부
     */
    function setAddonApproved(IStakingAddon addon, bool approved) external virtual onlyRole(DEFAULT_ADMIN_ROLE) {
        _validateAddress(address(addon));
        approvedAddons[address(addon)] = approved;
        emit AddonApprovalChanged(addon, approved);
    }

    /**
     * @notice 애드온 설정 (관리자 전용)
     * @param newAddon 새 애드온 주소 (0이면 제거)
     * @dev 승인된 애드온만 설정 가능 (allowlist)
     */
    function setStakingAddon(IStakingAddon newAddon) external virtual onlyRole(DEFAULT_ADMIN_ROLE) {
        require(
            (address(newAddon) == address(0) || approvedAddons[address(newAddon)]), StakingPoolBaseAddonNotApproved()
        );
        IStakingAddon oldAddon = stakingAddon;
        stakingAddon = newAddon;
        emit AddonSet(oldAddon, newAddon);
    }
}
