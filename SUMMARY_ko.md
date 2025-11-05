# Cross Staking Protocol - 프로젝트 요약

## 🎯 프로젝트 개요

Cross Staking Protocol은 Native CROSS와 ERC-20 토큰을 동시에 지원하는 다중 풀 스테이킹 시스템입니다. Factory + Router 구조를 통해 새로운 풀을 빠르게 배포하고, 보상 토큰을 유연하게 구성할 수 있습니다.

### 핵심 가치

- ✅ **Native + ERC-20 지원**: Router가 WCROSS를 통해 자동 래핑/언래핑
- ✅ **확장성**: 하나의 팩토리에서 무제한 풀 생성
- ✅ **정확한 보상 분배**: `rewardPerToken` 누적 방식으로 O(1) 가스 비용
- ✅ **간소화된 권한 관리**: Owner/StakingRoot 기반 접근 제어
- ✅ **3단계 풀 관리**: Active/Inactive/Paused로 세밀한 제어
- ✅ **공정한 보상**: Zero-stake 보호 및 제거된 토큰 정산
- ✅ **개선된 API**: 토큰 주소와 보상을 한 번에 조회
- ✅ **안전성**: 7계층 보안 + UUPS 업그레이드

---

## 📐 시스템 구성

```
사용자                (Native CROSS / ERC-20)
   │
   ▼
CrossStakingRouter ──► WCROSS (래핑)
   │
   ▼
CrossStaking (UUPS) ──► CrossStakingPool × N (UUPS)
```

| 컴포넌트            | 역할                                                           |
|---------------------|----------------------------------------------------------------|
| CrossStaking        | 풀 생성/관리, 보상 토큰 추가, 풀 상태 설정, Router 지정       |
| CrossStakingPool    | 개별 스테이킹 풀, stake/unstake/claim, 3-state 관리           |
| CrossStakingRouter  | Native/ERC-20 스테이킹 인터페이스, WCROSS 자동 처리           |
| WCROSS              | Native CROSS ↔ ERC-20 변환, Router 전용 `deposit/withdraw`    |

---

## 🔄 주요 플로우

### 1. Native CROSS 스테이킹
1. 사용자는 Router에 WCROSS allowance 설정
2. `stakeNative` 호출 → Router가 WCROSS로 래핑 후 풀에 `stakeFor`
3. 언스테이킹 시 Router가 풀에서 `unstakeFor` 실행 후 WCROSS를 언래핑

### 2. ERC-20 스테이킹
1. 사용자가 스테이킹 토큰을 Router에 허용
2. Router가 토큰을 Pool로 전송 후 `stakeFor`
3. `unstakeERC20` 시 Router가 보상은 사용자에게, 원금은 다시 사용자에게 전송

### 3. 보상 분배 및 조회
- 누구나 Pool 주소로 보상 토큰을 전송할 수 있음
- 다음 스테이킹/언스테이킹/클레임 시 `_syncReward`가 자동 감지
- `rewardPerTokenStored`를 통해 가스 비용을 상수로 유지

**보상 조회 API:**
- `pendingRewards(user)`: 모든 활성 보상 토큰과 보상 조회 → `(address[] tokens, uint[] rewards)`
- `pendingReward(user, token)`: 특정 토큰의 보상 조회 → `uint amount`
- `getUserStakingInfo(poolId, user)`: 스테이킹 정보 통합 조회 → `(uint stakedAmount, address[] tokens, uint[] rewards)`

### 4. Zero-Stake 보호
- `totalStaked=0` 상태에서 예치된 보상은 `withdrawableAmount`로 분류
- 첫 스테이커가 이러한 보상을 받지 못하도록 보호
- Owner가 `CrossStaking.withdrawFromPool()`로 회수 가능

### 5. 제거된 보상 토큰 처리
- `removeRewardToken` 호출 시 해당 토큰의 분배 가능 보상은 `distributedAmount`로 고정
- 사용자는 `claimReward(removedToken)`으로 계속 청구 가능
- 제거 후 추가 예치된 토큰은 `withdrawableAmount`에 추가되어 owner가 회수

---

## 🎯 주요 기능

### 1. 3-State 풀 관리
- **Active (0)**: 모든 작업 허용 (stake, unstake, claim)
- **Inactive (1)**: unstake와 claim만 허용
- **Paused (2)**: 모든 작업 중지

설정: `CrossStaking.setPoolStatus(poolId, status)`

### 2. 보상 메커니즘
- **O(1) 가스**: `rewardPerToken` 누적 방식
- **Zero-stake 보호**: 첫 스테이커에게 불공정한 보상 방지
- **제거된 토큰**: `distributedAmount` (사용자 claim 가능) vs `withdrawableAmount` (owner 회수 가능)
- **정확성**: 수학적으로 보장된 비례 분배

### 3. 접근 제어
**CrossStaking:**
- `DEFAULT_ADMIN_ROLE` (owner): Router 지정, Pool Implementation 설정, 업그레이드
- `MANAGER_ROLE`: 풀 생성, 보상 토큰 관리, 풀 상태 변경, 보상 출금

