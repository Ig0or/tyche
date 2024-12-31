package dtos

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type AccountDto struct {
	ID           int
	AccountId    uuid.UUID
	Email        string
	Cpf          string
	Password     string
	Balance      decimal.Decimal
	Transactions []*TransactionDto
	CreatedAt    time.Time
}
