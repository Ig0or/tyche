package presenters_interface

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/entities"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/jackc/pgx/v5"
)

type AccountPresenterInterface interface {
	FromDatabaseResultToModel(row pgx.Rows) (*models.AccountModel, *custom_errors.BaseCustomError)
	FromModelToEntity(model *models.AccountModel) *entities.AccountEntity
}
