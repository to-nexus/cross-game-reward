// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardRouter.sol";
import "forge-std/Script.sol";

/**
 * @title DeployRouter
 * @notice CrossGameRewardRouter 배포 및 설정 스크립트
 * @dev 새로운 Router를 배포하고 CrossGameReward에 등록합니다.
 *      기존 Router를 교체하거나, 초기 배포 후 Router만 추가로 배포할 때 사용합니다.
 *
 * .env 파일에서 읽어오는 환경 변수:
 * - CROSS_GAME_REWARD (required, 필수): CrossGameReward 프록시 주소
 *
 * 사용법:
 * forge script script/DeployRouter.s.sol:DeployRouter \
 *   --rpc-url <RPC_URL> \
 *   --broadcast
 *
 * 주의사항:
 * - 이 스크립트는 CrossGameReward의 admin 또는 적절한 권한을 가진 주소로 실행해야 합니다.
 * - 기존 Router가 있다면 교체됩니다.
 */
contract DeployRouter is Script {
    function run() external {
        address deployer = msg.sender;

        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);

        // .env에서 CrossGameReward 주소 읽기 (필수)
        address crossGameRewardAddress = vm.envAddress("CROSS_GAME_REWARD");
        require(crossGameRewardAddress != address(0), "CROSS_GAME_REWARD address is required");

        console.log("\nCrossGameReward Address:", crossGameRewardAddress);

        CrossGameReward crossGameReward = CrossGameReward(crossGameRewardAddress);

        // 기존 Router 확인
        address existingRouter = crossGameReward.router();
        if (existingRouter != address(0)) {
            console.log("Existing Router:", existingRouter);
            console.log("Warning: Existing router will be replaced");
        } else {
            console.log("No existing router found");
        }

        vm.startBroadcast();

        // 1. Router 배포
        console.log("\n=== Step 1: Deploy Router ===");
        CrossGameRewardRouter router = new CrossGameRewardRouter(crossGameRewardAddress);
        console.log("Router deployed:", address(router));

        // 2. CrossGameReward에 Router 등록
        console.log("\n=== Step 2: Set Router in CrossGameReward ===");
        crossGameReward.setRouter(address(router));
        console.log("Router registered successfully");

        vm.stopBroadcast();

        // 배포 요약
        console.log("\n=== Deployment Summary ===");
        console.log("CrossGameReward:", crossGameRewardAddress);
        console.log("Router (New):", address(router));
        if (existingRouter != address(0)) console.log("Router (Old):", existingRouter);
        console.log("WCROSS:", address(crossGameReward.wcross()));

        // 다음 단계
        console.log("\n=== Next Steps ===");
        console.log("1. Users can now deposit via:");
        console.log("   - router.depositNative() for native CROSS");
        console.log("   - router.depositERC20() for ERC20 tokens");
        console.log("2. Users can withdraw via:");
        console.log("   - router.withdrawNative() for native CROSS");
        console.log("   - router.withdrawERC20() for ERC20 tokens");
        console.log("3. Users can claim rewards via:");
        console.log("   - router.claimRewards()");

        // 검증 정보
        console.log("\n=== Verification ===");
        console.log("Verify Router is correctly set:");
        console.log("crossGameReward.router() ==", address(router));
    }
}

/*
 * ===================================
 * DeployRouter.env.example
 * ===================================
 *
 * # --------------------------------------------------
 * # CrossGameReward 주소 (필수)
 * # --------------------------------------------------
 * # 기존에 배포된 CrossGameReward 프록시 주소를 지정합니다.
 * # 이 주소의 admin 권한을 가진 계정으로 스크립트를 실행해야 합니다.
 *
 * CROSS_GAME_REWARD=0x...  # CrossGameReward 프록시 주소 (필수)
 *
 * # --------------------------------------------------
 * # 사용 예시
 * # --------------------------------------------------
 *
 * # 1. 환경변수 설정
 * export CROSS_GAME_REWARD=0x1234567890abcdef1234567890abcdef12345678
 *
 * # 2. 스크립트 실행
 * forge script script/DeployRouter.s.sol:DeployRouter \
 *   --rpc-url $RPC_URL \
 *   --private-key $PRIVATE_KEY \
 *   --broadcast \
 *   --verify
 *
 * # 3. 또는 .env 파일 사용
 * forge script script/DeployRouter.s.sol:DeployRouter \
 *   --rpc-url $RPC_URL \
 *   --broadcast \
 *   --verify
 *
 * # --------------------------------------------------
 * # 주의사항
 * # --------------------------------------------------
 * # - deployer가 CrossGameReward의 admin이거나 setRouter 권한이 있어야 합니다.
 * # - 기존 Router가 있다면 교체됩니다. 사용자들이 사용 중인지 확인하세요.
 * # - Router 교체 시 사용자 자금은 안전합니다 (Pool에 저장되어 있음)
 */
