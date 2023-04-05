package repository

import (
	"context"
	"testing"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
	"github.com/peixoto-leonardo/accounts/pkg/utils/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAccount(t *testing.T) {
	t.Run("should delete an account ", func(t *testing.T) {
		sqlMock := &postgres.SQLMock{}
		sqlMock.On("ExecuteContext").Return(postgres.Result{RowsAffected: 1}, nil)
		repositoy := New(sqlMock)

		err := repositoy.Delete(context.TODO(), uuid.New())

		assert.Nil(t, err)
		sqlMock.AssertExpectations(t)
	})

	t.Run("should return an ErrAccountNotFound when RowsAffected is zero", func(t *testing.T) {
		sqlMock := &postgres.SQLMock{}
		sqlMock.On("ExecuteContext").Return(postgres.Result{RowsAffected: 0}, nil)
		repositoy := New(sqlMock)

		err := repositoy.Delete(context.TODO(), uuid.New())

		assert.EqualError(t, err, domain.ErrAccountNotFound.Error())
		sqlMock.AssertExpectations(t)
	})
}
