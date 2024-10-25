package main

import (
	"os"
	"os/signal"
	"syscall"

	net_api "github.com/yusubond/netgin/http"
)

func main() {
	// create api server
	api := net_api.NewHttpServer()
	err := api.Init()
	if err != nil {
		panic(err)
	}
	api.Start()

	//
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGTERM)

	<-sigChan

	api.Infof("received shutdown signal")
	api.Stop()

	// quit
	os.Exit(0)
}
