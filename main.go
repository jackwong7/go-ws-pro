package main

import (
	"os"
	"os/signal"
	"syscall"
	"ws-pro/http_server"
)

func main() {

	go http_server.HttpServerRun()
	go http_server.WebsocketServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	http_server.HttpServerStop()
	http_server.WebsocketServerStop()
}
