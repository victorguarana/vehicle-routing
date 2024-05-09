package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/slc"
)

func BestInsertion(itineraryList []routes.Itinerary, m gps.Map) {
	orderedClientsListsByRoute := orderedClientsByItinerary(itineraryList, m.Clients)
	for itinerary, orderedClients := range orderedClientsListsByRoute {
		fillRoute(itinerary, orderedClients, m.Deposits)
	}
	finishItineraryOnClosestDeposits(itineraryList, m)
}

func orderedClientsByItinerary(itineraryList []routes.Itinerary, clients []gps.Point) map[routes.Itinerary][]gps.Point {
	orderedClientsByItinerary := map[routes.Itinerary][]gps.Point{}
	for i, client := range clients {
		itinerary := slc.CircularSelection(itineraryList, i)
		initialPoint := itinerary.Car.ActualPosition()
		orderedClients := orderedClientsByItinerary[itinerary]
		orderedClientsByItinerary[itinerary] = insertInBestPosition(initialPoint, client, orderedClients)
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

func fillRoute(itinerary routes.Itinerary, orderedClients []gps.Point, deposits []gps.Point) {
	route := itinerary.Route
	car := itinerary.Car
	for _, client := range orderedClients {
		closestDeposit := closestPoint(client, deposits)
		if !car.Support(client, closestDeposit) {
			moveCarAndAppendRoute(car, route, closestDeposit)
		}
		moveCarAndAppendRoute(car, route, client)
	}
}

func calcAdditionalDistance(from, through, to gps.Point) float64 {
	return gps.DistanceBetweenPoints(from, through, to) - gps.DistanceBetweenPoints(from, to)
}
