# 배포 가이드

Cross-Staking Protocol 배포 절차 및 설정 가이드

## 사전 요구사항

### 1. 환경 설정
```bash
# Foundry 설치
curl -L https://foundry.paradigm.xyz | bash
foundryup

# 의존성 설치
forge install

# 환경 변수 설정 (.env)
PRIVATE_KEY=your_private_key
RPC_URL=your_rpc_url
ETHERSCAN_API_KEY=your_etherscan_key
```

### 2. 컴파일 확인
```bash
forge build
```

### 3. 테스트 실행
```bash
forge test
forge test --gas-report  # 가스 리포트 포함
```

## 배포 순서

### 1단계: WCROSS 배포 (Native CROSS Wrapper)

```bash
forge create src/WCROSS.sol:WCROSS \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --verify
```

**주의**: WCROSS 주소를 기록해두세요!

### 2단계: Code 컨트랙트 배포

#### StakingPoolCode
```bash
forge create src/StakingPool.sol:StakingPoolCode \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --verify
```

#### RewardPoolCode
```bash
forge create src/RewardPool.sol:RewardPoolCode \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --verify
```

**주의**: 각 Code 컨트랙트 주소를 기록해두세요!

### 3단계: StakingProtocol 배포

```bash
forge create src/StakingProtocol.sol:StakingProtocol \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --constructor-args \
        <WCROSS_ADDRESS> \
        <STAKING_POOL_CODE_ADDRESS> \
        <REWARD_POOL_CODE_ADDRESS> \
        <ADMIN_ADDRESS> \
    --verify
```

**파라미터**:
- `WCROSS_ADDRESS`: WCROSS 컨트랙트 주소
- `STAKING_POOL_CODE_ADDRESS`: StakingPoolCode 주소
- `REWARD_POOL_CODE_ADDRESS`: RewardPoolCode 주소
- `ADMIN_ADDRESS`: 초기 관리자 주소

### 4단계: StakingRouter 배포

```bash
forge create src/StakingRouter.sol:StakingRouter \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --constructor-args \
        <WCROSS_ADDRESS> \
        <STAKING_PROTOCOL_ADDRESS> \
    --verify
```

## 배포 후 설정

### 1. Protocol 설정

#### Router 승인
```solidity
// StakingProtocol에서 StakingRouter를 승인된 Router로 등록
// 각 프로젝트별로 설정 필요
protocol.setApprovedRouter(projectId, routerAddress, true);
```

#### 기본 시즌 길이 설정 (선택)
```solidity
protocol.setDefaultSeasonBlocks(blocks); // 기본: 2592000 블록 (약 30일)
```

### 2. 프로젝트 생성

```solidity
uint projectId = protocol.createProject(
    "프로젝트 이름",
    seasonBlocks,           // 시즌 길이 (블록 수), 0이면 기본값 사용
    firstSeasonStartBlock,  // 첫 시즌 시작 블록
    poolEndBlock            // 풀 종료 블록 (0이면 무한)
);
```

**예시**:
```solidity
// 30일 시즌, 현재 블록에서 시작, 무한 진행
uint projectId = protocol.createProject(
    "My DeFi Project",
    2592000,           // 30일 (1초/블록 기준)
    block.number,      // 현재 블록부터 시작
    0                  // 무한 진행
);
```

### 3. 보상 설정

```solidity
// 시즌별 보상 예치
// 프로젝트 크리에이터만 호출 가능
protocol.fundProjectSeason(
    projectId,
    seasonNumber,
    rewardTokenAddress,
    amount
);
```

**주의**: 보상 토큰을 먼저 approve해야 합니다!

```solidity
// ERC20 approve
IERC20(rewardToken).approve(rewardPoolAddress, amount);

// 보상 예치
protocol.fundProjectSeason(projectId, 1, rewardToken, 1000 ether);
```

## 배포 스크립트 예시

```solidity
// script/Deploy.s.sol
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/WCROSS.sol";
import "../src/StakingProtocol.sol";
import "../src/StakingRouter.sol";
import "../src/StakingPool.sol";
import "../src/RewardPool.sol";

contract DeployScript is Script {
    function run() external {
        uint deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address admin = vm.envAddress("ADMIN_ADDRESS");
        
        vm.startBroadcast(deployerPrivateKey);

        // 1. WCROSS 배포
        WCROSS wcross = new WCROSS();
        console.log("WCROSS deployed:", address(wcross));

        // 2. Code 컨트랙트 배포
        StakingPoolCode stakingPoolCode = new StakingPoolCode();
        RewardPoolCode rewardPoolCode = new RewardPoolCode();
        console.log("StakingPoolCode:", address(stakingPoolCode));
        console.log("RewardPoolCode:", address(rewardPoolCode));

        // 3. StakingProtocol 배포
        StakingProtocol protocol = new StakingProtocol(
            address(wcross),
            address(stakingPoolCode),
            address(rewardPoolCode),
            admin
        );
        console.log("StakingProtocol:", address(protocol));

        // 4. StakingRouter 배포
        StakingRouter router = new StakingRouter(
            address(wcross),
            address(protocol)
        );
        console.log("StakingRouter:", address(router));

        vm.stopBroadcast();
    }
}
```

