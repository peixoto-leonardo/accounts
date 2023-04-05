package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/peixoto-leonardo/accounts/pkg/http/router"
	"github.com/peixoto-leonardo/accounts/pkg/utils/os"
)

type (
	server struct {
		server *http.Server
	}

	Server interface {
		Start()
		GracefullyShutdown() error
	}
)

const (
	ShutdownTimeout = 5 * time.Second
)

func NewServer() Server {
	config := NewConfig()

	return &server{
		server: &http.Server{
			Addr:    config.GetAddr(),
			Handler: router.NewRouter(),
		},
	}

}

func (s *server) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("listen error")
		}
	}()

	os.WaitForInterruptSignal()

	if err := s.GracefullyShutdown(); err != nil {
		log.Fatal("server forced to shutdown")
	}
}

func (s *server) GracefullyShutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to gracefully shuts down the server: %w", err)
	}

	return nil
}
