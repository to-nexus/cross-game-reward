// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./BaseTest.sol";

/**
 * @title StakingTest
 * @notice Stake/Unstake 기능 테스트
 */
contract StakingTest is BaseTest {
    function test_BasicStake() public {
        vm.startPrank(user1);

        uint balanceBefore = user1.balance;
        router.stake{value: 10 ether}(PROJECT_ID);

        assertEq(user1.balance, balanceBefore - 10 ether);
        (uint balance,,) = stakingPool.getStakePosition(user1);
        assertEq(balance, 10 ether);

        vm.stopPrank();
    }

    function test_MultipleStakes() public {
        stakeFor(user1, 5 ether);
        (uint balance1,,) = stakingPool.getStakePosition(user1);
        assertEq(balance1, 5 ether);

        stakeFor(user1, 3 ether);
        (uint balance2,,) = stakingPool.getStakePosition(user1);
        assertEq(balance2, 8 ether);
    }

    function test_MinimumStake() public {
        // 최소 스테이크 정확히
        stakeFor(user1, 1 ether);
        (uint balance,,) = stakingPool.getStakePosition(user1);
        assertEq(balance, 1 ether);
    }

    function test_RevertWhen_BelowMinimumStake() public {
        vm.startPrank(user1);
        vm.expectRevert();
        router.stake{value: 0.5 ether}(PROJECT_ID);
        vm.stopPrank();
    }

    function test_WithdrawAll() public {
        stakeFor(user1, 10 ether);
        withdrawFor(user1);

        (uint balance,,) = stakingPool.getStakePosition(user1);
        assertEq(balance, 0);
        assertEq(wcross.balanceOf(user1), 10 ether);
    }

    function test_RevertWhen_WithdrawWithoutStake() public {
        vm.prank(user1);
        vm.expectRevert();
        stakingPool.withdrawAll();
    }

    function test_RouterUnstake() public {
        vm.startPrank(user1);

        uint balanceBefore = user1.balance;
        router.stake{value: 10 ether}(PROJECT_ID);

        // router.unstake()가 내부에서 withdrawAllFor를 호출하므로 별도 withdrawAll 불필요
        router.unstake(PROJECT_ID);

        assertGt(user1.balance, balanceBefore - 10.01 ether);

        vm.stopPrank();
    }

    function test_MultipleUsersStake() public {
        stakeFor(user1, 10 ether);
        stakeFor(user2, 5 ether);
        stakeFor(user3, 3 ether);

        (uint balance1,,) = stakingPool.getStakePosition(user1);
        (uint balance2,,) = stakingPool.getStakePosition(user2);
        (uint balance3,,) = stakingPool.getStakePosition(user3);

        assertEq(balance1, 10 ether);
        assertEq(balance2, 5 ether);
        assertEq(balance3, 3 ether);
        assertEq(stakingPool.getTotalStakingPower(), 18 ether);
    }

    function test_StakeAfterWithdraw() public {
        stakeFor(user1, 10 ether);
        withdrawFor(user1);

        stakeFor(user1, 5 ether);
        (uint balance,,) = stakingPool.getStakePosition(user1);
        assertEq(balance, 5 ether);
    }
}
