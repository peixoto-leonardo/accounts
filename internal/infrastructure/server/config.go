package server

import (
	"fmt"

	"github.com/peixoto-leonardo/accounts/pkg/utils/env"
)

type Config struct {
	addr string
}

const DefaultPort = 3000

func NewConfig() *Config {
	return &Config{
		addr: fmt.Sprintf(`:%d`, env.GetEnvOrDefaultInt("PORT", DefaultPort)),
	}
}

func (c *Config) GetAddr() string {
	return c.addr
}
