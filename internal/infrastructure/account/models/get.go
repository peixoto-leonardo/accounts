package models

type AccountResponse struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	CPF       string  `json:"cpf"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}
