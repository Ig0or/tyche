package presenters_interface

import (
	"github.com/Ig0or/tyche/src/domain/entities"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/google/uuid"
)

type TransactionPresenterInterface interface {
	CreateInitialBalanceTransactionEntity(accountId uuid.UUID, amount float64) *entities.TransactionEntity
	FromEntityToModel(entity *entities.TransactionEntity) *models.TransactionModel
}
