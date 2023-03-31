package repository

import (
	"context"
	"time"
)

func (r repository) Delete(ctx context.Context, accountID string) error {
	if err := r.db.ExecuteContext(
		ctx,
		`UPDATE accounts SET deleted_at = $1 WHERE id = $2`,
		time.Now(),
		accountID,
	); err != nil {
		return err
	}

	return nil
}
