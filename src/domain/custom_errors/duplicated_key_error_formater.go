package custom_errors

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"strings"
)

func FormatDuplicatedKeyError(
	error *pgconn.PgError,
) string {
	if error.Code == "23505" {
		duplicatedKey := getDuplicatedKey(error)

		message := fmt.Sprintf("Fail to create account because this %s already exists.", duplicatedKey)

		return message
	}

	return ""
}

func getDuplicatedKey(error *pgconn.PgError) string {
	keys := []string{"cpf", "email"}

	for _, key := range keys {
		if strings.Contains(error.Detail, key) {
			return key
		}
	}

	return ""
}
