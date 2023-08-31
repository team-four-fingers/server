package routes

import "server/server/handler/common"

type RequestBody struct {
	Origin      common.Coordinate `json:"origin"`
	Destination common.Coordinate `json:"destination"`
	Waypoints   []common.Store    `json:"waypoints"`
}

type Route struct {
	Origin      common.Coordinate `json:"origin"`
	Destination common.Coordinate `json:"destination"`
	Waypoints   []common.Store    `json:"waypoints"`
}

type Response struct {
	Origin             common.Coordinate    `json:"origin"`
	Destination        common.Coordinate    `json:"destination"`
	Waypoints          []common.Store       `json:"waypoints"`
	CoordinatesInOrder []common.Coordinate  `json:"coordinates_in_order"`
	DurationInSeconds  int                  `json:"duration_in_minutes"`
	DistanceInMeters   int                  `json:"distance_in_meters"`
	Comparison         NavigationComparison `json:"comparison"`
}

type NavigationComparison struct {
	SavedTimeInMinutes int `json:"saved_time_in_minutes"`
	SavedGasCost       int `json:"saved_gas_cost"`
	Control            struct {
		DurationInMinutes int   `json:"duration_in_minutes"`
		DistanceInMeters  int   `json:"distance_in_meters"`
		GasCost           int   `json:"gas_cost"`
		Route             Route `json:"route"`
	} `json:"control"`
	Treatment struct {
		DurationInMinutes int     `json:"duration_in_minutes"`
		DistanceInMeters  int     `json:"distance_in_meters"`
		GasCost           int     `json:"gas_cost"`
		Routes            []Route `json:"routes"`
	} `json:"treatment"`
}
