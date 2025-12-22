package errors

// AppError HTTP 상태코드를 포함한 에러
type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

// 헬퍼 함수
func NotFound(msg string) error {
	return &AppError{Code: 404, Message: msg}
}

func BadRequest(msg string) error {
	return &AppError{Code: 400, Message: msg}
}

func InternalError(msg string) error {
	return &AppError{Code: 500, Message: msg}
}

// 공통 에러 변수 (선택사항)
var (
	ErrNotFound     = &AppError{Code: 404, Message: "not found"}
	ErrBadRequest   = &AppError{Code: 400, Message: "bad request"}
	ErrUnauthorized = &AppError{Code: 401, Message: "unauthorized"}
	ErrInvalidName  = &AppError{Code: 400, Message: "invalid name"}
	ErrInvalidAge   = &AppError{Code: 400, Message: "invalid age"}
)
