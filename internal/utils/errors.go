package utils

import (
	"errors"
	"net/http"
)

// Định nghĩa các loại lỗi
var (
	ErrConflict            = NewAppError(http.StatusConflict, "Dữ liệu đã tồn tại")
	ErrNotFound            = NewAppError(http.StatusNotFound, "Không tìm thấy dữ liệu")
	ErrInternalServerError = NewAppError(http.StatusInternalServerError, "Lỗi máy chủ nội bộ")
	ErrUnauthorized        = NewAppError(http.StatusUnauthorized, "Không được phép truy cập")
)

type AppError struct {
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(statusCode int, message string) error {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
	}
}

// MapErrorToStatusCode ánh xạ lỗi sang mã HTTP status code
func MapErrorToStatusCode(err error) int {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.StatusCode
	}
	return http.StatusInternalServerError
}
