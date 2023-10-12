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
func DistanceBetweenPoints(p1, p2 Point) float64 {
	return math.Abs(p1.Latitude-p2.Latitude) + math.Abs(p1.Longitude-p2.Longitude)
}
