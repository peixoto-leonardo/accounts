package domain

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
)

type (
	MoneySuite struct {
		suite.Suite
		amount float64
	}
)

func (s *MoneySuite) SetupTest() {
	s.amount = gofakeit.Float64Range(0.0, 100_000.0)
}

func (s *MoneySuite) TestFloatToMoney() {
	s.Require().Equal(Money(s.amount*100), FloatToMoney(s.amount))
}

func (s *MoneySuite) TestMoneyToFloat64() {
	expected := float64(Money(s.amount*100)) / 100

	s.Require().Equal(expected, FloatToMoney(s.amount).Float64())
}

func TestMoneySuite(t *testing.T) {
	suite.Run(t, new(MoneySuite))
}
