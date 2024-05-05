package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
)

func BestInsertion(routesList []routes.IRoute, m gps.Map) error {
	orderedClientsListsByRoute := orderedClientsByRoutes(routesList, m.Clients)

	var closestDeposit gps.Point
	for i, orderedClients := range orderedClientsListsByRoute {
		route := routesList[i]
		car := route.Car()

		for j := range orderedClients {
			client := orderedClients[j]
			closestDeposit = closestPoint(client, m.Deposits)

			var destination gps.Point
			if car.Support(client, closestDeposit) {
				destination = client
			} else {
				destination = closestDeposit
				j--
			}

			route.Car().Move(destination)
			route.Append(destination)
		}

		route.Car().Move(closestDeposit)
		route.Append(closestDeposit)
	}

	return nil
}

func orderedClientsByRoutes(routes []routes.IRoute, clients []gps.Point) [][]gps.Point {
	size := len(routes)
	orderedClientsLists := make([][]gps.Point, size)
	initialPoints := make([]gps.Point, size)
	for i := range routes {
		initialPoints[i] = routes[i].Car().ActualPosition()
	}

	for i, client := range clients {
		i = i % size
		index := findBestInsertionIndex(initialPoints[i], client, orderedClientsLists[i])
		orderedClientsLists[i] = insertAt(index, client, orderedClientsLists[i])
	}

	return orderedClientsLists
}

func findBestInsertionIndex(initialPoint gps.Point, client gps.Point, orderedClients []gps.Point) int {
	if len(orderedClients) == 0 {
		return 0
	}

	bestIndex := 0
	bestAdicionalDistance := gps.DistanceBetweenPoints(initialPoint, client, orderedClients[0]) - gps.DistanceBetweenPoints(initialPoint, orderedClients[0])

	for i := 0; i < len(orderedClients)-1; i++ {
		adicionalDistance := gps.DistanceBetweenPoints(orderedClients[i], client, orderedClients[i+1]) - gps.DistanceBetweenPoints(orderedClients[i], orderedClients[i+1])

		if adicionalDistance < bestAdicionalDistance {
			bestIndex = i + 1
			bestAdicionalDistance = adicionalDistance
		}
	}

	adicionalDistance := gps.DistanceBetweenPoints(orderedClients[len(orderedClients)-1], client)
	if adicionalDistance < bestAdicionalDistance {
		bestIndex = len(orderedClients)
	}

	return bestIndex
}

func insertAt(index int, client gps.Point, orderedClients []gps.Point) []gps.Point {
	newClientsList := make([]gps.Point, len(orderedClients)+1)
	for i := 0; i < len(newClientsList); i++ {
		switch {
		case i < index:
			newClientsList[i] = orderedClients[i]
		case i == index:
			newClientsList[i] = client
		case i > index:
			newClientsList[i] = orderedClients[i-1]
		}
	}
	return newClientsList
}
