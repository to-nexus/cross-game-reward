# Cross Staking 배포 스크립트

이 폴더에는 Cross Staking Protocol을 배포하고 관리하기 위한 Foundry 스크립트가 포함되어 있습니다.

## 📚 스크립트 목록

### 🚀 배포 스크립트

| 스크립트 | 설명 | 환경변수 파일 | 권장도 |
|---------|------|--------------|--------|
| **DeployWithFirstProject.s.sol** | 통합 배포 (시스템+프로젝트+리워드) | DeployWithFirstProject.env | ⭐⭐⭐ 최고 권장 |
| Deploy.s.sol | 전체 시스템 배포 (WCROSS, Protocol) | Deploy.env | ⭐ 기본 |
| CreateProject.s.sol | 프로젝트 생성 | CreateProject.env | ⭐ 기본 |
| DeployRouter.s.sol | Router & Viewer 배포 및 승인 | DeployRouter.env | ⭐⭐ 권장 |
| DeployRankingAddon.s.sol | RankingAddon 배포 | DeployRankingAddon.env | ⭐ 선택 |

### 🔧 관리 스크립트

| 스크립트 | 설명 | 환경변수 파일 | 권한 |
|---------|------|--------------|------|
| **SweepRewardPool.s.sol** | RewardPool 토큰 회수 | SweepRewardPool.env | Admin 전용 |
| TestScenario.s.sol | 테스트 시나리오 실행 | TestScenario.env | 테스터 |

---

## 🎯 빠른 시작

### 최소 배포 (2단계)

```bash
# 1. 통합 배포
forge script script/DeployWithFirstProject.s.sol:DeployWithFirstProjectScript \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --verify -vvvv

# 2. Router 배포
forge script script/DeployRouter.s.sol:DeployRouterScript \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -vvvv
```

**완료! 🎉**

---

## 📖 상세 가이드

### DeployWithFirstProject.s.sol ⭐ 신규 추가

**가장 빠르고 간편한 배포 방법!**

한 번에 모든 것을 배포:
- WCROSS (또는 기존 주소 재사용)
- StakingProtocol
- StakingRouter
- StakingViewer
- 첫 프로젝트 생성
- 선택적으로 리워드 예치

**환경변수:**
```bash
PROJECT_NAME="My Staking Project"
SEASON_BLOCKS=86400
FIRST_SEASON_START_BLOCK=1100
POOL_END_BLOCK=0

# 선택 (리워드 예치)
REWARD_TOKEN=0x...
REWARD_AMOUNT=1000000000000000000
REWARD_SEASON=1
```

**실행:**
```bash
source script/DeployWithFirstProject.env
forge script script/DeployWithFirstProject.s.sol:DeployWithFirstProjectScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast --verify -vvvv
```

**출력:**
- 웹앱용 주소 (VITE_*) - WCROSS, Protocol, Router, Viewer
- 추가 스크립트용 주소
- 프로젝트 정보

**자세한 내용:** [NEW_SCRIPTS.md](NEW_SCRIPTS.md)

---

### Deploy.s.sol

전체 시스템만 배포 (프로젝트 생성 없음)

**배포:**
- WCROSS
- StakingPoolCode
- RewardPoolCode
- StakingProtocol
- StakingRouter
- StakingViewer

**환경변수:**
```bash
# 선택
WCROSS_ADDRESS=0x...  # 기존 WCROSS 재사용
```

**실행:**
```bash
forge script script/Deploy.s.sol:DeployScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast --verify -vvvv
```

---

### CreateProject.s.sol

StakingProtocol에 새 프로젝트 생성

**환경변수:**
```bash
STAKING_PROTOCOL_ADDRESS=0x...
PROJECT_NAME="My Project"
SEASON_BLOCKS=86400
FIRST_SEASON_START_BLOCK=1100
POOL_END_BLOCK=0

# 선택
PROJECT_ADMIN=0x...  # 없으면 실행자가 관리자
```

**실행:**
```bash
source script/CreateProject.env
forge script script/CreateProject.s.sol:CreateProjectScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast -vvvv
```

---

### DeployRouter.s.sol

Native CROSS 스테이킹을 위한 Router 및 조회 전용 Viewer 배포 및 승인

**환경변수:**
```bash
WCROSS_ADDRESS=0x...
STAKING_PROTOCOL_ADDRESS=0x...
PROJECT_ID=1
```

**실행:**
```bash
source script/DeployRouter.env
forge script script/DeployRouter.s.sol:DeployRouterScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast -vvvv
```

---

### DeployRankingAddon.s.sol

Top 10 랭킹 추적을 위한 Addon 배포

**환경변수:**
```bash
STAKING_PROTOCOL_ADDRESS=0x...
STAKING_POOL_ADDRESS=0x...
PROJECT_ID=1
```

**실행:**
```bash
source script/DeployRankingAddon.env
forge script script/DeployRankingAddon.s.sol:DeployRankingAddonScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast -vvvv
```

---

### SweepRewardPool.s.sol ⭐ 신규 추가

