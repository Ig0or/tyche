package controllers

import (
	"github.com/Ig0or/tyche/src/application/data_types/responses"
	"github.com/Ig0or/tyche/src/application/ports/presenters_interface"
	"github.com/Ig0or/tyche/src/application/ports/use_cases_interface"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/logger_interface"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type AccountController struct {
	logger                   logger_interface.LoggerInterface
	createAccountPresenter   presenters_interface.CreateAccountPresenterInterface
	getAccountTokenPresenter presenters_interface.GetAccountTokenPresenterInterface
	createAccountUseCase     use_cases_interface.CreateAccountUseCaseInterface
	getAccountTokenUseCase   use_cases_interface.GetAccountTokenUseCaseInterface
}

type AccountControllerDependencies struct {
	dig.In

	Logger                   logger_interface.LoggerInterface                       `name:"Logger"`
	CreateAccountPresenter   presenters_interface.CreateAccountPresenterInterface   `name:"CreateAccountPresenter"`
	GetAccountTokenPresenter presenters_interface.GetAccountTokenPresenterInterface `name:"GetAccountTokenPresenter"`
	CreateAccountUseCase     use_cases_interface.CreateAccountUseCaseInterface      `name:"CreateAccountUseCase"`
	GetAccountTokenUseCase   use_cases_interface.GetAccountTokenUseCaseInterface    `name:"GetAccountTokenUseCase"`
}

func NewAccountController(dependencies AccountControllerDependencies) *AccountController {
	controller := &AccountController{
		logger:                   dependencies.Logger,
		createAccountPresenter:   dependencies.CreateAccountPresenter,
		getAccountTokenPresenter: dependencies.GetAccountTokenPresenter,
		createAccountUseCase:     dependencies.CreateAccountUseCase,
		getAccountTokenUseCase:   dependencies.GetAccountTokenUseCase,
	}

	return controller
}

func (controller *AccountController) CreateAccount(context *gin.Context) (*responses.BaseApiResponse, *custom_errors.BaseCustomError) {
	request, customError := controller.createAccountPresenter.FromContextToRequest(context)

	if customError != nil {
		controller.logger.Error(customError.Message, customError.OriginalError)

		return nil, customError
	}

	dto, customError := controller.createAccountUseCase.CreateAccount(request)

	if customError != nil {
		controller.logger.Error(customError.Message, customError.OriginalError)

		return nil, customError
	}

	response := controller.createAccountPresenter.FromDtoToResponse(dto)

	return response, nil
}

func (controller *AccountController) GetAccountToken(context *gin.Context) (*responses.BaseApiResponse, *custom_errors.BaseCustomError) {
	request, customError := controller.getAccountTokenPresenter.FromContextToRequest(context)

	if customError != nil {
		controller.logger.Error(customError.Message, customError.OriginalError)

		return nil, customError
	}

	dto, customError := controller.getAccountTokenUseCase.GetAccountToken(request)

	if customError != nil {
		controller.logger.Error(customError.Message, customError.OriginalError)

		return nil, customError
	}

	response := controller.getAccountTokenPresenter.FromDtoToResponse(dto)

	return response, nil
}
