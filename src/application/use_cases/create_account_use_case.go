package use_cases

import (
	"github.com/Ig0or/tyche/src/application/data_types/dtos"
	"github.com/Ig0or/tyche/src/application/data_types/requests"
	"github.com/Ig0or/tyche/src/application/ports/presenters_interface"
	"github.com/Ig0or/tyche/src/application/ports/repositories_interface"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/enums"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/google/uuid"
	"go.uber.org/dig"
)

type CreateAccountUseCase struct {
	createAccountPresenter presenters_interface.CreateAccountPresenterInterface
	transactionPresenter   presenters_interface.TransactionPresenterInterface
	accountRepository      repositories_interface.AccountRepositoryInterface
	transactionRepository  repositories_interface.TransactionRepositoryInterface
}

type CreateAccountUseCaseDependencies struct {
	dig.In

	CreateAccountPresenter presenters_interface.CreateAccountPresenterInterface  `name:"CreateAccountPresenter"`
	TransactionPresenter   presenters_interface.TransactionPresenterInterface    `name:"TransactionPresenter"`
	AccountRepository      repositories_interface.AccountRepositoryInterface     `name:"AccountRepository"`
	TransactionRepository  repositories_interface.TransactionRepositoryInterface `name:"TransactionRepository"`
}

func NewCreateAccountUseCase(dependencies CreateAccountUseCaseDependencies) *CreateAccountUseCase {
	createAccountUseCase := &CreateAccountUseCase{
		createAccountPresenter: dependencies.CreateAccountPresenter,
		transactionPresenter:   dependencies.TransactionPresenter,
		accountRepository:      dependencies.AccountRepository,
		transactionRepository:  dependencies.TransactionRepository,
	}

	return createAccountUseCase
}

func (useCase *CreateAccountUseCase) createAccount(request *requests.CreateAccountRequest) (*models.AccountModel, *custom_errors.BaseCustomError) {
	accountEntity, customError := useCase.createAccountPresenter.FromRequestToEntity(request)

	if customError != nil {
		return nil, customError
	}

	accountModel := useCase.createAccountPresenter.FromEntityToModel(accountEntity)

	customError = useCase.accountRepository.CreateAccount(accountModel)

	if customError != nil {
		return nil, customError
	}

	return accountModel, nil
}

func (useCase *CreateAccountUseCase) createTransaction(accountId uuid.UUID, amount float64) (*models.TransactionModel, *custom_errors.BaseCustomError) {
	transactionEntity := useCase.transactionPresenter.FromRequestToEntity(accountId, enums.Deposit, enums.CashIn, amount, uuid.Nil)

	transactionModel := useCase.transactionPresenter.FromEntityToModel(transactionEntity)

	customError := useCase.transactionRepository.CreateTransaction(transactionModel)

	if customError != nil {
		return nil, customError
	}

	return transactionModel, nil
}

func (useCase *CreateAccountUseCase) CreateAccount(request *requests.CreateAccountRequest) (*dtos.AccountDto, *custom_errors.BaseCustomError) {
	accountModel, customError := useCase.createAccount(request)

	if customError != nil {
		return nil, customError
	}

	transactionModel, customError := useCase.createTransaction(accountModel.AccountId, request.InitialBalance)

	if customError != nil {
		return nil, customError
	}

	accountDto := useCase.createAccountPresenter.FromModelToDto(accountModel, transactionModel)

	return accountDto, nil
}
