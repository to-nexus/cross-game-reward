// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/StakingProtocol.sol";

import {Script, console} from "forge-std/Script.sol";

/**
 * @title CreateProject
 * @notice 새로운 스테이킹 프로젝트 생성 스크립트
 * @dev 환경변수:
 *      - STAKING_PROTOCOL_ADDRESS: StakingProtocol 주소
 *      - PROJECT_NAME: 프로젝트 이름
 *      - SEASON_DURATION: 시즌 길이 (초 단위, 0이면 기본값)
 *      - FIRST_SEASON_START_TIME: 첫 시즌 시작 타임스탬프
 *      - POOL_END_TIME: 풀 종료 타임스탬프 (0이면 무한)
 *      - PROJECT_ADMIN: (선택) 프로젝트 관리자 주소 (없으면 생성자)
 * @dev 스크립트를 실행하는 계정이 프로젝트를 생성합니다
 */
contract CreateProjectScript is Script {
    function run() external {
        address protocolAddress = vm.envAddress("STAKING_PROTOCOL_ADDRESS");
        string memory projectName = vm.envString("PROJECT_NAME");
        uint seasonDuration = vm.envUint("SEASON_DURATION");
        uint firstSeasonStartTime = vm.envUint("FIRST_SEASON_START_TIME");
        uint poolEndTime = vm.envUint("POOL_END_TIME");

        // PROJECT_ADMIN이 설정되어 있으면 사용, 없으면 msg.sender 사용
        address projectAdmin;
        try vm.envAddress("PROJECT_ADMIN") returns (address admin) {
            projectAdmin = admin;
        } catch {
            projectAdmin = msg.sender;
        }

        vm.startBroadcast();

        console.log("=== Creating New Project ===");
        console.log("Creator:", msg.sender);
        console.log("Protocol:", protocolAddress);
        console.log("Project Name:", projectName);
        console.log("Season Duration (seconds):", seasonDuration);
        console.log("First Season Start Time:", firstSeasonStartTime);
        console.log("Pool End Time:", poolEndTime);
        console.log("Project Admin:", projectAdmin);

        StakingProtocol protocol = StakingProtocol(protocolAddress);

        (uint projectID, address stakingPool, address rewardPool) =
            protocol.createProject(projectName, seasonDuration, firstSeasonStartTime, poolEndTime, projectAdmin, 0);

        console.log("\n=== Project Created ===");
        console.log("Project ID:", projectID);
        console.log("StakingPool:", stakingPool);
        console.log("RewardPool:", rewardPool);

        console.log("\n=== Save these addresses ===");
        console.log("PROJECT_ID=", projectID);
        console.log("STAKING_POOL_ADDRESS=", stakingPool);
        console.log("REWARD_POOL_ADDRESS=", rewardPool);

        vm.stopBroadcast();
    }
}
