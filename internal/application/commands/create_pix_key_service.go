package commands

import (
	"github.com/danyukod/chave-pix-bff/internal/application/commands/dto"
	"github.com/danyukod/chave-pix-bff/internal/application/domain/model"
	"github.com/danyukod/chave-pix-bff/internal/application/domain/shared/aggregate"
	"github.com/danyukod/chave-pix-bff/internal/application/port/output"
)

type CreatePixKeyService struct {
	output.PixKeyGateway
}

func NewCreatePixKeyService(pixKeyGateway output.PixKeyGateway) *CreatePixKeyService {
	return &CreatePixKeyService{
		PixKeyGateway: pixKeyGateway,
	}
}

func (s *CreatePixKeyService) Execute(input *dto.CreatePixKeyInputDTO) error {

	holderDomain, err := aggregate.NewHolderDomain(input.AccountHolderName, input.AccountHolderLastName)
	if err != nil {
		return err
	}
	accountDomain, err := aggregate.NewAccountDomain(input.AccountNumber, input.AgencyNumber, input.AccountType, holderDomain)
	if err != nil {
		return err
	}
	pixKey, err := model.NewPixKeyDomain(input.PixKey, input.PixKeyType, accountDomain)

	err = s.PixKeyGateway.CreatePixKey(pixKey)
	if err != nil {
		return err
	}

	return nil
}
