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
- Protects deposit/withdraw/claim entry points (native + router paths).
- Uses EIP-1153 transient storage; cheap and wipe-on-return.

### SafeERC20 Everywhere
```solidity
depositToken.safeTransferFrom(msg.sender, address(this), amount);
depositToken.safeTransfer(user, amount);
```
- Handles non-standard ERC-20s, validates return data, prevents silent failures.

### Role Matrix
- `CrossGameReward`
  - `DEFAULT_ADMIN_ROLE`: upgrades, router changes.
  - `MANAGER_ROLE`: pool creation, reward token management, pause control.
- `CrossGameRewardPool`
  - `DEFAULT_ADMIN_ROLE`: resolves to factory admin via `owner()`.
  - `REWARD_ROOT_ROLE`: granted to factory for pool management.
  - `REWARD_MANAGER_ROLE`, `PAUSER_ROLE` available for delegation.

### Pausable
- `setPoolActive(poolId, false)` pauses deposit, withdraw, claiming.
- Resume with `setPoolActive(poolId, true)` which calls `unpause`.

### Upgrade Gates
- `_authorizeUpgrade` restricted to `DEFAULT_ADMIN_ROLE` in both factory and pools.
- Storage gaps reserved for forward compatibility.

### Removed Reward Token Settlement
- Removed tokens move to `_removedRewardTokenAddresses` while remaining claimable.
- `_withdraw` triggers `_updateRemovedRewards` and `_claimRemovedRewards`, so every outstanding reward is paid when a user exits.
- Active deposits can use `claimReward` / `claimRewards` to claim rewards without withdrawing; these routes operate on the active token set only.
- Regression tests `testRemovedRewardTokenClaimedOnUndeposit` and `testClaimRemovedRewardAfterUndepositDoesNotRevert` cover the flow (`src/CrossGameRewardPool.sol`).

### Router Check
- Router can call `depositFor`, `withdrawFor`, `claimRewardsFor`, and `claimRewardFor` on behalf of users.
- `_checkDelegate` enforces that only the authorized router can call these `For` functions.

### WCROSS - WETH9 Pattern
- Follows the WETH9 standard: anyone can wrap/unwrap
- Open deposit/withdraw pattern
  ```solidity
  function deposit() public payable { ... }
  function withdraw(uint amount) public { ... }
  ```
- Security maintained by ERC20 transfer rules - only token owner can transfer
- Enables direct DEX integration and better composability

---

## 🧪 Test Suite
- Foundry-based: 9 test files / **233 test cases**.
- Test Coverage by Contract:
  - **CrossGameRewardRouter**: 39 tests
  - **CrossGameRewardPool**: 142 tests
  - **WCROSS**: 10 tests
  - **CrossGameReward**: 30 tests
  - **Integration Tests**: 12 tests
- Categories:
  - **Functional**: deposit flows, reward accrual, view functions, claim operations.
  - **Integration**: end-to-end journeys, multi-pool coordination.
  - **Security**: reentrancy attempts, role enforcement, invariant checks.
- Helpers: `_userDeposit`, `_depositReward`, `_warpDays` enable comprehensive scenario coverage.

---

## 🗝️ Operational & Governance Notes
- `DEFAULT_ADMIN_ROLE` holders can replace the router, swap pool implementations, or upgrade contracts. Transition to a multi-sig or governance module before production.
- Pausing freezes claims and withdrawals. Define emergency procedures that either unpause for exits or provide alternative withdrawal paths.
- `setPoolActive` reverts if called repeatedly with the same status because OZ `pause/unpause` are not idempotent—check status before toggling.

---

## ✅ Summary
- **233/233 tests passing** (Foundry).
- Layered security controls built on well-audited OpenZeppelin modules.
- Removed-reward locking risk is mitigated via automatic settlement on withdraw.
- Router claim wrapper enables reward claiming without full withdrawal.
- WCROSS WETH9 pattern improves composability while maintaining security.

See also: [../test/README.md](../test/README.md)