**실행**:
```bash
forge script script/Deploy.s.sol:DeployScript \
    --rpc-url $RPC_URL \
    --broadcast \
    --verify \
    --slow  # 검증 안정성을 위해 slow 옵션 권장
```

## 설정 체크리스트

배포 후 다음 항목들을 확인하세요:

- [ ] WCROSS 배포 및 동작 확인
- [ ] StakingProtocol 배포 및 관리자 권한 확인
- [ ] StakingRouter 배포 및 WCROSS 연동 확인
- [ ] 테스트 프로젝트 생성 및 스테이킹 테스트
- [ ] Router 승인 설정 확인
- [ ] 보상 예치 및 청구 테스트
- [ ] 시즌 롤오버 테스트
- [ ] 가스비 프로필링
- [ ] Etherscan 검증 확인

## 시즌 관리

### 시즌 시작
- 자동: `firstSeasonStartBlock`에 도달하면 자동 시작
- 수동: 풀 종료 후 `setNextSeasonStart()`로 재시작

### 시즌 롤오버
```solidity
// 시즌 종료 블록 도달 후 누구나 호출 가능
stakingPool.rolloverSeason();
```

### 시즌 종료
```solidity
// 풀 전체 종료
protocol.setPoolEndBlock(projectId, endBlock);

// 또는 StakingPool에서 직접
stakingPool.setPoolEndBlock(endBlock);
```

## 프로젝트별 파라미터 설정

### 포인트 계산 시간 단위
```solidity
// 기본: 1 hour
protocol.setPoolPointsTimeUnit(projectId, 1 hours);

// 예: 1일로 변경
protocol.setPoolPointsTimeUnit(projectId, 1 days);
```

### 블록 시간
```solidity
// 기본: 1초/블록
protocol.setPoolBlockTime(projectId, 1);

// 예: 2초/블록 (다른 체인)
protocol.setPoolBlockTime(projectId, 2);
```

## 모니터링 및 유지보수

### 주요 모니터링 지표
1. **프로젝트별 총 스테이킹**: `stakingPool.totalStaked()`
2. **시즌 진행 상황**: `stakingPool.getCurrentSeasonInfo()`
3. **시즌 총 포인트**: `stakingPool.seasonTotalPointsSnapshot(seasonNum)`
4. **보상 풀 잔액**: 각 RewardPool의 토큰 잔액

### 일반적인 유지보수 작업
1. 시즌 종료 후 rolloverSeason 호출 (자동화 권장)
2. 다음 시즌 보상 예치
3. 미청구 보상 모니터링
4. 가스비 최적화 검토

## 트러블슈팅

### "StakingPoolNoActiveSeason" 에러
- **원인**: 시즌이 시작되지 않음
- **해결**: `nextSeasonStartBlock`이 도달했는지 확인, 필요시 `setNextSeasonStart()` 호출

### "StakingPoolSeasonNotEnded" 에러
- **원인**: 시즌이 아직 진행 중
- **해결**: `endBlock`을 초과한 후 rolloverSeason 호출

### "StakingPoolAlreadyClaimed" 에러
- **원인**: 이미 해당 시즌 보상 청구함
- **해결**: 정상 동작, 중복 청구 불가

### 가스비가 너무 높음
- rolloverSeason의 lazy evaluation 확인
- 불필요한 staker 순회 제거 확인
- 배치 작업 시 가스 한도 조정

## 보안 권장사항

1. **Multi-sig 사용**: 관리자 권한을 multi-sig 지갑으로 설정
2. **Timelock**: 중요한 변경사항에 대해 timelock 적용
3. **모니터링**: 이상 거래 자동 감지 시스템 구축
4. **감사**: 배포 전 외부 보안 감사 수행
5. **업그레이드 계획**: Proxy 패턴 적용 고려

## 지원

문의사항이나 이슈가 있을 경우:
- GitHub Issues: [repository URL]
- 개발자 문서: [docs URL]
- 디스코드: [discord URL]
