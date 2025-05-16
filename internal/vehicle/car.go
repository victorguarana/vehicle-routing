package vehicle

import "github.com/victorguarana/vehicle-routing/internal/gps"

const CarDefaultSpeed = 10.0
const CarDefaultEfficiency = 5.0

//go:generate mockgen -source=car.go -destination=mock/carmock.go
type ICar interface {
	ActualPoint() gps.Point
	Clone() ICar
	Drones() []IDrone
	Efficiency() float64
	Move(destination gps.Point)
	Name() string
	NewDefaultDrone(name string)
	NewDroneWithParams(params DroneParams)
	Range() float64
	Speed() float64
	Storage() float64
	Support(...gps.Point) bool
}

type car struct {
	actualPoint gps.Point
	efficiency  float64
	name        string
	speed       float64
	drones      []*drone
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

func (c *car) Efficiency() float64 {
	return c.efficiency
}

func (c *car) NewDefaultDrone(name string) {
	d := newDefaultDrone(name)
	c.drones = append(c.drones, d)
}

func (c *car) NewDroneWithParams(params DroneParams) {
	d := newDroneWithParams(params)
	c.drones = append(c.drones, d)
}

func (c *car) Name() string {
	return c.name
}

func (c *car) Speed() float64 {
	return c.speed
}

func (c *car) moveDockedDrones(destination gps.Point) {
	for _, d := range c.drones {
		if !d.isFlying {
			d.actualPoint = destination
		}
	}
}
