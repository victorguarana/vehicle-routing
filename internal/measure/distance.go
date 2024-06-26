package measure

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/route"
)

func TotalDistance(itineraryInfo itinerary.Info) float64 {
	iterator := itineraryInfo.RouteIterator()
	var totalDistance float64
	for iterator.HasNext() {
		actual := iterator.Actual()
		next := iterator.Next()
		totalDistance += gps.ManhattanDistanceBetweenPoints(actual.Point(), next.Point())
		if subRoutes := actual.StartingSubRoutes(); len(subRoutes) > 0 {
			for _, subRoute := range subRoutes {
				totalDistance += calcSubRouteDistance(subRoute)
			}
		}
		iterator.GoToNext()
	}
	return totalDistance
}

func calcSubRouteDistance(subRoute route.ISubRoute) float64 {
	var totalSubRouteDistance float64
	iterator := subRoute.Iterator()
	actualPoint := subRoute.StartingStop().Point()
	nextPoint := subRoute.First().Point()
	totalSubRouteDistance += gps.EuclideanDistanceBetweenPoints(actualPoint, nextPoint)
	for iterator.HasNext() {
		actualPoint = iterator.Actual().Point()
		nextPoint = iterator.Next().Point()
		totalSubRouteDistance += gps.EuclideanDistanceBetweenPoints(actualPoint, nextPoint)
		iterator.GoToNext()
	}
	actualPoint = nextPoint
	nextPoint = subRoute.ReturningStop().Point()
	totalSubRouteDistance += gps.EuclideanDistanceBetweenPoints(actualPoint, nextPoint)
	return totalSubRouteDistance
}
