package models

import (
	"github.com/Ig0or/tyche/src/domain/enums"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type TransactionModel struct {
	ID          int
	AccountId   uuid.UUID
	Operation   enums.Operation
	Type        enums.Type
	Amount      decimal.Decimal
	ToAccountId uuid.UUID
	CreatedAt   time.Time
}

func (model *TransactionModel) GetArgumentsToInsert() []interface{} {
	var toAccountId interface{} = model.ToAccountId

	if model.ToAccountId == uuid.Nil {
		toAccountId = nil
	}

	arguments := []interface{}{model.AccountId, model.Operation, model.Type, model.Amount, toAccountId, model.CreatedAt}

	return arguments
}
