package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
	"github.com/pkg/errors"
)

func (r repository) FindByID(ctx context.Context, accountID string) (*domain.Account, error) {
	tx, ok := ctx.Value("TransactionContextKey").(postgres.Tx)
	if !ok {
		var err error
		tx, err = r.db.BeginTx(ctx)
		if err != nil {
			return &domain.Account{}, errors.Wrap(err, "error find account by id")
		}
	}

	var (
		query = `
			SELECT 
				*
			FROM 
				accounts
			WHERE
				id = $1 and deleted_at IS NULL FOR NO KEY UPDATE
		`
		id        string
		name      string
		CPF       string
		balance   float64
		createdAt time.Time
		deletedAt sql.NullTime
	)

	err := tx.QueryRowContext(ctx, query, accountID).Scan(&id, &name, &CPF, &balance, &createdAt, &deletedAt)

	switch {
	case err == sql.ErrNoRows:
		return &domain.Account{}, domain.ErrAccountNotFound
	default:
		return domain.NewAccount(
			accountID,
			name,
			CPF,
			balance,
			createdAt,
			deletedAt.Time,
		), err
	}
}
