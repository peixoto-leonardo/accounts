package repository

import (
	"context"
	"strings"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
	"github.com/pkg/errors"
)

func (r repository) CreateTransaction(ctx context.Context, transaction domain.Transaction) error {
	tx, ok := ctx.Value("TransactionContextKey").(postgres.Tx)

	if !ok {
		var err error
		tx, err = r.db.BeginTx(ctx)

		if err != nil {
			return errors.Wrap(err, "error on insert transaction")
		}
	}

	var query = `
		INSERT INTO
			transactions(id, account_id, amount, type, created_at)
		VALUES
			($1, $2, $3, $4, $5)
	`

	if _, err := tx.ExecuteContext(
		ctx,
		query,
		transaction.GetId(),
		transaction.GetAccountId(),
		transaction.GetAmount(),
		strings.ToUpper(transaction.GetType().String()),
		transaction.GetCreatedAt(),
	); err != nil {
		return errors.Wrap(err, "error creating transaction")
	}

	return nil
}
