package main

import (
	"mustHaveGoRest/student"

	"github.com/gin-gonic/gin"
)

// MakeWebHandler 웹 핸들러 초기화
func MakeWebHandler() *gin.Engine {
	engin := gin.Default()

	// di
	repo := student.NewRepository()
	service := student.NewService(repo)
	handler := student.NewHandler(service)

	// 각 도메인 라우팅 등록
	handler.RegisterRoutes(engin)

	return engin
}
