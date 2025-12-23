package student

import "fmt"

const (
	// MinAge is the minimum valid age for a student
	MinAge = 0
	// MaxAge is the maximum valid age for a student
	MaxAge = 150
	// MinScore is the minimum valid score for a student
	MinScore = 0
	// MaxScore is the maximum valid score for a student
	MaxScore = 100
)

// Student 학생 정보
type Student struct {
	ID    int    `json:"id,omitempty"`
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
	return s[i].ID < s[j].ID
}

// Validate 도메인 규칙
func (s Student) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("invalid name: name is required")
	}
	if s.Age < MinAge || s.Age > MaxAge {
		return fmt.Errorf("invalid age: !(%d <= %d <= %d)", MinAge, s.Age, MaxAge)
	}
	if s.Score < MinScore || s.Score > MaxScore {
		return fmt.Errorf("invalid score: !(%d <= %d <= %d)", MinScore, s.Score, MaxScore)
	}
	return nil
}
