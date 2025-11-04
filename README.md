# Cross Staking Protocol

Native CROSS와 ERC20 토큰을 위한 다중 풀 스테이킹 프로토콜

## 🎯 개요

Cross Staking Protocol은 확장 가능한 다중 풀 아키텍처를 통해 다양한 토큰의 스테이킹을 지원하는 프로토콜입니다.

### 핵심 구성

```
사용자 인터페이스
    ↓
CrossStakingRouter → WCROSS
    ↓
CrossStaking (풀 관리)
    ↓
CrossStakingPool × n (개별 풀)
```

## ✨ 주요 특징

- ✅ **Native CROSS 지원**: 자동 래핑/언래핑
- ✅ **다중 풀**: 토큰별 독립적인 스테이킹 풀
- ✅ **다중 보상**: 풀당 여러 보상 토큰
- ✅ **rewardPerToken 누적**: 효율적인 보상 분배 (O(1))
- ✅ **UUPS 업그레이더블**: 시스템 업그레이드 가능
- ✅ **역할 기반 권한**: 세밀한 접근 제어

## 🚀 빠른 시작

### 설치

```bash
forge install
```

### 테스트

```bash
forge test
```

**결과: 159/159 테스트 통과 (100%)**

### 배포

```bash
forge script script/DeployFullSystem.s.sol:DeployFullSystem \
  --rpc-url <RPC_URL> \
  --private-key <PRIVATE_KEY> \
  --broadcast
```

## 💡 사용 예시

### 사용자: Native CROSS 스테이킹

```solidity
// WCROSS approve (최초 1회)
wcross.approve(address(router), type(uint).max);

// Native CROSS 스테이킹
router.stakeNative{value: 100 ether}(poolId);

// 언스테이킹 (보상 포함)
router.unstakeNative(poolId);
// → Native CROSS + 모든 보상 수령
```

### 관리자: 풀 생성

```solidity
// Native CROSS 풀 생성
(uint poolId, address poolAddr) = crossStaking.createPool(
    address(wcross),
    2 days
);

// 보상 토큰 추가
crossStaking.addRewardToken(poolId, address(usdt));

// 보상 입금 (누구나)
usdt.transfer(poolAddr, 1000 ether);
```

## 🏗️ 아키텍처

### 4개 핵심 컨트랙트

#### 1. WCROSS
- Native CROSS를 ERC20으로 래핑
- Router 전용 (deposit/withdraw)

#### 2. CrossStaking (UUPS)
- 풀 팩토리 및 관리자
- 풀 생성 (POOL_MANAGER_ROLE)
- Router 설정 (DEFAULT_ADMIN_ROLE)

#### 3. CrossStakingPool (UUPS)
- 개별 스테이킹 풀
- rewardPerToken 누적 보상 분배
- stakeFor/unstakeFor (Router용)

#### 4. CrossStakingRouter
- 사용자 인터페이스
- Native CROSS/ERC20 스테이킹
- 재배포 가능

## 🔑 역할 (Roles)

### CrossStaking
- **DEFAULT_ADMIN_ROLE**: 시스템 관리
- **POOL_MANAGER_ROLE**: 풀 생성/관리

### CrossStakingPool
- **DEFAULT_ADMIN_ROLE**: 풀 관리 (CrossStaking)
- **REWARD_MANAGER_ROLE**: 보상 토큰 관리
- **PAUSER_ROLE**: 긴급 정지

## 📊 보상 메커니즘

### rewardPerToken 누적 방식

```
누적 토큰당 보상 = 모든 보상의 합계
사용자 보상 = 예치량 × (현재 누적 - 사용자 체크포인트)
```

**특징:**
- O(1) 가스 비용 (사용자 수 무관)
- 예치 시점 이후 보상만 수령
- 지분율에 따른 공정한 분배
- 스테이커 없을 때 입금된 보상은 첫 스테이커가 받음

## 🔒 보안

- ✅ **ReentrancyGuardTransient**: 재진입 방지
- ✅ **SafeERC20**: 안전한 토큰 전송
- ✅ **AccessControl**: 역할 기반 권한
- ✅ **Pausable**: 긴급 정지
- ✅ **UUPS**: 안전한 업그레이드
- ✅ **Router Check**: 권한 검증
- ✅ **Custom Errors**: 타입 안전

## 📚 문서

- [Architecture (ko)](overview/ko/01_architecture.md) · [Architecture (en)](overview/en/01_architecture.md)
- [Reward Mechanism (ko)](overview/ko/02_reward_mechanism.md) · [Reward Mechanism (en)](overview/en/02_reward_mechanism.md)
- [Security & Testing (ko)](overview/ko/03_security_and_testing.md) · [Security & Testing (en)](overview/en/03_security_and_testing.md)
- [Test Guide](test/README.md) - 테스트 가이드

## 🧪 테스트

### 실행

```bash
# 전체 테스트
forge test

# 특정 컨트랙트
forge test --match-contract WCROSS

# Gas 리포트
forge test --gas-report
```

### 통계

- **총 테스트**: 159개
- **성공률**: 100%
- **커버리지**: ~100%

## 🔄 업그레이드

### CrossStaking

```solidity
CrossStaking newImpl = new CrossStaking();
crossStaking.upgradeToAndCall(address(newImpl), "");
```

### CrossStakingPool

```solidity
CrossStakingPool newImpl = new CrossStakingPool();
pool.upgradeToAndCall(address(newImpl), "");
```

### Router

```solidity
// 새 Router 배포 및 교체
CrossStakingRouter newRouter = new CrossStakingRouter(address(crossStaking));
crossStaking.setRouter(address(newRouter));
```

## 📜 라이선스

MIT

## 🔗 참고

- [OpenZeppelin Contracts](https://docs.openzeppelin.com/contracts/)
- [Foundry Book](https://book.getfoundry.sh/)
