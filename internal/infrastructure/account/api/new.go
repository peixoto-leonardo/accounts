package api

import (
	"context"

	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/models"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/validator"
	usecase "github.com/peixoto-leonardo/accounts/internal/usecases/account"
	"github.com/peixoto-leonardo/accounts/pkg/utils/response"
)

type (
	api struct {
		usecase   usecase.AccountUseCase
		validator validator.Interface
	}

	Interface interface {
		Create(context.Context, models.CreateAccountRequest) response.Response
		Delete(context.Context, string) response.Response
		Get(context.Context, string) response.Response
		GetStatement(context.Context, string) response.Response
		Deposit(context.Context, string, models.DepositRequest) response.Response
		Withdraw(context.Context, string, models.WithdrawRequest) response.Response
	}
)

func New(usecase usecase.AccountUseCase, validator validator.Interface) Interface {
	return &api{usecase, validator}
}

func (a *api) validate(request interface{}) (msgs []string) {
	err := a.validator.Validate(request)

	if err != nil {
		msgs = append(msgs, a.validator.Messages()...)
	}

	return
}
