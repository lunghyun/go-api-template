package student

import "sync"

type Repository struct {
	students map[int]Student
	lastId   int
	mutex    sync.RWMutex
}

func NewRepository() *Repository {
	repo := &Repository{
		students: make(map[int]Student),
		lastId:   2,
	}
	// 초기 데이터
	repo.students[1] = Student{1, "aaa", 16, 87}
	repo.students[2] = Student{2, "bbb", 18, 98}
	return repo
}

// GetAll 전체 조회
func (repo *Repository) GetAll() Students {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	list := make(Students, 0, len(repo.students))
	for _, v := range repo.students {
		list = append(list, v)
	}
	return list
}

// GetById Id로 조회
func (repo *Repository) GetById(id int) (Student, bool) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	s, ok := repo.students[id]
	return s, ok
}

// Create 학생 추가
func (repo *Repository) Create(s Student) Student {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.lastId++
	s.Id = repo.lastId
	repo.students[repo.lastId] = s
	return s
}

// Delete 학생 삭제
func (repo *Repository) Delete(id int) bool {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	if _, ok := repo.students[id]; !ok {
		return false
	}
	delete(repo.students, id)
	return true
}
