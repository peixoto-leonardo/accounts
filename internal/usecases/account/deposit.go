package account

import (
	"context"
)

func (c *usecase) Deposit(ctx context.Context, accountID string, amount float64) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.repository.WithTransaction(ctx, func(ctxTx context.Context) error {
		err := c.repository.Deposit(ctxTx, accountID, amount)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
