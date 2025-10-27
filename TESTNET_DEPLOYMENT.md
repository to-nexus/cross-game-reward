# 테스트넷 배포 및 테스트 가이드

## 📋 목차

1. [사전 준비](#사전-준비)
2. [배포 순서](#배포-순서)
3. [테스트 시나리오](#테스트-시나리오)
4. [검증 방법](#검증-방법)
5. [트러블슈팅](#트러블슈팅)

---

## 사전 준비

### 1. 환경 설정

```bash
# .env 파일 생성
cp .env.example .env

# 필수 환경변수 설정
RPC_URL=https://your-testnet-rpc-url
ETHERSCAN_API_KEY=your_etherscan_api_key
```

### 2. 테스트넷 준비물

- ✅ 테스트넷 Native 토큰 (가스비용)
- ✅ 배포자 계정 개인키
- ✅ 프로토콜 관리자 주소
- ✅ 테스터 계정 (최소 3개 권장)

### 3. 필요한 정보 체크리스트

```
[ ] RPC_URL
[ ] PRIVATE_KEY
[ ] PROTOCOL_ADMIN 주소
[ ] 현재 블록 번호
[ ] 테스트 계정 주소들
```

---

## 배포 순서

### Step 1: 전체 시스템 배포

**스크립트**: `Deploy.s.sol`

```bash
# 1. 환경변수 설정
cp script/Deploy.env .env.deploy
# .env.deploy 파일 편집
nano .env.deploy

# 2. 배포 실행
source .env.deploy
forge script script/Deploy.s.sol:DeployScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast \
    --verify \
    -vvvv

# 3. 배포 결과 저장
# 출력된 주소들을 메모장에 저장:
# - WCROSS_ADDRESS
# - STAKING_PROTOCOL_ADDRESS
```

**예상 출력**:
```
=== Deployment Summary ===
WCROSS: 0x1234...
StakingPoolCode: 0x5678...
RewardPoolCode: 0x9abc...
StakingProtocol: 0xdef0...
```

### Step 2: 프로젝트 생성

**스크립트**: `CreateProject.s.sol`

```bash
# 1. 환경변수 설정
cp script/CreateProject.env .env.project
nano .env.project

# 필수 설정:
# - STAKING_PROTOCOL_ADDRESS (Step 1 결과)
# - PROJECT_NAME
# - FIRST_SEASON_START_BLOCK (현재 블록 + 100)

# 2. 프로젝트 생성
source .env.project
forge script script/CreateProject.s.sol:CreateProjectScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast \
    -vvvv

# 3. 결과 저장
# - PROJECT_ID
# - STAKING_POOL_ADDRESS
# - REWARD_POOL_ADDRESS
```

**예상 출력**:
```
=== Project Created ===
Project ID: 1
StakingPool: 0xabcd...
RewardPool: 0xef12...
```

### Step 3: Router 배포

**스크립트**: `DeployRouter.s.sol`

```bash
# 1. 환경변수 설정
cp script/DeployRouter.env .env.router
nano .env.router

# 필수 설정:
# - WCROSS_ADDRESS
# - STAKING_PROTOCOL_ADDRESS
# - PROJECT_ID

# 2. Router 배포 및 승인
source .env.router
forge script script/DeployRouter.s.sol:DeployRouterScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast \
    -vvvv

# 3. 결과 저장
# - STAKING_ROUTER_ADDRESS
```

### Step 4: RankingAddon 배포 (선택)

**스크립트**: `DeployRankingAddon.s.sol`

```bash
# 1. 환경변수 설정
cp script/DeployRankingAddon.env .env.addon
nano .env.addon

# 필수 설정:
# - STAKING_PROTOCOL_ADDRESS
# - STAKING_POOL_ADDRESS
# - PROJECT_ID

# 2. Addon 배포 및 설정
source .env.addon
forge script script/DeployRankingAddon.s.sol:DeployRankingAddonScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast \
    -vvvv

# 3. 결과 저장
# - RANKING_ADDON_ADDRESS
```

---

## 테스트 시나리오

### 시나리오 1: 기본 스테이킹 플로우

#### 목표
- Router를 통한 Native CROSS 스테이킹 (자동 랩핑)
- 스테이킹 포지션 확인
- 포인트 누적 확인

#### 실행
```bash
# 1. 환경변수 설정
cp script/TestScenario.env .env.test
nano .env.test

# 기본 설정만 사용 (RANKING_ADDON, REWARD_TOKEN은 주석)
STAKE_AMOUNT=1000000000000000000  # 1 ether

# 2. 시나리오 실행
source .env.test
forge script script/TestScenario.s.sol:TestScenarioScript \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast \
    -vvvv
```

#### 예상 결과
```
[Scenario 1] Staking with Native CROSS
Native balance after: [감소한 잔액]

[Scenario 2] Checking Stake Position
Staked balance: 1000000000000000000
Current points: 0

[Scenario 5] Checking Points
User Points: [증가하는 포인트]
Staking Power: [증가하는 파워]
```

### 시나리오 2: Top 10 랭킹 테스트

#### 목표
- 여러 계정으로 스테이킹
- Top 10 랭커 확인
- 순위 변동 테스트

#### 실행 (3개 계정 필요)

```bash
# Account 1: 10 CROSS
STAKE_AMOUNT=10000000000000000000 \
forge script script/TestScenario.s.sol:TestScenarioScript \
    --rpc-url $RPC_URL \
    --private-key $ACCOUNT1_KEY \
    --broadcast -vvvv

# Account 2: 5 CROSS
STAKE_AMOUNT=5000000000000000000 \
forge script script/TestScenario.s.sol:TestScenarioScript \
    --rpc-url $RPC_URL \
    --private-key $ACCOUNT2_KEY \
    --broadcast -vvvv

# Account 3: 15 CROSS (최고 순위 진입)
STAKE_AMOUNT=15000000000000000000 \
forge script script/TestScenario.s.sol:TestScenarioScript \
    --rpc-url $RPC_URL \
    --private-key $ACCOUNT3_KEY \
    --broadcast -vvvv
```

#### 검증
```bash
# RankingAddon 조회 (cast 사용)
cast call $RANKING_ADDON_ADDRESS \
    "getTopRankers(uint256)(address[],uint256[])" \
    1 \
    --rpc-url $RPC_URL

# 예상: Account 3이 1위
```

### 시나리오 3: 시즌 전환 테스트

#### 목표
- 시즌 종료까지 대기
- 시즌 롤오버
- 이전 시즌 데이터 확인

#### 실행

```bash
# 1. 첫 시즌 종료까지 대기 (블록 확인)
cast block-number --rpc-url $RPC_URL

# 2. 시즌 종료 블록 계산
# FIRST_SEASON_START_BLOCK + SEASON_BLOCKS

# 3. 종료 블록 도달 후 롤오버
cast send $STAKING_POOL_ADDRESS \
    "rolloverSeason()" \
    --private-key $PRIVATE_KEY \
    --rpc-url $RPC_URL

# 4. 현재 시즌 확인
cast call $STAKING_POOL_ADDRESS \
    "currentSeason()(uint256)" \
    --rpc-url $RPC_URL
# 예상: 2
```

### 시나리오 4: 보상 분배 테스트

#### 사전 준비
- ERC20 테스트 토큰 배포
- 프로젝트 admin이 토큰 보유

#### 실행

```bash
# 1. 보상 토큰 승인
cast send $REWARD_TOKEN_ADDRESS \
    "approve(address,uint256)" \
    $REWARD_POOL_ADDRESS \
    10000000000000000000 \
    --private-key $ADMIN_KEY \
    --rpc-url $RPC_URL

# 2. 보상 예치
cast send $STAKING_PROTOCOL_ADDRESS \
    "fundProjectSeason(uint256,uint256,address,uint256)" \
    $PROJECT_ID \
    1 \
    $REWARD_TOKEN_ADDRESS \
    10000000000000000000 \
    --private-key $ADMIN_KEY \
    --rpc-url $RPC_URL

# 3. 시즌 종료 후 보상 청구
cast send $STAKING_POOL_ADDRESS \
    "claimSeason(uint256,address)" \
    1 \
    $REWARD_TOKEN_ADDRESS \
    --private-key $USER_KEY \
    --rpc-url $RPC_URL
```

### 시나리오 5: 출금 및 순위 변동

#### 실행

```bash
# 1. 현재 순위 확인
cast call $RANKING_ADDON_ADDRESS \
    "getUserRank(uint256,address)(uint256)" \
    1 \
    $USER_ADDRESS \
    --rpc-url $RPC_URL

# 2. 전액 출금
cast send $STAKING_ROUTER_ADDRESS \
    "unstake(uint256)" \
    $PROJECT_ID \
    --private-key $USER_KEY \
    --rpc-url $RPC_URL

# 3. 순위 재확인 (Top 10에서 제거되었는지)
cast call $RANKING_ADDON_ADDRESS \
    "isTopRanker(uint256,address)(bool)" \
    1 \
    $USER_ADDRESS \
    --rpc-url $RPC_URL
# 예상: false
```

---

## 검증 방법

### 1. Etherscan에서 확인

```
1. 각 컨트랙트 주소를 Etherscan에서 검색
2. Contract 탭 → Verify 완료 확인
3. Read Contract 기능으로 상태 확인
4. Write Contract로 직접 호출 테스트
```

### 2. Cast 명령어로 확인

```bash
# WCROSS 잔액 확인
cast call $WCROSS_ADDRESS \
    "balanceOf(address)(uint256)" \
    $USER_ADDRESS \
    --rpc-url $RPC_URL

# 스테이킹 잔액 확인
cast call $STAKING_POOL_ADDRESS \
    "userStakes(address)(uint256,uint256,uint256)" \
    $USER_ADDRESS \
    --rpc-url $RPC_URL

# 현재 시즌 확인
cast call $STAKING_POOL_ADDRESS \
    "currentSeason()(uint256)" \
    --rpc-url $RPC_URL

# Top 10 랭커 확인
cast call $RANKING_ADDON_ADDRESS \
    "getTopRankers(uint256)(address[],uint256[])" \
    1 \
    --rpc-url $RPC_URL

# 프로젝트 정보 확인
cast call $STAKING_PROTOCOL_ADDRESS \
    "getProject(uint256)((address,address,string,bool,uint256,address,address))" \
    1 \
    --rpc-url $RPC_URL
```

### 3. 이벤트 모니터링

```bash
# 스테이킹 이벤트 확인
cast logs \
    --from-block $START_BLOCK \
    --to-block latest \
    --address $STAKING_POOL_ADDRESS \
    "Staked(address,uint256)" \
    --rpc-url $RPC_URL

# 랭킹 업데이트 이벤트
cast logs \
    --from-block $START_BLOCK \
    --to-block latest \
    --address $RANKING_ADDON_ADDRESS \
    "TopRankerAdded(uint256,address,uint256,uint256)" \
    --rpc-url $RPC_URL
```

---

## 검증 체크리스트

### 배포 검증
- [ ] 모든 컨트랙트 Etherscan Verify 완료
- [ ] StakingProtocol의 projectCount > 0
- [ ] Router가 프로젝트에 승인됨
- [ ] RankingAddon이 활성화됨

### 기능 검증
- [ ] WCROSS wrapping/unwrapping 정상 작동
- [ ] 스테이킹/출금 정상 작동
- [ ] 포인트가 시간에 따라 증가
- [ ] Top 10 랭킹이 실시간 업데이트
- [ ] 시즌 롤오버 정상 작동
- [ ] 보상 예치 및 청구 정상 작동

### 보안 검증
- [ ] admin만 관리 함수 호출 가능
- [ ] 일반 유저는 자신의 데이터만 조회/수정
- [ ] RewardPool 재설정 불가능
- [ ] Reentrancy 공격 방어

---

## 트러블슈팅

### 문제 1: "Season not active" 에러

**원인**: 첫 시즌이 아직 시작되지 않음

**해결**:
```bash
# 현재 블록 확인
cast block-number --rpc-url $RPC_URL

# 시즌 시작 블록 확인
cast call $STAKING_POOL_ADDRESS \
    "nextSeasonStartBlock()(uint256)" \
    --rpc-url $RPC_URL

# 시작 블록까지 대기 또는 블록 시간 조정
```

### 문제 2: "Only project admin" 에러

**원인**: 권한이 없는 계정으로 관리 함수 호출

**해결**:
```bash
# 프로젝트 admin 확인
cast call $STAKING_PROTOCOL_ADDRESS \
    "projects(uint256)(address,address,string,bool,uint256,address,address)" \
    $PROJECT_ID \
    --rpc-url $RPC_URL

# admin 계정으로 다시 시도
```

### 문제 3: Router 승인 실패

**원인**: 프로토콜 admin이 아닌 계정으로 시도

**해결**:
```bash
# 프로토콜 admin 확인
cast call $STAKING_PROTOCOL_ADDRESS \
    "hasRole(bytes32,address)(bool)" \
    0x0000000000000000000000000000000000000000000000000000000000000000 \
    $YOUR_ADDRESS \
    --rpc-url $RPC_URL
```

### 문제 4: 랭킹 업데이트 안 됨

**원인**: RankingAddon이 제대로 연결되지 않음

**해결**:
```bash
# Addon 설정 확인
cast call $STAKING_POOL_ADDRESS \
    "stakingAddon()(address)" \
    --rpc-url $RPC_URL

# Addon 재설정 (프로젝트 admin)
cast send $STAKING_PROTOCOL_ADDRESS \
    "setPoolStakingAddon(uint256,address)" \
    $PROJECT_ID \
    $RANKING_ADDON_ADDRESS \
    --private-key $ADMIN_KEY \
    --rpc-url $RPC_URL
```

---

## 테스트넷별 추천 설정

### Ethereum Sepolia
```
RPC_URL=https://sepolia.infura.io/v3/YOUR_KEY
SEASON_BLOCKS=7200  # 약 1일 (12초/블록)
```

### Polygon Mumbai
```
RPC_URL=https://rpc-mumbai.maticvigil.com
SEASON_BLOCKS=43200  # 약 1일 (2초/블록)
```

### Arbitrum Sepolia
```
RPC_URL=https://sepolia-rollup.arbitrum.io/rpc
SEASON_BLOCKS=86400  # 약 1일 (1초/블록)
```

---

## 다음 단계

### 테스트 완료 후
1. ✅ 모든 시나리오 검증 완료
2. ✅ 가스 비용 측정 및 기록
3. ✅ 버그 리포트 작성
4. ✅ 외부 감사 준비

### 메인넷 배포 전 체크리스트
- [ ] 테스트넷 완전 검증
- [ ] 외부 감사 완료
- [ ] 가스 최적화 확인
- [ ] 문서 최종 검토
- [ ] 백업 계획 수립
- [ ] 긴급 중단 메커니즘 테스트

---

**Happy Testing! 🚀**

