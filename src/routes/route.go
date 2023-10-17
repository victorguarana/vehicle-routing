package routes

import (
	"fmt"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

type IRoute interface {
	CompleteRoute() []*gps.Point
	First() *gps.Point
	Last() *gps.Point
	Append(*gps.Point) error
	// InsertAt(int, *gps.Point) error
	Car() vehicles.ICar
	String() string
}

type route struct {
	car    vehicles.ICar
	points []*gps.Point
}

func NewRoute(car vehicles.ICar) IRoute {
	return &route{
		car:    car,
		points: make([]*gps.Point, 0),
	}
}

func (r *route) CompleteRoute() []*gps.Point {
	return r.points
}

func (r *route) Car() vehicles.ICar {
	return r.car
}

func (r *route) First() *gps.Point {
	return r.points[0]
}

func (r *route) Last() *gps.Point {
	return r.points[0]
}

func (r *route) Append(point *gps.Point) error {
	r.points = append(r.points, point)
	return nil
}

func (r *route) String() string {
	str := "Route:\n"
	for i, point := range r.points {
		str += fmt.Sprintf("#%d - %s", i, point.String())
	}

	return str
}
