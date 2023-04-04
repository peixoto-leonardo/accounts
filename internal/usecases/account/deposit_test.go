package account

import (
	"context"
	"testing"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/pkg/utils/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDepositAccount(t *testing.T) {
	t.Run("should deposit amount to an account", func(t *testing.T) {
		repository := &domain.AccountRepositoryMock{}
		usecase := New(time.Second, repository)
		account := &domain.Account{}
		repository.On("FindByID").Return(account, nil)
		repository.On("UpdateBalance").Return(nil)
		repository.On("CreateTransaction").Return(nil)

		err := usecase.Deposit(context.TODO(), uuid.New(), domain.Money(100_000))

		assert.NoError(t, err)
		assert.Equal(t, domain.Money(100_000), account.GetBalance())
	})
}
