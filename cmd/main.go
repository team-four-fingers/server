package main

import (
	"log"
	"net/http"
	"server/server"
	"server/server/handler"
	"server/server/handler/routes"
)

func main() {
	cfg := NewConfig()
	portNumber := cfg.PortNumber()

	httpServer := server.NewHTTPServer()

	if err := httpServer.RegisterHTTPHandlers(
		&handler.HealthHandler{},
		&routes.RoutesHandler{},
	); err != nil {
		log.Fatal(err)
	}

	if err := httpServer.Start(portNumber); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
