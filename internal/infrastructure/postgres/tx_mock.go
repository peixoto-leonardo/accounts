package postgres

import (
	"context"

	mockUtils "github.com/peixoto-leonardo/accounts/pkg/utils/mock"
	"github.com/stretchr/testify/mock"
)

type TxMock struct {
	mock.Mock
}

func (m *TxMock) ExecuteContext(_ context.Context, _ string, _ ...interface{}) (Result, error) {
	args := m.MethodCalled("ExecuteContext")
	return args.Get(0).(Result), mockUtils.ReturnNilOrError(args, 1)
}

func (m *TxMock) QueryRowContext(_ context.Context, _ string, _ ...interface{}) Row {
	args := m.MethodCalled("QueryRowContext")
	return args.Get(0).(Row)
}

func (m *TxMock) Commit() error {
	args := m.MethodCalled("Commit")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *TxMock) Rollback() error {
	args := m.MethodCalled("Rollback")
	return mockUtils.ReturnNilOrError(args, 0)
}
