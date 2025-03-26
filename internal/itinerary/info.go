package itinerary

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/route"
	"github.com/victorguarana/vehicle-routing/internal/slc"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

//go:generate mockgen -source=info.go -destination=mock/infomock.go
type Info interface {
	ActualCarPoint() gps.Point
	ActualCarStop() route.IMainStop
	Car() vehicle.ICar
	MainRoute() route.IMainRoute
	RouteIterator() slc.Iterator[route.IMainStop]
	SubItineraryList() []SubItinerary
}

var _ Info = (*info)(nil)

type info struct {
	*itinerary
}

func (i *info) ActualCarPoint() gps.Point {
	return i.car.ActualPoint()
}

func (i *info) ActualCarStop() route.IMainStop {
	return i.route.Last()
}

func (i *info) Car() vehicle.ICar {
	return i.car
}

func (i *info) MainRoute() route.IMainRoute {
	return i.route
}

func (i *info) RouteIterator() slc.Iterator[route.IMainStop] {
	return i.route.Iterator()
}

func (i *info) SubItineraryList() []SubItinerary {
	return i.completedSubItineraryList
}
