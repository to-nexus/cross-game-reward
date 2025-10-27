// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/RewardPool.sol";

import "../src/StakingPool.sol";
import "../src/StakingProtocol.sol";
import "../src/StakingRouter.sol";
import "../src/StakingViewer.sol";
import "../src/WCROSS.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title Deploy
 * @notice 전체 시스템 배포 스크립트
 * @dev 환경변수:
 *      - WCROSS_ADDRESS: (선택) 기존 WCROSS 주소 (없으면 새로 배포)
 * @dev 스크립트를 실행하는 계정이 Protocol Admin이 됩니다
 */
contract DeployScript is Script {
    function run() external {
        vm.startBroadcast();

        console.log("=== Cross Staking System Deployment ===");
        console.log("Deployer:", msg.sender);
        console.log("Protocol Admin:", msg.sender);

        // 1. WCROSS 배포 (또는 기존 주소 사용)
        WCROSS wcross;
        try vm.envAddress("WCROSS_ADDRESS") returns (address wcrossAddr) {
            console.log("\n[Using existing WCROSS]");
            console.log("WCROSS Address:", wcrossAddr);
            wcross = WCROSS(payable(wcrossAddr));
        } catch {
            console.log("\n[1/4] Deploying WCROSS...");
            wcross = new WCROSS();
            console.log("WCROSS deployed at:", address(wcross));
        }

        // 2. Code 컨트랙트 배포
        console.log("\n[2/4] Deploying Code Contracts...");
        StakingPoolCode stakingPoolCode = new StakingPoolCode();
        RewardPoolCode rewardPoolCode = new RewardPoolCode();
        console.log("StakingPoolCode:", address(stakingPoolCode));
        console.log("RewardPoolCode:", address(rewardPoolCode));

        // 3. StakingProtocol 배포
        console.log("\n[3/6] Deploying StakingProtocol...");
        StakingProtocol protocol =
            new StakingProtocol(address(wcross), address(stakingPoolCode), address(rewardPoolCode), msg.sender);
        console.log("StakingProtocol:", address(protocol));

        // 4. StakingRouter 배포
        console.log("\n[4/6] Deploying StakingRouter...");
        StakingRouter router = new StakingRouter(address(wcross), address(protocol));
        console.log("StakingRouter:", address(router));

        // 5. StakingViewer 배포
        console.log("\n[5/6] Deploying StakingViewer...");
        StakingViewer viewer = new StakingViewer(address(protocol));
        console.log("StakingViewer:", address(viewer));

        // 6. Router를 글로벌 승인
        console.log("\n[6/6] Approving Router globally...");
        protocol.setGlobalApprovedRouter(address(router), true);
        console.log("Router globally approved");

        // 배포 정보 요약
        console.log("\n=== Deployment Summary ===");
        console.log("WCROSS:", address(wcross));
        console.log("StakingPoolCode:", address(stakingPoolCode));
        console.log("RewardPoolCode:", address(rewardPoolCode));
        console.log("StakingProtocol:", address(protocol));
        console.log("StakingRouter:", address(router));
        console.log("StakingViewer:", address(viewer));
        console.log("Protocol Admin:", msg.sender);

        console.log("\n=== Save these addresses ===");
        console.log("WCROSS_ADDRESS=", address(wcross));
        console.log("STAKING_PROTOCOL_ADDRESS=", address(protocol));
        console.log("STAKING_ROUTER_ADDRESS=", address(router));
        console.log("STAKING_VIEWER_ADDRESS=", address(viewer));

        vm.stopBroadcast();
    }
}
