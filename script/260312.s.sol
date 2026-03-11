// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardPool.sol";
import "../src/interfaces/ICrossGameReward.sol";
import "../src/interfaces/ICrossGameRewardPool.sol";
import "forge-std/Script.sol";

/**
 * @title Deploy260312
 * @notice 2026.03.12 배포 스크립트 - CrossPool 선택적 업그레이드
 * @dev ERC1967 프록시의 구현체 슬롯을 조회하여, 타겟 구현체와 다른 풀만
 *      선택적으로 업그레이드합니다. owner(defaultAdmin)가 직접 각 풀의
 *      upgradeToAndCall을 호출하므로 Legacy V1 풀도 업그레이드 가능합니다.
 *
 *      1. 모든 CrossPool 프록시의 현재 구현체 주소 조회 (ERC1967 슬롯)
 *      2. 타겟 구현체와 다른 풀만 upgradeToAndCall 호출
 *      3. setPoolImplementation으로 팩토리 레퍼런스 갱신
 *
 * 필수 환경변수:
 * - CROSS_GAME_REWARD: CrossGameReward Proxy 주소
 * - POOL_IMPLEMENTATION: 타겟 CrossPool 구현체 주소
 *
 * 사용법:
 * forge script script/260312.s.sol:Deploy260312 \
 *   --rpc-url <RPC_URL> \
 *   --broadcast
 */
contract Deploy260312 is Script {
    /// @dev ERC1967 implementation storage slot
    bytes32 internal constant _IMPLEMENTATION_SLOT = 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;

    function _getImplementation(address proxy) internal view returns (address) {
        bytes32 slot = vm.load(proxy, _IMPLEMENTATION_SLOT);
        return address(uint160(uint256(slot)));
    }

    function run() external {
        address crossGameRewardAddr = vm.envAddress("CROSS_GAME_REWARD");
        address targetImpl = vm.envAddress("POOL_IMPLEMENTATION");

        CrossGameReward crossGameReward = CrossGameReward(crossGameRewardAddr);

        console.log("===========================================");
        console.log("  2026.03.12 CrossPool Selective Upgrade");
        console.log("===========================================");
        console.log("Deployer:", msg.sender);
        console.log("Chain ID:", block.chainid);
        console.log("CrossGameReward:", crossGameRewardAddr);
        console.log("Target Implementation:", targetImpl);

        // ============================================
        // Step 1: 모든 CrossPool 프록시의 구현체 조회
        // ============================================
        console.log("\n[Step 1] Scanning CrossPool implementations...");

        uint[] memory poolIds = crossGameReward.getPoolIdsByType(ICrossGameReward.PoolType.CrossPool);
        console.log("  Total CrossPools found:", poolIds.length);

        uint upgradeCount = 0;
        uint skipCount = 0;

        // 업그레이드 대상 식별 (CrossPool 이중 검증 포함)
        bool[] memory needsUpgrade = new bool[](poolIds.length);
        for (uint i = 0; i < poolIds.length; i++) {
            ICrossGameReward.PoolType poolType = crossGameReward.getPoolType(poolIds[i]);
            require(poolType == ICrossGameReward.PoolType.CrossPool, "SAFETY: non-CrossPool detected");

            address poolAddr = address(crossGameReward.getPoolAddress(poolIds[i]));
            address currentImpl = _getImplementation(poolAddr);
            bool outdated = currentImpl != targetImpl;
            needsUpgrade[i] = outdated;

            if (outdated) {
                console.log("  [OUTDATED] Pool ID:", poolIds[i]);
                console.log("    Address:", poolAddr);
                console.log("    Current Impl:", currentImpl);
                upgradeCount++;
            } else {
                console.log("  [OK]      Pool ID:", poolIds[i]);
                console.log("    Address:", poolAddr);
                skipCount++;
            }
        }

        console.log("\n  Pools to upgrade:", upgradeCount);
        console.log("  Pools to skip:", skipCount);

        // ============================================
        // Step 2: 다른 구현체를 가진 풀만 업그레이드
        // ============================================
        console.log("\n[Step 2] Upgrading outdated CrossPools...");

        vm.startBroadcast();

        for (uint i = 0; i < poolIds.length; i++) {
            if (!needsUpgrade[i]) continue;

            address poolAddr = address(crossGameReward.getPoolAddress(poolIds[i]));
            CrossGameRewardPool(poolAddr).upgradeToAndCall(targetImpl, "");
            console.log("  Upgraded Pool ID:", poolIds[i]);
        }

        // ============================================
        // Step 3: setPoolImplementation으로 팩토리 레퍼런스 갱신
        // ============================================
        console.log("\n[Step 3] Updating pool implementation reference...");

        address currentFactoryImpl = address(crossGameReward.poolImplementation());
        if (currentFactoryImpl != targetImpl) {
            crossGameReward.setPoolImplementation(ICrossGameRewardPool(targetImpl));
            console.log("  poolImplementation updated:");
            console.log("    Before:", currentFactoryImpl);
            console.log("    After: ", targetImpl);
        } else {
            console.log("  poolImplementation already up to date:", currentFactoryImpl);
        }

        vm.stopBroadcast();

        // ============================================
        // Summary
        // ============================================
        console.log("\n===========================================");
        console.log("  Upgrade Summary");
        console.log("===========================================");
        console.log("Total CrossPools:  ", poolIds.length);
        console.log("Upgraded:          ", upgradeCount);
        console.log("Skipped (up-to-date):", skipCount);
        console.log("Target Implementation:", targetImpl);
    }
}
