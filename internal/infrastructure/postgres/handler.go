package postgres

import (
	"context"
	"database/sql"
)

type handler struct {
	db *sql.DB
}

func (p handler) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := p.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (p handler) BeginTx(ctx context.Context) (Tx, error) {
	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return handler{}, err
	}

	return newTxHandler(tx), nil
}
