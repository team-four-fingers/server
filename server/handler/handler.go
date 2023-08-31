package handler

import (
	"server/server"
	"server/server/handler/common"
	"server/server/handler/mockroutes"
)

func MakeServerHandlers() []server.HTTPHandler {
	return []server.HTTPHandler{
		&common.HealthHandler{},
		&mockroutes.Handler{},
	}
}
