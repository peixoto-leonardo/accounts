package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/peixoto-leonardo/accounts/internal/domain"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
	"github.com/peixoto-leonardo/accounts/pkg/utils/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {
	t.Run("should return an account", func(t *testing.T) {
		txMock := &postgres.TxMock{}
		sqlMock := &postgres.SQLMock{}
		rowMock := &postgres.RowMock{}
		sqlMock.On("BeginTx").Return(txMock, nil)
		txMock.On("QueryRowContext").Return(rowMock, nil)
		rowMock.On("Scan").Return(nil)
		repository := New(sqlMock)

		_, err := repository.FindByID(context.TODO(), uuid.New())

		assert.Nil(t, err)
		sqlMock.AssertExpectations(t)
		txMock.AssertExpectations(t)
		rowMock.AssertExpectations(t)
	})

	t.Run("should return ErrAccountNotFound", func(t *testing.T) {
		txMock := &postgres.TxMock{}
		sqlMock := &postgres.SQLMock{}
		rowMock := &postgres.RowMock{}
		sqlMock.On("BeginTx").Return(txMock, nil)
		txMock.On("QueryRowContext").Return(rowMock, nil)
		rowMock.On("Scan").Return(sql.ErrNoRows)
		repository := New(sqlMock)

		_, err := repository.FindByID(context.TODO(), uuid.New())

		assert.EqualError(t, err, domain.ErrAccountNotFound.Error())
		sqlMock.AssertExpectations(t)
		txMock.AssertExpectations(t)
		rowMock.AssertExpectations(t)
	})
}
