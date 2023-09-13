package input

import (
	"github.com/danyukod/chave-pix-bff/internal/application/commands/dto"
	"github.com/danyukod/chave-pix-bff/internal/application/domain/model"
)

type CreatePixKeyUsecase interface {
	Execute(*dto.CreatePixKeyInputDTO) (model.PixKeyDomainInterface, error)
}
