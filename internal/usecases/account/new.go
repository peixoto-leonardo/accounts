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

	AccountUseCase interface {
		Create(context.Context, CreateAccountInput) (CreateAccountOutput, error)
		Delete(context.Context, string) error
	}

	usecase struct {
		contextTimeout time.Duration
		repository     domain.AccountRepository
	}
)

func New(contextTimeout time.Duration, repository domain.AccountRepository) AccountUseCase {
	return &usecase{contextTimeout, repository}
}
