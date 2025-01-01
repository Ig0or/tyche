package entities

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/google/uuid"
	"github.com/paemuri/brdoc"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AccountEntity struct {
	ID             int
	AccountId      uuid.UUID
	Email          string
	Cpf            string
	HashedPassword string
	CreatedAt      time.Time
	Balance        decimal.Decimal
	Transactions   []*TransactionEntity
}

func (entity *AccountEntity) ValidatePassword(password string) *custom_errors.BaseCustomError {
	passwordBytes := []byte(password)
	hashedPasswordBytes := []byte(entity.HashedPassword)

	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)

	if err != nil {
		customError := custom_errors.NewBadRequestError("Fail to get account token because the email or password is incorrect.", err)

		return customError
	}

	return nil
}

func VerifyValidCpf(cpf string) *custom_errors.BaseCustomError {
	validCpf := brdoc.IsCPF(cpf)

	if !validCpf {
		customError := custom_errors.NewBadRequestError("Fail to create account because this CPF is invalid.", nil)

		return customError
	}

	return nil
}

func HashPassword(password string) (string, *custom_errors.BaseCustomError) {
	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		customError := custom_errors.NewInternalServerError("Error while trying to hash password in AccountEntity.", err)

		return "", customError
	}

	hashedPasswordString := string(hashedPassword)

	return hashedPasswordString, nil
}
