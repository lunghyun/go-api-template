package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"mustHaveGoRest/student"

	"github.com/stretchr/testify/assert"
)

func TestGetStudentsHandler(t *testing.T) {
	asserts := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	asserts.Equal(http.StatusOK, res.Code)
	var list student.Students
	err := json.NewDecoder(res.Body).Decode(&list)
	asserts.Nil(err)
	asserts.Equal(2, len(list))
	asserts.Equal("aaa", list[0].Name)
	asserts.Equal("bbb", list[1].Name)
}

func TestGetStudentHandler(t *testing.T) {
	asserts := assert.New(t)

	var stu student.Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)

	mux.ServeHTTP(res, req)
	asserts.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&stu)
	asserts.Nil(err)
	asserts.Equal("aaa", stu.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/2", nil)

	mux.ServeHTTP(res, req)
	asserts.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&stu)
	asserts.Nil(err)
	asserts.Equal("bbb", stu.Name)
}

func TestPostStudentHandler(t *testing.T) {
	asserts := assert.New(t)

	var student student.Student
	muxes := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"Id":3,"name":"ccc","Age":20,"Score":78}`))

	muxes.ServeHTTP(res, req)
	asserts.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)

	muxes.ServeHTTP(res, req)
	asserts.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	asserts.Nil(err)
	asserts.Equal("ccc", student.Name)
	asserts.Equal(78, student.Score)
}

func TestDeleteStudentHandler(t *testing.T) {
	asserts := assert.New(t)
	muxes := MakeWebHandler()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/2", nil)
	muxes.ServeHTTP(res, req)
	asserts.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students", nil)
	muxes.ServeHTTP(res, req)
	asserts.Equal(http.StatusOK, res.Code)

	var students student.Students
	err := json.NewDecoder(res.Body).Decode(&students)
	asserts.Nil(err)
	asserts.Equal(2, len(students))
	asserts.Equal("ccc", students[1].Name)
}
