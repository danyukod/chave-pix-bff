package input

import "github.com/danyukod/chave-pix-bff/internal/application/commands/dto"

type CreatePixKeyUsecase interface {
	Execute(*dto.CreatePixKeyInputDTO) error
}
