package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
)

func closestPoint(originPoint *gps.Point, candidatePoints []*gps.Point) *gps.Point {
	var closestPoint *gps.Point
	var closestDistance float64

	for _, candidatePoint := range candidatePoints {
		if closestPoint == nil || gps.DistanceBetweenPoints(originPoint, candidatePoint) < closestDistance {
			closestPoint = candidatePoint
			closestDistance = gps.DistanceBetweenPoints(originPoint, candidatePoint)
		}
	}

	return closestPoint
}

func finishRoutesOnClosestDeposits(routesList []routes.IRoute, m gps.Map) {
	for _, route := range routesList {
		closestDeposit := closestPoint(route.Last().Point(), m.Deposits)
		route.Car().Move(closestDeposit)
		route.Append(closestDeposit)
	}
}

func swapBetween[T any](list []T, index int) T {
	i := index % len(list)
	return list[i]
}
