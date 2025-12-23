package student

import (
	"fmt"
	"sync"
)

type Repository interface {
	FindAll() Students
	FindById(id int) (Student, error)
	Save(s Student) Student
	DeleteById(id int) error
}

type memoryRepository struct {
	students map[int]Student
	lastId   int
	mutex    sync.RWMutex
}

func NewMemRepository() Repository {
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

	s, exists := r.students[id] // id에 해당하는 키가 없음
	if !exists {
		return s, fmt.Errorf("student id(%d) not found", id)
	}
	return s, nil
}

// Save 학생 생성 및 수정
func (r *memoryRepository) Save(s Student) Student {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Create id가 비어있을 경우 생성(id == 0) -> id 값 생성하여 lastId로 수정
	if _, exists := r.students[s.Id]; !exists {
		r.lastId++
		s.Id = r.lastId
	}

	// Update id에 해당하는 키가 있을 경우 수정
	r.students[s.Id] = s
	return s
}

// DeleteById id를 통한 학생 삭제
func (r *memoryRepository) DeleteById(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.students[id]; !exists { // id에 해당하는 키가 없음
		return fmt.Errorf("student id(%d) not found", id)
	}
	delete(r.students, id)
	return nil
}
