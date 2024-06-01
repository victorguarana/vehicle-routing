package itinerary

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/route"
	"github.com/victorguarana/vehicle-routing/src/slc"
	"github.com/victorguarana/vehicle-routing/src/vehicle"
)

type Info interface {
	ActualCarPoint() gps.Point
	ActualCarStop() route.IMainStop
	CarEfficiency() float64
	CarSpeed() float64
	CarSupport(nextPoints ...gps.Point) bool
	DroneCanReach(droneNumber DroneNumber, nextPoints ...gps.Point) bool
	DroneEfficiency() float64
	DroneIsFlying(droneNumber DroneNumber) bool
	DroneNumbers() []DroneNumber
	DroneSpeed() float64
	DroneSupport(droneNumber DroneNumber, deliveryPoint gps.Point, landingPoint gps.Point) bool
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

func (i *info) DroneCanReach(droneNumber DroneNumber, nextPoints ...gps.Point) bool {
	drone := i.droneByNumber(droneNumber)
	return drone.CanReach(nextPoints...)
}

func (i *info) DroneEfficiency() float64 {
	return vehicle.DroneEfficiency
}

func (i *info) DroneIsFlying(droneNumber DroneNumber) bool {
	drone := i.droneByNumber(droneNumber)
	return drone.IsFlying()
}

func (i *info) DroneNumbers() []DroneNumber {
	var droneNumbers []DroneNumber
	for droneNumber := range i.droneNumbersMap {
		droneNumbers = append(droneNumbers, droneNumber)
	}
	return droneNumbers
}

func (i *info) DroneSpeed() float64 {
	return vehicle.DroneSpeed
}

func (i *info) DroneSupport(droneNumber DroneNumber, deliveryPoint gps.Point, landingPoint gps.Point) bool {
	drone := i.droneByNumber(droneNumber)
	return drone.Support(deliveryPoint) && drone.CanReach(deliveryPoint, landingPoint)
}

func (i *info) RouteIterator() slc.Iterator[route.IMainStop] {
	return i.route.Iterator()
}

func (i *info) SubItineraryList() []SubItinerary {
	return i.completedSubItineraryList
}

func (i *info) droneByNumber(droneNumber DroneNumber) vehicle.IDrone {
	return i.droneNumbersMap[droneNumber]
}
