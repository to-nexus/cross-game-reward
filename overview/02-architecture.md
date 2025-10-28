# System Architecture

## Architecture Overview

Cross Staking system is designed with a hierarchical and modular architecture.

```
┌─────────────────────────────────────────────────────────────┐
│                       User Layer                              │
│  (End users using Native CROSS, Frontend applications)        │
└─────────────────────────────────────────────────────────────┘
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                      Router Layer                             │
│                    StakingRouter                              │
│  • Native CROSS handling                                      │
│  • User convenience functions                                 │
│  • Batch queries                                              │
└─────────────────────────────────────────────────────────────┘
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                    Protocol Layer                             │
│                   StakingProtocol                             │
│  • Project Factory                                            │
│  • Central Management                                         │
│  • Configuration                                              │
└─────────────────────────────────────────────────────────────┘
                            ▼
┌──────────────────────────┬──────────────────────────────────┐
│    Project A             │       Project B                   │
├──────────┬───────────────┼──────────┬───────────────────────┤
│ Staking  │  Reward       │ Staking  │  Reward               │
│  Pool    │   Pool        │  Pool    │   Pool                │
└──────────┴───────────────┴──────────┴───────────────────────┘
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                      Token Layer                              │
│                        WCROSS                                 │
│  • Native CROSS ↔ ERC20 conversion                           │
└─────────────────────────────────────────────────────────────┘
```

## Contract Hierarchy

### 1. Interface Layer (interfaces/)

All core contracts are defined through interfaces.

```
IStakingProtocol         Protocol factory interface
IStakingPool             Staking pool interface
IRewardPool              Reward pool interface
IStakingPoolCode         StakingPool code contract interface
IRewardPoolCode          RewardPool code contract interface
```

**Interface Roles:**
- Clarify inter-contract dependencies
- Enable upgradability
- Facilitate testing and mocking

### 2. Library Layer (libraries/)

Pure logic separated into libraries for reusability and gas efficiency.

#### PointsLib
```solidity
// Points calculation pure functions
- calculatePoints()           // Basic points calculation
- calculateAggregatedPoints() // Aggregated points calculation
- calculateProRata()          // Proportional distribution
```

**Features:**
- Composed of `internal pure` functions
- Inline optimization possible
- Provides precision constant (1e6)

#### SeasonLib
```solidity
// Season validation and helper functions
- isSeasonActive()        // Check season active status
- isSeasonEnded()         // Check season ended
- validateSeasonBlocks()  // Validate season blocks
- isBlockInSeason()       // Check if block in season range
- calculateEffectiveStart() // Calculate effective start block
```

**Features:**
- Encapsulates season-related business logic
- Consistent season state management

### 3. Base Contract Layer (base/)

Common logic extracted into abstract contracts for code reusability.

#### CrossStakingBase

```solidity
abstract contract CrossStakingBase is 
    AccessControlDefaultAdminRules,
    ReentrancyGuardTransient
```

**Provided Features:**
- Common error definitions
- Input validation functions
- Access control foundation
- Reentrancy protection foundation

**Inheritance:**
```
CrossStakingBase
    ├─ StakingPoolBase
    └─ RewardPoolBase
```

#### StakingPoolBase

```solidity
abstract contract StakingPoolBase is 
    IStakingPool,
    CrossStakingBase
```

**Core Structure:**

1. **State Variables**
```solidity
struct StakePosition {
    uint balance;           // Staked balance
    uint points;            // Accumulated points
    uint lastUpdateBlock;   // Last update block
}

struct Season {
    uint seasonNumber;      // Season number
    uint startBlock;        // Start block
    uint endBlock;          // End block
    bool isFinalized;       // Finalization status
    uint totalPoints;       // Total points (cached)
    uint seasonTotalStaked; // Total staked during season
    uint lastAggregatedBlock; // Last aggregation block
    uint aggregatedPoints;  // Aggregated points
}

struct UserSeasonData {
    uint points;           // User season points
    uint balance;          // User season balance
    uint joinBlock;        // Join block
    bool claimed;          // Claim status
    bool finalized;        // Snapshot completed
}
```

