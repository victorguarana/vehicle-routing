package vehicles

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
)

const CarSpeed = 10.0

type ICar interface {
	ActualPoint() gps.Point
	Drones() []IDrone
	Move(destination gps.Point)
	Name() string
	NewDrone(name string)
	Speed() float64
	Support(...gps.Point) bool
}

type car struct {
	actualPoint gps.Point
	drones      []*drone
	name        string
	speed       float64
}

func NewCar(name string, startingPoint gps.Point) ICar {
	return &car{
		actualPoint: startingPoint,
		drones:      []*drone{},
		name:        name,
		speed:       CarSpeed,
	}
}

func (c *car) ActualPoint() gps.Point {
	return c.actualPoint
}

func (c *car) Drones() []IDrone {
	drones := []IDrone{}
	for _, d := range c.drones {
		drones = append(drones, d)
	}
	return drones
}

func (c *car) Move(destination gps.Point) {
	c.actualPoint = destination
	c.moveDockedDrones(destination)
}

func (c *car) Name() string {
	return c.name
}

func (c *car) NewDrone(name string) {
	d := newDrone(name, c)
	c.drones = append(c.drones, d)
}

func (c *car) Speed() float64 {
	return c.speed
}

func (c *car) Support(destination ...gps.Point) bool {
	return true
}

func (c *car) moveDockedDrones(destination gps.Point) {
	for _, d := range c.drones {
		if !d.isFlying {
			d.actualPoint = destination
		}
	}
}
