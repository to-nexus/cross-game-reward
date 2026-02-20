// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import "../src/CrossGameReward.sol";
import "../src/CrossGameRewardPool.sol";
import "../src/GamePool.sol";
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
 *      V1/GamePool Implementation 모두 지원
 *
 * .env 파일에서 읽어오는 환경 변수:
 * - CROSS_GAME_REWARD_ROOT_IMPLEMENTATION (required)
 * - POOL_IMPLEMENTATION (required)
 * - GAME_POOL_IMPLEMENTATION (optional, GamePool 지원 시 필요)
 * - INITIAL_DELAY (optional, default: 1 days)
 * - ADMIN_ADDRESS (optional, default: deployer)
 *
 * Pool 생성 관련:
 * - CREATE_POOL (optional, true/false)
 * - POOL_VERSION (optional, "1" 또는 "2", default: "1")
 *   V1 전용:
 *   - DEPOSIT_TOKEN (0x1=native token)
 *   - POOL_NAME, MIN_DEPOSIT_AMOUNT
 *   - REWARD_TOKEN (optional, V1 pool에 보상 토큰 등록)
 *   GamePool 전용:
 *   - DEPOSIT_TOKEN (게임토큰 주소)
 *   - REWARD_TOKEN (required, CROSSD 등)
 *   - POOL_NAME, MIN_DEPOSIT_AMOUNT
 *   - SPONSOR_ADDRESS (optional, 스폰서 지갑)
 *
 * 사용법:
 * forge script script/DeployFullSystem.s.sol:DeployFullSystem \
 *   --rpc-url <RPC_URL> \
 *   --broadcast
 */
