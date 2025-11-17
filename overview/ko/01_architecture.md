# Cross GameReward Protocol - 아키텍처

## 📐 개요

Cross GameReward Protocol은 **rewardPerToken 누적 방식**을 사용하는 다중 풀 디파짓 시스템입니다.

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
│    CrossGameRewardRouter            │
│  • depositNative/withdrawNative     │
│  • depositERC20/withdrawERC20       │
│  • claimRewards/claimReward         │
│  • 재배포 가능                    │
└──────┬───────────────────────────┘
       │
       ├──► WCROSS (WETH9 패턴)
       │    • 누구나 deposit/withdraw
       │
       ▼
┌──────────────────────────────────┐
│     CrossGameReward                 │
│   • UUPS 업그레이더블             │
│   • createPool                   │
│   • setRouter                    │
└──────┬───────────────────────────┘
       │ creates
       ▼
┌──────────────────────────────────┐
│   CrossGameRewardPool × n           │
│   • UUPS 업그레이더블             │
│   • depositFor/withdrawFor          │
│   • rewardPerToken 누적          │
└──────────────────────────────────┘
```

---

## 🔧 컴포넌트 상세

### 1. WCROSS

**역할:** Native CROSS를 ERC20으로 래핑 (WETH9 패턴)

**주요 함수:**
```solidity
deposit() public payable              // 누구나 가능
withdraw(uint amount) external        // 누구나 가능
withdrawTo(address to, uint) public   // 누구나 가능
```

**특징:**
- WETH9 표준 패턴 준수
- Router 검사 제거 (접근성 향상)
- DEX 통합 용이
- 1:1 parity 유지

---

### 2. CrossGameReward

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
createPool(address depositToken, uint minDepositAmount)
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

### 3. CrossGameRewardPool

**역할:** 개별 디파짓 풀

**상태 변수:**
```solidity
IERC20 public depositToken;                        // 디파짓 토큰
ICrossGameReward public crossDeposit;                 // CrossGameReward 참조
uint public minDepositAmount;                      // 최소 디파짓 수량
uint public totalDeposited;                        // 전체 예치량
PoolStatus public poolStatus;                      // Active/Inactive/Paused
mapping(address => uint) public balances;          // 사용자 예치량

EnumerableSet.AddressSet private _rewardTokenAddresses;         // 활성 보상 토큰 목록
EnumerableSet.AddressSet private _removedRewardTokenAddresses;  // 제거된 보상 토큰 목록
mapping(IERC20 => RewardToken) private _rewardTokenData;        // 보상 토큰 데이터
mapping(address => mapping(IERC20 => UserReward)) public userRewards; // 사용자 보상
```

**주요 함수:**
```solidity
// Deposit/Withdraw
deposit(uint amount)                        // Active 상태에서만 가능
depositFor(address account, uint amount)    // Router 전용, Active 상태만
withdraw()                                 // Active/Inactive 상태 가능
withdrawFor(address account)               // Router 전용

// Claim (리팩토링 완료)
claimRewards()                            // 모든 보상 claim
claimRewardsFor(address account)           // Router 전용
claimReward(IERC20 token)                 // 특정 토큰만 claim
claimRewardFor(address account, token)     // Router 전용

// Admin
addRewardToken(IERC20 token)              // CrossGameReward만 호출 가능
removeRewardToken(IERC20 token)           // CrossGameReward만 호출 가능
withdraw(IERC20 token, address to)        // CrossGameReward만 호출 가능
setPoolStatus(uint8 status)               // CrossGameReward만 호출 가능
```

**Pool Status:**
- **Active**: 모든 작업 가능 (deposit, withdraw, claim)
- **Inactive**: deposit 불가, withdraw/claim만 가능
- **Paused**: 모든 작업 불가

> 제거된 보상 토큰은 `_removedRewardTokenAddresses`로 이동하며, `_withdraw` 과정에서 자동 정산·지급됩니다.
> totalDeposited=0 일 때 예치된 보상은 `withdrawableAmount`로 분류되어 owner가 회수할 수 있습니다.

**Access Control:**
- `onlyOwner()`: CrossGameReward의 owner만 가능 (upgrade 등)
- `onlyRewardRoot()`: CrossGameReward 컨트랙트만 가능 (관리 함수들)

---

### 4. CrossGameRewardRouter

**역할:** 사용자 인터페이스

**상태 변수:**
```solidity
CrossGameReward public immutable crossGameReward;
IWCROSS public immutable wcross;
```

**주요 함수:**
```solidity
// Deposit/Withdraw
depositNative(uint poolId) payable
withdrawNative(uint poolId)
depositERC20(uint poolId, uint amount)
depositERC20WithPermit(uint poolId, uint amount, ...) // EIP-2612
withdrawERC20(uint poolId)

