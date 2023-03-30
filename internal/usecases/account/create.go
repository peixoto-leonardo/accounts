package account

import (
	"context"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/pkg/utils/uuid"
)

func (c *usecase) Create(ctx context.Context, input CreateAccountInput) (CreateAccountOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	account, err := c.repository.Create(
		ctx,
		domain.NewAccount(
			uuid.New(),
			input.Name,
			input.CPF,
			0,
			time.Now(),
			time.Time{},
		),
	)

	if err != nil {
		return CreateAccountOutput{}, err
	}

	return CreateAccountOutput{account.GetId()}, nil
}
