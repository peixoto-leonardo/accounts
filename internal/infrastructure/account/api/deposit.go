package api

import (
	"context"
)

func (a *api) Deposit(ctx context.Context, accountId string, amount float64) error {
	if err := a.usecase.Deposit(ctx, accountId, amount); err != nil {
		return err
	}

	return nil
}
