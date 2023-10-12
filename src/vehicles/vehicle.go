package vehicles

import (
	"errors"

	"github.com/victorguarana/go-vehicle-route/src/gps"
)

var (
	ErrInvalidParams = errors.New("invalid param")
)

type ivehicle interface {
	ActualPosition() *gps.Point
	Move(*gps.Point) error
	Support(...gps.Point) bool
}

type vehicle struct {
	speed          float64
	name           string
	actualPosition *gps.Point
}

func (v vehicle) ActualPosition() *gps.Point {
	return v.actualPosition
}
