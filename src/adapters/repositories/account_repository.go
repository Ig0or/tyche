package repositories

import (
	"context"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/database_interface"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/dig"
)

type AccountRepository struct {
	database database_interface.DatabaseInterface
}

type AccountRepositoryDependencies struct {
	dig.In

	Database database_interface.DatabaseInterface `name:"PostgresDatabase"`
}

func NewAccountRepository(dependencies AccountRepositoryDependencies) *AccountRepository {
	return &AccountRepository{database: dependencies.Database}
}

func (repository *AccountRepository) createAccount(account *models.AccountModel, connection *pgxpool.Pool) *custom_errors.BaseCustomError {
	defer connection.Close()

	query := `
		INSERT INTO accounts 
		(account_id, email, cpf, password, created_at, updated_at)
		VALUES
		($1, $2, $3, $4, $5, $6)
		`

	arguments := []interface{}{account.AccountId, account.Email, account.Cpf, account.Password, account.CreatedAt, account.UpdatedAt}

	_, err := connection.Exec(context.TODO(), query, arguments...)

	// todo: verify error type to send the right message

	if err != nil {
		customErr := custom_errors.NewInternalServerError("Error while trying to create account", err)

		return customErr
	}

	return nil
}

func (repository *AccountRepository) CreateAccount(account *models.AccountModel) *custom_errors.BaseCustomError {
	connection, err := repository.database.GetConnection()

	if err != nil {
		return err
	}

	err = repository.createAccount(account, connection)

	if err != nil {
		return err
	}

	return nil
}
