package database_interface

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseInterface interface {
	GetConnection() *pgxpool.Pool
}
