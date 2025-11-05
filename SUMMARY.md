# Cross Staking Protocol - Project Summary

## 🎯 Overview

Cross Staking Protocol is a multi-pool staking system for native CROSS and ERC-20 tokens. A single factory deploys upgradeable pools, the router handles wrap/unwrap logic, and rewards are distributed with flat gas cost.

### Value proposition

- ✅ Native CROSS and ERC-20 in one framework  
- ✅ Unlimited pools per token, each with independent configuration  
- ✅ Accurate reward allocation with `rewardPerToken` accumulator  
- ✅ Simplified access control with Owner/StakingRoot modifiers
- ✅ 3-state pool management (Active/Inactive/Paused)
- ✅ Fair reward distribution with zero-stake protection
- ✅ Enhanced reward query APIs with token addresses
- ✅ Hardened security stack with upgrade gates and role separation

---

## 📐 Architecture

```
User (Native CROSS / ERC-20)
    │
    ▼
CrossStakingRouter ──► WCROSS (wrap)
    │
    ▼
CrossStaking (UUPS factory)
    │ creates
    ▼
CrossStakingPool × N (UUPS pools)
```

| Component            | Responsibility                                                           |
|----------------------|---------------------------------------------------------------------------|
| CrossStaking         | Creates pools, manages reward tokens, sets pool status, configures router |
| CrossStakingPool     | Holds stake balances, updates rewards, 3-state management                |
| CrossStakingRouter   | User entry point for native and ERC-20 staking, drives WCROSS            |
| WCROSS               | Wraps native CROSS; only the router may call `deposit/withdraw`          |

---

## 🔄 Operational flows

### Native staking
1. User approves WCROSS to the router  
2. `stakeNative` wraps native CROSS into WCROSS and stakes via `stakeFor`  
3. `unstakeNative` claims rewards, unwraps WCROSS, and returns native CROSS

### ERC-20 staking
1. User approves the router for the staking token  
2. Router transfers tokens, calls `stakeFor`, and records the position  
3. `unstakeERC20` returns the principal to the user and rewards directly from the pool

### Reward funding & queries
- Any address can transfer reward tokens to the pool  
- `_syncReward` detects balance deltas during the next interaction  
- `rewardPerTokenStored` keeps the per-stake reward up to date with O(1) gas

**Reward Query APIs:**
- `pendingRewards(user)`: Returns all active reward tokens and amounts → `(address[] tokens, uint[] rewards)`
- `pendingReward(user, token)`: Query specific token reward → `uint amount`
- `getUserStakingInfo(poolId, user)`: Unified staking info → `(uint stakedAmount, address[] tokens, uint[] rewards)`

### Zero-stake protection
- Rewards deposited when `totalStaked=0` are classified as `withdrawableAmount`
- Protects first staker from receiving unallocated pre-staking rewards
- Owner can recover via `CrossStaking.withdrawFromPool()`

### Removed reward tokens
- `removeRewardToken` freezes distributable balance as `distributedAmount`
- Users can still `claimReward(removedToken)` to collect these rewards  
- New deposits after removal are added to `withdrawableAmount` for owner recovery

---

## 🎯 Key Features

### 1. 3-State Pool Management
- **Active (0)**: All operations allowed (stake, unstake, claim)
- **Inactive (1)**: Only unstake and claim allowed
- **Paused (2)**: All operations stopped

Control: `CrossStaking.setPoolStatus(poolId, status)`

### 2. Reward Mechanics
- **O(1) gas**: `rewardPerToken` accumulation pattern
- **Zero-stake protection**: Prevents unfair rewards to first staker
- **Removed tokens**: `distributedAmount` (user-claimable) vs `withdrawableAmount` (owner-recoverable)
- **Accuracy**: Mathematically guaranteed proportional distribution

### 3. Access Control
**CrossStaking:**
- `DEFAULT_ADMIN_ROLE` (owner): Router assignment, pool implementation, upgrades
- `MANAGER_ROLE`: Pool creation, reward tokens, pool status, withdrawals

**CrossStakingPool:**
- `onlyOwner()`: CrossStaking's owner, upgrade authorization
- `onlyStakingRoot()`: CrossStaking contract, all management functions
- `onlyRouter`: Router-only, `stakeFor/unstakeFor`

---

## 🔒 Security stack

