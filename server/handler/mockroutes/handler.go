package mockroutes

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"net/http"
	"server/server"
	"server/server/handler/common"
)

var _ server.HTTPHandler = (*Handler)(nil)

type Handler struct{}

func (h *Handler) Method() string {
	return http.MethodPost
}

func (h *Handler) Path() string {
	return "/mock-routes"
}

func (h *Handler) HandleFunc() func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := &RoutesResponse{
			//"name": "일진빌딩",
			//{         "x": "126.946362033068",         "y": "37.5404741779088"     }
			Origin: &common.Coordinate{
				X: 126.946362033068,
				Y: 37.5404741779088,
			},
			// {         "name": "카카오모빌리티", "x": "127.1101250888609",         "y": "37.39407843730005"     }
			Destination: &common.Coordinate{
				X: 127.1101250888609,
				Y: 37.39407843730005,
			},
			// [{             "name": "파크원타워1",             "x": "126.92716700037366",         "y": "37.5266641708316"         }]
			Waypoints: []*common.Coordinate{
				{
					X: 126.92716700037366,
					Y: 37.5266641708316,
				},
			},
			CoordinatesInOrder: roadsToCoordinatesInOrder(LoadExampleRoadData()),
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func roadsToCoordinatesInOrder(roads []*Road) []*common.Coordinate {
	if len(roads) == 0 {
		return nil
	}

	coordinatesInOrder := make([]*common.Coordinate, 0, lo.SumBy(roads, func(road *Road) int {
		return len(road.Vertexes) / 2
	}))
	for _, road := range roads {
		for i := 0; i < len(road.Vertexes); i += 2 {
			coordinatesInOrder = append(coordinatesInOrder, &common.Coordinate{
				X: road.Vertexes[i],
				Y: road.Vertexes[i+1],
			})
		}
	}

	return coordinatesInOrder
}
