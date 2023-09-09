package pix_key

import (
	"github.com/danyukod/chave-pix-bff/internal/application/domain/model"
	"github.com/danyukod/chave-pix-bff/internal/application/domain/shared/aggregate"
)

type PixKeyRequestDTO struct {
	PixKeyType            string `json:"pixKeyType"`
	PixKey                string `json:"response"`
	AccountType           string `json:"accountType"`
	AccountNumber         int    `json:"accountNumber"`
	AgencyNumber          int    `json:"agencyNumber"`
	AccountHolderName     string `json:"accountHolderName"`
	AccountHolderLastName string `json:"accountHolderLastName"`
}

func NewPixKeyRequestDTOFromDomain(domain model.PixKeyDomainInterface) *PixKeyRequestDTO {
	return &PixKeyRequestDTO{
		PixKeyType:            domain.GetPixKeyType().GetType(),
		PixKey:                domain.GetPixKey(),
		AccountType:           domain.GetAccount().GetAccountType().String(),
		AccountNumber:         domain.GetAccount().GetNumber(),
		AgencyNumber:          domain.GetAccount().GetAgency(),
		AccountHolderName:     domain.GetAccount().GetHolder().GetName(),
		AccountHolderLastName: domain.GetAccount().GetHolder().GetLastName(),
	}
}

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
