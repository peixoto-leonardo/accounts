package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peixoto-leonardo/accounts/pkg/http/enums"
)

type (
	Response struct {
		Code int
		Data any
	}

	Handler func() Response

	router struct {
		engine *gin.Engine
	}

	Router interface {
		ServeHTTP(http.ResponseWriter, *http.Request)
		Bind(method enums.Method, path string, handler Handler)
	}
)

func NewRouter() Router {
	engine := gin.New()

	return &router{
		engine: engine,
	}
}

func (r *router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.engine.ServeHTTP(writer, request)
}

func (r *router) Bind(method enums.Method, path string, handler Handler) {
	switch method {
	case enums.POST:
		r.engine.POST(path, adapterHandler(handler))
	}
}

func adapterHandler(handler Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := handler()

		ctx.Negotiate(response.Code, gin.Negotiate{Data: response.Data})
	}
}
