package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student // 학생 목록
var lastId int

func MakeWebHandler() http.Handler {
	muxes := mux.NewRouter()
	muxes.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	muxes.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")
	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2

	return muxes
}

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

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0) // 1. 학생 목록을 id로 정렬
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list) // list를 json으로 변경
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // id(파라미터)를 가져옴
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 에러코드 처리
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(student)
}

func main() {
	http.ListenAndServe(":8080", MakeWebHandler())
}
