// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardPool.sol";
import "../src/CrossGameRewardRouter.sol";
import "../src/WCROSS.sol";
import "../src/interfaces/ICrossGameRewardPool.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Script.sol";

/**
 * @title DeployFullSystem
 * @notice 전체 시스템 배포 스크립트
 * @dev WCROSS, CrossGameReward (UUPS), CrossGameRewardRouter 배포
 *
 * 사용법:
 * forge script script/DeployFullSystem.s.sol:DeployFullSystem \
 *   --rpc-url <RPC_URL> \
 *   --private-key <PRIVATE_KEY> \
 *   --broadcast
 */
contract DeployFullSystem is Script {
    uint48 public constant INITIAL_DELAY = 2 days;

    function run() external {
        address deployer = msg.sender;

        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);

        vm.startBroadcast();

        // 1. CrossGameRewardPool Implementation 배포
        CrossGameRewardPool poolImplementation = new CrossGameRewardPool();
        console.log("\n1. Pool Implementation deployed:", address(poolImplementation));

        // 2. CrossGameReward Implementation & Proxy 배포 (WCROSS를 생성함)
        CrossGameReward crossGameRewardImpl = new CrossGameReward();
        bytes memory initData = abi.encodeWithSelector(
            CrossGameReward.initialize.selector,
            ICrossGameRewardPool(address(poolImplementation)),
            deployer,
            INITIAL_DELAY
        );
        ERC1967Proxy crossGameRewardProxy = new ERC1967Proxy(address(crossGameRewardImpl), initData);
        CrossGameReward crossGameReward = CrossGameReward(address(crossGameRewardProxy));
        console.log("2. CrossGameReward Proxy deployed:", address(crossGameReward));
        console.log("   CrossGameReward Implementation:", address(crossGameRewardImpl));

        // 3. CrossGameRewardRouter 배포
        CrossGameRewardRouter router = new CrossGameRewardRouter(address(crossGameReward));
        console.log("3. Router deployed:", address(router));

        // 4. WCROSS 주소 확인 (CrossGameReward가 생성함)
        IWCROSS wcross = crossGameReward.wcross();
        console.log("4. WCROSS deployed by CrossGameReward:", address(wcross));

        // 5. Router 등록
        crossGameReward.setRouter(address(router));
        console.log("5. Router set in CrossGameReward");

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("WCROSS:", address(wcross));
        console.log("CrossGameReward (Proxy):", address(crossGameReward));
        console.log("CrossGameReward (Impl):", address(crossGameRewardImpl));
        console.log("Pool Implementation:", address(poolImplementation));
        console.log("Router:", address(router));
        console.log("Admin:", deployer);
        console.log("\n=== Next Steps ===");
        console.log("1. Create pools via CrossGameReward.createPool()");
        console.log("2. Add reward tokens via CrossGameReward.addRewardToken()");
        console.log("3. Users deposit via Router.depositNative() or Router.depositERC20()");
    }
}

/**
 * @title DeployWithPools
 * @notice 풀까지 포함한 전체 배포
 */
contract DeployWithPools is Script {
    uint48 public constant INITIAL_DELAY = 2 days;
    uint48 public constant POOL_DELAY = 1 days;

    function run() external {
        address deployer = msg.sender;

        vm.startBroadcast();

        // 1. Pool Implementation 배포
        CrossGameRewardPool poolImplementation = new CrossGameRewardPool();

        // 2. CrossGameReward 배포 (UUPS - WCROSS를 생성함)
        CrossGameReward crossGameRewardImpl = new CrossGameReward();
        bytes memory initData = abi.encodeWithSelector(
            CrossGameReward.initialize.selector,
            ICrossGameRewardPool(address(poolImplementation)),
            deployer,
            INITIAL_DELAY
        );
        ERC1967Proxy crossGameRewardProxy = new ERC1967Proxy(address(crossGameRewardImpl), initData);
        CrossGameReward crossGameReward = CrossGameReward(address(crossGameRewardProxy));

        // 3. Router 배포
        CrossGameRewardRouter router = new CrossGameRewardRouter(address(crossGameReward));

        // 4. WCROSS 주소 확인
        IWCROSS wcross = crossGameReward.wcross();

        // 5. Router 등록
        crossGameReward.setRouter(address(router));

        // 6. Native CROSS 풀 생성
        (uint nativePoolId, ICrossGameRewardPool nativePool) =
            crossGameReward.createPool("Native CROSS Pool", IERC20(address(wcross)), 1 ether);

        console.log("\n=== Full Deployment Summary ===");
        console.log("WCROSS:", address(wcross));
        console.log("CrossGameReward:", address(crossGameReward));
        console.log("Router:", address(router));
        console.log("Native Pool ID:", nativePoolId);
        console.log("Native Pool Address:", address(nativePool));

        vm.stopBroadcast();
    }
}
