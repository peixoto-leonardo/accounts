package sql

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string
}

func NewConfigPostgres() *config {
	return &config{
		host:     "127.0.0.1",
		database: "bank",
		port:     "5432",
		driver:   "postgres",
		user:     "dev",
		password: "dev",
	}
}
