package postgres

import (
	"github.com/stretchr/testify/mock"
)

type RowsMock struct {
	mock.Mock
}

func (m *RowsMock) Scan(_ ...interface{}) error {
	return m.MethodCalled("Scan").Error(0)
}

func (m *RowsMock) Next() bool {
	return m.MethodCalled("Next").Bool(0)
}

func (m *RowsMock) Error() error {
	return m.MethodCalled("Error").Error(0)
}

func (m *RowsMock) Close() error {
	return m.MethodCalled("Close").Error(0)
}
