package student

import "fmt"

// Student 학생 정보
type Student struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score,omitempty"`
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
	return s[i].Id < s[j].Id
}

// Validate 도메인 규칙
func (s Student) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("invalid name: name is required")
	}
	if s.Age < 0 || s.Age > 150 {
		return fmt.Errorf("invalid age: 0 !< %d !< 150", s.Age)
	}
	if s.Score < 0 || s.Score > 100 {
		return fmt.Errorf("invalid score: 0 !< %d !< 100", s.Score)
	}
	return nil
}
