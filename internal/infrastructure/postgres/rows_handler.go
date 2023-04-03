package postgres

import "database/sql"

type rowsHandler struct {
	rows *sql.Rows
}

func (r *rowsHandler) Scan(dest ...interface{}) error {
	if err := r.rows.Scan(dest...); err != nil {
		return err
	}

	return nil
}

func (r *rowsHandler) Next() bool {
	return r.rows.Next()
}

func (r *rowsHandler) Error() error {
	return r.rows.Err()
}

func (r *rowsHandler) Close() error {
	return r.rows.Close()
}
