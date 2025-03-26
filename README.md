# InsuBridge

보험 브릿지 서비스

## 프로젝트 구조

```
.
├── api/        # API 정의
├── cmd/        # 메인 애플리케이션
│   └── server/ # 서버 진입점
├── configs/    # 설정 파일
├── docs/       # 문서
├── internal/   # 내부 패키지
└── pkg/        # 외부에서 사용 가능한 패키지
```

## 실행 방법

### 개발 모드
```bash
make dev
```

### 빌드
```bash
make build
```

### 실행
```bash
make run
```

### 테스트
```bash
make test
```

## 요구사항
- Go 1.16 이상
- Make
- Air (개발 모드용, 선택사항) 