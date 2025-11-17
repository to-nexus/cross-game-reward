# 보안 및 테스트

## 🛡️ 보안 메커니즘

### 7개 보안 계층

```
Layer 1: ReentrancyGuardTransient
Layer 2: SafeERC20
Layer 3: AccessControl
Layer 4: Pausable
Layer 5: UUPS
Layer 6: Reward Claim Recovery
Layer 7: Removed Token Handling
Layer 8: Custom Errors
Layer 9: Router Check
```

---

## 🔒 보안 기능 상세

### 1. ReentrancyGuardTransient

**보호 대상:**
- deposit, depositFor
- withdraw, withdrawFor
- claimRewards, claimReward

**특징:**
- Transient storage 사용 (EIP-1153)
- 99.5% gas 절약

### 2. SafeERC20

**적용:**
```solidity
using SafeERC20 for IERC20;

depositToken.safeTransferFrom(msg.sender, address(this), amount);
depositToken.safeTransfer(user, amount);
```

**보호:**
- 반환값 확인
- revert 처리
- 비표준 ERC20 대응

### 3. AccessControl

**CrossGameReward:**
```solidity
DEFAULT_ADMIN_ROLE      // 시스템 관리
POOL_MANAGER_ROLE       // 풀 생성/관리
```

**CrossGameRewardPool:**
```solidity
DEFAULT_ADMIN_ROLE      // 풀 관리
REWARD_MANAGER_ROLE     // 보상 토큰 관리
PAUSER_ROLE             // 긴급 정지
```

**특징:**
- 시간 지연 관리자 변경
- 2단계 변경 프로세스
- 안전한 권한 위임

### 4. Pausable

**적용 함수:**
- deposit, depositFor
- withdraw, withdrawFor
- claimRewards, claimReward

**사용:**
```solidity
// 긴급 정지
crossDeposit.setPoolActive(poolId, false);
// → pool.pause() 자동 호출

// 재개
crossDeposit.setPoolActive(poolId, true);
// → pool.unpause() 자동 호출
```

### 5. UUPS Upgradeable

**업그레이드 권한:**
```solidity
function _authorizeUpgrade(address newImplementation) 
    internal 
    override 
    onlyRole(DEFAULT_ADMIN_ROLE) 
{}
```

**Storage Gap:**
```solidity
uint[50] private __gap;  // CrossGameReward
uint[41] private __gap;  // CrossGameRewardPool
```

### 6. 보상 청구 복구 메커니즘

**배경:**
- 보상 토큰 전송 실패가 원금 출금을 막아서는 안 됨 (감사 지적 사항)
- `trySafeTransfer` 사용으로 전송 실패 시에도 withdraw 가능
- 전송 실패 후 withdraw하면 stored rewards는 유지되지만 `balances[user]`가 0이 되어 재청구 불가능한 문제 발생

**해결책:**

```solidity
// 1. 헬퍼 함수로 stored rewards 확인
function _hasStoredRewards(address user) private view returns (bool) {
    // active 및 removed 토큰의 stored rewards 체크
}

// 2. claimRewards() - balance 0이어도 stored rewards 있으면 청구 가능
function claimRewards() external {
    uint userBalance = balances[msg.sender];
    bool hasRewards = _hasStoredRewards(msg.sender);
    
    require(userBalance > 0 || hasRewards, CGRPNoDepositFound());
    
    // ⚡ 가스 최적화: balance가 0이면 sync/update 생략
    if (userBalance > 0) {
        _syncRewards();
        _updateRewards(msg.sender);
        _updateRemovedRewards(msg.sender);
    }
    
    _claimRewards(msg.sender);
    _claimRemovedRewards(msg.sender);
}

// 3. claimReward(token) - 특정 토큰 stored reward 체크
function claimReward(IERC20 token) external {
    uint userBalance = balances[msg.sender];
    uint storedReward = userRewards[msg.sender][token].rewards;
    
    require(userBalance > 0 || storedReward > 0, CGRPNoDepositFound());
    
    if (userBalance > 0) {
        if (_rewardTokenAddresses.contains(address(token))) _syncReward(token);
        _updateReward(token, msg.sender);
    }
    
    _claimReward(token, msg.sender);
}
```

**보안 장점:**
- ✅ **DoS 방지**: balance와 stored rewards 모두 0인 사용자 차단
- ✅ **무제한 슬롯 생성 방지**: 조건부 접근 제어 유지
- ✅ **가스 최적화**: balance 0일 때 불필요한 sync/update 생략
- ✅ **removed tokens 지원**: 모든 토큰 타입에 대해 복구 가능

**시나리오 예시:**
```solidity
// 1. 사용자 deposit 및 보상 적립
pool.deposit(10 ether);
// rewardToken 100 ether 적립

// 2. 보상 토큰 일시적 작동 중단 (예: 업그레이드, 버그)
rewardToken.setTransferShouldFail(true);

// 3. claim 실패 → 하지만 원금은 인출 가능
pool.claimReward(rewardToken); // ❌ 전송 실패, rewards 유지
pool.withdraw();               // ✅ 원금 인출 성공

// 4. 토큰 수정 완료
rewardToken.setTransferShouldFail(false);

// 5. balance 0 상태에서도 보상 청구 가능
pool.claimReward(rewardToken); // ✅ stored rewards 수령
```

