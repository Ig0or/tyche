package requests

import (
	"github.com/google/uuid"
)

type CreateTransactionRequest struct {
	AccountId            string
	Operation            string    `json:"operation" binding:"required,oneof=deposit withdraw transfer"`
	Amount               float64   `json:"amount" binding:"required,gte=0"`
	DestinationAccountId uuid.UUID `json:"destination_account_id" binding:"uuid"`
}
