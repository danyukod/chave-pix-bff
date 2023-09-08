package input

import (
	"github.com/danyukod/chave-pix-bff/internal/application/commands/dto"
)

type FindPixKeyByKeyUsecase interface {
	Execute(*dto.FindPixKeyByKeyInputDTO) (*dto.FindPixKeyByKeyOutputDTO, error)
}
