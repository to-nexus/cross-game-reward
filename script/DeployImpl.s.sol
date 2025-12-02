// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardPool.sol";
import "forge-std/Script.sol";

/**
 * @title DeployImpl
 * @notice Implementation 컨트랙트만 배포하는 스크립트
 * @dev CrossGameReward와 CrossGameRewardPool의 Implementation을 배포합니다.
 *      업그레이드나 새로운 시스템 배포 시 사용할 수 있습니다.
 *
 * 사용법:
 * 1. 전체 배포 (기본):
 *    forge script script/DeployImpl.s.sol:DeployImpl \
 *      --rpc-url <RPC_URL> \
 *      --broadcast
 *
 * 2. Pool Implementation만 배포:
 *    forge script script/DeployImpl.s.sol:DeployImpl \
 *      --sig "deployPool()" \
 *      --rpc-url <RPC_URL> \
 *      --broadcast
 *
 * 3. CrossGameReward Implementation만 배포:
 *    forge script script/DeployImpl.s.sol:DeployImpl \
 *      --sig "deployCrossGameReward()" \
 *      --rpc-url <RPC_URL> \
 *      --broadcast
 */
contract DeployImpl is Script {
    /**
     * @notice 기본 실행 함수 - 모든 Implementation 배포
     */
    function run() external {
        address deployer = msg.sender;

        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);

        vm.startBroadcast();

        // 1. CrossGameRewardPool Implementation 배포
        CrossGameRewardPool poolImplementation = new CrossGameRewardPool();
        console.log("\n1. Pool Implementation deployed:", address(poolImplementation));

        // 2. CrossGameReward Implementation 배포
        CrossGameReward crossGameRewardImpl = new CrossGameReward();
        console.log("2. CrossGameReward Implementation deployed:", address(crossGameRewardImpl));

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("CrossGameReward Implementation:", address(crossGameRewardImpl));
        console.log("Pool Implementation:", address(poolImplementation));
        console.log("\n=== Next Steps ===");
        console.log("1. Use these addresses in DeployFullSystem with environment variables:");
        console.log("   CROSS_GAME_REWARD_ROOT_IMPLEMENTATION=", address(crossGameRewardImpl));
        console.log("   POOL_IMPLEMENTATION=", address(poolImplementation));
        console.log("2. Or use them for upgrading existing proxy contracts");
    }

    /**
     * @notice CrossGameRewardPool Implementation만 배포
     * @dev --sig "deployPool()" 플래그와 함께 사용
     */
    function deployPool() external {
        address deployer = msg.sender;

        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);

        vm.startBroadcast();

        CrossGameRewardPool poolImplementation = new CrossGameRewardPool();
        console.log("\nPool Implementation deployed:", address(poolImplementation));

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("Pool Implementation:", address(poolImplementation));
        console.log("\n=== Next Steps ===");
        console.log("Use this address for:");
        console.log("1. Upgrading existing pool proxies");
        console.log("2. Setting as poolImplementation in CrossGameReward");
        console.log("   - Call setPoolImplementation(", address(poolImplementation), ")");
    }

    /**
     * @notice CrossGameReward Implementation만 배포
     * @dev --sig "deployCrossGameReward()" 플래그와 함께 사용
     */
    function deployCrossGameReward() external {
        address deployer = msg.sender;

        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);

        vm.startBroadcast();

        CrossGameReward crossGameRewardImpl = new CrossGameReward();
        console.log("\nCrossGameReward Implementation deployed:", address(crossGameRewardImpl));

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("CrossGameReward Implementation:", address(crossGameRewardImpl));
        console.log("\n=== Next Steps ===");
        console.log("Use this address for:");
        console.log("1. Upgrading existing CrossGameReward proxy");
        console.log("   - Call upgradeToAndCall(", address(crossGameRewardImpl), ", \"\")");
        console.log("2. Creating new CrossGameReward proxy with this implementation");
    }
}
