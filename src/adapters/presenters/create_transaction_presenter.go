package presenters

import (
	"github.com/Ig0or/tyche/src/application/data_types/dtos"
	"github.com/Ig0or/tyche/src/application/data_types/requests"
	"github.com/Ig0or/tyche/src/application/data_types/responses"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type CreateTransactionPresenter struct {
}

func NewCreateTransactionPresenter() *CreateTransactionPresenter {
	return &CreateTransactionPresenter{}
}

func (presenter *CreateTransactionPresenter) FromContextToRequest(context *gin.Context) (*requests.CreateTransactionRequest, *custom_errors.BaseCustomError) {
	var request requests.CreateTransactionRequest

	err := context.ShouldBindJSON(&request)

	if err != nil {
		message := custom_errors.FormatRequestValidationError(err)
		customError := custom_errors.NewBadRequestError(message, err)

		return nil, customError
	}

	accountId := context.Param("accountId")

	request.AccountId = accountId

	return &request, nil
}

func (presenter *CreateTransactionPresenter) FromModelToDto(transactionModel *models.TransactionModel) *dtos.TransactionDto {
	transactionDto := &dtos.TransactionDto{
		AccountId:            transactionModel.AccountId,
		Operation:            transactionModel.Operation,
		Type:                 transactionModel.Type,
		Amount:               transactionModel.Amount,
		DestinationAccountId: transactionModel.DestinationAccountId,
		CreatedAt:            transactionModel.CreatedAt,
	}

	return transactionDto
}

func (presenter *CreateTransactionPresenter) FromDtoToResponse(dto *dtos.TransactionDto) *responses.BaseApiResponse {
	amountInFloat, _ := dto.Amount.Float64()

	destinationAccountId := dto.DestinationAccountId.String()

	if dto.DestinationAccountId == uuid.Nil {
		destinationAccountId = ""
	}

	transactionResponse := &responses.CreateTransactionResponse{
		AccountId:            dto.AccountId.String(),
		Operation:            dto.Operation,
		Type:                 dto.Type,
		Amount:               amountInFloat,
		DestinationAccountId: destinationAccountId,
	}

	apiResponse := &responses.BaseApiResponse{
		Payload:    transactionResponse,
		Success:    true,
		StatusCode: http.StatusCreated,
	}

	return apiResponse
}
