package domain

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
)

type (
	DepositAccountSuite struct {
		suite.Suite
		account Account
		amount  Money
	}

	WithdrawAccountSuite struct {
		suite.Suite
		account Account
		amount  Money
	}
)

func (s *DepositAccountSuite) SetupTest() {
	s.amount = Money(gofakeit.Number(1, 100_000))
	s.account = Account{balance: Money(gofakeit.Number(0, 100_000))}
}

func (s *DepositAccountSuite) TestDeposit() {
	expectBalance := s.account.balance + s.amount

	s.account.Deposit(s.amount)

	s.Require().Equal(Money(expectBalance), s.account.balance)
}

func TestDepositAccountSuite(t *testing.T) {
	suite.Run(t, new(DepositAccountSuite))
}

func (s *WithdrawAccountSuite) SetupTest() {
	s.amount = Money(gofakeit.Number(1, 100))
	s.account = Account{balance: Money(gofakeit.Number(200, 100_000))}
}

func (s *WithdrawAccountSuite) TestWithdraw() {
	expectBalance := s.account.balance - s.amount

	s.account.Withdraw(s.amount)

	s.Require().Equal(Money(expectBalance), s.account.balance)
}

func (s *WithdrawAccountSuite) TestWithdraw_ShouldReturnErrInsufficientBalance() {
	s.Require().Equal(ErrInsufficientBalance, s.account.Withdraw(s.account.balance+s.amount))
}

func TestWithdrawAccountSuite(t *testing.T) {
	suite.Run(t, new(WithdrawAccountSuite))
}
