package dto

type CreatePixKeyInputDTO struct {
	PixKeyType            string `json:"pixKeyType"`
	PixKey                string `json:"response"`
	AccountType           string `json:"accountType"`
	AgencyNumber          int    `json:"agencyNumber"`
	AccountNumber         int    `json:"accountNumber"`
	AccountHolderName     string `json:"accountHolderName"`
	AccountHolderLastName string `json:"accountHolderLastName"`
}
