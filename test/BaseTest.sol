// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../src/RewardPool.sol";
import "../src/StakingPool.sol";
import "../src/StakingProtocol.sol";
import "../src/StakingRouter.sol";
import "../src/WCROSS.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Test, console} from "forge-std/Test.sol";

contract MockERC20 is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _mint(msg.sender, 1000000 ether);
    }

    function mint(address to, uint amount) external {
        _mint(to, amount);
    }
}

/**
 * @title BaseTest
 * @notice 모든 테스트의 기본 설정
 */
contract BaseTest is Test {
    // 컨트랙트
    WCROSS public wcross;
    StakingPool public stakingPool;
    RewardPool public rewardPool;
    StakingRouter public router;
    MockERC20 public rewardToken;
    StakingProtocol public protocol;

    // 계정
    address public owner = address(1);
    address public user1 = address(2);
    address public user2 = address(3);
    address public user3 = address(4);
    address public rewardProvider = address(5);

    // 상수
    uint public constant INITIAL_BALANCE = 10000 ether;
    uint public constant SEASON_BLOCKS = 100;
    uint public constant PROJECT_ID = 1;

    function setUp() public virtual {
        vm.startPrank(owner);

        // 1. WCROSS 배포
        wcross = new WCROSS();

        // 2. Code 컨트랙트 배포
        StakingPoolCode stakingPoolCode = new StakingPoolCode();
        RewardPoolCode rewardPoolCode = new RewardPoolCode();

        // 3. StakingProtocol 배포
        protocol = new StakingProtocol(address(wcross), address(stakingPoolCode), address(rewardPoolCode), owner);

        // 4. 프로젝트 생성 (Code 패턴 사용)
        uint firstSeasonStartBlock = block.number; // 현재 블록에서 시작
        uint poolEndBlock = 0; // 무한 진행
        (, address stakingPoolAddr, address rewardPoolAddr) =
            protocol.createProject("TestProject", SEASON_BLOCKS, firstSeasonStartBlock, poolEndBlock, address(0)); // admin은 msg.sender(owner)로 자동 설정

        stakingPool = StakingPool(stakingPoolAddr);
        rewardPool = RewardPool(rewardPoolAddr);

        // 5. StakingRouter 배포 및 승인
        router = new StakingRouter(address(wcross), address(protocol));
        protocol.setApprovedRouter(PROJECT_ID, address(router), true);

        // 6. Reward Token 배포
        rewardToken = new MockERC20("Reward Token", "REWARD");
        rewardToken.mint(rewardProvider, 1000000 ether);

        vm.stopPrank();

        // 7. 첫 시즌 실제 생성 (Virtual → Real)
        // owner에게 ETH 지급 후 트리거
        vm.deal(owner, 100 ether);
        vm.startPrank(owner);
        wcross.deposit{value: 10 ether}();
        wcross.approve(address(stakingPool), 1 ether);
        stakingPool.stake(1 ether); // _ensureSeason() → _startFirstSeason() 호출
        stakingPool.withdrawAll(); // 즉시 출금하여 초기 상태 유지
        vm.stopPrank();

        // 사용자에게 Native CROSS 지급
        vm.deal(user1, INITIAL_BALANCE);
        vm.deal(user2, INITIAL_BALANCE);
        vm.deal(user3, INITIAL_BALANCE);
        vm.deal(rewardProvider, INITIAL_BALANCE);
    }

    // Helper 함수들
    function stakeFor(address user, uint amount) internal {
        vm.prank(user);
        router.stake{value: amount}(PROJECT_ID);
    }

    function withdrawFor(address user) internal {
        vm.prank(user);
        stakingPool.withdrawAll();
    }

    function rolloverSeason() internal {
        // 실제 컨트랙트의 seasonBlocks 값을 사용
        uint seasonBlocks = stakingPool.seasonBlocks();
        vm.roll(block.number + seasonBlocks + 1);
        stakingPool.rolloverSeason();
    }

    function fundSeason(uint season, uint amount) internal {
        vm.startPrank(rewardProvider);
        rewardToken.approve(address(rewardPool), amount);
        rewardPool.fundSeason(season, address(rewardToken), amount);
        vm.stopPrank();
    }

    function claimSeasonFor(address user, uint season) internal {
        vm.prank(user);
        stakingPool.claimSeason(season, address(rewardToken));
    }

    function updatePointsFor(address user) internal {
        // updatePoints는 protocol 또는 rewardPool만 호출 가능
        vm.prank(address(protocol));
        stakingPool.updatePoints(user);
    }
}
