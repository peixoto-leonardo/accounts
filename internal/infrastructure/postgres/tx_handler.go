package postgres

import (
	"context"
	"database/sql"
)

type (
	txHandler struct {
		tx *sql.Tx
	}
)

func (t txHandler) ExecuteContext(ctx context.Context, query string, args ...interface{}) (Result, error) {
	result, err := t.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return Result{}, err
	}

	rows, _ := result.RowsAffected()

	return Result{rows}, nil

}

func (t txHandler) QueryRowContext(ctx context.Context, query string, args ...interface{}) Row {
	return newRowHandler(t.tx.QueryRowContext(ctx, query, args...))
}

func (t txHandler) Commit() error {
	return t.tx.Commit()
}

func (t txHandler) Rollback() error {
	return t.tx.Rollback()
}
