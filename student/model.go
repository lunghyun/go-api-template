package student

// Student 학생 정보
type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

// Students 학생 목록 (정렬을 위한 커스텀 타입)
type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}
