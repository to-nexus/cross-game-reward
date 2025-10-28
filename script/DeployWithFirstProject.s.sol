// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/RewardPool.sol";

import "../src/StakingPool.sol";
import "../src/StakingProtocol.sol";
import "../src/StakingRouter.sol";
import "../src/StakingViewer.sol";
import "../src/WCROSS.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title DeployWithFirstProject
 * @notice 전체 시스템 배포 + 첫 프로젝트 생성 + 선택적 리워드 예치를 한번에 실행
 * @dev 환경변수:
 *      - WCROSS_ADDRESS: (선택) 기존 WCROSS 주소 (없으면 새로 배포)
 *      - PROJECT_NAME: 프로젝트 이름
 *      - SEASON_BLOCKS: 시즌 길이 (블록 수, 0이면 기본값)
 *      - FIRST_SEASON_START_BLOCK: 첫 시즌 시작 블록
 *      - PRE_DEPOSIT_START_BLOCK: 사전 디폴트 블록
 *      - PROJECT_ADMIN: (선택) 프로젝트 관리자 주소 (없으면 생성자)
 *      - REWARD_TOKEN: (선택) 리워드 토큰 주소
 *      - REWARD_AMOUNT: (선택) 리워드 수량 (wei 단위)
 *      - REWARD_SEASON: (선택) 리워드를 예치할 시즌 번호 (기본값: 1)
 * @dev 스크립트를 실행하는 계정이 Protocol Admin이 됩니다
 */
