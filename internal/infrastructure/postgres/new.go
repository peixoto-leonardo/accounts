package postgres

import (
	"context"
	"database/sql"

	"github.com/peixoto-leonardo/accounts/internal/infrastructure/logger"
	_ "github.com/lib/pq"
)

var log = logger.WithPrefix(context.TODO(), "database")

func New() (SQL, error) {
	config := NewConfig()
	db, err := sql.Open(config.driver, config.GetUrl())

	if err != nil {
		return &handler{}, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &handler{db}, nil
}

func newTxHandler(tx *sql.Tx) Tx {
	return &txHandler{tx}
}
