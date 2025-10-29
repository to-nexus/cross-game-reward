# Cross Staking Protocol v1.0

> Season-based blockchain staking protocol with simplified points calculation

## Overview

Cross Staking Protocol is a decentralized staking platform with a season-based reward distribution system. It enables creation of independent staking pools for each project, providing fair reward distribution per season using simplified timestamp-based point calculations.

### Key Features

- ⏱️ **Time-based Seasons**: Clear reward periods defined by timestamps
- 🎯 **Simple Points**: 1 token × 1 second = 1 point (intuitive!)
- 🏭 **Project Independence**: Factory pattern creates isolated pools per project
- 🔄 **Native Token Support**: Automatic WCROSS wrapping for user convenience
- 🔐 **Enhanced Security**: Reentrancy Guard, Access Control, Pausable patterns
- ⚡ **Gas Optimized**: Custom errors, transient storage (EIP-1153), minimal calculations

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

## Points System

### Simple & Intuitive Formula

```solidity
points = balance × timeElapsed

Where:
- balance: Amount staked (in wei)
- timeElapsed: Time duration (in seconds)
- Result: token-seconds

Example:
- Stake 100 CROSS for 1 hour (3600 seconds)
- Points = 100e18 × 3600 = 3.6e23 (raw)
- Display = 3.6e23 / 1e18 = 360,000 token-seconds
```

### Display Formatting

**POINT_DECIMALS = 18**

```typescript
// Frontend
const rawPoints = 3.6e23;  // from contract
const displayPoints = rawPoints / 1e18;  // 360,000 token-seconds

// Or in Solidity for display
displayPoints = rawPoints / 10**PointsLib.POINT_DECIMALS;
```

### Why This Is Better

**Before (Complex)**:
```solidity
points = (balance × time × 1e6) / timeUnit
// Requires timeUnit parameter
// Division operation (more gas)
// Less intuitive
```

**After (Simple)**:
```solidity
points = balance × time
// No extra parameters
// No division (gas efficient)
// Crystal clear: 1 token × 1 second = 1 point
```

## Technical Implementation

### Season System

**Timestamp-based Seasons**
- Each season has fixed time duration (in seconds)
- Automatic rollover support (up to 50 seasons)
- Lazy snapshot for gas efficiency

**Season Lifecycle**
```solidity
Season 1: [startTime, endTime]
- Users stake and earn points
- Points = balance × time (directly!)
- At endTime: Season can be finalized
- After finalize: Rewards claimable

Season 2: Auto-starts after Season 1 ends
```

### O(1) Aggregation Algorithm

```solidity
totalPoints = aggregatedPoints + (totalStaked × timeSinceLastUpdate)

- Incremental updates on stake/unstake
- No iteration over users
- Constant time complexity
- Gas-efficient for any number of users
```

### Virtual Seasons

Supports season information queries without on-chain rollover:

```solidity
// Season 1 ended but not rolled over
getCurrentSeasonInfo() returns (1, startTime, endTime, duration)

// Frontend can display and calculate without transactions
previewClaim(1, user, token) returns (points, totalPoints, expectedReward)
```

### Pre-deposit Feature

Optional feature allowing staking before Season 1:

```solidity
preDepositStartTime = 1704067200  // 2024-01-01 00:00:00 UTC
firstSeasonStartTime = 1706745600  // 2024-02-01 00:00:00 UTC

// Users can stake between these times
// Points accumulate from firstSeasonStartTime
```

### Lazy Snapshot

User season data is finalized only when needed:

```solidity
// User A claims Season 1 → snapshot taken
// User B doesn't claim → no snapshot (gas saved)

// Snapshot taken on:
- Reward claim
- Manual snapshot call
- Season rollover with aggregation
```

## Usage Examples

### Basic Staking Flow

```solidity
// 1. User stakes 1000 WCROSS
wcross.approve(address(router), 1000e18);
router.stakeNative{value: 1000e18}(projectId);

// 2. Season progresses...
// Points accumulate: 1 token × 1 second = 1 point

// 3. Season ends and rewards are deposited
rewardToken.approve(address(rewardPool), 10000e18);
rewardPool.fundSeason(1, address(rewardToken), 10000e18);

// 4. User claims rewards
pool.claimSeason(1, address(rewardToken));

// 5. User unstakes
router.unstakeNative(projectId, 1000e18);
```

