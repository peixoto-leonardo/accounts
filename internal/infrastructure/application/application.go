package application

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/peixoto-leonardo/accounts/internal/infrastructure/ioc"
	"github.com/peixoto-leonardo/accounts/pkg/database/sql"
	"github.com/peixoto-leonardo/accounts/pkg/server"
)

type Application struct {
	Server server.Server
	DB     sql.SQL
	IOC    *ioc.Container
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) ConnectDb() *Application {
	db, err := sql.NewPostgres()

	if err != nil {
		log.Fatalln(err, "Could not make a connection to the database")
	}

	fmt.Println("Successfully connected to the SQL database")

	a.DB = db

	return a
}

func (a *Application) ConfigWebServer() *Application {
	server, err := server.NewServer(server.GinServer)

	if err != nil {
		log.Fatalln(err)
	}

	a.Server = server

	return a
}

func (a *Application) ConfigIOC() *Application {
	a.IOC = ioc.NewContainer(a.DB)

	return a
}

func (a *Application) StartServer() {
	a.Server.Start()

	waitForInterruptSignal()

	if err := a.Server.GracefullyShutdown(); err != nil {
		fmt.Print("server forced to shutdown")
	}

	fmt.Print("server exiting")
}

func waitForInterruptSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
