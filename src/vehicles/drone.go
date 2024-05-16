package vehicles

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
)

var defaultDroneRange = 150.0
var defaultDroneSpeed = 25.0
var defaultDroneStorage = 10.0

type IDrone interface {
	CanReach(...gps.Point) bool
	IsFlying() bool
	Land()
	Move(from, to gps.Point)
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

func (d *drone) CanReach(route ...gps.Point) bool {
	distance := gps.DistanceBetweenPoints(route...)
	return distance <= d.remaningRange
}

func (d *drone) IsFlying() bool {
	return d.isFlying
}

func (d *drone) Land() {
	d.isFlying = false
	d.resetAttributes()
}

func (d *drone) Move(from, to gps.Point) {
	d.isFlying = true
	d.remaningRange -= gps.DistanceBetweenPoints(from, to)
}

func (d *drone) Name() string {
	return d.name
}

func (d *drone) Speed() float64 {
	return d.speed
}

func (d *drone) Support(route ...gps.Point) bool {
	distance := gps.DistanceBetweenPoints(route...)
	packagesSize := 0.0
	for _, destination := range route {
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

func (d *drone) resetAttributes() {
	d.remaningRange = d.totalRange
	d.remaningStorage = d.totalStorage
}
