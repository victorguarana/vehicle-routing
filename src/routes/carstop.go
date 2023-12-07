package routes

import (
	"errors"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

var (
	ErrInvalidCarStop = errors.New("invalid car stop")
)

type ICarStop interface {
	Car() vehicles.ICar
	IsClient() bool
	IsDeposit() bool
	Flights() []IFlight
	Point() *gps.Point
}

type carStop struct {
	point   *gps.Point
	car     vehicles.ICar
	flights []*flight
}

func newCarStop(car vehicles.ICar, point *gps.Point) *carStop {
	return &carStop{
		point:   point,
		car:     car,
		flights: []*flight{},
	}
}

func (cs *carStop) Car() vehicles.ICar {
	return cs.car
}

func (cs *carStop) IsClient() bool {
	return cs.point.PackageSize != 0
}

func (cs *carStop) IsDeposit() bool {
	return cs.point.PackageSize == 0
}

func (cs *carStop) Flights() []IFlight {
	var flights []IFlight
	for _, f := range cs.flights {
		flights = append(flights, f)
	}

	return flights
}

func (cs *carStop) Point() *gps.Point {
	return cs.point
}
