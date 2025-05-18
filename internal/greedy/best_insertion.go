package greedy

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/slc"
)

func BestInsertion(constructorList []itinerary.Constructor, m gps.Map) {
	orderedCustomersListsByRoute := orderCustomersByItinerary(constructorList, m.Customers)
	for index, orderedCustomers := range orderedCustomersListsByRoute {
		constructor := constructorList[index]
		fillRoute(constructor, orderedCustomers, m.Warehouses)
	}
	finishOnClosestWarehouses(constructorList, m)
}

func orderCustomersByItinerary(constructorList []itinerary.Constructor, customers []gps.Point) map[int][]gps.Point {
	orderedCustomersByItinerary := map[int][]gps.Point{}
	for i, customer := range customers {
		constructor, constructorIndex := slc.CircularSelectionWithIndex(constructorList, i)
		initialPoint := constructor.ActualCarPoint()
		orderedCustomers := orderedCustomersByItinerary[constructorIndex]
		orderedCustomersByItinerary[constructorIndex] = insertInBestPosition(initialPoint, customer, orderedCustomers)
	}
	return orderedCustomersByItinerary
}

func insertInBestPosition(initialPoint gps.Point, customer gps.Point, orderedCustomers []gps.Point) []gps.Point {
	if len(orderedCustomers) == 0 {
		return []gps.Point{customer}
	}
	bestIndex := findBestPosition(initialPoint, customer, orderedCustomers)
	return slc.InsertAt(orderedCustomers, customer, bestIndex)
}

func findBestPosition(initialPoint gps.Point, customer gps.Point, orderedCustomers []gps.Point) int {
	var bestIndex int
	shortestAdditionalDistance := gps.AdditionalDistancePassingThrough(initialPoint, customer, orderedCustomers[0])
	for i := 1; i < len(orderedCustomers); i++ {
		addictionalDistance := gps.AdditionalDistancePassingThrough(orderedCustomers[i-1], customer, orderedCustomers[i])
		if addictionalDistance < shortestAdditionalDistance {
			bestIndex = i
			shortestAdditionalDistance = addictionalDistance
		}
	}

	addictionalDistance := gps.AdditionalDistancePassingThrough(orderedCustomers[len(orderedCustomers)-1], customer, initialPoint)
	if addictionalDistance < shortestAdditionalDistance {
		bestIndex = len(orderedCustomers)
	}

	return bestIndex
}

func fillRoute(constructor itinerary.Constructor, orderedCustomers []gps.Point, warehouses []gps.Point) {
	for _, customer := range orderedCustomers {
		closestWarehouse := gps.ClosestPoint(customer, warehouses)
		if !constructor.Car().Support(customer, closestWarehouse) {
			constructor.MoveCar(closestWarehouse)
		}
		constructor.MoveCar(customer)
	}
}
