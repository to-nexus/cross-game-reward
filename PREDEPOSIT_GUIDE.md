# Pre-deposit 기능 가이드

## 개요

Pre-deposit는 시즌 1이 시작되기 전에 미리 스테이킹할 수 있는 기능입니다. Pre-deposit를 통해 스테이킹하면 시즌 1 시작 블록부터 포인트가 누적되기 시작합니다.

## 주요 특징

### 1. 시즌 1 전용 기능
- Pre-deposit는 **오직 시즌 1에만** 적용됩니다.
- 시즌 2 이후부터는 일반 스테이킹만 가능합니다.

### 2. 블록 기반 타이밍
```
Timeline:
[preDepositStartBlock] -----> [firstSeasonStartBlock] -----> [Season 1 End]
     |                              |                              |
     Pre-deposit 시작           시즌 1 시작                    시즌 1 종료
     (스테이킹 가능)            (포인트 누적 시작)
```

### 3. 포인트 누적
- Pre-deposit로 스테이킹한 경우: **시즌 1 시작 블록**부터 포인트 누적
- 시즌 1 시작 후 스테이킹한 경우: **스테이킹한 블록**부터 포인트 누적

## 설정 방법

### 프로젝트 생성 시 설정

```solidity
// StakingProtocol.createProject() 호출 시
function createProject(
    string calldata projectName,
    uint seasonBlocks,
    uint firstSeasonStartBlock,  // 시즌 1 시작 블록
    uint poolEndBlock,
    address projectAdmin,
    uint preDepositStartBlock    // Pre-deposit 시작 블록 (0이면 비활성화)
) external returns (uint projectID, address stakingPool, address rewardPool)
```

### 예시

```solidity
// 현재 블록: 1000
// Pre-deposit 시작: 블록 1100
// 시즌 1 시작: 블록 1200

protocol.createProject(
    "MyProject",
    100,        // seasonBlocks: 100 블록
    1200,       // firstSeasonStartBlock: 블록 1200부터 시즌 1 시작
    0,          // poolEndBlock: 무한 (0)
    msg.sender, // projectAdmin
    1100        // preDepositStartBlock: 블록 1100부터 pre-deposit 가능
);
```

## 사용자 경험 (UI)

### 1. Pre-deposit 대기 중
```
현재 블록 < preDepositStartBlock
```
- 🟡 노란색 배너 표시
- "Pre-deposit는 블록 #1100부터 시작됩니다"
- 남은 블록 수 표시

### 2. Pre-deposit 기간
```
preDepositStartBlock <= 현재 블록 < firstSeasonStartBlock
```
- 🔵 파란색 배너 표시 (강조)
- "🎉 Pre-deposit 기간"
- "시즌 1 시작 전에 미리 스테이킹할 수 있습니다!"
- 시즌 시작까지 남은 블록 수 표시
- 스테이킹 가능

### 3. 시즌 1 시작 준비
```
현재 블록 >= firstSeasonStartBlock && currentSeason == 0
```
- 🟢 초록색 배너 표시
- "시즌 1 시작 준비 완료"
- "시즌 1이 곧 시작됩니다. 지금 스테이킹하세요!"

### 4. 시즌 1 진행 중
```
currentSeason >= 1
```
- 일반 스테이킹 UI 표시
- Pre-deposit 배너 숨김

## 스마트 컨트랙트 로직

### StakingPoolBase._stakeFor()

```solidity
function _stakeFor(address user, uint amount, address from) internal virtual {
    // ...
    
    _ensureSeason();
    
    // 시즌 활성 체크
    if (currentSeason == 0) {
        // 첫 시즌이 아직 생성되지 않음
        if (preDepositStartBlock > 0 && block.number >= preDepositStartBlock) {
            // ✅ preDeposit 기간: 스테이킹 가능
        } else {
            // ❌ preDeposit이 없거나 아직 preDeposit 블록 이전
            require(block.number >= nextSeasonStartBlock, StakingPoolBaseNoActiveSeason());
        }
    } else {
        // 시즌이 생성되었으면 일반 스테이킹
        require(isSeasonActive(), StakingPoolBaseNoActiveSeason());
    }
    
    // ... 스테이킹 로직
}
```

### 포인트 계산

Pre-deposit로 스테이킹한 경우, `_calculateCurrentSeasonPoints()`에서:

```solidity
uint lastUpdate = position.lastUpdateBlock;
if (lastUpdate < current.startBlock && position.balance > 0) {
    // Pre-deposit 케이스: 시즌 시작 블록부터 계산
    return PointsLib.calculatePoints(
        position.balance, 
        current.startBlock,  // ✅ 시즌 시작 블록부터
        block.number, 
        blockTime, 
        pointsTimeUnit
    );
}
```

## 웹앱 구현

### StakingPanel.tsx

```typescript
// Pre-deposit 정보 로드
const poolInfo = await stakingViewer.getPoolInfo(project.id);
const preDepositBlock = poolInfo[5]; // preDepositStartBlock
const firstSeasonBlock = poolInfo[6]; // firstSeasonStartBlock

// 현재 블록
const currentBlock = await provider.getBlockNumber();

// 상태 판단
if (currentSeason === 0) {
  if (preDepositBlock > 0) {
    if (currentBlock >= preDepositBlock && currentBlock < firstSeasonBlock) {
      // Pre-deposit 기간
      setIsPreDepositPeriod(true);
    } else if (currentBlock < preDepositBlock) {
      // Pre-deposit 대기
      setIsBeforePreDeposit(true);
    }
  }
}
```

