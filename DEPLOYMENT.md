# Deployment Guide

Cross-Staking Protocol deployment architecture and configuration guide.

## Architecture Overview

### Deployment Components

The Cross-Staking Protocol consists of the following contracts that need to be deployed:

```
1. WCROSS (Wrapped CROSS Token)
   └─ ERC20 wrapper for native CROSS

2. Code Contracts (for CREATE2 deployment)
   ├─ StakingPoolCode
   └─ RewardPoolCode

3. StakingProtocol (Factory)
   └─ Creates project-specific pools

4. StakingRouter (Optional)
   └─ User-friendly native token interface

5. StakingViewer (Optional)
   └─ Unified view functions
```

### Deployment Flow

```
Step 1: Deploy WCROSS
   ↓
Step 2: Deploy Code Contracts
   ↓
Step 3: Deploy StakingProtocol(WCROSS, Codes, Admin)
   ↓
Step 4: Deploy StakingRouter(WCROSS, Protocol)
   ↓
Step 5: Deploy StakingViewer(Protocol)
```

## Contract Initialization

### 1. WCROSS

**Parameters**: None

**Purpose**: Wraps native CROSS token into ERC20 standard.

```solidity
contract WCROSS {
    // No constructor parameters
}
```

### 2. Code Contracts

**StakingPoolCode**
```solidity
contract StakingPoolCode {
    function code() external pure returns (bytes memory) {
        return type(StakingPool).creationCode;
    }
}
```

**RewardPoolCode**
```solidity
contract RewardPoolCode {
    function code() external pure returns (bytes memory) {
        return type(RewardPool).creationCode;
    }
}
```

**Purpose**: Store creation bytecode for CREATE2 deployments, reducing factory contract size.

### 3. StakingProtocol

**Constructor Parameters**:
```solidity
constructor(
    address _wcross,              // WCROSS token address
    address _stakingPoolCode,     // StakingPoolCode address
    address _rewardPoolCode,      // RewardPoolCode address
    address _initialAdmin         // Initial protocol admin
)
```

**Purpose**: Factory contract for creating project-specific staking pools.

**Key Features**:
- CREATE2 deployment for deterministic addresses
- Global settings management
- Access control with 3-day timelock

### 4. StakingRouter

**Constructor Parameters**:
```solidity
constructor(
    address _wcross,              // WCROSS token address
    address _stakingProtocol      // StakingProtocol address
)
```

**Purpose**: User-friendly interface for native CROSS staking operations.

**Key Features**:
- Automatic WCROSS wrapping/unwrapping
- Batch operations support
- Reentrancy protection

### 5. StakingViewer

**Constructor Parameters**:
```solidity
constructor(
    address _protocol             // StakingProtocol address
)
```

**Purpose**: Unified view functions for all staking data queries.

**Key Features**:
- Virtual season calculations
- Batch queries
- Gas-free read operations

## Post-Deployment Configuration

### Project Creation

```solidity
function createProject(
    string calldata projectName,      // Unique project name
    uint seasonDuration,                 // Blocks per season (0 = use default)
    uint firstSeasonStartBlock,       // Season 1 start block
    uint poolEndTime,                 // Pool end block (0 = infinite)
    address projectAdmin,              // Project administrator
    uint preDepositStartBlock          // Pre-deposit start block (0 = disabled)
) external returns (uint projectID, address stakingPool, address rewardPool)
```

**Example**:
```solidity
// Create project with 30-day seasons
protocol.createProject(
    "My DeFi Project",
    2592000,           // 30 days (assuming 1 sec/block)
    block.number + 100,  // Start in 100 blocks
    0,                 // Infinite duration
    msg.sender,        // Admin
    0                  // No pre-deposit
);
```

### Reward Token Setup

```solidity
// 1. Approve reward token
IERC20(rewardToken).approve(rewardPoolAddress, amount);

// 2. Fund season
protocol.fundProjectSeason(
    projectId,
    seasonNumber,
    rewardTokenAddress,
    amount
);
```

## Configuration Parameters

### Global Settings

**Default Season Blocks**
```solidity
// Default season length if project doesn't specify
protocol.setDefaultSeasonBlocks(2592000); // 30 days
```

### Project-Specific Settings

**Points Time Unit**
```solidity
// Time unit for points calculation (default: 1 hour)
protocol.setPoolPointsTimeUnit(projectId, 1 hours);
```

