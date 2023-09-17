package vehicles

import "github.com/victorguarana/go-vehicle-route/gps"

type ICar interface {
	ivehicle
}

type car struct {
	vehicle
}

func NewCar(name string, startingPoint *gps.Point) ICar {
	c := car{
		vehicle{
			actualPosition: startingPoint,
			name:           name,
			speed:          defaultSpeed,
		},
	}

	return &c
}

func (c *car) Move(destination *gps.Point) error {
	if c.actualPosition == nil || destination == nil {
		return ErrInvalidParams
	}

	c.actualPosition = destination

	return nil
}

func (c *car) Reachable(destination gps.Point) bool {
	return true
}
