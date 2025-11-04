# Security & Testing

## 🛡️ Defense-in-Depth Layers
```
Layer 1: ReentrancyGuardTransient
Layer 2: SafeERC20
Layer 3: AccessControlDefaultAdminRules
Layer 4: Pausable
Layer 5: UUPS upgrade gates
Layer 6: Custom errors
Layer 7: Router caller validation
```

---

## 🔒 Security Features

### ReentrancyGuardTransient
- Protects stake/unstake/claim entry points (native + router paths).
- Uses EIP-1153 transient storage; cheap and wipe-on-return.

### SafeERC20 Everywhere
```solidity
stakingToken.safeTransferFrom(msg.sender, address(this), amount);
stakingToken.safeTransfer(user, amount);
```
- Handles non-standard ERC-20s, validates return data, prevents silent failures.

### Role Matrix
- `CrossStaking`
  - `DEFAULT_ADMIN_ROLE`: upgrades, router changes.
  - `MANAGER_ROLE`: pool creation, reward token management, pause control.
- `CrossStakingPool`
  - `DEFAULT_ADMIN_ROLE`: resolves to factory admin via `owner()`.
  - `STAKING_ROOT_ROLE`: granted to factory for pool management.
  - `REWARD_MANAGER_ROLE`, `PAUSER_ROLE` available for delegation.

### Pausable
- `setPoolActive(poolId, false)` pauses staking, unstaking, claiming.
- Resume with `setPoolActive(poolId, true)` which calls `unpause`.

### Upgrade Gates
- `_authorizeUpgrade` restricted to `DEFAULT_ADMIN_ROLE` in both factory and pools.
- Storage gaps reserved for forward compatibility.

### Removed Reward Token Settlement
- Removed tokens move to `_removedRewardTokenAddresses` while remaining claimable.
- `_unstake` triggers `_updateRemovedRewards` and `_claimRemovedRewards`, so every outstanding reward is paid when a user exits.
- Active stakes still rely on `claimReward` / `claimRewards`; these routes operate on the active token set only.
- Regression tests `testRemovedRewardTokenClaimedOnUnstake` and `testClaimRemovedRewardAfterUnstakeDoesNotRevert` cover the flow (`src/CrossStakingPool.sol`).

---

## 🧪 Test Suite
- Foundry-based: 9 files / 159 test cases (`forge test`, 2025-11-03).
- Categories:
  - **Functional**: staking flows, reward accrual, view functions.
  - **Integration**: end-to-end journeys, multi-pool coordination.
  - **Security**: reentrancy attempts, role enforcement, invariant checks.
- Helpers: `_userStake`, `_depositReward`, `_warpDays` enable scenario coverage.

---

## 🗝️ Operational & Governance Notes
- `DEFAULT_ADMIN_ROLE` holders can replace the router, swap pool implementations, or upgrade contracts. Transition to a multi-sig or governance module before production.
- Pausing freezes claims and withdrawals. Define emergency procedures that either unpause for exits or provide alternative withdrawal paths.
- `setPoolActive` reverts if called repeatedly with the same status because OZ `pause/unpause` are not idempotent—check status before toggling.

---

## ✅ Summary
- 159/159 tests passing as of 2025-11-03 (Foundry).
- Layered security controls built on well-audited OpenZeppelin modules.
- Removed-reward locking risk is mitigated via automatic settlement on unstake.

See also: [../test/README.md](../test/README.md)
