package parser

import (
	"encoding/json"
	"io"
)

func RequestBody(body io.ReadCloser, data any) error {
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return err
	}

	defer body.Close()

	return nil
}
