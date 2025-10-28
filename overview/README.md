# Cross Staking Protocol - Documentation

## Overview

This directory contains comprehensive documentation for the Cross Staking Protocol, organized for external audit and technical review.

## Document Structure

### 1. [Introduction](./01-introduction.md)
- Project overview
- Core values and features
- Technology stack
- Use cases

### 2. [Architecture](./02-architecture.md)
- System architecture
- Contract hierarchy
- Design patterns
- Security architecture
- Gas optimization strategies

### 3. [Core Concepts](./03-concepts.md)
- Points system
- Season system
- Aggregation system
- Lazy snapshot system
- Reward distribution
- Pre-deposit feature

### 4. [Contract Details](./04-contracts.md)
- Individual contract specifications
- Function documentation
- State variables
- Access control structure
- Security features

### 5. [Workflows](./05-workflows.md)
- Project creation workflow
- Staking/unstaking workflows
- Season rollover process
- Reward distribution flow
- Points calculation flow
- Emergency procedures

### 6. [Technical Implementation](./06-technical.md)
- Points calculation
- O(1) aggregation system
- Lazy snapshot implementation
- CREATE2 deployment
- Code contract pattern
- Hook pattern
- Virtual season calculation
- EIP-1153 transient storage
- Access control with timelock
- SafeERC20 usage
- Custom errors
- Unchecked arithmetic

## Additional Documentation

### In Project Root

- **[README.md](../README.md)**: Quick start and overview
- **[DEPLOYMENT.md](../DEPLOYMENT.md)**: Deployment architecture guide
- **[PREDEPOSIT_GUIDE.md](../PREDEPOSIT_GUIDE.md)**: Pre-deposit feature guide
- **[TESTS.md](../TESTS.md)**: Test documentation and coverage

## For Auditors

### Key Areas to Review

1. **Security Critical**
   - Reentrancy protection (EIP-1153 transient storage)
   - Access control (3-day timelock)
   - Points calculation accuracy
   - Reward distribution fairness
   - Overflow safety

2. **Gas Efficiency**
   - O(1) aggregation system
   - Lazy snapshot mechanism
   - Unchecked arithmetic usage
   - Transient storage usage

3. **Core Logic**
   - Season management and rollover
   - Points accumulation
   - Reward proportional distribution
   - Pre-deposit feature

4. **State Management**
   - Season finalization
   - User data snapshots
   - Claim tracking
   - Aggregation updates

### Test Coverage

- Total Tests: 68
- Pass Rate: 100%
- Core Logic Coverage: 100%
- Test Files: `test/` directory

### Solidity Version

- **Version**: 0.8.28
- **EVM**: Supports EIP-1153 (transient storage)
- **Framework**: Foundry

### External Dependencies

- OpenZeppelin Contracts 5.x
  - `AccessControlDefaultAdminRules`
  - `ReentrancyGuardTransient`
  - `SafeERC20`
  - `Pausable`

## Contract Addresses

Contracts are not yet deployed. This documentation is for audit purposes.

## License

MIT License

## Version

**v1.0.0** - Ready for Audit

