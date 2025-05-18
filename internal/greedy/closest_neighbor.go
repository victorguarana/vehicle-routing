package greedy

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/slc"
)

func ClosestNeighbor(constructorList []itinerary.Constructor, m gps.Map) {
	remaningCustomers := slc.Copy(m.Customers)
	for i := 0; len(remaningCustomers) > 0; i++ {
		constructor := slc.CircularSelection(constructorList, i)
		carActualPoint := constructor.ActualCarPoint()
		closestCustomer := gps.ClosestPoint(carActualPoint, remaningCustomers)
		closestWarehouseFromClosestCustomer := gps.ClosestPoint(closestCustomer, m.Warehouses)

		if constructor.Car().Support(closestCustomer, closestWarehouseFromClosestCustomer) {
			constructor.MoveCar(closestCustomer)
			remaningCustomers = slc.RemoveElement(remaningCustomers, closestCustomer)
			continue
		}

		closestWarehouseFromActualPosition := gps.ClosestPoint(carActualPoint, m.Warehouses)
		constructor.MoveCar(closestWarehouseFromActualPosition)
	}

	finishOnClosestWarehouses(constructorList, m)
}
