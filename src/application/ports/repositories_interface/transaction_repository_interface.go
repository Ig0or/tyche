package repositories_interface

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/google/uuid"
)

type TransactionRepositoryInterface interface {
	CreateTransaction(transaction *models.TransactionModel) *custom_errors.BaseCustomError
	GetTransactionsByAccountID(accountId uuid.UUID) (*[]*models.TransactionModel, *custom_errors.BaseCustomError)
}
