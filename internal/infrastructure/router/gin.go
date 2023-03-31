package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ap "github.com/peixoto-leonardo/accounts/internal/infrastructure/account/api"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/di"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/validator"
	"github.com/peixoto-leonardo/accounts/pkg/utils/response"
)

type ginEngine struct {
	router *gin.Engine
	db     postgres.SQL
}

func New(db postgres.SQL) *ginEngine {
	engine := &ginEngine{gin.New(), db}

	engine.configHandlers()

	return engine
}

func (g *ginEngine) GetHandler() http.Handler {
	return g.router
}

func (g *ginEngine) configHandlers() {
	g.router.POST("/v1/accounts", g.buildCreateAccountHandler())
	g.router.DELETE("/v1/accounts/:account_id", g.buildDeleteAccountHandler())
}

func (g *ginEngine) buildCreateAccountHandler() gin.HandlerFunc {
	api := ap.New(di.GetCreateAccountUseCase(g.db), validator.New())

	return func(ctx *gin.Context) {
		api.Create(ctx.Writer, ctx.Request)
	}
}

func (g *ginEngine) buildDeleteAccountHandler() gin.HandlerFunc {
	api := ap.New(di.GetCreateAccountUseCase(g.db), validator.New())

	return func(ctx *gin.Context) {
		if err := api.Delete(ctx.Request.Context(), ctx.Param("account_id")); err != nil {
			response.NewError(err, http.StatusInternalServerError).Send(ctx.Writer)
			return
		}

		response.NewSuccess(nil, http.StatusNoContent).Send(ctx.Writer)
	}
}
