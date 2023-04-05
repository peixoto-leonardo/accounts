package main

import (
	a "github.com/peixoto-leonardo/accounts/pkg/app"
)

func main() {
	app := a.New()

	app.Server.Start()
}
