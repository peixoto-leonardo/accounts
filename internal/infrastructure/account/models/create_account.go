package models

import "github.com/peixoto-leonardo/accounts/pkg/documents"

type CreateAccountRequest struct {
	Name string        `json:"name" validate:"required"`
	CPF  documents.CPF `json:"cpf" validate:"required"`
}

type CreateAccountResponse struct {
	Id string `json:"id"`
}
