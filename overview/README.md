# Cross Staking Protocol – Documentation Hub

## 📚 Language Packs
- **한국어 (ko)**
  - [00_overview.md](./ko/00_overview.md) – 개요 및 빠른 시작
  - [01_architecture.md](./ko/01_architecture.md) – 시스템 아키텍처
  - [02_reward_mechanism.md](./ko/02_reward_mechanism.md) – 보상 수학
  - [03_security_and_testing.md](./ko/03_security_and_testing.md) – 보안·테스트 및 알려진 이슈
- **English (en)**
  - [00_overview.md](./en/00_overview.md) – Overview & Quick start
  - [01_architecture.md](./en/01_architecture.md) – System architecture
  - [02_reward_mechanism.md](./en/02_reward_mechanism.md) – Reward mechanics
  - [03_security_and_testing.md](./en/03_security_and_testing.md) – Security, testing, known issues

---

## 🚀 Quick Reference
- `rewardPerToken` accumulation powers all reward distribution:
  ```
  rewardPerTokenStored += (newReward × 1e18) / totalStaked
  userReward = userBalance × (rewardPerTokenStored - userCheckpoint) / 1e18
  ```
- 4 core contracts: `CrossStaking`, `CrossStakingPool`, `CrossStakingRouter`, `WCROSS`
- Foundry test-suite: 9 files / 159 cases (`forge test`)

---

## 🧭 How to Use These Docs
1. **Product/PM** – read 00_overview → grasp UX flow and module map.
2. **Integrators** – use 01_architecture for contract APIs and role matrix.
3. **Quants & Auditors** – 02_reward_mechanism for proofs, 03_security_and_testing for guarantees vs. open risks.
4. **Developers** – rely on language-specific pack that matches your audience; keep both versions in sync when updating specs.

---

## ⚠️ Current Status Highlights
- High-severity issue H-01 (removed reward token claims) fixed on 2025-11-03; removed tokens are now auto-settled during unstake while active stakes continue to use the standard claim functions. See both language versions of `03_security_and_testing.md` for helper details and regression coverage.
- Operational controls (router assignment, UUPS upgrades, pausing) remain centralized under the `DEFAULT_ADMIN_ROLE`; multi-sig governance is recommended before production rollout.

---

### Related Resources
- [../test/README.md](../test/README.md) – 테스트 실행 가이드
- OpenZeppelin Contracts v5.4.0, Foundry Book, ConsenSys smart contract best practices
