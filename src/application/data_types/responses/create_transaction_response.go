package responses

import (
	"github.com/Ig0or/tyche/src/domain/enums"
)

type CreateTransactionResponse struct {
	AccountId            string          `json:"account_id,omitempty"`
	Operation            enums.Operation `json:"operation"`
	Type                 enums.Type      `json:"type"`
	Amount               float64         `json:"amount"`
	DestinationAccountId string          `json:"destination_account_id,omitempty"`
}
