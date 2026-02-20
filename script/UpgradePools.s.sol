// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardPool.sol";
import "../src/GamePool.sol";
import "../src/interfaces/ICrossGameReward.sol";
import "../src/interfaces/ICrossGameRewardPool.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "forge-std/Script.sol";

/**
 * @title UpgradePools
 * @notice CrossGameReward에 등록된 Pool을 타입별로 일괄 업그레이드하는 스크립트
 * @dev upgradePoolsByType()를 사용하여 V1 또는 GamePool 풀을 일괄 업그레이드합니다.
 *
 * 사용법:
 * # V1 Pool (CrossPool) 일괄 업그레이드
 * forge script script/UpgradePools.s.sol:UpgradePools \
 *   --rpc-url <RPC_URL> \
 *   --broadcast
 *
 * # GamePool 일괄 업그레이드
 * forge script script/UpgradePools.s.sol:UpgradePools \
 *   --sig "upgradeGamePool()" \
 *   --rpc-url <RPC_URL> \
 *   --broadcast
 *
 * 필수 환경변수:
 * - CROSS_GAME_REWARD: CrossGameReward 컨트랙트 주소
 * - NEW_POOL_IMPLEMENTATION: 새로운 Pool 구현체 주소
 *
 * 선택 환경변수:
 * - REINIT_DATA: reinitializer calldata (hex, 기본값: 없음)
 *
 * [주의: V1 Pool 최초 일괄 업그레이드]
 * 기존 배포된 V1 Pool의 _authorizeUpgrade가 onlyOwner 전용인 경우,
 * upgradePoolsByType()가 실패합니다. 이 경우:
 * 1. owner(defaultAdmin)가 각 V1 Pool에 대해 수동으로 1회 업그레이드:
 *    pool.upgradeToAndCall(newImpl, "")
 * 2. 새 구현체에서 _authorizeUpgrade가 owner + rewardRoot를 허용
 * 3. 이후부터는 이 스크립트로 일괄 업그레이드 가능
 */
contract UpgradePools is Script {
    /**
     * @notice V1 Pool (CrossPool) 일괄 업그레이드 - 기본 함수
     */
    function run() external {
        _upgrade(ICrossGameReward.PoolType.CrossPool, "V1 (CrossPool)");
    }

    /**
     * @notice GamePool 일괄 업그레이드
     * @dev --sig "upgradeGamePool()" 플래그와 함께 사용
     */
    function upgradeGamePool() external {
        _upgrade(ICrossGameReward.PoolType.GamePool, "GamePool");
    }

    function _upgrade(ICrossGameReward.PoolType poolType, string memory typeName) internal {
        address crossGameRewardAddr = vm.envAddress("CROSS_GAME_REWARD");
        address newImpl = vm.envAddress("NEW_POOL_IMPLEMENTATION");

        CrossGameReward crossGameReward = CrossGameReward(crossGameRewardAddr);

        // reinitializer calldata (선택적)
        bytes memory data;
        try vm.envBytes("REINIT_DATA") returns (bytes memory reinitData) {
            data = reinitData;
            console.log("Reinit data provided, length:", reinitData.length);
        } catch {
            data = "";
        }

        // 해당 타입의 풀 개수 확인
        uint[] memory poolIds = crossGameReward.getPoolIdsByType(poolType);

        console.log("\n=== Pool Upgrade Configuration ===");
        console.log("Type:", typeName);
        console.log("CrossGameReward:", crossGameRewardAddr);
        console.log("New Implementation:", newImpl);
        console.log("Pools to upgrade:", poolIds.length);

        for (uint i = 0; i < poolIds.length; i++) {
            ICrossGameRewardPool pool = crossGameReward.getPoolAddress(poolIds[i]);
            console.log("  Pool ID:", poolIds[i], "Address:", address(pool));
        }

        vm.startBroadcast();

        crossGameReward.upgradePoolsByType(poolType, newImpl, data);

        vm.stopBroadcast();

        console.log("\n=== Upgrade Complete ===");
        console.log("Pools upgraded:", poolIds.length);
        console.log("New Implementation:", newImpl);
    }
}
