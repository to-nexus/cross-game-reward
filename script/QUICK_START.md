# 🚀 Quick Start - 테스트넷 배포 가이드

## 📝 간단 요약

5단계로 테스트넷에 배포하고 테스트하세요!

```
1. Deploy.s.sol          → 전체 시스템 배포
2. CreateProject.s.sol   → 프로젝트 생성
3. DeployRouter.s.sol    → Router 배포
4. DeployRankingAddon.s.sol → Ranking Addon 배포 (선택)
5. TestScenario.s.sol    → 테스트 시나리오 실행
```

---

## ⚙️ 환경 설정

### .env 파일 생성 (프로젝트 루트)

```bash
# Private Key (스크립트를 실행할 계정의 private key)
PRIVATE_KEY=0xYourPrivateKeyHere

# RPC URL
RPC_URL=https://your-testnet-rpc-url
```

> ⚠️ **중요**: `.env` 파일을 git에 커밋하지 마세요! `.gitignore`에 추가되어 있는지 확인하세요.

### 계정 준비

스크립트는 환경변수로 제공된 `PRIVATE_KEY`의 주소를 사용하여 트랜잭션을 전송합니다.

- 배포 스크립트 (Deploy, CreateProject, DeployRouter, DeployRankingAddon): 관리자 권한을 가질 주소의 private key 사용
- 테스트 스크립트 (TestScenario): 일반 사용자 주소의 private key 사용 (또는 동일 주소 사용 가능)

---

## 🎯 Step 1: 전체 시스템 배포

```bash
# .env 파일 로드
source .env

# 배포 (스크립트 실행 계정이 Protocol Admin이 됩니다)
forge script script/Deploy.s.sol:DeployScript \
    --rpc-url $RPC_URL \
    --broadcast \
    --verify \
    -vvvv

# 결과 저장 (출력에서 복사)
export WCROSS_ADDRESS=0x...
export STAKING_PROTOCOL_ADDRESS=0x...
```

> 💡 **보안 Tip**: 메인넷에서는 `--ledger` (하드웨어 월렛) 사용 권장

---

## 🎯 Step 2: 프로젝트 생성

```bash
# 환경변수 설정
export PROJECT_NAME="My Test Project"
export SEASON_BLOCKS=86400  # 1일
export FIRST_SEASON_START_BLOCK=$(cast block-number --rpc-url $RPC_URL | awk '{print $1 + 100}')
export POOL_END_BLOCK=0  # 무한

# 프로젝트 생성 (스크립트 실행 계정이 Project Admin이 됩니다)
forge script script/CreateProject.s.sol:CreateProjectScript \
    --rpc-url $RPC_URL \
    --broadcast \
    -vvvv

# 결과 저장
export PROJECT_ID=1
export STAKING_POOL_ADDRESS=0x...
export REWARD_POOL_ADDRESS=0x...
```

> 💡 **Tip**: `PROJECT_ADMIN` 환경변수를 설정하면 다른 주소를 Project Admin으로 지정할 수 있습니다.

---

## 🎯 Step 3: Router 배포

```bash
# Router 배포 및 승인
forge script script/DeployRouter.s.sol:DeployRouterScript \
    --rpc-url $RPC_URL \
    --broadcast \
    -vvvv

# 결과 저장
export STAKING_ROUTER_ADDRESS=0x...
```

---

## 🎯 Step 4: RankingAddon 배포 (선택)

```bash
# RankingAddon 배포
forge script script/DeployRankingAddon.s.sol:DeployRankingAddonScript \
    --rpc-url $RPC_URL \
    --broadcast \
    -vvvv

# 결과 저장
export RANKING_ADDON_ADDRESS=0x...
```

---

## 🎯 Step 5: 테스트 시나리오 실행

```bash
# 스테이킹 금액 설정 (1 ether)
export STAKE_AMOUNT=1000000000000000000

# 시나리오 실행
# 테스트할 계정의 private key로 실행 (기존 PRIVATE_KEY 사용 가능)
forge script script/TestScenario.s.sol:TestScenarioScript \
    --rpc-url $RPC_URL \
    --broadcast \
    -vvvv
```

