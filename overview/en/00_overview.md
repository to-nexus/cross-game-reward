# Cross Staking Protocol v2.0 – Overview

## 🎯 Introduction
Cross Staking Protocol is a **multi-pool staking system** designed for native CROSS and ERC-20 tokens. It wraps native assets automatically and distributes rewards through an efficient `rewardPerToken` accumulator.

---

## 🏗️ System at a Glance

```
┌───────────────────────────────┐
│            User               │
│ (Native CROSS / ERC-20 stake) │
└──────────────┬────────────────┘
               │
               ▼
┌───────────────────────────────┐
│ CrossStakingRouter (CSR)      │
│ • stakeNative / unstakeNative │
│ • stakeERC20 / unstakeERC20   │
│ • redeployable front-door     │
└──────┬────────────────────────┘
       │
       ├──► WCROSS (router only wrapper)
       │
       ▼
┌───────────────────────────────┐
│ CrossStaking (CS)             │
│ • UUPS upgradeable factory    │
│ • createPool / setRouter      │
└──────┬────────────────────────┘
       │ creates
       ▼
┌───────────────────────────────┐
│ CrossStakingPool (CSP) × N    │
│ • UUPS upgradeable pools      │
│ • stakeFor / unstakeFor       │
│ • rewardPerToken accumulator  │
│ • multi reward tokens         │
└───────────────────────────────┘
```

---

## ✨ Key Capabilities
1. **Unlimited pools** – multiple pools per staking token supported.
2. **Native token UX** – automatic wrap/unwrap via WCROSS.
3. **Multi-reward** – each pool can emit several ERC-20 reward tokens.
4. **Upgradeable** – CrossStaking & CrossStakingPool follow UUPS; router is redeployable.
5. **O(1) reward accounting** – `rewardPerToken` accumulator keeps gas flat per deposit.

---

## 📊 Tech Stack
- Solidity 0.8.28, Foundry toolchain
- OpenZeppelin Contracts Upgradeable v5.4.0
- UUPS proxies (EIP-1967 slots)
- ReentrancyGuardTransient (EIP-1153)

---

## 📚 Documentation Map
- [Architecture](./01_architecture.md)
- [Reward Mechanics](./02_reward_mechanism.md)
- [Security & Testing](./03_security_and_testing.md)

---

## 🚀 Quick Start

### User: Stake native CROSS
```solidity
// One-time approval
wcross.approve(address(router), type(uint256).max);

// Stake native CROSS
router.stakeNative{value: 100 ether}(poolId);

// Unstake + claim all rewards
router.unstakeNative(poolId);
```

### Admin: Create pools & rewards
```solidity
// Create native CROSS pool
(uint256 poolId, address poolAddr) =
    crossStaking.createPool(address(wcross), 2 days);

// Add reward token
crossStaking.addRewardToken(poolId, address(usdt));
```

---

## 📈 Current Metrics
- Tests: **159 / 159 passing** (`forge test`, 2025-11-03)
- Gas footprint: all contracts < 24 KB, stake/unstake ~140–280k gas
- Reward distribution: O(1) per deposit, proportional to stake share

---

**Next:** [01_architecture.md](./01_architecture.md)
