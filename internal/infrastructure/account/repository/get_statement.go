package repository

import (
	"context"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/pkg/errors"
)

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
			amount          int64
			typeTransaction string
		)

		if err = rows.Scan(&id, &amount, &createdAt, &typeTransaction); err != nil {
			return []domain.Transaction{}, errors.Wrap(err, "error listing transactions")
		}

		if typeTransaction == "WITHDRAW" {
			transactions = append(transactions, domain.NewWithdraw(id, accountId, domain.Money(amount), createdAt))
		} else {
			transactions = append(transactions, domain.NewDeposit(id, accountId, domain.Money(amount), createdAt))
		}
	}

	return transactions, nil
}
