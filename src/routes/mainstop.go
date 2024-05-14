package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
)

type IMainStop interface {
	IsClient() bool
	IsDeposit() bool
	Point() gps.Point
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

func (ms *mainStop) IsDeposit() bool {
	return ms.point.PackageSize == 0
}

func (ms *mainStop) Point() gps.Point {
	return ms.point
}
