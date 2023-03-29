package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Error struct {
	code   int
	Errors []string `json:"errors"`
}

var errInternalServerError = errors.New("there was an internal error")

func NewBadRequestError(err error) *Error {
	return NewError(err, http.StatusBadRequest)
}

func NewInternalServerError() *Error {
	return NewError(errInternalServerError, http.StatusInternalServerError)
}

func NewError(err error, code int) *Error {
	return &Error{
		code:   code,
		Errors: []string{err.Error()},
	}
}

func (e Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.code)

	return json.NewEncoder(w).Encode(e)
}
