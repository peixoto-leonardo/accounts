package api

import (
	"context"
	"net/http"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/models"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/logger"
	"github.com/peixoto-leonardo/accounts/pkg/utils/response"
)

var logDeposit = logger.WithPrefix(context.TODO(), "deposit")

func (a *api) Deposit(ctx context.Context, accountId string, request models.DepositRequest) response.Response {
	if errs := a.validate(request); len(errs) > 0 {
		logDeposit.WithField("errors", errs).Error("invalid input")

		return response.New(http.StatusBadRequest, response.NewErrorMessage(errs))
	}

	err := a.usecase.Deposit(ctx, accountId, domain.FloatToMoney(request.Amount))

	if err == domain.ErrAccountNotFound {
		return response.New(http.StatusNotFound, response.NewError(err))
	}

	if err != nil {
		return response.New(http.StatusInternalServerError, response.NewError(err))
	}

	return response.New(http.StatusNoContent, nil)
}
