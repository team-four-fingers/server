package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/team-four-fingers/kakao/mobility"
	kakaocommon "github.com/team-four-fingers/kakao/mobility/common"
	kakaowaypoints "github.com/team-four-fingers/kakao/mobility/waypoints"
	"net/http"
	"server/pkg/grouper"
	"server/server"
	"server/server/handler/common"
)

const (
	gastCostPerLiter         = 1735
	fuelEfficiencyLiterPerKm = 10
)

var _ server.HTTPHandler = (*Handler)(nil)

type Handler struct {
	mobilityCli mobility.Client

	groupFunc grouper.GroupFunc
}

func NewHandler(mobilityCli mobility.Client, groupFunc grouper.GroupFunc) *Handler {
	return &Handler{mobilityCli: mobilityCli, groupFunc: groupFunc}
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

		g := h.groupFunc()

		allWaypoints := lo.Map(waypoints, func(store common.Store, _ int) kakaocommon.Location {
			return kakaocommon.Location{
				X: store.Coordinate.X,
				Y: store.Coordinate.Y,
			}
		})
		var bestRoute kakaowaypoints.Route
		g.Go(func() error {
			navigationResp, err := h.mobilityCli.NavigateRouteThroughWaypoints(&kakaowaypoints.NavigateRouteThroughWaypointsRequest{
				Origin: kakaocommon.Location{
					X: origin.X,
					Y: origin.Y,
				},
				Destination: kakaocommon.Location{
					X: destination.X,
					Y: destination.Y,
				},
				Waypoints: allWaypoints,
				Priority:  lo.ToPtr(kakaocommon.PathPriority.Recommend),
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if len(navigationResp.Routes) == 0 {
				return echo.NewHTTPError(http.StatusBadRequest, errors.New("no route found"))
			}

			bestRoute = navigationResp.Routes[0]
			return nil
		})

		waypointsOnFirstVisit := []common.Store{}
		waypointsOnSecondVisit := waypoints
		if len(waypoints) != 1 {
			waypointsOnFirstVisit = waypoints[:len(waypoints)/2]
			waypointsOnSecondVisit = waypoints[len(waypoints)/2:]
		}

		log.Info("fine")

		var initialFirstRoute, initialSecondRoute kakaowaypoints.Route
		// 반만 거치고 집으로 가는 길
		g.Go(func() error {
			navigationResp, err := h.mobilityCli.NavigateRouteThroughWaypoints(&kakaowaypoints.NavigateRouteThroughWaypointsRequest{
				Origin: kakaocommon.Location{
					X: origin.X,
					Y: origin.Y,
				},
				Destination: kakaocommon.Location{
					X: destination.X,
					Y: destination.Y,
				},
				Waypoints: lo.Map(waypointsOnFirstVisit, func(store common.Store, _ int) kakaocommon.Location {
					return kakaocommon.Location{
						X: store.Coordinate.X,
						Y: store.Coordinate.Y,
					}
				}),
				Priority: lo.ToPtr(kakaocommon.PathPriority.Recommend),
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if len(navigationResp.Routes) == 0 {
				return echo.NewHTTPError(http.StatusBadRequest, errors.New("no route found"))
			}

			initialFirstRoute = navigationResp.Routes[0]
			return nil
		})

		// 나머지 반을 거치고 집에서 왔다 가는 길
		g.Go(func() error {
			navigationResp, err := h.mobilityCli.NavigateRouteThroughWaypoints(&kakaowaypoints.NavigateRouteThroughWaypointsRequest{
				Origin: kakaocommon.Location{
					X: destination.X,
					Y: destination.Y,
				},
				Destination: kakaocommon.Location{
					X: destination.X,
					Y: destination.Y,
				},
				Waypoints: lo.Map(waypointsOnSecondVisit, func(store common.Store, _ int) kakaocommon.Location {
					return kakaocommon.Location{
						X: store.Coordinate.X,
						Y: store.Coordinate.Y,
					}
				}),
				Priority: lo.ToPtr(kakaocommon.PathPriority.Recommend),
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if len(navigationResp.Routes) == 0 {
				return echo.NewHTTPError(http.StatusBadRequest, errors.New("no route found"))
			}

			initialSecondRoute = navigationResp.Routes[0]
			return nil
		})

		if err := g.Wait(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		treatmentTotalDuration := initialFirstRoute.Summary.Duration + initialSecondRoute.Summary.Duration
		treatmentTotalDistance := initialFirstRoute.Summary.Distance + initialSecondRoute.Summary.Distance

		treatmentGasCost := treatmentTotalDistance / fuelEfficiencyLiterPerKm * gastCostPerLiter
		controlGastCost := bestRoute.Summary.Distance / fuelEfficiencyLiterPerKm * gastCostPerLiter

		resp := &Response{
			Origin:             origin,
			Destination:        destination,
			Waypoints:          waypoints,
			CoordinatesInOrder: sectionsToCoordinatesInOrder(bestRoute.Sections),
			DurationInSeconds:  bestRoute.Summary.Duration,
			DistanceInMeters:   bestRoute.Summary.Distance,
			Comparison: NavigationComparison{
				SavedTimeInMinutes: treatmentTotalDuration - bestRoute.Summary.Duration,
				SavedGasCost:       treatmentGasCost - controlGastCost,
				Control: Control{
					DurationInMinutes: bestRoute.Summary.Duration,
					DistanceInMeters:  bestRoute.Summary.Distance,
					GasCost:           0,
					Route: Route{
						Origin:      origin,
						Destination: destination,
						Waypoints:   waypoints,
					},
				},
				Treatment: Treatment{
					DurationInMinutes: treatmentTotalDuration,
					DistanceInMeters:  treatmentTotalDistance,
					GasCost:           treatmentTotalDistance / fuelEfficiencyLiterPerKm * gastCostPerLiter,
					Routes: []Route{
						{
							Origin:      origin,
							Destination: destination,
							Waypoints:   waypointsOnFirstVisit,
						},
						{
							Origin:      destination,
							Destination: destination,
							Waypoints:   waypointsOnSecondVisit,
						},
					},
				},
			},
		}

		return c.JSON(http.StatusOK, resp)
	}
}
