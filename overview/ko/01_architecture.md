# Cross Staking Protocol - 아키텍처

## 📐 개요

Cross Staking Protocol은 **rewardPerToken 누적 방식**을 사용하는 다중 풀 스테이킹 시스템입니다.

### 핵심 특징

- ✅ 다중 풀 지원 (무제한)
- ✅ Native CROSS 지원 (자동 래핑)
- ✅ 다중 보상 토큰
- ✅ UUPS 업그레이더블
- ✅ 역할 기반 접근 제어
- ✅ Gas 최적화 (O(1))

---

## 🏗️ 시스템 구조

### 4개 핵심 컨트랙트

```
┌──────────────────────────────────┐
│          User                     │
│    (Native CROSS / ERC20)         │
└──────────────┬───────────────────┘
               │
               ▼
┌──────────────────────────────────┐
│    CrossStakingRouter            │
│  • stakeNative/unstakeNative     │
│  • stakeERC20/unstakeERC20       │
│  • 재배포 가능                    │
└──────┬───────────────────────────┘
       │
       ├──► WCROSS
       │    • Router 전용
       │
       ▼
┌──────────────────────────────────┐
│     CrossStaking                 │
│   • UUPS 업그레이더블             │
│   • createPool                   │
│   • setRouter                    │
└──────┬───────────────────────────┘
       │ creates
       ▼
┌──────────────────────────────────┐
│   CrossStakingPool × n           │
│   • UUPS 업그레이더블             │
│   • stakeFor/unstakeFor          │
│   • rewardPerToken 누적          │
└──────────────────────────────────┘
```

---

## 🔧 컴포넌트 상세

### 1. WCROSS

**역할:** Native CROSS를 ERC20으로 래핑

**상태 변수:**
```solidity
CrossStaking public staking;  // CrossStaking 참조
```

**주요 함수:**
```solidity
deposit() public payable      // Router만 가능
withdraw(uint amount)         // Router만 가능
```

**접근 제어:**
- `msg.sender == staking.router()` 검증

---

### 2. CrossStaking

**역할:** 풀 팩토리 및 관리자

**상태 변수:**
```solidity
address public wcross;                    // WCROSS 주소
address public router;                    // Router 주소
address public poolImplementation;        // Pool implementation
mapping(uint => PoolInfo) public pools;   // 풀 정보
```

**주요 함수:**
```solidity
createPool(address stakingToken, uint minStakeAmount)
  returns (uint poolId, address poolAddress)

addRewardToken(uint poolId, address rewardToken)
removeRewardToken(uint poolId, address rewardToken)
setPoolStatus(uint poolId, uint8 status)  // 0=Active, 1=Inactive, 2=Paused
withdrawFromPool(uint poolId, address token, address to)
setRouter(address _router)
```

**Roles:**
- DEFAULT_ADMIN_ROLE (owner)
- MANAGER_ROLE (풀 및 보상 관리)

---

### 3. CrossStakingPool

**역할:** 개별 스테이킹 풀

**상태 변수:**
```solidity
IERC20 public stakingToken;                        // 스테이킹 토큰
ICrossStaking public crossStaking;                 // CrossStaking 참조
uint public minStakeAmount;                        // 최소 스테이킹 수량
uint public totalStaked;                           // 전체 예치량
PoolStatus public poolStatus;                      // Active/Inactive/Paused
mapping(address => uint) public balances;          // 사용자 예치량

EnumerableSet.AddressSet private _rewardTokenAddresses;         // 활성 보상 토큰 목록
EnumerableSet.AddressSet private _removedRewardTokenAddresses;  // 제거된 보상 토큰 목록
mapping(IERC20 => RewardToken) private _rewardTokenData;        // 보상 토큰 데이터
mapping(address => mapping(IERC20 => UserReward)) public userRewards; // 사용자 보상
```

**주요 함수:**
```solidity
stake(uint amount)                        // Active 상태에서만 가능
stakeFor(address account, uint amount)    // Router 전용, Active 상태만
unstake()                                 // Active/Inactive 상태 가능
unstakeFor(address account)               // Router 전용
claimRewards()                            // Active/Inactive 상태 가능
claimReward(IERC20 token)
addRewardToken(IERC20 token)              // CrossStaking만 호출 가능
removeRewardToken(IERC20 token)           // CrossStaking만 호출 가능
withdraw(IERC20 token, address to)        // CrossStaking만 호출 가능
setPoolStatus(uint8 status)               // CrossStaking만 호출 가능
```

**Pool Status:**
- **Active**: 모든 작업 가능 (stake, unstake, claim)
- **Inactive**: stake 불가, unstake/claim만 가능
- **Paused**: 모든 작업 불가

> 제거된 보상 토큰은 `_removedRewardTokenAddresses`로 이동하며, `_unstake` 과정에서 자동 정산·지급됩니다.
> totalStaked=0 일 때 예치된 보상은 `withdrawableAmount`로 분류되어 owner가 회수할 수 있습니다.

