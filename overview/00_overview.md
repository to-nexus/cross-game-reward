# Cross GameReward Protocol – Overview

## 🎯 Introduction
Cross GameReward Protocol is a **multi-pool deposit system** designed for native CROSS and ERC-20 tokens. It wraps native assets automatically and distributes rewards through an efficient `rewardPerToken` accumulator.

---

## 🏗️ System at a Glance

```
┌───────────────────────────────┐
│            User               │
│ (Native CROSS / ERC-20 deposit) │
└──────────────┬────────────────┘
               │
               ▼
┌───────────────────────────────┐
│ CrossGameRewardRouter (CGRR)      │
│ • depositNative / withdrawNative │
│ • depositERC20 / withdrawERC20   │
│ • claimRewards / claimReward     │
│ • redeployable front-door     │
└──────┬────────────────────────┘
       │
       ├──► WCROSS (WETH9 pattern)
       │    • anyone can deposit/withdraw
       │
       ▼
┌───────────────────────────────┐
│ CrossGameReward (CGR)             │
│ • UUPS upgradeable factory    │
│ • createPool / setRouter      │
└──────┬────────────────────────┘
       │ creates
       ▼
┌───────────────────────────────┐
│ CrossGameRewardPool (CGRP) × N    │
│ • UUPS upgradeable pools      │
│ • depositFor / withdrawFor       │
│ • rewardPerToken accumulator  │
│ • multi reward tokens         │
└───────────────────────────────┘
```

---

## ✨ Key Capabilities
1. **Unlimited pools** – multiple pools per deposit token supported.
2. **Native token UX** – automatic wrap/unwrap via WCROSS.
3. **Multi-reward** – each pool can emit several ERC-20 reward tokens.
4. **Upgradeable** – CrossGameReward & CrossGameRewardPool follow UUPS; router is redeployable.
5. **O(1) reward accounting** – `rewardPerToken` accumulator keeps gas flat per deposit.
6. **Simplified access control** – Owner and RewardRoot based permissions.
7. **3-state pool management** – Active/Inactive/Paused for granular control.
8. **Fair reward distribution** – Zero-deposit deposits automatically marked as withdrawable.
9. **Partial withdrawals** – withdraw specific amounts while remaining balance keeps earning.

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

### User: Deposit native CROSS
```solidity
// Deposit native CROSS (no approval needed - Router auto-wraps)
router.depositNative{value: 100 ether}(poolId);

// Claim rewards only (keep deposit)
router.claimRewards(poolId);

// Partial withdraw (30 ether) + claim all rewards
router.withdrawNative(poolId, 30 ether);

// Withdraw all remaining + claim all rewards
router.withdrawNative(poolId, 0);  // 0 = withdraw all
```

### Admin: Create pools & rewards
```solidity
// Create native CROSS pool
(uint256 poolId, address poolAddr) =
    crossDeposit.createPool("My Game Pool", address(wcross), 1 ether);

// Add reward token
crossDeposit.addRewardToken(poolId, address(usdt));
```

---

## 📈 Metrics
- Tests: **244 / 244 passing** (comprehensive edge case coverage including partial withdrawals)
- Gas footprint: all contracts < 24 KB, deposit/withdraw ~140–280k gas, claim O(1)
- Reward distribution: O(1) per deposit, proportional to deposit share
- Security: Multi-layered defense (reentrancy, access control, zero-deposit protection)

---

**Next:** [01_architecture.md](./01_architecture.md)