### StakingViewer.getPoolInfo()

```solidity
function getPoolInfo(uint projectID)
    external
    view
    returns (
        uint blockTime,
        uint pointsTimeUnit,
        uint seasonBlocks,
        uint poolEndBlock,
        uint currentSeason,
        uint preDepositStartBlock,      // ✅ Pre-deposit 시작 블록
        uint firstSeasonStartBlock      // ✅ 첫 시즌 시작 블록
    )
{
    (IStakingPool pool,) = _getPools(projectID);
    
    // ... 정보 조회
    preDepositStartBlock = pool.preDepositStartBlock();
    
    // firstSeasonStartBlock 계산
    if (currentSeason > 0) {
        // 시즌이 시작된 경우: 역산
        (uint season, uint startBlock,,) = pool.getCurrentSeasonInfo();
        firstSeasonStartBlock = startBlock - ((season - 1) * seasonBlocks);
    } else {
        // 시즌 시작 전: nextSeasonStartBlock
        firstSeasonStartBlock = pool.nextSeasonStartBlock();
    }
    
    return (...);
}
```

## 테스트 시나리오

### 1. Pre-deposit 전 스테이킹 시도
```
현재 블록: 1000
preDepositStartBlock: 1100
firstSeasonStartBlock: 1200

❌ 스테이킹 실패: StakingPoolBaseNoActiveSeason()
```

### 2. Pre-deposit 기간 스테이킹
```
현재 블록: 1150
preDepositStartBlock: 1100
firstSeasonStartBlock: 1200

✅ 스테이킹 성공
- position.lastUpdateBlock = 1150
- 포인트는 아직 0 (시즌 미시작)
```

### 3. 시즌 1 시작 후 포인트 확인
```
현재 블록: 1250
시즌 1 시작 블록: 1200
position.lastUpdateBlock: 1150 (pre-deposit)

✅ 포인트 계산:
- fromBlock = 1200 (시즌 시작 블록)
- toBlock = 1250 (현재 블록)
- points = balance * (1250 - 1200) * PRECISION / timeUnit
```

### 4. 시즌 1 시작 후 스테이킹
```
현재 블록: 1250
시즌 1 시작 블록: 1200

✅ 스테이킹 성공
- position.lastUpdateBlock = 1250
- 포인트 계산: fromBlock = 1250 (스테이킹 블록)
```

## 주의사항

### 1. Pre-deposit는 시즌 1 전용
- `preDepositStartBlock`은 시즌 1에만 적용
- 시즌 2 이후는 일반 스테이킹만 가능

### 2. 블록 순서
```
preDepositStartBlock < firstSeasonStartBlock
```
- 이 순서가 지켜지지 않으면 로직 오류 발생 가능
- 프로젝트 생성 시 검증 필요

### 3. 0 값 처리
```solidity
if (preDepositStartBlock == 0) {
    // Pre-deposit 비활성화
    // firstSeasonStartBlock부터만 스테이킹 가능
}
```

### 4. 포인트 표시
- Pre-deposit 후 시즌 시작 전: 포인트 0 (정상)
- 시즌 시작 후: 시즌 시작 블록부터 계산된 포인트 표시

## FAQ

### Q1: Pre-deposit와 일반 스테이킹의 차이는?
**A:** Pre-deposit는 시즌 시작 전에 스테이킹하지만, 포인트는 시즌 시작 블록부터 누적됩니다. 일반 스테이킹은 스테이킹한 블록부터 포인트가 누적됩니다.

### Q2: Pre-deposit 후 언스테이킹하면?
**A:** 일반 언스테이킹과 동일하게 처리됩니다. 시즌 시작 전이면 포인트가 0이므로 손실 없이 출금 가능합니다.

### Q3: Pre-deposit를 비활성화하려면?
**A:** `preDepositStartBlock`을 0으로 설정하면 됩니다. 이 경우 `firstSeasonStartBlock`부터만 스테이킹 가능합니다.

### Q4: 시즌 2에서도 pre-deposit가 가능한가요?
**A:** 아니요, pre-deposit는 시즌 1 전용 기능입니다. 시즌 2부터는 일반 스테이킹만 가능합니다.

## 관련 파일

### 스마트 컨트랙트
- `src/base/StakingPoolBase.sol` - Pre-deposit 로직 구현
- `src/StakingPool.sol` - 포인트 계산 로직
- `src/StakingViewer.sol` - Pre-deposit 정보 조회

### 웹앱
- `webapp/src/components/StakingPanel.tsx` - Pre-deposit UI
- `webapp/src/hooks/useContracts.ts` - 컨트랙트 연결

### 테스트
- `test/BaseTest.sol` - Pre-deposit 테스트 설정
- `test/Season.t.sol` - Pre-deposit 시나리오 테스트

## 결론

Pre-deposit 기능은 사용자가 시즌 시작 전에 미리 준비할 수 있도록 하여 더 나은 사용자 경험을 제공합니다. 시즌 1 시작과 동시에 포인트 누적이 시작되므로, 초기 참여자에게 공정한 기회를 제공합니다.

