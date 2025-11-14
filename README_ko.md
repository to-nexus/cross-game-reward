# Cross GameReward Protocol

Native CROSS와 ERC-20 토큰을 위한 다중 풀 디파짓 프로토콜

## 🎯 개요

Cross GameReward Protocol은 확장 가능한 다중 풀 아키텍처를 통해 다양한 토큰의 디파짓을 지원하는 시스템입니다.

### 핵심 구성요소

```
사용자 인터페이스
    ↓
CrossGameRewardRouter → WCROSS
    ↓
CrossGameReward (풀 관리)
    ↓
CrossGameRewardPool × n (개별 풀)
```

## ✨ 주요 특징

- ✅ **Native CROSS 지원**: Router가 자동으로 래핑/언래핑
- ✅ **다중 풀**: 동일 토큰으로도 여러 풀 생성 가능
- ✅ **다중 보상**: 풀당 여러 ERC-20 보상 토큰 지원
- ✅ **O(1) 보상 분배**: `rewardPerToken` 누적 방식
- ✅ **UUPS 업그레이드**: CrossGameReward 및 Pool 업그레이드 가능
- ✅ **간소화된 권한 관리**: Owner와 RewardRoot 기반 권한
- ✅ **3단계 풀 상태**: Active/Inactive/Paused로 세밀한 제어
- ✅ **공정한 보상 분배**: 디파짓 前 예치된 보상 자동 회수 가능
- ✅ **제거된 보상 자동 정산**: 출금 시 제거된 토큰 보상 자동 지급
- ✅ **보상 청구 복구**: 전송 실패 시에도 원금 출금 가능, 이후 재청구 지원

## 🚀 빠른 시작

### 설치

```bash
forge install
```

### 테스트

```bash
forge test
```

**현재: 222/222 테스트 통과**

### 배포 예시

```bash
forge script script/DeployFullSystem.s.sol:DeployFullSystem \
  --rpc-url <RPC_URL> \
  --private-key <PRIVATE_KEY> \
  --broadcast
```

## 💡 사용 예시

### 사용자: Native CROSS 디파짓

```solidity
// 1) 디파짓 (approve 불필요 - Router가 자동 래핑)
router.depositNative{value: 100 ether}(poolId);

// 2) 출금 + 보상
router.withdrawNative(poolId);
```

### 관리자: 풀 생성 및 보상 토큰 설정

```solidity
// Native CROSS 풀 생성
(uint poolId, ICrossGameRewardPool pool) =
    crossDeposit.createPool(IERC20(address(wcross)), 1 ether);

// 보상 토큰 추가
crossDeposit.addRewardToken(poolId, IERC20(address(usdt)));

// 보상 입금 (누구나 가능)
usdt.transfer(address(pool), 1000 ether);
```

## 🏗️ 아키텍처

### 핵심 컨트랙트

1. **WCROSS**
   - Native CROSS ↔ ERC-20 변환
   - Router 전용 `deposit` / `withdraw`

2. **CrossGameReward (UUPS)**
   - 풀 생성/관리 팩토리
   - `createPool`, `addRewardToken`, `setPoolStatus`, `withdrawFromPool`, `setRouter`

3. **CrossGameRewardPool (UUPS)**
   - 개별 디파짓 풀
   - `deposit`, `withdraw`, `claimRewards`, `claimReward`
   - 3단계 풀 상태: Active/Inactive/Paused
   - totalDeposited=0 시 예치된 보상은 자동으로 withdrawable로 처리
   - 제거된 보상 토큰은 `_removedRewardTokenAddresses`로 추적 후 `_withdraw` 시 자동 정산

4. **CrossGameRewardRouter**
   - 사용자 인터페이스
   - Native CROSS 및 일반 ERC-20 디파짓 지원
   - WCROSS 자동 래핑/언래핑

## 🔑 권한 모델

### CrossGameReward
| 역할                    | 기능                                              |
|------------------------|---------------------------------------------------|
| DEFAULT_ADMIN_ROLE (owner) | Router 지정, Pool Implementation 설정, 업그레이드 승인 |
| MANAGER_ROLE            | 풀 생성, 보상 토큰 추가/삭제, 풀 상태 변경, 보상 출금 |

### CrossGameRewardPool
| 함수 타입               | 권한                     | 설명                                    |
|------------------------|-------------------------|------------------------------------------|
| `onlyOwner()`          | CrossGameReward의 owner    | 업그레이드 승인                          |
| `onlyRewardRoot()`    | CrossGameReward 컨트랙트   | 보상 토큰 관리, 풀 상태 설정, 보상 출금  |
| `depositFor/withdrawFor`  | Router (검증됨)         | 사용자 대신 deposit/withdraw               |

