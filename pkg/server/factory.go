package server

import (
	"errors"
	"net/http"
	"time"
)

const (
	ShutdownTimeout = 5 * time.Second
)

type (
	Server interface {
		POST(string, Handler)
		Start() Server
		GracefullyShutdown() error
	}

	Handler func(http.ResponseWriter, *http.Request)
)

const (
	GinServer = iota
)

var (
	errInvalidServerInstance = errors.New("invalid server instance")
)

func NewServer(instance int) (Server, error) {
	switch instance {
	case GinServer:
		return newGinEngine(), nil
	default:
		return nil, errInvalidServerInstance
	}
}
