package entities

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/google/uuid"
	"github.com/paemuri/brdoc"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AccountEntity struct {
	ID           int
	AccountId    uuid.UUID
	Email        string
	Cpf          string
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Transactions []*TransactionEntity
}

func VerifyValidCpf(cpf string) *custom_errors.BaseCustomError {
	validCpf := brdoc.IsCPF(cpf)

	if !validCpf {
		customError := custom_errors.NewBadRequestError("This CPF is invalid.", nil)

		return customError
	}

	return nil
}

func HashPassword(password string) (string, *custom_errors.BaseCustomError) {
	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		customError := custom_errors.NewInternalServerError("Error while trying to hash password.", err)

		return "", customError
	}

	hashedPasswordString := string(hashedPassword)

	return hashedPasswordString, nil
}
