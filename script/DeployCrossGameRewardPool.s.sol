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
 *   --broadcast
 *
 * 필수 환경변수:
 * - CROSS_GAME_REWARD: CrossGameReward 컨트랙트 주소
 * - POOL_NAME: 생성할 풀 이름
 * - DEPOSIT_TOKEN: 예치 토큰 주소 (0x1=native token, 다른 주소=ERC20 token)
 * - MIN_DEPOSIT_AMOUNT: 최소 예치 금액 (wei 단위)
 *
 * 선택 환경변수:
 * - REWARD_TOKEN: 보상 토큰 주소 (없으면 보상 토큰을 등록하지 않음)
 */
contract CreatePool is Script {
    address public constant NATIVE_TOKEN_ADDRESS = address(0x1);

    function run() external {
        // 환경변수에서 설정 읽기
        address crossGameRewardAddress = vm.envAddress("CROSS_GAME_REWARD");
        string memory poolName = vm.envString("POOL_NAME");
        address depositTokenAddress = vm.envAddress("DEPOSIT_TOKEN");
        uint minDepositAmount = vm.envUint("MIN_DEPOSIT_AMOUNT");

        // CrossGameReward 컨트랙트 인스턴스 생성
        ICrossGameReward crossGameReward = ICrossGameReward(crossGameRewardAddress);

        console.log("\n=== Pool Creation Configuration ===");
        console.log("CrossGameReward Address:", crossGameRewardAddress);
        console.log("Pool Name:", poolName);
        console.log("Min Deposit Amount:", minDepositAmount);
        console.log("Deployer:", msg.sender);

        // Deposit Token 확인 (0x1 = native token, 그 외 = ERC20)
        IERC20 depositToken;
        if (depositTokenAddress == NATIVE_TOKEN_ADDRESS) {
            address wcrossAddress = address(crossGameReward.wcross());
            depositToken = IERC20(wcrossAddress);
            console.log("Deposit Token: WCROSS (Native)");
            console.log("Token Address:", wcrossAddress);
        } else {
            depositToken = IERC20(depositTokenAddress);
            console.log("Deposit Token: ERC20");
            console.log("Token Address:", depositTokenAddress);
        }

        vm.startBroadcast();

        // 1. Pool 생성
        (uint poolId, ICrossGameRewardPool pool) = crossGameReward.createPool(poolName, depositToken, minDepositAmount);

        console.log("\n=== Pool Created ===");
        console.log("Pool ID:", poolId);
        console.log("Pool Address:", address(pool));

        // 2. 보상 토큰 추가 (환경변수가 있는 경우에만)
        try vm.envAddress("REWARD_TOKEN") returns (address rewardTokenAddress) {
            console.log("\n=== Adding Reward Token ===");
            console.log("Reward Token Address:", rewardTokenAddress);

            crossGameReward.addRewardToken(poolId, IERC20(rewardTokenAddress));

            console.log("Reward token added successfully");
        } catch {
            console.log("\n=== No Reward Token ===");
            console.log("REWARD_TOKEN not set, skipping reward token registration");
        }

        vm.stopBroadcast();

        console.log("\n=== Summary ===");
        console.log("Pool created successfully!");
        console.log("Pool ID:", poolId);
        console.log("Pool Address:", address(pool));
        console.log("Deposit Token:", address(depositToken));
        console.log("Min Deposit Amount:", minDepositAmount);
    }
}

/*
 * ===================================
 * CreatePool.env.example
 * ===================================
 *
 * # --------------------------------------------------
 * # 필수 환경변수
 * # --------------------------------------------------
 *
 * CROSS_GAME_REWARD=0x1234567890abcdef1234567890abcdef12345678  # CrossGameReward 주소
 * POOL_NAME="My Pool"                                           # Pool 이름
 * MIN_DEPOSIT_AMOUNT=1000000000000000000                        # 최소 예치 금액 (1 ether)
 *
 * # --------------------------------------------------
 * # Deposit Token 설정 (필수)
 * # --------------------------------------------------
 * # 0x1 = Native Token (WCROSS 사용)
 * # 다른 주소 = ERC20 토큰 주소
 *
 * # Native Token Pool 생성
 * DEPOSIT_TOKEN=0x1
 *
 * # 또는 ERC20 Token Pool 생성
 * # DEPOSIT_TOKEN=0xabcdefabcdefabcdefabcdefabcdefabcdefabcd
 *
 * # --------------------------------------------------
 * # 선택 환경변수
 * # --------------------------------------------------
 * # Reward Token을 설정하면 Pool 생성 후 자동으로 등록됩니다.
 *
 * REWARD_TOKEN=0x9999999999999999999999999999999999999999
 *
 * # --------------------------------------------------
 * # 사용 예시
 * # --------------------------------------------------
 *
 * # 1. Native Token Pool 생성
 * forge script script/DeployCrossGameRewardPool.s.sol:CreatePool \
 *   --rpc-url $RPC_URL \
 *   --broadcast
 *
 * # 2. ERC20 Token Pool 생성
 * export DEPOSIT_TOKEN=0xabcd...
 * forge script script/DeployCrossGameRewardPool.s.sol:CreatePool \
 *   --rpc-url $RPC_URL \
 *   --broadcast
 */
