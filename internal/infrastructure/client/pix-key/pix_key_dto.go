package pix_key

import (
	"github.com/danyukod/chave-pix-bff/internal/application/domain/model"
	"github.com/danyukod/chave-pix-bff/internal/application/domain/shared/aggregate"
)

type PixKeyResponseDTO struct {
	Id                    string `json:"id"`
	PixKeyType            string `json:"pixKeyType"`
	PixKey                string `json:"response"`
	AccountType           string `json:"accountType"`
	AccountNumber         int    `json:"accountNumber"`
	AgencyNumber          int    `json:"agencyNumber"`
	AccountHolderName     string `json:"accountHolderName"`
	AccountHolderLastName string `json:"accountHolderLastName"`
}

func (d PixKeyResponseDTO) ToDomain() (model.PixKeyDomainInterface, error) {
	holder, err := aggregate.NewHolderDomain(d.AccountHolderName, d.AccountHolderLastName)
	if err != nil {
		return nil, err
	}
	account, err := aggregate.NewAccountDomain(d.AccountNumber, d.AgencyNumber, d.AccountType, holder)
	if err != nil {
		return nil, err
	}
	pixKey, err := model.NewPixKeyDomain(d.PixKeyType, d.PixKey, account)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
