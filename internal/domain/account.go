package domain

import (
	"context"
	"errors"
	"time"

	"github.com/peixoto-leonardo/accounts/pkg/utils/uuid"
)

var (
	ErrAccountNotFound     = errors.New("account not found")
	ErrInsufficientBalance = errors.New("origin account does not have sufficient balance")
)

type (
	AccountRepository interface {
		Create(context.Context, *Account) (*Account, error)
		Delete(context.Context, string) error
		UpdateBalance(context.Context, string, Money) error
		FindByID(context.Context, string) (*Account, error)
		WithTx(context.Context, func(context.Context) error) error
		CreateTransaction(context.Context, Transaction) error
		GetStatement(context.Context, string) ([]Transaction, error)
	}

	Account struct {
		id           string
		name         string
		cpf          string
		balance      Money
		transactions []Transaction
		deletedAt    time.Time
		createdAt    time.Time
	}
)

func NewAccount(
	id string,
	name string,
	cpf string,
	balance Money,
	createdAt time.Time,
	deletedAt time.Time,
) *Account {
	return &Account{
		id:           id,
		name:         name,
		cpf:          cpf,
		balance:      balance,
		transactions: nil,
		deletedAt:    deletedAt,
		createdAt:    createdAt,
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

func (a *Account) GetBalance() Money {
	return a.balance
}

func (a *Account) Deposit(amount Money) {
	a.balance += amount

	a.transactions = append(
		a.transactions,
		NewDeposit(uuid.New(), a.id, amount, time.Now()),
	)
}

func (a *Account) Withdraw(amount Money) error {
	if a.balance < amount {
		return ErrInsufficientBalance
	}

	a.balance -= amount

	a.transactions = append(
		a.transactions,
		NewWithdraw(uuid.New(), a.id, amount, time.Now()),
	)

	return nil
}

func (a *Account) GetLastTransaction() Transaction {
	return a.transactions[len(a.transactions)-1]
}
