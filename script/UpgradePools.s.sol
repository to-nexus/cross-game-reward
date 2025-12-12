// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardPool.sol";
import "../src/interfaces/ICrossGameRewardPool.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "forge-std/Script.sol";

/**
 * @title UpgradePools
 * @notice CrossGameReward에 등록된 Pool의 로직을 업그레이드하는 스크립트
 * @dev UUPS 업그레이드 패턴을 사용하여 기존 풀들을 새로운 구현체로 업그레이드
 *
 * 사용법:
 * forge script script/UpgradePools.s.sol:UpgradePools \
 *   --rpc-url <RPC_URL> \
 *   --broadcast
 *
 * 필수 환경변수:
 * - CROSS_GAME_REWARD: CrossGameReward 컨트랙트 주소
 * - NEW_POOL_IMPLEMENTATION: 새로운 Pool 구현체 주소
 */
contract UpgradePools is Script {
    function run() external {
        // 환경변수에서 설정 읽기
        address crossGameRewardAddress = vm.envAddress("CROSS_GAME_REWARD");
        address newPoolImplementation = vm.envAddress("NEW_POOL_IMPLEMENTATION");

        // CrossGameReward 컨트랙트 인스턴스 생성
        CrossGameReward crossGameReward = CrossGameReward(crossGameRewardAddress);

        console.log("\n=== Pool Upgrade Configuration ===");
        console.log("CrossGameReward Address:", crossGameRewardAddress);
        console.log("New Pool Implementation:", newPoolImplementation);
        console.log("Deployer:", msg.sender);

        vm.startBroadcast();

        // 1. CrossGameReward의 poolImplementation 업데이트
        console.log("\n=== Step 1: Update Pool Implementation in CrossGameReward ===");
        crossGameReward.setPoolImplementation(ICrossGameRewardPool(newPoolImplementation));
        console.log("Pool implementation updated in CrossGameReward");

        // 2. 모든 풀 ID 가져오기
        console.log("\n=== Step 2: Get All Pool IDs ===");
        uint[] memory poolIds = crossGameReward.getAllPoolIds();
        console.log("Total Pools:", poolIds.length);

        // 3. 각 풀을 새로운 구현체로 업그레이드
        console.log("\n=== Step 3: Upgrade Each Pool ===");
        for (uint i = 0; i < poolIds.length; i++) {
            uint poolId = poolIds[i];
            ICrossGameRewardPool pool = crossGameReward.getPoolAddress(poolId);

            console.log("\nUpgrading Pool ID:", poolId);
            console.log("Pool Address:", address(pool));

            // UUPS 업그레이드 (reinitialize 없이)
            CrossGameRewardPool(address(pool)).upgradeToAndCall(newPoolImplementation, "");

            console.log("Pool ID", poolId, "upgraded successfully");
        }

        vm.stopBroadcast();

        console.log("\n=== Upgrade Summary ===");
        console.log("Total Pools Upgraded:", poolIds.length);
        console.log("New Implementation:", newPoolImplementation);
        console.log("\n=== Upgrade Complete ===");
    }
}
