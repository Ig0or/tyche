package presenters_interface

import (
	"github.com/Ig0or/tyche/src/application/data_types/dtos"
	"github.com/Ig0or/tyche/src/application/data_types/requests"
	"github.com/Ig0or/tyche/src/application/data_types/responses"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/gin-gonic/gin"
)

type CreateTransactionPresenterInterface interface {
	FromContextToRequest(context *gin.Context) (*requests.CreateTransactionRequest, *custom_errors.BaseCustomError)
	FromModelToDto(transactionModel *models.TransactionModel) *dtos.TransactionDto
	FromDtoToResponse(dto *dtos.TransactionDto) *responses.BaseApiResponse
}
