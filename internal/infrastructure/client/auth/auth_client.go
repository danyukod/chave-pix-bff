package auth

import "github.com/danyukod/chave-pix-bff/internal/application/domain/model"

type AuthClient struct {
}

func NewAuthClient() *AuthClient {
	return &AuthClient{}
}

func (a *AuthClient) GenerateToken(model.AuthDomainInterface) (string, error) {
	return "", nil
}
