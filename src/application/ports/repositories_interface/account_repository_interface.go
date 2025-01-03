package repositories_interface

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/google/uuid"
)

type AccountRepositoryInterface interface {
	CreateAccount(account *models.AccountModel) *custom_errors.BaseCustomError
	GetAccountByEmail(email string) (*models.AccountModel, *custom_errors.BaseCustomError)
	GetAccountByAccountId(accountId uuid.UUID) (*models.AccountModel, *custom_errors.BaseCustomError)
}
