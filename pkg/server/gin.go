package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginEngine struct {
	handler *gin.Engine
	server  *http.Server
}

func newGinEngine() *ginEngine {
	handler := gin.New()

	return &ginEngine{
		handler,
		&http.Server{
			Addr:    fmt.Sprintf(`:%d`, 3000),
			Handler: handler,
		},
	}
}

func (g ginEngine) Start() Server {
	go func() {
		if err := g.server.ListenAndServe(); err != nil {
			fmt.Print("Error starting http server")
		}
	}()

	return g
}

func (g ginEngine) GracefullyShutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := g.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to gracefully shuts down the server: %w", err)
	}

	return nil
}

func (g ginEngine) POST(path string, handler Handler) {
	g.handler.POST(path, func(ctx *gin.Context) {
		handler(ctx.Writer, ctx.Request)
	})
}