**주요 변경사항:**
- AccessControlDefaultAdminRules 제거, 간소화된 modifier 기반 권한
- Pool의 모든 관리 기능은 CrossGameReward 컨트랙트를 통해서만 실행
- IERC5313 표준 준수 (`owner()` 함수)

## 📊 보상 메커니즘

### 기본 원리
- `rewardPerToken` 누적 방식을 사용하여 가스 비용을 O(1)로 유지
- 보상 입금은 누구나 가능 (ERC-20 `transfer`)
- 디파짓 중에는 `claimReward(token)` / `claimRewards()`를 통해 활성 토큰에 대한 보상만 수령

### 보상 조회
- `pendingRewards(user)`: 모든 활성 보상 토큰과 대기 중인 보상 반환 `(address[] tokens, uint[] rewards)`
- `pendingReward(user, token)`: 특정 토큰에 대한 대기 중인 보상 조회 `uint amount`

### Zero-deposit 보호
- `totalDeposited=0` 상태에서 예치된 보상은 `withdrawableAmount`로 분류
- 첫 디파짓터가 이러한 보상을 받지 않도록 보호
- Owner가 `CrossGameReward.withdrawFromPool()`로 회수 가능

### 제거된 토큰 정산
- 보상 토큰 제거 시점의 잔액은 `distributedAmount`로 고정
- 사용자는 제거된 토큰도 `claimReward(removedToken)`으로 계속 claim 가능
- 제거 후 새로 예치된 양은 `withdrawableAmount`에 추가되어 owner가 회수

## 🔒 보안

- ReentrancyGuardTransient (EIP-1153) 적용
- SafeERC20 기반 토큰 전송
- 간소화된 권한 관리 (Owner/RewardRoot)
- 3단계 풀 상태 제어 (Active/Inactive/Paused)
- UUPS 업그레이드 권한 제한
- Custom Errors로 가스 절감 및 명확한 revert
- Router 호출자 검증
- Zero-deposit 보상 보호
- 보상 전송 실패 시 복구 메커니즘 (trySafeTransfer + stored rewards)

## 📚 문서

- [Architecture](overview/ko/01_architecture.md)
- [Reward Mechanism](overview/ko/02_reward_mechanism.md)
- [Security & Testing](overview/ko/03_security_and_testing.md)
- [Test Guide](test/README.md)

## 🧪 테스트

- `forge test`
- `forge test --match-contract CrossGameReward`
- `forge test --gas-report`

**테스트 통계**

| Suite                         | Tests |
|-------------------------------|-------|
| WCROSS                        | 10    |
| CrossGameReward                  | 33    |
| CrossGameRewardRouter            | 28    |
| CrossGameRewardPoolDeposit       | 18    |
| CrossGameRewardPoolRewards       | 27    |
| CrossGameRewardPoolAdmin         | 34    |
| CrossGameRewardPoolIntegration   | 11    |
| CrossGameRewardPoolPendingRewards| 9     |
| CrossGameRewardPoolSecurity      | 21    |
| CrossGameRewardPoolEdgeCases     | 12    |
| FullIntegration               | 9     |
| **Total**                     | **212**|

## 🔄 업그레이드

```solidity
// CrossGameReward 업그레이드
CrossGameReward newImpl = new CrossGameReward();
crossDeposit.upgradeToAndCall(address(newImpl), "");

// Pool 업그레이드
CrossGameRewardPool newPoolImpl = new CrossGameRewardPool();
pool.upgradeToAndCall(address(newPoolImpl), "");

// Router 교체
CrossGameRewardRouter newRouter = new CrossGameRewardRouter(address(crossDeposit));
crossDeposit.setRouter(address(newRouter));
```

## ⚠️ 운영 시 고려 사항

- Router, 보상 토큰, 업그레이드는 멀티시그 등으로 보호 권장
- `setPoolStatus(poolId, status)`: 0=Active, 1=Inactive (claim/withdraw만), 2=Paused (모든 작업 중지)
- 제거된 보상 토큰은 사용자가 `claimReward(removedToken)`으로 개별 청구 가능
- Zero-deposit 상태에서 예치된 보상은 `withdrawFromPool`로 회수 가능

## 📜 라이선스

MIT

## 🔗 참고 링크

- [OpenZeppelin Contracts](https://docs.openzeppelin.com/contracts/)
- [Foundry Book](https://book.getfoundry.sh/)
- [Cross GameReward Protocol Docs](overview/README.md)
