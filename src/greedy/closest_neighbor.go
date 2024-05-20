package greedy

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/slc"
)

func ClosestNeighbor(itineraryList []itinerary.Itinerary, m gps.Map) {
	remaningClients := slc.Copy(m.Clients)
	for i := 0; len(remaningClients) > 0; i++ {
		itinerary := slc.CircularSelection(itineraryList, i)
		itineraryActualPosition := itinerary.ActualCarPoint()
		closestClient := gps.ClosestPoint(itineraryActualPosition, remaningClients)
		closestWarehouseFromClosestClient := gps.ClosestPoint(closestClient, m.Warehouses)

		if itinerary.CarSupport(closestClient, closestWarehouseFromClosestClient) {
			itinerary.MoveCar(closestClient)
			remaningClients = slc.RemoveElement(remaningClients, closestClient)
			continue
		}

		closestWarehouseFromActualPosition := gps.ClosestPoint(itineraryActualPosition, m.Warehouses)
		itinerary.MoveCar(closestWarehouseFromActualPosition)
	}

	finishItineraryOnClosestWarehouses(itineraryList, m)
}
