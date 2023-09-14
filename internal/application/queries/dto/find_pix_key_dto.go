package dto

type FindPixKeyQueryDTO struct {
	Id                    string  `json:"id"`
	PixKeyType            string  `json:"pixKeyType"`
	PixKey                string  `json:"response"`
	AccountType           string  `json:"accountType"`
	AccountNumber         int     `json:"accountNumber"`
	AgencyNumber          int     `json:"agencyNumber"`
	AccountHolderName     string  `json:"accountHolderName"`
	AccountHolderLastName *string `json:"accountHolderLastName"`
}
