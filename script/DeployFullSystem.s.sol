// SPDX-License-Identifier: BUSL-1.1
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
 * .env 파일에서 읽어오는 환경 변수:
 * - CROSS_GAME_REWARD_ROOT_IMPLEMENTATION (required, 필수)
 * - POOL_IMPLEMENTATION (required, 필수)
 * - INITIAL_DELAY (optional, default: 2 days)
 * - ADMIN_ADDRESS (optional, default: deployer)
 * - CREATE_POOL (optional, true/false - Pool 생성 여부)
 *   CREATE_POOL=true일 때 필요한 환경 변수:
 *   - DEPOSIT_TOKEN (required, 0x1=native token, 다른 주소=ERC20 token)
 *   - POOL_NAME (optional, default: "Cross Game Reward Pool")
 *   - MIN_DEPOSIT_AMOUNT (optional, default: 1 ether)
 * - REWARD_TOKEN (optional, 설정시 자동 등록)
 *   CREATE_POOL=true이면 생성된 pool에 등록
 *   CREATE_POOL=false이면 POOL_ID에 등록 (POOL_ID 필수)
 * - POOL_ID (optional, CREATE_POOL=false일 때 reward token 등록할 기존 pool ID)
 *
 * 사용법:
 * forge script script/DeployFullSystem.s.sol:DeployFullSystem \
 *   --rpc-url <RPC_URL> \
 *   --private-key <PRIVATE_KEY> \
 *   --broadcast
 */
contract DeployFullSystem is Script {
    uint48 public constant DEFAULT_INITIAL_DELAY = 2 days;
    address public constant NATIVE_TOKEN_ADDRESS = address(0x1);

    function run() external {
        address deployer = msg.sender;

        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);

        // .env에서 설정값 읽기
        uint48 initialDelay = uint48(vm.envOr("INITIAL_DELAY", uint(DEFAULT_INITIAL_DELAY)));
        address adminAddress = vm.envOr("ADMIN_ADDRESS", deployer);

        console.log("Initial Delay:", initialDelay);
        console.log("Admin Address:", adminAddress);

        // Implementation 주소 읽기 (필수)
        address poolImplAddr = vm.envAddress("POOL_IMPLEMENTATION");
        address crossGameRewardImplAddr = vm.envAddress("CROSS_GAME_REWARD_ROOT_IMPLEMENTATION");

        console.log("\n1. Pool Implementation:", poolImplAddr);
        console.log("2. CrossGameReward Implementation:", crossGameRewardImplAddr);

        vm.startBroadcast();

        bytes memory initData = abi.encodeWithSelector(
            CrossGameReward.initialize.selector, ICrossGameRewardPool(poolImplAddr), adminAddress, initialDelay
        );
        ERC1967Proxy crossGameRewardProxy = new ERC1967Proxy(crossGameRewardImplAddr, initData);
        CrossGameReward crossGameReward = CrossGameReward(address(crossGameRewardProxy));
        console.log("3. CrossGameReward Proxy deployed:", address(crossGameReward));

        // 4. CrossGameRewardRouter 배포
        CrossGameRewardRouter router = new CrossGameRewardRouter(address(crossGameReward));
        console.log("4. Router deployed:", address(router));

        // 5. WCROSS 주소 확인 (CrossGameReward가 생성함)
        IWCROSS wcross = crossGameReward.wcross();
        console.log("5. WCROSS deployed by CrossGameReward:", address(wcross));

        // 6. Router 등록
        crossGameReward.setRouter(address(router));
        console.log("6. Router set in CrossGameReward");

        // 7. Pool 생성 (선택적)
        uint poolId;
        bool poolCreated = false;
        try vm.envBool("CREATE_POOL") returns (bool shouldCreatePool) {
            if (shouldCreatePool) {
                // DEPOSIT_TOKEN 읽기 (필수)
                address depositTokenAddr = vm.envAddress("DEPOSIT_TOKEN");
                require(depositTokenAddr != address(0), "DEPOSIT_TOKEN is required when CREATE_POOL=true");

                // Pool 설정 읽기
                string memory poolName = vm.envOr("POOL_NAME", string("Cross Game Reward Pool"));
                uint minDepositAmount = vm.envOr("MIN_DEPOSIT_AMOUNT", uint(1 ether));

                console.log("\n7. Creating Pool...");
                console.log("   Pool Name:", poolName);
                console.log("   Min Deposit Amount:", minDepositAmount);

                // Deposit Token 확인 (0x1 = native token, 그 외 = ERC20)
                IERC20 depositToken;
                if (depositTokenAddr == NATIVE_TOKEN_ADDRESS) {
                    depositToken = IERC20(address(wcross));
                    console.log("   Deposit Token: WCROSS (Native)");
                    console.log("   Token Address:", address(wcross));
                } else {
                    depositToken = IERC20(depositTokenAddr);
                    console.log("   Deposit Token: ERC20");
                    console.log("   Token Address:", depositTokenAddr);
                }

                (poolId,) = crossGameReward.createPool(poolName, depositToken, minDepositAmount);
                poolCreated = true;
                console.log("   Pool created with ID:", poolId);
            }
        } catch {
            console.log("\n7. CREATE_POOL not set or false, skipping pool creation");
        }

        // 8. Reward Token 등록 (선택적)
        try vm.envAddress("REWARD_TOKEN") returns (address rewardTokenAddress) {
            if (rewardTokenAddress != address(0)) {
                uint targetPoolId;

                // Pool이 생성되었으면 해당 pool에 등록
                if (poolCreated) {
                    targetPoolId = poolId;
                    console.log("\n8. Adding Reward Token to newly created pool...");
                    console.log("   Pool ID:", targetPoolId);
                    console.log("   Token Address:", rewardTokenAddress);

                    crossGameReward.addRewardToken(targetPoolId, IERC20(rewardTokenAddress));
                    console.log("   Reward Token added successfully!");
                }
                // Pool이 생성되지 않았으면 POOL_ID 필요
                else {
                    try vm.envUint("POOL_ID") returns (uint envPoolId) {
                        targetPoolId = envPoolId;
                        console.log("\n8. Adding Reward Token to existing pool...");
                        console.log("   Pool ID:", targetPoolId);
                        console.log("   Token Address:", rewardTokenAddress);

                        crossGameReward.addRewardToken(targetPoolId, IERC20(rewardTokenAddress));
                        console.log("   Reward Token added successfully!");
                    } catch {
                        console.log("\n8. REWARD_TOKEN set but POOL_ID not provided.");
                        console.log("   Skipping reward token registration (CREATE_POOL=false and no POOL_ID)");
                    }
                }
            }
        } catch {
            console.log("\n8. No REWARD_TOKEN found, skipping reward token registration");
        }

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("WCROSS:", address(wcross));
        console.log("CrossGameReward (Proxy):", address(crossGameReward));
        console.log("CrossGameReward (Impl):", crossGameRewardImplAddr);
        console.log("Pool Implementation:", poolImplAddr);
        console.log("Router:", address(router));
        console.log("Admin:", adminAddress);
        console.log("\n=== Next Steps ===");
        if (poolCreated) {
            console.log("1. Add more reward tokens via CrossGameReward.addRewardToken()");
            console.log("2. Users deposit via Router.depositNative() or Router.depositERC20()");
        } else {
            console.log("1. Create pools via CrossGameReward.createPool()");
            console.log("2. Add reward tokens via CrossGameReward.addRewardToken()");
            console.log("3. Users deposit via Router.depositNative() or Router.depositERC20()");
        }
    }
}

