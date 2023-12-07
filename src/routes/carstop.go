package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

type ICarStop interface {
	IsClient() bool
	IsDeposit() bool
	Point() *gps.Point
	Car() vehicles.ICar
	Flights() []*flight
}

type carStop struct {
	point   *gps.Point
	car     vehicles.ICar
	flights []*flight
}

func NewCarStop(point *gps.Point, car vehicles.ICar) ICarStop {
	return &carStop{
		point:   point,
		car:     car,
		flights: []*flight{},
	}
}

func (cs *carStop) Point() *gps.Point {
	return cs.point
}

func (cs *carStop) Car() vehicles.ICar {
	return cs.car
}

func (cs *carStop) Flights() []*flight {
	return cs.flights
}

func (cs *carStop) IsClient() bool {
	return cs.point.PackageSize != 0
}

func (cs *carStop) IsDeposit() bool {
	return cs.point.PackageSize == 0
}
