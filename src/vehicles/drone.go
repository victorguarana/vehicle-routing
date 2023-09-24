package vehicles

import (
	"errors"

	"github.com/victorguarana/go-vehicle-route/src/gps"
)

var (
	ErrWithoutRange   = errors.New("vehicle does not support to move so far")
	ErrWithoutStorage = errors.New("vehicle does not have enough storage")
)

type IDrone interface {
	ivehicle
	IsFlying() bool
}

type drone struct {
	vehicle
	car             *car
	isFlying        bool
	totalStorage    float64
	remaningStorage float64
	totalRange      float64
	remaningRange   float64
}

func (d *drone) Move(destination *gps.Point) error {
	if d.actualPosition == nil || destination == nil {
		return ErrInvalidParams
	}

	if destination.PackageSize > d.remaningStorage {
		return ErrWithoutStorage
	}

	distance := gps.DistanceBetweenPoints(*d.actualPosition, *destination)
	if distance >= d.remaningRange {
		return ErrWithoutRange
	}

	d.remaningRange -= distance
	d.actualPosition = destination

	return nil
}

func (d *drone) Support(destinations ...gps.Point) bool {
	distance := 0.0
	packagesSize := 0.0
	position := *d.actualPosition

	for _, destination := range destinations {
		distance += gps.DistanceBetweenPoints(position, destination)
		packagesSize += destination.PackageSize
		position = destination
	}

	if distance > d.remaningRange {
		return false
	}

	if packagesSize > d.remaningStorage {
		return false
	}

	return true
}

func (d *drone) IsFlying() bool {
	return d.isFlying
}
