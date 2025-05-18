package route

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
)

//go:generate mockgen -source=substop.go -destination=mock/substopmock.go
type ISubStop interface {
	IsCustomer() bool
	IsWarehouse() bool
	Latitude() float64
	Longitude() float64
	Name() string
	Point() gps.Point
}

type subStop struct {
	point gps.Point
}

func NewSubStop(p gps.Point) ISubStop {
	return &subStop{point: p}
}

func (ss *subStop) IsCustomer() bool {
	return ss.point.PackageSize != 0
}

func (ss *subStop) IsWarehouse() bool {
	return ss.point.PackageSize == 0
}

func (ss *subStop) Latitude() float64 {
	return ss.point.Latitude
}

func (ss *subStop) Longitude() float64 {
	return ss.point.Longitude
}

func (ss *subStop) Name() string {
	return ss.point.Name
}

func (ss *subStop) Point() gps.Point {
	return ss.point
}
