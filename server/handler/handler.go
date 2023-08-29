package handler

import (
	"server/server"
	"server/server/handler/common"
	"server/server/handler/routes"
)

func MakeServerHandlers() []server.HTTPHandler {
	return []server.HTTPHandler{
		&common.HealthHandler{},
		&routes.RoutesHandler{},
	}
}
