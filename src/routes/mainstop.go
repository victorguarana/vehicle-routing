package routes

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
)

type IMainStop interface {
	IsClient() bool
	IsWarehouse() bool
	Latitude() float64
	Longitude() float64
	Name() string
	Point() gps.Point
	StartingSubRoutes() []ISubRoute
}

type mainStop struct {
	point              gps.Point
	startingSubRoutes  []*subRoute
	returningSubRoutes []*subRoute
}

func NewMainStop(point gps.Point) IMainStop {
	return &mainStop{
		point:              point,
		startingSubRoutes:  []*subRoute{},
		returningSubRoutes: []*subRoute{},
	}
}

func (ms *mainStop) IsClient() bool {
	return ms.point.PackageSize != 0
}

func (ms *mainStop) IsWarehouse() bool {
	return ms.point.PackageSize == 0
}

func (ms *mainStop) Latitude() float64 {
	return ms.point.Latitude
}

func (ms *mainStop) Longitude() float64 {
	return ms.point.Longitude
}

func (ms *mainStop) Name() string {
	return ms.point.Name
}

func (ms *mainStop) Point() gps.Point {
	return ms.point
}

func (ms *mainStop) StartingSubRoutes() []ISubRoute {
	subRoutes := make([]ISubRoute, len(ms.startingSubRoutes))
	for i, sr := range ms.startingSubRoutes {
		subRoutes[i] = sr
	}
	return subRoutes
}
