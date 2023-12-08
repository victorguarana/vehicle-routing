package routes

import (
	"errors"
	"fmt"
	"strings"

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
	String() string

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

func (r *route) String() string {
	str := []string{"Route:"}

	for i, carStop := range r.stops {
		str = append(str, fmt.Sprintf("  CarStop #%d (%s)", i, carStop.point))

		for j, flight := range carStop.flights {
			if flight.takeoffPoint != carStop {
				continue
			}

			str = append(str, fmt.Sprintf("    Flight #%d.%d (%s):", i, j, flight.drone.Name()))
			str = append(str, fmt.Sprintf("     Takeoff #%d.%d (%s)", i, j, flight.takeoffPoint.point))

			for k, droneStop := range flight.stops {
				str = append(str, fmt.Sprintf("      DroneStop #%d.%d.%d (%s)", i, j, k, droneStop.point))
			}

			str = append(str, fmt.Sprintf("     Landing #%d.%d (%s)", i, j, flight.landingPoint.point))
		}
		str = append(str, "")
	}

	return strings.Join(str, "\n")
}
