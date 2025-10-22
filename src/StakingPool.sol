// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./base/StakingPoolBase.sol";
import "./interfaces/IRewardPool.sol";
import "./interfaces/IStakingProtocol.sol";
import "./libraries/PointsLib.sol";

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
contract StakingPool is StakingPoolBase {
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
        uint _poolEndBlock
    ) StakingPoolBase(_stakingToken, address(_protocol), _seasonBlocks, _firstSeasonStartBlock, _poolEndBlock) {
        projectID = _projectID;
        protocol = _protocol;
    }

    // ============================================
    // Season Functions
    // ============================================

    /**
     * @notice 시즌 롤오버
     */
    function rolloverSeason() external {
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
        Season storage season = seasons[seasonNumber];
        require(season.isFinalized, StakingPoolBaseSeasonNotEnded());

        _ensureUserSeasonSnapshot(msg.sender, seasonNumber);

        UserSeasonData storage userData = userSeasonData[seasonNumber][msg.sender];
        require(!userData.claimed, StakingPoolAlreadyClaimed());

        uint userPoints = userData.points;

        uint totalPoints = season.totalPoints;
        if (totalPoints == 0) {
            totalPoints = _calculateSeasonTotalPoints(seasonNumber);
            season.totalPoints = totalPoints;
        }

        if (userPoints > 0 && totalPoints > 0) {
            userData.claimed = true;
            rewardPool.payUser(msg.sender, seasonNumber, rewardToken, userPoints, totalPoints);
            emit SeasonClaimed(msg.sender, seasonNumber, userPoints);
        }
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
        } else {
            userData.balance = position.balance;
        }

        if (lastUpdate < current.startBlock) {
            uint currentSeasonPoints =
                PointsLib.calculatePoints(position.balance, current.startBlock, currentBlock, blockTime, pointsTimeUnit);
            position.points = currentSeasonPoints;
        } else {
            uint newPoints =
                PointsLib.calculatePoints(position.balance, lastUpdate, currentBlock, blockTime, pointsTimeUnit);
            position.points += newPoints;
        }

        position.lastUpdateBlock = currentBlock;
        emit PointsUpdated(user, position.points);
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
            uint effectiveStart = userData.joinBlock > current.startBlock ? userData.joinBlock : current.startBlock;
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
            if (season.totalPoints > 0) return season.totalPoints;
            return season.aggregatedPoints;
        }

        // 현재 시즌: 집계된 포인트 + 마지막 집계 이후 포인트
        if (seasonNum == currentSeason) {
            uint blocksSinceAggregation =
                block.number > season.lastAggregatedBlock ? block.number - season.lastAggregatedBlock : 0;

            if (blocksSinceAggregation == 0) return season.aggregatedPoints;

            uint additionalPoints = PointsLib.calculatePoints(
                season.seasonTotalStaked, season.lastAggregatedBlock, block.number, blockTime, pointsTimeUnit
            );

            return season.aggregatedPoints + additionalPoints;
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
        if (nextSeasonStartBlock > 0 && (currentSeason == 0 || seasons[currentSeason].isFinalized)) {
            season = currentSeason + 1;
            startBlock = nextSeasonStartBlock;
            endBlock = startBlock + seasonBlocks;
            if (poolEndBlock > 0 && endBlock > poolEndBlock) endBlock = poolEndBlock;

            if (block.number >= startBlock) blocksElapsed = block.number - startBlock;
            return (season, startBlock, endBlock, blocksElapsed);
        }

        if (currentSeason == 0) return (0, 0, 0, 0);

        Season storage current = seasons[currentSeason];
        season = currentSeason;
        startBlock = current.startBlock;
        endBlock = current.endBlock;
        blocksElapsed = block.number - current.startBlock;
    }

    function getSeasonUserPoints(uint seasonNumber, address user) external view returns (uint) {
        UserSeasonData storage userData = userSeasonData[seasonNumber][user];

        if (userData.finalized) return userData.points;

        Season storage season = seasons[seasonNumber];
        if (!season.isFinalized) return 0;

        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return 0;

        uint lastUpdate = position.lastUpdateBlock;

        if (lastUpdate < season.startBlock) {
            return PointsLib.calculatePoints(
                position.balance, season.startBlock, season.endBlock, blockTime, pointsTimeUnit
            );
        } else if (lastUpdate <= season.endBlock) {
            return PointsLib.calculatePoints(position.balance, lastUpdate, season.endBlock, blockTime, pointsTimeUnit);
        }

        return 0;
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

    function getExpectedSeasonPoints(uint seasonNumber, address user) external view returns (uint) {
        Season storage season = seasons[seasonNumber];

        if (seasonNumber == currentSeason) return this.getUserPoints(user);

        if (!season.isFinalized) return 0;

        UserSeasonData storage userData = userSeasonData[seasonNumber][user];

        if (userData.finalized) return userData.points;

        StakePosition storage position = userStakes[user];
        if (position.balance == 0) return 0;

        uint lastUpdate = position.lastUpdateBlock;

        if (lastUpdate > season.endBlock) return 0;

        if (lastUpdate < season.startBlock) {
            return PointsLib.calculatePoints(
                position.balance, season.startBlock, season.endBlock, blockTime, pointsTimeUnit
            );
        } else {
            return PointsLib.calculatePoints(position.balance, lastUpdate, season.endBlock, blockTime, pointsTimeUnit);
        }
    }

    function getExpectedSeasonReward(uint seasonNumber, address user, address rewardToken)
        external
        view
        returns (uint)
    {
        if (address(rewardPool) == address(0)) return 0;

        Season storage season = seasons[seasonNumber];

        if (seasonNumber == currentSeason) return 0;

        if (!season.isFinalized) return 0;

        UserSeasonData storage userData = userSeasonData[seasonNumber][user];

        if (userData.claimed) return 0;

        uint userPoints = userData.points;
        if (userPoints == 0 && !userData.finalized) {
            StakePosition storage position = userStakes[user];
            if (position.balance == 0) return 0;

            uint lastUpdate = position.lastUpdateBlock;
            if (lastUpdate > season.endBlock) return 0;

            if (lastUpdate < season.startBlock) {
                userPoints = PointsLib.calculatePoints(
                    position.balance, season.startBlock, season.endBlock, blockTime, pointsTimeUnit
                );
            } else {
                userPoints =
                    PointsLib.calculatePoints(position.balance, lastUpdate, season.endBlock, blockTime, pointsTimeUnit);
            }
        }

        if (userPoints == 0) return 0;

        uint totalPoints = season.totalPoints;
        if (totalPoints == 0) totalPoints = _calculateSeasonTotalPoints(seasonNumber);

        if (totalPoints == 0) return 0;

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

        userPoints = userData.finalized ? userData.points : this.getExpectedSeasonPoints(seasonNumber, user);
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
     * @notice 유저의 이전 시즌들을 배치 처리 (외부 호출 가능)
     * @dev 최적화 설명:
     *      - 각 시즌은 개별적으로 UserSeasonData에 저장되어야 함 (개별 청구를 위해)
     *      - 연속 시즌 최적화는 불가능: 각 시즌마다 startBlock/endBlock이 다르고 개별 포인트 필요
     *      - 현재 최적화: lastFinalizedSeason을 루프 후 한 번만 업데이트
     *      - 완전 지연 평가: 청구하지 않는 시즌은 영원히 처리 안 함
     *      - 가스 효율: maxSeasons로 가스 한도 내에서 처리
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
    // Configuration Functions
    // ============================================

    /**
     * @notice RewardPool 설정 (단 한 번만 가능)
     * @dev 보안상 이유로 재설정 불가능:
     *      - 기존 보상 데이터 손실 방지
     *      - 사용자 청구 기록 보호
     *      - 악의적인 RewardPool 교체 방지
     *      ⚠️  신중하게 올바른 RewardPool 주소를 설정해야 함
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
