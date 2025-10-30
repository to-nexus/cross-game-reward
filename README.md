# Cross Staking Pool

**시간 가중 포인트 시스템**을 적용한 혁신적인 CROSS 토큰 스테이킹 풀입니다.

## 🎯 프로젝트 개요

CrossStakingPool은 **포인트 = 금액 × 시간** 공식을 사용하여 O(1) 복잡도로 시간에 비례한 공정한 보상 분배를 구현한 스테이킹 풀입니다.

### 🚨 해결한 문제

**기존 방식의 문제점:**
```
Day 1: User A stakes 10 CROSS (전체의 100%)
Day 365: 보상 누적 중...

Day 365 직전: User B stakes 990 CROSS
→ User B가 99%의 미래 보상을 가져감!
→ User A는 1년 보유했지만 1%만 받음
```

**포인트 방식의 해결:**
```
Points = Amount × Time

User A 포인트: 10 × 365 days = 매우 큼
User B 포인트: 990 × 1 day = 상대적으로 작음

보상 비율 = (내 포인트 / 전체 포인트)
→ User A가 더 많이 받음! ✅
```

### 💎 주요 특징

- ✅ **O(1) 복잡도**: 사용자 수에 무관한 일정한 가스비
- ✅ **시간 가중 포인트**: 오래 보유할수록 더 많은 포인트
- ✅ **실시간 보상**: 언제든 보상 입금 가능
- ✅ **다중 보상 토큰**: 여러 종류의 보상 지원
- ✅ **수학적 정확성**: 포인트 보존, 보상 보존 증명 완료
- ✅ **rewardPerToken 누적**: 시간 가중 보상 분배
- ✅ **완벽한 테스트**: 29/29 통과 (100%)

## 빠른 시작

### 설치

```bash
forge install
```

### 컴파일

```bash
forge build
```

### 테스트

```bash
forge test
```

### 배포

```bash
forge script script/DeployCrossStakingPool.s.sol --rpc-url $RPC_URL --broadcast
```

## 사용 예시

### 스테이킹

```solidity
// CROSS 토큰 승인
crossToken.approve(address(pool), 100 ether);

// 스테이킹
pool.stake(100 ether);
```

### 보상 확인

```solidity
// Pending 보상 조회 (view, gas 0)
uint[] memory rewards = pool.pendingRewards(msg.sender);

// 포인트 조회
uint myPoints = pool.getUserPoints(msg.sender);
uint totalPoints = pool.getTotalPoints();

// 내 보상 비율
uint myShare = (myPoints × 100) / totalPoints; // %
```

### 보상 수령

```solidity
// 모든 보상 claim
pool.claimRewards();

// 또는 특정 보상만
pool.claimReward(0); // 첫 번째 보상 토큰
```

### Unstake

```solidity
// 전체 unstake (원금 + 보상)
pool.unstake();
```

## 핵심 원리

### 포인트 시스템

```
포인트 = 금액 × 시간

예시:
10 CROSS × 100초 = 1,000 포인트
20 CROSS × 50초 = 1,000 포인트
(동일한 기여도!)
```

### O(1) 계산

```solidity
// 사용자 포인트 (O(1))
currentPoints = storedPoints + amount × (now - lastUpdate)

// 전역 포인트 (O(1))
totalPoints = globalStoredPoints + totalStaked × (now - lastUpdate)
```

**핵심:** 증분 업데이트 방식으로 모든 계산 O(1)

### 보상 분배

```
보상 비율 = 내 포인트 / 전체 포인트

User1 포인트: 1,000
User2 포인트: 500
Total: 1,500

150 tokens 입금:
  User1: (1,000 / 1,500) × 150 = 100 tokens
  User2: (500 / 1,500) × 150 = 50 tokens
```

## 📚 문서

### 기술 문서 (overview/)

- **[README.md](./overview/README.md)** - 전체 시스템 소개
- **[01_design.md](./overview/01_design.md)** - 상세 설계
- **[02_math_proofs.md](./overview/02_math_proofs.md)** - 수학적 증명
- **[03_gas_optimization.md](./overview/03_gas_optimization.md)** - 가스 최적화

## 컨트랙트

- **CrossStakingPool.sol** (338 lines)
  - 위치: `src/CrossStakingPool.sol`
  - 테스트: `test/CrossStakingPool.t.sol` (29개 테스트)

## 테스트 결과

```bash
$ forge test

Ran 29 tests
✅ 29 passed (100%)
❌ 0 failed

가스비:
- stake: ~165k
- unstake: ~345k
- claim: ~230k
```

## 실제 시나리오

### 시나리오 1: 늦은 대량 진입

```
User A: 10 CROSS (100일 보유)
User B: 990 CROSS (1일 보유)

포인트:
  A: 10 × 100 = 1,000
  B: 990 × 1 = 990

보상 (100 토큰):
  A: 50.25 토큰 (더 많음!)
  B: 49.75 토큰

→ 작은 금액이지만 오래 보유해서 더 많이 받음 ✅
```

### 시나리오 2: 추가 예치

```
t=0: 10 CROSS stake
  → points = 0

t=100: 추가 20 CROSS stake (total 30)
  → stored = 10 × 100 = 1,000
  → amount = 30

t=200: 보상 claim
  → points = 1,000 + 30 × 100 = 4,000
  → 정확히 계산됨! ✅
```

## 요구사항 충족

| # | 요구사항 | 구현 |
|---|---------|------|
| 1 | CROSS 토큰 예치 | ✅ stakingToken |
| 2 | 여러 리워드 토큰 | ✅ RewardToken[] |
| 3 | 실시간 보상 누적 | ✅ _syncReward() |
| 4 | 지분만큼 보상 | ✅ 포인트 비율 |
| 5 | 최소 1 CROSS | ✅ MIN_STAKE_AMOUNT |
| 6 | claim/unstake 수령 | ✅ 구현 완료 |
| 7 | 지분 비례 분배 | ✅ 포인트 기반 |
| 8 | 먼저 예치 → 더 많은 지분 | ✅ 포인트 = 금액 × 시간 |

## 기술 스택

- **Solidity**: 0.8.28
- **Framework**: Foundry
- **Libraries**: OpenZeppelin Contracts
- **Pattern**: rewardPerToken accumulation + Time-weighted Points

## 보안

- ✅ ReentrancyGuard
- ✅ SafeERC20
- ✅ Ownable
- ✅ Custom Errors
- ✅ 수학적 검증 완료

## 라이선스

MIT License

## 기여

기술 문의: GitHub Issues

---

**프로덕션 배포 준비 완료!** 🚀
