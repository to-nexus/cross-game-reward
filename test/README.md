# Cross GameReward 테스트 가이드

## 📁 테스트 구조

```
test/
├── base/
│   └── CrossGameRewardPoolBase.t.sol          # 공통 Base 컨트랙트
├── mocks/
│   └── MockERC20.sol                       # 테스트용 ERC20 토큰
├── WCROSS.t.sol                            # WCROSS 테스트 (21개)
├── CrossGameReward.t.sol                      # CrossGameReward 테스트 (15개)
├── CrossGameRewardRouter.t.sol                # CrossGameRewardRouter 테스트 (15개)
├── FullIntegration.t.sol                   # 전체 통합 테스트 (9개)
├── CrossGameRewardPoolDeposit.t.sol           # 디파짓 테스트 (18개)
├── CrossGameRewardPoolRewards.t.sol           # 보상 테스트 (18개)
├── CrossGameRewardPoolAdmin.t.sol             # 관리자 테스트 (24개)
├── CrossGameRewardPoolIntegration.t.sol       # Pool 통합 테스트 (11개)
└── CrossGameRewardPoolSecurity.t.sol          # 보안 테스트 (21개)
```

---

## 🧪 테스트 스위트 (총 212개)

### 1. WCROSS Test (10개)

**테스트 대상:** Wrapped CROSS 토큰

#### 기본 기능
- `testDeposit` - Native CROSS 래핑 (Router만 가능)
- `testWithdraw` - WCROSS 언래핑 (Router만 가능)
- `testReceiveFunction` - receive() 자동 래핑
- `testWithdrawTo` - 지정된 주소로 언래핑
- `testTransferBetweenUsers` - 사용자 간 WCROSS 전송

#### 권한 검증
- `testCannotDepositWithoutRouter` - Router 외 예치 방지
- `testCannotWithdrawWithoutRouter` - Router 외 인출 방지

#### 에러 케이스
- `testCannotDepositZero` - 0 입금 방지
- `testCannotWithdrawMoreThanBalance` - 잔액 초과 방지

#### 이벤트
- `testDepositEvent`, `testWithdrawalEvent`

**참고**: WCROSS는 Router 전용 설계로, CrossGameReward.router()로 등록된 주소만 deposit/withdraw 가능합니다.

---

### 2. CrossGameReward Test (33개)

**테스트 대상:** 풀 팩토리 및 관리

#### 풀 생성
- `testCreatePool` - 기본 풀 생성
- `testCreateMultiplePools` - 다중 풀 생성
- `testMultiplePoolsWithSameDepositToken` - 같은 토큰 다중 풀
- `testCannotCreatePoolWithZeroAddress` - 검증

#### 풀 조회
- `testGetPoolInfo` - 풀 정보 조회
- `testGetPoolAddress` - 주소로 조회
- `testGetPoolId` - ID로 조회
- `testGetTotalPoolCount` - 전체 개수
- `testGetAllPoolIds` - 모든 ID
- `testGetPoolIdsByDepositToken` - 토큰별 풀 조회
- `testCannotGetNonExistentPool` - 존재하지 않는 풀

#### Router 관리
- `testSetRouter` - Router 주소 설정
- `testOnlyOwnerCanSetRouter` - 권한 검증

#### 통합
- `testPoolsAreIndependent` - 풀 독립성 검증

---

### 3. CrossGameRewardRouter Test (28개)

**테스트 대상:** 사용자 인터페이스 라우터

#### Native CROSS 디파짓
- `testDepositNative` - Native CROSS 디파짓
- `testDepositNativeMultipleTimes` - 다중 디파짓
- `testWithdrawNative` - Native CROSS 언디파짓
- `testCannotDepositNativeZero` - 0 방지
- `testCannotDepositNativeOnERC20Pool` - 풀 검증
- `testCannotWithdrawNativeWithoutDeposit` - 디파짓 없음

#### ERC20 디파짓
- `testDepositERC20` - ERC20 디파짓
- `testDepositERC20MultipleTimes` - 다중 디파짓
- `testWithdrawERC20` - ERC20 언디파짓
- `testCannotDepositERC20Zero` - 0 방지
- `testCannotWithdrawERC20WithoutDeposit` - 디파짓 없음

#### View 함수
- `testGetUserDepositInfo` - 사용자 정보 조회
- `testIsNativePool` - Native 풀 확인

#### 복잡한 시나리오
- `testMultiUserNativeDeposit` - 다중 사용자
- `testMixedPoolUsage` - 혼합 사용

---

### 4. Full Integration Test (9개)

**테스트 대상:** 전체 시스템 통합

#### 전체 사용자 여정
- `testCompleteUserJourney` - 완전한 사용 시나리오
- `testMultiplePoolsSimultaneously` - 다중 풀 동시 사용
- `testRealWorldScenario` - 실전 시나리오

#### 에지 케이스
- `testDepositWithdrawDeposit` - 재예치 시나리오
- `testMultipleRewardRounds` - 다중 보상 라운드

#### 보안 검증
- `testCannotWithdrawOthersDeposit` - 타인 자금 보호
- `testReentrancyProtection` - 재진입 방지

#### 일관성 검증
- `testRewardDistributionAccuracy` - 보상 분배 정확성
- `testViewFunctionsConsistency` - View 함수 일관성

---

### 5. CrossGameRewardPool Test (92개)

기존 CrossGameRewardPool 단위 테스트

#### Deposit (18개)
- 기본 기능, 금액 검증, 추가 기능, 상태 추적, 에러 케이스

#### Rewards (18개)
- 보상 계산, 청구, 다중 사용자, 다중 토큰, 직접 transfer 감지

