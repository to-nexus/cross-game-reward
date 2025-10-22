# 테스트 문서

Cross-Staking Protocol 테스트 케이스 및 커버리지 문서

## 테스트 개요

- **총 테스트 수**: 68개
- **통과율**: 100%
- **프레임워크**: Foundry Test
- **커버리지**: 핵심 로직 100%

## 테스트 구조

```
test/
├── BaseTest.sol          # 공통 테스트 헬퍼 및 유틸리티
├── Staking.t.sol         # 기본 스테이킹 기능 테스트 (9)
├── Season.t.sol          # 시즌 관리 테스트 (7)
├── Points.t.sol          # 포인트 계산 테스트 (9)
├── Rewards.t.sol         # 보상 분배 테스트 (8)
├── MultiPool.t.sol       # 다중 프로젝트 테스트 (6)
├── Advanced.t.sol        # 고급 시나리오 테스트 (8)
├── Integrated.t.sol      # 통합 테스트 (8)
└── Fuzz.t.sol            # Fuzz 테스트 (13)
```

## 테스트 실행

### 기본 테스트
```bash
forge test
```

### 가스 리포트 포함
```bash
forge test --gas-report
```

### 특정 테스트 실행
```bash
forge test --match-contract StakingTest
forge test --match-test test_BasicStake
```

### 상세 로그
```bash
forge test -vvv  # 매우 상세한 로그
```

### Fuzz 테스트
```bash
forge test --match-contract FuzzTest
```

## 테스트 케이스 상세

### 1. Staking.t.sol - 기본 스테이킹 기능 (9 tests)

#### test_BasicStake
- **목적**: 기본 스테이킹 동작 확인
- **검증**:
  - 토큰 전송 정상 동작
  - 스테이킹 수량 기록
  - 이벤트 발생 확인

#### test_MinimumStake
- **목적**: 최소 스테이킹 수량 검증
- **검증**:
  - MIN_STAKE (1 CROSS) 이상만 허용
  - 정확히 MIN_STAKE일 때 성공

#### test_MultipleStakes
- **목적**: 분할 스테이킹 기능 확인
- **검증**:
  - 여러 번 스테이킹 가능
  - 누적 합산 정상 동작
  - 포인트 누적 정확성

#### test_MultipleUsersStake
- **목적**: 다수 사용자 스테이킹
- **검증**:
  - 사용자별 독립적인 포지션 관리
  - totalStaked 정확성

#### test_WithdrawAll
- **목적**: 전액 출금 기능
- **검증**:
  - 토큰 반환
  - 포인트 몰수
  - 포지션 초기화

#### test_StakeAfterWithdraw
- **목적**: 출금 후 재스테이킹
- **검증**:
  - 출금 후 다시 스테이킹 가능
  - 이전 시즌 포인트는 유지

#### test_RouterUnstake
- **목적**: Router를 통한 Native CROSS 출금
- **검증**:
  - WCROSS → Native CROSS 변환
  - 올바른 수량 반환

#### test_RevertWhen_BelowMinimumStake
- **목적**: 최소 수량 미만 스테이킹 거부
- **검증**: StakingPoolBelowMinStake 에러 발생

#### test_RevertWhen_WithdrawWithoutStake
- **목적**: 잔액 없이 출금 시도 거부
- **검증**: StakingPoolNoPosition 에러 발생

---

### 2. Season.t.sol - 시즌 관리 (7 tests)

#### test_SeasonRollover
- **목적**: 기본 시즌 롤오버 기능
- **검증**:
  - currentSeason 증가
  - 새 시즌 정보 생성

#### test_PointsResetAfterRollover
- **목적**: 롤오버 후 포인트 관리
- **검증**:
  - 이전 시즌 포인트 스냅샷
  - 새 시즌 포인트 독립적으로 시작

#### test_TotalPointsSnapshot
- **목적**: 시즌별 총 포인트 스냅샷
- **검증**:
  - 롤오버 시 totalPoints 계산
  - 스냅샷 값 정확성

#### test_MultipleSeasons
- **목적**: 다중 시즌 진행
- **검증**:
  - 연속적인 시즌 롤오버
  - 각 시즌 데이터 독립성

