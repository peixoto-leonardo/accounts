package postgres

import (
	"context"
	"database/sql"
)

type (
	Result struct {
		RowsAffected int64
	}

	SQL interface {
		ExecuteContext(context.Context, string, ...interface{}) (Result, error)
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

func newResult(r sql.Result) Result {
	rows, _ := r.RowsAffected()
	return Result{rows}
}
