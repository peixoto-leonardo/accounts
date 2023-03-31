package api

import (
	"context"
	"net/http"

	"github.com/peixoto-leonardo/accounts/internal/infrastructure/validator"
	usecase "github.com/peixoto-leonardo/accounts/internal/usecases/account"
)

type (
	api struct {
		usecase   usecase.AccountUseCase
		validator validator.Interface
	}

	Interface interface {
		Create(http.ResponseWriter, *http.Request)
		Delete(context.Context, string) error
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
