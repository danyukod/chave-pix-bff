package commands

import (
	"github.com/danyukod/chave-pix-bff/internal/application/commands/dto"
	"github.com/danyukod/chave-pix-bff/internal/application/port/input"
	"github.com/danyukod/chave-pix-bff/internal/application/port/output"
)

func NewFindPixKeyByKeyUsecase(pixKeyGateway output.PixKeyGateway) input.FindPixKeyByKeyUsecase {
	return &FindPixKeyByKeyService{PixKeyGateway: pixKeyGateway}
}

type FindPixKeyByKeyService struct {
	output.PixKeyGateway
}

func (f *FindPixKeyByKeyService) Execute(input *dto.FindPixKeyByKeyInputDTO) (*dto.FindPixKeyByKeyOutputDTO, error) {
	pixKeyDomain, err := f.PixKeyGateway.FindPixKeyByKey(input.Key)
	if err != nil {
		return nil, err
	}

	outputDto := &dto.FindPixKeyByKeyOutputDTO{
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
