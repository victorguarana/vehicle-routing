package gps

import (
	"github.com/victorguarana/vehicle-routing/internal/slc"
)

type Neighborhood map[Point][]Point

func MapNeighborhood(points []Point, maxDistance float64) Neighborhood {
	neighborhood := make(Neighborhood, len(points))
	for _, point := range points {
		neighborhood[point] = []Point{}
		for _, neighbor := range points {
			if point == neighbor {
				continue
			}
			if ManhattanDistanceBetweenPoints(point, neighbor) <= maxDistance {
				neighborhood[point] = append(neighborhood[point], neighbor)
			}
		}
	}
	return neighborhood
}

func RemovePointFromNearbyMap(pointToBeRemoved Point, neighborhood Neighborhood) {
	delete(neighborhood, pointToBeRemoved)
	for point, neighbors := range neighborhood {
		neighborhood[point] = slc.RemoveElement(neighbors, pointToBeRemoved)
	}
}

func ClosestPointWithMostNeighbors(initialPoint Point, neighborhood Neighborhood) Point {
	var bestPoint Point
	maxNeighborhoodSize := -1
	for point, neighborhood := range neighborhood {
		neighborhoodSize := len(neighborhood)
		if neighborhoodSize > maxNeighborhoodSize {
			bestPoint = point
			maxNeighborhoodSize = neighborhoodSize
		}
		if neighborhoodSize == maxNeighborhoodSize {
			if ManhattanDistanceBetweenPoints(initialPoint, point) < ManhattanDistanceBetweenPoints(initialPoint, bestPoint) {
				bestPoint = point
			}
		}
	}
	return bestPoint
}
