package vehicle

import (
	"math"

	"github.com/victorguarana/vehicle-routing/internal/gps"
)

type CarUnlimited struct {
	*car
}

func NewCarUnlimited(name string, startingPoint gps.Point) *CarUnlimited {
	return &CarUnlimited{
		car: &car{
			actualPoint: startingPoint,
			efficiency:  CarDefaultEfficiency,
			name:        name,
			speed:       CarDefaultSpeed,
			drones:      []*drone{},
		},
	}
}

func (c *CarUnlimited) Clone() ICar {
	clonedDrones := make([]*drone, len(c.drones))
	for i, d := range c.drones {
		clonedDrones[i] = d.clone()
	}

	return &CarUnlimited{
		car: &car{
			actualPoint: c.actualPoint,
			efficiency:  c.efficiency,
			name:        c.name,
			speed:       c.speed,
			drones:      clonedDrones,
		},
	}
}

func (c *CarUnlimited) Move(destination gps.Point) {
	c.actualPoint = destination
	c.moveDockedDrones(destination)
}

func (*CarUnlimited) Range() float64 {
	return math.Inf(1)
}

func (*CarUnlimited) Storage() float64 {
	return math.Inf(1)
}

func (*CarUnlimited) Support(destination ...gps.Point) bool {
	return true
}
