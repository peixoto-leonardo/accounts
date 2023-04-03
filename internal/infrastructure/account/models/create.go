package models

type CreateAccountRequest struct {
	Name string `json:"name" validate:"required"`
	CPF  string `json:"cpf" validate:"required"`
}

type CreateAccountResponse struct {
	Id string `json:"id"`
}
