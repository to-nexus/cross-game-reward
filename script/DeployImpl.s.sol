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
 * forge script script/DeployImpl.s.sol:DeployImpl \
 *   --rpc-url <RPC_URL> \
 *   --private-key <PRIVATE_KEY> \
 *   --broadcast
 */
contract DeployImpl is Script {
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
}

/*
 * ===================================
 * DeployImpl.env.example
 * ===================================
 *
 * # DeployImpl 스크립트는 환경변수가 필요 없습니다.
 * # Implementation 컨트랙트만 배포하며, 배포 후 주소를 출력합니다.
 * # 출력된 주소를 DeployFullSystem의 환경변수로 사용하거나,
 * # 기존 Proxy 컨트랙트의 업그레이드에 사용할 수 있습니다.
 *
 * # 사용 예시:
 * # forge script script/DeployImpl.s.sol:DeployImpl \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --broadcast
 */
