// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/StakingViewer.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title DeployViewer
 * @notice StakingViewer만 재배포하는 스크립트
 * @dev 환경변수:
 *      - STAKING_PROTOCOL_ADDRESS: 기존 StakingProtocol 주소 (필수)
 */
contract DeployViewerScript is Script {
    function run() external {
        address protocolAddress = vm.envAddress("STAKING_PROTOCOL_ADDRESS");

        vm.startBroadcast();

        console.log("=== StakingViewer Redeployment ===");
        console.log("Deployer:", msg.sender);
        console.log("StakingProtocol:", protocolAddress);

        // StakingViewer 배포
        console.log("\nDeploying StakingViewer...");
        StakingViewer viewer = new StakingViewer(protocolAddress);
        console.log("StakingViewer deployed at:", address(viewer));

        console.log("\n=== Deployment Complete ===");
        console.log("STAKING_VIEWER_ADDRESS=", address(viewer));

        vm.stopBroadcast();
    }
}
