package postgres

import (
	mockUtils "github.com/peixoto-leonardo/accounts/pkg/utils/mock"
	"github.com/stretchr/testify/mock"
)

type RowMock struct {
	mock.Mock
}

func (m *RowMock) Scan(_ ...interface{}) error {
	args := m.MethodCalled("Scan")
	return mockUtils.ReturnNilOrError(args, 0)
}
