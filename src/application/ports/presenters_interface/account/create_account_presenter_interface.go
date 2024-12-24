package presenters_interface

import (
	"github.com/Ig0or/tyche/src/application/data_types/requests/account"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/gin-gonic/gin"
)

type CreateAccountPresenterInterface interface {
	FromContextToRequest(context *gin.Context) (*requests.CreateAccountRequest, *custom_errors.BaseCustomError)
}
