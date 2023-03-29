package account

import (
	"context"
	"time"

	"github.com/peixoto-leonardo/accounts/pkg/documents"
	"github.com/peixoto-leonardo/accounts/pkg/money"
)

type AccountId string

func (a AccountId) String() string {
	return string(a)
}

type (
	AccountGateway interface {
		Create(context.Context, Account) (Account, error)
	}

	Account struct {
		id        AccountId
		balance   money.Money
		cpf       documents.CPF
		name      string
		createdAt time.Time
	}
)

func NewAccount(id AccountId, name string, cpf documents.CPF, balance money.Money, createdAt time.Time) Account {
	return Account{
		id:        id,
		balance:   balance,
		cpf:       cpf,
		name:      name,
		createdAt: createdAt,
	}
}

func (a Account) GetId() AccountId {
	return a.id
}

func (a Account) GetBalance() money.Money {
	return a.balance
}

func (a Account) GetCPF() documents.CPF {
	return a.cpf
}

func (a Account) GetName() string {
	return a.name
}

func (a Account) GetCreatedAt() time.Time {
	return a.createdAt
}
