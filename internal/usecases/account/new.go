package account

import (
	"context"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain"
)

type (
	CreateAccountInput struct {
		Name string
		CPF  string
	}

	CreateAccountOutput struct {
		Id string
	}

	AccountOutput struct {
		Id        string
		Name      string
		CPF       string
		Balance   float64
		CreatedAt string
	}

	TransactionOutput struct {
		Amount          float64   `json:"amount"`
		TransactionType string    `json:"type"`
		CreateAt        time.Time `json:"create_at"`
	}

	AccountUseCase interface {
		Create(context.Context, CreateAccountInput) (CreateAccountOutput, error)
		Delete(context.Context, string) error
		Get(context.Context, string) (AccountOutput, error)
		GetStatement(context.Context, string) ([]TransactionOutput, error)
		Deposit(context.Context, string, float64) error
		Withdraw(context.Context, string, float64) error
	}

	usecase struct {
		contextTimeout time.Duration
		repository     domain.AccountRepository
	}
)

func New(contextTimeout time.Duration, repository domain.AccountRepository) AccountUseCase {
	return &usecase{contextTimeout, repository}
}
