package csp

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/slc"
)

func CoveringWithDrones(itineraryList []itinerary.Itinerary, m gps.Map, neighborhoodDistance float64) {
	clientsNeighborhood := gps.MapNeighborhood(m.Clients, neighborhoodDistance)
	for itineraryIndex := 0; len(clientsNeighborhood) > 0; itineraryIndex++ {
		itn := slc.CircularSelection(itineraryList, itineraryIndex)
		actualClient := gps.ClosestPointWithMostNeighbors(itn.ActualCarPoint(), clientsNeighborhood)
		itn.MoveCar(actualClient)
		deliverNeighborsWithDrones(itn, clientsNeighborhood[actualClient])
		removeClientAndItsNeighborsFromMap(actualClient, clientsNeighborhood)
		gps.RemovePointFromNearbyMap(actualClient, clientsNeighborhood)
	}

	finishItineraryOnClosestWarehouses(itineraryList, m)
}

// TODO: Implement an function similar to this one
// but finish when drone can not support next client
func deliverNeighborsWithDrones(itn itinerary.Itinerary, neighbors []gps.Point) {
	if len(neighbors) <= 0 {
		return
	}
	droneNumbers := itn.DroneNumbers()
	actualCarPoint := itn.ActualCarPoint()
	actualCarStop := itn.ActualCarStop()
	for neighborIndex, droneIndex := 0, 0; neighborIndex < len(neighbors); droneIndex++ {
		actualNeighbor := neighbors[neighborIndex]
		actualDroneNumber := slc.CircularSelection(droneNumbers, droneIndex)
		if itn.DroneSupport(actualDroneNumber, actualNeighbor, actualCarPoint) {
			if !itn.DroneIsFlying(actualDroneNumber) {
				itn.StartDroneFlight(actualDroneNumber, actualCarStop)
			}
			itn.MoveDrone(actualDroneNumber, actualNeighbor)
			neighborIndex++
		} else {
			itn.LandDrone(actualDroneNumber, actualCarStop)
		}
	}
	itn.LandAllDrones(actualCarStop)
}

func removeClientAndItsNeighborsFromMap(client gps.Point, clientsNeighborhood gps.Neighborhood) {
	neighbors := clientsNeighborhood[client]
	for _, neighbor := range neighbors {
		gps.RemovePointFromNearbyMap(neighbor, clientsNeighborhood)
	}
	gps.RemovePointFromNearbyMap(client, clientsNeighborhood)
}

func finishItineraryOnClosestWarehouses(itineraryList []itinerary.Itinerary, m gps.Map) {
	for _, itinerary := range itineraryList {
		position := itinerary.ActualCarPoint()
		closestWarehouse := gps.ClosestPoint(position, m.Warehouses)
		itinerary.MoveCar(closestWarehouse)
	}
}
