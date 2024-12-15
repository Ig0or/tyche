package database_config

import (
	"context"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/i_logger_config"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/dig"
	"os"
)

type PostgresDatabaseConfig struct {
	connectionPool *pgxpool.Pool

	logger i_logger_config.ILoggerConfig
}

type PostgresDatabaseConfigDependencies struct {
	dig.In

	Logger i_logger_config.ILoggerConfig `name:"LoggerConfig"`
}

func NewPostgresDatabaseConfig(dependencies PostgresDatabaseConfigDependencies) *PostgresDatabaseConfig {
	database := &PostgresDatabaseConfig{logger: dependencies.Logger}

	database.openConnection()

	return database
}

func (database *PostgresDatabaseConfig) openConnection() {
	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")

	connectionPool, err := pgxpool.New(context.Background(), connectionString)

	defer connectionPool.Close()

	if err != nil {
		database.logger.Fatal("Fail to open database connection - Error: " + err.Error())
	}

	database.connectionPool = connectionPool
}

func (database *PostgresDatabaseConfig) GetConnection() *pgxpool.Pool {
	return database.connectionPool
}
