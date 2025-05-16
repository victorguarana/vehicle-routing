package greedy

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/slc"
)

func BestInsertion(constructorList []itinerary.Constructor, m gps.Map) {
	orderedClientsListsByRoute := orderClientsByItinerary(constructorList, m.Clients)
	for index, orderedClients := range orderedClientsListsByRoute {
		constructor := constructorList[index]
		fillRoute(constructor, orderedClients, m.Warehouses)
	}
	finishOnClosestWarehouses(constructorList, m)
}

func orderClientsByItinerary(constructorList []itinerary.Constructor, clients []gps.Point) map[int][]gps.Point {
	orderedClientsByItinerary := map[int][]gps.Point{}
	for i, client := range clients {
		constructor, constructorIndex := slc.CircularSelectionWithIndex(constructorList, i)
		initialPoint := constructor.ActualCarPoint()
		orderedClients := orderedClientsByItinerary[constructorIndex]
		orderedClientsByItinerary[constructorIndex] = insertInBestPosition(initialPoint, client, orderedClients)
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

func fillRoute(constructor itinerary.Constructor, orderedClients []gps.Point, warehouses []gps.Point) {
	for _, client := range orderedClients {
		closestWarehouse := gps.ClosestPoint(client, warehouses)
		if !constructor.Car().Support(client, closestWarehouse) {
			constructor.MoveCar(closestWarehouse)
		}
		constructor.MoveCar(client)
	}
}
