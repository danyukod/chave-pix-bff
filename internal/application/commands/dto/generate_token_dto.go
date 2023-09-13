package dto

type GenerateTokenDTO struct {
	Email    string
	Password string
}

func (g GenerateTokenDTO) GetEmail() string {
	return g.Email
}

func (g GenerateTokenDTO) GetPassword() string {
	return g.Password
}
