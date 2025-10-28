# Contract Details

## Contract Overview

The Cross Staking Protocol consists of 11 main contracts organized in a hierarchical structure.

## Core Contracts

### 1. WCROSS

**Purpose**: ERC20 wrapper for native CROSS token

**Key Features**:
- 1:1 wrapping ratio
- Standard WETH-style implementation
- Automatic wrapping via `receive()` function

**Main Functions**:
```solidity
function wrap() external payable
function unwrap(uint amount) external
function deposit() external payable  // Alias for wrap
function withdraw(uint amount) external  // Alias for unwrap
```

**Security**:
- Simple, battle-tested design
- No admin functions
- No upgradability (immutable)

---

### 2. StakingProtocol

**Purpose**: Factory and central management contract

**Key Features**:
- Creates project-specific staking pools
- Uses CREATE2 for deterministic addresses
- Manages global settings
- Access control with 3-day timelock

**Main Functions**:
```solidity
// Project Management
function createProject(
    string calldata projectName,
    uint seasonBlocks,
    uint firstSeasonStartBlock,
    uint poolEndBlock,
    address projectAdmin,
    uint preDepositStartBlock
) external returns (uint projectID, address stakingPool, address rewardPool)

// Funding
function fundProjectSeason(
    uint projectID,
    uint seasonNumber,
    address rewardToken,
    uint amount
) external

// Configuration
function setPoolPointsTimeUnit(uint projectID, uint _pointsTimeUnit) external
function setPoolBlockTime(uint projectID, uint _blockTime) external
function setPoolEndBlock(uint projectID, uint endBlock) external
```

**State Variables**:
```solidity
mapping(uint => ProjectInfo) public projects;
mapping(string => uint) public projectIDByName;
mapping(address => uint[]) public projectsByAddress;
uint public projectCount;
uint public defaultSeasonBlocks;
```

**CREATE2 Salt Structure**:
```solidity
// StakingPool salt
bytes32 salt = keccak256(abi.encodePacked(projectName, "StakingPool"));

// RewardPool salt
bytes32 salt = keccak256(abi.encodePacked(projectName, "RewardPool"));
```

---

### 3. StakingPool

**Purpose**: Project-specific staking pool

**Key Features**:
- Manages user stakes
- Calculates points with O(1) aggregation
- Automatic season rollover
- Lazy snapshot system
- Pausable for emergencies

**Main Functions**:
```solidity
// Staking Operations
function stake(uint amount) external
function stakeFor(address user, uint amount) external
function withdrawAll() external
function withdrawAllFor(address user) external

// Reward Claims
function claimSeason(uint season, address rewardToken) external
function claimMultipleSeasons(
    uint[] calldata seasons,
    address[] calldata rewardTokens
) external

// Season Management
function rolloverSeason() external
function manualRolloverSeasons(uint count) external

// View Functions
function getCurrentSeasonInfo() external view returns (
    uint currentSeason,
    uint seasonStartBlock,
    uint seasonEndBlock,
    uint blocksElapsed
)
function getUserPoints(address user) external view returns (uint)
function getSeasonUserPoints(uint season, address user) external view returns (uint)
```

**State Variables**:
```solidity
mapping(address => StakePosition) public userStakes;
mapping(uint => Season) public seasons;
mapping(address => mapping(uint => UserSeasonData)) public userSeasonData;
uint public currentSeason;
uint public totalStaked;
```

---

### 4. RewardPool

**Purpose**: Season-based reward distribution

**Key Features**:
- Multi-token support per season
- Proportional distribution
- Double-claim prevention
- Emergency token recovery

**Main Functions**:
```solidity
// Funding
function fundSeason(uint season, address token, uint amount) external
function depositReward(address token, uint amount) external

// Distribution
function payUser(
    address user,
    uint season,
    address rewardToken,
    uint userPoints,
    uint totalPoints
) external returns (uint rewardAmount)

// Recovery
function recoverRemaining(uint season, address token) external

// View Functions
function getSeasonRewardTokens(uint season) external view returns (address[] memory)
function getSeasonTokenInfo(uint season, address token) external view returns (
    uint total,
    uint claimed,
    uint remaining
)
```

