package greedy

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/slc"
)

func ClosestNeighbor(constructorList []itinerary.Constructor, m gps.Map) {
	remaningClients := slc.Copy(m.Clients)
	for i := 0; len(remaningClients) > 0; i++ {
		constructor := slc.CircularSelection(constructorList, i)
		carActualPoint := constructor.ActualCarPoint()
		closestClient := gps.ClosestPoint(carActualPoint, remaningClients)
		closestWarehouseFromClosestClient := gps.ClosestPoint(closestClient, m.Warehouses)

		if constructor.CarSupport(closestClient, closestWarehouseFromClosestClient) {
			constructor.MoveCar(closestClient)
			remaningClients = slc.RemoveElement(remaningClients, closestClient)
			continue
		}

		closestWarehouseFromActualPosition := gps.ClosestPoint(carActualPoint, m.Warehouses)
		constructor.MoveCar(closestWarehouseFromActualPosition)
	}

	finishOnClosestWarehouses(constructorList, m)
}
