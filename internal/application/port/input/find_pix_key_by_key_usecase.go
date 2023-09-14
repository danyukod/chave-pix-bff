package input

import (
	"github.com/danyukod/chave-pix-bff/internal/application/queries/dto"
)

type FindPixKeyByKeyUsecase interface {
	Execute(string) (*dto.FindPixKeyQueryDTO, error)
}
