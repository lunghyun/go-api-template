package student

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes student 라우팅 등록
func RegisterRoutes(g *gin.Engine) {
	//r.HandleFunc("/students", GetListHandler).Methods("GET")
	g.GET("/students", GetStudentsHandler)
	//r.HandleFunc("/students/{id:[0-9]+}", GetHandler).Methods("GET")
	g.GET("/students/:id", GetStudentHandler)
	//r.HandleFunc("/students", PostHandler).Methods("POST")
	g.POST("/students", PostStudentHandler)
	//r.HandleFunc("/students/{id:[0-9]+}", DeleteHandler).Methods("DELETE")
	g.DELETE("/students/:id", DeleteStudentHandler)
}
