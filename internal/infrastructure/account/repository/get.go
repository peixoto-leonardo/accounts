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
				id, name, cpf, balance, created_at, deleted_at
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

func (r repository) GetStatement(ctx context.Context, accountId string) ([]domain.Transaction, error) {
	var query = `
		SELECT 
			id, amount, created_at, type
		FROM 
			transactions 
		WHERE 
			account_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, accountId)

	if err != nil {
		return []domain.Transaction{}, err
	}

	transactions := make([]domain.Transaction, 0)
	for rows.Next() {
		var (
			id              string
			createdAt       time.Time
			amount          float64
			typeTransaction string
		)

		if err = rows.Scan(&id, &amount, &createdAt, &typeTransaction); err != nil {
			return []domain.Transaction{}, errors.Wrap(err, "error listing transactions")
		}

		if typeTransaction == "WITHDRAW" {
			transactions = append(transactions, domain.NewWithdraw(id, accountId, amount, createdAt))
		} else {
			transactions = append(transactions, domain.NewDeposit(id, accountId, amount, createdAt))
		}
	}

	return transactions, nil
}