contract DeployFullSystem is Script {
    uint48 public constant DEFAULT_INITIAL_DELAY = 1 days;
    address public constant NATIVE_TOKEN_ADDRESS = address(0x1);

    function run() external {
        address deployer = msg.sender;

        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);

        uint48 initialDelay = uint48(vm.envOr("INITIAL_DELAY", uint(DEFAULT_INITIAL_DELAY)));
        address adminAddress = vm.envOr("ADMIN_ADDRESS", deployer);

        console.log("Initial Delay:", initialDelay);
        console.log("Admin Address:", adminAddress);

        // Implementation 주소 읽기
        address poolImplAddr = vm.envAddress("POOL_IMPLEMENTATION");
        address crossGameRewardImplAddr = vm.envAddress("CROSS_GAME_REWARD_ROOT_IMPLEMENTATION");

        console.log("\n1. Pool V1 Implementation:", poolImplAddr);
        console.log("2. CrossGameReward Implementation:", crossGameRewardImplAddr);

        vm.startBroadcast();

        // 3. CrossGameReward Proxy 배포
        bytes memory initData = abi.encodeWithSelector(
            CrossGameReward.initialize.selector, ICrossGameRewardPool(poolImplAddr), adminAddress, initialDelay
        );
        ERC1967Proxy crossGameRewardProxy = new ERC1967Proxy(crossGameRewardImplAddr, initData);
        CrossGameReward crossGameReward = CrossGameReward(address(crossGameRewardProxy));
        console.log("3. CrossGameReward Proxy deployed:", address(crossGameReward));

        // 4. GamePool Implementation 등록 (선택적)
        try vm.envAddress("GAME_POOL_IMPLEMENTATION") returns (address gamePoolImplAddr) {
            crossGameReward.setGamePoolImplementation(ICrossGameRewardPool(gamePoolImplAddr));
            console.log("4. Pool GamePool Implementation set:", gamePoolImplAddr);
        } catch {
            console.log("4. GAME_POOL_IMPLEMENTATION not set, skipping GamePool setup");
        }

        // 5. Router 배포
        CrossGameRewardRouter router = new CrossGameRewardRouter(address(crossGameReward));
        console.log("5. Router deployed:", address(router));

        // 6. WCROSS 확인
        IWCROSS wcross = crossGameReward.wcross();
        console.log("6. WCROSS:", address(wcross));

        // 7. Router 등록
        crossGameReward.setRouter(address(router));
        console.log("7. Router registered");

        // 8. Pool 생성 (선택적)
        uint poolId;
        bool poolCreated = false;
        try vm.envBool("CREATE_POOL") returns (bool shouldCreatePool) {
            if (shouldCreatePool) {
                uint poolVersion = vm.envOr("POOL_VERSION", uint(1));

                if (poolVersion == 2) {
                    poolId = _createGamePool(crossGameReward);
                } else {
                    poolId = _createPoolV1(crossGameReward, wcross);
                }
                poolCreated = true;
            }
        } catch {
            console.log("\n8. CREATE_POOL not set, skipping pool creation");
        }

        // 9. V1 Reward Token 등록 (CREATE_POOL=true + POOL_VERSION=1일 때만)
        if (poolCreated) {
            uint poolVersion = vm.envOr("POOL_VERSION", uint(1));
            if (poolVersion == 1) {
                try vm.envAddress("REWARD_TOKEN") returns (address rewardTokenAddr) {
                    crossGameReward.addRewardToken(poolId, IERC20(rewardTokenAddr));
                    console.log("9. Reward Token added to V1 pool:", rewardTokenAddr);
                } catch {
                    console.log("9. No REWARD_TOKEN for V1 pool, skipping");
                }
            }
        }

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("WCROSS:", address(wcross));
        console.log("CrossGameReward (Proxy):", address(crossGameReward));
        console.log("Router:", address(router));
        console.log("Admin:", adminAddress);
        if (poolCreated) {
            console.log("Pool ID:", poolId);
        }
    }

    function _createPoolV1(CrossGameReward crossGameReward, IWCROSS wcross) internal returns (uint poolId) {
        address depositTokenAddr = vm.envAddress("DEPOSIT_TOKEN");
        string memory poolName = vm.envOr("POOL_NAME", string("Cross Game Reward Pool"));
        uint minDepositAmount = vm.envOr("MIN_DEPOSIT_AMOUNT", uint(1 ether));

        IERC20 depositToken;
        if (depositTokenAddr == NATIVE_TOKEN_ADDRESS) {
            depositToken = IERC20(address(wcross));
            console.log("\n8. Creating V1 Pool (Native/WCROSS)...");
        } else {
            depositToken = IERC20(depositTokenAddr);
            console.log("\n8. Creating V1 Pool (ERC20)...");
        }

        console.log("   Pool Name:", poolName);
        console.log("   Deposit Token:", address(depositToken));
        console.log("   Min Deposit:", minDepositAmount);

        (poolId,) = crossGameReward.createPool(poolName, depositToken, minDepositAmount);
        console.log("   Pool ID:", poolId);
    }

    function _createGamePool(CrossGameReward crossGameReward) internal returns (uint poolId) {
        address depositTokenAddr = vm.envAddress("DEPOSIT_TOKEN");
        address rewardTokenAddr = vm.envAddress("REWARD_TOKEN");
        string memory poolName = vm.envOr("POOL_NAME", string("Game Token Pool"));
        uint minDepositAmount = vm.envOr("MIN_DEPOSIT_AMOUNT", uint(1 ether));

        require(depositTokenAddr != NATIVE_TOKEN_ADDRESS, "GamePool does not support native token deposit");

        console.log("\n8. Creating GamePool...");
        console.log("   Pool Name:", poolName);
        console.log("   Deposit Token:", depositTokenAddr);
        console.log("   Reward Token:", rewardTokenAddr);
        console.log("   Min Deposit:", minDepositAmount);

        (poolId,) = crossGameReward.createGamePool(
            poolName, IERC20(depositTokenAddr), IERC20(rewardTokenAddr), minDepositAmount
        );
        console.log("   Pool ID:", poolId);

        // Sponsor Role 부여 (선택적)
        try vm.envAddress("SPONSOR_ADDRESS") returns (address sponsorAddr) {
            crossGameReward.grantSponsorRole(poolId, sponsorAddr);
            console.log("   Sponsor Role Granted:", sponsorAddr);
        } catch {
            console.log("   SPONSOR_ADDRESS not set, skipping");
        }
    }
}
