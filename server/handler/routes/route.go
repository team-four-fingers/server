package routes

import "server/server/handler/common"

type Road struct {
	Name         string    `json:"name"`
	Distance     int       `json:"distance"`
	Duration     int       `json:"duration"`
	TrafficSpeed int       `json:"traffic_speed"`
	TrafficState int       `json:"traffic_state"`
	Vertexes     []float64 `json:"vertexes"`
}

type StraightPath struct {
	Origin      *common.Coordinate
	Destination *common.Coordinate
	Distance    int
	Duration    int
}

type WaypointToWaypoint struct {
	Distance    int
	Duration    int
	Origin      *common.Coordinate
	Destination *common.Coordinate
	paths       []*StraightPath
}

type Route struct {
	WaypointToWaypoint []*WaypointToWaypoint
	Distance           int
	Duration           int
}

type RoutesResponse struct {
	Origin             *common.Coordinate
	Destination        *common.Coordinate
	Waypoints          []*common.Coordinate
	CoordinatesInOrder []*common.Coordinate
}
