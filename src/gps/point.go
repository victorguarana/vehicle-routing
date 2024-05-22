package gps

import (
	"math"
)

type Point struct {
	Latitude    float64
	Longitude   float64
	PackageSize float64
	Name        string
}

// Manhattan distance between two points
func DistanceBetweenPoints(points ...Point) float64 {
	var totalDistance float64
	for i := 0; i < len(points)-1; i++ {
		totalDistance += math.Abs(points[i].Latitude-points[i+1].Latitude) + math.Abs(points[i].Longitude-points[i+1].Longitude)
	}
	return totalDistance
}

// Returns the closest candidate point to the origin point
func ClosestPoint(originPoint Point, candidatePoints []Point) Point {
	var closestPoint Point
	closestDistance := -1.0
	for _, candidatePoint := range candidatePoints {
		if closestDistance < 0 || DistanceBetweenPoints(originPoint, candidatePoint) < closestDistance {
			closestPoint = candidatePoint
			closestDistance = DistanceBetweenPoints(originPoint, closestPoint)
		}
	}
	return closestPoint
}

// Returns the distance that is added when passing through the middle point
func AdditionalDistancePassingThrough(from, through, to Point) float64 {
	return DistanceBetweenPoints(from, through, to) - DistanceBetweenPoints(from, to)
}
