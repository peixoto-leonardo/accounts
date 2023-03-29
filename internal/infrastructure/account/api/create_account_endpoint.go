package api

import (
	"context"

	usecase "github.com/peixoto-leonardo/accounts/internal/application/usecases/account"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/models"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/presenters"
)

type CreateAccountEndpoint struct {
	usecase usecase.CreateAccountUseCase
}

func NewCreateAccountEndpoint(
	usecase usecase.CreateAccountUseCase,
) CreateAccountEndpoint {
	return CreateAccountEndpoint{usecase}
}

func (c CreateAccountEndpoint) Execute(ctx context.Context, request models.CreateAccountRequest) (models.CreateAccountResponse, error) {
	output, err := c.usecase.Execute(ctx, usecase.CreateAccountInput(request))

	if err != nil {
		return models.CreateAccountResponse{}, err
	}

	return presenters.CreateAccountOutputToResponse(output), nil
}
