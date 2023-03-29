package account

import (
	"context"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain/account"
	"github.com/peixoto-leonardo/accounts/pkg/documents"
	"github.com/peixoto-leonardo/accounts/pkg/uuid"
)

type (
	CreateAccountUseCase interface {
		Execute(context.Context, CreateAccountInput) (CreateAccountOutput, error)
	}

	CreateAccountInput struct {
		Name string
		CPF  documents.CPF
	}

	CreateAccountOutput struct {
		Id string
	}

	createAccountUseCase struct {
		repository     account.AccountGateway
		contextTimeout time.Duration
	}
)

func NewCreateAccountUseCase(
	gateway account.AccountGateway,
	contextTimeout time.Duration,
) CreateAccountUseCase {
	return createAccountUseCase{
		gateway,
		contextTimeout,
	}
}

func (c createAccountUseCase) Execute(ctx context.Context, input CreateAccountInput) (CreateAccountOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	account, err := c.repository.Create(
		ctx,
		account.NewAccount(
			account.AccountId(uuid.New()),
			input.Name,
			input.CPF,
			0,
			time.Now(),
		),
	)

	if err != nil {
		return CreateAccountOutput{}, err
	}

	return CreateAccountOutput{account.GetId().String()}, nil
}
