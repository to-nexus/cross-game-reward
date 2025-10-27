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
 *      - SEASON_BLOCKS: 시즌 길이 (블록 수, 0이면 기본값)
 *      - FIRST_SEASON_START_BLOCK: 첫 시즌 시작 블록
 *      - POOL_END_BLOCK: 풀 종료 블록 (0이면 무한)
 *      - PROJECT_ADMIN: (선택) 프로젝트 관리자 주소 (없으면 생성자)
 * @dev 스크립트를 실행하는 계정이 프로젝트를 생성합니다
 */
contract CreateProjectScript is Script {
    function run() external {
        address protocolAddress = vm.envAddress("STAKING_PROTOCOL_ADDRESS");
        string memory projectName = vm.envString("PROJECT_NAME");
        uint seasonBlocks = vm.envUint("SEASON_BLOCKS");
        uint firstSeasonStartBlock = vm.envUint("FIRST_SEASON_START_BLOCK");
        uint poolEndBlock = vm.envUint("POOL_END_BLOCK");

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
        console.log("Season Blocks:", seasonBlocks);
        console.log("First Season Start:", firstSeasonStartBlock);
        console.log("Pool End Block:", poolEndBlock);
        console.log("Project Admin:", projectAdmin);

        StakingProtocol protocol = StakingProtocol(protocolAddress);

        (uint projectID, address stakingPool, address rewardPool) =
            protocol.createProject(projectName, seasonBlocks, firstSeasonStartBlock, poolEndBlock, projectAdmin, 0);

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
