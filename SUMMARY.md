# Cross GameReward Protocol - Project Summary

## 🎯 Overview

Cross GameReward Protocol is a multi-pool deposit system for native CROSS and ERC-20 tokens. A single factory deploys upgradeable pools, the router handles wrap/unwrap logic, and rewards are distributed with flat gas cost.

### Value proposition

- ✅ Native CROSS and ERC-20 in one framework  
- ✅ Unlimited pools per token, each with independent configuration  
- ✅ Accurate reward allocation with `rewardPerToken` accumulator  
- ✅ Simplified access control with Owner/RewardRoot modifiers
- ✅ 3-state pool management (Active/Inactive/Paused)
- ✅ Fair reward distribution with zero-deposit protection
- ✅ Enhanced reward query APIs with token addresses
- ✅ Hardened security stack with upgrade gates and role separation
- ✅ Partial withdrawal support for flexible liquidity management

---

## 📐 Architecture

```
User (Native CROSS / ERC-20)
    │
    ▼
CrossGameRewardRouter ──► WCROSS (wrap)
    │
    ▼
CrossGameReward (UUPS factory)
    │ creates
    ▼
CrossGameRewardPool × N (UUPS pools)
```

| Component            | Responsibility                                                           |
|----------------------|---------------------------------------------------------------------------|
| CrossGameReward         | Creates pools, manages reward tokens, sets pool status, configures router |
| CrossGameRewardPool     | Holds deposit balances, updates rewards, 3-state management                |
| CrossGameRewardRouter   | User entry point for native and ERC-20 deposit, drives WCROSS            |
| WCROSS               | Wraps native CROSS; only the router may call `deposit/withdraw`          |

---

## 🔄 Operational flows

### Native deposit
1. User approves WCROSS to the router  
2. `depositNative` wraps native CROSS into WCROSS and deposits via `depositFor`  
3. `withdrawNative(poolId, amount)` claims all rewards, unwraps WCROSS, and returns native CROSS
   - `amount > 0`: Partial withdrawal of specified amount
   - `amount = 0`: Full withdrawal of all deposited tokens

### ERC-20 deposit
1. User approves the router for the deposit token  
2. Router transfers tokens, calls `depositFor`, and records the position  
3. `withdrawERC20(poolId, amount)` returns the principal to the user and all rewards directly from the pool
   - `amount > 0`: Partial withdrawal of specified amount
   - `amount = 0`: Full withdrawal of all deposited tokens

### Reward funding & queries
- Any address can transfer reward tokens to the pool  
- `_syncReward` detects balance deltas during the next interaction  
- `rewardPerTokenStored` keeps the per-deposit reward up to date with O(1) gas

**Reward Query APIs:**
- `pendingRewards(user)`: Returns all active reward tokens and amounts → `(address[] tokens, uint[] rewards)`
- `pendingReward(user, token)`: Query specific token reward → `uint amount`
- `getUserDepositInfo(poolId, user)`: Unified deposit info → `(uint depositedAmount, address[] tokens, uint[] rewards)`

### Zero-deposit protection
- Rewards deposited when `totalDeposited=0` are classified as `withdrawableAmount`
- Protects first depositor from receiving unallocated pre-deposit rewards
- Owner can recover via `CrossGameReward.withdrawFromPool()`

### Removed reward tokens
- `removeRewardToken` freezes distributable balance as `distributedAmount`
- Users can still `claimReward(removedToken)` to collect these rewards  
- New deposits after removal are added to `withdrawableAmount` for owner recovery

---

## 🎯 Key Features

### 1. 3-State Pool Management
- **Active (0)**: All operations allowed (deposit, withdraw, claim)
- **Inactive (1)**: Only withdraw and claim allowed
- **Paused (2)**: All operations stopped

Control: `CrossGameReward.setPoolStatus(poolId, status)`

