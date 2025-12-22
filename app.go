package main

import (
	"net/http"

	"mustHaveGoRest/student"

	"github.com/gorilla/mux"
)

// MakeWebHandler 웹 핸들러 초기화
func MakeWebHandler() http.Handler {
	muxes := mux.NewRouter()

	// 각 도메인 라우팅 등록
	student.RegisterRoutes(muxes)

	return muxes
}
