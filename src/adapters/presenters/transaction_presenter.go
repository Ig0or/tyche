package presenters

import (
	"github.com/Ig0or/tyche/src/domain/entities"
	"github.com/Ig0or/tyche/src/domain/enums"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type TransactionPresenter struct {
}

func NewTransactionPresenter() *TransactionPresenter {
	return &TransactionPresenter{}
}

func (presenter *TransactionPresenter) CreateInitialBalanceTransactionEntity(accountId uuid.UUID, amount float64) *entities.TransactionEntity {
	currentTime := time.Now().UTC()

	amountInDecimal := decimal.NewFromFloat(amount)

	transactionEntity := entities.TransactionEntity{
		AccountId: accountId,
		Operation: enums.Deposit,
		Type:      enums.CashIn,
		Amount:    amountInDecimal,
		CreatedAt: currentTime,
	}

	return &transactionEntity
}

func (presenter *TransactionPresenter) FromEntityToModel(entity *entities.TransactionEntity) *models.TransactionModel {
	transactionModel := &models.TransactionModel{
		AccountId:   entity.AccountId,
		Operation:   entity.Operation,
		Type:        entity.Type,
		Amount:      entity.Amount,
		ToAccountId: entity.ToAccountId,
		CreatedAt:   entity.CreatedAt,
	}

	return transactionModel
}
