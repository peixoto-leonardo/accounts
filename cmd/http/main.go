package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/peixoto-leonardo/accounts/internal/infrastructure/logger"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/postgres"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/router"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/server"
)

var log = logger.WithPrefix(context.TODO(), "main")

func main() {
	conn, err := postgres.New()
	if err != nil {
		log.WithError(err).Fatal("cannot connect to the database")
	}

	srv := server.New(router.New(conn).GetHandler()).Start()

	waitForInterruptSignal()

	if err := srv.GracefullyShutdown(); err != nil {
		log.WithError(err).Fatal("server forced to shutdown")
	}

	log.Info("server exiting")
}

func waitForInterruptSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
