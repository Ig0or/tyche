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

type CreateDepositTransactionUseCase struct {
	createTransactionPresenter presenters_interface.CreateTransactionPresenterInterface
	transactionPresenter       presenters_interface.TransactionPresenterInterface
	accountRepository          repositories_interface.AccountRepositoryInterface
	transactionRepository      repositories_interface.TransactionRepositoryInterface
}

type CreateDepositTransactionUseCaseDependencies struct {
	dig.In

	CreateTransactionPresenter presenters_interface.CreateTransactionPresenterInterface `name:"CreateTransactionPresenter"`
	TransactionPresenter       presenters_interface.TransactionPresenterInterface       `name:"TransactionPresenter"`
	AccountRepository          repositories_interface.AccountRepositoryInterface        `name:"AccountRepository"`
	TransactionRepository      repositories_interface.TransactionRepositoryInterface    `name:"TransactionRepository"`
}

func NewCreateDepositTransactionUseCase(dependencies CreateDepositTransactionUseCaseDependencies) *CreateDepositTransactionUseCase {
	createDepositTransactionUseCase := &CreateDepositTransactionUseCase{
		createTransactionPresenter: dependencies.CreateTransactionPresenter,
		transactionPresenter:       dependencies.TransactionPresenter,
		accountRepository:          dependencies.AccountRepository,
		transactionRepository:      dependencies.TransactionRepository,
	}

	return createDepositTransactionUseCase
}

func (useCase *CreateDepositTransactionUseCase) createDepositTransaction(accountId uuid.UUID, amount float64) (*models.TransactionModel, *custom_errors.BaseCustomError) {
	transactionEntity := useCase.transactionPresenter.FromRequestToEntity(accountId, enums.Deposit, enums.CashIn, amount, uuid.Nil)

	transactionModel := useCase.transactionPresenter.FromEntityToModel(transactionEntity)

	customError := useCase.transactionRepository.CreateTransaction(transactionModel)

	if customError != nil {
		return nil, customError
	}

	return transactionModel, nil
}

func (useCase *CreateDepositTransactionUseCase) getAccountModel(accountId uuid.UUID) (*models.AccountModel, *custom_errors.BaseCustomError) {
	accountModel, customError := useCase.accountRepository.GetAccountByAccountId(accountId)

	if customError != nil {
		return nil, customError
	}

	return accountModel, nil
}

func (useCase *CreateDepositTransactionUseCase) verifyDestinationAccountId(accountId uuid.UUID) *custom_errors.BaseCustomError {
	if accountId == uuid.Nil {
		customError := custom_errors.NewBadRequestError("Fail to create transaction because the deposit needs the destination account.", nil)

		return customError
	}

	return nil
}

func (useCase *CreateDepositTransactionUseCase) CreateTransaction(request *requests.CreateTransactionRequest) (*dtos.TransactionDto, *custom_errors.BaseCustomError) {
	customError := useCase.verifyDestinationAccountId(request.DestinationAccountId)

	if customError != nil {
		return nil, customError
	}

	accountModel, customError := useCase.getAccountModel(request.DestinationAccountId)

	if customError != nil {
		return nil, customError
	}

	transactionModel, customError := useCase.createDepositTransaction(accountModel.AccountId, request.Amount)

	if customError != nil {
		return nil, customError
	}

	transactionDto := useCase.createTransactionPresenter.FromModelToDto(transactionModel)

	return transactionDto, nil
}
