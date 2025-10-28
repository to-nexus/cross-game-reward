# Cross Staking Protocol v1.0

> 블록체인 기반 시즌제 스테이킹 프로토콜

## 개요

Cross Staking Protocol은 시즌 기반의 분산형 스테이킹 플랫폼입니다. 프로젝트별로 독립적인 스테이킹 풀을 생성하고, 시즌마다 공정한 보상 분배를 제공합니다.

### 주요 특징

- ⏱️ **시즌 기반 시스템**: 블록 기반 시즌으로 명확한 보상 구간
- 🎯 **포인트 시스템**: 스테이킹 금액 × 시간으로 공정한 보상 계산
- 🏭 **프로젝트별 독립**: Factory 패턴으로 프로젝트마다 독립적인 풀
- 🔄 **Native Token 지원**: WCROSS 자동 래핑으로 편리한 사용성
- 🔐 **보안 강화**: Reentrancy Guard, Access Control, Pausable 패턴
- ⚡ **가스 최적화**: Custom Error, Storage 최적화로 10-15% 절감

## 아키텍처

```
┌─────────────────────────────────────────────────────────┐
│                   StakingProtocol                       │
│              (Factory & Global Manager)                  │
└────────────────────┬────────────────────────────────────┘
                     │ CREATE2
            ┌────────┴────────┐
            │                 │
    ┌───────▼──────┐   ┌─────▼────────┐
    │ StakingPool  │───│ RewardPool   │
    │ (Project 1)  │   │ (Project 1)  │
    └──────────────┘   └──────────────┘
    
    ┌──────────────┐   ┌──────────────┐
    │ StakingPool  │───│ RewardPool   │
    │ (Project 2)  │   │ (Project 2)  │
    └──────────────┘   └──────────────┘
    
┌──────────────┐    ┌──────────────┐
│StakingRouter │    │StakingViewer │
│ (TX Handler) │    │(View Queries)│
└──────────────┘    └──────────────┘
```

## 핵심 컨트랙트

### StakingProtocol (Factory)
프로젝트별 스테이킹 풀 생성 및 전역 설정 관리

### StakingPool
- 토큰 스테이킹 및 출금
- 시즌 자동 롤오버
- 포인트 계산 및 집계
- 보상 청구

### RewardPool
- 보상 토큰 예치
- 보상 분배
- 시즌별 토큰 관리

### StakingRouter
- Native CROSS ↔ WCROSS 자동 변환
- 편의 함수 제공

### StakingViewer
- 모든 조회 함수 통합
- 가상 시즌 계산
- Batch 조회 지원

## 설치 및 실행

### Prerequisites
```bash
# Foundry 설치
curl -L https://foundry.paradigm.xyz | bash
foundryup

# 의존성 설치
forge install
```

### 컴파일
```bash
forge build
```

### 테스트
```bash
# 전체 테스트
forge test

# 가스 리포트
forge test --gas-report

# 커버리지
forge coverage
```

## 배포

### Testnet 배포
```bash
# 환경변수 설정
cp script/DeployWithFirstProject.env .env
# .env 파일 수정 후

# 배포 실행
forge script script/DeployWithFirstProject.s.sol:DeployWithFirstProjectScript \
    --rpc-url $RPC_URL \
    --sender $DEPLOYER \
    --keystore $KEYSTORE_PATH \
    --broadcast \
    --slow -vv
```

### 배포된 컨트랙트 (Testnet)
```
WCROSS: 0x494DC6816D77a77eBd7E3a28f6671Ab15586d577
StakingProtocol: 0x5404C56dC66Cf685A9b85F0B131Aa27e55828fF5
StakingRouter: 0xd87030275A699D4D301E31e89f9D43657dB19000
StakingViewer: 0x1cb1941c0452c844FFD2c4F446e2B06325219338

Project ID 1:
  StakingPool: 0xa862629377933063954E2e814667208b5B95f477
  RewardPool: 0xC07C614ebDB17e438cb3d7CC9566c4015F2BF09D
```

## 사용 예시

### 스테이킹
```solidity
// Native CROSS로 스테이킹
stakingRouter.stake{value: 5 ether}(projectID);

// WCROSS로 직접 스테이킹
wcross.approve(address(stakingPool), 5 ether);
stakingPool.stake(5 ether);
```

### 출금
```solidity
// Native CROSS로 출금
stakingRouter.unstake(projectID);

// WCROSS로 직접 출금
stakingPool.withdrawAll();
```

