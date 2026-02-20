// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/interfaces/IGamePool.sol";
import "../src/interfaces/IWCROSS.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Script.sol";

/**
 * @title CreatePool
 * @notice CrossGameReward에 새로운 풀을 추가하는 스크립트
 * @dev V1 (CrossPool) 또는 GamePool 풀을 생성할 수 있습니다.
 *
 * 사용법:
 * # V1 Pool (CrossPool) 생성
 * forge script script/CreatePool.s.sol:CreatePool \
 *   --rpc-url <RPC_URL> \
 *   --broadcast
 *
 * # GamePool 생성
 * forge script script/CreatePool.s.sol:CreatePool \
 *   --sig "createGamePool()" \
 *   --rpc-url <RPC_URL> \
 *   --broadcast
 *
 * V1 필수 환경변수:
 * - CROSS_GAME_REWARD: CrossGameReward 컨트랙트 주소
 * - POOL_NAME: 생성할 풀 이름
 * - DEPOSIT_TOKEN: 예치 토큰 주소 (0x1 = Native Token)
 * - MIN_DEPOSIT_AMOUNT: 최소 예치 금액 (wei 단위)
 * V1 선택 환경변수:
 * - REWARD_TOKEN: 보상 토큰 주소 (없으면 보상 토큰을 등록하지 않음)
 *
 * GamePool 필수 환경변수:
 * - CROSS_GAME_REWARD: CrossGameReward 컨트랙트 주소
 * - POOL_NAME: 생성할 풀 이름
 * - DEPOSIT_TOKEN: 예치 토큰 주소 (게임토큰)
 * - REWARD_TOKEN: 보상 토큰 주소 (CROSSD 등)
 * - MIN_DEPOSIT_AMOUNT: 최소 예치 금액 (wei 단위)
 * GamePool 선택 환경변수:
 * - SPONSOR_ADDRESS: 라운드 생성 권한을 부여할 스폰서 지갑 주소
 */
contract CreatePool is Script {
    address public constant NATIVE_TOKEN_ADDRESS = address(0x1);

    /**
     * @notice V1 Pool (CrossPool) 생성 - 기본 함수
     */
    function run() external {
        ICrossGameReward crossGameReward = _getCrossGameReward();

        string memory poolName = vm.envString("POOL_NAME");
        IERC20 depositToken = _resolveDepositToken(crossGameReward);
        uint minDepositAmount = vm.envUint("MIN_DEPOSIT_AMOUNT");

        console.log("\n=== V1 Pool (CrossPool) Creation ===");
        console.log("CrossGameReward:", address(crossGameReward));
        console.log("Pool Name:", poolName);
        console.log("Deposit Token:", address(depositToken));
        console.log("Min Deposit Amount:", minDepositAmount);

        vm.startBroadcast();

        (uint poolId, ICrossGameRewardPool pool) = crossGameReward.createPool(poolName, depositToken, minDepositAmount);

        console.log("\nPool Created:");
        console.log("  Pool ID:", poolId);
        console.log("  Pool Address:", address(pool));
        console.log("  Pool Type: CrossPool (V1)");

        // V1: 보상 토큰 추가 (선택적)
        try vm.envAddress("REWARD_TOKEN") returns (address rewardTokenAddress) {
            crossGameReward.addRewardToken(poolId, IERC20(rewardTokenAddress));
            console.log("  Reward Token Added:", rewardTokenAddress);
        } catch {
            console.log("  REWARD_TOKEN not set, skipping");
        }

        vm.stopBroadcast();

        console.log("\n=== Done ===");
    }

    /**
     * @notice GamePool 생성
     * @dev --sig "createGamePool()" 플래그와 함께 사용
     */
    function createGamePool() external {
        ICrossGameReward crossGameReward = _getCrossGameReward();

        string memory poolName = vm.envString("POOL_NAME");
        address depositTokenAddr = vm.envAddress("DEPOSIT_TOKEN");
        address rewardTokenAddr = vm.envAddress("REWARD_TOKEN");
        uint minDepositAmount = vm.envUint("MIN_DEPOSIT_AMOUNT");

        require(depositTokenAddr != address(0), "DEPOSIT_TOKEN is required");
        require(depositTokenAddr != NATIVE_TOKEN_ADDRESS, "GamePool does not support native token deposit");
        require(rewardTokenAddr != address(0), "REWARD_TOKEN is required for GamePool");

        IERC20 depositToken = IERC20(depositTokenAddr);
        IERC20 rewardToken = IERC20(rewardTokenAddr);

        console.log("\n=== GamePool Creation ===");
        console.log("CrossGameReward:", address(crossGameReward));
        console.log("Pool Name:", poolName);
        console.log("Deposit Token:", depositTokenAddr);
        console.log("Reward Token:", rewardTokenAddr);
        console.log("Min Deposit Amount:", minDepositAmount);

        vm.startBroadcast();

        (uint poolId, ICrossGameRewardPool pool) =
            crossGameReward.createGamePool(poolName, depositToken, rewardToken, minDepositAmount);

        console.log("\nPool Created:");
        console.log("  Pool ID:", poolId);
        console.log("  Pool Address:", address(pool));
        console.log("  Pool Type: GamePool");

        // GamePool: Sponsor Role 부여 (선택적)
        try vm.envAddress("SPONSOR_ADDRESS") returns (address sponsorAddr) {
            crossGameReward.grantSponsorRole(poolId, sponsorAddr);
            console.log("  Sponsor Role Granted:", sponsorAddr);
        } catch {
            console.log("  SPONSOR_ADDRESS not set, skipping");
        }

        vm.stopBroadcast();

        console.log("\n=== Done ===");
    }

    // ==================== Internal Helpers ====================

    function _getCrossGameReward() internal view returns (ICrossGameReward) {
        address addr = vm.envAddress("CROSS_GAME_REWARD");
        return ICrossGameReward(addr);
    }

    function _resolveDepositToken(ICrossGameReward crossGameReward) internal view returns (IERC20) {
        address depositTokenAddr = vm.envAddress("DEPOSIT_TOKEN");
        if (depositTokenAddr == NATIVE_TOKEN_ADDRESS) {
            IWCROSS wcross = crossGameReward.wcross();
            console.log("Deposit Token: WCROSS (Native):", address(wcross));
            return IERC20(address(wcross));
        }
        return IERC20(depositTokenAddr);
    }
}
