package dtos

import (
	"github.com/Ig0or/tyche/src/domain/enums"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type TransactionDto struct {
	ID          int
	AccountId   uuid.UUID
	Operation   enums.Operation
	Type        enums.Type
	Amount      decimal.Decimal
	ToAccountId uuid.UUID
	CreatedAt   time.Time
}