**CrossStakingPool:**
- `onlyOwner()`: CrossStaking의 owner, 업그레이드 승인
- `onlyStakingRoot()`: CrossStaking 컨트랙트, 모든 관리 기능
- `onlyRouter`: Router 전용, `stakeFor/unstakeFor`

---

## 🔒 보안 계층

1. **ReentrancyGuardTransient (EIP-1153)** – 모든 state 변경 함수 보호
2. **SafeERC20** – 비표준 토큰 대응 및 안전한 전송
3. **간소화된 권한 관리** – Owner/StakingRoot modifier 기반
4. **3-State 풀 관리** – Active/Inactive/Paused로 세밀한 제어
5. **UUPS 업그레이드** – Owner만 업그레이드 가능
6. **Custom Errors** – 가스 절약 및 명확한 revert
7. **Router Check** – Router 전용 함수 접근 통제
8. **Zero-Stake 보호** – 불공정한 보상 방지
9. **이벤트 최적화** – 중복 이벤트 제거로 가스 절감

권장 운영 방안:
- Router 교체 / 업그레이드 키는 멀티시그로 보호
- `setPoolStatus(poolId, status)`: 0=Active, 1=Inactive, 2=Paused
- Zero-stake 예치 보상은 `withdrawFromPool`로 회수

---

## 🧪 테스트 & 품질

```bash
forge test                                  # 전체 테스트
forge test --match-contract CrossStaking   # 특정 컨트랙트
forge test --gas-report                    # 가스 리포트
```

### 테스트 통계 (Foundry)

| Suite                          | 통과 |
|--------------------------------|------|
| WCROSS                         | 10   |
| CrossStaking                   | 33   |
| CrossStakingRouter             | 28   |
| CrossStakingPoolStaking        | 18   |
| CrossStakingPoolRewards        | 27   |
| CrossStakingPoolAdmin          | 34   |
| CrossStakingPoolIntegration    | 11   |
| CrossStakingPoolPendingRewards | 9    |
| CrossStakingPoolSecurity       | 21   |
| CrossStakingPoolEdgeCases      | 12   |
| FullIntegration                | 9    |
| **총합**                       | **212** |

**Coverage:** ~100%, 주요 시나리오(멀티 풀, 제거 토큰, 라우터, Zero-stake, 엣지 케이스) 포함

### 최근 개선사항
1. ✅ API 개선: `pendingRewards()` 토큰 주소 포함 반환
2. ✅ `pendingReward()` 단일 토큰 조회 함수 추가
3. ✅ 이벤트 최적화: 중복 이벤트 제거
4. ✅ 6개 새로운 테스트 추가 (PendingRewards)
5. ✅ 전체 문서 최신화

---

## 📊 코드 품질

### 통계
- **컨트랙트**: 4개 (main) + 4개 (interfaces)
- **테스트 스위트**: 11개
- **총 테스트**: 212개 (100% 통과)
- **코드 라인**: ~3,500 (테스트 포함)
- **경고**: 0개
- **가스 최적화**: 이벤트 중복 제거, Custom errors

### 배포 준비 체크리스트
- ✅ 212개 테스트 100% 통과
- ✅ 컴파일 경고 0개
- ✅ 재진입 방지 및 역할 검증 완료
- ✅ Zero-stake 보호 구현
- ✅ 제거된 보상 토큰 정산 확인
- ✅ 업그레이드 경로 (`upgradeToAndCall`) 검증
- ✅ 문서 최신화 완료
- ✅ API 개선 완료
- ✅ 이벤트 최적화 완료
- [ ] 외부 감사 (권장)

---

## 🚀 주요 변경사항 (Latest)

### Breaking Changes
⚠️ 프론트엔드 마이그레이션 필요:

```solidity
// 이전
uint[] memory rewards = pool.pendingRewards(user);
address[] memory tokens = pool.getRewardTokens();
// 수동으로 매칭

// 현재
(address[] memory tokens, uint[] memory rewards) = pool.pendingRewards(user);
// 자동 매칭!

// 또는 특정 토큰만
uint amount = pool.pendingReward(user, specificToken);
```

### 새로운 기능
1. **Enhanced Reward Queries**: 토큰 주소와 보상을 한 번에 조회
2. **Event Optimization**: 중복 이벤트 제거로 가스 절감
3. **Comprehensive Testing**: 212개 테스트로 모든 엣지 케이스 커버

---

## 📚 문서 & 링크

- [README 한국어](README_ko.md)
- [README English](README.md)
- [Architecture](overview/ko/01_architecture.md)
- [Reward Mechanism](overview/ko/02_reward_mechanism.md)
- [Security & Testing](overview/ko/03_security_and_testing.md)
- [Test Guide](test/README.md)
- [OpenZeppelin Contracts](https://docs.openzeppelin.com/contracts/)
- [Foundry Book](https://book.getfoundry.sh/)

---

## 📜 라이선스

MIT

---

## ✨ 결론

Cross Staking Protocol은 **완전히 테스트되고 문서화된 프로덕션 준비 상태**입니다:
- 212개 테스트 100% 통과
- 개선된 API로 더 나은 UX
- 최적화된 이벤트 로깅
- 최신 문서 완비
- 제로 컴파일 경고

배포 준비 완료! 🎉
