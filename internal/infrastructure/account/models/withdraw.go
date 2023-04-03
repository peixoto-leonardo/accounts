package models

type WithdrawRequest struct {
	Amount float64 `json:"amount" validate:"required"`
}
