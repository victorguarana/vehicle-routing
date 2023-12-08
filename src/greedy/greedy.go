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

func finishRoutes(routesList []routes.IRoute, m gps.Map) error {
	for _, route := range routesList {
		closestDeposit := closestPoint(route.Last().Point(), m.Deposits)
		err := moveAndAppend(route, closestDeposit)
		if err != nil {
			return err
		}
	}

	return nil
}

func moveAndAppend(route routes.IRoute, point *gps.Point) error {
	err := route.Car().Move(point)
	if err != nil {
		return err
	}

	err = route.Append(point)
	if err != nil {
		return err
	}

	return nil
}

func swapBetween[T any](list []T, index int) T {
	i := index % len(list)
	return list[i]
}
