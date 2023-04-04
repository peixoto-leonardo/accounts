package account

import (
	"context"
	"testing"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/pkg/utils/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T) {
	t.Run("should get an account", func(t *testing.T) {
		repository := &domain.AccountRepositoryMock{}
		repository.On("FindByID").Return(&domain.Account{}, nil)
		usecase := New(time.Second, repository)

		_, err := usecase.Get(context.TODO(), uuid.New())

		assert.NoError(t, err)
	})
}

func TestGetStatement(t *testing.T) {
	t.Run("should get statement of an amount", func(t *testing.T) {
		repository := &domain.AccountRepositoryMock{}
		repository.On("GetStatement").Return([]domain.Transaction{}, nil)
		usecase := New(time.Second, repository)

		_, err := usecase.GetStatement(context.TODO(), uuid.New())

		assert.NoError(t, err)
	})
}
