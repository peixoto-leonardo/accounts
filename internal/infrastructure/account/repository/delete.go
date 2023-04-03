package repository

import (
	"context"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain"
)

func (r repository) Delete(ctx context.Context, accountID string) error {
	result, err := r.db.ExecuteContext(
		ctx,
		`UPDATE accounts SET deleted_at = $1 WHERE id = $2`,
		time.Now(),
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
