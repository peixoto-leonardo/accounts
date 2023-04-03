package account

import (
	"context"
)

func (c *usecase) Withdraw(ctx context.Context, accountID string, amount float64) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.repository.WithTx(ctx, func(ctxTx context.Context) error {
		account, err := c.repository.FindByID(ctxTx, accountID)

		if err != nil {
			return err
		}

		if err = account.Withdraw(amount); err != nil {
			return err
		}

		if err := c.repository.UpdateBalance(ctxTx, account.GetId(), account.GetBalance()); err != nil {
			return err
		}

		if err = c.repository.CreateTransaction(ctxTx, account.GetLastTransaction()); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
