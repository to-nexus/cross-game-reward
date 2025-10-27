// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/RewardPool.sol";

import "../src/StakingProtocol.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title SweepRewardPool
 * @notice RewardPool에서 잘못 전송되거나 남은 토큰을 회수하는 스크립트
 * @dev 환경변수:
 *      - STAKING_PROTOCOL_ADDRESS: StakingProtocol 주소
 *      - REWARD_POOL_ADDRESS: RewardPool 주소
 *      - TOKEN_ADDRESS: 회수할 토큰 주소
 *      - SWEEP_AMOUNT: 회수할 수량 (wei 단위)
 * @dev 명령줄 플래그:
 *      - --sig "run(address)" <수신자주소>
 * @dev 이 스크립트는 Protocol Admin만 실행할 수 있습니다
 * @dev 주의: sweep은 되돌릴 수 없습니다. 신중하게 사용하세요!
 * @dev 사용 예시:
 *      forge script script/SweepRewardPool.s.sol:SweepRewardPoolScript \
 *          --sig "run(address)" 0x수신자주소 \
 *          --broadcast --rpc-url $RPC_URL
 */
contract SweepRewardPoolScript is Script {
    function run(address sweepTo) external {
        address protocolAddress = vm.envAddress("STAKING_PROTOCOL_ADDRESS");
        address rewardPoolAddress = vm.envAddress("REWARD_POOL_ADDRESS");
        address tokenAddress = vm.envAddress("TOKEN_ADDRESS");
        uint sweepAmount = vm.envUint("SWEEP_AMOUNT");

        vm.startBroadcast();

        console.log("=== RewardPool Sweep ===");
        console.log("Caller:", msg.sender);
        console.log("Protocol:", protocolAddress);
        console.log("RewardPool:", rewardPoolAddress);
        console.log("Token:", tokenAddress);
        console.log("Sweep To:", sweepTo);
        console.log("Amount:", sweepAmount);

        // 토큰 잔액 확인
        IERC20 token = IERC20(tokenAddress);
        uint poolBalance = token.balanceOf(rewardPoolAddress);
        console.log("\nRewardPool Token Balance:", poolBalance);

        require(poolBalance >= sweepAmount, "Insufficient token balance in RewardPool");
        require(sweepAmount > 0, "Sweep amount must be greater than 0");

        // 수신자 잔액 확인 (before)
        uint recipientBalanceBefore = token.balanceOf(sweepTo);
        console.log("Recipient Balance (before):", recipientBalanceBefore);

        // Protocol을 통해 sweep 호출
        // sweep(address token, address to, uint amount)
        StakingProtocol protocol = StakingProtocol(protocolAddress);

        console.log("\nCalling sweep...");
        // RewardPool의 sweep은 Protocol을 통해서만 호출 가능
        // Protocol에서 직접 RewardPool.sweep()을 호출해야 하므로
        // 여기서는 low-level call 사용
        (bool success,) = address(protocol).call(
            abi.encodeWithSignature(
                "sweepRewardPool(address,address,address,uint256)",
                rewardPoolAddress,
                tokenAddress,
                sweepTo,
                sweepAmount
            )
        );

        // Protocol에 sweepRewardPool 함수가 없다면 RewardPool에 직접 호출
        if (!success) {
            console.log("Protocol sweepRewardPool not found, calling RewardPool directly...");
            RewardPool rewardPool = RewardPool(rewardPoolAddress);
            rewardPool.sweep(tokenAddress, sweepTo, sweepAmount);
        }

        // 수신자 잔액 확인 (after)
        uint recipientBalanceAfter = token.balanceOf(sweepTo);
        uint poolBalanceAfter = token.balanceOf(rewardPoolAddress);

        console.log("\n=== Sweep Complete ===");
        console.log("RewardPool Balance (after):", poolBalanceAfter);
        console.log("Recipient Balance (after):", recipientBalanceAfter);
        console.log("Amount Swept:", recipientBalanceAfter - recipientBalanceBefore);

        require(recipientBalanceAfter >= recipientBalanceBefore + sweepAmount, "Sweep failed: incorrect amount");

        vm.stopBroadcast();

        console.log("\n=== Success! ===");
        console.log("Swept", sweepAmount, "tokens to", sweepTo);
    }
}

