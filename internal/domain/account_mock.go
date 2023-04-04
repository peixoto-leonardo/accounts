package domain

import (
	"context"

	mockUtils "github.com/peixoto-leonardo/accounts/pkg/utils/mock"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (m *AccountRepositoryMock) Create(_ context.Context, _ *Account) (*Account, error) {
	args := m.MethodCalled("Create")
	return args.Get(0).(*Account), mockUtils.ReturnNilOrError(args, 1)
}

func (m *AccountRepositoryMock) Delete(_ context.Context, _ string) error {
	args := m.MethodCalled("Delete")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *AccountRepositoryMock) UpdateBalance(_ context.Context, _ string, _ Money) error {
	args := m.MethodCalled("UpdateBalance")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *AccountRepositoryMock) FindByID(_ context.Context, _ string) (*Account, error) {
	args := m.MethodCalled("FindByID")
	return args.Get(0).(*Account), mockUtils.ReturnNilOrError(args, 1)
}

func (m *AccountRepositoryMock) WithTx(_ context.Context, _ func(context.Context) error) error {
	args := m.MethodCalled("WithTx")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *AccountRepositoryMock) CreateTransaction(_ context.Context, _ Transaction) error {
	args := m.MethodCalled("CreateTransaction")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *AccountRepositoryMock) GetStatement(_ context.Context, _ string) ([]Transaction, error) {
	args := m.MethodCalled("GetStatement")
	return args.Get(0).([]Transaction), mockUtils.ReturnNilOrError(args, 1)
}
