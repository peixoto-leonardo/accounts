package repository

import (
	"context"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
	"github.com/pkg/errors"
)

func (r repository) Deposit(ctx context.Context, accountID string, amount float64) error {
	tx, ok := ctx.Value("TransactionContextKey").(postgres.Tx)
	if !ok {
		var err error
		tx, err = r.db.BeginTx(ctx)
		if err != nil {
			return errors.Wrap(err, "error on find account by id")
		}
	}
	result, err := tx.ExecuteContext(
		ctx,
		`UPDATE accounts SET balance = balance + $1 WHERE id = $2 AND deleted_at IS NULL`,
		amount,
		accountID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return domain.ErrAccountNotFound
	}

	return nil
}
