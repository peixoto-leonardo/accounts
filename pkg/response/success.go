package response

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	code   int
	result interface{}
}

func NewSuccess(result interface{}, code int) Success {
	return Success{
		code:   code,
		result: result,
	}
}

func (r Success) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.code)

	return json.NewEncoder(w).Encode(r.result)
}
