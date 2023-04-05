package http

import (
	"fmt"

	"github.com/peixoto-leonardo/accounts/pkg/utils/env"
)

type config struct {
	addr string
}

const defaultPort = 3000

func NewConfig() *config {
	return &config{
		addr: fmt.Sprintf(`:%d`, env.GetEnvOrDefaultInt("PORT", defaultPort)),
	}
}

func (c *config) GetAddr() string {
	return c.addr
}
