package input

import (
	"github.com/danyukod/chave-pix-bff/internal/application/queries/dto"
)

type FindPixKeyUsecase interface {
	Execute() (*[]dto.FindPixKeyQueryDTO, error)
}