#### test_StakePreservedAcrossSeasons
- **목적**: 시즌 간 스테이크 유지
- **검증**:
  - 롤오버 후에도 스테이크 유지
  - 자동 다음 시즌 참여

#### test_SeasonInfo
- **목적**: 시즌 정보 조회
- **검증**:
  - startBlock, endBlock 정확성
  - blocksElapsed 계산

#### test_RevertWhen_RolloverBeforeSeasonEnd
- **목적**: 시즌 종료 전 롤오버 거부
- **검증**: StakingPoolSeasonNotEnded 에러

---

### 3. Points.t.sol - 포인트 계산 (9 tests)

#### test_PointsCalculation
- **목적**: 기본 포인트 계산 로직
- **검증**:
  - points = (balance × time × PRECISION) / timeUnit
  - 시간에 비례하여 증가

#### test_PointsAccumulation
- **목적**: 포인트 누적
- **검증**:
  - 시간 경과에 따른 누적
  - 실시간 조회 정확성

#### test_PointsProportionalToStake
- **목적**: 스테이크에 비례한 포인트
- **검증**:
  - 스테이크가 2배면 포인트도 2배
  - 비율 정확성

#### test_PointsTimeUnitAffectsCalculation
- **목적**: 시간 단위 변경 효과
- **검증**:
  - timeUnit 변경 시 포인트 계산 변화
  - 설정값 반영 정확성

#### test_BlockTimeAffectsPoints
- **목적**: 블록 시간 변경 효과
- **검증**:
  - blockTime 변경 시 포인트 계산 변화
  - 다른 체인 대응

#### test_SnapshotPlusAdditionalPoints
- **목적**: 스냅샷 + 실시간 포인트
- **검증**:
  - 분할 스테이킹 시 스냅샷 동작
  - 누적 포인트 정확성

#### test_ExpectedSeasonPointsWithSnapshot
- **목적**: 시즌별 예상 포인트 조회
- **검증**:
  - 과거 시즌 스냅샷
  - 현재 시즌 실시간 계산

#### test_PointsResetOnWithdraw
- **목적**: 출금 시 포인트 몰수
- **검증**:
  - 현재 시즌 포인트 0으로 초기화
  - 이전 시즌은 유지

#### test_NoPointsWithoutStake
- **목적**: 스테이크 없이는 포인트 없음
- **검증**: 스테이크 0이면 포인트도 0

---

### 4. Rewards.t.sol - 보상 분배 (8 tests)

#### test_FundSeason
- **목적**: 시즌 보상 예치
- **검증**:
  - 토큰 전송
  - 보상 풀 잔액 증가

#### test_EqualStakeEqualReward
- **목적**: 동일 스테이크 → 동일 보상
- **검증**:
  - 포인트 비율 1:1
  - 보상 비율 1:1

#### test_ProportionalRewardDistribution
- **목적**: 비례 보상 분배
- **검증**:
  - 포인트 비율에 따른 보상
  - 정확한 비율 계산

#### test_MultipleSeasonRewards
- **목적**: 다중 시즌 보상
- **검증**:
  - 시즌별 독립적인 보상
  - 각 시즌 청구 가능

#### test_NoRewardForNoStake
- **목적**: 스테이크 없으면 보상 없음
- **검증**:
  - 포인트 0이면 보상 0
  - 청구 시도 시 아무 일도 일어나지 않음

#### test_RemainingRewards
- **목적**: 미청구 보상 회수
- **검증**:
  - recoverRemaining 동작
  - 크리에이터에게 반환

#### test_RevertWhen_ClaimBeforeSeasonEnd
- **목적**: 시즌 종료 전 청구 거부
- **검증**: StakingPoolSeasonNotEnded 에러

#### test_RevertWhen_DuplicateClaim
- **목적**: 중복 청구 거부
- **검증**: StakingPoolAlreadyClaimed 에러

---

### 5. MultiPool.t.sol - 다중 프로젝트 (6 tests)

#### test_MultiplePoolsCreated
- **목적**: 여러 프로젝트 생성
- **검증**:
  - 각 프로젝트 독립적인 풀
  - projectCount 증가

#### test_StakeInDifferentPools
- **목적**: 여러 풀에 동시 스테이킹
- **검증**:
  - 사용자가 여러 프로젝트 참여 가능
  - 각 포지션 독립적

