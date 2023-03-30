package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/logger"
)

type handler struct {
	db *sql.DB
}

var log = logger.WithPrefix(context.TODO(), "database")

func NewConnection() (*handler, error) {
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

func (p handler) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := p.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return nil
}
