# Cross GameReward Protocol

Multi-pool deposit for native CROSS and ERC-20 tokens

## 🎯 Overview

Cross GameReward Protocol lets you launch and manage multiple deposit pools under a single factory. It wraps native CROSS when needed, supports arbitrary ERC-20 rewards, and keeps accounting gas costs flat.

### Architecture at a glance

```
User (Native CROSS / ERC-20)
    │
    ▼
CrossGameRewardRouter ──► WCROSS (wrap/unwrap)
    │
    ▼
CrossGameReward (factory, UUPS)
    │ creates
    ▼
CrossGameRewardPool × N (UUPS)
```

## ✨ Key features

- ✅ **Native CROSS support** – router auto-wraps and unwraps via WCROSS
- ✅ **Unlimited pools** – multiple pools per deposit token
- ✅ **Multi-reward** – each pool can emit several ERC-20 rewards
- ✅ **O(1) accounting** – reward distribution uses a `rewardPerToken` accumulator
- ✅ **Upgradeable** – CrossGameReward and pools follow the UUPS pattern
- ✅ **Simplified access control** – Owner and RewardRoot based permissions
- ✅ **3-state pool management** – Active/Inactive/Paused for granular control
- ✅ **Fair reward distribution** – Pre-deposit rewards automatically marked as withdrawable
- ✅ **Removed reward settlement** – rewards for removed tokens are auto-claimed during withdraw

## 🚀 Quick start

```bash
forge install
forge test
```

Deployment example:

```bash
forge script script/DeployFullSystem.s.sol:DeployFullSystem \
  --rpc-url <RPC_URL> \
  --private-key <PRIVATE_KEY> \
  --broadcast
```

## 💡 Usage snippets

### User – deposit native CROSS

```solidity
// Approve once
wcross.approve(address(router), type(uint256).max);

// Deposit
router.depositNative{value: 100 ether}(poolId);

// Withdraw + collect rewards
router.withdrawNative(poolId);
```

### Admin – create pool & configure rewards

```solidity
(uint256 poolId, ICrossGameRewardPool pool) =
    crossDeposit.createPool(IERC20(address(wcross)), 1 ether);

crossDeposit.addRewardToken(poolId, IERC20(address(usdt)));

// Anyone can fund rewards
usdt.transfer(address(pool), 1000 ether);
```

## 🏗️ Contracts

| Contract            | Role                                                                            |
|---------------------|---------------------------------------------------------------------------------|
| **WCROSS**          | Wraps native CROSS; router-only `deposit` / `withdraw`                         |
| **CrossGameReward**    | Factory (UUPS) – creates pools, manages reward tokens, sets pool status, withdraws unallocated rewards |
| **CrossGameRewardPool**| Individual pool (UUPS) – handles deposit/withdraw/claim with 3-state management (Active/Inactive/Paused), auto-settles removed rewards, handles zero-deposit deposits |
| **CrossGameRewardRouter** | User entry point – native and ERC-20 deposit with automatic wrap/unwrap     |

Key storage notes (pool):

```solidity
IERC20 public depositToken;
address public crossDeposit;
uint256 public minDepositAmount;
EnumerableSet.AddressSet private _rewardTokenAddresses;
EnumerableSet.AddressSet private _removedRewardTokenAddresses;
mapping(IERC20 => RewardToken) private _rewardTokenData;
mapping(address => mapping(IERC20 => UserReward)) public userRewards;
```

Removed reward tokens stay in `_removedRewardTokenAddresses` and are synchronised through
`_updateRemovedRewards` / `_claimRemovedRewards` whenever an account calls `_withdraw`.

## 🔑 Access Control Model

### CrossGameReward
| Role                        | Purpose                                                |
|-----------------------------|--------------------------------------------------------|
| `DEFAULT_ADMIN_ROLE` (owner) | Router assignment, pool implementation, upgrades       |
| `MANAGER_ROLE`              | Pool creation, reward tokens, pool status, withdrawals |

