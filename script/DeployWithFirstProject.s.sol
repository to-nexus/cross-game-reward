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
 *      - POOL_END_BLOCK: 풀 종료 블록 (0이면 무한)
 *      - PROJECT_ADMIN: (선택) 프로젝트 관리자 주소 (없으면 생성자)
 *      - REWARD_TOKEN: (선택) 리워드 토큰 주소
 *      - REWARD_AMOUNT: (선택) 리워드 수량 (wei 단위)
 *      - REWARD_SEASON: (선택) 리워드를 예치할 시즌 번호 (기본값: 1)
 * @dev 스크립트를 실행하는 계정이 Protocol Admin이 됩니다
 */
contract DeployWithFirstProjectScript is Script {
    function run() external {
        // 필수 환경변수
        string memory projectName = vm.envString("PROJECT_NAME");
        uint seasonBlocks = vm.envUint("SEASON_BLOCKS");
        uint firstSeasonStartBlock = vm.envUint("FIRST_SEASON_START_BLOCK");
        uint poolEndBlock = vm.envUint("POOL_END_BLOCK");

        // 선택 환경변수
        address projectAdmin;
        try vm.envAddress("PROJECT_ADMIN") returns (address admin) {
            projectAdmin = admin;
        } catch {
            projectAdmin = msg.sender;
        }

        vm.startBroadcast();

        console.log("=== Cross Staking Complete Deployment ===");
        console.log("Deployer:", msg.sender);
        console.log("Protocol Admin:", msg.sender);

        // ============================================
        // Step 1: WCROSS 배포 (또는 기존 주소 사용)
        // ============================================
        WCROSS wcross;
        try vm.envAddress("WCROSS_ADDRESS") returns (address wcrossAddr) {
            console.log("\n[Using existing WCROSS]");
            console.log("WCROSS Address:", wcrossAddr);
            wcross = WCROSS(payable(wcrossAddr));
        } catch {
            console.log("\n[1/5] Deploying WCROSS...");
            wcross = new WCROSS();
            console.log("WCROSS deployed at:", address(wcross));
        }

        // ============================================
        // Step 2: Code 컨트랙트 배포
        // ============================================
        console.log("\n[2/5] Deploying Code Contracts...");
        StakingPoolCode stakingPoolCode = new StakingPoolCode();
        RewardPoolCode rewardPoolCode = new RewardPoolCode();
        console.log("StakingPoolCode:", address(stakingPoolCode));
        console.log("RewardPoolCode:", address(rewardPoolCode));

        // ============================================
        // Step 3: StakingProtocol 배포
        // ============================================
        console.log("\n[3/6] Deploying StakingProtocol...");
        StakingProtocol protocol =
            new StakingProtocol(address(wcross), address(stakingPoolCode), address(rewardPoolCode), msg.sender);
        console.log("StakingProtocol:", address(protocol));

        // ============================================
        // Step 4: StakingRouter 배포
        // ============================================
        console.log("\n[4/7] Deploying StakingRouter...");
        StakingRouter router = new StakingRouter(address(wcross), address(protocol));
        console.log("StakingRouter:", address(router));

        // ============================================
        // Step 5: StakingViewer 배포
        // ============================================
        console.log("\n[5/7] Deploying StakingViewer...");
        StakingViewer viewer = new StakingViewer(address(protocol));
        console.log("StakingViewer:", address(viewer));

        // Router를 글로벌 승인
        console.log("\n[6/7] Approving Router globally...");
        protocol.setGlobalApprovedRouter(address(router), true);
        console.log("Router globally approved");

        // ============================================
        // Step 6: 첫 프로젝트 생성
        // ============================================
        console.log("\n[6/7] Creating First Project...");
        console.log("Project Name:", projectName);
        console.log("Season Blocks:", seasonBlocks);
        console.log("First Season Start:", firstSeasonStartBlock);
        console.log("Pool End Block:", poolEndBlock);
        console.log("Project Admin:", projectAdmin);

        (uint projectID, address stakingPool, address rewardPool) = protocol.createProject(
            projectName, seasonBlocks, firstSeasonStartBlock, poolEndBlock, projectAdmin, 12105615
        );

        console.log("Project ID:", projectID);
        console.log("StakingPool:", stakingPool);
        console.log("RewardPool:", rewardPool);

        // ============================================
        // Step 7: 선택적 리워드 예치
        // ============================================
        try vm.envAddress("REWARD_TOKEN") returns (address rewardToken) {
            uint rewardAmount = vm.envUint("REWARD_AMOUNT");
            uint rewardSeason;
            try vm.envUint("REWARD_SEASON") returns (uint season) {
                rewardSeason = season;
            } catch {
                rewardSeason = 1;
            }

            console.log("\n[7/7] Funding Reward Pool...");
            console.log("Reward Token:", rewardToken);
            console.log("Reward Amount:", rewardAmount);
            console.log("Target Season:", rewardSeason);

            IERC20 token = IERC20(rewardToken);

            uint balance = token.balanceOf(msg.sender);
            console.log("Your Token Balance:", balance);
            require(balance >= rewardAmount, "Insufficient token balance");

            token.approve(rewardPool, rewardAmount);
            console.log("Token approved");

            IRewardPool(rewardPool).fundSeason(rewardSeason, rewardToken, rewardAmount);
            console.log("Reward funded successfully");
        } catch {
            console.log("\n[7/7] Skipping reward funding (optional)");
            console.log("No REWARD_TOKEN specified");
        }

        // ============================================
        // 배포 정보 요약
        // ============================================
        console.log("\n=== Deployment Summary ===");
        console.log("WCROSS:", address(wcross));
        console.log("StakingPoolCode:", address(stakingPoolCode));
        console.log("RewardPoolCode:", address(rewardPoolCode));
        console.log("StakingProtocol:", address(protocol));
        console.log("StakingRouter:", address(router));
        console.log("StakingViewer:", address(viewer));
        console.log("Protocol Admin:", msg.sender);
        console.log("");
        console.log("Project ID:", projectID);
        console.log("StakingPool:", stakingPool);
        console.log("RewardPool:", rewardPool);
        console.log("Project Admin:", projectAdmin);

        console.log("\n=== Save these addresses for webapp ===");
        console.log("VITE_WCROSS_ADDRESS=", address(wcross));
        console.log("VITE_STAKING_PROTOCOL_ADDRESS=", address(protocol));
        console.log("VITE_STAKING_ROUTER_ADDRESS=", address(router));
        console.log("VITE_STAKING_VIEWER_ADDRESS=", address(viewer));
        console.log("VITE_DEFAULT_PROJECT_ID=", projectID);

        console.log("\n=== Save these for future scripts ===");
        console.log("WCROSS_ADDRESS=", address(wcross));
        console.log("STAKING_PROTOCOL_ADDRESS=", address(protocol));
        console.log("STAKING_ROUTER_ADDRESS=", address(router));
        console.log("STAKING_VIEWER_ADDRESS=", address(viewer));
        console.log("PROJECT_ID=", projectID);
        console.log("STAKING_POOL_ADDRESS=", stakingPool);
        console.log("REWARD_POOL_ADDRESS=", rewardPool);

        vm.stopBroadcast();

        console.log("\n=== Deployment Complete! ===");
    }
}
