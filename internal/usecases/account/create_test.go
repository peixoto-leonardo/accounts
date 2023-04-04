package account

import (
	"context"
	"testing"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/pkg/utils/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	t.Run("should create an account", func(t *testing.T) {
		account := domain.NewAccount(
			uuid.New(),
			"",
			"",
			0,
			time.Time{},
			time.Time{},
		)
		repository := &domain.AccountRepositoryMock{}
		repository.On("Create").Return(account, nil)
		usecase := New(time.Second, repository)

		output, err := usecase.Create(context.TODO(), CreateAccountInput{Name: "foo", CPF: "bar"})

		assert.NoError(t, err)
		assert.Equal(t, account.GetId(), output.Id)
	})
}
