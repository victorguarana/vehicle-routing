package csp

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/route"
	"github.com/victorguarana/vehicle-routing/internal/slc"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
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
	drones := constructor.Drones()
	actualCarPoint := constructor.ActualCarPoint()
	actualCarStop := constructor.ActualCarStop()
	for neighborIndex, droneIndex := 0, 0; neighborIndex < len(neighbors); droneIndex++ {
		actualNeighbor := neighbors[neighborIndex]
		actualDroneNumber := slc.CircularSelection(drones, droneIndex)
		if shouldRetry := tryToDeliver(constructor, actualDroneNumber, actualCarStop, actualCarPoint, actualNeighbor); !shouldRetry {
			neighborIndex++
		}
	}
	constructor.LandAllDrones(actualCarStop)
}

func tryToDeliver(constructor itinerary.Constructor, drone vehicle.IDrone, returningStop route.IMainStop, returningPoint gps.Point, destination gps.Point) bool {
	if constructor.DroneSupport(drone, destination, returningPoint) {
		if !constructor.DroneIsFlying(drone) {
			constructor.StartDroneFlight(drone, returningStop)
		}
		constructor.MoveDrone(drone, destination)
		return false
	}

	if constructor.DroneIsFlying(drone) {
		constructor.LandDrone(drone, returningStop)
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
