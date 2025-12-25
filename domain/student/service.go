package student

import (
	"sort"
)

type Service struct {
	repository Repository
}

// NewService 생성자
func NewService(r Repository) *Service {
	return &Service{repository: r}
}

// GetStudents 전체 조회 (+ 정렬)
func (s *Service) GetStudents() Students {
	list := s.repository.FindAll()
	sort.Sort(list)
	return list
}

// GetStudent id 조회
func (s *Service) GetStudent(id int) (*Student, error) {
	student, err := s.repository.FindById(id)
	if err != nil { // id에 해당하는 키가 없음
		return nil, err
	}
	return &student, nil
}

// CreateStudent 학생 생성
func (s *Service) CreateStudent(student Student) (*Student, error) {
	created := s.repository.Save(student)
	return &created, nil
}

// UpdateStudent 학생 수정
func (s *Service) UpdateStudent(id int, student Student) (*Student, error) {
	_, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	student.ID = id
	updated := s.repository.Save(student)
	return &updated, nil
}

// DeleteStudent 학생 삭제
func (s *Service) DeleteStudent(id int) error {
	if err := s.repository.DeleteById(id); err != nil { // id에 해당하는 키가 없음
		return err
	}
	return nil
}