// Claim (신규 추가)
claimRewards(uint poolId)                    // 모든 보상 claim
claimReward(uint poolId, address token)       // 특정 토큰만 claim

// View
getUserDepositInfo(uint poolId, address user)
getPendingRewards(uint poolId, address user)  // 모든 pending rewards
getPendingReward(uint poolId, address user, token) // 특정 토큰 pending
isNativePool(uint poolId)
```

**내부 함수:**
```solidity
_getPool(uint poolId) internal view
_getPoolAndValidateWCROSS(uint poolId) internal view
```

---

## 🔐 보안 메커니즘

### AccessControl

**CrossGameReward:**
```solidity
DEFAULT_ADMIN_ROLE      // 시스템 관리, 업그레이드
POOL_MANAGER_ROLE       // 풀 생성/관리
```

**CrossGameRewardPool:**
```solidity
DEFAULT_ADMIN_ROLE      // 풀 관리 (CrossGameReward이 보유)
REWARD_MANAGER_ROLE     // 보상 토큰 추가
PAUSER_ROLE             // 긴급 정지 (CrossGameReward이 보유)
```

### Router 권한 체크

```solidity
// CrossGameRewardPool
function _checkDelegate(address account) internal view {
    require(account != address(0), CGRPCanNotZeroAddress());
    require(msg.sender == ICrossGameReward(crossDeposit).router(), CGRPOnlyRouter());
}
```

**적용:**
- depositFor()
- withdrawFor()
- claimRewardsFor()
- claimRewardFor()

### WCROSS - WETH9 패턴

**Router 검사 제거:**
```solidity
function deposit() public payable {
    if (msg.value != 0) _mint(msg.sender, msg.value);
}

function withdrawTo(address to, uint amount) public {
    require(to != address(0), WCROSSInvalidAddress());
    _burn(msg.sender, amount);
    (bool success,) = to.call{value: amount}("");
    require(success, WCROSSTransferFailed());
}
```

**특징:**
- 누구나 deposit/withdraw 가능 (WETH9 표준)
- ERC20 메커니즘으로 보호
- DEX 통합 용이
- 보안성 유지 (검증된 패턴)

---

## 📊 보상 메커니즘

### rewardPerToken 누적 방식

**핵심 공식:**
```
rewardPerTokenStored += (newReward × 1e18) / totalDeposited
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
    if (totalDeposited == 0) return;
    
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
require(balances[msg.sender] > 0, CGRPNoDepositFound());

// 2. Effects
balances[msg.sender] = 0;
totalDeposited -= amount;

// 3. Interactions
depositToken.safeTransfer(msg.sender, amount);
```

### 3. 이벤트 기반 투명성

**모든 주요 액션에 이벤트:**
- Deposited, Withdrawn
- RewardSynced
- RewardClaimed
- PoolCreated
- PoolStatusChanged

---

## 🔄 업그레이드 전략

### UUPS Proxy 패턴

**CrossGameReward:**
```solidity
function _authorizeUpgrade(address newImplementation) 
    internal 
    override 
    onlyRole(DEFAULT_ADMIN_ROLE) 
{}
```

**CrossGameRewardPool:**
```solidity
function _authorizeUpgrade(address newImplementation) 
    internal 
    override 
    onlyRole(DEFAULT_ADMIN_ROLE) 
{}
```

**Storage Gap:**
```solidity
uint[50] private __gap;  // CrossGameReward
uint[41] private __gap;  // CrossGameRewardPool
```

### Router 교체

```solidity
// 새 Router 배포
CrossGameRewardRouter newRouter = new CrossGameRewardRouter(address(crossDeposit));

// CrossGameReward에서 Router 변경
crossDeposit.setRouter(address(newRouter));
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

**Cross GameReward Protocol은:**

1. Multi-Pool 디파짓 시스템
2. Native CROSS 지원
3. UUPS 업그레이더블
4. 역할 기반 보안
5. Gas 최적화
6. Production-ready

**다음**: [02_reward_mechanism.md](./02_reward_mechanism.md)
