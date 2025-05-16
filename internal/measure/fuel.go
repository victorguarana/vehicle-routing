package measure

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
)

func SpentFuel(itineraryInfo itinerary.Info) float64 {
	iterator := itineraryInfo.RouteIterator()
	var totalFuelSpent float64
	carEfficiency := itineraryInfo.Car().Efficiency()
	for iterator.HasNext() {
		actual := iterator.Actual()
		next := iterator.Next()
		totalFuelSpent += gps.ManhattanDistanceBetweenPoints(actual.Point(), next.Point()) / carEfficiency
		iterator.GoToNext()
	}

	for _, subItn := range itineraryInfo.SubItineraryList() {
		droneEfficiency := subItn.Drone.Efficiency()
		totalFuelSpent += calcSubRouteDistance(subItn.Flight) / droneEfficiency
	}
	return totalFuelSpent
}
