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
	defaultSpeed   = 30.0
)

type ivehicle interface {
	Move(*gps.Point) error
	Reachable(gps.Point) bool
}

type vehicle struct {
	speed          float64
	name           string
	actualPosition *gps.Point
}
