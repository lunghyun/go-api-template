package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestJsonHandler(t *testing.T) {
	asserts := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	asserts.Equal(http.StatusOK, res.Code)
	var list Students
	err := json.NewDecoder(res.Body).Decode(&list)
	asserts.Nil(err)
	asserts.Equal(2, len(list))
	asserts.Equal("aaa", list[0].Name)
	asserts.Equal("bbb", list[1].Name)
}

func TestJsonHandler2(t *testing.T) {
	asserts := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)

	mux.ServeHTTP(res, req)
	asserts.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	asserts.Nil(err)
	asserts.Equal("aaa", student.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/2", nil)

	mux.ServeHTTP(res, req)
	asserts.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&student)
	asserts.Nil(err)
	asserts.Equal("bbb", student.Name)
}
