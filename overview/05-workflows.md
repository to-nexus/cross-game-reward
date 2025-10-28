# Workflows

## 1. Project Creation Workflow

```
Creator
  ↓
Calls StakingProtocol.createProject()
  ↓
Factory retrieves bytecode from Code contracts
  ↓
Combines with constructor arguments
  ↓
Deploys StakingPool via CREATE2
  ↓
Deploys RewardPool via CREATE2
  ↓
Configures mutual references
  ↓
Grants roles
  ↓
Stores project info
  ↓
Returns (projectID, stakingPool, rewardPool)
```

**Key Points**:
- Deterministic addresses using CREATE2
- Salt: `keccak256(projectName, contractType)`
- StakingPool and RewardPool are linked
- Creator becomes project admin

---

## 2. Staking Workflow (Native CROSS)

```
User
  ↓
Sends native CROSS to StakingRouter.stake()
  ↓
Router wraps CROSS to WCROSS
  ↓
Router approves StakingPool for WCROSS
  ↓
Router calls StakingPool.stakeFor(user)
  ↓
StakingPool._ensureSeason()
  ├─ If season 0 and time reached → start first season
  ├─ If season ended → rollover season(s)
  └─ Ensure current season is active
  ↓
Update season aggregation
  ├─ Calculate incremental points
  ├─ aggregatedPoints += (totalStaked × elapsed blocks)
  └─ lastAggregatedBlock = current block
  ↓
Update user position
  ├─ Calculate user points since last update
  ├─ userPoints += (balance × elapsed blocks)
  ├─ balance += staking amount
  └─ lastUpdateBlock = current block
  ↓
Transfer WCROSS from Router to StakingPool
  ↓
Update totalStaked
  ↓
Emit Staked event
  ↓
Notify addon (if configured)
```

**Security Checks**:
- ✅ Season active
- ✅ Non-zero amount
- ✅ Token transfer success
- ✅ Reentrancy guard

---

## 3. Withdrawal Workflow

```
User
  ↓
Calls StakingRouter.unstake(projectID)
  ↓
Router calls StakingPool.withdrawAllFor(user)
  ↓
StakingPool._ensureSeason()
  ↓
Update season aggregation
  ├─ Calculate incremental points
  └─ Update aggregatedPoints
  ↓
Finalize user previous seasons (lazy snapshots)
  ├─ For each unfinalized previous season:
  │   ├─ Calculate user points for that season
  │   ├─ Store in UserSeasonData
  │   └─ Mark as finalized
  └─ Current season: calculate and store points
  ↓
Update season aggregation (after withdrawal)
  ↓
Calculate withdrawal amount (user's full balance)
  ↓
Reset user position
  ├─ points = 0
  ├─ balance = 0
  └─ lastUpdateBlock = current block
  ↓
Update totalStaked -= withdrawal amount
  ↓
Transfer WCROSS from StakingPool to Router
  ↓
Router unwraps WCROSS to native CROSS
  ↓
Router transfers native CROSS to user
  ↓
Emit Withdrawn event
  ↓
Notify addon (if configured)
```

**Key Points**:
- Finalizes all previous seasons (lazy snapshot)
- Forfeits current season points
- Full withdrawal only
- Automatic WCROSS unwrapping

---

## 4. Season Rollover Workflow

### Automatic Rollover

```
Any user transaction
  ↓
StakingPool._ensureSeason() called
  ↓
Check if current season ended
  ↓
If ended, start rollover loop (max 50 seasons)
  ↓
For each ended season:
  ├─ Finalize season aggregation
  │   ├─ Calculate final points
  │   ├─ aggregatedPoints += (totalStaked × remaining blocks)
  │   ├─ totalPoints = aggregatedPoints (cache)
  │   └─ isFinalized = true
  ├─ Notify addon of season end
  ├─ Increment currentSeason
  └─ Create next season
      ├─ seasonNumber = currentSeason
      ├─ startBlock = previous endBlock + 1
      ├─ endBlock = startBlock + seasonBlocks
      ├─ isFinalized = false
      ├─ totalPoints = 0
      ├─ aggregatedPoints = 0
      ├─ seasonTotalStaked = current totalStaked
      └─ lastAggregatedBlock = startBlock
```

### Manual Rollover

```
Admin
  ↓
Calls StakingPool.manualRolloverSeasons(count)
  ↓
Same process as automatic, but explicitly triggered
  ↓
Can rollover up to 'count' seasons
```

**Key Points**:
- Lazy evaluation (no user iteration)
- Batched rollovers (up to 50)
- Gas-efficient
- Deterministic

---

## 5. Reward Distribution Workflow

### Funding Season

```
Project Admin
  ↓
Approves RewardPool for ERC20 token
  ↓
Calls StakingProtocol.fundProjectSeason()
  ↓
Protocol verifies sender is project creator or admin
  ↓
Protocol calls RewardPool.fundSeason()
  ↓
RewardPool transfers tokens from sender
  ↓
seasonRewards[season][token] += amount
  ↓
If first funding for this token, add to seasonRewardTokens array
  ↓
Emit SeasonFunded event
```

### Claiming Rewards

