package database

import (
	"context"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type PostgresDatabase struct {
}

func NewPostgresDatabase() *PostgresDatabase {
	database := &PostgresDatabase{}

	return database
}

func (database *PostgresDatabase) openConnection() (*pgxpool.Pool, *custom_errors.BaseCustomError) {
	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")

	connectionPool, err := pgxpool.New(context.Background(), connectionString)

	if err != nil {
		customError := custom_errors.NewInternalServerError("Error while trying to open database connection", err)

		return nil, customError
	}

	return connectionPool, nil
}

func (database *PostgresDatabase) GetConnection() (*pgxpool.Pool, *custom_errors.BaseCustomError) {
	connection, err := database.openConnection()

	if err != nil {
		return nil, err
	}

	return connection, nil
}
