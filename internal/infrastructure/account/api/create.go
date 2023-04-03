package api

import (
	"context"
	"net/http"

	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/models"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/logger"
	"github.com/peixoto-leonardo/accounts/internal/usecases/account"
	"github.com/peixoto-leonardo/accounts/pkg/utils/response"
)

var logCreateAccount = logger.WithPrefix(context.TODO(), "create-account")

func (a *api) Create(ctx context.Context, request models.CreateAccountRequest) response.Response {
	if errs := a.validate(request); len(errs) > 0 {
		logCreateAccount.WithField("errors", errs).Error("invalid input")
		return response.New(http.StatusBadRequest, response.NewErrorMessage(errs))
	}

	output, err := a.usecase.Create(ctx, account.CreateAccountInput(request))

	if err != nil {
		logCreateAccount.WithError(err).Error("error when creating a new account")
		return response.New(http.StatusInternalServerError, response.NewError(err))
	}

	return response.New(http.StatusCreated, models.CreateAccountResponse(output))
}
