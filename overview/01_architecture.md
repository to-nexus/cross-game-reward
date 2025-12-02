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
│ • claimRewards / claimReward     │
│ • stateless, redeployable     │
└──────┬────────────────────────┘
       │
       ├──► WCROSS (WETH9 pattern)
       │    • anyone can deposit/withdraw
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
- Purpose: wrap native CROSS into an ERC-20 for deposit (WETH9 pattern).
- Key functions:
  ```solidity
  deposit() public payable              // anyone can call
  withdraw(uint amount) external        // anyone can call
  withdrawTo(address to, uint) public   // anyone can call
  ```
- Features:
  - Follows WETH9 standard pattern
  - Open access (anyone can deposit/withdraw)
  - Easy DEX integration
  - 1:1 parity maintained

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
  ```solidity
  // Deposit/Withdraw
  deposit(uint amount)                        // Active state only
  depositFor(address account, uint amount)    // Router only, Active state
  withdraw(uint amount)                      // Active/Inactive state (0 = withdraw all)
  withdrawFor(address account, uint amount)  // Router only (0 = withdraw all)

  // Claim (refactored)
  claimRewards()                            // claim all rewards
  claimRewardsFor(address account)           // Router only
  claimReward(IERC20 token)                 // claim specific token
  claimRewardFor(address account, token)     // Router only

  // Admin
  addRewardToken(IERC20 token)              // CrossGameReward only
  removeRewardToken(IERC20 token)           // CrossGameReward only (auto-claims removed tokens on withdraw)
  withdraw(IERC20 token, address to)        // CrossGameReward only
  setPoolStatus(uint8 status)               // CrossGameReward only
  ```

### 4. CrossGameRewardRouter
- Immutable references:
  ```solidity
  CrossGameReward public immutable crossGameReward;
  IWCROSS public immutable wcross;
  ```
- Key functions:
  ```solidity
  // Deposit/Withdraw
  depositNative(uint poolId) payable
  withdrawNative(uint poolId, uint amount)      // 0 = withdraw all
  depositERC20(uint poolId, uint amount)
  depositERC20WithPermit(uint poolId, uint amount, ...) // EIP-2612
  withdrawERC20(uint poolId, uint amount)       // 0 = withdraw all

  // Claim (newly added)
  claimRewards(uint poolId)                    // claim all rewards
  claimReward(uint poolId, address token)       // claim specific token

  // View
  getUserDepositInfo(uint poolId, address user)
  getPendingRewards(uint poolId, address user)  // all pending rewards
  getPendingReward(uint poolId, address user, token) // specific token pending
  isNativePool(uint poolId)
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
