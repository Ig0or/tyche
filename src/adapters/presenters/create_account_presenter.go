package presenters

import (
	"github.com/Ig0or/tyche/src/application/data_types/requests/account"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/entities"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type CreateAccountPresenter struct {
}

func NewCreateAccountPresenter() *CreateAccountPresenter {
	return &CreateAccountPresenter{}
}

func (presenter *CreateAccountPresenter) FromContextToRequest(context *gin.Context) (*requests.CreateAccountRequest, *custom_errors.BaseCustomError) {
	var request requests.CreateAccountRequest

	err := context.ShouldBindJSON(&request)

	if err != nil {
		message := custom_errors.FormatRequestValidationError(err)
		customError := custom_errors.NewBadRequestError(message, err)

		return nil, customError
	}

	return &request, nil
}

func (presenter *CreateAccountPresenter) FromRequestToEntity(request *requests.CreateAccountRequest) (*entities.AccountEntity, *custom_errors.BaseCustomError) {
	err := entities.VerifyValidCpf(request.Cpf)

	if err != nil {
		return nil, err
	}

	hashedPassword, err := entities.HashPassword(request.Password)

	if err != nil {
		return nil, err
	}

	accountId := uuid.New()

	currentTime := time.Now().UTC()

	accountEntity := &entities.AccountEntity{
		AccountId: accountId,
		Email:     request.Email,
		Cpf:       request.Cpf,
		Password:  hashedPassword,
		CreatedAt: currentTime,
	}

	return accountEntity, nil
}

func (presenter *CreateAccountPresenter) FromEntityToModel(entity *entities.AccountEntity) *models.AccountModel {
	accountModel := &models.AccountModel{
		AccountId: entity.AccountId,
		Email:     entity.Email,
		Cpf:       entity.Cpf,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
	}

	return accountModel
}