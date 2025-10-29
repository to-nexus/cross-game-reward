# Test Documentation

Cross-Staking Protocol test cases and coverage documentation.

## Test Overview

- **Total Tests**: 68
- **Pass Rate**: 100%
- **Framework**: Foundry Test
- **Core Logic Coverage**: 100%

## Test Structure

```
test/
├── BaseTest.sol          # Common test helpers and utilities
├── Staking.t.sol         # Basic staking functionality tests (9)
├── Season.t.sol          # Season management tests (7)
├── Points.t.sol          # Points calculation tests (9)
├── Rewards.t.sol         # Reward distribution tests (8)
├── MultiPool.t.sol       # Multi-project tests (6)
├── Advanced.t.sol        # Advanced scenario tests (8)
├── Integrated.t.sol      # Integration tests (8)
└── Fuzz.t.sol            # Fuzz tests (13)
```

## Running Tests

### Basic Tests
```bash
forge test
```

### With Gas Report
```bash
forge test --gas-report
```

### Specific Test
```bash
forge test --match-contract StakingTest
forge test --match-test test_BasicStake
```

### Verbose Output
```bash
forge test -vvv  # Very verbose logs
```

### Fuzz Tests
```bash
forge test --match-contract FuzzTest
```

## Test Cases Detail

### 1. Staking.t.sol - Basic Staking Functionality (9 tests)

#### test_BasicStake
- **Purpose**: Verify basic staking operation
- **Validates**:
  - Token transfer works correctly
  - Staking amount is recorded
  - Event emission

#### test_MinimumStake
- **Purpose**: Verify minimum staking amount
- **Validates**:
  - Only amounts >= MIN_STAKE (1 CROSS) allowed
  - Exactly MIN_STAKE succeeds

#### test_MultipleStakes
- **Purpose**: Verify split staking functionality
- **Validates**:
  - Multiple stakes allowed
  - Cumulative sum works correctly
  - Points accumulation accuracy

#### test_MultipleUsersStake
- **Purpose**: Multiple user staking
- **Validates**:
  - Independent position management per user
  - totalStaked accuracy

#### test_WithdrawAll
- **Purpose**: Full withdrawal functionality
- **Validates**:
  - Token return
  - Points forfeit
  - Position reset

#### test_StakeAfterWithdraw
- **Purpose**: Re-staking after withdrawal
- **Validates**:
  - Staking possible after withdrawal
  - Previous season points retained

#### test_RouterUnstake
- **Purpose**: Native CROSS withdrawal via Router
- **Validates**:
  - WCROSS → Native CROSS conversion
  - Correct amount returned

#### test_RevertWhen_BelowMinimumStake
- **Purpose**: Reject below minimum stake
- **Validates**: StakingPoolBelowMinStake error

#### test_RevertWhen_WithdrawWithoutStake
- **Purpose**: Reject withdrawal without balance
- **Validates**: StakingPoolNoPosition error

---

### 2. Season.t.sol - Season Management (7 tests)

#### test_SeasonRollover
- **Purpose**: Basic season rollover functionality
- **Validates**:
  - currentSeason increment
  - New season info creation

#### test_PointsResetAfterRollover
- **Purpose**: Points management after rollover
- **Validates**:
  - Previous season points snapshot
  - New season points start independently

#### test_TotalPointsSnapshot
- **Purpose**: Total points snapshot per season
- **Validates**:
  - totalPoints calculation during rollover
  - Snapshot value accuracy

#### test_MultipleSeasons
- **Purpose**: Multiple season progression
- **Validates**:
  - Consecutive season rollovers
  - Data independence per season

#### test_StakePreservedAcrossSeasons
- **Purpose**: Stake preservation across seasons
- **Validates**:
  - Stake maintained after rollover
  - Automatic participation in next season

#### test_SeasonInfo
- **Purpose**: Season information query
- **Validates**:
  - startTime, endTime accuracy
  - blocksElapsed calculation

#### test_RevertWhen_RolloverBeforeSeasonEnd
- **Purpose**: Reject rollover before season end
- **Validates**: StakingPoolSeasonNotEnded error

---

### 3. Points.t.sol - Points Calculation (9 tests)

#### test_PointsCalculation
- **Purpose**: Basic points calculation logic
- **Validates**:
  - points = (balance × time × PRECISION) / timeUnit
  - Proportional increase over time

#### test_PointsAccumulation
- **Purpose**: Points accumulation
- **Validates**:
  - Accumulation over time
  - Real-time query accuracy

