package postgres

import "context"

type (
	Result struct {
		RowsAffected int64
	}

	SQL interface {
		ExecuteContext(context.Context, string, ...interface{}) error
		BeginTx(context.Context) (Tx, error)
	}

	Tx interface {
		ExecuteContext(context.Context, string, ...interface{}) (Result, error)
		QueryRowContext(context.Context, string, ...interface{}) Row
		Commit() error
		Rollback() error
	}

	Row interface {
		Scan(dest ...interface{}) error
	}
)