/*
 * ===================================
 * DeployFullSystem.env.example
 * ===================================
 *
 * # --------------------------------------------------
 * # Implementation 주소 (필수)
 * # --------------------------------------------------
 * # DeployImpl.s.sol로 먼저 배포한 주소를 사용하거나,
 * # 기존에 배포된 Implementation 주소를 사용합니다.
 * # 필수 환경 변수이므로, 설정하지 않으면 스크립트가 실패합니다.
 *
 * CROSS_GAME_REWARD_ROOT_IMPLEMENTATION=0x...  # CrossGameReward Implementation 주소 (필수)
 * POOL_IMPLEMENTATION=0x...                     # Pool Implementation 주소 (필수)
 *
 * # --------------------------------------------------
 * # 기본 설정 (선택적)
 * # --------------------------------------------------
 * INITIAL_DELAY=172800                    # 2일 (초 단위) - 기본값: 172800
 * ADMIN_ADDRESS=0x...                     # Admin 주소 (기본값: deployer)
 *
 * # --------------------------------------------------
 * # 시나리오 1: 기본 시스템만 배포 (Pool 없음)
 * # --------------------------------------------------
 * CROSS_GAME_REWARD_ROOT_IMPLEMENTATION=0x1234...
 * POOL_IMPLEMENTATION=0x5678...
 * # CREATE_POOL=false (또는 설정 안 함)
 * # Implementation 주소를 사용하여 Proxy, Router만 배포합니다.
 *
 * # --------------------------------------------------
 * # 시나리오 2: Native Token Pool 생성
 * # --------------------------------------------------
 * CROSS_GAME_REWARD_ROOT_IMPLEMENTATION=0x1234...
 * POOL_IMPLEMENTATION=0x5678...
 *
 * CREATE_POOL=true
 * DEPOSIT_TOKEN=0x1                       # 0x1 = Native Token (WCROSS)
 * POOL_NAME="Native CROSS Pool"           # Pool 이름 (기본값: "Cross Game Reward Pool")
 * MIN_DEPOSIT_AMOUNT=1000000000000000000  # 1 ether (기본값: 1 ether)
 *
 * # --------------------------------------------------
 * # 시나리오 3: ERC20 Token Pool 생성
 * # --------------------------------------------------
 * CROSS_GAME_REWARD_ROOT_IMPLEMENTATION=0x1234...
 * POOL_IMPLEMENTATION=0x5678...
 * CREATE_POOL=true
 * DEPOSIT_TOKEN=0xabcd...                 # ERC20 토큰 주소
 * POOL_NAME="USDT Pool"
 * MIN_DEPOSIT_AMOUNT=1000000              # 1 USDT (decimals=6)
 *
 * # --------------------------------------------------
 * # 시나리오 4: Pool 생성 + Reward Token 등록
 * # --------------------------------------------------
 * CROSS_GAME_REWARD_ROOT_IMPLEMENTATION=0x1234...
 * POOL_IMPLEMENTATION=0x5678...
 * CREATE_POOL=true
 * DEPOSIT_TOKEN=0x1                       # Native Token
 * POOL_NAME="Native CROSS Pool"
 * MIN_DEPOSIT_AMOUNT=1000000000000000000
 * REWARD_TOKEN=0x9999...                  # Reward Token 주소 (생성된 pool에 자동 등록)
 *
 * # --------------------------------------------------
 * # 시나리오 5: 기존 Pool에 Reward Token 등록
 * # --------------------------------------------------
 * CROSS_GAME_REWARD_ROOT_IMPLEMENTATION=0x1234...
 * POOL_IMPLEMENTATION=0x5678...
 * # CREATE_POOL=false (또는 설정 안 함)
 * REWARD_TOKEN=0x9999...                  # Reward Token 주소
 * POOL_ID=0                                # 기존 Pool ID (필수)
 */
