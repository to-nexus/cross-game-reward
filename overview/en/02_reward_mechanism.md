# Reward Mechanism

## 📐 `rewardPerToken` Accumulator
The protocol distributes rewards with a global accumulator that tracks “reward per unit staked”. This delivers O(1) gas regardless of staker count.

---

## 🧮 Core Formulae

### 1. On reward deposit
```
if (totalStaked == 0) {
    withdrawableAmount += newReward  // Not distributed
} else {
    rewardPerTokenStored += (newReward × PRECISION) / totalStaked
}
PRECISION = 1e18
```

**Zero-Stake Protection:**
- Rewards deposited when `totalStaked=0` are classified as `withdrawableAmount`
- Owner can recover via `CrossStaking.withdrawFromPool()`
- Protects first staker from receiving unallocated pre-staking rewards

**After Token Removal:**
- Distributable balance at removal is frozen as `distributedAmount` (user-claimable)
- `withdrawableAmount` remains intact (owner-recoverable)
- New deposits after removal are also owner-recoverable

Interpretation: how much reward (scaled by `PRECISION`) each unit of staking tokens has earned.

### 2. On user accrual
```
earned = userBalance × (rewardPerTokenStored - userCheckpoint) / PRECISION
totalReward = storedRewards + earned
```
Where:
- `userCheckpoint` — the accumulator value last time the user was synced.
- `storedRewards` — rewards already stored for later withdrawal.

---

## 🔄 Distribution Walkthrough
Example:
1. Alice stakes 100 CROSS at _t0_ → `rewardPerTokenStored = 0`.
2. 100 reward tokens deposited at _t3_:
   - `rewardPerTokenStored = 1e18`
3. Bob stakes 100 CROSS at _t10_:
   - `userCheckpoint[bob] = 1e18`
4. Another 100 rewards at _t15_:
   - `rewardPerTokenStored = 1.5e18`
5. Claims:
   - Alice earns `100 × (1.5e18 - 0) / 1e18 = 150`
   - Bob earns `100 × (1.5e18 - 1e18) / 1e18 = 50`

---

## 🎯 Properties
- **Fairness** – rewards are proportional to stake share after the time of deposit.
- **No retroactive payouts** – deposits before a user joined never count toward their rewards.
- **Time independence** – claiming later yields the same amount as claiming immediately.

---

## 💡 Special Cases

### No stakers present (Zero-Stake Protection)
```solidity
function _syncReward(IERC20 token) internal {
    if (totalStaked == 0) {
        rt.withdrawableAmount += newReward;  // Mark as owner-recoverable
        rt.lastBalance = currentBalance;
        return;
    }
    // ...
}
```

**Behavior (current version):**
1. Rewards deposited when `totalStaked=0` are classified as `withdrawableAmount`
2. First staker does **not** receive these rewards
3. Owner can recover via `CrossStaking.withdrawFromPool()`

**Example:**
```
1. Pool is empty
2. 1000 reward deposited → withdrawableAmount = 1000
3. Alice stakes
4. Alice does NOT get the 1000 (fair distribution)
5. Later 100 reward deposited → Alice gets 100
6. Owner can recover withdrawableAmount of 1000
```

**Reward Queries:**
- `pendingRewards(user)`: Returns all active reward tokens and amounts `(address[] tokens, uint[] rewards)`
- `pendingReward(user, token)`: Query specific token reward `uint amount`

### Multiple reward tokens
- Pools maintain an `EnumerableSet` of active reward tokens.
- Each token has its own accumulator (`RewardToken.rewardPerTokenStored`).
- Removed tokens can still be claimed via `claimReward(removedToken)`.

---

## 🔍 Invariants
1. **Reward conservation**: deposits = claimed + pending + rounding dust.
2. **Monotonic accumulator**: `rewardPerTokenStored` never decreases.
3. **Exactness**: user earnings = Σ(stake share × deposit amount), within integer rounding.

---

## ⛽ Gas Profile
- `_syncReward` iterates reward tokens (linear in active token count).
- Claiming all rewards is O(n) in the number of reward tokens.
- No loops over staker addresses; supports large pools.

---

## ⚠️ Operational Notes
- Removed reward tokens remain claimable via `claimReward(removedToken)`
- Zero-stake deposits are owner-recoverable via `withdrawFromPool`
- Reward queries support both bulk (`pendingRewards`) and single-token (`pendingReward`) lookups

---

**Next:** [03_security_and_testing.md](./03_security_and_testing.md)
