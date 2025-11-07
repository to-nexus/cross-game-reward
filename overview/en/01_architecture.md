# Cross GameReward Protocol – Architecture

## 📐 Overview
The protocol exposes a modular multi-pool deposit topology built around a `rewardPerToken` accumulator. This document explains how the contracts interact and which responsibilities reside where.

---

## 🏗️ Core Components

```
┌───────────────────────────────┐
│            User               │
│  (Native CROSS / ERC-20)      │
└──────────────┬────────────────┘
               │
               ▼
┌───────────────────────────────┐
│ CrossGameRewardRouter            │
│ • depositNative / withdrawNative │
│ • depositERC20 / withdrawERC20   │
│ • stateless, redeployable     │
└──────┬────────────────────────┘
       │
       ├──► WCROSS (router-only)
       │
       ▼
┌───────────────────────────────┐
│ CrossGameReward (factory)        │
│ • UUPS upgradeable            │
│ • createPool / setRouter      │
└──────┬────────────────────────┘
       │ creates
       ▼
┌───────────────────────────────┐
│ CrossGameRewardPool × N          │
│ • UUPS upgradeable            │
│ • rewardPerToken accounting   │
│ • multi reward token support  │
└───────────────────────────────┘
```

---

## 🔧 Contract Details

### 1. WCROSS
- Purpose: wrap native CROSS into an ERC-20 for deposit.
- Storage: `CrossGameReward public deposit`.
- Router-only operations enforced through `require(msg.sender == deposit.router())`.

### 2. CrossGameReward (factory)
- Storage highlights:
  ```solidity
  address public wcross;
  address public router;
  address public poolImplementation;
  mapping(uint256 => PoolInfo) public pools;
  ```
- Roles:
  - `DEFAULT_ADMIN_ROLE` – upgrades, router assignment.
  - `MANAGER_ROLE` – pool lifecycle, reward token management.
- Responsibilities:
  - Deploy new pools via ERC1967 proxies.
  - Track pool metadata (`PoolInfo`).
  - Pause/unpause pools via `setPoolActive`.

### 3. CrossGameRewardPool
- Storage highlights:
  ```solidity
  IERC20 public depositToken;
  address public crossDeposit;
  uint256 public minDepositAmount;
  mapping(address => uint256) public balances;
  mapping(address => mapping(IERC20 => UserReward)) public userRewards;
  EnumerableSet.AddressSet private _rewardTokenAddresses;
  EnumerableSet.AddressSet private _removedRewardTokenAddresses;
  mapping(IERC20 => RewardToken) private _rewardTokenData;
  ```
- Roles:
  - `DEFAULT_ADMIN_ROLE` → CrossGameReward admin (via `owner()` override).
  - `REWARD_ROOT_ROLE` → CrossGameReward contract.
  - `REWARD_MANAGER_ROLE`, `PAUSER_ROLE` delegated as needed.
- Key functions:
  - `deposit` / `depositFor` (router enforced via `_checkDelegate`).
  - `withdraw` / `withdrawFor` – full withdrawal plus reward claim.
  - `addRewardToken`, `removeRewardToken` (auto-claims removed tokens on withdraw), `withdraw`.

### 4. CrossGameRewardRouter
- Immutable references:
  ```solidity
  CrossGameReward public immutable crossDeposit;
  IWCROSS public immutable wcross;
  ```
- Wraps native deposits, forwards ERC-20 deposits via `depositFor`, handles `withdraw` flows and reward delivery.

---

## 🔐 Access Control & Security
- AccessControlDefaultAdminRules across factory/pools (time-delayed admin transfers on the factory).
- Pausable on pools; `setPoolActive` toggles `pause/unpause`.
- ReentrancyGuardTransient applied to every state-changing pool function.
- SafeERC20 used for all token interactions; `forceApprove` avoids stuck allowances.

---

## 🧭 Data Interactions
- Pools keep an address-set of active reward tokens; removed tokens remain in storage for historical claims.
- `CrossGameReward` maintains:
  - `_allPoolIds` enumerable set (global list).
  - `_poolsByDepositToken` mapping (token → pool IDs).
- Router performs no storage writes beyond allowances; all accounting lives in pools.

---

## 🔄 Upgrade Strategy
- Both `CrossGameReward` and `CrossGameRewardPool` implement `_authorizeUpgrade` guarded by `DEFAULT_ADMIN_ROLE`.
- Storage gap reserved (50 slots total, 8 used in factory, 9 used in pool).
- Router is intentionally immutable; deploy a new instance and call `setRouter`.

---

## 🧪 Integration Notes
- Direct pool interactions (`deposit`) remain available for power users (e.g., WCROSS LPs).
- Native deposit path requires router assignment on the factory before first deposit.
- Removed reward tokens are auto-settled during `_withdraw`; active positions still use `claimReward`/`claimRewards` for partial withdrawals.

---

**Next:** [02_reward_mechanism.md](./02_reward_mechanism.md)
