// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/CrossStaking.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title MockERC20
 * @notice 테스트용 ERC20 토큰
 */
contract MockERC20 is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _mint(msg.sender, 1000000 * 10 ** 18);
    }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

/**
 * @title CrossStakingTest
 * @notice CrossStaking 컨트랙트 테스트
 */
contract CrossStakingTest is Test {
    CrossStaking public staking;
    MockERC20 public stakingToken;
    MockERC20 public rewardToken1;
    MockERC20 public rewardToken2;

    address public owner;
    address public user1;
    address public user2;
    address public user3;

    uint256 constant INITIAL_BALANCE = 10000 * 10 ** 18;
    uint256 constant STAKE_AMOUNT = 1000 * 10 ** 18;

    event Staked(address indexed user, uint256 amount, uint256 expiry, uint256 stakingPower);
    event Unstaked(address indexed user, uint256 amount);
    event AmountIncreased(address indexed user, uint256 additionalAmount, uint256 newStakingPower);
    event DurationExtended(address indexed user, uint256 newExpiry, uint256 newStakingPower);
    event RewardsClaimed(address indexed user, address indexed token, uint256 amount);

    function setUp() public {
        owner = address(this);
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        user3 = makeAddr("user3");

        // 토큰 배포
        stakingToken = new MockERC20("Staking Token", "STK");
        rewardToken1 = new MockERC20("Reward Token 1", "RWD1");
        rewardToken2 = new MockERC20("Reward Token 2", "RWD2");

        // CrossStaking 컨트랙트 배포
        staking = new CrossStaking(address(stakingToken), owner);

        // 사용자들에게 토큰 전송
        stakingToken.transfer(user1, INITIAL_BALANCE);
        stakingToken.transfer(user2, INITIAL_BALANCE);
        stakingToken.transfer(user3, INITIAL_BALANCE);

        // 사용자들이 스테이킹 컨트랙트에 approve
        vm.prank(user1);
        stakingToken.approve(address(staking), type(uint256).max);

        vm.prank(user2);
        stakingToken.approve(address(staking), type(uint256).max);

        vm.prank(user3);
        stakingToken.approve(address(staking), type(uint256).max);

        // 보상 토큰 추가
        staking.addRewardToken(address(rewardToken1));
        staking.addRewardToken(address(rewardToken2));
    }

    // ============ 기본 스테이킹 테스트 ============

    function testStakeSuccess() public {
        uint256 duration = 2 weeks;

        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        (uint256 amount, uint256 expiry, uint256 stakingPower) = staking.getStakePosition(user1);

        assertEq(amount, STAKE_AMOUNT);
        assertEq(expiry, block.timestamp + duration);
        assertGt(stakingPower, 0);

        // 토큰이 전송되었는지 확인
        assertEq(stakingToken.balanceOf(user1), INITIAL_BALANCE - STAKE_AMOUNT);
        assertEq(stakingToken.balanceOf(address(staking)), STAKE_AMOUNT);
    }

    function testStakeMinDuration() public {
        uint256 duration = staking.MIN_LOCK_TIME();

        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        (uint256 amount,,) = staking.getStakePosition(user1);
        assertEq(amount, STAKE_AMOUNT);
    }

    function testStakeMaxDuration() public {
        uint256 duration = staking.MAX_LOCK_TIME();

        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        (uint256 amount,,) = staking.getStakePosition(user1);
        assertEq(amount, STAKE_AMOUNT);
    }

    function testStakeFailTooShortDuration() public {
        uint256 duration = staking.MIN_LOCK_TIME() - 1;

        vm.prank(user1);
        vm.expectRevert("Duration too short");
        staking.stake(STAKE_AMOUNT, duration);
    }

    function testStakeFailTooLongDuration() public {
        uint256 duration = staking.MAX_LOCK_TIME() + 1;

        vm.prank(user1);
        vm.expectRevert("Duration too long");
        staking.stake(STAKE_AMOUNT, duration);
    }

    function testStakeFailZeroAmount() public {
        vm.prank(user1);
        vm.expectRevert("Amount must be greater than 0");
        staking.stake(0, 2 weeks);
    }

    function testStakeFailPositionAlreadyExists() public {
        vm.startPrank(user1);
        staking.stake(STAKE_AMOUNT, 2 weeks);

        vm.expectRevert("Position already exists");
        staking.stake(STAKE_AMOUNT, 2 weeks);
        vm.stopPrank();
    }

    // ============ 스테이킹 파워 계산 테스트 ============

    function testStakingPowerCalculation() public {
        uint256 duration = staking.MAX_LOCK_TIME();

        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        // 최대 기간으로 스테이킹하면 스테이킹 파워 = 스테이킹 수량
        uint256 stakingPower = staking.getStakingPower(user1);
        assertApproxEqRel(stakingPower, STAKE_AMOUNT, 0.01e18); // 1% 오차 허용
    }

    function testStakingPowerDecreaseOverTime() public {
        uint256 duration = 4 weeks;

        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        uint256 initialPower = staking.getStakingPower(user1);

        // 2주 경과
        vm.warp(block.timestamp + 2 weeks);

        uint256 powerAfter2Weeks = staking.getStakingPower(user1);

        // 스테이킹 파워가 감소해야 함
        assertLt(powerAfter2Weeks, initialPower);
    }

    function testStakingPowerZeroAfterExpiry() public {
        uint256 duration = 2 weeks;

        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        // 만료 후
        vm.warp(block.timestamp + duration + 1);

        uint256 stakingPower = staking.getStakingPower(user1);
        assertEq(stakingPower, 0);
    }

    // ============ 수량 증가 테스트 ============

    function testIncreaseAmount() public {
        vm.startPrank(user1);
        staking.stake(STAKE_AMOUNT, 2 weeks);

        uint256 additionalAmount = 500 * 10 ** 18;
        uint256 oldPower = staking.getStakingPower(user1);

        staking.increaseAmount(additionalAmount);

        (uint256 amount,,) = staking.getStakePosition(user1);
        uint256 newPower = staking.getStakingPower(user1);

        assertEq(amount, STAKE_AMOUNT + additionalAmount);
        assertGt(newPower, oldPower);
        vm.stopPrank();
    }

    function testIncreaseAmountFailNoPosition() public {
        vm.prank(user1);
        vm.expectRevert("No staking position");
        staking.increaseAmount(100 * 10 ** 18);
    }

    function testIncreaseAmountFailExpired() public {
        uint256 duration = 2 weeks;

        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        // 만료 후
        vm.warp(block.timestamp + duration + 1);

        vm.prank(user1);
        vm.expectRevert("Position expired");
        staking.increaseAmount(100 * 10 ** 18);
    }

    // ============ 기간 연장 테스트 ============

    function testExtendDuration() public {
        uint256 duration = 2 weeks;

        vm.startPrank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        uint256 oldPower = staking.getStakingPower(user1);

        uint256 newExpiry = block.timestamp + 4 weeks;
        staking.extendDuration(newExpiry);

        (, uint256 expiry,) = staking.getStakePosition(user1);
        uint256 newPower = staking.getStakingPower(user1);

        assertEq(expiry, newExpiry);
        assertGt(newPower, oldPower);
        vm.stopPrank();
    }

    function testExtendDurationFailNotGreater() public {
        uint256 duration = 2 weeks;

        vm.startPrank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        uint256 currentExpiry = block.timestamp + duration;

        vm.expectRevert("New expiry must be greater");
        staking.extendDuration(currentExpiry);
        vm.stopPrank();
    }

    function testExtendDurationFailExceedsMax() public {
        uint256 duration = 2 weeks;

        vm.startPrank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        uint256 tooLongExpiry = block.timestamp + staking.MAX_LOCK_TIME() + 1;

        vm.expectRevert("Exceeds max lock time");
        staking.extendDuration(tooLongExpiry);
        vm.stopPrank();
    }

    // ============ 언스테이킹 테스트 ============

    function testUnstakeSuccess() public {
        uint256 duration = 2 weeks;

        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        // 만료 후
        vm.warp(block.timestamp + duration + 1);

        uint256 balanceBefore = stakingToken.balanceOf(user1);

        vm.prank(user1);
        staking.unstake();

        uint256 balanceAfter = stakingToken.balanceOf(user1);

        assertEq(balanceAfter - balanceBefore, STAKE_AMOUNT);

        // 포지션이 삭제되었는지 확인
        (uint256 amount,,) = staking.getStakePosition(user1);
        assertEq(amount, 0);
    }

    function testUnstakeFailNotExpired() public {
        uint256 duration = 2 weeks;

        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        vm.prank(user1);
        vm.expectRevert("Position not expired");
        staking.unstake();
    }

    function testUnstakeFailNoPosition() public {
        vm.prank(user1);
        vm.expectRevert("No staking position");
        staking.unstake();
    }

    // ============ 보상 분배 테스트 ============

    function testDistributeRewards() public {
        // user1과 user2가 스테이킹
        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, 3 weeks);

        vm.prank(user2);
        staking.stake(STAKE_AMOUNT * 2, 3 weeks);

        // 보상 분배
        uint256 rewardAmount = 1000 * 10 ** 18;
        rewardToken1.approve(address(staking), rewardAmount);
        staking.distributeRewards(address(rewardToken1), rewardAmount);

        // 보상이 올바르게 계산되는지 확인
        uint256 user1Pending = staking.getPendingReward(user1, address(rewardToken1));
        uint256 user2Pending = staking.getPendingReward(user2, address(rewardToken1));

        assertGt(user1Pending, 0);
        assertGt(user2Pending, 0);

        // user2의 보상이 user1보다 많아야 함 (2배 스테이킹)
        assertGt(user2Pending, user1Pending);
    }

    function testClaimReward() public {
        // 스테이킹
        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, 3 weeks);

        // 보상 분배
        uint256 rewardAmount = 1000 * 10 ** 18;
        rewardToken1.approve(address(staking), rewardAmount);
        staking.distributeRewards(address(rewardToken1), rewardAmount);

        // 보상 청구
        uint256 pendingBefore = staking.getPendingReward(user1, address(rewardToken1));
        uint256 balanceBefore = rewardToken1.balanceOf(user1);

        vm.prank(user1);
        staking.claimReward(address(rewardToken1));

        uint256 balanceAfter = rewardToken1.balanceOf(user1);
        uint256 pendingAfter = staking.getPendingReward(user1, address(rewardToken1));

        assertEq(balanceAfter - balanceBefore, pendingBefore);
        assertEq(pendingAfter, 0);
    }

    function testClaimAllRewards() public {
        // 스테이킹
        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, 3 weeks);

        // 두 가지 보상 토큰 분배
        uint256 rewardAmount1 = 1000 * 10 ** 18;
        rewardToken1.approve(address(staking), rewardAmount1);
        staking.distributeRewards(address(rewardToken1), rewardAmount1);

        uint256 rewardAmount2 = 500 * 10 ** 18;
        rewardToken2.approve(address(staking), rewardAmount2);
        staking.distributeRewards(address(rewardToken2), rewardAmount2);

        // 모든 보상 청구
        uint256 balance1Before = rewardToken1.balanceOf(user1);
        uint256 balance2Before = rewardToken2.balanceOf(user1);

        vm.prank(user1);
        staking.claimAllRewards();

        uint256 balance1After = rewardToken1.balanceOf(user1);
        uint256 balance2After = rewardToken2.balanceOf(user1);

        assertGt(balance1After - balance1Before, 0);
        assertGt(balance2After - balance2Before, 0);
    }

    function testMultipleUsersRewardDistribution() public {
        // 3명의 사용자가 다른 수량으로 스테이킹
        vm.prank(user1);
        staking.stake(1000 * 10 ** 18, 3 weeks);

        vm.prank(user2);
        staking.stake(2000 * 10 ** 18, 3 weeks);

        vm.prank(user3);
        staking.stake(3000 * 10 ** 18, 3 weeks);

        // 보상 분배
        uint256 rewardAmount = 6000 * 10 ** 18;
        rewardToken1.approve(address(staking), rewardAmount);
        staking.distributeRewards(address(rewardToken1), rewardAmount);

        uint256 user1Reward = staking.getPendingReward(user1, address(rewardToken1));
        uint256 user2Reward = staking.getPendingReward(user2, address(rewardToken1));
        uint256 user3Reward = staking.getPendingReward(user3, address(rewardToken1));

        // 보상 비율이 스테이킹 비율과 유사해야 함
        // user1:user2:user3 = 1:2:3
        assertApproxEqRel(user2Reward, user1Reward * 2, 0.01e18);
        assertApproxEqRel(user3Reward, user1Reward * 3, 0.01e18);
    }

    // ============ 관리자 함수 테스트 ============

    function testAddRewardToken() public {
        MockERC20 newToken = new MockERC20("New Reward", "NEW");

        staking.addRewardToken(address(newToken));

        address[] memory tokens = staking.getRewardTokenList();
        bool found = false;
        for (uint256 i = 0; i < tokens.length; i++) {
            if (tokens[i] == address(newToken)) {
                found = true;
                break;
            }
        }
        assertTrue(found);
    }

    function testAddRewardTokenFailNotOwner() public {
        MockERC20 newToken = new MockERC20("New Reward", "NEW");

        vm.prank(user1);
        vm.expectRevert();
        staking.addRewardToken(address(newToken));
    }

    function testRemoveRewardToken() public {
        staking.removeRewardToken(address(rewardToken1));

        (,, bool isActive) = staking.rewardTokens(address(rewardToken1));
        assertFalse(isActive);
    }

    // ============ 레거시 함수 테스트 ============

    function testGetStakedAmount() public {
        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, 2 weeks);

        uint256 stakedAmount = staking.getStakedAmount(user1);
        assertEq(stakedAmount, STAKE_AMOUNT);
    }

    function testGetUnstakedAmount() public {
        uint256 duration = 2 weeks;

        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, duration);

        // 만료 전
        uint256 unstakedBefore = staking.getUnstakedAmount(user1);
        assertEq(unstakedBefore, 0);

        // 만료 후
        vm.warp(block.timestamp + duration + 1);
        uint256 unstakedAfter = staking.getUnstakedAmount(user1);
        assertEq(unstakedAfter, STAKE_AMOUNT);
    }

    // ============ 통합 시나리오 테스트 ============

    function testCompleteStakingCycle() public {
        // 1. 스테이킹
        vm.prank(user1);
        staking.stake(STAKE_AMOUNT, 3 weeks);

        // 2. 수량 증가
        vm.prank(user1);
        staking.increaseAmount(500 * 10 ** 18);

        // 3. 기간 연장
        vm.prank(user1);
        staking.extendDuration(block.timestamp + 4 weeks);

        // 4. 보상 분배
        uint256 rewardAmount = 1000 * 10 ** 18;
        rewardToken1.approve(address(staking), rewardAmount);
        staking.distributeRewards(address(rewardToken1), rewardAmount);

        // 5. 보상 청구
        vm.prank(user1);
        staking.claimReward(address(rewardToken1));

        // 6. 시간 경과 후 언스테이킹
        vm.warp(block.timestamp + 4 weeks + 1);

        vm.prank(user1);
        staking.unstake();

        // 스테이킹 토큰이 모두 반환되었는지 확인
        // user1은 1000 + 500 = 1500 토큰을 스테이킹했고, 모두 반환받아야 함
        assertEq(stakingToken.balanceOf(user1), INITIAL_BALANCE);

        // 보상 토큰을 받았는지 확인
        assertGt(rewardToken1.balanceOf(user1), 0);
    }

    function testMultipleStakersWithDifferentDurations() public {
        // user1: 짧은 기간, 많은 수량
        vm.prank(user1);
        staking.stake(3000 * 10 ** 18, 1 weeks);

        // user2: 긴 기간, 적은 수량
        vm.prank(user2);
        staking.stake(1000 * 10 ** 18, 5 weeks);

        uint256 power1 = staking.getStakingPower(user1);
        uint256 power2 = staking.getStakingPower(user2);

        // 스테이킹 파워가 수량과 기간 모두에 영향을 받는지 확인
        assertGt(power1, 0);
        assertGt(power2, 0);

        // 보상 분배
        uint256 rewardAmount = 1000 * 10 ** 18;
        rewardToken1.approve(address(staking), rewardAmount);
        staking.distributeRewards(address(rewardToken1), rewardAmount);

        uint256 reward1 = staking.getPendingReward(user1, address(rewardToken1));
        uint256 reward2 = staking.getPendingReward(user2, address(rewardToken1));

        // 두 사용자 모두 보상을 받아야 함
        assertGt(reward1, 0);
        assertGt(reward2, 0);
    }
}
