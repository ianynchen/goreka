package util

import (
	"os"
	"os/signal"
	"syscall"
)

func SigTermHandler() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		eureka.Unregister()
		os.Exit(1)
	}()
}
