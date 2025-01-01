package routers

import (
	"fmt"
	"github.com/Ig0or/tyche/src/application/data_types/responses"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/logger_interface"
	"github.com/Ig0or/tyche/src/externals/ports/routers_interface"
	"github.com/Ig0or/tyche/src/externals/ports/services_interface"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/dig"
	"net/http"
	"os"
)

type Middleware struct {
	logger       logger_interface.LoggerInterface
	tokenService services_interface.TokenServiceInterface
}

type MiddlewareDependencies struct {
	dig.In

	Logger       logger_interface.LoggerInterface         `name:"Logger"`
	TokenService services_interface.TokenServiceInterface `name:"TokenService"`
}

func NewMiddleware(dependencies MiddlewareDependencies) *Middleware {
	middleware := &Middleware{
		logger:       dependencies.Logger,
		tokenService: dependencies.TokenService,
	}

	return middleware
}

func (middleware *Middleware) buildErrorResponse(customError *custom_errors.BaseCustomError) *responses.BaseApiResponse {
	responseMessage := customError.Message

	middleware.logger.Error(responseMessage, customError.OriginalError)

	if customError.StatusCode == http.StatusInternalServerError {
		responseMessage = "An unexpected error occurred."
	}

	apiResponse := &responses.BaseApiResponse{
		Message:    responseMessage,
		Success:    false,
		StatusCode: customError.StatusCode,
	}

	return apiResponse
}

func (middleware *Middleware) processControllerFunction(function routers_interface.ControllerFunction, context *gin.Context) {
	response, customError := function(context)

	if customError != nil {
		response = middleware.buildErrorResponse(customError)
	}

	context.JSON(response.StatusCode, response)
}

func (middleware *Middleware) ResponseHandlerMiddleware(function routers_interface.ControllerFunction) gin.HandlerFunc {
	return func(context *gin.Context) {
		middleware.processControllerFunction(function, context)
	}
}

func (middleware *Middleware) verifyAccountId(context *gin.Context, claims jwt.MapClaims) *custom_errors.BaseCustomError {
	accountId := context.Param("accountId")

	if accountId != claims["account_id"] {
		customError := custom_errors.NewUnauthorizedError("Fail to validate JWT because this token doesnt belong to this account.", nil)

		return customError
	}

	return nil
}

func (middleware *Middleware) abortRequest(context *gin.Context, customError *custom_errors.BaseCustomError) {
	response := middleware.buildErrorResponse(customError)

	context.JSON(response.StatusCode, response)
	context.Abort()
}

func (middleware *Middleware) getAccessTokenFromHeader(context *gin.Context) (string, *custom_errors.BaseCustomError) {
	headerKeyFormat := os.Getenv("JWT_HEADER_KEY")

	accessToken := context.GetHeader(headerKeyFormat)

	if accessToken == "" {
		message := fmt.Sprintf("Fail to validate JWT because the key '%s' is missing.", headerKeyFormat)
		customError := custom_errors.NewBadRequestError(message, nil)

		return "", customError
	}

	return accessToken, nil

}

func (middleware *Middleware) ResponseHandlerWithJwtAuthMiddleware(function routers_interface.ControllerFunction) gin.HandlerFunc {
	return func(context *gin.Context) {
		accessToken, customError := middleware.getAccessTokenFromHeader(context)

		if customError != nil {
			middleware.abortRequest(context, customError)

			return
		}

		claims, customError := middleware.tokenService.ValidateToken(accessToken)

		if customError != nil {
			middleware.abortRequest(context, customError)

			return
		}

		customError = middleware.verifyAccountId(context, claims)

		if customError != nil {
			middleware.abortRequest(context, customError)

			return
		}

		if customError != nil {
			middleware.abortRequest(context, customError)

			return
		}

		middleware.processControllerFunction(function, context)
	}
}
