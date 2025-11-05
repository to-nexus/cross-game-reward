# 보상 메커니즘

## 📐 rewardPerToken 누적 방식

### 핵심 개념

**문제:** N명의 사용자에게 보상을 O(1) 가스비로 분배하려면?

**해결책:** 누적 "토큰당 보상" 추적

```
누적 계산 방식:
  전역 rewardPerTokenStored 사용
  → O(1) 가스 비용
  → 사용자 수 무관
```

---

## 🧮 수학적 원리

### 기본 공식

#### 1. 보상 입금 시

```
if (totalStaked == 0) {
    withdrawableAmount += newReward  // 분배하지 않음
} else {
    // withdrawableAmount가 이미 있어도, 그건 lastBalance에 포함됨
    rewardPerTokenStored += (newReward × PRECISION) / totalStaked
}

PRECISION = 1e18
```

**Zero-Stake 보호:**
- totalStaked=0 일 때 예치된 보상은 분배하지 않고 `withdrawableAmount`로 분류
- owner가 `withdrawFromPool`을 통해 회수 가능
- 첫 번째 staker가 과거 보상을 독점하지 못하도록 방지

**보상 토큰 제거 후:**
- 제거 시점의 분배 가능한 보상은 `distributedAmount`에 저장 (사용자 claim 가능)
- `withdrawableAmount`는 그대로 유지 (owner 회수 가능)
- 제거 후 추가 예치된 토큰도 owner가 회수 가능

**의미:**
- "1개의 스테이킹 토큰이 받을 수 있는 누적 보상량"
- PRECISION으로 스케일업하여 정밀도 유지

**예시:**
```
상황: 100 토큰 스테이킹 중, 50 보상 입금
계산: rewardPerTokenStored += (50 × 1e18) / 100
     = 0.5 × 1e18
의미: 스테이킹 토큰 1개당 0.5 보상
```

#### 2. 사용자 보상 계산 시

```
earned = userBalance × (rewardPerTokenStored - userCheckpoint) / PRECISION
totalReward = storedRewards + earned
```

**변수:**
- `userBalance`: 사용자의 스테이킹 수량
- `rewardPerTokenStored`: 현재 누적 토큰당 보상
- `userCheckpoint`: 사용자가 마지막으로 정산한 시점의 값
- `storedRewards`: 이미 계산되어 저장된 보상

---

## 🔄 보상 분배 시뮬레이션

### 시나리오

```
Day 0: Alice 100 토큰 예치
  totalStaked = 100
  rewardPerTokenStored = 0
  Alice.checkpoint = 0

Day 3: 보상 100 입금
  rewardPerTokenStored = 0 + (100 × 1e18) / 100 = 1e18
  (Alice는 아직 업데이트 안함)

Day 5: Bob 100 토큰 예치
  Bob.checkpoint = 1e18
  totalStaked = 200

Day 10: 보상 100 입금
  rewardPerTokenStored = 1e18 + (100 × 1e18) / 200 = 1.5e18

Alice 보상 = 100 × (1.5e18 - 0) / 1e18 = 150
Bob 보상 = 100 × (1.5e18 - 1e18) / 1e18 = 50

결과: Alice 150, Bob 50 ✅
```

---

## 🎯 주요 특징

### 1. 예치 시점 이후 보상만

```
Alice가 Day 0에 예치
  → Day 3 이후의 보상만 받음
  → 과거 보상 소급 없음
```

### 2. 지분율에 따른 분배

```
Alice: 100/200 = 50%
Bob: 100/200 = 50%

보상 100 입금 시:
  Alice: 50
  Bob: 50
```

### 3. 시간 독립적

```
언제 보상을 수령하든
누적된 보상은 동일
```

---

## 💡 특수 케이스

### 스테이커 없을 때 보상 (Zero-Stake 보호)

```solidity
function _syncReward(IERC20 token) internal {
    if (totalStaked == 0) {
        rt.withdrawableAmount += newReward;  // 분배하지 않고 회수 가능하게 표시
        rt.lastBalance = currentBalance;
        return;
    }
    // ...
}
```

**동작 (현재 버전):**
1. totalStaked = 0일 때 보상 입금
2. 보상이 `withdrawableAmount`로 분류됨
3. 첫 스테이커는 이 보상을 받지 **못함**
4. Owner가 `CrossStaking.withdrawFromPool()`로 회수 가능

**예시:**
```
1. 풀 비어있음
2. 보상 1000 입금 → withdrawableAmount = 1000
3. Alice 스테이킹
4. Alice는 1000을 받지 못함 (공정한 분배)
5. 이후 100 보상 입금 → Alice만 100을 받음
6. Owner가 withdrawableAmount 1000을 회수 가능
```

**보상 조회:**
- `pendingRewards(user)`: 모든 활성 보상 토큰과 보상 조회 `(address[] tokens, uint[] rewards)`
- `pendingReward(user, token)`: 특정 토큰의 보상 조회 `uint amount`

---

## 🔍 불변성 (Invariants)

### 1. 보상 보존

```
전체 입금된 보상 = 분배된 보상 + 대기 중인 보상
```

### 2. 누적 보상 단조 증가

```
새 rewardPerTokenStored >= 이전 rewardPerTokenStored
```

### 3. 사용자 보상 정확성

```
사용자 보상 = Σ(사용자 지분율 × 각 입금 보상)
```

---

## 📊 Gas 효율성

### O(1) 복잡도

**보상 분배:**
```
사용자 수와 무관하게 일정한 가스
  - 1명: ~150k gas
  - 100명: ~150k gas
  - 10,000명: ~150k gas
```

**이유:**
```
개별 사용자마다 계산하지 않음
전역 rewardPerTokenStored만 업데이트
```

---

## 🎓 정리

**rewardPerToken 누적 방식은:**

- ✅ 효율적 (O(1) 가스)
- ✅ 공정함 (지분율 기반)
- ✅ 확장 가능 (무제한 사용자)
- ✅ 정확함 (수학적 보장)

**다음**: [03_security_and_testing.md](./03_security_and_testing.md)
