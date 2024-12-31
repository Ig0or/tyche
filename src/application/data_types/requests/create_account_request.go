package requests

type CreateAccountRequest struct {
	Email          string  `json:"email" binding:"required,email"`
	Cpf            string  `json:"cpf" binding:"required,len=11"`
	Password       string  `json:"password" binding:"required"`
	InitialBalance float64 `json:"initial_balance" binding:"required,gte=0"`
}
