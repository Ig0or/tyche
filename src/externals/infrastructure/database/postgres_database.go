package database

import (
	"context"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/logger_interface"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/dig"
	"os"
)

type PostgresDatabase struct {
	connectionPool *pgxpool.Pool
	logger         logger_interface.LoggerInterface
}

type PostgresDatabaseDependencies struct {
	dig.In

	Logger logger_interface.LoggerInterface `name:"Logger"`
}

func NewPostgresDatabase(dependencies PostgresDatabaseDependencies) *PostgresDatabase {
	database := &PostgresDatabase{logger: dependencies.Logger}

	database.openConnection()

	return database
}

func (database *PostgresDatabase) openConnection() {
	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")

	connectionPool, err := pgxpool.New(context.Background(), connectionString)

	defer connectionPool.Close()

	if err != nil {
		database.logger.Fatal("Fail to open database connection - Error: " + err.Error())
	}

	database.connectionPool = connectionPool
}

func (database *PostgresDatabase) GetConnection() *pgxpool.Pool {
	return database.connectionPool
}
