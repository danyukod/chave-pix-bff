package queries

import (
	"github.com/danyukod/chave-pix-bff/internal/application/port/input"
	"github.com/danyukod/chave-pix-bff/internal/application/port/output"
	"github.com/danyukod/chave-pix-bff/internal/application/queries/dto"
)

type FindPixKeyQuery struct {
	pixKeyGateway output.PixKeyGateway
}

func NewFindPixKeyQuery(pixKeyGateway output.PixKeyGateway) input.FindPixKeyUsecase {
	return &FindPixKeyQuery{pixKeyGateway: pixKeyGateway}
}

func (f *FindPixKeyQuery) Execute() (*[]dto.FindPixKeyQueryDTO, error) {
	findPixKeyDTO, err := f.pixKeyGateway.FindPixKey()
	if err != nil {
		return nil, err
	}

	return &findPixKeyDTO, nil
}
