package csp

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/route"
	"github.com/victorguarana/vehicle-routing/src/slc"
)

func CoveringWithDrones(constructorList []itinerary.Constructor, m gps.Map, neighborhoodDistance float64) {
	clientsNeighborhood := gps.MapNeighborhood(m.Clients, neighborhoodDistance)
	for i := 0; len(clientsNeighborhood) > 0; i++ {
		constructor := slc.CircularSelection(constructorList, i)
		actualClient := gps.ClosestPointWithMostNeighbors(constructor.ActualCarPoint(), clientsNeighborhood)
		constructor.MoveCar(actualClient)
		deliverNeighborsWithDrones(constructor, clientsNeighborhood[actualClient])
		removeClientAndItsNeighborsFromMap(actualClient, clientsNeighborhood)
	}

	finishOnClosestWarehouses(constructorList, m)
}

// TODO: Implement an function similar to this one
// but finish when drone can not support next client
func deliverNeighborsWithDrones(constructor itinerary.Constructor, neighbors []gps.Point) {
	if len(neighbors) <= 0 {
		return
	}
	droneNumbers := constructor.DroneNumbers()
	actualCarPoint := constructor.ActualCarPoint()
	actualCarStop := constructor.ActualCarStop()
	for neighborIndex, droneIndex := 0, 0; neighborIndex < len(neighbors); droneIndex++ {
		actualNeighbor := neighbors[neighborIndex]
		actualDroneNumber := slc.CircularSelection(droneNumbers, droneIndex)
		if shouldRetry := tryToDeliver(constructor, actualDroneNumber, actualCarStop, actualCarPoint, actualNeighbor); !shouldRetry {
			neighborIndex++
		}
	}
	constructor.LandAllDrones(actualCarStop)
}

func tryToDeliver(constructor itinerary.Constructor, droneNumber itinerary.DroneNumber, returningStop route.IMainStop, returningPoint gps.Point, destination gps.Point) bool {
	if constructor.DroneSupport(droneNumber, destination, returningPoint) {
		if !constructor.DroneIsFlying(droneNumber) {
			constructor.StartDroneFlight(droneNumber, returningStop)
		}
		constructor.MoveDrone(droneNumber, destination)
		return false
	}

	if constructor.DroneIsFlying(droneNumber) {
		constructor.LandDrone(droneNumber, returningStop)
		return true
	}
	return false
}

func removeClientAndItsNeighborsFromMap(client gps.Point, clientsNeighborhood gps.Neighborhood) {
	neighbors := clientsNeighborhood[client]
	for _, neighbor := range neighbors {
		gps.RemovePointFromNearbyMap(neighbor, clientsNeighborhood)
	}
	gps.RemovePointFromNearbyMap(client, clientsNeighborhood)
}

func finishOnClosestWarehouses(constructorList []itinerary.Constructor, m gps.Map) {
	for _, constructor := range constructorList {
		position := constructor.ActualCarPoint()
		closestWarehouse := gps.ClosestPoint(position, m.Warehouses)
		constructor.MoveCar(closestWarehouse)
	}
}
