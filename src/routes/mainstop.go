package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
)

type IMainStop interface {
	SubRoutes() []ISubRoute
	IsClient() bool
	IsDeposit() bool
	Point() gps.Point
}

type mainStop struct {
	point     gps.Point
	subRoutes []*subRoute
}

func NewMainStop(point gps.Point) IMainStop {
	return &mainStop{
		point:     point,
		subRoutes: []*subRoute{},
	}
}

func (ms *mainStop) SubRoutes() []ISubRoute {
	subroutes := make([]ISubRoute, len(ms.subRoutes))
	for _, f := range ms.subRoutes {
		subroutes = append(subroutes, f)
	}
	return subroutes
}

func (ms *mainStop) IsClient() bool {
	return ms.point.PackageSize != 0
}

func (ms *mainStop) IsDeposit() bool {
	return ms.point.PackageSize == 0
}

func (ms *mainStop) Point() gps.Point {
	return ms.point
}