**Block Time**
```solidity
// Block time in seconds (default: 1 second)
protocol.setPoolBlockTime(projectId, 1);
```

**Pool End Block**
```solidity
// Set or update pool end block
protocol.setPoolEndBlock(projectId, endTime);
```

**Season Start**
```solidity
// Set next season start block (for restart after end)
protocol.setNextSeasonStart(projectId, startTime);
```

## Access Control

### Roles

**DEFAULT_ADMIN_ROLE**
- Protocol-level administrator
- 3-day timelock for role transfers
- Can modify global settings

**Project Creator**
- Can fund their project's seasons
- Can view project settings

**Project Admin**
- Can modify project-specific settings
- Set via project creation or transfer

### Permission Examples

```solidity
// Grant admin role (requires existing admin)
protocol.grantRole(DEFAULT_ADMIN_ROLE, newAdmin);

// Set project admin (requires protocol admin)
protocol.setProjectAdmin(projectId, newAdmin);
```

## Season Management

### Automatic Season Start

Seasons automatically start when `firstSeasonStartBlock` is reached. No manual intervention needed.

### Season Rollover

```solidity
// Anyone can call after season end block
stakingPool.rolloverSeason();
```

**Features**:
- Automatic rollover support (up to 50 seasons)
- Lazy snapshot for gas efficiency
- Deterministic season transitions

### Pool Termination

```solidity
// Set pool end block
protocol.setPoolEndBlock(projectId, endTime);
```

### Pool Restart

```solidity
// Set new season start after termination
protocol.setNextSeasonStart(projectId, startTime);
```

## Security Considerations

### Multi-signature Wallet

Recommended for admin roles:
- Use multi-sig wallet for `DEFAULT_ADMIN_ROLE`
- Protects against single point of failure
- Requires multiple signatures for critical operations

### Timelock

Built-in protection:
- 3-day timelock for admin role transfers
- Prevents immediate malicious takeover
- Allows community to detect suspicious activity

### Pausable

Emergency stop mechanism:
```solidity
// Pause staking operations (admin only)
stakingPool.pause();

// Resume operations
stakingPool.unpause();
```

### Access Control Verification

```solidity
// Check if address has admin role
bool isAdmin = protocol.hasRole(DEFAULT_ADMIN_ROLE, address);

// Check project admin
ProjectInfo memory info = protocol.getProject(projectId);
address admin = info.admin;
```

## CREATE2 Address Prediction

### Deterministic Addresses

StakingPool and RewardPool addresses can be predicted before deployment:

**Salt Structure**:
```solidity
// StakingPool salt
bytes32 salt = keccak256(abi.encodePacked(projectName, "StakingPool"));

// RewardPool salt
bytes32 salt = keccak256(abi.encodePacked(projectName, "RewardPool"));
```

**Benefits**:
- Frontend can calculate addresses before deployment
- Cross-chain deployment with same addresses
- Easier contract verification

## Monitoring

### Key Metrics

**Protocol Level**:
- Total projects created
- Active projects count
- Global settings

**Project Level**:
- Total staked amount
- Current season
- Season progress
- Total points

**Reward Level**:
- Season reward tokens
- Total rewards deposited
- Claimed vs unclaimed rewards

### View Functions

```solidity
// Project info
ProjectInfo memory info = protocol.getProject(projectId);

// Pool stats
uint totalStaked = viewer.getTotalStaked(projectId);
uint currentSeason = viewer.getCurrentSeason(projectId);

// User position
(uint balance, uint points, uint lastUpdate) = 
    viewer.getStakeInfo(projectId, user);

// Season info
(uint season, uint start, uint end, uint elapsed) = 
    viewer.getSeasonInfo(projectId);
```

## Verification

### Etherscan Verification

After deployment, verify all contracts on block explorer for transparency:

- WCROSS
- StakingPoolCode
- RewardPoolCode
- StakingProtocol
- StakingRouter
- StakingViewer

### Verification Checklist

- [ ] All contracts verified on block explorer
- [ ] Admin roles configured correctly
- [ ] Test project created successfully
- [ ] Staking/unstaking works
- [ ] Season rollover tested
- [ ] Reward deposit and claim tested

## Reference

For detailed implementation examples, see:
- Test files in `test/` directory
- Contract interfaces in `src/interfaces/`
- Base contracts in `src/base/`
