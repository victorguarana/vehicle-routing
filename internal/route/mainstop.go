package route

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/slc"
)

//go:generate mockgen -source=mainstop.go -destination=mock/mainstopmock.go
type IMainStop interface {
	IsCustomer() bool
	IsWarehouse() bool
	Latitude() float64
	Longitude() float64
	Name() string
	Point() gps.Point
	ReturningSubRoutes() []ISubRoute
	StartingSubRoutes() []ISubRoute
	RemoveStartingSubRoute(subRoute ISubRoute)
	RemoveReturningSubRoute(subRoute ISubRoute)
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

func (ms *mainStop) IsCustomer() bool {
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

func (ms *mainStop) ReturningSubRoutes() []ISubRoute {
	subRoutes := make([]ISubRoute, len(ms.returningSubRoutes))
	for i, sr := range ms.returningSubRoutes {
		subRoutes[i] = sr
	}
	return subRoutes
}

func (ms *mainStop) StartingSubRoutes() []ISubRoute {
	subRoutes := make([]ISubRoute, len(ms.startingSubRoutes))
	for i, sr := range ms.startingSubRoutes {
		subRoutes[i] = sr
	}
	return subRoutes
}

func (ms *mainStop) RemoveStartingSubRoute(iSubRoute ISubRoute) {
	subRoute := iSubRoute.(*subRoute)
	ms.startingSubRoutes = slc.RemoveElement(ms.startingSubRoutes, subRoute)
}

func (ms *mainStop) RemoveReturningSubRoute(iSubRoute ISubRoute) {
	subRoute := iSubRoute.(*subRoute)
	ms.returningSubRoutes = slc.RemoveElement(ms.returningSubRoutes, subRoute)
}
