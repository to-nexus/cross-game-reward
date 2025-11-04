// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../../src/CrossStaking.sol";
import "../../src/CrossStakingPool.sol";
import "../../src/interfaces/ICrossStakingPool.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Test.sol";

/**
 * @title MockERC20
 * @notice Simple ERC20 token for testing
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
 * @notice Shared base contract used by all pool tests
 * @dev Provides setup and convenience helpers
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
        // Deploy mock tokens
        crossToken = new MockERC20("CROSS Token", "CROSS");
        rewardToken1 = new MockERC20("Reward Token 1", "RWD1");
        rewardToken2 = new MockERC20("Reward Token 2", "RWD2");

        // Deploy CrossStaking via UUPS pattern
        CrossStakingPool poolImplementation = new CrossStakingPool();
        CrossStaking stakingImplementation = new CrossStaking();

        bytes memory initData =
            abi.encodeCall(CrossStaking.initialize, (ICrossStakingPool(address(poolImplementation)), owner, 2 days));

        ERC1967Proxy proxy = new ERC1967Proxy(address(stakingImplementation), initData);
        crossStaking = CrossStaking(address(proxy));

        // Create pool through the CrossStaking factory
        (uint poolId, ICrossStakingPool poolInterface) = crossStaking.createPool(IERC20(address(crossToken)), 1 ether);
        pool = CrossStakingPool(address(poolInterface));

        // Distribute CROSS tokens to users
        crossToken.transfer(user1, 1000 ether);
        crossToken.transfer(user2, 1000 ether);
        crossToken.transfer(user3, 1000 ether);

        // Allocate reward tokens to the owner account
        rewardToken1.transfer(owner, 100000 ether);
        rewardToken2.transfer(owner, 100000 ether);

        // Register reward tokens via CrossStaking
        crossStaking.addRewardToken(poolId, IERC20(address(rewardToken1)));
        crossStaking.addRewardToken(poolId, IERC20(address(rewardToken2)));
    }

    // ==================== Helper functions ====================

    /**
     * @notice Helper to stake on behalf of a user
     */
    function _userStake(address user, uint amount) internal {
        vm.startPrank(user);
        crossToken.approve(address(pool), amount);
        pool.stake(amount);
        vm.stopPrank();
    }

    /**
     * @notice Helper to deposit rewards into the pool
     * @dev Direct transfers are detected automatically by _syncReward
     */
    function _depositReward(address rewardToken, uint amount) internal {
        vm.startPrank(owner);
        IERC20(rewardToken).transfer(address(pool), amount);
        vm.stopPrank();
    }

    /**
     * @notice Helper to advance time by whole days
     */
    function _warpDays(uint days_) internal {
        vm.warp(block.timestamp + days_ * 1 days);
    }

    /**
     * @notice Helper to advance time by seconds
     */
    function _warpSeconds(uint seconds_) internal {
        vm.warp(block.timestamp + seconds_);
    }
}
