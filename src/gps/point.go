package gps

import "math"

type Point struct {
	Latitude  float64
	Longitude float64
	name      string
}

// Euclidean distance between two points
func DistanceBetweenPoints(p1, p2 Point) float64 {
	return math.Abs(p1.Latitude-p2.Latitude) + math.Abs(p1.Longitude-p2.Longitude)
}
