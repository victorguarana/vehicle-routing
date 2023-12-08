package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

func ClosestNeighbor(routeList []routes.IRoute, m gps.Map) error {
	var route routes.IRoute
	var car vehicles.ICar
	var carActualPosition, closestClient, closestDepositFromClosestClient *gps.Point

	remaningClients := make([]*gps.Point, len(m.Clients))
	copy(remaningClients, m.Clients)

	for i := 0; len(remaningClients) > 0; i++ {
		route = swapBetween(routeList, i)
		car = route.Car()
		carActualPosition = car.ActualPosition()

		closestClient = closestPoint(carActualPosition, remaningClients)
		closestDepositFromClosestClient = closestPoint(closestClient, m.Deposits)

		// Move to closest deposit when car does not support closest client
		if !car.Support(closestClient, closestDepositFromClosestClient) {
			closestDepositFromActualPosition := closestPoint(carActualPosition, m.Deposits)
			err := moveAndAppend(route, closestDepositFromActualPosition)
			if err != nil {
				return err
			}
			continue
		}

		// Move to closest client
		err := moveAndAppend(route, closestClient)
		if err != nil {
			return err
		}

		remaningClients = removePoint(remaningClients, closestClient)
	}

	// Finish routes in closest deposit
	finishRoutes(routeList, m)

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
