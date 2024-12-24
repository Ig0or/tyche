package presenters

import (
	"github.com/Ig0or/tyche/src/adapters/presenters"
	"github.com/Ig0or/tyche/src/application/data_types/requests/account"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/gin-gonic/gin"
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
		customError := presenters.FormatRequestValidationError(err)

		return &request, customError
	}

	return &request, nil
}
