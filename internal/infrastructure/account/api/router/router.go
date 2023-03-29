package router

import (
	"net/http"

	"github.com/peixoto-leonardo/accounts/internal/application/usecases/account"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/models"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/presenters"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/application"
	"github.com/peixoto-leonardo/accounts/pkg/parser"
	"github.com/peixoto-leonardo/accounts/pkg/response"
	"github.com/peixoto-leonardo/accounts/pkg/server"
)

type accountRouter struct {
	app *application.Application
}

func NewAccountRouter(app *application.Application) *accountRouter {
	return &accountRouter{app}
}

func (r accountRouter) ConfigEndpoints() {
	r.app.Server.POST("/v1/accounts", r.buildCreateAccountEndpoint())
}

func (router accountRouter) buildCreateAccountEndpoint() server.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models.CreateAccountRequest

		if err := parser.RequestBody(r.Body, &request); err != nil {
			response.NewBadRequestError(err).Send(w)
			return
		}

		output, err := router.app.IOC.GetCreateAccountUseCaseInstance().Execute(r.Context(), account.CreateAccountInput(request))

		if err != nil {
			response.NewInternalServerError().Send(w)
			return
		}

		response.NewSuccess(presenters.CreateAccountOutputToResponse(output), http.StatusCreated).Send(w)
	}
}
