package repository

import (
	"context"

	"github.com/pkg/errors"
)

type ctxKey string

const contextKey ctxKey = "TransactionContextKey"

func (r repository) WithTransaction(ctx context.Context, fn func(ctxTx context.Context) error) error {
	tx, err := r.db.BeginTx(ctx)

	if err != nil {
		return errors.Wrap(err, "error begin tx")
	}

	ctxTx := context.WithValue(ctx, contextKey, tx)

	err = fn(ctxTx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Wrap(err, "rollback error")
		}
		return err
	}

	return tx.Commit()
}
