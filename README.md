# Restapi 구현
## 목표
RESTful API를 알아보자. 그리고 api 서버 만들기
## 내용
- Restful api 정의
- restful api 서버 만들기
- 웹 서버 발전 과정
- gin으로 서버 만들기

## Restful api란?
REST는 자원을 이름으로 구분, 자원 상태를 주고 받는 소프트웨어 아키텍처
Restful api는 rest 규약을 따르는 api를 말함
웹서버에서는 url과 http 메서드로 데이터와 동작을 정의하는 방식을 의미

## 효과
- url 과 메서드를 사용해 데이터, 동작 정의 -> 동일한 방식으로 처리
- fe, be가 분리 -> restapi가 보편화, 데이터 프로바이더로써 웹 서버 역할이 중요
- restful api를 사용 -> 여러 서버가 동일한 인터페이스로 통신

# 과정
1. gorilla/mux 같은 restful api 웹서버 지원 패키지 설치
2. restful api에 맞추서 웹 핸들러 함수 정의
3. restful api 테스트 코드
4. 웹 브라우저로 데이터 조회

## 대충
GET /student
GET /student/id
POST /student/id
DELETE /student/id

# 시작
- 이전보다 간단하게 패키지를 써보자잉
- gorilla/mux → gin으로 전환

---

## 프로젝트 구조

```
mustHaveGoRest/
├── main.go              # 서버 시작점
├── app.go               # 라우터 설정
├── main_test.go         # 통합 테스트
└── student/             # student 도메인
    ├── model.go         # Student 구조체, JSON 태그
    ├── repository.go    # 데이터 저장소 (메모리)
    ├── handler.go       # HTTP 핸들러
    └── routes.go        # 라우팅 등록
```

---

## 실행 플로우

### 1. 서버 시작 (main.go → app.go)
```
main()
  ↓
MakeWebHandler()
  ↓ gin.Default()
gin.Engine 생성 (로깅, 복구 미들웨어 포함)
  ↓
student 패키지 import
  ↓
init() 자동 실행 (데이터 초기화)
  ↓
RegisterRoutes() (라우팅 등록)
  ↓
engin.Run(":8080") - 서버 시작
```

### 2. 요청 처리 흐름
```
클라이언트 요청 (예: GET /students/1)
  ↓
Gin Router (URL 패턴 매칭)
  ↓
Middleware (로깅, 인증 등)
  ↓
Handler (GetStudentHandler)
  ↓
Repository (students 데이터 조회)
  ↓
JSON 응답 반환
```

---

## Gorilla Mux → Gin 전환: 추상화 비교

### Gorilla Mux (저수준)
- Go 표준 라이브러리(`net/http`) 위에 라우팅만 추가
- **코드 예시**:
  ```go
  func Handler(w http.ResponseWriter, r *http.Request) {
      w.WriteHeader(http.StatusOK)
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(data)  // 3단계
  }
  ```

### Gin (고수준)
- 웹 프레임워크 - 많은 기능 자동화
- **코드 예시**:
  ```go
  func Handler(c *gin.Context) {
      c.JSON(http.StatusOK, data)  // 1단계
  }
  ```

### 추상화 수준 비교

| 항목 | Gorilla Mux | Gin                             |
|------|-------------|---------------------------------|
| **요청/응답** | `ResponseWriter`, `Request` 분리 | `Context` 하나로 통합                |
| **JSON 처리** | 수동 (`json.Encoder`, 헤더 설정) | 자동 (`c.JSON()`)                 |
| **파라미터 바인딩** | 수동 파싱 + 검증 | `ShouldBindJSON()` 자동           |
| **라우팅** | `.HandleFunc().Methods()` | `.GET()`, `.POST()`             |
| **에러 처리** | `WriteHeader()` | `c.JSON()`, `AbortWithStatus()` |
| **미들웨어** | 수동 구현 필요 | 로깅/복구 기본 제공                     |
| **성능** | 보통 | 비교적 빠름                          |

### 코드 라인 수 비교 (예: POST 핸들러)

**Gorilla Mux: ~15줄**
```go
func PostHandler(w http.ResponseWriter, r *http.Request) {
    var student Student
    err := json.NewDecoder(r.Body).Decode(&student)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    lastId++
    student.Id = lastId
    students[lastId] = student
    w.WriteHeader(http.StatusOK)
}
```

**Gin: ~10줄**
```go
func PostStudentHandler(c *gin.Context) {
    var student Student
    if err := c.ShouldBindJSON(&student); err != nil {
        c.JSON(http.StatusBadRequest, err.Error())
        return
    }
    lastId++
    student.Id = lastId
    students[lastId] = student
    c.String(http.StatusCreated, "Success to add id: %d", lastId)
}
```

### 결론
- **Gorilla Mux**: 표준 라이브러리에 가까움
- **Gin**: 간단
- **추상화 정도**: gin이 간결

---

## 테스트 실행
```bash
# 모든 테스트 실행
go test ./... -v

# 커버리지 확인
go test ./... -cover

# 그냥 통합 테스트
go test
```

## 서버 실행
```bash
go run .
```

