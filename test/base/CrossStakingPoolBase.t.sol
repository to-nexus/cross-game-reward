// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../../src/CrossStaking.sol";
import "../../src/CrossStakingPool.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "forge-std/Test.sol";

/**
 * @title MockERC20
 * @notice 테스트용 ERC20 토큰
 */
contract MockERC20 is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _mint(msg.sender, 1000000 ether);
    }

    function mint(address to, uint amount) external {
        _mint(to, amount);
    }
}

/**
 * @title CrossStakingPoolBase
 * @notice 모든 테스트의 공통 base 컨트랙트
 * @dev setUp, helper 함수 제공
 */
abstract contract CrossStakingPoolBase is Test {
    CrossStaking public crossStaking;
    CrossStakingPool public pool;
    MockERC20 public crossToken;
    MockERC20 public rewardToken1;
    MockERC20 public rewardToken2;

    address public owner = address(this); // Test contract is the owner
    address public user1 = address(2);
    address public user2 = address(3);
    address public user3 = address(4);

    function setUp() public virtual {
        // 토큰 배포
        crossToken = new MockERC20("CROSS Token", "CROSS");
        rewardToken1 = new MockERC20("Reward Token 1", "RWD1");
        rewardToken2 = new MockERC20("Reward Token 2", "RWD2");

        // CrossStaking 배포 (UUPS 패턴)
        CrossStakingPool poolImplementation = new CrossStakingPool();
        CrossStaking stakingImplementation = new CrossStaking();

        bytes memory initData = abi.encodeCall(CrossStaking.initialize, (address(poolImplementation), owner, 2 days));

        ERC1967Proxy proxy = new ERC1967Proxy(address(stakingImplementation), initData);
        crossStaking = CrossStaking(address(proxy));

        // 풀 생성 (CrossStaking을 통해)
        (uint poolId, address poolAddress) = crossStaking.createPool(address(crossToken), 1 ether);
        pool = CrossStakingPool(poolAddress);

        // 사용자들에게 CROSS 토큰 전송
        crossToken.transfer(user1, 1000 ether);
        crossToken.transfer(user2, 1000 ether);
        crossToken.transfer(user3, 1000 ether);

        // owner에게 보상 토큰 전송
        rewardToken1.transfer(owner, 100000 ether);
        rewardToken2.transfer(owner, 100000 ether);

        // 풀에 보상 토큰 등록 (CrossStaking을 통해)
        crossStaking.addRewardToken(poolId, address(rewardToken1));
        crossStaking.addRewardToken(poolId, address(rewardToken2));
    }

    // ==================== Helper 함수 ====================

    /**
     * @notice 사용자의 스테이킹 helper
     */
    function _userStake(address user, uint amount) internal {
        vm.startPrank(user);
        crossToken.approve(address(pool), amount);
        pool.stake(amount);
        vm.stopPrank();
    }

    /**
     * @notice 보상 입금 helper
     * @dev 직접 transfer하면 _syncReward가 자동 감지
     */
    function _depositReward(address rewardToken, uint amount) internal {
        vm.startPrank(owner);
        IERC20(rewardToken).transfer(address(pool), amount);
        vm.stopPrank();
    }

    /**
     * @notice 시간 이동 helper
     */
    function _warpDays(uint days_) internal {
        vm.warp(block.timestamp + days_ * 1 days);
    }

    /**
     * @notice 시간 이동 helper (초 단위)
     */
    function _warpSeconds(uint seconds_) internal {
        vm.warp(block.timestamp + seconds_);
    }
}
