package domain

import "time"

type TransactionType int

const (
	Deposit TransactionType = iota
	Withdraw
)

func (t TransactionType) String() string {
	if t == Deposit {
		return "Deposit"
	}

	return "Withdraw"
}

type (
	Transaction struct {
		id              string
		accountId       string
		amount          float64
		transactionType TransactionType
		createdAt       time.Time
	}
)

func NewTransaction(id, accountId string, amount float64, transactionType TransactionType, createdAt time.Time) Transaction {
	return Transaction{id, accountId, amount, transactionType, createdAt}
}

func NewDeposit(id, accountId string, amount float64, createdAt time.Time) Transaction {
	return NewTransaction(id, accountId, amount, Deposit, createdAt)
}

func NewWithdraw(id, accountId string, amount float64, createdAt time.Time) Transaction {
	return NewTransaction(id, accountId, amount, Withdraw, createdAt)
}

func (t *Transaction) GetId() string {
	return t.id
}

func (t *Transaction) GetAccountId() string {
	return t.accountId
}

func (t *Transaction) GetAmount() float64 {
	return t.amount
}

func (t *Transaction) GetType() TransactionType {
	return t.transactionType
}

func (t *Transaction) GetCreatedAt() time.Time {
	return t.createdAt
}
