# 애드온 시스템 가이드

## 📋 개요

애드온 시스템은 코어 스테이킹 로직을 수정하지 않고 추가 기능을 연결할 수 있는 플러그인 아키텍처입니다.

## 🎯 사용 사례

### 1. **랭킹 시스템** (`RankingAddon`)
- 스테이킹 금액 기반 리더보드
- 시즌별 Top 스테이커 추적
- 경쟁 요소 추가

### 2. **배지/업적 시스템**
- 마일스톤 달성 추적
- NFT 배지 발급
- 게이미피케이션

### 3. **통계 대시보드**
- 실시간 통계 집계
- 히스토리 추적
- 분석 데이터 수집

### 4. **알림 시스템**
- 온체인 이벤트 트리거
- 오프체인 알림 연동
- Discord/Telegram 봇 연동

## 🔧 구현 방법

### Step 1: IStakingAddon 인터페이스 구현

```solidity
contract MyAddon is IStakingAddon {
    address public immutable stakingPool;
    
    constructor(address _stakingPool) {
        stakingPool = _stakingPool;
    }
    
    modifier onlyPool() {
        require(msg.sender == stakingPool, "Only pool");
        _;
    }
    
    function onStake(
        address user, 
        uint amount, 
        uint oldBalance, 
        uint newBalance, 
        uint season
    ) external onlyPool {
        // 스테이킹 시 로직
    }
    
    function onWithdraw(
        address user, 
        uint amount, 
        uint season
    ) external onlyPool {
        // 출금 시 로직
    }
    
    function onSeasonEnd(
        uint season, 
        uint totalStaked, 
        uint totalPoints
    ) external onlyPool {
        // 시즌 종료 시 로직
    }
    
    function onClaim(
        address user, 
        uint season, 
        uint points, 
        uint rewardAmount
    ) external onlyPool {
        // 보상 청구 시 로직
    }
}
```

### Step 2: 애드온 배포 및 설정

```solidity
// 1. 애드온 배포
MyAddon addon = new MyAddon(address(stakingPool));

// 2. StakingPool에 애드온 설정 (Admin 권한 필요)
stakingPool.setStakingAddon(address(addon));

// 3. 제거 시
stakingPool.setStakingAddon(address(0));
```

## ⚙️ 작동 방식

### 호출 흐름

```
User -> StakingPool.stake()
  └─> _stakeFor()
      └─> _afterStake() [Hook]
          └─> _callAddonSafe()
              └─> addon.onStake() [try/catch로 안전하게 호출]
```

### 실패 처리

- 애드온 호출 실패 시 **메인 로직은 영향 없음**
- `AddonCallFailed` 이벤트로 실패 로그 기록
- 가스 부족이나 revert는 격리됨

## 📊 가스 비용

| 작업 | 추가 가스 비용 |
|------|--------------|
| 애드온 없음 | 0 gas |
| 애드온 호출 (성공) | ~2,000-5,000 gas |
| 애드온 호출 (실패) | ~3,000 gas |

## 🛡️ 보안 고려사항

### 1. **Only Pool Modifier 필수**
```solidity
modifier onlyPool() {
    require(msg.sender == stakingPool, "Only pool");
    _;
}
```

### 2. **Reentrancy 방지**
- 애드온에서 StakingPool 다시 호출 금지
- 외부 호출 시 주의

### 3. **가스 제한**
- 과도한 연산 피하기
- 무한 루프 방지

### 4. **권한 관리**
- Admin만 애드온 설정 가능
- 애드온 변경 전 충분한 검증

## 💡 Best Practices

### 1. **최소 상태 변경**
```solidity
// ❌ Bad: 과도한 스토리지 쓰기
function onStake(...) external onlyPool {
    for (uint i = 0; i < 1000; i++) {
        data[i] = value;
    }
}

// ✅ Good: 필요한 것만 저장
function onStake(...) external onlyPool {
    summary[season] += amount;
}
```

### 2. **이벤트 활용**
```solidity
// 애드온 내부 상태 변화는 이벤트로 추적
emit RankingUpdated(season, user, score);
```

### 3. **View 함수 제공**
```solidity
// 외부에서 조회 가능한 함수 제공
function getUserRankingScore(uint season, address user) 
    external view returns (uint);
```

## 🔄 업그레이드 패턴

### 애드온 교체
```solidity
// 1. 새 애드온 배포
MyAddonV2 addonV2 = new MyAddonV2(address(stakingPool));

// 2. 기존 데이터 마이그레이션 (필요시)
addonV2.migrateData(address(oldAddon));

// 3. 교체
stakingPool.setStakingAddon(address(addonV2));
```

## 📚 예제 코드

### 예제 1: 간단한 카운터
```solidity
contract StakeCounterAddon is IStakingAddon {
    mapping(uint => uint) public stakeCount;
    
    function onStake(...) external onlyPool {
        stakeCount[season]++;
    }
    
    function onWithdraw(...) external onlyPool {}
    function onSeasonEnd(...) external onlyPool {}
    function onClaim(...) external onlyPool {}
}
```

### 예제 2: 마일스톤 배지
```solidity
contract MilestoneAddon is IStakingAddon {
    mapping(address => uint) public milestoneLevel;
    
    function onStake(
        address user, 
        uint, 
        uint, 
        uint newBalance, 
        uint
    ) external onlyPool {
        if (newBalance >= 1000e18 && milestoneLevel[user] < 1) {
            milestoneLevel[user] = 1;
            emit MilestoneReached(user, 1);
        }
        if (newBalance >= 10000e18 && milestoneLevel[user] < 2) {
            milestoneLevel[user] = 2;
            emit MilestoneReached(user, 2);
        }
    }
    
    // ... 나머지 구현
}
```

## 🚀 다음 단계

1. 커스텀 애드온 설계
2. 테스트 작성
3. 가스 최적화
4. 배포 및 설정
5. 모니터링

---

**참고**: 애드온은 강력하지만 신중하게 사용해야 합니다. 충분한 테스트와 감사를 거친 후 프로덕션에 적용하세요.