### 보상 청구
```solidity
// 단일 시즌 청구
stakingPool.claimSeason(seasonNumber, rewardTokenAddress);

// 다중 시즌 청구
uint[] memory seasons = [1, 2, 3];
address[] memory tokens = [token1, token2, token3];
stakingRouter.claimMultipleRewards(projectID, seasons, tokens);
```

### 조회
```solidity
// 현재 포인트 조회
uint points = stakingViewer.getUserPoints(projectID, userAddress);

// 시즌 정보 조회
(uint season, uint startBlock, uint endBlock, uint blocksElapsed) = 
    stakingViewer.getSeasonInfo(projectID);

// 예상 보상 조회
uint expectedReward = stakingViewer.getClaimableReward(
    projectID, userAddress, seasonNumber, rewardTokenAddress
);
```

## 보안

### 적용된 보안 패턴
- ✅ ReentrancyGuardTransient (EIP-1153)
- ✅ AccessControlDefaultAdminRules (3-day timelock)
- ✅ Pausable Pattern
- ✅ SafeERC20
- ✅ Custom Error (가스 효율)
- ✅ Checks-Effects-Interactions Pattern

### 테스트 커버리지
- 총 테스트: 94개 (Security 테스트 포함)
- 통과율: 89/94 (94.7%)
- 주요 시나리오 커버리지: 100%

### 감사 상태
- ⏳ 내부 감사: 완료
- ⏳ 외부 감사: 진행 예정

## 가스 최적화

### 적용된 최적화 기법
1. **Custom Error**: 문자열 대비 15-20% 절감
2. **Named Import**: 컴파일 효율 향상
3. **Unchecked Arithmetic**: 안전한 연산에 5-10% 절감
4. **ReentrancyGuardTransient**: 기존 대비 30% 절감
5. **Immutable Variables**: Storage 접근 비용 절감

### 예상 가스 비용
| 작업 | 가스 비용 | 비고 |
|------|-----------|------|
| Stake | ~130k gas | Native CROSS 사용 시 |
| Unstake | ~155k gas | Native CROSS 수령 시 |
| Claim Reward | ~105k gas | 단일 시즌 |
| Season Rollover | ~260k gas | 자동 롤오버 |

## 개발 가이드

### 프로젝트 구조
```
src/
├── base/                  # 추상 컨트랙트
│   ├── CrossStakingBase.sol
│   ├── StakingPoolBase.sol
│   └── RewardPoolBase.sol
├── interfaces/            # 인터페이스
│   ├── IStakingPool.sol
│   ├── IRewardPool.sol
│   └── IStakingProtocol.sol
├── libraries/             # 라이브러리
│   ├── PointsLib.sol
│   └── SeasonLib.sol
├── StakingProtocol.sol    # Factory
├── StakingPool.sol        # 스테이킹 풀
├── RewardPool.sol         # 보상 풀
├── StakingRouter.sol      # Native Token 라우터
├── StakingViewer.sol      # View 함수 통합
└── WCROSS.sol            # Wrapped Token

test/
├── BaseTest.sol          # 기본 설정
├── Security.t.sol        # 보안 테스트
├── Staking.t.sol         # 스테이킹 테스트
├── Season.t.sol          # 시즌 테스트
├── Points.t.sol          # 포인트 테스트
├── Rewards.t.sol         # 보상 테스트
└── ...
```

### 코딩 규칙
1. Solidity 0.8.28 사용
2. Named Import 패턴
3. Custom Error 사용
4. NatSpec 주석 작성
5. 100자 줄 길이 제한

### 테스트 작성
```solidity
// test/MyFeature.t.sol
contract MyFeatureTest is BaseTest {
    function test_MyFeature() public {
        // Arrange
        vm.startPrank(user1);
        
        // Act
        uint result = contract.myFunction();
        
        // Assert
        assertEq(result, expectedValue);
        vm.stopPrank();
    }
}
```

## 문서

- [상세 문서](docs/project-info/README.md)
- [배포 가이드](DEPLOYMENT.md)
- [테스트 가이드](TESTS.md)
- [최적화 보고서](OPTIMIZATION_REPORT.md)
- [웹앱 연동](WEBAPP_INTEGRATION_META.md)

## 라이센스

MIT License

## 기여

기여를 환영합니다! PR을 제출하기 전에:
1. 모든 테스트 통과 확인
2. 코딩 규칙 준수
3. 상세한 커밋 메시지 작성

## 연락처

- GitHub: [to-nexus/cross-staking](https://github.com/to-nexus/cross-staking)
- Documentation: [docs/](docs/)

---

**v1.0.0** - Production Ready
