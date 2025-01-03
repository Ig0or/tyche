package presenters

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/entities"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/jackc/pgx/v5"
)

type AccountPresenter struct {
}

func NewAccountPresenter() *AccountPresenter {
	return &AccountPresenter{}
}

func (presenter *AccountPresenter) FromDatabaseResultToModel(rows pgx.Rows) (*models.AccountModel, *custom_errors.BaseCustomError) {
	var accountModel models.AccountModel

	rows.Next()

	err := rows.Scan(
		&accountModel.ID,
		&accountModel.AccountId,
		&accountModel.Email,
		&accountModel.Cpf,
		&accountModel.HashedPassword,
		&accountModel.CreatedAt,
	)

	if err != nil {
		customError := custom_errors.NewBadRequestError("Fail to get account because this account doesnt exist.", err)

		return nil, customError
	}

	return &accountModel, nil
}

func (presenter *AccountPresenter) FromModelToEntity(model *models.AccountModel) *entities.AccountEntity {
	accountEntity := &entities.AccountEntity{
		ID:             model.ID,
		AccountId:      model.AccountId,
		Email:          model.Email,
		Cpf:            model.Cpf,
		HashedPassword: model.HashedPassword,
		CreatedAt:      model.CreatedAt,
	}

	return accountEntity
}
