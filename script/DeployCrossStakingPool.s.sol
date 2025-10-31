// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "../src/CrossStakingPool.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "forge-std/Script.sol";

/**
 * @title DeployCrossStakingPool
 * @notice CrossStakingPool UUPS 프록시 배포 스크립트
 * @dev 사용법:
 * forge script script/DeployCrossStakingPool.s.sol:DeployCrossStakingPool \
 *   --rpc-url <RPC_URL> \
 *   --private-key <PRIVATE_KEY> \
 *   --broadcast
 */
contract DeployCrossStakingPool is Script {
    // 배포할 네트워크의 CROSS 토큰 주소를 여기에 설정
    address public constant CROSS_TOKEN = address(0); // TODO: 실제 CROSS 토큰 주소로 변경
    uint48 public constant INITIAL_DELAY = 2 days; // 관리자 변경 딜레이

    function run() external {
        // 환경 변수에서 배포자 개인키 로드
        uint deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        vm.startBroadcast(deployerPrivateKey);

        // 1. Implementation 배포
        CrossStakingPool implementation = new CrossStakingPool();
        console.log("Implementation deployed at:", address(implementation));

        // 2. Initialize data 준비
        bytes memory initData =
            abi.encodeWithSelector(CrossStakingPool.initialize.selector, IERC20(CROSS_TOKEN), deployer, INITIAL_DELAY);

        // 3. Proxy 배포
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), initData);
        console.log("Proxy deployed at:", address(proxy));

        // 4. Proxy를 통해 컨트랙트 접근
        CrossStakingPool pool = CrossStakingPool(address(proxy));
        console.log("\n=== Deployment Summary ===");
        console.log("Proxy Address:", address(pool));
        console.log("Implementation Address:", address(implementation));
        console.log("Staking token (CROSS):", address(pool.stakingToken()));
        console.log("Default Admin:", deployer);
        console.log("Admin Delay:", INITIAL_DELAY);

        vm.stopBroadcast();
    }
}

/**
 * @title DeployWithRewards
 * @notice 보상 토큰을 포함한 전체 설정 배포 스크립트
 */
contract DeployWithRewards is Script {
    // 네트워크별 토큰 주소 설정
    address public crossToken;
    address[] public rewardTokens;
    uint[] public initialRewardAmounts;
    uint48 public constant INITIAL_DELAY = 2 days;

    function run() external {
        // 환경 변수 로드
        uint deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        // 네트워크별 설정 로드
        loadNetworkConfig();

        vm.startBroadcast(deployerPrivateKey);

        // 1. Implementation 배포
        CrossStakingPool implementation = new CrossStakingPool();
        console.log("Implementation deployed at:", address(implementation));

        // 2. Initialize data 준비
        bytes memory initData =
            abi.encodeWithSelector(CrossStakingPool.initialize.selector, IERC20(crossToken), deployer, INITIAL_DELAY);

        // 3. Proxy 배포
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), initData);
        console.log("Proxy deployed at:", address(proxy));

        // 4. Proxy를 통해 컨트랙트 접근
        CrossStakingPool pool = CrossStakingPool(address(proxy));

        // 5. 보상 토큰 추가
        for (uint i = 0; i < rewardTokens.length; i++) {
            pool.addRewardToken(rewardTokens[i]);
            console.log("Added reward token:", rewardTokens[i]);
        }

        // 6. 초기 보상 토큰 공급 (선택사항)
        // 직접 transfer하면 _syncReward가 자동 감지
        for (uint i = 0; i < rewardTokens.length; i++) {
            if (initialRewardAmounts[i] > 0) {
            IERC20 rewardToken = IERC20(rewardTokens[i]);

                // 직접 transfer
                rewardToken.transfer(address(pool), initialRewardAmounts[i]);

                console.log("Transferred reward:", rewardTokens[i]);
            console.log("  Amount:", initialRewardAmounts[i]);
            }
        }

        console.log("\n=== Deployment Summary ===");
        console.log("Pool Proxy Address:", address(pool));
        console.log("Implementation Address:", address(implementation));
        console.log("CROSS Token:", crossToken);
        console.log("Number of reward tokens:", rewardTokens.length);
        console.log("Default Admin:", deployer);

        vm.stopBroadcast();
    }

    function loadNetworkConfig() internal {
        // 체인 ID에 따라 다른 설정 로드
        uint chainId = block.chainid;

        if (chainId == 1) {
            // Ethereum Mainnet
            loadMainnetConfig();
        } else if (chainId == 11155111) {
            // Sepolia Testnet
            loadSepoliaConfig();
        } else if (chainId == 31337) {
            // Anvil Local
            loadLocalConfig();
        } else {
            revert("Unsupported network");
        }
    }

    function loadMainnetConfig() internal {
        // TODO: 메인넷 주소로 변경
        crossToken = address(0);

        rewardTokens = new address[](2);
        rewardTokens[0] = address(0); // 보상 토큰 1
        rewardTokens[1] = address(0); // 보상 토큰 2

        initialRewardAmounts = new uint[](2);
        initialRewardAmounts[0] = 10000 ether;
        initialRewardAmounts[1] = 5000 ether;
    }

    function loadSepoliaConfig() internal {
        // TODO: Sepolia 테스트넷 주소로 변경
        crossToken = address(0);

        rewardTokens = new address[](1);
        rewardTokens[0] = address(0);

        initialRewardAmounts = new uint[](1);
        initialRewardAmounts[0] = 1000 ether;
    }

    function loadLocalConfig() internal {
        // 로컬 개발 환경 - 실제 주소는 배포 후 설정
        crossToken = vm.envAddress("CROSS_TOKEN_ADDRESS");

        rewardTokens = new address[](1);
        rewardTokens[0] = vm.envAddress("REWARD_TOKEN_ADDRESS");

        initialRewardAmounts = new uint[](1);
        initialRewardAmounts[0] = 10000 ether;
    }
}
