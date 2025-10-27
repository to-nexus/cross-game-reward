// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/StakingProtocol.sol";

import "../src/StakingRouter.sol";
import "../src/StakingViewer.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title DeployRouter
 * @notice StakingRouter 및 StakingViewer 배포 및 승인 스크립트
 * @dev 환경변수:
 *      - WCROSS_ADDRESS: WCROSS 주소
 *      - STAKING_PROTOCOL_ADDRESS: StakingProtocol 주소
 *      - PROJECT_ID: 라우터를 승인할 프로젝트 ID
 * @dev 스크립트를 실행하는 계정이 Router를 배포하고 승인합니다
 *      (프로젝트 admin 또는 프로토콜 admin 권한 필요)
 */
contract DeployRouterScript is Script {
    function run() external {
        address wcrossAddress = vm.envAddress("WCROSS_ADDRESS");
        address protocolAddress = vm.envAddress("STAKING_PROTOCOL_ADDRESS");
        uint projectID = vm.envUint("PROJECT_ID");

        vm.startBroadcast();

        console.log("=== Deploying StakingRouter & StakingViewer ===");
        console.log("Deployer:", msg.sender);
        console.log("WCROSS:", wcrossAddress);
        console.log("Protocol:", protocolAddress);
        console.log("Project ID:", projectID);

        // 1. Router 배포
        console.log("\n[1/3] Deploying Router...");
        StakingRouter router = new StakingRouter(wcrossAddress, protocolAddress);
        console.log("Router deployed at:", address(router));

        // 2. Viewer 배포
        console.log("\n[2/3] Deploying Viewer...");
        StakingViewer viewer = new StakingViewer(protocolAddress);
        console.log("Viewer deployed at:", address(viewer));

        // 3. Router 승인
        console.log("\n[3/3] Approving Router...");
        StakingProtocol protocol = StakingProtocol(protocolAddress);
        protocol.setApprovedRouter(projectID, address(router), true);
        console.log("Router approved for Project", projectID);

        console.log("\n=== Deployment Summary ===");
        console.log("StakingRouter:", address(router));
        console.log("StakingViewer:", address(viewer));
        console.log("Approved for Project ID:", projectID);

        console.log("\n=== Save these addresses ===");
        console.log("STAKING_ROUTER_ADDRESS=", address(router));
        console.log("STAKING_VIEWER_ADDRESS=", address(viewer));

        vm.stopBroadcast();
    }
}