**테스트 커버리지:**
- `CrossGameRewardPoolClaimRecovery.t.sol` (10개 테스트)
- 기본 복구, 다중 사용자, 엣지 케이스, 가스 최적화 검증

### 7. 제거된 보상 토큰 자동 정산

- 보상 토큰을 제거하면 주소가 `_removedRewardTokenAddresses`에 보관되고 활성 목록에서 제외됩니다.
- `_withdraw` 흐름은 `_updateRemovedRewards`와 `_claimRemovedRewards`를 호출해 제거된 토큰까지 자동 정산·지급합니다.
- 디파짓을 유지한 채 부분 청구하려면 기존과 동일하게 `claimReward`/`claimRewards`를 호출해야 하며, 이때는 활성 토큰만 동기화됩니다.
- 회귀 테스트 `testRemovedRewardTokenClaimedOnUndeposit`와 `testClaimRemovedRewardAfterUndepositDoesNotRevert`가 동작을 검증합니다.

### 8. Custom Errors

**장점:**
- 가스 절약 (~100-200 gas/호출)
- 타입 안전
- 명확한 에러 출처

**Naming Convention:**
```
CGR   - CrossGameReward
CGRP  - CrossGameRewardPool
CGRR  - CrossGameRewardRouter
WCROSS - WCROSS

예: CGRPNoDepositFound, CGRRInvalidAmount
```

### 9. Router Check

**CrossGameRewardPool:**
```solidity
function _checkDelegate(address account) internal view {
    require(account != address(0), CGRPCanNotZeroAddress());
    require(msg.sender == ICrossGameReward(crossDeposit).router(), CGRPOnlyRouter());
}
```

**적용 함수:**
- depositFor / withdrawFor
- claimRewardsFor / claimRewardFor

**WCROSS - WETH9 패턴:**
```solidity
function deposit() public payable {
    if (msg.value != 0) _mint(msg.sender, msg.value);
}
```

**특징:**
- Router 검사 제거 (누구나 사용 가능)
- WETH9 표준 준수
- ERC20 메커니즘으로 보호
- DEX 통합 용이

---

## 🧪 테스트 체계

### 테스트 구조

```
test/
├── WCROSS.t.sol                           (10개)
├── CrossGameReward.t.sol                     (33개)
├── CrossGameRewardRouter.t.sol               (39개) ← claim 테스트 추가
├── FullIntegration.t.sol                  (9개)
├── CrossGameRewardPoolDeposit.t.sol            (18개)
├── CrossGameRewardPoolRewards.t.sol            (18개)
├── CrossGameRewardPoolAdmin.t.sol              (24개)
├── CrossGameRewardPoolIntegration.t.sol        (11개)
├── CrossGameRewardPoolSecurity.t.sol           (21개)
├── CrossGameRewardPoolClaimRecovery.t.sol      (10개)
└── CrossGameRewardPoolStressTest.t.sol         (40개)

총 233개 테스트
```

### 테스트 카테고리

#### 기능 테스트 (Functional)
- 디파짓/언디파짓
- 보상 계산/분배
- 풀 생성/관리
- Router 기능

#### 통합 테스트 (Integration)
- 전체 사용자 여정
- 다중 풀 시나리오
- 실전 사용 패턴

#### 보안 테스트 (Security)
- 재진입 방어
- 권한 검증
- 불변성 체크
- 오버플로우 방지

---

## 🎯 주요 검증 사항

### 1. 보상 분배 정확성

```solidity
assertApproxEqAbs(userReward, expectedReward, 1 ether);
```

**검증:**
- 지분율에 따른 정확한 분배
- 총 보상 = 입금 보상
- 시간 독립성

### 2. 상태 일관성

```solidity
assertEq(pool.totalDeposited(), depositToken.balanceOf(address(pool)));
```

**검증:**
- totalDeposited == 실제 잔액
- 보상 토큰 잔액 일치

### 3. rewardPerToken 누적

```solidity
assertGe(newRewardPerToken, oldRewardPerToken);
```

**검증:**
- 증가만 함 (절대 감소 없음)
- 정확한 누적

### 4. Router 권한

```solidity
vm.expectRevert(CGRPOnlyRouter.selector);
pool.depositFor(user, amount);  // Non-router call
```

**검증:**
- Router만 depositFor/withdrawFor 호출 가능
- 권한 없는 접근 차단

---

## 🔬 테스트 커버리지

### 함수 커버리지

- **Line Coverage:** ~100%
- **Branch Coverage:** ~100%
- **Function Coverage:** 100%

### 시나리오 커버리지

