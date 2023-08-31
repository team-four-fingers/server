package main

import (
	"github.com/team-four-fingers/kakao/core"
	"github.com/team-four-fingers/kakao/mobility"
	"log"
	"net/http"
	"server/server"
	"server/server/handler"
)

func main() {
	setting := NewSetting()

	cfg := NewConfig(
		setting,
		mobility.NewClient(core.NewClient(core.WithRestAPIKey(setting.KakaoRESTAPIKey))),
	)
	portNumber := cfg.setting.PortNumber

	httpServer := server.NewHTTPServer()

	if err := httpServer.RegisterHTTPHandlers(
		handler.MakeServerHandlers()...,
	); err != nil {
		log.Fatal(err)
	}

	if err := httpServer.Start(portNumber); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
