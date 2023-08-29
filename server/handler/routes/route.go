package routes

//type WaypointsResponse struct {
//	TransId string `json:"trans_id"`
//	Routes  []struct {
//		ResultCode int    `json:"result_code"`
//		ResultMsg  string `json:"result_msg"`
//		Summary    struct {
//			Origin struct {
//				Name string  `json:"name"`
//				X    float64 `json:"x"`
//				Y    float64 `json:"y"`
//			} `json:"origin"`
//			Destination struct {
//				Name string  `json:"name"`
//				X    float64 `json:"x"`
//				Y    float64 `json:"y"`
//			} `json:"destination"`
//			Waypoints []struct {
//				Name string  `json:"name"`
//				X    float64 `json:"x"`
//				Y    float64 `json:"y"`
//			} `json:"waypoints"`
//			Priority string `json:"priority"`
//			Bound    struct {
//				MinX float64 `json:"min_x"`
//				MinY float64 `json:"min_y"`
//				MaxX float64 `json:"max_x"`
//				MaxY float64 `json:"max_y"`
//			} `json:"bound"`
//			Fare struct {
//				Taxi int `json:"taxi"`
//				Toll int `json:"toll"`
//			} `json:"fare"`
//			Distance int `json:"distance"`
//			Duration int `json:"duration"`
//		} `json:"summary"`
//		Sections []struct {
//			Distance int `json:"distance"`
//			Duration int `json:"duration"`
//			Bound    struct {
//				MinX float64 `json:"min_x"`
//				MinY float64 `json:"min_y"`
//				MaxX float64 `json:"max_x"`
//				MaxY float64 `json:"max_y"`
//			} `json:"bound"`
//			Roads []struct {
//				Name         string    `json:"name"`
//				Distance     int       `json:"distance"`
//				Duration     int       `json:"duration"`
//				TrafficSpeed int       `json:"traffic_speed"`
//				TrafficState int       `json:"traffic_state"`
//				Vertexes     []float64 `json:"vertexes"`
//			} `json:"roads"`
//			Guides []struct {
//				Name      string  `json:"name"`
//				X         float64 `json:"x"`
//				Y         float64 `json:"y"`
//				Distance  int     `json:"distance"`
//				Duration  int     `json:"duration"`
//				Type      int     `json:"type"`
//				Guidance  string  `json:"guidance"`
//				RoadIndex int     `json:"road_index"`
//			} `json:"guides"`
//		} `json:"sections"`
//	} `json:"routes"`
//}

type Road struct {
	Name         string    `json:"name"`
	Distance     int       `json:"distance"`
	Duration     int       `json:"duration"`
	TrafficSpeed int       `json:"traffic_speed"`
	TrafficState int       `json:"traffic_state"`
	Vertexes     []float64 `json:"vertexes"`
}

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type StraightPath struct {
	Origin      *Coordinate
	Destination *Coordinate
	Distance    int
	Duration    int
}

type WaypointToWaypoint struct {
	Distance    int
	Duration    int
	Origin      *Coordinate
	Destination *Coordinate
	paths       []*StraightPath
}

type Route struct {
	WaypointToWaypoint []*WaypointToWaypoint
	Distance           int
	Duration           int
}

type RoutesResponse struct {
	Origin             *Coordinate
	Destination        *Coordinate
	Waypoints          []*Coordinate
	CoordinatesInOrder []*Coordinate
}
