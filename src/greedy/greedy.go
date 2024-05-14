package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

func finishItineraryOnClosestDeposits(carsList []vehicles.ICar, m gps.Map) {
	for _, car := range carsList {
		position := car.ActualPoint()
		closestDeposit := gps.ClosestPoint(position, m.Deposits)
		car.Move(routes.NewMainStop(closestDeposit))
	}
}
