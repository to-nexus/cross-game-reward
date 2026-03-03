// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardPool.sol";
import "../src/GamePool.sol";
import "../src/interfaces/ICrossGameReward.sol";
import "../src/interfaces/ICrossGameRewardPool.sol";
import "forge-std/Script.sol";

/**
 * @title Deploy260303
 * @notice 2026.03.03 배포 스크립트 - GamePool 지원을 위한 전체 업그레이드
 * @dev 다음 작업을 순차적으로 수행:
 *      1. Implementation 전체 배포 (CrossPool + GamePool + CrossGameReward)
 *      2. CrossGameReward (Factory) 업그레이드
 *      3. 기존 CrossPool 개별 업그레이드 (기존 풀은 onlyOwner로 개별 업그레이드 필요)
 *      4. GamePool Implementation 등록
 *
 * 필수 환경변수:
 * - CROSS_GAME_REWARD: 기존 CrossGameReward Proxy 주소
 *
 * 사용법:
 * forge script script/260303.s.sol:Deploy260303 \
 *   --rpc-url <RPC_URL> \
 *   --broadcast
 */
contract Deploy260303 is Script {
    function run() external {
        address deployer = msg.sender;
        address crossGameRewardAddr = vm.envAddress("CROSS_GAME_REWARD");

        console.log("===========================================");
        console.log("  2026.03.03 GamePool Deployment Script");
        console.log("===========================================");
        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);
        console.log("CrossGameReward (existing):", crossGameRewardAddr);

        CrossGameReward crossGameReward = CrossGameReward(crossGameRewardAddr);

        vm.startBroadcast();

        // ============================================
        // Step 1: Deploy All Implementations
        // ============================================
        console.log("\n[Step 1] Deploying Implementations...");

        CrossGameRewardPool poolImpl = new CrossGameRewardPool();
        console.log("  CrossPool Implementation:", address(poolImpl));

        GamePool gamePoolImpl = new GamePool();
        console.log("  GamePool Implementation:", address(gamePoolImpl));

        CrossGameReward cgrImpl = new CrossGameReward();
        console.log("  CrossGameReward Implementation:", address(cgrImpl));

        // ============================================
        // Step 2: Upgrade CrossGameReward (Factory)
        // ============================================
        console.log("\n[Step 2] Upgrading CrossGameReward...");
        crossGameReward.upgradeToAndCall(address(cgrImpl), "");
        console.log("  CrossGameReward upgraded to:", address(cgrImpl));

        // ============================================
        // Step 3: Upgrade CrossPools (Individual)
        // ============================================
        console.log("\n[Step 3] Upgrading CrossPools individually...");

        // 업그레이드 후 새 함수 getPoolIdsByType 사용 가능
        uint[] memory crossPoolIds = crossGameReward.getPoolIdsByType(ICrossGameReward.PoolType.CrossPool);
        console.log("  Found CrossPools:", crossPoolIds.length);
        for (uint i = 0; i < crossPoolIds.length; i++) {
            CrossGameRewardPool pool = CrossGameRewardPool(address(crossGameReward.getPoolAddress(crossPoolIds[i])));
            console.log("  Pool ID:", crossPoolIds[i], "Address:", address(pool));
            pool.upgradeToAndCall(address(poolImpl), "");
            console.log("    -> upgraded");
        }
        console.log("  Total CrossPools upgraded:", crossPoolIds.length);

        // ============================================
        // Step 4: Register GamePool Implementation
        // ============================================
        console.log("\n[Step 4] Registering GamePool Implementation...");
        crossGameReward.setGamePoolImplementation(ICrossGameRewardPool(address(gamePoolImpl)));
        console.log("  GamePool Implementation set:", address(gamePoolImpl));

        vm.stopBroadcast();

        // ============================================
        // Summary
        // ============================================
        console.log("\n===========================================");
        console.log("  Deployment Summary");
        console.log("===========================================");
        console.log("CrossPool Implementation:        ", address(poolImpl));
        console.log("GamePool Implementation:       ", address(gamePoolImpl));
        console.log("CrossGameReward Implementation:", address(cgrImpl));
        console.log("CrossGameReward Proxy:         ", crossGameRewardAddr);
        console.log("CrossPools Upgraded:             ", crossPoolIds.length);
        console.log("\n=== Environment Variables for Future Use ===");
        console.log("POOL_IMPLEMENTATION=", address(poolImpl));
        console.log("GAME_POOL_IMPLEMENTATION=", address(gamePoolImpl));
        console.log("CROSS_GAME_REWARD_ROOT_IMPLEMENTATION=", address(cgrImpl));
    }
}
