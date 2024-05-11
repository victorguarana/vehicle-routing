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

func finishItineraryOnClosestDeposits(carsList []vehicles.ICar, m gps.Map) {
	for _, car := range carsList {
		position := car.ActualPoint()
		closestDeposit := closestPoint(position, m.Deposits)
		car.Move(routes.NewMainStop(closestDeposit))
	}
}
