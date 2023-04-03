package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peixoto-leonardo/accounts/internal/domain"
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
	g.router.GET("/v1/accounts/:account_id", g.buildGetAccountHandler())
	g.router.PATCH("/v1/accounts/:account_id", g.buildDepositInAccountHandler())
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

func (g *ginEngine) buildGetAccountHandler() gin.HandlerFunc {
	api := ap.New(di.GetCreateAccountUseCase(g.db), validator.New())

	return func(ctx *gin.Context) {
		r, err := api.Get(ctx.Request.Context(), ctx.Param("account_id"))

		if err == domain.ErrAccountNotFound {
			response.NewError(err, http.StatusNotFound).Send(ctx.Writer)
			return
		}

		if err != nil {
			response.NewError(err, http.StatusInternalServerError).Send(ctx.Writer)
			return
		}

		response.NewSuccess(r, http.StatusOK).Send(ctx.Writer)
	}
}

func (g *ginEngine) buildDepositInAccountHandler() gin.HandlerFunc {
	api := ap.New(di.GetCreateAccountUseCase(g.db), validator.New())

	return func(ctx *gin.Context) {
		err := api.Deposit(ctx.Request.Context(), ctx.Param("account_id"), 100.0)

		if err != nil {
			response.NewError(err, http.StatusInternalServerError).Send(ctx.Writer)
			return
		}

		response.NewSuccess(nil, http.StatusNoContent).Send(ctx.Writer)
	}
}
