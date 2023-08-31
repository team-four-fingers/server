package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/team-four-fingers/kakao/mobility"
	kakaocommon "github.com/team-four-fingers/kakao/mobility/common"
	kakaowaypoints "github.com/team-four-fingers/kakao/mobility/waypoints"
	"net/http"
	"server/server"
	"server/server/handler/common"
)

var _ server.HTTPHandler = (*Handler)(nil)

type Handler struct {
	mobilityCli mobility.Client
}

func NewHandler(mobilityCli mobility.Client) *Handler {
	return &Handler{mobilityCli: mobilityCli}
}

func (h *Handler) Method() string {
	return http.MethodPost
}

func (h *Handler) Path() string {
	return "/routes"
}

func (h *Handler) HandleFunc() func(c echo.Context) error {
	return func(c echo.Context) error {
		body := &RequestBody{}
		if err := c.Bind(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var (
			origin      = body.Origin
			destination = body.Destination
			waypoints   = body.Waypoints
		)

		navigationResp, err := h.mobilityCli.NavigateRouteThroughWaypoints(&kakaowaypoints.NavigateRouteThroughWaypointsRequest{
			Origin: kakaocommon.Location{
				X: origin.X,
				Y: origin.Y,
			},
			Destination: kakaocommon.Location{
				X: destination.X,
				Y: destination.Y,
			},
			Waypoints: lo.Map(waypoints, func(store common.Store, _ int) kakaocommon.Location {
				return kakaocommon.Location{
					X: store.Coordinate.X,
					Y: store.Coordinate.Y,
				}
			}),
			Priority: lo.ToPtr(kakaocommon.PathPriority.Recommend),
			Summary:  lo.ToPtr(true),
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if len(navigationResp.Routes) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("no route found"))
		}

		bestRoute := navigationResp.Routes[0]

		resp := &Response{
			Origin:             origin,
			Destination:        destination,
			Waypoints:          waypoints,
			CoordinatesInOrder: sectionsToCoordinatesInOrder(bestRoute.Sections),
			DurationInSeconds:  bestRoute.Summary.Duration,
			DistanceInMeters:   bestRoute.Summary.Distance,
			Comparison:         NavigationComparison{},
		}

		return c.JSON(http.StatusOK, resp)
	}
}
