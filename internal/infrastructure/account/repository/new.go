package repository

import (
	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
)

type repository struct {
	db postgres.SQL
}

func New(db postgres.SQL) domain.AccountRepository {
	return &repository{db}
}
