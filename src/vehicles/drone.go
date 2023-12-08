package vehicles

import (
	"errors"

	"github.com/victorguarana/go-vehicle-route/src/gps"
)

var (
	ErrWithoutRange   = errors.New("vehicle does not support to move so far")
	ErrWithoutStorage = errors.New("vehicle does not have enough storage")
)

var (
	defaultDroneStorage = 10.0
	defaultDroneRange   = 150.0
	defaultDroneSpeed   = 25.0
)

type IDrone interface {
	ivehicle
}

type drone struct {
	vehicle
	car *car

	totalStorage    float64
	remaningStorage float64
	totalRange      float64
	remaningRange   float64
}

func newDrone(name string, car *car) *drone {
	d := drone{
		totalStorage:    defaultDroneStorage,
		remaningStorage: defaultDroneStorage,
		totalRange:      defaultDroneRange,
		remaningRange:   defaultDroneRange,
		vehicle: vehicle{
			speed:          defaultDroneSpeed,
			name:           name,
			actualPosition: car.actualPosition,
		},
		car: car,
	}

	return &d
}

func (d *drone) Move(destination *gps.Point) error {
	if d.actualPosition == nil || destination == nil {
		return ErrInvalidParams
	}

	if destination.PackageSize > d.remaningStorage {
		return ErrWithoutStorage
	}

	distance := gps.DistanceBetweenPoints(d.actualPosition, destination)
	if distance >= d.remaningRange {
		return ErrWithoutRange
	}

	d.remaningRange -= distance
	d.actualPosition = destination

	return nil
}

func (d *drone) Support(destinations ...*gps.Point) bool {
	distance := gps.DistanceBetweenPoints(append([]*gps.Point{d.actualPosition}, destinations...)...)
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
