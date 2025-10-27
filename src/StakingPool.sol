// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./base/StakingPoolBase.sol";

import "./interfaces/IRewardPool.sol";
import "./interfaces/IStakingProtocol.sol";
import "./libraries/PointsLib.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";

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
 * @notice 프로젝트별 $CROSS 스테이킹 풀 (Base 상속 버전)
 * @dev StakingPoolBase를 상속하여 추가 기능 구현
 */
contract StakingPool is StakingPoolBase, Pausable {
    // ============================================
    // Errors
    // ============================================

    error StakingPoolAlreadyClaimed();

    // ============================================
    // State Variables
    // ============================================

    /// @notice 프로젝트 ID
    uint public immutable projectID;

    /// @notice 프로토콜 컨트랙트
    IStakingProtocol public immutable protocol;

    /// @notice 연결된 RewardPool
    IRewardPool public rewardPool;

    // ============================================
    // Events
    // ============================================

    event RewardPoolSet(IRewardPool indexed rewardPool);
    event RouterApprovalSet(address indexed router, bool approved);
    event PointsTimeUnitUpdated(uint oldValue, uint newValue);
    event BlockTimeUpdated(uint oldValue, uint newValue);
    event NextSeasonStartUpdated(uint oldValue, uint newValue);
    event PoolEndBlockUpdated(uint oldValue, uint newValue);

    // ============================================
    // Constructor
    // ============================================

    constructor(
        uint _projectID,
        IERC20 _stakingToken,
        IStakingProtocol _protocol,
        uint _seasonBlocks,
        uint _firstSeasonStartBlock,
        uint _poolEndBlock,
        uint _preDepositStartBlock
    )
        StakingPoolBase(
            _stakingToken,
            address(_protocol),
            _seasonBlocks,
            _firstSeasonStartBlock,
            _poolEndBlock,
            _preDepositStartBlock
        )
    {
        projectID = _projectID;
        protocol = _protocol;
    }

    // ============================================
    // Pausable Functions
    // ============================================

    /**
     * @notice 긴급 중지
     */
    function pause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
    }

    /**
     * @notice 긴급 중지 해제
     */
    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
    }

    // ============================================
    // Override Functions with Pausable
    // ============================================

    /**
     * @notice 스테이킹 (pausable)
     */
    function stake(uint amount) external override nonReentrant whenNotPaused {
        _stakeFor(msg.sender, amount, msg.sender);
    }

    /**
     * @notice 대리 스테이킹 (pausable)
     */
    function stakeFor(address user, uint amount) external override nonReentrant whenNotPaused {
        require(_isApprovedRouter(msg.sender), CrossStakingBaseNotAuthorized());
        _stakeFor(user, amount, msg.sender);
    }

    /**
     * @notice 전액 출금 (pausable)
     */
    function withdrawAll() external override nonReentrant whenNotPaused {
        _withdrawAll(msg.sender, msg.sender);
    }

    /**
     * @notice 대리 전액 출금 (pausable)
     * @dev Router가 호출 시 토큰을 router에게 전송 (router가 unwrap 처리)
     */
    function withdrawAllFor(address user) external override nonReentrant whenNotPaused {
        require(_isApprovedRouter(msg.sender), CrossStakingBaseNotAuthorized());
        _withdrawAll(user, msg.sender); // recipient를 router(msg.sender)로 설정
    }

    // ============================================
    // Season Functions
    // ============================================

    /**
     * @notice 시즌 롤오버
     */
    function rolloverSeason() external override {
        require(currentSeason != 0, StakingPoolBaseNoActiveSeason());
        Season storage current = seasons[currentSeason];
        require(block.number > current.endBlock, StakingPoolBaseSeasonNotEnded());
        require(nextSeasonStartBlock == 0 || block.number >= nextSeasonStartBlock, StakingPoolBaseSeasonNotEnded());
        _rolloverSeason();
    }

    /**
     * @notice 시즌 보상 청구
     */
    function claimSeason(uint seasonNumber, address rewardToken) external nonReentrant {
        _claimSeasonFor(msg.sender, seasonNumber, rewardToken);
    }

    /**
     * @notice 라우터가 사용자를 대신해 시즌 보상 청구
     */
    function claimSeasonFor(address user, uint seasonNumber, address rewardToken) external nonReentrant {
        require(_isApprovedRouter(msg.sender), CrossStakingBaseNotAuthorized());
        _claimSeasonFor(user, seasonNumber, rewardToken);
    }

    /**
     * @notice 사용자 포인트 업데이트
     */
    function updatePoints(address user) external {
        require(
            hasRole(REWARD_POOL_ROLE, msg.sender) || hasRole(DEFAULT_ADMIN_ROLE, msg.sender),
            CrossStakingBaseNotAuthorized()
        );
        _updatePoints(user);
    }

    // ============================================
    // Internal Functions
    // ============================================

    /**
     * @notice 사용자 포인트 업데이트 (내부)
     */
    function _updatePoints(address user) internal {
        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return;

        uint lastUpdate = position.lastUpdateBlock;
        uint currentBlock = block.number;

        if (lastUpdate >= currentBlock) return;

        _ensureUserAllPreviousSeasons(user);

        Season storage current = seasons[currentSeason];

        UserSeasonData storage userData = userSeasonData[currentSeason][user];
        if (userData.balance == 0) {
            uint joinBlock = lastUpdate < current.startBlock ? current.startBlock : lastUpdate;
            userData.balance = position.balance;
            userData.joinBlock = joinBlock;
            userData.lastPointsBlock = currentBlock;
        } else {
            userData.balance = position.balance;
        }

        uint effectiveStart = lastUpdate < current.startBlock ? current.startBlock : lastUpdate;
        uint newPoints =
            PointsLib.calculatePoints(position.balance, effectiveStart, currentBlock, blockTime, pointsTimeUnit);

        userData.points += newPoints;
        userData.lastPointsBlock = currentBlock;

        position.points = 0;
        position.lastUpdateBlock = currentBlock;
        emit PointsUpdated(user, userData.points);
    }

    /**
     * @notice 내부 공통 로직: 특정 사용자의 시즌 보상 청구 처리
     */
    function _claimSeasonFor(address user, uint seasonNumber, address rewardToken) internal {
        Season storage season = seasons[seasonNumber];
        require(season.isFinalized, StakingPoolBaseSeasonNotEnded());

        _ensureUserSeasonSnapshot(user, seasonNumber);

        UserSeasonData storage userData = userSeasonData[seasonNumber][user];
        require(!userData.claimed, StakingPoolAlreadyClaimed());

        uint userPoints = userData.points;

        uint totalPoints = season.totalPoints;
        if (totalPoints == 0) {
            totalPoints = _calculateSeasonTotalPoints(seasonNumber);
            season.totalPoints = totalPoints;
        }

        if (userPoints > 0 && totalPoints > 0) {
            userData.claimed = true;
            rewardPool.payUser(user, seasonNumber, rewardToken, userPoints, totalPoints);
            emit SeasonClaimed(user, seasonNumber, userPoints);
        }
    }

    /**
     * @notice 현재 시즌의 사용자 포인트 계산
     */
    function _calculateCurrentSeasonPoints(address user) internal view returns (uint) {
        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return 0;

        if (nextSeasonStartBlock > 0 && (currentSeason == 0 || seasons[currentSeason].isFinalized)) {
            if (block.number < nextSeasonStartBlock) return 0;

            uint posLastUpdate = position.lastUpdateBlock;
            uint seasonStart = nextSeasonStartBlock;

            uint virtualSeasonNum = currentSeason + 1;
            UserSeasonData storage virtualData = userSeasonData[virtualSeasonNum][user];

            if (posLastUpdate < seasonStart) {
                return virtualData.points
                    + PointsLib.calculatePoints(position.balance, seasonStart, block.number, blockTime, pointsTimeUnit);
            } else {
                return virtualData.points
                    + PointsLib.calculatePoints(position.balance, posLastUpdate, block.number, blockTime, pointsTimeUnit);
            }
        }

        if (currentSeason == 0) return 0;

        Season storage current = seasons[currentSeason];
        UserSeasonData storage userData = userSeasonData[currentSeason][user];

        if (userData.balance > 0) {
            // lastPointsBlock부터 추가 포인트 계산 (이중 계산 방지)
            uint effectiveStart = userData.lastPointsBlock > 0 ? userData.lastPointsBlock : userData.joinBlock;
            if (effectiveStart < current.startBlock) effectiveStart = current.startBlock;

            uint additionalPoints =
                PointsLib.calculatePoints(userData.balance, effectiveStart, block.number, blockTime, pointsTimeUnit);
            return userData.points + additionalPoints;
        }

        uint lastUpdate = position.lastUpdateBlock;
        if (lastUpdate < current.startBlock && position.balance > 0) {
            return
                PointsLib.calculatePoints(position.balance, current.startBlock, block.number, blockTime, pointsTimeUnit);
        }

        return 0;
    }

    /**
     * @notice 시즌별 총 포인트 계산 (O(1) 최적화 버전)
     */
    function _calculateSeasonTotalPoints(uint seasonNum) internal view returns (uint) {
        Season storage season = seasons[seasonNum];
        if (!season.isFinalized && seasonNum != currentSeason) return 0;

        // 종료된 시즌: 캐시된 totalPoints 반환
        if (season.isFinalized) {
            if (season.totalPoints > 0) {
                // forfeitedPoints 차감 (몰수된 포인트 제외)
                return season.totalPoints > season.forfeitedPoints ? season.totalPoints - season.forfeitedPoints : 0;
            }
            // aggregatedPoints에서 forfeitedPoints 차감
            return
                season.aggregatedPoints > season.forfeitedPoints ? season.aggregatedPoints - season.forfeitedPoints : 0;
        }

        // 현재 시즌: 집계된 포인트 + 마지막 집계 이후 포인트 - 몰수된 포인트
        if (seasonNum == currentSeason) {
            uint blocksSinceAggregation =
                block.number > season.lastAggregatedBlock ? block.number - season.lastAggregatedBlock : 0;

            uint totalAggregated = season.aggregatedPoints;

            if (blocksSinceAggregation > 0) {
                uint additionalPoints = PointsLib.calculatePoints(
                    season.seasonTotalStaked, season.lastAggregatedBlock, block.number, blockTime, pointsTimeUnit
                );
                totalAggregated += additionalPoints;
            }

            // forfeitedPoints 차감
            return totalAggregated > season.forfeitedPoints ? totalAggregated - season.forfeitedPoints : 0;
        }

        return 0;
    }

    // ============================================
    // View Functions
    // ============================================

    function getStakePosition(address user) external view returns (uint balance, uint points, uint lastUpdateBlock) {
        StakePosition storage position = userStakes[user];
        uint currentPoints = getUserPoints(user);
        return (position.balance, currentPoints, position.lastUpdateBlock);
    }

    function getCurrentSeasonInfo()
        external
        view
        returns (uint season, uint startBlock, uint endBlock, uint blocksElapsed)
    {
        // Case 1: No season started yet, but next season is scheduled
        if (currentSeason == 0 && nextSeasonStartBlock > 0) {
            if (block.number >= nextSeasonStartBlock) {
                // Calculate which season we're actually in based on elapsed blocks
                uint blocksSinceStart = block.number - nextSeasonStartBlock;
                uint seasonIndex = blocksSinceStart / seasonBlocks; // 0-based index

                season = seasonIndex + 1; // Convert to 1-based season number
                startBlock = nextSeasonStartBlock + (seasonIndex * seasonBlocks);
                endBlock = _calculateEndBlock(startBlock);
                blocksElapsed = block.number - startBlock;
            } else {
                // Season not started yet
                return (0, 0, 0, 0);
            }
            return (season, startBlock, endBlock, blocksElapsed);
        }

        // Case 2: No season at all
        if (currentSeason == 0) return (0, 0, 0, 0);

        // Case 3: Season exists - calculate actual current season based on block number
        Season storage current = seasons[currentSeason];

        // Check if we're still in the current season
        if (block.number <= current.endBlock) {
            // Still in current season
            season = currentSeason;
            startBlock = current.startBlock;
            endBlock = current.endBlock;
            blocksElapsed = block.number - startBlock;
        } else {
            // Current season ended, calculate which season we're actually in
            // Each season starts at: current.startBlock + (n * seasonBlocks)
            // where n = 0, 1, 2, 3...

            uint blocksSinceFirstSeason = block.number - current.startBlock;
            uint seasonIndex = blocksSinceFirstSeason / seasonBlocks;

            // Calculate the actual season number
            season = currentSeason + seasonIndex;

            // Calculate start and end blocks for this season
            startBlock = current.startBlock + (seasonIndex * seasonBlocks);
            endBlock = startBlock + seasonBlocks;

            blocksElapsed = block.number - startBlock;
        }
    }

    function getSeasonUserPoints(uint seasonNumber, address user)
        external
        view
        returns (uint userPoints, uint totalPoints)
    {
        UserSeasonData storage userData = userSeasonData[seasonNumber][user];

        // 이미 finalize된 유저 데이터가 있으면 반환
        if (userData.finalized) return (userData.points, this.seasonTotalPointsSnapshot(seasonNumber));

        // 현재 시즌이면 실시간 포인트 반환
        if (seasonNumber == currentSeason && currentSeason > 0) {
            return (getUserPoints(user), this.seasonTotalPointsSnapshot(seasonNumber));
        }

        // 과거 시즌: finalize되지 않았어도 계산해서 반환
        Season storage season = seasons[seasonNumber];

        // 시즌이 존재하지 않으면 getExpectedSeasonPoints 사용 (가상 시즌 계산)
        if (season.startBlock == 0) return (getExpectedSeasonPoints(seasonNumber, user), 0);

        // 아직 시작하지 않은 경우
        if (block.number < season.startBlock) return (0, 0);

        // 과거 시즌에서 이미 저장된 포인트가 있으면 사용
        if (userData.points > 0) return (userData.points, this.seasonTotalPointsSnapshot(seasonNumber));

        // 과거 시즌에서 해당 시즌의 잔액이 있으면 그것으로 계산
        if (userData.balance > 0) {
            uint effectiveStart = userData.joinBlock > season.startBlock ? userData.joinBlock : season.startBlock;
            userPoints =
                PointsLib.calculatePoints(userData.balance, effectiveStart, season.endBlock, blockTime, pointsTimeUnit);
            return (userPoints, this.seasonTotalPointsSnapshot(seasonNumber));
        }

        // userData가 없으면 position으로 계산 (자동 참여 케이스)
        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return (0, this.seasonTotalPointsSnapshot(seasonNumber));

        uint lastUpdate = position.lastUpdateBlock;

        // 시즌 종료 후 스테이킹한 경우
        if (lastUpdate > season.endBlock) return (0, this.seasonTotalPointsSnapshot(seasonNumber));

        // 시즌 시작 전에 스테이킹한 경우: 시즌 전체 기간 계산
        if (lastUpdate < season.startBlock) {
            userPoints = PointsLib.calculatePoints(
                position.balance, season.startBlock, season.endBlock, blockTime, pointsTimeUnit
            );
            return (userPoints, this.seasonTotalPointsSnapshot(seasonNumber));
        }

        // 시즌 중간에 스테이킹한 경우: lastUpdate부터 시즌 종료까지 계산
        userPoints = PointsLib.calculatePoints(position.balance, lastUpdate, season.endBlock, blockTime, pointsTimeUnit);
        return (userPoints, this.seasonTotalPointsSnapshot(seasonNumber));
    }

    function seasonTotalPointsSnapshot(uint seasonNumber) external view returns (uint) {
        Season storage season = seasons[seasonNumber];

        if (season.isFinalized) {
            if (season.totalPoints == 0) return _calculateSeasonTotalPoints(seasonNumber);
            return season.totalPoints;
        }

        if (seasonNumber == currentSeason) return _calculateSeasonTotalPoints(seasonNumber);

        return 0;
    }

    function getExpectedSeasonPoints(uint seasonNumber, address user) public view returns (uint) {
        Season storage season = seasons[seasonNumber];

        if (seasonNumber == currentSeason) return getUserPoints(user);

        if (!season.isFinalized) return 0;

        UserSeasonData storage userData = userSeasonData[seasonNumber][user];

        if (userData.finalized) return userData.points;

        // userData에 잔액이 있으면 그것으로 계산
        if (userData.balance > 0) {
            uint basePoints = userData.points;
            uint effectiveStart = userData.joinBlock > season.startBlock ? userData.joinBlock : season.startBlock;

            // 전체 기간의 포인트 계산
            uint totalPoints =
                PointsLib.calculatePoints(userData.balance, effectiveStart, season.endBlock, blockTime, pointsTimeUnit);

            // basePoints는 중간에 스냅샷된 값일 수 있으므로, 더 큰 값을 사용
            return totalPoints > basePoints ? totalPoints : basePoints;
        }

        // userData가 없으면 position으로 계산 (자동 참여 케이스)
        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return 0;

        uint lastUpdate = position.lastUpdateBlock;

        // 시즌 종료 후 스테이킹한 경우
        if (lastUpdate > season.endBlock) return 0;

        // 시즌 시작 전에 스테이킹한 경우: 시즌 전체 기간 계산
        if (lastUpdate < season.startBlock) {
            return PointsLib.calculatePoints(
                position.balance, season.startBlock, season.endBlock, blockTime, pointsTimeUnit
            );
        }

        // 시즌 중간에 스테이킹한 경우: lastUpdate부터 시즌 종료까지 계산
        return PointsLib.calculatePoints(position.balance, lastUpdate, season.endBlock, blockTime, pointsTimeUnit);
    }

    /**
     * @notice 시즌의 예상 보상 조회
     * @dev RewardPool의 getExpectedReward를 호출하여 계산
     * @param seasonNumber 시즌 번호
     * @param user 사용자 주소
     * @param rewardToken 보상 토큰 주소
     * @return 예상 보상 수량
     */
    function getExpectedSeasonReward(uint seasonNumber, address user, address rewardToken)
        external
        view
        returns (uint)
    {
        if (address(rewardPool) == address(0)) return 0;

        // 현재 시즌은 아직 종료되지 않았으므로 예상 보상 계산 불가
        if (seasonNumber == currentSeason) return 0;

        // RewardPool의 getExpectedReward를 직접 호출
        // RewardPool이 내부적으로 포인트, 총 포인트, 보상 풀 등을 확인하여 계산
        return rewardPool.getExpectedReward(user, seasonNumber, rewardToken);
    }

    function getUserPoints(address user) public view returns (uint) {
        return _calculateCurrentSeasonPoints(user);
    }

    /**
     * @notice 사용자의 시즌 데이터 조회
     */
    function getUserSeasonData(uint seasonNumber, address user)
        external
        view
        returns (uint points, uint balance, uint joinBlock, bool claimed, bool finalized)
    {
        UserSeasonData storage data = userSeasonData[seasonNumber][user];
        return (data.points, data.balance, data.joinBlock, data.claimed, data.finalized);
    }

    /**
     * @notice 시즌 청구 미리보기
     */
    function previewClaim(uint seasonNumber, address user, address rewardToken)
        external
        view
        returns (uint userPoints, uint totalPoints, uint expectedReward, bool alreadyClaimed, bool canClaim)
    {
        if (address(rewardPool) == address(0)) return (0, 0, 0, false, false);

        Season storage season = seasons[seasonNumber];
        canClaim = season.isFinalized;

        UserSeasonData storage userData = userSeasonData[seasonNumber][user];
        alreadyClaimed = userData.claimed;

        userPoints = userData.finalized ? userData.points : getExpectedSeasonPoints(seasonNumber, user);
        totalPoints = season.totalPoints > 0 ? season.totalPoints : this.seasonTotalPointsSnapshot(seasonNumber);

        if (userPoints > 0 && totalPoints > 0 && !alreadyClaimed) {
            expectedReward = rewardPool.getExpectedReward(user, seasonNumber, rewardToken);
        }
    }

    /**
     * @notice 스테이커 목록 페이지네이션 조회
     */
    function getStakers(uint offset, uint limit) external view returns (address[] memory stakerList, uint total) {
        total = stakers.length;
        if (offset >= total) return (new address[](0), total);

        uint end = offset + limit;
        if (end > total) end = total;

        uint count = end - offset;
        stakerList = new address[](count);

        for (uint i = 0; i < count; i++) {
            stakerList[i] = stakers[offset + i];
        }
    }

    /**
     * @notice 유저의 이전 시즌들을 배치 처리
     * @dev 가스 효율을 위해 maxSeasons로 한도 제한
     */
    function finalizeUserSeasonsBatch(address user, uint maxSeasons) external returns (uint processed) {
        if (currentSeason == 0) return 0;

        uint startSeason = lastFinalizedSeason[user];
        if (startSeason == 0) startSeason = 1;

        uint endSeason = currentSeason > 0 ? currentSeason - 1 : 0;
        if (startSeason > endSeason) return 0;

        processed = 0;
        uint lastProcessed = startSeason - 1;

        for (uint i = startSeason; i <= endSeason && processed < maxSeasons;) {
            _ensureUserSeasonSnapshot(user, i);
            lastProcessed = i;
            unchecked {
                ++i;
                ++processed;
            }
        }

        // 루프 끝에서 한 번만 업데이트 (가스 절약)
        if (processed > 0) lastFinalizedSeason[user] = lastProcessed;

        return processed;
    }

    // ============================================
    // Router Approval Check (Override)
    // ============================================

    /**
     * @notice Router 승인 여부 확인 (Protocol 글로벌 승인 + 로컬 승인)
     * @param router Router 주소
     * @return approved 글로벌 또는 로컬에서 승인되면 true
     */
    function _isApprovedRouter(address router) internal view override returns (bool) {
        return protocol.isGlobalApprovedRouter(router) || hasRole(ROUTER_ROLE, router);
    }

    // ============================================
    // Configuration Functions
    // ============================================

    /**
     * @notice RewardPool 설정 (단 한 번만 가능, 재설정 불가)
     */
    function setRewardPool(IRewardPool _rewardPool) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(address(rewardPool) == address(0), CrossStakingBaseAlreadySet());
        _validateAddress(address(_rewardPool));
        rewardPool = _rewardPool;
        _grantRole(REWARD_POOL_ROLE, address(_rewardPool));
        emit RewardPoolSet(_rewardPool);
    }

    function setApprovedRouter(address router, bool approved) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (approved) _grantRole(ROUTER_ROLE, router);
        else _revokeRole(ROUTER_ROLE, router);
        emit RouterApprovalSet(router, approved);
    }

    function setPointsTimeUnit(uint _timeUnit) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_timeUnit != 0, StakingPoolBaseInvalidTimeUnit());
        uint oldValue = pointsTimeUnit;
        pointsTimeUnit = _timeUnit;
        emit PointsTimeUnitUpdated(oldValue, _timeUnit);
    }

    function setBlockTime(uint _blockTime) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_blockTime != 0, StakingPoolBaseInvalidTimeUnit());
        uint oldValue = blockTime;
        blockTime = _blockTime;
        emit BlockTimeUpdated(oldValue, _blockTime);
    }

    function setNextSeasonStart(uint _startBlock) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_startBlock == 0 || _startBlock > block.number, StakingPoolBaseInvalidStartBlock());
        uint oldValue = nextSeasonStartBlock;
        nextSeasonStartBlock = _startBlock;
        emit NextSeasonStartUpdated(oldValue, _startBlock);
    }

    function setPoolEndBlock(uint _endBlock) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_endBlock == 0 || _endBlock > block.number, StakingPoolBaseInvalidEndBlock());
        uint oldValue = poolEndBlock;
        poolEndBlock = _endBlock;
        emit PoolEndBlockUpdated(oldValue, _endBlock);
    }
}