2. **Core Function Flow**

**Staking Flow:**
```
stake() / stakeFor()
    ↓
_stakeFor()
    ↓
_beforeStake() [Hook]
    ↓
Points calculation and update
    ↓
Token transfer (SafeERC20)
    ↓
Aggregation update
    ↓
_afterStake() [Hook]
    ↓
Addon notification (safe call)
```

**Withdrawal Flow:**
```
withdrawAll() / withdrawAllFor()
    ↓
_withdrawAll()
    ↓
_beforeWithdraw() [Hook]
    ↓
Aggregation update
    ↓
Token transfer (SafeERC20)
    ↓
_afterWithdraw() [Hook]
    ↓
Addon notification (safe call)
```

**Season Transition Flow:**
```
_ensureSeason()
    ↓
Check current block
    ↓
Season 0? → _startFirstSeason()
    ↓
Season ended? → _rolloverSeason()
    ↓
_finalizeSeasonAggregation()
    ↓
Addon notification (season end)
    ↓
Create new season
```

3. **Aggregation System**

**Purpose:** Calculate total season points with O(1) complexity

**Mechanism:**
```
Aggregated points = Σ(total staked × time interval)

Time t1: Total staked 100, aggregation 0
  ↓ 10 blocks elapsed
Time t2: aggregation += 100 × 10 = 1000
  ↓ Stake +50 (total 150)
  ↓ 20 blocks elapsed
Time t3: aggregation += 150 × 20 = 4000
  Total aggregation = 1000 + 4000 = 5000
```

**Functions:**
- `_updateSeasonAggregation()`: Incremental update
- `_finalizeSeasonAggregation()`: Final aggregation at season end

4. **Lazy Snapshot System**

**Purpose:** Gas-efficient user season data management

**Mechanism:**
```
Staking → Record current season only
Season transition → Auto transition, no user data processing
Claim/Staking → _ensureUserAllPreviousSeasons()
                  → Retroactively calculate previous seasons
                  → Store in UserSeasonData
                  → Set finalized = true
```

**Functions:**
- `_ensureUserAllPreviousSeasons()`: Process all previous seasons
- `_ensureUserSeasonSnapshot()`: Snapshot specific season

#### RewardPoolBase

```solidity
abstract contract RewardPoolBase is 
    IRewardPool,
    CrossStakingBase
```

**Core Structure:**

1. **State Variables**
```solidity
mapping(uint => mapping(address => uint)) seasonRewards;
// Total rewards per season per token

mapping(uint => mapping(address => uint)) seasonClaimed;
// Claimed rewards per season per token

mapping(address => mapping(uint => mapping(address => bool))) 
    hasClaimedSeasonReward;
// Claim status per user per season per token
```

2. **Reward Distribution Flow**

```
payUser()
    ↓
Check claim status (prevent double-claiming)
    ↓
Verify points
    ↓
Calculate base reward (proportional distribution)
    ↓
_calculateBonusReward() [Hook]
    ↓
Check balance
    ↓
_beforePayUser() [Hook]
    ↓
Token transfer (SafeERC20)
    ↓
_afterPayUser() [Hook]
```

3. **Hook System**

```solidity
// Extension points
function _beforeFundSeason(uint season, address token, uint amount)
function _afterFundSeason(uint season, address token, uint amount)
function _calculateBonusReward(...) returns (uint bonusAmount)
function _beforePayUser(address user, uint season, address token, uint amount)
function _afterPayUser(address user, uint season, address token, uint amount)
```

### 4. Implementation Contract Layer

#### StakingProtocol

**Role:** Project factory and central management

