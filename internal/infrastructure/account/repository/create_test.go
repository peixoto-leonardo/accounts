package repository

import (
	"context"
	"testing"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
	"github.com/peixoto-leonardo/accounts/pkg/utils/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	t.Run("should create an account ", func(t *testing.T) {
		sqlMock := &postgres.SQLMock{}
		sqlMock.On("ExecuteContext").Return(postgres.Result{}, nil)
		repositoy := New(sqlMock)
		account := domain.NewAccount(
			uuid.New(),
			"Test",
			"123455",
			0,
			time.Time{},
			time.Time{},
		)

		_, err := repositoy.Create(context.TODO(), account)

		assert.Nil(t, err)
		sqlMock.AssertExpectations(t)
	})
}
