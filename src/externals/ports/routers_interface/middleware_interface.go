package routers_interface

import (
	"github.com/Ig0or/tyche/src/application/data_types/responses"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/gin-gonic/gin"
)

type ControllerFunction func(context *gin.Context) (*responses.BaseApiResponse, *custom_errors.BaseCustomError)

type MiddlewareInterface interface {
	ResponseHandlerMiddleware(function ControllerFunction) gin.HandlerFunc
	ResponseHandlerWithJwtAuthMiddleware(function ControllerFunction) gin.HandlerFunc
}
