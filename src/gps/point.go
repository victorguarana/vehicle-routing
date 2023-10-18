package gps

import (
	"fmt"
	"math"
)

type Point struct {
	Latitude    float64
	Longitude   float64
	PackageSize float64
}

func (p Point) IsClient() bool {
	return p.PackageSize != 0
}

func (p Point) IsDeposit() bool {
	return p.PackageSize == 0
}

func (p Point) String() string {
	return fmt.Sprintf("Lat: %f, Long: %f, PackageSize: %f\n", p.Latitude, p.Longitude, p.PackageSize)
}

// Euclidean distance between two points
func DistanceBetweenPoints(points ...Point) float64 {
	var totalDistance float64
	for i := 0; i < len(points)-1; i++ {
		totalDistance += math.Abs(points[i].Latitude-points[i+1].Latitude) + math.Abs(points[i].Longitude-points[i+1].Longitude)
	}
	return totalDistance
}
