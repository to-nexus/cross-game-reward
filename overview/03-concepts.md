# Core Concepts

## 1. Points System

### What are Points?

Points are a measurement unit that **fairly reflects staking amount and duration**.

### Points Calculation Formula

```
points = (staking amount × elapsed time × POINTS_PRECISION) / time unit

Where:
- Staking amount: Number of tokens staked
- Elapsed time: (end block - start block) × block time (seconds)
- POINTS_PRECISION: 1e6 (precision)
- Time unit: Default 3600 seconds (1 hour)
```

### Example Calculation

**Scenario:**
- Alice stakes 100 CROSS
- 720 blocks elapse (1 sec/block = 720 seconds)
- Time unit: 3600 seconds (1 hour)

**Calculation:**
```
points = (100 × 720 × 1e6) / 3600
      = 72,000 × 1e6 / 3600
      = 20,000,000
      = 20 points (in 1e6 units)
```

**Interpretation:**
- 100 CROSS staked for 0.2 hours (12 minutes)
- 100 CROSS × 0.2 hours = 20 points

### Precision

**Why multiply by 1e6?**
- Solidity doesn't support decimals
- Multiply by 1e6 = 1,000,000 to represent 6 decimal places
- Example: 0.123456 points = 123,456 (internal representation)

### Time Unit Adjustment

**Default: 1 hour (3600 seconds)**
- 1 CROSS staked for 1 hour = 1 point

**Adjustable:**
```solidity
setPointsTimeUnit(1 days)  // 1 CROSS × 1 day = 1 point
setPointsTimeUnit(1 weeks) // 1 CROSS × 1 week = 1 point
```

### Points Accumulation Method

**Incremental Update:**
```
Block 100: Alice stakes 100 CROSS
    points = 0
    lastUpdateBlock = 100

Block 200: Points update
    additional points = (100 × (200-100) × blockTime) / timeUnit
    total points = 0 + additional points
    lastUpdateBlock = 200

Block 250: Alice stakes additional 50 CROSS
    additional points = (100 × (250-200) × blockTime) / timeUnit
    total points += additional points
    balance = 100 + 50 = 150 CROSS
    lastUpdateBlock = 250

Block 300: Points update
    additional points = (150 × (300-250) × blockTime) / timeUnit
    total points += additional points
```

### Points Usage

1. **Reward Distribution Ratio**
```
User reward = Total reward × (User points / Total points)
```

2. **Contribution Measurement**
```
Contribution = User points / Total points × 100%
```

3. **Fairness Guarantee**
```
- Late participants rewarded proportionally to participation period
- Higher stakes = higher rewards
- Considers both time × amount
```

## 2. Season System

### What is a Season?

A season is the **fundamental cycle for reward distribution**. Rewards are distributed independently for each season.

### Season Structure

```solidity
struct Season {
    uint seasonNumber;        // Season number (starts at 1)
    uint startBlock;          // Start block
    uint endBlock;            // End block
    bool isFinalized;         // Finalization status
    uint totalPoints;         // Total points (cached after finalize)
    uint seasonTotalStaked;   // Total staked during season
    uint lastAggregatedBlock; // Last aggregation block
    uint aggregatedPoints;    // Aggregated points (real-time)
}
```

**totalPoints vs aggregatedPoints:**
- `aggregatedPoints`: **Real-time cumulative value** that continuously increases during season
- `totalPoints`: **Finalized cached value** after season ends (immutable once set)
- **Separation reason**: 
  - Identify seasons not yet calculated with `totalPoints == 0` (lazy)
  - Gas efficiency: Reading totalPoints after finalize is cheaper than recalculating
  - Clarity: Distinguishes in-progress vs finalized state
  - Safety: Prevents value changes after finalize

### Season Lifecycle

```
[Waiting] → [Active] → [Ended] → [Finalized]

Waiting (Before Start):
- currentSeason = 0 or previous season finalized
- block.number < nextSeasonStartBlock

Active:
- isFinalized = false
- startBlock ≤ block.number ≤ endBlock
- Staking/withdrawal allowed
- Points accumulating

Ended:
- block.number > endBlock
- Not yet finalized = false
- Auto-rollover on next transaction

Finalized:
- isFinalized = true
- totalPoints confirmed
- Rewards claimable
- No further changes
```

### Season Transition Mechanism

**Automatic Rollover:**
```solidity
function _ensureSeason() internal {
    // When season 0
    if (currentSeason == 0) {
        if (block.number >= nextSeasonStartBlock) {
            _startFirstSeason();
        }
        return;
    }
    
    // Rollover all passed seasons (max 50)
    uint maxRollovers = 50;
    uint rolloversPerformed = 0;
    
    while (currentSeason > 0 && rolloversPerformed < maxRollovers) {
        Season storage current = seasons[currentSeason];
        if (block.number <= current.endBlock) break;
        
        _rolloverSeason();
        rolloversPerformed++;
    }
}
```

**When called?**
- On `stake()` call
- On `withdrawAll()` call
- On all state-changing transactions

**Multi-season Handling:**
- Automatically rollovers all seasons even after long inactivity
- Processes up to 50 seasons at once (considering gas limit)
- In practice, usually only 1-2 seasons rollover

### Season User Data

```solidity
struct UserSeasonData {
    uint points;        // Season points
    uint balance;       // Season balance
    uint joinBlock;     // Season join block
    bool claimed;       // Claim status
    bool finalized;     // Snapshot completed
}
```

## 3. Season Aggregation System

### Problem: Total Points Calculation Complexity

**Simple Method (O(N)):**
```
Total points = Σ(each staker's points)

Problem:
- Need to calculate 1000 times for 1000 stakers
- Very high gas cost
- Not practical
```

