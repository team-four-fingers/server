package main

import (
	"log"
	"net/http"
	"server/server"
	"server/server/handler"
)

func main() {
	cfg := NewConfig()
	portNumber := cfg.PortNumber()

	httpServer := server.NewHTTPServer()

	if err := httpServer.RegisterHTTPHandlers(
		&handler.HealthHandler{},
	); err != nil {
		log.Fatal(err)
	}

	if err := httpServer.Start(portNumber); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
