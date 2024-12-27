package repositories_interface

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/models"
)

type AccountRepositoryInterface interface {
	CreateAccount(account *models.AccountModel) *custom_errors.BaseCustomError
}
