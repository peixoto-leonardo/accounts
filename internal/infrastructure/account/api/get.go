package api

import (
	"context"

	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/models"
)

func (a *api) Get(ctx context.Context, accountId string) (models.AccountResponse, error) {
	output, err := a.usecase.Get(ctx, accountId)

	if err != nil {
		return models.AccountResponse{}, err
	}

	return models.AccountResponse(output), nil
}
