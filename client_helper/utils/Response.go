package utils

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type AccountResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Address string `json:"address"`
}
