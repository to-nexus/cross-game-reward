// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardRouter.sol";
import "forge-std/Script.sol";

/**
 * @title DeployRouter
 * @notice CrossGameRewardRouter 배포 스크립트
 * @dev Router 배포와 CrossGameReward 등록을 지원합니다.
 *
 * .env 파일에서 읽어오는 환경 변수:
 * - CROSS_GAME_REWARD (required, 필수): CrossGameReward 프록시 주소
 *
 * 사용법:
 * 1. Router 배포 + 자동 등록 (기본):
 *    forge script script/DeployRouter.s.sol:DeployRouter \
 *      --rpc-url <RPC_URL> \
 *      --broadcast
 *
 * 2. Router만 배포 (수동 등록):
 *    forge script script/DeployRouter.s.sol:DeployRouter \
 *      --sig "deployRouterOnly()" \
 *      --rpc-url <RPC_URL> \
 *      --broadcast
 */
contract DeployRouter is Script {
    /**
     * @notice 기본 실행 함수 - Router 배포 + CrossGameReward에 자동 등록
     * @dev CrossGameReward.setRouter()를 호출하여 Router를 등록합니다.
     */
    function run() external {
        address deployer = msg.sender;

        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);

        // CrossGameReward 주소 읽기 (필수)
        address crossGameRewardAddr = vm.envAddress("CROSS_GAME_REWARD");
        require(crossGameRewardAddr != address(0), "CROSS_GAME_REWARD is required");

        console.log("\nCrossGameReward Address:", crossGameRewardAddr);

        vm.startBroadcast();

        // 1. CrossGameRewardRouter 배포
        CrossGameRewardRouter router = new CrossGameRewardRouter(crossGameRewardAddr);
        console.log("1. Router deployed:", address(router));

        // 2. CrossGameReward에 Router 등록
        CrossGameReward crossGameReward = CrossGameReward(crossGameRewardAddr);
        crossGameReward.setRouter(address(router));
        console.log("2. Router registered in CrossGameReward");

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("Router:", address(router));
        console.log("CrossGameReward:", crossGameRewardAddr);
        console.log("WCROSS:", address(crossGameReward.wcross()));
        console.log("\n=== Next Steps ===");
        console.log("1. Users can now deposit via Router.depositNative() or Router.depositERC20()");
        console.log("2. Users can withdraw via Router.withdrawNative() or Router.withdrawERC20()");
        console.log("3. Users can claim rewards via Router.claimRewards()");
    }

    /**
     * @notice Router만 배포 (등록 없음)
     * @dev Router만 배포하고 CrossGameReward 등록은 수동으로 처리합니다.
     *      --sig "deployRouterOnly()" 플래그와 함께 사용
     */
    function deployRouterOnly() external {
        address deployer = msg.sender;

        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);

        // CrossGameReward 주소 읽기 (필수)
        address crossGameRewardAddr = vm.envAddress("CROSS_GAME_REWARD");
        require(crossGameRewardAddr != address(0), "CROSS_GAME_REWARD is required");

        console.log("\nCrossGameReward Address:", crossGameRewardAddr);

        vm.startBroadcast();

        // CrossGameRewardRouter 배포
        CrossGameRewardRouter router = new CrossGameRewardRouter(crossGameRewardAddr);
        console.log("\nRouter deployed:", address(router));

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("Router:", address(router));
        console.log("CrossGameReward:", crossGameRewardAddr);
        console.log("\n=== Next Steps ===");
        console.log("Register Router manually by calling:");
        console.log("CrossGameReward.setRouter(", address(router), ")");
        console.log("\nAfter registration:");
        console.log("1. Users can deposit via Router.depositNative() or Router.depositERC20()");
        console.log("2. Users can withdraw via Router.withdrawNative() or Router.withdrawERC20()");
        console.log("3. Users can claim rewards via Router.claimRewards()");
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
 * # DeployFullSystem.s.sol로 배포한 CrossGameReward 프록시 주소를 사용합니다.
 * # 필수 환경 변수이므로, 설정하지 않으면 스크립트가 실패합니다.
 *
 * CROSS_GAME_REWARD=0x...  # CrossGameReward Proxy 주소 (필수)
 *
 * # --------------------------------------------------
 * # 시나리오 1: Router 배포 + 자동 등록 (기본)
 * # --------------------------------------------------
 * # 기본 run() 함수를 사용하여 Router를 배포하고
 * # CrossGameReward에 자동으로 등록합니다.
 * #
 * # 사용법:
 * # forge script script/DeployRouter.s.sol:DeployRouter \
 * #   --rpc-url <RPC_URL> \
 * #   --broadcast
 *
 * CROSS_GAME_REWARD=0x1234...
 *
 * # --------------------------------------------------
 * # 시나리오 2: Router만 배포 (수동 등록)
 * # --------------------------------------------------
 * # deployRouterOnly() 함수를 사용하여 Router만 배포하고
 * # 등록은 나중에 수동으로 처리합니다.
 * #
 * # 사용법:
 * # forge script script/DeployRouter.s.sol:DeployRouter \
 * #   --sig "deployRouterOnly()" \
 * #   --rpc-url <RPC_URL> \
 * #   --broadcast
 * #
 * # 배포 후 수동으로 다음을 실행:
 * # cast send $CROSS_GAME_REWARD \
 * #   "setRouter(address)" $ROUTER_ADDRESS \
 * #   --rpc-url <RPC_URL> \
 * #   --private-key <PRIVATE_KEY>
 *
 * CROSS_GAME_REWARD=0x1234...
 */
