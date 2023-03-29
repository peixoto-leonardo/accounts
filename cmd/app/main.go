package main

import (
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/application"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/router"
)

func main() {
	app := application.
		NewApplication().
		ConfigWebServer().
		ConnectDb().
		ConfigIOC()

	router.Bind(app)

	app.StartServer()
}
