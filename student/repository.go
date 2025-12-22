package student

// [SRP 위반] 전역 변수로 데이터 관리 → 캡슐화 없음
// [DIP 위반] 구체적인 구현(map)만 존재, 추상화(Repository Interface) 없음
// [동시성 문제] 전역 변수에 여러 goroutine이 동시 접근 시 race condition 발생
//
// 개선 방안:
// 1. Repository Interface 정의로 추상화 (DIP)
// 2. InMemoryRepository 구조체로 캡슐화 (SRP)
// 3. sync.RWMutex로 동시성 보장 (동시성 안전)
// 4. 테스트 시 Mock Repository 주입 가능 (DIP)
var students map[int]Student // 학생 목록
var lastId int               // 마지막 학생 ID

// init 패키지 초기화 시 자동 실행
func init() {
	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2
}
