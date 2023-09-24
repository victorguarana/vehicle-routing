package vehicles

import (
	"errors"

	"github.com/victorguarana/go-vehicle-route/src/gps"
)

var (
	ErrInvalidParams = errors.New("invalid param")
)

type ivehicle interface {
	Move(*gps.Point) error
	Support(...gps.Point) bool
}

type vehicle struct {
	speed          float64
	name           string
	actualPosition *gps.Point
}
