package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

// var ErrIndexOutOfRange = errors.New("Index out of range")

type IRoute interface {
	// CompleteRoute() []ICarStop
	First() ICarStop
	Last() ICarStop
	Append(*gps.Point) error
	// InsertAt(int, ICarStop) error
	Car() vehicles.ICar
	// String() string
}

type route struct {
	car   vehicles.ICar
	stops []ICarStop
}

func NewRoute(car vehicles.ICar) IRoute {
	return &route{
		car:   car,
		stops: []ICarStop{},
	}
}

// func (r *route) CompleteRoute() []ICarStop {
// 	return r.stops
// }

// func (r *route) GetAtIndex(index int) (ICarStop, error) {
// 	if index < 0 || index >= len(r.stops) {
// 		return nil, ErrIndexOutOfRange
// 	}

// 	return r.stops[index], nil
// }

func (r *route) Car() vehicles.ICar {
	return r.car
}

func (r *route) First() ICarStop {
	return r.stops[0]
}

func (r *route) Last() ICarStop {
	return r.stops[len(r.stops)-1]
}

func (r *route) Append(point *gps.Point) error {
	carStop := NewCarStop(point, r.car)
	r.stops = append(r.stops, carStop)
	return nil
}

// func (r *route) String() string {
// 	str := "Route:\n"
// 	for i, point := range r.stops {
// 		str += fmt.Sprintf("#%d - %s", i, point.String())
// 	}

// 	return str
// }
