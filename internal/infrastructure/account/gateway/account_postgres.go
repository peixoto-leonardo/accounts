package gateway

import (
	"context"

	domain "github.com/peixoto-leonardo/accounts/internal/domain/account"
	"github.com/peixoto-leonardo/accounts/pkg/database/sql"
	"github.com/pkg/errors"
)

type AccountPostgresGateway struct {
	db sql.SQL
}

func NewAccountPostgresGateway(db sql.SQL) AccountPostgresGateway {
	return AccountPostgresGateway{db}
}

func (a AccountPostgresGateway) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	var query = `
		INSERT INTO
			accounts (id, name, cpf, balance, created_at)
		VALUES
			($1, $2, $3, $4, $5)
	`

	if err := a.db.ExecuteContext(
		ctx,
		query,
		account.GetId(),
		account.GetName(),
		account.GetCPF(),
		account.GetBalance(),
		account.GetCreatedAt(),
	); err != nil {
		return domain.Account{}, errors.Wrap(err, "error creating account")
	}

	return account, nil
}
