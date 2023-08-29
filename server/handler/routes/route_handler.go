package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/server"
)

var _ server.HTTPHandler = (*RoutesHandler)(nil)

type RoutesHandler struct{}

func (h *RoutesHandler) Method() string {
	return http.MethodPost
}

func (h *RoutesHandler) Path() string {
	return "/routes"
}

func (h *RoutesHandler) HandleFunc() func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := &RoutesResponse{
			Origin: &Coordinate{
				X: 37.5403831,
				Y: 126.9463611,
			},
			Destination: &Coordinate{
				X: 37.5513831,
				Y: 126.9573611,
			},
			Waypoints: []*Coordinate{
				{
					X: 37.5413831,
					Y: 126.9473611,
				},
			},
			CoordinatesInOrder: []*Coordinate{
				{
					X: 37.5403831,
					Y: 126.9463611,
				},
				{
					X: 37.5413831,
					Y: 126.9473611,
				},
				{
					X: 37.5513831,
					Y: 126.9573611,
				},
			},
		}

		return c.JSON(http.StatusOK, resp)
	}
}
