package response

import (
	"github.com/pkg/errors"
)

var (
	ErrParameterInvalid = errors.New("parameter invalid")

	ErrInvalidInput = errors.New("invalid input")
)

type Error struct {
	Errors []string `json:"errors"`
}

func NewError(err error) *Error {
	return &Error{[]string{err.Error()}}
}

func NewErrorMessage(messages []string) *Error {
	return &Error{messages}
}
