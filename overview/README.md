# CrossStakingPool 기술 문서

## 📚 문서 구성

이 폴더는 CrossStakingPool 컨트랙트의 상세 기술 문서를 포함합니다.

### 문서 목록

1. **[01_architecture.md](./01_architecture.md)** - 아키텍처 및 설계
2. **[02_reward_mechanism.md](./02_reward_mechanism.md)** - 보상 메커니즘 상세
3. **[03_security_and_testing.md](./03_security_and_testing.md)** - 보안 및 테스트

---

## 📖 빠른 시작

### 핵심 개념 이해

CrossStakingPool은 **rewardPerToken 누적 방식**을 사용합니다:

```
핵심 공식:
rewardPerTokenStored += (newReward × 1e18) / totalStaked
userReward = userBalance × (rewardPerTokenStored - rewardPerTokenPaid) / 1e18
```

**특징:**
- ✅ O(1) 가스 비용 (사용자 수와 무관)
- ✅ 예치 비율에 따른 공정한 분배
- ✅ 예치 이후 보상만 수령

### 주요 기능

| 기능 | 설명 | 문서 |
|------|------|------|
| **Staking** | CROSS 토큰 예치 | [Architecture](./01_architecture.md#-핵심-플로우) |
| **Rewards** | 실시간 보상 분배 | [Reward Mechanism](./02_reward_mechanism.md) |
| **Security** | 다층 보안 시스템 | [Security](./03_security_and_testing.md) |
| **Upgradeability** | UUPS 업그레이드 | [Architecture](./01_architecture.md#3-uups-업그레이더블) |

---

## 🎯 문서 사용 가이드

### 개발자용

**처음 시작:**
1. [Architecture](./01_architecture.md) 읽기
2. [Reward Mechanism](./02_reward_mechanism.md) 수학 이해
3. 테스트 코드 실행

**통합 개발:**
1. [Architecture - 사용 시나리오](./01_architecture.md#-사용-시나리오) 참고
2. 배포 가이드 따라하기

### 감사자(Auditor)용

**보안 검토:**
1. [Security](./03_security_and_testing.md#️-보안-메커니즘) 확인
2. [Testing](./03_security_and_testing.md#-테스트-체계) 커버리지 검토
3. [Reward Mechanism](./02_reward_mechanism.md#-불변성-invariants) 불변성 검증

**수학적 검증:**
1. [Reward Mechanism - 수학적 증명](./02_reward_mechanism.md#-수학적-증명)
2. [Simulation](./02_reward_mechanism.md#-보상-분배-시뮬레이션)

### 사용자용

**이해하기:**
1. [Architecture - 사용 시나리오](./01_architecture.md#-사용-시나리오)
2. [Reward Mechanism - 기본 공식](./02_reward_mechanism.md#-수학적-원리)

---

## 📊 프로젝트 통계

### 컨트랙트

- **Lines of Code:** 408
- **Functions:** 19
- **Events:** 6
- **Roles:** 3

### 테스트

- **Test Files:** 6
- **Test Cases:** 93
- **Success Rate:** 100%
- **Coverage:** ~100%

### 보안

- **Audits:** 진행 예정
- **Known Issues:** 0
- **Fixed Issues:** 2
  - `_updateCheckpoints` 중복 제거
  - CROSS를 보상 토큰으로 사용 방지

---

## 🔗 관련 링크

### 프로젝트

- **Repository:** [GitHub](https://github.com/to-nexus/cross-staking)
- **Tests:** [../test/README.md](../test/README.md)

### 참고 구현

- **Synthetix:** [StakingRewards.sol](https://github.com/Synthetixio/synthetix)
- **OpenZeppelin:** [Contracts v5.4.0](https://docs.openzeppelin.com/contracts/5.x/)

### 표준

- **EIP-1967:** Proxy Storage Slots
- **EIP-1153:** Transient Storage
- **ERC-20:** Token Standard

---

## 📝 문서 업데이트

**최종 업데이트:** 2025년 10월 30일

**버전:**
- 컨트랙트: v1.0
- 문서: v1.0

**기여자:**
- 초기 설계 및 구현
- 보안 검토 및 최적화
- 문서 작성

---

## 💡 질문 및 지원

### FAQ

**Q: 왜 rewardPerToken 누적 방식을 사용하나요?**  
A: O(1) 가스 비용으로 무제한 사용자 지원 가능

**Q: 다중 보상 토큰은 몇 개까지?**  
A: 이론적으로 무제한, 실전에서는 3-5개 권장

**Q: UUPS vs Transparent Proxy?**  
A: UUPS가 가스비 저렴하고 더 안전

**Q: 직접 transfer해도 작동하나요?**  
A: 네, 자동으로 감지되지만 `depositReward` 사용 권장

**Q: 시간 Lock이 있나요?**  
A: 현재 없음, 업그레이드로 추가 가능

### 추가 정보

더 자세한 내용은 각 문서를 참고하세요:
- 설계: [01_architecture.md](./01_architecture.md)
- 보상: [02_reward_mechanism.md](./02_reward_mechanism.md)
- 보안: [03_security_and_testing.md](./03_security_and_testing.md)

