package greedy

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/slc"
)

func BestInsertion(itineraryList []itinerary.Itinerary, m gps.Map) {
	orderedClientsListsByRoute := orderClientsByItinerary(itineraryList, m.Clients)
	for index, orderedClients := range orderedClientsListsByRoute {
		itn := itineraryList[index]
		fillRoute(itn, orderedClients, m.Deposits)
	}
	finishItineraryOnClosestDeposits(itineraryList, m)
}

func orderClientsByItinerary(itineraryList []itinerary.Itinerary, clients []gps.Point) map[int][]gps.Point {
	orderedClientsByItinerary := map[int][]gps.Point{}
	for i, client := range clients {
		itn, itnIndex := slc.CircularSelectionWithIndex(itineraryList, i)
		initialPoint := itn.ActualCarPoint()
		orderedClients := orderedClientsByItinerary[itnIndex]
		orderedClientsByItinerary[itnIndex] = insertInBestPosition(initialPoint, client, orderedClients)
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
	shortestAdditionalDistance := gps.AdditionalDistancePassingThrough(initialPoint, client, orderedClients[0])
	for i := 1; i < len(orderedClients); i++ {
		addictionalDistance := gps.AdditionalDistancePassingThrough(orderedClients[i-1], client, orderedClients[i])
		if addictionalDistance < shortestAdditionalDistance {
			bestIndex = i
			shortestAdditionalDistance = addictionalDistance
		}
	}

	addictionalDistance := gps.AdditionalDistancePassingThrough(orderedClients[len(orderedClients)-1], client, initialPoint)
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