### CrossGameRewardPool
| Function Type        | Authority                | Description                              |
|---------------------|--------------------------|------------------------------------------|
| `onlyOwner()`       | CrossGameReward's owner     | Upgrade authorization                    |
| `onlyRewardRoot()` | CrossGameReward contract    | Reward management, pool status, withdraw |
| `depositFor/withdrawFor` | Router (verified)      | deposit/withdraw on behalf of users         |

**Key Changes:**
- Removed AccessControlDefaultAdminRules, simplified to modifier-based access
- All pool management functions callable only through CrossGameReward contract
- IERC5313 compliant (`owner()` function)

## 📊 Reward mechanics

### Core principles
- Uses `rewardPerToken` accumulation (`PRECISION = 1e18`) keeping gas cost constant
- Anyone can fund rewards by transferring tokens to the pool
- During deposit, use `claimReward(token)` / `claimRewards()` to collect active token rewards

### Reward queries
- `pendingRewards(user)`: Returns all active reward tokens and pending amounts `(address[] tokens, uint[] rewards)`
- `pendingReward(user, token)`: Query pending reward for a specific token `uint amount`

### Zero-deposit protection
- Rewards deposited when `totalDeposited=0` are classified as `withdrawableAmount`
- Protects the first depositor from receiving these unallocated rewards
- Owner can recover via `CrossGameReward.withdrawFromPool()`

### Removed token settlement
- Token balance at removal time is frozen as `distributedAmount`
- Users can still `claimReward(removedToken)` to collect these rewards
- New deposits after removal are added to `withdrawableAmount` for owner recovery

## 🔒 Security layers

1. ReentrancyGuardTransient (EIP-1153)
2. SafeERC20 transfers
3. Simplified access control (Owner/RewardRoot modifiers)
4. 3-state pool management (Active/Inactive/Paused)
5. UUPS upgrade gates
6. Custom errors for gas efficiency
7. Router caller verification
8. Zero-deposit reward protection

## 📚 Documentation

- [Architecture](overview/en/01_architecture.md)
- [Reward Mechanism](overview/en/02_reward_mechanism.md)
- [Security & Testing](overview/en/03_security_and_testing.md)
- [Test Guide](test/README.md)

## 🧪 Testing

```bash
forge test                    # full suite
forge test --match-contract CrossGameReward
forge test --gas-report
```

**Current coverage:** 212 tests across 11 suites:

| Suite                          | Tests |
|--------------------------------|-------|
| WCROSS                         | 10    |
| CrossGameReward                   | 33    |
| CrossGameRewardRouter             | 28    |
| CrossGameRewardPoolDeposit        | 18    |
| CrossGameRewardPoolRewards        | 27    |
| CrossGameRewardPoolAdmin          | 34    |
| CrossGameRewardPoolIntegration    | 11    |
| CrossGameRewardPoolPendingRewards | 9     |
| CrossGameRewardPoolSecurity       | 21    |
| CrossGameRewardPoolEdgeCases      | 12    |
| FullIntegration                | 9     |
| **Total**                      | **212** |

## 🔄 Upgrades

```solidity
// CrossGameReward upgrade
CrossGameReward newImpl = new CrossGameReward();
crossDeposit.upgradeToAndCall(address(newImpl), "");

// Pool upgrade
CrossGameRewardPool newPoolImpl = new CrossGameRewardPool();
pool.upgradeToAndCall(address(newPoolImpl), "");

// Router replacement
CrossGameRewardRouter newRouter = new CrossGameRewardRouter(address(crossDeposit));
crossDeposit.setRouter(address(newRouter));
```

## ⚙️ Operational notes

- Protect admin keys (router assignment, upgrades) with a multisig or governance module
- `setPoolStatus(poolId, status)`: 0=Active, 1=Inactive (claim/withdraw only), 2=Paused (all operations stopped)
- Removed reward tokens can be claimed individually via `claimReward(removedToken)`
- Zero-deposit deposits can be recovered via `withdrawFromPool`

## 📜 License

MIT

## 🔗 References

- [OpenZeppelin Contracts](https://docs.openzeppelin.com/contracts/)
- [Foundry Book](https://book.getfoundry.sh/)
- [Protocol docs](overview/README.md)
