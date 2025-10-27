// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./interfaces/IRewardPool.sol";
import "./interfaces/IStakingPool.sol";
import "./interfaces/IStakingProtocol.sol";
import "./libraries/PointsLib.sol";

/**
 * @title StakingViewer
 * @notice Router 등에서 사용하던 조회 전용 기능을 모은 뷰어 컨트랙트
 */
contract StakingViewer {
    using PointsLib for *;

    IStakingProtocol public immutable protocol;

    constructor(address _protocol) {
        require(_protocol != address(0), "Invalid protocol");
        protocol = IStakingProtocol(_protocol);
    }

    function _getPools(uint projectID) internal view returns (IStakingPool pool, IRewardPool rewardPool) {
        (address stakingPool, address rewardPoolAddr,,,,,) = protocol.projects(projectID);
        pool = IStakingPool(stakingPool);
        rewardPool = IRewardPool(rewardPoolAddr);
    }

    function _calcFirstSeasonStart(IStakingPool pool) internal view returns (bool ok, uint firstStart) {
        (uint season, uint startBlock,,) = pool.getCurrentSeasonInfo();
        if (season > 0 && startBlock > 0) {
            uint sb = pool.seasonBlocks();
            // season은 1-based
            firstStart = startBlock - ((season - 1) * sb);
            return (true, firstStart);
        }
        // 시즌이 아직 시작 전인 경우(nextSeasonStartBlock 기반)
        uint nextStart = pool.nextSeasonStartBlock();
        if (nextStart > 0) return (true, nextStart);
        return (false, 0);
    }

    function _calcSeasonRange(IStakingPool pool, uint season)
        internal
        view
        returns (bool ok, uint startBlock, uint endBlock)
    {
        (bool okStart, uint firstStart) = _calcFirstSeasonStart(pool);
        if (!okStart || season == 0) return (false, 0, 0);

        uint sb = pool.seasonBlocks();
        startBlock = firstStart + ((season - 1) * sb);
        endBlock = startBlock + sb - 1;

        uint pe = pool.poolEndBlock();
        if (pe > 0) {
            if (startBlock > pe) return (false, 0, 0);
            if (endBlock > pe) endBlock = pe;
        }
        return (true, startBlock, endBlock);
    }

    /**
     * @notice 가상 시즌 데이터 계산 (온체인에 없는 시즌)
     * @param pool StakingPool 컨트랙트
     * @param season 시즌 번호
     * @param user 사용자 주소 (address(0)이면 totalPoints만 계산)
     * @return userPoints 유저 포인트
     * @return totalPoints 토탈 포인트
     */
    function _calculateVirtualSeasonData(IStakingPool pool, uint season, address user)
        internal
        view
        returns (uint userPoints, uint totalPoints)
    {
        (bool ok, uint startBlock, uint endBlock) = _calcSeasonRange(pool, season);
        if (!ok) return (0, 0);

        if (block.number < startBlock) return (0, 0);

        uint currentSeason = pool.currentSeason();

        // 과거 시즌: endBlock 사용 (종료된 시즌은 고정)
        // 현재/미래 시즌: 현재 블록까지만 반영
        uint toBlock = (season < currentSeason) ? endBlock : (block.number < endBlock ? block.number : endBlock);

        uint blockTime = pool.blockTime();
        uint timeUnit = pool.pointsTimeUnit();

        // 유저 포인트 계산 (user가 address(0)이 아닌 경우만)
        if (user != address(0)) {
            (uint balance,, uint lastUpdate) = pool.getStakePosition(user);

            if (balance > 0) {
                // 유저의 시작 블록: lastUpdate와 시즌 시작 블록 중 큰 값
                uint fromBlock = lastUpdate > startBlock ? lastUpdate : startBlock;

                // lastUpdate가 시즌 종료 후면 참여 안 함
                if (fromBlock <= toBlock) {
                    userPoints = PointsLib.calculatePoints(balance, fromBlock, toBlock, blockTime, timeUnit);
                }
            }
        }

        // 토탈 포인트 계산: 현재 totalStaked 기준
        uint totalStaked = pool.totalStaked();
        if (totalStaked > 0) {
            totalPoints = PointsLib.calculatePoints(totalStaked, startBlock, toBlock, blockTime, timeUnit);
        }

        return (userPoints, totalPoints);
    }

    // =====================
    // 조회 함수
    // =====================

    function getStakeInfo(uint projectID, address user)
        external
        view
        returns (uint balance, uint points, uint lastUpdateBlock)
    {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getStakePosition(user);
    }

    function getCurrentSeason(uint projectID) external view returns (uint season) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.currentSeason();
    }

    function getUserPoints(uint projectID, address user) external view returns (uint points) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getUserPoints(user);
    }

    function getStakingPower(uint projectID, address user) external view returns (uint stakingPower) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getStakingPower(user);
    }

    function getTotalStakingPower(uint projectID) external view returns (uint totalPower) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getTotalStakingPower();
    }

    function getSeasonInfo(uint projectID)
        external
        view
        returns (uint currentSeason, uint seasonStartBlock, uint seasonEndBlock, uint blocksElapsed)
    {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.getCurrentSeasonInfo();
    }

    function getProjectInfo(uint projectID)
        external
        view
        returns (address stakingPool, address rewardPool, string memory name, bool isActive, uint createdAt)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), "InvalidProjectID");
        (stakingPool, rewardPool, name, isActive, createdAt,,) = protocol.projects(projectID);
    }

    function getClaimableReward(uint projectID, address user, uint season, address rewardToken)
        external
        view
        returns (uint claimableAmount)
    {
        (, IRewardPool rewardPool) = _getPools(projectID);
        return rewardPool.getExpectedReward(user, season, rewardToken);
    }

    function getTotalStaked(uint projectID) external view returns (uint totalStaked) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.totalStaked();
    }

    function getSeasonTotalPoints(uint projectID, uint season) external view returns (uint totalPoints) {
        (IStakingPool pool,) = _getPools(projectID);
        uint currentSeason = pool.currentSeason();

        // 가상 시즌인 경우 직접 계산
        if (season > currentSeason) {
            (, totalPoints) = _calculateVirtualSeasonData(pool, season, address(0));
            return totalPoints;
        }

        // 온체인 시즌
        totalPoints = pool.seasonTotalPointsSnapshot(season);

        // totalPoints가 0이고 과거 시즌이면 계산 시도
        if (totalPoints == 0 && season < currentSeason) {
            (, totalPoints) = _calculateVirtualSeasonData(pool, season, address(0));
        }

        return totalPoints;
    }

    function getSeasonUserPoints(uint projectID, uint season, address user)
        external
        view
        returns (uint userPoints, uint totalPoints)
    {
        (IStakingPool pool,) = _getPools(projectID);
        uint currentSeason = pool.currentSeason();

        // 가상 시즌 (아직 온체인에 없는 미래 시즌)인 경우 직접 계산
        if (season > currentSeason) return _calculateVirtualSeasonData(pool, season, user);

        // 온체인 시즌은 pool에 위임
        (userPoints, totalPoints) = pool.getSeasonUserPoints(season, user);

        // totalPoints가 0인 경우에도 가상 계산 시도 (시즌은 있지만 finalize 안된 경우)
        if (totalPoints == 0 && season < currentSeason) {
            // 과거 시즌인데 totalPoints가 0이면 계산
            (, totalPoints) = _calculateVirtualSeasonData(pool, season, address(0));
        }
    }

    function isSeasonActive(uint projectID) external view returns (bool isActive) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.isSeasonActive();
    }

    function getPoolEndBlock(uint projectID) external view returns (uint endBlock) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.poolEndBlock();
    }

    function getNextSeasonStartBlock(uint projectID) external view returns (uint startBlock) {
        (IStakingPool pool,) = _getPools(projectID);
        return pool.nextSeasonStartBlock();
    }

    function getUserSeasonData(uint projectID, address user, uint season)
        external
        view
        returns (uint points, uint balance, uint joinBlock, bool claimed, bool finalized)
    {
        (IStakingPool pool,) = _getPools(projectID);
        // StakingPool의 getUserSeasonData 호출 (low-level staticcall과 동일 인터페이스)
        (bool success, bytes memory data) =
            address(pool).staticcall(abi.encodeWithSignature("getUserSeasonData(uint256,address)", season, user));
        if (success && data.length >= 160) {
            (points, balance, joinBlock, claimed, finalized) = abi.decode(data, (uint, uint, uint, bool, bool));
        }
    }

    function previewClaim(uint projectID, address user, uint season, address rewardToken)
        external
        view
        returns (uint expectedReward, uint userPoints, uint totalPoints, bool alreadyClaimed)
    {
        (IStakingPool pool,) = _getPools(projectID);
        (bool success, bytes memory data) = address(pool).staticcall(
            abi.encodeWithSignature("previewClaim(uint256,address,address)", season, user, rewardToken)
        );
        if (success && data.length >= 128) {
            (expectedReward, userPoints, totalPoints, alreadyClaimed) = abi.decode(data, (uint, uint, uint, bool));
        }
    }

    function getExpectedRewardsBatch(uint projectID, address user, uint[] calldata seasons, address rewardToken)
        external
        view
        returns (uint[] memory expectedRewards)
    {
        (IStakingPool pool,) = _getPools(projectID);
        expectedRewards = new uint[](seasons.length);
        for (uint i = 0; i < seasons.length;) {
            expectedRewards[i] = pool.getExpectedSeasonReward(seasons[i], user, rewardToken);
            unchecked {
                ++i;
            }
        }
    }

    function getActiveProjects(uint offset, uint limit)
        external
        view
        returns (uint[] memory projectIDs, address[] memory stakingPools, string[] memory names)
    {
        uint totalProjects = protocol.projectCount();
        require(offset < totalProjects, "InvalidOffset");

        uint end = offset + limit;
        if (end > totalProjects) end = totalProjects;
        uint count = end - offset;
        projectIDs = new uint[](count);
        stakingPools = new address[](count);
        names = new string[](count);

        uint resultIndex = 0;
        for (uint i = offset; i < end;) {
            uint projectID = i + 1;
            (address stakingPool,, string memory name, bool isActive,,,) = protocol.projects(projectID);
            if (isActive) {
                projectIDs[resultIndex] = projectID;
                stakingPools[resultIndex] = stakingPool;
                names[resultIndex] = name;
                unchecked {
                    ++resultIndex;
                }
            }
            unchecked {
                ++i;
            }
        }

        if (resultIndex < count) {
            assembly {
                mstore(projectIDs, resultIndex)
                mstore(stakingPools, resultIndex)
                mstore(names, resultIndex)
            }
        }
    }

    function getUserStakingSummary(address user, uint[] calldata projectIDs)
        external
        view
        returns (uint[] memory balances, uint[] memory points, uint[] memory currentSeasons)
    {
        balances = new uint[](projectIDs.length);
        points = new uint[](projectIDs.length);
        currentSeasons = new uint[](projectIDs.length);

        for (uint i = 0; i < projectIDs.length;) {
            if (projectIDs[i] > 0 && projectIDs[i] <= protocol.projectCount()) {
                (address stakingPool,,,,,,) = protocol.projects(projectIDs[i]);
                (balances[i], points[i],) = IStakingPool(stakingPool).getStakePosition(user);
                currentSeasons[i] = IStakingPool(stakingPool).currentSeason();
            }
            unchecked {
                ++i;
            }
        }
    }

    function getSeasonHistory(uint projectID, uint fromSeason, uint toSeason)
        external
        view
        returns (uint[] memory seasons, uint[] memory totalPoints)
    {
        require(projectID > 0 && projectID <= protocol.projectCount(), "InvalidProjectID");
        require(fromSeason <= toSeason, "InvalidRange");

        (IStakingPool pool,) = _getPools(projectID);
        uint currentSeason = pool.currentSeason();
        if (toSeason > currentSeason) toSeason = currentSeason;
        if (fromSeason > toSeason) return (new uint[](0), new uint[](0));

        uint count = toSeason - fromSeason + 1;
        seasons = new uint[](count);
        totalPoints = new uint[](count);
        for (uint i = 0; i < count;) {
            uint season = fromSeason + i;
            seasons[i] = season;
            totalPoints[i] = pool.seasonTotalPointsSnapshot(season);
            unchecked {
                ++i;
            }
        }
    }

    function getSeasonRewardTokens(uint projectID, uint season) external view returns (address[] memory tokens) {
        (, IRewardPool rewardPool) = _getPools(projectID);
        return rewardPool.getSeasonRewardTokens(season);
    }

    function getSeasonTokenInfo(uint projectID, uint season, address token)
        external
        view
        returns (uint total, uint claimed, uint remaining)
    {
        (, IRewardPool rewardPool) = _getPools(projectID);
        return rewardPool.getSeasonTokenInfo(season, token);
    }

    function getSeasonAllRewards(uint projectID, uint season)
        external
        view
        returns (address[] memory tokens, uint[] memory totals, uint[] memory claimeds, uint[] memory remainings)
    {
        (, IRewardPool rewardPool) = _getPools(projectID);
        return rewardPool.getSeasonAllRewards(season);
    }

    /**
     * @notice 시즌별 예상 포인트(Virtual). 풀에서 finalize되지 않은 시즌도 계산하여 반환
     */
    function getExpectedSeasonPoints(uint projectID, uint season, address user)
        external
        view
        returns (uint userPoints, uint totalPoints)
    {
        (IStakingPool pool,) = _getPools(projectID);
        uint currentSeason = pool.currentSeason();

        // 가상 시즌이면 직접 계산
        if (season > currentSeason) return _calculateVirtualSeasonData(pool, season, user);

        // 시즌 범위 계산
        (bool ok, uint startBlock, uint endBlock) = _calcSeasonRange(pool, season);
        if (!ok) return (0, 0);
        if (block.number < startBlock) return (0, 0);

        // 과거 시즌: 항상 endBlock 사용 (종료된 시즌이므로 고정)
        // 현재 시즌: 현재 블록까지만 반영
        uint toBlock = (season < currentSeason) ? endBlock : (block.number < endBlock ? block.number : endBlock);

        if (toBlock <= startBlock) return (0, 0);

        // 유저 포인트 계산
        (uint balance,, uint lastUpdate) = pool.getStakePosition(user);
        if (balance == 0) return (0, 0);
        if (lastUpdate > toBlock) return (0, 0);

        uint fromBlock = lastUpdate > startBlock ? lastUpdate : startBlock;
        userPoints = PointsLib.calculatePoints(balance, fromBlock, toBlock, pool.blockTime(), pool.pointsTimeUnit());

        // 토탈 포인트: 온체인 값 우선, 없으면 계산
        totalPoints = pool.seasonTotalPointsSnapshot(season);
        if (totalPoints == 0) {
            // 가상 계산
            (, totalPoints) = _calculateVirtualSeasonData(pool, season, address(0));
        }

        return (userPoints, totalPoints);
    }

    function getExpectedSeasonReward(uint projectID, uint season, address user, address rewardToken)
        external
        view
        returns (uint expectedReward)
    {
        (IStakingPool pool,) = _getPools(projectID);
        // StakingPool의 getExpectedSeasonReward를 호출
        return pool.getExpectedSeasonReward(season, user, rewardToken);
    }

    /**
     * @notice 풀의 기본 정보 조회 (predeposit 정보 포함)
     * @param projectID 프로젝트 ID
     * @return blockTime 블록 타임
     * @return pointsTimeUnit 포인트 계산 시간 단위
     * @return seasonBlocks 시즌 블록 수
     * @return poolEndBlock 풀 종료 블록
     * @return currentSeason 현재 시즌
     * @return preDepositStartBlock Pre-deposit 시작 블록
     * @return firstSeasonStartBlock 첫 시즌 시작 블록
     */
    function getPoolInfo(uint projectID)
        external
        view
        returns (
            uint blockTime,
            uint pointsTimeUnit,
            uint seasonBlocks,
            uint poolEndBlock,
            uint currentSeason,
            uint preDepositStartBlock,
            uint firstSeasonStartBlock
        )
    {
        (IStakingPool pool,) = _getPools(projectID);

        blockTime = pool.blockTime();
        pointsTimeUnit = pool.pointsTimeUnit();
        seasonBlocks = pool.seasonBlocks();
        poolEndBlock = pool.poolEndBlock();
        currentSeason = pool.currentSeason();
        preDepositStartBlock = pool.preDepositStartBlock();

        // firstSeasonStartBlock 계산
        if (currentSeason > 0) {
            // 시즌이 시작된 경우: 현재 시즌 정보에서 역산
            (uint season, uint startBlock,,) = pool.getCurrentSeasonInfo();
            if (season > 0 && startBlock > 0) firstSeasonStartBlock = startBlock - ((season - 1) * seasonBlocks);
        } else {
            // 시즌 시작 전: nextSeasonStartBlock 사용
            firstSeasonStartBlock = pool.nextSeasonStartBlock();
        }

        return (
            blockTime,
            pointsTimeUnit,
            seasonBlocks,
            poolEndBlock,
            currentSeason,
            preDepositStartBlock,
            firstSeasonStartBlock
        );
    }
}
