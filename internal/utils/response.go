package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

// ApiResponseSuccess trả về response thành công cho client
func ApiResponseSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  true,
		"message": message,
		"data":    data,
	})
}

// ApiResponseError trả về response lỗi cho client
func ApiResponseError(c *gin.Context, err error) {
	statusCode := MapErrorToStatusCode(err)
	c.JSON(statusCode, gin.H{
		"status":  false,
		"message": err.Error(),
	})
}

// ValidationErrorResponse struct cho lỗi validation
type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ErrorResponse struct cho response lỗi
type ErrorResponse struct {
	StatusCode int                       `json:"status_code"`
	Message    string                    `json:"message"`
	Errors     []ValidationErrorResponse `json:"errors,omitempty"`
}

// NewErrorResponse tạo ErrorResponse mới
func NewErrorResponse(statusCode int, message string, errors []ValidationErrorResponse) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		Errors:     errors,
	}
}

// HandleValidationError xử lý lỗi validation
func HandleValidationError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if ok := errors.As(err, &ve); ok {
		out := make([]ValidationErrorResponse, len(ve))
		for i, fe := range ve {
			out[i] = ValidationErrorResponse{fe.Field(), removeDataTag(fe.Tag())}
		}
		errResponse := NewErrorResponse(http.StatusBadRequest, "Validation Error", out)
		c.JSON(errResponse.StatusCode, errResponse)
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi server nội bộ"})
}

func removeDataTag(tag string) string {
	return strings.Replace(tag, "Data", "", -1)
}
