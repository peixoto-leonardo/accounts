package os

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForInterruptSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