**RewardPool에서 토큰 회수 (Protocol Admin 전용)**

잘못 전송되거나 남은 토큰을 안전하게 회수합니다.

**⚠️ 주의:**
- Protocol Admin만 실행 가능
- 되돌릴 수 없음
- 시즌 보상은 회수하지 말 것

**방법 1: RewardPool 주소 직접**
```bash
export STAKING_PROTOCOL_ADDRESS=0x...
export REWARD_POOL_ADDRESS=0x...
export TOKEN_ADDRESS=0x...
export SWEEP_TO=0x...
export SWEEP_AMOUNT=1000000000000000000

forge script script/SweepRewardPool.s.sol:SweepRewardPoolScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast -vvvv
```

**방법 2: Project ID 사용**
```bash
export STAKING_PROTOCOL_ADDRESS=0x...
export PROJECT_ID=1
export TOKEN_ADDRESS=0x...
export SWEEP_TO=0x...
export SWEEP_AMOUNT=1000000000000000000

forge script script/SweepRewardPool.s.sol:SweepRewardPoolByProjectIDScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast -vvvv
```

**자세한 내용:** [NEW_SCRIPTS.md](NEW_SCRIPTS.md)

---

### TestScenario.s.sol

전체 플로우 테스트 (스테이킹, 랭킹, 보상 등)

**환경변수:**
```bash
WCROSS_ADDRESS=0x...
STAKING_ROUTER_ADDRESS=0x...
PROJECT_ID=1
STAKE_AMOUNT=1000000000000000000

# 선택
RANKING_ADDON_ADDRESS=0x...
REWARD_TOKEN=0x...
```

**실행:**
```bash
source script/TestScenario.env
forge script script/TestScenario.s.sol:TestScenarioScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast -vvvv
```

---

## 🔄 배포 플로우

### ⭐ 추천: 빠른 배포

```bash
# Step 1: 통합 배포 (리워드 포함)
DeployWithFirstProject.s.sol
    ↓
# Step 2: Router 배포
DeployRouter.s.sol
    ↓
# 완료! 테스트 시작
```

**시간:** 5-10분  
**단계:** 2단계

### 전통적 배포

```bash
# Step 1: 시스템 배포
Deploy.s.sol
    ↓
# Step 2: 프로젝트 생성
CreateProject.s.sol
    ↓
# Step 3: Router 배포
DeployRouter.s.sol
    ↓
# Step 4 (선택): Ranking
DeployRankingAddon.s.sol
```

**시간:** 15-20분  
**단계:** 3-4단계

---

## 💡 사용 팁

### 환경변수 관리

**방법 1: .env 파일**
```bash
cp script/DeployWithFirstProject.env .env.deploy
nano .env.deploy
source .env.deploy
forge script ...
```

**방법 2: 직접 export**
```bash
export PROJECT_NAME="My Project"
export SEASON_BLOCKS=86400
forge script ...
```

### 주소 저장

배포 후 주소를 파일에 저장:
```bash
forge script ... | tee deployment.log
```

웹앱 설정:
```bash
# 출력된 VITE_* 주소를 webapp/.env에 복사
VITE_WCROSS_ADDRESS=0x...
VITE_STAKING_PROTOCOL_ADDRESS=0x...
```

### 여러 네트워크 배포

네트워크별 환경변수 파일 생성:
```bash
.env.testnet
.env.mainnet
```

---

## 📋 체크리스트

### 배포 전

- [ ] RPC_URL 설정
- [ ] PRIVATE_KEY 준비 (.env 파일)
- [ ] 계정에 가스비용 확보
- [ ] 필요한 환경변수 설정
- [ ] Etherscan API Key (verify용)

### 배포 후

- [ ] 모든 컨트랙트 주소 저장
- [ ] Etherscan에서 verify 확인
- [ ] 웹앱 .env 파일 업데이트
- [ ] 테스트 시나리오 실행
- [ ] 문서에 주소 기록

---

## 🆘 문제 해결

### "Insufficient funds"
→ 계정에 Native 토큰 확보

### "Season not active"
→ FIRST_SEASON_START_BLOCK 확인, 블록 도달까지 대기

### "Only protocol admin"
→ Protocol Admin 계정으로 실행

### 스크립트 실패
→ `-vvvv` 옵션으로 상세 로그 확인

---

## 📚 관련 문서

- **[QUICK_START.md](QUICK_START.md)** - 5분 빠른 시작
- **[NEW_SCRIPTS.md](NEW_SCRIPTS.md)** - 신규 스크립트 상세 가이드
- **[../TESTNET_DEPLOYMENT.md](../TESTNET_DEPLOYMENT.md)** - 전체 배포 가이드
- **각 .env 파일** - 환경변수 예제 및 사용법

---

## 🔐 보안

- Private Key를 Git에 커밋하지 마세요
- `.env` 파일은 `.gitignore`에 포함
- 메인넷 배포 시 하드웨어 월렛 사용 권장
- Admin 권한을 신중하게 관리

---

**Happy Deploying! 🚀**

