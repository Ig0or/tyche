package models

import (
	"github.com/Ig0or/tyche/src/domain/enums"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type TransactionModel struct {
	ID                   int
	AccountId            uuid.UUID
	Operation            enums.Operation
	Type                 enums.Type
	Amount               decimal.Decimal
	DestinationAccountId uuid.UUID
	CreatedAt            time.Time
}

func (model *TransactionModel) GetArgumentsToInsert() []interface{} {
	var destinationAccountId interface{} = model.DestinationAccountId

	if model.DestinationAccountId == uuid.Nil {
		destinationAccountId = nil
	}

	arguments := []interface{}{model.AccountId, model.Operation, model.Type, model.Amount, destinationAccountId, model.CreatedAt}

	return arguments
}
