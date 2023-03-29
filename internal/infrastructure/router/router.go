package router

import (
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/api/router"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/application"
)

func Bind(app *application.Application) {
	accountRouter := router.NewAccountRouter(app)

	accountRouter.ConfigEndpoints()
}
