package postgres

import (
	"context"

	"github.com/stretchr/testify/mock"

	mockUtils "github.com/peixoto-leonardo/accounts/pkg/utils/mock"
)

type SQLMock struct {
	mock.Mock
}

func (m *SQLMock) ExecuteContext(_ context.Context, _ string, _ ...interface{}) (Result, error) {
	args := m.MethodCalled("ExecuteContext")
	return args.Get(0).(Result), mockUtils.ReturnNilOrError(args, 1)
}

func (m *SQLMock) QueryContext(_ context.Context, _ string, _ ...interface{}) (Rows, error) {
	args := m.MethodCalled("QueryContext")
	return args.Get(0).(Rows), mockUtils.ReturnNilOrError(args, 1)
}

func (m *SQLMock) BeginTx(context.Context) (Tx, error) {
	args := m.MethodCalled("BeginTx")
	return args.Get(0).(Tx), mockUtils.ReturnNilOrError(args, 1)
}
