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

func TestCreateTransaction(t *testing.T) {
	t.Run("should create a transaction of an account ", func(t *testing.T) {
		txMock := &postgres.TxMock{}
		sqlMock := &postgres.SQLMock{}
		sqlMock.On("BeginTx").Return(txMock, nil)
		txMock.On("ExecuteContext").Return(postgres.Result{}, nil)
		repositoy := New(sqlMock)
		transaction := domain.NewDeposit(
			uuid.New(),
			uuid.New(),
			10,
			time.Time{},
		)

		err := repositoy.CreateTransaction(context.TODO(), transaction)

		assert.Nil(t, err)
		sqlMock.AssertExpectations(t)
		txMock.AssertExpectations(t)
	})
}
