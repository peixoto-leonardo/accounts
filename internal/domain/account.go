package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrAccountNotFound = errors.New("account not found")
)

type (
	AccountRepository interface {
		Create(context.Context, *Account) (*Account, error)
		Delete(context.Context, string) error
		Deposit(context.Context, string, float64) error
		Withdraw(context.Context, string, float64) error
		FindByID(ctx context.Context, accountID string) (*Account, error)
		WithTransaction(context.Context, func(context.Context) error) error
	}

	Account struct {
		id        string
		name      string
		cpf       string
		balance   float64
		deletedAt time.Time
		createdAt time.Time
	}
)

func NewAccount(
	id string,
	name string,
	cpf string,
	balance float64,
	createdAt time.Time,
	deletedAt time.Time,
) *Account {
	return &Account{
		id,
		name,
		cpf,
		balance,
		deletedAt,
		createdAt,
	}
}

func (a *Account) GetId() string {
	return a.id
}

func (a *Account) GetName() string {
	return a.name
}

func (a *Account) GetCpf() string {
	return a.cpf
}

func (a *Account) IsActive() bool {
	return a.deletedAt.IsZero()
}

func (a *Account) GetCreatedAt() time.Time {
	return a.createdAt
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}
