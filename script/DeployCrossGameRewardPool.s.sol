// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Script.sol";

/**
 * @title CreatePool
 * @notice CrossGameReward에 새로운 풀을 추가하는 스크립트
 * @dev 사용법:
 * forge script script/DeployCrossGameRewardPool.s.sol:CreatePool \
 *   --rpc-url <RPC_URL> \
 *   --private-key <PRIVATE_KEY> \
 *   --broadcast
 *
 * 필수 환경변수:
 * - CROSS_GAME_REWARD_ADDRESS: CrossGameReward 컨트랙트 주소
 * - POOL_NAME: 생성할 풀 이름
 * - DEPOSIT_TOKEN_ADDRESS: 예치 토큰 주소
 * - MIN_DEPOSIT_AMOUNT: 최소 예치 금액 (wei 단위)
 *
 * 선택 환경변수:
 * - REWARD_TOKEN_ADDRESS: 보상 토큰 주소 (없으면 보상 토큰을 등록하지 않음)
 */
contract CreatePool is Script {
    function run() external {
        // 환경변수에서 설정 읽기
        address crossGameRewardAddress = vm.envAddress("CROSS_GAME_REWARD_ADDRESS");
        string memory poolName = vm.envString("POOL_NAME");
        address depositTokenAddress = vm.envAddress("DEPOSIT_TOKEN_ADDRESS");
        uint minDepositAmount = vm.envUint("MIN_DEPOSIT_AMOUNT");

        // CrossGameReward 컨트랙트 인스턴스 생성
        ICrossGameReward crossGameReward = ICrossGameReward(crossGameRewardAddress);

        console.log("\n=== Pool Creation Configuration ===");
        console.log("CrossGameReward Address:", crossGameRewardAddress);
        console.log("Pool Name:", poolName);
        console.log("Deposit Token:", depositTokenAddress);
        console.log("Min Deposit Amount:", minDepositAmount);
        console.log("Deployer:", msg.sender);

        vm.startBroadcast();

        // 1. Pool 생성
        (uint poolId, ICrossGameRewardPool pool) =
            crossGameReward.createPool(poolName, IERC20(depositTokenAddress), minDepositAmount);

        console.log("\n=== Pool Created ===");
        console.log("Pool ID:", poolId);
        console.log("Pool Address:", address(pool));

        // 2. 보상 토큰 추가 (환경변수가 있는 경우에만)
        try vm.envAddress("REWARD_TOKEN_ADDRESS") returns (address rewardTokenAddress) {
            console.log("\n=== Adding Reward Token ===");
            console.log("Reward Token Address:", rewardTokenAddress);

            crossGameReward.addRewardToken(poolId, IERC20(rewardTokenAddress));

            console.log("Reward token added successfully");
        } catch {
            console.log("\n=== No Reward Token ===");
            console.log("REWARD_TOKEN_ADDRESS not set, skipping reward token registration");
        }

        vm.stopBroadcast();

        console.log("\n=== Summary ===");
        console.log("Pool created successfully!");
        console.log("Pool ID:", poolId);
        console.log("Pool Address:", address(pool));
        console.log("Deposit Token:", depositTokenAddress);
        console.log("Min Deposit Amount:", minDepositAmount);
    }
}
