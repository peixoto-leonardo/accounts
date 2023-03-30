package postgres

import (
	"fmt"
)

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string
}

func NewConfig() *config {
	return &config{
		host:     "127.0.0.1",
		database: "bank",
		port:     "5432",
		driver:   "postgres",
		user:     "dev",
		password: "dev",
	}
}

func (c *config) GetUrl() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.host,
		c.port,
		c.user,
		c.database,
		c.password,
	)
}
