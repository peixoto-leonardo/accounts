package postgres

import "context"

type SQL interface {
	ExecuteContext(context.Context, string, ...interface{}) error
	BeginTx(context.Context) (Tx, error)
}

type Tx interface {
	ExecuteContext(context.Context, string, ...interface{}) error
	QueryRowContext(context.Context, string, ...interface{}) Row
	Commit() error
	Rollback() error
}

type Row interface {
	Scan(dest ...interface{}) error
}
