package entities

import (
	"github.com/Ig0or/tyche/src/domain/enums"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type TransactionEntity struct {
	ID          int
	AccountId   uuid.UUID
	Operation   enums.Operation
	Type        enums.Type
	Amount      decimal.Decimal
	ToAccountId uuid.UUID
	CreatedAt   time.Time
}
