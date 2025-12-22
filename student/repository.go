package student

var students map[int]Student // 학생 목록
var lastId int               // 마지막 학생 ID

// init 패키지 초기화 시 자동 실행
func init() {
	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2
}
