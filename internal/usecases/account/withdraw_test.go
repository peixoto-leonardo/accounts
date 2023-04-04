package account

import (
	"context"
	"testing"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/pkg/utils/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWithdrawAccount(t *testing.T) {
	t.Run("should withdraw amount to an account", func(t *testing.T) {
		repository := &domain.AccountRepositoryMock{}
		usecase := New(time.Second, repository)
		account := domain.NewAccount(
			uuid.New(),
			"Test",
			"32778071008",
			domain.Money(100),
			time.Time{},
			time.Time{},
		)
		repository.On("FindByID").Return(account, nil)
		repository.On("UpdateBalance").Return(nil)
		repository.On("CreateTransaction").Return(nil)

		err := usecase.Withdraw(context.TODO(), uuid.New(), domain.Money(50))

		assert.NoError(t, err)
		assert.Equal(t, domain.Money(50), account.GetBalance())
	})
}
