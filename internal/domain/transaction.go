package domain

import "time"

type TransactionType int

const (
	Deposit TransactionType = iota
	Withdraw
)

type (
	Transaction struct {
		id              string
		accountId       string
		amount          float64
		transactionType TransactionType
		createdAt       time.Time
	}
)

func newTransaction(id, accountId string, amount float64, transactionType TransactionType, createdAt time.Time) Transaction {
	return Transaction{id, accountId, amount, transactionType, createdAt}
}

func newDeposit(id, accountId string, amount float64, createdAt time.Time) Transaction {
	return newTransaction(id, accountId, amount, Deposit, createdAt)
}

func newWithdraw(id, accountId string, amount float64, createdAt time.Time) Transaction {
	return newTransaction(id, accountId, amount, Withdraw, createdAt)
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
