package custom_errors

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func FormatRequestValidationError(
	validationError error,
) []string {
	var invalidFields []string

	for _, err := range validationError.(validator.ValidationErrors) {
		invalidField := fmt.Sprintf("Invalid value for field: %s - It must be a valid %s type", err.Field(), err.Type())

		invalidFields = append(invalidFields, invalidField)
	}

	return invalidFields
}
