package responses

type CreateAccountResponse struct {
	AccountId string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}
