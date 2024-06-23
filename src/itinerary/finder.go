package itinerary

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/route"
	"github.com/victorguarana/vehicle-routing/src/slc"
)

//go:generate mockgen -source=finder.go -destination=mock/findermock.go
type Finder interface {
	FindWorstDroneStop() DroneStopCost
	FindWorstSwappableCarStopsOrdered() []CarStopCost
}

type finder struct {
	*info
}

type DroneStopCost struct {
	Index  int
	Stop   route.ISubStop
	Flight route.ISubRoute
	cost   float64
}

type CarStopCost struct {
	Stop  route.IMainStop
	Index int
	cost  float64
}

func (f finder) FindWorstDroneStop() DroneStopCost {
	worstDroneStop := DroneStopCost{}
	subItineraryList := f.SubItineraryList()
	for _, subItinerary := range subItineraryList {
		worstDroneStopFromSubItinerary := findWorstDroneStopInFlight(subItinerary.Flight)
		if worstDroneStopFromSubItinerary.cost > worstDroneStop.cost {
			worstDroneStop = worstDroneStopFromSubItinerary
		}
	}

	return worstDroneStop
}

func findWorstDroneStopInFlight(flight route.ISubRoute) DroneStopCost {
	// When the flight has only one stop
	if flight.First() == flight.Last() {
		return DroneStopCost{
			Index:  0,
			Stop:   flight.First(),
			Flight: flight,
			cost: gps.AdditionalDistancePassingThrough(
				flight.StartingStop().Point(),
				flight.First().Point(),
				flight.ReturningStop().Point(),
			),
		}
	}

	iterator := flight.Iterator()
	// iterator.GoToNext()
	worstDroneStop := DroneStopCost{}
	// When the flight has more than one stop
	previousPoint := flight.StartingStop().Point()
	for iterator.HasNext() {
		actualPoint := iterator.Actual().Point()
		nextPoint := iterator.Next().Point()
		cost := gps.AdditionalDistancePassingThrough(previousPoint, actualPoint, nextPoint)
		if cost > worstDroneStop.cost {
			worstDroneStop = DroneStopCost{
				Index:  iterator.Index(),
				Stop:   iterator.Actual(),
				Flight: flight,
				cost:   cost,
			}
		}
		previousPoint = actualPoint
		iterator.GoToNext()
	}

	actualPoint := iterator.Actual().Point()
	nextPoint := flight.ReturningStop().Point()
	cost := gps.AdditionalDistancePassingThrough(previousPoint, actualPoint, nextPoint)
	if cost > worstDroneStop.cost {
		worstDroneStop = DroneStopCost{
			Index:  iterator.Index(),
			Stop:   iterator.Actual(),
			Flight: flight,
			cost:   cost,
		}
	}

	return worstDroneStop
}

func (f finder) FindWorstSwappableCarStopsOrdered() []CarStopCost {
	swappableCarStopsOrdered := []CarStopCost{}
	iterator := f.RouteIterator()
	for iterator.HasNext() {
		iterator.GoToNext()
		actual := iterator.Actual()
		if !carStopIsSwappable(actual) {
			continue
		}
		carStopCost := CarStopCost{
			Stop: actual,
			cost: gps.AdditionalDistancePassingThrough(
				iterator.Previous().Point(),
				actual.Point(),
				iterator.Next().Point(),
			),
			Index: iterator.Index(),
		}
		swappableCarStopsOrdered = insertCarStopCostOrdered(swappableCarStopsOrdered, carStopCost)
	}
	return swappableCarStopsOrdered
}

func carStopIsSwappable(carStop route.IMainStop) bool {
	if carStop.IsWarehouse() {
		return false
	}
	if len(carStop.StartingSubRoutes()) > 0 {
		return false
	}
	if len(carStop.ReturningSubRoutes()) > 0 {
		return false
	}
	return true
}

func insertCarStopCostOrdered(carStopCosts []CarStopCost, newCarStopCost CarStopCost) []CarStopCost {
	for i, csc := range carStopCosts {
		if newCarStopCost.cost < csc.cost {
			return slc.InsertAt(carStopCosts, newCarStopCost, i)
		}
	}
	return append(carStopCosts, newCarStopCost)
}
