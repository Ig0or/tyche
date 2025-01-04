package repositories

import (
	"context"
	"github.com/Ig0or/tyche/src/application/ports/presenters_interface"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/database_interface"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/logger_interface"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/dig"
)

type TransactionRepository struct {
	database             database_interface.DatabaseInterface
	logger               logger_interface.LoggerInterface
	transactionPresenter presenters_interface.TransactionPresenterInterface `name:"TransactionPresenter"`
}

type TransactionRepositoryDependencies struct {
	dig.In

	Database             database_interface.DatabaseInterface               `name:"PostgresDatabase"`
	Logger               logger_interface.LoggerInterface                   `name:"Logger"`
	TransactionPresenter presenters_interface.TransactionPresenterInterface `name:"TransactionPresenter"`
}

func NewTransactionRepository(dependencies TransactionRepositoryDependencies) *TransactionRepository {
	transactionRepository := &TransactionRepository{
		database:             dependencies.Database,
		logger:               dependencies.Logger,
		transactionPresenter: dependencies.TransactionPresenter,
	}

	return transactionRepository
}

func (repository *TransactionRepository) createTransaction(transaction *models.TransactionModel, connection *pgxpool.Pool) *custom_errors.BaseCustomError {
	defer connection.Close()

	query := `
		INSERT INTO transactions 
		(account_id, operation, type, amount, destination_account_id, created_at)
		VALUES
		($1, $2, $3, $4, $5, $6)
		`

	arguments := transaction.GetArgumentsToInsert()

	_, err := connection.Exec(context.TODO(), query, arguments...)

	if err != nil {
		customError := custom_errors.NewInternalServerError("Error while trying to create transaction in TransactionRepository.", err)

		return customError
	}

	return nil
}

func (repository *TransactionRepository) CreateTransaction(transaction *models.TransactionModel) *custom_errors.BaseCustomError {
	connection, err := repository.database.GetConnection()

	if err != nil {
		return err
	}

	err = repository.createTransaction(transaction, connection)

	if err != nil {
		return err
	}

	repository.logger.Info("New transaction created successfully. AccountId: %s", transaction.AccountId)

	return nil
}

func (repository *TransactionRepository) getTransactions(accountId uuid.UUID, connection *pgxpool.Pool) (*[]*models.TransactionModel, *custom_errors.BaseCustomError) {
	query := `SELECT * FROM transactions WHERE account_id = $1`

	rows, err := connection.Query(context.TODO(), query, accountId)

	defer rows.Close()

	if err != nil {
		customError := custom_errors.NewInternalServerError("Error while trying to access repository in TransactionRepository.", err)

		return nil, customError
	}

	transactionModels, customError := repository.transactionPresenter.FromDatabaseResultToModels(rows)

	if customError != nil {
		return nil, customError
	}

	return transactionModels, nil
}

func (repository *TransactionRepository) GetTransactionsByAccountID(accountId uuid.UUID) (*[]*models.TransactionModel, *custom_errors.BaseCustomError) {
	connection, customError := repository.database.GetConnection()

	if customError != nil {
		return nil, customError
	}

	accountModel, customError := repository.getTransactions(accountId, connection)

	if customError != nil {
		return nil, customError
	}

	return accountModel, nil
}
