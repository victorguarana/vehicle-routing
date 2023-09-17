package vehicles

import (
	"errors"

	"github.com/victorguarana/go-vehicle-route/gps"
)

var (
	ErrSoFar         = errors.New("vehicle does not support to move so far")
	ErrInvalidParams = errors.New("invalid param")

	defaultStorage = 3000.0
	defaultRange   = 1000.0
	defaultSpeed   = 30
)

type IVehicle interface {
	Move(*gps.Point) error
}

type vehicle struct {
	totalStorage    float64
	remaningStorage float64
	totalRange      float64
	remaningRange   float64
	speed           int
	name            string
	actualPosition  *gps.Point
}

func NewVehicle(name string, startingPoint *gps.Point) IVehicle {
	return &vehicle{
		totalStorage:    defaultStorage,
		remaningStorage: defaultStorage,
		totalRange:      defaultRange,
		remaningRange:   defaultRange,
		speed:           defaultSpeed,
		name:            name,
		actualPosition:  startingPoint,
	}
}

func (v *vehicle) Move(destination *gps.Point) error {
	if v.actualPosition == nil || destination == nil {
		return ErrInvalidParams
	}

	distance := gps.DistanceBetweenPoints(*v.actualPosition, *destination)
	if distance >= v.remaningRange {
		return ErrSoFar
	}

	v.remaningRange -= distance
	v.actualPosition = destination

	return nil
}
