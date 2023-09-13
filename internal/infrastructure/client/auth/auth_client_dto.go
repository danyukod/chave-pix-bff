package auth

type GenerateTokenDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
