package responses

import (
	"github.com/google/uuid"
)

type CreateAccountResponse struct {
	AccountId uuid.UUID `json:"account_id"`
	Balance   float64   `json:"balance"`
}
