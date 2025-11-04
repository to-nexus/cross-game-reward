# Reward Mechanism

## 📐 `rewardPerToken` Accumulator
The protocol distributes rewards with a global accumulator that tracks “reward per unit staked”. This delivers O(1) gas regardless of staker count.

---

## 🧮 Core Formulae

### 1. On reward deposit
```
rewardPerTokenStored += (newReward × PRECISION) / totalStaked
PRECISION = 1e18
```
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

### No stakers present
```solidity
if (totalStaked == 0) return; // _syncReward
```
- Rewards transferred while no one is staked are not synced.
- The first staker after such a deposit will claim the entire amount (documented behavior).

### Multiple reward tokens
- Pools maintain an `EnumerableSet` of active reward tokens.
- Each token has its own accumulator (`RewardToken.rewardPerTokenStored`).

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

## ⚠️ Operational Reminder
While reward tokens can be removed, the current implementation requires users to claim them **before** unstaking (see `03_security_and_testing.md` for the H-01 issue). Until patched, communicate this requirement in your UI or operational runbook.

---

**Next:** [03_security_and_testing.md](./03_security_and_testing.md)
