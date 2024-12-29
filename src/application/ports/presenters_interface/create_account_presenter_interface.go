package presenters_interface

import (
	"github.com/Ig0or/tyche/src/application/data_types/requests/account"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/entities"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/gin-gonic/gin"
)

type CreateAccountPresenterInterface interface {
	FromContextToRequest(context *gin.Context) (*requests.CreateAccountRequest, *custom_errors.BaseCustomError)
	FromRequestToEntity(request *requests.CreateAccountRequest) (*entities.AccountEntity, *custom_errors.BaseCustomError)
	FromEntityToModel(entity *entities.AccountEntity) *models.AccountModel
}
