package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lunghyun/go-api-template/internal/domain/student"
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

func TestPostStudentHandler_ValidationFail(t *testing.T) {
	assertion := assert.New(t)
	engin := MakeWebHandler()

	// 1. Name 누락 (required)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"age":20,"score":90}`))
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusBadRequest, res.Code)

	var errorResp map[string]string
	err := json.NewDecoder(res.Body).Decode(&errorResp)
	assertion.Nil(err)
	assertion.Contains(errorResp["error"], "Name")
	t.Logf("[Name 누락] 에러: %s", errorResp["error"])

	// 2. Age 범위 초과 (max=100)
	res = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"name":"test","age":101,"score":90}`))
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusBadRequest, res.Code)

	errorResp = map[string]string{}
	err = json.NewDecoder(res.Body).Decode(&errorResp)
	assertion.Nil(err)
	assertion.Contains(errorResp["error"], "Age")
	t.Logf("[Age 범위 초과] 에러: %s", errorResp["error"])

	// 3. Age 최소값 미만 (min=1)
	res = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"name":"test","age":0,"score":90}`))
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusBadRequest, res.Code)

	errorResp = map[string]string{}
	err = json.NewDecoder(res.Body).Decode(&errorResp)
	assertion.Nil(err)
	assertion.Contains(errorResp["error"], "Age")
	t.Logf("[Age 최소값 미만] 에러: %s", errorResp["error"])

	// 4. Score 범위 초과 (max=100)
	res = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"name":"test","age":20,"score":101}`))
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusBadRequest, res.Code)

	errorResp = map[string]string{}
	err = json.NewDecoder(res.Body).Decode(&errorResp)
	assertion.Nil(err)
	assertion.Contains(errorResp["error"], "Score")
	t.Logf("[Score 범위 초과] 에러: %s", errorResp["error"])

	// 5. Score 음수 (min=0)
	res = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"name":"test","age":20,"score":-1}`))
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusBadRequest, res.Code)

	errorResp = map[string]string{}
	err = json.NewDecoder(res.Body).Decode(&errorResp)
	assertion.Nil(err)
	assertion.Contains(errorResp["error"], "Score")
	t.Logf("[Score 음수] 에러: %s", errorResp["error"])
}

func TestPutStudentHandler_ValidationFail(t *testing.T) {
	assertion := assert.New(t)
	engin := MakeWebHandler()

	// 1. Name 누락 (required)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/students/1",
		strings.NewReader(`{"age":20,"score":90}`))
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusBadRequest, res.Code)

	var errorResp map[string]string
	err := json.NewDecoder(res.Body).Decode(&errorResp)
	assertion.Nil(err)
	assertion.Contains(errorResp["error"], "Name")
	t.Logf("[PUT Name 누락] 에러: %s", errorResp["error"])

	// 2. Age 범위 초과 (max=100)
	res = httptest.NewRecorder()
	req = httptest.NewRequest("PUT", "/students/1",
		strings.NewReader(`{"name":"test","age":101,"score":90}`))
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusBadRequest, res.Code)

	errorResp = map[string]string{}
	err = json.NewDecoder(res.Body).Decode(&errorResp)
	assertion.Nil(err)
	assertion.Contains(errorResp["error"], "Age")
	t.Logf("[PUT Age 범위 초과] 에러: %s", errorResp["error"])

	// 3. Age 음수 (min=1)
	res = httptest.NewRecorder()
	req = httptest.NewRequest("PUT", "/students/1",
		strings.NewReader(`{"name":"test","age":-1,"score":90}`))
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusBadRequest, res.Code)

	errorResp = map[string]string{}
	err = json.NewDecoder(res.Body).Decode(&errorResp)
	assertion.Nil(err)
	assertion.Contains(errorResp["error"], "Age")
	t.Logf("[PUT Age 음수] 에러: %s", errorResp["error"])

	// 4. Score 범위 초과 (max=100)
	res = httptest.NewRecorder()
	req = httptest.NewRequest("PUT", "/students/1",
		strings.NewReader(`{"name":"test","age":99,"score":101}`))
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusBadRequest, res.Code)

	errorResp = map[string]string{}
	err = json.NewDecoder(res.Body).Decode(&errorResp)
	assertion.Nil(err)
	assertion.Contains(errorResp["error"], "Score")
	t.Logf("[PUT Score 범위 초과] 에러: %s", errorResp["error"])

	// 5. Score 음수 (min=0)
	res = httptest.NewRecorder()
	req = httptest.NewRequest("PUT", "/students/1",
		strings.NewReader(`{"name":"test","age":99,"score":-1}`))
	engin.ServeHTTP(res, req)
	assertion.Equal(http.StatusBadRequest, res.Code)

	errorResp = map[string]string{}
	err = json.NewDecoder(res.Body).Decode(&errorResp)
	assertion.Nil(err)
	assertion.Contains(errorResp["error"], "Score")
	t.Logf("[PUT Score 음수] 에러: %s", errorResp["error"])
}
