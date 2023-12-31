package handler

import (
	"server/config"
	"server/pkg/grouper"
	"server/server"
	"server/server/handler/common"
	"server/server/handler/mockroutes"
	"server/server/handler/routes"
	"server/server/handler/search"
)

func MakeServerHandlers(cfg *config.Config) []server.HTTPHandler {
	return []server.HTTPHandler{
		&common.HealthHandler{},
		&mockroutes.Handler{},
		routes.NewHandler(cfg.MobilityCli(), grouper.NewPanicSafeGroup),
		search.NewHandler(cfg.LocalCli()),
	}
}
