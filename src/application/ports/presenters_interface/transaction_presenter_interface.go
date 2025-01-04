package presenters_interface

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/entities"
	"github.com/Ig0or/tyche/src/domain/enums"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TransactionPresenterInterface interface {
	FromRequestToEntity(accountId uuid.UUID, operation enums.Operation, type_ enums.Type, amount float64, destinationAccountId uuid.UUID) *entities.TransactionEntity
	FromEntityToModel(entity *entities.TransactionEntity) *models.TransactionModel
	FromDatabaseResultToModels(rows pgx.Rows) (*[]*models.TransactionModel, *custom_errors.BaseCustomError)
}
