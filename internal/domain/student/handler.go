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
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	var req PostStudentRequest

	// 1. req에 Json 바인딩 + validate in binding tag
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2, DTO -> Student 변환
	student := req.ToStudent()

	// 3. 서비스 호출
	created, err := h.service.CreateStudent(student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// PutStudentHandler 학생 수정 핸들러
func (h *Handler) PutStudentHandler(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := req.ToStudent()
	updated, err := h.service.UpdateStudent(id, student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DeleteStudentHandler 학생 삭제 핸들러
func (h *Handler) DeleteStudentHandler(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = h.service.DeleteStudent(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Student deleted id: %d", id)
}
