# Cross Staking Protocol v1.0

> Season-based blockchain staking protocol

## Overview

Cross Staking Protocol is a decentralized staking platform with a season-based reward distribution system. It enables creation of independent staking pools for each project, providing fair reward distribution per season.

### Key Features

- ⏱️ **Season-based System**: Clear reward periods defined by block ranges
- 🎯 **Points System**: Fair reward calculation based on staking amount × time
- 🏭 **Project Independence**: Factory pattern creates isolated pools per project
- 🔄 **Native Token Support**: Automatic WCROSS wrapping for user convenience
- 🔐 **Enhanced Security**: Reentrancy Guard, Access Control, Pausable patterns
- ⚡ **Gas Optimized**: Custom errors, storage optimization, transient storage (EIP-1153)

## Architecture

```
┌─────────────────────────────────────────────────────────┐
│                   StakingProtocol                       │
│              (Factory & Global Manager)                  │
└────────────────────┬────────────────────────────────────┘
                     │ CREATE2
            ┌────────┴────────┐
            │                 │
    ┌───────▼──────┐   ┌─────▼────────┐
    │ StakingPool  │───│ RewardPool   │
    │ (Project 1)  │   │ (Project 1)  │
    └──────────────┘   └──────────────┘
    
    ┌──────────────┐   ┌──────────────┐
    │ StakingPool  │───│ RewardPool   │
    │ (Project 2)  │   │ (Project 2)  │
    └──────────────┘   └──────────────┘
    
┌──────────────┐    ┌──────────────┐
│StakingRouter │    │StakingViewer │
│ (TX Handler) │    │(View Queries)│
└──────────────┘    └──────────────┘
```

## Core Contracts

### StakingProtocol (Factory)
Creates project-specific staking pools and manages global settings using CREATE2 for deterministic addresses.

### StakingPool
- Token staking and withdrawal
- Automatic season rollover
- Points calculation and aggregation
- Reward claims

### RewardPool
- Reward token deposits
- Reward distribution
- Season-based token management

### StakingRouter
- Native CROSS ↔ WCROSS automatic conversion
- User-friendly interface for native token operations

### StakingViewer
- Unified view functions
- Virtual season calculations
- Batch query support

### WCROSS
- ERC20 wrapper for native CROSS token
- 1:1 wrapping ratio
- Standard WETH-style implementation

## Technical Implementation

### Season System

**Block-based Seasons**
- Each season has fixed block duration
- Automatic rollover support (up to 50 seasons)
- Lazy snapshot for gas efficiency

**Season Lifecycle**
```solidity
Season 1: [startBlock, endBlock]
- Users stake and earn points
- Points = balance × time × PRECISION / timeUnit
- At endBlock: Season can be finalized
- After finalize: Rewards claimable

Season 2: Auto-starts after Season 1 ends
```

### Points System

**Calculation Formula**
```solidity
points = (balance × timeElapsed × POINTS_PRECISION) / timeUnit

Where:
- balance: Staked amount
- timeElapsed: (currentBlock - lastUpdateBlock) × blockTime
- POINTS_PRECISION: 1e6 (6 decimal places)
- timeUnit: Configurable (default: 1 hour)
```

**O(1) Aggregation**
- Incremental total points update
- No iteration over all users
- Gas-efficient design

**Lazy Snapshot**
- User points finalized only when needed
- Reduces gas costs during rollovers
- Maintains accuracy for reward distribution

### Pre-deposit Feature

Allows users to stake before Season 1 starts:

```solidity
Timeline:
[preDepositStartBlock] → [firstSeasonStartBlock] → [Season 1 End]
     |                        |                         |
   Staking allowed      Points start accumulating   Season ends
```

**Key Points**
- Only applicable to Season 1
- Points accumulate from season start block
- Optional feature (disabled if preDepositStartBlock = 0)

### Reward Distribution

**Proportional Distribution**
```solidity
userReward = (totalReward × userPoints) / totalPoints
```

**Multi-token Support**
- Multiple reward tokens per season
- Independent tracking per token
- Prevents double-claiming

### Security Patterns

**Reentrancy Protection**
- `ReentrancyGuardTransient` (EIP-1153)
- Transient storage for 30% gas savings
- Protects all state-changing functions

**Access Control**
- `AccessControlDefaultAdminRules`
- 3-day timelock for admin role transfers
- Role-based permissions:
  - `DEFAULT_ADMIN_ROLE`: Protocol admin
  - `STAKING_POOL_ROLE`: StakingPool contract
  - `REWARD_POOL_ROLE`: RewardPool contract

**Pausable**
- Emergency stop mechanism
- Only admin can pause/unpause
- Protects staking/unstaking functions

**Safe ERC20**
- Uses OpenZeppelin's `SafeERC20`
- Handles non-standard ERC20 tokens
- Prevents transfer failures

## Usage Examples

### Staking

