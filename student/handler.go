package student

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetStudentsHandler 학생 목록 조회 핸들러
// [SRP 위반] HTTP 처리 + 비즈니스 로직(정렬) + 데이터 접근을 한 함수에서 처리
func GetStudentsHandler(c *gin.Context) {
	list := make(Students, 0)
	// [DIP 위반] 전역 변수 students에 직접 의존 (추상화 없음)
	// [동시성 문제] 다른 goroutine에서 students 수정 시 race condition 가능
	for _, student := range students {
		list = append(list, student)
	}
	// [SRP 위반] 비즈니스 로직(정렬)이 핸들러에 포함됨 → Service 계층으로 분리 필요
	sort.Sort(list)             // Id 기준 정렬
	c.JSON(http.StatusOK, list) // json 포멧 변환
}

// GetStudentHandler 특정 학생 조회 핸들러
// [SRP 위반] HTTP 처리 + 데이터 접근을 한 함수에서 처리
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
	// [DIP 위반] 전역 변수 students에 직접 의존 (추상화 없음)
	// [동시성 문제] 읽는 동안 다른 goroutine이 수정하면 데이터 경합 발생 가능
	student, ok := students[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student)
}

// PostStudentHandler 학생 추가 핸들러
// [SRP 위반] HTTP 처리 + 비즈니스 로직(ID 생성) + 데이터 접근을 한 함수에서 처리
func PostStudentHandler(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// [DIP 위반] 전역 변수 lastId, students에 직접 의존
	// [동시성 문제] lastId++ 연산이 atomic하지 않음 → 동시 요청 시 같은 ID 생성 가능
	lastId++
	// [SRP 위반] ID 생성 로직이 핸들러에 포함됨 → Repository/Service로 분리 필요
	student.Id = lastId
	// [DIP 위반] 구체적인 구현(map)에 직접 의존 (추상화 없음)
	// [OCP 위반] DB로 변경 시 이 코드를 모두 수정해야 함
	students[lastId] = student
	c.String(http.StatusCreated, "Success to add id: %d", lastId)
}

// DeleteStudentHandler 학생 삭제 핸들러
// [SRP 위반] HTTP 처리 + 데이터 접근을 한 함수에서 처리
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
	// [DIP 위반] 전역 변수 students에 직접 의존 (추상화 없음)
	// [동시성 문제] 삭제하는 동안 다른 goroutine이 접근하면 데이터 경합 발생 가능
	// [OCP 위반] DB로 변경 시 이 코드를 모두 수정해야 함
	delete(students, id)
	c.String(http.StatusOK, "Success to delete id: %d", id)
}
