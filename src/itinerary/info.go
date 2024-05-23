package itinerary

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/route"
	"github.com/victorguarana/vehicle-routing/src/slc"
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
