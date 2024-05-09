package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

func closestPoint(originPoint gps.Point, candidatePoints []gps.Point) gps.Point {
	var closestPoint gps.Point
	closestDistance := -1.0
	for _, candidatePoint := range candidatePoints {
		if closestDistance < 0 || gps.DistanceBetweenPoints(originPoint, candidatePoint) < closestDistance {
			closestPoint = candidatePoint
			closestDistance = gps.DistanceBetweenPoints(originPoint, closestPoint)
		}
	}
	return closestPoint
}

func finishItineraryOnClosestDeposits(itineraryList []routes.Itinerary, m gps.Map) {
	for _, itinerary := range itineraryList {
		route := itinerary.Route
		car := itinerary.Car
		lastStop := route.Last()
		closestDeposit := closestPoint(lastStop.Point(), m.Deposits)
		moveCarAndAppendRoute(car, route, closestDeposit)
	}
}

func moveCarAndAppendRoute(car vehicles.ICar, route routes.IMainRoute, point gps.Point) {
	car.Move(point)
	route.Append(routes.NewMainStop(point))
}