1. **ReentrancyGuardTransient (EIP-1153)** wraps every state-changing entry point  
2. **SafeERC20** handles token transfers safely  
3. **Simplified access control** with Owner/StakingRoot modifiers
4. **3-state pool management** for granular control
5. **UUPS upgrade paths** restricted to owner
6. **Custom errors** to reduce gas and clarify revert reasons  
7. **Router caller checks** on pool and WCROSS methods
8. **Zero-stake protection** prevents unfair reward allocation
9. **Event optimization** removes duplicate emissions

Operational guidance:
- Protect admin keys (router changes, upgrades) with a multisig/governance process  
- `setPoolStatus(poolId, status)`: 0=Active, 1=Inactive, 2=Paused
- Zero-stake deposits recoverable via `withdrawFromPool`

---

## 🧪 Testing & quality

```bash
forge test                                  # full suite
forge test --match-contract CrossStaking   # specific contract
forge test --gas-report                    # gas report
```

### Test Statistics (Foundry)

| Suite                          | Passed tests |
|--------------------------------|--------------|
| WCROSS                         | 10           |
| CrossStaking                   | 33           |
| CrossStakingRouter             | 28           |
| CrossStakingPoolStaking        | 18           |
| CrossStakingPoolRewards        | 27           |
| CrossStakingPoolAdmin          | 34           |
| CrossStakingPoolIntegration    | 11           |
| CrossStakingPoolPendingRewards | 9            |
| CrossStakingPoolSecurity       | 21           |
| CrossStakingPoolEdgeCases      | 12           |
| FullIntegration                | 9            |
| **Total**                      | **212**      |

**Coverage:** ~100%, covering multi-pool deployment, reward removal, router flows, zero-stake scenarios, stress cases, and invariant checks.

### Recent Improvements
1. ✅ Enhanced API: `pendingRewards()` returns token addresses
2. ✅ Added `pendingReward()` for single-token queries
3. ✅ Event optimization: removed duplicate emissions
4. ✅ Added 6 new tests (PendingRewards suite)
5. ✅ Complete documentation update

---

## 📊 Code Quality

### Statistics
- **Contracts**: 4 main + 4 interfaces
- **Test Suites**: 11
- **Total Tests**: 212 (100% passing)
- **Lines of Code**: ~3,500 (including tests)
- **Warnings**: 0
- **Gas Optimizations**: Event deduplication, custom errors

### Deployment Checklist
- ✅ 212/212 tests passing  
- ✅ Zero compilation warnings
- ✅ Reentrancy protection and role checks verified  
- ✅ Zero-stake protection implemented
- ✅ Removed reward token settlement validated  
- ✅ UUPS upgrade paths tested (`upgradeToAndCall`)  
- ✅ Documentation up-to-date
- ✅ API improvements completed
- ✅ Event optimization done
- [ ] External third-party audit (recommended)

---

## 🚀 Latest Changes

### Breaking Changes
⚠️ Frontend migration required:

```solidity
// Old way
uint[] memory rewards = pool.pendingRewards(user);
address[] memory tokens = pool.getRewardTokens();
// Then match arrays manually

// New way
(address[] memory tokens, uint[] memory rewards) = pool.pendingRewards(user);
// Already matched!

// Or for specific token
uint amount = pool.pendingReward(user, specificToken);
```

### New Features
1. **Enhanced Reward Queries**: Token addresses and rewards in single call
2. **Event Optimization**: Removed duplicate events for gas savings
3. **Comprehensive Testing**: 212 tests covering all edge cases

---

## 📚 Documentation & links

- [README Korean](README_ko.md)
- [README English](README.md)
- [Architecture](overview/en/01_architecture.md)  
- [Reward Mechanism](overview/en/02_reward_mechanism.md)  
- [Security & Testing](overview/en/03_security_and_testing.md)  
- [Test Guide](test/README.md)  
- [OpenZeppelin Contracts](https://docs.openzeppelin.com/contracts/)  
- [Foundry Book](https://book.getfoundry.sh/)

---

## 📜 License

MIT

---

## ✨ Conclusion

Cross Staking Protocol is **fully tested and documented, production-ready**:
- 212 tests passing (100%)
- Enhanced API for better UX
- Optimized event logging
- Complete documentation
- Zero compilation warnings

Ready for deployment! 🎉
