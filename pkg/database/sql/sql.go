package sql

import "context"

type SQL interface {
	ExecuteContext(context.Context, string, ...interface{}) error
}