contract DeployWithFirstProjectScript is Script {
    // Deployment state
    struct DeployedContracts {
        WCROSS wcross;
        StakingPoolCode stakingPoolCode;
        RewardPoolCode rewardPoolCode;
        StakingProtocol protocol;
        StakingRouter router;
        StakingViewer viewer;
    }

    struct ProjectConfig {
        string name;
        uint seasonBlocks;
        uint firstSeasonStartBlock;
        uint preDepositStartBlock;
        address admin;
    }

    function run() external {
        console.log("=== Cross Staking Complete Deployment ===");
        console.log("Deployer:", msg.sender);

        ProjectConfig memory config = _loadProjectConfig();
        
        vm.startBroadcast();
        
        DeployedContracts memory contracts = _deployContracts();
        
        (uint projectID, address stakingPool, address rewardPool) = 
            _createFirstProject(contracts.protocol, config);
        
        _fundRewardIfConfigured(rewardPool);
        
        vm.stopBroadcast();
        
        _printSummary(contracts, projectID, stakingPool, rewardPool, config.admin);
    }

    function _loadProjectConfig() internal view returns (ProjectConfig memory config) {
        config.name = vm.envString("PROJECT_NAME");
        config.seasonBlocks = vm.envUint("SEASON_BLOCKS");
        config.firstSeasonStartBlock = vm.envUint("FIRST_SEASON_START_BLOCK");
        config.preDepositStartBlock = vm.envUint("PRE_DEPOSIT_START_BLOCK");
        
        try vm.envAddress("PROJECT_ADMIN") returns (address admin) {
            config.admin = admin;
        } catch {
            config.admin = msg.sender;
        }
    }

    function _deployContracts() internal returns (DeployedContracts memory contracts) {
        // WCROSS
        try vm.envAddress("WCROSS_ADDRESS") returns (address wcrossAddr) {
            console.log("\n[Using existing WCROSS]");
            console.log("WCROSS Address:", wcrossAddr);
            contracts.wcross = WCROSS(payable(wcrossAddr));
        } catch {
            console.log("\n[1/5] Deploying WCROSS...");
            contracts.wcross = new WCROSS();
            console.log("WCROSS deployed at:", address(contracts.wcross));
        }

        // Code Contracts
        console.log("\n[2/5] Deploying Code Contracts...");
        contracts.stakingPoolCode = new StakingPoolCode();
        contracts.rewardPoolCode = new RewardPoolCode();
        console.log("StakingPoolCode:", address(contracts.stakingPoolCode));
        console.log("RewardPoolCode:", address(contracts.rewardPoolCode));

        // Protocol
        console.log("\n[3/5] Deploying StakingProtocol...");
        contracts.protocol = new StakingProtocol(
            address(contracts.wcross),
            address(contracts.stakingPoolCode),
            address(contracts.rewardPoolCode),
            msg.sender
        );
        console.log("StakingProtocol:", address(contracts.protocol));

        // Router
        console.log("\n[4/5] Deploying StakingRouter...");
        contracts.router = new StakingRouter(address(contracts.wcross), address(contracts.protocol));
        console.log("StakingRouter:", address(contracts.router));

        // Viewer
        console.log("\n[5/5] Deploying StakingViewer...");
        contracts.viewer = new StakingViewer(address(contracts.protocol));
        console.log("StakingViewer:", address(contracts.viewer));

        // Approve Router
        console.log("\nApproving Router globally...");
        contracts.protocol.setGlobalApprovedRouter(address(contracts.router), true);
        console.log("Router globally approved");
    }

    function _createFirstProject(StakingProtocol protocol, ProjectConfig memory config)
        internal
        returns (uint projectID, address stakingPool, address rewardPool)
    {
        console.log("\nCreating First Project...");
        console.log("Project Name:", config.name);
        console.log("Season Blocks:", config.seasonBlocks);
        console.log("Pre-deposit Start:", config.preDepositStartBlock);
        console.log("First Season Start:", config.firstSeasonStartBlock);
        console.log("Project Admin:", config.admin);

        (projectID, stakingPool, rewardPool) = protocol.createProject(
            config.name,
            config.seasonBlocks,
            config.firstSeasonStartBlock,
            0,
            config.admin,
            config.preDepositStartBlock
        );

        console.log("Project ID:", projectID);
        console.log("StakingPool:", stakingPool);
        console.log("RewardPool:", rewardPool);
    }

    function _fundRewardIfConfigured(address rewardPool) internal {
        try vm.envAddress("REWARD_TOKEN") returns (address rewardToken) {
            uint rewardAmount = vm.envUint("REWARD_AMOUNT");
            uint rewardSeason;
            try vm.envUint("REWARD_SEASON") returns (uint season) {
                rewardSeason = season;
            } catch {
                rewardSeason = 1;
            }

            console.log("\nFunding Reward Pool...");
            console.log("Reward Token:", rewardToken);
            console.log("Reward Amount:", rewardAmount);
            console.log("Target Season:", rewardSeason);

            IERC20 token = IERC20(rewardToken);
            require(token.balanceOf(msg.sender) >= rewardAmount, "Insufficient token balance");

            token.approve(rewardPool, rewardAmount);
            IRewardPool(rewardPool).fundSeason(rewardSeason, rewardToken, rewardAmount);
            console.log("Reward funded successfully");
        } catch {
            console.log("\nSkipping reward funding (REWARD_TOKEN not specified)");
        }
    }

    function _printSummary(
        DeployedContracts memory contracts,
        uint projectID,
        address stakingPool,
        address rewardPool,
        address projectAdmin
    ) internal view {
        console.log("\n=== Deployment Summary ===");
        console.log("WCROSS:", address(contracts.wcross));
        console.log("StakingPoolCode:", address(contracts.stakingPoolCode));
        console.log("RewardPoolCode:", address(contracts.rewardPoolCode));
        console.log("StakingProtocol:", address(contracts.protocol));
        console.log("StakingRouter:", address(contracts.router));
        console.log("StakingViewer:", address(contracts.viewer));
        console.log("Protocol Admin:", msg.sender);
        console.log("");
        console.log("Project ID:", projectID);
        console.log("StakingPool:", stakingPool);
        console.log("RewardPool:", rewardPool);
        console.log("Project Admin:", projectAdmin);

        console.log("\n=== Save these addresses for webapp ===");
        console.log("VITE_WCROSS_ADDRESS=", address(contracts.wcross));
        console.log("VITE_STAKING_PROTOCOL_ADDRESS=", address(contracts.protocol));
        console.log("VITE_STAKING_ROUTER_ADDRESS=", address(contracts.router));
        console.log("VITE_STAKING_VIEWER_ADDRESS=", address(contracts.viewer));
        console.log("VITE_DEFAULT_PROJECT_ID=", projectID);

        console.log("\n=== Save these for future scripts ===");
        console.log("WCROSS_ADDRESS=", address(contracts.wcross));
        console.log("STAKING_PROTOCOL_ADDRESS=", address(contracts.protocol));
        console.log("STAKING_ROUTER_ADDRESS=", address(contracts.router));
        console.log("STAKING_VIEWER_ADDRESS=", address(contracts.viewer));
        console.log("PROJECT_ID=", projectID);
        console.log("STAKING_POOL_ADDRESS=", stakingPool);
        console.log("REWARD_POOL_ADDRESS=", rewardPool);

        console.log("\n=== Deployment Complete! ===");
    }
}
