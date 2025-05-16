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
	CarEfficiency() float64
	CarSpeed() float64
	CarSupport(nextPoints ...gps.Point) bool
	Drones() []vehicle.IDrone
	DroneCanReach(drone vehicle.IDrone, nextPoints ...gps.Point) bool
	DroneEfficiency() float64
	DroneIsFlying(drone vehicle.IDrone) bool
	DroneSpeed() float64
	DroneSupport(drone vehicle.IDrone, deliveryPoint gps.Point, landingPoint gps.Point) bool
	RouteIterator() slc.Iterator[route.IMainStop]
	SubItineraryList() []SubItinerary
}

type info struct {
	*itinerary
}

func (i *info) ActualCarPoint() gps.Point {
	return i.car.ActualPoint()
}

func (i *info) ActualCarStop() route.IMainStop {
	return i.route.Last()
}

func (i *info) CarEfficiency() float64 {
	return vehicle.CarEfficiency
}

func (i *info) CarSpeed() float64 {
	return vehicle.CarSpeed
}

func (i *info) CarSupport(nextPoints ...gps.Point) bool {
	return i.car.Support(nextPoints...)
}

func (i *info) DroneCanReach(drone vehicle.IDrone, nextPoints ...gps.Point) bool {
	return drone.CanReach(nextPoints...)
}

func (i *info) DroneEfficiency() float64 {
	return vehicle.DroneEfficiency
}

func (i *info) DroneIsFlying(drone vehicle.IDrone) bool {
	return drone.IsFlying()
}

func (i *info) Drones() []vehicle.IDrone {
	return i.car.Drones()
}

func (i *info) DroneSpeed() float64 {
	return vehicle.DroneSpeed
}

func (i *info) DroneSupport(drone vehicle.IDrone, deliveryPoint gps.Point, landingPoint gps.Point) bool {
	return drone.Support(deliveryPoint) && drone.CanReach(deliveryPoint, landingPoint)
}

func (i *info) RouteIterator() slc.Iterator[route.IMainStop] {
	return i.route.Iterator()
}

func (i *info) SubItineraryList() []SubItinerary {
	return i.completedSubItineraryList
}
