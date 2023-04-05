package app

import (
	"github.com/peixoto-leonardo/accounts/pkg/http"
)

type (
	App struct {
		Server http.Server
	}
)

func New() *App {
	return &App{
		Server: http.NewServer(),
	}
}
