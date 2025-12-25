package student

// Student 학생 정보
type Student struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

// Students 학생 목록 (정렬을 위한 커스텀 타입)
type Students []Student

// Len Swap Less -> []Student(Students) 에 대한 Sort 인터페이스의 메서드
func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}
