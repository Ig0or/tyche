package presenters

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/entities"
	"github.com/Ig0or/tyche/src/domain/enums"
	"github.com/Ig0or/tyche/src/domain/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
	"time"
)

type TransactionPresenter struct {
}

func NewTransactionPresenter() *TransactionPresenter {
	return &TransactionPresenter{}
}

func (presenter *TransactionPresenter) FromRequestToEntity(accountId uuid.UUID, operation enums.Operation, type_ enums.Type, amount float64, destinationAccountId uuid.UUID) *entities.TransactionEntity {
	currentTime := time.Now().UTC()

	amountInDecimal := decimal.NewFromFloat(amount)

	transactionEntity := entities.TransactionEntity{
		AccountId:            accountId,
		Operation:            operation,
		Type:                 type_,
		Amount:               amountInDecimal,
		DestinationAccountId: destinationAccountId,
		CreatedAt:            currentTime,
	}

	return &transactionEntity
}

func (presenter *TransactionPresenter) FromEntityToModel(entity *entities.TransactionEntity) *models.TransactionModel {
	transactionModel := &models.TransactionModel{
		AccountId:            entity.AccountId,
		Operation:            entity.Operation,
		Type:                 entity.Type,
		Amount:               entity.Amount,
		DestinationAccountId: entity.DestinationAccountId,
		CreatedAt:            entity.CreatedAt,
	}

	return transactionModel
}

func (presenter *TransactionPresenter) fromDatabaseResultToModel(rows pgx.Rows) (*models.TransactionModel, *custom_errors.BaseCustomError) {
	var transactionModel models.TransactionModel

	err := rows.Scan(
		&transactionModel.ID,
		&transactionModel.AccountId,
		&transactionModel.Operation,
		&transactionModel.Type,
		&transactionModel.Amount,
		&transactionModel.DestinationAccountId,
		&transactionModel.CreatedAt,
	)

	if err != nil {
		customError := custom_errors.NewInternalServerError("Error while trying to create model in TransactionPresenter.", err)

		return nil, customError
	}

	return &transactionModel, nil
}

func (presenter *TransactionPresenter) FromDatabaseResultToModels(rows pgx.Rows) (*[]*models.TransactionModel, *custom_errors.BaseCustomError) {
	var transactionModels []*models.TransactionModel

	for rows.Next() {
		transactionModel, customError := presenter.fromDatabaseResultToModel(rows)

		if customError != nil {
			return nil, customError
		}

		transactionModels = append(transactionModels, transactionModel)

	}

	return &transactionModels, nil
}
