package requests

type GetAccountTokenRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5,max=15"`
}
