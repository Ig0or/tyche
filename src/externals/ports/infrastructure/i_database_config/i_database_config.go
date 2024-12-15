package i_database_config

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type IDatabaseConfig interface {
	GetConnection() *pgxpool.Pool
}
