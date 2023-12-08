package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
)

func BestInsertion(route routes.IRoute, m gps.Map) error {
	car := route.Car()
	initialPosition := car.ActualPosition()
	orderedClients := orderedClients(initialPosition, m.Clients)

	var closestDeposit *gps.Point
	for i := range orderedClients {
		client := orderedClients[i]
		closestDeposit = closestPoint(client, m.Deposits)
		if car.Support(client, closestDeposit) {
			err := moveAndAppend(route, client)
			if err != nil {
				return err
			}
		} else {
			err := moveAndAppend(route, closestDeposit)
			if err != nil {
				return err
			}
			i--
		}
	}

	err := moveAndAppend(route, closestDeposit)
	if err != nil {
		return err
	}

	return nil
}

func orderedClients(initialPoint *gps.Point, clients []*gps.Point) []*gps.Point {
	orderedClients := []*gps.Point{}

	for _, client := range clients {
		index := findBestInsertionIndex(initialPoint, client, orderedClients)
		orderedClients = insertAt(index, client, orderedClients)
	}

	return orderedClients
}

func findBestInsertionIndex(initialPoint *gps.Point, client *gps.Point, orderedClients []*gps.Point) int {
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

func insertAt(index int, client *gps.Point, orderedClients []*gps.Point) []*gps.Point {
	newClientsList := make([]*gps.Point, len(orderedClients)+1)
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
