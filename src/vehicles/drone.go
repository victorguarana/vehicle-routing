package vehicles

import (
	"errors"

	"github.com/victorguarana/go-vehicle-route/src/gps"
)

var (
	ErrWithoutRange = errors.New("vehicle does not support to move so far")
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

	distance := gps.DistanceBetweenPoints(*d.actualPosition, *destination)
	if distance >= d.remaningRange {
		return ErrWithoutRange
	}

	d.remaningRange -= distance
	d.actualPosition = destination

	return nil
}

func (d *drone) Reachable(destinations ...gps.Point) bool {
	distance := 0.0
	position := *d.actualPosition

	for _, destination := range destinations {
		distance += gps.DistanceBetweenPoints(position, destination)
		position = destination
	}

	return distance <= d.remaningRange
}

func (d *drone) IsFlying() bool {
	return d.isFlying
}
