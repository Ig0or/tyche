package models

import (
	"github.com/google/uuid"
	"time"
)

type AccountModel struct {
	ID             int
	AccountId      uuid.UUID
	Email          string
	Cpf            string
	HashedPassword string
	CreatedAt      time.Time
}

func (model *AccountModel) GetArgumentsToInsert() []interface{} {
	arguments := []interface{}{model.AccountId, model.Email, model.Cpf, model.HashedPassword, model.CreatedAt}

	return arguments
}
