package student

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetStudentsHandler 학생 목록 조회 핸들러
func GetStudentsHandler(c *gin.Context) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)             // Id 기준 정렬
	c.JSON(http.StatusOK, list) // json 포멧 변환
}

// GetStudentHandler 특정 학생 조회 핸들러
func GetStudentHandler(c *gin.Context) {
	idstr := c.Params.ByName("id")
	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	student, ok := students[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student)
}

// PostStudentHandler 학생 추가 핸들러
func PostStudentHandler(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	lastId++
	student.Id = lastId
	students[lastId] = student
	c.String(http.StatusCreated, "Success to add id: %d", lastId)
}

// DeleteStudentHandler 학생 삭제 핸들러
func DeleteStudentHandler(c *gin.Context) {
	idstr := c.Params.ByName("id")
	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	delete(students, id)
	c.String(http.StatusOK, "Success to delete id: %d", id)
}
