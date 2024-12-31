package repositories

import (
	"context"
	"errors"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/database_interface"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/logger_interface"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/dig"
)

type AccountRepository struct {
	database database_interface.DatabaseInterface
	logger   logger_interface.LoggerInterface
}

type AccountRepositoryDependencies struct {
	dig.In

	Database database_interface.DatabaseInterface `name:"PostgresDatabase"`
	Logger   logger_interface.LoggerInterface     `name:"Logger"`
}

func NewAccountRepository(dependencies AccountRepositoryDependencies) *AccountRepository {
	return &AccountRepository{database: dependencies.Database, logger: dependencies.Logger}
}

func (repository *AccountRepository) handleErrorType(err error) *custom_errors.BaseCustomError {
	var pgError *pgconn.PgError
	var customError *custom_errors.BaseCustomError

	if errors.Is(err, pgError) {
		message := custom_errors.FormatDuplicatedKeyError(pgError)
		customError = custom_errors.NewBadRequestError(message, err)

	} else {
		customError = custom_errors.NewInternalServerError("Error while trying to create account in AccountRepository", err)
	}

	return customError
}

func (repository *AccountRepository) createAccount(account *models.AccountModel, connection *pgxpool.Pool) *custom_errors.BaseCustomError {
	defer connection.Close()

	query := `
		INSERT INTO accounts 
		(account_id, email, cpf, password, created_at)
		VALUES
		($1, $2, $3, $4, $5)
		`

	arguments := account.GetArgumentsToInsert()

	_, err := connection.Exec(context.TODO(), query, arguments...)

	if err != nil {
		customError := repository.handleErrorType(err)

		return customError
	}

	return nil
}

func (repository *AccountRepository) CreateAccount(account *models.AccountModel) *custom_errors.BaseCustomError {
	connection, customError := repository.database.GetConnection()

	if customError != nil {
		return customError
	}

	customError = repository.createAccount(account, connection)

	if customError != nil {
		return customError
	}

	repository.logger.Info("New account created successfully. AccountId: %s", account.AccountId)

	return nil
}
