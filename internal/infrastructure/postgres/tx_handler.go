package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type txHandler struct {
	tx *sql.Tx
}

func (t txHandler) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	if result, err := t.tx.ExecContext(ctx, query, args...); err != nil {
		return err
	} else {
		fmt.Println(result.RowsAffected())
		fmt.Println(query, args)
		return nil
	}

}