#### test_IndependentSeasons
- **목적**: 프로젝트별 독립적인 시즌
- **검증**:
  - 시즌 스케줄 독립성
  - 롤오버 독립성

#### test_IndependentRewards
- **목적**: 프로젝트별 독립적인 보상
- **검증**:
  - 각 RewardPool 독립성
  - 보상 분배 독립성

#### test_SameUserMultiplePools
- **목적**: 동일 사용자 다중 프로젝트
- **검증**:
  - 각 프로젝트에서 포인트 누적
  - 각 프로젝트에서 보상 청구

#### test_DifferentSeasonLengths
- **목적**: 프로젝트별 다른 시즌 길이
- **검증**:
  - seasonBlocks 설정 독립성
  - 각각 올바르게 동작

---

### 6. Advanced.t.sol - 고급 시나리오 (8 tests)

#### test_VirtualToRealSeasonTransition
- **목적**: Virtual 시즌에서 실제 시즌으로 전환
- **검증**:
  - 첫 시즌 lazy 생성
  - 자동 참여 메커니즘

#### test_MultiSeasonAutoParticipation
- **목적**: 자동 시즌 참여
- **검증**:
  - 스테이크 유지 시 자동 다음 시즌 참여
  - 중간에 액션 없어도 포인트 누적

#### test_ZeroBalanceUserInRollover
- **목적**: 잔액 0인 사용자의 롤오버 처리
- **검증**:
  - 에러 없이 처리
  - 불필요한 가스비 없음

#### test_SeasonGapBetweenRollovers
- **목적**: 롤오버 간 시간 간격
- **검증**:
  - 다음 시즌 시작 블록 설정
  - 지연된 시작 지원

#### test_PoolEndBlockAtSeasonBoundary
- **목적**: 풀 종료 블록이 시즌 경계일 때
- **검증**:
  - 정확한 종료 처리
  - 마지막 시즌 포인트 계산

#### test_PoolRestartAfterEnd
- **목적**: 풀 종료 후 재시작
- **검증**:
  - setNextSeasonStart로 재시작
  - 새로운 시즌 시작

#### test_ClaimAfterSeasonEnd
- **목적**: 시즌 종료 후 청구
- **검증**:
  - Lazy snapshot 동작
  - 정확한 포인트 계산

#### test_AdminEmergencyFunctions
- **목적**: 관리자 긴급 기능
- **검증**:
  - 설정 변경 권한
  - AccessControl 동작

---

### 7. Integrated.t.sol - 통합 테스트 (8 tests)

#### test_ProtocolCompleteFlow
- **목적**: 전체 프로토콜 흐름
- **검증**:
  - 프로젝트 생성 → 스테이킹 → 시즌 → 보상
  - 모든 단계 정상 동작

#### test_CompleteStakingFlow
- **목적**: 완전한 스테이킹 사이클
- **검증**:
  - stake → points → claim → withdraw
  - 전체 라이프사이클

#### test_MultiUserMultiSeasonFlow
- **목적**: 다수 사용자 다중 시즌
- **검증**:
  - 복잡한 시나리오 처리
  - 정확한 포인트 및 보상 분배

#### test_MultiSeasonFlow
- **목적**: 여러 시즌 진행
- **검증**:
  - 시즌 간 전환
  - 각 시즌 독립성

#### test_StakeWithdrawRestake
- **목적**: 스테이크 → 출금 → 재스테이크
- **검증**:
  - 출금 후 재참여 가능
  - 이전 데이터 정리

#### test_FullNativeTokenFlow
- **목적**: Native CROSS 전체 흐름
- **검증**:
  - Router를 통한 Native → WCROSS
  - 출금 시 WCROSS → Native

#### test_ProtocolCreateProject
- **목적**: Protocol을 통한 프로젝트 생성
- **검증**:
  - Factory 패턴 동작
  - 풀 연결 정상

#### test_EdgeCase_ClaimWithoutPoints
- **목적**: 포인트 없이 청구 시도
- **검증**:
  - 에러 없이 처리
  - 아무 일도 일어나지 않음

---

### 8. Fuzz.t.sol - Fuzz 테스트 (13 tests)

Fuzz 테스트는 랜덤 입력값으로 267회씩 실행됩니다.

