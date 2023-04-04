package account

import (
	"context"
	"testing"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/pkg/utils/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAccount(t *testing.T) {
	t.Run("should delete an account", func(t *testing.T) {
		accountId := uuid.New()
		repository := &domain.AccountRepositoryMock{}
		repository.On("Delete").Return(nil)
		usecase := New(time.Second, repository)

		err := usecase.Delete(context.TODO(), accountId)

		assert.NoError(t, err)
	})
}
