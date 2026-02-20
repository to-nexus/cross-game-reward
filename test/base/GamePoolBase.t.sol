// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../../src/CrossGameReward.sol";
import "../../src/CrossGameRewardPool.sol";
import "../../src/GamePool.sol";
import "../../src/interfaces/ICrossGameRewardPool.sol";
import "../../src/interfaces/IGamePool.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Test.sol";

/**
 * @title MockERC20GP
 * @notice Simple ERC20 token for GamePool testing
 */
contract MockERC20GP is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _mint(msg.sender, 10000000 ether);
    }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

/**
 * @title GamePoolBase
 * @notice Shared base contract used by all GamePool tests
 * @dev Provides setup and convenience helpers for GamePool testing
 */
abstract contract GamePoolBase is Test {
    CrossGameReward public crossGameReward;
    GamePool public gamePool;
    uint256 public poolId;

    MockERC20GP public gameToken; // Deposit token (game token)
    MockERC20GP public crossdToken; // Reward token (CROSSD)

    address public owner = address(this);
    address public sponsor = address(0x1001);
    address public sponsor2 = address(0x1002);
    address public user1 = address(0x2001);
    address public user2 = address(0x2002);
    address public user3 = address(0x2003);

    uint256 public constant MIN_DEPOSIT = 1 ether;
    uint256 public constant DEFAULT_ROUND_DURATION = 150 days * 86400; // 150 days in blocks (1 block = 1 second)

    function setUp() public virtual {
        // Deploy mock tokens
        gameToken = new MockERC20GP("Game Token", "GAME");
        crossdToken = new MockERC20GP("CROSSD Stablecoin", "CROSSD");

        // Deploy CrossGameReward with both V1 and GamePool implementations
        CrossGameRewardPool poolImplementation = new CrossGameRewardPool();
        GamePool gamePoolImpl = new GamePool();
        CrossGameReward gameRewardImplementation = new CrossGameReward();

        bytes memory initData = abi.encodeCall(
            CrossGameReward.initialize, (ICrossGameRewardPool(address(poolImplementation)), owner, 2 days)
        );

        ERC1967Proxy proxy = new ERC1967Proxy(address(gameRewardImplementation), initData);
        crossGameReward = CrossGameReward(address(proxy));

        // Set GamePool implementation
        crossGameReward.setGamePoolImplementation(ICrossGameRewardPool(address(gamePoolImpl)));

        // Create GamePool through the CrossGameReward factory
        (poolId,) = crossGameReward.createGamePool(
            "Test Game Pool", IERC20(address(gameToken)), IERC20(address(crossdToken)), MIN_DEPOSIT
        );
        gamePool = GamePool(address(crossGameReward.getPoolAddress(poolId)));

        // Grant sponsor roles
        crossGameReward.grantSponsorRole(poolId, sponsor);
        crossGameReward.grantSponsorRole(poolId, sponsor2);

        // Distribute game tokens to users
        gameToken.transfer(user1, 10000 ether);
        gameToken.transfer(user2, 10000 ether);
        gameToken.transfer(user3, 10000 ether);

        // Allocate CROSSD to sponsors for creating rounds
        crossdToken.transfer(sponsor, 1000000 ether);
        crossdToken.transfer(sponsor2, 1000000 ether);
    }

    // ==================== Helper functions ====================

    /**
     * @notice Helper to deposit on behalf of a user
     */
    function _userDeposit(address user, uint256 amount) internal {
        vm.startPrank(user);
        gameToken.approve(address(gamePool), amount);
        gamePool.deposit(amount);
        vm.stopPrank();
    }

    /**
     * @notice Helper to create a round as sponsor
     * @param amount Total reward amount
     * @param startBlockOffset Blocks from current block to start
     * @param durationBlocks Duration in blocks
     */
    function _createRound(uint256 amount, uint256 startBlockOffset, uint256 durationBlocks)
        internal
        returns (uint256 roundId)
    {
        vm.startPrank(sponsor);
        crossdToken.approve(address(gamePool), amount);
        roundId = gamePool.createRound(amount, block.number + startBlockOffset, durationBlocks);
        vm.stopPrank();
    }

    /**
     * @notice Helper to create a round with default duration
     */
    function _createRoundDefault(uint256 amount, uint256 startBlockOffset) internal returns (uint256 roundId) {
        return _createRound(amount, startBlockOffset, DEFAULT_ROUND_DURATION);
    }

    /**
     * @notice Helper to advance blocks
     */
    function _advanceBlocks(uint256 blocks) internal {
        vm.roll(block.number + blocks);
        vm.warp(block.timestamp + blocks); // Assume 1 block = 1 second
    }

    /**
     * @notice Helper to advance to specific block
     */
    function _advanceToBlock(uint256 targetBlock) internal {
        require(targetBlock > block.number, "Target block must be in future");
        uint256 blocksToAdvance = targetBlock - block.number;
        _advanceBlocks(blocksToAdvance);
    }

    /**
     * @notice Helper to get pending reward for a user
     */
    function _getPendingReward(address user) internal view returns (uint256) {
        return gamePool.pendingReward(user, crossdToken);
    }
}
