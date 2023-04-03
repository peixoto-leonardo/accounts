package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ap "github.com/peixoto-leonardo/accounts/internal/infrastructure/account/api"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/di"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/models"
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
	accounts := g.router.Group("/v1/accounts")
	{
		accounts.POST("", g.buildCreateAccountHandler())
		accounts.DELETE("/:account_id", g.buildDeleteAccountHandler())
		accounts.GET("/:account_id", g.buildGetAccountHandler())
		accounts.GET("/:account_id/statement", g.buildGetAccountStatementHandler())
		accounts.PATCH("/:account_id/deposit", g.buildDepositInAccountHandler())
		accounts.PATCH("/:account_id/withdraw", g.buildWithdrawInAccountHandler())
	}
}

func (g *ginEngine) buildCreateAccountHandler() gin.HandlerFunc {
	api := ap.New(di.GetCreateAccountUseCase(g.db), validator.New())

	return func(c *gin.Context) {
		var request models.CreateAccountRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, response.NewError(err))
		}

		response := api.Create(c.Request.Context(), request)

		c.JSON(response.StatusCode, response.Data)
	}
}

func (g *ginEngine) buildDeleteAccountHandler() gin.HandlerFunc {
	api := ap.New(di.GetCreateAccountUseCase(g.db), validator.New())

	return func(c *gin.Context) {
		response := api.Delete(c.Request.Context(), c.Param("account_id"))

		c.JSON(response.StatusCode, response.Data)
	}
}

func (g *ginEngine) buildGetAccountHandler() gin.HandlerFunc {
	api := ap.New(di.GetCreateAccountUseCase(g.db), validator.New())

	return func(c *gin.Context) {
		response := api.Get(c.Request.Context(), c.Param("account_id"))

		c.JSON(response.StatusCode, response.Data)
	}
}

func (g *ginEngine) buildGetAccountStatementHandler() gin.HandlerFunc {
	api := ap.New(di.GetCreateAccountUseCase(g.db), validator.New())

	return func(c *gin.Context) {
		response := api.GetStatement(c.Request.Context(), c.Param("account_id"))

		c.JSON(response.StatusCode, response.Data)
	}
}

func (g *ginEngine) buildDepositInAccountHandler() gin.HandlerFunc {
	api := ap.New(di.GetCreateAccountUseCase(g.db), validator.New())

	return func(c *gin.Context) {
		var request models.DepositRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, response.NewError(err))
		}

		response := api.Deposit(c.Request.Context(), c.Param("account_id"), request)

		c.JSON(response.StatusCode, response.Data)
	}
}

func (g *ginEngine) buildWithdrawInAccountHandler() gin.HandlerFunc {
	api := ap.New(di.GetCreateAccountUseCase(g.db), validator.New())

	return func(c *gin.Context) {
		var request models.WithdrawRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, response.NewError(err))
		}

		response := api.Withdraw(c.Request.Context(), c.Param("account_id"), request)

		c.JSON(response.StatusCode, response.Data)
	}
}
