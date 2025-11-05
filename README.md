# Cross Staking Protocol

Multi-pool staking for native CROSS and ERC-20 tokens

## 🎯 Overview

Cross Staking Protocol lets you launch and manage multiple staking pools under a single factory. It wraps native CROSS when needed, supports arbitrary ERC-20 rewards, and keeps accounting gas costs flat.

### Architecture at a glance

```
User (Native CROSS / ERC-20)
    │
    ▼
CrossStakingRouter ──► WCROSS (wrap/unwrap)
    │
    ▼
CrossStaking (factory, UUPS)
    │ creates
    ▼
CrossStakingPool × N (UUPS)
```

## ✨ Key features

- ✅ **Native CROSS support** – router auto-wraps and unwraps via WCROSS
- ✅ **Unlimited pools** – multiple pools per staking token
- ✅ **Multi-reward** – each pool can emit several ERC-20 rewards
- ✅ **O(1) accounting** – reward distribution uses a `rewardPerToken` accumulator
- ✅ **Upgradeable** – CrossStaking and pools follow the UUPS pattern
- ✅ **Simplified access control** – Owner and StakingRoot based permissions
- ✅ **3-state pool management** – Active/Inactive/Paused for granular control
- ✅ **Fair reward distribution** – Pre-staking rewards automatically marked as withdrawable
- ✅ **Removed reward settlement** – rewards for removed tokens are auto-claimed during unstake

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

### User – stake native CROSS

```solidity
// Approve once
wcross.approve(address(router), type(uint256).max);

// Stake
router.stakeNative{value: 100 ether}(poolId);

// Unstake + collect rewards
router.unstakeNative(poolId);
```

### Admin – create pool & configure rewards

```solidity
(uint256 poolId, ICrossStakingPool pool) =
    crossStaking.createPool(IERC20(address(wcross)), 1 ether);

crossStaking.addRewardToken(poolId, IERC20(address(usdt)));

// Anyone can fund rewards
usdt.transfer(address(pool), 1000 ether);
```

## 🏗️ Contracts

| Contract            | Role                                                                            |
|---------------------|---------------------------------------------------------------------------------|
| **WCROSS**          | Wraps native CROSS; router-only `deposit` / `withdraw`                         |
| **CrossStaking**    | Factory (UUPS) – creates pools, manages reward tokens, sets pool status, withdraws unallocated rewards |
| **CrossStakingPool**| Individual pool (UUPS) – handles stake/unstake/claim with 3-state management (Active/Inactive/Paused), auto-settles removed rewards, handles zero-stake deposits |
| **CrossStakingRouter** | User entry point – native and ERC-20 staking with automatic wrap/unwrap     |

Key storage notes (pool):

```solidity
IERC20 public stakingToken;
address public crossStaking;
uint256 public minStakeAmount;
EnumerableSet.AddressSet private _rewardTokenAddresses;
EnumerableSet.AddressSet private _removedRewardTokenAddresses;
mapping(IERC20 => RewardToken) private _rewardTokenData;
mapping(address => mapping(IERC20 => UserReward)) public userRewards;
```

Removed reward tokens stay in `_removedRewardTokenAddresses` and are synchronised through
`_updateRemovedRewards` / `_claimRemovedRewards` whenever an account calls `_unstake`.

## 🔑 Access Control Model

### CrossStaking
| Role                        | Purpose                                                |
|-----------------------------|--------------------------------------------------------|
| `DEFAULT_ADMIN_ROLE` (owner) | Router assignment, pool implementation, upgrades       |
| `MANAGER_ROLE`              | Pool creation, reward tokens, pool status, withdrawals |

### CrossStakingPool
| Function Type        | Authority                | Description                              |
|---------------------|--------------------------|------------------------------------------|
| `onlyOwner()`       | CrossStaking's owner     | Upgrade authorization                    |
| `onlyStakingRoot()` | CrossStaking contract    | Reward management, pool status, withdraw |
| `stakeFor/unstakeFor` | Router (verified)      | Stake/unstake on behalf of users         |

**Key Changes:**
- Removed AccessControlDefaultAdminRules, simplified to modifier-based access
- All pool management functions callable only through CrossStaking contract
- IERC5313 compliant (`owner()` function)

## 📊 Reward mechanics

### Core principles
- Uses `rewardPerToken` accumulation (`PRECISION = 1e18`) keeping gas cost constant
- Anyone can fund rewards by transferring tokens to the pool
- During staking, use `claimReward(token)` / `claimRewards()` to collect active token rewards

### Reward queries
- `pendingRewards(user)`: Returns all active reward tokens and pending amounts `(address[] tokens, uint[] rewards)`
- `pendingReward(user, token)`: Query pending reward for a specific token `uint amount`

### Zero-stake protection
- Rewards deposited when `totalStaked=0` are classified as `withdrawableAmount`
- Protects the first staker from receiving these unallocated rewards
- Owner can recover via `CrossStaking.withdrawFromPool()`

### Removed token settlement
- Token balance at removal time is frozen as `distributedAmount`
- Users can still `claimReward(removedToken)` to collect these rewards
- New deposits after removal are added to `withdrawableAmount` for owner recovery

## 🔒 Security layers

1. ReentrancyGuardTransient (EIP-1153)
2. SafeERC20 transfers
3. Simplified access control (Owner/StakingRoot modifiers)
4. 3-state pool management (Active/Inactive/Paused)
5. UUPS upgrade gates
6. Custom errors for gas efficiency
7. Router caller verification
8. Zero-stake reward protection

## 📚 Documentation

- [Architecture](overview/en/01_architecture.md)
- [Reward Mechanism](overview/en/02_reward_mechanism.md)
- [Security & Testing](overview/en/03_security_and_testing.md)
- [Test Guide](test/README.md)

## 🧪 Testing

```bash
forge test                    # full suite
forge test --match-contract CrossStaking
forge test --gas-report
```

**Current coverage:** 212 tests across 11 suites:

| Suite                          | Tests |
|--------------------------------|-------|
| WCROSS                         | 10    |
| CrossStaking                   | 33    |
| CrossStakingRouter             | 28    |
| CrossStakingPoolStaking        | 18    |
| CrossStakingPoolRewards        | 27    |
| CrossStakingPoolAdmin          | 34    |
| CrossStakingPoolIntegration    | 11    |
| CrossStakingPoolPendingRewards | 9     |
| CrossStakingPoolSecurity       | 21    |
| CrossStakingPoolEdgeCases      | 12    |
| FullIntegration                | 9     |
| **Total**                      | **212** |

## 🔄 Upgrades

```solidity
// CrossStaking upgrade
CrossStaking newImpl = new CrossStaking();
crossStaking.upgradeToAndCall(address(newImpl), "");

// Pool upgrade
CrossStakingPool newPoolImpl = new CrossStakingPool();
pool.upgradeToAndCall(address(newPoolImpl), "");

// Router replacement
CrossStakingRouter newRouter = new CrossStakingRouter(address(crossStaking));
crossStaking.setRouter(address(newRouter));
```

## ⚙️ Operational notes

- Protect admin keys (router assignment, upgrades) with a multisig or governance module
- `setPoolStatus(poolId, status)`: 0=Active, 1=Inactive (claim/unstake only), 2=Paused (all operations stopped)
- Removed reward tokens can be claimed individually via `claimReward(removedToken)`
- Zero-stake deposits can be recovered via `withdrawFromPool`

## 📜 License

MIT

## 🔗 References

- [OpenZeppelin Contracts](https://docs.openzeppelin.com/contracts/)
- [Foundry Book](https://book.getfoundry.sh/)
- [Protocol docs](overview/README.md)
