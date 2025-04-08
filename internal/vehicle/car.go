package vehicle

import (
	"math"

	"github.com/victorguarana/vehicle-routing/internal/gps"
)

const CarSpeed = 10.0
const CarEfficiency = 5.0

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
	drones      []*drone
	efficiency  float64
	name        string
	speed       float64
}

func NewDefaultCar(name string, startingPoint gps.Point) ICar {
	return &car{
		actualPoint: startingPoint,
		drones:      []*drone{},
		efficiency:  CarEfficiency,
		name:        name,
		speed:       CarSpeed,
	}
}

func (c *car) ActualPoint() gps.Point {
	return c.actualPoint
}

func (c *car) Clone() ICar {
	clonedDrones := make([]*drone, len(c.drones))
	for i, d := range c.drones {
		clonedDrones[i] = d.clone()
	}

	return &car{
		actualPoint: c.actualPoint,
		drones:      clonedDrones,
		efficiency:  c.efficiency,
		name:        c.name,
		speed:       c.speed,
	}
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

func (c *car) Move(destination gps.Point) {
	c.actualPoint = destination
	c.moveDockedDrones(destination)
}

func (c *car) Name() string {
	return c.name
}

func (c *car) NewDefaultDrone(name string) {
	d := newDefaultDrone(name)
	c.drones = append(c.drones, d)
}

func (c *car) NewDroneWithParams(params DroneParams) {
	d := newDroneWithParams(params)
	c.drones = append(c.drones, d)
}

func (*car) Range() float64 {
	return math.Inf(1)
}

func (c *car) Speed() float64 {
	return c.speed
}

func (*car) Storage() float64 {
	return math.Inf(1)
}

func (*car) Support(destination ...gps.Point) bool {
	return true
}

func (c *car) moveDockedDrones(destination gps.Point) {
	for _, d := range c.drones {
		if !d.isFlying {
			d.actualPoint = destination
		}
	}
}
