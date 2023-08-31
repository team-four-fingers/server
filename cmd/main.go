package main

import (
	"github.com/team-four-fingers/kakao/core"
	"github.com/team-four-fingers/kakao/local"
	"github.com/team-four-fingers/kakao/mobility"
	"log"
	"net/http"
	"server/config"
	"server/server"
	"server/server/handler"
)

func main() {
	setting := config.NewSetting()

	makeCoreCli := func() *core.Client {
		return core.NewClient(core.WithRestAPIKey(setting.KakaoRESTAPIKey))
	}

	cfg := config.NewConfig(
		setting,
		mobility.NewClient(makeCoreCli()),
		local.NewClient(makeCoreCli()),
	)

	httpServer := server.NewHTTPServer(cfg)

	if err := httpServer.RegisterHTTPHandlers(
		handler.MakeServerHandlers(cfg)...,
	); err != nil {
		log.Fatal(err)
	}

	if err := httpServer.Start(cfg.Setting().PortNumber); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
