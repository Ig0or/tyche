package custom_errors

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func FormatRequestValidationError(
	validationError error,
) string {
	var invalidFields []string

	for _, err := range validationError.(validator.ValidationErrors) {
		invalidField := fmt.Sprintf("Invalid value for field: %s - It must be a valid %s type.", err.Field(), err.Type())

		invalidFields = append(invalidFields, invalidField)
	}

	message := fmt.Sprintf(
		"Fail to create account because some fields are invalid: %s",
		strings.Join(invalidFields, " | "),
	)

	return message
}
