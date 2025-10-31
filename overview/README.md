# Cross Staking Protocol - 기술 문서

## 📚 문서 구성

이 폴더는 Cross Staking Protocol의 상세 기술 문서를 포함합니다.

### 문서 목록

0. **[00_overview.md](./00_overview.md)** - 빠른 시작
1. **[01_architecture.md](./01_architecture.md)** - 시스템 아키텍처
2. **[02_reward_mechanism.md](./02_reward_mechanism.md)** - 보상 메커니즘
3. **[03_security_and_testing.md](./03_security_and_testing.md)** - 보안 및 테스트

---

## 📖 빠른 시작

### 핵심 개념

Cross Staking Protocol은 **rewardPerToken 누적 방식**을 사용합니다:

```
rewardPerTokenStored += (newReward × 1e18) / totalStaked
userReward = userBalance × (rewardPerTokenStored - userCheckpoint) / 1e18
```

**특징:**
- ✅ O(1) 가스 비용
- ✅ 공정한 분배
- ✅ 예치 이후 보상만

### 주요 기능

| 기능 | 설명 | 문서 |
|------|------|------|
| **Staking** | Native/ERC20 스테이킹 | [Architecture](./01_architecture.md) |
| **Rewards** | 실시간 보상 분배 | [Reward Mechanism](./02_reward_mechanism.md) |
| **Security** | 7개 보안 계층 | [Security](./03_security_and_testing.md) |

---

## 🎯 문서 사용 가이드

### 개발자용

**처음 시작:**
1. [Overview](./00_overview.md) 읽기
2. [Architecture](./01_architecture.md) 구조 이해
3. [Reward Mechanism](./02_reward_mechanism.md) 수학 이해
4. 테스트 코드 실행

**통합 개발:**
1. [Architecture](./01_architecture.md#사용-시나리오) 참고
2. 배포 가이드 따라하기

### 감사자(Auditor)용

**보안 검토:**
1. [Security](./03_security_and_testing.md#보안-메커니즘) 확인
2. [Testing](./03_security_and_testing.md#테스트-체계) 커버리지 검토
3. [Reward Mechanism](./02_reward_mechanism.md#불변성) 불변성 검증

**수학적 검증:**
1. [Reward Mechanism - 수학적 원리](./02_reward_mechanism.md#수학적-원리)
2. [Simulation](./02_reward_mechanism.md#보상-분배-시뮬레이션)

### 사용자용

**이해하기:**
1. [Overview](./00_overview.md)
2. [Reward Mechanism - 기본 공식](./02_reward_mechanism.md#수학적-원리)

---

## 📊 프로젝트 통계

### 컨트랙트

- **Files:** 8개 (4개 컨트랙트 + 4개 Interface)
- **Lines:** ~1,150 라인
- **Size:** ~35 KB

### 테스트

- **Test Files:** 9개
- **Test Cases:** 159개
- **Success Rate:** 100%
- **Coverage:** ~100%

---

## 🔗 관련 링크

### 프로젝트

- **Repository:** GitHub
- **Tests:** [../test/README.md](../test/README.md)

### 참고

- **OpenZeppelin:** [Contracts v5.4.0](https://docs.openzeppelin.com/contracts/5.x/)
- **Foundry:** [Book](https://book.getfoundry.sh/)

### 표준

- **EIP-1967:** Proxy Storage Slots
- **EIP-1153:** Transient Storage
- **ERC-20:** Token Standard

---

## 💡 질문 및 지원

### FAQ

**Q: 왜 rewardPerToken 누적 방식인가?**  
A: O(1) 가스 비용으로 무제한 사용자 지원 가능

**Q: 다중 보상 토큰은 몇 개까지?**  
A: 이론적으로 무제한, 실전에서는 3-5개 권장

**Q: UUPS vs Transparent Proxy?**  
A: UUPS가 가스비 저렴하고 더 안전

**Q: 스테이커 없을 때 보상은?**  
A: 첫 스테이커가 모두 받음

**Q: Router 교체 가능?**  
A: 네, setRouter()로 언제든 교체 가능

---

**버전:** 1.0.0  
**최종 업데이트:** 2025-10-31
