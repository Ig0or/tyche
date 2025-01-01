package requests

type CreateAccountRequest struct {
	Email          string  `json:"email" binding:"required,email"`
	Cpf            string  `json:"cpf" binding:"required"`
	Password       string  `json:"password" binding:"required,min=5,max=15"`
	InitialBalance float64 `json:"initial_balance" binding:"required,gte=0"`
}
