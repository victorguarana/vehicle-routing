package vehicles

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
)

var defaultCarSpeed = 10.0

type ICar interface {
	Drones() []IDrone
	Name() string
	NewDrone(params DroneParams)
	Speed() float64
	Support(...gps.Point) bool
}

type car struct {
	drones []*drone
	name   string
	speed  float64
}

func NewCar(name string) ICar {
	return &car{
		drones: []*drone{},
		name:   name,
		speed:  defaultCarSpeed,
	}
}

func (c *car) Drones() []IDrone {
	drones := []IDrone{}
	for _, d := range c.drones {
		drones = append(drones, d)
	}
	return drones
}

func (c *car) Name() string {
	return c.name
}

func (c *car) NewDrone(params DroneParams) {
	params.car = c
	d := newDrone(params)
	c.drones = append(c.drones, d)
}

func (c *car) Speed() float64 {
	return c.speed
}

func (c *car) Support(destination ...gps.Point) bool {
	return true
}
