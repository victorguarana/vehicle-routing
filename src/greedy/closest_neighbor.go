package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/slc"
)

func ClosestNeighbor(itineraryList []routes.Itinerary, m gps.Map) {
	remaningClients := slc.Copy(m.Clients)
	for i := 0; len(remaningClients) > 0; i++ {
		itinerary := slc.CircularSelection(itineraryList, i)
		car := itinerary.Car
		route := itinerary.Route
		carActualPosition := car.ActualPosition()
		closestClient := closestPoint(carActualPosition, remaningClients)
		closestDepositFromClosestClient := closestPoint(closestClient, m.Deposits)

		if car.Support(closestClient, closestDepositFromClosestClient) {
			moveCarAndAppendRoute(car, route, closestClient)
			remaningClients = slc.RemoveElement(remaningClients, closestClient)
			continue
		}

		closestDepositFromActualPosition := closestPoint(carActualPosition, m.Deposits)
		moveCarAndAppendRoute(car, route, closestDepositFromActualPosition)
	}

	finishItineraryOnClosestDeposits(itineraryList, m)
}
