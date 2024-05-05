package vehicles

import "github.com/victorguarana/go-vehicle-route/src/gps"

var (
	defaultCarSpeed = 10.0
)

type ICar interface {
	ActualPosition() gps.Point
	Drones() []IDrone
	Move(gps.Point)
	Name() string
	NewDrone(string)
	Speed() float64
	Support(...gps.Point) bool
}

type car struct {
	speed          float64
	name           string
	actualPosition gps.Point
	drones         []*drone
}

func NewCar(name string, startingPoint gps.Point) ICar {
	c := car{
		actualPosition: startingPoint,
		name:           name,
		speed:          defaultCarSpeed,
		drones:         []*drone{},
	}

	return &c
}

func (c *car) ActualPosition() gps.Point {
	return c.actualPosition
}

func (c *car) Name() string {
	return c.name
}

func (c *car) Speed() float64 {
	return c.speed
}

func (c *car) NewDrone(name string) {
	d := newDrone(name, c)

	c.drones = append(c.drones, d)
}

func (c *car) Move(destination gps.Point) {
	c.actualPosition = destination
	c.moveDockedDrones()
}

func (c *car) Support(destination ...gps.Point) bool {
	return true
}

func (c *car) Drones() []IDrone {
	drones := []IDrone{}
	for _, d := range c.drones {
		drones = append(drones, d)
	}
	return drones
}

func (c *car) moveDockedDrones() {
	for _, d := range c.drones {
		if !d.isFlying {
			d.actualPosition = c.actualPosition
		}
	}
}
