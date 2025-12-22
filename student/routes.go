package student

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes student 라우팅 등록
func (h *Handler) RegisterRoutes(g *gin.Engine) {
	//r.HandleFunc("/students", GetListHandler).Methods("GET")
	g.GET("/students", h.GetStudentsHandler)
	//r.HandleFunc("/students/{id:[0-9]+}", GetHandler).Methods("GET")
	g.GET("/students/:id", h.GetStudentHandler)
	//r.HandleFunc("/students", PostHandler).Methods("POST")
	g.POST("/students", h.PostStudentHandler)
	//r.HandleFunc("/students/{id:[0-9]+}", DeleteHandler).Methods("DELETE")
	g.DELETE("/students/:id", h.DeleteStudentHandler)
}
