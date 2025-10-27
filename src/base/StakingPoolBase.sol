// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

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
    uint public constant MAX_AUTO_ROLLOVERS = 50; // 자동 롤오버 최대 시즌 수

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
    error StakingPoolBaseTooManySeasons();
    error StakingPoolBaseOnlySelf();
    error StakingPoolBaseInvalidMaxRollovers();

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
        uint forfeitedPoints; // unstake로 몰수된 포인트 (totalPoints 계산 시 제외)
    }

    /// @notice 시즌별 유저 데이터
    struct UserSeasonData {
        uint points;
        uint balance;
        uint joinBlock; // 실제 참여 블록 (정보용)
        uint lastPointsBlock; // 마지막 포인트 계산 블록
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

    /// @notice 사전 예치 시작 블록 (첫 시즌 시작 전 보상 예치 가능 시점)
    uint public preDepositStartBlock;

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

    // ============================================
    // Events
    // ============================================

    event Staked(address indexed user, uint amount, uint newBalance);
    event WithdrawnAll(address indexed user, uint amount);
    event PointsUpdated(address indexed user, uint points);
    event SeasonRolledOver(uint indexed oldSeason, uint indexed newSeason, uint totalPoints);
    event SeasonClaimed(address indexed user, uint indexed season, uint points);
    event ManualRolloverCompleted(uint rolloversPerformed, uint fromSeason, uint toSeason);
    event PointsForfeited(address indexed user, uint indexed season, uint amount);
    event SeasonAggregationUpdated(uint indexed season, uint aggregatedPoints, uint lastAggregatedBlock);

    // ============================================
    // Constructor
    // ============================================

    constructor(
        IERC20 _stakingToken,
        address admin,
        uint _seasonBlocks,
        uint _firstSeasonStartBlock,
        uint _poolEndBlock,
        uint _preDepositStartBlock
    ) CrossStakingBase(admin) {
        _validateAddress(address(_stakingToken));
        require(_seasonBlocks != 0, StakingPoolBaseInvalidSeasonBlocks());
        require(_firstSeasonStartBlock != 0, StakingPoolBaseInvalidStartBlock());
        require(
            _poolEndBlock == 0 || _poolEndBlock > _firstSeasonStartBlock + _seasonBlocks,
            StakingPoolBaseInvalidEndBlock()
        );
        require(
            _preDepositStartBlock == 0 || _preDepositStartBlock <= _firstSeasonStartBlock,
            "preDepositStartBlock must be before or equal to firstSeasonStartBlock"
        );

        stakingToken = _stakingToken;
        seasonBlocks = _seasonBlocks;
        nextSeasonStartBlock = _firstSeasonStartBlock;
        poolEndBlock = _poolEndBlock;
        preDepositStartBlock = _preDepositStartBlock;

        // Protocol(admin)에게 MANAGER_ROLE 부여
        _grantRole(MANAGER_ROLE, admin);
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
    function stakeFor(address user, uint amount) external virtual nonReentrant {
        require(_isApprovedRouter(msg.sender), CrossStakingBaseNotAuthorized());
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
    function withdrawAllFor(address user) external virtual nonReentrant {
        require(_isApprovedRouter(msg.sender), CrossStakingBaseNotAuthorized());
        _withdrawAll(user, msg.sender);
    }

    // ============================================
    // Internal Core Functions
    // ============================================

    /**
     * @notice 스테이킹 내부 로직
     */
    function _stakeFor(address user, uint amount, address from) internal virtual {
        // 풀 종료 체크
        if (poolEndBlock > 0 && block.number >= poolEndBlock) revert StakingPoolBaseNoActiveSeason();

        _ensureSeason();

        // 시즌 활성 체크
        if (currentSeason == 0) {
            // 첫 시즌이 아직 생성되지 않음
            if (preDepositStartBlock > 0 && block.number >= preDepositStartBlock) {
                // preDeposit 기간: preDepositStartBlock이 설정되어 있고 해당 블록 이후
                // 시즌 시작 전에도 스테이킹 가능
            } else {
                // preDeposit이 없거나 아직 preDeposit 블록 이전이면 시즌 시작 블록 이후부터만 가능
                require(block.number >= nextSeasonStartBlock, StakingPoolBaseNoActiveSeason());
            }
        } else {
            // 시즌이 생성되었지만, 활성 시즌인지 확인
            require(isSeasonActive(), StakingPoolBaseNoActiveSeason());
        }

        _ensureUserAllPreviousSeasons(user);

        StakePosition storage position = userStakes[user];
        uint oldBalance = position.balance;

        if (oldBalance > 0) {
            uint effectiveStart = position.lastUpdateBlock < seasons[currentSeason].startBlock
                ? seasons[currentSeason].startBlock
                : position.lastUpdateBlock;
            uint additionalPoints =
                PointsLib.calculatePoints(position.balance, effectiveStart, block.number, blockTime, pointsTimeUnit);
            UserSeasonData storage currentUserData = userSeasonData[currentSeason][user];
            currentUserData.points += additionalPoints;
            currentUserData.lastPointsBlock = block.number;
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
            seasonData.lastPointsBlock = block.number;
        } else {
            seasonData.balance = newBalance;
            seasonData.lastPointsBlock = block.number;
        }

        if (!isStaker[user]) {
            stakers.push(user);
            isStaker[user] = true;
        }

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

        if (currentSeason > 0) {
            Season storage currentSeasonData = seasons[currentSeason];
            UserSeasonData storage seasonData = userSeasonData[currentSeason][user];

            // 현재 시즌의 집계 업데이트
            _updateSeasonAggregation(currentSeason);

            // 출금하는 사용자의 현재 시즌 포인트 계산 (몰수될 포인트)
            uint userForfeitedPoints = 0;
            if (seasonData.balance > 0) {
                userForfeitedPoints = seasonData.points;

                uint effectiveStart = seasonData.lastPointsBlock > 0 ? seasonData.lastPointsBlock : seasonData.joinBlock;
                if (effectiveStart < currentSeasonData.startBlock) effectiveStart = currentSeasonData.startBlock;

                uint additionalPoints = PointsLib.calculatePoints(
                    seasonData.balance, effectiveStart, block.number, blockTime, pointsTimeUnit
                );
                userForfeitedPoints += additionalPoints;
            } else if (position.balance > 0) {
                uint lastUpdate = position.lastUpdateBlock;
                uint effectiveStart =
                    lastUpdate < currentSeasonData.startBlock ? currentSeasonData.startBlock : lastUpdate;
                userForfeitedPoints =
                    PointsLib.calculatePoints(position.balance, effectiveStart, block.number, blockTime, pointsTimeUnit);
            }

            // 사용자의 시즌 데이터 초기화 (포인트 몰수)
            seasonData.points = 0;
            seasonData.balance = 0;

            // 몰수된 포인트를 시즌의 forfeitedPoints에 누적
            if (userForfeitedPoints > 0) {
                currentSeasonData.forfeitedPoints += userForfeitedPoints;
                emit PointsForfeited(user, currentSeason, userForfeitedPoints);
            }

            // seasonTotalStaked 감소
            currentSeasonData.seasonTotalStaked -= amount;
        }

        position.balance = 0;
        position.points = 0;
        position.lastUpdateBlock = block.number;
        totalStaked -= amount;

        stakingToken.safeTransfer(recipient, amount);

        emit WithdrawnAll(user, amount);
    }

    // ============================================
    // Season Management
    // ============================================

    /**
     * @notice 시즌 자동 전환 체크
     * @dev 여러 시즌이 지나간 경우 현재 블록에 해당하는 시즌까지 모두 롤오버
     * @dev 가스 한도를 고려하여 최대 50개 시즌까지 한 번에 처리
     */
    function _ensureSeason() internal virtual {
        // 풀 종료 체크
        if (poolEndBlock > 0 && (nextSeasonStartBlock == 0 || nextSeasonStartBlock <= poolEndBlock)) {
            if (block.number >= poolEndBlock) return;
        }

        // 첫 시즌 시작
        if (currentSeason == 0) {
            if (block.number >= nextSeasonStartBlock) _startFirstSeason();
            // 첫 시즌 생성 후에도 계속 확인하여 추가 롤오버 필요 여부 체크
            else return;
        }

        // 현재 시즌이 끝났다면 필요한 시즌까지 모두 롤오버
        // 최대 50개 시즌까지 한 번에 처리 (가스 한도 고려)
        uint maxRollovers = MAX_AUTO_ROLLOVERS;
        uint rolloversPerformed = 0;

        while (currentSeason > 0 && rolloversPerformed < maxRollovers) {
            Season storage current = seasons[currentSeason];

            // 현재 시즌이 아직 진행 중이면 종료
            if (block.number <= current.endBlock) break;

            // 다음 시즌 시작 블록이 설정되어 있고 아직 도달하지 않았으면 대기
            if (nextSeasonStartBlock > 0 && block.number < nextSeasonStartBlock) break;

            // 시즌 롤오버 실행
            _rolloverSeason();
            unchecked {
                ++rolloversPerformed;
            }
        }

        // 50개를 초과하는 시즌이 쌓인 경우 에러 (비정상 상황)
        // 이 경우 관리자가 manualRolloverSeasons()를 호출하여 처리해야 함
        require(rolloversPerformed < maxRollovers, StakingPoolBaseTooManySeasons());
    }

    /**
     * @notice 첫 시즌 시작
     */
    function _startFirstSeason() internal virtual {
        require(currentSeason == 0, StakingPoolBaseSeasonNotEnded());

        uint startBlock = nextSeasonStartBlock;
        nextSeasonStartBlock = 0;

        currentSeason = 1;
        _createSeason(1, startBlock);

        emit SeasonRolledOver(0, 1, 0);
    }

    /**
     * @notice 시즌 롤오버
     */
    function _rolloverSeason() internal virtual {
        uint oldSeasonNumber = currentSeason;
        Season storage oldSeason = seasons[oldSeasonNumber];

        _finalizeSeasonAggregation(oldSeasonNumber);

        oldSeason.isFinalized = true;

        uint newSeasonNumber = oldSeasonNumber + 1;
        currentSeason = newSeasonNumber;

        uint nextStart;
        if (nextSeasonStartBlock > 0) {
            nextStart = nextSeasonStartBlock;
            nextSeasonStartBlock = 0;
        } else {
            // 이전 시즌 endBlock 다음 블록부터 시작 (블록 겹침 없음)
            nextStart = oldSeason.endBlock + 1;
        }

        _createSeason(newSeasonNumber, nextStart);

        emit SeasonRolledOver(oldSeasonNumber, newSeasonNumber, 0);
    }

    /**
     * @notice endBlock 계산 헬퍼 함수 (중복 로직 공통화)
     * @param startBlock 시작 블록
     * @return endBlock 계산된 종료 블록
     */
    function _calculateEndBlock(uint startBlock) internal view returns (uint) {
        // endBlock 계산: inclusive이므로 startBlock + seasonBlocks - 1
        // 예: startBlock=1, seasonBlocks=100 → endBlock=100 (블록 1~100, 총 100개)
        uint endBlock = startBlock + seasonBlocks - 1;

        // poolEndBlock 제한 적용
        if (poolEndBlock > 0 && endBlock > poolEndBlock) endBlock = poolEndBlock;

        return endBlock;
    }

    /**
     * @notice 시즌 생성 헬퍼 함수 (중복 로직 공통화)
     * @param seasonNumber 시즌 번호
     * @param startBlock 시작 블록
     */
    function _createSeason(uint seasonNumber, uint startBlock) internal {
        uint endBlock = _calculateEndBlock(startBlock);

        // 시즌 생성
        seasons[seasonNumber] = Season({
            seasonNumber: seasonNumber,
            startBlock: startBlock,
            endBlock: endBlock,
            isFinalized: false,
            totalPoints: 0,
            seasonTotalStaked: totalStaked,
            lastAggregatedBlock: startBlock,
            aggregatedPoints: 0,
            forfeitedPoints: 0
        });
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
        emit SeasonAggregationUpdated(seasonNum, season.aggregatedPoints, block.number);
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

        // userData.balance가 이미 기록되어 있으면 그것을 사용 (과거 시즌에 참여함)
        if (userData.balance > 0) {
            uint userJoinBlock = userData.joinBlock > 0 ? userData.joinBlock : season.startBlock;

            if (userJoinBlock < season.startBlock) {
                userData.joinBlock = season.startBlock;
                userData.points = PointsLib.calculatePoints(
                    userData.balance, season.startBlock, season.endBlock, blockTime, pointsTimeUnit
                );
            } else if (userJoinBlock <= season.endBlock) {
                userData.points = PointsLib.calculatePoints(
                    userData.balance, userJoinBlock, season.endBlock, blockTime, pointsTimeUnit
                );
            }
            userData.finalized = true;
            return;
        }

        // userData.balance가 0인 경우:
        // - 시즌 시작 전에 스테이킹하고 시즌 종료 전까지 유지한 경우만 position.balance 사용
        // - 시즌 종료 후에 스테이킹한 경우 참여하지 않은 것으로 처리
        if (lastUpdate > season.endBlock) {
            // 시즌 종료 후에 스테이킹 → 이 시즌에 참여하지 않음
            userData.finalized = true;
            return;
        }

        if (position.balance == 0) {
            // 스테이킹이 없었음
            userData.finalized = true;
            return;
        }

        // position.balance 사용 (자동 참여 케이스)
        // 단, lastUpdate가 시즌 기간 내에 있는 경우만
        uint positionJoinBlock = lastUpdate;

        if (positionJoinBlock < season.startBlock) {
            // 시즌 시작 전에 스테이킹 → 전체 시즌 참여
            userData.balance = position.balance;
            userData.joinBlock = season.startBlock;
            userData.points = PointsLib.calculatePoints(
                position.balance, season.startBlock, season.endBlock, blockTime, pointsTimeUnit
            );
        } else if (positionJoinBlock <= season.endBlock) {
            // 시즌 중에 스테이킹 → joinBlock부터 endBlock까지 참여
            userData.balance = position.balance;
            userData.joinBlock = positionJoinBlock;
            userData.points = PointsLib.calculatePoints(
                position.balance, positionJoinBlock, season.endBlock, blockTime, pointsTimeUnit
            );
        }

        userData.finalized = true;
    }

    // ============================================
    // Router Approval Check (확장 포인트)
    // ============================================

    /**
     * @notice Router 승인 여부 확인
     * @param router Router 주소
     * @return approved 승인 여부 (로컬 ROUTER_ROLE만 체크, 자식에서 override)
     * @dev 자식 컨트랙트에서 Protocol 글로벌 승인도 함께 체크하도록 override
     */
    function _isApprovedRouter(address router) internal view virtual returns (bool) {
        return hasRole(ROUTER_ROLE, router);
    }

    // ============================================
    // View Functions
    // ============================================

    /**
     * @notice 현재 시즌이 활성화되어 있는지 확인
     * @return true면 스테이킹 가능, false면 불가능
     */
    function isSeasonActive() public view virtual returns (bool) {
        // 풀 종료 체크
        if (poolEndBlock > 0 && block.number >= poolEndBlock) return false;

        // 첫 시즌 시작 전
        if (currentSeason == 0) {
            // nextSeasonStartBlock 도달 확인
            if (block.number < nextSeasonStartBlock) return false;

            // poolEndBlock 체크
            if (poolEndBlock > 0 && nextSeasonStartBlock >= poolEndBlock) return false;

            // 가상 첫 시즌 활성 기간 계산
            uint virtualEndBlock = _calculateEndBlock(nextSeasonStartBlock);
            return block.number <= virtualEndBlock;
        }

        Season storage season = seasons[currentSeason];

        // 현재 시즌이 종료되고 다음 시즌이 아직 롤오버되지 않은 경우
        if (block.number > season.endBlock) {
            // nextSeasonStartBlock이 설정되어 있고 아직 도달하지 않았으면 비활성
            if (nextSeasonStartBlock > 0 && block.number < nextSeasonStartBlock) return false;

            // 가상 다음 시즌 계산
            uint nextStart = nextSeasonStartBlock > 0 ? nextSeasonStartBlock : season.endBlock + 1;

            // poolEndBlock 체크
            if (poolEndBlock > 0 && nextStart >= poolEndBlock) return false;

            uint virtualEndBlock = _calculateEndBlock(nextStart);
            return block.number <= virtualEndBlock;
        }

        // 일반 시즌 활성 체크
        return !season.isFinalized && block.number >= season.startBlock && block.number <= season.endBlock;
    }

    function getStakingPower(address user) external view virtual returns (uint) {
        return userStakes[user].balance;
    }

    function getTotalStakingPower() external view virtual returns (uint) {
        return totalStaked;
    }

    // ============================================
    // Manual Season Management (Admin)
    // ============================================

    /**
     * @notice 수동으로 여러 시즌 롤오버 (관리자 전용)
     * @param maxRollovers 최대 롤오버 횟수
     * @return rolloversPerformed 실제로 수행된 롤오버 횟수
     * @dev 50개 이상의 시즌이 쌓인 경우 관리자가 여러 번 호출하여 처리
     */
    function manualRolloverSeasons(uint maxRollovers)
        external
        virtual
        onlyRole(MANAGER_ROLE)
        returns (uint rolloversPerformed)
    {
        require(maxRollovers > 0 && maxRollovers <= 100, StakingPoolBaseInvalidMaxRollovers());

        uint startSeason = currentSeason;

        // 풀 종료 체크
        if (poolEndBlock > 0 && block.number >= poolEndBlock) {
            emit ManualRolloverCompleted(0, startSeason, currentSeason);
            return 0;
        }

        // 첫 시즌이 아직 시작 안 됐으면 시작
        if (currentSeason == 0) {
            if (block.number >= nextSeasonStartBlock) {
                _startFirstSeason();
                rolloversPerformed = 1;
                emit ManualRolloverCompleted(1, 0, currentSeason);
            } else {
                emit ManualRolloverCompleted(0, 0, 0);
                return 0;
            }
        }

        // 여러 시즌 롤오버
        uint count = 0;
        while (currentSeason > 0 && count < maxRollovers) {
            Season storage current = seasons[currentSeason];

            // 현재 시즌이 아직 진행 중이면 종료
            if (block.number <= current.endBlock) break;

            // 다음 시즌 시작 블록이 설정되어 있고 아직 도달하지 않았으면 대기
            if (nextSeasonStartBlock > 0 && block.number < nextSeasonStartBlock) break;

            // 시즌 롤오버
            _rolloverSeason();

            unchecked {
                ++count;
                ++rolloversPerformed;
            }
        }

        // 실제 처리량 로깅
        emit ManualRolloverCompleted(rolloversPerformed, startSeason, currentSeason);

        return rolloversPerformed;
    }

    /**
     * @notice 롤오버가 필요한 시즌 개수 조회
     * @return pendingSeasons 롤오버 대기 중인 시즌 수
     */
    function getPendingSeasonRollovers() external view returns (uint pendingSeasons) {
        // 풀 종료됨
        if (poolEndBlock > 0 && block.number >= poolEndBlock) return 0;

        // 시즌 0 (첫 시즌 롤오버)
        if (currentSeason == 0) {
            if (block.number < nextSeasonStartBlock) return 0;

            // 첫 시즌 롤오버부터 현재 블록까지 필요한 시즌 개수 계산
            uint currentBlk = block.number;
            uint startBlock = nextSeasonStartBlock;
            uint count = 0;

            // 첫 시즌
            count = 1;
            uint endBlock = _calculateEndBlock(startBlock);

            // poolEndBlock에 도달하면 종료
            if (poolEndBlock > 0 && endBlock >= poolEndBlock) return count;

            // 추가 시즌 계산
            if (currentBlk > endBlock) {
                uint blocksAfterEnd = currentBlk - endBlock;
                // 각 시즌은 endBlock + 1부터 시작
                uint additionalSeasons = (blocksAfterEnd + seasonBlocks - 1) / seasonBlocks;

                // poolEndBlock 고려하여 실제 가능한 시즌만 카운트
                if (poolEndBlock > 0) {
                    uint nextStartBlk = endBlock + 1;
                    for (uint i = 0; i < additionalSeasons; i++) {
                        if (nextStartBlk >= poolEndBlock) break;
                        count++;
                        nextStartBlk += seasonBlocks;
                    }
                } else {
                    count += additionalSeasons;
                }
            }

            return count;
        }

        // 다음 시즌 시작 대기 중 (수동으로 설정된 경우)
        if (nextSeasonStartBlock > 0 && block.number < nextSeasonStartBlock) return 0;

        Season storage current = seasons[currentSeason];

        // 현재 시즌이 진행 중
        if (block.number <= current.endBlock) return 0;

        // 현재 시즌 종료 후 롤오버 필요한 개수 계산
        // nextSeasonStartBlock이 설정되어 있으면 그것을 사용, 아니면 current.endBlock + 1 사용
        uint nextStart = nextSeasonStartBlock > 0 ? nextSeasonStartBlock : current.endBlock + 1;

        // poolEndBlock 체크
        if (poolEndBlock > 0 && nextStart >= poolEndBlock) return 0;

        uint rolloverCount = 0;
        uint currentBlock = block.number;

        // 다음 시즌부터 현재 블록까지 필요한 시즌 개수 계산
        while (nextStart <= currentBlock) {
            uint endBlock = _calculateEndBlock(nextStart);
            rolloverCount++;

            // poolEndBlock에 도달하거나 현재 블록을 커버하면 종료
            if ((poolEndBlock > 0 && endBlock >= poolEndBlock) || currentBlock <= endBlock) break;

            nextStart = endBlock + 1;

            // 무한 루프 방지 (최대 100개 시즌)
            if (rolloverCount >= 100) break;
        }

        return rolloverCount;
    }
}
