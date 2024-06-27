package greedy

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
)

func finishOnClosestWarehouses(constructorList []itinerary.Constructor, m gps.Map) {
	for _, constructor := range constructorList {
		position := constructor.ActualCarPoint()
		closestWarehouse := gps.ClosestPoint(position, m.Warehouses)
		constructor.MoveCar(closestWarehouse)
	}
}
