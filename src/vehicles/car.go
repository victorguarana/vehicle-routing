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
	actualPosition gps.Point
	drones         []*drone
	name           string
	speed          float64
}

func NewCar(name string, startingPoint gps.Point) ICar {
	return &car{
		actualPosition: startingPoint,
		drones:         []*drone{},
		name:           name,
		speed:          defaultCarSpeed,
	}
}

func (c *car) ActualPosition() gps.Point {
	return c.actualPosition
}

func (c *car) Drones() []IDrone {
	drones := []IDrone{}
	for _, d := range c.drones {
		drones = append(drones, d)
	}
	return drones
}

func (c *car) Move(destination gps.Point) {
	c.actualPosition = destination
	c.moveDockedDrones()
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

func (c *car) moveDockedDrones() {
	for _, d := range c.drones {
		if !d.isFlying {
			d.actualPosition = c.actualPosition
		}
	}
}
