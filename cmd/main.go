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

	if err := core.InitializeSDK(setting.KakaoRESTAPIKey); err != nil {
		log.Fatal(err)
	}

	mobilityCli, err := mobility.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	localCli, err := local.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewConfig(
		setting,
		mobilityCli,
		localCli,
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