#### Admin (24개)
- 보상 토큰 관리, Pause/Unpause, 역할 관리, 업그레이드

#### Integration (11개)
- 복잡한 시나리오, 실전 사용 패턴

#### Security (21개)
- 재진입, 오버플로우, 정밀도, 불변성, 시간 독립성

---

## 🚀 테스트 실행

### 전체 테스트 실행

```bash
forge test
```

**예상 결과:**
```
Test Suite                      | Passed | Failed
================================+========+========
WCROSS                         |   10   |   0
CrossGameReward                   |   33   |   0
CrossGameRewardRouter             |   28   |   0
FullIntegration                |    9   |   0
CrossGameRewardPoolDeposit        |   21   |   0
CrossGameRewardPoolRewards        |   27   |   0
CrossGameRewardPoolAdmin          |   34   |   0
CrossGameRewardPoolIntegration    |   11   |   0
CrossGameRewardPoolPendingRewards |    9   |   0
CrossGameRewardPoolSecurity       |   21   |   0
CrossGameRewardPoolEdgeCases      |   12   |   0
-----------------------------------+--------+--------
Total                          |  212   |   0
```

### 특정 스위트 실행

```bash
# WCROSS 테스트만
forge test --match-contract WCROSSTest

# CrossGameReward 테스트만
forge test --match-contract CrossGameRewardTest

# Router 테스트만
forge test --match-contract CrossGameRewardRouterTest

# 통합 테스트만
forge test --match-contract FullIntegrationTest

# Pool 테스트만
forge test --match-contract CrossGameRewardPool
```

### 특정 테스트 실행

```bash
# 함수명으로 검색
forge test --match-test testDepositNative

# Verbose 모드
forge test --match-test testCompleteUserJourney -vvv

# Gas 리포트
forge test --gas-report
```

### 커버리지 확인

```bash
forge coverage
```

---

## 📊 Helper 함수

### CrossGameRewardPoolBase

```solidity
// 사용자 디파짓
function _userDeposit(address user, uint amount) internal;

// 보상 입금 (직접 transfer)
function _depositReward(address rewardToken, uint amount) internal;

// 시간 이동
function _warpDays(uint days_) internal;
function _warpSeconds(uint seconds_) internal;
```

---

## 🎯 테스트 카테고리

### 기능 테스트 (Functional)
- WCROSS: 래핑/언래핑
- CrossGameReward: 풀 생성/관리
- CrossGameRewardRouter: 사용자 상호작용
- CrossGameRewardPool: 디파짓/보상

### 통합 테스트 (Integration)
- 전체 시스템 플로우
- 다중 사용자 시나리오
- 실전 사용 패턴

### 보안 테스트 (Security)
- 재진입 공격 방지
- 권한 검증
- 불변성 체크
- 오버플로우 방지

---

## 🔍 주요 검증 사항

### 1. 보상 분배 정확성
```solidity
// 지분율에 따른 정확한 분배
assertApproxEqAbs(userReward, expectedReward, 1 ether);
```

### 2. 상태 일관성
```solidity
// totalDeposited == 실제 잔액
assertEq(pool.totalDeposited(), depositToken.balanceOf(address(pool)));
```

### 3. rewardPerToken 누적
```solidity
// 증가만 함 (절대 감소 없음)
assertGe(newRewardPerToken, oldRewardPerToken);
```

### 4. Native CROSS 플로우
```solidity
// Native -> WCROSS -> Deposit -> Withdraw -> WCROSS -> Native
assertEq(userNativeBalance, expectedNativeBalance);
```

---

## 🐛 알려진 제약사항

### 1. 정밀도
- PRECISION = 1e18
- 매우 작은 보상(<1 wei)은 손실 가능

### 2. Gas 한계
- 보상 토큰 3-5개 권장 (이론상 무제한)

---

## 📚 추가 리소스

### Foundry 문서
- [Testing](https://book.getfoundry.sh/forge/tests)
- [Cheatcodes](https://book.getfoundry.sh/cheatcodes/)
- [Gas Snapshots](https://book.getfoundry.sh/forge/gas-snapshots)

### 프로젝트 문서
- [Overview](../overview/00_overview.md)
- [Architecture](../overview/01_architecture.md)
- [Reward Mechanism](../overview/02_reward_mechanism.md)
- [Security & Testing](../overview/03_security_and_testing.md)

---

## ✅ 테스트 체크리스트

배포 전 확인:

- 212/212 테스트 통과
- Gas 최적화 확인
- 커버리지 ~100%
- 보안 검증 완료
- 외부 감사 권장

---

## 🎓 테스트 작성 가이드

새 테스트 추가 시:

1. **적절한 파일 선택**
   - WCROSS 관련 → `WCROSS.t.sol`
   - 풀 관리 → `CrossGameReward.t.sol`
   - 사용자 상호작용 → `CrossGameRewardRouter.t.sol`
   - 전체 플로우 → `FullIntegration.t.sol`
   - Pool 기능 → `CrossGameRewardPool*.t.sol`

2. **Helper 함수 활용**
   - `CrossGameRewardPoolBase`의 helper 사용
   - 코드 중복 최소화

3. **명확한 테스트명**
   - `test<Action><Condition>` 형식
   - 예: `testDepositNativeMultipleTimes`

4. **충분한 검증**
   - 상태 변경 확인
   - 이벤트 발생 확인
   - 에러 케이스 확인

---

## 🔬 테스트 통계

- **총 테스트:** 212개
- **성공률:** 100%
- **커버리지:** ~100%
- **실행 시간:** ~2초
- **평균 Gas:** 최적화됨
