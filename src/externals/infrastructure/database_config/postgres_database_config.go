package database_config

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type PostgresDatabaseConfig struct {
	connectionPool *pgxpool.Pool
}

func NewPostgresDatabaseConfig() *PostgresDatabaseConfig {
	database := &PostgresDatabaseConfig{}

	database.openConnection()

	return database
}

func (database *PostgresDatabaseConfig) openConnection() {
	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")

	connectionPool, err := pgxpool.New(context.Background(), connectionString)

	defer connectionPool.Close()

	if err != nil {
		log.Fatal(err)
	}

	database.connectionPool = connectionPool
}

func (database *PostgresDatabaseConfig) GetConnection() *pgxpool.Pool {
	return database.connectionPool
}
