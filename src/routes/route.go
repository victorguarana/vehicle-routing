package routes

import (
	"errors"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

var (
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrNilCar          = errors.New("car can not be nil")
)

type IRoute interface {
	Car() vehicles.ICar
	First() ICarStop
	Last() ICarStop
	Len() int

	Append(*gps.Point) error
	AtIndex(int) (ICarStop, error)

	RemoveCarStop(int) error
}

type route struct {
	car   vehicles.ICar
	stops []*carStop
}

func NewRoute(car vehicles.ICar) (IRoute, error) {
	if car == nil {
		return nil, ErrNilCar
	}

	return &route{
		car:   car,
		stops: []*carStop{},
	}, nil
}

func (r *route) Car() vehicles.ICar {
	return r.car
}

func (r *route) First() ICarStop {
	return r.stops[0]
}

func (r *route) Last() ICarStop {
	return r.stops[len(r.stops)-1]
}

func (r *route) Len() int {
	return len(r.stops)
}

func (r *route) Append(point *gps.Point) error {
	carStop := newCarStop(r.car, point)
	r.stops = append(r.stops, carStop)
	return nil
}

func (r *route) AtIndex(index int) (ICarStop, error) {
	if index < 0 || index >= len(r.stops) {
		return nil, ErrIndexOutOfRange
	}

	return r.stops[index], nil
}

func (r *route) RemoveCarStop(index int) error {
	if index < 0 || index >= len(r.stops) {
		return ErrIndexOutOfRange
	}

	r.stops = append(r.stops[:index], r.stops[index+1:]...)
	return nil
}
