// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardPool.sol";
import "../src/GamePool.sol";
import "forge-std/Script.sol";

/**
 * @title DeployImpl
 * @notice Implementation 컨트랙트만 배포하는 스크립트
 * @dev CrossGameReward, CrossGameRewardPool (V1), GamePool를 배포합니다.
 *
 * 사용법:
 * 1. 전체 배포 (기본 - V1 + CrossGameReward):
 *    forge script script/DeployImpl.s.sol:DeployImpl \
 *      --rpc-url <RPC_URL> --broadcast
 *
 * 2. 전체 배포 (V1 + GamePool + CrossGameReward):
 *    forge script script/DeployImpl.s.sol:DeployImpl \
 *      --sig "deployAll()" \
 *      --rpc-url <RPC_URL> --broadcast
 *
 * 3. Pool V1 Implementation만 배포:
 *    forge script script/DeployImpl.s.sol:DeployImpl \
 *      --sig "deployPool()" \
 *      --rpc-url <RPC_URL> --broadcast
 *
 * 4. Pool GamePool Implementation만 배포:
 *    forge script script/DeployImpl.s.sol:DeployImpl \
 *      --sig "deployGamePool()" \
 *      --rpc-url <RPC_URL> --broadcast
 *
 * 5. CrossGameReward Implementation만 배포:
 *    forge script script/DeployImpl.s.sol:DeployImpl \
 *      --sig "deployCrossGameReward()" \
 *      --rpc-url <RPC_URL> --broadcast
 */
contract DeployImpl is Script {
    /**
     * @notice 기본 실행 - V1 Pool + CrossGameReward Implementation 배포
     */
    function run() external {
        console.log("Deployer:", msg.sender);
        console.log("Chain ID:", block.chainid);

        vm.startBroadcast();

        CrossGameRewardPool poolImpl = new CrossGameRewardPool();
        console.log("\n1. Pool V1 Implementation:", address(poolImpl));

        CrossGameReward cgrImpl = new CrossGameReward();
        console.log("2. CrossGameReward Implementation:", address(cgrImpl));

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("Pool V1 Implementation:", address(poolImpl));
        console.log("CrossGameReward Implementation:", address(cgrImpl));
        console.log("\n=== Environment Variables ===");
        console.log("POOL_IMPLEMENTATION=", address(poolImpl));
        console.log("CROSS_GAME_REWARD_ROOT_IMPLEMENTATION=", address(cgrImpl));
    }

    /**
     * @notice 전체 배포 - V1 + GamePool + CrossGameReward Implementation
     * @dev --sig "deployAll()" 플래그와 함께 사용
     */
    function deployAll() external {
        console.log("Deployer:", msg.sender);
        console.log("Chain ID:", block.chainid);

        vm.startBroadcast();

        CrossGameRewardPool poolImpl = new CrossGameRewardPool();
        console.log("\n1. Pool V1 Implementation:", address(poolImpl));

        GamePool gamePoolImpl = new GamePool();
        console.log("2. Pool GamePool Implementation:", address(gamePoolImpl));

        CrossGameReward cgrImpl = new CrossGameReward();
        console.log("3. CrossGameReward Implementation:", address(cgrImpl));

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("Pool V1 Implementation:", address(poolImpl));
        console.log("Pool GamePool Implementation:", address(gamePoolImpl));
        console.log("CrossGameReward Implementation:", address(cgrImpl));
        console.log("\n=== Environment Variables ===");
        console.log("POOL_IMPLEMENTATION=", address(poolImpl));
        console.log("GAME_POOL_IMPLEMENTATION=", address(gamePoolImpl));
        console.log("CROSS_GAME_REWARD_ROOT_IMPLEMENTATION=", address(cgrImpl));
    }

    /**
     * @notice Pool V1 Implementation만 배포
     * @dev --sig "deployPool()" 플래그와 함께 사용
     */
    function deployPool() external {
        console.log("Deployer:", msg.sender);
        console.log("Chain ID:", block.chainid);

        vm.startBroadcast();

        CrossGameRewardPool poolImpl = new CrossGameRewardPool();

        vm.stopBroadcast();

        console.log("\nPool V1 Implementation:", address(poolImpl));
        console.log("Use: setPoolImplementation(", address(poolImpl), ")");
    }

    /**
     * @notice Pool GamePool Implementation만 배포
     * @dev --sig "deployGamePool()" 플래그와 함께 사용
     */
    function deployGamePool() external {
        console.log("Deployer:", msg.sender);
        console.log("Chain ID:", block.chainid);

        vm.startBroadcast();

        GamePool gamePoolImpl = new GamePool();

        vm.stopBroadcast();

        console.log("\nPool GamePool Implementation:", address(gamePoolImpl));
        console.log("Use: setGamePoolImplementation(", address(gamePoolImpl), ")");
    }

    /**
     * @notice CrossGameReward Implementation만 배포
     * @dev --sig "deployCrossGameReward()" 플래그와 함께 사용
     */
    function deployCrossGameReward() external {
        console.log("Deployer:", msg.sender);
        console.log("Chain ID:", block.chainid);

        vm.startBroadcast();

        CrossGameReward cgrImpl = new CrossGameReward();

        vm.stopBroadcast();

        console.log("\nCrossGameReward Implementation:", address(cgrImpl));
        console.log("Use: upgradeToAndCall(", address(cgrImpl), ', "")');
    }
}
