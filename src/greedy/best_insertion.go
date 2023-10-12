package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

func BestInsertion(car vehicles.ICar, m gps.Map) (routes.IRoute, error) {
	route := routes.NewRoute(car)
	orderedClients := orderedClients(m.Clients)

	for i := range orderedClients {
		client := orderedClients[i]
		closestDeposit := closestPoint(client, m.Deposits)
		if car.Support(*client, *closestDeposit) {
			err := moveAndAppend(route, client)
			if err != nil {
				return nil, err
			}
		} else {
			err := moveAndAppend(route, closestDeposit)
			if err != nil {
				return nil, err
			}
			i--
		}
	}

	return route, nil
}

func orderedClients(clients []*gps.Point) []*gps.Point {
	orderedClients := []*gps.Point{}

	for _, client := range clients {
		index := findBestInsertionIndex(client, orderedClients)
		orderedClients = insertAt(index, client, orderedClients)
	}

	return orderedClients
}

func findBestInsertionIndex(client *gps.Point, orderedClients []*gps.Point) int {
	var bestIndex int
	var bestDistance float64

	for index, c := range orderedClients {
		distance := gps.DistanceBetweenPoints(*client, *c)

		if bestDistance == 0 || distance < bestDistance {
			bestIndex = index
			bestDistance = distance
		}
	}

	return bestIndex
}

func insertAt(index int, client *gps.Point, orderedClients []*gps.Point) []*gps.Point {
	newClientsList := make([]*gps.Point, len(orderedClients)+1)
	newClientsList = append(newClientsList, orderedClients[:index]...)
	newClientsList = append(newClientsList, client)
	newClientsList = append(newClientsList, orderedClients[index:]...)
	return newClientsList
}
