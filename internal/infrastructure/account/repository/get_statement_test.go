package repository

import (
	"context"
	"testing"

	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
	"github.com/peixoto-leonardo/accounts/pkg/utils/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetStatement(t *testing.T) {
	t.Run("should return transactions of an account", func(t *testing.T) {
		sqlMock := &postgres.SQLMock{}
		rowsMock := &postgres.RowsMock{}
		sqlMock.On("QueryContext").Return(rowsMock, nil)
		rowsMock.On("Next").Return(false)
		repository := New(sqlMock)

		_, err := repository.GetStatement(context.TODO(), uuid.New())

		assert.Nil(t, err)
		sqlMock.AssertExpectations(t)
		rowsMock.AssertExpectations(t)
	})
}
