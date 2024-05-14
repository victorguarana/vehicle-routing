package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/slc"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

func ClosestNeighbor(cars []vehicles.ICar, m gps.Map) {
	remaningClients := slc.Copy(m.Clients)
	for i := 0; len(remaningClients) > 0; i++ {
		car := slc.CircularSelection(cars, i)
		carActualPosition := car.ActualPoint()
		closestClient := closestPoint(carActualPosition, remaningClients)
		closestDepositFromClosestClient := closestPoint(closestClient, m.Deposits)

		if car.Support(closestClient, closestDepositFromClosestClient) {
			car.Move(routes.NewMainStop(closestClient))
			remaningClients = slc.RemoveElement(remaningClients, closestClient)
			continue
		}

		closestDepositFromActualPosition := closestPoint(carActualPosition, m.Deposits)
		car.Move(routes.NewMainStop(closestDepositFromActualPosition))
	}

	finishItineraryOnClosestDeposits(cars, m)
}
