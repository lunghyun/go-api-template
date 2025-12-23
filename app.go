package main

import (
	"mustHaveGoRest/student"

	"github.com/gin-gonic/gin"
)

// MakeWebHandler 웹 핸들러 초기화
func MakeWebHandler() *gin.Engine {
	engin := gin.Default()

	// di
	studentRepo := student.NewRepository()
	studentService := student.NewService(studentRepo)
	studentHandler := student.NewHandler(studentService)

	// 각 도메인 라우팅 등록
	studentHandler.RegisterRoutes(engin)

	return engin
}
