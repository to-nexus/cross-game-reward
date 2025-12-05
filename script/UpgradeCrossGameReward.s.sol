// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "forge-std/Script.sol";

/**
 * @title UpgradeCrossGameReward
 * @notice CrossGameReward 프록시를 새로운 구현체로 업그레이드하는 스크립트
 * @dev UUPS 업그레이드 패턴을 사용하여 CrossGameReward의 로직을 업그레이드합니다.
 *
 * 사용법:
 * forge script script/UpgradeCrossGameReward.s.sol:UpgradeCrossGameReward \
 *   --rpc-url <RPC_URL> \
 *   --broadcast
 *
 * 필수 환경변수:
 * - CROSS_GAME_REWARD_PROXY: CrossGameReward 프록시 주소
 * - NEW_CROSS_GAME_REWARD_IMPLEMENTATION: 새로운 CrossGameReward 구현체 주소
 *
 * 주의사항:
 * - 이 스크립트는 CrossGameReward의 admin 권한을 가진 주소로 실행해야 합니다.
 * - 업그레이드 전 새로운 구현체가 올바르게 배포되었는지 확인하세요.
 * - 테스트넷에서 먼저 테스트하는 것을 권장합니다.
 */
contract UpgradeCrossGameReward is Script {
    function run() external {
        // 환경변수에서 설정 읽기
        address proxyAddress = vm.envAddress("CROSS_GAME_REWARD_PROXY");
        address newImplementation = vm.envAddress("NEW_CROSS_GAME_REWARD_IMPLEMENTATION");

        require(proxyAddress != address(0), "CROSS_GAME_REWARD_PROXY is required");
        require(newImplementation != address(0), "NEW_CROSS_GAME_REWARD_IMPLEMENTATION is required");

        // CrossGameReward 프록시 인스턴스 생성
        CrossGameReward crossGameReward = CrossGameReward(proxyAddress);

        console.log("\n=== CrossGameReward Upgrade Configuration ===");
        console.log("Proxy Address:", proxyAddress);
        console.log("New Implementation:", newImplementation);
        console.log("Deployer:", msg.sender);
        console.log("Chain ID:", block.chainid);

        vm.startBroadcast();

        // 1. 현재 상태 확인
        console.log("\n=== Step 1: Check Current State ===");
        address wcrossAddress = address(crossGameReward.wcross());
        address routerAddress = crossGameReward.router();
        uint totalPools = crossGameReward.getTotalPoolCount();

        console.log("Current WCROSS:", wcrossAddress);
        console.log("Current Router:", routerAddress);
        console.log("Total Pools:", totalPools);

        // 2. CrossGameReward 업그레이드
        console.log("\n=== Step 2: Upgrade CrossGameReward ===");
        console.log("Upgrading to new implementation...");

        // UUPS 업그레이드 (reinitialize 없이)
        crossGameReward.upgradeToAndCall(newImplementation, "");

        console.log("CrossGameReward upgraded successfully");

        // 3. 업그레이드 후 상태 확인
        console.log("\n=== Step 3: Verify After Upgrade ===");
        address wcrossAfter = address(crossGameReward.wcross());
        address routerAfter = crossGameReward.router();
        uint totalPoolsAfter = crossGameReward.getTotalPoolCount();

        console.log("WCROSS (after):", wcrossAfter);
        console.log("Router (after):", routerAfter);
        console.log("Total Pools (after):", totalPoolsAfter);

        // 상태 검증
        require(wcrossAddress == wcrossAfter, "WCROSS address changed unexpectedly");
        require(routerAddress == routerAfter, "Router address changed unexpectedly");
        require(totalPools == totalPoolsAfter, "Total pools changed unexpectedly");

        console.log("\nState verification passed!");

        vm.stopBroadcast();

        console.log("\n=== Upgrade Summary ===");
        console.log("Proxy Address:", proxyAddress);
        console.log("New Implementation:", newImplementation);
        console.log("WCROSS:", wcrossAddress);
        console.log("Router:", routerAddress);
        console.log("Total Pools:", totalPools);
        console.log("\n=== Upgrade Complete ===");
        console.log("All state preserved successfully");
    }
}

/*
 * ===================================
 * UpgradeCrossGameReward.env.example
 * ===================================
 *
 * # --------------------------------------------------
 * # CrossGameReward 프록시 주소 (필수)
 * # --------------------------------------------------
 * # 기존에 배포된 CrossGameReward 프록시 주소를 지정합니다.
 * # 이 주소의 admin 권한을 가진 계정으로 스크립트를 실행해야 합니다.
 *
 * CROSS_GAME_REWARD_PROXY=0x1234567890abcdef1234567890abcdef12345678
 *
 * # --------------------------------------------------
 * # 새로운 CrossGameReward 구현체 주소 (필수)
 * # --------------------------------------------------
 * # DeployImpl.s.sol의 deployCrossGameReward() 또는 run()으로
 * # 먼저 배포한 새로운 구현체 주소를 사용합니다.
 *
 * NEW_CROSS_GAME_REWARD_IMPLEMENTATION=0xabcdefabcdefabcdefabcdefabcdefabcdefabcd
 *
 * # --------------------------------------------------
 * # 사용 예시
 * # --------------------------------------------------
 *
 * # 1. 새로운 구현체 배포 (선택적 - 이미 배포된 경우 생략)
 * forge script script/DeployImpl.s.sol:DeployImpl \
 *   --sig "deployCrossGameReward()" \
 *   --rpc-url $RPC_URL \
 *   --broadcast
 *
 * # 2. 환경변수 설정
 * export CROSS_GAME_REWARD_PROXY=0x1234...
 * export NEW_CROSS_GAME_REWARD_IMPLEMENTATION=0xabcd...
 *
 * # 3. 업그레이드 실행
 * forge script script/UpgradeCrossGameReward.s.sol:UpgradeCrossGameReward \
 *   --rpc-url $RPC_URL \
 *   --private-key $PRIVATE_KEY \
 *   --broadcast
 *
 * # 또는 .env 파일 사용
 * forge script script/UpgradeCrossGameReward.s.sol:UpgradeCrossGameReward \
 *   --rpc-url $RPC_URL \
 *   --broadcast
 *
 * # --------------------------------------------------
 * # 주의사항
 * # --------------------------------------------------
 * # 1. deployer가 CrossGameReward의 DEFAULT_ADMIN_ROLE을 가지고 있어야 합니다.
 * # 2. 업그레이드 전 새로운 구현체를 테스트넷에서 충분히 테스트하세요.
 * # 3. 업그레이드는 되돌릴 수 없으므로 신중하게 진행하세요.
 * # 4. 업그레이드 후 모든 기능이 정상 작동하는지 확인하세요.
 * # 5. 기존 상태(WCROSS, Router, Pools)는 모두 보존됩니다.
 * # 6. 사용자 자금은 각 Pool 프록시에 안전하게 보관되어 있습니다.
 */
