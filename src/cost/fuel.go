package cost

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
)

func CalcTotalFuel(itn itinerary.Itinerary) float64 {
	iterator := itn.RouteIterator()
	var totalFuelSpent float64
	carEfficiency := itn.CarEfficiency()
	droneEfficiency := itn.DroneEfficiency()
	for iterator.HasNext() {
		actual := iterator.Actual()
		next := iterator.Next()
		totalFuelSpent += gps.DistanceBetweenPoints(actual.Point(), next.Point()) / carEfficiency
		if subRoutes := actual.StartingSubRoutes(); len(subRoutes) > 0 {
			for _, subRoute := range subRoutes {
				totalFuelSpent += calcSubRouteDistance(subRoute) / droneEfficiency
			}
		}
		iterator.GoToNext()
	}
	return totalFuelSpent
}
