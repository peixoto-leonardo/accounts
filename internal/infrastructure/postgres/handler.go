package postgres

import (
	"context"
	"database/sql"
)

type handler struct {
	db *sql.DB
}

func (h handler) ExecuteContext(ctx context.Context, query string, args ...interface{}) (Result, error) {
	result, err := h.db.ExecContext(ctx, query, args...)

	if err != nil {
		return Result{}, err
	}

	return newResult(result), nil
}

func (h handler) QueryRowContext(ctx context.Context, query string, args ...interface{}) Row {
	row := h.db.QueryRowContext(ctx, query, args...)

	return newRowHandler(row)
}

func (h handler) BeginTx(ctx context.Context) (Tx, error) {
	tx, err := h.db.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return txHandler{}, err
	}

	return newTxHandler(tx), nil
}
