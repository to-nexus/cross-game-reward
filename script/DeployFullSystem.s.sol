// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/CrossStaking.sol";
import "../src/CrossStakingPool.sol";
import "../src/CrossStakingRouter.sol";
import "../src/WCROSS.sol";
import "../src/interfaces/ICrossStakingPool.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Script.sol";

/**
 * @title DeployFullSystem
 * @notice 전체 시스템 배포 스크립트
 * @dev WCROSS, CrossStaking (UUPS), CrossStakingRouter 배포
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

        // 1. CrossStakingPool Implementation 배포
        CrossStakingPool poolImplementation = new CrossStakingPool();
        console.log("\n1. Pool Implementation deployed:", address(poolImplementation));

        // 2. CrossStaking Implementation & Proxy 배포 (WCROSS를 생성함)
        CrossStaking crossStakingImpl = new CrossStaking();
        bytes memory initData = abi.encodeWithSelector(
            CrossStaking.initialize.selector, ICrossStakingPool(address(poolImplementation)), deployer, INITIAL_DELAY
        );
        ERC1967Proxy crossStakingProxy = new ERC1967Proxy(address(crossStakingImpl), initData);
        CrossStaking crossStaking = CrossStaking(address(crossStakingProxy));
        console.log("2. CrossStaking Proxy deployed:", address(crossStaking));
        console.log("   CrossStaking Implementation:", address(crossStakingImpl));

        // 3. CrossStakingRouter 배포
        CrossStakingRouter router = new CrossStakingRouter(address(crossStaking));
        console.log("3. Router deployed:", address(router));

        // 4. WCROSS 주소 확인 (CrossStaking이 생성함)
        IWCROSS wcross = crossStaking.wcross();
        console.log("4. WCROSS deployed by CrossStaking:", address(wcross));

        // 5. Router 등록
        crossStaking.setRouter(address(router));
        console.log("5. Router set in CrossStaking");

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("WCROSS:", address(wcross));
        console.log("CrossStaking (Proxy):", address(crossStaking));
        console.log("CrossStaking (Impl):", address(crossStakingImpl));
        console.log("Pool Implementation:", address(poolImplementation));
        console.log("Router:", address(router));
        console.log("Admin:", deployer);
        console.log("\n=== Next Steps ===");
        console.log("1. Create pools via CrossStaking.createPool()");
        console.log("2. Add reward tokens via CrossStaking.addRewardToken()");
        console.log("3. Users stake via Router.stakeNative() or Router.stakeERC20()");
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
        CrossStakingPool poolImplementation = new CrossStakingPool();

        // 2. CrossStaking 배포 (UUPS - WCROSS를 생성함)
        CrossStaking crossStakingImpl = new CrossStaking();
        bytes memory initData = abi.encodeWithSelector(
            CrossStaking.initialize.selector, ICrossStakingPool(address(poolImplementation)), deployer, INITIAL_DELAY
        );
        ERC1967Proxy crossStakingProxy = new ERC1967Proxy(address(crossStakingImpl), initData);
        CrossStaking crossStaking = CrossStaking(address(crossStakingProxy));

        // 3. Router 배포
        CrossStakingRouter router = new CrossStakingRouter(address(crossStaking));

        // 4. WCROSS 주소 확인
        IWCROSS wcross = crossStaking.wcross();

        // 5. Router 등록
        crossStaking.setRouter(address(router));

        // 6. Native CROSS 풀 생성
        (uint nativePoolId, ICrossStakingPool nativePool) = crossStaking.createPool(IERC20(address(wcross)), 1 ether);

        console.log("\n=== Full Deployment Summary ===");
        console.log("WCROSS:", address(wcross));
        console.log("CrossStaking:", address(crossStaking));
        console.log("Router:", address(router));
        console.log("Native Pool ID:", nativePoolId);
        console.log("Native Pool Address:", address(nativePool));

        vm.stopBroadcast();
    }
}
