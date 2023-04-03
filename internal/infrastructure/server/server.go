package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/peixoto-leonardo/accounts/internal/infrastructure/logger"
)

const (
	ShutdownTimeout = 5 * time.Second
)

type (
	server struct {
		*http.Server
		Config *Config
	}

	Interface interface {
		Start() Interface
		GracefullyShutdown() error
	}
)

func New(handler http.Handler) Interface {
	config := NewConfig()

	return &server{
		Server: &http.Server{Addr: config.GetAddr(), Handler: handler},
		Config: config,
	}
}

func (s *server) Start() Interface {
	go func() {
		log := logger.WithPrefix(context.Background(), "server")

		log.WithField("addr", s.Server.Addr).Info("listening")

		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithError(err).Fatal("listen error")
		}
	}()

	return s
}

func (s *server) GracefullyShutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	logger.WithPrefix(ctx, "server").Warn("shutting down server")

	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to gracefully shuts down the server: %w", err)
	}

	return nil
}
