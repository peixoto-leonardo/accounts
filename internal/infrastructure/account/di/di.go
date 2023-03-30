package di

import (
	"time"

	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/repository"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
	usecases "github.com/peixoto-leonardo/accounts/internal/usecases/account"
)

func GetCreateAccountUseCase(conn postgres.SQL) usecases.AccountUseCase {
	return usecases.New(5*time.Second, repository.New(conn))
}