### Admin Operations

```solidity
// Create new staking pool
protocol.createProject(
    "Project Name",
    1 days,      // Season duration (86400 seconds)
    block.timestamp + 7 days,  // First season starts in 7 days
    0,           // Pool never ends
    0            // Pre-deposit disabled
);

// Emergency pool end
pool.setPoolEndTime(block.timestamp + 30 days);

// Manual season rollover
pool.rolloverSeason(10); // Rollover up to 10 seasons
```

### View Queries

```solidity
// Get current season info
(uint season, uint startTime, uint endTime, uint timeElapsed) = 
    pool.getCurrentSeasonInfo();

// Preview claim (before transaction)
(uint userPoints, uint totalPoints, uint expectedReward, bool claimed, bool canClaim) = 
    viewer.previewClaim(projectId, user, season, rewardToken);

// Get all claimable rewards
(address[] memory tokens, uint[] memory amounts) = 
    viewer.getClaimableRewards(projectId, user, season);
```

## Security Features

### Access Control
- **DEFAULT_ADMIN_ROLE**: Pool configuration
- **MANAGER_ROLE**: Season management
- **REWARD_POOL_ROLE**: Reward distribution
- **ROUTER_ROLE**: Batch operations

### Protection Mechanisms
- ✅ Reentrancy Guard (transient storage)
- ✅ Pausable (emergency stop)
- ✅ Minimum stake requirement (1 CROSS)
- ✅ Integer overflow protection (Solidity 0.8+)
- ✅ Safe ERC20 transfers
- ✅ Role-based access control

### Known Limitations
- Maximum 50 automatic rollovers per transaction
- Pool end time cannot be before current time
- Season duration must be > 0
- Points display decimals: 18 (for normalization)

## Deployment

### Prerequisites
- Foundry installed
- Private key configured
- RPC endpoint available

### Deploy Factory

```bash
# Set environment variables
export PRIVATE_KEY="0x..."
export RPC_URL="https://rpc.example.com"
export ADMIN_ADDRESS="0x..."

# Deploy
cd script
forge script Deploy.s.sol:Deploy --rpc-url $RPC_URL --broadcast
```

### Create First Project

```bash
forge script CreateProject.s.sol:CreateProject --rpc-url $RPC_URL --broadcast
```

### Configuration Files

See `script/` directory for deployment scripts and `.env` examples.

## Testing

```bash
# Run all tests
forge test

# Run specific test file
forge test --match-path test/Season.t.sol

# Run with gas report
forge test --gas-report

# Run with coverage
forge coverage
```

### Test Coverage
- ✅ 81 tests, 100% passing
- ✅ Unit tests for all core functions
- ✅ Integration tests for full workflows
- ✅ Fuzz tests for edge cases
- ✅ Multi-season scenarios
- ✅ Virtual season handling

## Gas Optimization

### Techniques Used
- Transient storage for reentrancy guard (EIP-1153)
- Custom errors instead of string reverts
- Simplified points calculation (no division!)
- Unchecked arithmetic where safe
- Storage variable packing
- Incremental aggregation (O(1))
- Lazy snapshot system

### Typical Gas Costs
| Operation | Gas Cost |
|-----------|----------|
| Stake | ~170,000 |
| Unstake | ~130,000 |
| Claim Reward | ~180,000 |
| Season Rollover | ~160,000 |
| View Functions | Free |

## License

MIT

## Documentation

For detailed documentation, see:
- [Overview](overview/README.md)
- [Architecture](overview/02-architecture.md)
- [Concepts](overview/03-concepts.md)
- [Contracts](overview/04-contracts.md)
- [Workflows](overview/05-workflows.md)
- [Technical Details](overview/06-technical.md)

## Support

For issues and questions:
- GitHub Issues: [github.com/to-nexus/cross-staking](https://github.com/to-nexus/cross-staking)
- Documentation: [docs](overview/)

## Version History

### v1.0.0
- Initial release with timestamp-based seasons
- Simplified points calculation: balance × time
- Points display decimals: 18
- O(1) aggregation algorithm
- Virtual season support
- Pre-deposit feature
- Multi-token rewards per season
- Gas optimizations

---

**Built with ❤️ using Solidity 0.8.28 and Foundry**
