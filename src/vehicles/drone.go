package vehicles

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
)

var defaultDroneRange = 150.0
var defaultDroneSpeed = 25.0
var defaultDroneStorage = 10.0

type IDrone interface {
	Flight() routes.ISubRoute
	IsFlying() bool
	Land(landingPoint routes.IMainStop)
	Move(destination routes.ISubStop)
	Name() string
	Speed() float64
	Support(...gps.Point) bool
}

type DroneParams struct {
	Name               string
	DroneFlightFactory func(routes.IMainStop) routes.ISubRoute
	car                *car
}

type drone struct {
	car             *car
	name            string
	speed           float64
	flight          routes.ISubRoute
	isFlying        bool
	remaningRange   float64
	remaningStorage float64
	totalRange      float64
	totalStorage    float64
	flightFactory   func(routes.IMainStop) routes.ISubRoute
}

func newDrone(params DroneParams) *drone {
	return &drone{
		car:             params.car,
		name:            params.Name,
		speed:           defaultDroneSpeed,
		remaningRange:   defaultDroneRange,
		remaningStorage: defaultDroneStorage,
		totalRange:      defaultDroneRange,
		totalStorage:    defaultDroneStorage,
		flightFactory:   params.DroneFlightFactory,
	}
}

func (d *drone) Flight() routes.ISubRoute {
	return d.flight
}

func (d *drone) IsFlying() bool {
	return d.isFlying
}

func (d *drone) Land(destination routes.IMainStop) {
	d.isFlying = false
	d.flight.Return(destination)
	d.flight = nil
	d.resetAttributes()
}

func (d *drone) Move(destination routes.ISubStop) {
	if d.isFlying {
		d.remaningRange -= gps.DistanceBetweenPoints(d.actualPoint(), destination.Point())
		d.flight.Append(destination)
		return
	}

	d.isFlying = true
	actualCarStop := d.car.route.Last()
	d.flight = d.flightFactory(actualCarStop)
	d.flight.Append(destination)
	d.remaningRange -= gps.DistanceBetweenPoints(actualCarStop.Point(), destination.Point())
}

func (d *drone) Name() string {
	return d.name
}

func (d *drone) Speed() float64 {
	return d.speed
}

func (d *drone) Support(destinations ...gps.Point) bool {
	distance := gps.DistanceBetweenPoints(append([]gps.Point{d.actualPoint()}, destinations...)...)
	packagesSize := 0.0
	for _, destination := range destinations {
		packagesSize += destination.PackageSize
	}
	if distance > d.remaningRange {
		return false
	}
	if packagesSize > d.remaningStorage {
		return false
	}
	return true
}

func (d *drone) actualPoint() gps.Point {
	if d.isFlying {
		return d.flight.Last().Point()
	}
	return d.car.ActualPoint()
}

func (d *drone) resetAttributes() {
	d.remaningRange = d.totalRange
	d.remaningStorage = d.totalStorage
}
