package controllers

import (
	"github.com/Ig0or/tyche/src/application/data_types/responses"
	"github.com/Ig0or/tyche/src/application/ports/presenters_interface"
	"github.com/Ig0or/tyche/src/application/ports/use_cases_interface"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/enums"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type TransactionController struct {
	createTransactionPresenter      presenters_interface.CreateTransactionPresenterInterface `name:"CreateTransactionPresenter"`
	createDepositTransactionUseCase use_cases_interface.CreateTransactionUseCaseInterface
	useCasesMap                     map[enums.Operation]use_cases_interface.CreateTransactionUseCaseInterface
}

type TransactionControllerDependencies struct {
	dig.In

	CreateTransactionPresenter      presenters_interface.CreateTransactionPresenterInterface `name:"CreateTransactionPresenter"`
	CreateDepositTransactionUseCase use_cases_interface.CreateTransactionUseCaseInterface    `name:"CreateDepositTransactionUseCase"`
}

func NewTransactionController(dependencies TransactionControllerDependencies) *TransactionController {
	useCasesMap := map[enums.Operation]use_cases_interface.CreateTransactionUseCaseInterface{
		enums.Deposit: dependencies.CreateDepositTransactionUseCase,
		//enums.Withdraw: "",
		//enums.Transfer: "",
	}

	controller := &TransactionController{
		createTransactionPresenter:      dependencies.CreateTransactionPresenter,
		createDepositTransactionUseCase: dependencies.CreateDepositTransactionUseCase,
		useCasesMap:                     useCasesMap,
	}

	return controller
}

func (controller *TransactionController) selectUseCaseByOperation(operation string) use_cases_interface.CreateTransactionUseCaseInterface {
	selectedUseCase := controller.useCasesMap[enums.Operation(operation)]

	return selectedUseCase
}

func (controller *TransactionController) CreateTransaction(context *gin.Context) (*responses.BaseApiResponse, *custom_errors.BaseCustomError) {
	request, customError := controller.createTransactionPresenter.FromContextToRequest(context)

	if customError != nil {
		return nil, customError
	}

	useCase := controller.selectUseCaseByOperation(request.Operation)

	dto, customError := useCase.CreateTransaction(request)

	if customError != nil {
		return nil, customError
	}

	response := controller.createTransactionPresenter.FromDtoToResponse(dto)

	return response, nil
}
