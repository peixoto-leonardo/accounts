package sql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type postgres struct {
	db *sql.DB
}

func NewPostgres() (*postgres, error) {
	config := NewConfigPostgres()

	dataSource := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.host,
		config.port,
		config.user,
		config.database,
		config.password,
	)

	db, err := sql.Open(config.driver, dataSource)
	if err != nil {
		return &postgres{}, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &postgres{db}, nil
}

func (p postgres) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := p.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return nil
}
