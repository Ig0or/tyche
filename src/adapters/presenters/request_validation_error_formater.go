package presenters

import (
	"fmt"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/go-playground/validator/v10"
)

func FormatRequestValidationError(
	validationError error,
) *custom_errors.BaseCustomError {
	originalError := getFieldErrors(validationError)

	customError := custom_errors.NewBadRequestError("Validation error. Some fields are invalid", originalError)

	return customError
}

func getFieldErrors(validationError error) []string {
	var invalidFields []string

	for _, err := range validationError.(validator.ValidationErrors) {
		invalidField := fmt.Sprintf("Invalid value for field: %s - It must be a valid %s type", err.Field(), err.Type())

		invalidFields = append(invalidFields, invalidField)
	}

	return invalidFields
}
