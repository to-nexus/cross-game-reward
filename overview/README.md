# Cross GameReward Protocol – Documentation Hub

## 📚 Documentation
- [00_overview.md](./00_overview.md) – Overview & Quick start
- [01_architecture.md](./01_architecture.md) – System architecture
- [02_reward_mechanism.md](./02_reward_mechanism.md) – Reward mechanics
- [03_security_and_testing.md](./03_security_and_testing.md) – Security, testing, and best practices

---

## 🚀 Quick Reference
- `rewardPerToken` accumulation powers all reward distribution:
  ```
  rewardPerTokenStored += (newReward × 1e18) / totalDeposited
  userReward = userBalance × (rewardPerTokenStored - userCheckpoint) / 1e18
  ```
- 4 core contracts: `CrossGameReward`, `CrossGameRewardPool`, `CrossGameRewardRouter`, `WCROSS`
- Foundry test-suite: 12 files / 233 cases

---

## 🧭 How to Use These Docs
1. **Product/PM** – read 00_overview → grasp UX flow and module map.
2. **Integrators** – use 01_architecture for contract APIs and role matrix.
3. **Quants & Auditors** – 02_reward_mechanism for proofs, 03_security_and_testing for guarantees vs. open risks.
4. **Developers** – refer to these docs for implementation details and best practices.

---

## ⚠️ Important Notes
- Operational controls (router assignment, UUPS upgrades, pausing) remain centralized under the `DEFAULT_ADMIN_ROLE`; multi-sig governance is recommended before production rollout.

---

### Related Resources
- [../test/README.md](../test/README.md) – Test execution guide
- OpenZeppelin Contracts v5.4.0, Foundry Book, ConsenSys smart contract best practices
