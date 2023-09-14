package queries

import (
	"github.com/danyukod/chave-pix-bff/internal/application/port/input"
	"github.com/danyukod/chave-pix-bff/internal/application/port/output"
	"github.com/danyukod/chave-pix-bff/internal/application/queries/dto"
)

func NewFindPixKeyByKeyQuery(pixKeyGateway output.PixKeyGateway) input.FindPixKeyByKeyUsecase {
	return &FindPixKeyByKeyQuery{PixKeyGateway: pixKeyGateway}
}

type FindPixKeyByKeyQuery struct {
	output.PixKeyGateway
}

func (f *FindPixKeyByKeyQuery) Execute(key string) (*dto.FindPixKeyQueryDTO, error) {
	pixKeyDomain, err := f.PixKeyGateway.FindPixKeyByKey(key)
	if err != nil {
		return nil, err
	}

	outputDto := &dto.FindPixKeyQueryDTO{
		Id:                    pixKeyDomain.GetID(),
		PixKeyType:            pixKeyDomain.GetPixKeyType().GetType(),
		PixKey:                pixKeyDomain.GetPixKey(),
		AccountType:           pixKeyDomain.GetAccount().GetAccountType().String(),
		AccountNumber:         pixKeyDomain.GetAccount().GetNumber(),
		AgencyNumber:          pixKeyDomain.GetAccount().GetAgency(),
		AccountHolderName:     pixKeyDomain.GetAccount().GetHolder().GetName(),
		AccountHolderLastName: pixKeyDomain.GetAccount().GetHolder().GetLastName(),
	}
	return outputDto, nil
}
