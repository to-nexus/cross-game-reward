# Changelog

## [1.0.0] - 2025-10-27

### 🎉 Initial Release

Cross Staking Protocol v1.0.0의 첫 번째 공식 릴리즈입니다.

### ✨ Features

#### Core Contracts
- **StakingProtocol**: 프로젝트 생성 및 관리를 위한 팩토리 컨트랙트
- **StakingPool**: 프로젝트별 스테이킹 풀 (시즌 기반, 포인트 시스템)
- **RewardPool**: 시즌별 보상 분배 시스템
- **StakingRouter**: 네이티브 토큰 래핑/언래핑 및 라우팅
- **StakingViewer**: 모든 조회 함수를 통합한 뷰어 컨트랙트
- **WCROSS**: WETH 스타일의 래핑된 CROSS 토큰

#### Key Features
- **시즌 기반 스테이킹**: 블록 기반 시즌 시스템으로 명확한 보상 주기 관리
- **포인트 시스템**: `balance * time` 기반의 공정한 포인트 계산
- **자동 롤오버**: 시즌 자동 전환 (최대 50시즌)
- **Pre-deposit**: 시즌 1 시작 전 사전 예치 기능
- **가상 시즌 계산**: 온체인 롤오버 없이도 미래 시즌 포인트 조회 가능
- **역할 기반 접근 제어**: OpenZeppelin AccessControl 사용
- **리엔트런시 방어**: EIP-1153 기반 Transient Storage 활용
- **일시 정지 기능**: 긴급 상황 대응을 위한 Pausable 패턴

#### Security
- 3일 Admin 지연 (AccessControlDefaultAdminRules)
- Reentrancy Guard (Transient Storage)
- SafeERC20 사용
- 최소 스테이킹 금액 (1 CROSS)
- 철저한 입력 검증

### 🔧 Technical Details

#### Architecture
- **CREATE2**: 결정적 주소 생성
- **Lazy Evaluation**: 필요 시에만 계산하여 가스 최적화
- **Virtual Season**: 온체인 상태 변경 없이 미래 시즌 데이터 계산
- **Separation of Concerns**: 조회 로직을 StakingViewer로 분리

#### Gas Optimization
- Transient Storage for reentrancy guard
- Batch operations support
- Efficient point calculation with caching
- Minimal storage updates

### 📚 Documentation
- 완전한 한글 문서 (`docs/` 디렉토리)
- 배포 가이드 (DEPLOYMENT.md, TESTNET_DEPLOYMENT.md)
- 테스트 가이드 (TESTS.md)
- Pre-deposit 가이드 (PREDEPOSIT_GUIDE.md)
- 스크립트 문서 (script/README.md, script/QUICK_START.md)

### 🧪 Testing
- 83개 테스트 통과
- Fuzz 테스트 포함
- 다양한 시나리오 커버리지

### 🌐 Web Application
- React + TypeScript + Vite
- Tailwind CSS
- ethers.js v6
- MetaMask 통합
- 실시간 데이터 업데이트
- Pre-deposit UI 지원

### 🔄 Changes from Development

#### Removed
- Addon 시스템 제거 (확장성보다 단순성 우선)
- Hook 함수 제거
- 불필요한 임시 문서 제거

#### Fixed
- 과거 시즌 포인트가 계속 증가하는 버그 수정
- 리워드 계산 정확도 개선
- `getSeasonUserPoints`가 `totalPoints`도 반환하도록 수정
- MetaMask 잠금 상태 처리 개선
- 웹앱 현재 블록 표시 오류 수정

#### Improved
- 모든 수정 내역 주석 제거
- 코드 정리 및 최적화
- 문서 통합 및 정리
- 일관된 NatSpec 주석

### 📦 Deployment

#### Mainnet (XDC Network)
- 배포 준비 완료
- 감사 대기 중

#### Testnet (Apothem)
- 완전히 배포 및 테스트 완료
- 웹앱 연동 확인

### 🙏 Acknowledgments
- OpenZeppelin for secure contract libraries
- Foundry for excellent development tools
- XDC Network for the blockchain infrastructure

---

## Future Plans

### v1.1.0 (Planned)
- 다중 토큰 스테이킹 지원
- 고급 통계 대시보드
- 모바일 최적화

### v2.0.0 (Planned)
- Layer 2 지원
- 크로스체인 브리지
- DAO 거버넌스

---

**Full Changelog**: Initial Release

