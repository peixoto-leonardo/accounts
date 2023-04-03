package postgres

import (
	"database/sql"
)

type rowHandler struct {
	row *sql.Row
}

func (pr *rowHandler) Scan(dest ...interface{}) error {
	if err := pr.row.Scan(dest...); err != nil {
		return err
	}

	return nil
}
