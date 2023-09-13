package input

import "github.com/danyukod/chave-pix-bff/internal/application/port/output"

type GenerateTokenUseCaseInterface interface {
	Execute(dto GenerateTokenDTO) (string, error)
}

type GenerateTokenDTO struct {
	Email    string
	Password string
}

type GenerateTokenUseCase struct {
	output.AuthGateway
}

func NewGenerateTokenUseCase(
	authGateway output.AuthGateway,
) GenerateTokenUseCaseInterface {
	return &GenerateTokenUseCase{
		authGateway,
	}
}

func (g GenerateTokenUseCase) Execute(dto GenerateTokenDTO) (string, error) {
	//TODO implement me
	panic("implement me")
}
