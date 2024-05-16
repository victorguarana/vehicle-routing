package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/itinerary"
	"github.com/victorguarana/go-vehicle-route/src/slc"
)

func BestInsertion(itineraryList []itinerary.Itinerary, m gps.Map) {
	orderedClientsListsByRoute := orderClientsByItinerary(itineraryList, m.Clients)
	for itinerary, orderedClients := range orderedClientsListsByRoute {
		fillRoute(itinerary, orderedClients, m.Deposits)
	}
	finishItineraryOnClosestDeposits(itineraryList, m)
}

func orderClientsByItinerary(itineraryList []itinerary.Itinerary, clients []gps.Point) map[itinerary.Itinerary][]gps.Point {
	orderedClientsByItinerary := map[itinerary.Itinerary][]gps.Point{}
	for i, client := range clients {
		itinerary := slc.CircularSelection(itineraryList, i)
		initialPoint := itinerary.ActualCarPoint()
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

func fillRoute(itinerary itinerary.Itinerary, orderedClients []gps.Point, deposits []gps.Point) {
	for _, client := range orderedClients {
		closestDeposit := gps.ClosestPoint(client, deposits)
		if !itinerary.CarSupport(client, closestDeposit) {
			itinerary.MoveCar(closestDeposit)
		}
		itinerary.MoveCar(client)
	}
}

// TODO: Move this method to gps package
func calcAdditionalDistance(from, through, to gps.Point) float64 {
	return gps.DistanceBetweenPoints(from, through, to) - gps.DistanceBetweenPoints(from, to)
}
