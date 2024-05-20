package greedy

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
)

func finishItineraryOnClosestWarehouses(itineraryList []itinerary.Itinerary, m gps.Map) {
	for _, itinerary := range itineraryList {
		position := itinerary.ActualCarPoint()
		closestWarehouse := gps.ClosestPoint(position, m.Warehouses)
		itinerary.MoveCar(closestWarehouse)
	}
}
