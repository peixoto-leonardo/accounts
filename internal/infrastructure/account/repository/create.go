package repository

import (
	"context"

	"github.com/peixoto-leonardo/accounts/internal/domain"
)

func (r repository) Create(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	var query = `
		INSERT INTO
			accounts(id, name, cpf, balance, created_at)
		VALUES
			($1, $2, $3, $4, $5)
	`

	if err := r.db.ExecuteContext(
		ctx,
		query,
		account.GetId(),
		account.GetName(),
		account.GetCpf(),
		account.GetBalance(),
		account.GetCreatedAt(),
	); err != nil {
		return &domain.Account{}, err
	}

	return account, nil
}
