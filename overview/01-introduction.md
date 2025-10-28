# Cross Staking Protocol - Introduction

## Project Overview

Cross Staking Protocol is a blockchain-based season staking platform. It creates independent staking pools for each project and distributes fair rewards every season.

## Core Values

### 1. Fairness
- Points system proportional to time and amount
- Transparent reward distribution algorithm
- All calculations verifiable on-chain

### 2. Flexibility
- Independent staking pools per project
- Customizable season length
- Multiple reward token support

### 3. Efficiency
- Gas optimization (10-15% reduction)
- O(1) aggregation-based points calculation
- ReentrancyGuardTransient (EIP-1153)

### 4. Security
- OpenZeppelin verified libraries
- Role-based Access Control
- Pausable Pattern
- 3-day Admin Timelock

## Key Features

### Season-based System
- Clear season distinction based on blocks
- Automatic rollover (up to 50 seasons)
- Manual rollover support

### Points System
```
points = balance × timeElapsed × PRECISION / timeUnit
```
- Precision: 1e6 (6 decimal places)
- Real-time calculation
- Virtual season support

### Factory Pattern
- Independent pool creation per project
- Predictable addresses with CREATE2
- Gas savings with Code Contract pattern

### Native Token Support
- Automatic WCROSS wrapping/unwrapping
- Users only interact with Native CROSS
- StakingRouter handles all conversions

## Technology Stack

- **Solidity**: 0.8.28
- **Framework**: Foundry
- **Libraries**: OpenZeppelin Contracts 5.x
- **Standards**: ERC20, EIP-1153
- **Patterns**: Factory, Proxy, Template Method

## Architecture Overview

```
StakingProtocol (Factory)
    ↓
StakingPool (Per Project)
    ↓
RewardPool (Per Project)

StakingRouter (Native Token)
StakingViewer (Read-only)
```

## Use Cases

### 1. Project Teams
- Operate independent staking pools
- Run seasonal reward programs
- Provide community incentives

### 2. Stakers
- Stake $CROSS tokens
- Earn seasonal rewards
- Auto-compound participation

### 3. Developers
- Smart contract integration
- Frontend application development
- Data analytics tool building

## Version Information

**Current Version**: v1.0.0  
**Status**: Ready for Audit  
**License**: MIT

## Documentation Structure

1. [Architecture](./02-architecture.md) - System structure details
2. [Core Concepts](./03-concepts.md) - Seasons, Points, Rewards
3. [Contract Details](./04-contracts.md) - Individual contract explanations
4. [Workflows](./05-workflows.md) - Key processes
5. [Technical Implementation](./06-technical.md) - Implementation details

