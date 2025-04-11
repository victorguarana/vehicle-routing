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

type CarLimited struct {
	*car
	remaningRange   float64
	remaningStorage float64
	totalRange      float64
	totalStorage    float64
}

func NewCarLimited(params CarParams) *CarLimited {
	return &CarLimited{
		car: &car{
			actualPoint: params.StartingPoint,
			efficiency:  params.Efficiency,
			name:        params.Name,
			speed:       params.Speed,
			drones:      []*drone{},
		},
		totalStorage:    params.Storage,
		remaningStorage: params.Storage,
		totalRange:      params.Range,
		remaningRange:   params.Range,
	}
}

func (c *CarLimited) Clone() *CarLimited {
	clonedDrones := make([]*drone, len(c.drones))
	for i, d := range c.drones {
		clonedDrones[i] = d.clone()
	}

	return &CarLimited{
		car: &car{
			actualPoint: c.actualPoint,
			efficiency:  c.efficiency,
			name:        c.name,
			speed:       c.speed,
			drones:      clonedDrones,
		},
		totalStorage:    c.totalStorage,
		remaningStorage: c.remaningStorage,
		totalRange:      c.totalRange,
		remaningRange:   c.remaningRange,
	}
}

func (c *CarLimited) Move(destination gps.Point) {
	c.remaningRange -= gps.ManhattanDistanceBetweenPoints(c.actualPoint, destination)
	c.remaningStorage -= destination.PackageSize
	c.actualPoint = destination
	c.moveDockedDrones(destination)
}

func (c *CarLimited) Range() float64 {
	return c.totalRange
}

func (c *CarLimited) Storage() float64 {
	return c.totalStorage
}

func (c *CarLimited) Support(destinations ...gps.Point) bool {
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
