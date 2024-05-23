package csp

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
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
		gps.RemovePointFromNearbyMap(actualClient, clientsNeighborhood)
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
		if constructor.DroneSupport(actualDroneNumber, actualNeighbor, actualCarPoint) {
			if !constructor.DroneIsFlying(actualDroneNumber) {
				constructor.StartDroneFlight(actualDroneNumber, actualCarStop)
			}
			constructor.MoveDrone(actualDroneNumber, actualNeighbor)
			neighborIndex++
		} else {
			constructor.LandDrone(actualDroneNumber, actualCarStop)
		}
	}
	constructor.LandAllDrones(actualCarStop)
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