#### test_PointsProportionalToStake
- **Purpose**: Points proportional to stake
- **Validates**:
  - 2x stake = 2x points
  - Ratio accuracy

#### test_PointsTimeUnitAffectsCalculation
- **Purpose**: Time unit change effects
- **Validates**:
  - Points calculation changes with timeUnit
  - Setting value reflection accuracy

#### test_BlockTimeAffectsPoints
- **Purpose**: Block time change effects
- **Validates**:
  - Points calculation changes with blockTime
  - Different chain adaptation

#### test_SnapshotPlusAdditionalPoints
- **Purpose**: Snapshot + real-time points
- **Validates**:
  - Snapshot behavior during split staking
  - Cumulative points accuracy

#### test_ExpectedSeasonPointsWithSnapshot
- **Purpose**: Season expected points query
- **Validates**:
  - Past season snapshots
  - Current season real-time calculation

#### test_PointsResetOnWithdraw
- **Purpose**: Points forfeiture on withdrawal
- **Validates**:
  - Current season points reset to 0
  - Previous seasons retained

#### test_NoPointsWithoutStake
- **Purpose**: No points without stake
- **Validates**: 0 stake = 0 points

---

### 4. Rewards.t.sol - Reward Distribution (8 tests)

#### test_FundSeason
- **Purpose**: Season reward deposit
- **Validates**:
  - Token transfer
  - Reward pool balance increase

#### test_EqualStakeEqualReward
- **Purpose**: Equal stake → equal reward
- **Validates**:
  - 1:1 points ratio
  - 1:1 reward ratio

#### test_ProportionalRewardDistribution
- **Purpose**: Proportional reward distribution
- **Validates**:
  - Rewards based on points ratio
  - Accurate ratio calculation

#### test_MultipleSeasonRewards
- **Purpose**: Multi-season rewards
- **Validates**:
  - Independent rewards per season
  - Each season claimable

#### test_NoRewardForNoStake
- **Purpose**: No reward without stake
- **Validates**:
  - 0 points = 0 reward
  - Claim attempt does nothing

#### test_RemainingRewards
- **Purpose**: Unclaimed reward recovery
- **Validates**:
  - recoverRemaining operation
  - Return to creator

#### test_RevertWhen_ClaimBeforeSeasonEnd
- **Purpose**: Reject claim before season end
- **Validates**: StakingPoolSeasonNotEnded error

#### test_RevertWhen_DuplicateClaim
- **Purpose**: Reject duplicate claim
- **Validates**: StakingPoolAlreadyClaimed error

---

### 5. MultiPool.t.sol - Multi-Project (6 tests)

#### test_MultiplePoolsCreated
- **Purpose**: Multiple project creation
- **Validates**:
  - Independent pool per project
  - projectCount increment

#### test_StakeInDifferentPools
- **Purpose**: Simultaneous staking in multiple pools
- **Validates**:
  - User participation in multiple projects
  - Independent positions

#### test_IndependentSeasons
- **Purpose**: Independent seasons per project
- **Validates**:
  - Season schedule independence
  - Rollover independence

#### test_IndependentRewards
- **Purpose**: Independent rewards per project
- **Validates**:
  - RewardPool independence
  - Reward distribution independence

#### test_SameUserMultiplePools
- **Purpose**: Same user in multiple projects
- **Validates**:
  - Points accumulation in each project
  - Reward claims in each project

#### test_DifferentSeasonLengths
- **Purpose**: Different season lengths per project
- **Validates**:
  - seasonDuration setting independence
  - Each operates correctly

---

### 6. Advanced.t.sol - Advanced Scenarios (8 tests)

#### test_VirtualToRealSeasonTransition
- **Purpose**: Virtual to actual season transition
- **Validates**:
  - Lazy first season creation
  - Automatic participation mechanism

#### test_MultiSeasonAutoParticipation
- **Purpose**: Automatic season participation
- **Validates**:
  - Auto-participation with maintained stake
  - Points accumulation without mid-action

#### test_ZeroBalanceUserInRollover
- **Purpose**: Rollover handling for zero-balance users
- **Validates**:
  - Error-free processing
  - No unnecessary gas cost

#### test_SeasonGapBetweenRollovers
- **Purpose**: Time gap between rollovers
- **Validates**:
  - Next season start block setting
  - Delayed start support

#### test_PoolEndBlockAtSeasonBoundary
- **Purpose**: Pool end block at season boundary
- **Validates**:
  - Accurate end handling
  - Last season points calculation

