// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/RewardPool.sol";
import "../src/StakingProtocol.sol";
import "../src/StakingRouter.sol";
import "../src/StakingViewer.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title TestScenario
 * @notice 테스트넷 시나리오 실행 스크립트
 * @dev 환경변수:
 *      - STAKING_ROUTER_ADDRESS: StakingRouter 주소
 *      - STAKING_PROTOCOL_ADDRESS: StakingProtocol 주소
 *      - PROJECT_ID: 프로젝트 ID
 *      - STAKE_AMOUNT: 스테이킹 금액 (wei 단위)
 *      - REWARD_TOKEN_ADDRESS: 보상 토큰 주소 (선택)
 *      - REWARD_AMOUNT: 보상 금액 (선택)
 * @dev 스크립트를 실행하는 계정으로 스테이킹 테스트를 실행합니다
 * @dev Router를 통해 Native CROSS로 스테이킹하므로 WCROSS 주소 불필요
 */
contract TestScenarioScript is Script {
    function run() external {
        address routerAddress = vm.envAddress("STAKING_ROUTER_ADDRESS");
        address protocolAddress = vm.envAddress("STAKING_PROTOCOL_ADDRESS");
        uint projectID = vm.envUint("PROJECT_ID");
        uint stakeAmount = vm.envUint("STAKE_AMOUNT");

        vm.startBroadcast();

        console.log("=== Test Scenario Execution ===");
        console.log("Tester:", msg.sender);
        console.log("Project ID:", projectID);

        StakingRouter router = StakingRouter(payable(routerAddress));
        StakingProtocol protocol = StakingProtocol(protocolAddress);
        StakingViewer viewer = new StakingViewer(protocolAddress);

        // 시나리오 1: Native CROSS로 스테이킹 (자동 wrapping)
        console.log("\n[Scenario 1] Staking with Native CROSS");
        console.log("Amount to stake:", stakeAmount);
        uint balanceBefore = msg.sender.balance;
        console.log("Native balance before:", balanceBefore);

        // Router.stake()는 payable이므로 Native CROSS 직접 전송
        router.stake{value: stakeAmount}(projectID);
        console.log("Native balance after:", msg.sender.balance);

        // 시나리오 2: 스테이킹 확인
        console.log("\n[Scenario 2] Checking Stake Position");

        (uint balance, uint points,) = viewer.getStakeInfo(projectID, msg.sender);
        console.log("Staked balance:", balance);
        console.log("Current points:", points);

        // 시나리오 3: 보상 예치 (선택)
        try vm.envAddress("REWARD_TOKEN_ADDRESS") returns (address rewardTokenAddress) {
            try vm.envUint("REWARD_AMOUNT") returns (uint rewardAmount) {
                console.log("\n[Scenario 4] Funding Rewards");

                IERC20 rewardToken = IERC20(rewardTokenAddress);

                console.log("Reward Token:", rewardTokenAddress);
                console.log("Reward Amount:", rewardAmount);
                console.log("Token balance:", rewardToken.balanceOf(msg.sender));

                // RewardPool 주소 조회
                (, address rewardPoolAddress,,,,,) = protocol.projects(projectID);

                // RewardPool에 직접 승인 및 예치
                rewardToken.approve(rewardPoolAddress, rewardAmount);

                uint currentSeason = viewer.getCurrentSeason(projectID);
                RewardPool rewardPool = RewardPool(rewardPoolAddress);
                rewardPool.fundSeason(currentSeason, rewardTokenAddress, rewardAmount);
                console.log("Rewards funded for Season", currentSeason);
            } catch {}
        } catch {
            console.log("\n[Scenario 4] Reward funding skipped");
        }

        // 시나리오 5: 포인트 확인
        console.log("\n[Scenario 5] Checking Points");
        uint userPoints = viewer.getUserPoints(projectID, msg.sender);
        uint stakingPower = viewer.getStakingPower(projectID, msg.sender);
        console.log("User Points:", userPoints);
        console.log("Staking Power:", stakingPower);

        console.log("\n=== Test Scenario Completed ===");

        vm.stopBroadcast();
    }
}
