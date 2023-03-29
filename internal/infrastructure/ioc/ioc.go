package ioc

import (
	"time"

	"github.com/peixoto-leonardo/accounts/internal/application/usecases/account"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/gateway"
	"github.com/peixoto-leonardo/accounts/pkg/database/sql"
)

type Container struct {
	db                     sql.SQL
	AccountPostgresGateway gateway.AccountPostgresGateway
}

func NewContainer(db sql.SQL) *Container {
	return &Container{
		db,
		gateway.NewAccountPostgresGateway(db),
	}
}

var createAccountUseCaseInstance account.CreateAccountUseCase

func (c *Container) GetCreateAccountUseCaseInstance() account.CreateAccountUseCase {
	if createAccountUseCaseInstance == nil {
		createAccountUseCaseInstance = account.NewCreateAccountUseCase(c.AccountPostgresGateway, 5*time.Second)
	}

	return createAccountUseCaseInstance
}