```solidity
// Stake with Native CROSS (via Router)
stakingRouter.stake{value: 5 ether}(projectID);

// Stake with WCROSS directly
wcross.approve(address(stakingPool), 5 ether);
stakingPool.stake(5 ether);
```

### Withdrawal

```solidity
// Withdraw to Native CROSS (via Router)
stakingRouter.unstake(projectID);

// Withdraw WCROSS directly
stakingPool.withdrawAll();
```

### Reward Claims

```solidity
// Claim single season
stakingPool.claimSeason(seasonNumber, rewardTokenAddress);

// Claim multiple seasons (via Router)
uint[] memory seasons = [1, 2, 3];
address[] memory tokens = [token1, token2, token3];
stakingRouter.claimMultipleRewards(projectID, seasons, tokens);
```

### View Functions

```solidity
// Get current points
uint points = stakingViewer.getUserPoints(projectID, userAddress);

// Get season information
(uint season, uint startBlock, uint endBlock, uint blocksElapsed) = 
    stakingViewer.getSeasonInfo(projectID);

// Get expected reward
uint expectedReward = stakingViewer.getClaimableReward(
    projectID, userAddress, seasonNumber, rewardTokenAddress
);
```

## Project Structure

```
src/
├── base/                    # Abstract contracts
│   ├── CrossStakingBase.sol     # Common base with access control
│   ├── StakingPoolBase.sol      # Core staking logic
│   └── RewardPoolBase.sol       # Core reward logic
├── interfaces/              # Contract interfaces
│   ├── IStakingPool.sol
│   ├── IRewardPool.sol
│   └── IStakingProtocol.sol
├── libraries/               # Pure logic libraries
│   ├── PointsLib.sol           # Points calculation
│   └── SeasonLib.sol           # Season validation
├── StakingProtocol.sol      # Factory contract
├── StakingPool.sol          # Staking pool implementation
├── RewardPool.sol           # Reward pool implementation
├── StakingRouter.sol        # Native token router
├── StakingViewer.sol        # View functions aggregator
└── WCROSS.sol              # Wrapped CROSS token

test/
├── BaseTest.sol            # Test base setup
├── Staking.t.sol           # Staking tests
├── Season.t.sol            # Season tests
├── Points.t.sol            # Points tests
├── Rewards.t.sol           # Reward tests
├── MultiPool.t.sol         # Multi-project tests
├── Advanced.t.sol          # Advanced scenarios
├── Integrated.t.sol        # Integration tests
└── Fuzz.t.sol              # Fuzz tests
```

## Testing

### Test Coverage

- **Total Tests**: 68
- **Pass Rate**: 100%
- **Core Logic Coverage**: 100%

### Test Categories

1. **Staking Tests** (9 tests)
   - Basic stake/unstake operations
   - Multiple users and stakes
   - Minimum stake requirements
   - Edge cases

2. **Season Tests** (7 tests)
   - Season rollover
   - Points snapshots
   - Multiple seasons
   - Season information queries

3. **Points Tests** (9 tests)
   - Points calculation accuracy
   - Time-based accumulation
   - Proportional distribution
   - Snapshot mechanisms

4. **Reward Tests** (8 tests)
   - Proportional distribution
   - Multi-season rewards
   - Claim prevention
   - Token recovery

5. **Fuzz Tests** (13 tests)
   - Random input validation
   - Overflow protection
   - Edge case handling

### Running Tests

```bash
# Run all tests
forge test

# Run with gas report
forge test --gas-report

# Run specific test file
forge test --match-contract StakingTest
```

## Security

### Applied Security Patterns

- ✅ **ReentrancyGuardTransient** (EIP-1153)
- ✅ **AccessControlDefaultAdminRules** (3-day timelock)
- ✅ **Pausable Pattern**
- ✅ **SafeERC20**
- ✅ **Custom Errors** (gas efficient)
- ✅ **Checks-Effects-Interactions Pattern**

### Gas Optimizations

1. **Custom Errors**: 15-20% savings vs string errors
2. **Transient Storage**: 30% savings in reentrancy guard
3. **Unchecked Arithmetic**: 5-10% savings where safe
4. **Immutable Variables**: Reduces storage access costs
5. **O(1) Aggregation**: Constant-time operations

### Audit Status

- ⏳ Internal audit: Completed
- ⏳ External audit: Pending

## Documentation

- [Architecture Overview](overview/architecture.md)
- [Core Concepts](overview/concepts.md)
- [Contract Details](overview/contracts.md)
- [Workflows](overview/workflows.md)
- [Technical Implementation](overview/technical.md)
- [Pre-deposit Guide](overview/predeposit.md)
- [Test Documentation](overview/tests.md)

## Development

### Prerequisites

```bash
# Install Foundry
curl -L https://foundry.paradigm.xyz | bash
foundryup

# Install dependencies
forge install
```

### Compile

```bash
forge build
```

### Test

```bash
# Run all tests
forge test

# With gas report
forge test --gas-report

# With coverage
forge coverage
```

## License

MIT License

---

**v1.0.0** - Ready for Audit