#### test_PoolRestartAfterEnd
- **Purpose**: Pool restart after end
- **Validates**:
  - Restart with setNextSeasonStart
  - New season start

#### test_ClaimAfterSeasonEnd
- **Purpose**: Claim after season end
- **Validates**:
  - Lazy snapshot operation
  - Accurate points calculation

#### test_AdminEmergencyFunctions
- **Purpose**: Admin emergency functions
- **Validates**:
  - Setting change permissions
  - AccessControl operation

---

### 7. Integrated.t.sol - Integration Tests (8 tests)

#### test_ProtocolCompleteFlow
- **Purpose**: Complete protocol flow
- **Validates**:
  - Project creation → staking → season → reward
  - All steps work correctly

#### test_CompleteStakingFlow
- **Purpose**: Complete staking cycle
- **Validates**:
  - stake → points → claim → withdraw
  - Full lifecycle

#### test_MultiUserMultiSeasonFlow
- **Purpose**: Multiple users, multiple seasons
- **Validates**:
  - Complex scenario handling
  - Accurate points and reward distribution

#### test_MultiSeasonFlow
- **Purpose**: Multiple season progression
- **Validates**:
  - Season transitions
  - Independence per season

#### test_StakeWithdrawRestake
- **Purpose**: Stake → withdraw → re-stake
- **Validates**:
  - Re-participation after withdrawal
  - Previous data cleanup

#### test_FullNativeTokenFlow
- **Purpose**: Native CROSS complete flow
- **Validates**:
  - Native → WCROSS via Router
  - WCROSS → Native on withdrawal

#### test_ProtocolCreateProject
- **Purpose**: Project creation via Protocol
- **Validates**:
  - Factory pattern operation
  - Pool connection normal

#### test_EdgeCase_ClaimWithoutPoints
- **Purpose**: Claim attempt without points
- **Validates**:
  - Error-free processing
  - Nothing happens

---

### 8. Fuzz.t.sol - Fuzz Tests (13 tests)

Fuzz tests run 267 times each with random inputs.

#### testFuzz_PointsCalculation
- **Purpose**: Points calculation with various inputs
- **Validates**: No overflow/underflow

#### testFuzz_RewardDistribution
- **Purpose**: Reward distribution with random stakes
- **Validates**: Ratio accuracy, no rounding errors

#### testFuzz_MultipleSeasonRewards
- **Purpose**: Random season progression
- **Validates**: No interference between seasons

#### testFuzz_IncrementalStaking
- **Purpose**: Random split staking
- **Validates**: Cumulative accuracy

#### testFuzz_SeasonRollover
- **Purpose**: Random timing rollover
- **Validates**: Normal operation regardless of rollover timing

#### testFuzz_LargeStake
- **Purpose**: Large amount staking
- **Validates**: Overflow prevention

#### testFuzz_LongDuration
- **Purpose**: Long period staking
- **Validates**: Time calculation accuracy

#### testFuzz_MinimumStake
- **Purpose**: Values near MIN_STAKE
- **Validates**: Boundary value handling

#### testFuzz_MultiUserPointsRatio
- **Purpose**: Multiple user points ratio
- **Validates**: Ratio calculation accuracy

#### testFuzz_PointsAcrossSeasons
- **Purpose**: Points tracking across seasons
- **Validates**: Data integrity between seasons

#### testFuzz_TimeParameters
- **Purpose**: Random time parameters
- **Validates**: Normal handling of timeUnit, blockTime changes

#### testFuzz_ManyIncrementalStakes
- **Purpose**: Many split stakes
- **Validates**: Gas cost, accuracy

#### testFuzz_ZeroAmount
- **Purpose**: Zero amount handling
- **Validates**: Error or normal rejection

## Coverage

### Core Functionality Coverage: 100%
- ✅ Staking/Withdrawal
- ✅ Points Calculation
- ✅ Season Management
- ✅ Reward Distribution
- ✅ Permission Management
- ✅ Router Functionality

### Edge Cases Coverage
- ✅ Minimum/Maximum values
- ✅ Season boundaries
- ✅ Duplicate actions
- ✅ Unauthorized calls
- ✅ Zero balance users
- ✅ Emergency situations

## Test Execution Time

- **All Tests**: ~520ms (CPU time)
- **Fuzz Tests**: ~580ms (CPU time)
- **Total**: ~1 second

## Reference

- [Foundry Book](https://book.getfoundry.sh/)
- [Forge Testing Guide](https://book.getfoundry.sh/forge/tests)
- [Fuzz Testing](https://book.getfoundry.sh/forge/fuzz-testing)
