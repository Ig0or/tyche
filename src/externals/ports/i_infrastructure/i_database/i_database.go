package i_database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type IDatabase interface {
	GetConnection() *pgxpool.Pool
}
