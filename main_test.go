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
	assertion := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	engin := MakeWebHandler()
	engin.ServeHTTP(res, req)

	assertion.Equal(http.StatusOK, res.Code)
	var list student.Students
	err := json.NewDecoder(res.Body).Decode(&list)
	assertion.Nil(err)
	assertion.Equal(2, len(list))
	assertion.Equal("aaa", list[0].Name)
	assertion.Equal("bbb", list[1].Name)
}

func TestGetStudentHandler(t *testing.T) {
	assertion := assert.New(t)

	var stu student.Student
	engin := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)

	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&stu)
	assertion.Nil(err)
	assertion.Equal("aaa", stu.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/2", nil)

	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&stu)
	assertion.Nil(err)
	assertion.Equal("bbb", stu.Name)
}

func TestPostStudentHandler(t *testing.T) {
	assertion := assert.New(t)

	var content student.Student
	engin := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"Id":3,"name":"ccc","Age":20,"Score":78}`))

	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)

	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&content)
	assertion.Nil(err)
	assertion.Equal("ccc", content.Name)
	assertion.Equal(78, content.Score)
}

func TestDeleteStudentHandler(t *testing.T) {
	assertion := assert.New(t)
	engin := MakeWebHandler()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/1", nil)
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students", nil)
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusOK, res.Code)

	var students student.Students
	err := json.NewDecoder(res.Body).Decode(&students)
	assertion.Nil(err)
	assertion.Equal(1, len(students))
	assertion.Equal("bbb", students[0].Name)
}

func TestPutStudentHandler(t *testing.T) {
	assertion := assert.New(t)
	engin := MakeWebHandler()

	var content student.Student
	res := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/students/1",
		strings.NewReader(`{"Id":1,"name":"ddd","Age":100,"Score":78}`))

	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/1", nil)

	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&content)
	assertion.Nil(err)
	assertion.Equal("ddd", content.Name)
	assertion.Equal(100, content.Age)
	assertion.Equal(78, content.Score)
}
