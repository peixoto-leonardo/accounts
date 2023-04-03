package account

import (
	"context"
)

func (c *usecase) Deposit(ctx context.Context, accountID string, amount float64) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.repository.WithTransaction(ctx, func(ctxTx context.Context) error {
		account, err := c.repository.FindByID(ctxTx, accountID)

		if err != nil {
			return err
		}

		account.Deposit(amount)

		if err := c.repository.UpdateBalance(ctxTx, account.GetId(), account.GetBalance()); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