### 2. Reward Mechanics
- **O(1) gas**: `rewardPerToken` accumulation pattern
- **Zero-deposit protection**: Prevents unfair rewards to first depositor
- **Removed tokens**: `distributedAmount` (user-claimable) vs `withdrawableAmount` (owner-recoverable)
- **Accuracy**: Mathematically guaranteed proportional distribution
- **Partial withdrawal**: Withdraw specific amounts (`withdraw(amount)`) while remaining balance continues earning
- **Automatic claim**: All accumulated rewards claimed during any withdrawal (partial or full)

### 3. Access Control
**CrossGameReward:**
- `DEFAULT_ADMIN_ROLE` (owner): Router assignment, pool implementation, upgrades
- `MANAGER_ROLE`: Pool creation, reward tokens, pool status, withdrawals

**CrossGameRewardPool:**
- `onlyOwner()`: CrossGameReward's owner, upgrade authorization
- `onlyRewardRoot()`: CrossGameReward contract, all management functions
- `onlyRouter`: Router-only, `depositFor/withdrawFor`

---

## 🔒 Security stack

1. **ReentrancyGuardTransient (EIP-1153)** wraps every state-changing entry point  
2. **SafeERC20** handles token transfers safely  
3. **Simplified access control** with Owner/RewardRoot modifiers
4. **3-state pool management** for granular control
5. **UUPS upgrade paths** restricted to owner
6. **Custom errors** to reduce gas and clarify revert reasons  
7. **Router caller checks** on pool and WCROSS methods
8. **Zero-deposit protection** prevents unfair reward allocation
9. **Event optimization** removes duplicate emissions

Operational guidance:
- Protect admin keys (router changes, upgrades) with a multisig/governance process  
- `setPoolStatus(poolId, status)`: 0=Active, 1=Inactive, 2=Paused
- Zero-deposit deposits recoverable via `withdrawFromPool`

---

## 🧪 Testing & quality

```bash
forge test                                  # full suite
forge test --match-contract CrossGameReward   # specific contract
forge test --gas-report                    # gas report
```

### Test Statistics (Foundry)

| Suite                          | Passed tests |
|--------------------------------|--------------|
| WCROSS                         | 10           |
| CrossGameReward                   | 33           |
| CrossGameRewardRouter             | 44           |
| CrossGameRewardPoolDeposit        | 24           |
| CrossGameRewardPoolRewards        | 27           |
| CrossGameRewardPoolAdmin          | 34           |
| CrossGameRewardPoolIntegration    | 11           |
| CrossGameRewardPoolPendingRewards | 9            |
| CrossGameRewardPoolSecurity       | 21           |
| CrossGameRewardPoolEdgeCases      | 12           |
| CrossGameRewardPoolClaimRecovery  | 10           |
| FullIntegration                | 9            |
| **Total**                      | **244**      |

**Coverage:** ~100%, covering multi-pool deployment, reward removal, router flows, zero-deposit scenarios, partial withdrawals, stress cases, and invariant checks.

---

## 📊 Code Quality

### Statistics
- **Contracts**: 4 main + 4 interfaces
- **Test Suites**: 12
- **Total Tests**: 244 (100% passing)
- **Lines of Code**: ~4,000 (including tests)
- **Warnings**: 0
- **Gas Optimizations**: Event deduplication, custom errors

### Pre-deployment Checklist
- All tests passing with zero warnings
- Reentrancy protection and role checks in place
- Zero-deposit protection implemented
- Removed reward token settlement validated
- UUPS upgrade paths tested
- External third-party audit recommended

---

## 📚 Documentation & links

- [Overview](overview/00_overview.md)
- [Architecture](overview/01_architecture.md)
- [Reward Mechanism](overview/02_reward_mechanism.md)
- [Security & Testing](overview/03_security_and_testing.md)
- [Test Guide](test/README.md)
- [OpenZeppelin Contracts](https://docs.openzeppelin.com/contracts/)
- [Foundry Book](https://book.getfoundry.sh/)

---

## 📜 License

BUSL 1.1 (Business Source License 1.1) - See [LICENSE](LICENSE) file for details
