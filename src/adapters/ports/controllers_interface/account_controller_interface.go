package controllers_interface

import (
	"github.com/Ig0or/tyche/src/application/data_types/responses"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/gin-gonic/gin"
)

type AccountControllerInterface interface {
	CreateAccount(context *gin.Context) (*responses.BaseApiResponse, *custom_errors.BaseCustomError)
}
