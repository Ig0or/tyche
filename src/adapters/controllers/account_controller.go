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
	createAccountPresenter presenters_interface.CreateAccountPresenterInterface
	createAccountUseCase   use_cases_interface.CreateAccountUseCaseInterface
	logger                 logger_interface.LoggerInterface
}

type AccountControllerDependencies struct {
	dig.In

	CreateAccountPresenter presenters_interface.CreateAccountPresenterInterface `name:"CreateAccountPresenter"`
	CreateAccountUseCase   use_cases_interface.CreateAccountUseCaseInterface    `name:"CreateAccountUseCase"`
	Logger                 logger_interface.LoggerInterface                     `name:"Logger"`
}

func NewAccountController(dependencies AccountControllerDependencies) *AccountController {
	controller := &AccountController{
		createAccountPresenter: dependencies.CreateAccountPresenter,
		createAccountUseCase:   dependencies.CreateAccountUseCase,
		logger:                 dependencies.Logger,
	}

	return controller
}

func (controller *AccountController) CreateAccount(context *gin.Context) (*responses.BaseApiResponse, *custom_errors.BaseCustomError) {
	request, customError := controller.createAccountPresenter.FromContextToRequest(context)

	if customError != nil {
		controller.logger.Error(customError.Message, customError.OriginalError)

		return nil, customError
	}

	accountDto, customError := controller.createAccountUseCase.CreateAccount(request)

	if customError != nil {
		controller.logger.Error(customError.Message, customError.OriginalError)

		return nil, customError
	}

	response := controller.createAccountPresenter.FromDtoToResponse(accountDto)

	return response, nil
}
