package csp

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/route"
	"github.com/victorguarana/vehicle-routing/internal/slc"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

func CoveringWithDrones(constructorList []itinerary.Constructor, m gps.Map, neighborhoodDistance float64) {
	customersNeighborhood := gps.MapNeighborhood(m.Customers, neighborhoodDistance)
	for i := 0; len(customersNeighborhood) > 0; i++ {
		constructor := slc.CircularSelection(constructorList, i)
		actualCustomer := gps.ClosestPointWithMostNeighbors(constructor.ActualCarPoint(), customersNeighborhood)
		constructor.MoveCar(actualCustomer)
		deliverNeighborsWithDrones(constructor, customersNeighborhood[actualCustomer])
		removeCustomerAndItsNeighborsFromMap(actualCustomer, customersNeighborhood)
	}

	finishOnClosestWarehouses(constructorList, m)
}

// TODO: Implement an function similar to this one
// but finish when drone can not support next customer
func deliverNeighborsWithDrones(constructor itinerary.Constructor, neighbors []gps.Point) {
	if len(neighbors) <= 0 {
		return
	}
	drones := constructor.Car().Drones()
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
	if drone.Support(destination, returningPoint) {
		if !drone.IsFlying() {
			constructor.StartDroneFlight(drone, returningStop)
		}
		constructor.MoveDrone(drone, destination)
		return false
	}

	if drone.IsFlying() {
		constructor.LandDrone(drone, returningStop)
		return true
	}
	return false
}

func removeCustomerAndItsNeighborsFromMap(customer gps.Point, customersNeighborhood gps.Neighborhood) {
	neighbors := customersNeighborhood[customer]
	for _, neighbor := range neighbors {
		gps.RemovePointFromNearbyMap(neighbor, customersNeighborhood)
	}
	gps.RemovePointFromNearbyMap(customer, customersNeighborhood)
}

func finishOnClosestWarehouses(constructorList []itinerary.Constructor, m gps.Map) {
	for _, constructor := range constructorList {
		position := constructor.ActualCarPoint()
		closestWarehouse := gps.ClosestPoint(position, m.Warehouses)
		constructor.MoveCar(closestWarehouse)
	}
}
