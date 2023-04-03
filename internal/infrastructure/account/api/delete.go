package api

import (
	"context"
	"net/http"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/pkg/utils/response"
)

func (a *api) Delete(ctx context.Context, accountId string) response.Response {
	err := a.usecase.Delete(ctx, accountId)

	if err == domain.ErrAccountNotFound {
		return response.New(http.StatusNotFound, response.NewError(err))
	}

	if err != nil {
		return response.New(http.StatusInternalServerError, response.NewError(err))
	}

	return response.New(http.StatusNoContent, nil)
}