**Structure:**
```solidity
struct ProjectInfo {
    address stakingPool;
    address rewardPool;
    string name;
    bool isActive;
    uint createdAt;
    address creator;
    address admin;
}

mapping(uint => ProjectInfo) projects;
mapping(string => uint) projectIDByName;
mapping(address => uint[]) projectsByAddress;
```

**Code Contract Pattern + CREATE2:**
```
StakingProtocol
    ↓
stakingPoolCodeContract.code()  // Get creation code
    ↓
salt = keccak256(projectName, "StakingPool")  // Generate predictable salt
    ↓
create2(code + constructorArgs, salt)  // Deploy with CREATE2
    ↓
StakingPool instance created (predictable address)
```

**Advantages:**
- Avoids code size limit
- Optimizes deployment cost
- Facilitates upgrades
- **Address predictability** (CREATE2)
  - Pre-calculate address from project name
  - Ensures same address on cross-chain deployment
  - Frontend can display address before deployment

#### StakingPool

Implements `StakingPoolBase` with:
- Pausable functionality
- Reward claim integration
- Advanced view functions

#### RewardPool

Implements `RewardPoolBase` with:
- StakingPool connection
- Emergency token recovery
- Pre-deposit season handling

#### StakingRouter

**Role:** Native token convenience interface

**Features:**
- Automatic WCROSS wrapping/unwrapping
- Batch operations
- Reentrancy protection

#### StakingViewer

**Role:** Unified view functions

**Features:**
- Virtual season calculations
- Batch queries
- Gas-free read operations

#### WCROSS

**Role:** Native CROSS wrapper

**Features:**
- Standard WETH-style implementation
- 1:1 wrapping ratio
- Full ERC20 compatibility

## Design Patterns

### 1. Factory Pattern
- StakingProtocol creates project-specific pools
- Deterministic addresses with CREATE2
- Efficient deployment

### 2. Code Contract Pattern
- Separate bytecode storage
- Reduces factory size
- Enables upgradability

### 3. Template Method Pattern
- Hook system in base contracts
- Extensibility without modification
- Clean separation of concerns

### 4. Lazy Evaluation
- Deferred snapshot calculation
- Gas optimization
- On-demand processing

### 5. Aggregation Pattern
- O(1) total points calculation
- Incremental updates
- Efficient for large user bases

## Security Architecture

### Access Control Layers

1. **Protocol Admin**
   - Global settings management
   - 3-day timelock for role transfer
   - Can set project admins

2. **Project Admin**
   - Project-specific settings
   - Pool configuration
   - Emergency pause

3. **Contract Roles**
   - STAKING_POOL_ROLE
   - REWARD_POOL_ROLE
   - Limited to specific functions

### Protection Mechanisms

1. **ReentrancyGuardTransient**
   - EIP-1153 transient storage
   - 30% gas savings vs traditional
   - Protects all state-changing functions

2. **Pausable**
   - Emergency stop capability
   - Admin-only control
   - Protects staking/unstaking

3. **SafeERC20**
   - Handles non-standard tokens
   - Prevents transfer failures
   - Used for all token operations

4. **Checks-Effects-Interactions**
   - State changes before external calls
   - Prevents reentrancy
   - Standard security pattern

## Gas Optimization Strategies

1. **O(1) Aggregation**
   - Constant-time total points
   - No user iteration
   - Scalable to millions of users

2. **Lazy Snapshots**
   - Deferred user data processing
   - Only process when needed
   - Reduces rollover costs

3. **Transient Storage**
   - EIP-1153 for reentrancy guard
   - 30% gas savings
   - No permanent storage cost

4. **Unchecked Arithmetic**
   - Safe overflow cases
   - 5-10% gas savings
   - Carefully audited

5. **Immutable Variables**
   - Compile-time constants
   - No storage reads
   - Reduced gas costs

6. **Custom Errors**
   - 15-20% gas savings vs strings
   - Better error information
   - ABI-encoded data

