package account

import (
	"context"

	"github.com/peixoto-leonardo/accounts/internal/domain"
)

func (c *usecase) Get(ctx context.Context, accountID string) (AccountOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	var (
		account *domain.Account
		err     error
	)
	err = c.repository.WithTx(ctx, func(ctxTx context.Context) error {
		account, err = c.repository.FindByID(ctxTx, accountID)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return AccountOutput{}, err
	}

	return AccountOutput{
		Id:        account.GetId(),
		Name:      account.GetName(),
		CPF:       account.GetCpf(),
		Balance:   account.GetBalance(),
		CreatedAt: account.GetCreatedAt().Format("2006-01-02"),
	}, nil
}
