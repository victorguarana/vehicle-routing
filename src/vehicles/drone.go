package vehicles

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
)

var (
	defaultDroneRange   = 150.0
	defaultDroneSpeed   = 25.0
	defaultDroneStorage = 10.0
)

type IDrone interface {
	ActualPosition() gps.Point
	Land(gps.Point)
	Move(gps.Point)
	Name() string
	Speed() float64
	Support(...gps.Point) bool
}

type drone struct {
	actualPosition  gps.Point
	car             *car
	name            string
	speed           float64
	isFlying        bool
	remaningRange   float64
	remaningStorage float64
	totalRange      float64
	totalStorage    float64
}

func newDrone(name string, car *car) *drone {
	return &drone{
		actualPosition:  car.actualPosition,
		car:             car,
		name:            name,
		speed:           defaultDroneSpeed,
		remaningRange:   defaultDroneRange,
		remaningStorage: defaultDroneStorage,
		totalRange:      defaultDroneRange,
		totalStorage:    defaultDroneStorage,
	}
}

func (d *drone) ActualPosition() gps.Point {
	return d.actualPosition
}

func (d *drone) Land(destination gps.Point) {
	d.Move(destination)
	d.isFlying = false
	d.remaningRange = d.totalRange
	d.remaningStorage = d.totalStorage
}

func (d *drone) Move(destination gps.Point) {
	d.isFlying = true
	d.remaningRange -= gps.DistanceBetweenPoints(d.actualPosition, destination)
	d.actualPosition = destination
}

func (d *drone) Name() string {
	return d.name
}

func (d *drone) Speed() float64 {
	return d.speed
}

func (d *drone) Support(destinations ...gps.Point) bool {
	distance := gps.DistanceBetweenPoints(append([]gps.Point{d.actualPosition}, destinations...)...)
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
