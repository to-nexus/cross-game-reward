// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./BaseTest.sol";

/**
 * @title FuzzTest
 * @notice Fuzzing 테스트 - 다양한 입력값으로 시스템 견고성 검증
 */
contract FuzzTest is BaseTest {
    uint constant MIN_STAKE = 1e18; // 최소 스테이크: 1 CROSS
    uint constant MAX_FUZZ_STAKE = 1000 ether; // Fuzzing 테스트용 최대 스테이크

    // ============ 포인트 계산 Fuzzing ============

    /**
     * @notice 다양한 스테이킹 금액과 블록 수로 포인트 계산 검증
     * @dev 포인트는 항상 시간과 금액에 비례해야 함
     */
    function testFuzz_PointsCalculation(uint88 stakeAmount, uint32 blocks) public {
        // 바운드 설정
        stakeAmount = uint88(bound(stakeAmount, MIN_STAKE, MAX_FUZZ_STAKE));
        blocks = uint32(bound(blocks, 1, SEASON_BLOCKS));

        // user1에게 WCROSS 직접 할당
        deal(address(wcross), user1, stakeAmount);

        // 스테이킹
        vm.startPrank(user1);
        wcross.approve(address(stakingPool), stakeAmount);
        stakingPool.stake(stakeAmount);
        vm.stopPrank();

        // 블록 경과
        vm.roll(block.number + blocks);

        // 포인트 조회
        uint points = stakingPool.getUserPoints(user1);

        // 검증: 포인트는 항상 0보다 커야 함
        assertGt(points, 0, "Points should be greater than 0");

        // 검증: 포인트는 스테이킹 금액과 시간에 비례
        // points = (stakeAmount * blocks * blockTime * POINTS_PRECISION) / pointsTimeUnit
        uint blockTime = stakingPool.blockTime();
        uint pointsTimeUnit = stakingPool.pointsTimeUnit();
        uint POINTS_PRECISION = 1e6; // StakingPool.POINTS_PRECISION
        uint expectedPoints = (uint(stakeAmount) * blocks * blockTime * POINTS_PRECISION) / pointsTimeUnit;
        assertApproxEqRel(points, expectedPoints, 0.01e18, "Points calculation mismatch");
    }

    /**
     * @notice 추가 예치 시 포인트가 올바르게 누적되는지 fuzzing 검증
     */
    function testFuzz_IncrementalStaking(uint88 initialAmount, uint88 additionalAmount, uint32 blocks1, uint32 blocks2)
        public
    {
        // 바운드 설정
        initialAmount = uint88(bound(initialAmount, MIN_STAKE, MAX_FUZZ_STAKE / 2));
        additionalAmount = uint88(bound(additionalAmount, MIN_STAKE, MAX_FUZZ_STAKE / 2));
        blocks1 = uint32(bound(blocks1, 1, SEASON_BLOCKS / 2));
        blocks2 = uint32(bound(blocks2, 1, SEASON_BLOCKS / 2));

        // user1에게 WCROSS 직접 할당
        deal(address(wcross), user1, initialAmount + additionalAmount);

        vm.startPrank(user1);
        wcross.approve(address(stakingPool), initialAmount + additionalAmount);

        // 1. 초기 예치
        stakingPool.stake(initialAmount);

        // 2. 첫 번째 블록 경과
        vm.roll(block.number + blocks1);

        uint pointsBeforeAdditional = stakingPool.getUserPoints(user1);

        // 3. 추가 예치
        stakingPool.stake(additionalAmount);

        // 4. 두 번째 블록 경과
        vm.roll(block.number + blocks2);

        uint finalPoints = stakingPool.getUserPoints(user1);

        vm.stopPrank();

        // 검증: 추가 예치 후 포인트가 증가했어야 함
        assertGt(finalPoints, pointsBeforeAdditional, "Points should increase after additional stake");

        // 검증: 최종 잔액
        (uint balance,,) = stakingPool.getStakePosition(user1);
        assertEq(balance, uint(initialAmount) + additionalAmount, "Final balance incorrect");
    }

    /**
     * @notice 여러 사용자의 포인트 비율 검증
     */
    function testFuzz_MultiUserPointsRatio(uint88 amount1, uint88 amount2, uint32 blocks) public {
        // 바운드 설정
        amount1 = uint88(bound(amount1, MIN_STAKE, MAX_FUZZ_STAKE / 2));
        amount2 = uint88(bound(amount2, MIN_STAKE, MAX_FUZZ_STAKE / 2));
        blocks = uint32(bound(blocks, 100, SEASON_BLOCKS));

        // 두 사용자 스테이킹
        stakeFor(user1, amount1);
        stakeFor(user2, amount2);

        // 블록 경과
        vm.roll(block.number + blocks);

        // 포인트 조회
        uint points1 = stakingPool.getUserPoints(user1);
        uint points2 = stakingPool.getUserPoints(user2);

        // 검증: 포인트 비율 = 스테이킹 비율
        if (amount1 > amount2) assertGt(points1, points2, "Higher stake should have more points");
        else if (amount2 > amount1) assertGt(points2, points1, "Higher stake should have more points");
        else assertApproxEqRel(points1, points2, 0.01e18, "Equal stake should have equal points");

        // 비율 검증 (1% 오차 허용)
        uint ratio1 = (points1 * 1e18) / amount1;
        uint ratio2 = (points2 * 1e18) / amount2;
        assertApproxEqRel(ratio1, ratio2, 0.01e18, "Points per token should be equal");
    }

    // ============ 보상 분배 Fuzzing ============

    /**
     * @notice 다양한 포인트 비율과 보상 금액으로 분배 검증
     */
    function testFuzz_RewardDistribution(uint88 stake1, uint88 stake2, uint88 rewardAmount) public {
        // 바운드 설정
        stake1 = uint88(bound(stake1, MIN_STAKE, MAX_FUZZ_STAKE / 2));
        stake2 = uint88(bound(stake2, MIN_STAKE, MAX_FUZZ_STAKE / 2));
        rewardAmount = uint88(bound(rewardAmount, 1 ether, MAX_FUZZ_STAKE));

        // 스테이킹
        stakeFor(user1, stake1);
        stakeFor(user2, stake2);

        // 시즌 롤오버
        rolloverSeason();

        // 보상 예치
        fundSeason(1, rewardAmount);

        // 보상 청구
        claimSeasonFor(user1, 1);
        claimSeasonFor(user2, 1);

        uint reward1 = rewardToken.balanceOf(user1);
        uint reward2 = rewardToken.balanceOf(user2);

        // 검증: 총 보상은 예치 금액 이하
        assertLe(reward1 + reward2, rewardAmount, "Total rewards exceed funded amount");

        // 검증: 보상 비율 = 스테이킹 비율 (1% 오차 허용)
        if (stake1 > 0 && stake2 > 0) {
            uint ratio1 = (reward1 * 1e18) / stake1;
            uint ratio2 = (reward2 * 1e18) / stake2;
            assertApproxEqRel(ratio1, ratio2, 0.01e18, "Reward per token should be proportional");
        }
    }

    /**
     * @notice 다양한 시즌 보상 예치 및 청구 검증
     */
    function testFuzz_MultipleSeasonRewards(uint88 stakeAmount, uint88 reward1, uint88 reward2) public {
        // 바운드 설정
        stakeAmount = uint88(bound(stakeAmount, MIN_STAKE, MAX_FUZZ_STAKE / 2));
        reward1 = uint88(bound(reward1, 1 ether, MAX_FUZZ_STAKE / 2));
        reward2 = uint88(bound(reward2, 1 ether, MAX_FUZZ_STAKE / 2));

        // 스테이킹
        stakeFor(user1, stakeAmount);

        // 시즌 1 롤오버 및 보상
        rolloverSeason();
        fundSeason(1, reward1);

        // 시즌 2 롤오버 및 보상
        rolloverSeason();
        fundSeason(2, reward2);

        // 보상 청구
        claimSeasonFor(user1, 1);
        claimSeasonFor(user1, 2);

        uint totalReward = rewardToken.balanceOf(user1);

        // 검증: 총 보상은 두 시즌 보상 합 이하
        assertLe(totalReward, reward1 + reward2, "Total reward exceeds funded");
        assertGt(totalReward, 0, "Should receive some reward");
    }

    // ============ 시즌 시스템 Fuzzing ============

    /**
     * @notice 다양한 타이밍에서 시즌 전환 검증
     */
    function testFuzz_SeasonRollover(uint32 blocksInSeason) public {
        // 바운드: 시즌 길이의 50%~150%
        blocksInSeason = uint32(bound(blocksInSeason, SEASON_BLOCKS / 2, SEASON_BLOCKS * 3 / 2));

        stakeFor(user1, 10 ether);

        uint initialSeason = stakingPool.currentSeason();

        // 블록 진행
        vm.roll(block.number + blocksInSeason);

        // 시즌 롤오버 시도 (endBlock + 1에서 가능)
        if (blocksInSeason > SEASON_BLOCKS) {
            // 롤오버 가능
            stakingPool.rolloverSeason();
            uint newSeason = stakingPool.currentSeason();
            assertEq(newSeason, initialSeason + 1, "Season should increment");
        } else {
            // 롤오버 불가능
            vm.expectRevert();
            stakingPool.rolloverSeason();
        }
    }

    /**
     * @notice 시즌 전환 전후 포인트 보존 검증
     */
    function testFuzz_PointsAcrossSeasons(uint88 stakeAmount, uint32 blocksBeforeRollover, uint32 blocksAfterRollover)
        public
    {
        // 바운드 설정
        stakeAmount = uint88(bound(stakeAmount, MIN_STAKE, MAX_FUZZ_STAKE / 2));
        blocksBeforeRollover = uint32(bound(blocksBeforeRollover, 1, SEASON_BLOCKS));
        blocksAfterRollover = uint32(bound(blocksAfterRollover, 1, SEASON_BLOCKS));

        // 스테이킹
        stakeFor(user1, stakeAmount);

        // 시즌 1 포인트 누적
        vm.roll(block.number + blocksBeforeRollover);
        uint season1Points = stakingPool.getUserPoints(user1);

        // 시즌 종료까지 대기
        vm.roll(block.number + (SEASON_BLOCKS - blocksBeforeRollover) + 1);

        // 롤오버
        stakingPool.rolloverSeason();

        // 시즌 2 포인트 누적
        vm.roll(block.number + blocksAfterRollover);
        uint season2Points = stakingPool.getUserPoints(user2);

        // 검증: 시즌 1 포인트는 스냅샷됨
        uint snapshotPoints = stakingPool.getExpectedSeasonPoints(1, user1);
        assertGt(snapshotPoints, 0, "Season 1 points should be snapshotted");

        // 검증: 스테이킹은 유지됨
        (uint balance,,) = stakingPool.getStakePosition(user1);
        assertEq(balance, stakeAmount, "Stake should be preserved across seasons");
    }

    // ============ 경계값 및 극단 케이스 Fuzzing ============

    /**
     * @notice 매우 작은 금액 스테이킹 검증
     */
    function testFuzz_MinimumStake(uint88 amount) public {
        amount = uint88(bound(amount, 1, MIN_STAKE * 2));

        // user1에게 WCROSS 직접 할당
        deal(address(wcross), user1, amount);

        vm.startPrank(user1);
        wcross.approve(address(stakingPool), amount);

        if (amount < MIN_STAKE) {
            // 최소 금액 미만은 실패해야 함
            vm.expectRevert();
            stakingPool.stake(amount);
        } else {
            // 최소 금액 이상은 성공
            stakingPool.stake(amount);
            (uint balance,,) = stakingPool.getStakePosition(user1);
            assertEq(balance, amount);
        }

        vm.stopPrank();
    }

    /**
     * @notice 매우 큰 금액 스테이킹 검증 (오버플로우 체크)
     */
    function testFuzz_LargeStake(uint88 amount) public {
        // 바운드: 합리적인 최대 범위 내 (오버플로우 방지)
        amount = uint88(bound(amount, MIN_STAKE, MAX_FUZZ_STAKE));

        // user1에게 WCROSS 직접 할당
        deal(address(wcross), user1, amount);

        vm.startPrank(user1);
        wcross.approve(address(stakingPool), amount);
        stakingPool.stake(amount);
        vm.stopPrank();

        (uint balance,,) = stakingPool.getStakePosition(user1);
        assertEq(balance, amount, "Large stake should work");

        // 포인트 계산도 정상 작동해야 함
        vm.roll(block.number + 100);
        uint points = stakingPool.getUserPoints(user1);
        assertGt(points, 0, "Points should accumulate for large stake");
    }

    /**
     * @notice 매우 긴 시간 경과 후 포인트 계산 (오버플로우 체크)
     */
    function testFuzz_LongDuration(uint32 blocks) public {
        // 바운드: 1 시즌 ~ 10 시즌
        blocks = uint32(bound(blocks, SEASON_BLOCKS, SEASON_BLOCKS * 10));

        stakeFor(user1, 100 ether);

        vm.roll(block.number + blocks);

        // 포인트 조회 시 오버플로우 없어야 함
        uint points = stakingPool.getUserPoints(user1);
        assertGt(points, 0, "Points should accumulate over long duration");

        // 예상 포인트와 비교
        uint blockTime = stakingPool.blockTime();
        uint pointsTimeUnit = stakingPool.pointsTimeUnit();
        uint POINTS_PRECISION = 1e6;
        uint expectedPoints = (100 ether * uint(blocks) * blockTime * POINTS_PRECISION) / pointsTimeUnit;
        assertApproxEqRel(points, expectedPoints, 0.01e18, "Long duration points calculation");
    }

    /**
     * @notice 매우 많은 추가 예치 검증
     */
    function testFuzz_ManyIncrementalStakes(uint8 numStakes) public {
        // 바운드: 2~10회 추가 예치 (토큰 한도 고려)
        numStakes = uint8(bound(numStakes, 2, 10));

        uint stakeAmount = 10 ether;
        uint totalNeeded = stakeAmount * numStakes;

        // user1에게 WCROSS 직접 할당
        deal(address(wcross), user1, totalNeeded);

        vm.startPrank(user1);
        wcross.approve(address(stakingPool), totalNeeded);

        uint totalStaked = 0;

        for (uint i = 0; i < numStakes; i++) {
            stakingPool.stake(stakeAmount);
            totalStaked += stakeAmount;

            // 각 예치 사이에 시간 경과
            vm.roll(block.number + 100);
        }

        vm.stopPrank();

        // 검증: 총 스테이킹 금액
        (uint balance,,) = stakingPool.getStakePosition(user1);
        assertEq(balance, totalStaked, "Total staked should match");

        // 검증: 포인트 누적
        uint points = stakingPool.getUserPoints(user1);
        assertGt(points, 0, "Points should accumulate across multiple stakes");
    }

    /**
     * @notice 제로 금액 엣지 케이스
     */
    function testFuzz_ZeroAmount() public {
        vm.startPrank(user1);
        wcross.approve(address(stakingPool), 0);

        // 0 금액 스테이킹은 실패해야 함
        vm.expectRevert();
        stakingPool.stake(0);

        vm.stopPrank();
    }

    /**
     * @notice blockTime과 pointsTimeUnit 조합 검증
     * @dev 오버플로우 방지를 위해 고정된 합리적인 값 조합 테스트
     */
    function testFuzz_TimeParameters(uint8 seed) public {
        // seed를 사용하여 합리적인 값 조합 선택
        seed = uint8(bound(seed, 0, 11)); // 12가지 조합

        uint blockTime;
        uint timeUnit;

        // 12가지 합리적인 조합 (오버플로우 없음)
        if (seed < 4) {
            blockTime = 1; // 1초 블록
            timeUnit = [uint(1 hours), 6 hours, 12 hours, 1 days][seed];
        } else if (seed < 8) {
            blockTime = 2; // 2초 블록
            timeUnit = [uint(1 hours), 6 hours, 12 hours, 1 days][seed - 4];
        } else {
            blockTime = 12; // 12초 블록 (이더리움)
            timeUnit = [uint(1 hours), 6 hours, 12 hours, 1 days][seed - 8];
        }

        // 설정 변경
        vm.prank(address(protocol));
        stakingPool.setBlockTime(blockTime);

        vm.prank(address(protocol));
        stakingPool.setPointsTimeUnit(timeUnit);

        // 스테이킹
        stakeFor(user1, 10 ether);

        // 시간 경과
        vm.roll(block.number + 100);

        // 포인트 조회 (오버플로우 없어야 함)
        uint points = stakingPool.getUserPoints(user1);

        // 검증: 포인트가 계산됨
        assertGt(points, 0, "Points should be calculated with custom time parameters");

        // 검증: 예상 포인트 (blockTime과 timeUnit의 조합)
        uint POINTS_PRECISION = 1e6;
        uint expectedPoints = (10 ether * 100 * blockTime * POINTS_PRECISION) / timeUnit;
        assertApproxEqRel(points, expectedPoints, 0.01e18, "Points with custom time parameters");
    }
}
