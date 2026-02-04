package measure

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/route"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

// type subRouteTimes map[route.ISubRoute]float64

func TimeSpentAgatz(itineraryInfo itinerary.Info) float64 {
	var subRoutesFlyingTimes = make(subRouteTimes)
	var mainRouteTravelTime = make(subRouteTimes)
	var totalTime float64
	carSpeed := itineraryInfo.Car().Speed()
	iterator := itineraryInfo.RouteIterator()
	droneByFlight := mapDroneByFlight(itineraryInfo)

	for {
		actual := iterator.Actual()
		if subRoutes := actual.ReturningSubRoutes(); len(subRoutes) > 0 {
			totalTime += additionalTimeWaitingSubRoutesAgatz(subRoutes, droneByFlight)
			removeReturningSubRoutes(mainRouteTravelTime, subRoutesFlyingTimes, subRoutes)
		}
		if !iterator.HasNext() {
			break
		}
		next := iterator.Next()
		travelTime := gps.EuclideanDistanceBetweenPoints(actual.Point(), next.Point()) / carSpeed
		updateMainRouteTravelTimes(mainRouteTravelTime, travelTime)
		totalTime += travelTime
		iterator.GoToNext()
	}

	return totalTime
}

// Considerando que existe apenas um drone
// E que o drone pode realizar múltiplas subrotas derivadas de um mesmo ponto
func additionalTimeWaitingSubRoutesAgatz(subRoutes []route.ISubRoute, droneBySubRoute map[route.ISubRoute]vehicle.IDrone) float64 {
	var waitingTime float64
	for _, subRoute := range subRoutes {
		waitingTime += calcSubRouteDistance(subRoute) / droneBySubRoute[subRoute].Speed()
	}
	return waitingTime
}
