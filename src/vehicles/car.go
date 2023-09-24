package vehicles

import "github.com/victorguarana/go-vehicle-route/src/gps"

var (
	defaultStorage = 3000.0
	defaultRange   = 1000.0
	defaultSpeed   = 30.0
)

type ICar interface {
	ivehicle
	NewDrone(string)
}

type car struct {
	vehicle
	drones []drone
}

func NewCar(name string, startingPoint *gps.Point) ICar {
	c := car{
		vehicle: vehicle{
			actualPosition: startingPoint,
			name:           name,
			speed:          defaultSpeed,
		},
		drones: []drone{},
	}

	return &c
}

func (c *car) NewDrone(name string) {
	d := drone{
		totalStorage:    defaultStorage,
		remaningStorage: defaultStorage,
		totalRange:      defaultRange,
		remaningRange:   defaultRange,
		vehicle: vehicle{
			speed:          defaultSpeed,
			name:           name,
			actualPosition: c.actualPosition,
		},
		car: c,
	}
	c.drones = append(c.drones, d)
}

func (c *car) Move(destination *gps.Point) error {
	if c.actualPosition == nil || destination == nil {
		return ErrInvalidParams
	}

	c.actualPosition = destination

	return nil
}

func (c *car) Support(destination ...gps.Point) bool {
	return true
}