/**
 * @title SweepRewardPoolByProjectID
 * @notice ProjectID를 사용하여 RewardPool에서 토큰을 회수하는 스크립트
 * @dev 환경변수:
 *      - STAKING_PROTOCOL_ADDRESS: StakingProtocol 주소
 *      - PROJECT_ID: 프로젝트 ID
 *      - TOKEN_ADDRESS: 회수할 토큰 주소
 *      - SWEEP_AMOUNT: 회수할 수량 (wei 단위)
 * @dev 명령줄 플래그:
 *      - --sig "run(address)" <수신자주소>
 * @dev 이 스크립트는 Protocol Admin만 실행할 수 있습니다
 * @dev 사용 예시:
 *      forge script script/SweepRewardPool.s.sol:SweepRewardPoolByProjectIDScript \
 *          --sig "run(address)" 0x수신자주소 \
 *          --broadcast --rpc-url $RPC_URL
 */
contract SweepRewardPoolByProjectIDScript is Script {
    function run(address sweepTo) external {
        address protocolAddress = vm.envAddress("STAKING_PROTOCOL_ADDRESS");
        uint projectID = vm.envUint("PROJECT_ID");
        address tokenAddress = vm.envAddress("TOKEN_ADDRESS");
        uint sweepAmount = vm.envUint("SWEEP_AMOUNT");

        vm.startBroadcast();

        console.log("=== RewardPool Sweep (by Project ID) ===");
        console.log("Caller:", msg.sender);
        console.log("Protocol:", protocolAddress);
        console.log("Project ID:", projectID);
        console.log("Token:", tokenAddress);
        console.log("Sweep To:", sweepTo);
        console.log("Amount:", sweepAmount);

        StakingProtocol protocol = StakingProtocol(protocolAddress);

        // 프로젝트 정보 조회
        (, address rewardPoolAddress,,,,,) = protocol.projects(projectID);
        console.log("RewardPool:", rewardPoolAddress);
        require(rewardPoolAddress != address(0), "Invalid project ID");

        // 토큰 잔액 확인
        IERC20 token = IERC20(tokenAddress);
        uint poolBalance = token.balanceOf(rewardPoolAddress);
        console.log("\nRewardPool Token Balance:", poolBalance);

        require(poolBalance >= sweepAmount, "Insufficient token balance in RewardPool");
        require(sweepAmount > 0, "Sweep amount must be greater than 0");

        // 수신자 잔액 확인 (before)
        uint recipientBalanceBefore = token.balanceOf(sweepTo);
        console.log("Recipient Balance (before):", recipientBalanceBefore);

        // Sweep 실행 (Protocol을 통해)
        console.log("\nCalling sweepRewardPool through Protocol...");
        protocol.sweepRewardPool(projectID, tokenAddress, sweepTo, sweepAmount);

        // 수신자 잔액 확인 (after)
        uint recipientBalanceAfter = token.balanceOf(sweepTo);
        uint poolBalanceAfter = token.balanceOf(rewardPoolAddress);

        console.log("\n=== Sweep Complete ===");
        console.log("RewardPool Balance (after):", poolBalanceAfter);
        console.log("Recipient Balance (after):", recipientBalanceAfter);
        console.log("Amount Swept:", recipientBalanceAfter - recipientBalanceBefore);

        require(recipientBalanceAfter >= recipientBalanceBefore + sweepAmount, "Sweep failed: incorrect amount");

        vm.stopBroadcast();

        console.log("\n=== Success! ===");
    }
}
