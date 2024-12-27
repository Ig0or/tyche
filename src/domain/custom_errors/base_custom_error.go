package custom_errors

import (
	"net/http"
)

type BaseCustomError struct {
	Message       string      `json:"message"`
	StatusCode    int         `json:"status_code"`
	OriginalError interface{} `json:"original_error"`
}

func NewBadRequestError(message string, originalError interface{}) *BaseCustomError {
	customError := &BaseCustomError{Message: message, StatusCode: http.StatusBadRequest, OriginalError: originalError}

	return customError
}

func NewInternalServerError(
	message string, originalError error,
) *BaseCustomError {
	customError := &BaseCustomError{Message: message, StatusCode: http.StatusInternalServerError, OriginalError: originalError}

	return customError
}
