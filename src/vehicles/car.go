package vehicles

import "github.com/victorguarana/go-vehicle-route/src/gps"

var (
	defaultCarSpeed = 10.0
)

type ICar interface {
	ivehicle

	NewDrone(string)

	Drones() []IDrone
}

type car struct {
	vehicle
	drones []*drone
}

func NewCar(name string, startingPoint *gps.Point) ICar {
	c := car{
		vehicle: vehicle{
			actualPosition: startingPoint,
			name:           name,
			speed:          defaultCarSpeed,
		},
		drones: []*drone{},
	}

	return &c
}

func (c *car) NewDrone(name string) {
	d := newDrone(name, c)

	c.drones = append(c.drones, d)
}

func (c *car) Move(destination *gps.Point) error {
	if c.actualPosition == nil || destination == nil {
		return ErrInvalidParams
	}

	c.actualPosition = destination

	return nil
}

func (c *car) Support(destination ...*gps.Point) bool {
	return true
}

func (c *car) Drones() []IDrone {
	drones := []IDrone{}
	for _, d := range c.drones {
		drones = append(drones, d)
	}
	return drones
}
