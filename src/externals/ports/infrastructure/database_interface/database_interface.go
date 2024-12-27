package database_interface

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseInterface interface {
	GetConnection() (*pgxpool.Pool, *custom_errors.BaseCustomError)
}
