package json

import (
	j "encoding/json"
	"io"
)

func ParserBody(body io.ReadCloser, to any) error {
	if err := j.NewDecoder(body).Decode(&to); err != nil {
		return nil
	}

	return nil
}