```
User
  ↓
Calls StakingPool.claimSeason(season, token)
  ↓
Verify season is finalized
  ↓
Finalize user's season data (lazy snapshot)
  ├─ If not already finalized
  ├─ Calculate user points for season
  └─ Store in UserSeasonData
  ↓
Get user points for season
  ↓
Get total points for season
  ↓
Call RewardPool.payUser()
  ├─ Check not already claimed
  ├─ Verify season finalized
  ├─ Calculate reward = (total × userPoints) / totalPoints
  ├─ Check reward balance available
  ├─ Transfer reward to user
  ├─ seasonClaimed[season][token] += reward
  └─ hasClaimedSeasonReward[user][season][token] = true
  ↓
Emit RewardPaid event
```

**Multiple Season Claims**:
```
User
  ↓
Calls StakingPool.claimMultipleSeasons(seasons[], tokens[])
  ↓
For each season:
  └─ Execute claim workflow
```

**Security Checks**:
- ✅ Season finalized
- ✅ Not already claimed
- ✅ Sufficient reward balance
- ✅ Valid points ratio

---

## 6. Points Calculation Workflow

### Real-time Points Query

```
Frontend
  ↓
Calls StakingViewer.getUserPoints(projectID, user)
  ↓
Get user's stake position
  ↓
Get current season info
  ↓
Calculate current season points
  ├─ If user has balance:
  │   ├─ fromBlock = max(lastUpdateBlock, seasonStartBlock)
  │   ├─ toBlock = min(current block, seasonEndBlock)
  │   └─ points = (balance × (toBlock - fromBlock) × blockTime) / timeUnit
  └─ Add to user's stored points
  ↓
Return total points
```

### Season Total Points Query

```
Frontend
  ↓
Calls StakingViewer.getSeasonUserPoints(projectID, season, user)
  ↓
If season is finalized:
  ├─ Return season.totalPoints (cached)
Else (current season):
  ├─ Calculate from aggregation
  ├─ aggregatedPoints + (totalStaked × elapsed blocks since last aggregation)
  └─ Return calculated value
```

**Key Points**:
- No state changes (view functions)
- Virtual season support
- Gas-free queries
- Accurate real-time data

---

## 7. Pre-deposit Workflow

### Before Season 1

```
Timeline:
[Current] < [preDepositStartBlock] < [firstSeasonStartBlock]
          |                         |
        Waiting                Pre-deposit period
```

### Pre-deposit Staking

```
User
  ↓
Calls stake() during pre-deposit period
  ↓
currentSeason == 0 check passes
  ├─ preDepositStartBlock > 0
  ├─ current block >= preDepositStartBlock
  └─ current block < firstSeasonStartBlock
  ↓
Stake is allowed
  ↓
User position created
  ├─ balance = staking amount
  ├─ points = 0 (season not started)
  └─ lastUpdateBlock = current block
```

### After Season 1 Starts

```
Season 1 Start Block Reached
  ↓
Next transaction triggers _ensureSeason()
  ↓
First season created
  ├─ currentSeason = 1
  ├─ startBlock = firstSeasonStartBlock
  └─ Pre-deposit users automatically included
  ↓
Points calculation for pre-deposit users
  ├─ fromBlock = firstSeasonStartBlock (not lastUpdateBlock!)
  ├─ toBlock = current block
  └─ points = (balance × (toBlock - fromBlock) × blockTime) / timeUnit
```

**Key Points**:
- Season 1 exclusive
- Points start from season start block
- Fair for all pre-deposit users
- Optional feature

---

## 8. Emergency Procedures

### Pause Staking

```
Admin
  ↓
Calls StakingPool.pause()
  ↓
Pausable state = true
  ↓
stake() and withdrawAll() revert
  ↓
View functions still work
  ↓
Reward claims still work
```

### Unpause

```
Admin
  ↓
Calls StakingPool.unpause()
  ↓
Pausable state = false
  ↓
Normal operations resume
```

### Recover Unclaimed Rewards

```
Project Admin
  ↓
Calls RewardPool.recoverRemaining(season, token)
  ↓
Verify season is finalized
  ↓
Calculate remaining = total - claimed
  ↓
Transfer remaining to protocol (creator can withdraw)
  ↓
Emit RemainingRecovered event
```

**Security Considerations**:
- Only admin can pause
- Cannot pause reward claims
- Cannot access user staked funds
- Recovery only for unclaimed rewards

---

## 9. Addon Integration Workflow

### StakingAddon Hook

```
StakingPool operation (stake/withdraw/season end)
  ↓
After core logic completes
  ↓
If stakingAddon configured:
  ├─ Try call addon notification function
  ├─ Use try-catch for safety
  └─ Continue even if addon fails
```

**Example: RankingAddon**
```
User stakes
  ↓
StakingPool completes stake
  ↓
Calls RankingAddon.afterStake(user, amount)
  ↓
RankingAddon updates Top 10 rankings
  ├─ Calculate user's total stake
  ├─ Update TopRanker array if needed
  └─ Emit ranking events
```

**Key Points**:
- Non-blocking (safe call)
- Extension without modification
- Optional feature
- Gas-efficient updates

---

## Common Error Scenarios

### "No Active Season"
- **Cause**: Staking before firstSeasonStartBlock
- **Solution**: Wait for season start or enable pre-deposit

### "Season Not Ended"
- **Cause**: Trying to claim before season end
- **Solution**: Wait for season to finalize

### "Already Claimed"
- **Cause**: Attempting to claim rewards twice
- **Solution**: Check claim status before claiming

### "Below Minimum Stake"
- **Cause**: Staking less than MIN_STAKE
- **Solution**: Stake at least 1 CROSS

### "No Position"
- **Cause**: Withdrawing without staked balance
- **Solution**: Check balance before withdrawing

