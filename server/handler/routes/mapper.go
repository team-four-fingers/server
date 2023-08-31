package routes

import (
	"github.com/samber/lo"
	"github.com/team-four-fingers/kakao/mobility/waypoints"
	"server/server/handler/common"
)

func roadsToCoordinatesInOrder(roads []waypoints.Road) []common.Coordinate {
	if len(roads) == 0 {
		return nil
	}

	coordinatesInOrder := make([]common.Coordinate, 0, lo.SumBy(roads, func(road waypoints.Road) int {
		return len(road.Vertexes) / 2
	}))
	for _, road := range roads {
		for i := 0; i < len(road.Vertexes); i += 2 {
			coordinatesInOrder = append(coordinatesInOrder, common.Coordinate{
				X: road.Vertexes[i],
				Y: road.Vertexes[i+1],
			})
		}
	}

	return coordinatesInOrder
}

func sectionsToCoordinatesInOrder(sections []waypoints.Section) []common.Coordinate {
	return lo.Flatten(lo.Map(sections, func(section waypoints.Section, _ int) []common.Coordinate {
		return roadsToCoordinatesInOrder(section.Roads)
	}))
}
