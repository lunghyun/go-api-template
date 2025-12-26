package main

import (
	"github.com/gin-gonic/gin"
	student2 "github.com/lunghyun/go-api-template/internal/domain/student"
)

// MakeWebHandler 웹 핸들러 초기화
func MakeWebHandler() *gin.Engine {
	engin := gin.Default()

	// di
	studentRepo := student2.NewMemRepository()
	studentService := student2.NewService(studentRepo)
	studentHandler := student2.NewHandler(studentService)

	// 각 도메인 라우팅 등록
	studentHandler.RegisterRoutes(engin)

	return engin
}
