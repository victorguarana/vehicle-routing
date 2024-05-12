package vehicles

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
)

var defaultCarSpeed = 10.0

type ICar interface {
	ActualPoint() gps.Point
	Drones() []IDrone
	Move(destination routes.IMainStop)
	Name() string
	NewDrone(params DroneParams)
	Route() routes.IMainRoute
	Speed() float64
	Support(...gps.Point) bool
}

type car struct {
	drones []*drone
	name   string
	route  routes.IMainRoute
	speed  float64
}

type CarParams struct {
	Name          string
	StartingPoint routes.IMainStop
	RouteFactory  func(routes.IMainStop) routes.IMainRoute
}

func NewCar(params CarParams) ICar {
	return &car{
		drones: []*drone{},
		name:   params.Name,
		route:  params.RouteFactory(params.StartingPoint),
		speed:  defaultCarSpeed,
	}
}

func (c *car) ActualPoint() gps.Point {
	return c.route.Last().Point()
}

func (c *car) Drones() []IDrone {
	drones := []IDrone{}
	for _, d := range c.drones {
		drones = append(drones, d)
	}
	return drones
}

func (c *car) Move(destination routes.IMainStop) {
	c.route.Append(destination)
}

func (c *car) Name() string {
	return c.name
}

func (c *car) NewDrone(params DroneParams) {
	params.car = c
	d := newDrone(params)
	c.drones = append(c.drones, d)
}

func (c *car) Route() routes.IMainRoute {
	return c.route
}

func (c *car) Speed() float64 {
	return c.speed
}

func (c *car) Support(destination ...gps.Point) bool {
	return true
}
