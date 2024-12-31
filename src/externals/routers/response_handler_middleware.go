package routers

import (
	"github.com/Ig0or/tyche/src/application/data_types/responses"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controllerFunction func(context *gin.Context) (*responses.BaseApiResponse, *custom_errors.BaseCustomError)

func responseHandlerMiddleware(function controllerFunction) gin.HandlerFunc {
	return func(context *gin.Context) {
		var response *responses.BaseApiResponse

		response, customError := function(context)

		if customError != nil {
			response = buildErrorResponse(customError)
		}

		context.JSON(response.StatusCode, response)
	}
}

func buildErrorResponse(customError *custom_errors.BaseCustomError) *responses.BaseApiResponse {
	responseMessage := customError.Message

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