**State Variables**:
```solidity
mapping(uint => mapping(address => uint)) public seasonRewards;
mapping(uint => mapping(address => uint)) public seasonClaimed;
mapping(address => mapping(uint => mapping(address => bool))) public hasClaimedSeasonReward;
mapping(uint => address[]) public seasonRewardTokens;
```

---

### 5. StakingRouter

**Purpose**: User-friendly native token interface

**Key Features**:
- Automatic WCROSS wrapping/unwrapping
- Batch operations
- Proxy staking support

**Main Functions**:
```solidity
// Native Token Operations
function stake(uint projectID) external payable
function unstake(uint projectID) external

// Batch Claims
function claimReward(uint projectID, uint season, address rewardToken) external
function claimMultipleRewards(
    uint projectID,
    uint[] calldata seasons,
    address[] calldata rewardTokens
) external

// Batch Finalization
function batchFinalizeSeasons(uint projectID, uint count) external
```

**Usage Flow**:
```
User sends native CROSS
    ↓
Router wraps to WCROSS
    ↓
Router approves StakingPool
    ↓
Router calls stakeFor()
    ↓
User's stake recorded
```

---

### 6. StakingViewer

**Purpose**: Unified view functions

**Key Features**:
- Virtual season calculations
- Batch queries
- Gas-free reads
- No state changes

**Main Functions**:
```solidity
// Pool Information
function getPoolInfo(uint projectID) external view returns (
    uint blockTime,
    uint pointsTimeUnit,
    uint seasonBlocks,
    uint poolEndBlock,
    uint currentSeason,
    uint preDepositStartBlock,
    uint firstSeasonStartBlock
)

// User Information
function getStakeInfo(uint projectID, address user) external view returns (
    uint balance,
    uint points,
    uint lastUpdateBlock
)

// Season Information
function getSeasonInfo(uint projectID) external view returns (
    uint currentSeason,
    uint seasonStartBlock,
    uint seasonEndBlock,
    uint blocksElapsed
)

// Points Information
function getUserPoints(uint projectID, address user) external view returns (uint)
function getSeasonUserPoints(uint projectID, uint season, address user) external view returns (
    uint userPoints,
    uint totalPoints
)

// Reward Preview
function previewClaim(
    uint projectID,
    address user,
    uint season,
    address rewardToken
) external view returns (
    uint userPoints,
    uint totalPoints,
    uint expectedReward,
    bool alreadyClaimed
)
```

**Virtual Season Support**:
- Calculates what currentSeason would be without on-chain rollover
- Allows frontends to show accurate data before expensive rollover tx
- Example: If 54 seasons passed but no rollover, viewer returns season 54

---

## Base Contracts

### CrossStakingBase

**Purpose**: Common base for all contracts

**Inherited Features**:
- `AccessControlDefaultAdminRules`: Role-based access with 3-day timelock
- `ReentrancyGuardTransient`: EIP-1153 transient storage reentrancy guard

**Provided Functions**:
```solidity
// Validation
function _requireNonZero(uint value) internal pure
function _requireNonZeroAddress(address addr) internal pure

// Safe Transfers
function _safeTransferFrom(IERC20 token, address from, address to, uint amount) internal
function _safeTransfer(IERC20 token, address to, uint amount) internal
```

---

### StakingPoolBase

**Purpose**: Core staking pool logic

**Key Features**:
- Season management
- Points calculation
- O(1) aggregation system
- Lazy snapshot system
- Hook pattern for extensibility

**Core Functions**:
```solidity
// Internal Core
function _stakeFor(address user, uint amount, address from) internal virtual
function _withdrawAll(address user, address to) internal virtual
function _ensureSeason() internal
function _rolloverSeason() internal

// Aggregation
function _updateSeasonAggregation(uint seasonNum) internal
function _finalizeSeasonAggregation(uint seasonNum) internal

// Snapshots
function _ensureUserSeasonSnapshot(address user, uint seasonNum) internal
function _ensureUserAllPreviousSeasons(address user) internal

// Hooks
function _beforeStake(address user, uint amount) internal virtual
function _afterStake(address user, uint amount) internal virtual
function _beforeWithdraw(address user, uint amount) internal virtual
function _afterWithdraw(address user, uint amount) internal virtual
```

---

### RewardPoolBase

**Purpose**: Core reward distribution logic

**Key Features**:
- Multi-token reward management
- Proportional distribution
- Double-claim prevention
- Hook pattern for bonuses

