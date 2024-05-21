package cost

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/routes"
)

// TODO: Calc Costs (Ex: Distance / vehicle fuel costs)
// TODO: Calc

func CalcTotalDistanceSpent(itn itinerary.Itinerary) float64 {
	iterator := itn.RouteIterator()
	var totalDistance float64
	for iterator.HasNext() {
		actual := iterator.Actual()
		next := iterator.Next()
		totalDistance += gps.DistanceBetweenPoints(actual.Point(), next.Point())
		if subRoutes := actual.StartingSubRoutes(); len(subRoutes) > 0 {
			for _, subRoute := range subRoutes {
				totalDistance += calcSubRouteDistance(subRoute)
			}
		}
		iterator.GoToNext()
	}
	return totalDistance
}

func calcSubRouteDistance(subRoute routes.ISubRoute) float64 {
	var totalSubRouteDistance float64
	iterator := subRoute.Iterator()
	actualPoint := subRoute.StartingStop().Point()
	nextPoint := subRoute.First().Point()
	totalSubRouteDistance += gps.DistanceBetweenPoints(actualPoint, nextPoint)
	for iterator.HasNext() {
		actualPoint = iterator.Actual().Point()
		nextPoint = iterator.Next().Point()
		totalSubRouteDistance += gps.DistanceBetweenPoints(actualPoint, nextPoint)
		iterator.GoToNext()
	}
	actualPoint = nextPoint
	nextPoint = subRoute.ReturningStop().Point()
	totalSubRouteDistance += gps.DistanceBetweenPoints(actualPoint, nextPoint)
	return totalSubRouteDistance
}
