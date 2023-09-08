package input

import (
	"github.com/danyukod/chave-pix-bff/internal/application/commands/dto"
)

type FindPixKeyUsecase interface {
	Execute() (*[]dto.FindPixKeyOutputDTO, error)
}
