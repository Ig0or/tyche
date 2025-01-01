package custom_errors

import (
	"net/http"
)

type BaseCustomError struct {
	Message       string
	OriginalError error
	StatusCode    int
}

func NewBadRequestError(message string, err error) *BaseCustomError {
	customError := &BaseCustomError{Message: message, OriginalError: err, StatusCode: http.StatusBadRequest}

	return customError
}

func NewInternalServerError(message string, err error) *BaseCustomError {
	customError := &BaseCustomError{Message: message, OriginalError: err, StatusCode: http.StatusInternalServerError}

	return customError
}

func NewUnauthorizedError(message string, err error) *BaseCustomError {
	customError := &BaseCustomError{Message: message, OriginalError: err, StatusCode: http.StatusUnauthorized}

	return customError
}
