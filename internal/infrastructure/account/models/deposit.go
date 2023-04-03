package models

type DepositRequest struct {
	Amount float64 `json:"amount" validate:"required"`
}
