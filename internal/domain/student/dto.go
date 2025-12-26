package student

// TODO : validate를 위한 dto 생성

// PostStudentRequest 생성 request
type PostStudentRequest struct {
	Name  string `json:"name" binding:"required,min=1,max=10"`
	Age   int    `json:"age" binding:"required,min=1,max=100"`
	Score int    `json:"score" binding:"required,min=0,max=100"`
}

func (r PostStudentRequest) ToStudent() Student {
	return Student{
		ID:    0,
		Name:  r.Name,
		Age:   r.Age,
		Score: r.Score,
	}
}

// UpdateStudentRequest 수정 request
type UpdateStudentRequest struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required,min=1,max=100"`
	Score int    `json:"score" binding:"required,min=0,max=100"`
}

func (r UpdateStudentRequest) ToStudent() Student {
	return Student{
		ID:    0,
		Name:  r.Name,
		Age:   r.Age,
		Score: r.Score,
	}
}
