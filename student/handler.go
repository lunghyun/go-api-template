package student

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

// NewHandler 생성자
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetStudentsHandler 학생 목록 조회 핸들러
func (h *Handler) GetStudentsHandler(c *gin.Context) {
	list := h.service.GetStudents()
	c.JSON(http.StatusOK, list) // json 포멧 변환
}

// GetStudentHandler 특정 학생 조회 핸들러
func (h *Handler) GetStudentHandler(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	student, err := h.service.GetStudent(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

// PostStudentHandler 학생 추가 핸들러
func (h *Handler) PostStudentHandler(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	created, err := h.service.CreateStudent(student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// DeleteStudentHandler 학생 삭제 핸들러
func (h *Handler) DeleteStudentHandler(c *gin.Context) {
	idstr := c.Params.ByName("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err = h.service.DeleteStudent(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.String(http.StatusOK, "Student deleted id: %d", id)
}