**Optimized Method (O(1)):**
```
Total points = Σ(total staked × time interval)

Advantages:
- O(1) regardless of staker count
- Efficient with incremental updates
- Gas cost savings
```

### Aggregation Principle

**Mathematical Basis:**
```
Sum of all users' points:
Σ(each user balance × time) = (sum of all user balances) × time
                            = total staked × time
```

**Example:**
```
Block 100:
- Alice: 50 CROSS
- Bob: 30 CROSS
- Total staked: 80 CROSS

Block 200 (100 blocks elapsed):
Individual calculation:
- Alice points: 50 × 100 = 5,000
- Bob points: 30 × 100 = 3,000
- Total points: 8,000

Aggregated calculation:
- Total points: 80 × 100 = 8,000 ✓ Same!
```

### Aggregation Update Points

**1. On Staking:**
```solidity
_updateSeasonAggregation(currentSeason);  // Aggregate up to now
currentSeasonData.seasonTotalStaked += amount;  // Increase total
```

**2. On Withdrawal:**
```solidity
_updateSeasonAggregation(currentSeason);  // Aggregate up to now
currentSeasonData.seasonTotalStaked -= amount;  // Decrease total
```

**3. On Season End:**
```solidity
_finalizeSeasonAggregation(seasonNum);  // Final aggregation
```

### Aggregation Calculation Logic

```solidity
function _updateSeasonAggregation(uint seasonNum) internal {
    Season storage season = seasons[seasonNum];
    
    // Skip if already up-to-date
    if (season.lastAggregatedBlock >= block.number) return;
    
    // If total staked is 0, only update time
    if (season.seasonTotalStaked == 0) {
        season.lastAggregatedBlock = block.number;
        return;
    }
    
    // Calculate incremental points
    uint additionalPoints = PointsLib.calculatePoints(
        season.seasonTotalStaked,
        season.lastAggregatedBlock,
        block.number,
        blockTime,
        pointsTimeUnit
    );
    
    // Add to aggregated points
    season.aggregatedPoints += additionalPoints;
    season.lastAggregatedBlock = block.number;
}
```

## 4. Lazy Snapshot System

### Purpose

Gas-efficient user season data management by deferring calculations.

### Mechanism

```
Staking → Record current season only
Season transition → Auto transition, no user data processing
Claim/Staking → _ensureUserAllPreviousSeasons()
                → Retroactively calculate previous seasons
                → Store in UserSeasonData
                → Set finalized = true
```

### When User Data is Snapshotted

**Trigger Events:**
1. User stakes in new season
2. User withdraws
3. User claims rewards

**Process:**
```
Check UserSeasonData[season].finalized
  ↓
If false → Calculate and store season data
  ↓
Calculate points from position data
  ↓
Store in UserSeasonData
  ↓
Set finalized = true
```

### Benefits

1. **Gas Efficiency**
   - No user iteration on season rollover
   - Only process when user actually needs it
   - O(1) rollover cost

2. **Scalability**
   - Supports millions of users
   - Gas cost independent of user count
   - No bottlenecks

3. **Fairness**
   - Accurate calculation for all users
   - No approximations
   - Retroactive calculation ensures correctness

## 5. Reward Distribution

### Proportional Distribution Formula

```
userReward = (totalSeasonReward × userPoints) / totalSeasonPoints
```

### Multi-token Support

- Multiple reward tokens per season
- Independent tracking per token
- Prevents double-claiming through mapping:
  ```solidity
  mapping(address => mapping(uint => mapping(address => bool))) hasClaimedSeasonReward;
  // user => season => token => claimed
  ```

### Reward Claim Process

```
1. User calls claimSeason(season, token)
   ↓
2. Check if season is finalized
   ↓
3. Check if already claimed
   ↓
4. Get user points for season
   ↓
5. Calculate proportional reward
   ↓
6. Transfer tokens to user
   ↓
7. Mark as claimed
```

### Unclaimed Reward Recovery

Project creators can recover unclaimed rewards after season ends:
```solidity
function recoverRemaining(uint season, address token) external
```

## 6. Pre-deposit Feature

### Concept

Allows users to stake before Season 1 starts, with points accumulating from the season start block.

### Timeline

```
[preDepositStartBlock] → [firstSeasonStartBlock] → [Season 1 End]
     |                        |                         |
   Staking allowed      Points start accumulating   Season ends
```

### Key Points

- **Season 1 exclusive**: Only applies to first season
- **Points timing**: Accumulate from season start block, not staking block
- **Optional**: Disabled if preDepositStartBlock = 0
- **Fairness**: All pre-deposit stakers start earning points simultaneously

### Implementation

```solidity
function _stakeFor(address user, uint amount, address from) internal virtual {
    _ensureSeason();
    
    if (currentSeason == 0) {
        // First season not yet created
        if (preDepositStartBlock > 0 && block.number >= preDepositStartBlock) {
            // ✅ Pre-deposit period: staking allowed
        } else {
            // ❌ No pre-deposit or before pre-deposit block
            require(block.number >= nextSeasonStartBlock, StakingPoolBaseNoActiveSeason());
        }
    }
    // ... staking logic
}
```

Points calculation for pre-deposit users:
```solidity
uint lastUpdate = position.lastUpdateBlock;
if (lastUpdate < current.startBlock && position.balance > 0) {
    // Pre-deposit case: calculate from season start block
    return PointsLib.calculatePoints(
        position.balance, 
        current.startBlock,  // ✅ From season start block
        block.number, 
        blockTime, 
        pointsTimeUnit
    );
}
```

