package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/slc"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

func BestInsertion(cars []vehicles.ICar, m gps.Map) {
	orderedClientsListsByRoute := orderedClientsByCars(cars, m.Clients)
	for itinerary, orderedClients := range orderedClientsListsByRoute {
		fillRoute(itinerary, orderedClients, m.Deposits)
	}
	finishItineraryOnClosestDeposits(cars, m)
}

func orderedClientsByCars(cars []vehicles.ICar, clients []gps.Point) map[vehicles.ICar][]gps.Point {
	orderedClientsByItinerary := map[vehicles.ICar][]gps.Point{}
	for i, client := range clients {
		car := slc.CircularSelection(cars, i)
		initialPoint := car.ActualPoint()
		orderedClients := orderedClientsByItinerary[car]
		orderedClientsByItinerary[car] = insertInBestPosition(initialPoint, client, orderedClients)
	}
	return orderedClientsByItinerary
}

func insertInBestPosition(initialPoint gps.Point, client gps.Point, orderedClients []gps.Point) []gps.Point {
	if len(orderedClients) == 0 {
		return []gps.Point{client}
	}
	bestIndex := findBestPosition(initialPoint, client, orderedClients)
	return slc.InsertAt(orderedClients, client, bestIndex)
}

func findBestPosition(initialPoint gps.Point, client gps.Point, orderedClients []gps.Point) int {
	var bestIndex int
	shortestAdditionalDistance := calcAdditionalDistance(initialPoint, client, orderedClients[0])
	for i := 1; i < len(orderedClients); i++ {
		addictionalDistance := calcAdditionalDistance(orderedClients[i-1], client, orderedClients[i])
		if addictionalDistance < shortestAdditionalDistance {
			bestIndex = i
			shortestAdditionalDistance = addictionalDistance
		}
	}

	addictionalDistance := calcAdditionalDistance(orderedClients[len(orderedClients)-1], client, initialPoint)
	if addictionalDistance < shortestAdditionalDistance {
		bestIndex = len(orderedClients)
	}

	return bestIndex
}

func fillRoute(car vehicles.ICar, orderedClients []gps.Point, deposits []gps.Point) {
	for _, client := range orderedClients {
		closestDeposit := gps.ClosestPoint(client, deposits)
		if !car.Support(client, closestDeposit) {
			car.Move(routes.NewMainStop(closestDeposit))
		}
		car.Move(routes.NewMainStop(client))
	}
}

func calcAdditionalDistance(from, through, to gps.Point) float64 {
	return gps.DistanceBetweenPoints(from, through, to) - gps.DistanceBetweenPoints(from, to)
}
