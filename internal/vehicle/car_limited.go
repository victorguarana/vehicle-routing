package vehicle

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
)

type CarParams struct {
	Efficiency    float64
	Speed         float64
	Storage       float64
	Range         float64
	Name          string
	StartingPoint gps.Point
}

type carLimtited struct {
	car
	remaningRange   float64
	remaningStorage float64
	totalRange      float64
	totalStorage    float64
}

func NewCarWithParams(params CarParams) ICar {
	return &carLimtited{
		car: car{
			actualPoint: params.StartingPoint,
			drones:      []*drone{},
			efficiency:  params.Efficiency,
			name:        params.Name,
			speed:       params.Speed,
		},
		totalStorage:    params.Storage,
		remaningStorage: params.Storage,
		totalRange:      params.Range,
		remaningRange:   params.Range,
	}
}

func (c *carLimtited) Move(destination gps.Point) {
	c.remaningRange -= gps.ManhattanDistanceBetweenPoints(c.actualPoint, destination)
	c.remaningStorage -= destination.PackageSize
	c.actualPoint = destination
	c.moveDockedDrones(destination)
}

func (c *carLimtited) Storage() float64 {
	return c.totalStorage
}

func (c *carLimtited) Support(destinations ...gps.Point) bool {
	distance := gps.ManhattanDistanceBetweenPoints(destinations...)
	packagesSize := 0.0
	for _, destination := range destinations {
		packagesSize += destination.PackageSize
	}
	if distance > c.remaningRange {
		return false
	}
	if packagesSize > c.remaningStorage {
		return false
	}
	return true
}
