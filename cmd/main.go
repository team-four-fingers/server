package main

import (
	"github.com/team-four-fingers/kakao/core"
	"github.com/team-four-fingers/kakao/mobility"
	"log"
	"net/http"
	"server/config"
	"server/server"
	"server/server/handler"
)

func main() {
	setting := config.NewSetting()

	cfg := config.NewConfig(
		setting,
		mobility.NewClient(core.NewClient(core.WithRestAPIKey(setting.KakaoRESTAPIKey))),
	)
	portNumber := cfg.Setting().PortNumber

	httpServer := server.NewHTTPServer(cfg)

	if err := httpServer.RegisterHTTPHandlers(
		handler.MakeServerHandlers(cfg)...,
	); err != nil {
		log.Fatal(err)
	}

	if err := httpServer.Start(portNumber); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
