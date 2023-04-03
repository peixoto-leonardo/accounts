package api

import (
	"context"
	"net/http"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/models"
	"github.com/peixoto-leonardo/accounts/pkg/utils/response"
)

func (a *api) Get(ctx context.Context, accountId string) response.Response {
	output, err := a.usecase.Get(ctx, accountId)

	if err == domain.ErrAccountNotFound {
		return response.New(http.StatusNotFound, response.NewError(err))
	}

	if err != nil {
		return response.New(http.StatusInternalServerError, response.NewError(err))
	}

	return response.New(http.StatusOK, models.AccountResponse(output))
}

func (a *api) GetStatement(ctx context.Context, accountId string) response.Response {
	output, err := a.usecase.GetStatement(ctx, accountId)

	if err != nil {
		return response.New(http.StatusInternalServerError, response.NewError(err))
	}

	return response.New(
		http.StatusOK,
		models.StatementResponse{Transactions: output},
	)
}
