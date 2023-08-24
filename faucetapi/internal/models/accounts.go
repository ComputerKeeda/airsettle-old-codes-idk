package models 


type accountDetails struct {
	Address string `json:"address"`
}

type jsonResponse struct {
	Success         bool   `json:"success"`
	TransactionHash string `json:"transactionhash"`
	Message         string `json:"message"`
	Description     string `json:"description"`
}