package common

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Store struct {
	Coordinate         Coordinate `json:"coordinate"`
	Name               string     `json:"name"`
	OperationHours     string     `json:"operation_hours"`
	HasParkingLot      bool       `json:"has_parking_lot"`
	DistanceFromOrigin int        `json:"distance_from_origin"`
}
