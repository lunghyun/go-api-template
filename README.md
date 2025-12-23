# go-api-template

다양한 아키텍처 패턴으로 구현된 Go REST API 템플릿 모음

## 개요

같은 기능(학생 관리 CRUD)을 여러 아키텍처 패턴으로 구현한 학습용 템플릿 프로젝트입니다.
각 브랜치가 하나의 아키텍처 패턴을 보여주며, 프로젝트 시작점으로 활용할 수 있습니다.

## 브랜치 구조

- `main`: 프로젝트 개요 및 문서 (코드 없음)
- `layered`: 3-tier 레이어드 아키텍처 (Handler → Service → Repository)

## 기술 스택

- **Language:** Go 1.25
- **Web Framework:** Gin
- **Storage:** In-Memory (map 기반)
- **Testing:** testify/assert

## 기능

학생 관리 REST API (CRUD)

- `GET /students` - 전체 학생 목록 조회
- `GET /students/:id` - 특정 학생 조회
- `POST /students` - 학생 생성
- `PUT /students/:id` - 학생 수정
- `DELETE /students/:id` - 학생 삭제

## 사용법

각 브랜치를 체크아웃하여 실행:

```bash
# layered 브랜치로 이동
git checkout layered

# 의존성 설치
go mod download

# 실행
go run main.go

# 테스트
go test -v
```

## 아키텍처 패턴 비교

### Layered Architecture (3-Tier)
- **구조:** Handler → Service → Repository
- **특징:** 전통적인 3계층 구조, 간단하고 직관적
- **적합:** 중소규모 CRUD API, 빠른 개발

## 프로젝트 구조 (layered 브랜치)

```
.
├── main.go              # 애플리케이션 진입점
├── app.go               # 의존성 주입 및 초기화
├── student/             # 학생 도메인
│   ├── model.go         # 도메인 모델 및 검증
│   ├── repository.go    # 데이터 접근 계층
│   ├── service.go       # 비즈니스 로직 계층
│   ├── handler.go       # HTTP 핸들러
│   └── routes.go        # 라우팅 등록
└── docs/                # 문서
```

## Go 관용구 적용

- Interface 기반 의존성 주입
- 명시적 에러 반환 패턴
- `sort.Interface` 구현
- `sync.RWMutex`를 통한 동시성 제어
- 도메인 모델 검증 (`Validate()` 메서드)

## 기여

이슈 및 PR 환영합니다!
