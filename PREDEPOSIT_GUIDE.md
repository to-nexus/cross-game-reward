# Pre-deposit Feature Guide

## Overview

Pre-deposit is a feature that allows users to stake before Season 1 begins. When users stake through pre-deposit, their points begin accumulating from the Season 1 start block.

## Key Characteristics

### 1. Season 1 Exclusive Feature
- Pre-deposit applies **only to Season 1**
- From Season 2 onwards, only regular staking is available

### 2. Block-based Timing
```
Timeline:
[preDepositStartBlock] -----> [firstSeasonStartBlock] -----> [Season 1 End]
     |                              |                              |
   Pre-deposit starts           Season 1 starts              Season 1 ends
   (Staking allowed)       (Points accumulation begins)
```

### 3. Points Accumulation
- **Pre-deposit staking**: Points accumulate from **Season 1 start block**
- **Regular staking**: Points accumulate from **staking block**

## Configuration

### Project Creation Setup

```solidity
// When calling StakingProtocol.createProject()
function createProject(
    string calldata projectName,
    uint seasonBlocks,
    uint firstSeasonStartBlock,  // Season 1 start block
    uint poolEndBlock,
    address projectAdmin,
    uint preDepositStartBlock    // Pre-deposit start block (0 to disable)
) external returns (uint projectID, address stakingPool, address rewardPool)
```

### Example

```solidity
// Current block: 1000
// Pre-deposit starts: Block 1100
// Season 1 starts: Block 1200

protocol.createProject(
    "MyProject",
    100,        // seasonBlocks: 100 blocks
    1200,       // firstSeasonStartBlock: Season 1 starts at block 1200
    0,          // poolEndBlock: Infinite (0)
    msg.sender, // projectAdmin
    1100        // preDepositStartBlock: Pre-deposit allowed from block 1100
);
```

## Smart Contract Logic

### StakingPoolBase._stakeFor()

```solidity
function _stakeFor(address user, uint amount, address from) internal virtual {
    // ...
    
    _ensureSeason();
    
    // Season active check
    if (currentSeason == 0) {
        // First season not yet created
        if (preDepositStartBlock > 0 && block.number >= preDepositStartBlock) {
            // ✅ Pre-deposit period: Staking allowed
        } else {
            // ❌ No pre-deposit or before pre-deposit block
            require(block.number >= nextSeasonStartBlock, StakingPoolBaseNoActiveSeason());
        }
    } else {
        // If season created, regular staking
        require(isSeasonActive(), StakingPoolBaseNoActiveSeason());
    }
    
    // ... staking logic
}
```

### Points Calculation

For pre-deposit stakes, `_calculateCurrentSeasonPoints()` uses:

```solidity
uint lastUpdate = position.lastUpdateBlock;
if (lastUpdate < current.startBlock && position.balance > 0) {
    // Pre-deposit case: Calculate from season start block
    return PointsLib.calculatePoints(
        position.balance, 
        current.startBlock,  // ✅ From season start block
        block.number, 
        blockTime, 
        pointsTimeUnit
    );
}
```

## Test Scenarios

### 1. Staking Attempt Before Pre-deposit
```
Current block: 1000
preDepositStartBlock: 1100
firstSeasonStartBlock: 1200

❌ Staking fails: StakingPoolBaseNoActiveSeason()
```

### 2. Staking During Pre-deposit Period
```
Current block: 1150
preDepositStartBlock: 1100
firstSeasonStartBlock: 1200

✅ Staking succeeds
- position.lastUpdateBlock = 1150
- Points are 0 (season not started yet)
```

### 3. Points Check After Season 1 Starts
```
Current block: 1250
Season 1 start block: 1200
position.lastUpdateBlock: 1150 (pre-deposit)

✅ Points calculation:
- fromBlock = 1200 (season start block)
- toBlock = 1250 (current block)
- points = balance * (1250 - 1200) * PRECISION / timeUnit
```

### 4. Staking After Season 1 Starts
```
Current block: 1250
Season 1 start block: 1200

✅ Staking succeeds
- position.lastUpdateBlock = 1250
- Points calculation: fromBlock = 1250 (staking block)
```

## Important Notes

### 1. Season 1 Exclusive
- `preDepositStartBlock` applies only to Season 1
- From Season 2 onwards, only regular staking is available

### 2. Block Ordering
```
preDepositStartBlock < firstSeasonStartBlock
```
- This order must be maintained to prevent logic errors
- Validation needed during project creation

### 3. Zero Value Handling
```solidity
if (preDepositStartBlock == 0) {
    // Pre-deposit disabled
    // Staking only allowed from firstSeasonStartBlock
}
```

### 4. Points Display
- After pre-deposit but before season start: Points = 0 (normal behavior)
- After season starts: Points calculated from season start block

## FAQ

### Q1: What's the difference between pre-deposit and regular staking?
**A:** Pre-deposit allows staking before season start, but points accumulate from the season start block. Regular staking accumulates points from the staking block.

### Q2: What happens if I unstake after pre-deposit?
**A:** Unstaking works the same way as regular unstaking. If done before season starts, points are 0 so you can withdraw without loss.

### Q3: How to disable pre-deposit?
**A:** Set `preDepositStartBlock` to 0. In this case, staking is only allowed from `firstSeasonStartBlock`.

### Q4: Is pre-deposit available for Season 2?
**A:** No, pre-deposit is a Season 1 exclusive feature. From Season 2 onwards, only regular staking is available.

## Related Files

### Smart Contracts
- `src/base/StakingPoolBase.sol` - Pre-deposit logic implementation
- `src/StakingPool.sol` - Points calculation logic
- `src/StakingViewer.sol` - Pre-deposit information queries

### Tests
- `test/BaseTest.sol` - Pre-deposit test setup
- `test/Season.t.sol` - Pre-deposit scenario tests

## Conclusion

The pre-deposit feature provides a better user experience by allowing users to prepare before the season starts. Points accumulation begins with Season 1 start, providing fair opportunities for early participants.
