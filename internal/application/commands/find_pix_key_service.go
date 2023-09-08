package commands

import (
	"github.com/danyukod/chave-pix-bff/internal/application/commands/dto"
	"github.com/danyukod/chave-pix-bff/internal/application/port/output"
)

type FindPixKeyService struct {
	pixKeyGateway output.PixKeyGateway
}

func NewFindPixKeyUsecase(pixKeyGateway output.PixKeyGateway) *FindPixKeyService {
	return &FindPixKeyService{pixKeyGateway: pixKeyGateway}
}

func (f *FindPixKeyService) Execute() (*[]dto.FindPixKeyOutputDTO, error) {
	domain, err := f.pixKeyGateway.FindPixKey()
	if err != nil {
		return nil, err
	}
	var response []dto.FindPixKeyOutputDTO
	for _, pixKeyDomain := range domain {
		response = append(response, dto.FindPixKeyOutputDTO{
			Id:                    pixKeyDomain.GetID(),
			PixKeyType:            pixKeyDomain.GetPixKeyType().GetType(),
			PixKey:                pixKeyDomain.GetPixKey(),
			AccountType:           pixKeyDomain.GetAccount().GetAccountType().String(),
			AccountNumber:         pixKeyDomain.GetAccount().GetNumber(),
			AgencyNumber:          pixKeyDomain.GetAccount().GetAgency(),
			AccountHolderName:     pixKeyDomain.GetAccount().GetHolder().GetName(),
			AccountHolderLastName: pixKeyDomain.GetAccount().GetHolder().GetLastName(),
		})
	}
	return &response, nil
}