> 💡 **Tip**: 다른 사용자 계정으로 테스트하려면 `.env` 파일의 `PRIVATE_KEY`를 변경하세요.

---

## ✅ 검증 방법

### 1. Etherscan에서 확인
```
각 컨트랙트 주소를 Etherscan에서 검색
→ "Contract" 탭 확인
→ "Read Contract" / "Write Contract" 기능 테스트
```

### 2. Cast 명령어로 확인

```bash
# 스테이킹 잔액 확인
cast call $STAKING_POOL_ADDRESS \
    "userStakes(address)(uint256,uint256,uint256)" \
    $YOUR_ADDRESS \
    --rpc-url $RPC_URL

# Top 10 랭커 확인
cast call $RANKING_ADDON_ADDRESS \
    "getTopRankers(uint256)(address[],uint256[])" \
    1 \
    --rpc-url $RPC_URL

# 현재 시즌 확인
cast call $STAKING_POOL_ADDRESS \
    "currentSeason()(uint256)" \
    --rpc-url $RPC_URL
```

---

## 🔧 한 줄 명령어 (All-in-One)

```bash
# .env 파일 로드
source .env

# 전체 배포를 순차 실행
forge script script/Deploy.s.sol:DeployScript \
    --rpc-url $RPC_URL --broadcast --verify -vvvv && \
sleep 5 && \
forge script script/CreateProject.s.sol:CreateProjectScript \
    --rpc-url $RPC_URL --broadcast -vvvv && \
sleep 5 && \
forge script script/DeployRouter.s.sol:DeployRouterScript \
    --rpc-url $RPC_URL --broadcast -vvvv && \
sleep 5 && \
forge script script/DeployRankingAddon.s.sol:DeployRankingAddonScript \
    --rpc-url $RPC_URL --broadcast -vvvv && \
sleep 10 && \
forge script script/TestScenario.s.sol:TestScenarioScript \
    --rpc-url $RPC_URL --broadcast -vvvv
```

> 💡 **Tip**: 모든 스크립트가 동일한 `PRIVATE_KEY`로 실행됩니다. 다른 계정으로 테스트하려면 별도로 실행하세요.

---

## 📊 체크리스트

### 배포 전
- [ ] 배포 계정에 테스트넷 Native 토큰 확보
- [ ] RPC_URL 설정
- [ ] `PRIVATE_KEY` 준비 (`.env` 파일)
- [ ] 필요한 환경변수 설정 (각 `.env` 파일)

### 배포 후
- [ ] 모든 컨트랙트 Etherscan Verify 완료
- [ ] WCROSS 작동 확인
- [ ] 스테이킹/출금 테스트
- [ ] Top 10 랭킹 확인
- [ ] 시즌 전환 테스트

---

## 🆘 빠른 트러블슈팅

### "Season not active" 에러
```bash
# 시즌 시작까지 대기
cast block-number --rpc-url $RPC_URL
# 시작 블록 도달 후 재시도
```

### "Insufficient balance" 에러
```bash
# 테스트넷 faucet에서 토큰 받기
# 잔액 확인
cast balance $YOUR_ADDRESS --rpc-url $RPC_URL
```

### Router 승인 실패
```bash
# Protocol admin인지 확인
cast call $STAKING_PROTOCOL_ADDRESS \
    "hasRole(bytes32,address)(bool)" \
    0x0000000000000000000000000000000000000000000000000000000000000000 \
    $YOUR_ADDRESS \
    --rpc-url $RPC_URL
```

---

## 📚 더 자세한 정보

- 전체 가이드: [TESTNET_DEPLOYMENT.md](../TESTNET_DEPLOYMENT.md)
- 스크립트 상세: 각 `.env` 파일 참고
- 프로젝트 문서: [docs/](../docs/)

---

## 🔐 Private Key 검증

스크립트 실행 전에 사용할 주소를 확인하세요:

```bash
# Private key로 주소 확인
cast wallet address --private-key $PRIVATE_KEY
# 예상 출력: 0xYourAddress...

# 잔액 확인
cast balance $(cast wallet address --private-key $PRIVATE_KEY) --rpc-url $RPC_URL
```

---

**Happy Testing! 🎉**

