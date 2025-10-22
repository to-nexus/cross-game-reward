// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RewardPool.sol";
import "./StakingProtocol.sol";
import "./interfaces/IRewardPool.sol";
import "./interfaces/IStakingPool.sol";

import "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

/**
 * @title StakingPoolCode
 * @notice StakingPool의 creation code를 반환하는 컨트랙트
 */
contract StakingPoolCode {
    function code() external pure returns (bytes memory) {
        return type(StakingPool).creationCode;
    }
}

/**
 * @title StakingPool
 * @notice 프로젝트별 $CROSS 스테이킹 관리 및 포인트 누적 (시즌 기반)
 * @dev 문서 기반 구현: 락업 없음, 전액 출금, 포인트 몰수, 시즌 시스템
 */
contract StakingPool is IStakingPool, AccessControlDefaultAdminRules, ReentrancyGuardTransient {
    using SafeERC20 for IERC20;

    // ============ Role 정의 ============

    bytes32 public constant REWARD_POOL_ROLE = keccak256("REWARD_POOL_ROLE");
    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");

    // ============ 에러 ============

    error StakingPoolCanNotZeroAddress();
    error StakingPoolAlreadySet();
    error StakingPoolBelowMinStake();
    error StakingPoolNoPosition();
    error StakingPoolNotAuthorized();
    error StakingPoolInvalidSeasonBlocks();
    error StakingPoolInvalidTimeUnit();
    error StakingPoolSeasonNotEnded();
    error StakingPoolAlreadyClaimed();
    error StakingPoolNoActiveSeason();
    error StakingPoolInvalidStartBlock();
    error StakingPoolInvalidEndBlock();

    // ============ 상수 ============

    uint public constant MIN_STAKE = 1e18; // 최소 스테이크: 1 CROSS
    uint public constant POINTS_PRECISION = 1e6; // 포인트 정밀도: 소수점 6자리

    // ============ 구조체 ============

    /// @notice 스테이킹 포지션
    struct StakePosition {
        uint balance; // 스테이킹 수량
        uint points; // 현재 시즌 누적 포인트
        uint lastUpdateBlock; // 마지막 업데이트 블록
    }

    /// @notice 시즌 정보
    struct Season {
        uint seasonNumber; // 시즌 번호
        uint startBlock; // 시작 블록
        uint endBlock; // 종료 블록
        bool isFinalized; // 종료 여부
        uint totalPoints; // 시즌 종료 시 계산된 총 포인트 (가스 최적화)
    }

    /// @notice 시즌별 유저 데이터 (통합)
    struct UserSeasonData {
        uint points; // 시즌별 포인트 (finalized 시 확정)
        uint balance; // 해당 시즌의 stake 수량
        uint joinBlock; // 해당 시즌 참여 시작 블록
        bool claimed; // 보상 청구 여부
        bool finalized; // 해당 유저의 시즌 데이터가 확정되었는지 (lazy snapshot 여부)
    }

    // ============ 상태 변수 ============

    /// @notice 프로젝트 ID
    uint public immutable projectId;

    /// @notice staking token
    IERC20 public immutable stakingToken;

    /// @notice 프로토콜 컨트랙트
    StakingProtocol public immutable protocol;

    /// @notice 연결된 RewardPool
    RewardPool public rewardPool;

    /// @notice 승인된 Router 주소
    mapping(address => bool) public approvedRouters;

    /// @notice 사용자별 스테이킹 포지션
    mapping(address => StakePosition) public userStakes;

    /// @notice 총 스테이킹 수량
    uint public totalStaked;

    /// @notice 현재 시즌 총 포인트
    uint public totalPointsInSeason;

    /// @notice 스테이커 목록
    address[] public stakers;
    mapping(address => bool) public isStaker;

    // ============ 시즌 시스템 ============

    /// @notice 현재 시즌 번호 (0이면 아직 시작 안함)
    uint public currentSeason;

    /// @notice 시즌 기간 (블록 수)
    uint public seasonBlocks;

    /// @notice 풀 종료 블록 (0이면 무한 진행)
    uint public poolEndBlock;

    /// @notice 다음 시즌 시작 블록 (0이면 자동, >0이면 해당 블록에서 시작)
    uint public nextSeasonStartBlock;

    /// @notice 포인트 계산 시간 단위 (기본: 1 hour = 3600 seconds)
    uint public pointsTimeUnit = 1 hours;

    /// @notice 블록 시간 (초 단위, 1초/블록)
    uint public blockTime = 1;

    /// @notice 시즌별 정보
    mapping(uint => Season) public seasons;

    /// @notice 시즌별 유저 데이터 (통합)
    mapping(uint => mapping(address => UserSeasonData)) public userSeasonData;

    // ============ 이벤트 ============

    event Staked(address indexed user, uint amount, uint newBalance);
    event WithdrawnAll(address indexed user, uint amount);
    event PointsUpdated(address indexed user, uint points);
    event SeasonRolledOver(uint indexed oldSeason, uint indexed newSeason, uint totalPoints);
    event SeasonClaimed(address indexed user, uint indexed season, uint points);

    // ============ 생성자 ============

    constructor(
        uint _projectId,
        address _stakingToken,
        address _protocol,
        uint _seasonBlocks,
        uint _firstSeasonStartBlock,
        uint _poolEndBlock
    ) AccessControlDefaultAdminRules(3 days, _protocol) {
        require(_stakingToken != address(0), StakingPoolCanNotZeroAddress());
        require(_protocol != address(0), StakingPoolCanNotZeroAddress());
        require(_seasonBlocks > 0, StakingPoolInvalidSeasonBlocks());
        require(_firstSeasonStartBlock > 0, StakingPoolInvalidStartBlock());
        if (_poolEndBlock > 0) {
            require(_poolEndBlock > _firstSeasonStartBlock + _seasonBlocks, StakingPoolInvalidEndBlock());
        }

        projectId = _projectId;
        stakingToken = IERC20(_stakingToken);
        protocol = StakingProtocol(_protocol);

        // 시즌 초기화
        seasonBlocks = _seasonBlocks;
        currentSeason = 0; // 아직 시작 안함
        nextSeasonStartBlock = _firstSeasonStartBlock;
        poolEndBlock = _poolEndBlock;

        // 첫 번째 시즌은 lazy 생성 (nextSeasonStartBlock에 도달하면 자동 생성)
    }

    // ============================================
    // Execute Functions
    // ============================================

    // ============ Staking Functions (External) ============

    /**
     * @notice 토큰 스테이킹 (분할 스테이킹 허용)
     * @param amount 스테이킹 수량
     */
    function stake(uint amount) external nonReentrant {
        _stakeFor(msg.sender, amount, msg.sender);
    }

    /**
     * @notice 다른 사용자를 위한 스테이킹 (Router 전용)
     */
    function stakeFor(address user, uint amount) external nonReentrant {
        require(approvedRouters[msg.sender], StakingPoolNotAuthorized());
        _stakeFor(user, amount, msg.sender);
    }

    /**
     * @notice 전액 출금 (포인트 몰수)
     */
    function withdrawAll() external nonReentrant {
        _withdrawAll(msg.sender, msg.sender);
    }

    /**
     * @notice 유저를 대신해 전액 출금 (Router 전용)
     * @param user 출금할 사용자 주소
     */
    function withdrawAllFor(address user) external nonReentrant {
        require(approvedRouters[msg.sender], StakingPoolNotAuthorized());
        _withdrawAll(user, msg.sender);
    }

    // ============ Season Management (External) ============

    /**
     * @notice 시즌 롤오버 (누구나 호출 가능)
     */
    function rolloverSeason() external {
        require(currentSeason > 0, StakingPoolNoActiveSeason());
        Season storage current = seasons[currentSeason];
        require(block.number > current.endBlock, StakingPoolSeasonNotEnded());
        require(nextSeasonStartBlock == 0 || block.number >= nextSeasonStartBlock, StakingPoolSeasonNotEnded());
        _rolloverSeason();
    }

    /**
     * @notice 시즌 보상 청구 (Lazy snapshot 및 totalPoints 계산)
     */
    function claimSeason(uint seasonNumber, address rewardToken) external nonReentrant {
        Season storage season = seasons[seasonNumber];

        require(season.isFinalized, StakingPoolSeasonNotEnded());

        // 해당 시즌의 유저 데이터를 lazy snapshot
        _ensureUserSeasonSnapshot(msg.sender, seasonNumber);

        UserSeasonData storage userData = userSeasonData[seasonNumber][msg.sender];
        require(!userData.claimed, StakingPoolAlreadyClaimed());

        uint userPoints = userData.points;

        // totalPoints가 아직 계산되지 않았으면 계산
        uint totalPoints = season.totalPoints;
        if (totalPoints == 0) {
            totalPoints = _calculateSeasonTotalPoints(seasonNumber);
            season.totalPoints = totalPoints; // 캐시 저장
        }

        if (userPoints > 0 && totalPoints > 0) {
            userData.claimed = true;

            // RewardPool에 보상 지급 요청
            rewardPool.payUser(msg.sender, seasonNumber, rewardToken, userPoints, totalPoints);

            emit SeasonClaimed(msg.sender, seasonNumber, userPoints);
        }
    }

    // ============ Points Management (External) ============

    /**
     * @notice 사용자 포인트 업데이트
     */
    function updatePoints(address user) external {
        require(
            hasRole(REWARD_POOL_ROLE, msg.sender) || hasRole(DEFAULT_ADMIN_ROLE, msg.sender), StakingPoolNotAuthorized()
        );
        _updatePoints(user);
    }

    // ============ Configuration Functions (External - Admin) ============

    function setRewardPool(address _rewardPool) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(address(rewardPool) == address(0), StakingPoolAlreadySet());
        require(_rewardPool != address(0), StakingPoolCanNotZeroAddress());
        rewardPool = RewardPool(_rewardPool);

        // RewardPool에 REWARD_POOL_ROLE 부여
        _grantRole(REWARD_POOL_ROLE, _rewardPool);
    }

    /**
     * @notice Router 승인/취소
     */
    function setApprovedRouter(address router, bool approved) external onlyRole(DEFAULT_ADMIN_ROLE) {
        approvedRouters[router] = approved;
    }

    /**
     * @notice 포인트 계산 시간 단위 설정
     */
    function setPointsTimeUnit(uint _timeUnit) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_timeUnit > 0, StakingPoolInvalidTimeUnit());
        pointsTimeUnit = _timeUnit;
    }

    /**
     * @notice 블록 시간 설정
     */
    function setBlockTime(uint _blockTime) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_blockTime > 0, StakingPoolInvalidTimeUnit());
        blockTime = _blockTime;
    }

    /**
     * @notice 다음 시즌 시작 블록 설정 (풀 재가동)
     * @param _startBlock 시작 블록 (0이면 자동)
     */
    function setNextSeasonStart(uint _startBlock) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_startBlock == 0 || _startBlock > block.number, StakingPoolInvalidStartBlock());
        nextSeasonStartBlock = _startBlock;
    }

    /**
     * @notice 풀 종료 블록 설정
     * @param _endBlock 종료 블록 (0이면 무한)
     */
    function setPoolEndBlock(uint _endBlock) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (_endBlock > 0) require(_endBlock > block.number, StakingPoolInvalidEndBlock());
        poolEndBlock = _endBlock;
    }

    // ============ 조회 함수 (시즌 상태) ============

    /**
     * @notice 진행중인 시즌이 있는지 확인
     * @return 진행중인 시즌 여부
     */
    function isSeasonActive() public view returns (bool) {
        // Virtual Season: nextSeasonStartBlock이 설정되어 있고, 현재 시즌이 없거나 종료됨
        if (nextSeasonStartBlock > 0) {
            // currentSeason이 0이거나, 명시적으로 finalize되었거나, endBlock이 지난 경우
            bool seasonEnded = currentSeason == 0 || seasons[currentSeason].isFinalized
                || block.number > seasons[currentSeason].endBlock;

            if (seasonEnded) {
                if (block.number < nextSeasonStartBlock) return false;

                // 시작 블록 ~ (시작 + seasonBlocks) 범위 내
                uint virtualEndBlock = nextSeasonStartBlock + seasonBlocks;

                // poolEndBlock이 nextSeasonStartBlock보다 뒤에 있으면 제한 적용
                if (poolEndBlock > 0 && poolEndBlock > nextSeasonStartBlock && virtualEndBlock > poolEndBlock) {
                    virtualEndBlock = poolEndBlock;
                }

                return block.number <= virtualEndBlock;
            }
        }

        // 실제 시즌이 생성된 경우
        if (currentSeason == 0) return false;

        Season storage season = seasons[currentSeason];

        // poolEndBlock 체크 (전역 종료)
        // 단, 시즌 startBlock > poolEndBlock 이면 재시작이므로 체크 무시
        if (poolEndBlock > 0 && season.startBlock <= poolEndBlock) if (block.number >= poolEndBlock) return false;

        return !season.isFinalized && block.number >= season.startBlock && block.number <= season.endBlock;
    }

    // ============ 스테이킹 함수 ============

    /**
     * @notice 토큰 스테이킹 (분할 스테이킹 허용)
     * @param amount 스테이킹 수량
     */
    function stake(uint amount) external nonReentrant {
        _stakeFor(msg.sender, amount, msg.sender);
    }

    /**
     * @notice 다른 사용자를 위한 스테이킹 (Router 전용)
     */
    function stakeFor(address user, uint amount) external nonReentrant {
        require(approvedRouters[msg.sender], StakingPoolNotAuthorized());
        _stakeFor(user, amount, msg.sender);
    }

    /**
     * @notice 스테이킹 내부 로직
     */
    function _stakeFor(address user, uint amount, address from) internal {
        // 시즌 자동 전환 체크
        _ensureSeason();

        // 진행중인 시즌이 있어야 stake 가능
        require(isSeasonActive(), StakingPoolNoActiveSeason());

        // 이전 시즌들 lazy snapshot
        _ensureUserAllPreviousSeasons(user);

        StakePosition storage position = userStakes[user];
        uint oldBalance = position.balance;

        // 추가 예치의 경우: 현재 시즌의 포인트를 스냅샷에 누적
        if (oldBalance > 0) {
            // 1. 기존 수량으로 현재까지의 포인트 계산
            uint additionalPoints = _calculatePoints(
                position.balance,
                position.lastUpdateBlock < seasons[currentSeason].startBlock
                    ? seasons[currentSeason].startBlock
                    : position.lastUpdateBlock,
                block.number
            );

            // 2. 현재 시즌의 누적 포인트 저장
            userSeasonData[currentSeason][user].points += additionalPoints;

            // 3. position.points 리셋 (새로운 balance로 다시 시작)
            position.points = 0;
        }

        // 최소 스테이크 체크 (신규 또는 추가 후)
        uint newBalance = oldBalance + amount;
        require(newBalance >= MIN_STAKE, StakingPoolBelowMinStake());

        // 토큰 전송
        stakingToken.safeTransferFrom(from, address(this), amount);

        // 스테이킹 업데이트
        position.balance = newBalance;
        position.lastUpdateBlock = block.number;
        totalStaked += amount;

        // 현재 시즌 데이터 기록
        UserSeasonData storage seasonData = userSeasonData[currentSeason][user];
        if (seasonData.balance == 0) {
            // 첫 기록
            Season storage current = seasons[currentSeason];
            seasonData.balance = newBalance;
            seasonData.joinBlock = block.number < current.startBlock ? current.startBlock : block.number;
        } else {
            // 이미 기록된 경우 balance만 업데이트
            seasonData.balance = newBalance;
        }

        // 스테이커 목록에 추가
        if (!isStaker[user]) {
            stakers.push(user);
            isStaker[user] = true;
        }

        emit Staked(user, amount, newBalance);
    }

    /**
     * @notice 현재 시즌의 사용자 포인트 계산 (스냅샷 + 실시간)
     */
    function _calculateCurrentSeasonPoints(address user) internal view returns (uint) {
        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return 0;

        // Virtual Season 처리 (첫 시즌 또는 재시작)
        if (nextSeasonStartBlock > 0 && (currentSeason == 0 || seasons[currentSeason].isFinalized)) {
            if (block.number < nextSeasonStartBlock) return 0;

            uint posLastUpdate = position.lastUpdateBlock;
            uint seasonStart = nextSeasonStartBlock;

            // Virtual Season의 시즌 번호 계산
            uint virtualSeasonNum = currentSeason + 1;
            UserSeasonData storage virtualData = userSeasonData[virtualSeasonNum][user];

            if (posLastUpdate < seasonStart) {
                return virtualData.points + _calculatePoints(position.balance, seasonStart, block.number);
            } else {
                return virtualData.points + _calculatePoints(position.balance, posLastUpdate, block.number);
            }
        }

        // 실제 시즌이 없으면 0
        if (currentSeason == 0) return 0;

        Season storage current = seasons[currentSeason];
        UserSeasonData storage userData = userSeasonData[currentSeason][user];

        // 시즌 데이터가 있으면 사용
        if (userData.balance > 0) {
            // 스냅샷된 포인트 + 추가 포인트
            uint effectiveStart = userData.joinBlock > current.startBlock ? userData.joinBlock : current.startBlock;
            uint additionalPoints = _calculatePoints(userData.balance, effectiveStart, block.number);
            return userData.points + additionalPoints;
        }

        // 시즌 데이터가 없으면 자동 참여로 계산
        uint lastUpdate = position.lastUpdateBlock;
        if (lastUpdate < current.startBlock && position.balance > 0) {
            // 이전 시즌부터 stake → 현재 시즌 시작부터 계산
            return _calculatePoints(position.balance, current.startBlock, block.number);
        }

        return 0;
    }

    /**
     * @notice 전액 출금 (포인트 몰수)
     */
    function withdrawAll() external nonReentrant {
        _withdrawAll(msg.sender, msg.sender);
    }

    /**
     * @notice 유저를 대신해 전액 출금 (Router 전용)
     * @param user 출금할 사용자 주소
     */
    function withdrawAllFor(address user) external nonReentrant {
        require(approvedRouters[msg.sender], StakingPoolNotAuthorized());
        _withdrawAll(user, msg.sender);
    }

    /**
     * @notice 내부 출금 로직
     * @dev 현재 시즌 포인트만 몰수, 이전 시즌 보상은 claim 가능
     */
    function _withdrawAll(address user, address recipient) internal {
        _ensureSeason();

        // 이전 시즌들 lazy snapshot (출금 전에 이전 시즌 보상 확정)
        _ensureUserAllPreviousSeasons(user);

        StakePosition storage position = userStakes[user];
        require(position.balance > 0, StakingPoolNoPosition());

        uint amount = position.balance;

        // 현재 시즌 포인트 몰수
        if (currentSeason > 0) {
            UserSeasonData storage seasonData = userSeasonData[currentSeason][user];
            seasonData.points = 0;
            seasonData.balance = 0;
        }

        if (position.points > 0) totalPointsInSeason -= position.points;

        // 스테이킹 정리
        position.balance = 0;
        position.points = 0;
        position.lastUpdateBlock = block.number;
        totalStaked -= amount;

        // 토큰 반환
        stakingToken.safeTransfer(recipient, amount);

        emit WithdrawnAll(user, amount);
    }

    // ============ 시즌 관리 ============

    /**
     * @notice 유저의 특정 시즌 데이터를 lazy하게 스냅샷
     * @dev 시즌이 종료되었지만 유저 데이터가 아직 확정되지 않은 경우에만 실행
     */
    function _ensureUserSeasonSnapshot(address user, uint seasonNum) internal {
        if (seasonNum == 0 || seasonNum > currentSeason) return;

        Season storage season = seasons[seasonNum];
        if (!season.isFinalized) return; // 시즌이 아직 진행 중

        UserSeasonData storage userData = userSeasonData[seasonNum][user];
        if (userData.finalized) return; // 이미 스냅샷됨

        StakePosition storage position = userStakes[user];

        // 현재 잔액이 0이어도, 해당 시즌에는 stake가 있었을 수 있음
        // lastUpdateBlock으로 판단
        uint lastUpdate = position.lastUpdateBlock;

        // 시즌 종료 후에 업데이트되었으면, 해당 시즌 동안 stake가 있었던 것
        if (lastUpdate > season.endBlock) {
            // 현재 balance는 0이지만, 시즌 동안에는 있었을 수 있음
            // seasonData에 balance가 없으면 이전에 stake가 있었다고 가정할 수 없으므로 0으로 처리
            if (userData.balance == 0) {
                userData.finalized = true;
                return;
            }
        }

        // 현재 balance가 0이면 해당 시즌에도 참여하지 않았음
        if (position.balance == 0 && userData.balance == 0) {
            userData.finalized = true;
            return;
        }

        // seasonData에 이미 balance가 기록되어 있으면 그것을 사용 (stake 중에 기록됨)
        uint balanceToUse = userData.balance > 0 ? userData.balance : position.balance;
        uint joinBlockToUse = userData.joinBlock > 0 ? userData.joinBlock : lastUpdate;

        // 해당 시즌에 참여했는지 확인
        if (joinBlockToUse < season.startBlock) {
            // 시즌 시작 전부터 stake → 전체 시즌 참여
            userData.balance = balanceToUse;
            userData.joinBlock = season.startBlock;
            userData.points = _calculatePoints(balanceToUse, season.startBlock, season.endBlock);
        } else if (joinBlockToUse <= season.endBlock) {
            // 시즌 중간에 참여
            userData.balance = balanceToUse;
            userData.joinBlock = joinBlockToUse;
            userData.points = _calculatePoints(balanceToUse, joinBlockToUse, season.endBlock);
        }
        // else: joinBlock > season.endBlock → 시즌 종료 후 참여 → 포인트 0

        userData.finalized = true;
    }

    /**
     * @notice 유저의 이전 시즌들을 모두 스냅샷
     * @dev stake, unstake, claim 등에서 호출
     */
    function _ensureUserAllPreviousSeasons(address user) internal {
        if (currentSeason == 0) return;

        // 현재 시즌 이전의 모든 finalized 시즌을 스냅샷
        for (uint i = 1; i < currentSeason; i++) {
            _ensureUserSeasonSnapshot(user, i);
        }
    }

    /**
     * @notice 시즌 자동 전환 체크
     * @dev Lazy 시즌 생성 및 롤오버
     */
    function _ensureSeason() internal {
        // 풀이 종료되었으면 시즌 생성/롤오버 안함
        // 단, nextSeasonStartBlock > poolEndBlock 이면 재시작이므로 체크 무시
        if (poolEndBlock > 0 && (nextSeasonStartBlock == 0 || nextSeasonStartBlock <= poolEndBlock)) {
            if (block.number >= poolEndBlock) return;
        }

        // 첫 시즌 시작 (currentSeason == 0)
        if (currentSeason == 0) {
            if (block.number >= nextSeasonStartBlock) _startFirstSeason();
            return;
        }

        // 기존 롤오버 로직
        Season storage current = seasons[currentSeason];
        if (block.number > current.endBlock) {
            // 다음 시즌 시작 블록이 설정되었거나 자동으로 시작
            if (nextSeasonStartBlock == 0 || block.number >= nextSeasonStartBlock) _rolloverSeason();
        }
    }

    /**
     * @notice 첫 시즌 시작 (내부)
     */
    function _startFirstSeason() internal {
        require(currentSeason == 0, StakingPoolSeasonNotEnded());

        uint startBlock = nextSeasonStartBlock;
        uint endBlock = startBlock + seasonBlocks;

        // poolEndBlock 체크
        if (poolEndBlock > 0 && endBlock > poolEndBlock) endBlock = poolEndBlock;

        currentSeason = 1;
        seasons[1] =
            Season({seasonNumber: 1, startBlock: startBlock, endBlock: endBlock, isFinalized: false, totalPoints: 0});

        nextSeasonStartBlock = 0; // 자동 시작됨

        emit SeasonRolledOver(0, 1, 0);
    }

    /**
     * @notice 시즌 롤오버 (누구나 호출 가능)
     */
    function rolloverSeason() external {
        require(currentSeason > 0, StakingPoolNoActiveSeason());
        Season storage current = seasons[currentSeason];
        require(block.number > current.endBlock, StakingPoolSeasonNotEnded());
        require(nextSeasonStartBlock == 0 || block.number >= nextSeasonStartBlock, StakingPoolSeasonNotEnded());
        _rolloverSeason();
    }

    /**
     * @notice 시즌 롤오버 내부 로직
     * @dev 이전 시즌의 totalPoints는 lazy하게 계산 (유저별 데이터도 lazy snapshot)
     */
    function _rolloverSeason() internal {
        uint oldSeasonNumber = currentSeason;
        Season storage oldSeason = seasons[oldSeasonNumber];

        // 이전 시즌 finalize
        oldSeason.isFinalized = true;

        // ✅ 최적화: staker 순회 제거 - 각 유저가 액션할 때 lazy하게 스냅샷됨
        // totalPoints도 필요할 때 lazy하게 계산됨

        // 새 시즌으로 전환
        uint newSeasonNumber = oldSeasonNumber + 1;
        currentSeason = newSeasonNumber;
        totalPointsInSeason = 0;

        // 다음 시즌 시작 블록 계산
        uint nextStart;
        if (nextSeasonStartBlock > 0) {
            // 수동 설정된 경우
            nextStart = nextSeasonStartBlock;
            nextSeasonStartBlock = 0; // 사용 후 리셋
        } else {
            // 자동: 이전 시즌 종료 + 1블록
            nextStart = oldSeason.endBlock + 1;
        }

        uint nextEnd = nextStart + seasonBlocks;

        // poolEndBlock 체크
        // nextStart > poolEndBlock이면 재시작이므로 체크 무시
        if (poolEndBlock > 0 && nextStart <= poolEndBlock) {
            // nextEnd가 poolEndBlock을 넘으면 제한
            if (nextEnd > poolEndBlock) nextEnd = poolEndBlock;
        }

        seasons[newSeasonNumber] = Season({
            seasonNumber: newSeasonNumber,
            startBlock: nextStart,
            endBlock: nextEnd,
            isFinalized: false,
            totalPoints: 0 // lazy 계산됨
        });

        emit SeasonRolledOver(oldSeasonNumber, newSeasonNumber, 0); // totalPoints는 lazy 계산
    }

    /**
     * @notice 시즌 보상 청구 (Lazy snapshot 및 totalPoints 계산)
     */
    function claimSeason(uint seasonNumber, address rewardToken) external nonReentrant {
        Season storage season = seasons[seasonNumber];

        require(season.isFinalized, StakingPoolSeasonNotEnded());

        // 해당 시즌의 유저 데이터를 lazy snapshot
        _ensureUserSeasonSnapshot(msg.sender, seasonNumber);

        UserSeasonData storage userData = userSeasonData[seasonNumber][msg.sender];
        require(!userData.claimed, StakingPoolAlreadyClaimed());

        uint userPoints = userData.points;

        // totalPoints가 아직 계산되지 않았으면 계산
        uint totalPoints = season.totalPoints;
        if (totalPoints == 0) {
            totalPoints = _calculateSeasonTotalPoints(seasonNumber);
            season.totalPoints = totalPoints; // 캐시 저장
        }

        if (userPoints > 0 && totalPoints > 0) {
            userData.claimed = true;

            // RewardPool에 보상 지급 요청
            rewardPool.payUser(msg.sender, seasonNumber, rewardToken, userPoints, totalPoints);

            emit SeasonClaimed(msg.sender, seasonNumber, userPoints);
        }
    }

    // ============ 포인트 계산 ============

    /**
     * @notice 사용자 포인트 업데이트
     */
    function updatePoints(address user) external {
        require(
            hasRole(REWARD_POOL_ROLE, msg.sender) || hasRole(DEFAULT_ADMIN_ROLE, msg.sender), StakingPoolNotAuthorized()
        );
        _updatePoints(user);
    }

    /**
     * @notice 사용자 포인트 업데이트 (내부)
     * @dev 시즌이 넘어갔으면 이전 시즌까지 스냅샷하고 새 시즌 시작
     */
    function _updatePoints(address user) internal {
        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return;

        uint lastUpdate = position.lastUpdateBlock;
        uint currentBlock = block.number;

        if (lastUpdate >= currentBlock) return;

        // 이전 시즌들 lazy snapshot
        _ensureUserAllPreviousSeasons(user);

        Season storage current = seasons[currentSeason];

        // 현재 시즌 데이터 기록 (없으면 생성)
        UserSeasonData storage userData = userSeasonData[currentSeason][user];
        if (userData.balance == 0) {
            // 첫 기록
            uint joinBlock = lastUpdate < current.startBlock ? current.startBlock : lastUpdate;
            userData.balance = position.balance;
            userData.joinBlock = joinBlock;
        } else {
            // 이미 기록된 경우 balance만 업데이트
            userData.balance = position.balance;
        }

        // 케이스 1: 마지막 업데이트가 이전 시즌들에 걸쳐있는 경우
        if (lastUpdate < current.startBlock) {
            // 현재 시즌의 포인트 계산 (시즌 시작부터 현재까지)
            uint currentSeasonPoints = _calculatePoints(position.balance, current.startBlock, currentBlock);
            position.points = currentSeasonPoints;
            totalPointsInSeason += currentSeasonPoints;
        }
        // 케이스 2: 같은 시즌 내에서 업데이트
        else {
            uint newPoints = _calculatePoints(position.balance, lastUpdate, currentBlock);
            position.points += newPoints;
            totalPointsInSeason += newPoints;
        }

        position.lastUpdateBlock = currentBlock;
        emit PointsUpdated(user, position.points);
    }

    /**
     * @notice 포인트 계산 헬퍼 함수
     */
    function _calculatePoints(uint balance, uint fromBlock, uint toBlock) internal view returns (uint) {
        if (fromBlock >= toBlock || balance == 0) return 0;

        uint blockElapsed = toBlock - fromBlock;
        uint timeElapsed = blockElapsed * blockTime;

        return (balance * timeElapsed * POINTS_PRECISION) / pointsTimeUnit;
    }

    /**
     * @notice 시즌별 총 포인트 동적 계산 (Lazy Evaluation)
     * @dev 모든 스테이커의 포인트를 계산하여 합산 (O(n))
     */
    function _calculateSeasonTotalPoints(uint seasonNum) internal view returns (uint) {
        Season storage season = seasons[seasonNum];
        if (!season.isFinalized && seasonNum != currentSeason) return 0;

        uint total = 0;

        for (uint i = 0; i < stakers.length; i++) {
            address staker = stakers[i];
            UserSeasonData storage userData = userSeasonData[seasonNum][staker];

            // finalized된 데이터가 있으면 사용
            if (userData.finalized) {
                total += userData.points;
                continue;
            }

            // 스냅샷이 없으면 계산
            StakePosition storage position = userStakes[staker];
            if (position.balance == 0) continue;

            uint lastUpdate = position.lastUpdateBlock;

            // 현재 시즌
            if (seasonNum == currentSeason) {
                // 스냅샷된 포인트
                total += userData.points;

                // 실시간 포인트
                if (lastUpdate >= season.startBlock && lastUpdate < block.number) {
                    uint additionalPoints = _calculatePoints(position.balance, lastUpdate, block.number);
                    total += position.points + additionalPoints;
                } else if (lastUpdate < season.startBlock) {
                    total += _calculatePoints(position.balance, season.startBlock, block.number);
                }
            }
            // 과거 시즌 (lazy 계산)
            else {
                // seasonData에 balance가 기록되어 있으면 그것을 사용
                if (userData.balance > 0) {
                    uint effectiveStart =
                        userData.joinBlock > season.startBlock ? userData.joinBlock : season.startBlock;
                    total += _calculatePoints(userData.balance, effectiveStart, season.endBlock);
                } else {
                    // seasonData에 없으면 현재 position으로 추정
                    // 시즌 종료 전에 stake했고, 아직 withdraw하지 않은 경우
                    if (lastUpdate <= season.endBlock) {
                        uint effectiveStart = lastUpdate < season.startBlock ? season.startBlock : lastUpdate;
                        total += _calculatePoints(position.balance, effectiveStart, season.endBlock);
                    }
                }
            }
        }

        return total;
    }

    // ============ 조회 함수 ============

    function getStakingPower(address user) external view returns (uint) {
        return userStakes[user].balance;
    }

    function getTotalStakingPower() external view returns (uint) {
        return totalStaked;
    }

    function getUserPoints(address user) public view returns (uint) {
        return _calculateCurrentSeasonPoints(user);
    }

    function getStakePosition(address user) external view returns (uint balance, uint points, uint lastUpdateBlock) {
        StakePosition storage position = userStakes[user];

        // getUserPoints 재사용으로 중복 로직 제거
        uint currentPoints = getUserPoints(user);

        return (position.balance, currentPoints, position.lastUpdateBlock);
    }

    function getCurrentSeasonInfo()
        external
        view
        returns (uint season, uint startBlock, uint endBlock, uint blocksElapsed)
    {
        // Virtual Season 처리 (첫 시즌 또는 재시작)
        if (nextSeasonStartBlock > 0 && (currentSeason == 0 || seasons[currentSeason].isFinalized)) {
            season = currentSeason + 1; // Virtual season 번호
            startBlock = nextSeasonStartBlock;
            endBlock = startBlock + seasonBlocks;
            if (poolEndBlock > 0 && endBlock > poolEndBlock) endBlock = poolEndBlock;

            if (block.number >= startBlock) blocksElapsed = block.number - startBlock;
            return (season, startBlock, endBlock, blocksElapsed);
        }

        // 실제 시즌이 없으면 0
        if (currentSeason == 0) return (0, 0, 0, 0);

        Season storage current = seasons[currentSeason];
        season = currentSeason;
        startBlock = current.startBlock;
        endBlock = current.endBlock;
        blocksElapsed = block.number - current.startBlock;
    }

    function getSeasonUserPoints(uint seasonNumber, address user) external view returns (uint) {
        UserSeasonData storage userData = userSeasonData[seasonNumber][user];

        // finalized된 데이터가 있으면 반환
        if (userData.finalized) return userData.points;

        // 아직 finalized되지 않았으면 계산
        Season storage season = seasons[seasonNumber];
        if (!season.isFinalized) return 0; // 시즌이 아직 진행 중이면 0

        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return 0;

        uint lastUpdate = position.lastUpdateBlock;

        if (lastUpdate < season.startBlock) {
            return _calculatePoints(position.balance, season.startBlock, season.endBlock);
        } else if (lastUpdate <= season.endBlock) {
            return _calculatePoints(position.balance, lastUpdate, season.endBlock);
        }

        return 0;
    }

    function seasonTotalPointsSnapshot(uint seasonNumber) external view returns (uint) {
        Season storage season = seasons[seasonNumber];

        // 종료된 시즌은 저장된 totalPoints 반환 (lazy 계산)
        if (season.isFinalized) {
            // totalPoints가 아직 계산되지 않았으면 계산
            if (season.totalPoints == 0) return _calculateSeasonTotalPoints(seasonNumber);
            return season.totalPoints;
        }

        // 현재 시즌은 동적 계산
        if (seasonNumber == currentSeason) return _calculateSeasonTotalPoints(seasonNumber);

        return 0;
    }

    /**
     * @notice 시즌별 예상 포인트 조회 (현재 블록 기준)
     * @dev lazy snapshot된 포인트 또는 계산된 포인트를 반환
     */
    function getExpectedSeasonPoints(uint seasonNumber, address user) external view returns (uint) {
        Season storage season = seasons[seasonNumber];

        // 현재 시즌이면 getUserPoints 재사용
        if (seasonNumber == currentSeason) return this.getUserPoints(user);

        // 유효하지 않은 시즌
        if (!season.isFinalized) return 0;

        UserSeasonData storage userData = userSeasonData[seasonNumber][user];

        // finalized된 데이터가 있으면 반환
        if (userData.finalized) return userData.points;

        // 아직 finalized되지 않았으면 계산
        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return 0;

        uint lastUpdate = position.lastUpdateBlock;

        // 시즌 종료 후 업데이트되었으면 0
        if (lastUpdate > season.endBlock) return 0;

        // 시즌 범위 내 포인트 계산
        if (lastUpdate < season.startBlock) {
            return _calculatePoints(position.balance, season.startBlock, season.endBlock);
        } else {
            return _calculatePoints(position.balance, lastUpdate, season.endBlock);
        }
    }

    /**
     * @notice 시즌별 예상 리워드 조회
     * @dev RewardPool과 연동하여 예상 리워드 계산
     */
    function getExpectedSeasonReward(uint seasonNumber, address user, address rewardToken)
        external
        view
        returns (uint)
    {
        if (address(rewardPool) == address(0)) return 0;

        Season storage season = seasons[seasonNumber];

        // 현재 시즌은 아직 리워드 계산 불가
        if (seasonNumber == currentSeason) return 0;

        // 종료되지 않은 시즌은 리워드 없음
        if (!season.isFinalized) return 0;

        UserSeasonData storage userData = userSeasonData[seasonNumber][user];

        // 이미 청구했으면 0
        if (userData.claimed) return 0;

        // 사용자 포인트 조회
        uint userPoints = userData.points;
        if (userPoints == 0 && !userData.finalized) {
            // finalized되지 않았으면 계산
            StakePosition storage position = userStakes[user];
            if (position.balance == 0) return 0;

            uint lastUpdate = position.lastUpdateBlock;
            if (lastUpdate > season.endBlock) return 0;

            if (lastUpdate < season.startBlock) {
                userPoints = _calculatePoints(position.balance, season.startBlock, season.endBlock);
            } else {
                userPoints = _calculatePoints(position.balance, lastUpdate, season.endBlock);
            }
        }

        if (userPoints == 0) return 0;

        // totalPoints 조회 (lazy 계산)
        uint totalPoints = season.totalPoints;
        if (totalPoints == 0) totalPoints = _calculateSeasonTotalPoints(seasonNumber);

        if (totalPoints == 0) return 0;

        // RewardPool에서 시즌 리워드 조회
        return rewardPool.getExpectedReward(user, seasonNumber, rewardToken);
    }
}
