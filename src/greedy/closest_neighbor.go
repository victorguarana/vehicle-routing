package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
)

func ClosestNeighbor(route routes.IRoute, m gps.Map) error {
	car := route.Car()

	// Add starting point to route
	carActualPosition := car.ActualPosition()
	route.Append(carActualPosition)

	remaningClients := make([]*gps.Point, len(m.Clients))
	copy(remaningClients, m.Clients)

	var closestClient, closestDepositFromClosestClient *gps.Point
	for len(remaningClients) > 0 {
		closestClient = closestPoint(carActualPosition, remaningClients)
		closestDepositFromClosestClient = closestPoint(closestClient, m.Deposits)

		if !car.Support(closestClient, closestDepositFromClosestClient) {
			// Move to closest deposit when car does not support closest client
			closestDepositFromActualPosition := closestPoint(carActualPosition, m.Deposits)
			err := moveAndAppend(route, closestDepositFromActualPosition)
			if err != nil {
				return err
			}
			carActualPosition = car.ActualPosition()
			continue
		}

		// Move to closest client
		err := moveAndAppend(route, closestClient)
		if err != nil {
			return err
		}

		remaningClients = removePoint(remaningClients, closestClient)
		carActualPosition = car.ActualPosition()
	}

	// Finish route in closest deposit
	err := moveAndAppend(route, closestDepositFromClosestClient)
	if err != nil {
		return err
	}

	return nil
}

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

func removePoint(points []*gps.Point, point *gps.Point) []*gps.Point {
	var newPoints []*gps.Point

	for _, p := range points {
		if p != point {
			newPoints = append(newPoints, p)
		}
	}

	return newPoints
}
