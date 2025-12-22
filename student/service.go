package student

import (
	"mustHaveGoRest/errors"
	"sort"
)

type Service struct {
	repo *Repository
}

// NewService 생성자
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetStudents 전체 조회 (+ 정렬)
func (s *Service) GetStudents() Students {
	list := s.repo.GetAll()
	sort.Sort(list)
	return list
}

// GetStudent id 조회
func (s *Service) GetStudent(id int) (*Student, error) {
	student, ok := s.repo.GetById(id)
	if !ok {
		return nil, errors.ErrNotFound
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
	created := s.repo.Create(student)
	return &created, nil
}

// DeleteStudent 학생 삭제
func (s *Service) DeleteStudent(id int) error {
	if !s.repo.Delete(id) {
		return errors.NotFound("student not found")
	}
	return nil
}
