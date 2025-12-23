package student

import (
	"fmt"
	"sync"
)

type Repository interface {
	FindAll() Students
	FindById(id int) (Student, error)
	Create(s Student) Student
	DeleteById(id int) error
}

type memoryRepository struct {
	students map[int]Student
	lastId   int
	mutex    sync.RWMutex
}

func NewRepository() Repository {
	repo := &memoryRepository{
		students: make(map[int]Student),
		lastId:   2,
	}
	// 초기 데이터
	repo.students[1] = Student{1, "aaa", 16, 87}
	repo.students[2] = Student{2, "bbb", 18, 98}
	return repo
}

// FindAll 전체 조회
func (r *memoryRepository) FindAll() Students {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	list := make(Students, 0, len(r.students))
	for _, v := range r.students {
		list = append(list, v)
	}
	return list
}

// FindById Id로 조회
func (r *memoryRepository) FindById(id int) (Student, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	s, ok := r.students[id]
	if !ok {
		return s, fmt.Errorf("student id(%d) not found", id)
	}
	return s, nil
}

// Create 학생 추가
func (r *memoryRepository) Create(s Student) Student {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.lastId++
	s.Id = r.lastId
	r.students[r.lastId] = s
	return s
}

// DeleteById id를 통한 학생 삭제
func (r *memoryRepository) DeleteById(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.students[id]; !ok {
		return fmt.Errorf("student id(%d) not found", id)
	}
	delete(r.students, id)
	return nil
}