**Core Functions**:
```solidity
// Internal Core
function _fundSeason(uint season, address token, uint amount) internal virtual
function _payUser(
    address user,
    uint season,
    address rewardToken,
    uint userPoints,
    uint totalPoints
) internal virtual returns (uint rewardAmount)

// Hooks
function _beforeFundSeason(uint season, address token, uint amount) internal virtual
function _afterFundSeason(uint season, address token, uint amount) internal virtual
function _calculateBonusReward(...) internal virtual returns (uint)
function _beforePayUser(...) internal virtual
function _afterPayUser(...) internal virtual
```

---

## Libraries

### PointsLib

**Purpose**: Pure points calculation logic

**Functions**:
```solidity
function calculatePoints(
    uint balance,
    uint fromBlock,
    uint toBlock,
    uint blockTime,
    uint timeUnit
) internal pure returns (uint points)

function calculateProRata(
    uint totalAmount,
    uint userShare,
    uint totalShare
) internal pure returns (uint userAmount)
```

---

### SeasonLib

**Purpose**: Season validation and helper functions

**Functions**:
```solidity
function isSeasonActive(
    uint startBlock,
    uint endBlock,
    bool isFinalized
) internal view returns (bool)

function isSeasonEnded(uint endBlock) internal view returns (bool)

function validateSeasonBlocks(uint startBlock, uint endBlock) internal pure

function isBlockInSeason(
    uint blockNumber,
    uint startBlock,
    uint endBlock
) internal pure returns (bool)
```

---

## Code Contracts

### StakingPoolCode

**Purpose**: Stores StakingPool creation bytecode

```solidity
contract StakingPoolCode {
    function code() external pure returns (bytes memory) {
        return type(StakingPool).creationCode;
    }
}
```

### RewardPoolCode

**Purpose**: Stores RewardPool creation bytecode

```solidity
contract RewardPoolCode {
    function code() external pure returns (bytes memory) {
        return type(RewardPool).creationCode;
    }
}
```

**Why Code Contracts?**
1. Reduces factory contract size (avoids 24KB limit)
2. Enables bytecode upgrades without redeploying factory
3. Gas-efficient storage of large bytecode
4. Clear separation of concerns

---

## Access Control Structure

```
StakingProtocol
├─ DEFAULT_ADMIN_ROLE (3-day timelock)
│  ├─ Can create projects
│  ├─ Can modify global settings
│  └─ Can grant/revoke roles
└─ MANAGER_ROLE
   └─ Can perform specific management tasks

StakingPool
├─ DEFAULT_ADMIN_ROLE (project admin)
│  ├─ Can pause/unpause
│  ├─ Can modify pool settings
│  └─ Cannot access user funds
└─ REWARD_POOL_ROLE (RewardPool contract)
   └─ Can call updatePoints()

RewardPool
├─ DEFAULT_ADMIN_ROLE (project admin)
│  └─ Can recover unclaimed rewards
└─ STAKING_POOL_ROLE (StakingPool contract)
   └─ Can call payUser()
```

---

## Security Features

### Reentrancy Protection
- `ReentrancyGuardTransient` on all state-changing functions
- EIP-1153 transient storage for 30% gas savings
- Automatic protection against reentrancy attacks

### Access Control
- Role-based permissions
- 3-day timelock for admin transfers
- Prevents immediate malicious takeover

### Pausable
- Emergency stop mechanism
- Admin can pause staking/unstaking
- Does not affect view functions or reward claims

### SafeERC20
- Handles non-standard ERC20 tokens
- Prevents silent failures
- Used for all token transfers

### Input Validation
- Zero address checks
- Zero amount checks
- Season validation
- Block range validation

### Immutable Variables
- Contract addresses stored as immutable
- Cannot be changed after deployment
- Prevents address manipulation

---

## Gas Optimization Techniques

1. **Custom Errors**: 15-20% savings vs string errors
2. **Transient Storage**: 30% savings in reentrancy guard  
3. **Unchecked Arithmetic**: 5-10% savings where safe
4. **O(1) Aggregation**: Constant-time operations
5. **Lazy Snapshots**: Deferred processing
6. **Immutable Variables**: No storage reads
7. **Named Imports**: Reduces compilation size
8. **Function Inlining**: Optimizer optimization