**Access Control:**
- `onlyOwner()`: CrossStaking의 owner만 가능 (upgrade 등)
- `onlyStakingRoot()`: CrossStaking 컨트랙트만 가능 (관리 함수들)

---

### 4. CrossStakingRouter

**역할:** 사용자 인터페이스

**상태 변수:**
```solidity
CrossStaking public immutable crossStaking;
IWCROSS public immutable wcross;
```

**주요 함수:**
```solidity
stakeNative(uint poolId) payable
unstakeNative(uint poolId)
stakeERC20(uint poolId, uint amount)
unstakeERC20(uint poolId)
```

**Helper 함수:**
```solidity
_getPool(uint poolId) internal view
_getPoolAndValidateWCROSS(uint poolId) internal view
```

---

## 🔐 보안 메커니즘

### AccessControl

**CrossStaking:**
```solidity
DEFAULT_ADMIN_ROLE      // 시스템 관리, 업그레이드
POOL_MANAGER_ROLE       // 풀 생성/관리
```

**CrossStakingPool:**
```solidity
DEFAULT_ADMIN_ROLE      // 풀 관리 (CrossStaking이 보유)
REWARD_MANAGER_ROLE     // 보상 토큰 추가
PAUSER_ROLE             // 긴급 정지 (CrossStaking이 보유)
```

### Router 권한 체크

```solidity
// CrossStakingPool
function _checkDelegate(address account) internal view {
    require(account != address(0), CSPCanNotZeroAddress());
    require(msg.sender == ICrossStaking(crossStaking).router(), CSPOnlyRouter());
}
```

**적용:**
- stakeFor()
- unstakeFor()

### WCROSS 권한 체크

```solidity
function deposit() public payable {
    require(msg.sender == staking.router(), WCROSSUnauthorized());
    // ...
}
```

---

## 📊 보상 메커니즘

### rewardPerToken 누적 방식

**핵심 공식:**
```
rewardPerTokenStored += (newReward × 1e18) / totalStaked
userReward = userBalance × (rewardPerTokenStored - userCheckpoint) / 1e18
```

**특징:**
- O(1) 가스 비용
- 사용자 수 무관
- 예치 시점 이후 보상만
- 공정한 분배

### 스테이커 없을 때 보상

```solidity
function _syncReward(address tokenAddress) internal {
    // 스테이커가 없으면 동기화하지 않음
    if (totalStaked == 0) return;
    
    // lastBalance 업데이트 안함
    // → 다음 스테이커가 모두 받음
}
```

---

## 🎯 설계 원칙

### 1. Pull over Push

**보상 수령 방식:**
- 사용자가 직접 claim
- 가스비 예측 가능
- 재진입 공격 방어

### 2. Checks-Effects-Interactions

**모든 함수에서 준수:**
```solidity
// 1. Checks
require(balances[msg.sender] > 0, CSPNoStakeFound());

// 2. Effects
balances[msg.sender] = 0;
totalStaked -= amount;

// 3. Interactions
stakingToken.safeTransfer(msg.sender, amount);
```

### 3. 이벤트 기반 투명성

**모든 주요 액션에 이벤트:**
- Staked, Unstaked
- RewardSynced
- RewardClaimed
- PoolCreated
- PoolStatusChanged

---

## 🔄 업그레이드 전략

### UUPS Proxy 패턴

**CrossStaking:**
```solidity
function _authorizeUpgrade(address newImplementation) 
    internal 
    override 
    onlyRole(DEFAULT_ADMIN_ROLE) 
{}
```

**CrossStakingPool:**
```solidity
function _authorizeUpgrade(address newImplementation) 
    internal 
    override 
    onlyRole(DEFAULT_ADMIN_ROLE) 
{}
```

**Storage Gap:**
```solidity
uint[50] private __gap;  // CrossStaking
uint[41] private __gap;  // CrossStakingPool
```

### Router 교체

```solidity
// 새 Router 배포
CrossStakingRouter newRouter = new CrossStakingRouter(address(crossStaking));

// CrossStaking에서 Router 변경
crossStaking.setRouter(address(newRouter));
```

---

## 📚 참고

### 사용된 컴포넌트

**OpenZeppelin Contracts v5.4.0:**
- AccessControlDefaultAdminRulesUpgradeable
- PausableUpgradeable
- ReentrancyGuardTransientUpgradeable
- UUPSUpgradeable
- SafeERC20
- EnumerableSet

### 표준

- **ERC-20**: Token Standard
- **EIP-1967**: Proxy Storage Slots
- **EIP-1153**: Transient Storage

---

## 📖 요약

**Cross Staking Protocol은:**

1. Multi-Pool 스테이킹 시스템
2. Native CROSS 지원
3. UUPS 업그레이더블
4. 역할 기반 보안
5. Gas 최적화
6. Production-ready

**다음**: [02_reward_mechanism.md](./02_reward_mechanism.md)
