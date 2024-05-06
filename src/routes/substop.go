package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
)

type ISubStop interface {
	Point() gps.Point
}

type subStop struct {
	point gps.Point
}

func NewSubStop(p gps.Point) ISubStop {
	return &subStop{point: p}
}

func (ss *subStop) Point() gps.Point {
	return ss.point
}
