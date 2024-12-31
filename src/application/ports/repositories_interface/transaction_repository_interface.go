package repositories_interface

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/models"
)

type TransactionRepositoryInterface interface {
	CreateTransaction(transaction *models.TransactionModel) *custom_errors.BaseCustomError
}
