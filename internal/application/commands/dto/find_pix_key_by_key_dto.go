package dto

type FindPixKeyByKeyInputDTO struct {
	Key string
}

type FindPixKeyByKeyOutputDTO struct {
	Id                    string `json:"id"`
	PixKeyType            string `json:"pixKeyType"`
	PixKey                string `json:"response"`
	AccountType           string `json:"accountType"`
	AccountNumber         int    `json:"accountNumber"`
	AgencyNumber          int    `json:"agencyNumber"`
	AccountHolderName     string `json:"accountHolderName"`
	AccountHolderLastName string `json:"accountHolderLastName"`
}