- ✅ 단일 사용자
- ✅ 다중 사용자
- ✅ 다중 보상 토큰
- ✅ Native CROSS 플로우
- ✅ ERC20 플로우
- ✅ 스테이커 없을 때
- ✅ 긴급 정지
- ✅ 업그레이드

---

## 📈 Gas 벤치마크

### 주요 함수

| 함수 | Gas | 비고 |
|------|-----|------|
| deposit | ~143k | 기본 디파짓 |
| depositFor | ~145k | Router용 |
| withdraw | ~288k | 보상 포함 |
| depositNative | ~177k | 래핑 포함 |
| withdrawNative | ~235k | 언래핑 포함 |
| createPool | ~571k | 풀 생성 |

### 최적화 기법

- **EnumerableSet**: O(1) lookup
- **ReentrancyGuardTransient**: 99.5% gas ↓
- **Helper 함수**: 중복 코드 제거
- **Custom Errors**: ~100-200 gas ↓

---

## 🎓 테스트 작성 원칙

### 1. AAA 패턴

```solidity
function testExample() public {
    // Arrange
    _userDeposit(user1, 100 ether);
    
    // Act
    _depositReward(address(rewardToken), 1000 ether);
    
    // Assert
    uint[] memory rewards = pool.pendingRewards(user1);
    assertApproxEqAbs(rewards[0], 1000 ether, 1 ether);
}
```

### 2. 독립성

각 테스트는 독립적으로 실행 가능

### 3. 명확한 네이밍

```solidity
testDepositNativeMultipleTimes()
testRewardDistributionWithZeroDeposited()
```

### 4. Helper 함수

```solidity
_userDeposit(address user, uint amount)
_depositReward(address token, uint amount)
_warpDays(uint days_)
```

---

## 🏆 테스트 통계

```
총 테스트: 233개
성공률: 100%
실행 시간: ~0.11초
커버리지: ~100%
```

### 테스트 분포

```
WCROSS (10개):
  - 누구나 deposit/withdraw (WETH9 패턴)
  - Transfer 기능
  - Integration

CrossGameReward (33개):
  - 풀 생성/관리
  - Router 관리
  - View 함수
  - 업그레이드

CrossGameRewardRouter (39개):
  - Native 디파짓/출금
  - ERC20 디파짓/출금
  - Claim 래퍼 함수 (신규 12개)
  - EIP-2612 Permit
  - 에러 케이스

FullIntegration (9개):
  - 전체 플로우
  - 다중 풀
  - 보상 정확성

CrossGameRewardPool (142개):
  - 디파짓 (18개)
  - 보상 (18개)
  - 관리자 (24개)
  - 통합 (11개)
  - 보안 (21개)
  - 청구 복구 (10개)
  - 스트레스 테스트 (40개)
```

---

## 🔍 감사 가이드

### 확인 사항

1. **재진입 공격**
   - 모든 state-changing 함수에 nonReentrant
   - Checks-Effects-Interactions 패턴

2. **권한 검증**
   - AccessControl 적용
   - Router 전용 함수 체크

3. **정밀도 손실**
   - PRECISION = 1e18
   - 충분한 정밀도

4. **불변성**
   - totalDeposited == 실제 잔액
   - 보상 보존

5. **업그레이드**
   - Storage gap
   - _authorizeUpgrade

---

## 📚 참고 자료

### 테스트 도구

- [Foundry - Testing](https://book.getfoundry.sh/forge/tests)
- [Foundry - Cheatcodes](https://book.getfoundry.sh/cheatcodes/)

### 보안

- [OpenZeppelin Security](https://docs.openzeppelin.com/contracts/security)
- [Smart Contract Security Best Practices](https://consensys.github.io/smart-contract-best-practices/)

---

## 🗝️ 운영 및 거버넌스 주의
- `DEFAULT_ADMIN_ROLE` 보유자는 Router 교체, 새 풀 구현 지정, 업그레이드 승인 등 핵심 권한을 독점함 (`CrossGameReward`, `CrossGameRewardPool`). 멀티시그 또는 거버넌스 설계를 권장.
- `pause` 상태에서는 디파짓·언디파짓·클레임 모두 차단되므로, 긴급 상황에서 자금 인출 정책을 사전에 정의해야 함.

---

## ✨ 결론

**현재 상태 요약**
- ✅ Foundry 기반 233개 테스트 통과 (2025-11-17)
- ✅ OZ 기반 방어 계층·재진입 보호 적용
- ✅ Router claim 래퍼 함수 추가 (deposit 유지하면서 보상만 claim 가능)
- ✅ WCROSS WETH9 패턴 적용 (DEX 통합 용이)
- ✅ Pool claim 함수 리팩토링 (중복 코드 48% 감소)
- ✅ 제거된 보상 토큰은 언디파짓 시 자동 정산되어 미지급 위험 제거
- ✅ 보상 전송 실패 시 복구 메커니즘으로 원금 출금 및 재청구 보장

**다음**: [test/README.md](../test/README.md)
