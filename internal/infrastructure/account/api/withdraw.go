package api

import (
	"context"
	"net/http"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/models"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/logger"
	"github.com/peixoto-leonardo/accounts/pkg/utils/response"
)

var logWithdraw = logger.WithPrefix(context.TODO(), "withdraw")

func (a *api) Withdraw(ctx context.Context, accountId string, request models.WithdrawRequest) response.Response {
	if errs := a.validate(request); len(errs) > 0 {
		logWithdraw.WithField("errors", errs).Error("invalid input")

		return response.New(http.StatusBadRequest, response.NewErrorMessage(errs))
	}

	err := a.usecase.Withdraw(ctx, accountId, request.Amount)

	if err == domain.ErrAccountNotFound {
		return response.New(http.StatusNotFound, response.NewError(err))
	}

	if err != nil {
		return response.New(http.StatusInternalServerError, response.NewError(err))
	}

	return response.New(http.StatusNoContent, nil)
}
