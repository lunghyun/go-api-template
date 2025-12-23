package student

import (
	"mustHaveGoRest/errors"
	"sort"
)

type Service struct {
	repository *Repository
}

// NewService 생성자
func NewService(r *Repository) *Service {
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
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// CreateStudent 학생 생성
func (s *Service) CreateStudent(student Student) (*Student, error) {
	if student.Name == "" {
		return nil, errors.ErrInvalidName
	}
	if student.Age < 0 || student.Age > 150 {
		return nil, errors.ErrInvalidAge
	}
	created := s.repository.Create(student)
	return &created, nil
}

// DeleteStudent 학생 삭제
func (s *Service) DeleteStudent(id int) error {
	if err := s.repository.DeleteById(id); err != nil {
		return err
	}
	return nil
}
