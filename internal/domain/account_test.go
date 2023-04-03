package domain

import (
	"testing"
)

func TestAccount_Deposit(t *testing.T) {
	t.Parallel()

	type args struct {
		amount Money
	}

	tests := []struct {
		name     string
		account  Account
		args     args
		expected Money
	}{
		{
			name:     "Successful depositing balance",
			args:     args{amount: 10},
			account:  Account{balance: 0},
			expected: 10,
		},
		{
			name:     "Successful depositing balance",
			args:     args{amount: 100},
			account:  Account{balance: 100},
			expected: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.account.Deposit(tt.args.amount)

			if tt.account.GetBalance() != tt.expected {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					tt.account.GetBalance(),
					tt.expected,
				)
			}
		})
	}
}

func TestAccount_Deposit_Transaction(t *testing.T) {
	t.Parallel()

	type args struct {
		amount Money
	}

	tests := []struct {
		name     string
		account  Account
		args     args
		expected Transaction
	}{
		{
			name:    "Successful depositing balance",
			args:    args{amount: 10},
			account: Account{balance: 0, id: "62587372-c8b8-4d3b-ad20-47d7ac8294a2"},
			expected: Transaction{
				accountId:       "62587372-c8b8-4d3b-ad20-47d7ac8294a2",
				amount:          10,
				transactionType: Deposit,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.account.Deposit(tt.args.amount)

			transaction := tt.account.GetLastTransaction()

			if transaction.accountId != tt.expected.accountId {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					transaction.accountId,
					tt.expected.accountId,
				)
			}
			if transaction.amount != tt.args.amount {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					transaction.accountId,
					tt.expected.accountId,
				)
			}
			if transaction.transactionType != tt.expected.transactionType {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					transaction.transactionType,
					tt.expected.transactionType,
				)
			}
		})
	}
}

func TestAccount_Withdraw(t *testing.T) {
	t.Parallel()

	type args struct {
		amount Money
	}

	tests := []struct {
		name        string
		account     Account
		args        args
		expected    Money
		expectedErr error
	}{
		{
			name:     "Successful withdrawing balance",
			args:     args{amount: 100},
			account:  Account{balance: 200},
			expected: 100,
		},
		{
			name: "error when withdrawing account balance without sufficient balance",
			args: args{
				amount: 500,
			},
			account:     Account{balance: 200},
			expectedErr: ErrInsufficientBalance,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.account.Withdraw(tt.args.amount); (err != nil) && (err.Error() != tt.expectedErr.Error()) {
				t.Errorf("[TestCase '%s'] ResultError: '%v' | ExpectedError: '%v'",
					tt.name,
					err,
					tt.expectedErr.Error(),
				)
				return
			}

			if tt.expectedErr == nil && tt.account.GetBalance() != tt.expected {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					tt.account.GetBalance(),
					tt.expected,
				)
			}
		})
	}
}

func TestAccount_Withdraw_Transaction(t *testing.T) {
	t.Parallel()

	type args struct {
		amount Money
	}

	tests := []struct {
		name        string
		account     Account
		args        args
		expected    Transaction
		expectedErr error
	}{
		{
			name:     "Successful withdrawing balance",
			args:     args{amount: 100},
			account:  Account{balance: 200, id: "62587372-c8b8-4d3b-ad20-47d7ac8294a2"},
			expected: Transaction{amount: 100, transactionType: Withdraw, accountId: "62587372-c8b8-4d3b-ad20-47d7ac8294a2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.account.Withdraw(tt.args.amount)

			transaction := tt.account.GetLastTransaction()

			if transaction.accountId != tt.expected.accountId {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					transaction.accountId,
					tt.expected.accountId,
				)
			}
			if transaction.amount != tt.args.amount {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					transaction.accountId,
					tt.expected.accountId,
				)
			}
			if transaction.transactionType != tt.expected.transactionType {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					transaction.transactionType,
					tt.expected.transactionType,
				)
			}
		})
	}
}
