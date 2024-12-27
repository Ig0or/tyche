package models

import (
	"github.com/google/uuid"
	"time"
)

type AccountModel struct {
	ID        int
	AccountId uuid.UUID
	Email     string
	Cpf       string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
