package repository

import (
	"context"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
	"github.com/pkg/errors"
)

func (r repository) UpdateBalance(ctx context.Context, accountId string, balance domain.Money) error {
	tx, ok := ctx.Value("TransactionContextKey").(postgres.Tx)

	if !ok {
		var err error
		tx, err = r.db.BeginTx(ctx)

		if err != nil {
			return errors.Wrap(err, "error on update balance")
		}
	}

	query := "UPDATE accounts SET balance = $1 WHERE id = $2 AND deleted_at IS NULL"

	result, err := tx.ExecuteContext(ctx, query, balance, accountId)

	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return domain.ErrAccountNotFound
	}

	return nil
}