#### testFuzz_PointsCalculation
- **목적**: 다양한 입력값으로 포인트 계산
- **검증**: 오버플로우, 언더플로우 없음

#### testFuzz_RewardDistribution
- **목적**: 랜덤 스테이크로 보상 분배
- **검증**: 비율 정확성, 반올림 오류 없음

#### testFuzz_MultipleSeasonRewards
- **목적**: 랜덤 시즌 진행
- **검증**: 시즌 간 간섭 없음

#### testFuzz_IncrementalStaking
- **목적**: 랜덤 분할 스테이킹
- **검증**: 누적 정확성

#### testFuzz_SeasonRollover
- **목적**: 랜덤 타이밍 롤오버
- **검증**: 언제 rollover해도 정상 동작

#### testFuzz_LargeStake
- **목적**: 큰 수량 스테이킹
- **검증**: 오버플로우 방지

#### testFuzz_LongDuration
- **목적**: 긴 기간 스테이킹
- **검증**: 시간 계산 정확성

#### testFuzz_MinimumStake
- **목적**: MIN_STAKE 근처 값
- **검증**: 경계값 처리

#### testFuzz_MultiUserPointsRatio
- **목적**: 다수 사용자 포인트 비율
- **검증**: 비율 계산 정확성

#### testFuzz_PointsAcrossSeasons
- **목적**: 여러 시즌 포인트 추적
- **검증**: 시즌 간 데이터 무결성

#### testFuzz_TimeParameters
- **목적**: 랜덤 시간 파라미터
- **검증**: timeUnit, blockTime 변경 정상 처리

#### testFuzz_ManyIncrementalStakes
- **목적**: 많은 분할 스테이킹
- **검증**: 가스비, 정확성

#### testFuzz_ZeroAmount
- **목적**: 0 수량 처리
- **검증**: 에러 또는 정상 거부

## 커버리지

### 핵심 기능 커버리지: 100%
- ✅ 스테이킹/출금
- ✅ 포인트 계산
- ✅ 시즌 관리
- ✅ 보상 분배
- ✅ 권한 관리
- ✅ Router 기능

### 엣지 케이스 커버리지
- ✅ 최소/최대 값
- ✅ 시즌 경계
- ✅ 중복 액션
- ✅ 권한 없는 호출
- ✅ 0 잔액 사용자
- ✅ 긴급 상황

## 가스 리포트

주요 함수 가스비 (평균):

| Function | Gas |
|----------|-----|
| stake | ~280,000 |
| withdrawAll | ~280,000 |
| rolloverSeason | ~120,000 |
| claimSeason | ~600,000 (첫 청구자, totalPoints 계산 포함) |
| claimSeason | ~100,000 (이후 청구자) |

## 테스트 실행 시간

- **전체 테스트**: ~520ms (CPU time)
- **Fuzz 테스트**: ~580ms (CPU time)
- **총**: ~1초

## 지속적 통합 (CI)

GitHub Actions 설정 예시:

```yaml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
        with:
          submodules: recursive

      - name: Install Foundry
        uses: foundry-rs/foundry-toolchain@v1

      - name: Run tests
        run: forge test --gas-report

      - name: Run coverage
        run: forge coverage
```

## 추가 테스트 권장사항

### 1. Invariant 테스트
```solidity
// 항상 유지되어야 하는 불변식
// - totalStaked = sum(all user balances)
// - season totalPoints = sum(all user points)
```

### 2. 시간 기반 테스트
- 실제 블록 시간 시뮬레이션
- 장기간 스테이킹 시나리오

### 3. 다중 체인 테스트
- 다른 blockTime 설정
- 다른 체인별 특성 고려

### 4. 부하 테스트
- 수천 명의 사용자 시뮬레이션
- 가스비 프로필링

## 버그 리포트

테스트 실패 시 다음 정보를 포함하여 리포트해주세요:
- 실패한 테스트 이름
- 에러 메시지
- 재현 단계
- 환경 정보 (Foundry 버전, 체인 등)

## 참고 자료

- [Foundry Book](https://book.getfoundry.sh/)
- [Forge Testing Guide](https://book.getfoundry.sh/forge/tests)
- [Fuzz Testing](https://book.getfoundry.sh/forge/fuzz-testing)
